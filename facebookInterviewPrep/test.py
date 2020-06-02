graph = {
  'A' : ['B','C', 'D'],
  'B' : ['D', 'E'],
  'C' : ['F'],
  'D' : [],
  'E' : ['F'],
  'F' : []
}

visited = [] # List to keep track of visited nodes.
queue = []     #Initialize a queue

def bfs(visited, graph, node):
  visited.append(node)
  queue.append(node)
  
  i = 0
  
  while queue:
    s = queue.pop(0) 
    print (s + " \n Level:" + str(i), end = " ") 

    for neighbour in graph[s]:
      if neighbour not in visited:
        visited.append(neighbour)
        queue.append(neighbour)
    
    i+= 1

# Driver Code
bfs(visited, graph, 'A')