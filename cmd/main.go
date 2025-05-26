package main

import (
	"fmt"
	"log"

	"github.com/SPVJ/document-service-go/config"
	"github.com/SPVJ/document-service-go/controller"
	"github.com/SPVJ/document-service-go/repository"
	"github.com/SPVJ/document-service-go/router"
	"github.com/SPVJ/document-service-go/service"
	"github.com/SPVJ/fs-common-lib/core/db"
	"github.com/SPVJ/fs-common-lib/core/logger"
	"github.com/SPVJ/fs-common-lib/core/middleware"
	"github.com/gin-gonic/gin"
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

	db := db.New(config.AppConfig.Database)

	fmt.Println(config.AppConfig.Database)

	app := gin.Default()
	app.Use(gin.Recovery())
	app.Use(middleware.New(newLogger))

	// Initialize repositories
	documentRepo := repository.NewDocumentRepositoryImpl(db)
	documentRequestRepo := repository.NewDocumentRequestRepositoryImpl(db)
	companyRepo := repository.NewCompanyRepositoryImpl(db)

	// Initialize services
	documentService := service.NewDocumentService(documentRepo)
	documentRequestService := service.NewDocumentRequestService(documentRequestRepo)
	companyService := service.NewCompanyService(companyRepo)

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
