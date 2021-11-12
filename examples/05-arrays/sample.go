package main

import f "fmt"

func array01() {
	f.Println("크기가 5인 배열을 생성하고 2번째 값을 5로 초기화 한다.")
	var v [5]int
	v[2] = 5
	f.Println(v)
}

func array02() {
	f.Println("정수형의 크기가 3인 배열을 선언하고 값을 초기화 한다.")
	var v [3]int = [3]int{1, 2, 3}
	f.Println(v)
}

func array03() {
	f.Println("문자열의 크기가 3인 배열을 선언하고 값을 초기화 한다.")
	var v = [3]string{"name", "age", "weight"}
	f.Println(v)
}

func array04() {
	f.Println("실수형의 크기가 5인 배열을 선언하고 값을 초기화 한다.")
	var v = [5]float32{
		1.14,
		2.14,
		3.14,
		4.14,
		5.14,
	}
	f.Println(v)
}

func array05() {
	f.Println("동적 배열을 선언하고 값을 할당 합니다.")
	v := [...]string{"apple", "lg", "samsung", "kt"}
	f.Println(v)
}

func array06() {
	f.Println("다 차원(중첩) 배열을 선언하고 값을 할당 하고 활용 합니다.")
	var v = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	f.Println(v)

	for _, array := range v {
		for _, value := range array {
			f.Println("array:", array, ", value:", value)
		}
	}

	f.Println("Array size", len(v))
}

func array07() {
	f.Println("배열을 copy 입니다.")
	original := [...]string{"ios", "linux", "unix", "android", "windows"}
	clone := original

	original[3] = "ubuntu"
	f.Println("original: ", original)
	f.Println("clone 된 객체는 레퍼런스 하지 않는 유일한 객체로 생성 됩니다.")
	f.Println("clone: ", clone)
}

func main() {
	array01()
	array02()
	array03()
	array04()
	array05()
	array06()
	array07()
}
