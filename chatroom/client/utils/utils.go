package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

// 这里将这些方法关联到结构体中
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte //这时传输时，使用缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	read, err := this.Conn.Read(this.Buf[:4])
	fmt.Println("读取客户端发送的数据..", read) //从conn读取字节到buf中

	if err != nil {
		fmt.Println("conn.Read err=", err)
		err = errors.New("读长度出错了")
		return
	}
	//根据buf[:4] 转成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	fmt.Println("读到的buf=", this.Buf[:4])

	//根据pkgLen读取消息内容
	//从conn读pkgLen个字节,扔到buf
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	fmt.Println("n=", n)
	if n != int(pkgLen) || err != nil {
		err = errors.New("读内容出错了")
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	// 发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	//	发送data本身
	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return

}
