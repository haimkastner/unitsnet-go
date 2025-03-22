// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAccelerationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MeterPerSecondSquared"}`
	
	factory := units.AccelerationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AccelerationMeterPerSecondSquared {
		t.Errorf("expected unit %v, got %v", units.AccelerationMeterPerSecondSquared, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MeterPerSecondSquared"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAccelerationDto_ToJSON(t *testing.T) {
	dto := units.AccelerationDto{
		Value: 45,
		Unit:  units.AccelerationMeterPerSecondSquared,
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
	if result["unit"].(string) != string(units.AccelerationMeterPerSecondSquared) {
		t.Errorf("expected unit %s, got %v", units.AccelerationMeterPerSecondSquared, result["unit"])
	}
}

func TestNewAcceleration_InvalidValue(t *testing.T) {
	factory := units.AccelerationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAcceleration(math.NaN(), units.AccelerationMeterPerSecondSquared)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAcceleration(math.Inf(1), units.AccelerationMeterPerSecondSquared)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAccelerationConversions(t *testing.T) {
	factory := units.AccelerationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAcceleration(180, units.AccelerationMeterPerSecondSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MetersPerSecondSquared.
		// No expected conversion value provided for MetersPerSecondSquared, verifying result is not NaN.
		result := a.MetersPerSecondSquared()
		cacheResult := a.MetersPerSecondSquared()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MetersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to InchesPerSecondSquared.
		// No expected conversion value provided for InchesPerSecondSquared, verifying result is not NaN.
		result := a.InchesPerSecondSquared()
		cacheResult := a.InchesPerSecondSquared()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InchesPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to FeetPerSecondSquared.
		// No expected conversion value provided for FeetPerSecondSquared, verifying result is not NaN.
		result := a.FeetPerSecondSquared()
		cacheResult := a.FeetPerSecondSquared()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to FeetPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to KnotsPerSecond.
		// No expected conversion value provided for KnotsPerSecond, verifying result is not NaN.
		result := a.KnotsPerSecond()
		cacheResult := a.KnotsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KnotsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KnotsPerMinute.
		// No expected conversion value provided for KnotsPerMinute, verifying result is not NaN.
		result := a.KnotsPerMinute()
		cacheResult := a.KnotsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KnotsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to KnotsPerHour.
		// No expected conversion value provided for KnotsPerHour, verifying result is not NaN.
		result := a.KnotsPerHour()
		cacheResult := a.KnotsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KnotsPerHour returned NaN")
		}
	}
	{
		// Test conversion to StandardGravity.
		// No expected conversion value provided for StandardGravity, verifying result is not NaN.
		result := a.StandardGravity()
		cacheResult := a.StandardGravity()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to StandardGravity returned NaN")
		}
	}
	{
		// Test conversion to NanometersPerSecondSquared.
		// No expected conversion value provided for NanometersPerSecondSquared, verifying result is not NaN.
		result := a.NanometersPerSecondSquared()
		cacheResult := a.NanometersPerSecondSquared()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanometersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to MicrometersPerSecondSquared.
		// No expected conversion value provided for MicrometersPerSecondSquared, verifying result is not NaN.
		result := a.MicrometersPerSecondSquared()
		cacheResult := a.MicrometersPerSecondSquared()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrometersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to MillimetersPerSecondSquared.
		// No expected conversion value provided for MillimetersPerSecondSquared, verifying result is not NaN.
		result := a.MillimetersPerSecondSquared()
		cacheResult := a.MillimetersPerSecondSquared()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimetersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to CentimetersPerSecondSquared.
		// No expected conversion value provided for CentimetersPerSecondSquared, verifying result is not NaN.
		result := a.CentimetersPerSecondSquared()
		cacheResult := a.CentimetersPerSecondSquared()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentimetersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to DecimetersPerSecondSquared.
		// No expected conversion value provided for DecimetersPerSecondSquared, verifying result is not NaN.
		result := a.DecimetersPerSecondSquared()
		cacheResult := a.DecimetersPerSecondSquared()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecimetersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to KilometersPerSecondSquared.
		// No expected conversion value provided for KilometersPerSecondSquared, verifying result is not NaN.
		result := a.KilometersPerSecondSquared()
		cacheResult := a.KilometersPerSecondSquared()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilometersPerSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to MillistandardGravity.
		// No expected conversion value provided for MillistandardGravity, verifying result is not NaN.
		result := a.MillistandardGravity()
		cacheResult := a.MillistandardGravity()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillistandardGravity returned NaN")
		}
	}
}

func TestAcceleration_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AccelerationFactory{}
	a, err := factory.CreateAcceleration(90, units.AccelerationMeterPerSecondSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AccelerationMeterPerSecondSquared {
		t.Errorf("expected default unit MeterPerSecondSquared, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AccelerationMeterPerSecondSquared
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AccelerationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AccelerationMeterPerSecondSquared {
		t.Errorf("expected unit MeterPerSecondSquared, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAccelerationFactory_FromDto(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationMeterPerSecondSquared,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.AccelerationDto{
        Value: math.NaN(),
        Unit:  units.AccelerationMeterPerSecondSquared,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test MeterPerSecondSquared conversion
    meters_per_second_squaredDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationMeterPerSecondSquared,
    }
    
    var meters_per_second_squaredResult *units.Acceleration
    meters_per_second_squaredResult, err = factory.FromDto(meters_per_second_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with MeterPerSecondSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_per_second_squaredResult.Convert(units.AccelerationMeterPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test InchPerSecondSquared conversion
    inches_per_second_squaredDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationInchPerSecondSquared,
    }
    
    var inches_per_second_squaredResult *units.Acceleration
    inches_per_second_squaredResult, err = factory.FromDto(inches_per_second_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with InchPerSecondSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_per_second_squaredResult.Convert(units.AccelerationInchPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test FootPerSecondSquared conversion
    feet_per_second_squaredDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationFootPerSecondSquared,
    }
    
    var feet_per_second_squaredResult *units.Acceleration
    feet_per_second_squaredResult, err = factory.FromDto(feet_per_second_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with FootPerSecondSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_per_second_squaredResult.Convert(units.AccelerationFootPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test KnotPerSecond conversion
    knots_per_secondDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationKnotPerSecond,
    }
    
    var knots_per_secondResult *units.Acceleration
    knots_per_secondResult, err = factory.FromDto(knots_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KnotPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = knots_per_secondResult.Convert(units.AccelerationKnotPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KnotPerSecond = %v, want %v", converted, 100)
    }
    // Test KnotPerMinute conversion
    knots_per_minuteDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationKnotPerMinute,
    }
    
    var knots_per_minuteResult *units.Acceleration
    knots_per_minuteResult, err = factory.FromDto(knots_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with KnotPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = knots_per_minuteResult.Convert(units.AccelerationKnotPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KnotPerMinute = %v, want %v", converted, 100)
    }
    // Test KnotPerHour conversion
    knots_per_hourDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationKnotPerHour,
    }
    
    var knots_per_hourResult *units.Acceleration
    knots_per_hourResult, err = factory.FromDto(knots_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with KnotPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = knots_per_hourResult.Convert(units.AccelerationKnotPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KnotPerHour = %v, want %v", converted, 100)
    }
    // Test StandardGravity conversion
    standard_gravityDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationStandardGravity,
    }
    
    var standard_gravityResult *units.Acceleration
    standard_gravityResult, err = factory.FromDto(standard_gravityDto)
    if err != nil {
        t.Errorf("FromDto() with StandardGravity returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_gravityResult.Convert(units.AccelerationStandardGravity)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardGravity = %v, want %v", converted, 100)
    }
    // Test NanometerPerSecondSquared conversion
    nanometers_per_second_squaredDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationNanometerPerSecondSquared,
    }
    
    var nanometers_per_second_squaredResult *units.Acceleration
    nanometers_per_second_squaredResult, err = factory.FromDto(nanometers_per_second_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with NanometerPerSecondSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanometers_per_second_squaredResult.Convert(units.AccelerationNanometerPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanometerPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test MicrometerPerSecondSquared conversion
    micrometers_per_second_squaredDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationMicrometerPerSecondSquared,
    }
    
    var micrometers_per_second_squaredResult *units.Acceleration
    micrometers_per_second_squaredResult, err = factory.FromDto(micrometers_per_second_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with MicrometerPerSecondSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrometers_per_second_squaredResult.Convert(units.AccelerationMicrometerPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrometerPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test MillimeterPerSecondSquared conversion
    millimeters_per_second_squaredDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationMillimeterPerSecondSquared,
    }
    
    var millimeters_per_second_squaredResult *units.Acceleration
    millimeters_per_second_squaredResult, err = factory.FromDto(millimeters_per_second_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterPerSecondSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_per_second_squaredResult.Convert(units.AccelerationMillimeterPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test CentimeterPerSecondSquared conversion
    centimeters_per_second_squaredDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationCentimeterPerSecondSquared,
    }
    
    var centimeters_per_second_squaredResult *units.Acceleration
    centimeters_per_second_squaredResult, err = factory.FromDto(centimeters_per_second_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with CentimeterPerSecondSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_per_second_squaredResult.Convert(units.AccelerationCentimeterPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test DecimeterPerSecondSquared conversion
    decimeters_per_second_squaredDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationDecimeterPerSecondSquared,
    }
    
    var decimeters_per_second_squaredResult *units.Acceleration
    decimeters_per_second_squaredResult, err = factory.FromDto(decimeters_per_second_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with DecimeterPerSecondSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimeters_per_second_squaredResult.Convert(units.AccelerationDecimeterPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimeterPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test KilometerPerSecondSquared conversion
    kilometers_per_second_squaredDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationKilometerPerSecondSquared,
    }
    
    var kilometers_per_second_squaredResult *units.Acceleration
    kilometers_per_second_squaredResult, err = factory.FromDto(kilometers_per_second_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with KilometerPerSecondSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometers_per_second_squaredResult.Convert(units.AccelerationKilometerPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilometerPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test MillistandardGravity conversion
    millistandard_gravityDto := units.AccelerationDto{
        Value: 100,
        Unit:  units.AccelerationMillistandardGravity,
    }
    
    var millistandard_gravityResult *units.Acceleration
    millistandard_gravityResult, err = factory.FromDto(millistandard_gravityDto)
    if err != nil {
        t.Errorf("FromDto() with MillistandardGravity returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millistandard_gravityResult.Convert(units.AccelerationMillistandardGravity)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillistandardGravity = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.AccelerationDto{
        Value: 0,
        Unit:  units.AccelerationMeterPerSecondSquared,
    }
    
    var zeroResult *units.Acceleration
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestAccelerationFactory_FromDtoJSON(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "MeterPerSecondSquared"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "MeterPerSecondSquared"}`)
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
    nanJSON, _ := json.Marshal(units.AccelerationDto{
        Value: nanValue,
        Unit:  units.AccelerationMeterPerSecondSquared,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with MeterPerSecondSquared unit
    meters_per_second_squaredJSON := []byte(`{"value": 100, "unit": "MeterPerSecondSquared"}`)
    meters_per_second_squaredResult, err := factory.FromDtoJSON(meters_per_second_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeterPerSecondSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_per_second_squaredResult.Convert(units.AccelerationMeterPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test JSON with InchPerSecondSquared unit
    inches_per_second_squaredJSON := []byte(`{"value": 100, "unit": "InchPerSecondSquared"}`)
    inches_per_second_squaredResult, err := factory.FromDtoJSON(inches_per_second_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InchPerSecondSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_per_second_squaredResult.Convert(units.AccelerationInchPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test JSON with FootPerSecondSquared unit
    feet_per_second_squaredJSON := []byte(`{"value": 100, "unit": "FootPerSecondSquared"}`)
    feet_per_second_squaredResult, err := factory.FromDtoJSON(feet_per_second_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FootPerSecondSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_per_second_squaredResult.Convert(units.AccelerationFootPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test JSON with KnotPerSecond unit
    knots_per_secondJSON := []byte(`{"value": 100, "unit": "KnotPerSecond"}`)
    knots_per_secondResult, err := factory.FromDtoJSON(knots_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KnotPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = knots_per_secondResult.Convert(units.AccelerationKnotPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KnotPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KnotPerMinute unit
    knots_per_minuteJSON := []byte(`{"value": 100, "unit": "KnotPerMinute"}`)
    knots_per_minuteResult, err := factory.FromDtoJSON(knots_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KnotPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = knots_per_minuteResult.Convert(units.AccelerationKnotPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KnotPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with KnotPerHour unit
    knots_per_hourJSON := []byte(`{"value": 100, "unit": "KnotPerHour"}`)
    knots_per_hourResult, err := factory.FromDtoJSON(knots_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KnotPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = knots_per_hourResult.Convert(units.AccelerationKnotPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KnotPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with StandardGravity unit
    standard_gravityJSON := []byte(`{"value": 100, "unit": "StandardGravity"}`)
    standard_gravityResult, err := factory.FromDtoJSON(standard_gravityJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with StandardGravity unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_gravityResult.Convert(units.AccelerationStandardGravity)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardGravity = %v, want %v", converted, 100)
    }
    // Test JSON with NanometerPerSecondSquared unit
    nanometers_per_second_squaredJSON := []byte(`{"value": 100, "unit": "NanometerPerSecondSquared"}`)
    nanometers_per_second_squaredResult, err := factory.FromDtoJSON(nanometers_per_second_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanometerPerSecondSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanometers_per_second_squaredResult.Convert(units.AccelerationNanometerPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanometerPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test JSON with MicrometerPerSecondSquared unit
    micrometers_per_second_squaredJSON := []byte(`{"value": 100, "unit": "MicrometerPerSecondSquared"}`)
    micrometers_per_second_squaredResult, err := factory.FromDtoJSON(micrometers_per_second_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrometerPerSecondSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrometers_per_second_squaredResult.Convert(units.AccelerationMicrometerPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrometerPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterPerSecondSquared unit
    millimeters_per_second_squaredJSON := []byte(`{"value": 100, "unit": "MillimeterPerSecondSquared"}`)
    millimeters_per_second_squaredResult, err := factory.FromDtoJSON(millimeters_per_second_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterPerSecondSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_per_second_squaredResult.Convert(units.AccelerationMillimeterPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test JSON with CentimeterPerSecondSquared unit
    centimeters_per_second_squaredJSON := []byte(`{"value": 100, "unit": "CentimeterPerSecondSquared"}`)
    centimeters_per_second_squaredResult, err := factory.FromDtoJSON(centimeters_per_second_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentimeterPerSecondSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_per_second_squaredResult.Convert(units.AccelerationCentimeterPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test JSON with DecimeterPerSecondSquared unit
    decimeters_per_second_squaredJSON := []byte(`{"value": 100, "unit": "DecimeterPerSecondSquared"}`)
    decimeters_per_second_squaredResult, err := factory.FromDtoJSON(decimeters_per_second_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecimeterPerSecondSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimeters_per_second_squaredResult.Convert(units.AccelerationDecimeterPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimeterPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test JSON with KilometerPerSecondSquared unit
    kilometers_per_second_squaredJSON := []byte(`{"value": 100, "unit": "KilometerPerSecondSquared"}`)
    kilometers_per_second_squaredResult, err := factory.FromDtoJSON(kilometers_per_second_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilometerPerSecondSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometers_per_second_squaredResult.Convert(units.AccelerationKilometerPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilometerPerSecondSquared = %v, want %v", converted, 100)
    }
    // Test JSON with MillistandardGravity unit
    millistandard_gravityJSON := []byte(`{"value": 100, "unit": "MillistandardGravity"}`)
    millistandard_gravityResult, err := factory.FromDtoJSON(millistandard_gravityJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillistandardGravity unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millistandard_gravityResult.Convert(units.AccelerationMillistandardGravity)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillistandardGravity = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "MeterPerSecondSquared"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromMetersPerSecondSquared function
func TestAccelerationFactory_FromMetersPerSecondSquared(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMetersPerSecondSquared(100)
    if err != nil {
        t.Errorf("FromMetersPerSecondSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationMeterPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMetersPerSecondSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMetersPerSecondSquared(math.NaN())
    if err == nil {
        t.Error("FromMetersPerSecondSquared() with NaN value should return error")
    }

    _, err = factory.FromMetersPerSecondSquared(math.Inf(1))
    if err == nil {
        t.Error("FromMetersPerSecondSquared() with +Inf value should return error")
    }

    _, err = factory.FromMetersPerSecondSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromMetersPerSecondSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMetersPerSecondSquared(0)
    if err != nil {
        t.Errorf("FromMetersPerSecondSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationMeterPerSecondSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMetersPerSecondSquared() with zero value = %v, want 0", converted)
    }
}
// Test FromInchesPerSecondSquared function
func TestAccelerationFactory_FromInchesPerSecondSquared(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInchesPerSecondSquared(100)
    if err != nil {
        t.Errorf("FromInchesPerSecondSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationInchPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInchesPerSecondSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInchesPerSecondSquared(math.NaN())
    if err == nil {
        t.Error("FromInchesPerSecondSquared() with NaN value should return error")
    }

    _, err = factory.FromInchesPerSecondSquared(math.Inf(1))
    if err == nil {
        t.Error("FromInchesPerSecondSquared() with +Inf value should return error")
    }

    _, err = factory.FromInchesPerSecondSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromInchesPerSecondSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInchesPerSecondSquared(0)
    if err != nil {
        t.Errorf("FromInchesPerSecondSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationInchPerSecondSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInchesPerSecondSquared() with zero value = %v, want 0", converted)
    }
}
// Test FromFeetPerSecondSquared function
func TestAccelerationFactory_FromFeetPerSecondSquared(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFeetPerSecondSquared(100)
    if err != nil {
        t.Errorf("FromFeetPerSecondSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationFootPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFeetPerSecondSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFeetPerSecondSquared(math.NaN())
    if err == nil {
        t.Error("FromFeetPerSecondSquared() with NaN value should return error")
    }

    _, err = factory.FromFeetPerSecondSquared(math.Inf(1))
    if err == nil {
        t.Error("FromFeetPerSecondSquared() with +Inf value should return error")
    }

    _, err = factory.FromFeetPerSecondSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromFeetPerSecondSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFeetPerSecondSquared(0)
    if err != nil {
        t.Errorf("FromFeetPerSecondSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationFootPerSecondSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFeetPerSecondSquared() with zero value = %v, want 0", converted)
    }
}
// Test FromKnotsPerSecond function
func TestAccelerationFactory_FromKnotsPerSecond(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKnotsPerSecond(100)
    if err != nil {
        t.Errorf("FromKnotsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationKnotPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKnotsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKnotsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKnotsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKnotsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKnotsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKnotsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKnotsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKnotsPerSecond(0)
    if err != nil {
        t.Errorf("FromKnotsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationKnotPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKnotsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKnotsPerMinute function
func TestAccelerationFactory_FromKnotsPerMinute(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKnotsPerMinute(100)
    if err != nil {
        t.Errorf("FromKnotsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationKnotPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKnotsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKnotsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromKnotsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromKnotsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromKnotsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromKnotsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromKnotsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKnotsPerMinute(0)
    if err != nil {
        t.Errorf("FromKnotsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationKnotPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKnotsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromKnotsPerHour function
func TestAccelerationFactory_FromKnotsPerHour(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKnotsPerHour(100)
    if err != nil {
        t.Errorf("FromKnotsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationKnotPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKnotsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKnotsPerHour(math.NaN())
    if err == nil {
        t.Error("FromKnotsPerHour() with NaN value should return error")
    }

    _, err = factory.FromKnotsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromKnotsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromKnotsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromKnotsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKnotsPerHour(0)
    if err != nil {
        t.Errorf("FromKnotsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationKnotPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKnotsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromStandardGravity function
func TestAccelerationFactory_FromStandardGravity(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromStandardGravity(100)
    if err != nil {
        t.Errorf("FromStandardGravity() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationStandardGravity)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromStandardGravity() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromStandardGravity(math.NaN())
    if err == nil {
        t.Error("FromStandardGravity() with NaN value should return error")
    }

    _, err = factory.FromStandardGravity(math.Inf(1))
    if err == nil {
        t.Error("FromStandardGravity() with +Inf value should return error")
    }

    _, err = factory.FromStandardGravity(math.Inf(-1))
    if err == nil {
        t.Error("FromStandardGravity() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromStandardGravity(0)
    if err != nil {
        t.Errorf("FromStandardGravity() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationStandardGravity)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromStandardGravity() with zero value = %v, want 0", converted)
    }
}
// Test FromNanometersPerSecondSquared function
func TestAccelerationFactory_FromNanometersPerSecondSquared(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanometersPerSecondSquared(100)
    if err != nil {
        t.Errorf("FromNanometersPerSecondSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationNanometerPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanometersPerSecondSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanometersPerSecondSquared(math.NaN())
    if err == nil {
        t.Error("FromNanometersPerSecondSquared() with NaN value should return error")
    }

    _, err = factory.FromNanometersPerSecondSquared(math.Inf(1))
    if err == nil {
        t.Error("FromNanometersPerSecondSquared() with +Inf value should return error")
    }

    _, err = factory.FromNanometersPerSecondSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromNanometersPerSecondSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanometersPerSecondSquared(0)
    if err != nil {
        t.Errorf("FromNanometersPerSecondSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationNanometerPerSecondSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanometersPerSecondSquared() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrometersPerSecondSquared function
func TestAccelerationFactory_FromMicrometersPerSecondSquared(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrometersPerSecondSquared(100)
    if err != nil {
        t.Errorf("FromMicrometersPerSecondSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationMicrometerPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrometersPerSecondSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrometersPerSecondSquared(math.NaN())
    if err == nil {
        t.Error("FromMicrometersPerSecondSquared() with NaN value should return error")
    }

    _, err = factory.FromMicrometersPerSecondSquared(math.Inf(1))
    if err == nil {
        t.Error("FromMicrometersPerSecondSquared() with +Inf value should return error")
    }

    _, err = factory.FromMicrometersPerSecondSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrometersPerSecondSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrometersPerSecondSquared(0)
    if err != nil {
        t.Errorf("FromMicrometersPerSecondSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationMicrometerPerSecondSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrometersPerSecondSquared() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimetersPerSecondSquared function
func TestAccelerationFactory_FromMillimetersPerSecondSquared(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimetersPerSecondSquared(100)
    if err != nil {
        t.Errorf("FromMillimetersPerSecondSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationMillimeterPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimetersPerSecondSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimetersPerSecondSquared(math.NaN())
    if err == nil {
        t.Error("FromMillimetersPerSecondSquared() with NaN value should return error")
    }

    _, err = factory.FromMillimetersPerSecondSquared(math.Inf(1))
    if err == nil {
        t.Error("FromMillimetersPerSecondSquared() with +Inf value should return error")
    }

    _, err = factory.FromMillimetersPerSecondSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimetersPerSecondSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimetersPerSecondSquared(0)
    if err != nil {
        t.Errorf("FromMillimetersPerSecondSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationMillimeterPerSecondSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimetersPerSecondSquared() with zero value = %v, want 0", converted)
    }
}
// Test FromCentimetersPerSecondSquared function
func TestAccelerationFactory_FromCentimetersPerSecondSquared(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentimetersPerSecondSquared(100)
    if err != nil {
        t.Errorf("FromCentimetersPerSecondSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationCentimeterPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentimetersPerSecondSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentimetersPerSecondSquared(math.NaN())
    if err == nil {
        t.Error("FromCentimetersPerSecondSquared() with NaN value should return error")
    }

    _, err = factory.FromCentimetersPerSecondSquared(math.Inf(1))
    if err == nil {
        t.Error("FromCentimetersPerSecondSquared() with +Inf value should return error")
    }

    _, err = factory.FromCentimetersPerSecondSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromCentimetersPerSecondSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentimetersPerSecondSquared(0)
    if err != nil {
        t.Errorf("FromCentimetersPerSecondSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationCentimeterPerSecondSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentimetersPerSecondSquared() with zero value = %v, want 0", converted)
    }
}
// Test FromDecimetersPerSecondSquared function
func TestAccelerationFactory_FromDecimetersPerSecondSquared(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecimetersPerSecondSquared(100)
    if err != nil {
        t.Errorf("FromDecimetersPerSecondSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationDecimeterPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecimetersPerSecondSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecimetersPerSecondSquared(math.NaN())
    if err == nil {
        t.Error("FromDecimetersPerSecondSquared() with NaN value should return error")
    }

    _, err = factory.FromDecimetersPerSecondSquared(math.Inf(1))
    if err == nil {
        t.Error("FromDecimetersPerSecondSquared() with +Inf value should return error")
    }

    _, err = factory.FromDecimetersPerSecondSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromDecimetersPerSecondSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecimetersPerSecondSquared(0)
    if err != nil {
        t.Errorf("FromDecimetersPerSecondSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationDecimeterPerSecondSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecimetersPerSecondSquared() with zero value = %v, want 0", converted)
    }
}
// Test FromKilometersPerSecondSquared function
func TestAccelerationFactory_FromKilometersPerSecondSquared(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilometersPerSecondSquared(100)
    if err != nil {
        t.Errorf("FromKilometersPerSecondSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationKilometerPerSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilometersPerSecondSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilometersPerSecondSquared(math.NaN())
    if err == nil {
        t.Error("FromKilometersPerSecondSquared() with NaN value should return error")
    }

    _, err = factory.FromKilometersPerSecondSquared(math.Inf(1))
    if err == nil {
        t.Error("FromKilometersPerSecondSquared() with +Inf value should return error")
    }

    _, err = factory.FromKilometersPerSecondSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromKilometersPerSecondSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilometersPerSecondSquared(0)
    if err != nil {
        t.Errorf("FromKilometersPerSecondSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationKilometerPerSecondSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilometersPerSecondSquared() with zero value = %v, want 0", converted)
    }
}
// Test FromMillistandardGravity function
func TestAccelerationFactory_FromMillistandardGravity(t *testing.T) {
    factory := units.AccelerationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillistandardGravity(100)
    if err != nil {
        t.Errorf("FromMillistandardGravity() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AccelerationMillistandardGravity)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillistandardGravity() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillistandardGravity(math.NaN())
    if err == nil {
        t.Error("FromMillistandardGravity() with NaN value should return error")
    }

    _, err = factory.FromMillistandardGravity(math.Inf(1))
    if err == nil {
        t.Error("FromMillistandardGravity() with +Inf value should return error")
    }

    _, err = factory.FromMillistandardGravity(math.Inf(-1))
    if err == nil {
        t.Error("FromMillistandardGravity() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillistandardGravity(0)
    if err != nil {
        t.Errorf("FromMillistandardGravity() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AccelerationMillistandardGravity)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillistandardGravity() with zero value = %v, want 0", converted)
    }
}

func TestAccelerationToString(t *testing.T) {
	factory := units.AccelerationFactory{}
	a, err := factory.CreateAcceleration(45, units.AccelerationMeterPerSecondSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AccelerationMeterPerSecondSquared, 2)
	expected := "45.00 " + units.GetAccelerationAbbreviation(units.AccelerationMeterPerSecondSquared)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AccelerationMeterPerSecondSquared, -1)
	expected = "45 " + units.GetAccelerationAbbreviation(units.AccelerationMeterPerSecondSquared)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAcceleration_EqualityAndComparison(t *testing.T) {
	factory := units.AccelerationFactory{}
	a1, _ := factory.CreateAcceleration(60, units.AccelerationMeterPerSecondSquared)
	a2, _ := factory.CreateAcceleration(60, units.AccelerationMeterPerSecondSquared)
	a3, _ := factory.CreateAcceleration(120, units.AccelerationMeterPerSecondSquared)

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

func TestAcceleration_Arithmetic(t *testing.T) {
	factory := units.AccelerationFactory{}
	a1, _ := factory.CreateAcceleration(30, units.AccelerationMeterPerSecondSquared)
	a2, _ := factory.CreateAcceleration(45, units.AccelerationMeterPerSecondSquared)

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


func TestGetAccelerationAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.AccelerationUnits
        want string
    }{
        {
            name: "MeterPerSecondSquared abbreviation",
            unit: units.AccelerationMeterPerSecondSquared,
            want: "m/s²",
        },
        {
            name: "InchPerSecondSquared abbreviation",
            unit: units.AccelerationInchPerSecondSquared,
            want: "in/s²",
        },
        {
            name: "FootPerSecondSquared abbreviation",
            unit: units.AccelerationFootPerSecondSquared,
            want: "ft/s²",
        },
        {
            name: "KnotPerSecond abbreviation",
            unit: units.AccelerationKnotPerSecond,
            want: "kn/s",
        },
        {
            name: "KnotPerMinute abbreviation",
            unit: units.AccelerationKnotPerMinute,
            want: "kn/min",
        },
        {
            name: "KnotPerHour abbreviation",
            unit: units.AccelerationKnotPerHour,
            want: "kn/h",
        },
        {
            name: "StandardGravity abbreviation",
            unit: units.AccelerationStandardGravity,
            want: "g",
        },
        {
            name: "NanometerPerSecondSquared abbreviation",
            unit: units.AccelerationNanometerPerSecondSquared,
            want: "nm/s²",
        },
        {
            name: "MicrometerPerSecondSquared abbreviation",
            unit: units.AccelerationMicrometerPerSecondSquared,
            want: "μm/s²",
        },
        {
            name: "MillimeterPerSecondSquared abbreviation",
            unit: units.AccelerationMillimeterPerSecondSquared,
            want: "mm/s²",
        },
        {
            name: "CentimeterPerSecondSquared abbreviation",
            unit: units.AccelerationCentimeterPerSecondSquared,
            want: "cm/s²",
        },
        {
            name: "DecimeterPerSecondSquared abbreviation",
            unit: units.AccelerationDecimeterPerSecondSquared,
            want: "dm/s²",
        },
        {
            name: "KilometerPerSecondSquared abbreviation",
            unit: units.AccelerationKilometerPerSecondSquared,
            want: "km/s²",
        },
        {
            name: "MillistandardGravity abbreviation",
            unit: units.AccelerationMillistandardGravity,
            want: "mg",
        },
        {
            name: "invalid unit",
            unit: units.AccelerationUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetAccelerationAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetAccelerationAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestAcceleration_String(t *testing.T) {
    factory := units.AccelerationFactory{}
    
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
            unit, err := factory.CreateAcceleration(tt.value, units.AccelerationMeterPerSecondSquared)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Acceleration.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestAcceleration_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.AccelerationFactory{}

	_, err := uf.CreateAcceleration(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}