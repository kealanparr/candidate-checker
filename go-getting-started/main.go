package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/kealanparr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goodcandidate.tmpl.html", nil)	
	}

	router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "badcandidate.tmpl.html", nil)
	})

	router.Run(":" + port)
}
