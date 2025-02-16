// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestFuelEfficiencyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "LiterPer100Kilometers"}`
	
	factory := units.FuelEfficiencyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.FuelEfficiencyLiterPer100Kilometers {
		t.Errorf("expected unit %v, got %v", units.FuelEfficiencyLiterPer100Kilometers, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "LiterPer100Kilometers"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestFuelEfficiencyDto_ToJSON(t *testing.T) {
	dto := units.FuelEfficiencyDto{
		Value: 45,
		Unit:  units.FuelEfficiencyLiterPer100Kilometers,
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
	if result["unit"].(string) != string(units.FuelEfficiencyLiterPer100Kilometers) {
		t.Errorf("expected unit %s, got %v", units.FuelEfficiencyLiterPer100Kilometers, result["unit"])
	}
}

func TestNewFuelEfficiency_InvalidValue(t *testing.T) {
	factory := units.FuelEfficiencyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateFuelEfficiency(math.NaN(), units.FuelEfficiencyLiterPer100Kilometers)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateFuelEfficiency(math.Inf(1), units.FuelEfficiencyLiterPer100Kilometers)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestFuelEfficiencyConversions(t *testing.T) {
	factory := units.FuelEfficiencyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateFuelEfficiency(180, units.FuelEfficiencyLiterPer100Kilometers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to LitersPer100Kilometers.
		// No expected conversion value provided for LitersPer100Kilometers, verifying result is not NaN.
		result := a.LitersPer100Kilometers()
		if math.IsNaN(result) {
			t.Errorf("conversion to LitersPer100Kilometers returned NaN")
		}
	}
	{
		// Test conversion to MilesPerUsGallon.
		// No expected conversion value provided for MilesPerUsGallon, verifying result is not NaN.
		result := a.MilesPerUsGallon()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilesPerUsGallon returned NaN")
		}
	}
	{
		// Test conversion to MilesPerUkGallon.
		// No expected conversion value provided for MilesPerUkGallon, verifying result is not NaN.
		result := a.MilesPerUkGallon()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilesPerUkGallon returned NaN")
		}
	}
	{
		// Test conversion to KilometersPerLiters.
		// No expected conversion value provided for KilometersPerLiters, verifying result is not NaN.
		result := a.KilometersPerLiters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilometersPerLiters returned NaN")
		}
	}
}

func TestFuelEfficiency_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.FuelEfficiencyFactory{}
	a, err := factory.CreateFuelEfficiency(90, units.FuelEfficiencyLiterPer100Kilometers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.FuelEfficiencyLiterPer100Kilometers {
		t.Errorf("expected default unit LiterPer100Kilometers, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.FuelEfficiencyLiterPer100Kilometers
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.FuelEfficiencyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.FuelEfficiencyLiterPer100Kilometers {
		t.Errorf("expected unit LiterPer100Kilometers, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestFuelEfficiencyToString(t *testing.T) {
	factory := units.FuelEfficiencyFactory{}
	a, err := factory.CreateFuelEfficiency(45, units.FuelEfficiencyLiterPer100Kilometers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.FuelEfficiencyLiterPer100Kilometers, 2)
	expected := "45.00 " + units.GetFuelEfficiencyAbbreviation(units.FuelEfficiencyLiterPer100Kilometers)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.FuelEfficiencyLiterPer100Kilometers, -1)
	expected = "45 " + units.GetFuelEfficiencyAbbreviation(units.FuelEfficiencyLiterPer100Kilometers)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestFuelEfficiency_EqualityAndComparison(t *testing.T) {
	factory := units.FuelEfficiencyFactory{}
	a1, _ := factory.CreateFuelEfficiency(60, units.FuelEfficiencyLiterPer100Kilometers)
	a2, _ := factory.CreateFuelEfficiency(60, units.FuelEfficiencyLiterPer100Kilometers)
	a3, _ := factory.CreateFuelEfficiency(120, units.FuelEfficiencyLiterPer100Kilometers)

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

func TestFuelEfficiency_Arithmetic(t *testing.T) {
	factory := units.FuelEfficiencyFactory{}
	a1, _ := factory.CreateFuelEfficiency(30, units.FuelEfficiencyLiterPer100Kilometers)
	a2, _ := factory.CreateFuelEfficiency(45, units.FuelEfficiencyLiterPer100Kilometers)

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