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

// CreateComment สร้าง comment ใหม่
// @Summary สร้าง comment ใหม่
// @Description สร้าง comment ใหม่จากข้อมูลที่ส่งมา
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body model.CreateCommentRequest true "ข้อมูล Comment ที่จะสร้าง"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /comments [post]
func (c *CommentHandlers) CreateComment(ctx *fiber.Ctx) error {
	var createCommentRequest model.CreateCommentRequest
	if err := ctx.BodyParser(&createCommentRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid input"})
	}

	if err := c.commentService.CreateComment(&createCommentRequest); err != nil {
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

// DeleteComment ลบ comment ตาม ID
// @Summary ลบ comment ตาม ID
// @Description ลบ comment โดยใช้ ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /comments/{id} [delete]
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

// GetPaginatedComments ดึงข้อมูล comment แบบแบ่งหน้า
// @Summary ดึงข้อมูล comment แบบแบ่งหน้า
// @Description ดึงข้อมูล comment แบบแบ่งหน้าตาม Post ID
// @Tags comments
// @Accept json
// @Produce json
// @Param request body model.CommentPaginatedRequest true "ข้อมูลการแบ่งหน้า"
// @Success 200 {object} model.PaginatedResponse
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /comments/paginated [post]
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
	if req.Limit == 0 || req.Limit > 100 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Limit must be between 1 and 100"})
	}

	paginatedResponse, err := c.commentService.GetPaginatedComments(int(req.Page), int(req.Limit), int(req.PostID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch paginated comments"})
	}

	return ctx.JSON(paginatedResponse)
}

// UpdateComment อัปเดต comment ตาม ID
// @Summary อัปเดต comment ตาม ID
// @Description อัปเดตข้อมูล comment โดยใช้ ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Param comment body model.UpdateCommentRequest true "ข้อมูล Comment ที่จะอัปเดต"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /comments/{id} [put]
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
