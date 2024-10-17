package handler

import (
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CommentHandlers struct {
	commentService *service.CommentService
}

func NewCommentHandlers(commentService *service.CommentService) *CommentHandlers {
	return &CommentHandlers{commentService: commentService}
}

func (c *CommentHandlers) CreateComment(ctx *fiber.Ctx) error {
	var createCommentRequest model.CreateCommentRequest
	if err := ctx.BodyParser(&createCommentRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid input"})
	}

	if err := c.commentService.CreateComment(&createCommentRequest); err != nil {
		if err.Error() == "post not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(map[string]interface{}{"error": err.Error()})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(map[string]interface{}{"message": "Comment created successfully"})
}

func (c *CommentHandlers) GetCommentByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	println(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	comment, err := c.commentService.GetCommentByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
	}

	return ctx.JSON(comment)
}

func (c *CommentHandlers) GetCommentByPostID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	println(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	comment, err := c.commentService.GetCommentByPostID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
	}

	return ctx.JSON(comment)
}

func (c *CommentHandlers) DeleteComment(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := c.commentService.DeleteComment(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Comment deleted successfully"})
}

func (c *CommentHandlers) GetPaginatedComments(ctx *fiber.Ctx) error {
	var req model.CommentPaginatedRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if req.PostID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "PostID must be greater than 0"})
	}
	if req.Page == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Page must be greater than 0"})
	}
	if req.Size == 0 || req.Size > 100 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Size must be between 1 and 100"})
	}

	paginatedResponse, err := c.commentService.GetPaginatedComments(int(req.Page), int(req.Size), int(req.PostID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch paginated comments"})
	}

	return ctx.JSON(paginatedResponse)
}

func (c *CommentHandlers) UpdateComment(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var updateRequest model.UpdateCommentRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.commentService.UpdateComment(uint(id), &updateRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Comment updated successfully"})
}

func (c *CommentHandlers) GetPaginatedCommentsByUserId(ctx *fiber.Ctx) error {
	var req model.CommentByUserIdPaginatedRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if req.CommentCreateBy == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "PostID must be greater than 0"})
	}
	if req.Page == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Page must be greater than 0"})
	}
	if req.Size == 0 || req.Size > 100 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Size must be between 1 and 100"})
	}

	paginatedResponse, err := c.commentService.GetPaginatedCommentsByUserId(int(req.Page), int(req.Size), req.CommentCreateBy)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch paginated comments"})
	}

	return ctx.JSON(paginatedResponse)
}
