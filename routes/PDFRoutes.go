package routes

import (
    "log"
    "os"
    "pccth/portal-blog/internal/handler"
    "pccth/portal-blog/internal/middleware"
    "github.com/gofiber/fiber/v2"
)

func PDFRoutes(app *fiber.App, pdfHandler *handler.PDFHandler, authMiddleware *middleware.AuthMiddleware) {
    uploadDir := "./uploads/pdfs"
    if err := os.MkdirAll(uploadDir, 0755); err != nil {
        log.Fatalf("ไม่สามารถสร้าง upload directory ได้: %v", err)
    }

    app.Static("/pdfs", "./uploads/pdfs")

    api := app.Group("/pdf")
    api.Post("/upload", authMiddleware.HasRole("client_user", "client_admin"), pdfHandler.UploadPDF)
    api.Get("/getpdfByName/:name", authMiddleware.HasRole("client_user", "client_admin"), pdfHandler.GetPDFByName)
} 