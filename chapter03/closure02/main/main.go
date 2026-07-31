package main

import (
	"fmt"
	"strings"
)

func main() {
	//suffix := strings.HasSuffix("aaa.txt", "txt")
	//fmt.Println(suffix)

	genName := func(s string) func(string) string {
		suffix := s
		return func(name string) string {
			if strings.HasSuffix(name, s) {
				return name
			}
			return name + "." + suffix
		}
	}
	f := genName("webp")
	fmt.Println(f("aaa.webp"))
	fmt.Println(f("dingyi"))
	fmt.Println(f("dingyi.png"))

}

//如果传入的名字有后缀，就返回原来的名字
//如果没有后缀，则加上指定的后缀defaultSuffix
