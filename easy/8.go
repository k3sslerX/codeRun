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
	verticesAmount, _ := strconv.Atoi(parts[0])
	edgesAmount, _ := strconv.Atoi(parts[1])
	adjList := make([][]int, verticesAmount+1)
	for i := 0; i < edgesAmount; i++ {
		scanner.Scan()
		parts = strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
	}
	var outStrings []string
	visited := make([]bool, verticesAmount+1)
	for i := 1; i < verticesAmount+1; i++ {
		var components []int
		var dfs func(int)
		dfs = func(u int) {
			visited[u] = true
			components = append(components, u)
			for _, v := range adjList[u] {
				if !visited[v] {
					dfs(v)
				}
			}
		}
		if !visited[i] {
			dfs(i)
			outString := strconv.Itoa(len(components)) + "\n"
			for j := 0; j < len(components); j++ {
				outString += strconv.Itoa(components[j]) + " "
			}
			outStrings = append(outStrings, outString)
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(len(outStrings)) + "\n")
	for _, str := range outStrings {
		writer.WriteString(str + "\n")
	}
	writer.Flush()
}
