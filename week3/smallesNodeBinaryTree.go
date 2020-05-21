package main;

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func bstSmallest(t *TreeNode, k, res *int){
    if t == nil || *k == 0 {
        return
    }
    
    bstSmallest(t.Left, k, res)
    
    (*k)--
    
    if *k == 0 {
        *res = t.Val
        return
    }

    bstSmallest(t.Right, k, res)
}

func kthSmallest(root *TreeNode, k int) int {
    var res int
    bstSmallest(root, &k, &res)
    return res
}