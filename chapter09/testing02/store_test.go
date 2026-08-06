package testing02

import "testing"

func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "牛魔王",
		Age:   500,
		Skill: "牛魔拳",
	}

	ok := monster.Store()
	if !ok {
		t.Fatalf("Store() 执行错误")
	}
	t.Logf("Store() 执行正确")

	monster2 := &Monster{}
	ok = monster2.ReStore()
	if !ok {
		t.Fatalf("ReStore() 执行错误")
	}

	if monster2.Name != "牛魔王" {
		t.Fatalf("ReStore() 错误，期望 Name=牛魔王，实际=%v", monster2.Name)
	}
	if monster2.Age != 500 {
		t.Fatalf("ReStore() 错误，期望 Age=500，实际=%v", monster2.Age)
	}
	if monster2.Skill != "牛魔拳" {
		t.Fatalf("ReStore() 错误，期望 Skill=牛魔拳，实际=%v", monster2.Skill)
	}

	t.Logf("ReStore() 执行正确，Name=%v Age=%v Skill=%v", monster2.Name, monster2.Age, monster2.Skill)
}
