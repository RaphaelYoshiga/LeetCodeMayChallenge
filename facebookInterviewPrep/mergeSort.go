package main;

func mergeSort(a []int) []int{
	temp := make([]int, len(a));

	mergeSortWorker(a, temp, 0, len(a) -1);

	return temp;
}

func mergeSortWorker(a []int, temp []int, left int, right int){
	if left >= right{
		return;
	}

	middle:= (left + right) / 2;
	mergeSortWorker(a, temp, left, middle);
	mergeSortWorker(a, temp, middle + 1, right);

	mergeHalves(a, temp, left, right);
}

func mergeHalves(a []int, temp []int, leftStart int, rightEnd int){
	leftEnd:= (leftStart + rightEnd) / 2;
	rightStart := leftEnd + 1;

	right:= rightStart;
	left := leftStart;
	index := leftStart;

	for left <= leftEnd && right <= rightEnd{
		if a[left] <= a[right]{
			temp[index] = a[left];
			left++;
		}else{
			temp[index] = a[right];
			right++;
		}
		index++;
	}

	for left <= leftEnd{
		temp[index] = a[left];
		left++;
		index++;
	}

	for right<= rightEnd{
		temp[index] = a[right];
		right++;
		index++;
	}

	i:= leftStart;
	size := rightEnd - leftStart + 1;
	for i < leftStart + size{
		a[i] = temp[i];
		i++;
	}
}
func main(){
	mergeSort([]int {4, 3, 2, 1});
}