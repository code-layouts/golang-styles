# Functions

GoLang 에서 함수를 사용하는 예시 입니다.

Examples
-------------------
기본 함수

```
func hello() {
    f.Println("Hello")
}
```

파라미터를 받는 함수

```
func world(value string) {
    f.Println("Hello,", value, "!")
}
```

리턴 타입이 있는 함수

```
func plus(a int, b int) int {
    return a + b
}
```

동적 파라미터를 받는 함수

```
func sliceSum(n ...int) int {
    result := 0
    for _, value := range n {
        result += value
    }
    return result
}
```

파라미터로 함수를 받는 고차 함수

```
func calc(f func(int, int) int, a int, b int) int {
    return f(a, b)
}

var result int = calc(func(a int, b int) int { return a + b    }, 1, 2)
f.Println( result )
```

재귀 함수

```
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
f.Println( fact(5) )
```

지연 호출 함수 defer

```
func main() {
    defer func() {
        f.Println("fruit")
    }()

    func() {
        f.Println("apple")
    }()
}
```

struct 타임 전용 함수

```
type Person struct {
    Message string
}

func (person *Person) message(message string) *Person {
    person.Message = message
    return person
}

person := Person{}
person.message("hello")
f.Println(person)
```

## Build

```
go build -o target/sample sample.go
```

## Run

```
./target/sample
```