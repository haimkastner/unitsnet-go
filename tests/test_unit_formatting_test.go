package unitsnet_go_test

import (
	"fmt"
	"testing"

	"github.com/haimkastner/unitsnet-go/unitsnet_go"
)

func TestFmtFormatDefault(t *testing.T) {
	angle, _ := unitsnet_go.AngleFactory{}.FromDegrees(180)
	expected := "180.00 °"
	if fmt.Sprintf("%v", angle) != expected {
		t.Errorf("Expected '%s', got '%s'", expected, fmt.Sprintf("%v", angle))
	}
}

func TestFormatBase(t *testing.T) {
	angle, _ := unitsnet_go.AngleFactory{}.FromDegrees(180)
	if angle.ToString(unitsnet_go.AngleDegree, 0) != "180 °" {
		t.Errorf("Expected '180 °', got '%s'", angle.ToString(unitsnet_go.AngleDegree, 0))
	}
}

func TestFormatOtherUnit(t *testing.T) {
	angle, _ := unitsnet_go.AngleFactory{}.FromDegrees(180)

	if angle.ToString(unitsnet_go.AngleDegree, 0) != "180 °" {
		t.Errorf("Expected '180 °', got '%s'", angle.ToString(unitsnet_go.AngleDegree, 0))
	}

	if angle.ToString(unitsnet_go.AngleRadian, -1) != "3.141592653589793 rad" {
		t.Errorf("Expected '3.141592653589793 rad', got '%s'", angle.ToString(unitsnet_go.AngleRadian, -1))
	}

	if angle.ToString(unitsnet_go.AngleMilliradian, -1) != "3141.592653589793 mrad" {
		t.Errorf("Expected '3141.592653589793 mrad', got '%s'", angle.ToString(unitsnet_go.AngleRadian, -1))
	}
}

func TestFormatWithDigitsLimit(t *testing.T) {
	angle, _ := unitsnet_go.AngleFactory{}.FromDegrees(180)

	// Without limit
	if result := angle.ToString(unitsnet_go.AngleRadian, -1); result != "3.141592653589793 rad" {
		t.Errorf("Expected '3.141592653589793 rad', got '%s'", result)
	}

	// Without limit - random negative number
	if result := angle.ToString(unitsnet_go.AngleRadian, -10); result != "3.141592653589793 rad" {
		t.Errorf("Expected '3.141592653589793 rad', got '%s'", result)
	}

	// Limit to 0 digits
	if result := angle.ToString(unitsnet_go.AngleRadian, 0); result != "3 rad" {
		t.Errorf("Expected '3 rad', got '%s'", result)
	}

	// Limit to 1 digit
	if result := angle.ToString(unitsnet_go.AngleRadian, 1); result != "3.1 rad" {
		t.Errorf("Expected '3.1 rad', got '%s'", result)
	}

	// Limit to 2 digits
	if result := angle.ToString(unitsnet_go.AngleRadian, 2); result != "3.14 rad" {
		t.Errorf("Expected '3.14 rad', got '%s'", result)
	}

	// No change if it's already formatted correctly
	if result := angle.ToString(unitsnet_go.AngleDegree, 2); result != "180.00 °" {
		t.Errorf("Expected '180.00 °', got '%s'", result)
	}
}
