package unitsnet_go_test

import (
	"fmt"
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

func nextPositionFormula(position *units.Length, speed *units.Speed, timePassed *units.Duration) *units.Length {
	lf := units.LengthFactory{}
	next, _ := lf.FromMeters(position.Meters() + (speed.MetersPerSecond() * timePassed.Seconds()))
	return next
}

func getCurrentPosition(position *units.Length, speed *units.Speed, timePassed *units.Duration) *units.Length {
	return nextPositionFormula(position, speed, timePassed)
}

func main() {
	lf := units.LengthFactory{}
	sf := units.SpeedFactory{}
	df := units.DurationFactory{}
	platformPosition, _ := lf.FromMeters(100)
	platformSpeed, _ := sf.FromMetersPerMinutes(1)
	timePassed, _ := df.FromMicroseconds(500)
	newPlatform := getCurrentPosition(platformPosition, platformSpeed, timePassed)

	fmt.Println(newPlatform.Meters())
}
