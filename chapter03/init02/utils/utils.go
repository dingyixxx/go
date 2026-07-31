package utils

import "fmt"

var UtilName string
var UtilAge int
var UtilNum = getNumInner()

func init() {
	UtilName = "打印机"
	UtilAge = 10
	fmt.Println("utils包...init...")
}

func getNumInner() int {
	fmt.Println("utils包...getNumInner...")
	return 1
}
