package response

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// errors response
var (
	ErrBadRequest     = errors.New("Bad Request, something wrong on your request") // 400
	ErrInternalServer = errors.New("Internal Server Error")                        // 500
)

// Error is
type Error struct {
	Message string `json:"remark"`
}

// NewError is
func NewError(f *fiber.Ctx, code int, m string) error {
	data := Error{
		Message: m,
	}
	return f.Status(code).JSON(data)
}

// BadRequest is | 400
func BadRequest(f *fiber.Ctx, s interface{}) error {
	return NewError(f, fiber.StatusBadRequest, ErrBadRequest.Error())
}

func HandleErrors(f *fiber.Ctx, s string) error {
	return NewError(f, fiber.StatusBadRequest, s)
}
