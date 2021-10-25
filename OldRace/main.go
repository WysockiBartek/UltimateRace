package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
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
	mainLength := 5
	subLenght := 5
	allCars := make([][]Car, 5)
	podium := make([]Car, 3)
	finalGr := make([]Car, 5)
	// gr1 := []Car{}
	// gr2 := []Car{}
	// gr3 := []Car{}
	// gr4 := []Car{}
	// gr5 := []Car{}
	// gr6 := []Car{}

	for i := 0; i < mainLength; i++ {
		allCars[i] = make([]Car, 5)
		for j := 0; j < subLenght; j++ {
			allCars[i][j] = Car{}
		}
	}

	for i := 0; i < mainLength; i++ {
		sort.Slice(allCars[i], func(a, b int) bool {
			return allCars[i][a].GetCarSpeed() > allCars[i][b].GetCarSpeed()
		})
	}

	sort.Slice(allCars, func(a, b int) bool {
		return allCars[a][0].GetCarSpeed() > allCars[b][0].GetCarSpeed()
	})

	podium[0] = allCars[0][0]

	finalGr = append(finalGr, allCars[0][1], allCars[0][2], allCars[1][0], allCars[1][1], allCars[2][0])

	sort.Slice(finalGr, func(a, b int) bool {
		return finalGr[a].GetCarSpeed() > finalGr[b].GetCarSpeed()
	})

	podium = append(podium, finalGr[0], finalGr[1])

	fmt.Println("All Cars: ", allCars)
	fmt.Println("Final Group: ", finalGr)
	fmt.Println("Podium: ", podium)
	// for i := 0; i < 11; i++ {
	// 	allCars = RaceSortRemove(allCars)
	// }

	// fmt.Println("Race results: ")
	// for _, c := range allCars {
	// 	fmt.Printf(" %d\n", c.speed)
	// }
}
