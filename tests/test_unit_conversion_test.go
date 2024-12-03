package unitsnet_go_test

import (
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/unitsnet_go"
)

var af = unitsnet_go.AngleFactory{}
var inf = unitsnet_go.InformationFactory{}
var maxDelta = 0.0000001

func TestConvertFromBaseToOtherUnit(t *testing.T) {
	angle, _ := af.FromDegrees(180)
	expected := float64(3.141592653589793)
	if angle.Radians() != expected {
		t.Errorf("Expected %f, but got %f", expected, angle.Radians())
	}
}

func TestConvertBaseToPrefixUnit(t *testing.T) {
	angle, _ := af.FromDegrees(180)
	expected := float64(180000000)
	if angle.Microdegrees() != expected {
		t.Errorf("Expected %f, but got %f", expected, angle.Microdegrees())
	}
}

func TestConvertBaseToOtherUnitPrefixUnit(t *testing.T) {
	angle, _ := af.FromDegrees(180)
	expected := float64(3141592.6535897935)
	if angle.Microradians() != expected {
		t.Errorf("Expected %f, but got %f", expected, angle.Microradians())
	}
}

func TestConvertUnitToBase(t *testing.T) {
	angle, _ := af.FromRadians(3.141592653589793)
	expected := float64(180)
	if angle.Degrees() != expected {
		t.Errorf("Expected %f, but got %f", expected, angle.Degrees())
	}
}

func TestConvertBasePrefixToBase(t *testing.T) {
	angle, _ := af.FromMicrodegrees(180000000)
	expected := float64(180)
	if angle.Degrees() != expected {
		t.Errorf("Expected %f, but got %f", expected, angle.Degrees())
	}
}

func TestConvertUnitPrefixToBase(t *testing.T) {
	angle, _ := af.FromMicroradians(3141592.65358979)
	expected := float64(180)
	if math.Abs(angle.Degrees()-expected) > maxDelta {
		t.Errorf("Expected %f, but got %f, difference is larger than %f", expected, angle.Degrees(), maxDelta)
	}
}

func TestConvertBitsPrefixToBase(t *testing.T) {
	data, _ := inf.FromKibibits(1)
	expected := float64(1024)
	if data.Bits() != expected {
		t.Errorf("Expected %f, but got %f", expected, data.Bits())
	}
}

func TestConvertBaseToBitsPrefix(t *testing.T) {
	data, _ := inf.FromBits(1024)
	expected := float64(1)
	if data.Kibibits() != expected {
		t.Errorf("Expected %f, but got %f", expected, data.Kibibits())
	}
}
