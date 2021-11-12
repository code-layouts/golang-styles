package main

import f "fmt"

func slice01() {
	var v []int
	f.Println("슬라이스는 배열과 달리 [] 안에 길이를 지정하지 않습니다.", v, ", size: ", len(v))
	var v1 []int = make([]int, 5)
	f.Println("make 함수로 int형에 길이가 5인 슬라이스에 공간을 할당 합니다.", v1, ", size: ", len(v1))
	var v2 = make([]string, 5)
	f.Println("make 함수로 string형에 길이가 5인 슬라이스에 공간을 할당 합니다.", v2, ", size: ", len(v2))
	v3 := make([]float64, 5, 200)
	f.Println("make 함수로 float64형에 길이가 5인 슬라이스에 공간을 할당 합니다.", v3, ", size: ", len(v3))

}

func slice02() {
	f.Println("슬라이스에 값을 추가로 지정하고 싶다면 append 함수를 사용합니다.")
	v := []string{"windows", "unix", "linux"}
	v = append(v, "ubuntu")
	v = append(v, "centos")
	v = append(v, "fedora")
	f.Println(v)
	f.Println("길이 len: ", len(v), " , 용량 cap: ", cap(v))
	v = append(v, "redhat")
	f.Println(v)
	f.Println("슬라이스의 용량은 배열 크기 대비 미리 정적(대게 배수)으로 할당하므로 메모리 및 성능에 영향을 줍니다.")
	f.Println("길이 len: ", len(v), " , 용량 cap: ", cap(v))

	f.Println("슬라이스에 또다른 슬라이스를 추가할 수 있습니다.")
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	a = append(a, b...)
	a = append(a, []int{7, 8, 9}...)
	f.Println(a)
}

func map01() {
	f.Println("string 타입의 key 와 int 타입의 value 를 선언하는 map 입니다.")
	v := make(map[string]int)
	v["age"] = 27
	v["height"] = 175
	f.Println("age: ", v["age"])
	f.Println("height: ", v["height"])
	f.Println("value: ", v, "len: ", len(v))
}

func map02() {
	f.Println("map 의 값의 컨트롤 및 순회")
	capacityUnit := make(map[string]string)

	capacityUnit["1byte"] = "8 bit"
	capacityUnit["1KB"] = "1024 Byte"
	capacityUnit["1MB"] = "1024 KB"
	capacityUnit["1GB"] = "1024 MB"
	capacityUnit["1TB"] = "1024 GB"
	capacityUnit["1PB"] = "1024 TB"
	f.Println(capacityUnit, ", size: ", len(capacityUnit))

	value, result := capacityUnit["1YB"]
	f.Println("key '1YB' 가 존재하지 않으므로 === ", value, result)
	value, result = capacityUnit["1GB"]
	f.Println("key '1GB' 가 존재 하므로 ========= ", value, result)

	f.Println("map 객체 loop 예시 ==============")
	for key, result := range capacityUnit {
		f.Println(key, " : ", result)
	}

	f.Println("map 객체의 key 삭제 ==============")
	delete(capacityUnit, "1PB")
	f.Println(capacityUnit, ", size: ", len(capacityUnit))
}

func map03() {
	f.Println("Complex map 의 사용 예시")
	v := map[string]map[string]float32{
		"Mercury": map[string]float32{
			"meanRadius":    2439.7,
			"mass":          3.3022e+23,
			"orbitalPeriod": 87.969,
		},
		"Venus": map[string]float32{
			"meanRadius":    6051.8,
			"mass":          4.8676e+24,
			"orbitalPeriod": 224.70069,
		},
		"Earth": map[string]float32{
			"meanRadius":    6371.0,
			"mass":          5.97219e+24,
			"orbitalPeriod": 365.25641,
		},
		"Mars": map[string]float32{
			"meanRadius":    3389.5,
			"mass":          6.4185e+23,
			"orbitalPeriod": 686.9600,
		},
	}

	f.Println(v["Earth"]["orbitalPeriod"])

}

func main() {
	f.Println("slice -----------------------------------------")
	slice01()
	slice02()
	f.Println("map -----------------------------------------")
	map01()
	map02()
	map03()
}
