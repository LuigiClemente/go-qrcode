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

func (h *Handler) ValidateQRCode(c *gin.Context) {
	// Get the token from the request query parameter
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	// Validate the QR code
	valid, err := h.service.ValidateQRCode(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"valid": valid})
}

func (h *Handler) InvalidateQRCode(c *gin.Context) {
	// Get the token from the URL path parameter
	token := c.Param("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	// Invalidate the QR code
	err := h.service.InvalidateQRCode(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "QR code invalidated successfully"})
}

func (h *Handler) DeleteQRCode(c *gin.Context) {
	// Get the token from the URL path parameter
	token := c.Param("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	// Delete the QR code
	err := h.service.DeleteQRCode(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "QR code deleted successfully"})
}
