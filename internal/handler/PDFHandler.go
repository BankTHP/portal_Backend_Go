package handler

import (
	"pccth/portal-blog/internal/service"
)

type PDFHandler struct {
    pdfService *service.PDFService
}

func NewPDFHandler(pdfService *service.PDFService) *PDFHandler {
    return &PDFHandler{
        pdfService: pdfService,
    }
}