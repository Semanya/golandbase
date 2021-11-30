package main

import (
	"errors"
	"fmt"
	"reflect"
)

type StructWithMoreNested struct {
	StructWithStructField `keyname:"struct_with_struct_field"`
	Message               string `keyname:"message"`
}

type StructWithStructField struct {
	Message string       `keyname:"message"`
	Simple  SimpleStruct `keyname:"simple"`
}

type SimpleStruct struct {
	Name string `keyname:"name"`
	ID   int    `keyname:"id"`
}

var simple = map[string]interface{}{
	"name": "Superman",
	"id":   3388,
}
var withStructField = map[string]interface{}{
	"simple":  SimpleStruct{Name: "Batman", ID: 123},
	"message": "hello",
}
var withAnotherStrunct = map[string]interface{}{
	"struct_with_struct_field": StructWithStructField{
		Simple: SimpleStruct{
			Name: "Batman",
			ID:   123,
		},
		Message: "i am",
	},
	"message": "the night",
}

func main() {
	///////////////////////////////////////////
	/////////////// FIRST CASE ////////////////
	///////////////////////////////////////////
	var res SimpleStruct
	MapToStruct(simple, &res)
	// fmt.Printf("%+[1]v\n", res.Name)
	// fmt.Printf("%+[1]v\n", res.ID)
	fmt.Printf("%+[1]v\n", res)
	///////////////////////////////////////////
	////////////// SECOND CASE ////////////////
	///////////////////////////////////////////
	// var res2 StructWithStructField
	// MapToStruct(withStructField, &res2)
	// // fmt.Printf("%+[1]v\n", res2.Message)
	// // fmt.Printf("%+[1]v\n", res2.Simple.Name)
	// // fmt.Printf("%+[1]v\n", res2.Simple.ID)
	// fmt.Printf("%+[1]v\n", res2)
	// ///////////////////////////////////////////
	// /////////////// THIRD CASE ////////////////
	// ///////////////////////////////////////////
	// var res3 StructWithMoreNested
	// MapToStruct(withAnotherStrunct, &res3)
	// // fmt.Printf("%+[1]v\n", res3.Message)
	// // fmt.Printf("%+[1]v\n", res3.StructWithStructField.Message)
	// // fmt.Printf("%+[1]v\n", res3.Simple.Name)
	// // fmt.Printf("%+[1]v\n", res3.Simple.ID)
	// fmt.Printf("%+[1]v\n", res3)
}

func MapToStruct(data map[string]interface{}, result interface{}) error {
	t := reflect.ValueOf(result).Elem()

	tag := "keyname"
	r := reflect.ValueOf(result)

	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	} else if r.Kind() != reflect.Struct {
		return nil
	}

	typ := r.Type()
	var tagv string
	for k, v := range data {
		for i := 0; i < r.NumField(); i++ {
			tagv = typ.Field(i).Tag.Get(tag)
			name := typ.Field(i).Name
			fmt.Println("k, tagv, name", k, tagv, name)
			if tagv != "" {
				if k == tagv {
					k = name
					val := t.FieldByName(k)
					val.Set(reflect.ValueOf(v))
				}
			} else {
				if k == name {
					val := t.FieldByName(k)
					val.Set(reflect.ValueOf(v))
				} else {
					return errors.New("No key in struct")
				}
			}
		}
	}
	return nil
}
