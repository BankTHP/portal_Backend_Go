package handler

import (
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/service"

	"github.com/gofiber/fiber/v2"
)

type NewsHandlers struct {
	newsService *service.NewsService
}

func NewsPHandlers(newsService *service.NewsService) *NewsHandlers {
	return &NewsHandlers{newsService: newsService}
}


func (c *NewsHandlers) CreateNews(ctx *fiber.Ctx) error {
	var createNewsRequest model.CreateNewsRequest
	if err := ctx.BodyParser(&createNewsRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_INPUT", "ข้อมูลไม่ถูกต้อง"))
	}

	if err := c.newsService.CreateNews(&createNewsRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("CREATE_NEWS_ERROR", err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.NewSuccessResponse(map[string]string{
		"message": "สร้างข่าวสำเร็จ",
	}))
}


func (c *NewsHandlers) GetNewsByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_ID", "รหัสไม่ถูกต้อง"))
	}

	news, err := c.newsService.GetNewsByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(model.NewErrorResponse("NEWS_NOT_FOUND", "ไม่พบข่าว"))
	}

	return ctx.JSON(model.NewSuccessResponse(news))
}


func (c *NewsHandlers) UpdateNews(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_ID", "รหัสไม่ถูกต้อง"))
	}

	var updateRequest model.UpdateNewsRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_INPUT", "ข้อมูลไม่ถูกต้อง"))
	}

	if err := c.newsService.UpdateNews(uint(id), &updateRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("UPDATE_NEWS_ERROR", err.Error()))
	}

	return ctx.JSON(model.NewSuccessResponse(map[string]string{
		"message": "อัปเดตข่าวสำเร็จ",
	}))
}


func (c *NewsHandlers) DeleteNews(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_ID", "รหัสไม่ถูกต้อง"))
	}

	if err := c.newsService.DeleteNews(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("DELETE_NEWS_ERROR", err.Error()))
	}

	return ctx.JSON(model.NewSuccessResponse(map[string]string{
		"message": "ลบข่าวสำเร็จ",
	}))
}


func (c *NewsHandlers) GetAllNews(ctx *fiber.Ctx) error {
	news, err := c.newsService.GetAllNews()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("FETCH_NEWS_ERROR", err.Error()))
	}

	return ctx.JSON(model.NewSuccessResponse(news))
}
