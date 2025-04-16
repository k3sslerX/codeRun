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
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	adjList := make([][]int, n+1)
	for i := 0; i < m; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		adjList[u] = append(adjList[u], v)
	}
	visited := make([]bool, n+1)
	inStack := make([]bool, n+1)
	var stack []int
	hasCycle := false
	var dfs func(int) bool
	dfs = func(u int) bool {
		if inStack[u] {
			return true
		}
		if visited[u] {
			return false
		}
		visited[u] = true
		inStack[u] = true
		for _, v := range adjList[u] {
			if dfs(v) {
				return true
			}
		}
		inStack[u] = false
		stack = append(stack, u)
		return false
	}
	for u := 1; u <= n; u++ {
		if !visited[u] {
			if dfs(u) {
				hasCycle = true
				break
			}
		}
	}
	if hasCycle {
		fmt.Println(-1)
	} else {
		for i := len(stack) - 1; i >= 0; i-- {
			fmt.Print(stack[i], " ")
		}
		fmt.Println()
	}
}
