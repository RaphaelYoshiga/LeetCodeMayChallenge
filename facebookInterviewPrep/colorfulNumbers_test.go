package main;

import "testing";
func TestMaxCircularArray(t *testing.T) {
	tables := []struct {
		a int
		exp bool
	}{
		{ 3245, true},
		{ 326, false},
    }
    
	for _, table := range tables {
		result := colorfulNumber(table.a);

        if result != table.exp {
			t.Errorf("Colorful number. %d actual %t should be %t", table.a, result, table.exp);
         }
	}
}
