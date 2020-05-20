package main;

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	min := root.Val;

	bst(root)

	return min;
}

const (
    MaxInt  = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
    MinInt  = -MaxInt - 1         // -1 << 31 or -1 << 63
    MaxUint = 1<<UintSize - 1     // 1<<32 - 1 or 1<<64 - 1
)

func bst(node *TreeNode) int{
	if node == nil {
		return MaxInt;
	}

	right := bst(node.Right);
	left := bst(node.Left);

	childMin := min(left, right);
	return min(node.Val, childMin);
	
}

func min(a int, b int) int{
	if a > b{
		return b;
	}
	return a;
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}