package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8888"
		log.Print("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/chapter-1", func(c *gin.Context) {
		sum := Add(2, 2)
		expected := 4

		values := []obj{{Key: "sum?",Value: sum}, {Key: "expected?",Value: expected}}
		c.HTML(http.StatusOK, "basic-learn.tmpl.html", gin.H{
			"chapter": "정수",
			"obj": values,
		})
	})

	router.GET("/chapter-2", func(c *gin.Context) {
		values := []obj{{Key: "",Value: ""}}
		c.HTML(http.StatusOK, "basic-learn.tmpl.html", gin.H{
			"chapter": "반복",
			"obj": values,
		})
	})

	router.Run(":" + port)
}