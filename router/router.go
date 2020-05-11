package main

import (
	"github.com/gin-gonic/gin"
)

func HomePage (c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func PostHomePage(c *gin.Context)  {
	a := c.PostForm("a")
	b := c.PostForm("b")

	c.JSON(200, gin.H{
		"a": a,
		"b": b,
	})
}

func QueryString(c *gin.Context)  {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}

func Path(c *gin.Context)  {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}


func main()  {
	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryString)
	r.GET("/path/:name/:age", Path)
	r.Run()
}
