package handlers

import (
	"net/http"
	"theorcshack/db/models"
	"theorcshack/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateDish(c *gin.Context) {
	var dish models.Dish
	if err := c.ShouldBindJSON(&dish); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	tenantID, _ := c.Get("tenantID")
	dish.TenantID = tenantID.(uuid.UUID)

	if err := models.DB.Create(&dish).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, dish)
}

func ListDishes(c *gin.Context) {
	var dishes []models.Dish
	tenantID, _ := c.Get("tenantID")

	models.DB.Where("tenant_id = ?", tenantID).Find(&dishes)
	c.JSON(http.StatusOK, dishes)
}

func GetDish(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid UUID")
		return
	}
	tenantID, _ := c.Get("tenantID")

	var dish models.Dish
	if err := models.DB.Where("id = ? AND tenant_id = ?", id, tenantID).First(&dish).Error; err != nil {
		helpers.RespondWithError(c, http.StatusNotFound, "Dish not found")
		return
	}
	c.JSON(http.StatusOK, dish)
}

func UpdateDish(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid UUID")
		return
	}
	tenantID, _ := c.Get("tenantID")

	var dish models.Dish
	if err := models.DB.Where("id = ? AND tenant_id = ?", id, tenantID).First(&dish).Error; err != nil {
		helpers.RespondWithError(c, http.StatusNotFound, "Dish not found")
		return
	}

	if err := c.ShouldBindJSON(&dish); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := models.DB.Save(&dish).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, dish)
}

func DeleteDish(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid UUID")
		return
	}
	tenantID, _ := c.Get("tenantID")

	var dish models.Dish
	if err := models.DB.Where("id = ? AND tenant_id = ?", id, tenantID).First(&dish).Error; err != nil {
		helpers.RespondWithError(c, http.StatusNotFound, "Dish not found")
		return
	}
	if err := models.DB.Delete(&dish).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}

func SearchDishes(c *gin.Context) {
	query := c.Query("query")
	tenantID, _ := c.Get("tenantID")

	var dishes []models.Dish
	if err := models.DB.Where("tenant_id = ? AND (name LIKE ? OR description LIKE ?)", tenantID, "%"+query+"%", "%"+query+"%").Find(&dishes).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, dishes)
}
