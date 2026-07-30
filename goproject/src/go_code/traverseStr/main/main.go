package main

func main() {
	str := "hello World!上海辽阳"

	//这么遍历,是按照字节遍历的,一个汉字占3个字节,所以会乱码
	for i := 0; i < len(str); i++ {
		//fmt.Printf("i=%d val=%c \n", i, str[i])
		//fmt.Println(i, str[i])
		//i=0 val=h
		//i=1 val=e
		//i=2 val=l
		//i=3 val=l
		//i=4 val=o
		//i=5 val=
		//	i=6 val=W
		//i=7 val=o
		//i=8 val=r
		//i=9 val=l
		//i=10 val=d
		//i=11 val=!
		//	i=12 val=ä
		//i=13 val=¸
		//i=14 val=
		//i=15 val=æ
		//i=16 val=µ
		//i=17 val=·
		//i=18 val=è
		//i=19 val=¾
		//i=20 val=½
		//i=21 val=é
		//i=22 val=
		//i=23 val=³
	}

	//for-range遍历字符串和数组(数组后面会讲,此时先说字符串)
	//for-range遍历字符串时,是按照字符来遍历的,而不是按照字节来的,请注意这点
	//这么遍历,是按照字符遍历的,因此,虽然一个汉字占3个字节,但是不会乱码
	//for idx, letter := range str {
	//	fmt.Printf("%d %c\n", idx, letter)
	//	//fmt.Println(idx, letter)
	//	//0 104
	//	//1 101
	//	//2 108
	//	//3 108
	//	//4 111
	//	//5 32
	//	//6 87
	//	//7 111
	//	//8 114
	//	//9 108
	//	//10 100
	//	//11 33
	//	//0 h
	//	//1 e
	//	//2 l
	//	//3 l
	//	//4 o
	//	//5
	//	//6 W
	//	//7 o
	//	//8 r
	//	//9 l
	//	//10 d
	//	//11 !
	//	//	12 上
	//	//15 海
	//	//18 辽
	//	//21 阳
	//}

	//转换方式 底层类型 含义
	//[]byte(str) []uint8 将字符串按 UTF-8 字节 拆分
	//[]rune(str) []int32 将字符串按 Unicode 字符 拆分
	//因此,只要带了中文,就要用Unicode来转
	//	用[]rune也不会乱码 神奇
	//str2 := []rune(str)
	//for i := 0; i < len(str2); i++ {
	//	fmt.Printf("i=%d val=%c \n", i, str2[i])
	//	//i=0 val=h
	//	//i=1 val=e
	//	//i=2 val=l
	//	//i=3 val=l
	//	//i=4 val=o
	//	//i=5 val=
	//	//	i=6 val=W
	//	//i=7 val=o
	//	//i=8 val=r
	//	//i=9 val=l
	//	//i=10 val=d
	//	//i=11 val=!
	//	//	i=12 val=上
	//	//i=13 val=海
	//	//i=14 val=辽
	//	//i=15 val=阳
	//}

}
