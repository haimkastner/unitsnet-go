// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricPotentialAcDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltAc"}`
	
	factory := units.ElectricPotentialAcDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricPotentialAcVoltAc {
		t.Errorf("expected unit %v, got %v", units.ElectricPotentialAcVoltAc, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltAc"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricPotentialAcDto_ToJSON(t *testing.T) {
	dto := units.ElectricPotentialAcDto{
		Value: 45,
		Unit:  units.ElectricPotentialAcVoltAc,
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
	if result["unit"].(string) != string(units.ElectricPotentialAcVoltAc) {
		t.Errorf("expected unit %s, got %v", units.ElectricPotentialAcVoltAc, result["unit"])
	}
}

func TestNewElectricPotentialAc_InvalidValue(t *testing.T) {
	factory := units.ElectricPotentialAcFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricPotentialAc(math.NaN(), units.ElectricPotentialAcVoltAc)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricPotentialAc(math.Inf(1), units.ElectricPotentialAcVoltAc)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricPotentialAcConversions(t *testing.T) {
	factory := units.ElectricPotentialAcFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricPotentialAc(180, units.ElectricPotentialAcVoltAc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltsAc.
		// No expected conversion value provided for VoltsAc, verifying result is not NaN.
		result := a.VoltsAc()
		if math.IsNaN(result) {
			t.Errorf("conversion to VoltsAc returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsAc.
		// No expected conversion value provided for MicrovoltsAc, verifying result is not NaN.
		result := a.MicrovoltsAc()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrovoltsAc returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsAc.
		// No expected conversion value provided for MillivoltsAc, verifying result is not NaN.
		result := a.MillivoltsAc()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillivoltsAc returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsAc.
		// No expected conversion value provided for KilovoltsAc, verifying result is not NaN.
		result := a.KilovoltsAc()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilovoltsAc returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsAc.
		// No expected conversion value provided for MegavoltsAc, verifying result is not NaN.
		result := a.MegavoltsAc()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegavoltsAc returned NaN")
		}
	}
}

func TestElectricPotentialAc_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricPotentialAcFactory{}
	a, err := factory.CreateElectricPotentialAc(90, units.ElectricPotentialAcVoltAc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricPotentialAcVoltAc {
		t.Errorf("expected default unit VoltAc, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricPotentialAcVoltAc
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricPotentialAcDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricPotentialAcVoltAc {
		t.Errorf("expected unit VoltAc, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricPotentialAcToString(t *testing.T) {
	factory := units.ElectricPotentialAcFactory{}
	a, err := factory.CreateElectricPotentialAc(45, units.ElectricPotentialAcVoltAc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricPotentialAcVoltAc, 2)
	expected := "45.00 " + units.GetElectricPotentialAcAbbreviation(units.ElectricPotentialAcVoltAc)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricPotentialAcVoltAc, -1)
	expected = "45 " + units.GetElectricPotentialAcAbbreviation(units.ElectricPotentialAcVoltAc)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricPotentialAc_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricPotentialAcFactory{}
	a1, _ := factory.CreateElectricPotentialAc(60, units.ElectricPotentialAcVoltAc)
	a2, _ := factory.CreateElectricPotentialAc(60, units.ElectricPotentialAcVoltAc)
	a3, _ := factory.CreateElectricPotentialAc(120, units.ElectricPotentialAcVoltAc)

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

func TestElectricPotentialAc_Arithmetic(t *testing.T) {
	factory := units.ElectricPotentialAcFactory{}
	a1, _ := factory.CreateElectricPotentialAc(30, units.ElectricPotentialAcVoltAc)
	a2, _ := factory.CreateElectricPotentialAc(45, units.ElectricPotentialAcVoltAc)

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