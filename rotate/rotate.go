package rotate

func rotate(nums []int, k int) {
	for i, numModified := 0, 0; i < len(nums) && numModified < len(nums); i++ {
		temp := nums[i]

		for j := i; ; {
			offset := (j + k) % len(nums)

			nums[offset], temp = temp, nums[offset]
			numModified++

			j = offset
			if j == i {
				break
			}
		}
	}
}

func Program() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	//       Offset
	//       3
	// 1 2 3 4 5 6 7
	// 5 6 7 1 2 3 4

	rotate([]int{-1, -100, 3, 99}, 2)
	//          Offset
	//          2
	// -1 -100  3  99
	//  3  99  -1 -100
}
