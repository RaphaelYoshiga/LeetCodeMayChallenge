package main;

import "sort";

type kv struct {
	Key   rune
	Value int
}

func frequencySort(s string) string {
	hashMap := make(map[rune]int);
    for _, c := range s{
		hashMap[c]++;
	}

    var ss []kv
    for k, v := range hashMap {
        ss = append(ss, kv{Key:k, Value:v})
	}
	
	
    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value > ss[j].Value
	});

	result := make([]rune, len(s));
	i := 0;
	for _, kv := range ss{
		for kv.Value > 0{
			result[i] = kv.Key;
			i++;
			kv.Value--;
		}
	}

	return string(result);
}