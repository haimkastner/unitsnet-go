// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestSpecificEntropyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerKilogramKelvin"}`
	
	factory := units.SpecificEntropyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.SpecificEntropyJoulePerKilogramKelvin {
		t.Errorf("expected unit %v, got %v", units.SpecificEntropyJoulePerKilogramKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerKilogramKelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestSpecificEntropyDto_ToJSON(t *testing.T) {
	dto := units.SpecificEntropyDto{
		Value: 45,
		Unit:  units.SpecificEntropyJoulePerKilogramKelvin,
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
	if result["unit"].(string) != string(units.SpecificEntropyJoulePerKilogramKelvin) {
		t.Errorf("expected unit %s, got %v", units.SpecificEntropyJoulePerKilogramKelvin, result["unit"])
	}
}

func TestNewSpecificEntropy_InvalidValue(t *testing.T) {
	factory := units.SpecificEntropyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateSpecificEntropy(math.NaN(), units.SpecificEntropyJoulePerKilogramKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateSpecificEntropy(math.Inf(1), units.SpecificEntropyJoulePerKilogramKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestSpecificEntropyConversions(t *testing.T) {
	factory := units.SpecificEntropyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateSpecificEntropy(180, units.SpecificEntropyJoulePerKilogramKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerKilogramKelvin.
		// No expected conversion value provided for JoulesPerKilogramKelvin, verifying result is not NaN.
		result := a.JoulesPerKilogramKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerKilogramKelvin returned NaN")
		}
	}
	{
		// Test conversion to JoulesPerKilogramDegreeCelsius.
		// No expected conversion value provided for JoulesPerKilogramDegreeCelsius, verifying result is not NaN.
		result := a.JoulesPerKilogramDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerKilogramDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to CaloriesPerGramKelvin.
		// No expected conversion value provided for CaloriesPerGramKelvin, verifying result is not NaN.
		result := a.CaloriesPerGramKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to CaloriesPerGramKelvin returned NaN")
		}
	}
	{
		// Test conversion to BtusPerPoundFahrenheit.
		// No expected conversion value provided for BtusPerPoundFahrenheit, verifying result is not NaN.
		result := a.BtusPerPoundFahrenheit()
		if math.IsNaN(result) {
			t.Errorf("conversion to BtusPerPoundFahrenheit returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerKilogramKelvin.
		// No expected conversion value provided for KilojoulesPerKilogramKelvin, verifying result is not NaN.
		result := a.KilojoulesPerKilogramKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerKilogramKelvin returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerKilogramKelvin.
		// No expected conversion value provided for MegajoulesPerKilogramKelvin, verifying result is not NaN.
		result := a.MegajoulesPerKilogramKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegajoulesPerKilogramKelvin returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerKilogramDegreeCelsius.
		// No expected conversion value provided for KilojoulesPerKilogramDegreeCelsius, verifying result is not NaN.
		result := a.KilojoulesPerKilogramDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerKilogramDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerKilogramDegreeCelsius.
		// No expected conversion value provided for MegajoulesPerKilogramDegreeCelsius, verifying result is not NaN.
		result := a.MegajoulesPerKilogramDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegajoulesPerKilogramDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerGramKelvin.
		// No expected conversion value provided for KilocaloriesPerGramKelvin, verifying result is not NaN.
		result := a.KilocaloriesPerGramKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilocaloriesPerGramKelvin returned NaN")
		}
	}
}

func TestSpecificEntropy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.SpecificEntropyFactory{}
	a, err := factory.CreateSpecificEntropy(90, units.SpecificEntropyJoulePerKilogramKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.SpecificEntropyJoulePerKilogramKelvin {
		t.Errorf("expected default unit JoulePerKilogramKelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.SpecificEntropyJoulePerKilogramKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.SpecificEntropyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.SpecificEntropyJoulePerKilogramKelvin {
		t.Errorf("expected unit JoulePerKilogramKelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestSpecificEntropyToString(t *testing.T) {
	factory := units.SpecificEntropyFactory{}
	a, err := factory.CreateSpecificEntropy(45, units.SpecificEntropyJoulePerKilogramKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.SpecificEntropyJoulePerKilogramKelvin, 2)
	expected := "45.00 " + units.GetSpecificEntropyAbbreviation(units.SpecificEntropyJoulePerKilogramKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.SpecificEntropyJoulePerKilogramKelvin, -1)
	expected = "45 " + units.GetSpecificEntropyAbbreviation(units.SpecificEntropyJoulePerKilogramKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestSpecificEntropy_EqualityAndComparison(t *testing.T) {
	factory := units.SpecificEntropyFactory{}
	a1, _ := factory.CreateSpecificEntropy(60, units.SpecificEntropyJoulePerKilogramKelvin)
	a2, _ := factory.CreateSpecificEntropy(60, units.SpecificEntropyJoulePerKilogramKelvin)
	a3, _ := factory.CreateSpecificEntropy(120, units.SpecificEntropyJoulePerKilogramKelvin)

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

func TestSpecificEntropy_Arithmetic(t *testing.T) {
	factory := units.SpecificEntropyFactory{}
	a1, _ := factory.CreateSpecificEntropy(30, units.SpecificEntropyJoulePerKilogramKelvin)
	a2, _ := factory.CreateSpecificEntropy(45, units.SpecificEntropyJoulePerKilogramKelvin)

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