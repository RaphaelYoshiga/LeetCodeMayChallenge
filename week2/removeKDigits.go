package main;

func removeKdigits(num string, k int) string {

	out := []rune(num)
	for k > 0{

		nextI := 1;
		
		if len(out) == 1{
			return "0";
		}

		out = handle(out, nextI)
		k--;
	}

	for {
		if len(out) == 1 && out[0] == '0'{
			return "0";
		}

		if out[0] != '0'{
			break;
		}

		out = delChar(out, 0)
	}
	return string(out)
}	

func handle(out []rune, nextI int) []rune{
	currentI := nextI-1;
	current := out[currentI];

	for current == '0'{
		out = delChar(out, currentI);
		current = out[currentI];
	}

	if nextI >= len(out) {
		return delChar(out, currentI)
	}else if out[1] == '0' || current > out[nextI]{
		return delChar(out, currentI)
	}else {
		return handle(out, nextI + 1);
	}
}	

func delChar(s []rune, index int) []rune {
 	return append(s[0:index], s[index+1:]...)
}

