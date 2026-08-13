package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的内容（序列化后的JSON）
}

//定义两个消息..后面需要再增加

type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

type LoginResMes struct {
	Code   int    `json:"code"`   // 返回状态码 500 表示该用户未注册 200表示登录成功
	UserId []int  `json:"userId"` // 增加字段，保存用户id的切片
	Error  string `json:"error"`  // 返回错误信息
}

type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体.
}
type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回状态码 400 表示该用户已经占有 200表示注册成功
	Error string `json:"error"` // 返回错误信息
}
