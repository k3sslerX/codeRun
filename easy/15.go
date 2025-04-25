package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	cave := make([][][]string, n)
	visited := make([][][]bool, n)
	lengths := make([][][]int, n)
	var startPosition [3]int

	for i := 0; i < n; i++ {
		cave[i] = make([][]string, n)
		visited[i] = make([][]bool, n)
		lengths[i] = make([][]int, n)
		scanner.Scan()
		for j := 0; j < n; j++ {
			cave[i][j] = make([]string, n)
			visited[i][j] = make([]bool, n)
			lengths[i][j] = make([]int, n)
			scanner.Scan()
			parts := strings.Trim(scanner.Text(), "")
			for p := 0; p < len(parts); p++ {
				cave[i][j][p] = string(parts[p])
				if parts[p] == 'S' {
					startPosition[0] = i
					startPosition[1] = j
					startPosition[2] = p

					visited[i][j][p] = true
				}
			}
		}
	}
	queue := make([][3]int, 0)
	queue = append(queue, startPosition)
	for len(queue) > 0 {
		position := queue[0]
		queue = queue[1:]
		for i := 0; i < 3; i++ {
			newPosition := position
			newPosition[i] = position[i] + 1
			if newPosition[0] >= 0 && newPosition[0] < n && newPosition[1] >= 0 && newPosition[1] < n && newPosition[2] >= 0 && newPosition[2] < n {
				if visited[newPosition[0]][newPosition[1]][newPosition[2]] == false && cave[newPosition[0]][newPosition[1]][newPosition[2]] == "." {
					lengths[newPosition[0]][newPosition[1]][newPosition[2]] = lengths[position[0]][position[1]][position[2]] + 1
					visited[newPosition[0]][newPosition[1]][newPosition[2]] = true
					queue = append(queue, newPosition)
				}
			}
			newPosition[i] = position[i] - 1
			if newPosition[0] >= 0 && newPosition[0] < n && newPosition[1] >= 0 && newPosition[1] < n && newPosition[2] >= 0 && newPosition[2] < n {
				if visited[newPosition[0]][newPosition[1]][newPosition[2]] == false && cave[newPosition[0]][newPosition[1]][newPosition[2]] == "." {
					lengths[newPosition[0]][newPosition[1]][newPosition[2]] = lengths[position[0]][position[1]][position[2]] + 1
					visited[newPosition[0]][newPosition[1]][newPosition[2]] = true
					queue = append(queue, newPosition)
				}
			}
		}
	}

	minLength := math.MaxInt
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if lengths[0][i][j] < minLength && visited[0][i][j] {
				minLength = lengths[0][i][j]
			}
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString(strconv.Itoa(minLength))
	_ = writer.Flush()
}
