package main;

func checkInclusion(s1 string, s2 string) bool {
	s1Len := len(s1);

	compare := s2[0:s1Len];

	expectedHash := permutationHash(s1);
	compareHash:= permutationHash(compare);
	if expectedHash == compareHash{
		return true;
	}

	for i := s1Len - 1; i < len(s2); i++{

		remove := compare[0];
		compare = append(compare[1:], s2[i]);
		compareHash = compareHash - (remove - 'a' + 1) + (s2[i] - 'a' + 1); 

		if expectedHash == compareHash{
			return true;
		}
	} 
	return false;
}

func permutationHash(s string) int{

 	sum := 0;
	for _, c := range s{
		sum+= c - 'a' + 1;
	}
	return i;
}

func main(){
	checkInclusion("aa", "abaa")
}