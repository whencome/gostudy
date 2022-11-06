package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type APIRequest interface{}

type APIResponse interface{}

type LogicFunc func(c *gin.Context, r APIRequest) (APIResponse, error)

type APIRequestFunc func() APIRequest

func registerApi(rf APIRequestFunc, h LogicFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r APIRequest
		if rf != nil {
			r = rf()
			if err := c.ShouldBind(r); err != nil {
				fmt.Fprintf(c.Writer, "bind param err: %s", err)
				return
			}
		}
		resp, err := h(c, r)
		if err != nil {
			fmt.Fprintf(c.Writer, "logic handle err: %s", err)
			return
		}
		fmt.Fprintf(c.Writer, "logic handle succ: %+v", resp)
		return
	}
}

func Api(rf APIRequestFunc, h LogicFunc) gin.HandlerFunc {
	return registerApi(rf, h)
}
