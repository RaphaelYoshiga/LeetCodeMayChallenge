package main;

import "math";

func isPerfectSquare(num int) bool {
	sq := sqrt(float64(num));
	

	sq = math.Round(sq*10000) / 10000;
	return sq == float64(int64(sq));
}

const DELTA = 0.0000001
const INITIAL_Z = 100.0

func sqrt(x float64) (z float64) {
    z = INITIAL_Z
    
    step := func() float64 {
    	return z - (z*z - x) / (2 * z)
    }
    
    for zz := step(); math.Abs(zz - z) > DELTA; {
    	z = zz
		zz = step()
    }
    return
}
