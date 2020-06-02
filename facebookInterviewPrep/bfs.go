package main;

type Node struct{
	Val string
	Neighbours []*Node
}

func bfs(node *Node, target string) int{
	visited := make(map[string]int);
	visited[node.Val] = 0;

	level := 1;
	
	frontier := []*Node { node };
	for len(frontier) > 0{
		next := []*Node {}
		for _, u := range frontier{
			for _, v := range u.Neighbours{
				_, found := visited[v.Val];
				if !found {
					visited[v.Val] = level;
					next = append(next, v);
				}
			}
		}
		frontier = next;
		level++;
	}
	return visited[target];
}

func main(){
	head:= &Node { 
		Neighbours: []*Node {
			&Node { 
				Val: "C", 
				Neighbours: []*Node {
					&Node { Val:"D", Neighbours: []*Node {}},
				},
			},}};
	bfs(head, "D");
}