package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// id请求
type IDRequest struct {
	ID int64 `json:"id" form:"id" label:"ID" binding:"required,gt=0"`
}

func IDRequestLogic(c *gin.Context, r APIRequest) (APIResponse, error) {
	fmt.Println("----- ok -----")
	req, ok := r.(*IDRequest)
	if !ok {
		return nil, fmt.Errorf("request convert failed")
	}
	fmt.Println("IDRequest.ID = ", req.ID)
	return "222", nil
}

func main() {
	r := gin.Default()
	r.GET("/test", Api(func() APIRequest {
		return new(IDRequest)
	}, func(c *gin.Context, r APIRequest) (APIResponse, error) {
		fmt.Println("----- ok -----")
		req, ok := r.(*IDRequest)
		if !ok {
			return nil, fmt.Errorf("request convert failed")
		}
		fmt.Println("IDRequest.ID = ", req.ID)
		return "111", nil
	}))
	r.GET("/test1", Api(func() APIRequest {
		return new(IDRequest)
	}, IDRequestLogic))
	r.Run()
}
