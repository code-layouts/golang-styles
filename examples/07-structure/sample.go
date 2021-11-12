package main

import (
	f "fmt"
	"reflect"
)

type person struct {
	name    string
	age     int
	hobby   []string
	enabled bool
}

func printPerson() {
	v := person{
		name:    "long",
		age:     22,
		hobby:   []string{"baseball", "fish", "netflex"},
		enabled: true,
	}
	f.Println("person : ", v)
	v2 := person{"josh", 44, []string{"game", "coding"}, true}
	f.Println("person2 : ", v2)
	f.Println("& 키워드를 변수 앞에 두면 메모리 레퍼런스를 가리킨다.")
	f.Println("v.age 인스턴스의 메모리 참조: ", &v.age, ", v2.age 인스턴스의 메모리 참조: ", &v2.age)

}

/**
 * struct 를 정의할 때 맴버 변수의 첫 문자가 대문자인 경우 public 을 의미한다.
 * 만약 소문자로 지정하게되면 reflect 등을 사요할 때 참조를 할 수 없게 된다.
 * avg 함수 참고
 */
type subject struct {
	Os       float32
	Database float32
	Network  float32
	Coding   float32
}

func total(stc subject) float32 {
	return stc.Os + stc.Database + stc.Network + stc.Coding
}

func avg(sub subject) float32 {
	e := reflect.ValueOf(&sub).Elem()
	var total float32 = 0
	length := e.NumField()
	for i := 0; i < length; i++ {
		// varName := e.Type().Field(i).Name
		// varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		// f.Println(varName, varType, varValue)
		total += varValue.(float32)
	}
	return total / float32(length)
}

func main() {
	f.Println("struct -----------------------------------------")
	printPerson()
	apple := subject{90, 80, 70, 90}
	f.Println(apple)

	f.Println()
	cpApple := apple
	cpApple.Os = 10
	cpApple.Coding = 10
	f.Println("cpApple 의 경우 별도의 인스턴스로 메모리에 할당 된다.")
	f.Println("CHK1 apple :", apple)
	f.Println("CHK1 cpApple :", cpApple)
	f.Println()

	refApple := &apple
	refApple.Os = 50
	refApple.Coding = 30
	f.Println("refApple 의 경우 기존 apple 인스턴스 메모리를 참조 한다.")
	f.Println("CHK2 apple :", apple)
	f.Println("CHK2 refApple :", refApple)

	f.Println()
	f.Println("TOTAL SCORE -----------------")
	apple = subject{100, 100, 80, 90}
	banana := subject{80, 70, 90, 90}
	f.Println("TOTAL SCORE apple :", total(apple))
	f.Println("TOTAL SCORE banana :", total(banana))

	f.Println()
	f.Println("AVERAGE SCORE -----------------")
	f.Println("TOTAL SCORE apple :", avg(apple))
	f.Println("TOTAL SCORE banana :", avg(banana))
}
