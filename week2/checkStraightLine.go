package main;

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true;
	}

	x1 := coordinates[0][0];
	y1 := coordinates[0][1];

	xDiff := coordinates[1][0] - x1;
	yDiff := coordinates[1][1] - y1;

	if xDiff == 0{
		return false;
	}

	slope := float64(yDiff) / float64(xDiff);

	for i:= 2; i < len(coordinates); i++ {
		
		yDiff =  coordinates[i][1] - y1;
		xDiff = coordinates[i][0] - x1;
		currentSlope := float64(yDiff) / float64(xDiff);

		if currentSlope != slope{
			return false;
		}
	}

	return true;
}

func main(){
	checkStraightLine([][]int {{1,1},{2,2},{3,4},{4,5},{5,6},{7,7}})
}