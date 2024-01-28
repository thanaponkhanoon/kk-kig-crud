package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanaponkhanoon/sa-65-example/entity"
	"net/http"
)

// POST /Product
func CreateProduct(c *gin.Context) {
	var pcode entity.Pcode
	var detail entity.Detail
	var product entity.Product
	//ผลลัพทธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปล product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//9: ค้นหา pcode ด้วย id
	if tx := entity.DB().Where("id = ?", product.PcodeID).First(&pcode); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pcode not found"})
		return
	}
	//10: ค้นหา detail ด้วย id
	if tx := entity.DB().Where("id = ?", product.DetailID).First(&detail); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "detail not found"})
		return
	}
	//11: สร้าง Product
	pd := entity.Product{
		Pcode:  pcode,
		Detail: detail,
		Time:   product.Time,
		Name:   product.Name,
		Price:  product.Price,
		Amount: product.Amount,
	}
	//12: บันทึก
	if err := entity.DB().Create(&pd).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pd})
}

// GET /product/:id
func GetProduct(c *gin.Context) {
	var product entity.Product
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM product WHERE id = ?", id).
		Preload("Pcode").
		Preload("Detail").
		Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GET /product
func ListProducts(c *gin.Context) {
	var products []entity.Product
	if err := entity.DB().Raw("SELECT * FROM products").
		Preload("Pcode").
		Preload("Detail").
		Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// DELETE /product/:id
func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    if tx := entity.DB().Exec("DELETE FROM product WHERE id = ?", id); tx.RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /product
func UpdateProduct(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", product.ID).First(&product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}
	if err := entity.DB().Save(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}
