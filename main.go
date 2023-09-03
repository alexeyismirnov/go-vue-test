package main

import (
	"net/http"

	// "html/template"

	"github.com/gin-gonic/gin"
  "github.com/alexeyismirnov/go-vue-test/admin"
)


func main() {
  db := admin.ConnectDB()
	api := admin.ApiController{db}

	router := gin.Default()
	router.Use(corsMiddleware())

	v1 := router.Group("api/v1")
	{
		v1.GET("getMessage", api.GetMessage)
		v1.POST("setMessage", api.SetMessage)
	}

	router.NoRoute(gin.WrapH(http.FileServer(gin.Dir("web/dist", false))))

	/*
	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl"))
	router.SetHTMLTemplate(templ)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
*/



	router.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
