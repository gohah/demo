package main

import(
	//"fmt"
	"reflect"
)

type Human struct {
	name string
	age int
}

func main() {


	//var x float64 = 3.4
	//v := reflect.ValueOf(x)
	//fmt.Println("type:", v.Type())
	//fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	//fmt.Println("value:", v.Float())

	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(7.1)
}