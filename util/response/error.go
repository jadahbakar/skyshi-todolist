package response

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// errors response
var (
	ErrNotFound       = errors.New("Not Found")             //404
	ErrBadRequest     = errors.New("Bad Request")           // 400
	ErrInternalServer = errors.New("Internal Server Error") // 500
)

// Error is
type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// NewError is
func NewError(f *fiber.Ctx, code int, m string) error {
	var status string
	switch code {
	case 400:
		status = ErrBadRequest.Error()
	case 401:
		status = ErrNotFound.Error()
	case 500:
		status = ErrInternalServer.Error()
	default:
		status = ErrInternalServer.Error()
	}

	data := Error{
		Status:  status,
		Message: m,
	}
	return f.Status(code).JSON(data)
}

// BadRequest is | 400
func BadRequest(f *fiber.Ctx, s string) error {
	return NewError(f, fiber.StatusBadRequest, s)
}

func HandleErrors(f *fiber.Ctx, s string) error {
	return NewError(f, fiber.StatusInternalServerError, s)
}

func NotFound(f *fiber.Ctx, s string) error {
	return NewError(f, fiber.StatusNotFound, s)
}
