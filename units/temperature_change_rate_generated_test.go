// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

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
		cacheResult := a.DegreesCelsiusPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DegreesCelsiusPerMinute.
		// No expected conversion value provided for DegreesCelsiusPerMinute, verifying result is not NaN.
		result := a.DegreesCelsiusPerMinute()
		cacheResult := a.DegreesCelsiusPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesCelsiusPerMinute returned NaN")
		}
	}
	{
		// Test conversion to DegreesKelvinPerMinute.
		// No expected conversion value provided for DegreesKelvinPerMinute, verifying result is not NaN.
		result := a.DegreesKelvinPerMinute()
		cacheResult := a.DegreesKelvinPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesKelvinPerMinute returned NaN")
		}
	}
	{
		// Test conversion to DegreesFahrenheitPerMinute.
		// No expected conversion value provided for DegreesFahrenheitPerMinute, verifying result is not NaN.
		result := a.DegreesFahrenheitPerMinute()
		cacheResult := a.DegreesFahrenheitPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesFahrenheitPerMinute returned NaN")
		}
	}
	{
		// Test conversion to DegreesFahrenheitPerSecond.
		// No expected conversion value provided for DegreesFahrenheitPerSecond, verifying result is not NaN.
		result := a.DegreesFahrenheitPerSecond()
		cacheResult := a.DegreesFahrenheitPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesFahrenheitPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DegreesKelvinPerSecond.
		// No expected conversion value provided for DegreesKelvinPerSecond, verifying result is not NaN.
		result := a.DegreesKelvinPerSecond()
		cacheResult := a.DegreesKelvinPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesKelvinPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DegreesCelsiusPerHour.
		// No expected conversion value provided for DegreesCelsiusPerHour, verifying result is not NaN.
		result := a.DegreesCelsiusPerHour()
		cacheResult := a.DegreesCelsiusPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesCelsiusPerHour returned NaN")
		}
	}
	{
		// Test conversion to DegreesKelvinPerHour.
		// No expected conversion value provided for DegreesKelvinPerHour, verifying result is not NaN.
		result := a.DegreesKelvinPerHour()
		cacheResult := a.DegreesKelvinPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesKelvinPerHour returned NaN")
		}
	}
	{
		// Test conversion to DegreesFahrenheitPerHour.
		// No expected conversion value provided for DegreesFahrenheitPerHour, verifying result is not NaN.
		result := a.DegreesFahrenheitPerHour()
		cacheResult := a.DegreesFahrenheitPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesFahrenheitPerHour returned NaN")
		}
	}
	{
		// Test conversion to NanodegreesCelsiusPerSecond.
		// No expected conversion value provided for NanodegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.NanodegreesCelsiusPerSecond()
		cacheResult := a.NanodegreesCelsiusPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanodegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicrodegreesCelsiusPerSecond.
		// No expected conversion value provided for MicrodegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.MicrodegreesCelsiusPerSecond()
		cacheResult := a.MicrodegreesCelsiusPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrodegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillidegreesCelsiusPerSecond.
		// No expected conversion value provided for MillidegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.MillidegreesCelsiusPerSecond()
		cacheResult := a.MillidegreesCelsiusPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillidegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CentidegreesCelsiusPerSecond.
		// No expected conversion value provided for CentidegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.CentidegreesCelsiusPerSecond()
		cacheResult := a.CentidegreesCelsiusPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentidegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecidegreesCelsiusPerSecond.
		// No expected conversion value provided for DecidegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.DecidegreesCelsiusPerSecond()
		cacheResult := a.DecidegreesCelsiusPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecidegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecadegreesCelsiusPerSecond.
		// No expected conversion value provided for DecadegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.DecadegreesCelsiusPerSecond()
		cacheResult := a.DecadegreesCelsiusPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecadegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to HectodegreesCelsiusPerSecond.
		// No expected conversion value provided for HectodegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.HectodegreesCelsiusPerSecond()
		cacheResult := a.HectodegreesCelsiusPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to HectodegreesCelsiusPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilodegreesCelsiusPerSecond.
		// No expected conversion value provided for KilodegreesCelsiusPerSecond, verifying result is not NaN.
		result := a.KilodegreesCelsiusPerSecond()
		cacheResult := a.KilodegreesCelsiusPerSecond()
		if math.IsNaN(result) || cacheResult != result {
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

func TestTemperatureChangeRateFactory_FromDto(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateDegreeCelsiusPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.TemperatureChangeRateDto{
        Value: math.NaN(),
        Unit:  units.TemperatureChangeRateDegreeCelsiusPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test DegreeCelsiusPerSecond conversion
    degrees_celsius_per_secondDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateDegreeCelsiusPerSecond,
    }
    
    var degrees_celsius_per_secondResult *units.TemperatureChangeRate
    degrees_celsius_per_secondResult, err = factory.FromDto(degrees_celsius_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeCelsiusPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateDegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test DegreeCelsiusPerMinute conversion
    degrees_celsius_per_minuteDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateDegreeCelsiusPerMinute,
    }
    
    var degrees_celsius_per_minuteResult *units.TemperatureChangeRate
    degrees_celsius_per_minuteResult, err = factory.FromDto(degrees_celsius_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeCelsiusPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsius_per_minuteResult.Convert(units.TemperatureChangeRateDegreeCelsiusPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsiusPerMinute = %v, want %v", converted, 100)
    }
    // Test DegreeKelvinPerMinute conversion
    degrees_kelvin_per_minuteDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateDegreeKelvinPerMinute,
    }
    
    var degrees_kelvin_per_minuteResult *units.TemperatureChangeRate
    degrees_kelvin_per_minuteResult, err = factory.FromDto(degrees_kelvin_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeKelvinPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_kelvin_per_minuteResult.Convert(units.TemperatureChangeRateDegreeKelvinPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeKelvinPerMinute = %v, want %v", converted, 100)
    }
    // Test DegreeFahrenheitPerMinute conversion
    degrees_fahrenheit_per_minuteDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateDegreeFahrenheitPerMinute,
    }
    
    var degrees_fahrenheit_per_minuteResult *units.TemperatureChangeRate
    degrees_fahrenheit_per_minuteResult, err = factory.FromDto(degrees_fahrenheit_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeFahrenheitPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_fahrenheit_per_minuteResult.Convert(units.TemperatureChangeRateDegreeFahrenheitPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeFahrenheitPerMinute = %v, want %v", converted, 100)
    }
    // Test DegreeFahrenheitPerSecond conversion
    degrees_fahrenheit_per_secondDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateDegreeFahrenheitPerSecond,
    }
    
    var degrees_fahrenheit_per_secondResult *units.TemperatureChangeRate
    degrees_fahrenheit_per_secondResult, err = factory.FromDto(degrees_fahrenheit_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeFahrenheitPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_fahrenheit_per_secondResult.Convert(units.TemperatureChangeRateDegreeFahrenheitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeFahrenheitPerSecond = %v, want %v", converted, 100)
    }
    // Test DegreeKelvinPerSecond conversion
    degrees_kelvin_per_secondDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateDegreeKelvinPerSecond,
    }
    
    var degrees_kelvin_per_secondResult *units.TemperatureChangeRate
    degrees_kelvin_per_secondResult, err = factory.FromDto(degrees_kelvin_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeKelvinPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_kelvin_per_secondResult.Convert(units.TemperatureChangeRateDegreeKelvinPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeKelvinPerSecond = %v, want %v", converted, 100)
    }
    // Test DegreeCelsiusPerHour conversion
    degrees_celsius_per_hourDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateDegreeCelsiusPerHour,
    }
    
    var degrees_celsius_per_hourResult *units.TemperatureChangeRate
    degrees_celsius_per_hourResult, err = factory.FromDto(degrees_celsius_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeCelsiusPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsius_per_hourResult.Convert(units.TemperatureChangeRateDegreeCelsiusPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsiusPerHour = %v, want %v", converted, 100)
    }
    // Test DegreeKelvinPerHour conversion
    degrees_kelvin_per_hourDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateDegreeKelvinPerHour,
    }
    
    var degrees_kelvin_per_hourResult *units.TemperatureChangeRate
    degrees_kelvin_per_hourResult, err = factory.FromDto(degrees_kelvin_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeKelvinPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_kelvin_per_hourResult.Convert(units.TemperatureChangeRateDegreeKelvinPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeKelvinPerHour = %v, want %v", converted, 100)
    }
    // Test DegreeFahrenheitPerHour conversion
    degrees_fahrenheit_per_hourDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateDegreeFahrenheitPerHour,
    }
    
    var degrees_fahrenheit_per_hourResult *units.TemperatureChangeRate
    degrees_fahrenheit_per_hourResult, err = factory.FromDto(degrees_fahrenheit_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeFahrenheitPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_fahrenheit_per_hourResult.Convert(units.TemperatureChangeRateDegreeFahrenheitPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeFahrenheitPerHour = %v, want %v", converted, 100)
    }
    // Test NanodegreeCelsiusPerSecond conversion
    nanodegrees_celsius_per_secondDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateNanodegreeCelsiusPerSecond,
    }
    
    var nanodegrees_celsius_per_secondResult *units.TemperatureChangeRate
    nanodegrees_celsius_per_secondResult, err = factory.FromDto(nanodegrees_celsius_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with NanodegreeCelsiusPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanodegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateNanodegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanodegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test MicrodegreeCelsiusPerSecond conversion
    microdegrees_celsius_per_secondDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateMicrodegreeCelsiusPerSecond,
    }
    
    var microdegrees_celsius_per_secondResult *units.TemperatureChangeRate
    microdegrees_celsius_per_secondResult, err = factory.FromDto(microdegrees_celsius_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MicrodegreeCelsiusPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microdegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateMicrodegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrodegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test MillidegreeCelsiusPerSecond conversion
    millidegrees_celsius_per_secondDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateMillidegreeCelsiusPerSecond,
    }
    
    var millidegrees_celsius_per_secondResult *units.TemperatureChangeRate
    millidegrees_celsius_per_secondResult, err = factory.FromDto(millidegrees_celsius_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MillidegreeCelsiusPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millidegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateMillidegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillidegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test CentidegreeCelsiusPerSecond conversion
    centidegrees_celsius_per_secondDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateCentidegreeCelsiusPerSecond,
    }
    
    var centidegrees_celsius_per_secondResult *units.TemperatureChangeRate
    centidegrees_celsius_per_secondResult, err = factory.FromDto(centidegrees_celsius_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with CentidegreeCelsiusPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centidegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateCentidegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentidegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test DecidegreeCelsiusPerSecond conversion
    decidegrees_celsius_per_secondDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateDecidegreeCelsiusPerSecond,
    }
    
    var decidegrees_celsius_per_secondResult *units.TemperatureChangeRate
    decidegrees_celsius_per_secondResult, err = factory.FromDto(decidegrees_celsius_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DecidegreeCelsiusPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decidegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateDecidegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecidegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test DecadegreeCelsiusPerSecond conversion
    decadegrees_celsius_per_secondDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateDecadegreeCelsiusPerSecond,
    }
    
    var decadegrees_celsius_per_secondResult *units.TemperatureChangeRate
    decadegrees_celsius_per_secondResult, err = factory.FromDto(decadegrees_celsius_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DecadegreeCelsiusPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decadegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateDecadegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecadegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test HectodegreeCelsiusPerSecond conversion
    hectodegrees_celsius_per_secondDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateHectodegreeCelsiusPerSecond,
    }
    
    var hectodegrees_celsius_per_secondResult *units.TemperatureChangeRate
    hectodegrees_celsius_per_secondResult, err = factory.FromDto(hectodegrees_celsius_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with HectodegreeCelsiusPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectodegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateHectodegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectodegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test KilodegreeCelsiusPerSecond conversion
    kilodegrees_celsius_per_secondDto := units.TemperatureChangeRateDto{
        Value: 100,
        Unit:  units.TemperatureChangeRateKilodegreeCelsiusPerSecond,
    }
    
    var kilodegrees_celsius_per_secondResult *units.TemperatureChangeRate
    kilodegrees_celsius_per_secondResult, err = factory.FromDto(kilodegrees_celsius_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilodegreeCelsiusPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilodegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateKilodegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilodegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.TemperatureChangeRateDto{
        Value: 0,
        Unit:  units.TemperatureChangeRateDegreeCelsiusPerSecond,
    }
    
    var zeroResult *units.TemperatureChangeRate
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestTemperatureChangeRateFactory_FromDtoJSON(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "DegreeCelsiusPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "DegreeCelsiusPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.TemperatureChangeRateDto{
        Value: nanValue,
        Unit:  units.TemperatureChangeRateDegreeCelsiusPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with DegreeCelsiusPerSecond unit
    degrees_celsius_per_secondJSON := []byte(`{"value": 100, "unit": "DegreeCelsiusPerSecond"}`)
    degrees_celsius_per_secondResult, err := factory.FromDtoJSON(degrees_celsius_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeCelsiusPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateDegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeCelsiusPerMinute unit
    degrees_celsius_per_minuteJSON := []byte(`{"value": 100, "unit": "DegreeCelsiusPerMinute"}`)
    degrees_celsius_per_minuteResult, err := factory.FromDtoJSON(degrees_celsius_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeCelsiusPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsius_per_minuteResult.Convert(units.TemperatureChangeRateDegreeCelsiusPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsiusPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeKelvinPerMinute unit
    degrees_kelvin_per_minuteJSON := []byte(`{"value": 100, "unit": "DegreeKelvinPerMinute"}`)
    degrees_kelvin_per_minuteResult, err := factory.FromDtoJSON(degrees_kelvin_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeKelvinPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_kelvin_per_minuteResult.Convert(units.TemperatureChangeRateDegreeKelvinPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeKelvinPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeFahrenheitPerMinute unit
    degrees_fahrenheit_per_minuteJSON := []byte(`{"value": 100, "unit": "DegreeFahrenheitPerMinute"}`)
    degrees_fahrenheit_per_minuteResult, err := factory.FromDtoJSON(degrees_fahrenheit_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeFahrenheitPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_fahrenheit_per_minuteResult.Convert(units.TemperatureChangeRateDegreeFahrenheitPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeFahrenheitPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeFahrenheitPerSecond unit
    degrees_fahrenheit_per_secondJSON := []byte(`{"value": 100, "unit": "DegreeFahrenheitPerSecond"}`)
    degrees_fahrenheit_per_secondResult, err := factory.FromDtoJSON(degrees_fahrenheit_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeFahrenheitPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_fahrenheit_per_secondResult.Convert(units.TemperatureChangeRateDegreeFahrenheitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeFahrenheitPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeKelvinPerSecond unit
    degrees_kelvin_per_secondJSON := []byte(`{"value": 100, "unit": "DegreeKelvinPerSecond"}`)
    degrees_kelvin_per_secondResult, err := factory.FromDtoJSON(degrees_kelvin_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeKelvinPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_kelvin_per_secondResult.Convert(units.TemperatureChangeRateDegreeKelvinPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeKelvinPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeCelsiusPerHour unit
    degrees_celsius_per_hourJSON := []byte(`{"value": 100, "unit": "DegreeCelsiusPerHour"}`)
    degrees_celsius_per_hourResult, err := factory.FromDtoJSON(degrees_celsius_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeCelsiusPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsius_per_hourResult.Convert(units.TemperatureChangeRateDegreeCelsiusPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsiusPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeKelvinPerHour unit
    degrees_kelvin_per_hourJSON := []byte(`{"value": 100, "unit": "DegreeKelvinPerHour"}`)
    degrees_kelvin_per_hourResult, err := factory.FromDtoJSON(degrees_kelvin_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeKelvinPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_kelvin_per_hourResult.Convert(units.TemperatureChangeRateDegreeKelvinPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeKelvinPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeFahrenheitPerHour unit
    degrees_fahrenheit_per_hourJSON := []byte(`{"value": 100, "unit": "DegreeFahrenheitPerHour"}`)
    degrees_fahrenheit_per_hourResult, err := factory.FromDtoJSON(degrees_fahrenheit_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeFahrenheitPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_fahrenheit_per_hourResult.Convert(units.TemperatureChangeRateDegreeFahrenheitPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeFahrenheitPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with NanodegreeCelsiusPerSecond unit
    nanodegrees_celsius_per_secondJSON := []byte(`{"value": 100, "unit": "NanodegreeCelsiusPerSecond"}`)
    nanodegrees_celsius_per_secondResult, err := factory.FromDtoJSON(nanodegrees_celsius_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanodegreeCelsiusPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanodegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateNanodegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanodegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicrodegreeCelsiusPerSecond unit
    microdegrees_celsius_per_secondJSON := []byte(`{"value": 100, "unit": "MicrodegreeCelsiusPerSecond"}`)
    microdegrees_celsius_per_secondResult, err := factory.FromDtoJSON(microdegrees_celsius_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrodegreeCelsiusPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microdegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateMicrodegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrodegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MillidegreeCelsiusPerSecond unit
    millidegrees_celsius_per_secondJSON := []byte(`{"value": 100, "unit": "MillidegreeCelsiusPerSecond"}`)
    millidegrees_celsius_per_secondResult, err := factory.FromDtoJSON(millidegrees_celsius_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillidegreeCelsiusPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millidegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateMillidegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillidegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with CentidegreeCelsiusPerSecond unit
    centidegrees_celsius_per_secondJSON := []byte(`{"value": 100, "unit": "CentidegreeCelsiusPerSecond"}`)
    centidegrees_celsius_per_secondResult, err := factory.FromDtoJSON(centidegrees_celsius_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentidegreeCelsiusPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centidegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateCentidegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentidegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DecidegreeCelsiusPerSecond unit
    decidegrees_celsius_per_secondJSON := []byte(`{"value": 100, "unit": "DecidegreeCelsiusPerSecond"}`)
    decidegrees_celsius_per_secondResult, err := factory.FromDtoJSON(decidegrees_celsius_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecidegreeCelsiusPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decidegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateDecidegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecidegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DecadegreeCelsiusPerSecond unit
    decadegrees_celsius_per_secondJSON := []byte(`{"value": 100, "unit": "DecadegreeCelsiusPerSecond"}`)
    decadegrees_celsius_per_secondResult, err := factory.FromDtoJSON(decadegrees_celsius_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecadegreeCelsiusPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decadegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateDecadegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecadegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with HectodegreeCelsiusPerSecond unit
    hectodegrees_celsius_per_secondJSON := []byte(`{"value": 100, "unit": "HectodegreeCelsiusPerSecond"}`)
    hectodegrees_celsius_per_secondResult, err := factory.FromDtoJSON(hectodegrees_celsius_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HectodegreeCelsiusPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectodegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateHectodegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectodegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilodegreeCelsiusPerSecond unit
    kilodegrees_celsius_per_secondJSON := []byte(`{"value": 100, "unit": "KilodegreeCelsiusPerSecond"}`)
    kilodegrees_celsius_per_secondResult, err := factory.FromDtoJSON(kilodegrees_celsius_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilodegreeCelsiusPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilodegrees_celsius_per_secondResult.Convert(units.TemperatureChangeRateKilodegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilodegreeCelsiusPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "DegreeCelsiusPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromDegreesCelsiusPerSecond function
func TestTemperatureChangeRateFactory_FromDegreesCelsiusPerSecond(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesCelsiusPerSecond(100)
    if err != nil {
        t.Errorf("FromDegreesCelsiusPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateDegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesCelsiusPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesCelsiusPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDegreesCelsiusPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDegreesCelsiusPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesCelsiusPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDegreesCelsiusPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesCelsiusPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesCelsiusPerSecond(0)
    if err != nil {
        t.Errorf("FromDegreesCelsiusPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateDegreeCelsiusPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesCelsiusPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesCelsiusPerMinute function
func TestTemperatureChangeRateFactory_FromDegreesCelsiusPerMinute(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesCelsiusPerMinute(100)
    if err != nil {
        t.Errorf("FromDegreesCelsiusPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateDegreeCelsiusPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesCelsiusPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesCelsiusPerMinute(math.NaN())
    if err == nil {
        t.Error("FromDegreesCelsiusPerMinute() with NaN value should return error")
    }

    _, err = factory.FromDegreesCelsiusPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesCelsiusPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromDegreesCelsiusPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesCelsiusPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesCelsiusPerMinute(0)
    if err != nil {
        t.Errorf("FromDegreesCelsiusPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateDegreeCelsiusPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesCelsiusPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesKelvinPerMinute function
func TestTemperatureChangeRateFactory_FromDegreesKelvinPerMinute(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesKelvinPerMinute(100)
    if err != nil {
        t.Errorf("FromDegreesKelvinPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateDegreeKelvinPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesKelvinPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesKelvinPerMinute(math.NaN())
    if err == nil {
        t.Error("FromDegreesKelvinPerMinute() with NaN value should return error")
    }

    _, err = factory.FromDegreesKelvinPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesKelvinPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromDegreesKelvinPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesKelvinPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesKelvinPerMinute(0)
    if err != nil {
        t.Errorf("FromDegreesKelvinPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateDegreeKelvinPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesKelvinPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesFahrenheitPerMinute function
func TestTemperatureChangeRateFactory_FromDegreesFahrenheitPerMinute(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesFahrenheitPerMinute(100)
    if err != nil {
        t.Errorf("FromDegreesFahrenheitPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateDegreeFahrenheitPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesFahrenheitPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesFahrenheitPerMinute(math.NaN())
    if err == nil {
        t.Error("FromDegreesFahrenheitPerMinute() with NaN value should return error")
    }

    _, err = factory.FromDegreesFahrenheitPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesFahrenheitPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromDegreesFahrenheitPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesFahrenheitPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesFahrenheitPerMinute(0)
    if err != nil {
        t.Errorf("FromDegreesFahrenheitPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateDegreeFahrenheitPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesFahrenheitPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesFahrenheitPerSecond function
func TestTemperatureChangeRateFactory_FromDegreesFahrenheitPerSecond(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesFahrenheitPerSecond(100)
    if err != nil {
        t.Errorf("FromDegreesFahrenheitPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateDegreeFahrenheitPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesFahrenheitPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesFahrenheitPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDegreesFahrenheitPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDegreesFahrenheitPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesFahrenheitPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDegreesFahrenheitPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesFahrenheitPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesFahrenheitPerSecond(0)
    if err != nil {
        t.Errorf("FromDegreesFahrenheitPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateDegreeFahrenheitPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesFahrenheitPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesKelvinPerSecond function
func TestTemperatureChangeRateFactory_FromDegreesKelvinPerSecond(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesKelvinPerSecond(100)
    if err != nil {
        t.Errorf("FromDegreesKelvinPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateDegreeKelvinPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesKelvinPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesKelvinPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDegreesKelvinPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDegreesKelvinPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesKelvinPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDegreesKelvinPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesKelvinPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesKelvinPerSecond(0)
    if err != nil {
        t.Errorf("FromDegreesKelvinPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateDegreeKelvinPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesKelvinPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesCelsiusPerHour function
func TestTemperatureChangeRateFactory_FromDegreesCelsiusPerHour(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesCelsiusPerHour(100)
    if err != nil {
        t.Errorf("FromDegreesCelsiusPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateDegreeCelsiusPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesCelsiusPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesCelsiusPerHour(math.NaN())
    if err == nil {
        t.Error("FromDegreesCelsiusPerHour() with NaN value should return error")
    }

    _, err = factory.FromDegreesCelsiusPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesCelsiusPerHour() with +Inf value should return error")
    }

    _, err = factory.FromDegreesCelsiusPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesCelsiusPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesCelsiusPerHour(0)
    if err != nil {
        t.Errorf("FromDegreesCelsiusPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateDegreeCelsiusPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesCelsiusPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesKelvinPerHour function
func TestTemperatureChangeRateFactory_FromDegreesKelvinPerHour(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesKelvinPerHour(100)
    if err != nil {
        t.Errorf("FromDegreesKelvinPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateDegreeKelvinPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesKelvinPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesKelvinPerHour(math.NaN())
    if err == nil {
        t.Error("FromDegreesKelvinPerHour() with NaN value should return error")
    }

    _, err = factory.FromDegreesKelvinPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesKelvinPerHour() with +Inf value should return error")
    }

    _, err = factory.FromDegreesKelvinPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesKelvinPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesKelvinPerHour(0)
    if err != nil {
        t.Errorf("FromDegreesKelvinPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateDegreeKelvinPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesKelvinPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesFahrenheitPerHour function
func TestTemperatureChangeRateFactory_FromDegreesFahrenheitPerHour(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesFahrenheitPerHour(100)
    if err != nil {
        t.Errorf("FromDegreesFahrenheitPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateDegreeFahrenheitPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesFahrenheitPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesFahrenheitPerHour(math.NaN())
    if err == nil {
        t.Error("FromDegreesFahrenheitPerHour() with NaN value should return error")
    }

    _, err = factory.FromDegreesFahrenheitPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesFahrenheitPerHour() with +Inf value should return error")
    }

    _, err = factory.FromDegreesFahrenheitPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesFahrenheitPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesFahrenheitPerHour(0)
    if err != nil {
        t.Errorf("FromDegreesFahrenheitPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateDegreeFahrenheitPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesFahrenheitPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromNanodegreesCelsiusPerSecond function
func TestTemperatureChangeRateFactory_FromNanodegreesCelsiusPerSecond(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanodegreesCelsiusPerSecond(100)
    if err != nil {
        t.Errorf("FromNanodegreesCelsiusPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateNanodegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanodegreesCelsiusPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanodegreesCelsiusPerSecond(math.NaN())
    if err == nil {
        t.Error("FromNanodegreesCelsiusPerSecond() with NaN value should return error")
    }

    _, err = factory.FromNanodegreesCelsiusPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromNanodegreesCelsiusPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromNanodegreesCelsiusPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromNanodegreesCelsiusPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanodegreesCelsiusPerSecond(0)
    if err != nil {
        t.Errorf("FromNanodegreesCelsiusPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateNanodegreeCelsiusPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanodegreesCelsiusPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrodegreesCelsiusPerSecond function
func TestTemperatureChangeRateFactory_FromMicrodegreesCelsiusPerSecond(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrodegreesCelsiusPerSecond(100)
    if err != nil {
        t.Errorf("FromMicrodegreesCelsiusPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateMicrodegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrodegreesCelsiusPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrodegreesCelsiusPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMicrodegreesCelsiusPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMicrodegreesCelsiusPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMicrodegreesCelsiusPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMicrodegreesCelsiusPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrodegreesCelsiusPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrodegreesCelsiusPerSecond(0)
    if err != nil {
        t.Errorf("FromMicrodegreesCelsiusPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateMicrodegreeCelsiusPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrodegreesCelsiusPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMillidegreesCelsiusPerSecond function
func TestTemperatureChangeRateFactory_FromMillidegreesCelsiusPerSecond(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillidegreesCelsiusPerSecond(100)
    if err != nil {
        t.Errorf("FromMillidegreesCelsiusPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateMillidegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillidegreesCelsiusPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillidegreesCelsiusPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMillidegreesCelsiusPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMillidegreesCelsiusPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMillidegreesCelsiusPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMillidegreesCelsiusPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMillidegreesCelsiusPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillidegreesCelsiusPerSecond(0)
    if err != nil {
        t.Errorf("FromMillidegreesCelsiusPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateMillidegreeCelsiusPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillidegreesCelsiusPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromCentidegreesCelsiusPerSecond function
func TestTemperatureChangeRateFactory_FromCentidegreesCelsiusPerSecond(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentidegreesCelsiusPerSecond(100)
    if err != nil {
        t.Errorf("FromCentidegreesCelsiusPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateCentidegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentidegreesCelsiusPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentidegreesCelsiusPerSecond(math.NaN())
    if err == nil {
        t.Error("FromCentidegreesCelsiusPerSecond() with NaN value should return error")
    }

    _, err = factory.FromCentidegreesCelsiusPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromCentidegreesCelsiusPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromCentidegreesCelsiusPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromCentidegreesCelsiusPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentidegreesCelsiusPerSecond(0)
    if err != nil {
        t.Errorf("FromCentidegreesCelsiusPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateCentidegreeCelsiusPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentidegreesCelsiusPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDecidegreesCelsiusPerSecond function
func TestTemperatureChangeRateFactory_FromDecidegreesCelsiusPerSecond(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecidegreesCelsiusPerSecond(100)
    if err != nil {
        t.Errorf("FromDecidegreesCelsiusPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateDecidegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecidegreesCelsiusPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecidegreesCelsiusPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDecidegreesCelsiusPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDecidegreesCelsiusPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDecidegreesCelsiusPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDecidegreesCelsiusPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDecidegreesCelsiusPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecidegreesCelsiusPerSecond(0)
    if err != nil {
        t.Errorf("FromDecidegreesCelsiusPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateDecidegreeCelsiusPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecidegreesCelsiusPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDecadegreesCelsiusPerSecond function
func TestTemperatureChangeRateFactory_FromDecadegreesCelsiusPerSecond(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecadegreesCelsiusPerSecond(100)
    if err != nil {
        t.Errorf("FromDecadegreesCelsiusPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateDecadegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecadegreesCelsiusPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecadegreesCelsiusPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDecadegreesCelsiusPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDecadegreesCelsiusPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDecadegreesCelsiusPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDecadegreesCelsiusPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDecadegreesCelsiusPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecadegreesCelsiusPerSecond(0)
    if err != nil {
        t.Errorf("FromDecadegreesCelsiusPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateDecadegreeCelsiusPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecadegreesCelsiusPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromHectodegreesCelsiusPerSecond function
func TestTemperatureChangeRateFactory_FromHectodegreesCelsiusPerSecond(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectodegreesCelsiusPerSecond(100)
    if err != nil {
        t.Errorf("FromHectodegreesCelsiusPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateHectodegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectodegreesCelsiusPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectodegreesCelsiusPerSecond(math.NaN())
    if err == nil {
        t.Error("FromHectodegreesCelsiusPerSecond() with NaN value should return error")
    }

    _, err = factory.FromHectodegreesCelsiusPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromHectodegreesCelsiusPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromHectodegreesCelsiusPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromHectodegreesCelsiusPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectodegreesCelsiusPerSecond(0)
    if err != nil {
        t.Errorf("FromHectodegreesCelsiusPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateHectodegreeCelsiusPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectodegreesCelsiusPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilodegreesCelsiusPerSecond function
func TestTemperatureChangeRateFactory_FromKilodegreesCelsiusPerSecond(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilodegreesCelsiusPerSecond(100)
    if err != nil {
        t.Errorf("FromKilodegreesCelsiusPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureChangeRateKilodegreeCelsiusPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilodegreesCelsiusPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilodegreesCelsiusPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilodegreesCelsiusPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilodegreesCelsiusPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilodegreesCelsiusPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilodegreesCelsiusPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilodegreesCelsiusPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilodegreesCelsiusPerSecond(0)
    if err != nil {
        t.Errorf("FromKilodegreesCelsiusPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureChangeRateKilodegreeCelsiusPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilodegreesCelsiusPerSecond() with zero value = %v, want 0", converted)
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


func TestGetTemperatureChangeRateAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.TemperatureChangeRateUnits
        want string
    }{
        {
            name: "DegreeCelsiusPerSecond abbreviation",
            unit: units.TemperatureChangeRateDegreeCelsiusPerSecond,
            want: "°C/s",
        },
        {
            name: "DegreeCelsiusPerMinute abbreviation",
            unit: units.TemperatureChangeRateDegreeCelsiusPerMinute,
            want: "°C/min",
        },
        {
            name: "DegreeKelvinPerMinute abbreviation",
            unit: units.TemperatureChangeRateDegreeKelvinPerMinute,
            want: "K/min",
        },
        {
            name: "DegreeFahrenheitPerMinute abbreviation",
            unit: units.TemperatureChangeRateDegreeFahrenheitPerMinute,
            want: "°F/min",
        },
        {
            name: "DegreeFahrenheitPerSecond abbreviation",
            unit: units.TemperatureChangeRateDegreeFahrenheitPerSecond,
            want: "°F/s",
        },
        {
            name: "DegreeKelvinPerSecond abbreviation",
            unit: units.TemperatureChangeRateDegreeKelvinPerSecond,
            want: "K/s",
        },
        {
            name: "DegreeCelsiusPerHour abbreviation",
            unit: units.TemperatureChangeRateDegreeCelsiusPerHour,
            want: "°C/h",
        },
        {
            name: "DegreeKelvinPerHour abbreviation",
            unit: units.TemperatureChangeRateDegreeKelvinPerHour,
            want: "K/h",
        },
        {
            name: "DegreeFahrenheitPerHour abbreviation",
            unit: units.TemperatureChangeRateDegreeFahrenheitPerHour,
            want: "°F/h",
        },
        {
            name: "NanodegreeCelsiusPerSecond abbreviation",
            unit: units.TemperatureChangeRateNanodegreeCelsiusPerSecond,
            want: "n°C/s",
        },
        {
            name: "MicrodegreeCelsiusPerSecond abbreviation",
            unit: units.TemperatureChangeRateMicrodegreeCelsiusPerSecond,
            want: "μ°C/s",
        },
        {
            name: "MillidegreeCelsiusPerSecond abbreviation",
            unit: units.TemperatureChangeRateMillidegreeCelsiusPerSecond,
            want: "m°C/s",
        },
        {
            name: "CentidegreeCelsiusPerSecond abbreviation",
            unit: units.TemperatureChangeRateCentidegreeCelsiusPerSecond,
            want: "c°C/s",
        },
        {
            name: "DecidegreeCelsiusPerSecond abbreviation",
            unit: units.TemperatureChangeRateDecidegreeCelsiusPerSecond,
            want: "d°C/s",
        },
        {
            name: "DecadegreeCelsiusPerSecond abbreviation",
            unit: units.TemperatureChangeRateDecadegreeCelsiusPerSecond,
            want: "da°C/s",
        },
        {
            name: "HectodegreeCelsiusPerSecond abbreviation",
            unit: units.TemperatureChangeRateHectodegreeCelsiusPerSecond,
            want: "h°C/s",
        },
        {
            name: "KilodegreeCelsiusPerSecond abbreviation",
            unit: units.TemperatureChangeRateKilodegreeCelsiusPerSecond,
            want: "k°C/s",
        },
        {
            name: "invalid unit",
            unit: units.TemperatureChangeRateUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetTemperatureChangeRateAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetTemperatureChangeRateAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestTemperatureChangeRate_String(t *testing.T) {
    factory := units.TemperatureChangeRateFactory{}
    
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
            unit, err := factory.CreateTemperatureChangeRate(tt.value, units.TemperatureChangeRateDegreeCelsiusPerSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("TemperatureChangeRate.String() = %v, want %v", got, tt.want)
            }
        })
    }
}