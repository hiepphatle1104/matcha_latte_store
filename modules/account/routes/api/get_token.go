package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hiepphatle1104/matcha_latte_store/modules/account/domain/dto"
	"net/http"
)

func (h *accountHandler) GetToken(c *gin.Context) {
	var data dto.SignInAccount
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	tokenString, err := h.service.LoginAccount(c.Request.Context(), &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("auth_token", tokenString, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"data": tokenString, "message": "Token retrieved"})
}
