package handlers

// import (
// 	"database/sql"
// 	"net/http"

// 	// "central-server/config"
// 	// "central-server/models"

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/crypto/bcrypt"
// )

// // Signup User
// func SignupUser(c *gin.Context) {
// 	var user models.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	// Hash password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
// 		return
// 	}

// 	_, err = config.DB.Exec("INSERT INTO users (email, password, role) VALUES ($1, $2, $3)", user.Email, string(hashedPassword), "user")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
// }

// // Signin User
// func SigninUser(c *gin.Context) {
// 	var user models.User
// 	var storedPassword string

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	err := config.DB.QueryRow("SELECT id, password, role FROM users WHERE email = $1", user.Email).
// 		Scan(&user.ID, &storedPassword, &user.Role)
// 	if err == sql.ErrNoRows {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
// 		return
// 	} else if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
// 		return
// 	}

// 	// Check password
// 	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "role": user.Role})
// }

// // Update User Role
// func UpdateUserRole(c *gin.Context) {
// 	var request struct {
// 		Email string `json:"email"`
// 		Role  string `json:"role"`
// 	}

// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	_, err := config.DB.Exec("UPDATE users SET role = $1 WHERE email = $2", request.Role, request.Email)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User role updated successfully"})
// }
