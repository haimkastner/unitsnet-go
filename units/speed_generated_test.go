// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestSpeedDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MeterPerSecond"}`
	
	factory := units.SpeedDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.SpeedMeterPerSecond {
		t.Errorf("expected unit %v, got %v", units.SpeedMeterPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MeterPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestSpeedDto_ToJSON(t *testing.T) {
	dto := units.SpeedDto{
		Value: 45,
		Unit:  units.SpeedMeterPerSecond,
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
	if result["unit"].(string) != string(units.SpeedMeterPerSecond) {
		t.Errorf("expected unit %s, got %v", units.SpeedMeterPerSecond, result["unit"])
	}
}

func TestNewSpeed_InvalidValue(t *testing.T) {
	factory := units.SpeedFactory{}
	// NaN value should return an error.
	_, err := factory.CreateSpeed(math.NaN(), units.SpeedMeterPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateSpeed(math.Inf(1), units.SpeedMeterPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestSpeedConversions(t *testing.T) {
	factory := units.SpeedFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateSpeed(180, units.SpeedMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MetersPerSecond.
		// No expected conversion value provided for MetersPerSecond, verifying result is not NaN.
		result := a.MetersPerSecond()
		cacheResult := a.MetersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MetersPerMinutes.
		// No expected conversion value provided for MetersPerMinutes, verifying result is not NaN.
		result := a.MetersPerMinutes()
		cacheResult := a.MetersPerMinutes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MetersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MetersPerHour.
		// No expected conversion value provided for MetersPerHour, verifying result is not NaN.
		result := a.MetersPerHour()
		cacheResult := a.MetersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MetersPerHour returned NaN")
		}
	}
	{
		// Test conversion to FeetPerSecond.
		// No expected conversion value provided for FeetPerSecond, verifying result is not NaN.
		result := a.FeetPerSecond()
		cacheResult := a.FeetPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to FeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to FeetPerMinute.
		// No expected conversion value provided for FeetPerMinute, verifying result is not NaN.
		result := a.FeetPerMinute()
		cacheResult := a.FeetPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to FeetPerMinute returned NaN")
		}
	}
	{
		// Test conversion to FeetPerHour.
		// No expected conversion value provided for FeetPerHour, verifying result is not NaN.
		result := a.FeetPerHour()
		cacheResult := a.FeetPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to FeetPerHour returned NaN")
		}
	}
	{
		// Test conversion to UsSurveyFeetPerSecond.
		// No expected conversion value provided for UsSurveyFeetPerSecond, verifying result is not NaN.
		result := a.UsSurveyFeetPerSecond()
		cacheResult := a.UsSurveyFeetPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsSurveyFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to UsSurveyFeetPerMinute.
		// No expected conversion value provided for UsSurveyFeetPerMinute, verifying result is not NaN.
		result := a.UsSurveyFeetPerMinute()
		cacheResult := a.UsSurveyFeetPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsSurveyFeetPerMinute returned NaN")
		}
	}
	{
		// Test conversion to UsSurveyFeetPerHour.
		// No expected conversion value provided for UsSurveyFeetPerHour, verifying result is not NaN.
		result := a.UsSurveyFeetPerHour()
		cacheResult := a.UsSurveyFeetPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsSurveyFeetPerHour returned NaN")
		}
	}
	{
		// Test conversion to InchesPerSecond.
		// No expected conversion value provided for InchesPerSecond, verifying result is not NaN.
		result := a.InchesPerSecond()
		cacheResult := a.InchesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InchesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to InchesPerMinute.
		// No expected conversion value provided for InchesPerMinute, verifying result is not NaN.
		result := a.InchesPerMinute()
		cacheResult := a.InchesPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InchesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to InchesPerHour.
		// No expected conversion value provided for InchesPerHour, verifying result is not NaN.
		result := a.InchesPerHour()
		cacheResult := a.InchesPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InchesPerHour returned NaN")
		}
	}
	{
		// Test conversion to YardsPerSecond.
		// No expected conversion value provided for YardsPerSecond, verifying result is not NaN.
		result := a.YardsPerSecond()
		cacheResult := a.YardsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to YardsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to YardsPerMinute.
		// No expected conversion value provided for YardsPerMinute, verifying result is not NaN.
		result := a.YardsPerMinute()
		cacheResult := a.YardsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to YardsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to YardsPerHour.
		// No expected conversion value provided for YardsPerHour, verifying result is not NaN.
		result := a.YardsPerHour()
		cacheResult := a.YardsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to YardsPerHour returned NaN")
		}
	}
	{
		// Test conversion to Knots.
		// No expected conversion value provided for Knots, verifying result is not NaN.
		result := a.Knots()
		cacheResult := a.Knots()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Knots returned NaN")
		}
	}
	{
		// Test conversion to MilesPerHour.
		// No expected conversion value provided for MilesPerHour, verifying result is not NaN.
		result := a.MilesPerHour()
		cacheResult := a.MilesPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilesPerHour returned NaN")
		}
	}
	{
		// Test conversion to Mach.
		// No expected conversion value provided for Mach, verifying result is not NaN.
		result := a.Mach()
		cacheResult := a.Mach()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Mach returned NaN")
		}
	}
	{
		// Test conversion to NanometersPerSecond.
		// No expected conversion value provided for NanometersPerSecond, verifying result is not NaN.
		result := a.NanometersPerSecond()
		cacheResult := a.NanometersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanometersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicrometersPerSecond.
		// No expected conversion value provided for MicrometersPerSecond, verifying result is not NaN.
		result := a.MicrometersPerSecond()
		cacheResult := a.MicrometersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrometersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillimetersPerSecond.
		// No expected conversion value provided for MillimetersPerSecond, verifying result is not NaN.
		result := a.MillimetersPerSecond()
		cacheResult := a.MillimetersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CentimetersPerSecond.
		// No expected conversion value provided for CentimetersPerSecond, verifying result is not NaN.
		result := a.CentimetersPerSecond()
		cacheResult := a.CentimetersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentimetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecimetersPerSecond.
		// No expected conversion value provided for DecimetersPerSecond, verifying result is not NaN.
		result := a.DecimetersPerSecond()
		cacheResult := a.DecimetersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecimetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilometersPerSecond.
		// No expected conversion value provided for KilometersPerSecond, verifying result is not NaN.
		result := a.KilometersPerSecond()
		cacheResult := a.KilometersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilometersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to NanometersPerMinutes.
		// No expected conversion value provided for NanometersPerMinutes, verifying result is not NaN.
		result := a.NanometersPerMinutes()
		cacheResult := a.NanometersPerMinutes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanometersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MicrometersPerMinutes.
		// No expected conversion value provided for MicrometersPerMinutes, verifying result is not NaN.
		result := a.MicrometersPerMinutes()
		cacheResult := a.MicrometersPerMinutes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrometersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MillimetersPerMinutes.
		// No expected conversion value provided for MillimetersPerMinutes, verifying result is not NaN.
		result := a.MillimetersPerMinutes()
		cacheResult := a.MillimetersPerMinutes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimetersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to CentimetersPerMinutes.
		// No expected conversion value provided for CentimetersPerMinutes, verifying result is not NaN.
		result := a.CentimetersPerMinutes()
		cacheResult := a.CentimetersPerMinutes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentimetersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to DecimetersPerMinutes.
		// No expected conversion value provided for DecimetersPerMinutes, verifying result is not NaN.
		result := a.DecimetersPerMinutes()
		cacheResult := a.DecimetersPerMinutes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecimetersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to KilometersPerMinutes.
		// No expected conversion value provided for KilometersPerMinutes, verifying result is not NaN.
		result := a.KilometersPerMinutes()
		cacheResult := a.KilometersPerMinutes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilometersPerMinutes returned NaN")
		}
	}
	{
		// Test conversion to MillimetersPerHour.
		// No expected conversion value provided for MillimetersPerHour, verifying result is not NaN.
		result := a.MillimetersPerHour()
		cacheResult := a.MillimetersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimetersPerHour returned NaN")
		}
	}
	{
		// Test conversion to CentimetersPerHour.
		// No expected conversion value provided for CentimetersPerHour, verifying result is not NaN.
		result := a.CentimetersPerHour()
		cacheResult := a.CentimetersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentimetersPerHour returned NaN")
		}
	}
	{
		// Test conversion to KilometersPerHour.
		// No expected conversion value provided for KilometersPerHour, verifying result is not NaN.
		result := a.KilometersPerHour()
		cacheResult := a.KilometersPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilometersPerHour returned NaN")
		}
	}
}

func TestSpeed_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.SpeedFactory{}
	a, err := factory.CreateSpeed(90, units.SpeedMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.SpeedMeterPerSecond {
		t.Errorf("expected default unit MeterPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.SpeedMeterPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.SpeedDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.SpeedMeterPerSecond {
		t.Errorf("expected unit MeterPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestSpeedFactory_FromDto(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedMeterPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.SpeedDto{
        Value: math.NaN(),
        Unit:  units.SpeedMeterPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test MeterPerSecond conversion
    meters_per_secondDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedMeterPerSecond,
    }
    
    var meters_per_secondResult *units.Speed
    meters_per_secondResult, err = factory.FromDto(meters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MeterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_per_secondResult.Convert(units.SpeedMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterPerSecond = %v, want %v", converted, 100)
    }
    // Test MeterPerMinute conversion
    meters_per_minutesDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedMeterPerMinute,
    }
    
    var meters_per_minutesResult *units.Speed
    meters_per_minutesResult, err = factory.FromDto(meters_per_minutesDto)
    if err != nil {
        t.Errorf("FromDto() with MeterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_per_minutesResult.Convert(units.SpeedMeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterPerMinute = %v, want %v", converted, 100)
    }
    // Test MeterPerHour conversion
    meters_per_hourDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedMeterPerHour,
    }
    
    var meters_per_hourResult *units.Speed
    meters_per_hourResult, err = factory.FromDto(meters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MeterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_per_hourResult.Convert(units.SpeedMeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterPerHour = %v, want %v", converted, 100)
    }
    // Test FootPerSecond conversion
    feet_per_secondDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedFootPerSecond,
    }
    
    var feet_per_secondResult *units.Speed
    feet_per_secondResult, err = factory.FromDto(feet_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with FootPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_per_secondResult.Convert(units.SpeedFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootPerSecond = %v, want %v", converted, 100)
    }
    // Test FootPerMinute conversion
    feet_per_minuteDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedFootPerMinute,
    }
    
    var feet_per_minuteResult *units.Speed
    feet_per_minuteResult, err = factory.FromDto(feet_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with FootPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_per_minuteResult.Convert(units.SpeedFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootPerMinute = %v, want %v", converted, 100)
    }
    // Test FootPerHour conversion
    feet_per_hourDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedFootPerHour,
    }
    
    var feet_per_hourResult *units.Speed
    feet_per_hourResult, err = factory.FromDto(feet_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with FootPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_per_hourResult.Convert(units.SpeedFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootPerHour = %v, want %v", converted, 100)
    }
    // Test UsSurveyFootPerSecond conversion
    us_survey_feet_per_secondDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedUsSurveyFootPerSecond,
    }
    
    var us_survey_feet_per_secondResult *units.Speed
    us_survey_feet_per_secondResult, err = factory.FromDto(us_survey_feet_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with UsSurveyFootPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_survey_feet_per_secondResult.Convert(units.SpeedUsSurveyFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsSurveyFootPerSecond = %v, want %v", converted, 100)
    }
    // Test UsSurveyFootPerMinute conversion
    us_survey_feet_per_minuteDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedUsSurveyFootPerMinute,
    }
    
    var us_survey_feet_per_minuteResult *units.Speed
    us_survey_feet_per_minuteResult, err = factory.FromDto(us_survey_feet_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with UsSurveyFootPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_survey_feet_per_minuteResult.Convert(units.SpeedUsSurveyFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsSurveyFootPerMinute = %v, want %v", converted, 100)
    }
    // Test UsSurveyFootPerHour conversion
    us_survey_feet_per_hourDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedUsSurveyFootPerHour,
    }
    
    var us_survey_feet_per_hourResult *units.Speed
    us_survey_feet_per_hourResult, err = factory.FromDto(us_survey_feet_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with UsSurveyFootPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_survey_feet_per_hourResult.Convert(units.SpeedUsSurveyFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsSurveyFootPerHour = %v, want %v", converted, 100)
    }
    // Test InchPerSecond conversion
    inches_per_secondDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedInchPerSecond,
    }
    
    var inches_per_secondResult *units.Speed
    inches_per_secondResult, err = factory.FromDto(inches_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with InchPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_per_secondResult.Convert(units.SpeedInchPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchPerSecond = %v, want %v", converted, 100)
    }
    // Test InchPerMinute conversion
    inches_per_minuteDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedInchPerMinute,
    }
    
    var inches_per_minuteResult *units.Speed
    inches_per_minuteResult, err = factory.FromDto(inches_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with InchPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_per_minuteResult.Convert(units.SpeedInchPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchPerMinute = %v, want %v", converted, 100)
    }
    // Test InchPerHour conversion
    inches_per_hourDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedInchPerHour,
    }
    
    var inches_per_hourResult *units.Speed
    inches_per_hourResult, err = factory.FromDto(inches_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with InchPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_per_hourResult.Convert(units.SpeedInchPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchPerHour = %v, want %v", converted, 100)
    }
    // Test YardPerSecond conversion
    yards_per_secondDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedYardPerSecond,
    }
    
    var yards_per_secondResult *units.Speed
    yards_per_secondResult, err = factory.FromDto(yards_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with YardPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = yards_per_secondResult.Convert(units.SpeedYardPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for YardPerSecond = %v, want %v", converted, 100)
    }
    // Test YardPerMinute conversion
    yards_per_minuteDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedYardPerMinute,
    }
    
    var yards_per_minuteResult *units.Speed
    yards_per_minuteResult, err = factory.FromDto(yards_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with YardPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = yards_per_minuteResult.Convert(units.SpeedYardPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for YardPerMinute = %v, want %v", converted, 100)
    }
    // Test YardPerHour conversion
    yards_per_hourDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedYardPerHour,
    }
    
    var yards_per_hourResult *units.Speed
    yards_per_hourResult, err = factory.FromDto(yards_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with YardPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = yards_per_hourResult.Convert(units.SpeedYardPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for YardPerHour = %v, want %v", converted, 100)
    }
    // Test Knot conversion
    knotsDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedKnot,
    }
    
    var knotsResult *units.Speed
    knotsResult, err = factory.FromDto(knotsDto)
    if err != nil {
        t.Errorf("FromDto() with Knot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = knotsResult.Convert(units.SpeedKnot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Knot = %v, want %v", converted, 100)
    }
    // Test MilePerHour conversion
    miles_per_hourDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedMilePerHour,
    }
    
    var miles_per_hourResult *units.Speed
    miles_per_hourResult, err = factory.FromDto(miles_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MilePerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = miles_per_hourResult.Convert(units.SpeedMilePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilePerHour = %v, want %v", converted, 100)
    }
    // Test Mach conversion
    machDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedMach,
    }
    
    var machResult *units.Speed
    machResult, err = factory.FromDto(machDto)
    if err != nil {
        t.Errorf("FromDto() with Mach returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = machResult.Convert(units.SpeedMach)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mach = %v, want %v", converted, 100)
    }
    // Test NanometerPerSecond conversion
    nanometers_per_secondDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedNanometerPerSecond,
    }
    
    var nanometers_per_secondResult *units.Speed
    nanometers_per_secondResult, err = factory.FromDto(nanometers_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with NanometerPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanometers_per_secondResult.Convert(units.SpeedNanometerPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanometerPerSecond = %v, want %v", converted, 100)
    }
    // Test MicrometerPerSecond conversion
    micrometers_per_secondDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedMicrometerPerSecond,
    }
    
    var micrometers_per_secondResult *units.Speed
    micrometers_per_secondResult, err = factory.FromDto(micrometers_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MicrometerPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrometers_per_secondResult.Convert(units.SpeedMicrometerPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrometerPerSecond = %v, want %v", converted, 100)
    }
    // Test MillimeterPerSecond conversion
    millimeters_per_secondDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedMillimeterPerSecond,
    }
    
    var millimeters_per_secondResult *units.Speed
    millimeters_per_secondResult, err = factory.FromDto(millimeters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_per_secondResult.Convert(units.SpeedMillimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterPerSecond = %v, want %v", converted, 100)
    }
    // Test CentimeterPerSecond conversion
    centimeters_per_secondDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedCentimeterPerSecond,
    }
    
    var centimeters_per_secondResult *units.Speed
    centimeters_per_secondResult, err = factory.FromDto(centimeters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with CentimeterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_per_secondResult.Convert(units.SpeedCentimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterPerSecond = %v, want %v", converted, 100)
    }
    // Test DecimeterPerSecond conversion
    decimeters_per_secondDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedDecimeterPerSecond,
    }
    
    var decimeters_per_secondResult *units.Speed
    decimeters_per_secondResult, err = factory.FromDto(decimeters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DecimeterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimeters_per_secondResult.Convert(units.SpeedDecimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimeterPerSecond = %v, want %v", converted, 100)
    }
    // Test KilometerPerSecond conversion
    kilometers_per_secondDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedKilometerPerSecond,
    }
    
    var kilometers_per_secondResult *units.Speed
    kilometers_per_secondResult, err = factory.FromDto(kilometers_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilometerPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometers_per_secondResult.Convert(units.SpeedKilometerPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilometerPerSecond = %v, want %v", converted, 100)
    }
    // Test NanometerPerMinute conversion
    nanometers_per_minutesDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedNanometerPerMinute,
    }
    
    var nanometers_per_minutesResult *units.Speed
    nanometers_per_minutesResult, err = factory.FromDto(nanometers_per_minutesDto)
    if err != nil {
        t.Errorf("FromDto() with NanometerPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanometers_per_minutesResult.Convert(units.SpeedNanometerPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanometerPerMinute = %v, want %v", converted, 100)
    }
    // Test MicrometerPerMinute conversion
    micrometers_per_minutesDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedMicrometerPerMinute,
    }
    
    var micrometers_per_minutesResult *units.Speed
    micrometers_per_minutesResult, err = factory.FromDto(micrometers_per_minutesDto)
    if err != nil {
        t.Errorf("FromDto() with MicrometerPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrometers_per_minutesResult.Convert(units.SpeedMicrometerPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrometerPerMinute = %v, want %v", converted, 100)
    }
    // Test MillimeterPerMinute conversion
    millimeters_per_minutesDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedMillimeterPerMinute,
    }
    
    var millimeters_per_minutesResult *units.Speed
    millimeters_per_minutesResult, err = factory.FromDto(millimeters_per_minutesDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_per_minutesResult.Convert(units.SpeedMillimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterPerMinute = %v, want %v", converted, 100)
    }
    // Test CentimeterPerMinute conversion
    centimeters_per_minutesDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedCentimeterPerMinute,
    }
    
    var centimeters_per_minutesResult *units.Speed
    centimeters_per_minutesResult, err = factory.FromDto(centimeters_per_minutesDto)
    if err != nil {
        t.Errorf("FromDto() with CentimeterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_per_minutesResult.Convert(units.SpeedCentimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterPerMinute = %v, want %v", converted, 100)
    }
    // Test DecimeterPerMinute conversion
    decimeters_per_minutesDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedDecimeterPerMinute,
    }
    
    var decimeters_per_minutesResult *units.Speed
    decimeters_per_minutesResult, err = factory.FromDto(decimeters_per_minutesDto)
    if err != nil {
        t.Errorf("FromDto() with DecimeterPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimeters_per_minutesResult.Convert(units.SpeedDecimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimeterPerMinute = %v, want %v", converted, 100)
    }
    // Test KilometerPerMinute conversion
    kilometers_per_minutesDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedKilometerPerMinute,
    }
    
    var kilometers_per_minutesResult *units.Speed
    kilometers_per_minutesResult, err = factory.FromDto(kilometers_per_minutesDto)
    if err != nil {
        t.Errorf("FromDto() with KilometerPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometers_per_minutesResult.Convert(units.SpeedKilometerPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilometerPerMinute = %v, want %v", converted, 100)
    }
    // Test MillimeterPerHour conversion
    millimeters_per_hourDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedMillimeterPerHour,
    }
    
    var millimeters_per_hourResult *units.Speed
    millimeters_per_hourResult, err = factory.FromDto(millimeters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_per_hourResult.Convert(units.SpeedMillimeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterPerHour = %v, want %v", converted, 100)
    }
    // Test CentimeterPerHour conversion
    centimeters_per_hourDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedCentimeterPerHour,
    }
    
    var centimeters_per_hourResult *units.Speed
    centimeters_per_hourResult, err = factory.FromDto(centimeters_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with CentimeterPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_per_hourResult.Convert(units.SpeedCentimeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterPerHour = %v, want %v", converted, 100)
    }
    // Test KilometerPerHour conversion
    kilometers_per_hourDto := units.SpeedDto{
        Value: 100,
        Unit:  units.SpeedKilometerPerHour,
    }
    
    var kilometers_per_hourResult *units.Speed
    kilometers_per_hourResult, err = factory.FromDto(kilometers_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with KilometerPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometers_per_hourResult.Convert(units.SpeedKilometerPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilometerPerHour = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.SpeedDto{
        Value: 0,
        Unit:  units.SpeedMeterPerSecond,
    }
    
    var zeroResult *units.Speed
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestSpeedFactory_FromDtoJSON(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "MeterPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "MeterPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.SpeedDto{
        Value: nanValue,
        Unit:  units.SpeedMeterPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with MeterPerSecond unit
    meters_per_secondJSON := []byte(`{"value": 100, "unit": "MeterPerSecond"}`)
    meters_per_secondResult, err := factory.FromDtoJSON(meters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_per_secondResult.Convert(units.SpeedMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MeterPerMinute unit
    meters_per_minutesJSON := []byte(`{"value": 100, "unit": "MeterPerMinute"}`)
    meters_per_minutesResult, err := factory.FromDtoJSON(meters_per_minutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_per_minutesResult.Convert(units.SpeedMeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MeterPerHour unit
    meters_per_hourJSON := []byte(`{"value": 100, "unit": "MeterPerHour"}`)
    meters_per_hourResult, err := factory.FromDtoJSON(meters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_per_hourResult.Convert(units.SpeedMeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with FootPerSecond unit
    feet_per_secondJSON := []byte(`{"value": 100, "unit": "FootPerSecond"}`)
    feet_per_secondResult, err := factory.FromDtoJSON(feet_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FootPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_per_secondResult.Convert(units.SpeedFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with FootPerMinute unit
    feet_per_minuteJSON := []byte(`{"value": 100, "unit": "FootPerMinute"}`)
    feet_per_minuteResult, err := factory.FromDtoJSON(feet_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FootPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_per_minuteResult.Convert(units.SpeedFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with FootPerHour unit
    feet_per_hourJSON := []byte(`{"value": 100, "unit": "FootPerHour"}`)
    feet_per_hourResult, err := factory.FromDtoJSON(feet_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FootPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_per_hourResult.Convert(units.SpeedFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with UsSurveyFootPerSecond unit
    us_survey_feet_per_secondJSON := []byte(`{"value": 100, "unit": "UsSurveyFootPerSecond"}`)
    us_survey_feet_per_secondResult, err := factory.FromDtoJSON(us_survey_feet_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsSurveyFootPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_survey_feet_per_secondResult.Convert(units.SpeedUsSurveyFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsSurveyFootPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with UsSurveyFootPerMinute unit
    us_survey_feet_per_minuteJSON := []byte(`{"value": 100, "unit": "UsSurveyFootPerMinute"}`)
    us_survey_feet_per_minuteResult, err := factory.FromDtoJSON(us_survey_feet_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsSurveyFootPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_survey_feet_per_minuteResult.Convert(units.SpeedUsSurveyFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsSurveyFootPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with UsSurveyFootPerHour unit
    us_survey_feet_per_hourJSON := []byte(`{"value": 100, "unit": "UsSurveyFootPerHour"}`)
    us_survey_feet_per_hourResult, err := factory.FromDtoJSON(us_survey_feet_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsSurveyFootPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_survey_feet_per_hourResult.Convert(units.SpeedUsSurveyFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsSurveyFootPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with InchPerSecond unit
    inches_per_secondJSON := []byte(`{"value": 100, "unit": "InchPerSecond"}`)
    inches_per_secondResult, err := factory.FromDtoJSON(inches_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InchPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_per_secondResult.Convert(units.SpeedInchPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with InchPerMinute unit
    inches_per_minuteJSON := []byte(`{"value": 100, "unit": "InchPerMinute"}`)
    inches_per_minuteResult, err := factory.FromDtoJSON(inches_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InchPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_per_minuteResult.Convert(units.SpeedInchPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with InchPerHour unit
    inches_per_hourJSON := []byte(`{"value": 100, "unit": "InchPerHour"}`)
    inches_per_hourResult, err := factory.FromDtoJSON(inches_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InchPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_per_hourResult.Convert(units.SpeedInchPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with YardPerSecond unit
    yards_per_secondJSON := []byte(`{"value": 100, "unit": "YardPerSecond"}`)
    yards_per_secondResult, err := factory.FromDtoJSON(yards_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with YardPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = yards_per_secondResult.Convert(units.SpeedYardPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for YardPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with YardPerMinute unit
    yards_per_minuteJSON := []byte(`{"value": 100, "unit": "YardPerMinute"}`)
    yards_per_minuteResult, err := factory.FromDtoJSON(yards_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with YardPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = yards_per_minuteResult.Convert(units.SpeedYardPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for YardPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with YardPerHour unit
    yards_per_hourJSON := []byte(`{"value": 100, "unit": "YardPerHour"}`)
    yards_per_hourResult, err := factory.FromDtoJSON(yards_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with YardPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = yards_per_hourResult.Convert(units.SpeedYardPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for YardPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with Knot unit
    knotsJSON := []byte(`{"value": 100, "unit": "Knot"}`)
    knotsResult, err := factory.FromDtoJSON(knotsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Knot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = knotsResult.Convert(units.SpeedKnot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Knot = %v, want %v", converted, 100)
    }
    // Test JSON with MilePerHour unit
    miles_per_hourJSON := []byte(`{"value": 100, "unit": "MilePerHour"}`)
    miles_per_hourResult, err := factory.FromDtoJSON(miles_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilePerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = miles_per_hourResult.Convert(units.SpeedMilePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilePerHour = %v, want %v", converted, 100)
    }
    // Test JSON with Mach unit
    machJSON := []byte(`{"value": 100, "unit": "Mach"}`)
    machResult, err := factory.FromDtoJSON(machJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Mach unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = machResult.Convert(units.SpeedMach)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mach = %v, want %v", converted, 100)
    }
    // Test JSON with NanometerPerSecond unit
    nanometers_per_secondJSON := []byte(`{"value": 100, "unit": "NanometerPerSecond"}`)
    nanometers_per_secondResult, err := factory.FromDtoJSON(nanometers_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanometerPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanometers_per_secondResult.Convert(units.SpeedNanometerPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanometerPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicrometerPerSecond unit
    micrometers_per_secondJSON := []byte(`{"value": 100, "unit": "MicrometerPerSecond"}`)
    micrometers_per_secondResult, err := factory.FromDtoJSON(micrometers_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrometerPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrometers_per_secondResult.Convert(units.SpeedMicrometerPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrometerPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterPerSecond unit
    millimeters_per_secondJSON := []byte(`{"value": 100, "unit": "MillimeterPerSecond"}`)
    millimeters_per_secondResult, err := factory.FromDtoJSON(millimeters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_per_secondResult.Convert(units.SpeedMillimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with CentimeterPerSecond unit
    centimeters_per_secondJSON := []byte(`{"value": 100, "unit": "CentimeterPerSecond"}`)
    centimeters_per_secondResult, err := factory.FromDtoJSON(centimeters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentimeterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_per_secondResult.Convert(units.SpeedCentimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DecimeterPerSecond unit
    decimeters_per_secondJSON := []byte(`{"value": 100, "unit": "DecimeterPerSecond"}`)
    decimeters_per_secondResult, err := factory.FromDtoJSON(decimeters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecimeterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimeters_per_secondResult.Convert(units.SpeedDecimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimeterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilometerPerSecond unit
    kilometers_per_secondJSON := []byte(`{"value": 100, "unit": "KilometerPerSecond"}`)
    kilometers_per_secondResult, err := factory.FromDtoJSON(kilometers_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilometerPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometers_per_secondResult.Convert(units.SpeedKilometerPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilometerPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with NanometerPerMinute unit
    nanometers_per_minutesJSON := []byte(`{"value": 100, "unit": "NanometerPerMinute"}`)
    nanometers_per_minutesResult, err := factory.FromDtoJSON(nanometers_per_minutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanometerPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanometers_per_minutesResult.Convert(units.SpeedNanometerPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanometerPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MicrometerPerMinute unit
    micrometers_per_minutesJSON := []byte(`{"value": 100, "unit": "MicrometerPerMinute"}`)
    micrometers_per_minutesResult, err := factory.FromDtoJSON(micrometers_per_minutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrometerPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrometers_per_minutesResult.Convert(units.SpeedMicrometerPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrometerPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterPerMinute unit
    millimeters_per_minutesJSON := []byte(`{"value": 100, "unit": "MillimeterPerMinute"}`)
    millimeters_per_minutesResult, err := factory.FromDtoJSON(millimeters_per_minutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_per_minutesResult.Convert(units.SpeedMillimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with CentimeterPerMinute unit
    centimeters_per_minutesJSON := []byte(`{"value": 100, "unit": "CentimeterPerMinute"}`)
    centimeters_per_minutesResult, err := factory.FromDtoJSON(centimeters_per_minutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentimeterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_per_minutesResult.Convert(units.SpeedCentimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with DecimeterPerMinute unit
    decimeters_per_minutesJSON := []byte(`{"value": 100, "unit": "DecimeterPerMinute"}`)
    decimeters_per_minutesResult, err := factory.FromDtoJSON(decimeters_per_minutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecimeterPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimeters_per_minutesResult.Convert(units.SpeedDecimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimeterPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with KilometerPerMinute unit
    kilometers_per_minutesJSON := []byte(`{"value": 100, "unit": "KilometerPerMinute"}`)
    kilometers_per_minutesResult, err := factory.FromDtoJSON(kilometers_per_minutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilometerPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometers_per_minutesResult.Convert(units.SpeedKilometerPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilometerPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterPerHour unit
    millimeters_per_hourJSON := []byte(`{"value": 100, "unit": "MillimeterPerHour"}`)
    millimeters_per_hourResult, err := factory.FromDtoJSON(millimeters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_per_hourResult.Convert(units.SpeedMillimeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with CentimeterPerHour unit
    centimeters_per_hourJSON := []byte(`{"value": 100, "unit": "CentimeterPerHour"}`)
    centimeters_per_hourResult, err := factory.FromDtoJSON(centimeters_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentimeterPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_per_hourResult.Convert(units.SpeedCentimeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with KilometerPerHour unit
    kilometers_per_hourJSON := []byte(`{"value": 100, "unit": "KilometerPerHour"}`)
    kilometers_per_hourResult, err := factory.FromDtoJSON(kilometers_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilometerPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometers_per_hourResult.Convert(units.SpeedKilometerPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilometerPerHour = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "MeterPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromMetersPerSecond function
func TestSpeedFactory_FromMetersPerSecond(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMetersPerSecond(100)
    if err != nil {
        t.Errorf("FromMetersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMetersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMetersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMetersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMetersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMetersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMetersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMetersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMetersPerSecond(0)
    if err != nil {
        t.Errorf("FromMetersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedMeterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMetersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMetersPerMinutes function
func TestSpeedFactory_FromMetersPerMinutes(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMetersPerMinutes(100)
    if err != nil {
        t.Errorf("FromMetersPerMinutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedMeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMetersPerMinutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMetersPerMinutes(math.NaN())
    if err == nil {
        t.Error("FromMetersPerMinutes() with NaN value should return error")
    }

    _, err = factory.FromMetersPerMinutes(math.Inf(1))
    if err == nil {
        t.Error("FromMetersPerMinutes() with +Inf value should return error")
    }

    _, err = factory.FromMetersPerMinutes(math.Inf(-1))
    if err == nil {
        t.Error("FromMetersPerMinutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMetersPerMinutes(0)
    if err != nil {
        t.Errorf("FromMetersPerMinutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedMeterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMetersPerMinutes() with zero value = %v, want 0", converted)
    }
}
// Test FromMetersPerHour function
func TestSpeedFactory_FromMetersPerHour(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMetersPerHour(100)
    if err != nil {
        t.Errorf("FromMetersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedMeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMetersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMetersPerHour(math.NaN())
    if err == nil {
        t.Error("FromMetersPerHour() with NaN value should return error")
    }

    _, err = factory.FromMetersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMetersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMetersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMetersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMetersPerHour(0)
    if err != nil {
        t.Errorf("FromMetersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedMeterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMetersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromFeetPerSecond function
func TestSpeedFactory_FromFeetPerSecond(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFeetPerSecond(100)
    if err != nil {
        t.Errorf("FromFeetPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFeetPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFeetPerSecond(math.NaN())
    if err == nil {
        t.Error("FromFeetPerSecond() with NaN value should return error")
    }

    _, err = factory.FromFeetPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromFeetPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromFeetPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromFeetPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFeetPerSecond(0)
    if err != nil {
        t.Errorf("FromFeetPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedFootPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFeetPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromFeetPerMinute function
func TestSpeedFactory_FromFeetPerMinute(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFeetPerMinute(100)
    if err != nil {
        t.Errorf("FromFeetPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFeetPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFeetPerMinute(math.NaN())
    if err == nil {
        t.Error("FromFeetPerMinute() with NaN value should return error")
    }

    _, err = factory.FromFeetPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromFeetPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromFeetPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromFeetPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFeetPerMinute(0)
    if err != nil {
        t.Errorf("FromFeetPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedFootPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFeetPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromFeetPerHour function
func TestSpeedFactory_FromFeetPerHour(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFeetPerHour(100)
    if err != nil {
        t.Errorf("FromFeetPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFeetPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFeetPerHour(math.NaN())
    if err == nil {
        t.Error("FromFeetPerHour() with NaN value should return error")
    }

    _, err = factory.FromFeetPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromFeetPerHour() with +Inf value should return error")
    }

    _, err = factory.FromFeetPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromFeetPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFeetPerHour(0)
    if err != nil {
        t.Errorf("FromFeetPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedFootPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFeetPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromUsSurveyFeetPerSecond function
func TestSpeedFactory_FromUsSurveyFeetPerSecond(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsSurveyFeetPerSecond(100)
    if err != nil {
        t.Errorf("FromUsSurveyFeetPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedUsSurveyFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsSurveyFeetPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsSurveyFeetPerSecond(math.NaN())
    if err == nil {
        t.Error("FromUsSurveyFeetPerSecond() with NaN value should return error")
    }

    _, err = factory.FromUsSurveyFeetPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromUsSurveyFeetPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromUsSurveyFeetPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromUsSurveyFeetPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsSurveyFeetPerSecond(0)
    if err != nil {
        t.Errorf("FromUsSurveyFeetPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedUsSurveyFootPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsSurveyFeetPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromUsSurveyFeetPerMinute function
func TestSpeedFactory_FromUsSurveyFeetPerMinute(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsSurveyFeetPerMinute(100)
    if err != nil {
        t.Errorf("FromUsSurveyFeetPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedUsSurveyFootPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsSurveyFeetPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsSurveyFeetPerMinute(math.NaN())
    if err == nil {
        t.Error("FromUsSurveyFeetPerMinute() with NaN value should return error")
    }

    _, err = factory.FromUsSurveyFeetPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromUsSurveyFeetPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromUsSurveyFeetPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromUsSurveyFeetPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsSurveyFeetPerMinute(0)
    if err != nil {
        t.Errorf("FromUsSurveyFeetPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedUsSurveyFootPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsSurveyFeetPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromUsSurveyFeetPerHour function
func TestSpeedFactory_FromUsSurveyFeetPerHour(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsSurveyFeetPerHour(100)
    if err != nil {
        t.Errorf("FromUsSurveyFeetPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedUsSurveyFootPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsSurveyFeetPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsSurveyFeetPerHour(math.NaN())
    if err == nil {
        t.Error("FromUsSurveyFeetPerHour() with NaN value should return error")
    }

    _, err = factory.FromUsSurveyFeetPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromUsSurveyFeetPerHour() with +Inf value should return error")
    }

    _, err = factory.FromUsSurveyFeetPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromUsSurveyFeetPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsSurveyFeetPerHour(0)
    if err != nil {
        t.Errorf("FromUsSurveyFeetPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedUsSurveyFootPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsSurveyFeetPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromInchesPerSecond function
func TestSpeedFactory_FromInchesPerSecond(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInchesPerSecond(100)
    if err != nil {
        t.Errorf("FromInchesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedInchPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInchesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInchesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromInchesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromInchesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromInchesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromInchesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromInchesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInchesPerSecond(0)
    if err != nil {
        t.Errorf("FromInchesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedInchPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInchesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromInchesPerMinute function
func TestSpeedFactory_FromInchesPerMinute(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInchesPerMinute(100)
    if err != nil {
        t.Errorf("FromInchesPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedInchPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInchesPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInchesPerMinute(math.NaN())
    if err == nil {
        t.Error("FromInchesPerMinute() with NaN value should return error")
    }

    _, err = factory.FromInchesPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromInchesPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromInchesPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromInchesPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInchesPerMinute(0)
    if err != nil {
        t.Errorf("FromInchesPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedInchPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInchesPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromInchesPerHour function
func TestSpeedFactory_FromInchesPerHour(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInchesPerHour(100)
    if err != nil {
        t.Errorf("FromInchesPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedInchPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInchesPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInchesPerHour(math.NaN())
    if err == nil {
        t.Error("FromInchesPerHour() with NaN value should return error")
    }

    _, err = factory.FromInchesPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromInchesPerHour() with +Inf value should return error")
    }

    _, err = factory.FromInchesPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromInchesPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInchesPerHour(0)
    if err != nil {
        t.Errorf("FromInchesPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedInchPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInchesPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromYardsPerSecond function
func TestSpeedFactory_FromYardsPerSecond(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromYardsPerSecond(100)
    if err != nil {
        t.Errorf("FromYardsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedYardPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromYardsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromYardsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromYardsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromYardsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromYardsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromYardsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromYardsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromYardsPerSecond(0)
    if err != nil {
        t.Errorf("FromYardsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedYardPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromYardsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromYardsPerMinute function
func TestSpeedFactory_FromYardsPerMinute(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromYardsPerMinute(100)
    if err != nil {
        t.Errorf("FromYardsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedYardPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromYardsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromYardsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromYardsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromYardsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromYardsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromYardsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromYardsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromYardsPerMinute(0)
    if err != nil {
        t.Errorf("FromYardsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedYardPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromYardsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromYardsPerHour function
func TestSpeedFactory_FromYardsPerHour(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromYardsPerHour(100)
    if err != nil {
        t.Errorf("FromYardsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedYardPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromYardsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromYardsPerHour(math.NaN())
    if err == nil {
        t.Error("FromYardsPerHour() with NaN value should return error")
    }

    _, err = factory.FromYardsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromYardsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromYardsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromYardsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromYardsPerHour(0)
    if err != nil {
        t.Errorf("FromYardsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedYardPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromYardsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromKnots function
func TestSpeedFactory_FromKnots(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKnots(100)
    if err != nil {
        t.Errorf("FromKnots() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedKnot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKnots() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKnots(math.NaN())
    if err == nil {
        t.Error("FromKnots() with NaN value should return error")
    }

    _, err = factory.FromKnots(math.Inf(1))
    if err == nil {
        t.Error("FromKnots() with +Inf value should return error")
    }

    _, err = factory.FromKnots(math.Inf(-1))
    if err == nil {
        t.Error("FromKnots() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKnots(0)
    if err != nil {
        t.Errorf("FromKnots() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedKnot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKnots() with zero value = %v, want 0", converted)
    }
}
// Test FromMilesPerHour function
func TestSpeedFactory_FromMilesPerHour(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilesPerHour(100)
    if err != nil {
        t.Errorf("FromMilesPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedMilePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilesPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilesPerHour(math.NaN())
    if err == nil {
        t.Error("FromMilesPerHour() with NaN value should return error")
    }

    _, err = factory.FromMilesPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMilesPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMilesPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMilesPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilesPerHour(0)
    if err != nil {
        t.Errorf("FromMilesPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedMilePerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilesPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromMach function
func TestSpeedFactory_FromMach(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMach(100)
    if err != nil {
        t.Errorf("FromMach() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedMach)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMach() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMach(math.NaN())
    if err == nil {
        t.Error("FromMach() with NaN value should return error")
    }

    _, err = factory.FromMach(math.Inf(1))
    if err == nil {
        t.Error("FromMach() with +Inf value should return error")
    }

    _, err = factory.FromMach(math.Inf(-1))
    if err == nil {
        t.Error("FromMach() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMach(0)
    if err != nil {
        t.Errorf("FromMach() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedMach)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMach() with zero value = %v, want 0", converted)
    }
}
// Test FromNanometersPerSecond function
func TestSpeedFactory_FromNanometersPerSecond(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanometersPerSecond(100)
    if err != nil {
        t.Errorf("FromNanometersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedNanometerPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanometersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanometersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromNanometersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromNanometersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromNanometersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromNanometersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromNanometersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanometersPerSecond(0)
    if err != nil {
        t.Errorf("FromNanometersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedNanometerPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanometersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrometersPerSecond function
func TestSpeedFactory_FromMicrometersPerSecond(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrometersPerSecond(100)
    if err != nil {
        t.Errorf("FromMicrometersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedMicrometerPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrometersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrometersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMicrometersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMicrometersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMicrometersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMicrometersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrometersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrometersPerSecond(0)
    if err != nil {
        t.Errorf("FromMicrometersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedMicrometerPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrometersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimetersPerSecond function
func TestSpeedFactory_FromMillimetersPerSecond(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimetersPerSecond(100)
    if err != nil {
        t.Errorf("FromMillimetersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedMillimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimetersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimetersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMillimetersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMillimetersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMillimetersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMillimetersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimetersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimetersPerSecond(0)
    if err != nil {
        t.Errorf("FromMillimetersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedMillimeterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimetersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromCentimetersPerSecond function
func TestSpeedFactory_FromCentimetersPerSecond(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentimetersPerSecond(100)
    if err != nil {
        t.Errorf("FromCentimetersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedCentimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentimetersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentimetersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromCentimetersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromCentimetersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromCentimetersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromCentimetersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromCentimetersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentimetersPerSecond(0)
    if err != nil {
        t.Errorf("FromCentimetersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedCentimeterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentimetersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDecimetersPerSecond function
func TestSpeedFactory_FromDecimetersPerSecond(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecimetersPerSecond(100)
    if err != nil {
        t.Errorf("FromDecimetersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedDecimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecimetersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecimetersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDecimetersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDecimetersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDecimetersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDecimetersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDecimetersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecimetersPerSecond(0)
    if err != nil {
        t.Errorf("FromDecimetersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedDecimeterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecimetersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilometersPerSecond function
func TestSpeedFactory_FromKilometersPerSecond(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilometersPerSecond(100)
    if err != nil {
        t.Errorf("FromKilometersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedKilometerPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilometersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilometersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilometersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilometersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilometersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilometersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilometersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilometersPerSecond(0)
    if err != nil {
        t.Errorf("FromKilometersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedKilometerPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilometersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromNanometersPerMinutes function
func TestSpeedFactory_FromNanometersPerMinutes(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanometersPerMinutes(100)
    if err != nil {
        t.Errorf("FromNanometersPerMinutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedNanometerPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanometersPerMinutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanometersPerMinutes(math.NaN())
    if err == nil {
        t.Error("FromNanometersPerMinutes() with NaN value should return error")
    }

    _, err = factory.FromNanometersPerMinutes(math.Inf(1))
    if err == nil {
        t.Error("FromNanometersPerMinutes() with +Inf value should return error")
    }

    _, err = factory.FromNanometersPerMinutes(math.Inf(-1))
    if err == nil {
        t.Error("FromNanometersPerMinutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanometersPerMinutes(0)
    if err != nil {
        t.Errorf("FromNanometersPerMinutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedNanometerPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanometersPerMinutes() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrometersPerMinutes function
func TestSpeedFactory_FromMicrometersPerMinutes(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrometersPerMinutes(100)
    if err != nil {
        t.Errorf("FromMicrometersPerMinutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedMicrometerPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrometersPerMinutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrometersPerMinutes(math.NaN())
    if err == nil {
        t.Error("FromMicrometersPerMinutes() with NaN value should return error")
    }

    _, err = factory.FromMicrometersPerMinutes(math.Inf(1))
    if err == nil {
        t.Error("FromMicrometersPerMinutes() with +Inf value should return error")
    }

    _, err = factory.FromMicrometersPerMinutes(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrometersPerMinutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrometersPerMinutes(0)
    if err != nil {
        t.Errorf("FromMicrometersPerMinutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedMicrometerPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrometersPerMinutes() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimetersPerMinutes function
func TestSpeedFactory_FromMillimetersPerMinutes(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimetersPerMinutes(100)
    if err != nil {
        t.Errorf("FromMillimetersPerMinutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedMillimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimetersPerMinutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimetersPerMinutes(math.NaN())
    if err == nil {
        t.Error("FromMillimetersPerMinutes() with NaN value should return error")
    }

    _, err = factory.FromMillimetersPerMinutes(math.Inf(1))
    if err == nil {
        t.Error("FromMillimetersPerMinutes() with +Inf value should return error")
    }

    _, err = factory.FromMillimetersPerMinutes(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimetersPerMinutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimetersPerMinutes(0)
    if err != nil {
        t.Errorf("FromMillimetersPerMinutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedMillimeterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimetersPerMinutes() with zero value = %v, want 0", converted)
    }
}
// Test FromCentimetersPerMinutes function
func TestSpeedFactory_FromCentimetersPerMinutes(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentimetersPerMinutes(100)
    if err != nil {
        t.Errorf("FromCentimetersPerMinutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedCentimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentimetersPerMinutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentimetersPerMinutes(math.NaN())
    if err == nil {
        t.Error("FromCentimetersPerMinutes() with NaN value should return error")
    }

    _, err = factory.FromCentimetersPerMinutes(math.Inf(1))
    if err == nil {
        t.Error("FromCentimetersPerMinutes() with +Inf value should return error")
    }

    _, err = factory.FromCentimetersPerMinutes(math.Inf(-1))
    if err == nil {
        t.Error("FromCentimetersPerMinutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentimetersPerMinutes(0)
    if err != nil {
        t.Errorf("FromCentimetersPerMinutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedCentimeterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentimetersPerMinutes() with zero value = %v, want 0", converted)
    }
}
// Test FromDecimetersPerMinutes function
func TestSpeedFactory_FromDecimetersPerMinutes(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecimetersPerMinutes(100)
    if err != nil {
        t.Errorf("FromDecimetersPerMinutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedDecimeterPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecimetersPerMinutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecimetersPerMinutes(math.NaN())
    if err == nil {
        t.Error("FromDecimetersPerMinutes() with NaN value should return error")
    }

    _, err = factory.FromDecimetersPerMinutes(math.Inf(1))
    if err == nil {
        t.Error("FromDecimetersPerMinutes() with +Inf value should return error")
    }

    _, err = factory.FromDecimetersPerMinutes(math.Inf(-1))
    if err == nil {
        t.Error("FromDecimetersPerMinutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecimetersPerMinutes(0)
    if err != nil {
        t.Errorf("FromDecimetersPerMinutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedDecimeterPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecimetersPerMinutes() with zero value = %v, want 0", converted)
    }
}
// Test FromKilometersPerMinutes function
func TestSpeedFactory_FromKilometersPerMinutes(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilometersPerMinutes(100)
    if err != nil {
        t.Errorf("FromKilometersPerMinutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedKilometerPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilometersPerMinutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilometersPerMinutes(math.NaN())
    if err == nil {
        t.Error("FromKilometersPerMinutes() with NaN value should return error")
    }

    _, err = factory.FromKilometersPerMinutes(math.Inf(1))
    if err == nil {
        t.Error("FromKilometersPerMinutes() with +Inf value should return error")
    }

    _, err = factory.FromKilometersPerMinutes(math.Inf(-1))
    if err == nil {
        t.Error("FromKilometersPerMinutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilometersPerMinutes(0)
    if err != nil {
        t.Errorf("FromKilometersPerMinutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedKilometerPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilometersPerMinutes() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimetersPerHour function
func TestSpeedFactory_FromMillimetersPerHour(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimetersPerHour(100)
    if err != nil {
        t.Errorf("FromMillimetersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedMillimeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimetersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimetersPerHour(math.NaN())
    if err == nil {
        t.Error("FromMillimetersPerHour() with NaN value should return error")
    }

    _, err = factory.FromMillimetersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMillimetersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMillimetersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimetersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimetersPerHour(0)
    if err != nil {
        t.Errorf("FromMillimetersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedMillimeterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimetersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromCentimetersPerHour function
func TestSpeedFactory_FromCentimetersPerHour(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentimetersPerHour(100)
    if err != nil {
        t.Errorf("FromCentimetersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedCentimeterPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentimetersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentimetersPerHour(math.NaN())
    if err == nil {
        t.Error("FromCentimetersPerHour() with NaN value should return error")
    }

    _, err = factory.FromCentimetersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromCentimetersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromCentimetersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromCentimetersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentimetersPerHour(0)
    if err != nil {
        t.Errorf("FromCentimetersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedCentimeterPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentimetersPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromKilometersPerHour function
func TestSpeedFactory_FromKilometersPerHour(t *testing.T) {
    factory := units.SpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilometersPerHour(100)
    if err != nil {
        t.Errorf("FromKilometersPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpeedKilometerPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilometersPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilometersPerHour(math.NaN())
    if err == nil {
        t.Error("FromKilometersPerHour() with NaN value should return error")
    }

    _, err = factory.FromKilometersPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromKilometersPerHour() with +Inf value should return error")
    }

    _, err = factory.FromKilometersPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromKilometersPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilometersPerHour(0)
    if err != nil {
        t.Errorf("FromKilometersPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpeedKilometerPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilometersPerHour() with zero value = %v, want 0", converted)
    }
}

func TestSpeedToString(t *testing.T) {
	factory := units.SpeedFactory{}
	a, err := factory.CreateSpeed(45, units.SpeedMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.SpeedMeterPerSecond, 2)
	expected := "45.00 " + units.GetSpeedAbbreviation(units.SpeedMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.SpeedMeterPerSecond, -1)
	expected = "45 " + units.GetSpeedAbbreviation(units.SpeedMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestSpeed_EqualityAndComparison(t *testing.T) {
	factory := units.SpeedFactory{}
	a1, _ := factory.CreateSpeed(60, units.SpeedMeterPerSecond)
	a2, _ := factory.CreateSpeed(60, units.SpeedMeterPerSecond)
	a3, _ := factory.CreateSpeed(120, units.SpeedMeterPerSecond)

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

func TestSpeed_Arithmetic(t *testing.T) {
	factory := units.SpeedFactory{}
	a1, _ := factory.CreateSpeed(30, units.SpeedMeterPerSecond)
	a2, _ := factory.CreateSpeed(45, units.SpeedMeterPerSecond)

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


func TestGetSpeedAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.SpeedUnits
        want string
    }{
        {
            name: "MeterPerSecond abbreviation",
            unit: units.SpeedMeterPerSecond,
            want: "m/s",
        },
        {
            name: "MeterPerMinute abbreviation",
            unit: units.SpeedMeterPerMinute,
            want: "m/min",
        },
        {
            name: "MeterPerHour abbreviation",
            unit: units.SpeedMeterPerHour,
            want: "m/h",
        },
        {
            name: "FootPerSecond abbreviation",
            unit: units.SpeedFootPerSecond,
            want: "ft/s",
        },
        {
            name: "FootPerMinute abbreviation",
            unit: units.SpeedFootPerMinute,
            want: "ft/min",
        },
        {
            name: "FootPerHour abbreviation",
            unit: units.SpeedFootPerHour,
            want: "ft/h",
        },
        {
            name: "UsSurveyFootPerSecond abbreviation",
            unit: units.SpeedUsSurveyFootPerSecond,
            want: "ftUS/s",
        },
        {
            name: "UsSurveyFootPerMinute abbreviation",
            unit: units.SpeedUsSurveyFootPerMinute,
            want: "ftUS/min",
        },
        {
            name: "UsSurveyFootPerHour abbreviation",
            unit: units.SpeedUsSurveyFootPerHour,
            want: "ftUS/h",
        },
        {
            name: "InchPerSecond abbreviation",
            unit: units.SpeedInchPerSecond,
            want: "in/s",
        },
        {
            name: "InchPerMinute abbreviation",
            unit: units.SpeedInchPerMinute,
            want: "in/min",
        },
        {
            name: "InchPerHour abbreviation",
            unit: units.SpeedInchPerHour,
            want: "in/h",
        },
        {
            name: "YardPerSecond abbreviation",
            unit: units.SpeedYardPerSecond,
            want: "yd/s",
        },
        {
            name: "YardPerMinute abbreviation",
            unit: units.SpeedYardPerMinute,
            want: "yd/min",
        },
        {
            name: "YardPerHour abbreviation",
            unit: units.SpeedYardPerHour,
            want: "yd/h",
        },
        {
            name: "Knot abbreviation",
            unit: units.SpeedKnot,
            want: "kn",
        },
        {
            name: "MilePerHour abbreviation",
            unit: units.SpeedMilePerHour,
            want: "mph",
        },
        {
            name: "Mach abbreviation",
            unit: units.SpeedMach,
            want: "M",
        },
        {
            name: "NanometerPerSecond abbreviation",
            unit: units.SpeedNanometerPerSecond,
            want: "nm/s",
        },
        {
            name: "MicrometerPerSecond abbreviation",
            unit: units.SpeedMicrometerPerSecond,
            want: "μm/s",
        },
        {
            name: "MillimeterPerSecond abbreviation",
            unit: units.SpeedMillimeterPerSecond,
            want: "mm/s",
        },
        {
            name: "CentimeterPerSecond abbreviation",
            unit: units.SpeedCentimeterPerSecond,
            want: "cm/s",
        },
        {
            name: "DecimeterPerSecond abbreviation",
            unit: units.SpeedDecimeterPerSecond,
            want: "dm/s",
        },
        {
            name: "KilometerPerSecond abbreviation",
            unit: units.SpeedKilometerPerSecond,
            want: "km/s",
        },
        {
            name: "NanometerPerMinute abbreviation",
            unit: units.SpeedNanometerPerMinute,
            want: "nm/min",
        },
        {
            name: "MicrometerPerMinute abbreviation",
            unit: units.SpeedMicrometerPerMinute,
            want: "μm/min",
        },
        {
            name: "MillimeterPerMinute abbreviation",
            unit: units.SpeedMillimeterPerMinute,
            want: "mm/min",
        },
        {
            name: "CentimeterPerMinute abbreviation",
            unit: units.SpeedCentimeterPerMinute,
            want: "cm/min",
        },
        {
            name: "DecimeterPerMinute abbreviation",
            unit: units.SpeedDecimeterPerMinute,
            want: "dm/min",
        },
        {
            name: "KilometerPerMinute abbreviation",
            unit: units.SpeedKilometerPerMinute,
            want: "km/min",
        },
        {
            name: "MillimeterPerHour abbreviation",
            unit: units.SpeedMillimeterPerHour,
            want: "mm/h",
        },
        {
            name: "CentimeterPerHour abbreviation",
            unit: units.SpeedCentimeterPerHour,
            want: "cm/h",
        },
        {
            name: "KilometerPerHour abbreviation",
            unit: units.SpeedKilometerPerHour,
            want: "km/h",
        },
        {
            name: "invalid unit",
            unit: units.SpeedUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetSpeedAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetSpeedAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestSpeed_String(t *testing.T) {
    factory := units.SpeedFactory{}
    
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
            unit, err := factory.CreateSpeed(tt.value, units.SpeedMeterPerSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Speed.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestSpeed_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.SpeedFactory{}

	_, err := uf.CreateSpeed(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}