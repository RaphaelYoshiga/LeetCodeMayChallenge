package main;

func countSquares(matrix [][]int) int {

	sum := 0;
	for i, row := range matrix{
		for y, cell := range row{
			if cell == 1{
				value := cell;
				if i > 0 && y > 0{
					value += min(matrix[i-1][y], min(matrix[i-1][y-1], row[y-1]))
				}

				sum += value;
				row[y] = value;
			}
		}
	}

	return sum;
}

func min(a int, b int) int{
	if a > b{
		return b;
	}
	return a;
}