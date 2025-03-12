package services

import (
	"fmt"
	"qr-code-backend/database"
)

type Service struct {
	db *database.Database
}

func NewService() *Service {
	return &Service{
		db: database.GetDatabaseInstance(),
	}
}

// CustomError wraps the original error with a message
type CustomError struct {
	Message string
	Err     error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}
