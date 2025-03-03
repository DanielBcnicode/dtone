package main

import (
	"context"
	"database/sql"
	"dtonetest/config"
	_ "dtonetest/docs"
	"dtonetest/internal/controller"
	"dtonetest/internal/services"
	"dtonetest/internal/use_cases"
	"dtonetest/repositories"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"log"
)

// @title DTOne Swagger API
// @version 1.0
// @description Api test.

// @contact.name Daniel

// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /
// @schemes http
// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200
// @Router /health [get]
func main() {
	// Load configuration
	cnf, err := config.GetConfiguration()
	if err != nil {
		panic(err)
	}

	// Tracing initialization
	tracer := services.InitTracer(cnf.OTL.ServiceName, cnf.OTL.CollectorUrl, cnf.OTL.InsecureCollector)
	defer func() {
		err := tracer(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	// DB connection
	db, err := cnf.Database.DatabaseConnection()
	if err != nil {
		panic(err)
	}
	connection, err := db.DB()
	defer func(connection *sql.DB) {
		err := connection.Close()
		if err != nil {
			log.Print(err)
		}
	}(connection)

	mongoUserRepository := repositories.NewMongoUserRepository(db)
	mongoProductRepository := repositories.NewMongoProductRepository(db)
	mongoTransactionRepository := repositories.NewMongoTransactionRepository(db)

	createUser := use_cases.NewCreateUserUseCase(mongoUserRepository)
	userController := controller.NewRegisterController(createUser)

	webTokenService, err := services.NewWebTokenService(cnf.ApiSecret, cnf.TokenLifeSpan)
	if err != nil {
		panic(err)
	}
	login := use_cases.NewLoginUseCase(mongoUserRepository, webTokenService)
	loginController := controller.NewLoginController(login)

	topUpUseCase := use_cases.NewTopUpUserUseCase(mongoUserRepository, mongoTransactionRepository)
	topUpController := controller.NewTopUpUserController(topUpUseCase)

	getOneUserUseCase := use_cases.NewGetOneUserUseCase(mongoUserRepository)
	getOneUserController := controller.NewGetOneUserController(getOneUserUseCase)

	createProductUseCase := use_cases.NewCreateProductUseCase(mongoProductRepository, mongoUserRepository)
	createProductController := controller.NewCreateProductController(createProductUseCase, cnf.FolderRepository)

	uploadProductUseCase := use_cases.NewUploadProductUseCase(mongoProductRepository, mongoUserRepository, cnf.FolderRepository)
	uploadProductController := controller.NewUploadProductController(uploadProductUseCase, cnf.FolderRepository)

	getOneProductUseCase := use_cases.NewGetOneProductUseCase(mongoProductRepository)
	getOneProductController := controller.NewGetOneProductController(getOneProductUseCase)

	getAllProductsUseCase := use_cases.NewGetAllProductsUseCase(mongoProductRepository)
	getAllProductsController := controller.NewGetAllProductsController(getAllProductsUseCase)

	buyProductUseCase := use_cases.NewBuyProductUseCase(mongoProductRepository, mongoUserRepository, mongoTransactionRepository)
	buyProductController := controller.NewBuyProductController(buyProductUseCase)

	getUserTransactionsUseCase := use_cases.NewGetUserTransactionsUseCase(mongoTransactionRepository)
	getUserTransactionsController := controller.NewGetUserTransactionsController(getUserTransactionsUseCase)

	r := gin.Default()
	control := r.Group("/")
	control.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "running"})
	})
	public := r.Group("/api/v1")
	public.Use(otelgin.Middleware("DTOne"))
	public.POST("/register", userController.Handle)
	public.POST("/login", loginController.Login)
	public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	protected := r.Group("/api/v1")
	protected.Use(otelgin.Middleware("DTOne"))
	protected.Use(services.JwtAuthMiddleware(cnf.ApiSecret))
	protected.PUT("users/:user_id/topup", topUpController.Handle)
	protected.GET("users/:user_id", getOneUserController.Handle)
	protected.GET("users/:user_id/transactions", getUserTransactionsController.Handle)
	protected.POST("products", createProductController.Handle)
	protected.POST("products/:product_id/file", uploadProductController.Handle)
	protected.POST("products/:product_id/buy", buyProductController.HandleBuy)
	protected.POST("products/:product_id/gift", buyProductController.HandleGift)
	protected.GET("products", getAllProductsController.Handle)
	protected.GET("products/:product_id", getOneProductController.Handle)

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
