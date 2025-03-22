package services

import (
	"errors"
	"qr-code-backend/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SignUp(user models.User) (*models.User, error) {
	// password already hashed in the handler
	userSQL := `INSERT INTO users (first_name, last_name, email, password) 
                 VALUES (?, ?, ?, ?) RETURNING id`
	var userID uint
	if err := s.db.DB.Raw(userSQL, user.FirstName, user.LastName, user.Email, user.Password).Scan(&userID).Error; err != nil {
		return nil, err
	}
	user.ID = userID
	return &user, nil
}

func (s *Service) Login(email, password string) (*models.User, error) {
	var user models.User
	userSQL := `SELECT id, first_name, last_name, email, password FROM users WHERE email = ?`
	if err := s.db.DB.Raw(userSQL, email).Scan(&user).Error; err != nil {
		return nil, err
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return &user, nil
}

func (s *Service) ForgotPassword(email string) error {
	var user models.User
	userSQL := `SELECT id, email FROM users WHERE email = ?`
	if err := s.db.DB.Raw(userSQL, email).Scan(&user).Error; err != nil {
		return errors.New("email not found")
	}

	// Generate a reset token (you can use JWT or a UUID)
	resetToken := uuid.New().String()

	// Save the reset token in the database (or use Redis for temporary storage)
	tokenSQL := `INSERT INTO password_resets (user_id, token) VALUES (?, ?)`
	if err := s.db.DB.Exec(tokenSQL, user.ID, resetToken).Error; err != nil {
		return err
	}

	// TODO: Send resetToken via email to the user

	return nil
}
