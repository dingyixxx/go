package main

import (
	"log"
	"time"
)

func main() {
	//log.Fatalln("Fatal Level: log entry") //app exits here
	//log.Panic("log.Panic")
	log.Println("Normal Level: log entry")
	time.Sleep(time.Second * 5)
}
