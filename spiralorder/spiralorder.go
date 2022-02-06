package spiralorder

func SpiralOrder(matrix [][]int) (values []int) {
	width := len(matrix[0])
	height := len(matrix)

	for ring := 0; ; ring++ {
		margin := ring

		wStart, wEnd := margin, width-1-margin
		hStart, hEnd := margin, height-1-margin

		if wStart > wEnd {
			break
		}

		if hStart > hEnd {
			break
		}

		topRow := ring
		rightCol := wEnd
		bottomRow := hEnd
		leftCol := ring

		// Top
		for i := wStart; i <= wEnd; i++ {
			values = append(values, matrix[topRow][i])
		}

		// Right
		for i := hStart + 1; i <= hEnd; i++ {
			values = append(values, matrix[i][rightCol])
		}

		// Bottom
		if bottomRow != topRow {
			for i := wEnd - 1; i >= wStart; i-- {
				values = append(values, matrix[bottomRow][i])
			}
		}

		// Left
		if leftCol != rightCol {
			for i := hEnd - 1; i > hStart; i-- {
				values = append(values, matrix[i][leftCol])
			}
		}
	}

	return
}
