package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanaponkhanoon/sa-65-example/entity"
	"net/http"
)

// POST /Customer
func CreateCustomer(c *gin.Context) {
	var ccode entity.Ccode
	var cusdetail entity.Cusdetail
	var customer entity.Customer
	//ผลลัพทธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปล customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//9: ค้นหา ccode ด้วย id
	if tx := entity.DB().Where("id = ?", customer.CcodeID).First(&ccode); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ccode not found"})
		return
	}
	//10: ค้นหา cusdetail ด้วย id
	if tx := entity.DB().Where("id = ?", customer.Cusdetail).First(&cusdetail); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cusdetail not found"})
		return
	}
	//11: สร้าง Customer
	cm := entity.Customer{
		Ccode:  ccode,
		Cusdetail: cusdetail,
		Name:   customer.Name,
	}
	//12: บันทึก
	if err := entity.DB().Create(&cm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})
}

// GET /customer/:id
func GetCustomer(c *gin.Context) {
	var customer entity.Customer
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM customer WHERE id = ?", id).
		Preload("Ccode").
		Preload("Cusdetail").
		Find(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// GET /customer
func ListCustomer(c *gin.Context) {
	var customer []entity.Customer
	if err := entity.DB().Raw("SELECT * FROM customer").
		Preload("Ccode").
		Preload("Cusdetail").
		Find(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// DELETE /customer/:id
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM customer WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /customer
func UpdateCustomer(c *gin.Context) {
	var customer entity.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", customer.ID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}
	if err := entity.DB().Save(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}
