# struct
go-lang 에서 커스텀 자료구조를 정의하는 struct 에 대한 예시 입니다.

struct 를 정의할 때 맴버 변수의 첫 문자가 대문자인 경우 public 을 의미 합니다.  
만약 소문자로 지정하게되면 reflect 등을 사요할 때 참조를 할 수 없게 됩니다.

```
/* struct 구조체의 멤버 변수 이름은 대문자로 시작 */
type subject struct {
	Os       float32
	Database float32
	Network  float32
	Coding   float32
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