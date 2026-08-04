package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 方法
// 1.存款 都传入指针，记死了
func (acc *Account) Deposit(money float64, pwd string) {
	if pwd != acc.Pwd {
		fmt.Println("您输入的密码不正确")
		return
	}
	if money < 0 {
		fmt.Println("您输入的金额不正确")
		return
	}
	acc.Balance += money
	fmt.Println("存款成功")
}

// 2.取款 都传入指针，记死了
func (acc *Account) Withdraw(money float64, pwd string) {
	if pwd != acc.Pwd {
		fmt.Println("您输入的密码不正确")
		return
	}
	if money < 0 {
		fmt.Println("您输入的金额不正确")
		return
	}
	if money > acc.Balance {
		fmt.Println("您输入的金额大于了余额")
		return
	}
	acc.Balance -= money
	fmt.Println("取款成功")
}

// 3.余额
func (acc Account) Query(pwd string) {
	if pwd != acc.Pwd {
		fmt.Println("您输入的密码不正确")
		return
	}

	fmt.Printf("你的账号为%v，余额为%v", acc.AccountNo, acc.Balance)
}
func main() {
	account := Account{
		AccountNo: "dingyixxx的xx银行编号gs6226091210562496",
		Pwd:       "666666",
		Balance:   100.89,
	}
	account.Query("666666")
	account.Deposit(29.01, "666666")
	account.Query("666666")
	account.Withdraw(0.99, "666666")
	account.Query("666666")

}
