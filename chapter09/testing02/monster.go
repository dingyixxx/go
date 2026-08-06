package testing02

import (
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) Store() bool {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
		return false
	}
	filePath := "monster.dat"
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Printf("写入文件错误 err=%v\n", err)
		return false
	}
	return true
}

func (m *Monster) ReStore() bool {
	filePath := "monster.dat"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("读取文件错误 err=%v\n", err)
		return false
	}
	err = json.Unmarshal(data, m)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
		return false
	}
	return true
}
