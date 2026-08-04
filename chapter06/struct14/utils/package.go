package utils

import "fmt"

// 方法绑定的，一般都是“指针”，而不是“值类型”，因为“值类型”拷贝是比较慢的
type SliceType []int

func (self SliceType) speak(num int) {
	self = append(self, num)
}

func (self SliceType) BiggerSpeak(num int) {
	self = append(self, num)
}
func (self *SliceType) speakPtr(num int) {
	*self = append(*self, num)
}

func main() {
	slice := SliceType(make([]int, 0))
	slice = append(slice, 11)
	slice.speak(33)
	fmt.Println(slice)
	fmt.Println("------")

	slice1 := SliceType(make([]int, 0))
	slice1 = append(slice1, 11)
	slice1.speakPtr(33)
	fmt.Println(slice1)
}
