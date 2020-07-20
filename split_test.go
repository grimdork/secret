package secret_test

import (
	"math"
	"strings"
	"testing"

	"github.com/ncw/gmp"

	"github.com/grimdork/secret"
)

const testSecret = "secret string"

func TestSplit(t *testing.T) {
	t.Logf("Input string: '%s'", testSecret)
	points := secret.Split([]byte(testSecret), 2, 4)
	for _, p := range points {
		t.Logf("%s", secret.ShareToString(p))
	}
	t.Log("\nCombining with these points:")
	points2 := points[2:]
	t.Logf("%s", secret.ShareToString(points2[0]))
	t.Logf("%s", secret.ShareToString(points2[1]))
	s := secret.Combine(points2)
	t.Logf("Output string: '%s'", s)
	if strings.Compare(string(s), testSecret) != 0 {
		t.Errorf("Error: result does not match input: %s", s)
		t.FailNow()
	}
}

func TestConversion(t *testing.T) {
	n := gmp.NewInt(math.MaxInt64)
	t.Logf("Max int64 is %d", n.Int64())
	p := secret.Point{1, n}
	s := secret.ShareToString(p)
	t.Logf("ShareToString(): %s", s)
	p, err := secret.StringToShare(s)
	if err != nil {
		t.Errorf("Error converting string back: %s", err.Error())
		t.FailNow()
	}

	if p.Y.Int64() != math.MaxInt64 {
		t.Errorf("Expected %d, got %d", math.MaxInt64, p.Y.Int64())
		t.FailNow()
	}
}
