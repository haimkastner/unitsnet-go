// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestStandardVolumeFlowDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "StandardCubicMeterPerSecond"}`
	
	factory := units.StandardVolumeFlowDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.StandardVolumeFlowStandardCubicMeterPerSecond {
		t.Errorf("expected unit %v, got %v", units.StandardVolumeFlowStandardCubicMeterPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "StandardCubicMeterPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestStandardVolumeFlowDto_ToJSON(t *testing.T) {
	dto := units.StandardVolumeFlowDto{
		Value: 45,
		Unit:  units.StandardVolumeFlowStandardCubicMeterPerSecond,
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
	if result["unit"].(string) != string(units.StandardVolumeFlowStandardCubicMeterPerSecond) {
		t.Errorf("expected unit %s, got %v", units.StandardVolumeFlowStandardCubicMeterPerSecond, result["unit"])
	}
}

func TestNewStandardVolumeFlow_InvalidValue(t *testing.T) {
	factory := units.StandardVolumeFlowFactory{}
	// NaN value should return an error.
	_, err := factory.CreateStandardVolumeFlow(math.NaN(), units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateStandardVolumeFlow(math.Inf(1), units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestStandardVolumeFlowConversions(t *testing.T) {
	factory := units.StandardVolumeFlowFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateStandardVolumeFlow(180, units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to StandardCubicMetersPerSecond.
		// No expected conversion value provided for StandardCubicMetersPerSecond, verifying result is not NaN.
		result := a.StandardCubicMetersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicMetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicMetersPerMinute.
		// No expected conversion value provided for StandardCubicMetersPerMinute, verifying result is not NaN.
		result := a.StandardCubicMetersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicMetersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicMetersPerHour.
		// No expected conversion value provided for StandardCubicMetersPerHour, verifying result is not NaN.
		result := a.StandardCubicMetersPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicMetersPerHour returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicMetersPerDay.
		// No expected conversion value provided for StandardCubicMetersPerDay, verifying result is not NaN.
		result := a.StandardCubicMetersPerDay()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicMetersPerDay returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicCentimetersPerMinute.
		// No expected conversion value provided for StandardCubicCentimetersPerMinute, verifying result is not NaN.
		result := a.StandardCubicCentimetersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicCentimetersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to StandardLitersPerMinute.
		// No expected conversion value provided for StandardLitersPerMinute, verifying result is not NaN.
		result := a.StandardLitersPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardLitersPerMinute returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicFeetPerSecond.
		// No expected conversion value provided for StandardCubicFeetPerSecond, verifying result is not NaN.
		result := a.StandardCubicFeetPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicFeetPerMinute.
		// No expected conversion value provided for StandardCubicFeetPerMinute, verifying result is not NaN.
		result := a.StandardCubicFeetPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicFeetPerMinute returned NaN")
		}
	}
	{
		// Test conversion to StandardCubicFeetPerHour.
		// No expected conversion value provided for StandardCubicFeetPerHour, verifying result is not NaN.
		result := a.StandardCubicFeetPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to StandardCubicFeetPerHour returned NaN")
		}
	}
}

func TestStandardVolumeFlow_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.StandardVolumeFlowFactory{}
	a, err := factory.CreateStandardVolumeFlow(90, units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.StandardVolumeFlowStandardCubicMeterPerSecond {
		t.Errorf("expected default unit StandardCubicMeterPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.StandardVolumeFlowStandardCubicMeterPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.StandardVolumeFlowDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.StandardVolumeFlowStandardCubicMeterPerSecond {
		t.Errorf("expected unit StandardCubicMeterPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestStandardVolumeFlowFactory_FromDto(t *testing.T) {
    factory := units.StandardVolumeFlowFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.StandardVolumeFlowDto{
        Value: 100,
        Unit:  units.StandardVolumeFlowStandardCubicMeterPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.StandardVolumeFlowDto{
        Value: math.NaN(),
        Unit:  units.StandardVolumeFlowStandardCubicMeterPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test StandardCubicMeterPerSecond conversion
    standard_cubic_meters_per_secondDto := units.StandardVolumeFlowDto{
        Value: 100,
        Unit:  units.StandardVolumeFlowStandardCubicMeterPerSecond,
    }
    
    var standard_cubic_meters_per_secondResult *units.StandardVolumeFlow
    standard_cubic_meters_per_secondResult, err = factory.FromDto(standard_cubic_meters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with StandardCubicMeterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_meters_per_secondResult.Convert(units.StandardVolumeFlowStandardCubicMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicMeterPerSecond = %v, want %v", converted, 100)
    }
    // Test StandardCubicMeterPerMinute conversion
    standard_cubic_meters_per_minuteDto := units.StandardVolumeFlowDto{
        Value: 100,
        Unit:  units.StandardVolumeFlowStandardCubicMeterPerMinute,
    }
    
    var standard_cubic_meters_per_minuteResult *units.StandardVolumeFlow
    standard_cubic_meters_per_minuteResult, err = factory.FromDto(standard_cubic_meters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with StandardCubicMeterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_meters_per_minuteResult.Convert(units.StandardVolumeFlowStandardCubicMeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicMeterPerMinute = %v, want %v", converted, 100)
    }
    // Test StandardCubicMeterPerHour conversion
    standard_cubic_meters_per_hourDto := units.StandardVolumeFlowDto{
        Value: 100,
        Unit:  units.StandardVolumeFlowStandardCubicMeterPerHour,
    }
    
    var standard_cubic_meters_per_hourResult *units.StandardVolumeFlow
    standard_cubic_meters_per_hourResult, err = factory.FromDto(standard_cubic_meters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with StandardCubicMeterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_meters_per_hourResult.Convert(units.StandardVolumeFlowStandardCubicMeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicMeterPerHour = %v, want %v", converted, 100)
    }
    // Test StandardCubicMeterPerDay conversion
    standard_cubic_meters_per_dayDto := units.StandardVolumeFlowDto{
        Value: 100,
        Unit:  units.StandardVolumeFlowStandardCubicMeterPerDay,
    }
    
    var standard_cubic_meters_per_dayResult *units.StandardVolumeFlow
    standard_cubic_meters_per_dayResult, err = factory.FromDto(standard_cubic_meters_per_dayDto)
    if err != nil {
        t.Errorf("FromDto() with StandardCubicMeterPerDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_meters_per_dayResult.Convert(units.StandardVolumeFlowStandardCubicMeterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicMeterPerDay = %v, want %v", converted, 100)
    }
    // Test StandardCubicCentimeterPerMinute conversion
    standard_cubic_centimeters_per_minuteDto := units.StandardVolumeFlowDto{
        Value: 100,
        Unit:  units.StandardVolumeFlowStandardCubicCentimeterPerMinute,
    }
    
    var standard_cubic_centimeters_per_minuteResult *units.StandardVolumeFlow
    standard_cubic_centimeters_per_minuteResult, err = factory.FromDto(standard_cubic_centimeters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with StandardCubicCentimeterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_centimeters_per_minuteResult.Convert(units.StandardVolumeFlowStandardCubicCentimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicCentimeterPerMinute = %v, want %v", converted, 100)
    }
    // Test StandardLiterPerMinute conversion
    standard_liters_per_minuteDto := units.StandardVolumeFlowDto{
        Value: 100,
        Unit:  units.StandardVolumeFlowStandardLiterPerMinute,
    }
    
    var standard_liters_per_minuteResult *units.StandardVolumeFlow
    standard_liters_per_minuteResult, err = factory.FromDto(standard_liters_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with StandardLiterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_liters_per_minuteResult.Convert(units.StandardVolumeFlowStandardLiterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardLiterPerMinute = %v, want %v", converted, 100)
    }
    // Test StandardCubicFootPerSecond conversion
    standard_cubic_feet_per_secondDto := units.StandardVolumeFlowDto{
        Value: 100,
        Unit:  units.StandardVolumeFlowStandardCubicFootPerSecond,
    }
    
    var standard_cubic_feet_per_secondResult *units.StandardVolumeFlow
    standard_cubic_feet_per_secondResult, err = factory.FromDto(standard_cubic_feet_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with StandardCubicFootPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_feet_per_secondResult.Convert(units.StandardVolumeFlowStandardCubicFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicFootPerSecond = %v, want %v", converted, 100)
    }
    // Test StandardCubicFootPerMinute conversion
    standard_cubic_feet_per_minuteDto := units.StandardVolumeFlowDto{
        Value: 100,
        Unit:  units.StandardVolumeFlowStandardCubicFootPerMinute,
    }
    
    var standard_cubic_feet_per_minuteResult *units.StandardVolumeFlow
    standard_cubic_feet_per_minuteResult, err = factory.FromDto(standard_cubic_feet_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with StandardCubicFootPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_feet_per_minuteResult.Convert(units.StandardVolumeFlowStandardCubicFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicFootPerMinute = %v, want %v", converted, 100)
    }
    // Test StandardCubicFootPerHour conversion
    standard_cubic_feet_per_hourDto := units.StandardVolumeFlowDto{
        Value: 100,
        Unit:  units.StandardVolumeFlowStandardCubicFootPerHour,
    }
    
    var standard_cubic_feet_per_hourResult *units.StandardVolumeFlow
    standard_cubic_feet_per_hourResult, err = factory.FromDto(standard_cubic_feet_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with StandardCubicFootPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_feet_per_hourResult.Convert(units.StandardVolumeFlowStandardCubicFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicFootPerHour = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.StandardVolumeFlowDto{
        Value: 0,
        Unit:  units.StandardVolumeFlowStandardCubicMeterPerSecond,
    }
    
    var zeroResult *units.StandardVolumeFlow
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestStandardVolumeFlowFactory_FromDtoJSON(t *testing.T) {
    factory := units.StandardVolumeFlowFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "StandardCubicMeterPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "StandardCubicMeterPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.StandardVolumeFlowDto{
        Value: nanValue,
        Unit:  units.StandardVolumeFlowStandardCubicMeterPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with StandardCubicMeterPerSecond unit
    standard_cubic_meters_per_secondJSON := []byte(`{"value": 100, "unit": "StandardCubicMeterPerSecond"}`)
    standard_cubic_meters_per_secondResult, err := factory.FromDtoJSON(standard_cubic_meters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with StandardCubicMeterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_meters_per_secondResult.Convert(units.StandardVolumeFlowStandardCubicMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicMeterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with StandardCubicMeterPerMinute unit
    standard_cubic_meters_per_minuteJSON := []byte(`{"value": 100, "unit": "StandardCubicMeterPerMinute"}`)
    standard_cubic_meters_per_minuteResult, err := factory.FromDtoJSON(standard_cubic_meters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with StandardCubicMeterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_meters_per_minuteResult.Convert(units.StandardVolumeFlowStandardCubicMeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicMeterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with StandardCubicMeterPerHour unit
    standard_cubic_meters_per_hourJSON := []byte(`{"value": 100, "unit": "StandardCubicMeterPerHour"}`)
    standard_cubic_meters_per_hourResult, err := factory.FromDtoJSON(standard_cubic_meters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with StandardCubicMeterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_meters_per_hourResult.Convert(units.StandardVolumeFlowStandardCubicMeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicMeterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with StandardCubicMeterPerDay unit
    standard_cubic_meters_per_dayJSON := []byte(`{"value": 100, "unit": "StandardCubicMeterPerDay"}`)
    standard_cubic_meters_per_dayResult, err := factory.FromDtoJSON(standard_cubic_meters_per_dayJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with StandardCubicMeterPerDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_meters_per_dayResult.Convert(units.StandardVolumeFlowStandardCubicMeterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicMeterPerDay = %v, want %v", converted, 100)
    }
    // Test JSON with StandardCubicCentimeterPerMinute unit
    standard_cubic_centimeters_per_minuteJSON := []byte(`{"value": 100, "unit": "StandardCubicCentimeterPerMinute"}`)
    standard_cubic_centimeters_per_minuteResult, err := factory.FromDtoJSON(standard_cubic_centimeters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with StandardCubicCentimeterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_centimeters_per_minuteResult.Convert(units.StandardVolumeFlowStandardCubicCentimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicCentimeterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with StandardLiterPerMinute unit
    standard_liters_per_minuteJSON := []byte(`{"value": 100, "unit": "StandardLiterPerMinute"}`)
    standard_liters_per_minuteResult, err := factory.FromDtoJSON(standard_liters_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with StandardLiterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_liters_per_minuteResult.Convert(units.StandardVolumeFlowStandardLiterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardLiterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with StandardCubicFootPerSecond unit
    standard_cubic_feet_per_secondJSON := []byte(`{"value": 100, "unit": "StandardCubicFootPerSecond"}`)
    standard_cubic_feet_per_secondResult, err := factory.FromDtoJSON(standard_cubic_feet_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with StandardCubicFootPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_feet_per_secondResult.Convert(units.StandardVolumeFlowStandardCubicFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicFootPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with StandardCubicFootPerMinute unit
    standard_cubic_feet_per_minuteJSON := []byte(`{"value": 100, "unit": "StandardCubicFootPerMinute"}`)
    standard_cubic_feet_per_minuteResult, err := factory.FromDtoJSON(standard_cubic_feet_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with StandardCubicFootPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_feet_per_minuteResult.Convert(units.StandardVolumeFlowStandardCubicFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicFootPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with StandardCubicFootPerHour unit
    standard_cubic_feet_per_hourJSON := []byte(`{"value": 100, "unit": "StandardCubicFootPerHour"}`)
    standard_cubic_feet_per_hourResult, err := factory.FromDtoJSON(standard_cubic_feet_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with StandardCubicFootPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_cubic_feet_per_hourResult.Convert(units.StandardVolumeFlowStandardCubicFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardCubicFootPerHour = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "StandardCubicMeterPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromStandardCubicMetersPerSecond function
func TestStandardVolumeFlowFactory_FromStandardCubicMetersPerSecond(t *testing.T) {
    factory := units.StandardVolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromStandardCubicMetersPerSecond(100)
    if err != nil {
        t.Errorf("FromStandardCubicMetersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.StandardVolumeFlowStandardCubicMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromStandardCubicMetersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromStandardCubicMetersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromStandardCubicMetersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromStandardCubicMetersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromStandardCubicMetersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromStandardCubicMetersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromStandardCubicMetersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromStandardCubicMetersPerSecond(0)
    if err != nil {
        t.Errorf("FromStandardCubicMetersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.StandardVolumeFlowStandardCubicMeterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromStandardCubicMetersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromStandardCubicMetersPerMinute function
func TestStandardVolumeFlowFactory_FromStandardCubicMetersPerMinute(t *testing.T) {
    factory := units.StandardVolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromStandardCubicMetersPerMinute(100)
    if err != nil {
        t.Errorf("FromStandardCubicMetersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.StandardVolumeFlowStandardCubicMeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromStandardCubicMetersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromStandardCubicMetersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromStandardCubicMetersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromStandardCubicMetersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromStandardCubicMetersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromStandardCubicMetersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromStandardCubicMetersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromStandardCubicMetersPerMinute(0)
    if err != nil {
        t.Errorf("FromStandardCubicMetersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.StandardVolumeFlowStandardCubicMeterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromStandardCubicMetersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromStandardCubicMetersPerHour function
func TestStandardVolumeFlowFactory_FromStandardCubicMetersPerHour(t *testing.T) {
    factory := units.StandardVolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromStandardCubicMetersPerHour(100)
    if err != nil {
        t.Errorf("FromStandardCubicMetersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.StandardVolumeFlowStandardCubicMeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromStandardCubicMetersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromStandardCubicMetersPerHour(math.NaN())
    if err == nil {
        t.Error("FromStandardCubicMetersPerHour() with NaN value should return error")
    }

    _, err = factory.FromStandardCubicMetersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromStandardCubicMetersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromStandardCubicMetersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromStandardCubicMetersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromStandardCubicMetersPerHour(0)
    if err != nil {
        t.Errorf("FromStandardCubicMetersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.StandardVolumeFlowStandardCubicMeterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromStandardCubicMetersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromStandardCubicMetersPerDay function
func TestStandardVolumeFlowFactory_FromStandardCubicMetersPerDay(t *testing.T) {
    factory := units.StandardVolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromStandardCubicMetersPerDay(100)
    if err != nil {
        t.Errorf("FromStandardCubicMetersPerDay() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.StandardVolumeFlowStandardCubicMeterPerDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromStandardCubicMetersPerDay() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromStandardCubicMetersPerDay(math.NaN())
    if err == nil {
        t.Error("FromStandardCubicMetersPerDay() with NaN value should return error")
    }

    _, err = factory.FromStandardCubicMetersPerDay(math.Inf(1))
    if err == nil {
        t.Error("FromStandardCubicMetersPerDay() with +Inf value should return error")
    }

    _, err = factory.FromStandardCubicMetersPerDay(math.Inf(-1))
    if err == nil {
        t.Error("FromStandardCubicMetersPerDay() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromStandardCubicMetersPerDay(0)
    if err != nil {
        t.Errorf("FromStandardCubicMetersPerDay() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.StandardVolumeFlowStandardCubicMeterPerDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromStandardCubicMetersPerDay() with zero value = %v, want 0", converted)
    }
}
// Test FromStandardCubicCentimetersPerMinute function
func TestStandardVolumeFlowFactory_FromStandardCubicCentimetersPerMinute(t *testing.T) {
    factory := units.StandardVolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromStandardCubicCentimetersPerMinute(100)
    if err != nil {
        t.Errorf("FromStandardCubicCentimetersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.StandardVolumeFlowStandardCubicCentimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromStandardCubicCentimetersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromStandardCubicCentimetersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromStandardCubicCentimetersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromStandardCubicCentimetersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromStandardCubicCentimetersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromStandardCubicCentimetersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromStandardCubicCentimetersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromStandardCubicCentimetersPerMinute(0)
    if err != nil {
        t.Errorf("FromStandardCubicCentimetersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.StandardVolumeFlowStandardCubicCentimeterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromStandardCubicCentimetersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromStandardLitersPerMinute function
func TestStandardVolumeFlowFactory_FromStandardLitersPerMinute(t *testing.T) {
    factory := units.StandardVolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromStandardLitersPerMinute(100)
    if err != nil {
        t.Errorf("FromStandardLitersPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.StandardVolumeFlowStandardLiterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromStandardLitersPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromStandardLitersPerMinute(math.NaN())
    if err == nil {
        t.Error("FromStandardLitersPerMinute() with NaN value should return error")
    }

    _, err = factory.FromStandardLitersPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromStandardLitersPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromStandardLitersPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromStandardLitersPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromStandardLitersPerMinute(0)
    if err != nil {
        t.Errorf("FromStandardLitersPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.StandardVolumeFlowStandardLiterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromStandardLitersPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromStandardCubicFeetPerSecond function
func TestStandardVolumeFlowFactory_FromStandardCubicFeetPerSecond(t *testing.T) {
    factory := units.StandardVolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromStandardCubicFeetPerSecond(100)
    if err != nil {
        t.Errorf("FromStandardCubicFeetPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.StandardVolumeFlowStandardCubicFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromStandardCubicFeetPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromStandardCubicFeetPerSecond(math.NaN())
    if err == nil {
        t.Error("FromStandardCubicFeetPerSecond() with NaN value should return error")
    }

    _, err = factory.FromStandardCubicFeetPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromStandardCubicFeetPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromStandardCubicFeetPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromStandardCubicFeetPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromStandardCubicFeetPerSecond(0)
    if err != nil {
        t.Errorf("FromStandardCubicFeetPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.StandardVolumeFlowStandardCubicFootPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromStandardCubicFeetPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromStandardCubicFeetPerMinute function
func TestStandardVolumeFlowFactory_FromStandardCubicFeetPerMinute(t *testing.T) {
    factory := units.StandardVolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromStandardCubicFeetPerMinute(100)
    if err != nil {
        t.Errorf("FromStandardCubicFeetPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.StandardVolumeFlowStandardCubicFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromStandardCubicFeetPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromStandardCubicFeetPerMinute(math.NaN())
    if err == nil {
        t.Error("FromStandardCubicFeetPerMinute() with NaN value should return error")
    }

    _, err = factory.FromStandardCubicFeetPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromStandardCubicFeetPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromStandardCubicFeetPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromStandardCubicFeetPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromStandardCubicFeetPerMinute(0)
    if err != nil {
        t.Errorf("FromStandardCubicFeetPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.StandardVolumeFlowStandardCubicFootPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromStandardCubicFeetPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromStandardCubicFeetPerHour function
func TestStandardVolumeFlowFactory_FromStandardCubicFeetPerHour(t *testing.T) {
    factory := units.StandardVolumeFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromStandardCubicFeetPerHour(100)
    if err != nil {
        t.Errorf("FromStandardCubicFeetPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.StandardVolumeFlowStandardCubicFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromStandardCubicFeetPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromStandardCubicFeetPerHour(math.NaN())
    if err == nil {
        t.Error("FromStandardCubicFeetPerHour() with NaN value should return error")
    }

    _, err = factory.FromStandardCubicFeetPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromStandardCubicFeetPerHour() with +Inf value should return error")
    }

    _, err = factory.FromStandardCubicFeetPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromStandardCubicFeetPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromStandardCubicFeetPerHour(0)
    if err != nil {
        t.Errorf("FromStandardCubicFeetPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.StandardVolumeFlowStandardCubicFootPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromStandardCubicFeetPerHour() with zero value = %v, want 0", converted)
    }
}

func TestStandardVolumeFlowToString(t *testing.T) {
	factory := units.StandardVolumeFlowFactory{}
	a, err := factory.CreateStandardVolumeFlow(45, units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.StandardVolumeFlowStandardCubicMeterPerSecond, 2)
	expected := "45.00 " + units.GetStandardVolumeFlowAbbreviation(units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.StandardVolumeFlowStandardCubicMeterPerSecond, -1)
	expected = "45 " + units.GetStandardVolumeFlowAbbreviation(units.StandardVolumeFlowStandardCubicMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestStandardVolumeFlow_EqualityAndComparison(t *testing.T) {
	factory := units.StandardVolumeFlowFactory{}
	a1, _ := factory.CreateStandardVolumeFlow(60, units.StandardVolumeFlowStandardCubicMeterPerSecond)
	a2, _ := factory.CreateStandardVolumeFlow(60, units.StandardVolumeFlowStandardCubicMeterPerSecond)
	a3, _ := factory.CreateStandardVolumeFlow(120, units.StandardVolumeFlowStandardCubicMeterPerSecond)

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

func TestStandardVolumeFlow_Arithmetic(t *testing.T) {
	factory := units.StandardVolumeFlowFactory{}
	a1, _ := factory.CreateStandardVolumeFlow(30, units.StandardVolumeFlowStandardCubicMeterPerSecond)
	a2, _ := factory.CreateStandardVolumeFlow(45, units.StandardVolumeFlowStandardCubicMeterPerSecond)

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