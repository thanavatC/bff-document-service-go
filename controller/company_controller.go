package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanavatC/bff-document-service-go/model"
	"github.com/thanavatC/bff-document-service-go/service"
)

type CompanyController struct {
	service service.CompanyService
}

func NewCompanyController(service service.CompanyService) *CompanyController {
	return &CompanyController{
		service: service,
	}
}

func (c *CompanyController) GetCompanies(ctx *gin.Context) {
	companies, err := c.service.GetCompanies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}
