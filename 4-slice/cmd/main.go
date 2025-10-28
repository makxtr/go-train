package main

import "fmt"

type car struct {
	color   string
	mileage int
}

func main() {

	cars := []car{
		{
			color:   "red",
			mileage: 5000,
		},
		{
			color:   "green",
			mileage: 10000,
		},
		{
			color:   "blue",
			mileage: 7000,
		},
	}

	carPtr := &cars[0]
	carPtr.mileage += 100

	cars = append(cars, car{color: "yellow", mileage: 15000})
	carPtr.mileage += 50

	fmt.Println(cars[0].mileage, cars[0].color)
	fmt.Println(carPtr.mileage, carPtr.color)
}
