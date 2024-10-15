package handler

import (
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PostHandlers struct {
	postService *service.PostService
}

func NewPostHandlers(postService *service.PostService) *PostHandlers {
	return &PostHandlers{postService: postService}
}

func (c *PostHandlers) CreatePost(ctx *fiber.Ctx) error {
	var createPostRequest model.CreatePostRequest
	if err := ctx.BodyParser(&createPostRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.postService.CreatePost(&createPostRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Post created successfully"})
}

func (c *PostHandlers) GetPostByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	println(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	post, err := c.postService.GetPostByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
	}

	return ctx.JSON(post)
}

func (c *PostHandlers) UpdatePost(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var updateRequest model.UpdatePostRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.postService.UpdatePost(uint(id), &updateRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Post updated successfully"})
}

func (c *PostHandlers) DeletePost(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := c.postService.DeletePost(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Post deleted successfully"})
}

func (c *PostHandlers) GetAllPosts(ctx *fiber.Ctx) error {
	posts, err := c.postService.GetAllPosts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(posts)
}

func (c *PostHandlers) GetPaginatedPosts(ctx *fiber.Ctx) error {
    page, err := strconv.Atoi(ctx.Query("page", "1"))
    if err != nil || page < 1 {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page number"})
    }

    limit, err := strconv.Atoi(ctx.Query("limit", "10"))
    if err != nil || limit < 1 {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit number"})
    }

    paginatedResponse, err := c.postService.GetPaginatedPosts(page, limit)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return ctx.JSON(paginatedResponse)
}


