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

func NewServerHTTP(
	userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	otpHandler *handler.OtpHandler,
	productHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	orderHandler *handler.OrderHandler,
	paymentHanlder *handler.PaymentHandler,
) *ServerHTTP {
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

			adminPanel.GET("dashboard", adminHandler.AdminDashboard)

			adminPanel.GET("/users", userHandler.ListAllUsers)
			adminPanel.GET("/users/:id", userHandler.FindUserByID)
			adminPanel.PUT("/users/block-user", userHandler.BlockUser)
			adminPanel.PUT("users/unblock-user/:id", userHandler.UnblockUser)

			adminPanel.POST("/create-admin", adminHandler.CreateAdmin)
			adminPanel.PUT("/block-admin/:admin-id", adminHandler.BlockAdmin)
			adminPanel.PUT("/unblock-admin/:admin-id", adminHandler.UnblockAdmin)

			//	category management
			adminPanel.POST("/create-category", productHandler.CreateCategory)
			adminPanel.GET("/view-all-categories", productHandler.ViewAllCategories)
			adminPanel.GET("/view-category/:id", productHandler.FindCategoryByID)
			adminPanel.PUT("update-category", productHandler.UpdateCategory)
			adminPanel.DELETE("delete-category", productHandler.DeleteCategory)

			//brand management
			adminPanel.POST("/create-brand", productHandler.CreateBrand)
			adminPanel.GET("/view-all-brands", productHandler.ViewAllBrands)
			adminPanel.GET("/view-brand-by-id/:id", productHandler.ViewBrandByID)
			adminPanel.PUT("/update-brand", productHandler.UpdateBrand)
			adminPanel.DELETE("/delete-brand/:id", productHandler.DeleteBrand)

			//	product management
			adminPanel.POST("/create-product", productHandler.CreateProduct)
			adminPanel.GET("/view-all-products", productHandler.ViewAllProducts)
			adminPanel.GET("/view-product/:id", productHandler.FindProductByID)
			adminPanel.PUT("update-product", productHandler.UpdateProduct)
			adminPanel.DELETE("delete-product", productHandler.DeleteProduct)

			//	product item management
			adminPanel.POST("/create-product-item", productHandler.CreateProductItem)
			adminPanel.GET("/view-all-product-items", productHandler.ViewAllProductItems)
			adminPanel.GET("/view-product-item/:id", productHandler.FindProductItemByID)
			adminPanel.PUT("/update-product-item", productHandler.UpdateProductItem)
			adminPanel.DELETE("/delete-product-item", productHandler.DeleteProductItem)

			//	order management
			adminPanel.PUT("/update-order", orderHandler.UpdateOrder)
		}
	}

	//User routes
	user := engine.Group("/")
	{
		//	User routes that doesn't require middleware checking
		user.POST("signup", userHandler.CreateUser)
		user.POST("login-email", userHandler.LoginWithEmail)
		user.POST("login-phone", userHandler.LoginWithPhone)
		user.GET("logout", userHandler.UserLogout)
		user.POST("sendOTP", otpHandler.SendOtp)
		user.POST("verifyOTP", otpHandler.ValidateOtp)

		//Todo : write separate handlers
		user.GET("/view-all-categories", productHandler.ViewAllCategories)
		user.GET("/view-category/:id", productHandler.FindCategoryByID)

		user.GET("/view-all-brands", productHandler.ViewAllBrands)
		user.GET("/view-brand-by-id/:id", productHandler.ViewBrandByID)

		user.GET("/view-all-products", productHandler.ViewAllProducts)
		user.GET("/view-product/:id", productHandler.FindProductByID)

		user.GET("/view-all-product-items", productHandler.ViewAllProductItems)
		user.GET("/view-product-item/:id", productHandler.FindProductItemByID)

		//User routes that require middleware checking
		user.Use(middleware.UserAuth)

		{
			user.GET("profile", userHandler.UserProfile)

			user.POST("/add-address", userHandler.AddAddress)
			user.PUT("/update-address", userHandler.UpdateAddress)
			user.POST("/add-to-cart/:product_item_id", cartHandler.AddToCart)
			user.DELETE("/remove-from-cart/:product_item_id", cartHandler.RemoveFromCart)
			user.GET("view-cart", cartHandler.ViewCart)
			user.DELETE("/empty-cart", cartHandler.EmptyCart)

			user.POST("/buy-product-item", orderHandler.BuyProductItem)
			user.POST("/cart/buy-all", orderHandler.BuyAll)
			user.GET("/view-order-by-id/:order_id", orderHandler.ViewOrderByID)
			user.GET("/view-all-orders", orderHandler.ViewAllOrders)
			user.PUT("/cancel-order/:order_id", orderHandler.CancelOrder)

			user.GET("/order/razorpay/:order_id", paymentHanlder.CreateRazorpayPayment)
			user.GET("/payment-handler", paymentHanlder.PaymentSuccess)

		}

	}
	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.LoadHTMLGlob("template/*.html")
	err := sh.engine.Run(":3000")
	if err != nil {
		return
	}

}
