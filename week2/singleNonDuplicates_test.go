package main

import "testing"

func TestBinaryScan(t *testing.T) {
	tables := []struct {
		a []int
		exp int
	}{
		
		
		{ []int {1,1,2}, 2},
		{ []int {1,1,2,2,3}, 3},
		{ []int {1,1,2,3,3}, 2},
		{ []int {1,1,2,3,3,4,4,8,8}, 2},
		{ []int {3,3,7,7,10,11,11}, 10},
    }
    
	for _, table := range tables {
		result := singleNonDuplicate(table.a);

        if result != table.exp {
			t.Errorf("%d, actual: %d should be %d", table.a, result, table.exp);
         }
	}
}