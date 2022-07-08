package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wonpanu/learn-golang/service/pkg/entity"
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

func (b BlogHandler) CreateBlog(c *fiber.Ctx) error {
	var blog entity.Blog
	if err := c.BodyParser(&blog); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Status: "error",
			Data:   err.Error(),
		})
	}
	res, err := b.blogUsecase.CreateBlog(blog)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Status: "error",
			Data:   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(Response{
		Status: "ok",
		Data:   res,
	})
}

func (b BlogHandler) UpdateBlog(c *fiber.Ctx) error {
	var blog entity.Blog
	id := c.Params("id")
	if err := c.BodyParser(&blog); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Status: "error",
			Data:   err.Error(),
		})
	}
	res, err := b.blogUsecase.UpdateBlog(id, blog)
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

func (b BlogHandler) DeleteBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := b.blogUsecase.DeleteBlog(id)
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

func NewBlogHandler(blogUsecase usecase.BlogUsecase) BlogHandler {
	return BlogHandler{
		blogUsecase: blogUsecase,
	}
}
