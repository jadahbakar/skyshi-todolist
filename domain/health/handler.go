package health

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/skyshi-todolist/util/logger"
	"github.com/jadahbakar/skyshi-todolist/util/response"
)

func AddRoutes(router fiber.Router) {
	router.Get("/health", GetHealth)
}

func GetHealth(c *fiber.Ctx) error {
	// Return status 200 OK.
	logger.Info("healthy")
	return response.NewSuccess(c, fiber.StatusOK, "healthty")
}
