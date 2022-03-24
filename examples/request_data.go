package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2020-09-05" time_uct:"1"`
}

func startPage(ctx *gin.Context) {
	var person Person
	if ctx.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}
	ctx.String(200, "Success")
}
func main() {
	router := gin.Default()
	router.GET("/testing", startPage)
	router.POST("testing", startPage)
	router.Run()
}

/**
Testing
curl -X GET "localhost:8080/testing?name=Earl&address=zxc&birthday=2020-09-05"
curl -X POST http://localhost:8080/testing -d "name=Earl&address=zxc&birthday=2020-09-05"
*/
