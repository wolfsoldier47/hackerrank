package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	// "sort"
)

// There is a large pile of socks that must be paired by color. Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.

// Example
// n = 7
// ar = [1,2,1,2,1,3,2]

// There is one pair of color  and one of color . There are three odd socks left, one of each color. The number of pairs is .
// 2
/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */
func contains(item int, list []int) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
func sockMerchant(n int, ar []int) int {
	// Write your code here
	uniqueSock := []int{}

	for _, item := range ar {
		if !contains(item, uniqueSock) {
			uniqueSock = append(uniqueSock, item)
		}
	}
	count := 0
	pairs := 0
	for i := 0; i < len(uniqueSock); i++ {
		for _, item := range ar {
			if item == uniqueSock[i] {
				count++
			}
		}
		pairs += count / 2
		count = 0
	}

	return pairs

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)

	arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ar []int

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)
	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
