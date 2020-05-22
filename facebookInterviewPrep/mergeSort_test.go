package main;

import "testing";
func TestMergeSort(t *testing.T) {
	tables := []struct {
		a []int
		exp []int
	}{
		{ []int { 3, 2, 1}, []int {1, 2, 3}},
		{ []int { 3, 2}, []int {2, 3}},
		{ []int { 2, 1}, []int {1, 2}},
		{ []int { 4,3, 2, 1}, []int {1,2,3,4}},
		{ []int { 9,8,7,6,5,4,3,2,1}, []int {1,2,3,4,5,6,7,8,9}},
    }
    
	for _, table := range tables {
		result := mergeSort(table.a);

        if !SliceEqual(table.exp, result) {
			t.Errorf("Merge sort. %d actual %d should be %d", table.a, result, table.exp);
         }
	}
}

func SliceEqual(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}