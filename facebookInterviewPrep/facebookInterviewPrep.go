package main;

var (
    hashMap = make(map[rune]rune)
)

func init(){
	hashMap['{'] = '}';
	hashMap['('] = ')';
	hashMap['['] = ']';
}

func isBalanced(s string) string {
	stack := []rune {};

	for _, c := range s{
		_, ok := hashMap[c];
		if ok {
			stack = append(stack, c);
		}else{

			stackLen := len(stack);			
			if stackLen <= 0{
				return "NO"
			}
			lastSpecial := stack[stackLen-1];
			closing := hashMap[lastSpecial];
			if closing != c {
				return "NO";
			}
			stack = stack[0:stackLen - 1];
		}
	}

	if len(stack) == 0{
		return "YES"
	}
	return "NO";
}

func main(){
	result := isBalanced("[(})])}{}}]{({[]]]))]})]");
	print(result);
}