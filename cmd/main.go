package main

import (
	"context"
	"database/sql"
	"dtonetest/config"
	"dtonetest/internal/controller"
	"dtonetest/internal/services"
	"dtonetest/internal/use_cases"
	"dtonetest/repositories"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"log"
)

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

	createUser := use_cases.NewCreateUserUseCase(mongoUserRepository)
	userController := controller.NewRegisterController(createUser)

	webTokenService, err := services.NewWebTokenService(cnf.ApiSecret, cnf.TokenLifeSpan)
	if err != nil {
		panic(err)
	}
	login := use_cases.NewLoginUseCase(mongoUserRepository, webTokenService)
	loginController := controller.NewLoginController(login)

	topUpUseCase := use_cases.NewTopUpUserUseCase(mongoUserRepository)
	topUpController := controller.NewTopUpUserController(topUpUseCase)

	getOneUserUseCase := use_cases.NewGetOneUserUseCase(mongoUserRepository)
	getOneUserController := controller.NewGetOneUserController(getOneUserUseCase)

	createProductUseCase := use_cases.NewCreateProductUseCase(mongoProductRepository, mongoUserRepository)
	createProductController := controller.NewCreateProductController(createProductUseCase, cnf.FolderRepository)

	r := gin.Default()
	control := r.Group("/")
	control.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "running"})
	})
	public := r.Group("/api/v1")
	public.Use(otelgin.Middleware("DTOne"))
	public.POST("/register", userController.Handle)
	public.POST("/login", loginController.Login)

	protected := r.Group("/api/v1")
	protected.Use(otelgin.Middleware("DTOne"))
	protected.Use(services.JwtAuthMiddleware(cnf.ApiSecret))
	protected.PUT("users/:user_id/topup", topUpController.Handle)
	protected.GET("users/:user_id", getOneUserController.Handle)
	protected.POST("products", createProductController.Handle)
	protected.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "castaña"})
	})

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
