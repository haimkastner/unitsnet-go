// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricReactivePowerDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltampereReactive"}`
	
	factory := units.ElectricReactivePowerDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricReactivePowerVoltampereReactive {
		t.Errorf("expected unit %v, got %v", units.ElectricReactivePowerVoltampereReactive, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltampereReactive"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricReactivePowerDto_ToJSON(t *testing.T) {
	dto := units.ElectricReactivePowerDto{
		Value: 45,
		Unit:  units.ElectricReactivePowerVoltampereReactive,
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
	if result["unit"].(string) != string(units.ElectricReactivePowerVoltampereReactive) {
		t.Errorf("expected unit %s, got %v", units.ElectricReactivePowerVoltampereReactive, result["unit"])
	}
}

func TestNewElectricReactivePower_InvalidValue(t *testing.T) {
	factory := units.ElectricReactivePowerFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricReactivePower(math.NaN(), units.ElectricReactivePowerVoltampereReactive)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricReactivePower(math.Inf(1), units.ElectricReactivePowerVoltampereReactive)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricReactivePowerConversions(t *testing.T) {
	factory := units.ElectricReactivePowerFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricReactivePower(180, units.ElectricReactivePowerVoltampereReactive)
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

func TestElectricReactivePower_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricReactivePowerFactory{}
	a, err := factory.CreateElectricReactivePower(90, units.ElectricReactivePowerVoltampereReactive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricReactivePowerVoltampereReactive {
		t.Errorf("expected default unit VoltampereReactive, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricReactivePowerVoltampereReactive
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricReactivePowerDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricReactivePowerVoltampereReactive {
		t.Errorf("expected unit VoltampereReactive, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricReactivePowerToString(t *testing.T) {
	factory := units.ElectricReactivePowerFactory{}
	a, err := factory.CreateElectricReactivePower(45, units.ElectricReactivePowerVoltampereReactive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricReactivePowerVoltampereReactive, 2)
	expected := "45.00 " + units.GetElectricReactivePowerAbbreviation(units.ElectricReactivePowerVoltampereReactive)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricReactivePowerVoltampereReactive, -1)
	expected = "45 " + units.GetElectricReactivePowerAbbreviation(units.ElectricReactivePowerVoltampereReactive)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricReactivePower_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricReactivePowerFactory{}
	a1, _ := factory.CreateElectricReactivePower(60, units.ElectricReactivePowerVoltampereReactive)
	a2, _ := factory.CreateElectricReactivePower(60, units.ElectricReactivePowerVoltampereReactive)
	a3, _ := factory.CreateElectricReactivePower(120, units.ElectricReactivePowerVoltampereReactive)

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

func TestElectricReactivePower_Arithmetic(t *testing.T) {
	factory := units.ElectricReactivePowerFactory{}
	a1, _ := factory.CreateElectricReactivePower(30, units.ElectricReactivePowerVoltampereReactive)
	a2, _ := factory.CreateElectricReactivePower(45, units.ElectricReactivePowerVoltampereReactive)

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