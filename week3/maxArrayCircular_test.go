package main

import "testing"

func TestMaxCircularArray(t *testing.T) {
	tables := []struct {
		a []int
		exp int
	}{
		
		
		{ []int {5,-3,5}, 10},
		{ []int {1,-2,3,-2}, 3},
		{ []int {3,-2,2,-3}, 3},
		{ []int {-2,-3,-1}, -1},
		{ []int {-2}, -2},
    }
    
	for _, table := range tables {
		result := maxSubarraySumCircular(table.a);

        if result != table.exp {
			t.Errorf("Max subarray circular: %d, actual: %d should be %d", table.a, result, table.exp);
         }
	}
}