package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *accountHandler) DeleteToken(c *gin.Context) {
	c.SetCookie("auth_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Token deleted", "data": nil})
}
