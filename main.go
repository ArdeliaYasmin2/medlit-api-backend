package main

import (
	"net/http"
	"os"

	"medlit-api-backend/api/controllers"
	"medlit-api-backend/api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// dev env
	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	db := models.ConnectToDatabase()

	repo := models.NewRepo(db)
	controller := controllers.NewController(repo)

	router := gin.Default()

	api := router.Group("/api/medlit")

	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"error":   "false",
			"message": "Welcome to Medlit API",
		})
	})
	api.POST("/login", controller.Login)
	api.POST("/register", controller.Register)
	api.POST("/medicine/add", controller.AddMedicine)
	api.GET("/medicine/get/all", controller.GetAllMedicine)
	api.GET("/medicine/get/:id", controller.GetMedicineByID)
	api.GET("/medicine/search", controller.GetMedicineByQuery)

	port := ":" + os.Getenv("DB_PORT")
	router.Run(port)
}
