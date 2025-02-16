// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricCurrentDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Ampere"}`
	
	factory := units.ElectricCurrentDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricCurrentAmpere {
		t.Errorf("expected unit %v, got %v", units.ElectricCurrentAmpere, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Ampere"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricCurrentDto_ToJSON(t *testing.T) {
	dto := units.ElectricCurrentDto{
		Value: 45,
		Unit:  units.ElectricCurrentAmpere,
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
	if result["unit"].(string) != string(units.ElectricCurrentAmpere) {
		t.Errorf("expected unit %s, got %v", units.ElectricCurrentAmpere, result["unit"])
	}
}

func TestNewElectricCurrent_InvalidValue(t *testing.T) {
	factory := units.ElectricCurrentFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricCurrent(math.NaN(), units.ElectricCurrentAmpere)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricCurrent(math.Inf(1), units.ElectricCurrentAmpere)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricCurrentConversions(t *testing.T) {
	factory := units.ElectricCurrentFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricCurrent(180, units.ElectricCurrentAmpere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Amperes.
		// No expected conversion value provided for Amperes, verifying result is not NaN.
		result := a.Amperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Amperes returned NaN")
		}
	}
	{
		// Test conversion to Femtoamperes.
		// No expected conversion value provided for Femtoamperes, verifying result is not NaN.
		result := a.Femtoamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Femtoamperes returned NaN")
		}
	}
	{
		// Test conversion to Picoamperes.
		// No expected conversion value provided for Picoamperes, verifying result is not NaN.
		result := a.Picoamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Picoamperes returned NaN")
		}
	}
	{
		// Test conversion to Nanoamperes.
		// No expected conversion value provided for Nanoamperes, verifying result is not NaN.
		result := a.Nanoamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanoamperes returned NaN")
		}
	}
	{
		// Test conversion to Microamperes.
		// No expected conversion value provided for Microamperes, verifying result is not NaN.
		result := a.Microamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microamperes returned NaN")
		}
	}
	{
		// Test conversion to Milliamperes.
		// No expected conversion value provided for Milliamperes, verifying result is not NaN.
		result := a.Milliamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milliamperes returned NaN")
		}
	}
	{
		// Test conversion to Centiamperes.
		// No expected conversion value provided for Centiamperes, verifying result is not NaN.
		result := a.Centiamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Centiamperes returned NaN")
		}
	}
	{
		// Test conversion to Kiloamperes.
		// No expected conversion value provided for Kiloamperes, verifying result is not NaN.
		result := a.Kiloamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kiloamperes returned NaN")
		}
	}
	{
		// Test conversion to Megaamperes.
		// No expected conversion value provided for Megaamperes, verifying result is not NaN.
		result := a.Megaamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megaamperes returned NaN")
		}
	}
}

func TestElectricCurrent_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricCurrentFactory{}
	a, err := factory.CreateElectricCurrent(90, units.ElectricCurrentAmpere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricCurrentAmpere {
		t.Errorf("expected default unit Ampere, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricCurrentAmpere
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricCurrentDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricCurrentAmpere {
		t.Errorf("expected unit Ampere, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricCurrentToString(t *testing.T) {
	factory := units.ElectricCurrentFactory{}
	a, err := factory.CreateElectricCurrent(45, units.ElectricCurrentAmpere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricCurrentAmpere, 2)
	expected := "45.00 " + units.GetElectricCurrentAbbreviation(units.ElectricCurrentAmpere)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricCurrentAmpere, -1)
	expected = "45 " + units.GetElectricCurrentAbbreviation(units.ElectricCurrentAmpere)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricCurrent_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricCurrentFactory{}
	a1, _ := factory.CreateElectricCurrent(60, units.ElectricCurrentAmpere)
	a2, _ := factory.CreateElectricCurrent(60, units.ElectricCurrentAmpere)
	a3, _ := factory.CreateElectricCurrent(120, units.ElectricCurrentAmpere)

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

func TestElectricCurrent_Arithmetic(t *testing.T) {
	factory := units.ElectricCurrentFactory{}
	a1, _ := factory.CreateElectricCurrent(30, units.ElectricCurrentAmpere)
	a2, _ := factory.CreateElectricCurrent(45, units.ElectricCurrentAmpere)

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