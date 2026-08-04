package model

import "fmt"

// 结构体 A
type A struct {
	Name     string // 大写，包外可访问
	age      int    // 小写，仅包内可访问
	NianLing int    // 小写，仅包内可访问
}

func (a *A) SayOk() { // 大写方法，包外可访问
	fmt.Printf("A SayOk-%q\n", a.Name)
}

func (a *A) hello() { // 小写方法，仅包内可访问
	fmt.Printf("A hello\n", a.Name)
}

func (a *A) Hello() {
	fmt.Printf("A Hello-%q\n", a.Name)
}
