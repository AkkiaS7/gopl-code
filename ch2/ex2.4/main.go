// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	ans := 0
	for i := 0; i < 64; i++ {
		ans += int(x & 1)
		x >>= 1
	}
	return ans
}

//!-
