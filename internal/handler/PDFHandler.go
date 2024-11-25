package handler

import (
    "pccth/portal-blog/internal/model"
    "pccth/portal-blog/internal/service"
    "github.com/gofiber/fiber/v2"
)

type PDFHandler struct {
    pdfService *service.PDFService
}

func NewPDFHandler(pdfService *service.PDFService) *PDFHandler {
    return &PDFHandler{
        pdfService: pdfService,
    }
}

func (h *PDFHandler) UploadPDF(c *fiber.Ctx) error {
    file, err := c.FormFile("pdf")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(model.PDFResponse{
            Success: false,
            Error:   "ไม่พบไฟล์ PDF กรุณาเลือกไฟล์ที่ต้องการอัปโหลด",
        })
    }

    uploadedFile := &model.UploadedPDFFile{
        File:     file,
        Filename: file.Filename,
        Size:     file.Size,
        SaveFunc: func(path string) error {
            return c.SaveFile(file, path)
        },
    }

    response, err := h.pdfService.ProcessPDFUpload(uploadedFile)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(model.PDFResponse{
            Success: false,
            Error:   err.Error(),
        })
    }

    return c.JSON(response)
}

func (h *PDFHandler) GetPDFByName(c *fiber.Ctx) error {
    pdfName := c.Params("name")
    if pdfName == "" {
        return c.Status(fiber.StatusBadRequest).JSON(model.PDFResponse{
            Success: false,
            Error:   "กรุณาระบุชื่อไฟล์ PDF",
        })
    }

    pdf, err := h.pdfService.GetPDFByName(pdfName)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(model.PDFResponse{
            Success: false,
            Error:   err.Error(),
        })
    }

    return c.JSON(model.PDFResponse{
        Success: true,
        FullURL: pdf.PDFPath,
        Size:    pdf.PDFSize,
    })
} 