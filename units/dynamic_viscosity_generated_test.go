// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestDynamicViscosityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonSecondPerMeterSquared"}`
	
	factory := units.DynamicViscosityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.DynamicViscosityNewtonSecondPerMeterSquared {
		t.Errorf("expected unit %v, got %v", units.DynamicViscosityNewtonSecondPerMeterSquared, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonSecondPerMeterSquared"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestDynamicViscosityDto_ToJSON(t *testing.T) {
	dto := units.DynamicViscosityDto{
		Value: 45,
		Unit:  units.DynamicViscosityNewtonSecondPerMeterSquared,
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
	if result["unit"].(string) != string(units.DynamicViscosityNewtonSecondPerMeterSquared) {
		t.Errorf("expected unit %s, got %v", units.DynamicViscosityNewtonSecondPerMeterSquared, result["unit"])
	}
}

func TestNewDynamicViscosity_InvalidValue(t *testing.T) {
	factory := units.DynamicViscosityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateDynamicViscosity(math.NaN(), units.DynamicViscosityNewtonSecondPerMeterSquared)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateDynamicViscosity(math.Inf(1), units.DynamicViscosityNewtonSecondPerMeterSquared)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestDynamicViscosityConversions(t *testing.T) {
	factory := units.DynamicViscosityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateDynamicViscosity(180, units.DynamicViscosityNewtonSecondPerMeterSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonSecondsPerMeterSquared.
		// No expected conversion value provided for NewtonSecondsPerMeterSquared, verifying result is not NaN.
		result := a.NewtonSecondsPerMeterSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonSecondsPerMeterSquared returned NaN")
		}
	}
	{
		// Test conversion to PascalSeconds.
		// No expected conversion value provided for PascalSeconds, verifying result is not NaN.
		result := a.PascalSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to PascalSeconds returned NaN")
		}
	}
	{
		// Test conversion to Poise.
		// No expected conversion value provided for Poise, verifying result is not NaN.
		result := a.Poise()
		if math.IsNaN(result) {
			t.Errorf("conversion to Poise returned NaN")
		}
	}
	{
		// Test conversion to Reyns.
		// No expected conversion value provided for Reyns, verifying result is not NaN.
		result := a.Reyns()
		if math.IsNaN(result) {
			t.Errorf("conversion to Reyns returned NaN")
		}
	}
	{
		// Test conversion to PoundsForceSecondPerSquareInch.
		// No expected conversion value provided for PoundsForceSecondPerSquareInch, verifying result is not NaN.
		result := a.PoundsForceSecondPerSquareInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForceSecondPerSquareInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsForceSecondPerSquareFoot.
		// No expected conversion value provided for PoundsForceSecondPerSquareFoot, verifying result is not NaN.
		result := a.PoundsForceSecondPerSquareFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForceSecondPerSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerFootSecond.
		// No expected conversion value provided for PoundsPerFootSecond, verifying result is not NaN.
		result := a.PoundsPerFootSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerFootSecond returned NaN")
		}
	}
	{
		// Test conversion to MillipascalSeconds.
		// No expected conversion value provided for MillipascalSeconds, verifying result is not NaN.
		result := a.MillipascalSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillipascalSeconds returned NaN")
		}
	}
	{
		// Test conversion to MicropascalSeconds.
		// No expected conversion value provided for MicropascalSeconds, verifying result is not NaN.
		result := a.MicropascalSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicropascalSeconds returned NaN")
		}
	}
	{
		// Test conversion to Centipoise.
		// No expected conversion value provided for Centipoise, verifying result is not NaN.
		result := a.Centipoise()
		if math.IsNaN(result) {
			t.Errorf("conversion to Centipoise returned NaN")
		}
	}
}

func TestDynamicViscosity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.DynamicViscosityFactory{}
	a, err := factory.CreateDynamicViscosity(90, units.DynamicViscosityNewtonSecondPerMeterSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.DynamicViscosityNewtonSecondPerMeterSquared {
		t.Errorf("expected default unit NewtonSecondPerMeterSquared, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.DynamicViscosityNewtonSecondPerMeterSquared
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.DynamicViscosityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.DynamicViscosityNewtonSecondPerMeterSquared {
		t.Errorf("expected unit NewtonSecondPerMeterSquared, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestDynamicViscosityToString(t *testing.T) {
	factory := units.DynamicViscosityFactory{}
	a, err := factory.CreateDynamicViscosity(45, units.DynamicViscosityNewtonSecondPerMeterSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.DynamicViscosityNewtonSecondPerMeterSquared, 2)
	expected := "45.00 " + units.GetDynamicViscosityAbbreviation(units.DynamicViscosityNewtonSecondPerMeterSquared)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.DynamicViscosityNewtonSecondPerMeterSquared, -1)
	expected = "45 " + units.GetDynamicViscosityAbbreviation(units.DynamicViscosityNewtonSecondPerMeterSquared)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestDynamicViscosity_EqualityAndComparison(t *testing.T) {
	factory := units.DynamicViscosityFactory{}
	a1, _ := factory.CreateDynamicViscosity(60, units.DynamicViscosityNewtonSecondPerMeterSquared)
	a2, _ := factory.CreateDynamicViscosity(60, units.DynamicViscosityNewtonSecondPerMeterSquared)
	a3, _ := factory.CreateDynamicViscosity(120, units.DynamicViscosityNewtonSecondPerMeterSquared)

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

func TestDynamicViscosity_Arithmetic(t *testing.T) {
	factory := units.DynamicViscosityFactory{}
	a1, _ := factory.CreateDynamicViscosity(30, units.DynamicViscosityNewtonSecondPerMeterSquared)
	a2, _ := factory.CreateDynamicViscosity(45, units.DynamicViscosityNewtonSecondPerMeterSquared)

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