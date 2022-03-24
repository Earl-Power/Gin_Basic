package main

import "github.com/gin-gonic/gin"

// StructA 最基本的结构体
type StructA struct {
	FieldA string `form:"field_a"`
}

// StructB 嵌套结构体的结构体
type StructB struct {
	NestedStructPointer StructA
	FieldB              string `form:"field_b"`
}

// StructC 嵌套结构体-指针结构体
type StructC struct {
	NestedStruct *StructA
	FieldC       string `form:"field_c"`
}

// StructD 嵌套匿名结构体的结构体
type StructD struct {
	NestedAntonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

// GetDataB 返回StructB
func GetDataB(c *gin.Context) {
	var b StructB
	// 读取请求数据并并拼写结构体b
	c.Bind(&b)
	// 返回JSON 格式响应
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"b": b.FieldB,
	})
}

// GetDataC 返回StructC
func GetDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"c": b.FieldC,
	})
}

// GetDataD 返回StructD
func GetDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAntonyStruct,
		"d": b.FieldD,
	})
}

func main() {
	r := gin.Default()
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)
	r.Run()

}
