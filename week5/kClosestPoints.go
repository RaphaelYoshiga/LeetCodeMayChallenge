package main;

import "sort";
import "math";

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return calculateDistance(points[i]) < calculateDistance(points[j])
	});

	return points[0:K];
	
}

func calculateDistance(point []int) float64{
	X1 := 0;
	X2 := point[0];
	Y1 := 0;
	Y2 := point[1];
	
	intermediate:= float64((X1 - X2) * (X1 - X2) + (Y1 - Y2) * (Y1 - Y2));
	distance:=math.Sqrt(intermediate);
	return distance;
}

func main(){
	result:= kClosest([][]int{{1,3}, {-2,2}}, 1)
	print(result);
}