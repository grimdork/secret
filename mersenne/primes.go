// Package mersenne returns Mersenne primes as libgmp numbers.
package mersenne

import (
	"github.com/ncw/gmp"
)

var (
	validExponents = []uint{
		2, 3, 5, 7, 13,
		17, 19, 31, 61, 89,
		107, 127, 521, 607, 1279,
		2203, 2281, 3217, 4253, 4423,
		9689, 9941, 11213, 19937, 21701,
		23209, 44497, 86243, 110503, 132049,
		216091, 756839, 859433, 1257787, 1398269,
		2976221, 3021377, 6972593, 13466917, 20996011,
		24036583, 25964951, 30402457, 32582657, 37156667,
		42643801, 43112609, 57885161, 74207281, 77232917,
		82589933,
	}

	// Digits is the number of digits when printed as decimal.
	Digits = []uint{
		1, 1, 2, 3, 4,
		6, 6, 10, 19, 27,
		33, 39, 157, 183, 386,
		664, 687, 969, 1281, 1332,
		2917, 2993, 3376, 6002, 6533,
		6987, 13395, 25962, 33265, 39751,
		65050, 227832, 258716, 378632, 420921,
		895932, 909526, 2098960, 4053946, 6320430,
		7235733, 786230, 9152052, 9808358, 11185272,
		12837064, 12978189, 17425170, 22338618, 23249425,
		24862048,
	}

	// HexDigits is the number of digits when printed as headecimal.
	HexDigits = []uint{
		1, 1, 2, 2, 4,
		5, 5, 8, 16, 23,
		27, 32, 131, 152, 320,
		551, 571, 805, 1064, 1106,
		2423, 2486, 2804, 4985, 5426,
		5803, 1112, 2156, 2762, 33013,
		5402, 18920, 21489, 31447, 349568,
		7440, 7553, 17439, 33660, 5249003,
		60096, 64918, 76005, 81455, 9289167,
		106691, 107713, 144721, 185581, 19308230,
		20647484,
	}
)

const (
	// FIRST known Mersenne prime's position in the table.
	FIRST = 1
	// LAST known Mersenne prime's position in the table.
	// The 51st prime was found in October 2018.
	LAST = 51
)

var one = gmp.NewInt(1)

// Get prime n from the table.
func Get(n uint) *gmp.Int {
	p := gmp.NewInt(0)
	p.Lsh(one, validExponents[n-1])
	p.Sub(p, one)
	return p
}

// GetMinimum returns the first prime <= n.
// Returns nil if none are big enough.
func GetMinimum(n *gmp.Int) *gmp.Int {
	var i uint
	for i = FIRST; i < LAST; i++ {
		p := Get(i)
		if n.Cmp(p) <= 0 {
			return p
		}
	}

	return nil
}
