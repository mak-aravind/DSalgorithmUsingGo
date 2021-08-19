package main

import (
	"fmt"

	"github.com/mak-aravind/DSalgorithmUsingGo/map-problems/twosum"
)

func main() {
	//KEY: target value
	//VALUE :the array that has the couplets which will sum up to target
	inputNumsMap := map[int][]int{
		4:  {1, 2, 0},
		7:  {3, 4, 5},
		0:  {-1, -2, 0, 1, 4},
		10: {3, 7, 2, -3},
		12: {0, 2, -2, 4, 10, 7, 8},
	}
	for target, nums := range inputNumsMap {
		fmt.Println("Finding couplet for input: ", nums)
		fmt.Println("Target value:", target)
		if itsThere, indices := twosum.GetAddendIndics(target, nums); itsThere {
			fmt.Println("The couplet available at positions", indices)
		} else {
			fmt.Println("Not available")
		}
	}
}
