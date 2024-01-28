package controller

import (
	"github.com/thanaponkhanoon/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Pcodes
func CreatePcode(c *gin.Context) {
	var pcode entity.Pcode
	if err := c.ShouldBindJSON(&pcode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&pcode).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": pcode})
}

// GET /pcode/:id
func GetPcode(c *gin.Context) {
	var pcode entity.Pcode
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&pcode); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pcode not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pcode})
}

// GET /Product
func ListPcode(c *gin.Context) {
	var pcodes []entity.Pcode
	if err := entity.DB().Raw("SELECT * FROM pcodes").Scan(&pcodes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pcodes})
}

// DELETE /categories/:id
func DeletePcode(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM pcode WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pcode not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Pcode
func UpdatePcode(c *gin.Context) {
	var pcode entity.Pcode
	if err := c.ShouldBindJSON(&pcode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", pcode.ID).First(&pcode); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pcode not found"})
		return
	}

	if err := entity.DB().Save(&pcode).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pcode})
}