package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"lecture10/internal/user/entity"
	"lecture10/internal/user/repository"
	"log"
	"net/http"
	"strconv"
	"time"
)

var userRepository repository.UserRepositoryInterface

func RegisterHandler(c *gin.Context, userRepository repository.UserRepositoryInterface, redisClient *redis.Client) {
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

	userID, err := userRepository.Register(c, user, hashedPassword)
	if err != nil {
		log.Printf("Error registering user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	fmt.Printf("User registered successfully with ID: %d\n", userID)

	registeredUser, err := userRepository.GetById(c, userID)
	if err != nil {
		log.Printf("Error fetching registered user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch registered user"})
		return
	}

	if registeredUser == nil {
		log.Printf("Registered user not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Registered user not found"})
		return
	}

	SaveUserToRedis(redisClient, *registeredUser)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func SaveUserToRedis(redisClient *redis.Client, user entity.User) {
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error marshalling user: %v\n", err)
		return
	}

	ctx := context.Background()
	key := "user:" + strconv.Itoa(user.ID)

	_, err = redisClient.Set(ctx, key, userJSON, 5*time.Minute).Result()
	if err != nil {
		log.Printf("Error storing user in Redis: %v\n", err)
	}
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
func GetUserByIDHandler(c *gin.Context, userRepository repository.UserRepositoryInterface, redisClient *redis.Client) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	ctx := context.Background()
	key := "user:" + userIDStr
	userJSON, err := redisClient.Get(ctx, key).Result()

	if err == redis.Nil {
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

		userJSON, err := json.Marshal(user)
		if err != nil {
			log.Printf("Error marshalling user: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal user"})
			return
		}

		_, err = redisClient.Set(ctx, key, userJSON, 5*time.Minute).Result()
		if err != nil {
			log.Printf("Error storing user in Redis: %v\n", err)
		}
	} else if err != nil {
		log.Printf("Error fetching user from Redis: %v\n", err)
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(userJSON))
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
