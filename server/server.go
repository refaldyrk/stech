package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"kreditplus-test/config"
	"kreditplus-test/controller"
	"kreditplus-test/middleware"
	"kreditplus-test/repository"
	"kreditplus-test/service"
	"kreditplus-test/store"
)

func App(ctx context.Context) *gin.Engine {
	//Initial Config
	configBase := config.NewConfig()
	configBase.InitialConfig()

	//Initial Store Key
	stores := store.NewStore()

	//Repository
	customerRepository := repository.NewCustomerRepository(configBase.SqlDb)
	limitRepository := repository.NewLimitRepository(configBase.SqlDb)
	transactionRepository := repository.NewTransaksiRepository(configBase.SqlDb)

	//Service
	customerService := service.NewCustomerService(customerRepository, stores, configBase)
	limitService := service.NewLimitService(limitRepository)
	transactionService := service.NewTransaksiService(transactionRepository, limitRepository)

	//Controller
	customerController := controller.NewCustomerController(customerService)
	limitController := controller.NewLimitController(limitService)
	transactionController := controller.NewTransaksiController(transactionService)

	//Middleware
	pasetoMiddleware := middleware.PasetoMiddleware(configBase, customerRepository, stores.GetKey())
	app := gin.Default()

	app.Use(gin.Recovery())
	app.Use(gin.Logger())

	// cors	config
	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{"*"}
	cfg.AllowCredentials = true
	cfg.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	cfg.AllowHeaders = []string{"*"}

	app.Use(cors.New(cfg))

	apiRoutesV1 := app.Group("/api/v1")

	//Auth
	apiRoutesV1.POST("/login", customerController.Login)
	apiRoutesV1.POST("/register", customerController.Register)
	apiRoutesV1.POST("/kyc", pasetoMiddleware, customerController.UploadKYC)

	//User
	apiRoutesV1.GET("/user", pasetoMiddleware, customerController.GetCurrentUser)

	//Limit
	apiRoutesV1.GET("/user/limit", pasetoMiddleware, limitController.GetAllLimitByUserID)
	apiRoutesV1.GET("/user/limit/:tenor", pasetoMiddleware, limitController.GetLimitByTenor)

	//Transaction
	apiRoutesV1.POST("/transaction", pasetoMiddleware, transactionController.CreateTransaction)
	apiRoutesV1.GET("/transaction", pasetoMiddleware, transactionController.GetAllTransactionCurrentUser)
	return app
}
