package main

import (
	"admin_shop/pkg/api"
	"admin_shop/pkg/database"

	"github.com/gin-gonic/gin"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /api/v1

// @securityDefinitions.apikey JwtAuth
// @in header
// @name Authorization
func main() {

	db := database.NewDatabase()

	gin.SetMode(gin.DebugMode)

	router := api.NewRouter(db)

	if err := router.Run(":8002"); err != nil {
		panic(err)
	}
}
