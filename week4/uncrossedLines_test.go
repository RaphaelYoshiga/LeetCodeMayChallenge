package main;

import "testing";
func TestUncrossedLines(t *testing.T) {
	tables := []struct {
		a []int
		b []int
		exp int
	}{
		{ []int { 1 }, []int { 1}, 1},
		{ []int { 1,1 }, []int { 1,1}, 2},
		{ []int { 2,2 }, []int { 2,2}, 2},
		{ []int { 2,2,3 }, []int { 2,2,1,3}, 3},
		{ []int { 2,2,4 }, []int { 2,2,1,4}, 3},
		{ []int { 2,2,4,3 }, 
		  []int { 2,2,3}, 3},
		{ []int { 2,2,4,3 }, 
		  []int { 2,2,3}, 3},

		{ []int { 1,2 }, 
		  []int { 2,1 }, 1},
	
		{ []int { 2,5,1,2,5 }, 
		  []int { 10,5,2,1,5,2 }, 3},
		
		{ []int { 3,2 }, 
		  []int { 2, 2, 2, 3 }, 1},
		{ []int { 3,3 }, 
		  []int { 1,3,1,2,1 }, 1},

		{ []int { 1,2,2 }, 
		  []int { 3 }, 0},

		{ []int { 3,3,1,3 }, 
		  []int { 1,3,2,3,3,1 }, 3},


    }
    
	for _, table := range tables {
		result := maxUncrossedLines(table.a, table.b);

        if result != table.exp {
			t.Errorf("Uncrossed lines for %d %d\n  actual %d \n should be %d", table.a, table.b, result, table.exp);
         }
	}
}