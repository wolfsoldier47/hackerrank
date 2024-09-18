package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// func migratoryBirds(arr []int32) int {
//     // Write your code here

//     birdMap := map[int]int{
//         1 : 0,
//         2 : 0,
//         3 : 0,
//         4 : 0,
//         5 : 0 ,
//     }

//     for i:=0; i<len(arr);i++ {
//         switch arr[i] {
//         case 1:
//             birdMap[1]++
//         case 2:
//             birdMap[2]++
//         case 3:
//             birdMap[3]++
//         case 4:
//             birdMap[4]++
//         case 5:
//             birdMap[5]++
//         }
//     }
//     // Iterate over the map to find the largest value
//     largest := 0
//     mkey := 0
//     for key, value := range birdMap {
//         if value > largest {
//             largest = value
//             mkey = key
//                 }
//             }
//     return mkey
// }

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
type Types struct {
	type1 int32
	type2 int32
	type3 int32
	type4 int32
	type5 int32
}

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	var types Types

	for _, value := range arr {
		switch value {
		case 1:
			types.type1++
		case 2:
			types.type2++
		case 3:
			types.type3++
		case 4:
			types.type4++
		case 5:
			types.type5++
		}
	}
	b := findHighestType(types)
	return b
}
func findHighestType(types Types) int32 {
	maxValue := types.type1
	var maxType int32 = 1

	if types.type2 > maxValue {
		maxValue = types.type2
		maxType = 2
	}
	if types.type3 > maxValue {
		maxValue = types.type3
		maxType = 3
	}
	if types.type4 > maxValue {
		maxValue = types.type4
		maxType = 4
	}
	if types.type5 > maxValue {
		maxValue = types.type5
		maxType = 5
	}

	return maxType
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := migratoryBirds(arr)

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
