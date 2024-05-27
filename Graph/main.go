package main

import "fmt"

// Solution struct
type Solution struct{}

// dfs function to perform depth-first search
func (s *Solution) dfs(node int, adj [][]int, vis []bool, ls *[]int) {
	vis[node] = true
	*ls = append(*ls, node)
	for _, it := range adj[node] {
		if !vis[it] {
			s.dfs(it, adj, vis, ls)
		}
	}
}

// dfsOfGraph function to return a list containing the DFS traversal of the graph
func (s *Solution) dfsOfGraph(V int, adj [][]int) []int {
	vis := make([]bool, V)
	start := 0
	var ls []int
	s.dfs(start, adj, vis, &ls)
	return ls
}

// bfsOfGraph function to return Breadth First Traversal of the given graph
func (s *Solution) bfsOfGraph(V int, adj [][]int) []int {
	vis := make([]bool, V)
	vis[0] = true
	q := []int{0}

	var bfs []int
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		bfs = append(bfs, node)

		for _, it := range adj[node] {
			if !vis[it] {
				vis[it] = true
				q = append(q, it)
			}
		}
	}

	return bfs
}

func main() {
	V := 5

	adj := make([][]int, V)
	for i := range adj {
		adj[i] = []int{}
	}

	edges := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{2, 4},
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u) // Add this line for undirected graph
	}

	s := Solution{}
	fmt.Println("---------dfs----------")

	ans := s.dfsOfGraph(V, adj)
	for _, val := range ans {
		fmt.Print(val, " ")
	}

	fmt.Println()
	fmt.Println("---------bfs----------")

	adj2 := make([][]int, V)
	for i := range adj {
		adj2[i] = []int{}
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj2[u] = append(adj2[u], v)
		// adj2[v] = append(adj2[v], u) // Add this line for undirected graph
	}

	ans = s.bfsOfGraph(V, adj2)
	for _, val := range ans {
		fmt.Print(val, " ")
	}

	fmt.Println()
}
