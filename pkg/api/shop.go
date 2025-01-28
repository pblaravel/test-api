package api

import (
	"admin_shop/pkg/models"
	"net/http"
	"strconv"

	"github.com/google/uuid"

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
// @Router /shop/list [get]
func ShopList(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)

	var shops []models.Shop

	// if err := c.ShouldBindJSON(&userLogin); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	// 	return
	// }

	db.Find(&shops)

	c.JSON(200, gin.H{"shops": shops})
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
// @Param   id     path    int     true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /shop/show/{id} [get]
func ShopShow(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var shop models.Shop
	id := c.Param("id")
	db.Where(models.Shop{UUID: id}).First(&shop)

	var orders []models.Order
	db.Preload("Items", "shop_id = ?", id).Find(&orders)
	var sum float32 = 0
	for _, o := range orders {
		if len(o.Items) != 0 {
			sum = sum + o.Sum
		}
	}

	c.JSON(200, gin.H{"shop": shop, "sum": sum})
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
// @Param   id     path    int     true        "User registration object"
// @Param   shop     body    models.Shop     false        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /shop/update/{id} [post]
func ShopUpdate(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)

	var income models.Shop
	id := c.Param("id")

	if err := c.ShouldBindJSON(&income); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	db.Where(models.Shop{UUID: id}).Updates(&income)
	c.JSON(200, gin.H{"message": "update success"})
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
// @Param  status     path    int     true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /shop/by-status/{status} [get]
func ShopByStatus(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)

	var shops []models.Shop
	status := c.Param("status")
	idInt, _ := strconv.Atoi(status)

	db.Where(models.Shop{Status: idInt}).Find(&shops)
	c.JSON(200, gin.H{"shops": shops})
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
// @Router /shop/by-category/{category} [get]
func ShopByCategory(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var shops []models.Shop
	category := c.Param("category")

	db.Where(models.Shop{Category: category}).Find(&shops)
	c.JSON(200, gin.H{"shop": category})

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
// @Param  name     path    string    true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /shop/by-name/{name} [get]
func ShopName(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var shops []models.Shop
	name := c.Param("name")

	db.Where(models.Shop{Name: name}).Find(&shops)
	c.JSON(200, gin.H{"shops": shops})

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
// @Param   product     body    models.Shop     true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /shop/create [post]
func ShopCreate(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var shop models.Shop
	if err := c.ShouldBindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shop.UUID = uuid.New().String()
	db.Create(&shop)
	c.JSON(200, gin.H{"shops": shop})

}

// ShopList avtorizuet polizavatilea v sisteme
// ShopList godoc
// @Summary Register a new user
// @Schemes http
// @Description shoping
// @Tags category
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param   category     body    models.Category     true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /category/create [post]
func CategoryCreate(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var category models.Category
	var categoryDB models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Where(models.Category{Name: category.Name}).First(&categoryDB)
	if categoryDB.ID != 0 {
		c.JSON(404, gin.H{"error": "category is exist"})
		return
	}
	db.Create(&category)
	c.JSON(200, gin.H{"shops": category})

}

// ShopList avtorizuet polizavatilea v sisteme
// ShopList godoc
// @Summary Register a new user
// @Schemes http
// @Description shoping
// @Tags category
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /category/list [get]
func CategoryList(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var categores []models.Category

	db.Find(&categores)
	c.JSON(200, gin.H{"shops": categores})

}

// ShopList avtorizuet polizavatilea v sisteme
// ShopList godoc
// @Summary Register a new user
// @Schemes http
// @Description shoping
// @Tags subCategory
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param   subCategory     body    models.SubCategory     true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /subCategory/create [post]
func SubcategoryCreate(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var subCategory models.SubCategory
	if err := c.ShouldBindJSON(&subCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&subCategory)
	c.JSON(200, gin.H{"shop": subCategory})
}

// ShopList avtorizuet polizavatilea v sisteme
// ShopList godoc
// @Summary Register a new user
// @Schemes http
// @Description shoping
// @Tags subCategory
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /subCategory/list [get]
func SubCategoryList(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var subCategores []models.SubCategory

	db.Find(&subCategores)
	c.JSON(200, gin.H{"shops": subCategores})

}
