# channel select
채널 셀렉트(channel select)는 여러 채널로부터 메시지 수신을 대기 하면서 메시지가 수신된 채널이 있으면 해당 채널을 실행할 수 있습니다.


```
func main() {
  ch1 := make(chan bool)
  ch2 := make(chan bool)
  ch3 := make(chan bool)
 
  go run(ch1)
  go run(ch2)
  go run(ch3)
 
EXIT:
    for {
        select {
            case` <-ch1:
                println("ch1 완료")
            case <-ch2:
                println("ch2 완료")
            case <-ch3:
                println("ch3 완료")
                break EXIT
        }
    }
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