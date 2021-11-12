package main

import (
	f "fmt"
	"time"
)

func basic() {
	ch := make(chan int, 3)

	func() {
		// 채널에 메시지를 수신하기 전 반드시 channel 을 닫아 주어야 한다. defer 호출 활용
		defer close(ch)
		ch <- 1001
		ch <- 1002
		ch <- 1003
	}()

	for i := range ch {
		time.Sleep(time.Millisecond * 50)
		f.Println(i)
	}

}

func run(ch *chan bool, interval time.Duration) {
	time.Sleep(time.Second * interval)
	*ch <- true
}

func main() {
	f.Println("Channel Select 예제 입니다.")
	basic()

	f.Println()
	f.Println()
	f.Println("메시지가 수신된 채널들을 선택 합니다.")

	f.Println("1. Channel 등록")
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	defer close(ch1)
	defer close(ch2)
	defer close(ch3)

	f.Println("2. 경량 스레드 실행")
	go run(&ch1, 3)
	go run(&ch2, 2)
	go run(&ch3, 1)

	f.Println("3. 채널 셀렉트 대기")

SUB:
	for {
		select {
		case <-ch1:
			println("ch1 완료")
			break SUB
		case <-ch2:
			println("ch2 완료")
		case <-ch3:
			println("ch3 완료")
		}
	}

	f.Println()
	f.Println("선착순으로 처음 메시지가 수신된 채널만 선택합니다.")

	f.Println("2.1. Channel 등록")
	ch11 := make(chan bool)
	ch22 := make(chan bool)
	ch33 := make(chan bool)

	defer close(ch11)
	defer close(ch22)
	defer close(ch33)

	f.Println("2.2. 경량 스레드 실행")
	go run(&ch11, 3)
	go run(&ch22, 1)
	go run(&ch33, 2)

	f.Println("2.3. 채널 셀렉트 대기")
EXIT:
	for {
		select {
		case <-ch11:
			println("ch11 완료")
			break EXIT
		case <-ch22:
			println("ch22 완료")
			break EXIT
		case <-ch33:
			println("ch33 완료")
			break EXIT
		}
	}

	f.Println()
	f.Println("Channel 종료")
}
