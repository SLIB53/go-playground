package reverse

import "fmt"

func reverse(nums []int) {
	for start, end := 0, len(nums)-1; start < end; start, end = start+1, end-1 {
		nums[start], nums[end] = nums[end], nums[start]
	}
}

func Program() {
	a := []int{1, 2, 3, 4}
	reverse(a)
	fmt.Println(a)

	a = []int{1, 2, 3, 4, 5}
	reverse(a)
	fmt.Println(a)

	a = []int{}
	reverse(a)
	fmt.Println(a)

	a = []int{1}
	reverse(a)
	fmt.Println(a)
}
