package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 通过字典模拟Database
var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// 初始化 Gin框架默认实例，该实例包含了路由、中间件以及配置信息
	Routes := gin.Default()

	//Ping测试路由
	Routes.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 获取用户数据路由
	Routes.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no valus"})
		}
	})

	// 需要HTTP基本授权认证的子路由群组设置
	authorized := Routes.Group("/", gin.BasicAuth(gin.Accounts{
		"foot": "bar", //用户：foo 密码：bar
		"manu": "123", //用户：manu 密码：123
	}))

	// 保存用户路由信息
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		// 解析并验证 JSON 格式请求数据
		var json struct {
			Value string `json:"value" binding:"required"`
		}
		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})
	return Routes
}

func main() {
	// 设置路由信息
	Routes := setupRouter()
	// 启动服务器并监听 8080 端口
	Routes.Run(":8080")
	/**
	Testing
	go run main.go
	curl http://localhost:8080/ping
	curl http://localhost:8080/user/foo
	curl http://localhost:8080/user/manu
	curl -X POST -H "Content-Type:application/json" -d '{"user":"foo","value":"1"}' foo:bar@localhost:8080/admin
	curl -X POST -H "Content-Type:application/json" -d '{"user":"manu","value":"1"}' manu:123@localhost:8080/admin

	*/
}
