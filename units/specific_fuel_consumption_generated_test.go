// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestSpecificFuelConsumptionDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "GramPerKiloNewtonSecond"}`
	
	factory := units.SpecificFuelConsumptionDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.SpecificFuelConsumptionGramPerKiloNewtonSecond {
		t.Errorf("expected unit %v, got %v", units.SpecificFuelConsumptionGramPerKiloNewtonSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "GramPerKiloNewtonSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestSpecificFuelConsumptionDto_ToJSON(t *testing.T) {
	dto := units.SpecificFuelConsumptionDto{
		Value: 45,
		Unit:  units.SpecificFuelConsumptionGramPerKiloNewtonSecond,
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
	if result["unit"].(string) != string(units.SpecificFuelConsumptionGramPerKiloNewtonSecond) {
		t.Errorf("expected unit %s, got %v", units.SpecificFuelConsumptionGramPerKiloNewtonSecond, result["unit"])
	}
}

func TestNewSpecificFuelConsumption_InvalidValue(t *testing.T) {
	factory := units.SpecificFuelConsumptionFactory{}
	// NaN value should return an error.
	_, err := factory.CreateSpecificFuelConsumption(math.NaN(), units.SpecificFuelConsumptionGramPerKiloNewtonSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateSpecificFuelConsumption(math.Inf(1), units.SpecificFuelConsumptionGramPerKiloNewtonSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestSpecificFuelConsumptionConversions(t *testing.T) {
	factory := units.SpecificFuelConsumptionFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateSpecificFuelConsumption(180, units.SpecificFuelConsumptionGramPerKiloNewtonSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to PoundsMassPerPoundForceHour.
		// No expected conversion value provided for PoundsMassPerPoundForceHour, verifying result is not NaN.
		result := a.PoundsMassPerPoundForceHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsMassPerPoundForceHour returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerKilogramForceHour.
		// No expected conversion value provided for KilogramsPerKilogramForceHour, verifying result is not NaN.
		result := a.KilogramsPerKilogramForceHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerKilogramForceHour returned NaN")
		}
	}
	{
		// Test conversion to GramsPerKiloNewtonSecond.
		// No expected conversion value provided for GramsPerKiloNewtonSecond, verifying result is not NaN.
		result := a.GramsPerKiloNewtonSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerKiloNewtonSecond returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerKiloNewtonSecond.
		// No expected conversion value provided for KilogramsPerKiloNewtonSecond, verifying result is not NaN.
		result := a.KilogramsPerKiloNewtonSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerKiloNewtonSecond returned NaN")
		}
	}
}

func TestSpecificFuelConsumption_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.SpecificFuelConsumptionFactory{}
	a, err := factory.CreateSpecificFuelConsumption(90, units.SpecificFuelConsumptionGramPerKiloNewtonSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.SpecificFuelConsumptionGramPerKiloNewtonSecond {
		t.Errorf("expected default unit GramPerKiloNewtonSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.SpecificFuelConsumptionPoundMassPerPoundForceHour
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.SpecificFuelConsumptionDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.SpecificFuelConsumptionGramPerKiloNewtonSecond {
		t.Errorf("expected unit GramPerKiloNewtonSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestSpecificFuelConsumptionToString(t *testing.T) {
	factory := units.SpecificFuelConsumptionFactory{}
	a, err := factory.CreateSpecificFuelConsumption(45, units.SpecificFuelConsumptionGramPerKiloNewtonSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.SpecificFuelConsumptionGramPerKiloNewtonSecond, 2)
	expected := "45.00 " + units.GetSpecificFuelConsumptionAbbreviation(units.SpecificFuelConsumptionGramPerKiloNewtonSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.SpecificFuelConsumptionGramPerKiloNewtonSecond, -1)
	expected = "45 " + units.GetSpecificFuelConsumptionAbbreviation(units.SpecificFuelConsumptionGramPerKiloNewtonSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestSpecificFuelConsumption_EqualityAndComparison(t *testing.T) {
	factory := units.SpecificFuelConsumptionFactory{}
	a1, _ := factory.CreateSpecificFuelConsumption(60, units.SpecificFuelConsumptionGramPerKiloNewtonSecond)
	a2, _ := factory.CreateSpecificFuelConsumption(60, units.SpecificFuelConsumptionGramPerKiloNewtonSecond)
	a3, _ := factory.CreateSpecificFuelConsumption(120, units.SpecificFuelConsumptionGramPerKiloNewtonSecond)

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

func TestSpecificFuelConsumption_Arithmetic(t *testing.T) {
	factory := units.SpecificFuelConsumptionFactory{}
	a1, _ := factory.CreateSpecificFuelConsumption(30, units.SpecificFuelConsumptionGramPerKiloNewtonSecond)
	a2, _ := factory.CreateSpecificFuelConsumption(45, units.SpecificFuelConsumptionGramPerKiloNewtonSecond)

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