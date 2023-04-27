package main

import (
	"e-commerce-api/database"
	"e-commerce-api/models"
	"e-commerce-api/routes"
	"e-commerce-api/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var runEnv string

func init() {
	fmt.Println("Hii 😊")

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

func main() {
	app := gin.Default()

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.ErrorMessage(http.StatusText((http.StatusNotFound)), 404))
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

	app.Run(":8001")
}
