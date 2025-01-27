package api

import (
	"admin_shop/pkg/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ShopList avtorizuet polizavatilea v sisteme
// ShopList godoc
// @Summary Register a new user
// @Schemes http
// @Description shoping
// @Tags orders
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /order/list [get]
func OrderList(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)

	var orders []models.Order
	db.Preload("Items").Preload("Items.Shop").Preload("Items.Product").Find(&orders)
	c.JSON(200, gin.H{"orders": orders})

}

// ShopList avtorizuet polizavatilea v sisteme
// ShopList godoc
// @Summary Register a new user
// @Schemes http
// @Description shoping
// @Tags orders
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param  id     path    int     true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /order/show/{id} [get]
func OrderId(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var order models.Order
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	db.Preload("Items").Where(models.Order{ID: idInt}).First(&order)
	c.JSON(200, gin.H{"order": order})

}

// ShopList avtorizuet polizavatilea v sisteme
// ShopList godoc
// @Summary Register a new user
// @Schemes http
// @Description shoping
// @Tags orders
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param  email    path    string     true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /order/email/{email} [get]
func OrderEmail(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var orders []models.Order
	email := c.Param("email")
	db.Where(models.Order{Email: email}).Find(&orders)
	c.JSON(200, gin.H{"orders": orders})
}

// ShopList avtorizuet polizavatilea v sisteme
// ShopList godoc
// @Summary Register a new user
// @Schemes http
// @Description shoping
// @Tags orders
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param   order     body    models.CreateOrder     true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /order/create [post]
func OrderCreate(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var order models.Order
	var createOrder models.CreateOrder
	var orderItems []*models.OrderItem

	if err := c.ShouldBindJSON(&createOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.Email = createOrder.Email
	order.Status = "create"
	for _, i := range createOrder.Items {
		fmt.Println("fdsfds", i)
		var product models.Product
		db.Where(models.Product{ID: i.ProductId}).First(&product)
		order.Sum = order.Sum + (float32(product.Price * float64(i.Count)))
		orderItems = append(orderItems, &models.OrderItem{
			ProductId: product.ID,
			ShopId:    product.ShopID,
			Count:     i.Count,
		})
	}
	db.Create(&order)

	for _, o := range orderItems {
		o.OrderId = order.ID
	}
	db.Create(orderItems)

	c.JSON(200, gin.H{"order": order})
}
