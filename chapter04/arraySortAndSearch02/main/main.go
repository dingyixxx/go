package main

import "fmt"

func main() {
	names := []string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var input string
	fmt.Println("请输入你的名字：")
	fmt.Scanln(&input)

	hasFound := false
	for _, name := range names {
		if name == input {
			fmt.Println("找到了")
			hasFound = true
			break
		}
	}
	if hasFound == false {
		fmt.Println("没找着")
	}
}
