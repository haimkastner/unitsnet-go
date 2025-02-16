// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricPotentialDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Volt"}`
	
	factory := units.ElectricPotentialDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricPotentialVolt {
		t.Errorf("expected unit %v, got %v", units.ElectricPotentialVolt, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Volt"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricPotentialDto_ToJSON(t *testing.T) {
	dto := units.ElectricPotentialDto{
		Value: 45,
		Unit:  units.ElectricPotentialVolt,
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
	if result["unit"].(string) != string(units.ElectricPotentialVolt) {
		t.Errorf("expected unit %s, got %v", units.ElectricPotentialVolt, result["unit"])
	}
}

func TestNewElectricPotential_InvalidValue(t *testing.T) {
	factory := units.ElectricPotentialFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricPotential(math.NaN(), units.ElectricPotentialVolt)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricPotential(math.Inf(1), units.ElectricPotentialVolt)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricPotentialConversions(t *testing.T) {
	factory := units.ElectricPotentialFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricPotential(180, units.ElectricPotentialVolt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Volts.
		// No expected conversion value provided for Volts, verifying result is not NaN.
		result := a.Volts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Volts returned NaN")
		}
	}
	{
		// Test conversion to Nanovolts.
		// No expected conversion value provided for Nanovolts, verifying result is not NaN.
		result := a.Nanovolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanovolts returned NaN")
		}
	}
	{
		// Test conversion to Microvolts.
		// No expected conversion value provided for Microvolts, verifying result is not NaN.
		result := a.Microvolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microvolts returned NaN")
		}
	}
	{
		// Test conversion to Millivolts.
		// No expected conversion value provided for Millivolts, verifying result is not NaN.
		result := a.Millivolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millivolts returned NaN")
		}
	}
	{
		// Test conversion to Kilovolts.
		// No expected conversion value provided for Kilovolts, verifying result is not NaN.
		result := a.Kilovolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilovolts returned NaN")
		}
	}
	{
		// Test conversion to Megavolts.
		// No expected conversion value provided for Megavolts, verifying result is not NaN.
		result := a.Megavolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megavolts returned NaN")
		}
	}
}

func TestElectricPotential_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricPotentialFactory{}
	a, err := factory.CreateElectricPotential(90, units.ElectricPotentialVolt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricPotentialVolt {
		t.Errorf("expected default unit Volt, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricPotentialVolt
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricPotentialDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricPotentialVolt {
		t.Errorf("expected unit Volt, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricPotentialToString(t *testing.T) {
	factory := units.ElectricPotentialFactory{}
	a, err := factory.CreateElectricPotential(45, units.ElectricPotentialVolt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricPotentialVolt, 2)
	expected := "45.00 " + units.GetElectricPotentialAbbreviation(units.ElectricPotentialVolt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricPotentialVolt, -1)
	expected = "45 " + units.GetElectricPotentialAbbreviation(units.ElectricPotentialVolt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricPotential_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricPotentialFactory{}
	a1, _ := factory.CreateElectricPotential(60, units.ElectricPotentialVolt)
	a2, _ := factory.CreateElectricPotential(60, units.ElectricPotentialVolt)
	a3, _ := factory.CreateElectricPotential(120, units.ElectricPotentialVolt)

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

func TestElectricPotential_Arithmetic(t *testing.T) {
	factory := units.ElectricPotentialFactory{}
	a1, _ := factory.CreateElectricPotential(30, units.ElectricPotentialVolt)
	a2, _ := factory.CreateElectricPotential(45, units.ElectricPotentialVolt)

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