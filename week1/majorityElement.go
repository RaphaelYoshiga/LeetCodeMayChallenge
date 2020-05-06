package main;

func majorityElement(nums []int) int {

	midIndex := len(nums) / 2;

	countMap:= make(map[int]int)
    for _, n := range nums{
		countMap[n]++;
		if countMap[n] > midIndex{
			return n;
		}
	}
	return -1;
}