package controller

import (
	"github.com/thanaponkhanoon/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Ccodes
func CreateCcode(c *gin.Context) {
	var ccode entity.Ccode
	if err := c.ShouldBindJSON(&ccode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&ccode).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": ccode})
}

// GET /pcode/:id
func GetCcode(c *gin.Context) {
	var ccode entity.Ccode
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&ccode); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ccode not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ccode})
}

// GET /Ccode
func ListCcode(c *gin.Context) {
	var ccode []entity.Ccode
	if err := entity.DB().Raw("SELECT * FROM ccode").Scan(&ccode).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ccode})
}

// DELETE /Ccode/:id
func DeleteCcode(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM ccode WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ccode not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Ccode
func UpdateCcode(c *gin.Context) {
	var ccode entity.Ccode
	if err := c.ShouldBindJSON(&ccode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", ccode.ID).First(&ccode); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pcode not found"})
		return
	}

	if err := entity.DB().Save(&ccode).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ccode})
}