package main

import "testing"

func TestRemoveKDigits(t *testing.T) {
	tables := []struct {
		a string
		k int
		exp string
	}{
		{"101", 1 , "1"},
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
		{"10", 1, "0"},
		{"10", 2, "0"},
		{"100", 1, "0"},
		{"112", 1, "11"},
		{"111111", 3, "111"},
		{"111311", 3, "111"},
		{"1101311", 3, "111"},
    }
    
	for _, table := range tables {
		result :=removeKdigits(table.a, table.k);

        if result != table.exp {
			t.Errorf("%s, actual: %s should be %s", table.a, result, table.exp);
         }
	}
}