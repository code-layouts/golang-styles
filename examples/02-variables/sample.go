package main

import (
	f "fmt"
	"time"
)

type ByteSize uint64

var (
	i int
	b bool
	s string
)

var name, title, num1, num2 = "symple", "MR", 1, 2

const NICKNAME string = "symplesims"

const (
	GO = iota // 여러 상수의 값을 0부터 1씩 증가시킨다.
	JAVA
	PYTHON
	C
)

const (
	_           = iota             // 초기값이 0이기 때문에 버린다.
	KB ByteSize = 1 << (10 * iota) // << 연산자는 비트를 이동 시킨다. 본 문법에는 1을 왼쪽으로 10회 이동 한다.  2의 10승은 1024가 된다.
	MB
	GB
	TB
	PB
	EB
)

func main() {
	// 변수 선언과 동시에 값을 할당 (축약형)
	firstName := "apple"
	birthYear := 1976
	f.Println(firstName, birthYear)

	f.Println("실제 메모리가 할당된 주소는 &변수명을 통해 확인 가능 합니다. (예: &firstName)")
	f.Println("firstName 메모리주소: ", &firstName, ", birthYear 메모리주소: ", &birthYear)

	i, b, s = 1, true, "sample"
	f.Println(i, b, s)
	f.Println("Nickname : ", NICKNAME)
	f.Println(name, title, num1, num2)
	f.Println(GO, JAVA, PYTHON, C)
	f.Println(KB, MB, GB, TB, PB, EB)

	var job string = "Smith"
	var age uint32 = 30
	f.Println(job, age)
	f.Println()
	f.Println("Constants iterator example -----")
	var Weekdays = []time.Weekday{
		time.Sunday,
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	}
	for _, day := range Weekdays {
		f.Println("Weekdays:", day)
	}
	f.Println()
	f.Println("Enum example -----")

	en := Beta
	switch en {
	case Alpha:
		f.Println("Alpha")
	case Beta:
		f.Println("Beta")
	case Gamma:
		f.Println("Gamma")
	case Delta:
		f.Println("Delta")
	case Epsilon:
		f.Println("Epsilon")
	case Zeta:
		f.Println("Zeta")
	default:
		f.Println("Eta")
	}

	f.Println()
	f.Println("Custom Enum example -----")
	enumString := EnumString{Delta}
	f.Println("CHK1 EnumString:", enumString)
	f.Println("CHK2 EnumString:", enumString.X)
	switch enumString.X {
	case Alpha:
		f.Println("CHK3 EnumString: Alpha")
	case Delta:
		f.Println("CHK3 EnumString: Delta")
	}

	enumString = EnumString{Beta}
	f.Println(enumString.X)
	enumString = EnumString{Epsilon}
	f.Println(enumString.X)

}

type Enum int

const (
	Alpha Enum = iota
	Beta
	Gamma
	Delta
	Epsilon
	Zeta
	Eta
)

type EnumString struct {
	X Enum
}

func (en Enum) String() string {
	f.Println("call enum-string function....")
	switch en {
	case Alpha:
		return "Alpha"
	case Beta:
		return "Beta"
	case Delta:
		return "Delta"
	case Eta:
		return "Eta"
	default:
		return "Zeta"
	}
}
