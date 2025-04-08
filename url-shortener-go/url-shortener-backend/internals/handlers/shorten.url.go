package handlers

import (
    "url-shorter/server/internals/store"
    "url-shorter/server/internals/util"

    "github.com/gofiber/fiber/v2"
)

type UrlShortenHandler struct {
    urlStore store.UrlStore
}

type UrlShortenHandlerParam struct {
    UrlStore store.UrlStore
}

func NewUrlShortenHandler(params UrlShortenHandlerParam) *UrlShortenHandler {
    return &UrlShortenHandler{
        urlStore: params.UrlStore,
    }
}

func (h *UrlShortenHandler) ShortenUrl(c *fiber.Ctx) error {
    var requestBody struct {
        LongUrl string `json:"longUrl"`
    }

    if err := c.BodyParser(&requestBody); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    longUrl := requestBody.LongUrl
    if len(longUrl) == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid or missing long URL",
        })
    }

    shortUrl := util.GenerateShortUrl(6)
    err := h.urlStore.StoreUrl(longUrl, shortUrl)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to store the URL",
        })
    }
    fullShortUrl := "http://localhost:8080/" + shortUrl

    return c.Status(201).JSON(fiber.Map{
        "shortUrl": fullShortUrl,
    })
}