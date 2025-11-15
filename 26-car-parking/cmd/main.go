package main

import (
	"fmt"
	"sync"
	"time"
)

type ParkingLot struct {
	slots chan struct{}
}

func (p *ParkingLot) Park(carID int64) {
	p.slots <- struct{}{}
	fmt.Printf("Car %d: Parking...\n", carID)

	time.Sleep(time.Second)

	fmt.Printf("Car %d: Leaving\n", carID)
	<-p.slots
}

func NewParkingLot(slotsCount int64) *ParkingLot {
	p := &ParkingLot{
		slots: make(chan struct{}, slotsCount),
	}
	return p
}

func main() {
	parkingLot := NewParkingLot(3)

	var wg sync.WaitGroup

	carIDs := []int64{1, 2, 3, 4, 5, 6}

	for _, carID := range carIDs {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()
			parkingLot.Park(id)
		}(carID)
	}
	wg.Wait()
}
