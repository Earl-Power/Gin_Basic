package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "快速入门！"}) })

	router.GET("/student", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "学生信息查询成功"}) })

	router.POST("/create", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "学生信息创建成功"}) })

	router.PUT("/update", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "学生信息更新成功"}) })

	router.DELETE("/delete", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"msg": "学生信息删除成功"}) })
	router.Run()
}
