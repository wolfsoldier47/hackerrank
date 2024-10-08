package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */
// Sample Input

// 8
// UDDDUDUU
// Sample Output

// 1
// Explanation

// If we represent _ as sea level, a step up as /, and a step down as \, the hike can be drawn as:

// _/\      _
//
//	\    /
//	 \/\/
//
// The hiker enters and leaves one valley.
func countingValleys(steps int32, path string) int32 {
	// Write your code here
	valley := 0
	sealevel := 0
	onmountains := true
	for i := 0; i < int(steps); i++ {
		if path[i] == 'U' {
			sealevel++
			if sealevel == 0 && onmountains == false {
				valley++
			}
		} else if path[i] == 'D' {
			sealevel--
			if sealevel == 0 && onmountains == false {
				valley++
			}
		}
		if sealevel > 0 {
			onmountains = true
		} else if sealevel < 0 {
			onmountains = false
		}
	}

	return int32(valley)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	steps := int32(stepsTemp)

	path := readLine(reader)

	result := countingValleys(steps, path)

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
