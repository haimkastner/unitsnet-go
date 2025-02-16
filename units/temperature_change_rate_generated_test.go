// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestTemperatureChangeRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "DegreeCelsiusPerSecond"}`
	
	factory := units.TemperatureChangeRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.TemperatureChangeRateDegreeCelsiusPerSecond {
		t.Errorf("expected unit %v, got %v", units.TemperatureChangeRateDegreeCelsiusPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "DegreeCelsiusPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestTemperatureChangeRateDto_ToJSON(t *testing.T) {
	dto := units.TemperatureChangeRateDto{
		Value: 45,
		Unit:  units.TemperatureChangeRateDegreeCelsiusPerSecond,
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
	if result["unit"].(string) != string(units.TemperatureChangeRateDegreeCelsiusPerSecond) {
		t.Errorf("expected unit %s, got %v", units.TemperatureChangeRateDegreeCelsiusPerSecond, result["unit"])
	}
}

func TestNewTemperatureChangeRate_InvalidValue(t *testing.T) {
	factory := units.TemperatureChangeRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreateTemperatureChangeRate(math.NaN(), units.TemperatureChangeRateDegreeCelsiusPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateTemperatureChangeRate(math.Inf(1), units.TemperatureChangeRateDegreeCelsiusPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestTemperatureChangeRateConversions(t *testing.T) {
	factory := units.TemperatureChangeRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateTemperatureChangeRate(180, units.TemperatureChangeRateDegreeCelsiusPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to DegreesCelsiusPerSecond.
		// No expected conversion value provided for DegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.DegreesCelsiusPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DegreesCelsiusPerMinute.
		// No expected conversion value provided for DegreesCelsiusPerMinute, verifying result is not NaN.
		result := a.DegreesCelsiusPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesCelsiusPerMinute returned NaN")
		}
	}
	{
		// Test conversion to DegreesKelvinPerMinute.
		// No expected conversion value provided for DegreesKelvinPerMinute, verifying result is not NaN.
		result := a.DegreesKelvinPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesKelvinPerMinute returned NaN")
		}
	}
	{
		// Test conversion to DegreesFahrenheitPerMinute.
		// No expected conversion value provided for DegreesFahrenheitPerMinute, verifying result is not NaN.
		result := a.DegreesFahrenheitPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesFahrenheitPerMinute returned NaN")
		}
	}
	{
		// Test conversion to DegreesFahrenheitPerSecond.
		// No expected conversion value provided for DegreesFahrenheitPerSecond, verifying result is not NaN.
		result := a.DegreesFahrenheitPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesFahrenheitPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DegreesKelvinPerSecond.
		// No expected conversion value provided for DegreesKelvinPerSecond, verifying result is not NaN.
		result := a.DegreesKelvinPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesKelvinPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DegreesCelsiusPerHour.
		// No expected conversion value provided for DegreesCelsiusPerHour, verifying result is not NaN.
		result := a.DegreesCelsiusPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesCelsiusPerHour returned NaN")
		}
	}
	{
		// Test conversion to DegreesKelvinPerHour.
		// No expected conversion value provided for DegreesKelvinPerHour, verifying result is not NaN.
		result := a.DegreesKelvinPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesKelvinPerHour returned NaN")
		}
	}
	{
		// Test conversion to DegreesFahrenheitPerHour.
		// No expected conversion value provided for DegreesFahrenheitPerHour, verifying result is not NaN.
		result := a.DegreesFahrenheitPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to DegreesFahrenheitPerHour returned NaN")
		}
	}
	{
		// Test conversion to NanodegreesCelsiusPerSecond.
		// No expected conversion value provided for NanodegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.NanodegreesCelsiusPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanodegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicrodegreesCelsiusPerSecond.
		// No expected conversion value provided for MicrodegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.MicrodegreesCelsiusPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrodegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillidegreesCelsiusPerSecond.
		// No expected conversion value provided for MillidegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.MillidegreesCelsiusPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillidegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CentidegreesCelsiusPerSecond.
		// No expected conversion value provided for CentidegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.CentidegreesCelsiusPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentidegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecidegreesCelsiusPerSecond.
		// No expected conversion value provided for DecidegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.DecidegreesCelsiusPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecidegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecadegreesCelsiusPerSecond.
		// No expected conversion value provided for DecadegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.DecadegreesCelsiusPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecadegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to HectodegreesCelsiusPerSecond.
		// No expected conversion value provided for HectodegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.HectodegreesCelsiusPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to HectodegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilodegreesCelsiusPerSecond.
		// No expected conversion value provided for KilodegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.KilodegreesCelsiusPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilodegreesCelsiusPerSecond returned NaN")
		}
	}
}

func TestTemperatureChangeRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.TemperatureChangeRateFactory{}
	a, err := factory.CreateTemperatureChangeRate(90, units.TemperatureChangeRateDegreeCelsiusPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.TemperatureChangeRateDegreeCelsiusPerSecond {
		t.Errorf("expected default unit DegreeCelsiusPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.TemperatureChangeRateDegreeCelsiusPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.TemperatureChangeRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.TemperatureChangeRateDegreeCelsiusPerSecond {
		t.Errorf("expected unit DegreeCelsiusPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestTemperatureChangeRateToString(t *testing.T) {
	factory := units.TemperatureChangeRateFactory{}
	a, err := factory.CreateTemperatureChangeRate(45, units.TemperatureChangeRateDegreeCelsiusPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.TemperatureChangeRateDegreeCelsiusPerSecond, 2)
	expected := "45.00 " + units.GetTemperatureChangeRateAbbreviation(units.TemperatureChangeRateDegreeCelsiusPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.TemperatureChangeRateDegreeCelsiusPerSecond, -1)
	expected = "45 " + units.GetTemperatureChangeRateAbbreviation(units.TemperatureChangeRateDegreeCelsiusPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestTemperatureChangeRate_EqualityAndComparison(t *testing.T) {
	factory := units.TemperatureChangeRateFactory{}
	a1, _ := factory.CreateTemperatureChangeRate(60, units.TemperatureChangeRateDegreeCelsiusPerSecond)
	a2, _ := factory.CreateTemperatureChangeRate(60, units.TemperatureChangeRateDegreeCelsiusPerSecond)
	a3, _ := factory.CreateTemperatureChangeRate(120, units.TemperatureChangeRateDegreeCelsiusPerSecond)

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

func TestTemperatureChangeRate_Arithmetic(t *testing.T) {
	factory := units.TemperatureChangeRateFactory{}
	a1, _ := factory.CreateTemperatureChangeRate(30, units.TemperatureChangeRateDegreeCelsiusPerSecond)
	a2, _ := factory.CreateTemperatureChangeRate(45, units.TemperatureChangeRateDegreeCelsiusPerSecond)

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