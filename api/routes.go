package api

import (
	"qr-code-backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	// Register routes
	qrCode := router.Group("api/qr-code")
	{
		qrCodeHandler := handlers.NewHandler()
		qrCode.POST("/generate", qrCodeHandler.GenerateQRCode)
		qrCode.POST("/validate", qrCodeHandler.ValidateQRCode)

	}

	return router
}
