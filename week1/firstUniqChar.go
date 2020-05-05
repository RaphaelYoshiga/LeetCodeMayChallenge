package main;

type Char struct{
	i int
	count int
}

func firstUniqChar(s string) int {
	runeMap := make(map[rune]*Char)

	for i, c := range s {
		mapItem, ok := runeMap[c];
		if ok{
			mapItem.count++;
		}else {
			runeMap[c] = &Char { i: i, count: 1  }
		}
	}

	oneItem := false;
	firstIndex := 9999999;
	for _, mapItem := range runeMap{
		if mapItem.count == 1{
			if mapItem.i < firstIndex{
				firstIndex = mapItem.i;
				oneItem = true;
			}
		}
	}
	if oneItem{
		return firstIndex;
	}
	return -1;
}