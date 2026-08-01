package main

import "fmt"

func main() {
	//	数组 和 切片 不是 切丝儿

	hen1 := 3.0
	hen2 := 5.0
	hen3 := 1.0
	hen4 := 3.4
	hen5 := 2.0
	hen6 := 50.0

	totalWeight := hen1 + hen2 + hen3 + hen4 + hen5 + hen6
	avgWeight := fmt.Sprintf("%.2f", totalWeight/6)

	fmt.Printf("totalWeight=%v avgWeight=%v", totalWeight, avgWeight)

	//	在go中，数组是值类型
	//	在go中，数组是值类型
	//	在go中，数组是值类型

	// 1. 定义一个数组
	var hens [6]float64

	// 2. 给数组的每个元素赋值（下标从0开始，0~5）
	hens[0] = 3.0 // 第1只鸡
	hens[1] = 5.0 // 第2只鸡
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	// 3. 遍历数组求出总体重
	var totalWeight2 float64
	for i := 0; i < len(hens); i++ {
		totalWeight2 += hens[i]
	}
	fmt.Println()

	// 4. 求平均体重
	avgWeight2 := totalWeight2 / float64(len(hens))

	fmt.Printf("totalWeight=%.2f avgWeight=%.2f\n", totalWeight2, avgWeight2)

}
