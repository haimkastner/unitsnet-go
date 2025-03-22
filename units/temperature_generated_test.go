// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

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
		cacheResult := a.Kelvins()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kelvins returned NaN")
		}
	}
	{
		// Test conversion to DegreesCelsius.
		// No expected conversion value provided for DegreesCelsius, verifying result is not NaN.
		result := a.DegreesCelsius()
		cacheResult := a.DegreesCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesCelsius returned NaN")
		}
	}
	{
		// Test conversion to MillidegreesCelsius.
		// No expected conversion value provided for MillidegreesCelsius, verifying result is not NaN.
		result := a.MillidegreesCelsius()
		cacheResult := a.MillidegreesCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillidegreesCelsius returned NaN")
		}
	}
	{
		// Test conversion to DegreesDelisle.
		// No expected conversion value provided for DegreesDelisle, verifying result is not NaN.
		result := a.DegreesDelisle()
		cacheResult := a.DegreesDelisle()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesDelisle returned NaN")
		}
	}
	{
		// Test conversion to DegreesFahrenheit.
		// No expected conversion value provided for DegreesFahrenheit, verifying result is not NaN.
		result := a.DegreesFahrenheit()
		cacheResult := a.DegreesFahrenheit()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesFahrenheit returned NaN")
		}
	}
	{
		// Test conversion to DegreesNewton.
		// No expected conversion value provided for DegreesNewton, verifying result is not NaN.
		result := a.DegreesNewton()
		cacheResult := a.DegreesNewton()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesNewton returned NaN")
		}
	}
	{
		// Test conversion to DegreesRankine.
		// No expected conversion value provided for DegreesRankine, verifying result is not NaN.
		result := a.DegreesRankine()
		cacheResult := a.DegreesRankine()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesRankine returned NaN")
		}
	}
	{
		// Test conversion to DegreesReaumur.
		// No expected conversion value provided for DegreesReaumur, verifying result is not NaN.
		result := a.DegreesReaumur()
		cacheResult := a.DegreesReaumur()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesReaumur returned NaN")
		}
	}
	{
		// Test conversion to DegreesRoemer.
		// No expected conversion value provided for DegreesRoemer, verifying result is not NaN.
		result := a.DegreesRoemer()
		cacheResult := a.DegreesRoemer()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesRoemer returned NaN")
		}
	}
	{
		// Test conversion to SolarTemperatures.
		// No expected conversion value provided for SolarTemperatures, verifying result is not NaN.
		result := a.SolarTemperatures()
		cacheResult := a.SolarTemperatures()
		if math.IsNaN(result) || cacheResult != result {
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

func TestTemperatureFactory_FromDto(t *testing.T) {
    factory := units.TemperatureFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.TemperatureDto{
        Value: 100,
        Unit:  units.TemperatureKelvin,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.TemperatureDto{
        Value: math.NaN(),
        Unit:  units.TemperatureKelvin,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Kelvin conversion
    kelvinsDto := units.TemperatureDto{
        Value: 100,
        Unit:  units.TemperatureKelvin,
    }
    
    var kelvinsResult *units.Temperature
    kelvinsResult, err = factory.FromDto(kelvinsDto)
    if err != nil {
        t.Errorf("FromDto() with Kelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kelvinsResult.Convert(units.TemperatureKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kelvin = %v, want %v", converted, 100)
    }
    // Test DegreeCelsius conversion
    degrees_celsiusDto := units.TemperatureDto{
        Value: 100,
        Unit:  units.TemperatureDegreeCelsius,
    }
    
    var degrees_celsiusResult *units.Temperature
    degrees_celsiusResult, err = factory.FromDto(degrees_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsiusResult.Convert(units.TemperatureDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsius = %v, want %v", converted, 100)
    }
    // Test MillidegreeCelsius conversion
    millidegrees_celsiusDto := units.TemperatureDto{
        Value: 100,
        Unit:  units.TemperatureMillidegreeCelsius,
    }
    
    var millidegrees_celsiusResult *units.Temperature
    millidegrees_celsiusResult, err = factory.FromDto(millidegrees_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with MillidegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millidegrees_celsiusResult.Convert(units.TemperatureMillidegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillidegreeCelsius = %v, want %v", converted, 100)
    }
    // Test DegreeDelisle conversion
    degrees_delisleDto := units.TemperatureDto{
        Value: 100,
        Unit:  units.TemperatureDegreeDelisle,
    }
    
    var degrees_delisleResult *units.Temperature
    degrees_delisleResult, err = factory.FromDto(degrees_delisleDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeDelisle returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_delisleResult.Convert(units.TemperatureDegreeDelisle)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeDelisle = %v, want %v", converted, 100)
    }
    // Test DegreeFahrenheit conversion
    degrees_fahrenheitDto := units.TemperatureDto{
        Value: 100,
        Unit:  units.TemperatureDegreeFahrenheit,
    }
    
    var degrees_fahrenheitResult *units.Temperature
    degrees_fahrenheitResult, err = factory.FromDto(degrees_fahrenheitDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeFahrenheit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_fahrenheitResult.Convert(units.TemperatureDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeFahrenheit = %v, want %v", converted, 100)
    }
    // Test DegreeNewton conversion
    degrees_newtonDto := units.TemperatureDto{
        Value: 100,
        Unit:  units.TemperatureDegreeNewton,
    }
    
    var degrees_newtonResult *units.Temperature
    degrees_newtonResult, err = factory.FromDto(degrees_newtonDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeNewton returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_newtonResult.Convert(units.TemperatureDegreeNewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeNewton = %v, want %v", converted, 100)
    }
    // Test DegreeRankine conversion
    degrees_rankineDto := units.TemperatureDto{
        Value: 100,
        Unit:  units.TemperatureDegreeRankine,
    }
    
    var degrees_rankineResult *units.Temperature
    degrees_rankineResult, err = factory.FromDto(degrees_rankineDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeRankine returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_rankineResult.Convert(units.TemperatureDegreeRankine)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeRankine = %v, want %v", converted, 100)
    }
    // Test DegreeReaumur conversion
    degrees_reaumurDto := units.TemperatureDto{
        Value: 100,
        Unit:  units.TemperatureDegreeReaumur,
    }
    
    var degrees_reaumurResult *units.Temperature
    degrees_reaumurResult, err = factory.FromDto(degrees_reaumurDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeReaumur returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_reaumurResult.Convert(units.TemperatureDegreeReaumur)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeReaumur = %v, want %v", converted, 100)
    }
    // Test DegreeRoemer conversion
    degrees_roemerDto := units.TemperatureDto{
        Value: 100,
        Unit:  units.TemperatureDegreeRoemer,
    }
    
    var degrees_roemerResult *units.Temperature
    degrees_roemerResult, err = factory.FromDto(degrees_roemerDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeRoemer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_roemerResult.Convert(units.TemperatureDegreeRoemer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeRoemer = %v, want %v", converted, 100)
    }
    // Test SolarTemperature conversion
    solar_temperaturesDto := units.TemperatureDto{
        Value: 100,
        Unit:  units.TemperatureSolarTemperature,
    }
    
    var solar_temperaturesResult *units.Temperature
    solar_temperaturesResult, err = factory.FromDto(solar_temperaturesDto)
    if err != nil {
        t.Errorf("FromDto() with SolarTemperature returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = solar_temperaturesResult.Convert(units.TemperatureSolarTemperature)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SolarTemperature = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.TemperatureDto{
        Value: 0,
        Unit:  units.TemperatureKelvin,
    }
    
    var zeroResult *units.Temperature
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestTemperatureFactory_FromDtoJSON(t *testing.T) {
    factory := units.TemperatureFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Kelvin"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Kelvin"}`)
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
    nanJSON, _ := json.Marshal(units.TemperatureDto{
        Value: nanValue,
        Unit:  units.TemperatureKelvin,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Kelvin unit
    kelvinsJSON := []byte(`{"value": 100, "unit": "Kelvin"}`)
    kelvinsResult, err := factory.FromDtoJSON(kelvinsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kelvinsResult.Convert(units.TemperatureKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kelvin = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeCelsius unit
    degrees_celsiusJSON := []byte(`{"value": 100, "unit": "DegreeCelsius"}`)
    degrees_celsiusResult, err := factory.FromDtoJSON(degrees_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsiusResult.Convert(units.TemperatureDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with MillidegreeCelsius unit
    millidegrees_celsiusJSON := []byte(`{"value": 100, "unit": "MillidegreeCelsius"}`)
    millidegrees_celsiusResult, err := factory.FromDtoJSON(millidegrees_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillidegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millidegrees_celsiusResult.Convert(units.TemperatureMillidegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillidegreeCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeDelisle unit
    degrees_delisleJSON := []byte(`{"value": 100, "unit": "DegreeDelisle"}`)
    degrees_delisleResult, err := factory.FromDtoJSON(degrees_delisleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeDelisle unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_delisleResult.Convert(units.TemperatureDegreeDelisle)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeDelisle = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeFahrenheit unit
    degrees_fahrenheitJSON := []byte(`{"value": 100, "unit": "DegreeFahrenheit"}`)
    degrees_fahrenheitResult, err := factory.FromDtoJSON(degrees_fahrenheitJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeFahrenheit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_fahrenheitResult.Convert(units.TemperatureDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeFahrenheit = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeNewton unit
    degrees_newtonJSON := []byte(`{"value": 100, "unit": "DegreeNewton"}`)
    degrees_newtonResult, err := factory.FromDtoJSON(degrees_newtonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeNewton unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_newtonResult.Convert(units.TemperatureDegreeNewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeNewton = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeRankine unit
    degrees_rankineJSON := []byte(`{"value": 100, "unit": "DegreeRankine"}`)
    degrees_rankineResult, err := factory.FromDtoJSON(degrees_rankineJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeRankine unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_rankineResult.Convert(units.TemperatureDegreeRankine)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeRankine = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeReaumur unit
    degrees_reaumurJSON := []byte(`{"value": 100, "unit": "DegreeReaumur"}`)
    degrees_reaumurResult, err := factory.FromDtoJSON(degrees_reaumurJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeReaumur unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_reaumurResult.Convert(units.TemperatureDegreeReaumur)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeReaumur = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeRoemer unit
    degrees_roemerJSON := []byte(`{"value": 100, "unit": "DegreeRoemer"}`)
    degrees_roemerResult, err := factory.FromDtoJSON(degrees_roemerJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeRoemer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_roemerResult.Convert(units.TemperatureDegreeRoemer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeRoemer = %v, want %v", converted, 100)
    }
    // Test JSON with SolarTemperature unit
    solar_temperaturesJSON := []byte(`{"value": 100, "unit": "SolarTemperature"}`)
    solar_temperaturesResult, err := factory.FromDtoJSON(solar_temperaturesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SolarTemperature unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = solar_temperaturesResult.Convert(units.TemperatureSolarTemperature)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SolarTemperature = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Kelvin"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromKelvins function
func TestTemperatureFactory_FromKelvins(t *testing.T) {
    factory := units.TemperatureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKelvins(100)
    if err != nil {
        t.Errorf("FromKelvins() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKelvins() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKelvins(math.NaN())
    if err == nil {
        t.Error("FromKelvins() with NaN value should return error")
    }

    _, err = factory.FromKelvins(math.Inf(1))
    if err == nil {
        t.Error("FromKelvins() with +Inf value should return error")
    }

    _, err = factory.FromKelvins(math.Inf(-1))
    if err == nil {
        t.Error("FromKelvins() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKelvins(0)
    if err != nil {
        t.Errorf("FromKelvins() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKelvins() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesCelsius function
func TestTemperatureFactory_FromDegreesCelsius(t *testing.T) {
    factory := units.TemperatureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesCelsius(100)
    if err != nil {
        t.Errorf("FromDegreesCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesCelsius(math.NaN())
    if err == nil {
        t.Error("FromDegreesCelsius() with NaN value should return error")
    }

    _, err = factory.FromDegreesCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesCelsius() with +Inf value should return error")
    }

    _, err = factory.FromDegreesCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesCelsius(0)
    if err != nil {
        t.Errorf("FromDegreesCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromMillidegreesCelsius function
func TestTemperatureFactory_FromMillidegreesCelsius(t *testing.T) {
    factory := units.TemperatureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillidegreesCelsius(100)
    if err != nil {
        t.Errorf("FromMillidegreesCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureMillidegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillidegreesCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillidegreesCelsius(math.NaN())
    if err == nil {
        t.Error("FromMillidegreesCelsius() with NaN value should return error")
    }

    _, err = factory.FromMillidegreesCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromMillidegreesCelsius() with +Inf value should return error")
    }

    _, err = factory.FromMillidegreesCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromMillidegreesCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillidegreesCelsius(0)
    if err != nil {
        t.Errorf("FromMillidegreesCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureMillidegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillidegreesCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesDelisle function
func TestTemperatureFactory_FromDegreesDelisle(t *testing.T) {
    factory := units.TemperatureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesDelisle(100)
    if err != nil {
        t.Errorf("FromDegreesDelisle() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureDegreeDelisle)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesDelisle() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesDelisle(math.NaN())
    if err == nil {
        t.Error("FromDegreesDelisle() with NaN value should return error")
    }

    _, err = factory.FromDegreesDelisle(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesDelisle() with +Inf value should return error")
    }

    _, err = factory.FromDegreesDelisle(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesDelisle() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesDelisle(0)
    if err != nil {
        t.Errorf("FromDegreesDelisle() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureDegreeDelisle)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesDelisle() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesFahrenheit function
func TestTemperatureFactory_FromDegreesFahrenheit(t *testing.T) {
    factory := units.TemperatureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesFahrenheit(100)
    if err != nil {
        t.Errorf("FromDegreesFahrenheit() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesFahrenheit() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesFahrenheit(math.NaN())
    if err == nil {
        t.Error("FromDegreesFahrenheit() with NaN value should return error")
    }

    _, err = factory.FromDegreesFahrenheit(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesFahrenheit() with +Inf value should return error")
    }

    _, err = factory.FromDegreesFahrenheit(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesFahrenheit() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesFahrenheit(0)
    if err != nil {
        t.Errorf("FromDegreesFahrenheit() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureDegreeFahrenheit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesFahrenheit() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesNewton function
func TestTemperatureFactory_FromDegreesNewton(t *testing.T) {
    factory := units.TemperatureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesNewton(100)
    if err != nil {
        t.Errorf("FromDegreesNewton() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureDegreeNewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesNewton() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesNewton(math.NaN())
    if err == nil {
        t.Error("FromDegreesNewton() with NaN value should return error")
    }

    _, err = factory.FromDegreesNewton(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesNewton() with +Inf value should return error")
    }

    _, err = factory.FromDegreesNewton(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesNewton() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesNewton(0)
    if err != nil {
        t.Errorf("FromDegreesNewton() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureDegreeNewton)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesNewton() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesRankine function
func TestTemperatureFactory_FromDegreesRankine(t *testing.T) {
    factory := units.TemperatureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesRankine(100)
    if err != nil {
        t.Errorf("FromDegreesRankine() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureDegreeRankine)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesRankine() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesRankine(math.NaN())
    if err == nil {
        t.Error("FromDegreesRankine() with NaN value should return error")
    }

    _, err = factory.FromDegreesRankine(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesRankine() with +Inf value should return error")
    }

    _, err = factory.FromDegreesRankine(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesRankine() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesRankine(0)
    if err != nil {
        t.Errorf("FromDegreesRankine() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureDegreeRankine)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesRankine() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesReaumur function
func TestTemperatureFactory_FromDegreesReaumur(t *testing.T) {
    factory := units.TemperatureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesReaumur(100)
    if err != nil {
        t.Errorf("FromDegreesReaumur() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureDegreeReaumur)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesReaumur() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesReaumur(math.NaN())
    if err == nil {
        t.Error("FromDegreesReaumur() with NaN value should return error")
    }

    _, err = factory.FromDegreesReaumur(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesReaumur() with +Inf value should return error")
    }

    _, err = factory.FromDegreesReaumur(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesReaumur() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesReaumur(0)
    if err != nil {
        t.Errorf("FromDegreesReaumur() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureDegreeReaumur)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesReaumur() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesRoemer function
func TestTemperatureFactory_FromDegreesRoemer(t *testing.T) {
    factory := units.TemperatureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesRoemer(100)
    if err != nil {
        t.Errorf("FromDegreesRoemer() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureDegreeRoemer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesRoemer() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesRoemer(math.NaN())
    if err == nil {
        t.Error("FromDegreesRoemer() with NaN value should return error")
    }

    _, err = factory.FromDegreesRoemer(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesRoemer() with +Inf value should return error")
    }

    _, err = factory.FromDegreesRoemer(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesRoemer() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesRoemer(0)
    if err != nil {
        t.Errorf("FromDegreesRoemer() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureDegreeRoemer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesRoemer() with zero value = %v, want 0", converted)
    }
}
// Test FromSolarTemperatures function
func TestTemperatureFactory_FromSolarTemperatures(t *testing.T) {
    factory := units.TemperatureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSolarTemperatures(100)
    if err != nil {
        t.Errorf("FromSolarTemperatures() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureSolarTemperature)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSolarTemperatures() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSolarTemperatures(math.NaN())
    if err == nil {
        t.Error("FromSolarTemperatures() with NaN value should return error")
    }

    _, err = factory.FromSolarTemperatures(math.Inf(1))
    if err == nil {
        t.Error("FromSolarTemperatures() with +Inf value should return error")
    }

    _, err = factory.FromSolarTemperatures(math.Inf(-1))
    if err == nil {
        t.Error("FromSolarTemperatures() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSolarTemperatures(0)
    if err != nil {
        t.Errorf("FromSolarTemperatures() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureSolarTemperature)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSolarTemperatures() with zero value = %v, want 0", converted)
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


func TestGetTemperatureAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.TemperatureUnits
        want string
    }{
        {
            name: "Kelvin abbreviation",
            unit: units.TemperatureKelvin,
            want: "K",
        },
        {
            name: "DegreeCelsius abbreviation",
            unit: units.TemperatureDegreeCelsius,
            want: "°C",
        },
        {
            name: "MillidegreeCelsius abbreviation",
            unit: units.TemperatureMillidegreeCelsius,
            want: "m°C",
        },
        {
            name: "DegreeDelisle abbreviation",
            unit: units.TemperatureDegreeDelisle,
            want: "°De",
        },
        {
            name: "DegreeFahrenheit abbreviation",
            unit: units.TemperatureDegreeFahrenheit,
            want: "°F",
        },
        {
            name: "DegreeNewton abbreviation",
            unit: units.TemperatureDegreeNewton,
            want: "°N",
        },
        {
            name: "DegreeRankine abbreviation",
            unit: units.TemperatureDegreeRankine,
            want: "°R",
        },
        {
            name: "DegreeReaumur abbreviation",
            unit: units.TemperatureDegreeReaumur,
            want: "°Ré",
        },
        {
            name: "DegreeRoemer abbreviation",
            unit: units.TemperatureDegreeRoemer,
            want: "°Rø",
        },
        {
            name: "SolarTemperature abbreviation",
            unit: units.TemperatureSolarTemperature,
            want: "T⊙",
        },
        {
            name: "invalid unit",
            unit: units.TemperatureUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetTemperatureAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetTemperatureAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestTemperature_String(t *testing.T) {
    factory := units.TemperatureFactory{}
    
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
            unit, err := factory.CreateTemperature(tt.value, units.TemperatureKelvin)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Temperature.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestTemperature_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.TemperatureFactory{}

	_, err := uf.CreateTemperature(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}