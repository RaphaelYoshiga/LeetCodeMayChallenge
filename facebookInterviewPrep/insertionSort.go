package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the insertionSort2 function below.
func insertionSort2(n int32, arr []int32) {

	for i, current := range arr{
		if i == 0{
			continue;
		}

		for y:= 0; y < i; y++{
			comparing := arr[y];
			if current < comparing{
				arr = insertAt(arr, y, current);
				arr = removeAt(arr, i + 1)
				break;
			}
		}
		fmt.Println(arrayToString(arr, " "));
	}
}

func insertAt(a []int32, i int, value int32) []int32{
	return append(a[:i], append([]int32{value}, a[i:]...)...);
}

func removeAt(a []int32, i int) []int32{
	return append(a[:i], a[i+1:]...)
}

func arrayToString(a []int32, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}


func main() {
	insertionSort2(6, []int32 {1,4,3,5,6,2});

    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    insertionSort2(n, arr)
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
