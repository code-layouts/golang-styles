# channel
채널(channel)은 데이타를 주고 받는 통로라 볼 수 있는데, make() 함수를 통해 채널을 미리 생성되어야 하며, 채널 연산자'<-'를 통해 데이타를 주고 받습니다.  
채널은 대게 go-routine들 사이 데이타를 주고 받는데 이용되는데, 상대편이 준비될 때까지 채널에서 대기함으로써 별도의 lock을 걸지 않고 데이타를 동기화하는데 사용 됩니다.


```
func main() {
  ch := make(chan int) // 정수형 채널을 생성한다
 
  // 별도의 goroutine 을 통해 채널에 123을 보낸다
  go func() {
    ch <- 123  
  }()
 
  // goroutine 의 채널로부터 123을 받는다
  i := <- ch  
  println(i)
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