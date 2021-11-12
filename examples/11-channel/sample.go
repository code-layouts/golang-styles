package main

import (
	f "fmt"
	"time"
)

func integerChannel() {
	ch := make(chan int)
	go func() {
		ch <- 123
	}()

	i := <-ch
	println(i)
}

func wait() {
	done := make(chan bool)
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(time.Second)
			f.Println(i, "초를 대기 합니다.")
		}
		done <- true
	}()
	<-done
	f.Println("wait 프로세스가 done 처리 되었습니다.")
}

func main() {
	f.Println("Channel 예제 입니다.")
	integerChannel()
	wait()
}
