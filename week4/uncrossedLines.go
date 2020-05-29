package main;
func max(x ...int) int {
	max := x[0]
	for _, c := range x {
		if c > max {
			max = c
		}
	}
	return max
}


func maxUncrossedLines(A []int, B []int) int {
  L1 := len(A)
	L2 := len(B)
	mems := make([][]int, L1+1)
	for i := 0; i < L1+1; i++ {
		mems[i] = make([]int, L2+1)
	}

	for i := 1; i <= L1; i++ {
		for j := 1; j <= L2; j++ {
			if A[i-1] == B[j-1] {
				mems[i][j] = max(
					mems[i-1][j],
					mems[i][j-1],
					mems[i-1][j-1]+1,
				)
			} else {
				mems[i][j] = max(
					mems[i-1][j],
					mems[i][j-1],
					mems[i-1][j-1],
				)
			}
		}
	}

	return mems[L1][L2] 
  
}
// func maxUncrossedLines(A []int, B []int) int {

// 	sum:= 0;
// 	for i, value := range A{
// 		if i >= len(B){
// 			break;
// 		}
		
// 		if i < len(B) && value == B[i]{
// 			sum++;
// 			A[i] = -1;
// 			B[i] = -1;
// 		}else if i < len(B) -1 && value == B[i + 1]{
// 			sum++;
// 			A[i] = -1;
// 			B[i + 1] = -1;
// 		}else if i < len(A) -1 && A[i + 1] == B[i]{
// 			sum++;
// 			A[i+1] = -1;
// 			B[i] = -1;
// 		}
// 	}

//     return sum;
// }

func main(){
	maxUncrossedLines([]int {3,3}, []int {1,3,1,2,1});
}