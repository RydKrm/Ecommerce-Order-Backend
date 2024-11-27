package order_crud_controller

import (
	"delivery/database"
	order_model "delivery/modules/order/model"
	"delivery/services/product_service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/*
order = {
   address_id: 123,
   user_id: 123,
   order_date: 123,
   total_amount: 123,
   status: "pending"
   products: [
	  {
		 product_id: 123,
		 shop_id: 123,
		 quantity: 1,
		 price: 123
	  },
	  {
		 product_id: 12,
		 shop_id: 13,
		 quantity: 1,
		 price: 123
	  }
   ]
}
*/

// var validate = validator.New()

// Struct for Order input with validation rules
type OrderInput struct {
	Address_id   uint      `json:"address_id" binding:"required"`
	User_id      uint      `json:"user_id"`
	Order_date   time.Time `json:"order_date" binding:"required,datetime"`
	Total_amount uint64    `json:"total_amount" binding:"required,gt=0"`
	Products     []Product `json:"products" binding:"required,dive"` // Dive validates each element in the slice
}


// Struct for Product input with validation rules
type Product struct {
	Product_id string `json:"product_id" binding:"required"`
	Shop_id    string `json:"shop_id" binding:"required"`
	Quantity   uint   `json:"quantity" binding:"required,gt=0"`
	Price      uint   `json:"price" binding:"required,gt=0"`
}

func formatValidationErrors(errs validator.ValidationErrors) map[string]string {
	formattedErrors := make(map[string]string)
	for _, err := range errs {
		formattedErrors[err.Field()] = err.Tag() // e.g., "Total_amount": "required"
	}
	return formattedErrors
}

func CreateOrder(c *gin.Context) {
	var order OrderInput

	// Bind and validate input JSON
	if err := c.ShouldBindJSON(&order); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Input validation failed",
			"errors":  formatValidationErrors(validationErrors),
		})
		return
	}

	userId, isExists := c.Get("user_id")

	if !isExists {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "User not found"})
		return
	}

	tx := database.DB.Begin()

	newOrder := order_model.Order{
		User_id:      userId.(uint),
		Address_id:   order.Address_id,
		Order_date:   order.Order_date,
		Total_amount: order.Total_amount,
		Status:       "pending",
	}

	// Attempt to create the order in DB
	if err := tx.Create(&newOrder).Error; err != nil {
		tx.Rollback() // Rollback on error
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to create order", "error": err.Error()})
		return
	}

	// Create order items
	for _, product := range order.Products {
		newProduct := order_model.OrderItem{
			Order_id:   newOrder.ID,
			Product_id: product.Product_id,
			Shop_id:    product.Shop_id,
			Quantity:   product.Quantity,
			Price:      product.Price,
		}
		if err := tx.Create(&newProduct).Error; err != nil {
			tx.Rollback() // Rollback on error
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to create order item", "error": err.Error()})
			return
		}
	}

	// Commit transaction if everything is successful
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to commit transaction", "error": err.Error()})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Order created successfully",
		"order_id": newOrder.ID, // Return order ID for reference
	})
}

type OrderUpdateInput struct {
	Address_id   uint      `json:"address_id"`
	User_id      uint      `json:"user_id"`
	Order_date   time.Time `json:"order_date" building:"datetime"`
	Total_amount uint64    `json:"total_amount" ,building:"gt=0"`
	Products     []Product `json:"products" ,binding:"dive"` // Dive validates each element in the slice
}


// Struct for Product input with validation rules

func UpdateOrder(c *gin.Context){
	var input OrderUpdateInput;
	id := c.Param("id");
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message":"Input invalid", })
	}

	result := database.DB.Model(&order_model.Order{}).Where("id = ?",id).Updates(input);

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message":"Not updated"});
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status":false, "message":"Order not found by Id"})
		return;
	}

	c.JSON(http.StatusOK, gin.H{"status":true, "message":"Order updated", "order":result})
}

func GetAllOrder(c *gin.Context){
 userID := c.Param("id") // User ID from the URL
    var orders []order_model.Order
    var orderItems []order_model.OrderItem

    // Fetch orders for the user
    if err := database.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error fetching orders"})
        return
    }

    // Fetch order items for each order
    for i, order := range orders {
        if err := database.DB.Where("order_id = ?", order.ID).Find(&orderItems).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error fetching order items"})
            return
        }

        // Fetch the product name for each order item
        for j, item := range orderItems {
            productName, err := product_service.FetchSingleProduct(item.Product_id);
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Error fetching product details"})
                return
            }
            orderItems[j].Product_name = productName;
        }

        orders[i].OrderItems = orderItems // Attach the order items with product names to the order
    }

    c.JSON(http.StatusOK, gin.H{"status": true, "orders": orders})
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

	orderItem := []order_model.OrderItem{}

	orderItemList := database.DB.Where("order_id = ?", id).Find(&orderItem);

	if orderItemList.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status":false, "message":"Method not working"});
		return;
	}

	if orderItemList.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message":"order not found by id "});
		return;
	}

	c.JSON(http.StatusOK, gin.H{"status":true, "message":"Order list by user", "list":result, "orderItem":orderItem});
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