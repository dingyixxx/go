package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体
type Monster struct {
	Name     string  `json:"mname"`
	Age      int     `json:"mage"`
	Birthday string  `json:"mbirthday"`
	Sal      float64 `json:"msal"`
	Skill    string  `json:"mskill"`
}

func jsonSlice() {
	var slice []map[string]interface{}

	var m1 map[string]interface{}
	// 使用 map 前，需要先 make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	// 使用 map 前，需要先 make
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [3]string{"墨西哥", "上海", "苏州昆山"}
	slice = append(slice, m2)

	// 序列化为 JSON
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("序列化后=%v\n", string(data))
}

// 对基本数据类型序列化
func testFloat64() {
	var num1 float64 = 2345.67

	// 对 num1 进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("num1 序列化后=%v\n", string(data))
}

func main() {
	//testStruct()
	//testMap()
	jsonSlice()
	//testFloat64()
}

// 将 map 进行序列化
func testMap() {
	// 定义一个 map
	var a map[string]interface{}
	// 使用 map，需要 make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	// 将 a 这个 map 进行序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("map序列化后=%v\n", string(data))
}
func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}

	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Printf("序列化错误...%v", err)
	}
	//data是字节数组
	fmt.Printf("monster序列化错误...%v", string(data))
}
