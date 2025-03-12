package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GenerateQRCode(c *gin.Context) {
	qrCode, err := h.service.GenerateQRCode()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, qrCode)
}
