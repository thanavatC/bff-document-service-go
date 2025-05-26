package controller

import (
	"net/http"

	"github.com/SPVJ/document-service-go/model"
	"github.com/SPVJ/document-service-go/service"
	"github.com/gin-gonic/gin"
)

type DocumentController struct {
	service *service.DocumentService
}

func NewDocumentController(service *service.DocumentService) *DocumentController {
	return &DocumentController{
		service: service,
	}
}

func (c *DocumentController) GetDocumentStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	status, err := c.service.GetDocumentStatus(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Document not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": status})
}

func (c *DocumentController) ReTranslateDocument(ctx *gin.Context) {
	id := ctx.Param("id")
	doc, err := c.service.ReTranslateDocument(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Document not found"})
		return
	}

	response := model.DocumentResponse{
		ID:        doc.ID,
		CompanyID: doc.CompanyID,
		Name:      doc.Name,
		Type:      doc.Type,
		Status:    doc.Status,
		FileURL:   doc.FileURL,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: doc.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *DocumentController) UpdateDocument(ctx *gin.Context) {
	id := ctx.Param("id")
	var req model.UpdateDocumentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request"})
		return
	}

	doc, err := c.service.UpdateDocument(id, req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Document not found"})
		return
	}

	ctx.JSON(http.StatusOK, doc)
}

func (c *DocumentController) DeleteDocument(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteDocument(id); err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Document not found"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *DocumentController) ValidateDocument(ctx *gin.Context) {
	id := ctx.Param("id")
	var req model.ValidateDocumentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request"})
		return
	}

	doc, err := c.service.ValidateDocument(id, req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Document not found"})
		return
	}

	response := model.ValidateDocumentResponse{
		ID:        doc.ID,
		CompanyID: doc.CompanyID,
		Name:      doc.Name,
		Type:      doc.Type,
		Status:    doc.Status,
		FileURL:   doc.FileURL,
		CreatedAt: doc.CreatedAt,
		UpdatedAt: doc.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}
