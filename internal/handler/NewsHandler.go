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

// CreateNews สร้างข่าวใหม่
// @Summary สร้างข่าวใหม่
// @Description สร้างข่าวใหม่ในระบบ
// @Tags news
// @Accept json
// @Produce json
// @Param news body model.CreateNewsRequest true "ข้อมูลข่าวที่จะสร้าง"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /news [post]
func (c *NewsHandlers) CreateNews(ctx *fiber.Ctx) error {
	var createNewsRequest model.CreateNewsRequest
	if err := ctx.BodyParser(&createNewsRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.newsService.CreateNews(&createNewsRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "News created successfully"})
}


func (c *NewsHandlers) GetNewsByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	println(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	news, err := c.newsService.GetNewsByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
	}

	return ctx.JSON(news)
}

// UpdateNews อัพเดทข่าว
// @Summary อัพเดทข่าว
// @Description อัพเดทข้อมูลข่าวตาม ID ที่ระบุ
// @Tags news
// @Accept json
// @Produce json
// @Param id path int true "ID ของข่าว"
// @Param news body model.UpdateNewsRequest true "ข้อมูลข่าวที่จะอัพเดท"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /news/{id} [put]
func (c *NewsHandlers) UpdateNews(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var updateRequest model.UpdateNewsRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.newsService.UpdateNews(uint(id), &updateRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "News updated successfully"})
}

// DeleteNews ลบข่าว
// @Summary ลบข่าว
// @Description ลบข่าวตาม ID ที่ระบุ
// @Tags news
// @Accept json
// @Produce json
// @Param id path int true "ID ของข่าว"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /news/{id} [delete]
func (c *NewsHandlers) DeleteNews(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := c.newsService.DeleteNews(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "News deleted successfully"})
}


func (c *NewsHandlers) GetAllNews(ctx *fiber.Ctx) error {
	news, err := c.newsService.GetAllNews()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(news)
}
