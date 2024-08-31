package address_controller

import (
	"delivery/database"
	address_model "delivery/modules/address/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAddress(c *gin.Context){
	var address address_model.Address;

	// check the validate
	if err := c.ShouldBindJSON(&address); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "message": "input field is not valid"});
		return;
	}

	if err := database.DB.Create(&address).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server error",
		"status":false,
	})
	}

	c.JSON(http.StatusCreated, gin.H{"status":true, "message":"Address created"})

}

func UpdateAddress(c *gin.Context){
   var input struct{
		User_id uint `json:"user_id"`
		District string `json:"district"`
	    Division string `json:"division"`
	    Road string `json:"road"`
	    House string `json:"house"`
	    Description string `json:"description"`
	    Status bool `json:"status" gorm:"default:true"`
	}

	id := c.Param("id");

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "message":"Input invalid"})
	}

	result := database.DB.Model(&address_model.Address{}).Where("id = ?", id).Updates(input);

	if result.Error != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message":"Order not updated"})
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status":false, "message":"Address not found by Id"})
	}

	c.JSON(http.StatusOK, gin.H{"status":false, "message":"Address updated", "order":result})
}

func GetAllAddress(c *gin.Context){
	string_id :=  c.Param("id");
	id, err := strconv.Atoi(string_id);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "message" : "invalid id"});
		return;
	}
	var orders []address_model.Address;
	result := database.DB.Where("user_id = ?", id).Find(&orders)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false});
		return;
	}
	c.JSON(http.StatusOK, gin.H{"status":true, "message":"Order list by user", "list":result});
}

func GetSingleOrder(c *gin.Context){
	string_id :=  c.Param("id");
	id, err := strconv.Atoi(string_id);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false,"message":"Invalid id"});
		return;
	}
	var orders []address_model.Address;
	result := database.DB.First(&orders,id);

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false});
		return;
	}
	c.JSON(http.StatusOK, gin.H{"status":true, "message":"Order list by user", "list":result});
}

func DeleteSingleAddress(c *gin.Context){
	string_id :=  c.Param("id");
	id, err := strconv.Atoi(string_id);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false,"message":"Invalid id"});
		return;
	}
	result := database.DB.Delete(address_model.Address{}, id);

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":false, "message":"Server not working"});
		return;
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message":"order not found by id "});
		return;
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message":"order deleted"});
}