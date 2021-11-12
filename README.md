# golang-styles
golang 을 처음 시작하기위한 가이드 문서 입니다.


## Install
macos 에 go 를 설치 합니다.
```shell
brew install go
go env
go version
```

## Project-layout
golang 프로젝트 레이아웃을 참고 합니다.
[project-layout](https://github.com/golang-standards/project-layout)

## Project-init
golang-styles 프로젝트를 checkout 합니다.
```
git clone https://github.com/code-layouts/golang-styles.git
cd golang-style
```

## Examples
[examples](examples/README.md) 를 통해 기초적인 hello world 부터 다양한 예시를 보고 이해 합니다.

## go programming guide
- https://pronist.tistory.com/96?category=889834#reflect%--%EB%A-%-E%EC%-D%B-%--%EC%--%AC%EC%-A%A-%ED%--%--%EA%B-%B-%--%EC%--%-A%EA%B-%B-
- https://laeshiny.medium.com/go-code-review-comments-%EC%A0%95%EB%A6%AC-47d05fdb49f6

panic 은 피하고 error 는 확실하게 처리 하기
-----
panic 은 서비스가 중단되는 치명적인 애러 이므로 피해야 한다. 대표적으로 객체가 생성되지 않았는데 사용한다던지,
잘못된 배열 인덱스를 사용한다던지 프로그래머가 놓치는 휴먼 애러일 것이다.

error 는 오류가 발생하더라도 서비스가 중단되지 않는다. 대게 예외에 대한 로그를 error 로 기록 할 수 있다.
```
func hello() {
    _, err := os.Open("Unknown.txt")
    if err != nil {
        log.Fatal(err)
    }
}
```