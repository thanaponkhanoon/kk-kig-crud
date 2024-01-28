package controller

import (
	"github.com/thanaponkhanoon/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Cusdetail
func CreateCusdetail(c *gin.Context) {
	var cusdetail entity.Cusdetail
	if err := c.ShouldBindJSON(&cusdetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&cusdetail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": cusdetail})
}

// GET /cusdetail/:id
func GetCusdetail(c *gin.Context) {
	var cusdetail entity.Cusdetail
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&cusdetail); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cusdetail not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cusdetail})
}

// GET /cusdetail
func ListCusdetail(c *gin.Context) {
	var cusdetail []entity.Cusdetail
	if err := entity.DB().Raw("SELECT * FROM cusdetail").Scan(&cusdetail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cusdetail})
}

// DELETE /cusdetail/:id
func DeleteCusdetail(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM cusdetail WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cusdetail not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /cusdetail
func UpdateCusdetail(c *gin.Context) {
	var cusdetail entity.Cusdetail
	if err := c.ShouldBindJSON(&cusdetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", cusdetail.ID).First(&cusdetail); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cusdetail not found"})
		return
	}

	if err := entity.DB().Save(&cusdetail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cusdetail})
}