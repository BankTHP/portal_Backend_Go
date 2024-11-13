package routes

import (
	"pccth/portal-blog/internal/handler"
	"pccth/portal-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func FeedbackRoutes(app *fiber.App, feedbackHandler *handler.FeedbackHandler, authMiddleware *middleware.AuthMiddleware) {
	api := app.Group("/feedback")
	api.Post("/create",authMiddleware.HasRole("client_user", "client_admin"), feedbackHandler.CreateFeedback)
	api.Get("/getById/:id",authMiddleware.HasRole("client_user", "client_admin"), feedbackHandler.GetFeedbackById)
	api.Delete("/delete/:id",authMiddleware.HasRole("client_user", "client_admin"), feedbackHandler.DeleteFeedback)
	api.Get("/getPaginated",authMiddleware.HasRole("client_user", "client_admin"), feedbackHandler.GetPaginatedFeedbacks)
	api.Get("/getAll",authMiddleware.HasRole("client_user", "client_admin"), feedbackHandler.GetAllFeedbacks)
} 