package main

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

func main() {
	//b的平方-4ac
	//（-b +/- (b的平方-4ac)的平方根）/2a
	a := decimal.NewFromFloat(1.0)
	fmt.Printf("%T\n", a) //decimal.Decimal
	b := decimal.NewFromFloat(6.0)
	c := decimal.NewFromFloat(4.0)

	m := b.Mul(b).Sub(a.Mul(c).Mul(decimal.NewFromInt(4)))
	//不推荐“把m转换为float64”，即，不推荐用m.Float64()，
	//因为，如果这么转，那就相当于在最后一个步骤丢失了精度
	//（你好不容易选用了decimal，但却在最后一步转成float64时，丢失了精度，那不是前功尽弃了吗）

	compare := m.Cmp(decimal.Zero)
	fmt.Printf("%T\n", compare) //int

	isZero := m.IsZero()
	fmt.Printf("%T\n", isZero) //bool

	lessThan := m.LessThan(decimal.Zero)
	fmt.Printf("%T\n", lessThan) //bool

	greaterThan := m.GreaterThan(decimal.Zero)
	fmt.Printf("%T\n", greaterThan) //bool

	//因为math.sqrt只接受float64，所以，也就不用math.sqrt了
	if m.LessThan(decimal.Zero) {
		fmt.Println("无解")
	} else {
		// decimal没有Sqrt...所以还是转回float64()
		//先转float64，开平方，再转回decimal
		f, exact := m.Float64()                  //第二个字段，代表精度是否有丢失
		fmt.Printf("类型=%T,值=%v\n", f, f)         //float64
		fmt.Printf("类型=%T,值=%v\n", exact, exact) //bool
		sqrtFloat := math.Sqrt(f)
		//但decimal好像不像java的BigDecimal有直接开根号的api，所以，还是要用math类
		sqrtFloatDecimal := decimal.NewFromFloat(sqrtFloat)
		negB := b.Neg()
		denominator := a.Mul(decimal.NewFromInt(2))
		if m.IsZero() {
			fmt.Println("有且仅有一个解")
			x := negB.DivRound(denominator, 8)
			fmt.Printf("有且仅有一个解,x=%v", x)
		} else {
			numerator1 := negB.Add(sqrtFloatDecimal)
			numerator2 := negB.Sub(sqrtFloatDecimal)
			x1 := numerator1.DivRound(denominator, 8)
			x2 := numerator2.DivRound(denominator, 8)
			fmt.Printf("有两个解,x1=%v,x2=%v", x1, x2)
		}
	}
}
