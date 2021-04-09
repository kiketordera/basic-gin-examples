package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Path to the static files. /static is rendered in the HTML and /media is the link to the path to the  images, svg, css.. the static files
	r.StaticFS("/static", http.Dir("ui/css"))

	// Path to the HTML templates. * is a wildcard
	r.LoadHTMLGlob("ui/html/*/*.html")

	r.GET("/hello", RenderLanding)
	r.POST("/hello", GetForm)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// RenderLanding shows the landing page
func RenderLanding(c *gin.Context) {
	c.HTML(http.StatusOK, "landing.html", gin.H{})
}

// GetForm gets the
func GetForm(c *gin.Context) {
	// We define the data to fetch
	formData := &struct {
		FirstField  string `form:"first_field" binding:"required"`
		SecondField string `form:"second_field" binding:"required"`
		ThirdField  string `form:"third_field" binding:"required"`
	}{}
	// Now we fetch the data from the form
	if err := c.Bind(formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// And render the data we have fetched
	c.HTML(http.StatusOK, "result_form.html", gin.H{
		"hello":  "Hello world!",
		"first":  formData.FirstField,
		"second": formData.SecondField,
		"third":  formData.ThirdField,
	})
}
