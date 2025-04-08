package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "url-shorter/server/internals/handlers"
    "url-shorter/server/internals/store/dbstore"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/helmet"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    app := fiber.New()

    app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:5173", 
        AllowMethods: "GET,POST,OPTIONS",     
    }))
    app.Use(helmet.New())
    app.Use(logger.New())

    urlStore := dbstore.NewUrlStore()

    app.Get("/health", func(c *fiber.Ctx) error {
        return c.SendStatus(200)
    })

    app.Post("/", handlers.NewUrlShortenHandler(handlers.UrlShortenHandlerParam{
        UrlStore: urlStore,
    }).ShortenUrl)

    app.Get("/:shortUrl", handlers.GetOrignalUrlHandler(handlers.GetUrlHandlerParams{
        UrlStore: urlStore,
    }).GetUrl)
    killSig := make(chan os.Signal, 1)
    signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

    go func() {
        <-killSig
        fmt.Println("Gracefully shutting down...")
        _ = app.Shutdown()
    }()
    if err := app.Listen(":8080"); err != nil {
        log.Panic(err)
    }

    fmt.Println("Running cleanup tasks...")

}