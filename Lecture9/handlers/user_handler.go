package handlers

import (
	"Lecture9/internal/user/entity"
	"Lecture9/internal/user/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

var userRepository repository.UserRepositoryInterface

func RegisterHandler(c *gin.Context, userRepository repository.UserRepositoryInterface) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name, email, and password are required"})
		return
	}

	fmt.Printf("User data received: %+v\n", user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error generating hashed password: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	fmt.Printf("SQL Query: %s\n", query)

	userID, err := userRepository.Register(c, user, hashedPassword)
	if err != nil {
		log.Printf("Error registering user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	fmt.Printf("User registered successfully with ID: %d\n", userID)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginHandler(c *gin.Context, userRepository repository.UserRepositoryInterface) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	authenticatedUser, err := userRepository.Login(c, user)

	if err != nil {
		log.Printf("Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Log the error
		return
	}

	if authenticatedUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func GetAllUsersHandler(c *gin.Context, userRepository repository.UserRepositoryInterface) {
	users, err := userRepository.GetAll(c)
	if err != nil {
		log.Printf("Error fetching users: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
func GetUserByIDHandler(c *gin.Context, userRepository repository.UserRepositoryInterface) {
	userIDStr := c.Param("id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := userRepository.GetById(c, userID)
	if err != nil {
		log.Printf("Error fetching user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserHandler(c *gin.Context, userRepository repository.UserRepositoryInterface) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updatedUser entity.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		log.Printf("Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userRepository.Update(c, userID, updatedUser); err != nil {
		log.Printf("Error updating user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUserHandler(c *gin.Context, userRepository repository.UserRepositoryInterface) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := userRepository.Delete(c, userID); err != nil {
		log.Printf("Error deleting user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
