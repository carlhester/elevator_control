package main

import (
	"elevator/car"
	"fmt"
	"log"
)

func main() {
	floors := []int{1, 2, 3, 4, 5}

	Car1, err := car.NewCar(floors)
	if err != nil {
		log.Printf("Unable to create new Car", err)
	}

	fmt.Printf("%+v\n", Car1)
}
