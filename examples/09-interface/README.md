# interface
go-lang 에서 인터페이스를 활용하는 에시 예시 입니다.

type geometry interface {
area() float64
perim() float64
}

```
/* 인터페이스 정의 */
type drawable interface {
    draw() int
}

/* person 객체에 대해 drawable 인터페이스 구현 */
func draw(p person) int {
	x := p.xxx
    return x
}

/* airplane 객체에 대해 drawable 인터페이스 구현 */
func draw(a airplane) int {
	x := a.xxx
    return x
}
```

## Build
```
go build -o target/sample sample.go
```

## Run
```
./target/sample
```