// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestWarpingMomentOfInertiaDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MeterToTheSixth"}`
	
	factory := units.WarpingMomentOfInertiaDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.WarpingMomentOfInertiaMeterToTheSixth {
		t.Errorf("expected unit %v, got %v", units.WarpingMomentOfInertiaMeterToTheSixth, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MeterToTheSixth"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestWarpingMomentOfInertiaDto_ToJSON(t *testing.T) {
	dto := units.WarpingMomentOfInertiaDto{
		Value: 45,
		Unit:  units.WarpingMomentOfInertiaMeterToTheSixth,
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
	if result["unit"].(string) != string(units.WarpingMomentOfInertiaMeterToTheSixth) {
		t.Errorf("expected unit %s, got %v", units.WarpingMomentOfInertiaMeterToTheSixth, result["unit"])
	}
}

func TestNewWarpingMomentOfInertia_InvalidValue(t *testing.T) {
	factory := units.WarpingMomentOfInertiaFactory{}
	// NaN value should return an error.
	_, err := factory.CreateWarpingMomentOfInertia(math.NaN(), units.WarpingMomentOfInertiaMeterToTheSixth)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateWarpingMomentOfInertia(math.Inf(1), units.WarpingMomentOfInertiaMeterToTheSixth)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestWarpingMomentOfInertiaConversions(t *testing.T) {
	factory := units.WarpingMomentOfInertiaFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateWarpingMomentOfInertia(180, units.WarpingMomentOfInertiaMeterToTheSixth)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MetersToTheSixth.
		// No expected conversion value provided for MetersToTheSixth, verifying result is not NaN.
		result := a.MetersToTheSixth()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetersToTheSixth returned NaN")
		}
	}
	{
		// Test conversion to DecimetersToTheSixth.
		// No expected conversion value provided for DecimetersToTheSixth, verifying result is not NaN.
		result := a.DecimetersToTheSixth()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecimetersToTheSixth returned NaN")
		}
	}
	{
		// Test conversion to CentimetersToTheSixth.
		// No expected conversion value provided for CentimetersToTheSixth, verifying result is not NaN.
		result := a.CentimetersToTheSixth()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentimetersToTheSixth returned NaN")
		}
	}
	{
		// Test conversion to MillimetersToTheSixth.
		// No expected conversion value provided for MillimetersToTheSixth, verifying result is not NaN.
		result := a.MillimetersToTheSixth()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillimetersToTheSixth returned NaN")
		}
	}
	{
		// Test conversion to FeetToTheSixth.
		// No expected conversion value provided for FeetToTheSixth, verifying result is not NaN.
		result := a.FeetToTheSixth()
		if math.IsNaN(result) {
			t.Errorf("conversion to FeetToTheSixth returned NaN")
		}
	}
	{
		// Test conversion to InchesToTheSixth.
		// No expected conversion value provided for InchesToTheSixth, verifying result is not NaN.
		result := a.InchesToTheSixth()
		if math.IsNaN(result) {
			t.Errorf("conversion to InchesToTheSixth returned NaN")
		}
	}
}

func TestWarpingMomentOfInertia_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.WarpingMomentOfInertiaFactory{}
	a, err := factory.CreateWarpingMomentOfInertia(90, units.WarpingMomentOfInertiaMeterToTheSixth)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.WarpingMomentOfInertiaMeterToTheSixth {
		t.Errorf("expected default unit MeterToTheSixth, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.WarpingMomentOfInertiaMeterToTheSixth
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.WarpingMomentOfInertiaDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.WarpingMomentOfInertiaMeterToTheSixth {
		t.Errorf("expected unit MeterToTheSixth, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestWarpingMomentOfInertiaToString(t *testing.T) {
	factory := units.WarpingMomentOfInertiaFactory{}
	a, err := factory.CreateWarpingMomentOfInertia(45, units.WarpingMomentOfInertiaMeterToTheSixth)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.WarpingMomentOfInertiaMeterToTheSixth, 2)
	expected := "45.00 " + units.GetWarpingMomentOfInertiaAbbreviation(units.WarpingMomentOfInertiaMeterToTheSixth)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.WarpingMomentOfInertiaMeterToTheSixth, -1)
	expected = "45 " + units.GetWarpingMomentOfInertiaAbbreviation(units.WarpingMomentOfInertiaMeterToTheSixth)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestWarpingMomentOfInertia_EqualityAndComparison(t *testing.T) {
	factory := units.WarpingMomentOfInertiaFactory{}
	a1, _ := factory.CreateWarpingMomentOfInertia(60, units.WarpingMomentOfInertiaMeterToTheSixth)
	a2, _ := factory.CreateWarpingMomentOfInertia(60, units.WarpingMomentOfInertiaMeterToTheSixth)
	a3, _ := factory.CreateWarpingMomentOfInertia(120, units.WarpingMomentOfInertiaMeterToTheSixth)

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

func TestWarpingMomentOfInertia_Arithmetic(t *testing.T) {
	factory := units.WarpingMomentOfInertiaFactory{}
	a1, _ := factory.CreateWarpingMomentOfInertia(30, units.WarpingMomentOfInertiaMeterToTheSixth)
	a2, _ := factory.CreateWarpingMomentOfInertia(45, units.WarpingMomentOfInertiaMeterToTheSixth)

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