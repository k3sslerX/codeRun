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
	m, _ := strconv.Atoi(parts[1])
	matrix := make([][]int, n)
	costs := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		costs[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		scanner.Scan()
		parts = strings.Fields(scanner.Text())
		for j := 0; j < m; j++ {
			matrix[i][j], _ = strconv.Atoi(parts[j])
		}
	}
	costs[0][0] = matrix[0][0]
	for i := 1; i < n; i++ {
		costs[i][0] = costs[i-1][0] + matrix[i][0]
	}
	for j := 1; j < m; j++ {
		costs[0][j] = costs[0][j-1] + matrix[0][j]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			costs[i][j] = matrix[i][j] + min(costs[i-1][j], costs[i][j-1])
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(costs[n-1][m-1]))
	writer.Flush()
}
