// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricPotentialChangeRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltPerSecond"}`
	
	factory := units.ElectricPotentialChangeRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricPotentialChangeRateVoltPerSecond {
		t.Errorf("expected unit %v, got %v", units.ElectricPotentialChangeRateVoltPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricPotentialChangeRateDto_ToJSON(t *testing.T) {
	dto := units.ElectricPotentialChangeRateDto{
		Value: 45,
		Unit:  units.ElectricPotentialChangeRateVoltPerSecond,
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
	if result["unit"].(string) != string(units.ElectricPotentialChangeRateVoltPerSecond) {
		t.Errorf("expected unit %s, got %v", units.ElectricPotentialChangeRateVoltPerSecond, result["unit"])
	}
}

func TestNewElectricPotentialChangeRate_InvalidValue(t *testing.T) {
	factory := units.ElectricPotentialChangeRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricPotentialChangeRate(math.NaN(), units.ElectricPotentialChangeRateVoltPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricPotentialChangeRate(math.Inf(1), units.ElectricPotentialChangeRateVoltPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricPotentialChangeRateConversions(t *testing.T) {
	factory := units.ElectricPotentialChangeRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricPotentialChangeRate(180, units.ElectricPotentialChangeRateVoltPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltsPerSeconds.
		// No expected conversion value provided for VoltsPerSeconds, verifying result is not NaN.
		result := a.VoltsPerSeconds()
		cacheResult := a.VoltsPerSeconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltsPerSeconds returned NaN")
		}
	}
	{
		// Test conversion to VoltsPerMicroseconds.
		// No expected conversion value provided for VoltsPerMicroseconds, verifying result is not NaN.
		result := a.VoltsPerMicroseconds()
		cacheResult := a.VoltsPerMicroseconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltsPerMicroseconds returned NaN")
		}
	}
	{
		// Test conversion to VoltsPerMinutes.
		// No expected conversion value provided for VoltsPerMinutes, verifying result is not NaN.
		result := a.VoltsPerMinutes()
		cacheResult := a.VoltsPerMinutes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltsPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to VoltsPerHours.
		// No expected conversion value provided for VoltsPerHours, verifying result is not NaN.
		result := a.VoltsPerHours()
		cacheResult := a.VoltsPerHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltsPerHours returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsPerSeconds.
		// No expected conversion value provided for MicrovoltsPerSeconds, verifying result is not NaN.
		result := a.MicrovoltsPerSeconds()
		cacheResult := a.MicrovoltsPerSeconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrovoltsPerSeconds returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsPerSeconds.
		// No expected conversion value provided for MillivoltsPerSeconds, verifying result is not NaN.
		result := a.MillivoltsPerSeconds()
		cacheResult := a.MillivoltsPerSeconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillivoltsPerSeconds returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsPerSeconds.
		// No expected conversion value provided for KilovoltsPerSeconds, verifying result is not NaN.
		result := a.KilovoltsPerSeconds()
		cacheResult := a.KilovoltsPerSeconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilovoltsPerSeconds returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsPerSeconds.
		// No expected conversion value provided for MegavoltsPerSeconds, verifying result is not NaN.
		result := a.MegavoltsPerSeconds()
		cacheResult := a.MegavoltsPerSeconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegavoltsPerSeconds returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsPerMicroseconds.
		// No expected conversion value provided for MicrovoltsPerMicroseconds, verifying result is not NaN.
		result := a.MicrovoltsPerMicroseconds()
		cacheResult := a.MicrovoltsPerMicroseconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrovoltsPerMicroseconds returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsPerMicroseconds.
		// No expected conversion value provided for MillivoltsPerMicroseconds, verifying result is not NaN.
		result := a.MillivoltsPerMicroseconds()
		cacheResult := a.MillivoltsPerMicroseconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillivoltsPerMicroseconds returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsPerMicroseconds.
		// No expected conversion value provided for KilovoltsPerMicroseconds, verifying result is not NaN.
		result := a.KilovoltsPerMicroseconds()
		cacheResult := a.KilovoltsPerMicroseconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilovoltsPerMicroseconds returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsPerMicroseconds.
		// No expected conversion value provided for MegavoltsPerMicroseconds, verifying result is not NaN.
		result := a.MegavoltsPerMicroseconds()
		cacheResult := a.MegavoltsPerMicroseconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegavoltsPerMicroseconds returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsPerMinutes.
		// No expected conversion value provided for MicrovoltsPerMinutes, verifying result is not NaN.
		result := a.MicrovoltsPerMinutes()
		cacheResult := a.MicrovoltsPerMinutes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrovoltsPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsPerMinutes.
		// No expected conversion value provided for MillivoltsPerMinutes, verifying result is not NaN.
		result := a.MillivoltsPerMinutes()
		cacheResult := a.MillivoltsPerMinutes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillivoltsPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsPerMinutes.
		// No expected conversion value provided for KilovoltsPerMinutes, verifying result is not NaN.
		result := a.KilovoltsPerMinutes()
		cacheResult := a.KilovoltsPerMinutes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilovoltsPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsPerMinutes.
		// No expected conversion value provided for MegavoltsPerMinutes, verifying result is not NaN.
		result := a.MegavoltsPerMinutes()
		cacheResult := a.MegavoltsPerMinutes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegavoltsPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsPerHours.
		// No expected conversion value provided for MicrovoltsPerHours, verifying result is not NaN.
		result := a.MicrovoltsPerHours()
		cacheResult := a.MicrovoltsPerHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrovoltsPerHours returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsPerHours.
		// No expected conversion value provided for MillivoltsPerHours, verifying result is not NaN.
		result := a.MillivoltsPerHours()
		cacheResult := a.MillivoltsPerHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillivoltsPerHours returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsPerHours.
		// No expected conversion value provided for KilovoltsPerHours, verifying result is not NaN.
		result := a.KilovoltsPerHours()
		cacheResult := a.KilovoltsPerHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilovoltsPerHours returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsPerHours.
		// No expected conversion value provided for MegavoltsPerHours, verifying result is not NaN.
		result := a.MegavoltsPerHours()
		cacheResult := a.MegavoltsPerHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegavoltsPerHours returned NaN")
		}
	}
}

func TestElectricPotentialChangeRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricPotentialChangeRateFactory{}
	a, err := factory.CreateElectricPotentialChangeRate(90, units.ElectricPotentialChangeRateVoltPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricPotentialChangeRateVoltPerSecond {
		t.Errorf("expected default unit VoltPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricPotentialChangeRateVoltPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricPotentialChangeRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricPotentialChangeRateVoltPerSecond {
		t.Errorf("expected unit VoltPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricPotentialChangeRateFactory_FromDto(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateVoltPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricPotentialChangeRateDto{
        Value: math.NaN(),
        Unit:  units.ElectricPotentialChangeRateVoltPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test VoltPerSecond conversion
    volts_per_secondsDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateVoltPerSecond,
    }
    
    var volts_per_secondsResult *units.ElectricPotentialChangeRate
    volts_per_secondsResult, err = factory.FromDto(volts_per_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with VoltPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_secondsResult.Convert(units.ElectricPotentialChangeRateVoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerSecond = %v, want %v", converted, 100)
    }
    // Test VoltPerMicrosecond conversion
    volts_per_microsecondsDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateVoltPerMicrosecond,
    }
    
    var volts_per_microsecondsResult *units.ElectricPotentialChangeRate
    volts_per_microsecondsResult, err = factory.FromDto(volts_per_microsecondsDto)
    if err != nil {
        t.Errorf("FromDto() with VoltPerMicrosecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_microsecondsResult.Convert(units.ElectricPotentialChangeRateVoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test VoltPerMinute conversion
    volts_per_minutesDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateVoltPerMinute,
    }
    
    var volts_per_minutesResult *units.ElectricPotentialChangeRate
    volts_per_minutesResult, err = factory.FromDto(volts_per_minutesDto)
    if err != nil {
        t.Errorf("FromDto() with VoltPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_minutesResult.Convert(units.ElectricPotentialChangeRateVoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerMinute = %v, want %v", converted, 100)
    }
    // Test VoltPerHour conversion
    volts_per_hoursDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateVoltPerHour,
    }
    
    var volts_per_hoursResult *units.ElectricPotentialChangeRate
    volts_per_hoursResult, err = factory.FromDto(volts_per_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with VoltPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_hoursResult.Convert(units.ElectricPotentialChangeRateVoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerHour = %v, want %v", converted, 100)
    }
    // Test MicrovoltPerSecond conversion
    microvolts_per_secondsDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMicrovoltPerSecond,
    }
    
    var microvolts_per_secondsResult *units.ElectricPotentialChangeRate
    microvolts_per_secondsResult, err = factory.FromDto(microvolts_per_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with MicrovoltPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_secondsResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerSecond = %v, want %v", converted, 100)
    }
    // Test MillivoltPerSecond conversion
    millivolts_per_secondsDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMillivoltPerSecond,
    }
    
    var millivolts_per_secondsResult *units.ElectricPotentialChangeRate
    millivolts_per_secondsResult, err = factory.FromDto(millivolts_per_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with MillivoltPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_secondsResult.Convert(units.ElectricPotentialChangeRateMillivoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerSecond = %v, want %v", converted, 100)
    }
    // Test KilovoltPerSecond conversion
    kilovolts_per_secondsDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateKilovoltPerSecond,
    }
    
    var kilovolts_per_secondsResult *units.ElectricPotentialChangeRate
    kilovolts_per_secondsResult, err = factory.FromDto(kilovolts_per_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_secondsResult.Convert(units.ElectricPotentialChangeRateKilovoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerSecond = %v, want %v", converted, 100)
    }
    // Test MegavoltPerSecond conversion
    megavolts_per_secondsDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMegavoltPerSecond,
    }
    
    var megavolts_per_secondsResult *units.ElectricPotentialChangeRate
    megavolts_per_secondsResult, err = factory.FromDto(megavolts_per_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_secondsResult.Convert(units.ElectricPotentialChangeRateMegavoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerSecond = %v, want %v", converted, 100)
    }
    // Test MicrovoltPerMicrosecond conversion
    microvolts_per_microsecondsDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMicrovoltPerMicrosecond,
    }
    
    var microvolts_per_microsecondsResult *units.ElectricPotentialChangeRate
    microvolts_per_microsecondsResult, err = factory.FromDto(microvolts_per_microsecondsDto)
    if err != nil {
        t.Errorf("FromDto() with MicrovoltPerMicrosecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_microsecondsResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test MillivoltPerMicrosecond conversion
    millivolts_per_microsecondsDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMillivoltPerMicrosecond,
    }
    
    var millivolts_per_microsecondsResult *units.ElectricPotentialChangeRate
    millivolts_per_microsecondsResult, err = factory.FromDto(millivolts_per_microsecondsDto)
    if err != nil {
        t.Errorf("FromDto() with MillivoltPerMicrosecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_microsecondsResult.Convert(units.ElectricPotentialChangeRateMillivoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test KilovoltPerMicrosecond conversion
    kilovolts_per_microsecondsDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateKilovoltPerMicrosecond,
    }
    
    var kilovolts_per_microsecondsResult *units.ElectricPotentialChangeRate
    kilovolts_per_microsecondsResult, err = factory.FromDto(kilovolts_per_microsecondsDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltPerMicrosecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_microsecondsResult.Convert(units.ElectricPotentialChangeRateKilovoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test MegavoltPerMicrosecond conversion
    megavolts_per_microsecondsDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMegavoltPerMicrosecond,
    }
    
    var megavolts_per_microsecondsResult *units.ElectricPotentialChangeRate
    megavolts_per_microsecondsResult, err = factory.FromDto(megavolts_per_microsecondsDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltPerMicrosecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_microsecondsResult.Convert(units.ElectricPotentialChangeRateMegavoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test MicrovoltPerMinute conversion
    microvolts_per_minutesDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMicrovoltPerMinute,
    }
    
    var microvolts_per_minutesResult *units.ElectricPotentialChangeRate
    microvolts_per_minutesResult, err = factory.FromDto(microvolts_per_minutesDto)
    if err != nil {
        t.Errorf("FromDto() with MicrovoltPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_minutesResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerMinute = %v, want %v", converted, 100)
    }
    // Test MillivoltPerMinute conversion
    millivolts_per_minutesDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMillivoltPerMinute,
    }
    
    var millivolts_per_minutesResult *units.ElectricPotentialChangeRate
    millivolts_per_minutesResult, err = factory.FromDto(millivolts_per_minutesDto)
    if err != nil {
        t.Errorf("FromDto() with MillivoltPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_minutesResult.Convert(units.ElectricPotentialChangeRateMillivoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerMinute = %v, want %v", converted, 100)
    }
    // Test KilovoltPerMinute conversion
    kilovolts_per_minutesDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateKilovoltPerMinute,
    }
    
    var kilovolts_per_minutesResult *units.ElectricPotentialChangeRate
    kilovolts_per_minutesResult, err = factory.FromDto(kilovolts_per_minutesDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_minutesResult.Convert(units.ElectricPotentialChangeRateKilovoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerMinute = %v, want %v", converted, 100)
    }
    // Test MegavoltPerMinute conversion
    megavolts_per_minutesDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMegavoltPerMinute,
    }
    
    var megavolts_per_minutesResult *units.ElectricPotentialChangeRate
    megavolts_per_minutesResult, err = factory.FromDto(megavolts_per_minutesDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_minutesResult.Convert(units.ElectricPotentialChangeRateMegavoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerMinute = %v, want %v", converted, 100)
    }
    // Test MicrovoltPerHour conversion
    microvolts_per_hoursDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMicrovoltPerHour,
    }
    
    var microvolts_per_hoursResult *units.ElectricPotentialChangeRate
    microvolts_per_hoursResult, err = factory.FromDto(microvolts_per_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with MicrovoltPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_hoursResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerHour = %v, want %v", converted, 100)
    }
    // Test MillivoltPerHour conversion
    millivolts_per_hoursDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMillivoltPerHour,
    }
    
    var millivolts_per_hoursResult *units.ElectricPotentialChangeRate
    millivolts_per_hoursResult, err = factory.FromDto(millivolts_per_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with MillivoltPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_hoursResult.Convert(units.ElectricPotentialChangeRateMillivoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerHour = %v, want %v", converted, 100)
    }
    // Test KilovoltPerHour conversion
    kilovolts_per_hoursDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateKilovoltPerHour,
    }
    
    var kilovolts_per_hoursResult *units.ElectricPotentialChangeRate
    kilovolts_per_hoursResult, err = factory.FromDto(kilovolts_per_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_hoursResult.Convert(units.ElectricPotentialChangeRateKilovoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerHour = %v, want %v", converted, 100)
    }
    // Test MegavoltPerHour conversion
    megavolts_per_hoursDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMegavoltPerHour,
    }
    
    var megavolts_per_hoursResult *units.ElectricPotentialChangeRate
    megavolts_per_hoursResult, err = factory.FromDto(megavolts_per_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_hoursResult.Convert(units.ElectricPotentialChangeRateMegavoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerHour = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricPotentialChangeRateDto{
        Value: 0,
        Unit:  units.ElectricPotentialChangeRateVoltPerSecond,
    }
    
    var zeroResult *units.ElectricPotentialChangeRate
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricPotentialChangeRateFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "VoltPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "VoltPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricPotentialChangeRateDto{
        Value: nanValue,
        Unit:  units.ElectricPotentialChangeRateVoltPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with VoltPerSecond unit
    volts_per_secondsJSON := []byte(`{"value": 100, "unit": "VoltPerSecond"}`)
    volts_per_secondsResult, err := factory.FromDtoJSON(volts_per_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_secondsResult.Convert(units.ElectricPotentialChangeRateVoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with VoltPerMicrosecond unit
    volts_per_microsecondsJSON := []byte(`{"value": 100, "unit": "VoltPerMicrosecond"}`)
    volts_per_microsecondsResult, err := factory.FromDtoJSON(volts_per_microsecondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltPerMicrosecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_microsecondsResult.Convert(units.ElectricPotentialChangeRateVoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test JSON with VoltPerMinute unit
    volts_per_minutesJSON := []byte(`{"value": 100, "unit": "VoltPerMinute"}`)
    volts_per_minutesResult, err := factory.FromDtoJSON(volts_per_minutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_minutesResult.Convert(units.ElectricPotentialChangeRateVoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with VoltPerHour unit
    volts_per_hoursJSON := []byte(`{"value": 100, "unit": "VoltPerHour"}`)
    volts_per_hoursResult, err := factory.FromDtoJSON(volts_per_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_hoursResult.Convert(units.ElectricPotentialChangeRateVoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MicrovoltPerSecond unit
    microvolts_per_secondsJSON := []byte(`{"value": 100, "unit": "MicrovoltPerSecond"}`)
    microvolts_per_secondsResult, err := factory.FromDtoJSON(microvolts_per_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrovoltPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_secondsResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MillivoltPerSecond unit
    millivolts_per_secondsJSON := []byte(`{"value": 100, "unit": "MillivoltPerSecond"}`)
    millivolts_per_secondsResult, err := factory.FromDtoJSON(millivolts_per_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillivoltPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_secondsResult.Convert(units.ElectricPotentialChangeRateMillivoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltPerSecond unit
    kilovolts_per_secondsJSON := []byte(`{"value": 100, "unit": "KilovoltPerSecond"}`)
    kilovolts_per_secondsResult, err := factory.FromDtoJSON(kilovolts_per_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_secondsResult.Convert(units.ElectricPotentialChangeRateKilovoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltPerSecond unit
    megavolts_per_secondsJSON := []byte(`{"value": 100, "unit": "MegavoltPerSecond"}`)
    megavolts_per_secondsResult, err := factory.FromDtoJSON(megavolts_per_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_secondsResult.Convert(units.ElectricPotentialChangeRateMegavoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicrovoltPerMicrosecond unit
    microvolts_per_microsecondsJSON := []byte(`{"value": 100, "unit": "MicrovoltPerMicrosecond"}`)
    microvolts_per_microsecondsResult, err := factory.FromDtoJSON(microvolts_per_microsecondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrovoltPerMicrosecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_microsecondsResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test JSON with MillivoltPerMicrosecond unit
    millivolts_per_microsecondsJSON := []byte(`{"value": 100, "unit": "MillivoltPerMicrosecond"}`)
    millivolts_per_microsecondsResult, err := factory.FromDtoJSON(millivolts_per_microsecondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillivoltPerMicrosecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_microsecondsResult.Convert(units.ElectricPotentialChangeRateMillivoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltPerMicrosecond unit
    kilovolts_per_microsecondsJSON := []byte(`{"value": 100, "unit": "KilovoltPerMicrosecond"}`)
    kilovolts_per_microsecondsResult, err := factory.FromDtoJSON(kilovolts_per_microsecondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltPerMicrosecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_microsecondsResult.Convert(units.ElectricPotentialChangeRateKilovoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltPerMicrosecond unit
    megavolts_per_microsecondsJSON := []byte(`{"value": 100, "unit": "MegavoltPerMicrosecond"}`)
    megavolts_per_microsecondsResult, err := factory.FromDtoJSON(megavolts_per_microsecondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltPerMicrosecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_microsecondsResult.Convert(units.ElectricPotentialChangeRateMegavoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicrovoltPerMinute unit
    microvolts_per_minutesJSON := []byte(`{"value": 100, "unit": "MicrovoltPerMinute"}`)
    microvolts_per_minutesResult, err := factory.FromDtoJSON(microvolts_per_minutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrovoltPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_minutesResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MillivoltPerMinute unit
    millivolts_per_minutesJSON := []byte(`{"value": 100, "unit": "MillivoltPerMinute"}`)
    millivolts_per_minutesResult, err := factory.FromDtoJSON(millivolts_per_minutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillivoltPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_minutesResult.Convert(units.ElectricPotentialChangeRateMillivoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltPerMinute unit
    kilovolts_per_minutesJSON := []byte(`{"value": 100, "unit": "KilovoltPerMinute"}`)
    kilovolts_per_minutesResult, err := factory.FromDtoJSON(kilovolts_per_minutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_minutesResult.Convert(units.ElectricPotentialChangeRateKilovoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltPerMinute unit
    megavolts_per_minutesJSON := []byte(`{"value": 100, "unit": "MegavoltPerMinute"}`)
    megavolts_per_minutesResult, err := factory.FromDtoJSON(megavolts_per_minutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_minutesResult.Convert(units.ElectricPotentialChangeRateMegavoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MicrovoltPerHour unit
    microvolts_per_hoursJSON := []byte(`{"value": 100, "unit": "MicrovoltPerHour"}`)
    microvolts_per_hoursResult, err := factory.FromDtoJSON(microvolts_per_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrovoltPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_hoursResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MillivoltPerHour unit
    millivolts_per_hoursJSON := []byte(`{"value": 100, "unit": "MillivoltPerHour"}`)
    millivolts_per_hoursResult, err := factory.FromDtoJSON(millivolts_per_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillivoltPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_hoursResult.Convert(units.ElectricPotentialChangeRateMillivoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltPerHour unit
    kilovolts_per_hoursJSON := []byte(`{"value": 100, "unit": "KilovoltPerHour"}`)
    kilovolts_per_hoursResult, err := factory.FromDtoJSON(kilovolts_per_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_hoursResult.Convert(units.ElectricPotentialChangeRateKilovoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltPerHour unit
    megavolts_per_hoursJSON := []byte(`{"value": 100, "unit": "MegavoltPerHour"}`)
    megavolts_per_hoursResult, err := factory.FromDtoJSON(megavolts_per_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_hoursResult.Convert(units.ElectricPotentialChangeRateMegavoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerHour = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "VoltPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromVoltsPerSeconds function
func TestElectricPotentialChangeRateFactory_FromVoltsPerSeconds(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltsPerSeconds(100)
    if err != nil {
        t.Errorf("FromVoltsPerSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateVoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltsPerSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltsPerSeconds(math.NaN())
    if err == nil {
        t.Error("FromVoltsPerSeconds() with NaN value should return error")
    }

    _, err = factory.FromVoltsPerSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromVoltsPerSeconds() with +Inf value should return error")
    }

    _, err = factory.FromVoltsPerSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltsPerSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltsPerSeconds(0)
    if err != nil {
        t.Errorf("FromVoltsPerSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateVoltPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltsPerSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromVoltsPerMicroseconds function
func TestElectricPotentialChangeRateFactory_FromVoltsPerMicroseconds(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltsPerMicroseconds(100)
    if err != nil {
        t.Errorf("FromVoltsPerMicroseconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateVoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltsPerMicroseconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltsPerMicroseconds(math.NaN())
    if err == nil {
        t.Error("FromVoltsPerMicroseconds() with NaN value should return error")
    }

    _, err = factory.FromVoltsPerMicroseconds(math.Inf(1))
    if err == nil {
        t.Error("FromVoltsPerMicroseconds() with +Inf value should return error")
    }

    _, err = factory.FromVoltsPerMicroseconds(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltsPerMicroseconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltsPerMicroseconds(0)
    if err != nil {
        t.Errorf("FromVoltsPerMicroseconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateVoltPerMicrosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltsPerMicroseconds() with zero value = %v, want 0", converted)
    }
}
// Test FromVoltsPerMinutes function
func TestElectricPotentialChangeRateFactory_FromVoltsPerMinutes(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltsPerMinutes(100)
    if err != nil {
        t.Errorf("FromVoltsPerMinutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateVoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltsPerMinutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltsPerMinutes(math.NaN())
    if err == nil {
        t.Error("FromVoltsPerMinutes() with NaN value should return error")
    }

    _, err = factory.FromVoltsPerMinutes(math.Inf(1))
    if err == nil {
        t.Error("FromVoltsPerMinutes() with +Inf value should return error")
    }

    _, err = factory.FromVoltsPerMinutes(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltsPerMinutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltsPerMinutes(0)
    if err != nil {
        t.Errorf("FromVoltsPerMinutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateVoltPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltsPerMinutes() with zero value = %v, want 0", converted)
    }
}
// Test FromVoltsPerHours function
func TestElectricPotentialChangeRateFactory_FromVoltsPerHours(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltsPerHours(100)
    if err != nil {
        t.Errorf("FromVoltsPerHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateVoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltsPerHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltsPerHours(math.NaN())
    if err == nil {
        t.Error("FromVoltsPerHours() with NaN value should return error")
    }

    _, err = factory.FromVoltsPerHours(math.Inf(1))
    if err == nil {
        t.Error("FromVoltsPerHours() with +Inf value should return error")
    }

    _, err = factory.FromVoltsPerHours(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltsPerHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltsPerHours(0)
    if err != nil {
        t.Errorf("FromVoltsPerHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateVoltPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltsPerHours() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrovoltsPerSeconds function
func TestElectricPotentialChangeRateFactory_FromMicrovoltsPerSeconds(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrovoltsPerSeconds(100)
    if err != nil {
        t.Errorf("FromMicrovoltsPerSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMicrovoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrovoltsPerSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrovoltsPerSeconds(math.NaN())
    if err == nil {
        t.Error("FromMicrovoltsPerSeconds() with NaN value should return error")
    }

    _, err = factory.FromMicrovoltsPerSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromMicrovoltsPerSeconds() with +Inf value should return error")
    }

    _, err = factory.FromMicrovoltsPerSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrovoltsPerSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrovoltsPerSeconds(0)
    if err != nil {
        t.Errorf("FromMicrovoltsPerSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrovoltsPerSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromMillivoltsPerSeconds function
func TestElectricPotentialChangeRateFactory_FromMillivoltsPerSeconds(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillivoltsPerSeconds(100)
    if err != nil {
        t.Errorf("FromMillivoltsPerSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMillivoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillivoltsPerSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillivoltsPerSeconds(math.NaN())
    if err == nil {
        t.Error("FromMillivoltsPerSeconds() with NaN value should return error")
    }

    _, err = factory.FromMillivoltsPerSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromMillivoltsPerSeconds() with +Inf value should return error")
    }

    _, err = factory.FromMillivoltsPerSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromMillivoltsPerSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillivoltsPerSeconds(0)
    if err != nil {
        t.Errorf("FromMillivoltsPerSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMillivoltPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillivoltsPerSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltsPerSeconds function
func TestElectricPotentialChangeRateFactory_FromKilovoltsPerSeconds(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltsPerSeconds(100)
    if err != nil {
        t.Errorf("FromKilovoltsPerSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateKilovoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltsPerSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltsPerSeconds(math.NaN())
    if err == nil {
        t.Error("FromKilovoltsPerSeconds() with NaN value should return error")
    }

    _, err = factory.FromKilovoltsPerSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltsPerSeconds() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltsPerSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltsPerSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltsPerSeconds(0)
    if err != nil {
        t.Errorf("FromKilovoltsPerSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateKilovoltPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltsPerSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltsPerSeconds function
func TestElectricPotentialChangeRateFactory_FromMegavoltsPerSeconds(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltsPerSeconds(100)
    if err != nil {
        t.Errorf("FromMegavoltsPerSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMegavoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltsPerSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltsPerSeconds(math.NaN())
    if err == nil {
        t.Error("FromMegavoltsPerSeconds() with NaN value should return error")
    }

    _, err = factory.FromMegavoltsPerSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltsPerSeconds() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltsPerSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltsPerSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltsPerSeconds(0)
    if err != nil {
        t.Errorf("FromMegavoltsPerSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMegavoltPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltsPerSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrovoltsPerMicroseconds function
func TestElectricPotentialChangeRateFactory_FromMicrovoltsPerMicroseconds(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrovoltsPerMicroseconds(100)
    if err != nil {
        t.Errorf("FromMicrovoltsPerMicroseconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMicrovoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrovoltsPerMicroseconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrovoltsPerMicroseconds(math.NaN())
    if err == nil {
        t.Error("FromMicrovoltsPerMicroseconds() with NaN value should return error")
    }

    _, err = factory.FromMicrovoltsPerMicroseconds(math.Inf(1))
    if err == nil {
        t.Error("FromMicrovoltsPerMicroseconds() with +Inf value should return error")
    }

    _, err = factory.FromMicrovoltsPerMicroseconds(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrovoltsPerMicroseconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrovoltsPerMicroseconds(0)
    if err != nil {
        t.Errorf("FromMicrovoltsPerMicroseconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerMicrosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrovoltsPerMicroseconds() with zero value = %v, want 0", converted)
    }
}
// Test FromMillivoltsPerMicroseconds function
func TestElectricPotentialChangeRateFactory_FromMillivoltsPerMicroseconds(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillivoltsPerMicroseconds(100)
    if err != nil {
        t.Errorf("FromMillivoltsPerMicroseconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMillivoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillivoltsPerMicroseconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillivoltsPerMicroseconds(math.NaN())
    if err == nil {
        t.Error("FromMillivoltsPerMicroseconds() with NaN value should return error")
    }

    _, err = factory.FromMillivoltsPerMicroseconds(math.Inf(1))
    if err == nil {
        t.Error("FromMillivoltsPerMicroseconds() with +Inf value should return error")
    }

    _, err = factory.FromMillivoltsPerMicroseconds(math.Inf(-1))
    if err == nil {
        t.Error("FromMillivoltsPerMicroseconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillivoltsPerMicroseconds(0)
    if err != nil {
        t.Errorf("FromMillivoltsPerMicroseconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMillivoltPerMicrosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillivoltsPerMicroseconds() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltsPerMicroseconds function
func TestElectricPotentialChangeRateFactory_FromKilovoltsPerMicroseconds(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltsPerMicroseconds(100)
    if err != nil {
        t.Errorf("FromKilovoltsPerMicroseconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateKilovoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltsPerMicroseconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltsPerMicroseconds(math.NaN())
    if err == nil {
        t.Error("FromKilovoltsPerMicroseconds() with NaN value should return error")
    }

    _, err = factory.FromKilovoltsPerMicroseconds(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltsPerMicroseconds() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltsPerMicroseconds(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltsPerMicroseconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltsPerMicroseconds(0)
    if err != nil {
        t.Errorf("FromKilovoltsPerMicroseconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateKilovoltPerMicrosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltsPerMicroseconds() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltsPerMicroseconds function
func TestElectricPotentialChangeRateFactory_FromMegavoltsPerMicroseconds(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltsPerMicroseconds(100)
    if err != nil {
        t.Errorf("FromMegavoltsPerMicroseconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMegavoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltsPerMicroseconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltsPerMicroseconds(math.NaN())
    if err == nil {
        t.Error("FromMegavoltsPerMicroseconds() with NaN value should return error")
    }

    _, err = factory.FromMegavoltsPerMicroseconds(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltsPerMicroseconds() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltsPerMicroseconds(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltsPerMicroseconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltsPerMicroseconds(0)
    if err != nil {
        t.Errorf("FromMegavoltsPerMicroseconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMegavoltPerMicrosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltsPerMicroseconds() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrovoltsPerMinutes function
func TestElectricPotentialChangeRateFactory_FromMicrovoltsPerMinutes(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrovoltsPerMinutes(100)
    if err != nil {
        t.Errorf("FromMicrovoltsPerMinutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMicrovoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrovoltsPerMinutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrovoltsPerMinutes(math.NaN())
    if err == nil {
        t.Error("FromMicrovoltsPerMinutes() with NaN value should return error")
    }

    _, err = factory.FromMicrovoltsPerMinutes(math.Inf(1))
    if err == nil {
        t.Error("FromMicrovoltsPerMinutes() with +Inf value should return error")
    }

    _, err = factory.FromMicrovoltsPerMinutes(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrovoltsPerMinutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrovoltsPerMinutes(0)
    if err != nil {
        t.Errorf("FromMicrovoltsPerMinutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrovoltsPerMinutes() with zero value = %v, want 0", converted)
    }
}
// Test FromMillivoltsPerMinutes function
func TestElectricPotentialChangeRateFactory_FromMillivoltsPerMinutes(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillivoltsPerMinutes(100)
    if err != nil {
        t.Errorf("FromMillivoltsPerMinutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMillivoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillivoltsPerMinutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillivoltsPerMinutes(math.NaN())
    if err == nil {
        t.Error("FromMillivoltsPerMinutes() with NaN value should return error")
    }

    _, err = factory.FromMillivoltsPerMinutes(math.Inf(1))
    if err == nil {
        t.Error("FromMillivoltsPerMinutes() with +Inf value should return error")
    }

    _, err = factory.FromMillivoltsPerMinutes(math.Inf(-1))
    if err == nil {
        t.Error("FromMillivoltsPerMinutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillivoltsPerMinutes(0)
    if err != nil {
        t.Errorf("FromMillivoltsPerMinutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMillivoltPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillivoltsPerMinutes() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltsPerMinutes function
func TestElectricPotentialChangeRateFactory_FromKilovoltsPerMinutes(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltsPerMinutes(100)
    if err != nil {
        t.Errorf("FromKilovoltsPerMinutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateKilovoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltsPerMinutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltsPerMinutes(math.NaN())
    if err == nil {
        t.Error("FromKilovoltsPerMinutes() with NaN value should return error")
    }

    _, err = factory.FromKilovoltsPerMinutes(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltsPerMinutes() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltsPerMinutes(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltsPerMinutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltsPerMinutes(0)
    if err != nil {
        t.Errorf("FromKilovoltsPerMinutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateKilovoltPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltsPerMinutes() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltsPerMinutes function
func TestElectricPotentialChangeRateFactory_FromMegavoltsPerMinutes(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltsPerMinutes(100)
    if err != nil {
        t.Errorf("FromMegavoltsPerMinutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMegavoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltsPerMinutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltsPerMinutes(math.NaN())
    if err == nil {
        t.Error("FromMegavoltsPerMinutes() with NaN value should return error")
    }

    _, err = factory.FromMegavoltsPerMinutes(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltsPerMinutes() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltsPerMinutes(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltsPerMinutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltsPerMinutes(0)
    if err != nil {
        t.Errorf("FromMegavoltsPerMinutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMegavoltPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltsPerMinutes() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrovoltsPerHours function
func TestElectricPotentialChangeRateFactory_FromMicrovoltsPerHours(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrovoltsPerHours(100)
    if err != nil {
        t.Errorf("FromMicrovoltsPerHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMicrovoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrovoltsPerHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrovoltsPerHours(math.NaN())
    if err == nil {
        t.Error("FromMicrovoltsPerHours() with NaN value should return error")
    }

    _, err = factory.FromMicrovoltsPerHours(math.Inf(1))
    if err == nil {
        t.Error("FromMicrovoltsPerHours() with +Inf value should return error")
    }

    _, err = factory.FromMicrovoltsPerHours(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrovoltsPerHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrovoltsPerHours(0)
    if err != nil {
        t.Errorf("FromMicrovoltsPerHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrovoltsPerHours() with zero value = %v, want 0", converted)
    }
}
// Test FromMillivoltsPerHours function
func TestElectricPotentialChangeRateFactory_FromMillivoltsPerHours(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillivoltsPerHours(100)
    if err != nil {
        t.Errorf("FromMillivoltsPerHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMillivoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillivoltsPerHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillivoltsPerHours(math.NaN())
    if err == nil {
        t.Error("FromMillivoltsPerHours() with NaN value should return error")
    }

    _, err = factory.FromMillivoltsPerHours(math.Inf(1))
    if err == nil {
        t.Error("FromMillivoltsPerHours() with +Inf value should return error")
    }

    _, err = factory.FromMillivoltsPerHours(math.Inf(-1))
    if err == nil {
        t.Error("FromMillivoltsPerHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillivoltsPerHours(0)
    if err != nil {
        t.Errorf("FromMillivoltsPerHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMillivoltPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillivoltsPerHours() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltsPerHours function
func TestElectricPotentialChangeRateFactory_FromKilovoltsPerHours(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltsPerHours(100)
    if err != nil {
        t.Errorf("FromKilovoltsPerHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateKilovoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltsPerHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltsPerHours(math.NaN())
    if err == nil {
        t.Error("FromKilovoltsPerHours() with NaN value should return error")
    }

    _, err = factory.FromKilovoltsPerHours(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltsPerHours() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltsPerHours(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltsPerHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltsPerHours(0)
    if err != nil {
        t.Errorf("FromKilovoltsPerHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateKilovoltPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltsPerHours() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltsPerHours function
func TestElectricPotentialChangeRateFactory_FromMegavoltsPerHours(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltsPerHours(100)
    if err != nil {
        t.Errorf("FromMegavoltsPerHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMegavoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltsPerHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltsPerHours(math.NaN())
    if err == nil {
        t.Error("FromMegavoltsPerHours() with NaN value should return error")
    }

    _, err = factory.FromMegavoltsPerHours(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltsPerHours() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltsPerHours(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltsPerHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltsPerHours(0)
    if err != nil {
        t.Errorf("FromMegavoltsPerHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMegavoltPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltsPerHours() with zero value = %v, want 0", converted)
    }
}

func TestElectricPotentialChangeRateToString(t *testing.T) {
	factory := units.ElectricPotentialChangeRateFactory{}
	a, err := factory.CreateElectricPotentialChangeRate(45, units.ElectricPotentialChangeRateVoltPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricPotentialChangeRateVoltPerSecond, 2)
	expected := "45.00 " + units.GetElectricPotentialChangeRateAbbreviation(units.ElectricPotentialChangeRateVoltPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricPotentialChangeRateVoltPerSecond, -1)
	expected = "45 " + units.GetElectricPotentialChangeRateAbbreviation(units.ElectricPotentialChangeRateVoltPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricPotentialChangeRate_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricPotentialChangeRateFactory{}
	a1, _ := factory.CreateElectricPotentialChangeRate(60, units.ElectricPotentialChangeRateVoltPerSecond)
	a2, _ := factory.CreateElectricPotentialChangeRate(60, units.ElectricPotentialChangeRateVoltPerSecond)
	a3, _ := factory.CreateElectricPotentialChangeRate(120, units.ElectricPotentialChangeRateVoltPerSecond)

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

func TestElectricPotentialChangeRate_Arithmetic(t *testing.T) {
	factory := units.ElectricPotentialChangeRateFactory{}
	a1, _ := factory.CreateElectricPotentialChangeRate(30, units.ElectricPotentialChangeRateVoltPerSecond)
	a2, _ := factory.CreateElectricPotentialChangeRate(45, units.ElectricPotentialChangeRateVoltPerSecond)

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


func TestGetElectricPotentialChangeRateAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricPotentialChangeRateUnits
        want string
    }{
        {
            name: "VoltPerSecond abbreviation",
            unit: units.ElectricPotentialChangeRateVoltPerSecond,
            want: "V/s",
        },
        {
            name: "VoltPerMicrosecond abbreviation",
            unit: units.ElectricPotentialChangeRateVoltPerMicrosecond,
            want: "V/μs",
        },
        {
            name: "VoltPerMinute abbreviation",
            unit: units.ElectricPotentialChangeRateVoltPerMinute,
            want: "V/min",
        },
        {
            name: "VoltPerHour abbreviation",
            unit: units.ElectricPotentialChangeRateVoltPerHour,
            want: "V/h",
        },
        {
            name: "MicrovoltPerSecond abbreviation",
            unit: units.ElectricPotentialChangeRateMicrovoltPerSecond,
            want: "μV/s",
        },
        {
            name: "MillivoltPerSecond abbreviation",
            unit: units.ElectricPotentialChangeRateMillivoltPerSecond,
            want: "mV/s",
        },
        {
            name: "KilovoltPerSecond abbreviation",
            unit: units.ElectricPotentialChangeRateKilovoltPerSecond,
            want: "kV/s",
        },
        {
            name: "MegavoltPerSecond abbreviation",
            unit: units.ElectricPotentialChangeRateMegavoltPerSecond,
            want: "MV/s",
        },
        {
            name: "MicrovoltPerMicrosecond abbreviation",
            unit: units.ElectricPotentialChangeRateMicrovoltPerMicrosecond,
            want: "μV/μs",
        },
        {
            name: "MillivoltPerMicrosecond abbreviation",
            unit: units.ElectricPotentialChangeRateMillivoltPerMicrosecond,
            want: "mV/μs",
        },
        {
            name: "KilovoltPerMicrosecond abbreviation",
            unit: units.ElectricPotentialChangeRateKilovoltPerMicrosecond,
            want: "kV/μs",
        },
        {
            name: "MegavoltPerMicrosecond abbreviation",
            unit: units.ElectricPotentialChangeRateMegavoltPerMicrosecond,
            want: "MV/μs",
        },
        {
            name: "MicrovoltPerMinute abbreviation",
            unit: units.ElectricPotentialChangeRateMicrovoltPerMinute,
            want: "μV/min",
        },
        {
            name: "MillivoltPerMinute abbreviation",
            unit: units.ElectricPotentialChangeRateMillivoltPerMinute,
            want: "mV/min",
        },
        {
            name: "KilovoltPerMinute abbreviation",
            unit: units.ElectricPotentialChangeRateKilovoltPerMinute,
            want: "kV/min",
        },
        {
            name: "MegavoltPerMinute abbreviation",
            unit: units.ElectricPotentialChangeRateMegavoltPerMinute,
            want: "MV/min",
        },
        {
            name: "MicrovoltPerHour abbreviation",
            unit: units.ElectricPotentialChangeRateMicrovoltPerHour,
            want: "μV/h",
        },
        {
            name: "MillivoltPerHour abbreviation",
            unit: units.ElectricPotentialChangeRateMillivoltPerHour,
            want: "mV/h",
        },
        {
            name: "KilovoltPerHour abbreviation",
            unit: units.ElectricPotentialChangeRateKilovoltPerHour,
            want: "kV/h",
        },
        {
            name: "MegavoltPerHour abbreviation",
            unit: units.ElectricPotentialChangeRateMegavoltPerHour,
            want: "MV/h",
        },
        {
            name: "invalid unit",
            unit: units.ElectricPotentialChangeRateUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricPotentialChangeRateAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricPotentialChangeRateAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricPotentialChangeRate_String(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    
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
            unit, err := factory.CreateElectricPotentialChangeRate(tt.value, units.ElectricPotentialChangeRateVoltPerSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricPotentialChangeRate.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestElectricPotentialChangeRate_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.ElectricPotentialChangeRateFactory{}

	_, err := uf.CreateElectricPotentialChangeRate(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}