package main;

func main(){
	flipAndInvertImage([][]int { []int { 1, 1, 0, 0} })
}

func flipAndInvertImage(A [][]int) [][]int {
	rowLen := len(A[0]);
	isEven := rowLen % 2 == 0;

	middle := rowLen / 2;
	for i := 0; i < len(A); i++ {
		row := A[i];

		for y := 0; y < rowLen; y++ {
			if !isEven && y == middle {
				row[y] = flip(row[y]);
				break;
			}

			if isEven && y >= middle {
				break;
			}

			index := rowLen - y - 1;
			aux:= row[index];
			row[index] = flip(row[y]);

			row[y] = flip(aux);
		}
	}

    return A;
}

func flip(a int) int{
	if a == 1{
		return 0;
	}
	return 1;
}