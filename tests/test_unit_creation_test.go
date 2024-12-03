package unitsnet_go_test

import (
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
