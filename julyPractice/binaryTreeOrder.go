package main;
import "sort";

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	stack := make(map[int][]int);
	if root == nil{
		return [][]int {};
	}

	goDeep(root, stack, 1);

	return arrayFromMap(stack);
}

func arrayFromMap(stack map[int][]int) [][]int{
	var keys []int
    for k := range stack {
        keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	});

	array := [][]int{ };
	for _, k := range keys{
		array = append(array, stack[k]);
	}
	return array;
}

func goDeep(node *TreeNode, stack map[int][]int, level int){
	if node.Left!= nil {
		goDeep(node.Left, stack, level + 1);
	}
	if node.Right != nil {
		goDeep(node.Right, stack, level +1);
	}
	stack[level] = append(stack[level], node.Val);
}

func main(){
	node :=TreeNode {
		Val: 3,
		Left: &TreeNode { 
			Val: 2,
		},
		Right: &TreeNode {
			Val: 4,
		},
	};
	levelOrderBottom(&node);
}