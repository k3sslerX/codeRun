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
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}
	matrix[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i >= 2 && j >= 1 {
				matrix[i][j] += matrix[i-2][j-1]
			}
			if i >= 1 && j >= 2 {
				matrix[i][j] += matrix[i-1][j-2]
			}
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(matrix[n-1][m-1]))
	writer.Flush()
}
