package main;

func findAnagrams(s string, p string) []int {
    sArr, pArr := [26]int{}, [26]int{}
    ans := []int{}
    for _, v := range p {
        pArr[v - 'a']++
    }
    for i, v := range s {
        if i >= len(p) { sArr[s[i - len(p)] - 'a']-- }
        sArr[v - 'a']++
        if sArr == pArr { ans = append(ans, i - len(p) + 1) }
    }
    return ans
}

// func findAnagrams(s string, p string) []int {
// 	l := len(s);
// 	pLen := len(p);

// 	output := []int {};

// 	pHash := hash(p);
// 	for i, _ := range s{
// 		if i + pLen > l{
// 			break;
// 		}
		
// 		compare:= s[i:i+pLen];
// 		if len(compare) == pLen && pHash == hash(compare){
// 			output = append(output, i);
// 		}
// 	}	

// 	return output;
// }

// var prime = []int {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103};
// func hash(s string) int{
//     hash:= 1;
//     for _, c := range s{
//         hash *= prime[c - 'a'];
//     }
//     return hash;
// }