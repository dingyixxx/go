package model

type Student struct {
	Name  string
	Score float64
}

// person这个结构体，只能在model包使用
type person struct {
	Name  string
	score float64
}

func NewPerson(n string, s float64) *person {
	return &person{
		Name:  n,
		score: s,
	}
}

func (p person) GetScore() float64 {
	return p.score
}
