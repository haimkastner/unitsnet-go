// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMagneticFieldDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Tesla"}`
	
	factory := units.MagneticFieldDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MagneticFieldTesla {
		t.Errorf("expected unit %v, got %v", units.MagneticFieldTesla, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Tesla"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMagneticFieldDto_ToJSON(t *testing.T) {
	dto := units.MagneticFieldDto{
		Value: 45,
		Unit:  units.MagneticFieldTesla,
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
	if result["unit"].(string) != string(units.MagneticFieldTesla) {
		t.Errorf("expected unit %s, got %v", units.MagneticFieldTesla, result["unit"])
	}
}

func TestNewMagneticField_InvalidValue(t *testing.T) {
	factory := units.MagneticFieldFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMagneticField(math.NaN(), units.MagneticFieldTesla)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMagneticField(math.Inf(1), units.MagneticFieldTesla)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMagneticFieldConversions(t *testing.T) {
	factory := units.MagneticFieldFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMagneticField(180, units.MagneticFieldTesla)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Teslas.
		// No expected conversion value provided for Teslas, verifying result is not NaN.
		result := a.Teslas()
		if math.IsNaN(result) {
			t.Errorf("conversion to Teslas returned NaN")
		}
	}
	{
		// Test conversion to Gausses.
		// No expected conversion value provided for Gausses, verifying result is not NaN.
		result := a.Gausses()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gausses returned NaN")
		}
	}
	{
		// Test conversion to Nanoteslas.
		// No expected conversion value provided for Nanoteslas, verifying result is not NaN.
		result := a.Nanoteslas()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanoteslas returned NaN")
		}
	}
	{
		// Test conversion to Microteslas.
		// No expected conversion value provided for Microteslas, verifying result is not NaN.
		result := a.Microteslas()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microteslas returned NaN")
		}
	}
	{
		// Test conversion to Milliteslas.
		// No expected conversion value provided for Milliteslas, verifying result is not NaN.
		result := a.Milliteslas()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milliteslas returned NaN")
		}
	}
	{
		// Test conversion to Milligausses.
		// No expected conversion value provided for Milligausses, verifying result is not NaN.
		result := a.Milligausses()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milligausses returned NaN")
		}
	}
}

func TestMagneticField_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MagneticFieldFactory{}
	a, err := factory.CreateMagneticField(90, units.MagneticFieldTesla)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MagneticFieldTesla {
		t.Errorf("expected default unit Tesla, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MagneticFieldTesla
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MagneticFieldDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MagneticFieldTesla {
		t.Errorf("expected unit Tesla, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMagneticFieldToString(t *testing.T) {
	factory := units.MagneticFieldFactory{}
	a, err := factory.CreateMagneticField(45, units.MagneticFieldTesla)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MagneticFieldTesla, 2)
	expected := "45.00 " + units.GetMagneticFieldAbbreviation(units.MagneticFieldTesla)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MagneticFieldTesla, -1)
	expected = "45 " + units.GetMagneticFieldAbbreviation(units.MagneticFieldTesla)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMagneticField_EqualityAndComparison(t *testing.T) {
	factory := units.MagneticFieldFactory{}
	a1, _ := factory.CreateMagneticField(60, units.MagneticFieldTesla)
	a2, _ := factory.CreateMagneticField(60, units.MagneticFieldTesla)
	a3, _ := factory.CreateMagneticField(120, units.MagneticFieldTesla)

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

func TestMagneticField_Arithmetic(t *testing.T) {
	factory := units.MagneticFieldFactory{}
	a1, _ := factory.CreateMagneticField(30, units.MagneticFieldTesla)
	a2, _ := factory.CreateMagneticField(45, units.MagneticFieldTesla)

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