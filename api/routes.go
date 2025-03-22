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
	qrCode := router.Group("api/v1/qrcode-app")
	{
		qrCodeHandler := handlers.NewHandler()
		qrCode.POST("/generate", qrCodeHandler.GenerateQRCode)
		qrCode.GET("/validate/:token", qrCodeHandler.ValidateQRCode)
		qrCode.POST("/invalidate", qrCodeHandler.InvalidateQRCode)
		qrCode.DELETE("/qr/:token", qrCodeHandler.DeleteQRCode)
	}

	auth := router.Group("api/v1/auth")
	{
		auth.POST("/signup", handlers.SignupUser)
		auth.POST("/login", handlers.LoginUser)
		auth.POST("/forgot-password", handlers.ForgotPassword)
	}

	return router
}
