package pascalstriangle

func Generate(numRows int) (rows [][]int) {
	rows = append(rows, []int{1})

	for i := 0; i < numRows-1; i++ {
		prevRow := rows[i]

		curRow := []int{1}

		for j := 1; j < len(prevRow); j++ {
			curRow = append(curRow, prevRow[j-1]+prevRow[j])
		}

		curRow = append(curRow, 1)

		rows = append(rows, curRow)
	}

	return
}
