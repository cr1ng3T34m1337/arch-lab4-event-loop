package engine

import (
	"sync"
)

type Handler interface {
	Post(cmd Command)
}

type queue struct {
	mu      sync.Mutex
	s       []Command
	waiting bool

	isReady chan struct{}
}

func (q *queue) push(cmd Command) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.s = append(q.s, cmd)

	if q.waiting {
		q.waiting = false
		q.isReady <- struct{}{}
	}
}

func (q *queue) pull() Command {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.s) == 0 {
		q.waiting = true
		q.mu.Unlock()
		<-q.isReady
		q.mu.Lock()
	}

	res := q.s[0]
	q.s[0] = nil
	q.s = q.s[1:]

	return res
}

func (q *queue) isEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.s) == 0
}

type EventLoop struct {
	messageQueue *queue
	stop         bool
	stopSignal   chan struct{}
}

func (el *EventLoop) Start() {
	el.messageQueue = &queue{
		isReady: make(chan struct{}),
	}
	el.stopSignal = make(chan struct{})
	go func() {
		for !el.stop || !el.messageQueue.isEmpty() {
			cmd := el.messageQueue.pull()
			cmd.Execute(el)
		}
		el.stopSignal <- struct{}{}
	}()
}

func (el *EventLoop) Post(cmd Command) {
	el.messageQueue.push(cmd)
}

func (el *EventLoop) AwaitFinish() {
	el.stop = true
	<-el.stopSignal
}
