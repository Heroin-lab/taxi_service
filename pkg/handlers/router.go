package handlers

import (
	"github.com/Heroin-lab/taxi_service.git/pkg/services"
	"github.com/gin-gonic/gin"
)

type Router struct {
	services *services.Services
}

func NewRouter(services *services.Services) *Router {
	return &Router{services: services}
}

func (r *Router) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", r.signUp)
		auth.POST("/sign-in", r.signIn)
	}

	return router
}
