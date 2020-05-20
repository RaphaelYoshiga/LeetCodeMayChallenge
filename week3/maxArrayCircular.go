package main;

const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

func maxSubarraySumCircular(A []int) int {
	if len(A) == 1{
		return A[0];
	}

	maxKadane := kadane(A);

	maxWrap := 0
    for i, _:= range A {
		maxWrap += A[i] 
        A[i] = -A[i] 
	}

	maxNegative :=kadane(A);

	if maxNegative + maxWrap == 0{
		return maxKadane;
	}

    return max(maxKadane, maxWrap + maxNegative);
}

func kadane(A []int) int{
	n := len(A);

	dp := make([]int,n)

	dp[0] = A[0];

	answer := MinInt;

	for i:= 1; i < n; i++{

		previous := dp[i - 1];
		dp[i] = max(previous, 0) + A[i];
		answer = max(answer, dp[i]);
	}
	return answer;
}

func max(a int, b int) int{
	if a > b{
		return a;
	}
	return b;
}

