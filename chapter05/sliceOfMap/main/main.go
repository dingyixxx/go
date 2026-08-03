package main

import "fmt"

func main() {
	sliceOfMap := []map[string]string{}
	sliceOfMap = make([]map[string]string, 2)
	sliceOfMap[0] = make(map[string]string)
	sliceOfMap[0]["name"] = "小白"
	sliceOfMap[0]["age"] = "1"

	//和idea一样，也可以.if 然后按tab键
	if sliceOfMap[1] == nil {
		sliceOfMap[1] = make(map[string]string)
		sliceOfMap[1]["name"] = "小新"
		sliceOfMap[1]["age"] = "4"
	}

	m := make(map[string]string)
	m["name"] = "小葵"
	m["age"] = "2"

	sliceOfMap = append(sliceOfMap, m)

	newMonster := map[string]string{
		"name": "火云邪神",
		"age":  "33",
	}
	sliceOfMap = append(sliceOfMap, newMonster)

	for _, v := range sliceOfMap {
		fmt.Println(v)
	}

}
