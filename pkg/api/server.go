package http

import (
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/api/handler"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/api/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, adminHandler *handler.AdminHandler, otpHandler *handler.OtpHandler, productHandler *handler.ProductHandler) *ServerHTTP {
	engine := gin.New()

	// logger middleware logs following info for each request : http method, req URL, remote address of the client, res status code, elapsed time
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//Admin routes
	adminPanel := engine.Group("/adminPanel")
	{
		//	Admin routes that doesn't need middleware checking
		adminPanel.POST("/login", adminHandler.AdminLogin)
		adminPanel.GET("/logout", adminHandler.AdminLogout)

		//	Admin routes that requires middleware checking
		adminPanel.Use(middleware.AdminAuth)
		{

			adminPanel.POST("/create-admin", adminHandler.CreateAdmin)

			//	category management
			adminPanel.POST("/create-category", productHandler.CreateCategory)
			adminPanel.GET("/view-all-categories", productHandler.ViewAllCategories)
			adminPanel.GET("/view-category/:id", productHandler.FindCategoryByID)
			adminPanel.PUT("update-category", productHandler.UpdateCategory)
			adminPanel.DELETE("delete-category", productHandler.DeleteCategory)

			//	product management
			adminPanel.POST("/create-product", productHandler.CreateProduct)
			adminPanel.GET("/view-all-products", productHandler.ViewAllProducts)
			adminPanel.GET("/view-product/:id", productHandler.FindProductByID)
			adminPanel.PUT("update-product", productHandler.UpdateProduct)
			adminPanel.DELETE("delete-product", productHandler.DeleteProduct)

		}
	}

	//User routes
	user := engine.Group("/")
	{
		//	User routes that doesn't require middleware checking
		user.POST("signup", userHandler.CreateUser)
		user.POST("login-email", userHandler.LoginWithEmail)
		user.POST("login-phone", userHandler.LoginWithPhone)
		user.POST("sendOTP", otpHandler.SendOtp)
		user.POST("verifyOTP", otpHandler.ValidateOtp)

		//User routes that require middleware checking
		user.Use(middleware.UserAuth)

	}
	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	err := sh.engine.Run(":3000")
	if err != nil {
		return
	}

}
