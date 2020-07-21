package secret

import (
	"github.com/grimdork/secret/mersenne"
	"github.com/ncw/gmp"
)

// Point in the polynomial.
type Point struct {
	X uint32
	Y *gmp.Int
}

// SplitSecret into n parts, requiring threshold t parts to recover it.
// secret should be relatively small, otherwise the resulting shares end up huge.
// threshold must be minimum 2, and will be silently increased to 2 if less.
// n should be higher than t, and will be corrected to t+1 if not.
//
// The resulting slice of points will be numbered from 1 and up.
func Split(secret []byte, t, n uint) []Point {
	return FixedSplit(secret, t, n, nil)
}

// FixedSplit works like Split, with the option to specify a prime to use.
// This allows uniform share lengths across different input sizes. The prime still
// needs to be larger than the secret, and will be adjusted accordingly.
func FixedSplit(secret []byte, t, n uint, prime *gmp.Int) []Point {
	s := gmp.NewInt(0)
	s.SetBytes([]byte(secret))
	if prime == nil {
		prime = mersenne.GetMinimum(s)
	} else {
		if prime.Cmp(s) < 0 {
			prime = mersenne.GetMinimum(s)
		}
	}

	if t < 2 {
		t = 2
	}

	if n < t {
		n = t + 1
	}

	coeff := randomPolynomial(t-1, s, prime)
	return getPoints(coeff, n, prime)
}

func randomPolynomial(degree uint, secret, prime *gmp.Int) []*gmp.Int {
	coeff := []*gmp.Int{secret}
	coeff[0] = secret
	var i uint
	for i = 1; i <= degree; i++ {
		x := randomNumber(prime)
		coeff = append(coeff, x)
	}
	return coeff
}

func getPoints(coeff []*gmp.Int, count uint, prime *gmp.Int) []Point {
	points := []Point{}
	var x uint
	for x = 1; x <= count; x++ {
		y := gmp.NewInt(0)
		y.Set(coeff[0])
		for i := 1; i < len(coeff); i++ {
			exp := gmp.NewInt(int64(x))
			exp.Exp(exp, gmp.NewInt(int64(i)), nil)
			exp.Mod(exp, prime)

			term := gmp.NewInt(0)
			term.Set(coeff[i])
			term = term.Mul(term, exp)
			term = term.Mod(term, prime)

			y = y.Add(y, term)
			y.Mod(y, prime)
		}
		p := Point{uint32(x), y}
		points = append(points, p)
	}
	return points
}
