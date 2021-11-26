package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int    `keyname:"a"`
	B string `keyname:"b"`
	// A int
	// B string
	C string
	D string
}

func main() {
	f := Foo{A: 1, B: "hello", C: "world", D: "!!!!!!"}
	ff := &f
	fmt.Printf("%[1]T, %+[1]v\n", f)
	g := StructToMap(f)
	fmt.Printf("%[1]T, %+[1]v\n", g)
	kkk := map[string]interface{}{"A": 23, "B": "23", "C": "sada", "D": "saddd"}
	// fmt.Println("hello")

	FillStruct(kkk, ff)
	// MapToStruct(g, ff)

	fmt.Printf("%[1]T, %+[1]v\n", ff)

}

func FillStruct(data map[string]interface{}, result interface{}) {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)
		fmt.Println(v)
		fmt.Println(val)
		val.Set(reflect.ValueOf(v))
	}
}

func StructToMap(item interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	var res string
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		tagv := fi.Tag.Get("keyname")
		// fmt.Println(tagv)
		if tagv == "" {
			res = fi.Name
		} else {
			res = tagv
		}
		// fmt.Println(i)
		out[res] = v.Field(i).Interface()
	}
	return out
}
