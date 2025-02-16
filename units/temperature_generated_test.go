// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestTemperatureDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Kelvin"}`
	
	factory := units.TemperatureDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.TemperatureKelvin {
		t.Errorf("expected unit %v, got %v", units.TemperatureKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Kelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestTemperatureDto_ToJSON(t *testing.T) {
	dto := units.TemperatureDto{
		Value: 45,
		Unit:  units.TemperatureKelvin,
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
	if result["unit"].(string) != string(units.TemperatureKelvin) {
		t.Errorf("expected unit %s, got %v", units.TemperatureKelvin, result["unit"])
	}
}

func TestNewTemperature_InvalidValue(t *testing.T) {
	factory := units.TemperatureFactory{}
	// NaN value should return an error.
	_, err := factory.CreateTemperature(math.NaN(), units.TemperatureKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateTemperature(math.Inf(1), units.TemperatureKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestTemperatureConversions(t *testing.T) {
	factory := units.TemperatureFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateTemperature(180, units.TemperatureKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Kelvins.
		// No expected conversion value provided for Kelvins, verifying result is not NaN.
		result := a.Kelvins()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kelvins returned NaN")
		}
	}
	{
		// Test conversion to DegreesCelsius.
		// No expected conversion value provided for DegreesCelsius, verifying result is not NaN.
		result := a.DegreesCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesCelsius returned NaN")
		}
	}
	{
		// Test conversion to MillidegreesCelsius.
		// No expected conversion value provided for MillidegreesCelsius, verifying result is not NaN.
		result := a.MillidegreesCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillidegreesCelsius returned NaN")
		}
	}
	{
		// Test conversion to DegreesDelisle.
		// No expected conversion value provided for DegreesDelisle, verifying result is not NaN.
		result := a.DegreesDelisle()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesDelisle returned NaN")
		}
	}
	{
		// Test conversion to DegreesFahrenheit.
		// No expected conversion value provided for DegreesFahrenheit, verifying result is not NaN.
		result := a.DegreesFahrenheit()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesFahrenheit returned NaN")
		}
	}
	{
		// Test conversion to DegreesNewton.
		// No expected conversion value provided for DegreesNewton, verifying result is not NaN.
		result := a.DegreesNewton()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesNewton returned NaN")
		}
	}
	{
		// Test conversion to DegreesRankine.
		// No expected conversion value provided for DegreesRankine, verifying result is not NaN.
		result := a.DegreesRankine()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesRankine returned NaN")
		}
	}
	{
		// Test conversion to DegreesReaumur.
		// No expected conversion value provided for DegreesReaumur, verifying result is not NaN.
		result := a.DegreesReaumur()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesReaumur returned NaN")
		}
	}
	{
		// Test conversion to DegreesRoemer.
		// No expected conversion value provided for DegreesRoemer, verifying result is not NaN.
		result := a.DegreesRoemer()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesRoemer returned NaN")
		}
	}
	{
		// Test conversion to SolarTemperatures.
		// No expected conversion value provided for SolarTemperatures, verifying result is not NaN.
		result := a.SolarTemperatures()
		if math.IsNaN(result) {
			t.Errorf("conversion to SolarTemperatures returned NaN")
		}
	}
}

func TestTemperature_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.TemperatureFactory{}
	a, err := factory.CreateTemperature(90, units.TemperatureKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.TemperatureKelvin {
		t.Errorf("expected default unit Kelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.TemperatureKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.TemperatureDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.TemperatureKelvin {
		t.Errorf("expected unit Kelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestTemperatureToString(t *testing.T) {
	factory := units.TemperatureFactory{}
	a, err := factory.CreateTemperature(45, units.TemperatureKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.TemperatureKelvin, 2)
	expected := "45.00 " + units.GetTemperatureAbbreviation(units.TemperatureKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.TemperatureKelvin, -1)
	expected = "45 " + units.GetTemperatureAbbreviation(units.TemperatureKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestTemperature_EqualityAndComparison(t *testing.T) {
	factory := units.TemperatureFactory{}
	a1, _ := factory.CreateTemperature(60, units.TemperatureKelvin)
	a2, _ := factory.CreateTemperature(60, units.TemperatureKelvin)
	a3, _ := factory.CreateTemperature(120, units.TemperatureKelvin)

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

func TestTemperature_Arithmetic(t *testing.T) {
	factory := units.TemperatureFactory{}
	a1, _ := factory.CreateTemperature(30, units.TemperatureKelvin)
	a2, _ := factory.CreateTemperature(45, units.TemperatureKelvin)

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