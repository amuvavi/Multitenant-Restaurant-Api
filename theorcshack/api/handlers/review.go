package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"theorcshack/db/models"

	"theorcshack/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewAndRatingInput struct {
	Content string  `json:"content" binding:"required"`
	Rating  float64 `json:"rating" binding:"required,gt=0,lt=6"`
}

func getTenantID(c *gin.Context) (uuid.UUID, error) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		return uuid.Nil, fmt.Errorf("X-Tenant-ID header is required")
	}

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid X-Tenant-ID header")
	}

	return tenantUUID, nil
}

func getUserID(c *gin.Context) (uuid.UUID, error) {
	userID, exists := c.Get("userID")
	if !exists || userID == nil {
		return uuid.Nil, fmt.Errorf("userID not found in context")
	}

	userUUID, ok := userID.(uuid.UUID)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid userID")
	}

	return userUUID, nil
}

func ReviewAndRateDish(c *gin.Context) {
	var input ReviewAndRatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	dishID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid UUID for dish")
		return
	}

	tenantUUID, err := getTenantID(c)
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userUUID, err := getUserID(c)
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Rating < 1 || input.Rating > 5 {
		helpers.RespondWithError(c, http.StatusBadRequest, "Rating must be between 1 and 5")
		return
	}

	var review models.Review
	result := models.DB.Where("user_id = ? AND dish_id = ? AND tenant_id = ?", userUUID, dishID, tenantUUID).First(&review)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Database error during review lookup")
		return
	}

	review.Content = input.Content
	review.Rating = input.Rating
	review.SentimentScore = analyzeSentiment(input.Content)

	if result.RowsAffected == 0 {
		review.TenantID = tenantUUID
		review.UserID = userUUID
		review.DishID = dishID
		if err := models.DB.Create(&review).Error; err != nil {
			helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusCreated, review)
	} else {
		if err := models.DB.Save(&review).Error; err != nil {
			helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, review)
	}
}

func GetReview(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid UUID")
		return
	}

	tenantUUID, err := getTenantID(c)
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	var review models.Review
	if err := models.DB.Where("id = ? AND tenant_id = ?", id, tenantUUID).First(&review).Error; err != nil {
		helpers.RespondWithError(c, http.StatusNotFound, "Review not found")
		return
	}

	c.JSON(http.StatusOK, review)
}

func UpdateReview(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid UUID")
		return
	}

	tenantUUID, err := getTenantID(c)
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	var review models.Review
	if err := models.DB.Where("id = ? AND tenant_id = ?", id, tenantUUID).First(&review).Error; err != nil {
		helpers.RespondWithError(c, http.StatusNotFound, "Review not found")
		return
	}

	var updateInput ReviewAndRatingInput
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	review.Content = updateInput.Content
	review.Rating = updateInput.Rating
	review.SentimentScore = analyzeSentiment(updateInput.Content)

	if err := models.DB.Save(&review).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, review)
}

func DeleteReview(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid UUID")
		return
	}

	tenantUUID, err := getTenantID(c)
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	var review models.Review
	if err := models.DB.Where("id = ? AND tenant_id = ?", id, tenantUUID).First(&review).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.RespondWithError(c, http.StatusNotFound, "Review not found")
			return
		}
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := models.DB.Delete(&review).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}
