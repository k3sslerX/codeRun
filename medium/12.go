package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type edgeStruct struct {
	a int
	b int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	vertices := make([]int, 0)
	edges := make([]edgeStruct, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		for j := i; j < n; j++ {
			if parts[j] == "1" {
				vertices = append(vertices, i+1)
				vertices = append(vertices, j+1)
				edges = append(edges, edgeStruct{i + 1, j + 1})
			}
		}
	}
	scanner.Scan()
	//parts := strings.Fields(scanner.Text())
	//a, _ := strconv.Atoi(parts[0])
	//b, _ := strconv.Atoi(parts[1])
}
