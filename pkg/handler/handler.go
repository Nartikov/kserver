package handler

import "github.com/gin-gonic/gin"

type Handler struct{

}

func (h *Handler) InitRoutes() *gin.Engine{
	router:=gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}
	api:= router.Group("/api")
		listquiz:=api.Group("/listquizes")
	{	
		listquiz.GET("/:id")
		listquiz.GET("/:id/*id")


	}
}