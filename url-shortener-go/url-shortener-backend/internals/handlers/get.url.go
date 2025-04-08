package handlers

import (
    "log"
    "url-shorter/server/internals/store"
    "strings"

    "github.com/gofiber/fiber/v2"
)

type GetUrlHandler struct {
    urlStore store.UrlStore
}

type GetUrlHandlerParams struct {
    UrlStore store.UrlStore
}
func GetOrignalUrlHandler(params GetUrlHandlerParams) *GetUrlHandler {
    return &GetUrlHandler{
        urlStore: params.UrlStore,
    }
}

func (h *GetUrlHandler) GetUrl(c *fiber.Ctx) error {
    shortUrl := c.Params("shortUrl")
    if len(shortUrl) == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid short URL",
        })
    }

    url, err := h.urlStore.GetUrl(shortUrl)
    if err != nil {
        log.Printf("Error fetching URL for shortUrl %s: %v", shortUrl, err)
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "URL not found",
        })
    }

    if url.LongUrl == "" {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Invalid long URL",
        })
    }

    if !startsWithProtocol(url.LongUrl) {
        url.LongUrl = "http://" + url.LongUrl
    }

    log.Printf("Redirecting short URL %s to %s", shortUrl, url.LongUrl)
    return c.Redirect(url.LongUrl, fiber.StatusFound)
}

func startsWithProtocol(url string) bool {
    return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}