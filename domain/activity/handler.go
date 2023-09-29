package activity

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/skyshi-todolist/util/logger"
	"github.com/jadahbakar/skyshi-todolist/util/response"
	"github.com/jadahbakar/skyshi-todolist/util/validator"
)

type Handler struct {
	service Service
}

func NewHandler(r fiber.Router, s Service) {
	h := &Handler{service: s}
	r.Post("/", h.Create)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	req := &PostReq{}

	if err := c.BodyParser(req); err != nil {
		logger.Errorf("Error On Body Parser: ", err)
		return response.BadRequest(c, err.Error())
	}

	//---validation
	validate := validator.NewValidator()
	if err := validate.Struct(req); err != nil {
		logger.Errorf("Error On Validate: ", err)
		return response.BadRequest(c, err)
	}

	//---service
	data, err := h.service.Create(req)
	if err != nil {
		logger.Errorf("Error On Service: ", err)
		return response.HandleErrors(c, err.Error())
	}

	//---response
	return c.Status(fiber.StatusOK).JSON(data)
}
