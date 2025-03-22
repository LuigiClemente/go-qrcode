package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"qr-code-backend/models"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

// GenerateQRCode creates a QR code, converts it to a Data URL, and stores it in the database.
func (s *Service) GenerateQRCode() (*models.QRCode, error) {
	token := uuid.New().String()
	url := fmt.Sprintf("https://matomo.gutricious.com/validate/%s", token)

	// Generate QR code
	qr, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR code: %w", err)
	}

	// Encode QR code as PNG in a buffer
	var buf bytes.Buffer
	if err := qr.Write(256, &buf); err != nil {
		return nil, fmt.Errorf("failed to write QR code to buffer: %w", err)
	}

	// Convert buffer to Base64 and format as Data URL
	base64Image := "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())

	// Insert into database and return the created QRCode object
	qrCode := models.QRCode{
		Token: token,
		Url:   url,
		Image: base64Image, // Storing as Data URL
		Valid: true,
	}

	if err := s.db.Create(&qrCode).Error; err != nil {
		return nil, fmt.Errorf("failed to save QR code: %w", err)
	}

	return &qrCode, nil
}

// ValidateQRCode checks if the provided QR token is valid.
func (s *Service) ValidateQRCode(token string) (bool, error) {
	var valid bool
	if err := s.db.Model(&models.QRCode{}).Select("valid").Where("token = ?", token).Scan(&valid).Error; err != nil {
		return false, fmt.Errorf("failed to validate QR code: %w", err)
	}
	return valid, nil
}

// InvalidateQRCode marks a QR token as invalid.
func (s *Service) InvalidateQRCode(token string) error {
	return s.updateQRCodeValidity(token, false)
}

// DeleteQRCode permanently removes a QR token from the database.
func (s *Service) DeleteQRCode(token string) error {
	if err := s.db.Where("token = ?", token).Delete(&models.QRCode{}).Error; err != nil {
		return fmt.Errorf("failed to delete QR code: %w", err)
	}
	return nil
}

// updateQRCodeValidity updates the validity status of a QR code.
func (s *Service) updateQRCodeValidity(token string, valid bool) error {
	result := s.db.Model(&models.QRCode{}).Where("token = ?", token).Update("valid", valid)
	if result.Error != nil {
		return fmt.Errorf("failed to update QR code validity: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("QR code not found")
	}
	return nil
}
