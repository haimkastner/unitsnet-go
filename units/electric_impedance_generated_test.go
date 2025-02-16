// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricImpedanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Ohm"}`
	
	factory := units.ElectricImpedanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricImpedanceOhm {
		t.Errorf("expected unit %v, got %v", units.ElectricImpedanceOhm, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Ohm"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricImpedanceDto_ToJSON(t *testing.T) {
	dto := units.ElectricImpedanceDto{
		Value: 45,
		Unit:  units.ElectricImpedanceOhm,
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
	if result["unit"].(string) != string(units.ElectricImpedanceOhm) {
		t.Errorf("expected unit %s, got %v", units.ElectricImpedanceOhm, result["unit"])
	}
}

func TestNewElectricImpedance_InvalidValue(t *testing.T) {
	factory := units.ElectricImpedanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricImpedance(math.NaN(), units.ElectricImpedanceOhm)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricImpedance(math.Inf(1), units.ElectricImpedanceOhm)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricImpedanceConversions(t *testing.T) {
	factory := units.ElectricImpedanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricImpedance(180, units.ElectricImpedanceOhm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Ohms.
		// No expected conversion value provided for Ohms, verifying result is not NaN.
		result := a.Ohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Ohms returned NaN")
		}
	}
	{
		// Test conversion to Nanoohms.
		// No expected conversion value provided for Nanoohms, verifying result is not NaN.
		result := a.Nanoohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanoohms returned NaN")
		}
	}
	{
		// Test conversion to Microohms.
		// No expected conversion value provided for Microohms, verifying result is not NaN.
		result := a.Microohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microohms returned NaN")
		}
	}
	{
		// Test conversion to Milliohms.
		// No expected conversion value provided for Milliohms, verifying result is not NaN.
		result := a.Milliohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milliohms returned NaN")
		}
	}
	{
		// Test conversion to Kiloohms.
		// No expected conversion value provided for Kiloohms, verifying result is not NaN.
		result := a.Kiloohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kiloohms returned NaN")
		}
	}
	{
		// Test conversion to Megaohms.
		// No expected conversion value provided for Megaohms, verifying result is not NaN.
		result := a.Megaohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megaohms returned NaN")
		}
	}
	{
		// Test conversion to Gigaohms.
		// No expected conversion value provided for Gigaohms, verifying result is not NaN.
		result := a.Gigaohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigaohms returned NaN")
		}
	}
	{
		// Test conversion to Teraohms.
		// No expected conversion value provided for Teraohms, verifying result is not NaN.
		result := a.Teraohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Teraohms returned NaN")
		}
	}
}

func TestElectricImpedance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricImpedanceFactory{}
	a, err := factory.CreateElectricImpedance(90, units.ElectricImpedanceOhm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricImpedanceOhm {
		t.Errorf("expected default unit Ohm, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricImpedanceOhm
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricImpedanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricImpedanceOhm {
		t.Errorf("expected unit Ohm, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricImpedanceToString(t *testing.T) {
	factory := units.ElectricImpedanceFactory{}
	a, err := factory.CreateElectricImpedance(45, units.ElectricImpedanceOhm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricImpedanceOhm, 2)
	expected := "45.00 " + units.GetElectricImpedanceAbbreviation(units.ElectricImpedanceOhm)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricImpedanceOhm, -1)
	expected = "45 " + units.GetElectricImpedanceAbbreviation(units.ElectricImpedanceOhm)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricImpedance_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricImpedanceFactory{}
	a1, _ := factory.CreateElectricImpedance(60, units.ElectricImpedanceOhm)
	a2, _ := factory.CreateElectricImpedance(60, units.ElectricImpedanceOhm)
	a3, _ := factory.CreateElectricImpedance(120, units.ElectricImpedanceOhm)

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

func TestElectricImpedance_Arithmetic(t *testing.T) {
	factory := units.ElectricImpedanceFactory{}
	a1, _ := factory.CreateElectricImpedance(30, units.ElectricImpedanceOhm)
	a2, _ := factory.CreateElectricImpedance(45, units.ElectricImpedanceOhm)

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