// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLevelDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Decibel"}`
	
	factory := units.LevelDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LevelDecibel {
		t.Errorf("expected unit %v, got %v", units.LevelDecibel, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Decibel"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLevelDto_ToJSON(t *testing.T) {
	dto := units.LevelDto{
		Value: 45,
		Unit:  units.LevelDecibel,
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
	if result["unit"].(string) != string(units.LevelDecibel) {
		t.Errorf("expected unit %s, got %v", units.LevelDecibel, result["unit"])
	}
}

func TestNewLevel_InvalidValue(t *testing.T) {
	factory := units.LevelFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLevel(math.NaN(), units.LevelDecibel)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLevel(math.Inf(1), units.LevelDecibel)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLevelConversions(t *testing.T) {
	factory := units.LevelFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLevel(180, units.LevelDecibel)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Decibels.
		// No expected conversion value provided for Decibels, verifying result is not NaN.
		result := a.Decibels()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decibels returned NaN")
		}
	}
	{
		// Test conversion to Nepers.
		// No expected conversion value provided for Nepers, verifying result is not NaN.
		result := a.Nepers()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nepers returned NaN")
		}
	}
}

func TestLevel_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LevelFactory{}
	a, err := factory.CreateLevel(90, units.LevelDecibel)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LevelDecibel {
		t.Errorf("expected default unit Decibel, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LevelDecibel
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LevelDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LevelDecibel {
		t.Errorf("expected unit Decibel, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLevelToString(t *testing.T) {
	factory := units.LevelFactory{}
	a, err := factory.CreateLevel(45, units.LevelDecibel)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LevelDecibel, 2)
	expected := "45.00 " + units.GetLevelAbbreviation(units.LevelDecibel)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LevelDecibel, -1)
	expected = "45 " + units.GetLevelAbbreviation(units.LevelDecibel)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLevel_EqualityAndComparison(t *testing.T) {
	factory := units.LevelFactory{}
	a1, _ := factory.CreateLevel(60, units.LevelDecibel)
	a2, _ := factory.CreateLevel(60, units.LevelDecibel)
	a3, _ := factory.CreateLevel(120, units.LevelDecibel)

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

func TestLevel_Arithmetic(t *testing.T) {
	factory := units.LevelFactory{}
	a1, _ := factory.CreateLevel(30, units.LevelDecibel)
	a2, _ := factory.CreateLevel(45, units.LevelDecibel)

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