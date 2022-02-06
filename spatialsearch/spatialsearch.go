package spatialsearch

import "math"

// Search finds occupied adjacent spaces that are within a limited range.
// The spaces that are found will be stored in the map passed in the "results" parameter.
//
// Parameters:
//	x, y		the current space's index as 2d coordinates
//	distance	range of search
//	spaces		sparse table of spaces (key = spatial hash, value = occupant)
//	results		set of occupied adjacent spaces
func Search(
	x, y uint32,
	distance uint32,
	spaces map[uint64]interface{},
	results map[uint64]bool,
) {
	// spatial hash of current location
	current := uint64(x)<<32 | uint64(y)

	if distance == 0 {
		// current space is out of range
		return
	}

	if _, ok := results[current]; !ok {
		// current space has already been found
		return
	}

	if _, ok := spaces[current]; !ok {
		// current space is not occupied
		return
	}

	results[current] = true

	d := distance - 1

	// top
	if y < math.MaxUint32 {
		Search(x, y+1, d, spaces, results)
	}

	// left
	if x > 0 {
		Search(x-1, y, d, spaces, results)
	}

	// right
	if x < math.MaxUint32 {
		Search(x+1, y, d, spaces, results)
	}

	// bottom
	if y > 0 {
		Search(x, y-1, d, spaces, results)
	}
}
