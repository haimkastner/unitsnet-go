// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestCoefficientOfThermalExpansionDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "PerKelvin"}`
	
	factory := units.CoefficientOfThermalExpansionDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.CoefficientOfThermalExpansionPerKelvin {
		t.Errorf("expected unit %v, got %v", units.CoefficientOfThermalExpansionPerKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "PerKelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestCoefficientOfThermalExpansionDto_ToJSON(t *testing.T) {
	dto := units.CoefficientOfThermalExpansionDto{
		Value: 45,
		Unit:  units.CoefficientOfThermalExpansionPerKelvin,
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
	if result["unit"].(string) != string(units.CoefficientOfThermalExpansionPerKelvin) {
		t.Errorf("expected unit %s, got %v", units.CoefficientOfThermalExpansionPerKelvin, result["unit"])
	}
}

func TestNewCoefficientOfThermalExpansion_InvalidValue(t *testing.T) {
	factory := units.CoefficientOfThermalExpansionFactory{}
	// NaN value should return an error.
	_, err := factory.CreateCoefficientOfThermalExpansion(math.NaN(), units.CoefficientOfThermalExpansionPerKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateCoefficientOfThermalExpansion(math.Inf(1), units.CoefficientOfThermalExpansionPerKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestCoefficientOfThermalExpansionConversions(t *testing.T) {
	factory := units.CoefficientOfThermalExpansionFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateCoefficientOfThermalExpansion(180, units.CoefficientOfThermalExpansionPerKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to PerKelvin.
		// No expected conversion value provided for PerKelvin, verifying result is not NaN.
		result := a.PerKelvin()
		cacheResult := a.PerKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PerKelvin returned NaN")
		}
	}
	{
		// Test conversion to PerDegreeCelsius.
		// No expected conversion value provided for PerDegreeCelsius, verifying result is not NaN.
		result := a.PerDegreeCelsius()
		cacheResult := a.PerDegreeCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PerDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to PerDegreeFahrenheit.
		// No expected conversion value provided for PerDegreeFahrenheit, verifying result is not NaN.
		result := a.PerDegreeFahrenheit()
		cacheResult := a.PerDegreeFahrenheit()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PerDegreeFahrenheit returned NaN")
		}
	}
	{
		// Test conversion to PpmPerKelvin.
		// No expected conversion value provided for PpmPerKelvin, verifying result is not NaN.
		result := a.PpmPerKelvin()
		cacheResult := a.PpmPerKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PpmPerKelvin returned NaN")
		}
	}
	{
		// Test conversion to PpmPerDegreeCelsius.
		// No expected conversion value provided for PpmPerDegreeCelsius, verifying result is not NaN.
		result := a.PpmPerDegreeCelsius()
		cacheResult := a.PpmPerDegreeCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PpmPerDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to PpmPerDegreeFahrenheit.
		// No expected conversion value provided for PpmPerDegreeFahrenheit, verifying result is not NaN.
		result := a.PpmPerDegreeFahrenheit()
		cacheResult := a.PpmPerDegreeFahrenheit()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PpmPerDegreeFahrenheit returned NaN")
		}
	}
}

func TestCoefficientOfThermalExpansion_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.CoefficientOfThermalExpansionFactory{}
	a, err := factory.CreateCoefficientOfThermalExpansion(90, units.CoefficientOfThermalExpansionPerKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.CoefficientOfThermalExpansionPerKelvin {
		t.Errorf("expected default unit PerKelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.CoefficientOfThermalExpansionPerKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.CoefficientOfThermalExpansionDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.CoefficientOfThermalExpansionPerKelvin {
		t.Errorf("expected unit PerKelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestCoefficientOfThermalExpansionFactory_FromDto(t *testing.T) {
    factory := units.CoefficientOfThermalExpansionFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.CoefficientOfThermalExpansionDto{
        Value: 100,
        Unit:  units.CoefficientOfThermalExpansionPerKelvin,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.CoefficientOfThermalExpansionDto{
        Value: math.NaN(),
        Unit:  units.CoefficientOfThermalExpansionPerKelvin,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test PerKelvin conversion
    per_kelvinDto := units.CoefficientOfThermalExpansionDto{
        Value: 100,
        Unit:  units.CoefficientOfThermalExpansionPerKelvin,
    }
    
    var per_kelvinResult *units.CoefficientOfThermalExpansion
    per_kelvinResult, err = factory.FromDto(per_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with PerKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = per_kelvinResult.Convert(units.CoefficientOfThermalExpansionPerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PerKelvin = %v, want %v", converted, 100)
    }
    // Test PerDegreeCelsius conversion
    per_degree_celsiusDto := units.CoefficientOfThermalExpansionDto{
        Value: 100,
        Unit:  units.CoefficientOfThermalExpansionPerDegreeCelsius,
    }
    
    var per_degree_celsiusResult *units.CoefficientOfThermalExpansion
    per_degree_celsiusResult, err = factory.FromDto(per_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with PerDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = per_degree_celsiusResult.Convert(units.CoefficientOfThermalExpansionPerDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PerDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test PerDegreeFahrenheit conversion
    per_degree_fahrenheitDto := units.CoefficientOfThermalExpansionDto{
        Value: 100,
        Unit:  units.CoefficientOfThermalExpansionPerDegreeFahrenheit,
    }
    
    var per_degree_fahrenheitResult *units.CoefficientOfThermalExpansion
    per_degree_fahrenheitResult, err = factory.FromDto(per_degree_fahrenheitDto)
    if err != nil {
        t.Errorf("FromDto() with PerDegreeFahrenheit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = per_degree_fahrenheitResult.Convert(units.CoefficientOfThermalExpansionPerDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PerDegreeFahrenheit = %v, want %v", converted, 100)
    }
    // Test PpmPerKelvin conversion
    ppm_per_kelvinDto := units.CoefficientOfThermalExpansionDto{
        Value: 100,
        Unit:  units.CoefficientOfThermalExpansionPpmPerKelvin,
    }
    
    var ppm_per_kelvinResult *units.CoefficientOfThermalExpansion
    ppm_per_kelvinResult, err = factory.FromDto(ppm_per_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with PpmPerKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ppm_per_kelvinResult.Convert(units.CoefficientOfThermalExpansionPpmPerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PpmPerKelvin = %v, want %v", converted, 100)
    }
    // Test PpmPerDegreeCelsius conversion
    ppm_per_degree_celsiusDto := units.CoefficientOfThermalExpansionDto{
        Value: 100,
        Unit:  units.CoefficientOfThermalExpansionPpmPerDegreeCelsius,
    }
    
    var ppm_per_degree_celsiusResult *units.CoefficientOfThermalExpansion
    ppm_per_degree_celsiusResult, err = factory.FromDto(ppm_per_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with PpmPerDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ppm_per_degree_celsiusResult.Convert(units.CoefficientOfThermalExpansionPpmPerDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PpmPerDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test PpmPerDegreeFahrenheit conversion
    ppm_per_degree_fahrenheitDto := units.CoefficientOfThermalExpansionDto{
        Value: 100,
        Unit:  units.CoefficientOfThermalExpansionPpmPerDegreeFahrenheit,
    }
    
    var ppm_per_degree_fahrenheitResult *units.CoefficientOfThermalExpansion
    ppm_per_degree_fahrenheitResult, err = factory.FromDto(ppm_per_degree_fahrenheitDto)
    if err != nil {
        t.Errorf("FromDto() with PpmPerDegreeFahrenheit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ppm_per_degree_fahrenheitResult.Convert(units.CoefficientOfThermalExpansionPpmPerDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PpmPerDegreeFahrenheit = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.CoefficientOfThermalExpansionDto{
        Value: 0,
        Unit:  units.CoefficientOfThermalExpansionPerKelvin,
    }
    
    var zeroResult *units.CoefficientOfThermalExpansion
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestCoefficientOfThermalExpansionFactory_FromDtoJSON(t *testing.T) {
    factory := units.CoefficientOfThermalExpansionFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "PerKelvin"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "PerKelvin"}`)
    _, err = factory.FromDtoJSON(invalidJSON)
    if err == nil {
        t.Error("FromDtoJSON() with invalid JSON should return error")
    }

    // Test malformed JSON
    malformedJSON := []byte(`{malformed json`)
    _, err = factory.FromDtoJSON(malformedJSON)
    if err == nil {
        t.Error("FromDtoJSON() with malformed JSON should return error")
    }

    // Test empty JSON
    emptyJSON := []byte(`{}`)
    _, err = factory.FromDtoJSON(emptyJSON)
    if err == nil {
        t.Error("FromDtoJSON() with empty JSON should return error")
    }

    // Test JSON with invalid value (NaN)
    nanValue := math.NaN()
    nanJSON, _ := json.Marshal(units.CoefficientOfThermalExpansionDto{
        Value: nanValue,
        Unit:  units.CoefficientOfThermalExpansionPerKelvin,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with PerKelvin unit
    per_kelvinJSON := []byte(`{"value": 100, "unit": "PerKelvin"}`)
    per_kelvinResult, err := factory.FromDtoJSON(per_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PerKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = per_kelvinResult.Convert(units.CoefficientOfThermalExpansionPerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PerKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with PerDegreeCelsius unit
    per_degree_celsiusJSON := []byte(`{"value": 100, "unit": "PerDegreeCelsius"}`)
    per_degree_celsiusResult, err := factory.FromDtoJSON(per_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PerDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = per_degree_celsiusResult.Convert(units.CoefficientOfThermalExpansionPerDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PerDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with PerDegreeFahrenheit unit
    per_degree_fahrenheitJSON := []byte(`{"value": 100, "unit": "PerDegreeFahrenheit"}`)
    per_degree_fahrenheitResult, err := factory.FromDtoJSON(per_degree_fahrenheitJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PerDegreeFahrenheit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = per_degree_fahrenheitResult.Convert(units.CoefficientOfThermalExpansionPerDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PerDegreeFahrenheit = %v, want %v", converted, 100)
    }
    // Test JSON with PpmPerKelvin unit
    ppm_per_kelvinJSON := []byte(`{"value": 100, "unit": "PpmPerKelvin"}`)
    ppm_per_kelvinResult, err := factory.FromDtoJSON(ppm_per_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PpmPerKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ppm_per_kelvinResult.Convert(units.CoefficientOfThermalExpansionPpmPerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PpmPerKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with PpmPerDegreeCelsius unit
    ppm_per_degree_celsiusJSON := []byte(`{"value": 100, "unit": "PpmPerDegreeCelsius"}`)
    ppm_per_degree_celsiusResult, err := factory.FromDtoJSON(ppm_per_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PpmPerDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ppm_per_degree_celsiusResult.Convert(units.CoefficientOfThermalExpansionPpmPerDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PpmPerDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with PpmPerDegreeFahrenheit unit
    ppm_per_degree_fahrenheitJSON := []byte(`{"value": 100, "unit": "PpmPerDegreeFahrenheit"}`)
    ppm_per_degree_fahrenheitResult, err := factory.FromDtoJSON(ppm_per_degree_fahrenheitJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PpmPerDegreeFahrenheit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ppm_per_degree_fahrenheitResult.Convert(units.CoefficientOfThermalExpansionPpmPerDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PpmPerDegreeFahrenheit = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "PerKelvin"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromPerKelvin function
func TestCoefficientOfThermalExpansionFactory_FromPerKelvin(t *testing.T) {
    factory := units.CoefficientOfThermalExpansionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPerKelvin(100)
    if err != nil {
        t.Errorf("FromPerKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CoefficientOfThermalExpansionPerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPerKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPerKelvin(math.NaN())
    if err == nil {
        t.Error("FromPerKelvin() with NaN value should return error")
    }

    _, err = factory.FromPerKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromPerKelvin() with +Inf value should return error")
    }

    _, err = factory.FromPerKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromPerKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPerKelvin(0)
    if err != nil {
        t.Errorf("FromPerKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CoefficientOfThermalExpansionPerKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPerKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromPerDegreeCelsius function
func TestCoefficientOfThermalExpansionFactory_FromPerDegreeCelsius(t *testing.T) {
    factory := units.CoefficientOfThermalExpansionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPerDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromPerDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CoefficientOfThermalExpansionPerDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPerDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPerDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromPerDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromPerDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromPerDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromPerDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromPerDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPerDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromPerDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CoefficientOfThermalExpansionPerDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPerDegreeCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromPerDegreeFahrenheit function
func TestCoefficientOfThermalExpansionFactory_FromPerDegreeFahrenheit(t *testing.T) {
    factory := units.CoefficientOfThermalExpansionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPerDegreeFahrenheit(100)
    if err != nil {
        t.Errorf("FromPerDegreeFahrenheit() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CoefficientOfThermalExpansionPerDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPerDegreeFahrenheit() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPerDegreeFahrenheit(math.NaN())
    if err == nil {
        t.Error("FromPerDegreeFahrenheit() with NaN value should return error")
    }

    _, err = factory.FromPerDegreeFahrenheit(math.Inf(1))
    if err == nil {
        t.Error("FromPerDegreeFahrenheit() with +Inf value should return error")
    }

    _, err = factory.FromPerDegreeFahrenheit(math.Inf(-1))
    if err == nil {
        t.Error("FromPerDegreeFahrenheit() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPerDegreeFahrenheit(0)
    if err != nil {
        t.Errorf("FromPerDegreeFahrenheit() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CoefficientOfThermalExpansionPerDegreeFahrenheit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPerDegreeFahrenheit() with zero value = %v, want 0", converted)
    }
}
// Test FromPpmPerKelvin function
func TestCoefficientOfThermalExpansionFactory_FromPpmPerKelvin(t *testing.T) {
    factory := units.CoefficientOfThermalExpansionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPpmPerKelvin(100)
    if err != nil {
        t.Errorf("FromPpmPerKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CoefficientOfThermalExpansionPpmPerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPpmPerKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPpmPerKelvin(math.NaN())
    if err == nil {
        t.Error("FromPpmPerKelvin() with NaN value should return error")
    }

    _, err = factory.FromPpmPerKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromPpmPerKelvin() with +Inf value should return error")
    }

    _, err = factory.FromPpmPerKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromPpmPerKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPpmPerKelvin(0)
    if err != nil {
        t.Errorf("FromPpmPerKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CoefficientOfThermalExpansionPpmPerKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPpmPerKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromPpmPerDegreeCelsius function
func TestCoefficientOfThermalExpansionFactory_FromPpmPerDegreeCelsius(t *testing.T) {
    factory := units.CoefficientOfThermalExpansionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPpmPerDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromPpmPerDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CoefficientOfThermalExpansionPpmPerDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPpmPerDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPpmPerDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromPpmPerDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromPpmPerDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromPpmPerDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromPpmPerDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromPpmPerDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPpmPerDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromPpmPerDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CoefficientOfThermalExpansionPpmPerDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPpmPerDegreeCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromPpmPerDegreeFahrenheit function
func TestCoefficientOfThermalExpansionFactory_FromPpmPerDegreeFahrenheit(t *testing.T) {
    factory := units.CoefficientOfThermalExpansionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPpmPerDegreeFahrenheit(100)
    if err != nil {
        t.Errorf("FromPpmPerDegreeFahrenheit() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CoefficientOfThermalExpansionPpmPerDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPpmPerDegreeFahrenheit() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPpmPerDegreeFahrenheit(math.NaN())
    if err == nil {
        t.Error("FromPpmPerDegreeFahrenheit() with NaN value should return error")
    }

    _, err = factory.FromPpmPerDegreeFahrenheit(math.Inf(1))
    if err == nil {
        t.Error("FromPpmPerDegreeFahrenheit() with +Inf value should return error")
    }

    _, err = factory.FromPpmPerDegreeFahrenheit(math.Inf(-1))
    if err == nil {
        t.Error("FromPpmPerDegreeFahrenheit() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPpmPerDegreeFahrenheit(0)
    if err != nil {
        t.Errorf("FromPpmPerDegreeFahrenheit() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CoefficientOfThermalExpansionPpmPerDegreeFahrenheit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPpmPerDegreeFahrenheit() with zero value = %v, want 0", converted)
    }
}

func TestCoefficientOfThermalExpansionToString(t *testing.T) {
	factory := units.CoefficientOfThermalExpansionFactory{}
	a, err := factory.CreateCoefficientOfThermalExpansion(45, units.CoefficientOfThermalExpansionPerKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.CoefficientOfThermalExpansionPerKelvin, 2)
	expected := "45.00 " + units.GetCoefficientOfThermalExpansionAbbreviation(units.CoefficientOfThermalExpansionPerKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.CoefficientOfThermalExpansionPerKelvin, -1)
	expected = "45 " + units.GetCoefficientOfThermalExpansionAbbreviation(units.CoefficientOfThermalExpansionPerKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestCoefficientOfThermalExpansion_EqualityAndComparison(t *testing.T) {
	factory := units.CoefficientOfThermalExpansionFactory{}
	a1, _ := factory.CreateCoefficientOfThermalExpansion(60, units.CoefficientOfThermalExpansionPerKelvin)
	a2, _ := factory.CreateCoefficientOfThermalExpansion(60, units.CoefficientOfThermalExpansionPerKelvin)
	a3, _ := factory.CreateCoefficientOfThermalExpansion(120, units.CoefficientOfThermalExpansionPerKelvin)

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

func TestCoefficientOfThermalExpansion_Arithmetic(t *testing.T) {
	factory := units.CoefficientOfThermalExpansionFactory{}
	a1, _ := factory.CreateCoefficientOfThermalExpansion(30, units.CoefficientOfThermalExpansionPerKelvin)
	a2, _ := factory.CreateCoefficientOfThermalExpansion(45, units.CoefficientOfThermalExpansionPerKelvin)

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


func TestGetCoefficientOfThermalExpansionAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.CoefficientOfThermalExpansionUnits
        want string
    }{
        {
            name: "PerKelvin abbreviation",
            unit: units.CoefficientOfThermalExpansionPerKelvin,
            want: "K⁻¹",
        },
        {
            name: "PerDegreeCelsius abbreviation",
            unit: units.CoefficientOfThermalExpansionPerDegreeCelsius,
            want: "°C⁻¹",
        },
        {
            name: "PerDegreeFahrenheit abbreviation",
            unit: units.CoefficientOfThermalExpansionPerDegreeFahrenheit,
            want: "°F⁻¹",
        },
        {
            name: "PpmPerKelvin abbreviation",
            unit: units.CoefficientOfThermalExpansionPpmPerKelvin,
            want: "ppm/K",
        },
        {
            name: "PpmPerDegreeCelsius abbreviation",
            unit: units.CoefficientOfThermalExpansionPpmPerDegreeCelsius,
            want: "ppm/°C",
        },
        {
            name: "PpmPerDegreeFahrenheit abbreviation",
            unit: units.CoefficientOfThermalExpansionPpmPerDegreeFahrenheit,
            want: "ppm/°F",
        },
        {
            name: "invalid unit",
            unit: units.CoefficientOfThermalExpansionUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetCoefficientOfThermalExpansionAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetCoefficientOfThermalExpansionAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestCoefficientOfThermalExpansion_String(t *testing.T) {
    factory := units.CoefficientOfThermalExpansionFactory{}
    
    tests := []struct {
        name  string
        value float64
        want  string
    }{
        {
            name:  "positive integer",
            value: 100,
            want:  "100.00",
        },
        {
            name:  "negative integer",
            value: -100,
            want:  "-100.00",
        },
        {
            name:  "zero",
            value: 0,
            want:  "0.00",
        },
        {
            name:  "positive decimal",
            value: 123.456,
            want:  "123.46",
        },
        {
            name:  "negative decimal",
            value: -123.456,
            want:  "-123.46",
        },
        {
            name:  "small decimal",
            value: 0.123,
            want:  "0.12",
        },
        {
            name:  "large number",
            value: 1000000,
            want:  "1000000.00",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            unit, err := factory.CreateCoefficientOfThermalExpansion(tt.value, units.CoefficientOfThermalExpansionPerKelvin)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("CoefficientOfThermalExpansion.String() = %v, want %v", got, tt.want)
            }
        })
    }
}