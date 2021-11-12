package main

import f "fmt"

func hello() {
	f.Println("Hello World")
}

func main() {
	f.Println("goroutine 경량 실행 스레드")
	go hello()
	f.Scanln()
}
