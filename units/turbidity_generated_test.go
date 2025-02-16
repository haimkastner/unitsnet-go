// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestTurbidityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NTU"}`
	
	factory := units.TurbidityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.TurbidityNTU {
		t.Errorf("expected unit %v, got %v", units.TurbidityNTU, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NTU"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestTurbidityDto_ToJSON(t *testing.T) {
	dto := units.TurbidityDto{
		Value: 45,
		Unit:  units.TurbidityNTU,
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
	if result["unit"].(string) != string(units.TurbidityNTU) {
		t.Errorf("expected unit %s, got %v", units.TurbidityNTU, result["unit"])
	}
}

func TestNewTurbidity_InvalidValue(t *testing.T) {
	factory := units.TurbidityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateTurbidity(math.NaN(), units.TurbidityNTU)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateTurbidity(math.Inf(1), units.TurbidityNTU)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestTurbidityConversions(t *testing.T) {
	factory := units.TurbidityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateTurbidity(180, units.TurbidityNTU)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NTU.
		// No expected conversion value provided for NTU, verifying result is not NaN.
		result := a.NTU()
		if math.IsNaN(result) {
			t.Errorf("conversion to NTU returned NaN")
		}
	}
}

func TestTurbidity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.TurbidityFactory{}
	a, err := factory.CreateTurbidity(90, units.TurbidityNTU)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.TurbidityNTU {
		t.Errorf("expected default unit NTU, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.TurbidityNTU
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.TurbidityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.TurbidityNTU {
		t.Errorf("expected unit NTU, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestTurbidityToString(t *testing.T) {
	factory := units.TurbidityFactory{}
	a, err := factory.CreateTurbidity(45, units.TurbidityNTU)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.TurbidityNTU, 2)
	expected := "45.00 " + units.GetTurbidityAbbreviation(units.TurbidityNTU)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.TurbidityNTU, -1)
	expected = "45 " + units.GetTurbidityAbbreviation(units.TurbidityNTU)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestTurbidity_EqualityAndComparison(t *testing.T) {
	factory := units.TurbidityFactory{}
	a1, _ := factory.CreateTurbidity(60, units.TurbidityNTU)
	a2, _ := factory.CreateTurbidity(60, units.TurbidityNTU)
	a3, _ := factory.CreateTurbidity(120, units.TurbidityNTU)

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

func TestTurbidity_Arithmetic(t *testing.T) {
	factory := units.TurbidityFactory{}
	a1, _ := factory.CreateTurbidity(30, units.TurbidityNTU)
	a2, _ := factory.CreateTurbidity(45, units.TurbidityNTU)

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