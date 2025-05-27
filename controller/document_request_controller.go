package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanavatC/bff-document-service-go/model"
	"github.com/thanavatC/bff-document-service-go/service"
)

type DocumentRequestController struct {
	service service.DocumentRequestService
}

func NewDocumentRequestController(service service.DocumentRequestService) *DocumentRequestController {
	return &DocumentRequestController{
		service: service,
	}
}

func (c *DocumentRequestController) CreateDocumentRequest(ctx *gin.Context) {
	var req model.CreateDocumentRequestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request"})
		return
	}

	docRequest, err := c.service.CreateDocumentRequest(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	response := model.DocumentRequestResponse{
		ID:         docRequest.ID,
		CompanyID:  docRequest.CompanyID,
		DocumentID: docRequest.DocumentID,
		Status:     docRequest.Status,
		CreatedAt:  docRequest.CreatedAt,
		UpdatedAt:  docRequest.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *DocumentRequestController) ListDocumentRequests(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	requests, total, err := c.service.ListDocumentRequests(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	var responseRequests []model.DocumentRequestResponse
	for _, req := range requests {
		responseRequests = append(responseRequests, model.DocumentRequestResponse{
			ID:         req.ID,
			CompanyID:  req.CompanyID,
			DocumentID: req.DocumentID,
			Status:     req.Status,
			CreatedAt:  req.CreatedAt,
			UpdatedAt:  req.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, model.DocumentRequestListResponse{
		Requests: responseRequests,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

func (c *DocumentRequestController) DeleteDocumentRequest(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteDocumentRequest(id); err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Document request not found"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *DocumentRequestController) ValidateDocumentRequest(ctx *gin.Context) {
	id := ctx.Param("id")
	var req model.ValidateDocumentRequestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request"})
		return
	}

	docRequest, err := c.service.ValidateDocumentRequest(id, req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Document request not found"})
		return
	}

	response := model.ValidateDocumentRequestResponse{
		ID:         docRequest.ID,
		CompanyID:  docRequest.CompanyID,
		DocumentID: docRequest.DocumentID,
		Status:     string(docRequest.Status),
		CreatedAt:  docRequest.CreatedAt,
		UpdatedAt:  docRequest.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}
