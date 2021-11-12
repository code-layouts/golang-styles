# variables
go-lang 에서 변수 및 상수를 사용하는 예시 입니다.

## Examples
```
// 변수 선언
var fruit string = "apple"

// 변수 선언과 동시에 값을 할당 (축약형)
firstName := "apple"
birthYear := 1976

// 다중 변수 선언 및 값의 할당 
var (
	i int
	b bool
	s string
)
i, b, s = 1, true, "sample"

// 상수 선언
const NICKNAME string = "symplesims"

// 다중 상수 선언 및 iota 활용
const (
	GO = iota // 여러 상수의 값을 0부터 1씩 증가시킨다.
	JAVA
	PYTHON
	C
)
fmt.Println(GO, JAVA, PYTHON, C)

```

## Build
```
go build -o target/sample sample.go
```

## Run
```
./target/sample
```

## Reference
- [Getting-Started](https://golang.org/doc/tutorial/getting-started)