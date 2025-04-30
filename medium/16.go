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
	_, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	lines := make([][]int, m)

	for i := 0; i < m; i++ {
		lines[i] = make([]int, 0)
		scanner.Scan()
		parts := strings.Fields(scanner.Text())

		for j := 1; j < len(parts); j++ {
			station, _ := strconv.Atoi(parts[j])
			lines[i] = append(lines[i], station)
		}
	}

	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])

	visited := make([]bool, m)
	queue := make([]int, 0)
	queueReseats := make([]int, 0)
	reseatsFinal := 0
	reach := false

	for i, line := range lines {
		for _, station := range line {
			if station == a {
				queue = append(queue, i)
				visited[i] = true
				queueReseats = append(queueReseats, 0)
			}
		}
	}

	for len(queue) > 0 {
		line := queue[0]
		reseats := queueReseats[0]
		queue = queue[1:]
		queueReseats = queueReseats[1:]
		for _, station := range lines[line] {
			if station == b {
				reach = true
				reseatsFinal = reseats
				break
			}
			for i, queueLine := range lines {
				if !visited[i] {
					for _, stationLine := range queueLine {
						if stationLine == station {
							queue = append(queue, i)
							visited[i] = true
							queueReseats = append(queueReseats, reseats+1)
						}
					}
				}
			}
		}
		if reach {
			break
		}
	}

	reseatsCount := -1
	if reach {
		reseatsCount = reseatsFinal
	}

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString(strconv.Itoa(reseatsCount))
	_ = writer.Flush()
}
