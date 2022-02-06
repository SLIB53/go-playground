package bitstring

import (
	"fmt"
	"sync"
)

func Program() {
	var b uint32
	fmt.Println(binaryString(b))
	// output:
	// 00000000000000000000000000000000

	b = 0xFFFF_FFFF
	fmt.Println(binaryString(b))
	// output:
	// 11111111111111111111111111111111

	b >>= 1
	fmt.Println(binaryString(b))
	// output:
	// 01111111111111111111111111111111

	fmt.Println(binaryString((1 << 31) >> 1))
	// output:
	// 01000000000000000000000000000000
}

func binaryRuneAt(bin uint32, i int) (r rune) {
	const msbmask = 1 << 31

	if (bin<<i)&msbmask == 0 {
		r = '0'
	} else {
		r = '1'
	}

	return
}

func binaryString(bin uint32) string {
	// Create a shared buffer to store bits converted to unicode characters.
	bits := make([]rune, 32)

	// Create a WaitGroup, with a counter of 32 for each bit.
	var wg sync.WaitGroup
	wg.Add(len(bits))

	// Map each bit to unicode character.
	for i := range bits {
		// Spawn a goroutine to concurrently execute conversion from bit to unicode character.
		go func(j int) {
			// Decrement the WaitGroup counter at the end of goroutine.
			defer wg.Done()

			// Map the bit to unicode character and store in shared buffer.
			bits[j] = binaryRuneAt(bin, j)
		}(i)
	}

	// Block until all 32 bits have been converted.
	wg.Wait()

	// Return formatted buffer.
	return string(bits)
}
