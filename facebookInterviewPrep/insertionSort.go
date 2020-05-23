package main

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
	}
}

func insertAt(a []int32, i int, value int32) []int32{
	return append(a[:i], append([]int32{value}, a[i:]...)...);
}

func removeAt(a []int32, i int) []int32{
	return append(a[:i], a[i+1:]...)
}

