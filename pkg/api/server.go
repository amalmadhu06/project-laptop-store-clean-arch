package http

import (
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/api/handler"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/api/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	otpHandler *handler.OtpHandler,
	productHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	orderHandler *handler.OrderHandler,
	paymentHandler *handler.PaymentHandler,
	wishlistHandler *handler.WishlistHandler,
) *ServerHTTP {

	engine := gin.New()
	engine.Use(gin.Logger())

	// swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// set up routes
	routes.UserRoutes(engine.Group("/"), userHandler, productHandler, cartHandler, orderHandler, otpHandler, paymentHandler, wishlistHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, userHandler, productHandler, orderHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	//sh.engine.LoadHTMLGlob("template/*.html")
	err := sh.engine.Run(":3000")
	if err != nil {
		return
	}
}
