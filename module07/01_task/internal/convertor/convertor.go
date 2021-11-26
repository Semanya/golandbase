package convertor

import (
	"errors"
	"reflect"
)

func StructToMap(item interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	var res string
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		tagv := fi.Tag.Get("keyname")
		if tagv == "" {
			res = fi.Name
		} else {
			res = tagv
		}
		out[res] = v.Field(i).Interface()
	}
	return out
}

func MapToStruct(mp map[string]interface{}, item interface{}) error {
	t := reflect.ValueOf(item).Elem()

	tag := "keyname"
	r := reflect.ValueOf(item)

	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	} else if r.Kind() != reflect.Struct {
		return nil
	}

	typ := r.Type()
	var tagv string
	for i := 0; i < r.NumField(); i++ {
		tagv = typ.Field(i).Tag.Get(tag)
		name := typ.Field(i).Name
		if tagv != "" {
			for k, v := range mp {
				if k == tagv {
					k := name
					val := t.FieldByName(k)
					val.Set(reflect.ValueOf(v))
				}
			}
		} else {
			return errors.New("no tag")
		}
	}
	return nil
}
