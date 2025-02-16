// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricApparentPowerDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Voltampere"}`
	
	factory := units.ElectricApparentPowerDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricApparentPowerVoltampere {
		t.Errorf("expected unit %v, got %v", units.ElectricApparentPowerVoltampere, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Voltampere"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricApparentPowerDto_ToJSON(t *testing.T) {
	dto := units.ElectricApparentPowerDto{
		Value: 45,
		Unit:  units.ElectricApparentPowerVoltampere,
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
	if result["unit"].(string) != string(units.ElectricApparentPowerVoltampere) {
		t.Errorf("expected unit %s, got %v", units.ElectricApparentPowerVoltampere, result["unit"])
	}
}

func TestNewElectricApparentPower_InvalidValue(t *testing.T) {
	factory := units.ElectricApparentPowerFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricApparentPower(math.NaN(), units.ElectricApparentPowerVoltampere)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricApparentPower(math.Inf(1), units.ElectricApparentPowerVoltampere)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricApparentPowerConversions(t *testing.T) {
	factory := units.ElectricApparentPowerFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricApparentPower(180, units.ElectricApparentPowerVoltampere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Voltamperes.
		// No expected conversion value provided for Voltamperes, verifying result is not NaN.
		result := a.Voltamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Voltamperes returned NaN")
		}
	}
	{
		// Test conversion to Microvoltamperes.
		// No expected conversion value provided for Microvoltamperes, verifying result is not NaN.
		result := a.Microvoltamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microvoltamperes returned NaN")
		}
	}
	{
		// Test conversion to Millivoltamperes.
		// No expected conversion value provided for Millivoltamperes, verifying result is not NaN.
		result := a.Millivoltamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millivoltamperes returned NaN")
		}
	}
	{
		// Test conversion to Kilovoltamperes.
		// No expected conversion value provided for Kilovoltamperes, verifying result is not NaN.
		result := a.Kilovoltamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilovoltamperes returned NaN")
		}
	}
	{
		// Test conversion to Megavoltamperes.
		// No expected conversion value provided for Megavoltamperes, verifying result is not NaN.
		result := a.Megavoltamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megavoltamperes returned NaN")
		}
	}
	{
		// Test conversion to Gigavoltamperes.
		// No expected conversion value provided for Gigavoltamperes, verifying result is not NaN.
		result := a.Gigavoltamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigavoltamperes returned NaN")
		}
	}
}

func TestElectricApparentPower_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricApparentPowerFactory{}
	a, err := factory.CreateElectricApparentPower(90, units.ElectricApparentPowerVoltampere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricApparentPowerVoltampere {
		t.Errorf("expected default unit Voltampere, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricApparentPowerVoltampere
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricApparentPowerDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricApparentPowerVoltampere {
		t.Errorf("expected unit Voltampere, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricApparentPowerToString(t *testing.T) {
	factory := units.ElectricApparentPowerFactory{}
	a, err := factory.CreateElectricApparentPower(45, units.ElectricApparentPowerVoltampere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricApparentPowerVoltampere, 2)
	expected := "45.00 " + units.GetElectricApparentPowerAbbreviation(units.ElectricApparentPowerVoltampere)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricApparentPowerVoltampere, -1)
	expected = "45 " + units.GetElectricApparentPowerAbbreviation(units.ElectricApparentPowerVoltampere)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricApparentPower_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricApparentPowerFactory{}
	a1, _ := factory.CreateElectricApparentPower(60, units.ElectricApparentPowerVoltampere)
	a2, _ := factory.CreateElectricApparentPower(60, units.ElectricApparentPowerVoltampere)
	a3, _ := factory.CreateElectricApparentPower(120, units.ElectricApparentPowerVoltampere)

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

func TestElectricApparentPower_Arithmetic(t *testing.T) {
	factory := units.ElectricApparentPowerFactory{}
	a1, _ := factory.CreateElectricApparentPower(30, units.ElectricApparentPowerVoltampere)
	a2, _ := factory.CreateElectricApparentPower(45, units.ElectricApparentPowerVoltampere)

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