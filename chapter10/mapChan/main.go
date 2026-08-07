package main

func main() {
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string)
	m1 := make(map[string]string, 10)
	m1["city1"] = "北京"
	m1["city2"] = "上海"

	m2 := make(map[string]string, 10)
	m2["hero1"] = "宋江"
	m2["hero2"] = "武松"

	mapChan <- m1
	mapChan <- m2

}
