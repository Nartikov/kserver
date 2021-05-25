package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nartikov/kserver/pkg/models"
)

func (h *Handler) createQuiz(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id not fund")
	}
	var input models.Quiz
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Quiz.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllQuizes(c *gin.Context) {

}

func (h *Handler) getQuizById(c *gin.Context) {

}

func (h *Handler) deleteQuiz(c *gin.Context) {

}
