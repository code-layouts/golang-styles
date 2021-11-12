package main

import (
	f "fmt"
)

func hello() {
	f.Println("Hello")
}

func world(value string) {
	f.Println("Hello,", value, "!")
}

func division(a int, b int) (float32, int) {
	value := float32(a) / float32(b)
	remainder := a % b
	return value, remainder
}

func sliceSum(n ...int) int {
	result := 0
	for _, value := range n {
		result += value
	}
	return result
}

func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}

func calc(f func(int, int) int, a int, b int) int {
	return f(a, b)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func closer() {
	f.Println("클로저 함수 예시 입니다. (함수 내의 지역변수 값을 유지 할 수 있습니다.)")
	nextInt := intSeq()
	f.Println(nextInt())
	f.Println(nextInt())
	f.Println(nextInt())

	newInts := intSeq()
	f.Println(newInts())
}

func deferPrintName() {
	defer func() {
		f.Println("fruit")
	}()

	func() {
		f.Println("apple")
	}()

	array := []int{1, 2, 3, 4, 5}
	for _, value := range array {
		defer f.Println("defer value: ", value)
		f.Println("value: ", value)
	}
}

type Paper struct {
	Line string
}

/**
  Paper 스트럭쳐는 sketch 메서드를 가진다.
*/
func (paper *Paper) sketch(line string) string {
	return "Draw line " + line
}

func anonymous() {

	func() {
		f.Println("Welcome! to GeeksforGeeks")
	}()

	p1 := "anonymous"
	p2 := "function"
	result := func(a string, b string) string {
		return a + " " + b + "!!!"
	}(p1, p2)
	f.Println(result)

}

func main() {
	hello()
	world("World")
	var value, remainder = division(7777, 256)
	f.Println("7777 / 256 에 대한 value: ", value, ", remainder: ", remainder)
	f.Println("동적 아규먼트 함수 전달 예시 입니다.")
	f.Println("sum(1 to 5): ", sliceSum(1, 2, 3, 4, 5))
	p1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f.Println("sum(1 to 10): ", sliceSum(p1...))

	f.Println("고차 함수(First-Class) 예시 입니다.")
	f.Println("1 + 2 =", calc(plus, 1, 2))
	f.Println("10 - 2 =", calc(minus, 10, 2))
	f.Println("512 * 2 =", calc(multiply, 512, 2))
	f.Println("512 / 2 =", calc(divide, 512, 2))
	f.Println("팩토리얼 재귀 함수 예시 fact(5): ", fact(5))
	closer()

	f.Println("지연 호출 예시 입니다.")
	deferPrintName()

	f.Println()
	f.Println("struct 를 위한 함수 사용 예시 -----")
	p := &Paper{}
	f.Println(p.sketch("straight"))
	f.Println(p.sketch("curve"))

	f.Println()
	f.Println("anonymous 함수 사용 예시 -----")
	anonymous()

}
