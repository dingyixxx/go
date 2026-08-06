package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"mname"`
	Age      int     `json:"mage"`
	Birthday string  `json:"mbirthday"`
	Sal      float64 `json:"msal"`
	Skill    string  `json:"mskill"`
}

func unmarshalStruct() {
	str := "{\"mname\":\"牛魔王\",\"mage\":500,\"mbirthday\":\"2011-11-11\",\"msal\":8000,\"mskill\":\"牛魔拳\"}\n"
	var m Monster
	json.Unmarshal([]byte(str), &m) //引用的方式，传入
	fmt.Println(m.Name)
	fmt.Println(m)
}
func unmarshalMap() {
	str := "{\"address\":\"洪崖" +
		"洞\",\"age\":30,\"name\":\"红孩儿\"}"
	var m map[string]interface{}
	json.Unmarshal([]byte(str), &m) //反序列化，会自动make
	fmt.Println(m)
}
func unmarshalFloat() {
	str := "2345.67"
	var m float64
	json.Unmarshal([]byte(str), &m)
	fmt.Println(m)
}

func unmarshalSliceOfMap() {
	var slice []map[string]interface{} //slice和map都被封装在Unmarshal里了
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"},{\"address\":[\"墨西哥\",\"上海\",\"苏州昆山\"],\"age\":\"20\",\"name\":\"tom\"}]\n"
	json.Unmarshal([]byte(str), &slice)
	fmt.Println(slice)
}
func main() {
	//unmarshalStruct()
	//unmarshalMap()
	//unmarshalFloat()
	unmarshalSliceOfMap()
}
