package unitsnet_go_test

import (
	"testing"

	"github.com/haimkastner/unitsnet-go/unitsnet_go"
)

func TestUnitArithmetics(t *testing.T) {
	lf := unitsnet_go.LengthFactory{}
	length1, _ := lf.FromMeters(10)
	length2, _ := lf.FromMeters(3)

	t.Run("TestAdd", func(t *testing.T) {
		result := length1.Add(length2)
		if result.Meters() != 13 {
			t.Errorf("Expected 13 meters, got %f", result.Meters())
		}
	})

	t.Run("TestSubtract", func(t *testing.T) {
		result := length1.Subtract(length2)
		if result.Meters() != 7 {
			t.Errorf("Expected 7 meters, got %f", result.Meters())
		}
	})

	t.Run("TestMultiply", func(t *testing.T) {
		result := length1.Multiply(length2)
		if result.Meters() != 30 {
			t.Errorf("Expected 30 meters, got %f", result.Meters())
		}
	})

	t.Run("TestDivide", func(t *testing.T) {
		result := length1.Divide(length2)
		expected := 3.33333333333333
		if diff := result.Meters() - expected; diff < -0.000000001 || diff > 0.000000001 {
			t.Errorf("Expected %f meters, got %f", expected, result.Meters())
		}
	})
}
