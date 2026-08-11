package main

import (
	"reflect"
	"testing"
)

type user struct {
	UserId string
	Name   string
}

func TestReflectStruct(t *testing.T) {
	var (
		model *user
		sv    reflect.Value
	)

	model = &user{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.ValueOf()", sv.Kind().String()) //struct
	sv = sv.Elem()
	t.Log("reflect.ValueOf.Elem()", sv.Kind().String())
	sv.FieldByName("UserId").SetString("12345678")
	sv.FieldByName("Name").SetString("nickname")
	t.Log("model=", model)
}
