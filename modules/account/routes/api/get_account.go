package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *accountHandler) GetAccount(c *gin.Context) {
	accountID, exists := c.Get("account_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account ID is required"})
		return
	}

	account, err := h.service.GetAccount(c.Request.Context(), accountID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": account, "message": "Account retrieved"})
}
