package rle

import (
	"fmt"
)

func RLE(s string) (sEncoded string) {
	for start := 0; start < len(s); {
		end := start
		for ; end < len(s) && s[start] == s[end]; end++ {
		}

		sEncoded += fmt.Sprintf("%d%s", end-start, string(s[start]))

		start = end
	}

	return
}

func RLEDecode(nums []int) (numsDecoded []int) {
	for i := 0; i < len(nums); i += 2 {
		rlen, rval := nums[i], nums[i+1]

		for j := 0; j < rlen; j++ {
			numsDecoded = append(numsDecoded, rval)
		}
	}

	return
}

func Program() {
	fmt.Println(RLEDecode([]int{1, 2, 3, 4}))
}
