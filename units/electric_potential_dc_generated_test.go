// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricPotentialDcDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltDc"}`
	
	factory := units.ElectricPotentialDcDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricPotentialDcVoltDc {
		t.Errorf("expected unit %v, got %v", units.ElectricPotentialDcVoltDc, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltDc"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricPotentialDcDto_ToJSON(t *testing.T) {
	dto := units.ElectricPotentialDcDto{
		Value: 45,
		Unit:  units.ElectricPotentialDcVoltDc,
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
	if result["unit"].(string) != string(units.ElectricPotentialDcVoltDc) {
		t.Errorf("expected unit %s, got %v", units.ElectricPotentialDcVoltDc, result["unit"])
	}
}

func TestNewElectricPotentialDc_InvalidValue(t *testing.T) {
	factory := units.ElectricPotentialDcFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricPotentialDc(math.NaN(), units.ElectricPotentialDcVoltDc)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricPotentialDc(math.Inf(1), units.ElectricPotentialDcVoltDc)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricPotentialDcConversions(t *testing.T) {
	factory := units.ElectricPotentialDcFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricPotentialDc(180, units.ElectricPotentialDcVoltDc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltsDc.
		// No expected conversion value provided for VoltsDc, verifying result is not NaN.
		result := a.VoltsDc()
		if math.IsNaN(result) {
			t.Errorf("conversion to VoltsDc returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsDc.
		// No expected conversion value provided for MicrovoltsDc, verifying result is not NaN.
		result := a.MicrovoltsDc()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrovoltsDc returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsDc.
		// No expected conversion value provided for MillivoltsDc, verifying result is not NaN.
		result := a.MillivoltsDc()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillivoltsDc returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsDc.
		// No expected conversion value provided for KilovoltsDc, verifying result is not NaN.
		result := a.KilovoltsDc()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilovoltsDc returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsDc.
		// No expected conversion value provided for MegavoltsDc, verifying result is not NaN.
		result := a.MegavoltsDc()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegavoltsDc returned NaN")
		}
	}
}

func TestElectricPotentialDc_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricPotentialDcFactory{}
	a, err := factory.CreateElectricPotentialDc(90, units.ElectricPotentialDcVoltDc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricPotentialDcVoltDc {
		t.Errorf("expected default unit VoltDc, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricPotentialDcVoltDc
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricPotentialDcDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricPotentialDcVoltDc {
		t.Errorf("expected unit VoltDc, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricPotentialDcToString(t *testing.T) {
	factory := units.ElectricPotentialDcFactory{}
	a, err := factory.CreateElectricPotentialDc(45, units.ElectricPotentialDcVoltDc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricPotentialDcVoltDc, 2)
	expected := "45.00 " + units.GetElectricPotentialDcAbbreviation(units.ElectricPotentialDcVoltDc)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricPotentialDcVoltDc, -1)
	expected = "45 " + units.GetElectricPotentialDcAbbreviation(units.ElectricPotentialDcVoltDc)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricPotentialDc_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricPotentialDcFactory{}
	a1, _ := factory.CreateElectricPotentialDc(60, units.ElectricPotentialDcVoltDc)
	a2, _ := factory.CreateElectricPotentialDc(60, units.ElectricPotentialDcVoltDc)
	a3, _ := factory.CreateElectricPotentialDc(120, units.ElectricPotentialDcVoltDc)

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

func TestElectricPotentialDc_Arithmetic(t *testing.T) {
	factory := units.ElectricPotentialDcFactory{}
	a1, _ := factory.CreateElectricPotentialDc(30, units.ElectricPotentialDcVoltDc)
	a2, _ := factory.CreateElectricPotentialDc(45, units.ElectricPotentialDcVoltDc)

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