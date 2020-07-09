package main

import "fmt"

func myFunction(flightTime int, movieTimes []int) bool {
	diff := make(map[int]bool)

	for _, movieTime := range movieTimes {
		if _, ok := diff[flightTime-movieTime]; ok {
			return true
		} else {
			diff[movieTime] = true
		}
	}

	return false
}

func main() {
	flightTime := 360
	movieTimes := []int{120, 180, 150, 230}

	fmt.Println(myFunction(flightTime, movieTimes))
}
