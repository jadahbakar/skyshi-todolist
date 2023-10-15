package response

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/skyshi-todolist/util/errorlib"
)

// errors response
var (
	ErrNotFound       = errors.New("Not Found")             //404
	ErrBadRequest     = errors.New("Bad Request")           // 400
	ErrInternalServer = errors.New("internal server error") // 500
)

// Error is
type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func RenderError(f *fiber.Ctx, err error, msg string) error {
	code := fiber.StatusInternalServerError
	status := ErrInternalServer.Error()
	var ierr *errorlib.Error

	fmt.Printf("err : %v\n", err)
	// fmt.Printf("ierr.Code() : %d\n", ierr.Code())

	data := Error{}
	if !errors.As(err, &ierr) {
		code = fiber.StatusInternalServerError
		data = Error{
			Status:  status,
			Message: ierr.Error(),
		}
	} else {
		switch ierr.Code() {
		case errorlib.ErrorCodeNotFound:
			code = fiber.StatusNotFound
			status = ErrNotFound.Error()
		case errorlib.ErrorCodeInvalidArgument:
			code = fiber.StatusBadRequest
			status = ErrBadRequest.Error()
		}
	}
	data = Error{
		Status:  status,
		Message: ierr.Error(),
	}

	return f.Status(code).JSON(data)

}
