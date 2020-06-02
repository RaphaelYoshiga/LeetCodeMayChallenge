package main;

import "testing";
func TestBfs(t *testing.T) {
	head:= &Node { 
		Neighbours: []*Node {
			&Node { 
				Val: "C", 
				Neighbours: []*Node {
					&Node { Val:"D", Neighbours: []*Node {}},
				}},
				&Node { 
					Val: "F", 
					Neighbours: []*Node {
						&Node { Val:"G", Neighbours: []*Node {}},
					},
			},}};

	tables := []struct {
		Target string
		ExpectedLevel int
	}{
		{ "C", 1},
		{ "D", 2},
		{ "F", 1},
		{ "G", 2},
    }
    
	for _, table := range tables {
		result := bfs(head ,table.Target);

        if result != table.ExpectedLevel {
			t.Errorf("Wrong level for %s. actual %d should be %d", table.Target, result, table.ExpectedLevel);
         }
	}
}
