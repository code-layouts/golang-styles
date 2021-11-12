package main

import (
	f "fmt"
	"time"
)

// 조건문 예시
func ifcondition(num int) {
	if num <= 10 {
		f.Println("10 보다 작거나 같다")
	} else {
		f.Println("잘못 입력 했습니다.")
	}

}

// switch 조건문 예시
func switchCondition(num int) {

	switch num {
	case 1:
		f.Println("입력된 값: 1")
	case 2:
		f.Println("입력된 값: 2")
	case 3:
		f.Println("입력된 값: 3")
	default:
		f.Println("입력된 값: ", num)
	}

	switch num {
	case 1:
		f.Println("fallthrough 입력된 값: 1")
		fallthrough
	case 2:
		f.Println("fallthrough 입력된 값: 2")
		fallthrough
	case 3:
		f.Println("fallthrough 입력된 값: 3")
		fallthrough
	default:
		f.Println("fallthrough 입력된 값: ", num)
	}
}

func week() string {
	var week string
	currentTime := time.Now()
	f.Println("현재 시간 : ", currentTime)
	weekday := time.Now().Weekday()
	switch weekday {
	case time.Monday:
		week = "월요일"
	case time.Tuesday:
		week = "화요일"
	case time.Wednesday:
		week = "수요일"
	case time.Thursday:
		week = "목요일"
	case time.Friday:
		week = "금요일"
	case time.Saturday:
		week = "토요일"
	default:
		week = "일요일"
	}
	return week
}

func main() {
	var num int
	f.Println("1 ~ 10까지의 정수 하나를 입력")
	f.Scanln(&num)
	ifcondition(num)
	switchCondition(num)
	week := week()
	f.Println("요일 : ", week)
}
