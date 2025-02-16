// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricCurrentGradientDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "AmperePerSecond"}`
	
	factory := units.ElectricCurrentGradientDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricCurrentGradientAmperePerSecond {
		t.Errorf("expected unit %v, got %v", units.ElectricCurrentGradientAmperePerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "AmperePerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricCurrentGradientDto_ToJSON(t *testing.T) {
	dto := units.ElectricCurrentGradientDto{
		Value: 45,
		Unit:  units.ElectricCurrentGradientAmperePerSecond,
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
	if result["unit"].(string) != string(units.ElectricCurrentGradientAmperePerSecond) {
		t.Errorf("expected unit %s, got %v", units.ElectricCurrentGradientAmperePerSecond, result["unit"])
	}
}

func TestNewElectricCurrentGradient_InvalidValue(t *testing.T) {
	factory := units.ElectricCurrentGradientFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricCurrentGradient(math.NaN(), units.ElectricCurrentGradientAmperePerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricCurrentGradient(math.Inf(1), units.ElectricCurrentGradientAmperePerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricCurrentGradientConversions(t *testing.T) {
	factory := units.ElectricCurrentGradientFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricCurrentGradient(180, units.ElectricCurrentGradientAmperePerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to AmperesPerSecond.
		// No expected conversion value provided for AmperesPerSecond, verifying result is not NaN.
		result := a.AmperesPerSecond()
		cacheResult := a.AmperesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AmperesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to AmperesPerMinute.
		// No expected conversion value provided for AmperesPerMinute, verifying result is not NaN.
		result := a.AmperesPerMinute()
		cacheResult := a.AmperesPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AmperesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to AmperesPerMillisecond.
		// No expected conversion value provided for AmperesPerMillisecond, verifying result is not NaN.
		result := a.AmperesPerMillisecond()
		cacheResult := a.AmperesPerMillisecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AmperesPerMillisecond returned NaN")
		}
	}
	{
		// Test conversion to AmperesPerMicrosecond.
		// No expected conversion value provided for AmperesPerMicrosecond, verifying result is not NaN.
		result := a.AmperesPerMicrosecond()
		cacheResult := a.AmperesPerMicrosecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AmperesPerMicrosecond returned NaN")
		}
	}
	{
		// Test conversion to AmperesPerNanosecond.
		// No expected conversion value provided for AmperesPerNanosecond, verifying result is not NaN.
		result := a.AmperesPerNanosecond()
		cacheResult := a.AmperesPerNanosecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AmperesPerNanosecond returned NaN")
		}
	}
	{
		// Test conversion to MilliamperesPerSecond.
		// No expected conversion value provided for MilliamperesPerSecond, verifying result is not NaN.
		result := a.MilliamperesPerSecond()
		cacheResult := a.MilliamperesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliamperesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MilliamperesPerMinute.
		// No expected conversion value provided for MilliamperesPerMinute, verifying result is not NaN.
		result := a.MilliamperesPerMinute()
		cacheResult := a.MilliamperesPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliamperesPerMinute returned NaN")
		}
	}
}

func TestElectricCurrentGradient_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricCurrentGradientFactory{}
	a, err := factory.CreateElectricCurrentGradient(90, units.ElectricCurrentGradientAmperePerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricCurrentGradientAmperePerSecond {
		t.Errorf("expected default unit AmperePerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricCurrentGradientAmperePerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricCurrentGradientDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricCurrentGradientAmperePerSecond {
		t.Errorf("expected unit AmperePerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricCurrentGradientFactory_FromDto(t *testing.T) {
    factory := units.ElectricCurrentGradientFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricCurrentGradientDto{
        Value: 100,
        Unit:  units.ElectricCurrentGradientAmperePerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricCurrentGradientDto{
        Value: math.NaN(),
        Unit:  units.ElectricCurrentGradientAmperePerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test AmperePerSecond conversion
    amperes_per_secondDto := units.ElectricCurrentGradientDto{
        Value: 100,
        Unit:  units.ElectricCurrentGradientAmperePerSecond,
    }
    
    var amperes_per_secondResult *units.ElectricCurrentGradient
    amperes_per_secondResult, err = factory.FromDto(amperes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with AmperePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_secondResult.Convert(units.ElectricCurrentGradientAmperePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerSecond = %v, want %v", converted, 100)
    }
    // Test AmperePerMinute conversion
    amperes_per_minuteDto := units.ElectricCurrentGradientDto{
        Value: 100,
        Unit:  units.ElectricCurrentGradientAmperePerMinute,
    }
    
    var amperes_per_minuteResult *units.ElectricCurrentGradient
    amperes_per_minuteResult, err = factory.FromDto(amperes_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with AmperePerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_minuteResult.Convert(units.ElectricCurrentGradientAmperePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerMinute = %v, want %v", converted, 100)
    }
    // Test AmperePerMillisecond conversion
    amperes_per_millisecondDto := units.ElectricCurrentGradientDto{
        Value: 100,
        Unit:  units.ElectricCurrentGradientAmperePerMillisecond,
    }
    
    var amperes_per_millisecondResult *units.ElectricCurrentGradient
    amperes_per_millisecondResult, err = factory.FromDto(amperes_per_millisecondDto)
    if err != nil {
        t.Errorf("FromDto() with AmperePerMillisecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_millisecondResult.Convert(units.ElectricCurrentGradientAmperePerMillisecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerMillisecond = %v, want %v", converted, 100)
    }
    // Test AmperePerMicrosecond conversion
    amperes_per_microsecondDto := units.ElectricCurrentGradientDto{
        Value: 100,
        Unit:  units.ElectricCurrentGradientAmperePerMicrosecond,
    }
    
    var amperes_per_microsecondResult *units.ElectricCurrentGradient
    amperes_per_microsecondResult, err = factory.FromDto(amperes_per_microsecondDto)
    if err != nil {
        t.Errorf("FromDto() with AmperePerMicrosecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_microsecondResult.Convert(units.ElectricCurrentGradientAmperePerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerMicrosecond = %v, want %v", converted, 100)
    }
    // Test AmperePerNanosecond conversion
    amperes_per_nanosecondDto := units.ElectricCurrentGradientDto{
        Value: 100,
        Unit:  units.ElectricCurrentGradientAmperePerNanosecond,
    }
    
    var amperes_per_nanosecondResult *units.ElectricCurrentGradient
    amperes_per_nanosecondResult, err = factory.FromDto(amperes_per_nanosecondDto)
    if err != nil {
        t.Errorf("FromDto() with AmperePerNanosecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_nanosecondResult.Convert(units.ElectricCurrentGradientAmperePerNanosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerNanosecond = %v, want %v", converted, 100)
    }
    // Test MilliamperePerSecond conversion
    milliamperes_per_secondDto := units.ElectricCurrentGradientDto{
        Value: 100,
        Unit:  units.ElectricCurrentGradientMilliamperePerSecond,
    }
    
    var milliamperes_per_secondResult *units.ElectricCurrentGradient
    milliamperes_per_secondResult, err = factory.FromDto(milliamperes_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MilliamperePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliamperes_per_secondResult.Convert(units.ElectricCurrentGradientMilliamperePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliamperePerSecond = %v, want %v", converted, 100)
    }
    // Test MilliamperePerMinute conversion
    milliamperes_per_minuteDto := units.ElectricCurrentGradientDto{
        Value: 100,
        Unit:  units.ElectricCurrentGradientMilliamperePerMinute,
    }
    
    var milliamperes_per_minuteResult *units.ElectricCurrentGradient
    milliamperes_per_minuteResult, err = factory.FromDto(milliamperes_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with MilliamperePerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliamperes_per_minuteResult.Convert(units.ElectricCurrentGradientMilliamperePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliamperePerMinute = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricCurrentGradientDto{
        Value: 0,
        Unit:  units.ElectricCurrentGradientAmperePerSecond,
    }
    
    var zeroResult *units.ElectricCurrentGradient
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricCurrentGradientFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricCurrentGradientFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "AmperePerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "AmperePerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricCurrentGradientDto{
        Value: nanValue,
        Unit:  units.ElectricCurrentGradientAmperePerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with AmperePerSecond unit
    amperes_per_secondJSON := []byte(`{"value": 100, "unit": "AmperePerSecond"}`)
    amperes_per_secondResult, err := factory.FromDtoJSON(amperes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AmperePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_secondResult.Convert(units.ElectricCurrentGradientAmperePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with AmperePerMinute unit
    amperes_per_minuteJSON := []byte(`{"value": 100, "unit": "AmperePerMinute"}`)
    amperes_per_minuteResult, err := factory.FromDtoJSON(amperes_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AmperePerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_minuteResult.Convert(units.ElectricCurrentGradientAmperePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with AmperePerMillisecond unit
    amperes_per_millisecondJSON := []byte(`{"value": 100, "unit": "AmperePerMillisecond"}`)
    amperes_per_millisecondResult, err := factory.FromDtoJSON(amperes_per_millisecondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AmperePerMillisecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_millisecondResult.Convert(units.ElectricCurrentGradientAmperePerMillisecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerMillisecond = %v, want %v", converted, 100)
    }
    // Test JSON with AmperePerMicrosecond unit
    amperes_per_microsecondJSON := []byte(`{"value": 100, "unit": "AmperePerMicrosecond"}`)
    amperes_per_microsecondResult, err := factory.FromDtoJSON(amperes_per_microsecondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AmperePerMicrosecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_microsecondResult.Convert(units.ElectricCurrentGradientAmperePerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerMicrosecond = %v, want %v", converted, 100)
    }
    // Test JSON with AmperePerNanosecond unit
    amperes_per_nanosecondJSON := []byte(`{"value": 100, "unit": "AmperePerNanosecond"}`)
    amperes_per_nanosecondResult, err := factory.FromDtoJSON(amperes_per_nanosecondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AmperePerNanosecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_nanosecondResult.Convert(units.ElectricCurrentGradientAmperePerNanosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerNanosecond = %v, want %v", converted, 100)
    }
    // Test JSON with MilliamperePerSecond unit
    milliamperes_per_secondJSON := []byte(`{"value": 100, "unit": "MilliamperePerSecond"}`)
    milliamperes_per_secondResult, err := factory.FromDtoJSON(milliamperes_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliamperePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliamperes_per_secondResult.Convert(units.ElectricCurrentGradientMilliamperePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliamperePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MilliamperePerMinute unit
    milliamperes_per_minuteJSON := []byte(`{"value": 100, "unit": "MilliamperePerMinute"}`)
    milliamperes_per_minuteResult, err := factory.FromDtoJSON(milliamperes_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliamperePerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliamperes_per_minuteResult.Convert(units.ElectricCurrentGradientMilliamperePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliamperePerMinute = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "AmperePerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromAmperesPerSecond function
func TestElectricCurrentGradientFactory_FromAmperesPerSecond(t *testing.T) {
    factory := units.ElectricCurrentGradientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAmperesPerSecond(100)
    if err != nil {
        t.Errorf("FromAmperesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentGradientAmperePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAmperesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAmperesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromAmperesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromAmperesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromAmperesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromAmperesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromAmperesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAmperesPerSecond(0)
    if err != nil {
        t.Errorf("FromAmperesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentGradientAmperePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAmperesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromAmperesPerMinute function
func TestElectricCurrentGradientFactory_FromAmperesPerMinute(t *testing.T) {
    factory := units.ElectricCurrentGradientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAmperesPerMinute(100)
    if err != nil {
        t.Errorf("FromAmperesPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentGradientAmperePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAmperesPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAmperesPerMinute(math.NaN())
    if err == nil {
        t.Error("FromAmperesPerMinute() with NaN value should return error")
    }

    _, err = factory.FromAmperesPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromAmperesPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromAmperesPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromAmperesPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAmperesPerMinute(0)
    if err != nil {
        t.Errorf("FromAmperesPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentGradientAmperePerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAmperesPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromAmperesPerMillisecond function
func TestElectricCurrentGradientFactory_FromAmperesPerMillisecond(t *testing.T) {
    factory := units.ElectricCurrentGradientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAmperesPerMillisecond(100)
    if err != nil {
        t.Errorf("FromAmperesPerMillisecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentGradientAmperePerMillisecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAmperesPerMillisecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAmperesPerMillisecond(math.NaN())
    if err == nil {
        t.Error("FromAmperesPerMillisecond() with NaN value should return error")
    }

    _, err = factory.FromAmperesPerMillisecond(math.Inf(1))
    if err == nil {
        t.Error("FromAmperesPerMillisecond() with +Inf value should return error")
    }

    _, err = factory.FromAmperesPerMillisecond(math.Inf(-1))
    if err == nil {
        t.Error("FromAmperesPerMillisecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAmperesPerMillisecond(0)
    if err != nil {
        t.Errorf("FromAmperesPerMillisecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentGradientAmperePerMillisecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAmperesPerMillisecond() with zero value = %v, want 0", converted)
    }
}
// Test FromAmperesPerMicrosecond function
func TestElectricCurrentGradientFactory_FromAmperesPerMicrosecond(t *testing.T) {
    factory := units.ElectricCurrentGradientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAmperesPerMicrosecond(100)
    if err != nil {
        t.Errorf("FromAmperesPerMicrosecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentGradientAmperePerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAmperesPerMicrosecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAmperesPerMicrosecond(math.NaN())
    if err == nil {
        t.Error("FromAmperesPerMicrosecond() with NaN value should return error")
    }

    _, err = factory.FromAmperesPerMicrosecond(math.Inf(1))
    if err == nil {
        t.Error("FromAmperesPerMicrosecond() with +Inf value should return error")
    }

    _, err = factory.FromAmperesPerMicrosecond(math.Inf(-1))
    if err == nil {
        t.Error("FromAmperesPerMicrosecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAmperesPerMicrosecond(0)
    if err != nil {
        t.Errorf("FromAmperesPerMicrosecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentGradientAmperePerMicrosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAmperesPerMicrosecond() with zero value = %v, want 0", converted)
    }
}
// Test FromAmperesPerNanosecond function
func TestElectricCurrentGradientFactory_FromAmperesPerNanosecond(t *testing.T) {
    factory := units.ElectricCurrentGradientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAmperesPerNanosecond(100)
    if err != nil {
        t.Errorf("FromAmperesPerNanosecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentGradientAmperePerNanosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAmperesPerNanosecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAmperesPerNanosecond(math.NaN())
    if err == nil {
        t.Error("FromAmperesPerNanosecond() with NaN value should return error")
    }

    _, err = factory.FromAmperesPerNanosecond(math.Inf(1))
    if err == nil {
        t.Error("FromAmperesPerNanosecond() with +Inf value should return error")
    }

    _, err = factory.FromAmperesPerNanosecond(math.Inf(-1))
    if err == nil {
        t.Error("FromAmperesPerNanosecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAmperesPerNanosecond(0)
    if err != nil {
        t.Errorf("FromAmperesPerNanosecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentGradientAmperePerNanosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAmperesPerNanosecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliamperesPerSecond function
func TestElectricCurrentGradientFactory_FromMilliamperesPerSecond(t *testing.T) {
    factory := units.ElectricCurrentGradientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliamperesPerSecond(100)
    if err != nil {
        t.Errorf("FromMilliamperesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentGradientMilliamperePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliamperesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliamperesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMilliamperesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMilliamperesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMilliamperesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMilliamperesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliamperesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliamperesPerSecond(0)
    if err != nil {
        t.Errorf("FromMilliamperesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentGradientMilliamperePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliamperesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliamperesPerMinute function
func TestElectricCurrentGradientFactory_FromMilliamperesPerMinute(t *testing.T) {
    factory := units.ElectricCurrentGradientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliamperesPerMinute(100)
    if err != nil {
        t.Errorf("FromMilliamperesPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentGradientMilliamperePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliamperesPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliamperesPerMinute(math.NaN())
    if err == nil {
        t.Error("FromMilliamperesPerMinute() with NaN value should return error")
    }

    _, err = factory.FromMilliamperesPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromMilliamperesPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromMilliamperesPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliamperesPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliamperesPerMinute(0)
    if err != nil {
        t.Errorf("FromMilliamperesPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentGradientMilliamperePerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliamperesPerMinute() with zero value = %v, want 0", converted)
    }
}

func TestElectricCurrentGradientToString(t *testing.T) {
	factory := units.ElectricCurrentGradientFactory{}
	a, err := factory.CreateElectricCurrentGradient(45, units.ElectricCurrentGradientAmperePerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricCurrentGradientAmperePerSecond, 2)
	expected := "45.00 " + units.GetElectricCurrentGradientAbbreviation(units.ElectricCurrentGradientAmperePerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricCurrentGradientAmperePerSecond, -1)
	expected = "45 " + units.GetElectricCurrentGradientAbbreviation(units.ElectricCurrentGradientAmperePerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricCurrentGradient_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricCurrentGradientFactory{}
	a1, _ := factory.CreateElectricCurrentGradient(60, units.ElectricCurrentGradientAmperePerSecond)
	a2, _ := factory.CreateElectricCurrentGradient(60, units.ElectricCurrentGradientAmperePerSecond)
	a3, _ := factory.CreateElectricCurrentGradient(120, units.ElectricCurrentGradientAmperePerSecond)

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

func TestElectricCurrentGradient_Arithmetic(t *testing.T) {
	factory := units.ElectricCurrentGradientFactory{}
	a1, _ := factory.CreateElectricCurrentGradient(30, units.ElectricCurrentGradientAmperePerSecond)
	a2, _ := factory.CreateElectricCurrentGradient(45, units.ElectricCurrentGradientAmperePerSecond)

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


func TestGetElectricCurrentGradientAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricCurrentGradientUnits
        want string
    }{
        {
            name: "AmperePerSecond abbreviation",
            unit: units.ElectricCurrentGradientAmperePerSecond,
            want: "A/s",
        },
        {
            name: "AmperePerMinute abbreviation",
            unit: units.ElectricCurrentGradientAmperePerMinute,
            want: "A/min",
        },
        {
            name: "AmperePerMillisecond abbreviation",
            unit: units.ElectricCurrentGradientAmperePerMillisecond,
            want: "A/ms",
        },
        {
            name: "AmperePerMicrosecond abbreviation",
            unit: units.ElectricCurrentGradientAmperePerMicrosecond,
            want: "A/μs",
        },
        {
            name: "AmperePerNanosecond abbreviation",
            unit: units.ElectricCurrentGradientAmperePerNanosecond,
            want: "A/ns",
        },
        {
            name: "MilliamperePerSecond abbreviation",
            unit: units.ElectricCurrentGradientMilliamperePerSecond,
            want: "mA/s",
        },
        {
            name: "MilliamperePerMinute abbreviation",
            unit: units.ElectricCurrentGradientMilliamperePerMinute,
            want: "mA/min",
        },
        {
            name: "invalid unit",
            unit: units.ElectricCurrentGradientUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricCurrentGradientAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricCurrentGradientAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricCurrentGradient_String(t *testing.T) {
    factory := units.ElectricCurrentGradientFactory{}
    
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
            unit, err := factory.CreateElectricCurrentGradient(tt.value, units.ElectricCurrentGradientAmperePerSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricCurrentGradient.String() = %v, want %v", got, tt.want)
            }
        })
    }
}