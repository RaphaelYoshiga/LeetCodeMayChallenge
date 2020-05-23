package main;

func intervalIntersection(A [][]int, B [][]int) [][]int {
	intersections := [][]int {};
	for _, r :=range A{
		for _, c:= range B{
			if r[0] <= c[0] && c[1] >= r[1] && r[1] >= c[0]{
				intersections = append(intersections, []int { getStart(r, c), getEnding(r, c)});
			}else if r[0] >= c[0] && r[1] <= c[1]{
				intersections = append(intersections, []int { getStart(r, c), getEnding(r, c)});
			}else if c[1] >= r[0] && r[1] >= c[1] {
				intersections = append(intersections, []int { getStart(r, c), getEnding(r, c)});
			}
		}
	}
	
	return intersections;
}

func getStart(a []int, b []int) int{
	if a[0] > b[0]{
		return a[0];
	}
	return b[0];
}

func getEnding(a []int, b []int) int{
	if b[1] > a[1]{
		return a[1];
	}
	return b[1];
}

func main(){
	intervalIntersection([][]int { {3,5}, {9,20} }, [][]int { {4,5},{7,10},{11,12},{14,15},{16,20}});
}