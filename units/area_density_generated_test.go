// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAreaDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerSquareMeter"}`
	
	factory := units.AreaDensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AreaDensityKilogramPerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.AreaDensityKilogramPerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAreaDensityDto_ToJSON(t *testing.T) {
	dto := units.AreaDensityDto{
		Value: 45,
		Unit:  units.AreaDensityKilogramPerSquareMeter,
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
	if result["unit"].(string) != string(units.AreaDensityKilogramPerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.AreaDensityKilogramPerSquareMeter, result["unit"])
	}
}

func TestNewAreaDensity_InvalidValue(t *testing.T) {
	factory := units.AreaDensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAreaDensity(math.NaN(), units.AreaDensityKilogramPerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAreaDensity(math.Inf(1), units.AreaDensityKilogramPerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAreaDensityConversions(t *testing.T) {
	factory := units.AreaDensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAreaDensity(180, units.AreaDensityKilogramPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to KilogramsPerSquareMeter.
		// No expected conversion value provided for KilogramsPerSquareMeter, verifying result is not NaN.
		result := a.KilogramsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerSquareMeter.
		// No expected conversion value provided for GramsPerSquareMeter, verifying result is not NaN.
		result := a.GramsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerSquareMeter.
		// No expected conversion value provided for MilligramsPerSquareMeter, verifying result is not NaN.
		result := a.MilligramsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerSquareMeter returned NaN")
		}
	}
}

func TestAreaDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AreaDensityFactory{}
	a, err := factory.CreateAreaDensity(90, units.AreaDensityKilogramPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AreaDensityKilogramPerSquareMeter {
		t.Errorf("expected default unit KilogramPerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AreaDensityKilogramPerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AreaDensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AreaDensityKilogramPerSquareMeter {
		t.Errorf("expected unit KilogramPerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAreaDensityToString(t *testing.T) {
	factory := units.AreaDensityFactory{}
	a, err := factory.CreateAreaDensity(45, units.AreaDensityKilogramPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AreaDensityKilogramPerSquareMeter, 2)
	expected := "45.00 " + units.GetAreaDensityAbbreviation(units.AreaDensityKilogramPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AreaDensityKilogramPerSquareMeter, -1)
	expected = "45 " + units.GetAreaDensityAbbreviation(units.AreaDensityKilogramPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAreaDensity_EqualityAndComparison(t *testing.T) {
	factory := units.AreaDensityFactory{}
	a1, _ := factory.CreateAreaDensity(60, units.AreaDensityKilogramPerSquareMeter)
	a2, _ := factory.CreateAreaDensity(60, units.AreaDensityKilogramPerSquareMeter)
	a3, _ := factory.CreateAreaDensity(120, units.AreaDensityKilogramPerSquareMeter)

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

func TestAreaDensity_Arithmetic(t *testing.T) {
	factory := units.AreaDensityFactory{}
	a1, _ := factory.CreateAreaDensity(30, units.AreaDensityKilogramPerSquareMeter)
	a2, _ := factory.CreateAreaDensity(45, units.AreaDensityKilogramPerSquareMeter)

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