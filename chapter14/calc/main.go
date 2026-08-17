package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 迷宫回溯，不写了

// 使用数组来模拟一个栈
type Stack struct {
	MaxTop int     // 栈最大可以存放数的个数
	Top    int     // 栈顶，初始为-1表示空栈
	arr    [20]int // 数组模拟栈
}

// 入栈
func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

// 出栈
func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return
}

// 判断栈是否为空
func (this *Stack) IsEmpty() bool {
	return this.Top == -1
}

// 判断栈是否已满
func (this *Stack) IsFull() bool {
	return this.Top == this.MaxTop-1
}

// 显示栈中所有元素
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

// 判断一个字符是不是一个运算符[+, -, *, /]
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	}
	return false
}

// 运算的方法
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42: //*
		res = num2 * num1
	case 43: //+
		res = num2 + num1
	case 45: //-
		res = num2 - num1
	case 47: // /
		res = num2 / num1
	default:
		fmt.Println("运算符错误.")
	}
	return res
}

// 返回某个运算符的优先级[* / => 1, + - => 0]
func (this *Stack) Priority(oper int) int {
	if oper == '*' || oper == '/' {
		return 1
	} else if oper == '+' || oper == '-' {
		return 0
	}
	return -1
}

func main() {

	//stack := &Stack{
	//	MaxTop: 20,
	//	Top:    -1,
	//}
	//
	//stack.Push(10)
	//stack.Push(20)
	//stack.Push(30)
	//
	//stack.List()
	//
	//val, _ := stack.Pop()
	//fmt.Println("pop:", val)
	//
	//stack.List()
	//数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	exp := "230+130*106-114-226"
	//定义一个index，帮助扫描exp
	index := 0
	//为了配合运算，我们定义需要的变量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""

	for {

		ch := exp[index : index+1]
		temp := int(ch[0]) // byte转int
		if operStack.IsOper(temp) {

			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				if operStack.Priority(operStack.arr[operStack.Top]) >=
					operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					numStack.Push(result)
					operStack.Push(temp)

				} else {
					operStack.Push(temp)
				}

			}
		} else {

			keepNum += ch

			if index == len(exp)-1 {
				//和atoi那个力扣题一样的，java的charArray获取的char，要先转为字符串，再parseInt，再计算
				//char的ascii码值直接计算的话，那就不对了
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				if operStack.IsOper(int(exp[index+1 : index+2][0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}
		if index+1 == len(exp) {
			break
		}
		index++

	}

	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		numStack.Push(result)
	}

	res, _ := numStack.Pop()
	fmt.Printf("表达式%s=%v", exp, res)

}
