package main;

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote)>len(magazine) {
        return false
	}

	var noteLetters [128]byte;
	count := 0;

	for _, c := range ransomNote{
		noteLetters[c]++;
		count++;		
	}

	for _, c := range magazine {
		val := noteLetters[c];
		if val > 0{
			noteLetters[c]--;
			count--;
		}else if count == 0{
			return true;
		}
	}


	return count == 0;
    
}