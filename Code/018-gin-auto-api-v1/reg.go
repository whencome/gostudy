package main

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

type APIRequest interface{}

type APIResponse interface{}

type LogicFunc func(c *gin.Context, r APIRequest) (APIResponse, error)

type APIRequestFunc func() APIRequest

type NeedValidateRequest interface {
	Validate() error
}

func newAPIRequest(v APIRequest) interface{} {
	if v == nil {
		return nil
	}
	rt := reflect.TypeOf(v)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	return reflect.New(rt).Interface()
}

func registerApi(r APIRequest, h LogicFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req APIRequest
		if r != nil {
			req = newAPIRequest(r)
			fmt.Printf("new obj: %T => %+v\n", req, req)
			if err := c.ShouldBind(req); err != nil {
				fmt.Fprintf(c.Writer, "bind param err: %s", err)
				return
			}
			if nvr, ok := req.(NeedValidateRequest); ok {
				if err := nvr.Validate(); err != nil {
					fmt.Fprintf(c.Writer, "request validate err: %s", err)
					return
				}
			}
		}
		resp, err := h(c, req)
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
