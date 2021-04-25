package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nartikov/kserver/pkg/models"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context){
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

	fmt.Println("Create new user")
}

func (h *Handler) signIn(c *gin.Context){
	fmt.Println("start signIn")
}
