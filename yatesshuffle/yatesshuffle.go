package yatesshuffle

import (
	"fmt"
	"math/rand"
	"time"
)

func YatesShuffle(nums []int) {
	rand.Seed(time.Now().UnixNano())

	for pivot := 0; pivot < len(nums)-1; pivot++ {
		r := rand.Intn(len(nums)-pivot) + pivot

		nums[pivot], nums[r] = nums[r], nums[pivot]
	}
}

func Program() {
	// even
	nums := []int{10, 20, 30, 40}
	YatesShuffle(nums)
	fmt.Println(nums)

	// odd
	nums = []int{10, 20, 30, 40, 50}
	YatesShuffle(nums)
	fmt.Println(nums)

	// single element
	nums = []int{1}
	YatesShuffle(nums)
	fmt.Println(nums)

	// empty
	nums = []int{}
	YatesShuffle(nums)
	fmt.Println(nums)
}
