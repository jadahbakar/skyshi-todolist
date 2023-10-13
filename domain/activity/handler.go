package activity

import (
	"fmt"

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
	r.Patch("/:id", h.Update)
	r.Delete("/:id", h.Delete)
	r.Get("/:id", h.GetById)
	r.Get("/", h.GetAll)
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
	return response.NewSuccess(c, fiber.StatusCreated, data)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	req := &PatchReq{}

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

	param := c.Params("id")

	//---service
	data, err := h.service.Update(param, req.Title)
	if err != nil {
		logger.Errorf("Error On Service: ", err)
		return response.HandleErrors(c, err.Error())
	}

	//---response
	return response.NewSuccess(c, fiber.StatusOK, data)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	param := c.Params("id")

	//---service
	data, err := h.service.Delete(param)
	if err != nil {
		logger.Errorf("Error On Service: ", err)
		return response.HandleErrors(c, err.Error())
	}

	resp := fmt.Sprintf("Activity with ID %d Not Found", data)
	//---response
	return response.SuccessDelete(c, fiber.StatusOK, resp)
}

func (h *Handler) GetById(c *fiber.Ctx) error {
	param := c.Params("id")

	//---service
	data, err := h.service.FindActId(param)
	if err != nil {
		logger.Errorf("Error On Service: ", err)
		return response.HandleErrors(c, err.Error())
	}

	//---response
	return response.NewSuccess(c, fiber.StatusOK, data)
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	//---service
	data, err := h.service.FindAll()
	if err != nil {
		logger.Errorf("Error On Service: ", err)
		return response.HandleErrors(c, err.Error())
	}

	//---response
	return response.NewSuccess(c, fiber.StatusOK, data)
}
