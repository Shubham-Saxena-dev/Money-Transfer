package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"qonto/internal/config"
	"qonto/internal/database"
	"qonto/internal/repositories"
	"qonto/internal/routes"
	"qonto/pkg/handlers"
	"qonto/pkg/service"
)

var (
	controller handlers.AccountController
)

func init() {
	config.InitFromFile(".env")
}

func main() {
	log.Info("Hi, this is Qonto take home test")
	initializeDb()
	createServer()
}

func initializeDb() {
	db := database.NewDatabase()
	db.InitialiseDbConnection()
}

func createServer() {
	server := gin.Default()
	initializeLayers()
	routes.RegisterHandlers(server, controller).RegisterHandlers()
	err := server.Run(":" + config.EnvConfigs.App.AppPort)
	if err != nil {
		gin.Logger()
		log.Error(err)
		os.Exit(1)
	}
}

func initializeLayers() {
	bankAccountRepo := repositories.NewBankAccountRepository()
	transactionRepo := repositories.NewTransactionRepository(bankAccountRepo)
	transferService := service.NewTransferService(transactionRepo)
	controller = handlers.NewController(transferService)
}
