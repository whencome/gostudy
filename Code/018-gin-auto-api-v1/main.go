package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type IDRequest struct {
	ID   int64  `json:"id" form:"id" label:"ID" binding:"required,gt=0"`
	Name string `json:"name" form:"name"`
}

func (r IDRequest) Validate() error {
	return nil
}

type StuRequest struct {
	Name  string `json:"name" form:"name" label:"姓名" binding:"required"`
	Age   int    `json:"age" form:"age" label:"年龄" binding:"required,gt=0"`
	Class string `json:"class" form:"class" label:"班级"`
}

func (r StuRequest) Validate() error {
	if r.Age <= 0 || r.Age > 100 {
		return fmt.Errorf("invalid student age of value %d, the age should greater than 0 and less or equal 100", r.Age)
	}
	return nil
}

func IDRequestLogic(c *gin.Context, r APIRequest) (APIResponse, error) {
	fmt.Println("----- ok -----")
	fmt.Printf("request: %+v\n", r)
	req, ok := r.(*IDRequest)
	if !ok {
		return nil, fmt.Errorf("request convert failed")
	}
	fmt.Println("IDRequest.ID = ", req.ID)
	fmt.Println("IDRequest.Name = ", req.Name)
	return "success", nil
}

func StuRequestLogic(c *gin.Context, r APIRequest) (APIResponse, error) {
	fmt.Println("----- ok -----")
	fmt.Printf("request: %+v\n", r)
	req, ok := r.(*StuRequest)
	if !ok {
		return nil, fmt.Errorf("request convert failed")
	}
	fmt.Println("StuRequest.Name = ", req.Name)
	fmt.Println("StuRequest.Age = ", req.Age)
	fmt.Println("StuRequest.Class = ", req.Class)
	return "success", nil
}

func main() {
	r := gin.Default()
	r.GET("/test", Api(IDRequest{}, IDRequestLogic))
	r.GET("/stu", Api(StuRequest{}, StuRequestLogic))
	r.Run()
}
