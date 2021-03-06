package main;


func oddEvenList(head *ListNode) *ListNode {
    if head == nil { return nil }
    odd := head
    evenHead := odd.Next
    even := evenHead
    
    
    for even != nil && even.Next != nil {
        odd.Next = even.Next
        odd = odd.Next
        even.Next = odd.Next
        even = even.Next
    }
    odd.Next = evenHead
    return head    

}   

// func oddEvenList(head *ListNode) *ListNode {
// 	current := head;
// 	if head == nil || current.Next == nil{
// 		return head;
// 	}

// 	tailStart := current.Next;
// 	var tailEnd *ListNode
// 	tailEnd = current.Next;

// 	var lastOdd *ListNode

// 	i := 0;
// 	for current != nil {
// 		if (i + 1) % 2 == 0{
// 			if tailEnd != current{
// 				tailEnd.Next = current;
// 				tailEnd = current;
// 			}

// 			next := current.Next;

// 			lastOdd.Next = next;

// 			current.Next = nil;
// 			current = next;
// 		}else{
// 			lastOdd = current;
// 			current = current.Next;
// 		}
// 		i++;
// 	}
// 	lastOdd.Next = tailStart;
//     return head;
// }

type ListNode struct {
	Val int
	Next *ListNode
}