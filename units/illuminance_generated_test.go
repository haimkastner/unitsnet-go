// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestIlluminanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Lux"}`
	
	factory := units.IlluminanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.IlluminanceLux {
		t.Errorf("expected unit %v, got %v", units.IlluminanceLux, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Lux"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestIlluminanceDto_ToJSON(t *testing.T) {
	dto := units.IlluminanceDto{
		Value: 45,
		Unit:  units.IlluminanceLux,
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
	if result["unit"].(string) != string(units.IlluminanceLux) {
		t.Errorf("expected unit %s, got %v", units.IlluminanceLux, result["unit"])
	}
}

func TestNewIlluminance_InvalidValue(t *testing.T) {
	factory := units.IlluminanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateIlluminance(math.NaN(), units.IlluminanceLux)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateIlluminance(math.Inf(1), units.IlluminanceLux)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestIlluminanceConversions(t *testing.T) {
	factory := units.IlluminanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateIlluminance(180, units.IlluminanceLux)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Lux.
		// No expected conversion value provided for Lux, verifying result is not NaN.
		result := a.Lux()
		if math.IsNaN(result) {
			t.Errorf("conversion to Lux returned NaN")
		}
	}
	{
		// Test conversion to Millilux.
		// No expected conversion value provided for Millilux, verifying result is not NaN.
		result := a.Millilux()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millilux returned NaN")
		}
	}
	{
		// Test conversion to Kilolux.
		// No expected conversion value provided for Kilolux, verifying result is not NaN.
		result := a.Kilolux()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilolux returned NaN")
		}
	}
	{
		// Test conversion to Megalux.
		// No expected conversion value provided for Megalux, verifying result is not NaN.
		result := a.Megalux()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megalux returned NaN")
		}
	}
}

func TestIlluminance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.IlluminanceFactory{}
	a, err := factory.CreateIlluminance(90, units.IlluminanceLux)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.IlluminanceLux {
		t.Errorf("expected default unit Lux, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.IlluminanceLux
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.IlluminanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.IlluminanceLux {
		t.Errorf("expected unit Lux, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestIlluminanceToString(t *testing.T) {
	factory := units.IlluminanceFactory{}
	a, err := factory.CreateIlluminance(45, units.IlluminanceLux)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.IlluminanceLux, 2)
	expected := "45.00 " + units.GetIlluminanceAbbreviation(units.IlluminanceLux)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.IlluminanceLux, -1)
	expected = "45 " + units.GetIlluminanceAbbreviation(units.IlluminanceLux)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestIlluminance_EqualityAndComparison(t *testing.T) {
	factory := units.IlluminanceFactory{}
	a1, _ := factory.CreateIlluminance(60, units.IlluminanceLux)
	a2, _ := factory.CreateIlluminance(60, units.IlluminanceLux)
	a3, _ := factory.CreateIlluminance(120, units.IlluminanceLux)

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

func TestIlluminance_Arithmetic(t *testing.T) {
	factory := units.IlluminanceFactory{}
	a1, _ := factory.CreateIlluminance(30, units.IlluminanceLux)
	a2, _ := factory.CreateIlluminance(45, units.IlluminanceLux)

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