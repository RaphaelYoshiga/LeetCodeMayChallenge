package main;

func numJewelsInStones(J string, S string) int {
	jewelsMap := make(map[rune]bool);

	for _, c := range J{
		jewelsMap[c] = true;
	}

	count := 0;

	for _, c := range S{
		_, inMap := jewelsMap[c];
		if inMap {
			count++;
		}
	}
	return count;
}

func main(){}