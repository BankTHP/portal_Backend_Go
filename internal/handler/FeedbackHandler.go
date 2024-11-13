package handler

import (
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/service"

	"github.com/gofiber/fiber/v2"
)

type FeedbackHandler struct {
	feedbackService *service.FeedbackService
}

func NewFeedbackHandler(feedbackService *service.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{feedbackService: feedbackService}
}

func (h *FeedbackHandler) CreateFeedback(c *fiber.Ctx) error {
	var createRequest model.CreateFeedbackRequest
	if err := c.BodyParser(&createRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_INPUT", "ข้อมูลไม่ถูกต้อง"))
	}

	if err := h.feedbackService.CreateFeedback(&createRequest); err != nil {
		if err.Error() == "เบอร์โทรศัพท์ต้องไม่เกิน 10 หลัก" {
			return c.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_PHONE", err.Error()))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("CREATE_FEEDBACK_ERROR", err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(model.NewSuccessResponse(map[string]string{
		"message": "สร้าง Feedback สำเร็จ",
	}))
}

func (h *FeedbackHandler) GetFeedbackById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_ID", "รหัสไม่ถูกต้อง"))
	}

	feedback, err := h.feedbackService.GetFeedbackById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.NewErrorResponse("FEEDBACK_NOT_FOUND", "ไม่พบ Feedback"))
	}

	return c.JSON(model.NewSuccessResponse(feedback))
}

func (h *FeedbackHandler) DeleteFeedback(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_ID", "รหัสไม่ถูกต้อง"))
	}

	if err := h.feedbackService.DeleteFeedback(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("DELETE_FEEDBACK_ERROR", err.Error()))
	}

	return c.JSON(model.NewSuccessResponse(map[string]string{
		"message": "ลบ Feedback สำเร็จ",
	}))
}

func (h *FeedbackHandler) GetPaginatedFeedbacks(c *fiber.Ctx) error {
	var req model.FeedbackPaginatedRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_INPUT", "ข้อมูลไม่ถูกต้อง"))
	}

	if req.Page == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_PAGE", "หน้าต้องมากกว่า 0"))
	}

	if req.Size == 0 || req.Size > 100 {
		return c.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_SIZE", "ขนาดต้องอยู่ระหว่าง 1 ถึง 100"))
	}

	paginatedResponse, err := h.feedbackService.GetPaginatedFeedbacks(int(req.Page), int(req.Size))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("FETCH_FEEDBACKS_ERROR", err.Error()))
	}

	return c.JSON(model.NewSuccessResponse(paginatedResponse))
}

func (h *FeedbackHandler) GetAllFeedbacks(c *fiber.Ctx) error {
	feedbacks, err := h.feedbackService.GetAllFeedbacks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse(
			"FETCH_FEEDBACKS_ERROR",
			"ไม่สามารถดึงข้อมูล Feedback ได้: " + err.Error(),
		))
	}

	return c.JSON(model.NewSuccessResponse(feedbacks))
} 