package main
import (
    "bufio"
    "fmt"
    "io"

    "strings"
)

// func main() {

// 	test := []int{ 5, 8, 1, 3, 7, 9, 2};
// 	quickSortWork(test, 0, len(test) -1)
//     reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

//     nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//     checkError(err)
//     n := int32(nTemp)

//     arrTemp := strings.Split(readLine(reader), " ")

//     var arr []int

//     for i := 0; i < int(n); i++ {
//         arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
//         checkError(err)
//         arrItem := int(arrItemTemp)
//         arr = append(arr, arrItem)
//     }

//     quickSortWork(arr, 0, len(arr) - 1);
// }

func quickSortWork(arr []int, left int, right int){
	if left >= right{
		return;
	}

	if left < right -1{
		fmt.Println(arrayToString(arr[left:right], " "));
	}

	pivot := arr[left];

	index := partition(arr, left, right, pivot);
	quickSortWork(arr, left, index - 1);
	quickSortWork(arr, index, right);

	
}

func arrayToString(a []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func partition(arr []int, left int, right int, pivot int) int{
	for left <= right{
		for arr[left] < pivot{
			left++;
		}

		for arr[right] > pivot{
			right--;
		}

		if left <= right{
			swap(arr, left, right);
			left++;
			right--;
		}
	}
	return left;
}

func swap(arr []int, left int, right int){
	aux := arr[left];
	arr[left] = arr[right];
	arr[right] = aux;
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
