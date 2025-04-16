package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func sort(slice *[]int) {
	for i := 0; i < len(*slice); i++ {
		for j := i; j < len(*slice); j++ {
			if (*slice)[j] < (*slice)[i] {
				(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
			}
		}
	}
}

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
		adjList[v] = append(adjList[v], u)
	}
	visited := make([]bool, n+1)
	var component []int
	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true
		component = append(component, u)
		for _, v := range adjList[u] {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	dfs(1)
	sort(&component)
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(len(component)) + "\n")
	for _, v := range component {
		writer.WriteString(strconv.Itoa(v) + " ")
	}
	writer.Flush()
}
