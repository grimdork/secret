package secret_test

import (
	"math"
	"strings"
	"testing"

	"github.com/grimdork/secret"
	"github.com/grimdork/secret/mersenne"
	"github.com/ncw/gmp"
)

const testSecret = "secret string"

// TestSplit runs Split() and verifies the Combine() output.
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

// TestConversion() verifies that ShareToString() and StringToShare() reverse each other's output.
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

// TestFixed intentionally tries a too small prime and checks that FixedSplit() does the correct thing.
func TestFixed(t *testing.T) {
	points := secret.FixedSplit([]byte(testSecret+testSecret), 2, 5, mersenne.Get(9))
	for _, p := range points {
		t.Logf("%s", secret.ShareToString(p))
	}
	points = []secret.Point{points[1], points[3]}
	t.Logf("%#v", points)
	s := secret.Combine(points)
	t.Logf("Output string: '%s'", s)
	if strings.Compare(string(s), testSecret+testSecret) != 0 {
		t.Errorf("Error: result does not match input: %s", s)
		t.FailNow()
	}
}
