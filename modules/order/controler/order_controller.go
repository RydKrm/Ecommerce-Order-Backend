package order_crud_controller

import (
	"delivery/database"
	order_model "delivery/modules/order/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order order_model.Order;
	if err := c.ShouldBindJSON(&order); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message":"Input field required"});
		return;
	}

	if err := database.DB.Create(&order).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"status":false, "message":""});
		return;
	}

	c.JSON(http.StatusCreated, gin.H{"status":true,"message":"Order created","order":order });
}

func UpdateOrder(c *gin.Context){
	var input struct{
		User_id uint `json:"user_id"`
    	Address_id uint64 `json:"address_id"`
	    Order_date time.Time  `json:"order_date"`
	    Total_amount uint64 `json:"total_amount"`
	    Status string `json:"status"`
	}

	id := c.Param("id");

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message":"Input invalid", })
	}

	result := database.DB.Model(&order_model.Order{}).Where("id = ?",id).Updates(input);

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message":"Not updated"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status":false, "message":"Order not found by Id"})
	}

	c.JSON(http.StatusOK, gin.H{"status":false, "message":"Order updated", "order":result})
}

func GetAllOrder(c *gin.Context){
	string_id :=  c.Param("id");
	id, err := strconv.Atoi(string_id);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false});
		return;
	}
	var orders []order_model.Order;
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
	var orders []order_model.Order;
	result := database.DB.First(&orders,id);

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false});
		return;
	}
	c.JSON(http.StatusOK, gin.H{"status":true, "message":"Order list by user", "list":result});
}

func DeleteSingleOrder(c *gin.Context){
	string_id :=  c.Param("id");
	id, err := strconv.Atoi(string_id);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false,"message":"Invalid id"});
		return;
	}
	result := database.DB.Delete(order_model.Order{}, id);

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":false, "message":"Method not working"});
		return;
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message":"order not found by id "});
		return;
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message":"order deleted"});
}


