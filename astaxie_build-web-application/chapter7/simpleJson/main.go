package main

import (
	"fmt"

	"github.com/bitly/go-simplejson"
)

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	js, err := simplejson.NewJson(b)
	if err != nil {
		fmt.Println(err)
	}
	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()
	fmt.Println(arr)
	fmt.Println(i)
	fmt.Println(ms)
	return
}
