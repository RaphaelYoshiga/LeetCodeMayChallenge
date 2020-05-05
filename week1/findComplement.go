package main;

func findComplement(num int) int {
	maxN := 1
	for num > maxN {
		maxN *= 2
	}

	if maxN == num {
        // Flip leftiest switch to zero, and all rest to 1.
		return maxN - 1
	} else {
		return maxN - 1 - num
	}
}