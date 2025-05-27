package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thanavatC/bff-document-service-go/controller"
)

func SetupRouter(documentController *controller.DocumentController, documentRequestController *controller.DocumentRequestController, companyController *controller.CompanyController, app *gin.Engine) {
	// API versioning
	v1 := app.Group("/api/v1")
	{
		// Document routes
		documents := v1.Group("/documents")
		{
			documents.GET("/:id/status", documentController.GetDocumentStatus)
			documents.POST("/:id/retranslate", documentController.ReTranslateDocument)
			documents.PATCH("/:id", documentController.UpdateDocument)
			documents.DELETE("/:id", documentController.DeleteDocument)
			documents.POST("/:id/validate", documentController.ValidateDocument)
		}

		// Document request routes
		requests := v1.Group("/document-requests")
		{
			requests.POST("", documentRequestController.CreateDocumentRequest)
			requests.GET("", documentRequestController.ListDocumentRequests)
			requests.DELETE("/:id", documentRequestController.DeleteDocumentRequest)
			requests.POST("/:id/validate", documentRequestController.ValidateDocumentRequest)
		}

		// Company routes
		companies := v1.Group("/companies")
		{
			companies.GET("", companyController.GetCompanies)
		}

		// Health check endpoint
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"service": "document-service",
			})
		})
	}
}
