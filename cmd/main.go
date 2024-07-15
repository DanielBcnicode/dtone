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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	// Load configuration
	cnf, err := config.GetConfiguration()
	if err != nil {
		panic(err)
	}

	// Tracing initialization
	tracer := initTracer(cnf.OTL.ServiceName, cnf.OTL.CollectorUrl, cnf.OTL.InsecureCollector)
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
	createUser := use_cases.NewCreateUserUseCase(mongoUserRepository)
	userController := controller.NewRegisterController(createUser)

	webTokenService, err := services.NewWebTokenService(cnf.ApiSecret, cnf.TokenLifeSpan)
	if err != nil {
		panic(err)
	}
	login, err := use_cases.NewLoginUseCase(mongoUserRepository, webTokenService)
	if err != nil {
		panic(err)
	}
	loginController := controller.NewLoginController(login)

	r := gin.Default()
	r.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "running"})
	})
	public := r.Group("/api/v1")
	public.Use(otelgin.Middleware("DTOne"))
	public.POST("/register", userController.Register)
	public.POST("/login", loginController.Login)

	protected := r.Group("/api/v1")
	protected.Use(otelgin.Middleware("DTOne"))
	protected.Use(services.JwtAuthMiddleware(cnf.ApiSecret))
	protected.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "castaña"})
	})

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func initTracer(serviceName string, collectorUrl string, insecure string) func(context.Context) error {
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(insecure) > 0 {
		secureOption = otlptracegrpc.WithInsecure()
	}
	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(collectorUrl),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Printf("could not set resources: %v", err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resources),
		),
	)

	return exporter.Shutdown
}
