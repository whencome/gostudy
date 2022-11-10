package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	StrField string
	IntField int
}

func createObj(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	rt := reflect.TypeOf(v)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	newV := reflect.New(rt)
	return newV.Interface()
}

func main() {
	f1 := &Foo{
		StrField: "a string",
		IntField: 22,
	}
	f2 := createObj(f1)
	if f2 == nil {
		fmt.Println("create new object failed")
		return
	}
	fmt.Printf("new obj: %+v\n", f2)
	f3 := f2.(*Foo)
	f3.StrField = "new string"
	f3.IntField = 33
	fmt.Printf("f3: %+v\n", f3)
}
