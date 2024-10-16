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
// @Description สร้างโพสต์ใหม่ในระบบ
// @Tags posts
// @Accept json
// @Produce json
// @Param post body model.CreatePostRequest true "ข้อมูลโพสต์"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
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

// GetPostByID godoc
// @Summary ดึงข้อมูลโพสต์โดย ID
// @Description ดึงข้อมูลโพสต์โดย ID
// @Tags posts
// @Produce json
// @Param id path string true "ID ของโพสต์"
// @Success 200 {object} model.Post
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /posts/{id} [get]
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

// UpdatePost godoc
// @Summary อั���เดตโพสต์
// @Description อัปเดตโพสต์ในระบบ
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "ID ของโพสต์"
// @Param post body model.UpdatePostRequest true "ข้อมูลโพสต์"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /posts/{id} [patch]
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

// DeletePost godoc
// @Summary ลบโพสต์
// @Description ลบโพสต์ในระบบ
// @Tags posts
// @Produce json
// @Param id path string true "ID ของโพสต์"
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

// GetAllPosts godoc
// @Summary แสดงรายการโพสต์ทั้งหมด
// @Description ดึงข้อมูลโพสต์ทั้งหมดจากระบบ
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {array} model.Post
// @Failure 500 {object} map[string]interface{}
// @Router /posts/getAllPosts [get]
func (c *PostHandlers) GetAllPosts(ctx *fiber.Ctx) error {
	posts, err := c.postService.GetAllPosts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(posts)
}

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

