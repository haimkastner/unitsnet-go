// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestEntropyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerKelvin"}`
	
	factory := units.EntropyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.EntropyJoulePerKelvin {
		t.Errorf("expected unit %v, got %v", units.EntropyJoulePerKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerKelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestEntropyDto_ToJSON(t *testing.T) {
	dto := units.EntropyDto{
		Value: 45,
		Unit:  units.EntropyJoulePerKelvin,
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
	if result["unit"].(string) != string(units.EntropyJoulePerKelvin) {
		t.Errorf("expected unit %s, got %v", units.EntropyJoulePerKelvin, result["unit"])
	}
}

func TestNewEntropy_InvalidValue(t *testing.T) {
	factory := units.EntropyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateEntropy(math.NaN(), units.EntropyJoulePerKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateEntropy(math.Inf(1), units.EntropyJoulePerKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestEntropyConversions(t *testing.T) {
	factory := units.EntropyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateEntropy(180, units.EntropyJoulePerKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerKelvin.
		// No expected conversion value provided for JoulesPerKelvin, verifying result is not NaN.
		result := a.JoulesPerKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerKelvin returned NaN")
		}
	}
	{
		// Test conversion to CaloriesPerKelvin.
		// No expected conversion value provided for CaloriesPerKelvin, verifying result is not NaN.
		result := a.CaloriesPerKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to CaloriesPerKelvin returned NaN")
		}
	}
	{
		// Test conversion to JoulesPerDegreeCelsius.
		// No expected conversion value provided for JoulesPerDegreeCelsius, verifying result is not NaN.
		result := a.JoulesPerDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerKelvin.
		// No expected conversion value provided for KilojoulesPerKelvin, verifying result is not NaN.
		result := a.KilojoulesPerKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerKelvin returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerKelvin.
		// No expected conversion value provided for MegajoulesPerKelvin, verifying result is not NaN.
		result := a.MegajoulesPerKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegajoulesPerKelvin returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerKelvin.
		// No expected conversion value provided for KilocaloriesPerKelvin, verifying result is not NaN.
		result := a.KilocaloriesPerKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilocaloriesPerKelvin returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerDegreeCelsius.
		// No expected conversion value provided for KilojoulesPerDegreeCelsius, verifying result is not NaN.
		result := a.KilojoulesPerDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerDegreeCelsius returned NaN")
		}
	}
}

func TestEntropy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.EntropyFactory{}
	a, err := factory.CreateEntropy(90, units.EntropyJoulePerKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.EntropyJoulePerKelvin {
		t.Errorf("expected default unit JoulePerKelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.EntropyJoulePerKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.EntropyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.EntropyJoulePerKelvin {
		t.Errorf("expected unit JoulePerKelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestEntropyToString(t *testing.T) {
	factory := units.EntropyFactory{}
	a, err := factory.CreateEntropy(45, units.EntropyJoulePerKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.EntropyJoulePerKelvin, 2)
	expected := "45.00 " + units.GetEntropyAbbreviation(units.EntropyJoulePerKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.EntropyJoulePerKelvin, -1)
	expected = "45 " + units.GetEntropyAbbreviation(units.EntropyJoulePerKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestEntropy_EqualityAndComparison(t *testing.T) {
	factory := units.EntropyFactory{}
	a1, _ := factory.CreateEntropy(60, units.EntropyJoulePerKelvin)
	a2, _ := factory.CreateEntropy(60, units.EntropyJoulePerKelvin)
	a3, _ := factory.CreateEntropy(120, units.EntropyJoulePerKelvin)

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

func TestEntropy_Arithmetic(t *testing.T) {
	factory := units.EntropyFactory{}
	a1, _ := factory.CreateEntropy(30, units.EntropyJoulePerKelvin)
	a2, _ := factory.CreateEntropy(45, units.EntropyJoulePerKelvin)

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