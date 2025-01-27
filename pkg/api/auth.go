package api

import (
	"admin_shop/pkg/auth"
	"admin_shop/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Login avtorizuet polizavatilea v sisteme
// Login godoc
// @Summary Register a new user
// @Schemes http
// @Description Registers a new user with the given username and password
// @Tags user
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param   user     body    models.User     true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)

	var userLogin models.User
	var dbUser models.User

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	db.Where(models.User{Username: userLogin.Username}).First(&dbUser)

	if dbUser.ID == 0 {
		c.JSON(401, gin.H{"error": "user not found"})
		return
	}

	token, err := auth.GenerateToken(dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}
	c.JSON(200, gin.H{"token": token, "user": dbUser})
}

// Register - reghistriruet novovo polizavitilea
// Register godoc
// @Summary Register a new user
// @Schemes http
// @Description Registers a new user with the given username and password
// @Tags user
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param   user     body    models.User     true        "User registration object"
// @Success 200 {string} string "Successfully registered"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	db, _ := c.MustGet("db").(*gorm.DB)
	var user models.User
	var dbUser models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Where(models.User{Username: user.Username}).First(&dbUser)
	if dbUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userName iz uset"})
		return
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not hash password"})
		return
	}

	user.Password = hashedPassword
	// novaia structura dlea sahranenia v bazu
	newUser := models.User{Name: user.Name, Username: user.Username, Password: user.Password}

	db.Create(&newUser)

	token, err := auth.GenerateToken(newUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "message": "registration successful"})
}
