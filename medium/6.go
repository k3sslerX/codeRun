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
	n, _ := strconv.Atoi(scanner.Text())
	nSeq := make([]int, n)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	for i := 0; i < n; i++ {
		nSeq[i], _ = strconv.Atoi(parts[i])
	}
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	mSeq := make([]int, m)
	scanner.Scan()
	parts = strings.Fields(scanner.Text())
	for i := 0; i < m; i++ {
		mSeq[i], _ = strconv.Atoi(parts[i])
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if nSeq[i-1] == mSeq[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	subseq := make([]int, 0)
	i, j := n, m
	for i > 0 && j > 0 {
		if nSeq[i-1] == mSeq[j-1] {
			subseq = append(subseq, nSeq[i-1])
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	for k := len(subseq) - 1; k >= 0; k-- {
		writer.WriteString(strconv.Itoa(subseq[k]))
		if k > 0 {
			writer.WriteString(" ")
		}
	}
	writer.Flush()
}
