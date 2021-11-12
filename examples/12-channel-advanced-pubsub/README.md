# channel advanced pubsub 
채널(channel)을 통해 메시지 발행/구독(pubsub) 모델을 만들고 구현하는 예제 입니다.

## pubsub design
발행 구독(pubsub) 의 주요 객체에 대해 이해 합니다.

Pubsub struct
-----
발행 구독 객체로 동시성을 해결하기 위한 mu:mutex 와 구독자 정보를 담고 있는 subs:map 이 있습니다.
발행 구독 채널(파이프라인)이 현재 닫혔는지 확인하는 closed:bool 속성도 관리 합니다. 
```
type Pubsub struct {
    mu     sync.RWMutex
    subs   map[string][]chan string
    closed bool
}
```

NewPubsub
-----
Pubsub 객체를 생성하는 팩토리 메서드 입니다.  
```
pubsub := NewPubsub() // pubsub 객체 생성  

func NewPubsub() *Pubsub {
	ps := &Pubsub{}
	ps.subs = make(map[string][]chan string)
	ps.closed = false
	return ps
}
``` 

Subscribe
-----
구독자(sub) 객체를 생성 하고 각각의 구독자는 내부적으로 채널(channel) 파이프라인을 생성 하여 반환 합니다.  
메시지를 구독 하는 동한 다른 구독자가 동시에 처리 되지 않게 하기 위해 mutex.lock 를 걸고, 모든 작업이 완료 되었을 때 mutex.unlock 으로을 해제 합니다.  
특히 Pubsub 객체의 구독자의 토픽에 다수의 구독자를 추가하는 로직은 다음과 같습니다.
ps.subs[topic] = append(ps.subs[topic], ch)
```
ch1 := ps.Subscribe("tech") // tech 토픽의 ch1 구독자 생성

func (ps *Pubsub) Subscribe(topic string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string, 1)
	ps.subs[topic] = append(ps.subs[topic], ch)
	return ch
}
```

Publish
-----
Pubsub 객체를 통해 구독자에게 메시지를 발행하는 처리기 입니다.  
메시지 발행중에 충돌을 방지하기 위해 mutex lock/unlock 을 처리 하고 있흡니다.
```
ps.Pubsub("토픽이름", "발행메시지")

func (ps *Pubsub) Publish(topic string, msg string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ps.closed {
		return
	}

	for _, ch := range ps.subs[topic] {
		ch <- msg
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