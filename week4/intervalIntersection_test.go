package main;

import "testing";
func TestIntervalDetector(t *testing.T) {
	tables := []struct {
		a [][]int
		b [][]int
		exp [][]int
	}{
		{ [][]int { {3,5} }, [][]int { {1,5}}, [][]int { {3,5 }}},
		{ [][]int { {3,3} }, [][]int { {1,5}}, [][]int { {3,3}}},
		{ [][]int { {0,3} }, [][]int { {1,3}}, [][]int { {1,3}}},
		{ [][]int { {0,3} }, [][]int { {3,4}}, [][]int { {3,3}}},
		{ [][]int { {9,10} }, [][]int { {4,5}}, [][]int { }},
		{ [][]int { {9,20} }, [][]int { {11,12}}, [][]int { {11,12} }},
		{ [][]int { {0,3} }, [][]int { {2,4}}, [][]int { {2,3}}},
		{ [][]int { {0,2}, {5,10} }, [][]int { {1,5}, {8,12}}, [][]int { {1,2}, {5,5}, {8,10}}},


		{ [][]int { {3,5}, {9,20} }, 
		[][]int { {4,5},{7,10},{11,12},{14,15},{16,20}}, 
		[][]int { {4,5},{9,10},{11,12},{14,15},{16,20}}},

		{ [][]int { {0,2},{5,10},{13,23},{24,25} }, 
		  [][]int { {1,5},{8,12},{15,24},{25,26}},
		  [][]int { {1,2},{5,5},{8,10},{15,23},{24,24},{25,25}}},

    }
    
	for _, table := range tables {
		result := intervalIntersection(table.a, table.b);

        if !MatrixEquals(table.exp, result) {
			t.Errorf("IntervalIntersection sort. %d \n actual %d \n should be %d", table.a, result, table.exp);
         }
	}
}

func MatrixEquals(a, b [][]int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
		for y, cell := range v{
			if cell != b[i][y]{
				return false
			}
		}
    }
    return true
}