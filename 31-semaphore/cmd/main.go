package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	slots chan struct{}
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		slots: make(chan struct{}, n),
	}
}

func (s *Semaphore) Acquire() {
	s.slots <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.slots
}

func downloadFile(filename string) {
	fmt.Printf("Downloading file %s\n", filename)
	time.Sleep(time.Second)
	fmt.Printf("File %s downloaded\n", filename)
}

func main() {

	files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt"}

	sem := NewSemaphore(3)

	wg := sync.WaitGroup{}
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()
			downloadFile(file)
		}(file)
	}

	wg.Wait()
}
