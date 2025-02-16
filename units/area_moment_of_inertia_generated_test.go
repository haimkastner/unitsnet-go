// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAreaMomentOfInertiaDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MeterToTheFourth"}`
	
	factory := units.AreaMomentOfInertiaDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AreaMomentOfInertiaMeterToTheFourth {
		t.Errorf("expected unit %v, got %v", units.AreaMomentOfInertiaMeterToTheFourth, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MeterToTheFourth"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAreaMomentOfInertiaDto_ToJSON(t *testing.T) {
	dto := units.AreaMomentOfInertiaDto{
		Value: 45,
		Unit:  units.AreaMomentOfInertiaMeterToTheFourth,
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
	if result["unit"].(string) != string(units.AreaMomentOfInertiaMeterToTheFourth) {
		t.Errorf("expected unit %s, got %v", units.AreaMomentOfInertiaMeterToTheFourth, result["unit"])
	}
}

func TestNewAreaMomentOfInertia_InvalidValue(t *testing.T) {
	factory := units.AreaMomentOfInertiaFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAreaMomentOfInertia(math.NaN(), units.AreaMomentOfInertiaMeterToTheFourth)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAreaMomentOfInertia(math.Inf(1), units.AreaMomentOfInertiaMeterToTheFourth)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAreaMomentOfInertiaConversions(t *testing.T) {
	factory := units.AreaMomentOfInertiaFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAreaMomentOfInertia(180, units.AreaMomentOfInertiaMeterToTheFourth)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MetersToTheFourth.
		// No expected conversion value provided for MetersToTheFourth, verifying result is not NaN.
		result := a.MetersToTheFourth()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetersToTheFourth returned NaN")
		}
	}
	{
		// Test conversion to DecimetersToTheFourth.
		// No expected conversion value provided for DecimetersToTheFourth, verifying result is not NaN.
		result := a.DecimetersToTheFourth()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecimetersToTheFourth returned NaN")
		}
	}
	{
		// Test conversion to CentimetersToTheFourth.
		// No expected conversion value provided for CentimetersToTheFourth, verifying result is not NaN.
		result := a.CentimetersToTheFourth()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentimetersToTheFourth returned NaN")
		}
	}
	{
		// Test conversion to MillimetersToTheFourth.
		// No expected conversion value provided for MillimetersToTheFourth, verifying result is not NaN.
		result := a.MillimetersToTheFourth()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillimetersToTheFourth returned NaN")
		}
	}
	{
		// Test conversion to FeetToTheFourth.
		// No expected conversion value provided for FeetToTheFourth, verifying result is not NaN.
		result := a.FeetToTheFourth()
		if math.IsNaN(result) {
			t.Errorf("conversion to FeetToTheFourth returned NaN")
		}
	}
	{
		// Test conversion to InchesToTheFourth.
		// No expected conversion value provided for InchesToTheFourth, verifying result is not NaN.
		result := a.InchesToTheFourth()
		if math.IsNaN(result) {
			t.Errorf("conversion to InchesToTheFourth returned NaN")
		}
	}
}

func TestAreaMomentOfInertia_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AreaMomentOfInertiaFactory{}
	a, err := factory.CreateAreaMomentOfInertia(90, units.AreaMomentOfInertiaMeterToTheFourth)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AreaMomentOfInertiaMeterToTheFourth {
		t.Errorf("expected default unit MeterToTheFourth, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AreaMomentOfInertiaMeterToTheFourth
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AreaMomentOfInertiaDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AreaMomentOfInertiaMeterToTheFourth {
		t.Errorf("expected unit MeterToTheFourth, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAreaMomentOfInertiaToString(t *testing.T) {
	factory := units.AreaMomentOfInertiaFactory{}
	a, err := factory.CreateAreaMomentOfInertia(45, units.AreaMomentOfInertiaMeterToTheFourth)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AreaMomentOfInertiaMeterToTheFourth, 2)
	expected := "45.00 " + units.GetAreaMomentOfInertiaAbbreviation(units.AreaMomentOfInertiaMeterToTheFourth)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AreaMomentOfInertiaMeterToTheFourth, -1)
	expected = "45 " + units.GetAreaMomentOfInertiaAbbreviation(units.AreaMomentOfInertiaMeterToTheFourth)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAreaMomentOfInertia_EqualityAndComparison(t *testing.T) {
	factory := units.AreaMomentOfInertiaFactory{}
	a1, _ := factory.CreateAreaMomentOfInertia(60, units.AreaMomentOfInertiaMeterToTheFourth)
	a2, _ := factory.CreateAreaMomentOfInertia(60, units.AreaMomentOfInertiaMeterToTheFourth)
	a3, _ := factory.CreateAreaMomentOfInertia(120, units.AreaMomentOfInertiaMeterToTheFourth)

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

func TestAreaMomentOfInertia_Arithmetic(t *testing.T) {
	factory := units.AreaMomentOfInertiaFactory{}
	a1, _ := factory.CreateAreaMomentOfInertia(30, units.AreaMomentOfInertiaMeterToTheFourth)
	a2, _ := factory.CreateAreaMomentOfInertia(45, units.AreaMomentOfInertiaMeterToTheFourth)

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