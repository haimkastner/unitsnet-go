// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestFluidResistanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "PascalSecondPerCubicMeter"}`
	
	factory := units.FluidResistanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.FluidResistancePascalSecondPerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.FluidResistancePascalSecondPerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "PascalSecondPerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestFluidResistanceDto_ToJSON(t *testing.T) {
	dto := units.FluidResistanceDto{
		Value: 45,
		Unit:  units.FluidResistancePascalSecondPerCubicMeter,
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
	if result["unit"].(string) != string(units.FluidResistancePascalSecondPerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.FluidResistancePascalSecondPerCubicMeter, result["unit"])
	}
}

func TestNewFluidResistance_InvalidValue(t *testing.T) {
	factory := units.FluidResistanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateFluidResistance(math.NaN(), units.FluidResistancePascalSecondPerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateFluidResistance(math.Inf(1), units.FluidResistancePascalSecondPerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestFluidResistanceConversions(t *testing.T) {
	factory := units.FluidResistanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateFluidResistance(180, units.FluidResistancePascalSecondPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to PascalSecondsPerLiter.
		// No expected conversion value provided for PascalSecondsPerLiter, verifying result is not NaN.
		result := a.PascalSecondsPerLiter()
		cacheResult := a.PascalSecondsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PascalSecondsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to PascalMinutesPerLiter.
		// No expected conversion value provided for PascalMinutesPerLiter, verifying result is not NaN.
		result := a.PascalMinutesPerLiter()
		cacheResult := a.PascalMinutesPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PascalMinutesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to PascalSecondsPerMilliliter.
		// No expected conversion value provided for PascalSecondsPerMilliliter, verifying result is not NaN.
		result := a.PascalSecondsPerMilliliter()
		cacheResult := a.PascalSecondsPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PascalSecondsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to PascalMinutesPerMilliliter.
		// No expected conversion value provided for PascalMinutesPerMilliliter, verifying result is not NaN.
		result := a.PascalMinutesPerMilliliter()
		cacheResult := a.PascalMinutesPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PascalMinutesPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to PascalSecondsPerCubicMeter.
		// No expected conversion value provided for PascalSecondsPerCubicMeter, verifying result is not NaN.
		result := a.PascalSecondsPerCubicMeter()
		cacheResult := a.PascalSecondsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PascalSecondsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PascalMinutesPerCubicMeter.
		// No expected conversion value provided for PascalMinutesPerCubicMeter, verifying result is not NaN.
		result := a.PascalMinutesPerCubicMeter()
		cacheResult := a.PascalMinutesPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PascalMinutesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PascalSecondsPerCubicCentimeter.
		// No expected conversion value provided for PascalSecondsPerCubicCentimeter, verifying result is not NaN.
		result := a.PascalSecondsPerCubicCentimeter()
		cacheResult := a.PascalSecondsPerCubicCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PascalSecondsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to PascalMinutesPerCubicCentimeter.
		// No expected conversion value provided for PascalMinutesPerCubicCentimeter, verifying result is not NaN.
		result := a.PascalMinutesPerCubicCentimeter()
		cacheResult := a.PascalMinutesPerCubicCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PascalMinutesPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to DyneSecondsPerCentimeterToTheFifth.
		// No expected conversion value provided for DyneSecondsPerCentimeterToTheFifth, verifying result is not NaN.
		result := a.DyneSecondsPerCentimeterToTheFifth()
		cacheResult := a.DyneSecondsPerCentimeterToTheFifth()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DyneSecondsPerCentimeterToTheFifth returned NaN")
		}
	}
	{
		// Test conversion to MillimeterMercurySecondsPerLiter.
		// No expected conversion value provided for MillimeterMercurySecondsPerLiter, verifying result is not NaN.
		result := a.MillimeterMercurySecondsPerLiter()
		cacheResult := a.MillimeterMercurySecondsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimeterMercurySecondsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MillimeterMercuryMinutesPerLiter.
		// No expected conversion value provided for MillimeterMercuryMinutesPerLiter, verifying result is not NaN.
		result := a.MillimeterMercuryMinutesPerLiter()
		cacheResult := a.MillimeterMercuryMinutesPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimeterMercuryMinutesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MillimeterMercurySecondsPerMilliliter.
		// No expected conversion value provided for MillimeterMercurySecondsPerMilliliter, verifying result is not NaN.
		result := a.MillimeterMercurySecondsPerMilliliter()
		cacheResult := a.MillimeterMercurySecondsPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimeterMercurySecondsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to MillimeterMercuryMinutesPerMilliliter.
		// No expected conversion value provided for MillimeterMercuryMinutesPerMilliliter, verifying result is not NaN.
		result := a.MillimeterMercuryMinutesPerMilliliter()
		cacheResult := a.MillimeterMercuryMinutesPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimeterMercuryMinutesPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to MillimeterMercurySecondsPerCubicCentimeter.
		// No expected conversion value provided for MillimeterMercurySecondsPerCubicCentimeter, verifying result is not NaN.
		result := a.MillimeterMercurySecondsPerCubicCentimeter()
		cacheResult := a.MillimeterMercurySecondsPerCubicCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimeterMercurySecondsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MillimeterMercuryMinutesPerCubicCentimeter.
		// No expected conversion value provided for MillimeterMercuryMinutesPerCubicCentimeter, verifying result is not NaN.
		result := a.MillimeterMercuryMinutesPerCubicCentimeter()
		cacheResult := a.MillimeterMercuryMinutesPerCubicCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimeterMercuryMinutesPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MillimeterMercurySecondsPerCubicMeter.
		// No expected conversion value provided for MillimeterMercurySecondsPerCubicMeter, verifying result is not NaN.
		result := a.MillimeterMercurySecondsPerCubicMeter()
		cacheResult := a.MillimeterMercurySecondsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimeterMercurySecondsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MillimeterMercuryMinutesPerCubicMeter.
		// No expected conversion value provided for MillimeterMercuryMinutesPerCubicMeter, verifying result is not NaN.
		result := a.MillimeterMercuryMinutesPerCubicMeter()
		cacheResult := a.MillimeterMercuryMinutesPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimeterMercuryMinutesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to WoodUnits.
		// No expected conversion value provided for WoodUnits, verifying result is not NaN.
		result := a.WoodUnits()
		cacheResult := a.WoodUnits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WoodUnits returned NaN")
		}
	}
	{
		// Test conversion to MegapascalSecondsPerCubicMeter.
		// No expected conversion value provided for MegapascalSecondsPerCubicMeter, verifying result is not NaN.
		result := a.MegapascalSecondsPerCubicMeter()
		cacheResult := a.MegapascalSecondsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegapascalSecondsPerCubicMeter returned NaN")
		}
	}
}

func TestFluidResistance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.FluidResistanceFactory{}
	a, err := factory.CreateFluidResistance(90, units.FluidResistancePascalSecondPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.FluidResistancePascalSecondPerCubicMeter {
		t.Errorf("expected default unit PascalSecondPerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.FluidResistancePascalSecondPerLiter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.FluidResistanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.FluidResistancePascalSecondPerCubicMeter {
		t.Errorf("expected unit PascalSecondPerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestFluidResistanceFactory_FromDto(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistancePascalSecondPerCubicMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.FluidResistanceDto{
        Value: math.NaN(),
        Unit:  units.FluidResistancePascalSecondPerCubicMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test PascalSecondPerLiter conversion
    pascal_seconds_per_literDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistancePascalSecondPerLiter,
    }
    
    var pascal_seconds_per_literResult *units.FluidResistance
    pascal_seconds_per_literResult, err = factory.FromDto(pascal_seconds_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with PascalSecondPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_seconds_per_literResult.Convert(units.FluidResistancePascalSecondPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalSecondPerLiter = %v, want %v", converted, 100)
    }
    // Test PascalMinutePerLiter conversion
    pascal_minutes_per_literDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistancePascalMinutePerLiter,
    }
    
    var pascal_minutes_per_literResult *units.FluidResistance
    pascal_minutes_per_literResult, err = factory.FromDto(pascal_minutes_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with PascalMinutePerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_minutes_per_literResult.Convert(units.FluidResistancePascalMinutePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalMinutePerLiter = %v, want %v", converted, 100)
    }
    // Test PascalSecondPerMilliliter conversion
    pascal_seconds_per_milliliterDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistancePascalSecondPerMilliliter,
    }
    
    var pascal_seconds_per_milliliterResult *units.FluidResistance
    pascal_seconds_per_milliliterResult, err = factory.FromDto(pascal_seconds_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with PascalSecondPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_seconds_per_milliliterResult.Convert(units.FluidResistancePascalSecondPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalSecondPerMilliliter = %v, want %v", converted, 100)
    }
    // Test PascalMinutePerMilliliter conversion
    pascal_minutes_per_milliliterDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistancePascalMinutePerMilliliter,
    }
    
    var pascal_minutes_per_milliliterResult *units.FluidResistance
    pascal_minutes_per_milliliterResult, err = factory.FromDto(pascal_minutes_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with PascalMinutePerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_minutes_per_milliliterResult.Convert(units.FluidResistancePascalMinutePerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalMinutePerMilliliter = %v, want %v", converted, 100)
    }
    // Test PascalSecondPerCubicMeter conversion
    pascal_seconds_per_cubic_meterDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistancePascalSecondPerCubicMeter,
    }
    
    var pascal_seconds_per_cubic_meterResult *units.FluidResistance
    pascal_seconds_per_cubic_meterResult, err = factory.FromDto(pascal_seconds_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with PascalSecondPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_seconds_per_cubic_meterResult.Convert(units.FluidResistancePascalSecondPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalSecondPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test PascalMinutePerCubicMeter conversion
    pascal_minutes_per_cubic_meterDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistancePascalMinutePerCubicMeter,
    }
    
    var pascal_minutes_per_cubic_meterResult *units.FluidResistance
    pascal_minutes_per_cubic_meterResult, err = factory.FromDto(pascal_minutes_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with PascalMinutePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_minutes_per_cubic_meterResult.Convert(units.FluidResistancePascalMinutePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalMinutePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test PascalSecondPerCubicCentimeter conversion
    pascal_seconds_per_cubic_centimeterDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistancePascalSecondPerCubicCentimeter,
    }
    
    var pascal_seconds_per_cubic_centimeterResult *units.FluidResistance
    pascal_seconds_per_cubic_centimeterResult, err = factory.FromDto(pascal_seconds_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with PascalSecondPerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_seconds_per_cubic_centimeterResult.Convert(units.FluidResistancePascalSecondPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalSecondPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test PascalMinutePerCubicCentimeter conversion
    pascal_minutes_per_cubic_centimeterDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistancePascalMinutePerCubicCentimeter,
    }
    
    var pascal_minutes_per_cubic_centimeterResult *units.FluidResistance
    pascal_minutes_per_cubic_centimeterResult, err = factory.FromDto(pascal_minutes_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with PascalMinutePerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_minutes_per_cubic_centimeterResult.Convert(units.FluidResistancePascalMinutePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalMinutePerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test DyneSecondPerCentimeterToTheFifth conversion
    dyne_seconds_per_centimeter_to_the_fifthDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistanceDyneSecondPerCentimeterToTheFifth,
    }
    
    var dyne_seconds_per_centimeter_to_the_fifthResult *units.FluidResistance
    dyne_seconds_per_centimeter_to_the_fifthResult, err = factory.FromDto(dyne_seconds_per_centimeter_to_the_fifthDto)
    if err != nil {
        t.Errorf("FromDto() with DyneSecondPerCentimeterToTheFifth returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = dyne_seconds_per_centimeter_to_the_fifthResult.Convert(units.FluidResistanceDyneSecondPerCentimeterToTheFifth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DyneSecondPerCentimeterToTheFifth = %v, want %v", converted, 100)
    }
    // Test MillimeterMercurySecondPerLiter conversion
    millimeter_mercury_seconds_per_literDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistanceMillimeterMercurySecondPerLiter,
    }
    
    var millimeter_mercury_seconds_per_literResult *units.FluidResistance
    millimeter_mercury_seconds_per_literResult, err = factory.FromDto(millimeter_mercury_seconds_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterMercurySecondPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_seconds_per_literResult.Convert(units.FluidResistanceMillimeterMercurySecondPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercurySecondPerLiter = %v, want %v", converted, 100)
    }
    // Test MillimeterMercuryMinutePerLiter conversion
    millimeter_mercury_minutes_per_literDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistanceMillimeterMercuryMinutePerLiter,
    }
    
    var millimeter_mercury_minutes_per_literResult *units.FluidResistance
    millimeter_mercury_minutes_per_literResult, err = factory.FromDto(millimeter_mercury_minutes_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterMercuryMinutePerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_minutes_per_literResult.Convert(units.FluidResistanceMillimeterMercuryMinutePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercuryMinutePerLiter = %v, want %v", converted, 100)
    }
    // Test MillimeterMercurySecondPerMilliliter conversion
    millimeter_mercury_seconds_per_milliliterDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistanceMillimeterMercurySecondPerMilliliter,
    }
    
    var millimeter_mercury_seconds_per_milliliterResult *units.FluidResistance
    millimeter_mercury_seconds_per_milliliterResult, err = factory.FromDto(millimeter_mercury_seconds_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterMercurySecondPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_seconds_per_milliliterResult.Convert(units.FluidResistanceMillimeterMercurySecondPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercurySecondPerMilliliter = %v, want %v", converted, 100)
    }
    // Test MillimeterMercuryMinutePerMilliliter conversion
    millimeter_mercury_minutes_per_milliliterDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistanceMillimeterMercuryMinutePerMilliliter,
    }
    
    var millimeter_mercury_minutes_per_milliliterResult *units.FluidResistance
    millimeter_mercury_minutes_per_milliliterResult, err = factory.FromDto(millimeter_mercury_minutes_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterMercuryMinutePerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_minutes_per_milliliterResult.Convert(units.FluidResistanceMillimeterMercuryMinutePerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercuryMinutePerMilliliter = %v, want %v", converted, 100)
    }
    // Test MillimeterMercurySecondPerCubicCentimeter conversion
    millimeter_mercury_seconds_per_cubic_centimeterDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistanceMillimeterMercurySecondPerCubicCentimeter,
    }
    
    var millimeter_mercury_seconds_per_cubic_centimeterResult *units.FluidResistance
    millimeter_mercury_seconds_per_cubic_centimeterResult, err = factory.FromDto(millimeter_mercury_seconds_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterMercurySecondPerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_seconds_per_cubic_centimeterResult.Convert(units.FluidResistanceMillimeterMercurySecondPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercurySecondPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test MillimeterMercuryMinutePerCubicCentimeter conversion
    millimeter_mercury_minutes_per_cubic_centimeterDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistanceMillimeterMercuryMinutePerCubicCentimeter,
    }
    
    var millimeter_mercury_minutes_per_cubic_centimeterResult *units.FluidResistance
    millimeter_mercury_minutes_per_cubic_centimeterResult, err = factory.FromDto(millimeter_mercury_minutes_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterMercuryMinutePerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_minutes_per_cubic_centimeterResult.Convert(units.FluidResistanceMillimeterMercuryMinutePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercuryMinutePerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test MillimeterMercurySecondPerCubicMeter conversion
    millimeter_mercury_seconds_per_cubic_meterDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistanceMillimeterMercurySecondPerCubicMeter,
    }
    
    var millimeter_mercury_seconds_per_cubic_meterResult *units.FluidResistance
    millimeter_mercury_seconds_per_cubic_meterResult, err = factory.FromDto(millimeter_mercury_seconds_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterMercurySecondPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_seconds_per_cubic_meterResult.Convert(units.FluidResistanceMillimeterMercurySecondPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercurySecondPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test MillimeterMercuryMinutePerCubicMeter conversion
    millimeter_mercury_minutes_per_cubic_meterDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistanceMillimeterMercuryMinutePerCubicMeter,
    }
    
    var millimeter_mercury_minutes_per_cubic_meterResult *units.FluidResistance
    millimeter_mercury_minutes_per_cubic_meterResult, err = factory.FromDto(millimeter_mercury_minutes_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterMercuryMinutePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_minutes_per_cubic_meterResult.Convert(units.FluidResistanceMillimeterMercuryMinutePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercuryMinutePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test WoodUnit conversion
    wood_unitsDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistanceWoodUnit,
    }
    
    var wood_unitsResult *units.FluidResistance
    wood_unitsResult, err = factory.FromDto(wood_unitsDto)
    if err != nil {
        t.Errorf("FromDto() with WoodUnit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = wood_unitsResult.Convert(units.FluidResistanceWoodUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WoodUnit = %v, want %v", converted, 100)
    }
    // Test MegapascalSecondPerCubicMeter conversion
    megapascal_seconds_per_cubic_meterDto := units.FluidResistanceDto{
        Value: 100,
        Unit:  units.FluidResistanceMegapascalSecondPerCubicMeter,
    }
    
    var megapascal_seconds_per_cubic_meterResult *units.FluidResistance
    megapascal_seconds_per_cubic_meterResult, err = factory.FromDto(megapascal_seconds_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MegapascalSecondPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapascal_seconds_per_cubic_meterResult.Convert(units.FluidResistanceMegapascalSecondPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapascalSecondPerCubicMeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.FluidResistanceDto{
        Value: 0,
        Unit:  units.FluidResistancePascalSecondPerCubicMeter,
    }
    
    var zeroResult *units.FluidResistance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestFluidResistanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "PascalSecondPerCubicMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "PascalSecondPerCubicMeter"}`)
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
    nanJSON, _ := json.Marshal(units.FluidResistanceDto{
        Value: nanValue,
        Unit:  units.FluidResistancePascalSecondPerCubicMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with PascalSecondPerLiter unit
    pascal_seconds_per_literJSON := []byte(`{"value": 100, "unit": "PascalSecondPerLiter"}`)
    pascal_seconds_per_literResult, err := factory.FromDtoJSON(pascal_seconds_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PascalSecondPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_seconds_per_literResult.Convert(units.FluidResistancePascalSecondPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalSecondPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with PascalMinutePerLiter unit
    pascal_minutes_per_literJSON := []byte(`{"value": 100, "unit": "PascalMinutePerLiter"}`)
    pascal_minutes_per_literResult, err := factory.FromDtoJSON(pascal_minutes_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PascalMinutePerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_minutes_per_literResult.Convert(units.FluidResistancePascalMinutePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalMinutePerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with PascalSecondPerMilliliter unit
    pascal_seconds_per_milliliterJSON := []byte(`{"value": 100, "unit": "PascalSecondPerMilliliter"}`)
    pascal_seconds_per_milliliterResult, err := factory.FromDtoJSON(pascal_seconds_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PascalSecondPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_seconds_per_milliliterResult.Convert(units.FluidResistancePascalSecondPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalSecondPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with PascalMinutePerMilliliter unit
    pascal_minutes_per_milliliterJSON := []byte(`{"value": 100, "unit": "PascalMinutePerMilliliter"}`)
    pascal_minutes_per_milliliterResult, err := factory.FromDtoJSON(pascal_minutes_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PascalMinutePerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_minutes_per_milliliterResult.Convert(units.FluidResistancePascalMinutePerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalMinutePerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with PascalSecondPerCubicMeter unit
    pascal_seconds_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "PascalSecondPerCubicMeter"}`)
    pascal_seconds_per_cubic_meterResult, err := factory.FromDtoJSON(pascal_seconds_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PascalSecondPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_seconds_per_cubic_meterResult.Convert(units.FluidResistancePascalSecondPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalSecondPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PascalMinutePerCubicMeter unit
    pascal_minutes_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "PascalMinutePerCubicMeter"}`)
    pascal_minutes_per_cubic_meterResult, err := factory.FromDtoJSON(pascal_minutes_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PascalMinutePerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_minutes_per_cubic_meterResult.Convert(units.FluidResistancePascalMinutePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalMinutePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PascalSecondPerCubicCentimeter unit
    pascal_seconds_per_cubic_centimeterJSON := []byte(`{"value": 100, "unit": "PascalSecondPerCubicCentimeter"}`)
    pascal_seconds_per_cubic_centimeterResult, err := factory.FromDtoJSON(pascal_seconds_per_cubic_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PascalSecondPerCubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_seconds_per_cubic_centimeterResult.Convert(units.FluidResistancePascalSecondPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalSecondPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with PascalMinutePerCubicCentimeter unit
    pascal_minutes_per_cubic_centimeterJSON := []byte(`{"value": 100, "unit": "PascalMinutePerCubicCentimeter"}`)
    pascal_minutes_per_cubic_centimeterResult, err := factory.FromDtoJSON(pascal_minutes_per_cubic_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PascalMinutePerCubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_minutes_per_cubic_centimeterResult.Convert(units.FluidResistancePascalMinutePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalMinutePerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with DyneSecondPerCentimeterToTheFifth unit
    dyne_seconds_per_centimeter_to_the_fifthJSON := []byte(`{"value": 100, "unit": "DyneSecondPerCentimeterToTheFifth"}`)
    dyne_seconds_per_centimeter_to_the_fifthResult, err := factory.FromDtoJSON(dyne_seconds_per_centimeter_to_the_fifthJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DyneSecondPerCentimeterToTheFifth unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = dyne_seconds_per_centimeter_to_the_fifthResult.Convert(units.FluidResistanceDyneSecondPerCentimeterToTheFifth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DyneSecondPerCentimeterToTheFifth = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterMercurySecondPerLiter unit
    millimeter_mercury_seconds_per_literJSON := []byte(`{"value": 100, "unit": "MillimeterMercurySecondPerLiter"}`)
    millimeter_mercury_seconds_per_literResult, err := factory.FromDtoJSON(millimeter_mercury_seconds_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterMercurySecondPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_seconds_per_literResult.Convert(units.FluidResistanceMillimeterMercurySecondPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercurySecondPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterMercuryMinutePerLiter unit
    millimeter_mercury_minutes_per_literJSON := []byte(`{"value": 100, "unit": "MillimeterMercuryMinutePerLiter"}`)
    millimeter_mercury_minutes_per_literResult, err := factory.FromDtoJSON(millimeter_mercury_minutes_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterMercuryMinutePerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_minutes_per_literResult.Convert(units.FluidResistanceMillimeterMercuryMinutePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercuryMinutePerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterMercurySecondPerMilliliter unit
    millimeter_mercury_seconds_per_milliliterJSON := []byte(`{"value": 100, "unit": "MillimeterMercurySecondPerMilliliter"}`)
    millimeter_mercury_seconds_per_milliliterResult, err := factory.FromDtoJSON(millimeter_mercury_seconds_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterMercurySecondPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_seconds_per_milliliterResult.Convert(units.FluidResistanceMillimeterMercurySecondPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercurySecondPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterMercuryMinutePerMilliliter unit
    millimeter_mercury_minutes_per_milliliterJSON := []byte(`{"value": 100, "unit": "MillimeterMercuryMinutePerMilliliter"}`)
    millimeter_mercury_minutes_per_milliliterResult, err := factory.FromDtoJSON(millimeter_mercury_minutes_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterMercuryMinutePerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_minutes_per_milliliterResult.Convert(units.FluidResistanceMillimeterMercuryMinutePerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercuryMinutePerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterMercurySecondPerCubicCentimeter unit
    millimeter_mercury_seconds_per_cubic_centimeterJSON := []byte(`{"value": 100, "unit": "MillimeterMercurySecondPerCubicCentimeter"}`)
    millimeter_mercury_seconds_per_cubic_centimeterResult, err := factory.FromDtoJSON(millimeter_mercury_seconds_per_cubic_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterMercurySecondPerCubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_seconds_per_cubic_centimeterResult.Convert(units.FluidResistanceMillimeterMercurySecondPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercurySecondPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterMercuryMinutePerCubicCentimeter unit
    millimeter_mercury_minutes_per_cubic_centimeterJSON := []byte(`{"value": 100, "unit": "MillimeterMercuryMinutePerCubicCentimeter"}`)
    millimeter_mercury_minutes_per_cubic_centimeterResult, err := factory.FromDtoJSON(millimeter_mercury_minutes_per_cubic_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterMercuryMinutePerCubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_minutes_per_cubic_centimeterResult.Convert(units.FluidResistanceMillimeterMercuryMinutePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercuryMinutePerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterMercurySecondPerCubicMeter unit
    millimeter_mercury_seconds_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "MillimeterMercurySecondPerCubicMeter"}`)
    millimeter_mercury_seconds_per_cubic_meterResult, err := factory.FromDtoJSON(millimeter_mercury_seconds_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterMercurySecondPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_seconds_per_cubic_meterResult.Convert(units.FluidResistanceMillimeterMercurySecondPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercurySecondPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterMercuryMinutePerCubicMeter unit
    millimeter_mercury_minutes_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "MillimeterMercuryMinutePerCubicMeter"}`)
    millimeter_mercury_minutes_per_cubic_meterResult, err := factory.FromDtoJSON(millimeter_mercury_minutes_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterMercuryMinutePerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeter_mercury_minutes_per_cubic_meterResult.Convert(units.FluidResistanceMillimeterMercuryMinutePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterMercuryMinutePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with WoodUnit unit
    wood_unitsJSON := []byte(`{"value": 100, "unit": "WoodUnit"}`)
    wood_unitsResult, err := factory.FromDtoJSON(wood_unitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WoodUnit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = wood_unitsResult.Convert(units.FluidResistanceWoodUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WoodUnit = %v, want %v", converted, 100)
    }
    // Test JSON with MegapascalSecondPerCubicMeter unit
    megapascal_seconds_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "MegapascalSecondPerCubicMeter"}`)
    megapascal_seconds_per_cubic_meterResult, err := factory.FromDtoJSON(megapascal_seconds_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapascalSecondPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapascal_seconds_per_cubic_meterResult.Convert(units.FluidResistanceMegapascalSecondPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapascalSecondPerCubicMeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "PascalSecondPerCubicMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromPascalSecondsPerLiter function
func TestFluidResistanceFactory_FromPascalSecondsPerLiter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPascalSecondsPerLiter(100)
    if err != nil {
        t.Errorf("FromPascalSecondsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistancePascalSecondPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPascalSecondsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPascalSecondsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromPascalSecondsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromPascalSecondsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromPascalSecondsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromPascalSecondsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromPascalSecondsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPascalSecondsPerLiter(0)
    if err != nil {
        t.Errorf("FromPascalSecondsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistancePascalSecondPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPascalSecondsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromPascalMinutesPerLiter function
func TestFluidResistanceFactory_FromPascalMinutesPerLiter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPascalMinutesPerLiter(100)
    if err != nil {
        t.Errorf("FromPascalMinutesPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistancePascalMinutePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPascalMinutesPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPascalMinutesPerLiter(math.NaN())
    if err == nil {
        t.Error("FromPascalMinutesPerLiter() with NaN value should return error")
    }

    _, err = factory.FromPascalMinutesPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromPascalMinutesPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromPascalMinutesPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromPascalMinutesPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPascalMinutesPerLiter(0)
    if err != nil {
        t.Errorf("FromPascalMinutesPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistancePascalMinutePerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPascalMinutesPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromPascalSecondsPerMilliliter function
func TestFluidResistanceFactory_FromPascalSecondsPerMilliliter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPascalSecondsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromPascalSecondsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistancePascalSecondPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPascalSecondsPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPascalSecondsPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromPascalSecondsPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromPascalSecondsPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromPascalSecondsPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromPascalSecondsPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromPascalSecondsPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPascalSecondsPerMilliliter(0)
    if err != nil {
        t.Errorf("FromPascalSecondsPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistancePascalSecondPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPascalSecondsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromPascalMinutesPerMilliliter function
func TestFluidResistanceFactory_FromPascalMinutesPerMilliliter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPascalMinutesPerMilliliter(100)
    if err != nil {
        t.Errorf("FromPascalMinutesPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistancePascalMinutePerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPascalMinutesPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPascalMinutesPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromPascalMinutesPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromPascalMinutesPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromPascalMinutesPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromPascalMinutesPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromPascalMinutesPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPascalMinutesPerMilliliter(0)
    if err != nil {
        t.Errorf("FromPascalMinutesPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistancePascalMinutePerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPascalMinutesPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromPascalSecondsPerCubicMeter function
func TestFluidResistanceFactory_FromPascalSecondsPerCubicMeter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPascalSecondsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromPascalSecondsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistancePascalSecondPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPascalSecondsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPascalSecondsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromPascalSecondsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromPascalSecondsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromPascalSecondsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromPascalSecondsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromPascalSecondsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPascalSecondsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromPascalSecondsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistancePascalSecondPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPascalSecondsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPascalMinutesPerCubicMeter function
func TestFluidResistanceFactory_FromPascalMinutesPerCubicMeter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPascalMinutesPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromPascalMinutesPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistancePascalMinutePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPascalMinutesPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPascalMinutesPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromPascalMinutesPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromPascalMinutesPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromPascalMinutesPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromPascalMinutesPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromPascalMinutesPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPascalMinutesPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromPascalMinutesPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistancePascalMinutePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPascalMinutesPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPascalSecondsPerCubicCentimeter function
func TestFluidResistanceFactory_FromPascalSecondsPerCubicCentimeter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPascalSecondsPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromPascalSecondsPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistancePascalSecondPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPascalSecondsPerCubicCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPascalSecondsPerCubicCentimeter(math.NaN())
    if err == nil {
        t.Error("FromPascalSecondsPerCubicCentimeter() with NaN value should return error")
    }

    _, err = factory.FromPascalSecondsPerCubicCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromPascalSecondsPerCubicCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromPascalSecondsPerCubicCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromPascalSecondsPerCubicCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPascalSecondsPerCubicCentimeter(0)
    if err != nil {
        t.Errorf("FromPascalSecondsPerCubicCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistancePascalSecondPerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPascalSecondsPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPascalMinutesPerCubicCentimeter function
func TestFluidResistanceFactory_FromPascalMinutesPerCubicCentimeter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPascalMinutesPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromPascalMinutesPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistancePascalMinutePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPascalMinutesPerCubicCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPascalMinutesPerCubicCentimeter(math.NaN())
    if err == nil {
        t.Error("FromPascalMinutesPerCubicCentimeter() with NaN value should return error")
    }

    _, err = factory.FromPascalMinutesPerCubicCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromPascalMinutesPerCubicCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromPascalMinutesPerCubicCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromPascalMinutesPerCubicCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPascalMinutesPerCubicCentimeter(0)
    if err != nil {
        t.Errorf("FromPascalMinutesPerCubicCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistancePascalMinutePerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPascalMinutesPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromDyneSecondsPerCentimeterToTheFifth function
func TestFluidResistanceFactory_FromDyneSecondsPerCentimeterToTheFifth(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDyneSecondsPerCentimeterToTheFifth(100)
    if err != nil {
        t.Errorf("FromDyneSecondsPerCentimeterToTheFifth() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistanceDyneSecondPerCentimeterToTheFifth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDyneSecondsPerCentimeterToTheFifth() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDyneSecondsPerCentimeterToTheFifth(math.NaN())
    if err == nil {
        t.Error("FromDyneSecondsPerCentimeterToTheFifth() with NaN value should return error")
    }

    _, err = factory.FromDyneSecondsPerCentimeterToTheFifth(math.Inf(1))
    if err == nil {
        t.Error("FromDyneSecondsPerCentimeterToTheFifth() with +Inf value should return error")
    }

    _, err = factory.FromDyneSecondsPerCentimeterToTheFifth(math.Inf(-1))
    if err == nil {
        t.Error("FromDyneSecondsPerCentimeterToTheFifth() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDyneSecondsPerCentimeterToTheFifth(0)
    if err != nil {
        t.Errorf("FromDyneSecondsPerCentimeterToTheFifth() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistanceDyneSecondPerCentimeterToTheFifth)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDyneSecondsPerCentimeterToTheFifth() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimeterMercurySecondsPerLiter function
func TestFluidResistanceFactory_FromMillimeterMercurySecondsPerLiter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimeterMercurySecondsPerLiter(100)
    if err != nil {
        t.Errorf("FromMillimeterMercurySecondsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistanceMillimeterMercurySecondPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimeterMercurySecondsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimeterMercurySecondsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromMillimeterMercurySecondsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromMillimeterMercurySecondsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMillimeterMercurySecondsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromMillimeterMercurySecondsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimeterMercurySecondsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimeterMercurySecondsPerLiter(0)
    if err != nil {
        t.Errorf("FromMillimeterMercurySecondsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistanceMillimeterMercurySecondPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimeterMercurySecondsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimeterMercuryMinutesPerLiter function
func TestFluidResistanceFactory_FromMillimeterMercuryMinutesPerLiter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimeterMercuryMinutesPerLiter(100)
    if err != nil {
        t.Errorf("FromMillimeterMercuryMinutesPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistanceMillimeterMercuryMinutePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimeterMercuryMinutesPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimeterMercuryMinutesPerLiter(math.NaN())
    if err == nil {
        t.Error("FromMillimeterMercuryMinutesPerLiter() with NaN value should return error")
    }

    _, err = factory.FromMillimeterMercuryMinutesPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMillimeterMercuryMinutesPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromMillimeterMercuryMinutesPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimeterMercuryMinutesPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimeterMercuryMinutesPerLiter(0)
    if err != nil {
        t.Errorf("FromMillimeterMercuryMinutesPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistanceMillimeterMercuryMinutePerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimeterMercuryMinutesPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimeterMercurySecondsPerMilliliter function
func TestFluidResistanceFactory_FromMillimeterMercurySecondsPerMilliliter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimeterMercurySecondsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromMillimeterMercurySecondsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistanceMillimeterMercurySecondPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimeterMercurySecondsPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimeterMercurySecondsPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromMillimeterMercurySecondsPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromMillimeterMercurySecondsPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromMillimeterMercurySecondsPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromMillimeterMercurySecondsPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimeterMercurySecondsPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimeterMercurySecondsPerMilliliter(0)
    if err != nil {
        t.Errorf("FromMillimeterMercurySecondsPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistanceMillimeterMercurySecondPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimeterMercurySecondsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimeterMercuryMinutesPerMilliliter function
func TestFluidResistanceFactory_FromMillimeterMercuryMinutesPerMilliliter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimeterMercuryMinutesPerMilliliter(100)
    if err != nil {
        t.Errorf("FromMillimeterMercuryMinutesPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistanceMillimeterMercuryMinutePerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimeterMercuryMinutesPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimeterMercuryMinutesPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromMillimeterMercuryMinutesPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromMillimeterMercuryMinutesPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromMillimeterMercuryMinutesPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromMillimeterMercuryMinutesPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimeterMercuryMinutesPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimeterMercuryMinutesPerMilliliter(0)
    if err != nil {
        t.Errorf("FromMillimeterMercuryMinutesPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistanceMillimeterMercuryMinutePerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimeterMercuryMinutesPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimeterMercurySecondsPerCubicCentimeter function
func TestFluidResistanceFactory_FromMillimeterMercurySecondsPerCubicCentimeter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimeterMercurySecondsPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromMillimeterMercurySecondsPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistanceMillimeterMercurySecondPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimeterMercurySecondsPerCubicCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimeterMercurySecondsPerCubicCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMillimeterMercurySecondsPerCubicCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMillimeterMercurySecondsPerCubicCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMillimeterMercurySecondsPerCubicCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMillimeterMercurySecondsPerCubicCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimeterMercurySecondsPerCubicCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimeterMercurySecondsPerCubicCentimeter(0)
    if err != nil {
        t.Errorf("FromMillimeterMercurySecondsPerCubicCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistanceMillimeterMercurySecondPerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimeterMercurySecondsPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimeterMercuryMinutesPerCubicCentimeter function
func TestFluidResistanceFactory_FromMillimeterMercuryMinutesPerCubicCentimeter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimeterMercuryMinutesPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromMillimeterMercuryMinutesPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistanceMillimeterMercuryMinutePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimeterMercuryMinutesPerCubicCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimeterMercuryMinutesPerCubicCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMillimeterMercuryMinutesPerCubicCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMillimeterMercuryMinutesPerCubicCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMillimeterMercuryMinutesPerCubicCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMillimeterMercuryMinutesPerCubicCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimeterMercuryMinutesPerCubicCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimeterMercuryMinutesPerCubicCentimeter(0)
    if err != nil {
        t.Errorf("FromMillimeterMercuryMinutesPerCubicCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistanceMillimeterMercuryMinutePerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimeterMercuryMinutesPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimeterMercurySecondsPerCubicMeter function
func TestFluidResistanceFactory_FromMillimeterMercurySecondsPerCubicMeter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimeterMercurySecondsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMillimeterMercurySecondsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistanceMillimeterMercurySecondPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimeterMercurySecondsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimeterMercurySecondsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromMillimeterMercurySecondsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromMillimeterMercurySecondsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMillimeterMercurySecondsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromMillimeterMercurySecondsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimeterMercurySecondsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimeterMercurySecondsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromMillimeterMercurySecondsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistanceMillimeterMercurySecondPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimeterMercurySecondsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimeterMercuryMinutesPerCubicMeter function
func TestFluidResistanceFactory_FromMillimeterMercuryMinutesPerCubicMeter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimeterMercuryMinutesPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMillimeterMercuryMinutesPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistanceMillimeterMercuryMinutePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimeterMercuryMinutesPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimeterMercuryMinutesPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromMillimeterMercuryMinutesPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromMillimeterMercuryMinutesPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMillimeterMercuryMinutesPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromMillimeterMercuryMinutesPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimeterMercuryMinutesPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimeterMercuryMinutesPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromMillimeterMercuryMinutesPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistanceMillimeterMercuryMinutePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimeterMercuryMinutesPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromWoodUnits function
func TestFluidResistanceFactory_FromWoodUnits(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWoodUnits(100)
    if err != nil {
        t.Errorf("FromWoodUnits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistanceWoodUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWoodUnits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWoodUnits(math.NaN())
    if err == nil {
        t.Error("FromWoodUnits() with NaN value should return error")
    }

    _, err = factory.FromWoodUnits(math.Inf(1))
    if err == nil {
        t.Error("FromWoodUnits() with +Inf value should return error")
    }

    _, err = factory.FromWoodUnits(math.Inf(-1))
    if err == nil {
        t.Error("FromWoodUnits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWoodUnits(0)
    if err != nil {
        t.Errorf("FromWoodUnits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistanceWoodUnit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWoodUnits() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapascalSecondsPerCubicMeter function
func TestFluidResistanceFactory_FromMegapascalSecondsPerCubicMeter(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapascalSecondsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMegapascalSecondsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FluidResistanceMegapascalSecondPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapascalSecondsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapascalSecondsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromMegapascalSecondsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromMegapascalSecondsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMegapascalSecondsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromMegapascalSecondsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapascalSecondsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapascalSecondsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromMegapascalSecondsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FluidResistanceMegapascalSecondPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapascalSecondsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}

func TestFluidResistanceToString(t *testing.T) {
	factory := units.FluidResistanceFactory{}
	a, err := factory.CreateFluidResistance(45, units.FluidResistancePascalSecondPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.FluidResistancePascalSecondPerCubicMeter, 2)
	expected := "45.00 " + units.GetFluidResistanceAbbreviation(units.FluidResistancePascalSecondPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.FluidResistancePascalSecondPerCubicMeter, -1)
	expected = "45 " + units.GetFluidResistanceAbbreviation(units.FluidResistancePascalSecondPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestFluidResistance_EqualityAndComparison(t *testing.T) {
	factory := units.FluidResistanceFactory{}
	a1, _ := factory.CreateFluidResistance(60, units.FluidResistancePascalSecondPerCubicMeter)
	a2, _ := factory.CreateFluidResistance(60, units.FluidResistancePascalSecondPerCubicMeter)
	a3, _ := factory.CreateFluidResistance(120, units.FluidResistancePascalSecondPerCubicMeter)

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

func TestFluidResistance_Arithmetic(t *testing.T) {
	factory := units.FluidResistanceFactory{}
	a1, _ := factory.CreateFluidResistance(30, units.FluidResistancePascalSecondPerCubicMeter)
	a2, _ := factory.CreateFluidResistance(45, units.FluidResistancePascalSecondPerCubicMeter)

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


func TestGetFluidResistanceAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.FluidResistanceUnits
        want string
    }{
        {
            name: "PascalSecondPerLiter abbreviation",
            unit: units.FluidResistancePascalSecondPerLiter,
            want: "Pa·s/l",
        },
        {
            name: "PascalMinutePerLiter abbreviation",
            unit: units.FluidResistancePascalMinutePerLiter,
            want: "Pa·min/l",
        },
        {
            name: "PascalSecondPerMilliliter abbreviation",
            unit: units.FluidResistancePascalSecondPerMilliliter,
            want: "Pa·s/ml",
        },
        {
            name: "PascalMinutePerMilliliter abbreviation",
            unit: units.FluidResistancePascalMinutePerMilliliter,
            want: "Pa·min/ml",
        },
        {
            name: "PascalSecondPerCubicMeter abbreviation",
            unit: units.FluidResistancePascalSecondPerCubicMeter,
            want: "Pa·s/m³",
        },
        {
            name: "PascalMinutePerCubicMeter abbreviation",
            unit: units.FluidResistancePascalMinutePerCubicMeter,
            want: "Pa·min/m³",
        },
        {
            name: "PascalSecondPerCubicCentimeter abbreviation",
            unit: units.FluidResistancePascalSecondPerCubicCentimeter,
            want: "Pa·s/cm³",
        },
        {
            name: "PascalMinutePerCubicCentimeter abbreviation",
            unit: units.FluidResistancePascalMinutePerCubicCentimeter,
            want: "Pa·min/cm³",
        },
        {
            name: "DyneSecondPerCentimeterToTheFifth abbreviation",
            unit: units.FluidResistanceDyneSecondPerCentimeterToTheFifth,
            want: "dyn·s/cm⁵",
        },
        {
            name: "MillimeterMercurySecondPerLiter abbreviation",
            unit: units.FluidResistanceMillimeterMercurySecondPerLiter,
            want: "mmHg·s/l",
        },
        {
            name: "MillimeterMercuryMinutePerLiter abbreviation",
            unit: units.FluidResistanceMillimeterMercuryMinutePerLiter,
            want: "mmHg·min/l",
        },
        {
            name: "MillimeterMercurySecondPerMilliliter abbreviation",
            unit: units.FluidResistanceMillimeterMercurySecondPerMilliliter,
            want: "mmHg·s/ml",
        },
        {
            name: "MillimeterMercuryMinutePerMilliliter abbreviation",
            unit: units.FluidResistanceMillimeterMercuryMinutePerMilliliter,
            want: "mmHg·min/ml",
        },
        {
            name: "MillimeterMercurySecondPerCubicCentimeter abbreviation",
            unit: units.FluidResistanceMillimeterMercurySecondPerCubicCentimeter,
            want: "mmHg·s/cm³",
        },
        {
            name: "MillimeterMercuryMinutePerCubicCentimeter abbreviation",
            unit: units.FluidResistanceMillimeterMercuryMinutePerCubicCentimeter,
            want: "mmHg·min/cm³",
        },
        {
            name: "MillimeterMercurySecondPerCubicMeter abbreviation",
            unit: units.FluidResistanceMillimeterMercurySecondPerCubicMeter,
            want: "mmHg·s/m³",
        },
        {
            name: "MillimeterMercuryMinutePerCubicMeter abbreviation",
            unit: units.FluidResistanceMillimeterMercuryMinutePerCubicMeter,
            want: "mmHg·min/m³",
        },
        {
            name: "WoodUnit abbreviation",
            unit: units.FluidResistanceWoodUnit,
            want: "WU",
        },
        {
            name: "MegapascalSecondPerCubicMeter abbreviation",
            unit: units.FluidResistanceMegapascalSecondPerCubicMeter,
            want: "MPa·s/m³",
        },
        {
            name: "invalid unit",
            unit: units.FluidResistanceUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetFluidResistanceAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetFluidResistanceAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestFluidResistance_String(t *testing.T) {
    factory := units.FluidResistanceFactory{}
    
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
            unit, err := factory.CreateFluidResistance(tt.value, units.FluidResistancePascalSecondPerCubicMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("FluidResistance.String() = %v, want %v", got, tt.want)
            }
        })
    }
}