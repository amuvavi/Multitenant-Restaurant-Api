package middlewares

import (
	"log"
	"net/http"
	"os"
	"theorcshack/api/handlers"
	"theorcshack/helpers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			helpers.RespondWithError(c, http.StatusUnauthorized, "Authorization header is required")
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]
		log.Printf("Received Token: %s", tokenString)
		log.Printf("JWT Secret Key (Validate): %s", os.Getenv("JWT_SECRET_KEY"))

		claims := &handlers.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			log.Printf("Token validation error: %v", err)
			helpers.RespondWithError(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		// Ensuring that the userID in claims is a valid UUID
		userID, err := uuid.Parse(claims.UserID.String())
		if err != nil {
			log.Printf("Invalid userID in token: %v", err)
			helpers.RespondWithError(c, http.StatusUnauthorized, "Invalid userID in token")
			c.Abort()
			return
		}

		// Setting the userID in the context
		c.Set("userID", userID)
		c.Next()
	}
}
