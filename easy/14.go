package main

import (
	"bufio"
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
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	s, _ := strconv.Atoi(parts[2])
	t, _ := strconv.Atoi(parts[3])
	q, _ := strconv.Atoi(parts[4])
	coordinates := make([][]int, q)
	for i := 0; i < q; i++ {
		coordinates[i] = make([]int, 2)
		scanner.Scan()
		parts = strings.Fields(scanner.Text())
		coordinates[i][0], _ = strconv.Atoi(parts[0])
		coordinates[i][1], _ = strconv.Atoi(parts[1])
	}
	jumps := make([][]int, n+1)
	visited := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		jumps[i] = make([]int, m+1)
		visited[i] = make([]bool, m+1)
	}
	visited[s][t] = true
	queueX := make([]int, 0)
	queueY := make([]int, 0)
	queueX = append(queueX, s)
	queueY = append(queueY, t)
	for len(queueX) > 0 && len(queueY) > 0 {
		x := queueX[0]
		queueX = queueX[1:]
		y := queueY[0]
		queueY = queueY[1:]
		for p := -2; p <= 2; p++ {
			if p == 0 {
				continue
			}
			for o := -2; o <= 2; o++ {
				if o == 0 || abs(p) == abs(o) {
					continue
				}
				if x+p <= n && x+p > 0 && y+o <= m && y+o > 0 {
					if !visited[x+p][y+o] {
						visited[x+p][y+o] = true
						jumps[x+p][y+o] = jumps[x][y] + 1
						queueX = append(queueX, x+p)
						queueY = append(queueY, y+o)
					}
				}
			}
		}
	}
	totalJumps := 0
	all := true
	for i := 0; i < q; i++ {
		x := coordinates[i][0]
		y := coordinates[i][1]
		if !visited[x][y] {
			all = false
			break
		} else {
			totalJumps += jumps[x][y]
		}
	}
	if !all {
		totalJumps = -1
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(totalJumps))
	writer.Flush()
}
