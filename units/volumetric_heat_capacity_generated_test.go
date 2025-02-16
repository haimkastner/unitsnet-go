// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestVolumetricHeatCapacityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerCubicMeterKelvin"}`
	
	factory := units.VolumetricHeatCapacityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.VolumetricHeatCapacityJoulePerCubicMeterKelvin {
		t.Errorf("expected unit %v, got %v", units.VolumetricHeatCapacityJoulePerCubicMeterKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerCubicMeterKelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestVolumetricHeatCapacityDto_ToJSON(t *testing.T) {
	dto := units.VolumetricHeatCapacityDto{
		Value: 45,
		Unit:  units.VolumetricHeatCapacityJoulePerCubicMeterKelvin,
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
	if result["unit"].(string) != string(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin) {
		t.Errorf("expected unit %s, got %v", units.VolumetricHeatCapacityJoulePerCubicMeterKelvin, result["unit"])
	}
}

func TestNewVolumetricHeatCapacity_InvalidValue(t *testing.T) {
	factory := units.VolumetricHeatCapacityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateVolumetricHeatCapacity(math.NaN(), units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateVolumetricHeatCapacity(math.Inf(1), units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestVolumetricHeatCapacityConversions(t *testing.T) {
	factory := units.VolumetricHeatCapacityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateVolumetricHeatCapacity(180, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerCubicMeterKelvin.
		// No expected conversion value provided for JoulesPerCubicMeterKelvin, verifying result is not NaN.
		result := a.JoulesPerCubicMeterKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerCubicMeterKelvin returned NaN")
		}
	}
	{
		// Test conversion to JoulesPerCubicMeterDegreeCelsius.
		// No expected conversion value provided for JoulesPerCubicMeterDegreeCelsius, verifying result is not NaN.
		result := a.JoulesPerCubicMeterDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerCubicMeterDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to CaloriesPerCubicCentimeterDegreeCelsius.
		// No expected conversion value provided for CaloriesPerCubicCentimeterDegreeCelsius, verifying result is not NaN.
		result := a.CaloriesPerCubicCentimeterDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to CaloriesPerCubicCentimeterDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to BtusPerCubicFootDegreeFahrenheit.
		// No expected conversion value provided for BtusPerCubicFootDegreeFahrenheit, verifying result is not NaN.
		result := a.BtusPerCubicFootDegreeFahrenheit()
		if math.IsNaN(result) {
			t.Errorf("conversion to BtusPerCubicFootDegreeFahrenheit returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerCubicMeterKelvin.
		// No expected conversion value provided for KilojoulesPerCubicMeterKelvin, verifying result is not NaN.
		result := a.KilojoulesPerCubicMeterKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerCubicMeterKelvin returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerCubicMeterKelvin.
		// No expected conversion value provided for MegajoulesPerCubicMeterKelvin, verifying result is not NaN.
		result := a.MegajoulesPerCubicMeterKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegajoulesPerCubicMeterKelvin returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerCubicMeterDegreeCelsius.
		// No expected conversion value provided for KilojoulesPerCubicMeterDegreeCelsius, verifying result is not NaN.
		result := a.KilojoulesPerCubicMeterDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerCubicMeterDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerCubicMeterDegreeCelsius.
		// No expected conversion value provided for MegajoulesPerCubicMeterDegreeCelsius, verifying result is not NaN.
		result := a.MegajoulesPerCubicMeterDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegajoulesPerCubicMeterDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerCubicCentimeterDegreeCelsius.
		// No expected conversion value provided for KilocaloriesPerCubicCentimeterDegreeCelsius, verifying result is not NaN.
		result := a.KilocaloriesPerCubicCentimeterDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilocaloriesPerCubicCentimeterDegreeCelsius returned NaN")
		}
	}
}

func TestVolumetricHeatCapacity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.VolumetricHeatCapacityFactory{}
	a, err := factory.CreateVolumetricHeatCapacity(90, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.VolumetricHeatCapacityJoulePerCubicMeterKelvin {
		t.Errorf("expected default unit JoulePerCubicMeterKelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.VolumetricHeatCapacityJoulePerCubicMeterKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.VolumetricHeatCapacityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.VolumetricHeatCapacityJoulePerCubicMeterKelvin {
		t.Errorf("expected unit JoulePerCubicMeterKelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestVolumetricHeatCapacityToString(t *testing.T) {
	factory := units.VolumetricHeatCapacityFactory{}
	a, err := factory.CreateVolumetricHeatCapacity(45, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin, 2)
	expected := "45.00 " + units.GetVolumetricHeatCapacityAbbreviation(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin, -1)
	expected = "45 " + units.GetVolumetricHeatCapacityAbbreviation(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestVolumetricHeatCapacity_EqualityAndComparison(t *testing.T) {
	factory := units.VolumetricHeatCapacityFactory{}
	a1, _ := factory.CreateVolumetricHeatCapacity(60, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	a2, _ := factory.CreateVolumetricHeatCapacity(60, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	a3, _ := factory.CreateVolumetricHeatCapacity(120, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)

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

func TestVolumetricHeatCapacity_Arithmetic(t *testing.T) {
	factory := units.VolumetricHeatCapacityFactory{}
	a1, _ := factory.CreateVolumetricHeatCapacity(30, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	a2, _ := factory.CreateVolumetricHeatCapacity(45, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)

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