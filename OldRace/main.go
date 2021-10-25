package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

var raceCounter int = 0

// declaring a struct
type Car struct {
	// defining struct variables
	speed int
}

// function to assign speed
func (car *Car) GetCarSpeed() int {
	// rand.Seed(time.Now().UnixNano())
	// c := new(Car)
	if car.speed == 0 {
		min := 40
		max := 200
		car.speed = rand.Intn(max-min+1) + min
	}
	return car.speed
}

func RaceSortRemove(allCars []Car) []Car {
	swapF := reflect.Swapper(allCars)
	maxi, maxj := 4, 5
	//--------SORT-----------//
	for i := 0; i < maxi; i++ {
		for j := i + 1; j < maxj; j++ {
			if allCars[i].GetCarSpeed() > allCars[j].GetCarSpeed() {
				swapF(i, j)
			}
		}
	}
	//--------SORT-----------//
	raceCounter++
	fmt.Printf("Race no. %d:\n", raceCounter)
	for j := 0; j < maxj; j++ {
		fmt.Printf(" %d\n", allCars[j].GetCarSpeed())
	}

	allCars = append(allCars[2:])

	return allCars
}

func main() {
	mainLength := 25
	subLenght := 5
	allCars := []Car{}

	for i := 0; i < subLenght; i++ {
		for j := 0; j < 5; j++ {
			allCars = append(allCars, Car{})
		}

	}

	for i := 0; i < 11; i++ {
		allCars = RaceSortRemove(allCars)
	}

	fmt.Println("Race results: ")
	for _, c := range allCars {
		fmt.Printf(" %d\n", c.speed)
	}
}
