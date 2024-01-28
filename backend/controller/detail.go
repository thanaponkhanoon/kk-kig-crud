package controller

import (
	"github.com/thanaponkhanoon/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Detail
func CreateDetail(c *gin.Context) {
	var detail entity.Detail
	if err := c.ShouldBindJSON(&detail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&detail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": detail})
}

// GET /detail/:id
func GetDetail(c *gin.Context) {
	var detail entity.Detail
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&detail); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "detail not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": detail})
}

// GET /detail
func ListDetail(c *gin.Context) {
	var details []entity.Detail
	if err := entity.DB().Raw("SELECT * FROM details").Scan(&details).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": details})
}

// DELETE /detail/:id
func DeleteDetail(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM detail WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "detail not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /detail
func UpdateDetail(c *gin.Context) {
	var detail entity.Detail
	if err := c.ShouldBindJSON(&detail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", detail.ID).First(&detail); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "detail not found"})
		return
	}

	if err := entity.DB().Save(&detail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": detail})
}