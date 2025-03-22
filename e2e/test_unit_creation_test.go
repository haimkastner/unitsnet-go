package unitsnet_go_test

import (
	"strings"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"
)

func TestUnitCreation(t *testing.T) {
	// Create a factory instance
	af := units.AngleFactory{}

	angleCreateDriven, _ := af.CreateAngle(180, units.AngleDegree)
	angleFromDriven, _ := af.FromDegrees(180)

	if angleCreateDriven.BaseValue() != angleFromDriven.BaseValue() {
		t.Errorf("Expected angleCreateDriven and angleFromDriven to be equal")
	}
}

func TestBrokenUnitCreation(t *testing.T) {
	// Create a factory instance
	af := units.AngleFactory{}

	_, err := af.CreateAngle(180, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}
