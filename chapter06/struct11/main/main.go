package main

import (
	"encoding/json"
	"fmt"
)

//type Monster struct {
//	name  string
//	age   int
//	skill string
//}

type Monster struct {
	Name  string `json:"nameLa"`
	Age   int    `json:"ageTE"`
	Skill string `json:"skillJr"`
}

func main() {
	var monster Monster
	monster.Name = "红孩儿"
	monster.Age = 35
	monster.Skill = "吐火"
	//monster.name = "红孩儿"
	//monster.age = 35
	//monster.skill = "吐火"
	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Printf("json encoding err:", err)
		return
	}
	fmt.Printf("%T\n", data) //[]uint8

	fmt.Println("json后的数据=", string(data))
	//	json后的数据= {"Name":"红孩儿","Age":35,"Skill":"吐火"}

	//	json后的数据= {}
	//	把结构体改成小写了之后,返回空串了
	//	相当于私有属性,没有setter,无法设置注入

	//	加了tag之后,就是这样了
	//	json后的数据= {"nameLa":"红孩儿","ageTE":35,"skillJr":"吐火"}

}
