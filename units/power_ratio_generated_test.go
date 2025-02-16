// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPowerRatioDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "DecibelWatt"}`
	
	factory := units.PowerRatioDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PowerRatioDecibelWatt {
		t.Errorf("expected unit %v, got %v", units.PowerRatioDecibelWatt, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "DecibelWatt"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPowerRatioDto_ToJSON(t *testing.T) {
	dto := units.PowerRatioDto{
		Value: 45,
		Unit:  units.PowerRatioDecibelWatt,
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
	if result["unit"].(string) != string(units.PowerRatioDecibelWatt) {
		t.Errorf("expected unit %s, got %v", units.PowerRatioDecibelWatt, result["unit"])
	}
}

func TestNewPowerRatio_InvalidValue(t *testing.T) {
	factory := units.PowerRatioFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePowerRatio(math.NaN(), units.PowerRatioDecibelWatt)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePowerRatio(math.Inf(1), units.PowerRatioDecibelWatt)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPowerRatioConversions(t *testing.T) {
	factory := units.PowerRatioFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePowerRatio(180, units.PowerRatioDecibelWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to DecibelWatts.
		// No expected conversion value provided for DecibelWatts, verifying result is not NaN.
		result := a.DecibelWatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecibelWatts returned NaN")
		}
	}
	{
		// Test conversion to DecibelMilliwatts.
		// No expected conversion value provided for DecibelMilliwatts, verifying result is not NaN.
		result := a.DecibelMilliwatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecibelMilliwatts returned NaN")
		}
	}
}

func TestPowerRatio_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PowerRatioFactory{}
	a, err := factory.CreatePowerRatio(90, units.PowerRatioDecibelWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PowerRatioDecibelWatt {
		t.Errorf("expected default unit DecibelWatt, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PowerRatioDecibelWatt
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PowerRatioDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PowerRatioDecibelWatt {
		t.Errorf("expected unit DecibelWatt, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPowerRatioToString(t *testing.T) {
	factory := units.PowerRatioFactory{}
	a, err := factory.CreatePowerRatio(45, units.PowerRatioDecibelWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PowerRatioDecibelWatt, 2)
	expected := "45.00 " + units.GetPowerRatioAbbreviation(units.PowerRatioDecibelWatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PowerRatioDecibelWatt, -1)
	expected = "45 " + units.GetPowerRatioAbbreviation(units.PowerRatioDecibelWatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPowerRatio_EqualityAndComparison(t *testing.T) {
	factory := units.PowerRatioFactory{}
	a1, _ := factory.CreatePowerRatio(60, units.PowerRatioDecibelWatt)
	a2, _ := factory.CreatePowerRatio(60, units.PowerRatioDecibelWatt)
	a3, _ := factory.CreatePowerRatio(120, units.PowerRatioDecibelWatt)

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

func TestPowerRatio_Arithmetic(t *testing.T) {
	factory := units.PowerRatioFactory{}
	a1, _ := factory.CreatePowerRatio(30, units.PowerRatioDecibelWatt)
	a2, _ := factory.CreatePowerRatio(45, units.PowerRatioDecibelWatt)

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