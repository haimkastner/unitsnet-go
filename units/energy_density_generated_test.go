// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestEnergyDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerCubicMeter"}`
	
	factory := units.EnergyDensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.EnergyDensityJoulePerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.EnergyDensityJoulePerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestEnergyDensityDto_ToJSON(t *testing.T) {
	dto := units.EnergyDensityDto{
		Value: 45,
		Unit:  units.EnergyDensityJoulePerCubicMeter,
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
	if result["unit"].(string) != string(units.EnergyDensityJoulePerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.EnergyDensityJoulePerCubicMeter, result["unit"])
	}
}

func TestNewEnergyDensity_InvalidValue(t *testing.T) {
	factory := units.EnergyDensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateEnergyDensity(math.NaN(), units.EnergyDensityJoulePerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateEnergyDensity(math.Inf(1), units.EnergyDensityJoulePerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestEnergyDensityConversions(t *testing.T) {
	factory := units.EnergyDensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateEnergyDensity(180, units.EnergyDensityJoulePerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerCubicMeter.
		// No expected conversion value provided for JoulesPerCubicMeter, verifying result is not NaN.
		result := a.JoulesPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to WattHoursPerCubicMeter.
		// No expected conversion value provided for WattHoursPerCubicMeter, verifying result is not NaN.
		result := a.WattHoursPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattHoursPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerCubicMeter.
		// No expected conversion value provided for KilojoulesPerCubicMeter, verifying result is not NaN.
		result := a.KilojoulesPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerCubicMeter.
		// No expected conversion value provided for MegajoulesPerCubicMeter, verifying result is not NaN.
		result := a.MegajoulesPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegajoulesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to GigajoulesPerCubicMeter.
		// No expected conversion value provided for GigajoulesPerCubicMeter, verifying result is not NaN.
		result := a.GigajoulesPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigajoulesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to TerajoulesPerCubicMeter.
		// No expected conversion value provided for TerajoulesPerCubicMeter, verifying result is not NaN.
		result := a.TerajoulesPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerajoulesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PetajoulesPerCubicMeter.
		// No expected conversion value provided for PetajoulesPerCubicMeter, verifying result is not NaN.
		result := a.PetajoulesPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PetajoulesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattHoursPerCubicMeter.
		// No expected conversion value provided for KilowattHoursPerCubicMeter, verifying result is not NaN.
		result := a.KilowattHoursPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattHoursPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattHoursPerCubicMeter.
		// No expected conversion value provided for MegawattHoursPerCubicMeter, verifying result is not NaN.
		result := a.MegawattHoursPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattHoursPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to GigawattHoursPerCubicMeter.
		// No expected conversion value provided for GigawattHoursPerCubicMeter, verifying result is not NaN.
		result := a.GigawattHoursPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattHoursPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to TerawattHoursPerCubicMeter.
		// No expected conversion value provided for TerawattHoursPerCubicMeter, verifying result is not NaN.
		result := a.TerawattHoursPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerawattHoursPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PetawattHoursPerCubicMeter.
		// No expected conversion value provided for PetawattHoursPerCubicMeter, verifying result is not NaN.
		result := a.PetawattHoursPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PetawattHoursPerCubicMeter returned NaN")
		}
	}
}

func TestEnergyDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.EnergyDensityFactory{}
	a, err := factory.CreateEnergyDensity(90, units.EnergyDensityJoulePerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.EnergyDensityJoulePerCubicMeter {
		t.Errorf("expected default unit JoulePerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.EnergyDensityJoulePerCubicMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.EnergyDensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.EnergyDensityJoulePerCubicMeter {
		t.Errorf("expected unit JoulePerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestEnergyDensityToString(t *testing.T) {
	factory := units.EnergyDensityFactory{}
	a, err := factory.CreateEnergyDensity(45, units.EnergyDensityJoulePerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.EnergyDensityJoulePerCubicMeter, 2)
	expected := "45.00 " + units.GetEnergyDensityAbbreviation(units.EnergyDensityJoulePerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.EnergyDensityJoulePerCubicMeter, -1)
	expected = "45 " + units.GetEnergyDensityAbbreviation(units.EnergyDensityJoulePerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestEnergyDensity_EqualityAndComparison(t *testing.T) {
	factory := units.EnergyDensityFactory{}
	a1, _ := factory.CreateEnergyDensity(60, units.EnergyDensityJoulePerCubicMeter)
	a2, _ := factory.CreateEnergyDensity(60, units.EnergyDensityJoulePerCubicMeter)
	a3, _ := factory.CreateEnergyDensity(120, units.EnergyDensityJoulePerCubicMeter)

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

func TestEnergyDensity_Arithmetic(t *testing.T) {
	factory := units.EnergyDensityFactory{}
	a1, _ := factory.CreateEnergyDensity(30, units.EnergyDensityJoulePerCubicMeter)
	a2, _ := factory.CreateEnergyDensity(45, units.EnergyDensityJoulePerCubicMeter)

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