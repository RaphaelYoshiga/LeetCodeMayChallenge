package main;

func singleNonDuplicate(nums []int) int {
	left:= 0;
	right := len(nums) - 1;

	for{
		
		if left >= right -1{

			if left % 2 == 0{
				return nums[left];
			}
			return nums[right];
		}

		check := ((right - left) / 2) + left;

		toCompare := toCompareCalc(check);

		if nums[toCompare] != nums[check]{
			right = check;
		}else{
			left = check;
		}
	}
}

func toCompareCalc(i int) int{
	if i % 2 == 0{
		return i + 1;
	}else{
		return i - 1;
	}
}
