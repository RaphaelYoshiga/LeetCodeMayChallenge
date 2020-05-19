package main;

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    
    countsS1 := [26]int{}
    for _, c := range s1 {
        countsS1[c-'a']++
    }
    
    i := 0
    countsWin := [26]int{}
    for j := 0; j < len(s2); j++ {
        countsWin[s2[j]-'a']++
    
        if j-i+1 > len(s1) {
            countsWin[s2[i]-'a']--
            i++
        }

        if countsS1 == countsWin {
            return true
        }
        
    }
    return false
}

// func checkInclusion(s1 string, s2 string) bool {
// 	s1Len := len(s1);

// 	if s1Len > len(s2){
// 		return false;
// 	}

// 	compare := s2[0:s1Len];

// 	expectedHash := hash(s1);
// 	compareHash:= hash(compare);
// 	if expectedHash == compareHash{
// 		return true;
// 	}

// 	for i := s1Len; i < len(s2); i++{

// 		compare = compare[1:] + string(s2[i]);
// 		compareHash = hash(compare);

// 		if expectedHash == compareHash{
// 			return true;
// 		}
// 	} 
// 	return false;
// }

// var prime = []int {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103};
// func hash(s string) int{
//     hash:= 1;
//     for _, c := range s{
//         hash *= prime[c - 'a'];
//     }
//     return hash;
// }
