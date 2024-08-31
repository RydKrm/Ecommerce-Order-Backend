package return_controller

import (
	"delivery/database"
	return_model "delivery/modules/return/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateReturn(c *gin.Context){
	var input return_model.Return

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(400, gin.H{"status": false, "message":"Invalid input field"})
		return
	}

	if err := database.DB.Create(&input); err != nil{
		c.JSON(400, gin.H{"status":false, "message":"Not created"})
		return;
	}

	c.JSON(201, gin.H{"status": true, "message":"Successfully return created"})
}

func UpdateReturn(c *gin.Context){
   var input struct{
			Order_id uint `json:"order_id"`
	User_id uint `json:"user_id"`
	Order_item_id uint `json:"order_item_id"`
	Return_date time.Time `json:"return_date"`
	Reason string `json:"reason"`
	}

	id := c.Param("id");

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "message":"Input invalid"})
		return;
	}

	result := database.DB.Model(&return_model.Return{}).Where("id = ?", id).Updates(input);

	if result.Error != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message":"Not updated"})
		return;
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status":false, "message":"Return not found by Id"})
		return;
	}

	c.JSON(http.StatusOK, gin.H{"status":true, "message":"Return updated", "order":result})
}

func GetAllReturn(c *gin.Context){
	string_id :=  c.Param("id");
	id, err := strconv.Atoi(string_id);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "message" : "invalid id"});
		return;
	}
	var orders []return_model.Return;
	result := database.DB.Where("user_id = ?", id).Find(&orders)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false});
		return;
	}
	c.JSON(http.StatusOK, gin.H{"status":true, "message":"Order list by user", "list":result});
}

func GetSingleReturn(c *gin.Context){
	string_id :=  c.Param("id");
	id, err := strconv.Atoi(string_id);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false,"message":"Invalid id"});
		return;
	}
	var orders []return_model.Return;
	result := database.DB.First(&orders,id);

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false});
		return;
	}
	c.JSON(http.StatusOK, gin.H{"status":true, "message":"Return list by user", "list":result});
}

func DeleteSingleReturn(c *gin.Context){
	string_id :=  c.Param("id");
	id, err := strconv.Atoi(string_id);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false,"message":"Invalid id"});
		return;
	}
	result := database.DB.Delete(return_model.Return{}, id);

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":false, "message":"Server not working"});
		return;
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message":"Return not found by id "});
		return;
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message":"order deleted"});
}
