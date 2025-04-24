package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func findCubes(num int) (cubes []int) {
	for i := 1; i*i*i <= num; i++ {
		cubes = append(cubes, i*i*i)
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())

	cubes := findCubes(num)

	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		dp[i] = math.MaxInt
		for _, c := range cubes {
			if c > i {
				continue
			}
			if dp[i-c]+1 < dp[i] {
				dp[i] = dp[i-c] + 1
			}
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString(strconv.Itoa(dp[num]))
	_ = writer.Flush()
}
