package main

import (
	"fmt"
	"sync"
)

type Toilet struct {
	Mu    sync.Mutex
	Cond  *sync.Cond
	Ch    chan string
	Group string
}

func NewToilet() *Toilet {
	t := &Toilet{
		Ch: make(chan string, 3),
	}
	t.Cond = sync.NewCond(&t.Mu)
	return t
}

func (t *Toilet) Use(group string) {
	t.Mu.Lock()
	for t.Group != "" && t.Group != group {
		t.Cond.Wait()
	}
	t.Group = group
	t.Mu.Unlock()
	t.Ch <- group
	fmt.Println("Entered Group:", group)
	<-t.Ch
	t.Mu.Lock()
	if len(t.Ch) == 0 {
		t.Group = ""
		t.Cond.Broadcast()

	}
	t.Mu.Unlock()

}

func main() {
	t := NewToilet()
	var wg sync.WaitGroup
	wg.Add(6)
	go func() {
		t.Use("GroupA")
		wg.Done()
	}()
	go func() {
		t.Use("GroupA")
		wg.Done()
	}()
	go func() {
		t.Use("GroupA")
		wg.Done()
	}()
	go func() {
		t.Use("GroupB")
		wg.Done()
	}()
	go func() {
		t.Use("GroupB")
		wg.Done()
	}()
	go func() {
		t.Use("GroupB")
		wg.Done()
	}()

	wg.Wait()

}
