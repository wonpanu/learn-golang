package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wonpanu/learn-golang/service/pkg/usecase"
)

type BahRamHandler struct {
	bahRamUsecase usecase.BahRamUsecase
}

func (b BahRamHandler) BahRam(c *fiber.Ctx) error {
	n := c.Params("n")
	res, err := b.bahRamUsecase.BahRam(n)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Status: "error",
			Data:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status: "ok",
		Data:   res,
	})
}

func NewBahRamHandler(bahRamUsecase usecase.BahRamUsecase) BahRamHandler {
	return BahRamHandler{
		bahRamUsecase: bahRamUsecase,
	}
}
