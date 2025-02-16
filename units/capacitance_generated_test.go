// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestCapacitanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Farad"}`
	
	factory := units.CapacitanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.CapacitanceFarad {
		t.Errorf("expected unit %v, got %v", units.CapacitanceFarad, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Farad"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestCapacitanceDto_ToJSON(t *testing.T) {
	dto := units.CapacitanceDto{
		Value: 45,
		Unit:  units.CapacitanceFarad,
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
	if result["unit"].(string) != string(units.CapacitanceFarad) {
		t.Errorf("expected unit %s, got %v", units.CapacitanceFarad, result["unit"])
	}
}

func TestNewCapacitance_InvalidValue(t *testing.T) {
	factory := units.CapacitanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateCapacitance(math.NaN(), units.CapacitanceFarad)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateCapacitance(math.Inf(1), units.CapacitanceFarad)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestCapacitanceConversions(t *testing.T) {
	factory := units.CapacitanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateCapacitance(180, units.CapacitanceFarad)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Farads.
		// No expected conversion value provided for Farads, verifying result is not NaN.
		result := a.Farads()
		if math.IsNaN(result) {
			t.Errorf("conversion to Farads returned NaN")
		}
	}
	{
		// Test conversion to Picofarads.
		// No expected conversion value provided for Picofarads, verifying result is not NaN.
		result := a.Picofarads()
		if math.IsNaN(result) {
			t.Errorf("conversion to Picofarads returned NaN")
		}
	}
	{
		// Test conversion to Nanofarads.
		// No expected conversion value provided for Nanofarads, verifying result is not NaN.
		result := a.Nanofarads()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanofarads returned NaN")
		}
	}
	{
		// Test conversion to Microfarads.
		// No expected conversion value provided for Microfarads, verifying result is not NaN.
		result := a.Microfarads()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microfarads returned NaN")
		}
	}
	{
		// Test conversion to Millifarads.
		// No expected conversion value provided for Millifarads, verifying result is not NaN.
		result := a.Millifarads()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millifarads returned NaN")
		}
	}
	{
		// Test conversion to Kilofarads.
		// No expected conversion value provided for Kilofarads, verifying result is not NaN.
		result := a.Kilofarads()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilofarads returned NaN")
		}
	}
	{
		// Test conversion to Megafarads.
		// No expected conversion value provided for Megafarads, verifying result is not NaN.
		result := a.Megafarads()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megafarads returned NaN")
		}
	}
}

func TestCapacitance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.CapacitanceFactory{}
	a, err := factory.CreateCapacitance(90, units.CapacitanceFarad)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.CapacitanceFarad {
		t.Errorf("expected default unit Farad, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.CapacitanceFarad
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.CapacitanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.CapacitanceFarad {
		t.Errorf("expected unit Farad, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestCapacitanceToString(t *testing.T) {
	factory := units.CapacitanceFactory{}
	a, err := factory.CreateCapacitance(45, units.CapacitanceFarad)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.CapacitanceFarad, 2)
	expected := "45.00 " + units.GetCapacitanceAbbreviation(units.CapacitanceFarad)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.CapacitanceFarad, -1)
	expected = "45 " + units.GetCapacitanceAbbreviation(units.CapacitanceFarad)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestCapacitance_EqualityAndComparison(t *testing.T) {
	factory := units.CapacitanceFactory{}
	a1, _ := factory.CreateCapacitance(60, units.CapacitanceFarad)
	a2, _ := factory.CreateCapacitance(60, units.CapacitanceFarad)
	a3, _ := factory.CreateCapacitance(120, units.CapacitanceFarad)

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

func TestCapacitance_Arithmetic(t *testing.T) {
	factory := units.CapacitanceFactory{}
	a1, _ := factory.CreateCapacitance(30, units.CapacitanceFarad)
	a2, _ := factory.CreateCapacitance(45, units.CapacitanceFarad)

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