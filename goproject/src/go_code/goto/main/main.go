package main

import "fmt"

func main() {
	//	golang支持用goto，但，话锋一转，能不用goto就不用

	fmt.Println(1)
	fmt.Println(2)
	goto label1
	fmt.Println(3)
label1:
	fmt.Println(4)

}
