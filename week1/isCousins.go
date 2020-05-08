package main;

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
};

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type CousingStatus struct {
	XDepth int
	XParent int
	YDepth int
	YParent int
}

 func isCousins(root *TreeNode, x int, y int) bool {

	if root.Val == x || root.Val == y{
		return false;
	}

	status := &CousingStatus{ };
	goNextLevel(root, x, y, root, 1, status)

	return status.XDepth == status.YDepth && status.XParent != status.YParent;

}

func goNextLevel(node *TreeNode, x int, y int, parent *TreeNode, depth int, status *CousingStatus){

	if node.Val == x{
		status.XDepth = depth;
		status.XParent = parent.Val;
	}
	
	if node.Val == y{
		status.YDepth = depth;
		status.YParent = parent.Val;
	}

	depth++;
	if node.Left != nil{
		goNextLevel(node.Left, x, y, node, depth, status);
	}
	if node.Right != nil{
		goNextLevel(node.Right, x, y, node, depth, status);
	}
}