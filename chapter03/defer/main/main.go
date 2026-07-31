package main

import "fmt"

func sum(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	//defer 数据库连接、文件描述符、锁用完关闭，类似于java的finally
	//LIFO
	//入defer栈的是“拷贝”（而不是“引用”）
	res := sum(10, 20)
	fmt.Println("res=", res)

	//	defer写完之后，还可以继续写查询db或者读文件的代码
	//例如
	//	file := openFile("test.txt")
	//	defer file.close()        // ← 打开后立刻 defer
	//
	//	// 下面安心写业务逻辑，不用管关闭的事
	//	data := file.read()
	//	file.write("hello")
	//	return

}
