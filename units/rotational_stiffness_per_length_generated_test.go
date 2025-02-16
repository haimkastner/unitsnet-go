// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRotationalStiffnessPerLengthDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonMeterPerRadianPerMeter"}`
	
	factory := units.RotationalStiffnessPerLengthDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter {
		t.Errorf("expected unit %v, got %v", units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonMeterPerRadianPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRotationalStiffnessPerLengthDto_ToJSON(t *testing.T) {
	dto := units.RotationalStiffnessPerLengthDto{
		Value: 45,
		Unit:  units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter,
	}
	data, err := dto.ToJSON()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("error unmarshalling JSON: %v", err)
	}
	if result["value"].(float64) != 45 {
		t.Errorf("expected value 45, got %v", result["value"])
	}
	if result["unit"].(string) != string(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter) {
		t.Errorf("expected unit %s, got %v", units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter, result["unit"])
	}
}

func TestNewRotationalStiffnessPerLength_InvalidValue(t *testing.T) {
	factory := units.RotationalStiffnessPerLengthFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRotationalStiffnessPerLength(math.NaN(), units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRotationalStiffnessPerLength(math.Inf(1), units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRotationalStiffnessPerLengthConversions(t *testing.T) {
	factory := units.RotationalStiffnessPerLengthFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRotationalStiffnessPerLength(180, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonMetersPerRadianPerMeter.
		// No expected conversion value provided for NewtonMetersPerRadianPerMeter, verifying result is not NaN.
		result := a.NewtonMetersPerRadianPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonMetersPerRadianPerMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundForceFeetPerDegreesPerFeet.
		// No expected conversion value provided for PoundForceFeetPerDegreesPerFeet, verifying result is not NaN.
		result := a.PoundForceFeetPerDegreesPerFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundForceFeetPerDegreesPerFeet returned NaN")
		}
	}
	{
		// Test conversion to KilopoundForceFeetPerDegreesPerFeet.
		// No expected conversion value provided for KilopoundForceFeetPerDegreesPerFeet, verifying result is not NaN.
		result := a.KilopoundForceFeetPerDegreesPerFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundForceFeetPerDegreesPerFeet returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMetersPerRadianPerMeter.
		// No expected conversion value provided for KilonewtonMetersPerRadianPerMeter, verifying result is not NaN.
		result := a.KilonewtonMetersPerRadianPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonMetersPerRadianPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMetersPerRadianPerMeter.
		// No expected conversion value provided for MeganewtonMetersPerRadianPerMeter, verifying result is not NaN.
		result := a.MeganewtonMetersPerRadianPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonMetersPerRadianPerMeter returned NaN")
		}
	}
}

func TestRotationalStiffnessPerLength_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RotationalStiffnessPerLengthFactory{}
	a, err := factory.CreateRotationalStiffnessPerLength(90, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter {
		t.Errorf("expected default unit NewtonMeterPerRadianPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RotationalStiffnessPerLengthDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter {
		t.Errorf("expected unit NewtonMeterPerRadianPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRotationalStiffnessPerLengthToString(t *testing.T) {
	factory := units.RotationalStiffnessPerLengthFactory{}
	a, err := factory.CreateRotationalStiffnessPerLength(45, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter, 2)
	expected := "45.00 " + units.GetRotationalStiffnessPerLengthAbbreviation(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter, -1)
	expected = "45 " + units.GetRotationalStiffnessPerLengthAbbreviation(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRotationalStiffnessPerLength_EqualityAndComparison(t *testing.T) {
	factory := units.RotationalStiffnessPerLengthFactory{}
	a1, _ := factory.CreateRotationalStiffnessPerLength(60, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	a2, _ := factory.CreateRotationalStiffnessPerLength(60, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	a3, _ := factory.CreateRotationalStiffnessPerLength(120, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)

	if !a1.Equals(a2) {
		t.Error("expected a1 and a2 to be equal")
	}
	if a1.Equals(a3) {
		t.Error("expected a1 and a3 not to be equal")
	}
	if a1.CompareTo(a2) != 0 {
		t.Error("expected CompareTo to return 0 for equal values")
	}
	if a1.CompareTo(a3) != -1 {
		t.Errorf("expected a1 less than a3, got %d", a1.CompareTo(a3))
	}
	if a3.CompareTo(a1) != 1 {
		t.Errorf("expected a3 greater than a1, got %d", a3.CompareTo(a1))
	}
}

func TestRotationalStiffnessPerLength_Arithmetic(t *testing.T) {
	factory := units.RotationalStiffnessPerLengthFactory{}
	a1, _ := factory.CreateRotationalStiffnessPerLength(30, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	a2, _ := factory.CreateRotationalStiffnessPerLength(45, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)

	added := a1.Add(a2)
	if math.Abs(added.BaseValue()-75) > 1e-9 {
		t.Errorf("expected sum 75, got %v", added.BaseValue())
	}

	subtracted := a2.Subtract(a1)
	if math.Abs(subtracted.BaseValue()-15) > 1e-9 {
		t.Errorf("expected difference 15, got %v", subtracted.BaseValue())
	}

	multiplied := a1.Multiply(a2)
	if math.Abs(multiplied.BaseValue()-1350) > 1e-9 {
		t.Errorf("expected product 1350, got %v", multiplied.BaseValue())
	}

	divided := a2.Divide(a1)
	if math.Abs(divided.BaseValue()-1.5) > 1e-9 {
		t.Errorf("expected quotient 1.5, got %v", divided.BaseValue())
	}
}