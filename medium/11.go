package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	adjMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		adjMatrix[i] = make([]int, n)
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		for j := 0; j < n; j++ {
			adjMatrix[i][j], _ = strconv.Atoi(parts[j])
		}
	}
	visited := make([]bool, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
	var cycle []int
	var dfs func(int, int) bool
	dfs = func(u, p int) bool {
		visited[u] = true
		for v := 0; v < n; v++ {
			if adjMatrix[u][v] == 1 {
				if !visited[v] {
					parent[v] = u
					if dfs(v, u) {
						return true
					}
				} else if v != p {
					cycle = append(cycle, v)
					for x := u; x != v; x = parent[x] {
						cycle = append(cycle, x)
					}
					cycle = append(cycle, v)
					return true
				}
			}
		}
		return false
	}
	for u := 0; u < n; u++ {
		if !visited[u] {
			if dfs(u, -1) {
				break
			}
		}
	}
	if len(cycle) == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		fmt.Println(len(cycle) - 1)
		for i := len(cycle) - 1; i > 0; i-- {
			fmt.Print(cycle[i]+1, " ")
		}
		fmt.Println()
	}
}
