# hello world
Hello, World 예제 입니다.

main() 함수에서 시작 하며, [fmt](https://golang.org/src/fmt/print.go) 기본 패키지를 통해 콘솔에 메시지를 출력 합니다.

```
import "fmt"

func main() {
	fmt.Printf("hello %s", "world '")
}
```

## Build
```
go build -o target/hello hello.go
```

## Run
```
./target/hello
```

## Reference
- [Getting-Started](https://golang.org/doc/tutorial/getting-started)