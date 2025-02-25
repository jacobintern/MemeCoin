package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jacobintern/meme_coin/cmd/injection"
	"github.com/jacobintern/meme_coin/controller"
	"github.com/jacobintern/meme_coin/docs"
	"github.com/jacobintern/meme_coin/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func main() {
	i := injection.New()

	// logger
	logger := i.ZapLog

	// get gin
	r := gin.New()

	// recovery
	r.Use(gin.Recovery())

	// cors
	r.Use(cors.New(
		cors.Config{
			AllowAllOrigins:  i.Config.Cors.ALLOW.ALLOrigins,
			AllowMethods:     i.Config.Cors.ALLOW.Methods,
			AllowHeaders:     i.Config.Cors.ALLOW.Headers,
			AllowCredentials: i.Config.Cors.ALLOW.Credentials,
			MaxAge:           time.Duration(i.Config.Cors.MaxAge) * time.Second,
		}))

	// handlers
	registerRoutes(r, i)

	// start up
	err := r.Run(fmt.Sprintf(":%s", i.Config.Server.Port))
	if err != nil {
		logger.Panic("server start up failed.", zap.Error(err))
	}

	logger.Info("the shop backend service successfully start up")
}

func registerRoutes(r *gin.Engine, i *injection.Injection) {
	r.Use(cors.Default())

	// swagger
	r.GET("/swagger/*any", swaggerRouteHandler())

	// controller
	memeCoinController := controller.NewMemeCoinController(i.MemeCoinService)

	api := r.Group("/api")

	// api
	memeCoin := api.Group("/meme_coin", middleware.TimeoutMiddleware(time.Duration(i.Config.Http.TimeoutSec)*time.Second))
	{
		memeCoin.POST("", memeCoinController.Create)
		memeCoin.GET("/:id", memeCoinController.Get)
		memeCoin.PATCH("/:id", memeCoinController.Patch)
		memeCoin.DELETE("/:id", memeCoinController.Delete)
		memeCoin.POST("/:id/poke", memeCoinController.Poke)
	}
}

// swaggerRouteHandler sets up swagger settings
func swaggerRouteHandler() gin.HandlerFunc {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Meme Coin API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
