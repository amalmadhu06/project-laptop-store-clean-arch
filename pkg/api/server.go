package http

import (
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/api/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//_ "github.com/thnkrn/go-gin-clean-arch/cmd/api/docs"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, otpHandler *handler.OtpHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// request JWT

	engine.POST("/signup", userHandler.CreateUser)
	engine.POST("/login", userHandler.UserLogin)
	engine.POST("/sendOTP", otpHandler.SendOtp)
	engine.POST("/verifyOTP", otpHandler.ValidateOtp)

	//Auth middleware
	//api := engine.Group("/api", middleware.AuthorizationMiddleware)
	//
	//api.GET("users", userHandler.FindAll)
	//api.GET("users/:id", userHandler.FindByID)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
