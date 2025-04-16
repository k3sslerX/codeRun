package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printMatrix(matrix [][]int, prices []int, n int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		writer.WriteString(fmt.Sprintf("%3d", prices[i]) + " | ")
		for j := 0; j < n; j++ {
			writer.WriteString(fmt.Sprintf("%3d", matrix[i][j]) + " ")
		}
		writer.WriteString("\n")
	}
	writer.Flush()
}

//5
//110
//40
//120
//60
//90

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	prices := make([]int, n)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		prices[i], _ = strconv.Atoi(scanner.Text())
		matrix[i] = make([]int, n)
	}
	var coupons []int
	//usedCoupons := 0
	matrix[0][0] = prices[0]
	if prices[0] > 100 {
		coupons = append(coupons, 0)
	}
	for i := 1; i < n; i++ {
		matrix[i][0] = matrix[i-1][0] + prices[i]
		if prices[i] > 100 {
			coupons = append(coupons, i)
		}
	}
	for i := 1; i < n; i++ {
		for _, coupon := range coupons {
			for j := 1; j < i-coupon+1; j++ {
				if matrix[i-1][j] == 0 {
					matrix[i][j] = matrix[i-1][0]
					matrix[0][j] = i
				} else {
					if matrix[0][j] > len(coupons) && len(coupons) > 1 {
						matrix[i][j] = matrix[i-1][j]
					} else {
						matrix[i][j] = matrix[i-1][j] + prices[i]
					}
				}
			}
		}
	}
	minSum := matrix[n-1][0]
	minSumInd := 0
	for i := 0; i < n; i++ {
		if matrix[n-1][i] < minSum && matrix[n-1][i] != 0 {
			minSum = matrix[n-1][i]
			minSumInd = i
		}
	}
	couponsUsed := 0
	couponsUsedDays := ""
	for i := n - 1; i > 0; i-- {
		if matrix[i-1][minSumInd] == 0 {
			couponsUsed++
			couponsUsedDays = strconv.Itoa(i+1) + " " + couponsUsedDays
			break
		}
		if matrix[i][minSumInd]-matrix[i-1][minSumInd] == 0 {
			couponsUsed++
			couponsUsedDays = strconv.Itoa(i+1) + " " + couponsUsedDays
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(minSum) + "\n")
	writer.WriteString(strconv.Itoa(len(coupons)-couponsUsed) + " " + strconv.Itoa(couponsUsed) + "\n")
	writer.WriteString(couponsUsedDays)
	writer.Flush()
	//printMatrix(matrix, prices, n)
}
