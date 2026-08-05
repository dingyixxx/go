package main

// 接口 A
type AInterface interface {
	Test01()
	Test02()
}

// 接口 B
type BInterface interface {
	Test01()
	Test03()
}

// 接口 C 嵌入 A 和 B
type CInterface interface {
	AInterface
	BInterface
}

func main() {
}
