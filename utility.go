package secret

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ncw/gmp"
)

var randsource = rand.NewSource(time.Now().UnixNano())

func randomNumber(lessthan *gmp.Int) *gmp.Int {
	r := rand.New(randsource)
	n := gmp.NewInt(0)
	n.Set(lessthan)
	n.SubUint32(n, 1)
	n.Rand(r, n)
	return n
}

// ShareToString returns a nicely formatted hexadecimal representation of a share.
func ShareToString(p Point) string {
	return fmt.Sprintf("%d-%x", p.X, p.Y)
}

// ErrInvalidSyntax is returned if a share isn't formatted as a number, hyphen and string.
var ErrInvalidSyntax = errors.New("invalid syntax in share")

// StringToShare will split a share string into a Point.
func StringToShare(s string) (Point, error) {
	a := strings.Split(s, "-")
	if len(a) != 2 {
		return Point{0, nil}, ErrInvalidSyntax
	}

	x, err := strconv.Atoi(a[0])
	if err != nil {
		return Point{0, nil}, err
	}

	y := gmp.NewInt(0)
	y.SetString(a[1], 16)
	return Point{uint32(x), y}, nil
}
