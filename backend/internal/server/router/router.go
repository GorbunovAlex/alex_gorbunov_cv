package router

import (
	"go.mongodb.org/mongo-driver/mongo"
	//_ "gorbunov_alex.space/cv/docs"
	mLogger "gorbunov_alex.space/cv/internal/server/middleware/logger"
	"log/slog"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(log *slog.Logger, storage *mongo.Database) http.Handler {
	router := gin.Default()

	router.Use(mLogger.New(log))
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//v1 := router.Group("/api/v1")
	//{
	//	auth := v1.Group("/")
	//	auth.Use(token.TokenValidationMiddleware())
	//	{
	//		auth.POST("/operations/new", operations.New(log, storage))
	//		auth.GET("/operations/", operations.GetAll(log, storage))
	//		auth.PUT("/operations/:id", operations.Update(log, storage))
	//		auth.DELETE("/operations/:id", operations.Delete(log, storage))
	//
	//		auth.GET("/categories", categories.GetAll(log, storage))
	//		auth.POST("/categories/new", categories.New(log, storage))
	//		auth.PUT("/categories/:id", categories.Update(log, storage))
	//		auth.DELETE("/categories/:id", categories.Delete(log, storage))
	//	}
	//	v1.POST("/users/signup", users.Signup(log, storage))
	//	v1.POST("/users/login", users.Login(log, storage))
	//}

	return router
}
