// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestReactivePowerDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltampereReactive"}`
	
	factory := units.ReactivePowerDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ReactivePowerVoltampereReactive {
		t.Errorf("expected unit %v, got %v", units.ReactivePowerVoltampereReactive, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltampereReactive"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestReactivePowerDto_ToJSON(t *testing.T) {
	dto := units.ReactivePowerDto{
		Value: 45,
		Unit:  units.ReactivePowerVoltampereReactive,
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
	if result["unit"].(string) != string(units.ReactivePowerVoltampereReactive) {
		t.Errorf("expected unit %s, got %v", units.ReactivePowerVoltampereReactive, result["unit"])
	}
}

func TestNewReactivePower_InvalidValue(t *testing.T) {
	factory := units.ReactivePowerFactory{}
	// NaN value should return an error.
	_, err := factory.CreateReactivePower(math.NaN(), units.ReactivePowerVoltampereReactive)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateReactivePower(math.Inf(1), units.ReactivePowerVoltampereReactive)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestReactivePowerConversions(t *testing.T) {
	factory := units.ReactivePowerFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateReactivePower(180, units.ReactivePowerVoltampereReactive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltamperesReactive.
		// No expected conversion value provided for VoltamperesReactive, verifying result is not NaN.
		result := a.VoltamperesReactive()
		if math.IsNaN(result) {
			t.Errorf("conversion to VoltamperesReactive returned NaN")
		}
	}
	{
		// Test conversion to KilovoltamperesReactive.
		// No expected conversion value provided for KilovoltamperesReactive, verifying result is not NaN.
		result := a.KilovoltamperesReactive()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilovoltamperesReactive returned NaN")
		}
	}
	{
		// Test conversion to MegavoltamperesReactive.
		// No expected conversion value provided for MegavoltamperesReactive, verifying result is not NaN.
		result := a.MegavoltamperesReactive()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegavoltamperesReactive returned NaN")
		}
	}
	{
		// Test conversion to GigavoltamperesReactive.
		// No expected conversion value provided for GigavoltamperesReactive, verifying result is not NaN.
		result := a.GigavoltamperesReactive()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigavoltamperesReactive returned NaN")
		}
	}
}

func TestReactivePower_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ReactivePowerFactory{}
	a, err := factory.CreateReactivePower(90, units.ReactivePowerVoltampereReactive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ReactivePowerVoltampereReactive {
		t.Errorf("expected default unit VoltampereReactive, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ReactivePowerVoltampereReactive
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ReactivePowerDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ReactivePowerVoltampereReactive {
		t.Errorf("expected unit VoltampereReactive, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestReactivePowerToString(t *testing.T) {
	factory := units.ReactivePowerFactory{}
	a, err := factory.CreateReactivePower(45, units.ReactivePowerVoltampereReactive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ReactivePowerVoltampereReactive, 2)
	expected := "45.00 " + units.GetReactivePowerAbbreviation(units.ReactivePowerVoltampereReactive)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ReactivePowerVoltampereReactive, -1)
	expected = "45 " + units.GetReactivePowerAbbreviation(units.ReactivePowerVoltampereReactive)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestReactivePower_EqualityAndComparison(t *testing.T) {
	factory := units.ReactivePowerFactory{}
	a1, _ := factory.CreateReactivePower(60, units.ReactivePowerVoltampereReactive)
	a2, _ := factory.CreateReactivePower(60, units.ReactivePowerVoltampereReactive)
	a3, _ := factory.CreateReactivePower(120, units.ReactivePowerVoltampereReactive)

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

func TestReactivePower_Arithmetic(t *testing.T) {
	factory := units.ReactivePowerFactory{}
	a1, _ := factory.CreateReactivePower(30, units.ReactivePowerVoltampereReactive)
	a2, _ := factory.CreateReactivePower(45, units.ReactivePowerVoltampereReactive)

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