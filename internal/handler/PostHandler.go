package handler

import (
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/service"

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
	
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_FORM", "รูปแบบข้อมูลไม่ถูกต้อง"))
	}

	
	var createPostRequest model.CreatePostRequest
	if err := ctx.BodyParser(&createPostRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_INPUT", "ข้อมูลไม่ถูกต้อง"))
	}


	files := form.File["pdf"]

	if err := c.postService.CreatePost(&createPostRequest, files); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("CREATE_POST_ERROR", err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.NewSuccessResponse(map[string]string{
		"message": "สร้างโพสต์สำเร็จ",
	}))
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
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_ID", "รหัสไม่ถูกต้อง"))
	}

	post, err := c.postService.GetPostByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(model.NewErrorResponse("POST_NOT_FOUND", "ไม่พบโพสต์"))
	}

	return ctx.JSON(model.NewSuccessResponse(post))
}

// UpdatePost godoc
// @Summary อัพเดตโพสต์
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
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_ID", "รหัสไม่ถูกต้อง"))
	}

	var updateRequest model.UpdatePostRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_INPUT", "ข้อมูลไม่ถูกต้อง"))
	}

	if err := c.postService.UpdatePost(uint(id), &updateRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("UPDATE_POST_ERROR", err.Error()))
	}

	return ctx.JSON(model.NewSuccessResponse(map[string]string{
		"message": "อัปเดตโพสต์สำเร็จ",
	}))
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
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_ID", "รหัสไม่ถูกต้อง"))
	}

	if err := c.postService.DeletePost(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("DELETE_POST_ERROR", err.Error()))
	}

	return ctx.JSON(model.NewSuccessResponse(map[string]string{
		"message": "ลบโพสต์และความคิดเห็นที่เกี่ยวข้องสำเร็จ",
	}))
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("FETCH_POSTS_ERROR", err.Error()))
	}

	return ctx.JSON(model.NewSuccessResponse(posts))
}

// GetPaginatedPosts godoc
// @Summary แสดงรายการโพสต์แบบแบ่งหน้า
// @Description ดึงข้อมูลโพสต์แบบแบ่งหน้าจากระบบ
// @Tags posts
// @Accept json
// @Produce json
// @Param request body model.PostPaginatedRequest true "ข้อมูลการแบ่งหน้า"
// @Success 200 {object} model.PaginatedResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /posts/getPaginatedPosts [post]
func (c *PostHandlers) GetPaginatedPosts(ctx *fiber.Ctx) error {
	var request model.PostPaginatedRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_INPUT", "ข้อมูลไม่ถูกต้อง"))
	}

	if request.Page < 1 {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_PAGE", "หน้าต้องมากกว่า 0"))
	}

	if request.Size < 1 {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_SIZE", "ขนาดต้องมากกว่า 0"))
	}

	paginatedResponse, err := c.postService.GetPaginatedPosts(int(request.Page), int(request.Size))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("FETCH_PAGINATED_POSTS_ERROR", err.Error()))
	}

	return ctx.JSON(model.NewSuccessResponse(paginatedResponse))
}

func (c *PostHandlers) GetPaginatedPostsByUserId(ctx *fiber.Ctx) error {
	var req model.PostByUserIdPaginatedRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_INPUT", "ข้อมูลไม่ถูกต้อง"))
	}

	if req.PostCreateBy == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_USER_ID", "รหัสผู้ใช้ไม่ถูกต้อง"))
	}
	if req.Page == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_PAGE", "หน้าต้องมากกว่า 0"))
	}
	if req.Size == 0 || req.Size > 100 {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_SIZE", "ขนาดต้องอยู่ระหว่าง 1 ถึง 100"))
	}

	paginatedResponse, err := c.postService.GetPaginatedPostsByUserId(int(req.Page), int(req.Size), req.PostCreateBy)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("FETCH_USER_POSTS_ERROR", "ไม่สามารถดึงข้อมูลโพสต์ของผู้ใช้ได้"))
	}

	return ctx.JSON(model.NewSuccessResponse(paginatedResponse))
}
