package main

import (
	f "fmt"
	"sync"
	"time"
)

type Pubsub struct {
	mu     sync.RWMutex
	subs   map[string][]chan string
	closed bool
}

func NewPubsub() *Pubsub {
	ps := &Pubsub{}
	ps.subs = make(map[string][]chan string)
	ps.closed = false
	return ps
}

func (ps *Pubsub) Subscribe(topic string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string, 1)
	// Pubsub 토픽에 구독자 ch 를 추가 합니다.
	ps.subs[topic] = append(ps.subs[topic], ch)
	return ch
}

func (ps *Pubsub) Publish(topic string, msg string) {
	f.Printf("Publishing @%s: %s\n", topic, msg)
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ps.closed {
		return
	}

	for _, ch := range ps.subs[topic] {
		ch <- msg
	}
}

func (ps *Pubsub) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if !ps.closed {
		ps.closed = true
		for _, subs := range ps.subs {
			for _, ch := range subs {
				close(ch)
			}
		}
	}
}

func (ps *Pubsub) PublishInterval(topic string, msg string) {
	f.Println("wait 500ms...")
	time.Sleep(time.Millisecond * 500)
	ps.Publish(topic, msg)
}

func main() {

	ps := NewPubsub()
	ch1 := ps.Subscribe("tech")
	ch2 := ps.Subscribe("travel") // travel 토픽에 ch2 구독자가 추가됨
	ch3 := ps.Subscribe("travel") // travel 토픽에 ch3 구독자가 추가됨

	listener := func(name string, ch <-chan string) {
		for i := range ch {
			f.Printf("[%s] got %s\n", name, i)
		}
		f.Printf("[%s] done\n", name)
	}

	// goroutine 을 통해 경량 스레드로 구독자마다 리스너 추가
	go listener("sub-1", ch1)
	go listener("sub-2.1", ch2)
	go listener("sub-2.2", ch3)

	time.Sleep(time.Second)
	ps.Publish("tech", "1 tablets")
	ps.Publish("travel", "2 vitamins")
	ps.Publish("tech", "3 robots")
	time.Sleep(time.Millisecond * 300)

	// 500ms 만큼 시간 간격을 두고 topic 에 메시지를 발행
	ps.PublishInterval("travel", "4 beaches")
	ps.PublishInterval("travel", "5 hiking")
	ps.PublishInterval("tech", "6 drones")

	time.Sleep(1000 * time.Millisecond)
	ps.Close()
}
