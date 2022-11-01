package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type APIRequest interface{}

type APIResponse interface{}

type LogicFunc func(c *gin.Context, r APIRequest) (APIResponse, error)

func registerApi(r APIRequest, h LogicFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBind(r); err != nil {
			fmt.Fprintf(c.Writer, "bind param err: %s", err)
			return
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

func Api(r APIRequest, h LogicFunc) gin.HandlerFunc {
	return registerApi(r, h)
}
