package main;

type Person struct{
	I int
	TrustedBy []int
	Trusts []int
}

func findJudge(N int, trust [][]int) int {

	if N == 1{
		return 1;
	}

	peopleMap := make(map[int]*Person);

	for _, relation := range trust {

		source:= relation[0];
		trusted := relation[1];

		value, inMap := peopleMap[source];
		if inMap {
			value.Trusts = append(value.Trusts, trusted);
		}else{
			peopleMap[source] = &Person{ I: source, Trusts: []int { trusted }, TrustedBy: []int {}};
		}

		newVal, trustedInMap := peopleMap[trusted];
		if trustedInMap {
			newVal.TrustedBy = append(newVal.TrustedBy, source);
		}else{
			peopleMap[trusted] = &Person{ I: trusted, Trusts: []int { }, TrustedBy: []int { source }};
		}
	}


	for _, value := range peopleMap {

		if len(value.Trusts) == 0 && len(value.TrustedBy) == N - 1{
			return value.I;
		}
	}

	return -1;
}
