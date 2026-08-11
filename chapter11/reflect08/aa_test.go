package reflect08

import (
	"reflect"
	"testing"
)

type user struct {
	UserId string
	Name   string
}

func TestReflectStructPtr(t *testing.T) {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)
	st = reflect.TypeOf(model)                       //获取类型*user
	t.Log("reflect.TypeOf", st.Kind().String())      // *reflect08.user的大类是ptr
	st = st.Elem()                                   //st指向的类型，解引用，和之前那个set改变Student和int一样
	t.Log("reflect.TypeOf.Elem", st.Kind().String()) //struct

	elem = reflect.New(st)
	//st此时此刻是一个纯纯的user结构体了，不是指针了
	//New返回一个Value类型值，该值持有一个指向类型为typ的新申请的零值的指针

	t.Log("reflect.New", elem.Kind().String())             // ptr
	t.Log("reflect.New.Elem", elem.Elem().Kind().String()) //struct
	//model就是创建的user结构体变量(实例)
	model = elem.Interface().(*user) //model = "user 它的指向和elem是一样的 model是个指针
	elem = elem.Elem()               //解指针然后设置值
	t.Logf("model 地址: %p\n", model)
	t.Logf("elem 地址: %x\n", elem.UnsafeAddr())

	//取得elem指向的值
	elem.FieldByName("UserId").SetString("12345678") //赋值..
	elem.FieldByName("Name").SetString("nickname")
	t.Log("model model.Name", model, model.Name)
}
