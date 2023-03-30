package routes

import (
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/api/handler"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(
	api *gin.RouterGroup,
	adminHandler *handler.AdminHandler,
	userHandler *handler.UserHandler,
	productHandler *handler.ProductHandler,
	orderHandler *handler.OrderHandler,
) {

	api.POST("/login", adminHandler.AdminLogin)

	api.Use(middleware.AdminAuth)
	{
		api.GET("/logout", adminHandler.AdminLogout)
		api.GET("/dashboard", adminHandler.AdminDashboard)
		api.GET("/sales-report", adminHandler.SalesReport)

		//user management
		userRoutes := api.Group("/users")
		{
			userRoutes.GET("/", userHandler.ListAllUsers)
			userRoutes.GET("/:id", userHandler.FindUserByID)
			userRoutes.PUT("/block", userHandler.BlockUser)
			userRoutes.PUT("/unblock/:id", userHandler.UnblockUser)
		}

		//admin management
		adminManagement := api.Group("/admins")
		{
			adminManagement.POST("/", adminHandler.CreateAdmin)
			adminManagement.PUT("/:id/block", adminHandler.BlockAdmin)
			adminManagement.PUT("/:id/unblock", adminHandler.UnblockAdmin)

		}

		// Category management routes
		categoryRoutes := api.Group("/categories")
		{
			categoryRoutes.POST("/", productHandler.CreateCategory)
			categoryRoutes.GET("/", productHandler.ViewAllCategories)
			categoryRoutes.GET("/:id", productHandler.FindCategoryByID)
			categoryRoutes.PUT("/", productHandler.UpdateCategory)
			categoryRoutes.DELETE("/:id", productHandler.DeleteCategory)
		}

		// Brand management routes
		brandRoutes := api.Group("/brands")
		{
			brandRoutes.POST("/", productHandler.CreateBrand)
			brandRoutes.GET("/", productHandler.ViewAllBrands)
			brandRoutes.GET("/:id", productHandler.ViewBrandByID)
			brandRoutes.PUT("/", productHandler.UpdateBrand)
			brandRoutes.DELETE("/:id", productHandler.DeleteBrand)
		}
		// Product management routes
		productRoutes := api.Group("/products")
		{
			productRoutes.POST("/", productHandler.CreateProduct)
			productRoutes.GET("/", productHandler.ViewAllProducts)
			productRoutes.GET("/:id", productHandler.FindProductByID)
			productRoutes.PUT("/", productHandler.UpdateProduct)
			productRoutes.DELETE("/:id", productHandler.DeleteProduct)
		}

		// Product item management routes
		productItemRoutes := api.Group("/product-items")
		{
			productItemRoutes.POST("/", productHandler.CreateProductItem)
			productItemRoutes.GET("/", productHandler.ViewAllProductItems)
			productItemRoutes.GET("/:id", productHandler.FindProductItemByID)
			productItemRoutes.PUT("/", productHandler.UpdateProductItem)
			productItemRoutes.DELETE("/:id", productHandler.DeleteProductItem)
		}

		//	Coupon Management Routes
		couponRoutes := api.Group("/coupons")
		{
			couponRoutes.GET("/", productHandler.ViewAllCoupons)
			couponRoutes.GET("/:coupon_id", productHandler.ViewCouponByID)
			couponRoutes.POST("/", productHandler.CreateCoupon)
			couponRoutes.PUT("/", productHandler.UpdateCoupon)
			couponRoutes.DELETE("/:coupon_id", productHandler.DeleteCoupon)
		}

		order := api.Group("/orders")
		{
			order.PUT("/", orderHandler.UpdateOrder)
		}
	}
}
