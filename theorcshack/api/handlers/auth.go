package handlers

import (
	"log"
	"net/http"
	"os"
	"theorcshack/db/models"
	"time"

	"theorcshack/helpers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func getJWTKey() []byte {
	key := os.Getenv("JWT_SECRET_KEY")
	log.Printf("JWT Secret Key (Generate): %s", key)
	return []byte(key)
}

type Claims struct {
	UserID   uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	TenantID uuid.UUID `json:"tenant_id"`
	jwt.StandardClaims
}

func Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	tenant := models.Tenant{
		Name: input.Email + "'s Tenant",
	}
	if err := models.DB.Create(&tenant).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Creating a new user associated with the tenant
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		TenantID: tenant.ID,
	}
	if err := models.DB.Create(&user).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"name":     user.Name,
		"email":    user.Email,
		"tenantID": tenant.ID,
	})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		helpers.RespondWithError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		helpers.RespondWithError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   user.ID,
		Email:    user.Email,
		TenantID: user.TenantID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTKey())
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Generated Token: %s", tokenString)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
