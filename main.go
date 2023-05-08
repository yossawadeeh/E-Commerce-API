package main

import (
	"e-commerce-api/database"
	"e-commerce-api/docs"
	"e-commerce-api/models"
	"e-commerce-api/routes"
	"e-commerce-api/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var runEnv string

func init() {
	fmt.Println("Hii 😊")
	fmt.Println("Hii 🌲")

	runEnv = os.Getenv("RUN_ENV")
	if runEnv == "" {
		runEnv = "dev"
	}
	utils.InitViper()

	var err error
	if err = database.ConnectDB(); err != nil {
		log.Fatal(err)
	}
}

// @Security bearer
// @securityDefinitions.apikey bearer
// @in header
// @name Authorization

func main() {

	docs.SwaggerInfo.Title = "Ecommerce API"
	docs.SwaggerInfo.Description = "Ecommerce API with gin framework"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8001"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app := gin.Default()

	//app.Use(cors.Default())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"*"}
	app.Use(cors.New(config))

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.ErrorMessage("Not Found Route", 404))
	})
	app.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, utils.ErrorMessage(http.StatusText((http.StatusMethodNotAllowed)), 405))
	})

	// test end point
	app.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	app.GET("/test/test", (func(c *gin.Context) {

		res := []models.Test{}
		if err := database.DB.Find(&res).Error; err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorMessage(http.StatusText((http.StatusInternalServerError)), 500))
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"data": res,
		})
	}))
	// test end point

	// routes group
	v1 := app.Group("/v1")
	routes.ShopRoute(v1)
	routes.AuthUserRoute(v1)
	routes.CustomerRoute(v1)
	routes.OrderRoute(v1)
	routes.ProductRoute(v1)

	// http://localhost:8001/swagger/index.html
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Run(":8001")
}
