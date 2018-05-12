package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main 함수시작")

	go long()
	go short()
	time.Sleep(5 * time.Second)
	fmt.Println("main 끝")
}

func short(){
	fmt.Println("short 시작")
	time.Sleep(2 * time.Second)
	fmt.Println("short 끝")
}

func long(){
	fmt.Println("long 시작")
	time.Sleep(3 * time.Second)
	fmt.Println("long 끝")
}