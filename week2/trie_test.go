package main

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor();

	trie.Insert("apple");
	assert(t, trie.Search("apple"), true, "search");   // returns true
	assert(t, trie.Search("app"), false, "search");     // returns false
	assert(t, trie.StartsWith("app"), true, "starts with"); // returns true
	trie.Insert("app");  
	assert(t, trie.Search("app"), true, "search"); // returns true
}

func assert(t *testing.T, actual bool, expected bool, debug string){
	if actual != expected{
		t.Errorf("Trie assert -- %s, %t should be %t", debug, actual, expected);
	}
}