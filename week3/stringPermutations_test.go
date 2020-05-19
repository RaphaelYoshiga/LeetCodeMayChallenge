package main

import "testing"

func TestStringPermutations(t *testing.T) {
	tables := []struct {
		a string
		b string
		exp bool
	}{
		{ "abc", "ccccbbbbaaaa", false },
    }
	for _, table := range tables {

		result := checkInclusion(table.a, table.b);

		if result != table.exp {
			t.Errorf("%s - %s, contains a substring permutation actual %t, should be: %t", table.a, table.b, result, table.exp);
		}
	
	}
}