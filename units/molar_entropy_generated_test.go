// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMolarEntropyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerMoleKelvin"}`
	
	factory := units.MolarEntropyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MolarEntropyJoulePerMoleKelvin {
		t.Errorf("expected unit %v, got %v", units.MolarEntropyJoulePerMoleKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerMoleKelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMolarEntropyDto_ToJSON(t *testing.T) {
	dto := units.MolarEntropyDto{
		Value: 45,
		Unit:  units.MolarEntropyJoulePerMoleKelvin,
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
	if result["unit"].(string) != string(units.MolarEntropyJoulePerMoleKelvin) {
		t.Errorf("expected unit %s, got %v", units.MolarEntropyJoulePerMoleKelvin, result["unit"])
	}
}

func TestNewMolarEntropy_InvalidValue(t *testing.T) {
	factory := units.MolarEntropyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMolarEntropy(math.NaN(), units.MolarEntropyJoulePerMoleKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMolarEntropy(math.Inf(1), units.MolarEntropyJoulePerMoleKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMolarEntropyConversions(t *testing.T) {
	factory := units.MolarEntropyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMolarEntropy(180, units.MolarEntropyJoulePerMoleKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerMoleKelvin.
		// No expected conversion value provided for JoulesPerMoleKelvin, verifying result is not NaN.
		result := a.JoulesPerMoleKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerMoleKelvin returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerMoleKelvin.
		// No expected conversion value provided for KilojoulesPerMoleKelvin, verifying result is not NaN.
		result := a.KilojoulesPerMoleKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerMoleKelvin returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerMoleKelvin.
		// No expected conversion value provided for MegajoulesPerMoleKelvin, verifying result is not NaN.
		result := a.MegajoulesPerMoleKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegajoulesPerMoleKelvin returned NaN")
		}
	}
}

func TestMolarEntropy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MolarEntropyFactory{}
	a, err := factory.CreateMolarEntropy(90, units.MolarEntropyJoulePerMoleKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MolarEntropyJoulePerMoleKelvin {
		t.Errorf("expected default unit JoulePerMoleKelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MolarEntropyJoulePerMoleKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MolarEntropyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MolarEntropyJoulePerMoleKelvin {
		t.Errorf("expected unit JoulePerMoleKelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMolarEntropyToString(t *testing.T) {
	factory := units.MolarEntropyFactory{}
	a, err := factory.CreateMolarEntropy(45, units.MolarEntropyJoulePerMoleKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MolarEntropyJoulePerMoleKelvin, 2)
	expected := "45.00 " + units.GetMolarEntropyAbbreviation(units.MolarEntropyJoulePerMoleKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MolarEntropyJoulePerMoleKelvin, -1)
	expected = "45 " + units.GetMolarEntropyAbbreviation(units.MolarEntropyJoulePerMoleKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMolarEntropy_EqualityAndComparison(t *testing.T) {
	factory := units.MolarEntropyFactory{}
	a1, _ := factory.CreateMolarEntropy(60, units.MolarEntropyJoulePerMoleKelvin)
	a2, _ := factory.CreateMolarEntropy(60, units.MolarEntropyJoulePerMoleKelvin)
	a3, _ := factory.CreateMolarEntropy(120, units.MolarEntropyJoulePerMoleKelvin)

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

func TestMolarEntropy_Arithmetic(t *testing.T) {
	factory := units.MolarEntropyFactory{}
	a1, _ := factory.CreateMolarEntropy(30, units.MolarEntropyJoulePerMoleKelvin)
	a2, _ := factory.CreateMolarEntropy(45, units.MolarEntropyJoulePerMoleKelvin)

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