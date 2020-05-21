package main;


var (
	digitsMap = make(map[int]bool);
)

func getDigits(a int) []int{
	result := []int {};
	for a > 0{
		b := a%10;
		a = a/10;
		result = append(result, b);
	}
	return result;
}

func powerset(a []int, used []bool, x int) bool{
	if x == len(used) - 1{
		used[x] = false;
		b := insertInHash(a, used);
		if b {
			used[x] = true;
			return insertInHash(a, used);
		} else{
			return false;
		}
	}

	used[x] = false;
	p := powerset(a, used, x+1);
	used[x] = true;		
	return p && powerset(a, used, x+1);
}

func insertInHash(a []int, used []bool) bool{
	sum:= 0;
	for i, isUsed := range used{
		if isUsed{
			if sum == 0{
				sum = a[i];
			}else{
				sum *= a[i];
			}
		}
	}

	_,ok := digitsMap[sum];
	if ok {
		return false;
	}

	digitsMap[sum] = true;
	return true;
}

func colorfulNumber(a int) bool{
	digits := getDigits(a);
	used := make([]bool, len(digits));
	return powerset(digits, used, 0);
}

func main(){
	colorfulNumber(326);
}