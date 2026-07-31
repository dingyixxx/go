package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Name :="22222" 不能这么写，变量赋值不能在函数外
// fmt.Println(Name)
// 为什么我从未好奇过BigDecimal、Hutool、StringBuilder的各个API的源码呢，初始化了什么，怎么计算的，做了什么
func main() {
	fmt.Println(len("helloWorld网")) //13 utf-8 1个中文=3个字节

	sentence := "你好世界"
	strs := []rune(sentence)
	for i := 0; i < len(strs); i++ {
		fmt.Printf("%c\n", strs[i])
	}

	num, err := strconv.Atoi("123323.3243434")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

	num1 := strconv.Itoa(98765)
	fmt.Println(num1)

	str2 := "hello*go"
	bytes := []byte(str2)
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%v-%c  ", bytes[i], bytes[i])
	}
	fmt.Println()

	//	字节数组 转 字符串
	str3 := string([]byte{65, 66, 67, 68}) //ABCD
	fmt.Println(str3)

	fmt.Println()

	//	10进制 转 2、8、16进制
	num2 := strconv.FormatInt(123, 2)
	num8 := strconv.FormatInt(123, 8)
	num16 := strconv.FormatInt(123, 16)
	fmt.Println(num2)  //1111011
	fmt.Println(num8)  //173
	fmt.Println(num16) //7b

	fmt.Println()
	fmt.Printf("%b\n", 123)
	fmt.Printf("%O\n", 123)
	fmt.Printf("%X\n", 123)

	fmt.Println()
	n2 := fmt.Sprintf("%b", 123)
	n8 := fmt.Sprintf("%O", 123)
	n16 := fmt.Sprintf("%X", 123) //fmt.Sprintf 和 strconv.FormatInt 都可以把数字“转化为”二进制，而不仅仅是“打印出来”
	fmt.Println(n2)
	fmt.Println(n8)
	fmt.Println(n16)

	fmt.Println()
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "mary"))
	fmt.Println()

	count := strings.Count("hello", "l")
	fmt.Println(count)

	fmt.Println()
	fmt.Println(strings.EqualFold("abc", "aBC")) //忽略大小写，所以是true

	index := strings.Index("ABCSA_abc_abc", "abc")
	fmt.Println(index)
	fmt.Println(strings.Index("ABCSA_abc_abc", "hhg"))
	fmt.Println()
	fmt.Println(strings.LastIndex("go_goland_123_golang", "gol")) //14
	fmt.Println()

	replace := strings.Replace("go_goland_123_golang", "go", "小苹果", -1) //-1替换全部
	fmt.Println(replace)
	replace1 := strings.Replace("go_goland_123_golang", "go", "1小苹果1", 1)
	fmt.Println(replace1)

	fmt.Println()
	splitArr := strings.Split("hello,world,here,I,am", ",")
	for i := 0; i < len(splitArr); i++ {
		fmt.Printf("splitArr[%v]=%v\n", i, splitArr[i])
	}
	fmt.Println()
	fmt.Printf("%v\n", splitArr) //[hello world here I am]
	fmt.Println()

	str1 := "hello,wOrlD,heRe,I,am"
	fmt.Println(strings.ToLower(str1))
	fmt.Println(strings.ToUpper(str1))
	fmt.Println(str1)

	str13 := strings.Trim("! he!l!lo! ", " !") //he!l!lo
	fmt.Println(str13)

	strL := strings.TrimLeft("! he!l!lo! ", " !") //he!l!lo!
	fmt.Println(strL)

	strR := strings.TrimRight("! he!l!lo! ", " !") //! he!l!lo
	fmt.Println(strR)

	strS := strings.TrimSpace("    ! h  e ! l  !l o! ") //! h  e ! l  !l o!
	fmt.Println(strS)

	fmt.Println(strings.HasPrefix("http://aaa.jpg", "http:/"))
	fmt.Println(strings.HasSuffix("http://aaa.jpg", "a.jpg"))

}
