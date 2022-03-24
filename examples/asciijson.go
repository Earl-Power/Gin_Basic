package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/asciiJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Gin框架",
			"tag":  "<br>",
		}

		// 输出: {"lang":"Gin\u6846\u67b6","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
