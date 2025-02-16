// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestApparentEnergyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltampereHour"}`
	
	factory := units.ApparentEnergyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ApparentEnergyVoltampereHour {
		t.Errorf("expected unit %v, got %v", units.ApparentEnergyVoltampereHour, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltampereHour"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestApparentEnergyDto_ToJSON(t *testing.T) {
	dto := units.ApparentEnergyDto{
		Value: 45,
		Unit:  units.ApparentEnergyVoltampereHour,
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
	if result["unit"].(string) != string(units.ApparentEnergyVoltampereHour) {
		t.Errorf("expected unit %s, got %v", units.ApparentEnergyVoltampereHour, result["unit"])
	}
}

func TestNewApparentEnergy_InvalidValue(t *testing.T) {
	factory := units.ApparentEnergyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateApparentEnergy(math.NaN(), units.ApparentEnergyVoltampereHour)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateApparentEnergy(math.Inf(1), units.ApparentEnergyVoltampereHour)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestApparentEnergyConversions(t *testing.T) {
	factory := units.ApparentEnergyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateApparentEnergy(180, units.ApparentEnergyVoltampereHour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltampereHours.
		// No expected conversion value provided for VoltampereHours, verifying result is not NaN.
		result := a.VoltampereHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to VoltampereHours returned NaN")
		}
	}
	{
		// Test conversion to KilovoltampereHours.
		// No expected conversion value provided for KilovoltampereHours, verifying result is not NaN.
		result := a.KilovoltampereHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilovoltampereHours returned NaN")
		}
	}
	{
		// Test conversion to MegavoltampereHours.
		// No expected conversion value provided for MegavoltampereHours, verifying result is not NaN.
		result := a.MegavoltampereHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegavoltampereHours returned NaN")
		}
	}
}

func TestApparentEnergy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ApparentEnergyFactory{}
	a, err := factory.CreateApparentEnergy(90, units.ApparentEnergyVoltampereHour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ApparentEnergyVoltampereHour {
		t.Errorf("expected default unit VoltampereHour, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ApparentEnergyVoltampereHour
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ApparentEnergyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ApparentEnergyVoltampereHour {
		t.Errorf("expected unit VoltampereHour, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestApparentEnergyToString(t *testing.T) {
	factory := units.ApparentEnergyFactory{}
	a, err := factory.CreateApparentEnergy(45, units.ApparentEnergyVoltampereHour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ApparentEnergyVoltampereHour, 2)
	expected := "45.00 " + units.GetApparentEnergyAbbreviation(units.ApparentEnergyVoltampereHour)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ApparentEnergyVoltampereHour, -1)
	expected = "45 " + units.GetApparentEnergyAbbreviation(units.ApparentEnergyVoltampereHour)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestApparentEnergy_EqualityAndComparison(t *testing.T) {
	factory := units.ApparentEnergyFactory{}
	a1, _ := factory.CreateApparentEnergy(60, units.ApparentEnergyVoltampereHour)
	a2, _ := factory.CreateApparentEnergy(60, units.ApparentEnergyVoltampereHour)
	a3, _ := factory.CreateApparentEnergy(120, units.ApparentEnergyVoltampereHour)

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

func TestApparentEnergy_Arithmetic(t *testing.T) {
	factory := units.ApparentEnergyFactory{}
	a1, _ := factory.CreateApparentEnergy(30, units.ApparentEnergyVoltampereHour)
	a2, _ := factory.CreateApparentEnergy(45, units.ApparentEnergyVoltampereHour)

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