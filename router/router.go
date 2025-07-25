package router

import (
	ginHandlers "PaymentSystem/handlers/gin"
	"github.com/gin-gonic/gin"
)

type Router interface {
	SetupRoutes()
	AddMiddleware(middleware interface{})
	GetHandler() interface{}
}

type GinRouter struct {
	engine *gin.Engine
}

func NewGinRouter() *GinRouter {
	return &GinRouter{engine: gin.Default()}
}

func (r *GinRouter) SetupRoutes() {

	api := r.engine.Group("/api")
	{
		api.GET("/transactions", ginHandlers.GetLast)
		api.GET("/wallet/:wname/balance", ginHandlers.GetBalance)
		api.POST("/send", ginHandlers.PostSend)
	}

}

func (r *GinRouter) AddMiddleware(middleware interface{}) {
	if m, ok := middleware.(gin.HandlerFunc); ok {
		r.engine.Use(m)
	}
}

func (r *GinRouter) GetHandler() interface{} {
	return r.engine
}
