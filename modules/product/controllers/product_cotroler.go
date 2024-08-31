package product_controller

import (
	product_model "delivery/modules/product/models"
	product_service "delivery/modules/product/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context){
	var product product_model.Product
	if err := c.ShouldBindJSON(&product); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err:= product_service.CreateProduct(&product); err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return;
	}

	c.JSON(http.StatusCreated, product)
}

func GetAllProducts(c *gin.Context){
	products, err := product_service.GetAllProducts();
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message":"Product found", "product_list":products})
}

func GetProductByID(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    product, err := product_service.GetProductByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
    var product product_model.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := product_service.UpdateProduct(&product); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    if err := product_service.DeleteProduct(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusNoContent, gin.H{})
}

