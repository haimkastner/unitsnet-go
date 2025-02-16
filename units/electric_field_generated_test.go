// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricFieldDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltPerMeter"}`
	
	factory := units.ElectricFieldDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricFieldVoltPerMeter {
		t.Errorf("expected unit %v, got %v", units.ElectricFieldVoltPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricFieldDto_ToJSON(t *testing.T) {
	dto := units.ElectricFieldDto{
		Value: 45,
		Unit:  units.ElectricFieldVoltPerMeter,
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
	if result["unit"].(string) != string(units.ElectricFieldVoltPerMeter) {
		t.Errorf("expected unit %s, got %v", units.ElectricFieldVoltPerMeter, result["unit"])
	}
}

func TestNewElectricField_InvalidValue(t *testing.T) {
	factory := units.ElectricFieldFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricField(math.NaN(), units.ElectricFieldVoltPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricField(math.Inf(1), units.ElectricFieldVoltPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricFieldConversions(t *testing.T) {
	factory := units.ElectricFieldFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricField(180, units.ElectricFieldVoltPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltsPerMeter.
		// No expected conversion value provided for VoltsPerMeter, verifying result is not NaN.
		result := a.VoltsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to VoltsPerMeter returned NaN")
		}
	}
}

func TestElectricField_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricFieldFactory{}
	a, err := factory.CreateElectricField(90, units.ElectricFieldVoltPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricFieldVoltPerMeter {
		t.Errorf("expected default unit VoltPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricFieldVoltPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricFieldDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricFieldVoltPerMeter {
		t.Errorf("expected unit VoltPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricFieldToString(t *testing.T) {
	factory := units.ElectricFieldFactory{}
	a, err := factory.CreateElectricField(45, units.ElectricFieldVoltPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricFieldVoltPerMeter, 2)
	expected := "45.00 " + units.GetElectricFieldAbbreviation(units.ElectricFieldVoltPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricFieldVoltPerMeter, -1)
	expected = "45 " + units.GetElectricFieldAbbreviation(units.ElectricFieldVoltPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricField_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricFieldFactory{}
	a1, _ := factory.CreateElectricField(60, units.ElectricFieldVoltPerMeter)
	a2, _ := factory.CreateElectricField(60, units.ElectricFieldVoltPerMeter)
	a3, _ := factory.CreateElectricField(120, units.ElectricFieldVoltPerMeter)

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

func TestElectricField_Arithmetic(t *testing.T) {
	factory := units.ElectricFieldFactory{}
	a1, _ := factory.CreateElectricField(30, units.ElectricFieldVoltPerMeter)
	a2, _ := factory.CreateElectricField(45, units.ElectricFieldVoltPerMeter)

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