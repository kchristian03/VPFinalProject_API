package MusicController

import (
	"ZenZen_API/app"
	"ZenZen_API/framework/utils"
	"github.com/gofiber/fiber/v2"
)

// MusicHandler
// This is an example of a handler
// Modify it to suit your needs
func MusicHandler(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(utils.HttpResponse{
		Success: true,
		Message: "Hello World from " + app.Config.Application.Name,
		Data:    nil,
	})
}
