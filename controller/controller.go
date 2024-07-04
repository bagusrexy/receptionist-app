package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bagusrexy/test-dataon/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GuestController struct{}

func (h *GuestController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "SUCCESS"})
}

type CreateGuestInput struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	AccessFloor string `json:"access_floor" binding:"required"`
}

func (h *GuestController) RegisterGuest(c *gin.Context) {
	var input CreateGuestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Failed to register",
			"error":   err.Error(),
		})
		return
	}

	guest := models.Guest{
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		CheckIn:     time.Now(),
		AccessFloor: input.AccessFloor,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&guest)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Successfully registered",
		"Data":    guest})
}

func (h *GuestController) CheckOutGuest(c *gin.Context) {
	var guest models.Guest
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&guest).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Guest not found!"})
		return
	}

	guest.CheckOut = time.Now()
	db.Save(&guest)

	c.JSON(http.StatusOK, guest)
}

func (gc *GuestController) UploadPhoto(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Photo is required"})
		return
	}

	path := fmt.Sprintf("uploads/%s", file.Filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save photo"})
		return
	}

	if err := db.Model(&models.Guest{}).Where("id = ?", id).Update("photo", path).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update guest photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo uploaded successfully", "path": path})
}
