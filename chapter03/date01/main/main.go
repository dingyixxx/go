package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("类型：%T 值：%v\n", now, now) //time.Time

	fmt.Println("------")
	fmt.Println(now.Format("2006/01/02 15:04:05")) //"2006/01/02 15:04:05" go语言之父 突然想创作一门语言 这个想法 诞生的时刻
	fmt.Println(now.Format("2026/07/31 20:06:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2026-07-31"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Format("20:06:05"))
	fmt.Println("------")
	fmt.Println(now.Format("2006"))
	fmt.Println(now.Format("01"))
	fmt.Println(now.Format("02"))
	fmt.Println(now.Format("15"))
	fmt.Println(now.Format("04"))
	fmt.Println(now.Format("05"))
	fmt.Println("------")

	fmt.Println(now.Year())         //2026
	fmt.Println(now.Month())        //July
	fmt.Println(int(now.Month()))   //7
	fmt.Println(now.Weekday())      //Friday
	fmt.Println(int(now.Weekday())) //5
	fmt.Println(now.Day())          //31
	fmt.Println(now.Date())         //2026 July 31
	fmt.Println(now.Hour())         //19
	fmt.Println(now.Minute())       //42
	fmt.Println(now.Second())       //48

}
