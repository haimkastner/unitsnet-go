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
		// Test conversion to VoltsPerSecond.
		// No expected conversion value provided for VoltsPerSecond, verifying result is not NaN.
		result := a.VoltsPerSecond()
		cacheResult := a.VoltsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to VoltsPerMicrosecond.
		// No expected conversion value provided for VoltsPerMicrosecond, verifying result is not NaN.
		result := a.VoltsPerMicrosecond()
		cacheResult := a.VoltsPerMicrosecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltsPerMicrosecond returned NaN")
		}
	}
	{
		// Test conversion to VoltsPerMinute.
		// No expected conversion value provided for VoltsPerMinute, verifying result is not NaN.
		result := a.VoltsPerMinute()
		cacheResult := a.VoltsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to VoltsPerHour.
		// No expected conversion value provided for VoltsPerHour, verifying result is not NaN.
		result := a.VoltsPerHour()
		cacheResult := a.VoltsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltsPerHour returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsPerSecond.
		// No expected conversion value provided for MicrovoltsPerSecond, verifying result is not NaN.
		result := a.MicrovoltsPerSecond()
		cacheResult := a.MicrovoltsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrovoltsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsPerSecond.
		// No expected conversion value provided for MillivoltsPerSecond, verifying result is not NaN.
		result := a.MillivoltsPerSecond()
		cacheResult := a.MillivoltsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillivoltsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsPerSecond.
		// No expected conversion value provided for KilovoltsPerSecond, verifying result is not NaN.
		result := a.KilovoltsPerSecond()
		cacheResult := a.KilovoltsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilovoltsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsPerSecond.
		// No expected conversion value provided for MegavoltsPerSecond, verifying result is not NaN.
		result := a.MegavoltsPerSecond()
		cacheResult := a.MegavoltsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegavoltsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsPerMicrosecond.
		// No expected conversion value provided for MicrovoltsPerMicrosecond, verifying result is not NaN.
		result := a.MicrovoltsPerMicrosecond()
		cacheResult := a.MicrovoltsPerMicrosecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrovoltsPerMicrosecond returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsPerMicrosecond.
		// No expected conversion value provided for MillivoltsPerMicrosecond, verifying result is not NaN.
		result := a.MillivoltsPerMicrosecond()
		cacheResult := a.MillivoltsPerMicrosecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillivoltsPerMicrosecond returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsPerMicrosecond.
		// No expected conversion value provided for KilovoltsPerMicrosecond, verifying result is not NaN.
		result := a.KilovoltsPerMicrosecond()
		cacheResult := a.KilovoltsPerMicrosecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilovoltsPerMicrosecond returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsPerMicrosecond.
		// No expected conversion value provided for MegavoltsPerMicrosecond, verifying result is not NaN.
		result := a.MegavoltsPerMicrosecond()
		cacheResult := a.MegavoltsPerMicrosecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegavoltsPerMicrosecond returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsPerMinute.
		// No expected conversion value provided for MicrovoltsPerMinute, verifying result is not NaN.
		result := a.MicrovoltsPerMinute()
		cacheResult := a.MicrovoltsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrovoltsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsPerMinute.
		// No expected conversion value provided for MillivoltsPerMinute, verifying result is not NaN.
		result := a.MillivoltsPerMinute()
		cacheResult := a.MillivoltsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillivoltsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsPerMinute.
		// No expected conversion value provided for KilovoltsPerMinute, verifying result is not NaN.
		result := a.KilovoltsPerMinute()
		cacheResult := a.KilovoltsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilovoltsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsPerMinute.
		// No expected conversion value provided for MegavoltsPerMinute, verifying result is not NaN.
		result := a.MegavoltsPerMinute()
		cacheResult := a.MegavoltsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegavoltsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsPerHour.
		// No expected conversion value provided for MicrovoltsPerHour, verifying result is not NaN.
		result := a.MicrovoltsPerHour()
		cacheResult := a.MicrovoltsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrovoltsPerHour returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsPerHour.
		// No expected conversion value provided for MillivoltsPerHour, verifying result is not NaN.
		result := a.MillivoltsPerHour()
		cacheResult := a.MillivoltsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillivoltsPerHour returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsPerHour.
		// No expected conversion value provided for KilovoltsPerHour, verifying result is not NaN.
		result := a.KilovoltsPerHour()
		cacheResult := a.KilovoltsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilovoltsPerHour returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsPerHour.
		// No expected conversion value provided for MegavoltsPerHour, verifying result is not NaN.
		result := a.MegavoltsPerHour()
		cacheResult := a.MegavoltsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegavoltsPerHour returned NaN")
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
    volts_per_secondDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateVoltPerSecond,
    }
    
    var volts_per_secondResult *units.ElectricPotentialChangeRate
    volts_per_secondResult, err = factory.FromDto(volts_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with VoltPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_secondResult.Convert(units.ElectricPotentialChangeRateVoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerSecond = %v, want %v", converted, 100)
    }
    // Test VoltPerMicrosecond conversion
    volts_per_microsecondDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateVoltPerMicrosecond,
    }
    
    var volts_per_microsecondResult *units.ElectricPotentialChangeRate
    volts_per_microsecondResult, err = factory.FromDto(volts_per_microsecondDto)
    if err != nil {
        t.Errorf("FromDto() with VoltPerMicrosecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_microsecondResult.Convert(units.ElectricPotentialChangeRateVoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test VoltPerMinute conversion
    volts_per_minuteDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateVoltPerMinute,
    }
    
    var volts_per_minuteResult *units.ElectricPotentialChangeRate
    volts_per_minuteResult, err = factory.FromDto(volts_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with VoltPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_minuteResult.Convert(units.ElectricPotentialChangeRateVoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerMinute = %v, want %v", converted, 100)
    }
    // Test VoltPerHour conversion
    volts_per_hourDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateVoltPerHour,
    }
    
    var volts_per_hourResult *units.ElectricPotentialChangeRate
    volts_per_hourResult, err = factory.FromDto(volts_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with VoltPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_hourResult.Convert(units.ElectricPotentialChangeRateVoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerHour = %v, want %v", converted, 100)
    }
    // Test MicrovoltPerSecond conversion
    microvolts_per_secondDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMicrovoltPerSecond,
    }
    
    var microvolts_per_secondResult *units.ElectricPotentialChangeRate
    microvolts_per_secondResult, err = factory.FromDto(microvolts_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MicrovoltPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_secondResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerSecond = %v, want %v", converted, 100)
    }
    // Test MillivoltPerSecond conversion
    millivolts_per_secondDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMillivoltPerSecond,
    }
    
    var millivolts_per_secondResult *units.ElectricPotentialChangeRate
    millivolts_per_secondResult, err = factory.FromDto(millivolts_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MillivoltPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_secondResult.Convert(units.ElectricPotentialChangeRateMillivoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerSecond = %v, want %v", converted, 100)
    }
    // Test KilovoltPerSecond conversion
    kilovolts_per_secondDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateKilovoltPerSecond,
    }
    
    var kilovolts_per_secondResult *units.ElectricPotentialChangeRate
    kilovolts_per_secondResult, err = factory.FromDto(kilovolts_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_secondResult.Convert(units.ElectricPotentialChangeRateKilovoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerSecond = %v, want %v", converted, 100)
    }
    // Test MegavoltPerSecond conversion
    megavolts_per_secondDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMegavoltPerSecond,
    }
    
    var megavolts_per_secondResult *units.ElectricPotentialChangeRate
    megavolts_per_secondResult, err = factory.FromDto(megavolts_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_secondResult.Convert(units.ElectricPotentialChangeRateMegavoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerSecond = %v, want %v", converted, 100)
    }
    // Test MicrovoltPerMicrosecond conversion
    microvolts_per_microsecondDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMicrovoltPerMicrosecond,
    }
    
    var microvolts_per_microsecondResult *units.ElectricPotentialChangeRate
    microvolts_per_microsecondResult, err = factory.FromDto(microvolts_per_microsecondDto)
    if err != nil {
        t.Errorf("FromDto() with MicrovoltPerMicrosecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_microsecondResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test MillivoltPerMicrosecond conversion
    millivolts_per_microsecondDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMillivoltPerMicrosecond,
    }
    
    var millivolts_per_microsecondResult *units.ElectricPotentialChangeRate
    millivolts_per_microsecondResult, err = factory.FromDto(millivolts_per_microsecondDto)
    if err != nil {
        t.Errorf("FromDto() with MillivoltPerMicrosecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_microsecondResult.Convert(units.ElectricPotentialChangeRateMillivoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test KilovoltPerMicrosecond conversion
    kilovolts_per_microsecondDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateKilovoltPerMicrosecond,
    }
    
    var kilovolts_per_microsecondResult *units.ElectricPotentialChangeRate
    kilovolts_per_microsecondResult, err = factory.FromDto(kilovolts_per_microsecondDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltPerMicrosecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_microsecondResult.Convert(units.ElectricPotentialChangeRateKilovoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test MegavoltPerMicrosecond conversion
    megavolts_per_microsecondDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMegavoltPerMicrosecond,
    }
    
    var megavolts_per_microsecondResult *units.ElectricPotentialChangeRate
    megavolts_per_microsecondResult, err = factory.FromDto(megavolts_per_microsecondDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltPerMicrosecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_microsecondResult.Convert(units.ElectricPotentialChangeRateMegavoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test MicrovoltPerMinute conversion
    microvolts_per_minuteDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMicrovoltPerMinute,
    }
    
    var microvolts_per_minuteResult *units.ElectricPotentialChangeRate
    microvolts_per_minuteResult, err = factory.FromDto(microvolts_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with MicrovoltPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_minuteResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerMinute = %v, want %v", converted, 100)
    }
    // Test MillivoltPerMinute conversion
    millivolts_per_minuteDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMillivoltPerMinute,
    }
    
    var millivolts_per_minuteResult *units.ElectricPotentialChangeRate
    millivolts_per_minuteResult, err = factory.FromDto(millivolts_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with MillivoltPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_minuteResult.Convert(units.ElectricPotentialChangeRateMillivoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerMinute = %v, want %v", converted, 100)
    }
    // Test KilovoltPerMinute conversion
    kilovolts_per_minuteDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateKilovoltPerMinute,
    }
    
    var kilovolts_per_minuteResult *units.ElectricPotentialChangeRate
    kilovolts_per_minuteResult, err = factory.FromDto(kilovolts_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_minuteResult.Convert(units.ElectricPotentialChangeRateKilovoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerMinute = %v, want %v", converted, 100)
    }
    // Test MegavoltPerMinute conversion
    megavolts_per_minuteDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMegavoltPerMinute,
    }
    
    var megavolts_per_minuteResult *units.ElectricPotentialChangeRate
    megavolts_per_minuteResult, err = factory.FromDto(megavolts_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_minuteResult.Convert(units.ElectricPotentialChangeRateMegavoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerMinute = %v, want %v", converted, 100)
    }
    // Test MicrovoltPerHour conversion
    microvolts_per_hourDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMicrovoltPerHour,
    }
    
    var microvolts_per_hourResult *units.ElectricPotentialChangeRate
    microvolts_per_hourResult, err = factory.FromDto(microvolts_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MicrovoltPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_hourResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerHour = %v, want %v", converted, 100)
    }
    // Test MillivoltPerHour conversion
    millivolts_per_hourDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMillivoltPerHour,
    }
    
    var millivolts_per_hourResult *units.ElectricPotentialChangeRate
    millivolts_per_hourResult, err = factory.FromDto(millivolts_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MillivoltPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_hourResult.Convert(units.ElectricPotentialChangeRateMillivoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerHour = %v, want %v", converted, 100)
    }
    // Test KilovoltPerHour conversion
    kilovolts_per_hourDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateKilovoltPerHour,
    }
    
    var kilovolts_per_hourResult *units.ElectricPotentialChangeRate
    kilovolts_per_hourResult, err = factory.FromDto(kilovolts_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_hourResult.Convert(units.ElectricPotentialChangeRateKilovoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerHour = %v, want %v", converted, 100)
    }
    // Test MegavoltPerHour conversion
    megavolts_per_hourDto := units.ElectricPotentialChangeRateDto{
        Value: 100,
        Unit:  units.ElectricPotentialChangeRateMegavoltPerHour,
    }
    
    var megavolts_per_hourResult *units.ElectricPotentialChangeRate
    megavolts_per_hourResult, err = factory.FromDto(megavolts_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_hourResult.Convert(units.ElectricPotentialChangeRateMegavoltPerHour)
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
    volts_per_secondJSON := []byte(`{"value": 100, "unit": "VoltPerSecond"}`)
    volts_per_secondResult, err := factory.FromDtoJSON(volts_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_secondResult.Convert(units.ElectricPotentialChangeRateVoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with VoltPerMicrosecond unit
    volts_per_microsecondJSON := []byte(`{"value": 100, "unit": "VoltPerMicrosecond"}`)
    volts_per_microsecondResult, err := factory.FromDtoJSON(volts_per_microsecondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltPerMicrosecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_microsecondResult.Convert(units.ElectricPotentialChangeRateVoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test JSON with VoltPerMinute unit
    volts_per_minuteJSON := []byte(`{"value": 100, "unit": "VoltPerMinute"}`)
    volts_per_minuteResult, err := factory.FromDtoJSON(volts_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_minuteResult.Convert(units.ElectricPotentialChangeRateVoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with VoltPerHour unit
    volts_per_hourJSON := []byte(`{"value": 100, "unit": "VoltPerHour"}`)
    volts_per_hourResult, err := factory.FromDtoJSON(volts_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_hourResult.Convert(units.ElectricPotentialChangeRateVoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MicrovoltPerSecond unit
    microvolts_per_secondJSON := []byte(`{"value": 100, "unit": "MicrovoltPerSecond"}`)
    microvolts_per_secondResult, err := factory.FromDtoJSON(microvolts_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrovoltPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_secondResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MillivoltPerSecond unit
    millivolts_per_secondJSON := []byte(`{"value": 100, "unit": "MillivoltPerSecond"}`)
    millivolts_per_secondResult, err := factory.FromDtoJSON(millivolts_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillivoltPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_secondResult.Convert(units.ElectricPotentialChangeRateMillivoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltPerSecond unit
    kilovolts_per_secondJSON := []byte(`{"value": 100, "unit": "KilovoltPerSecond"}`)
    kilovolts_per_secondResult, err := factory.FromDtoJSON(kilovolts_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_secondResult.Convert(units.ElectricPotentialChangeRateKilovoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltPerSecond unit
    megavolts_per_secondJSON := []byte(`{"value": 100, "unit": "MegavoltPerSecond"}`)
    megavolts_per_secondResult, err := factory.FromDtoJSON(megavolts_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_secondResult.Convert(units.ElectricPotentialChangeRateMegavoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicrovoltPerMicrosecond unit
    microvolts_per_microsecondJSON := []byte(`{"value": 100, "unit": "MicrovoltPerMicrosecond"}`)
    microvolts_per_microsecondResult, err := factory.FromDtoJSON(microvolts_per_microsecondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrovoltPerMicrosecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_microsecondResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test JSON with MillivoltPerMicrosecond unit
    millivolts_per_microsecondJSON := []byte(`{"value": 100, "unit": "MillivoltPerMicrosecond"}`)
    millivolts_per_microsecondResult, err := factory.FromDtoJSON(millivolts_per_microsecondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillivoltPerMicrosecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_microsecondResult.Convert(units.ElectricPotentialChangeRateMillivoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltPerMicrosecond unit
    kilovolts_per_microsecondJSON := []byte(`{"value": 100, "unit": "KilovoltPerMicrosecond"}`)
    kilovolts_per_microsecondResult, err := factory.FromDtoJSON(kilovolts_per_microsecondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltPerMicrosecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_microsecondResult.Convert(units.ElectricPotentialChangeRateKilovoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltPerMicrosecond unit
    megavolts_per_microsecondJSON := []byte(`{"value": 100, "unit": "MegavoltPerMicrosecond"}`)
    megavolts_per_microsecondResult, err := factory.FromDtoJSON(megavolts_per_microsecondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltPerMicrosecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_microsecondResult.Convert(units.ElectricPotentialChangeRateMegavoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerMicrosecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicrovoltPerMinute unit
    microvolts_per_minuteJSON := []byte(`{"value": 100, "unit": "MicrovoltPerMinute"}`)
    microvolts_per_minuteResult, err := factory.FromDtoJSON(microvolts_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrovoltPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_minuteResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MillivoltPerMinute unit
    millivolts_per_minuteJSON := []byte(`{"value": 100, "unit": "MillivoltPerMinute"}`)
    millivolts_per_minuteResult, err := factory.FromDtoJSON(millivolts_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillivoltPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_minuteResult.Convert(units.ElectricPotentialChangeRateMillivoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltPerMinute unit
    kilovolts_per_minuteJSON := []byte(`{"value": 100, "unit": "KilovoltPerMinute"}`)
    kilovolts_per_minuteResult, err := factory.FromDtoJSON(kilovolts_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_minuteResult.Convert(units.ElectricPotentialChangeRateKilovoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltPerMinute unit
    megavolts_per_minuteJSON := []byte(`{"value": 100, "unit": "MegavoltPerMinute"}`)
    megavolts_per_minuteResult, err := factory.FromDtoJSON(megavolts_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_minuteResult.Convert(units.ElectricPotentialChangeRateMegavoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MicrovoltPerHour unit
    microvolts_per_hourJSON := []byte(`{"value": 100, "unit": "MicrovoltPerHour"}`)
    microvolts_per_hourResult, err := factory.FromDtoJSON(microvolts_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrovoltPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_per_hourResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MillivoltPerHour unit
    millivolts_per_hourJSON := []byte(`{"value": 100, "unit": "MillivoltPerHour"}`)
    millivolts_per_hourResult, err := factory.FromDtoJSON(millivolts_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillivoltPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_per_hourResult.Convert(units.ElectricPotentialChangeRateMillivoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltPerHour unit
    kilovolts_per_hourJSON := []byte(`{"value": 100, "unit": "KilovoltPerHour"}`)
    kilovolts_per_hourResult, err := factory.FromDtoJSON(kilovolts_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_per_hourResult.Convert(units.ElectricPotentialChangeRateKilovoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltPerHour unit
    megavolts_per_hourJSON := []byte(`{"value": 100, "unit": "MegavoltPerHour"}`)
    megavolts_per_hourResult, err := factory.FromDtoJSON(megavolts_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_per_hourResult.Convert(units.ElectricPotentialChangeRateMegavoltPerHour)
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
// Test FromVoltsPerSecond function
func TestElectricPotentialChangeRateFactory_FromVoltsPerSecond(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltsPerSecond(100)
    if err != nil {
        t.Errorf("FromVoltsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateVoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromVoltsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromVoltsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromVoltsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromVoltsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltsPerSecond(0)
    if err != nil {
        t.Errorf("FromVoltsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateVoltPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromVoltsPerMicrosecond function
func TestElectricPotentialChangeRateFactory_FromVoltsPerMicrosecond(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltsPerMicrosecond(100)
    if err != nil {
        t.Errorf("FromVoltsPerMicrosecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateVoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltsPerMicrosecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltsPerMicrosecond(math.NaN())
    if err == nil {
        t.Error("FromVoltsPerMicrosecond() with NaN value should return error")
    }

    _, err = factory.FromVoltsPerMicrosecond(math.Inf(1))
    if err == nil {
        t.Error("FromVoltsPerMicrosecond() with +Inf value should return error")
    }

    _, err = factory.FromVoltsPerMicrosecond(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltsPerMicrosecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltsPerMicrosecond(0)
    if err != nil {
        t.Errorf("FromVoltsPerMicrosecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateVoltPerMicrosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltsPerMicrosecond() with zero value = %v, want 0", converted)
    }
}
// Test FromVoltsPerMinute function
func TestElectricPotentialChangeRateFactory_FromVoltsPerMinute(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltsPerMinute(100)
    if err != nil {
        t.Errorf("FromVoltsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateVoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromVoltsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromVoltsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromVoltsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromVoltsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltsPerMinute(0)
    if err != nil {
        t.Errorf("FromVoltsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateVoltPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromVoltsPerHour function
func TestElectricPotentialChangeRateFactory_FromVoltsPerHour(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltsPerHour(100)
    if err != nil {
        t.Errorf("FromVoltsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateVoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltsPerHour(math.NaN())
    if err == nil {
        t.Error("FromVoltsPerHour() with NaN value should return error")
    }

    _, err = factory.FromVoltsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromVoltsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromVoltsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltsPerHour(0)
    if err != nil {
        t.Errorf("FromVoltsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateVoltPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrovoltsPerSecond function
func TestElectricPotentialChangeRateFactory_FromMicrovoltsPerSecond(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrovoltsPerSecond(100)
    if err != nil {
        t.Errorf("FromMicrovoltsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMicrovoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrovoltsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrovoltsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMicrovoltsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMicrovoltsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMicrovoltsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMicrovoltsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrovoltsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrovoltsPerSecond(0)
    if err != nil {
        t.Errorf("FromMicrovoltsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrovoltsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMillivoltsPerSecond function
func TestElectricPotentialChangeRateFactory_FromMillivoltsPerSecond(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillivoltsPerSecond(100)
    if err != nil {
        t.Errorf("FromMillivoltsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMillivoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillivoltsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillivoltsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMillivoltsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMillivoltsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMillivoltsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMillivoltsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMillivoltsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillivoltsPerSecond(0)
    if err != nil {
        t.Errorf("FromMillivoltsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMillivoltPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillivoltsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltsPerSecond function
func TestElectricPotentialChangeRateFactory_FromKilovoltsPerSecond(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltsPerSecond(100)
    if err != nil {
        t.Errorf("FromKilovoltsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateKilovoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilovoltsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilovoltsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltsPerSecond(0)
    if err != nil {
        t.Errorf("FromKilovoltsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateKilovoltPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltsPerSecond function
func TestElectricPotentialChangeRateFactory_FromMegavoltsPerSecond(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltsPerSecond(100)
    if err != nil {
        t.Errorf("FromMegavoltsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMegavoltPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMegavoltsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMegavoltsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltsPerSecond(0)
    if err != nil {
        t.Errorf("FromMegavoltsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMegavoltPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrovoltsPerMicrosecond function
func TestElectricPotentialChangeRateFactory_FromMicrovoltsPerMicrosecond(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrovoltsPerMicrosecond(100)
    if err != nil {
        t.Errorf("FromMicrovoltsPerMicrosecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMicrovoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrovoltsPerMicrosecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrovoltsPerMicrosecond(math.NaN())
    if err == nil {
        t.Error("FromMicrovoltsPerMicrosecond() with NaN value should return error")
    }

    _, err = factory.FromMicrovoltsPerMicrosecond(math.Inf(1))
    if err == nil {
        t.Error("FromMicrovoltsPerMicrosecond() with +Inf value should return error")
    }

    _, err = factory.FromMicrovoltsPerMicrosecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrovoltsPerMicrosecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrovoltsPerMicrosecond(0)
    if err != nil {
        t.Errorf("FromMicrovoltsPerMicrosecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerMicrosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrovoltsPerMicrosecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMillivoltsPerMicrosecond function
func TestElectricPotentialChangeRateFactory_FromMillivoltsPerMicrosecond(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillivoltsPerMicrosecond(100)
    if err != nil {
        t.Errorf("FromMillivoltsPerMicrosecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMillivoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillivoltsPerMicrosecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillivoltsPerMicrosecond(math.NaN())
    if err == nil {
        t.Error("FromMillivoltsPerMicrosecond() with NaN value should return error")
    }

    _, err = factory.FromMillivoltsPerMicrosecond(math.Inf(1))
    if err == nil {
        t.Error("FromMillivoltsPerMicrosecond() with +Inf value should return error")
    }

    _, err = factory.FromMillivoltsPerMicrosecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMillivoltsPerMicrosecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillivoltsPerMicrosecond(0)
    if err != nil {
        t.Errorf("FromMillivoltsPerMicrosecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMillivoltPerMicrosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillivoltsPerMicrosecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltsPerMicrosecond function
func TestElectricPotentialChangeRateFactory_FromKilovoltsPerMicrosecond(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltsPerMicrosecond(100)
    if err != nil {
        t.Errorf("FromKilovoltsPerMicrosecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateKilovoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltsPerMicrosecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltsPerMicrosecond(math.NaN())
    if err == nil {
        t.Error("FromKilovoltsPerMicrosecond() with NaN value should return error")
    }

    _, err = factory.FromKilovoltsPerMicrosecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltsPerMicrosecond() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltsPerMicrosecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltsPerMicrosecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltsPerMicrosecond(0)
    if err != nil {
        t.Errorf("FromKilovoltsPerMicrosecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateKilovoltPerMicrosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltsPerMicrosecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltsPerMicrosecond function
func TestElectricPotentialChangeRateFactory_FromMegavoltsPerMicrosecond(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltsPerMicrosecond(100)
    if err != nil {
        t.Errorf("FromMegavoltsPerMicrosecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMegavoltPerMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltsPerMicrosecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltsPerMicrosecond(math.NaN())
    if err == nil {
        t.Error("FromMegavoltsPerMicrosecond() with NaN value should return error")
    }

    _, err = factory.FromMegavoltsPerMicrosecond(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltsPerMicrosecond() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltsPerMicrosecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltsPerMicrosecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltsPerMicrosecond(0)
    if err != nil {
        t.Errorf("FromMegavoltsPerMicrosecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMegavoltPerMicrosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltsPerMicrosecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrovoltsPerMinute function
func TestElectricPotentialChangeRateFactory_FromMicrovoltsPerMinute(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrovoltsPerMinute(100)
    if err != nil {
        t.Errorf("FromMicrovoltsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMicrovoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrovoltsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrovoltsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromMicrovoltsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromMicrovoltsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromMicrovoltsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromMicrovoltsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrovoltsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrovoltsPerMinute(0)
    if err != nil {
        t.Errorf("FromMicrovoltsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrovoltsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromMillivoltsPerMinute function
func TestElectricPotentialChangeRateFactory_FromMillivoltsPerMinute(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillivoltsPerMinute(100)
    if err != nil {
        t.Errorf("FromMillivoltsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMillivoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillivoltsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillivoltsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromMillivoltsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromMillivoltsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromMillivoltsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromMillivoltsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromMillivoltsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillivoltsPerMinute(0)
    if err != nil {
        t.Errorf("FromMillivoltsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMillivoltPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillivoltsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltsPerMinute function
func TestElectricPotentialChangeRateFactory_FromKilovoltsPerMinute(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltsPerMinute(100)
    if err != nil {
        t.Errorf("FromKilovoltsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateKilovoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromKilovoltsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromKilovoltsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltsPerMinute(0)
    if err != nil {
        t.Errorf("FromKilovoltsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateKilovoltPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltsPerMinute function
func TestElectricPotentialChangeRateFactory_FromMegavoltsPerMinute(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltsPerMinute(100)
    if err != nil {
        t.Errorf("FromMegavoltsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMegavoltPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromMegavoltsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromMegavoltsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltsPerMinute(0)
    if err != nil {
        t.Errorf("FromMegavoltsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMegavoltPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrovoltsPerHour function
func TestElectricPotentialChangeRateFactory_FromMicrovoltsPerHour(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrovoltsPerHour(100)
    if err != nil {
        t.Errorf("FromMicrovoltsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMicrovoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrovoltsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrovoltsPerHour(math.NaN())
    if err == nil {
        t.Error("FromMicrovoltsPerHour() with NaN value should return error")
    }

    _, err = factory.FromMicrovoltsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMicrovoltsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMicrovoltsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrovoltsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrovoltsPerHour(0)
    if err != nil {
        t.Errorf("FromMicrovoltsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMicrovoltPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrovoltsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromMillivoltsPerHour function
func TestElectricPotentialChangeRateFactory_FromMillivoltsPerHour(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillivoltsPerHour(100)
    if err != nil {
        t.Errorf("FromMillivoltsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMillivoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillivoltsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillivoltsPerHour(math.NaN())
    if err == nil {
        t.Error("FromMillivoltsPerHour() with NaN value should return error")
    }

    _, err = factory.FromMillivoltsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMillivoltsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMillivoltsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMillivoltsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillivoltsPerHour(0)
    if err != nil {
        t.Errorf("FromMillivoltsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMillivoltPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillivoltsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltsPerHour function
func TestElectricPotentialChangeRateFactory_FromKilovoltsPerHour(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltsPerHour(100)
    if err != nil {
        t.Errorf("FromKilovoltsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateKilovoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltsPerHour(math.NaN())
    if err == nil {
        t.Error("FromKilovoltsPerHour() with NaN value should return error")
    }

    _, err = factory.FromKilovoltsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltsPerHour(0)
    if err != nil {
        t.Errorf("FromKilovoltsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateKilovoltPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltsPerHour function
func TestElectricPotentialChangeRateFactory_FromMegavoltsPerHour(t *testing.T) {
    factory := units.ElectricPotentialChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltsPerHour(100)
    if err != nil {
        t.Errorf("FromMegavoltsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialChangeRateMegavoltPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltsPerHour(math.NaN())
    if err == nil {
        t.Error("FromMegavoltsPerHour() with NaN value should return error")
    }

    _, err = factory.FromMegavoltsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltsPerHour(0)
    if err != nil {
        t.Errorf("FromMegavoltsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialChangeRateMegavoltPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltsPerHour() with zero value = %v, want 0", converted)
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