package main

import (
	"log"
	"os"

	"sharingvision-be-test/config"
	"sharingvision-be-test/controllers"
	"sharingvision-be-test/repository"
	"sharingvision-be-test/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	_ = godotenv.Load()

	db := config.InitDB()
	if db == nil {
		log.Println("Not Connected to Database")
	}

	repo := repository.NewArticleRepository(db)
	service := services.NewArticleService(repo)
	controller := controllers.NewArticleController(service)

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.POST("/article/", controller.CreateArticle)
	r.POST("/article", controller.CreateArticle)

	r.GET("/article/:param1", controller.GetArticleByID)
	r.GET("/article/:param1/:param2", controller.GetArticles)

	r.PUT("/article/:param1", controller.UpdateArticle)
	r.PATCH("/article/:param1", controller.UpdateArticle)

	r.DELETE("/article/:param1", controller.DeleteArticle)

	r.POST("/article/:param1", func(c *gin.Context) {
		action := c.Query("action")
		if action == "delete" {
			controller.DeleteArticle(c)
		} else {
			controller.UpdateArticle(c)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
