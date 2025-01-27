package api

import (
	"admin_shop/pkg/models"
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
// @Tags shops
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /shop/products/list [get]
func ProductList(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)

	var products []models.Product

	// if err := c.ShouldBindJSON(&userLogin); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	// 	return
	// }

	db.Find(&products)

	c.JSON(200, gin.H{"products": products})
}

// ShopList avtorizuet polizavatilea v sisteme
// ShopList godoc
// @Summary Register a new user
// @Schemes http
// @Description shoping
// @Tags products
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param  id     path    int     true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /shop/products/show/{id} [get]
func ID(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)

	var product models.Product

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	db.Where(models.Product{ID: idInt}).First(&product)
	c.JSON(200, gin.H{"product": product})

}

// ShopList avtorizuet polizavatilea v sisteme
// ShopList godoc
// @Summary Register a new user
// @Schemes http
// @Description shoping
// @Tags shops
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param  shop_id     path    int     true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /shop/{shop_id}/products/list [get]
func ProductShopList(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)

	var products []models.Product
	shopId := c.Param("shop_id")

	db.Where(models.Product{ShopID: shopId}).First(&products)
	c.JSON(200, gin.H{"products": products})

}

// ShopList avtorizuet polizavatilea v sisteme
// ShopList godoc
// @Summary Register a new user
// @Schemes http
// @Description shoping
// @Tags shops
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param  category     path    string    true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /shop/products/by-category/{category} [get]
func ProductShowCategory(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)

	var products []models.Product

	category := c.Param("category")
	db.Where(models.Product{Category: category}).Find(&products)
	c.JSON(200, gin.H{"products": products})

}

// ShopList avtorizuet polizavatilea v sisteme
// ShopList godoc
// @Summary Register a new user
// @Schemes http
// @Description shoping
// @Tags shops
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param  shop_id    path    int    true        "User registration object"
// @Param  category     path    string    true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /shop/{shop_id}/products/by-category/{category} [get]
func ProductByShopAndCategory(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var products []models.Product

	category := c.Param("category")
	shopId := c.Param("shop_id")

	db.Where(models.Product{Category: category, ShopID: shopId}).Find(&products)
	c.JSON(200, gin.H{"products": products})
}

// ShopList avtorizuet polizavatilea v sisteme
// ShopList godoc
// @Summary Register a new user
// @Schemes http
// @Description shoping
// @Tags shops
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param   product     body    models.Product     true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /shop/products/create [post]
func ProductCreate(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&product)
	c.JSON(200, gin.H{"products": product})
}
