package secret

import (
	"github.com/grimdork/secret/mersenne"
	"github.com/ncw/gmp"
)

// Combine points via Lagrange interpolation to form the secret, provided
// a number of points equivalent to the threshold are supplied.
func Combine(points []Point) []byte {
	max := gmp.NewInt(0)
	for _, n := range points {
		if max.Cmp(n.Y) < 0 {
			max.Set(n.Y)
		}
	}

	prime := mersenne.GetMinimum(max)
	x := gmp.NewInt(0)
	fx := gmp.NewInt(0)
	lagrange := gmp.NewInt(0)
	for i, p1 := range points {
		num, den := gmp.NewInt(1), gmp.NewInt(1)
		for j, p2 := range points {
			if i == j {
				continue
			}

			z := gmp.NewInt(0)
			z.SubUint32(x, p2.X)
			num.Mul(num, z)
			num.Mod(num, prime)

			z.SetUint64(uint64(p1.X))
			z.SubUint32(z, p2.X)
			den.Mul(den, z)
			den.Mod(den, prime)
		}

		lagrange.ModInverse(den, prime)
		lagrange.Mul(lagrange, num)

		fx.Add(fx, prime)
		l := gmp.NewInt(0)
		l.Set(p1.Y)
		l.Mul(l, lagrange)
		fx.Add(fx, l)
		fx.Mod(fx, prime)
	}

	return fx.Bytes()
}
