package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	adjMatrix := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		parts = strings.Fields(scanner.Text())
		adjMatrix[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			adjMatrix[i][j], _ = strconv.Atoi(parts[j-1])
		}
	}
	scanner.Scan()
	parts = strings.Fields(scanner.Text())
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	visited := make([]bool, n+1)
	distances := make([]int, n+1)
	routes := make([]int, n+1)
	queue := make([]int, 0)
	found := false
	queue = append(queue, a)
	visited[a] = true
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		if vertex == b {
			found = true
			break
		}
		for i := 1; i <= n; i++ {
			if adjMatrix[vertex][i] == 1 && !visited[i] {
				visited[i] = true
				distances[i] = distances[vertex] + 1
				routes[i] = vertex
				queue = append(queue, i)
			}
		}
	}
	minLength := -1
	route := ""
	if found {
		minLength = distances[b]
		if minLength > 0 {
			for v := b; v != a; v = routes[v] {
				route = strconv.Itoa(v) + " " + route
			}
			route = strconv.Itoa(a) + " " + route
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(minLength) + "\n")
	writer.WriteString(route)
	writer.Flush()
}
