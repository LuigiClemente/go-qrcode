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
	var qrCode models.QRCode
	if err := s.db.Raw(qrSQL, token, "", imagePath, true).Scan(&qrCode.ID).Error; err != nil {
		return nil, err
	}

	// Retrieve the stored QRCode object from the database
	if err := s.db.Raw(`SELECT id, token, url, image, valid, created_at FROM qr_codes WHERE id = ?`, qrCode.ID).
		Scan(&qrCode).Error; err != nil {
		return nil, err
	}

	return &qrCode, nil
}

// Checks if a provided QR token is valid
func (s *Service) ValidateQRCode(token string) (bool, error) {
	// Query the database to check if the QR code exists and is valid
	var valid bool
	err := s.db.Raw(`SELECT valid FROM qr_codes WHERE token = ?`, token).Scan(&valid).Error
	if err != nil {
		return false, fmt.Errorf("failed to validate QR code: %w", err)
	}

	return valid, nil
}

// Marks a QR token as invalid
func (s *Service) InvalidateQRCode(token string) error {
	// Update the database to mark the QR code as invalid
	err := s.db.Exec(`UPDATE qr_codes SET valid = false WHERE token = ?`, token).Error
	if err != nil {
		return fmt.Errorf("failed to invalidate QR code: %w", err)
	}

	return nil
}
