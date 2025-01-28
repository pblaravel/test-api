package api

import (
	"admin_shop/pkg/models"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "admin_shop/docs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ContextMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(ContextMiddleware(db))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/api/v1")
	{

		auth := v1.Group("/auth")
		{
			auth.POST("/register", Register)
			auth.POST("/login", Login)
		}

		v1.GET("/", func(ctx *gin.Context) {
			db, _ := ctx.MustGet("db").(*gorm.DB)
			var persons models.User
			db.Where(&models.User{Name: "Aliona"}).First(&persons)
			ctx.JSON(200, persons)
		})
		// name/:name
		v1.GET("/name/:name", func(ctx *gin.Context) {
			db, _ := ctx.MustGet("db").(*gorm.DB)
			name := ctx.Param("name")
			var persons models.User
			db.Where(&models.User{Name: name}).First(&persons)
			ctx.JSON(200, persons)
		})
		v1.GET("/person/add/:name/:age/:adres/:prof", func(ctx *gin.Context) {
			db, _ := ctx.MustGet("db").(*gorm.DB)
			var persons models.User

			if err := ctx.ShouldBindUri(&persons); err != nil {
				ctx.JSON(200, gin.H{"msg": err.Error()})
				return
			}
			db.Create(&persons)
			ctx.JSON(200, persons)
		})

		order := v1.Group("/order")
		{
			order.POST("create", OrderCreate)
			order.GET("/list", OrderList)
			order.GET("/show/:id", OrderId)
			order.GET("/email/:email", OrderEmail)
		}
		category := v1.Group("category")
		{
			category.GET("/list", CategoryList)
			category.POST("/create", CategoryCreate)
		}
		subCategory := v1.Group("subCategory")
		{
			subCategory.POST("/create", SubcategoryCreate)
			subCategory.GET("/list", SubCategoryList)
		}

		shop := v1.Group("/shop")
		{
			shop.POST("/create", ShopCreate)
			shop.GET("/show/:id", ShopShow)
			shop.POST("/update/:id", ShopUpdate)
			shop.GET("/by-status/:status", ShopByStatus)
			shop.GET("/by-name/:name", ShopName)

			// nado sdelati grupu
			shop.GET("/products/list", ProductList)
			shop.POST("/products/create", ProductCreate)
			shop.GET("/products/show/:id", ID)
			shop.GET("/products/by-category/:category", ProductShowCategory)

			shop.GET("/:shop_id/products/by-category/:category", ProductByShopAndCategory)
			shop.GET("/:shop_id/products/list", ProductShopList)
			shop.GET("/by-category/:category", ShopByCategory)
			shop.GET("delete/:id", func(ctx *gin.Context) {
				ctx.JSON(200, "delete/:id")
			})
			shop.GET("/list", ShopList)

			member := shop.Group("/member")
			{
				member.GET("/add", func(ctx *gin.Context) {
					ctx.JSON(200, "add")
				})
				member.GET("/list", func(ctx *gin.Context) {
					ctx.JSON(200, "list")
				})
				member.GET("/update", func(ctx *gin.Context) {
					ctx.JSON(200, "update")
				})
				member.GET("/delete", func(ctx *gin.Context) {
					ctx.JSON(200, "delete")
				})
			}
		}

		// v1.GET("/shop/create", func(ctx *gin.Context) {
		// 	ctx.JSON(200, "test")
		// })
		// v1.GET("/shop/show/:id", func(ctx *gin.Context) {
		// 	ctx.JSON(200, "text")
		// })
	}

	return r
}
