package routes

import (
	"microservice/infra"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutes(r, ctx)
	return r
}

func authRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	authGroup := r.Group("/customer")
	authGroup.Use(ctx.Middleware.Authentication())
	authGroup.POST("/check-email", ctx.Ctl.User.CheckDetailByEmailUserController)
	authGroup.PUT("/update", ctx.Ctl.User.UpdateUserByEmailController)
}
