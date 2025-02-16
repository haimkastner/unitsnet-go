// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPermittivityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "FaradPerMeter"}`
	
	factory := units.PermittivityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PermittivityFaradPerMeter {
		t.Errorf("expected unit %v, got %v", units.PermittivityFaradPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "FaradPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPermittivityDto_ToJSON(t *testing.T) {
	dto := units.PermittivityDto{
		Value: 45,
		Unit:  units.PermittivityFaradPerMeter,
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
	if result["unit"].(string) != string(units.PermittivityFaradPerMeter) {
		t.Errorf("expected unit %s, got %v", units.PermittivityFaradPerMeter, result["unit"])
	}
}

func TestNewPermittivity_InvalidValue(t *testing.T) {
	factory := units.PermittivityFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePermittivity(math.NaN(), units.PermittivityFaradPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePermittivity(math.Inf(1), units.PermittivityFaradPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPermittivityConversions(t *testing.T) {
	factory := units.PermittivityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePermittivity(180, units.PermittivityFaradPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to FaradsPerMeter.
		// No expected conversion value provided for FaradsPerMeter, verifying result is not NaN.
		result := a.FaradsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to FaradsPerMeter returned NaN")
		}
	}
}

func TestPermittivity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PermittivityFactory{}
	a, err := factory.CreatePermittivity(90, units.PermittivityFaradPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PermittivityFaradPerMeter {
		t.Errorf("expected default unit FaradPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PermittivityFaradPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PermittivityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PermittivityFaradPerMeter {
		t.Errorf("expected unit FaradPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPermittivityToString(t *testing.T) {
	factory := units.PermittivityFactory{}
	a, err := factory.CreatePermittivity(45, units.PermittivityFaradPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PermittivityFaradPerMeter, 2)
	expected := "45.00 " + units.GetPermittivityAbbreviation(units.PermittivityFaradPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PermittivityFaradPerMeter, -1)
	expected = "45 " + units.GetPermittivityAbbreviation(units.PermittivityFaradPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPermittivity_EqualityAndComparison(t *testing.T) {
	factory := units.PermittivityFactory{}
	a1, _ := factory.CreatePermittivity(60, units.PermittivityFaradPerMeter)
	a2, _ := factory.CreatePermittivity(60, units.PermittivityFaradPerMeter)
	a3, _ := factory.CreatePermittivity(120, units.PermittivityFaradPerMeter)

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

func TestPermittivity_Arithmetic(t *testing.T) {
	factory := units.PermittivityFactory{}
	a1, _ := factory.CreatePermittivity(30, units.PermittivityFaradPerMeter)
	a2, _ := factory.CreatePermittivity(45, units.PermittivityFaradPerMeter)

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