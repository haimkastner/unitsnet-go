// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestApparentPowerDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Voltampere"}`
	
	factory := units.ApparentPowerDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ApparentPowerVoltampere {
		t.Errorf("expected unit %v, got %v", units.ApparentPowerVoltampere, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Voltampere"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestApparentPowerDto_ToJSON(t *testing.T) {
	dto := units.ApparentPowerDto{
		Value: 45,
		Unit:  units.ApparentPowerVoltampere,
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
	if result["unit"].(string) != string(units.ApparentPowerVoltampere) {
		t.Errorf("expected unit %s, got %v", units.ApparentPowerVoltampere, result["unit"])
	}
}

func TestNewApparentPower_InvalidValue(t *testing.T) {
	factory := units.ApparentPowerFactory{}
	// NaN value should return an error.
	_, err := factory.CreateApparentPower(math.NaN(), units.ApparentPowerVoltampere)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateApparentPower(math.Inf(1), units.ApparentPowerVoltampere)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestApparentPowerConversions(t *testing.T) {
	factory := units.ApparentPowerFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateApparentPower(180, units.ApparentPowerVoltampere)
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

func TestApparentPower_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ApparentPowerFactory{}
	a, err := factory.CreateApparentPower(90, units.ApparentPowerVoltampere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ApparentPowerVoltampere {
		t.Errorf("expected default unit Voltampere, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ApparentPowerVoltampere
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ApparentPowerDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ApparentPowerVoltampere {
		t.Errorf("expected unit Voltampere, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestApparentPowerToString(t *testing.T) {
	factory := units.ApparentPowerFactory{}
	a, err := factory.CreateApparentPower(45, units.ApparentPowerVoltampere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ApparentPowerVoltampere, 2)
	expected := "45.00 " + units.GetApparentPowerAbbreviation(units.ApparentPowerVoltampere)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ApparentPowerVoltampere, -1)
	expected = "45 " + units.GetApparentPowerAbbreviation(units.ApparentPowerVoltampere)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestApparentPower_EqualityAndComparison(t *testing.T) {
	factory := units.ApparentPowerFactory{}
	a1, _ := factory.CreateApparentPower(60, units.ApparentPowerVoltampere)
	a2, _ := factory.CreateApparentPower(60, units.ApparentPowerVoltampere)
	a3, _ := factory.CreateApparentPower(120, units.ApparentPowerVoltampere)

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

func TestApparentPower_Arithmetic(t *testing.T) {
	factory := units.ApparentPowerFactory{}
	a1, _ := factory.CreateApparentPower(30, units.ApparentPowerVoltampere)
	a2, _ := factory.CreateApparentPower(45, units.ApparentPowerVoltampere)

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