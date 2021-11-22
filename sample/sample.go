package main

import (
	"fmt"

	"github.com/Chris-Behan/howfast"
)

func main() {
	// generate list of nums between 1 and 100 mil
	nums := generateNums(100000000)
	// set target to 77,777,777
	target := 77777777
	// initialize timer struct
	timer := howfast.Timer{}
	// start timer for linear sesarch
	timer.Start()
	linearAns := linearSearch(nums, target)
	// stop timer after linear search returns
	timer.Stop()
	linearMilliSeconds, _ := timer.MillisecondsElapsed()
	fmt.Printf("It took linear search %v milliseconds to find the target, which was at index %v.\n", linearMilliSeconds, linearAns)

	// reset the timer
	timer.Clear()

	timer.Start()
	binaryAns := binarySearch(nums, target)
	timer.Stop()
	binaryMilliSeconds, _ := timer.MillisecondsElapsed()
	fmt.Printf("It took binary search %v milliseconds to find the target, which was at index %v.\n", binaryMilliSeconds, binaryAns)
}

func generateNums(n int) []int {
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	return nums
}

func linearSearch(nums []int, target int) int {
	for idx, val := range nums {
		if val == target {
			return idx
		}
	}
	return -1
}

func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
