package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nartikov/kserver/pkg/models"
)

func (h *Handler) createQuiz(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok{
		newErrorResponse(c, http.StatusInternalServerError, "user id not fund")
	}
	var input models.Quiz
	if err:= c.BindJSON(&input); err!=nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//call metod services
}

func (h *Handler) getAllQuizes(c *gin.Context) {

}

func (h *Handler) getQuizById(c *gin.Context) {

}

func (h *Handler) deleteQuiz(c *gin.Context) {

}
