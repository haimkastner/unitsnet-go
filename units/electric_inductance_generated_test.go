// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricInductanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Henry"}`
	
	factory := units.ElectricInductanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricInductanceHenry {
		t.Errorf("expected unit %v, got %v", units.ElectricInductanceHenry, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Henry"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricInductanceDto_ToJSON(t *testing.T) {
	dto := units.ElectricInductanceDto{
		Value: 45,
		Unit:  units.ElectricInductanceHenry,
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
	if result["unit"].(string) != string(units.ElectricInductanceHenry) {
		t.Errorf("expected unit %s, got %v", units.ElectricInductanceHenry, result["unit"])
	}
}

func TestNewElectricInductance_InvalidValue(t *testing.T) {
	factory := units.ElectricInductanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricInductance(math.NaN(), units.ElectricInductanceHenry)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricInductance(math.Inf(1), units.ElectricInductanceHenry)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricInductanceConversions(t *testing.T) {
	factory := units.ElectricInductanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricInductance(180, units.ElectricInductanceHenry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Henries.
		// No expected conversion value provided for Henries, verifying result is not NaN.
		result := a.Henries()
		if math.IsNaN(result) {
			t.Errorf("conversion to Henries returned NaN")
		}
	}
	{
		// Test conversion to Picohenries.
		// No expected conversion value provided for Picohenries, verifying result is not NaN.
		result := a.Picohenries()
		if math.IsNaN(result) {
			t.Errorf("conversion to Picohenries returned NaN")
		}
	}
	{
		// Test conversion to Nanohenries.
		// No expected conversion value provided for Nanohenries, verifying result is not NaN.
		result := a.Nanohenries()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanohenries returned NaN")
		}
	}
	{
		// Test conversion to Microhenries.
		// No expected conversion value provided for Microhenries, verifying result is not NaN.
		result := a.Microhenries()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microhenries returned NaN")
		}
	}
	{
		// Test conversion to Millihenries.
		// No expected conversion value provided for Millihenries, verifying result is not NaN.
		result := a.Millihenries()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millihenries returned NaN")
		}
	}
}

func TestElectricInductance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricInductanceFactory{}
	a, err := factory.CreateElectricInductance(90, units.ElectricInductanceHenry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricInductanceHenry {
		t.Errorf("expected default unit Henry, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricInductanceHenry
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricInductanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricInductanceHenry {
		t.Errorf("expected unit Henry, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricInductanceToString(t *testing.T) {
	factory := units.ElectricInductanceFactory{}
	a, err := factory.CreateElectricInductance(45, units.ElectricInductanceHenry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricInductanceHenry, 2)
	expected := "45.00 " + units.GetElectricInductanceAbbreviation(units.ElectricInductanceHenry)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricInductanceHenry, -1)
	expected = "45 " + units.GetElectricInductanceAbbreviation(units.ElectricInductanceHenry)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricInductance_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricInductanceFactory{}
	a1, _ := factory.CreateElectricInductance(60, units.ElectricInductanceHenry)
	a2, _ := factory.CreateElectricInductance(60, units.ElectricInductanceHenry)
	a3, _ := factory.CreateElectricInductance(120, units.ElectricInductanceHenry)

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

func TestElectricInductance_Arithmetic(t *testing.T) {
	factory := units.ElectricInductanceFactory{}
	a1, _ := factory.CreateElectricInductance(30, units.ElectricInductanceHenry)
	a2, _ := factory.CreateElectricInductance(45, units.ElectricInductanceHenry)

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