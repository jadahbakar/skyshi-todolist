package response

import "github.com/gofiber/fiber/v2"

// Success is
type Success struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewSuccess is
func NewSuccess(c *fiber.Ctx, code int, d interface{}) error {
	data := Success{
		Status:  "Success",
		Message: "Success",
		Data:    d,
	}
	return c.Status(code).JSON(data)
}
