package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func A(c *gin.Context) {
	fmt.Println("A1")
	c.Next()
	fmt.Println("A2")
}

func B(c *gin.Context) {
	fmt.Println("B1")
	c.Next()
	fmt.Println("B2")
}

func hello(c *gin.Context) {}

func main() {
	r := gin.Default()
	// when a request comes in, it will print A1 B1 B2 A2
	r.Use(A, B)
	r.GET("/hello", hello)
	r.Run()
}
