package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	unlocked = false
	locked   = true
)

type Mutex struct {
	state atomic.Bool
}

func (m *Mutex) Lock() {
	for !m.state.CompareAndSwap(unlocked, locked) {
	}
}

func (m *Mutex) Unlock() {
	m.state.Store(unlocked)
}

const gNumber = 1000

func main() {
	var mutex Mutex
	var wg sync.WaitGroup
	wg.Add(gNumber)

	value := 0
	for i := 0; i < gNumber; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			value++
			defer mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(value)
}
