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

// CreatePost godoc
// @Summary สร้างโพสต์ใหม่
// @Description สร้างโพสต์ใหม่จากข้อมูลที่ส่งมา
// @Tags posts
// @Accept json
// @Produce json
// @Param post body model.CreatePostRequest true "ข้อมูลโพสต์ที่จะสร้าง"
// @Success 201 {object} map[string]interface{}{"message": "Post created successfully"}
// @Failure 400 {object} map[string]interface{}{"error": "Invalid input"}
// @Failure 500 {object} map[string]interface{}{"error": "Internal server error"}
// @Router /posts/createPost [post]
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

// @Summary อัปเดตโพสต์
// @Description อัปเดตข้อมูลโพสต์ตาม ID ที่ระบุ
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "ID ของโพสต์"
// @Param updatePostRequest body model.UpdatePostRequest true "ข้อมูลโพสต์ที่จะอัปเดต"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /posts/{id} [put]
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

// @Summary ลบโพสต์
// @Description ลบโพสต์ตาม ID ที่ระบุ
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "ID ของโพสต์"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /posts/{id} [delete]
func (c *PostHandlers) DeletePost(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid ID"})
	}

	if err := c.postService.DeletePost(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(map[string]interface{}{"message": "Post deleted successfully"})
}


func (c *PostHandlers) GetAllPosts(ctx *fiber.Ctx) error {
	posts, err := c.postService.GetAllPosts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(posts)
}

// @Summary ดึงข้อมูลโพสต์แบบแบ่งหน้า
// @Description ดึงข้อมูลโพสต์แบบแบ่งหน้าตามพารามิเตอร์ที่ระบุ
// @Tags posts
// @Accept json
// @Produce json
// @Param page query int false "หมายเลขหน้า (ค่าเริ่มต้น: 1)"
// @Param limit query int false "จำนวนรายการต่อหน้า (ค่าเริ่มต้น: 10)"
// @Success 200 {object} model.PaginatedResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /posts/paginated [get]
func (c *PostHandlers) GetPaginatedPosts(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil || page < 1 {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid page number"})
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
