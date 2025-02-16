// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestCoefficientOfThermalExpansionDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "PerKelvin"}`
	
	factory := units.CoefficientOfThermalExpansionDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.CoefficientOfThermalExpansionPerKelvin {
		t.Errorf("expected unit %v, got %v", units.CoefficientOfThermalExpansionPerKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "PerKelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestCoefficientOfThermalExpansionDto_ToJSON(t *testing.T) {
	dto := units.CoefficientOfThermalExpansionDto{
		Value: 45,
		Unit:  units.CoefficientOfThermalExpansionPerKelvin,
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
	if result["unit"].(string) != string(units.CoefficientOfThermalExpansionPerKelvin) {
		t.Errorf("expected unit %s, got %v", units.CoefficientOfThermalExpansionPerKelvin, result["unit"])
	}
}

func TestNewCoefficientOfThermalExpansion_InvalidValue(t *testing.T) {
	factory := units.CoefficientOfThermalExpansionFactory{}
	// NaN value should return an error.
	_, err := factory.CreateCoefficientOfThermalExpansion(math.NaN(), units.CoefficientOfThermalExpansionPerKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateCoefficientOfThermalExpansion(math.Inf(1), units.CoefficientOfThermalExpansionPerKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestCoefficientOfThermalExpansionConversions(t *testing.T) {
	factory := units.CoefficientOfThermalExpansionFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateCoefficientOfThermalExpansion(180, units.CoefficientOfThermalExpansionPerKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to PerKelvin.
		// No expected conversion value provided for PerKelvin, verifying result is not NaN.
		result := a.PerKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to PerKelvin returned NaN")
		}
	}
	{
		// Test conversion to PerDegreeCelsius.
		// No expected conversion value provided for PerDegreeCelsius, verifying result is not NaN.
		result := a.PerDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to PerDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to PerDegreeFahrenheit.
		// No expected conversion value provided for PerDegreeFahrenheit, verifying result is not NaN.
		result := a.PerDegreeFahrenheit()
		if math.IsNaN(result) {
			t.Errorf("conversion to PerDegreeFahrenheit returned NaN")
		}
	}
	{
		// Test conversion to PpmPerKelvin.
		// No expected conversion value provided for PpmPerKelvin, verifying result is not NaN.
		result := a.PpmPerKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to PpmPerKelvin returned NaN")
		}
	}
	{
		// Test conversion to PpmPerDegreeCelsius.
		// No expected conversion value provided for PpmPerDegreeCelsius, verifying result is not NaN.
		result := a.PpmPerDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to PpmPerDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to PpmPerDegreeFahrenheit.
		// No expected conversion value provided for PpmPerDegreeFahrenheit, verifying result is not NaN.
		result := a.PpmPerDegreeFahrenheit()
		if math.IsNaN(result) {
			t.Errorf("conversion to PpmPerDegreeFahrenheit returned NaN")
		}
	}
}

func TestCoefficientOfThermalExpansion_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.CoefficientOfThermalExpansionFactory{}
	a, err := factory.CreateCoefficientOfThermalExpansion(90, units.CoefficientOfThermalExpansionPerKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.CoefficientOfThermalExpansionPerKelvin {
		t.Errorf("expected default unit PerKelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.CoefficientOfThermalExpansionPerKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.CoefficientOfThermalExpansionDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.CoefficientOfThermalExpansionPerKelvin {
		t.Errorf("expected unit PerKelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestCoefficientOfThermalExpansionToString(t *testing.T) {
	factory := units.CoefficientOfThermalExpansionFactory{}
	a, err := factory.CreateCoefficientOfThermalExpansion(45, units.CoefficientOfThermalExpansionPerKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.CoefficientOfThermalExpansionPerKelvin, 2)
	expected := "45.00 " + units.GetCoefficientOfThermalExpansionAbbreviation(units.CoefficientOfThermalExpansionPerKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.CoefficientOfThermalExpansionPerKelvin, -1)
	expected = "45 " + units.GetCoefficientOfThermalExpansionAbbreviation(units.CoefficientOfThermalExpansionPerKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestCoefficientOfThermalExpansion_EqualityAndComparison(t *testing.T) {
	factory := units.CoefficientOfThermalExpansionFactory{}
	a1, _ := factory.CreateCoefficientOfThermalExpansion(60, units.CoefficientOfThermalExpansionPerKelvin)
	a2, _ := factory.CreateCoefficientOfThermalExpansion(60, units.CoefficientOfThermalExpansionPerKelvin)
	a3, _ := factory.CreateCoefficientOfThermalExpansion(120, units.CoefficientOfThermalExpansionPerKelvin)

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

func TestCoefficientOfThermalExpansion_Arithmetic(t *testing.T) {
	factory := units.CoefficientOfThermalExpansionFactory{}
	a1, _ := factory.CreateCoefficientOfThermalExpansion(30, units.CoefficientOfThermalExpansionPerKelvin)
	a2, _ := factory.CreateCoefficientOfThermalExpansion(45, units.CoefficientOfThermalExpansionPerKelvin)

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