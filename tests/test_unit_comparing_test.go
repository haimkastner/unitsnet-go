package unitsnet_go_test

import (
	"testing"

	"github.com/haimkastner/unitsnet-go/unitsnet_go"
)

func TestUnitComparing(t *testing.T) {
	lf := unitsnet_go.LengthFactory{}
	length1, _ := lf.FromMeters(10)
	length2, _ := lf.FromDecimeters(100)
	length3, _ := lf.FromMeters(3)

	t.Run("TestEqual", func(t *testing.T) {
		if !length1.Equals(length2) {
			t.Errorf("Expected length1 (%v) to equal length2 (%v)", length1, length2)
		}
	})

	t.Run("TestNotEqual", func(t *testing.T) {
		if length1.Equals(length3) {
			t.Errorf("Expected length1 (%v) not to equal length3 (%v)", length1, length3)
		}
	})

	t.Run("TestBigger", func(t *testing.T) {
		if length1.CompareTo(length3) <= 0 {
			t.Errorf("Expected length1 (%v) to be greater than length3 (%v)", length1, length3)
		}
	})

	t.Run("TestSmaller", func(t *testing.T) {
		if length3.CompareTo(length1) >= 0 {
			t.Errorf("Expected length3 (%v) to be less than length1 (%v)", length3, length1)
		}
	})

	t.Run("TestSmallerOrEqual", func(t *testing.T) {
		if length1.CompareTo(length2) > 0 {
			t.Errorf("Expected length1 (%v) to be less than or equal to length2 (%v)", length1, length2)
		}
	})

	t.Run("TestBiggerOrEqual", func(t *testing.T) {
		if length1.CompareTo(length2) < 0 {
			t.Errorf("Expected length1 (%v) to be greater than or equal to length2 (%v)", length1, length2)
		}
	})
}
