package main

import "testing"

func TestStraightLine(t *testing.T) {
	tables := []struct {
		a [][]int
		exp bool
	}{
		
		{ [][]int {{0,1},{1,3},{-4,-7},{5,11}}, true},
		{ [][]int {{-3,-2},{-1,-2},{2,-2},{-2,-2},{0,-2}}, true},
		{ [][]int {{1,1},{2,2},{3,4},{4,5},{5,6},{7,7}}, false},
		
    }
    
	for _, table := range tables {
		result :=checkStraightLine(table.a);

        if result != table.exp {
			t.Errorf("%d, actual: %t should be %t", table.a, result, table.exp);
         }
	}
}