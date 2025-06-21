package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/domain/dto"
	"net/http"
)

func (h *accountHandler) CreateAccount(c *gin.Context) {
	var data dto.SignUpAccount
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.RegisterAccount(c.Request.Context(), &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": true, "message": "Account created"})
}
