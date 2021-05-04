package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createQuiz(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
	})
}

func (h *Handler) getAllQuizes(c *gin.Context) {

}

func (h *Handler) getQuizById(c *gin.Context) {

}

func (h *Handler) deleteQuiz(c *gin.Context) {

}
