package main

import "testing"

func TestOddEvenList(t *testing.T) {
	tables := []struct {
		a []int
		exp[] int
	}{
		{ []int {2, 1, 3}, []int {2, 3, 1} },
		{ []int {2, 1, 3, 5, 6, 4, 7}, []int {2, 3, 6, 7, 1, 5, 4} },
    }
	for _, table := range tables {
		head := &ListNode { Val: table.a[0]};
		current := head;

		for	i := 1; i < len(table.a); i++{
			next := &ListNode{ Val: table.a[i]};
			current.Next = next;
			current = next;
		}

		result := oddEvenList(head);
		currentNode := result;
		for i, expected := range table.exp {

			if currentNode.Val != expected {
				t.Errorf("%d Current nodes should be the same at %d, actual: %d should be %d", table.a, i, currentNode.Val, expected);
			}
			currentNode = currentNode.Next;
		}

		if currentNode != nil {
			t.Errorf("%d, Last node should be nil!", table.a);
		}
	
	}
}