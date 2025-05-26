package router

import (
	"github.com/SPVJ/document-service-go/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(documentController *controller.DocumentController, documentRequestController *controller.DocumentRequestController, companyController *controller.CompanyController, app *gin.Engine) {
	// Document routes
	documents := app.Group("/api/documents")
	{
		documents.GET("/:id/status", documentController.GetDocumentStatus)
		documents.POST("/:id/retranslate", documentController.ReTranslateDocument)
		documents.PATCH("/:id", documentController.UpdateDocument)
		documents.DELETE("/:id", documentController.DeleteDocument)
		documents.POST("/:id/validate", documentController.ValidateDocument)
	}

	// Document request routes
	requests := app.Group("/api/document-requests")
	{
		requests.POST("", documentRequestController.CreateDocumentRequest)
		requests.GET("", documentRequestController.ListDocumentRequests)
		requests.DELETE("/:id", documentRequestController.DeleteDocumentRequest)
		requests.POST("/:id/validate", documentRequestController.ValidateDocumentRequest)
	}

	// Company routes
	companies := app.Group("/api/companies")
	{
		companies.GET("", companyController.GetCompanies)
	}
}
