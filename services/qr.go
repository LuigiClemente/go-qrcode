package services

import (
	"fmt"
	"os"
	"path/filepath"
	"qr-code-backend/models"
	"time"

	"github.com/skip2/go-qrcode"
)

func (s *Service) GenerateQRCode() (*models.QRCode, error) {
	// Generate unique token
	token := fmt.Sprintf("qr-%d", time.Now().UnixNano())

	// Define image directory and file path
	imageDir := "./qr_images"
	if err := os.MkdirAll(imageDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	imagePath := filepath.Join(imageDir, token+".png")

	// Generate and save the QR code image
	err := qrcode.WriteFile(token, qrcode.Medium, 256, imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR code: %w", err)
	}

	// Insert into database
	qrSQL := `INSERT INTO qr_codes (token, url, image, valid) VALUES (?, ?, ?, ?) RETURNING id`
	var id int
	if err := s.db.Raw(qrSQL, token, "", imagePath, true).Scan(&id).Error; err != nil {
		return nil, err
	}

	// Return the stored QRCode object
	return &models.QRCode{
		ID:    id,
		Token: token,
		Image: imagePath,
		Valid: true,
	}, nil
}
