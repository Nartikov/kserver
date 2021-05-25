package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nartikov/kserver/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		quiz := api.Group("/quiz")
		{
			quiz.GET("/", h.getAllQuizes)
			quiz.GET("/:id", h.getQuizById)
			quiz.POST("/", h.createQuiz)
			quiz.DELETE("/:id", h.deleteQuiz)
			question := quiz.Group("/question")
			{
				question.GET("/", h.getAllQuestion)
				question.GET("/:id", h.getQuestionById)
				question.POST("/", h.createQuestion)
				question.DELETE("/:id", h.deleteQuestion)
				answer := question.Group("answer")
				{
					answer.GET("/", h.getAnswersToQuestion)
					answer.POST("/", h.createAnswersToQuestion)
					answer.DELETE("/", h.deleteAnswersToQuestion)
				}
			}
		}
	}
	return router
}
