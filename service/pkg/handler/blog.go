package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wonpanu/learn-golang/service/pkg/usecase"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type BlogHandler struct {
	blogUsecase usecase.BlogUsecase
}

func (b BlogHandler) GetAll(c *fiber.Ctx) error {
	blogs, err := b.blogUsecase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Status: "error",
			Data:   err.Error(),
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Status: "ok",
		Data:   blogs,
	})
}

func NewBlogHandler(blogUsecase usecase.BlogUsecase) BlogHandler {
	return BlogHandler{
		blogUsecase: blogUsecase,
	}
}
