package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/mgierok/gokeeper/assets/views"
)

//go:generate go run generate.go

func loadViews() multitemplate.Render {
	r := multitemplate.New()

	for _, x := range views.AssetNames() {
		templateString, err := views.Asset(x)
		if err != nil {
			log.Fatal(err)
		}

		tmplMessage, err := template.New(x).Parse(string(templateString))
		if err != nil {
			log.Fatal(err)
		}

		r.Add(x, tmplMessage)
	}

	return r
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.HTMLRender = loadViews()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"appName": "GoKeeper",
		})
	})

	r.Run(":" + port)
}
