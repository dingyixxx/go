package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `dd:"monster_age"`
	Score float32
	Sex   string
}

func (s *Monster) Print() {
	fmt.Println("---start----")
	fmt.Println(s)
	fmt.Println("---end----")
}

func (s *Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s *Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	//获取到reflect.Type
	typ := reflect.TypeOf(a)
	//获取到reflect.Value
	val := reflect.ValueOf(a)
	//获取到a对应的大的类别kind
	kd := val.Kind()

	fmt.Printf("reflect.Ptr转为十进制是-%d\n", reflect.Ptr)
	//reflect.Ptr-22
	fmt.Printf("reflect.Ptr值是-%v\n", reflect.Ptr)
	//reflect.Ptr值是-ptr，重写了Kind的String方法，所以打出来ptr

	//传入的是结构体or指针
	if kd != reflect.Ptr {
		fmt.Println("expect Ptr")
		return
	}
	elem := val.Elem() // 解引用得到结构体

	num := elem.NumField()
	//NumField 其实是 所有的属性 而不是数值类型的属性

	fmt.Printf("struct has %d fields\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为%v\n", i, elem.Field(i)) //返回的是value

		//获取到struct标签，注意需要通过reflect.Type来获取tag标签的值
		structField := typ.Elem().Field(i)
		//reflect.StructField{}这个类型

		tagVal := structField.Tag.Get("dd") //json tag 为什么Marshal必须写json
		//tagVal := structField.Tag.Get("json") //json tag 为什么Marshal必须写json

		//返回的是字段的类型，后面跟着的json tag

		//这个类型的结构体的某个字段

		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	numOfMethod := val.NumMethod() //但是方法不需要“解指针”
	fmt.Printf("struct has %d methods\n", numOfMethod)
	//var params []reflect.Value
	val.Method(1).Call(nil) //Print方法 按照ascii码排序

	//调用结构体的第1个方法Method(0)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //GetSum方法
	//传入的参数是 []reflect.Value
	//应该是50
	//返回结果，返回的结果是 []reflect.Value

	//res[0]的值为50，类型为reflect.Value
	fmt.Printf("res[0]的值为%v，类型为%T\n", res[0], res[0])

	//res[0].Int()的值为50，类型为int64
	fmt.Printf("res[0].Int()的值为%v，类型为%T\n", res[0].Int(), res[0].Int())

	var params1 []reflect.Value
	params1 = append(params1, reflect.ValueOf("酒精"))
	params1 = append(params1, reflect.ValueOf(2))
	params1 = append(params1, reflect.ValueOf(float32(66.66)))
	params1 = append(params1, reflect.ValueOf("unknown"))

	val.Method(2).Call(params1)
	fmt.Println("res1=", elem.Interface())

}

func main() {
	//	使用“反射”来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值
	//	Method 按照方法名排序对应i值
	//	Call 传入参数和返回参数，是切片，[]reflect.Value
	var a *Monster = &Monster{
		Name:  "味精",
		Age:   400,
		Score: 30.8,
	}
	//为什么不传入指针呢
	TestStruct(a)
	fmt.Println(a)
}
