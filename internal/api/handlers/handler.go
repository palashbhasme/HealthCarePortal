package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/palashbhasme/healthcare-portal/config"
	"github.com/palashbhasme/healthcare-portal/internal/api/dto/request"
	"github.com/palashbhasme/healthcare-portal/internal/api/middleware"
	"github.com/palashbhasme/healthcare-portal/internal/domain/models"
	"github.com/palashbhasme/healthcare-portal/internal/services/patient_service"
	"github.com/palashbhasme/healthcare-portal/internal/services/user_service"
	"go.uber.org/zap"
)

type Handler struct {
	userService    *user_service.UserService
	patientService *patient_service.PatientService
	logger         *zap.Logger
	auth           *config.AuthConfig
}

func NewHandler(router *gin.Engine, logger *zap.Logger,
	userService *user_service.UserService,
	patientService *patient_service.PatientService,
	auth *config.AuthConfig) {

	handler := &Handler{
		userService:    userService,
		patientService: patientService,
		logger:         logger,
		auth:           auth,
	}

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/signup", handler.CreateUser)
			user.POST("/login", handler.LoginUser)

			user.PUT("/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(), handler.UpdateUserById)
			user.DELETE("/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(), handler.DeleteUserById)
		}

		patient := api.Group("/patient")
		{
			// Patient routes
			patient.POST("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(), handler.CreatePatient)
			patient.PUT("/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(), handler.UpdatePatientById)
			patient.GET("/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(), handler.GetPatientById)
		}
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		h.logger.Error("Failed to bind user request", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	userResponse, err := h.userService.CreateUser(&userRequest)
	if err != nil {
		if err.Error() == "user already exists" {
			h.logger.Error("User already exists", zap.String("username", userRequest.Username))
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		} else {
			h.logger.Error("Failed to create user", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	}

	h.logger.Info("User created successfully", zap.String("username", userRequest.Username))
	c.JSON(http.StatusCreated, gin.H{"user": userResponse})
}

func (h *Handler) LoginUser(c *gin.Context) {

	var userRequest request.UserLoginRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		h.logger.Error("Failed to bind user request", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	userResponse, err := h.userService.LoginUser(&userRequest)
	if err != nil {
		if err.Error() == "invalid credentials" {
			h.logger.Error("Invalid credentials", zap.String("username", userRequest.Username))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		h.logger.Error("Failed to login user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login user"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour).Unix()

	claims := models.UserClaims{
		Role: userResponse.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,

			Subject: fmt.Sprint(userResponse.ID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(h.auth.SecretKey))
	if err != nil {
		h.logger.Error("Failed to sign token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign token"})
		return
	}

	h.logger.Info("User logged in successfully", zap.String("username", userRequest.Username))
	c.JSON(http.StatusOK, gin.H{"user": userResponse, "token": tokenString})
}

func (h *Handler) UpdateUserById(c *gin.Context) {
	idParam := c.Param("id")

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		h.logger.Error("Failed to bind user update request", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	userResponse, err := h.userService.UpdateUserById(idParam, updates)
	if err != nil {
		if err.Error() == "invalid user ID" {
			h.logger.Error("Invalid user ID", zap.String("userID", idParam))
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		h.logger.Error("Failed to update user", zap.Error(err))
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	h.logger.Info("User updated successfully", zap.String("userID", idParam))
	c.JSON(200, gin.H{"user": userResponse})
}

func (h *Handler) DeleteUserById(c *gin.Context) {
	idParam := c.Param("id")

	err := h.userService.DeleteUserById(idParam)
	if err != nil {
		if err.Error() == "invalid user ID" {
			h.logger.Error("Invalid user ID", zap.String("userID", idParam))
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		h.logger.Error("Failed to delete user", zap.Error(err))
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	h.logger.Info("User deleted successfully", zap.String("userID", idParam))
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func (h *Handler) UpdatePatientById(c *gin.Context) {
	idParam := c.Param("id")
	role := c.GetString("role")

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		h.logger.Error("Failed to bind patient update request", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	patientResponse, err := h.patientService.UpdatePatientById(idParam, updates, role)
	if err != nil {
		if err.Error() == "invalid patient ID" {
			h.logger.Error("Invalid patient ID", zap.String("patientID", idParam))
			c.JSON(400, gin.H{"error": "Invalid patient ID"})
			return
		}
		if err.Error() == "permission denied" {
			h.logger.Warn("Permission denied to update patient", zap.String("role", role))
			c.JSON(403, gin.H{"error": "You do not have permission to update this patient"})
			return
		}
		h.logger.Error("Failed to update patient", zap.Error(err))
		c.JSON(500, gin.H{"error": "Failed to update patient"})
		return
	}

	h.logger.Info("Patient updated successfully", zap.String("patientID", idParam))
	c.JSON(200, gin.H{"patient": patientResponse})
}

func (h *Handler) GetPatientById(c *gin.Context) {
	idParam := c.Param("id")
	role := c.GetString("role")

	patientResponse, err := h.patientService.GetPatientById(idParam, role)
	if err != nil {
		if err.Error() == "invalid patient ID" {
			h.logger.Error("Invalid patient ID", zap.String("patientID", idParam))
			c.JSON(400, gin.H{"error": "Invalid patient ID"})
			return
		}
		if err.Error() == "permission denied" {
			h.logger.Warn("Permission denied to view patient", zap.String("role", role))
			c.JSON(403, gin.H{"error": "You do not have permission to view this patient"})
			return
		}
		h.logger.Error("Failed to get patient", zap.Error(err))
		c.JSON(500, gin.H{"error": "Failed to get patient"})
		return
	}

	h.logger.Info("Patient retrieved successfully", zap.String("patientID", idParam))
	c.JSON(200, gin.H{"patient": patientResponse})
}

func (h *Handler) CreatePatient(c *gin.Context) {
	var patientRequest request.PatientRequest
	role := c.GetString("role")

	if err := c.ShouldBindJSON(&patientRequest); err != nil {
		h.logger.Error("Failed to bind patient request", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	patientResponse, err := h.patientService.CreatePatient(&patientRequest, role)
	if err != nil {
		if err.Error() == "permission denied" {
			h.logger.Warn("Permission denied to create patient", zap.String("role", role))
			c.JSON(403, gin.H{"error": "You do not have permission to create a patient"})
			return
		}
		h.logger.Error("Failed to create patient", zap.Error(err))
		c.JSON(500, gin.H{"error": "Failed to create patient"})
		return
	}

	h.logger.Info("Patient created successfully", zap.String("patientID", fmt.Sprint(patientResponse.ID)))
	c.JSON(201, gin.H{"patient": patientResponse})
}
