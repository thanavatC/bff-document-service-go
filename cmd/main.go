package main

import (
	"log"

	core "github.com/SPVJ/fs-common-lib/core/client"
	"github.com/SPVJ/fs-common-lib/core/logger"
	"github.com/SPVJ/fs-common-lib/core/middleware"
	"github.com/gin-gonic/gin"
	"github.com/thanavatC/bff-document-service-go/client"
	"github.com/thanavatC/bff-document-service-go/config"
	"github.com/thanavatC/bff-document-service-go/controller"
	"github.com/thanavatC/bff-document-service-go/router"
	"github.com/thanavatC/bff-document-service-go/service"
)

func main() {
	err := config.LoadConfig("/config")
	if err != nil {
		log.Fatalln("Error loading config:", err)
	}

	loggerConfig := logger.Config{
		IsProductionMode: config.AppConfig.Server.IsProductionMode,
	}

	newLogger := logger.NewLogger(loggerConfig)

	app := gin.Default()

	// TODO: common middleware stack
	app.Use(gin.Recovery())
	app.Use(middleware.New(newLogger))

	// Initialize repositories
	documentClient := client.NewDocumentServiceClientImpl(
		core.NewHttpClient(config.AppConfig.Webclient.DocumentService.HttpClientConfig))
	documentRequestClient := client.NewDocumentRequestServiceClientImpl(
		core.NewHttpClient(config.AppConfig.Webclient.DocumentService.HttpClientConfig))
	companyClient := client.NewCompanyServiceClientImpl(
		core.NewHttpClient(config.AppConfig.Webclient.DocumentService.HttpClientConfig))

	// Initialize services
	documentService := service.NewDocumentServiceImpl(documentClient)
	documentRequestService := service.NewDocumentRequestServiceImpl(documentRequestClient)
	companyService := service.NewCompanyServiceImpl(companyClient)

	// Initialize controllers
	documentController := controller.NewDocumentController(documentService)
	documentRequestController := controller.NewDocumentRequestController(documentRequestService)
	companyController := controller.NewCompanyController(companyService)

	// Setup routes
	router.SetupRouter(documentController, documentRequestController, companyController, app)

	if err := app.Run(config.AppConfig.Server.Port); err != nil {
		log.Fatalln(err)
	}
}
