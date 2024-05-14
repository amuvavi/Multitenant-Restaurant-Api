package handlers

import (
	"net/http"
	"theorcshack/db/models"

	"theorcshack/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTenant(c *gin.Context) {
	var tenant models.Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := models.DB.Create(&tenant).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, tenant)
}

func ListTenants(c *gin.Context) {
	var tenants []models.Tenant
	if err := models.DB.Find(&tenants).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tenants)
}

func GetTenant(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid UUID")
		return
	}

	var tenant models.Tenant
	if err := models.DB.First(&tenant, "id = ?", id).Error; err != nil {
		helpers.RespondWithError(c, http.StatusNotFound, "Tenant not found")
		return
	}
	c.JSON(http.StatusOK, tenant)
}

func UpdateTenant(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid UUID")
		return
	}

	var tenant models.Tenant
	if err := models.DB.First(&tenant, "id = ?", id).Error; err != nil {
		helpers.RespondWithError(c, http.StatusNotFound, "Tenant not found")
		return
	}

	if err := c.ShouldBindJSON(&tenant); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := models.DB.Save(&tenant).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tenant)
}

func DeleteTenant(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid UUID")
		return
	}

	var tenant models.Tenant
	if err := models.DB.First(&tenant, "id = ?", id).Error; err != nil {
		helpers.RespondWithError(c, http.StatusNotFound, "Tenant not found")
		return
	}
	if err := models.DB.Delete(&tenant).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}
