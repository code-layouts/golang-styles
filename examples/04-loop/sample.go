package main

import f "fmt"

func multipleTable(number int) {
	f.Println("Multiplication table")
	for i := 2; i <= number; i++ {
		f.Println("Multiplication ", i, " ============")
		for j := 1; j < 10; j++ {
			f.Println(i, " * ", j, " = ", i*j)
		}
	}
}

func loopBreak() {
	f.Println("0 to 5 까지 loop 하지만 3 에서 break ")
	for i := 0; i < 5; i++ {
		if i == 3 {
			f.Println("3이 되었습니다.")
			break
		}
	}
}

func loopContinue() {
	f.Println("0 to 5 의 중첩 배열을 loop 하지만 j 값 3 에서 continue (3을 건너뜀) ")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				continue
			}
			f.Println("i:", i, " , j:", j)
		}
	}
}

func main() {
	var number int
	f.Println("2 ~ 9까지의 정수 하나를 입력 하세요")
	f.Scanln(&number)
	f.Println("-----------------------------------------")
	multipleTable(number)
	f.Println("-----------------------------------------")
	loopBreak()
	f.Println("-----------------------------------------")
	loopContinue()

}
