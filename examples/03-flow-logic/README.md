# flow logic
go-lang 조건, 분기 같은 흐름을 제어하는 예제 입니다.

## Examples

if 조건문
```
if num <= 10 {
    f.Println("10 보다 작거나 같다")
} else if num > 10 && num < 100 {
    f.Println("10 보다 크고 100 보다 작습니다.")
} else {
    f.Println("잘못 입력 했습니다.")
}
```

switch 다중 조건문
```
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
```

## Build
```
go build -o target/sample sample.go
```

## Run
```
./target/sample
```