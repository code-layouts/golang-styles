# go routine
고루틴 (goroutine)은 가벼운 스레드와 같은 것으로 수행 흐름과 별개로 병렬 처리가 가능하게 합니다.

**동시(Concurrent) 처리:** 싱글 코어에서 멀티 스레드를 동작시키는 개념으로 한번에 여러개가 동시에 실행되는 것 처럼 보이게 되지만, 실제로는 연산을 위한 context-switching 이 발생 합니다.  
**병렬(Parallel) 처리:** 멀티 코어에서 각각의 코어가 독립적으로 동시에 실행 됩니다. (멀티코어 멀티스레드 처리 가능)  

참고로 현재 HOST 의 모든 CPU를 사용 하려면 **runtime.GOMAXPROCS(runtime.NumCPU())** 함수를 호출 하여 시스템의 모든 코어를 사용하도록 설정할 수 있습니다.


```
/* go 의 경량 스레드를 처리하려면 함수 호출을 go 와 함게 실행 하면 됩니다. */
func func main() {
    go hello()
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