// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestDurationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Second"}`
	
	factory := units.DurationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.DurationSecond {
		t.Errorf("expected unit %v, got %v", units.DurationSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Second"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestDurationDto_ToJSON(t *testing.T) {
	dto := units.DurationDto{
		Value: 45,
		Unit:  units.DurationSecond,
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
	if result["unit"].(string) != string(units.DurationSecond) {
		t.Errorf("expected unit %s, got %v", units.DurationSecond, result["unit"])
	}
}

func TestNewDuration_InvalidValue(t *testing.T) {
	factory := units.DurationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateDuration(math.NaN(), units.DurationSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateDuration(math.Inf(1), units.DurationSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestDurationConversions(t *testing.T) {
	factory := units.DurationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateDuration(180, units.DurationSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Years365.
		// No expected conversion value provided for Years365, verifying result is not NaN.
		result := a.Years365()
		cacheResult := a.Years365()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Years365 returned NaN")
		}
	}
	{
		// Test conversion to Months30.
		// No expected conversion value provided for Months30, verifying result is not NaN.
		result := a.Months30()
		cacheResult := a.Months30()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Months30 returned NaN")
		}
	}
	{
		// Test conversion to Weeks.
		// No expected conversion value provided for Weeks, verifying result is not NaN.
		result := a.Weeks()
		cacheResult := a.Weeks()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Weeks returned NaN")
		}
	}
	{
		// Test conversion to Days.
		// No expected conversion value provided for Days, verifying result is not NaN.
		result := a.Days()
		cacheResult := a.Days()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Days returned NaN")
		}
	}
	{
		// Test conversion to Hours.
		// No expected conversion value provided for Hours, verifying result is not NaN.
		result := a.Hours()
		cacheResult := a.Hours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Hours returned NaN")
		}
	}
	{
		// Test conversion to Minutes.
		// No expected conversion value provided for Minutes, verifying result is not NaN.
		result := a.Minutes()
		cacheResult := a.Minutes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Minutes returned NaN")
		}
	}
	{
		// Test conversion to Seconds.
		// No expected conversion value provided for Seconds, verifying result is not NaN.
		result := a.Seconds()
		cacheResult := a.Seconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Seconds returned NaN")
		}
	}
	{
		// Test conversion to JulianYears.
		// No expected conversion value provided for JulianYears, verifying result is not NaN.
		result := a.JulianYears()
		cacheResult := a.JulianYears()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to JulianYears returned NaN")
		}
	}
	{
		// Test conversion to Sols.
		// No expected conversion value provided for Sols, verifying result is not NaN.
		result := a.Sols()
		cacheResult := a.Sols()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Sols returned NaN")
		}
	}
	{
		// Test conversion to Nanoseconds.
		// No expected conversion value provided for Nanoseconds, verifying result is not NaN.
		result := a.Nanoseconds()
		cacheResult := a.Nanoseconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanoseconds returned NaN")
		}
	}
	{
		// Test conversion to Microseconds.
		// No expected conversion value provided for Microseconds, verifying result is not NaN.
		result := a.Microseconds()
		cacheResult := a.Microseconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microseconds returned NaN")
		}
	}
	{
		// Test conversion to Milliseconds.
		// No expected conversion value provided for Milliseconds, verifying result is not NaN.
		result := a.Milliseconds()
		cacheResult := a.Milliseconds()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Milliseconds returned NaN")
		}
	}
}

func TestDuration_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.DurationFactory{}
	a, err := factory.CreateDuration(90, units.DurationSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.DurationSecond {
		t.Errorf("expected default unit Second, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.DurationYear365
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.DurationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.DurationSecond {
		t.Errorf("expected unit Second, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestDurationFactory_FromDto(t *testing.T) {
    factory := units.DurationFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.DurationDto{
        Value: 100,
        Unit:  units.DurationSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.DurationDto{
        Value: math.NaN(),
        Unit:  units.DurationSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Year365 conversion
    years365Dto := units.DurationDto{
        Value: 100,
        Unit:  units.DurationYear365,
    }
    
    var years365Result *units.Duration
    years365Result, err = factory.FromDto(years365Dto)
    if err != nil {
        t.Errorf("FromDto() with Year365 returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = years365Result.Convert(units.DurationYear365)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Year365 = %v, want %v", converted, 100)
    }
    // Test Month30 conversion
    months30Dto := units.DurationDto{
        Value: 100,
        Unit:  units.DurationMonth30,
    }
    
    var months30Result *units.Duration
    months30Result, err = factory.FromDto(months30Dto)
    if err != nil {
        t.Errorf("FromDto() with Month30 returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = months30Result.Convert(units.DurationMonth30)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Month30 = %v, want %v", converted, 100)
    }
    // Test Week conversion
    weeksDto := units.DurationDto{
        Value: 100,
        Unit:  units.DurationWeek,
    }
    
    var weeksResult *units.Duration
    weeksResult, err = factory.FromDto(weeksDto)
    if err != nil {
        t.Errorf("FromDto() with Week returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = weeksResult.Convert(units.DurationWeek)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Week = %v, want %v", converted, 100)
    }
    // Test Day conversion
    daysDto := units.DurationDto{
        Value: 100,
        Unit:  units.DurationDay,
    }
    
    var daysResult *units.Duration
    daysResult, err = factory.FromDto(daysDto)
    if err != nil {
        t.Errorf("FromDto() with Day returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = daysResult.Convert(units.DurationDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Day = %v, want %v", converted, 100)
    }
    // Test Hour conversion
    hoursDto := units.DurationDto{
        Value: 100,
        Unit:  units.DurationHour,
    }
    
    var hoursResult *units.Duration
    hoursResult, err = factory.FromDto(hoursDto)
    if err != nil {
        t.Errorf("FromDto() with Hour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hoursResult.Convert(units.DurationHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hour = %v, want %v", converted, 100)
    }
    // Test Minute conversion
    minutesDto := units.DurationDto{
        Value: 100,
        Unit:  units.DurationMinute,
    }
    
    var minutesResult *units.Duration
    minutesResult, err = factory.FromDto(minutesDto)
    if err != nil {
        t.Errorf("FromDto() with Minute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = minutesResult.Convert(units.DurationMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Minute = %v, want %v", converted, 100)
    }
    // Test Second conversion
    secondsDto := units.DurationDto{
        Value: 100,
        Unit:  units.DurationSecond,
    }
    
    var secondsResult *units.Duration
    secondsResult, err = factory.FromDto(secondsDto)
    if err != nil {
        t.Errorf("FromDto() with Second returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = secondsResult.Convert(units.DurationSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Second = %v, want %v", converted, 100)
    }
    // Test JulianYear conversion
    julian_yearsDto := units.DurationDto{
        Value: 100,
        Unit:  units.DurationJulianYear,
    }
    
    var julian_yearsResult *units.Duration
    julian_yearsResult, err = factory.FromDto(julian_yearsDto)
    if err != nil {
        t.Errorf("FromDto() with JulianYear returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = julian_yearsResult.Convert(units.DurationJulianYear)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JulianYear = %v, want %v", converted, 100)
    }
    // Test Sol conversion
    solsDto := units.DurationDto{
        Value: 100,
        Unit:  units.DurationSol,
    }
    
    var solsResult *units.Duration
    solsResult, err = factory.FromDto(solsDto)
    if err != nil {
        t.Errorf("FromDto() with Sol returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = solsResult.Convert(units.DurationSol)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Sol = %v, want %v", converted, 100)
    }
    // Test Nanosecond conversion
    nanosecondsDto := units.DurationDto{
        Value: 100,
        Unit:  units.DurationNanosecond,
    }
    
    var nanosecondsResult *units.Duration
    nanosecondsResult, err = factory.FromDto(nanosecondsDto)
    if err != nil {
        t.Errorf("FromDto() with Nanosecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanosecondsResult.Convert(units.DurationNanosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanosecond = %v, want %v", converted, 100)
    }
    // Test Microsecond conversion
    microsecondsDto := units.DurationDto{
        Value: 100,
        Unit:  units.DurationMicrosecond,
    }
    
    var microsecondsResult *units.Duration
    microsecondsResult, err = factory.FromDto(microsecondsDto)
    if err != nil {
        t.Errorf("FromDto() with Microsecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microsecondsResult.Convert(units.DurationMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microsecond = %v, want %v", converted, 100)
    }
    // Test Millisecond conversion
    millisecondsDto := units.DurationDto{
        Value: 100,
        Unit:  units.DurationMillisecond,
    }
    
    var millisecondsResult *units.Duration
    millisecondsResult, err = factory.FromDto(millisecondsDto)
    if err != nil {
        t.Errorf("FromDto() with Millisecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millisecondsResult.Convert(units.DurationMillisecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millisecond = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.DurationDto{
        Value: 0,
        Unit:  units.DurationSecond,
    }
    
    var zeroResult *units.Duration
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestDurationFactory_FromDtoJSON(t *testing.T) {
    factory := units.DurationFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Second"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Second"}`)
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
    nanJSON, _ := json.Marshal(units.DurationDto{
        Value: nanValue,
        Unit:  units.DurationSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Year365 unit
    years365JSON := []byte(`{"value": 100, "unit": "Year365"}`)
    years365Result, err := factory.FromDtoJSON(years365JSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Year365 unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = years365Result.Convert(units.DurationYear365)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Year365 = %v, want %v", converted, 100)
    }
    // Test JSON with Month30 unit
    months30JSON := []byte(`{"value": 100, "unit": "Month30"}`)
    months30Result, err := factory.FromDtoJSON(months30JSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Month30 unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = months30Result.Convert(units.DurationMonth30)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Month30 = %v, want %v", converted, 100)
    }
    // Test JSON with Week unit
    weeksJSON := []byte(`{"value": 100, "unit": "Week"}`)
    weeksResult, err := factory.FromDtoJSON(weeksJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Week unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = weeksResult.Convert(units.DurationWeek)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Week = %v, want %v", converted, 100)
    }
    // Test JSON with Day unit
    daysJSON := []byte(`{"value": 100, "unit": "Day"}`)
    daysResult, err := factory.FromDtoJSON(daysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Day unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = daysResult.Convert(units.DurationDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Day = %v, want %v", converted, 100)
    }
    // Test JSON with Hour unit
    hoursJSON := []byte(`{"value": 100, "unit": "Hour"}`)
    hoursResult, err := factory.FromDtoJSON(hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Hour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hoursResult.Convert(units.DurationHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hour = %v, want %v", converted, 100)
    }
    // Test JSON with Minute unit
    minutesJSON := []byte(`{"value": 100, "unit": "Minute"}`)
    minutesResult, err := factory.FromDtoJSON(minutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Minute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = minutesResult.Convert(units.DurationMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Minute = %v, want %v", converted, 100)
    }
    // Test JSON with Second unit
    secondsJSON := []byte(`{"value": 100, "unit": "Second"}`)
    secondsResult, err := factory.FromDtoJSON(secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Second unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = secondsResult.Convert(units.DurationSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Second = %v, want %v", converted, 100)
    }
    // Test JSON with JulianYear unit
    julian_yearsJSON := []byte(`{"value": 100, "unit": "JulianYear"}`)
    julian_yearsResult, err := factory.FromDtoJSON(julian_yearsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JulianYear unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = julian_yearsResult.Convert(units.DurationJulianYear)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JulianYear = %v, want %v", converted, 100)
    }
    // Test JSON with Sol unit
    solsJSON := []byte(`{"value": 100, "unit": "Sol"}`)
    solsResult, err := factory.FromDtoJSON(solsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Sol unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = solsResult.Convert(units.DurationSol)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Sol = %v, want %v", converted, 100)
    }
    // Test JSON with Nanosecond unit
    nanosecondsJSON := []byte(`{"value": 100, "unit": "Nanosecond"}`)
    nanosecondsResult, err := factory.FromDtoJSON(nanosecondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanosecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanosecondsResult.Convert(units.DurationNanosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanosecond = %v, want %v", converted, 100)
    }
    // Test JSON with Microsecond unit
    microsecondsJSON := []byte(`{"value": 100, "unit": "Microsecond"}`)
    microsecondsResult, err := factory.FromDtoJSON(microsecondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microsecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microsecondsResult.Convert(units.DurationMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microsecond = %v, want %v", converted, 100)
    }
    // Test JSON with Millisecond unit
    millisecondsJSON := []byte(`{"value": 100, "unit": "Millisecond"}`)
    millisecondsResult, err := factory.FromDtoJSON(millisecondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millisecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millisecondsResult.Convert(units.DurationMillisecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millisecond = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Second"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromYears365 function
func TestDurationFactory_FromYears365(t *testing.T) {
    factory := units.DurationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromYears365(100)
    if err != nil {
        t.Errorf("FromYears365() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DurationYear365)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromYears365() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromYears365(math.NaN())
    if err == nil {
        t.Error("FromYears365() with NaN value should return error")
    }

    _, err = factory.FromYears365(math.Inf(1))
    if err == nil {
        t.Error("FromYears365() with +Inf value should return error")
    }

    _, err = factory.FromYears365(math.Inf(-1))
    if err == nil {
        t.Error("FromYears365() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromYears365(0)
    if err != nil {
        t.Errorf("FromYears365() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DurationYear365)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromYears365() with zero value = %v, want 0", converted)
    }
}
// Test FromMonths30 function
func TestDurationFactory_FromMonths30(t *testing.T) {
    factory := units.DurationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMonths30(100)
    if err != nil {
        t.Errorf("FromMonths30() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DurationMonth30)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMonths30() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMonths30(math.NaN())
    if err == nil {
        t.Error("FromMonths30() with NaN value should return error")
    }

    _, err = factory.FromMonths30(math.Inf(1))
    if err == nil {
        t.Error("FromMonths30() with +Inf value should return error")
    }

    _, err = factory.FromMonths30(math.Inf(-1))
    if err == nil {
        t.Error("FromMonths30() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMonths30(0)
    if err != nil {
        t.Errorf("FromMonths30() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DurationMonth30)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMonths30() with zero value = %v, want 0", converted)
    }
}
// Test FromWeeks function
func TestDurationFactory_FromWeeks(t *testing.T) {
    factory := units.DurationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWeeks(100)
    if err != nil {
        t.Errorf("FromWeeks() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DurationWeek)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWeeks() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWeeks(math.NaN())
    if err == nil {
        t.Error("FromWeeks() with NaN value should return error")
    }

    _, err = factory.FromWeeks(math.Inf(1))
    if err == nil {
        t.Error("FromWeeks() with +Inf value should return error")
    }

    _, err = factory.FromWeeks(math.Inf(-1))
    if err == nil {
        t.Error("FromWeeks() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWeeks(0)
    if err != nil {
        t.Errorf("FromWeeks() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DurationWeek)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWeeks() with zero value = %v, want 0", converted)
    }
}
// Test FromDays function
func TestDurationFactory_FromDays(t *testing.T) {
    factory := units.DurationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDays(100)
    if err != nil {
        t.Errorf("FromDays() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DurationDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDays() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDays(math.NaN())
    if err == nil {
        t.Error("FromDays() with NaN value should return error")
    }

    _, err = factory.FromDays(math.Inf(1))
    if err == nil {
        t.Error("FromDays() with +Inf value should return error")
    }

    _, err = factory.FromDays(math.Inf(-1))
    if err == nil {
        t.Error("FromDays() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDays(0)
    if err != nil {
        t.Errorf("FromDays() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DurationDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDays() with zero value = %v, want 0", converted)
    }
}
// Test FromHours function
func TestDurationFactory_FromHours(t *testing.T) {
    factory := units.DurationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHours(100)
    if err != nil {
        t.Errorf("FromHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DurationHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHours(math.NaN())
    if err == nil {
        t.Error("FromHours() with NaN value should return error")
    }

    _, err = factory.FromHours(math.Inf(1))
    if err == nil {
        t.Error("FromHours() with +Inf value should return error")
    }

    _, err = factory.FromHours(math.Inf(-1))
    if err == nil {
        t.Error("FromHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHours(0)
    if err != nil {
        t.Errorf("FromHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DurationHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHours() with zero value = %v, want 0", converted)
    }
}
// Test FromMinutes function
func TestDurationFactory_FromMinutes(t *testing.T) {
    factory := units.DurationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMinutes(100)
    if err != nil {
        t.Errorf("FromMinutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DurationMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMinutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMinutes(math.NaN())
    if err == nil {
        t.Error("FromMinutes() with NaN value should return error")
    }

    _, err = factory.FromMinutes(math.Inf(1))
    if err == nil {
        t.Error("FromMinutes() with +Inf value should return error")
    }

    _, err = factory.FromMinutes(math.Inf(-1))
    if err == nil {
        t.Error("FromMinutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMinutes(0)
    if err != nil {
        t.Errorf("FromMinutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DurationMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMinutes() with zero value = %v, want 0", converted)
    }
}
// Test FromSeconds function
func TestDurationFactory_FromSeconds(t *testing.T) {
    factory := units.DurationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSeconds(100)
    if err != nil {
        t.Errorf("FromSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DurationSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSeconds(math.NaN())
    if err == nil {
        t.Error("FromSeconds() with NaN value should return error")
    }

    _, err = factory.FromSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromSeconds() with +Inf value should return error")
    }

    _, err = factory.FromSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSeconds(0)
    if err != nil {
        t.Errorf("FromSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DurationSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromJulianYears function
func TestDurationFactory_FromJulianYears(t *testing.T) {
    factory := units.DurationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJulianYears(100)
    if err != nil {
        t.Errorf("FromJulianYears() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DurationJulianYear)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJulianYears() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJulianYears(math.NaN())
    if err == nil {
        t.Error("FromJulianYears() with NaN value should return error")
    }

    _, err = factory.FromJulianYears(math.Inf(1))
    if err == nil {
        t.Error("FromJulianYears() with +Inf value should return error")
    }

    _, err = factory.FromJulianYears(math.Inf(-1))
    if err == nil {
        t.Error("FromJulianYears() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJulianYears(0)
    if err != nil {
        t.Errorf("FromJulianYears() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DurationJulianYear)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJulianYears() with zero value = %v, want 0", converted)
    }
}
// Test FromSols function
func TestDurationFactory_FromSols(t *testing.T) {
    factory := units.DurationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSols(100)
    if err != nil {
        t.Errorf("FromSols() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DurationSol)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSols() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSols(math.NaN())
    if err == nil {
        t.Error("FromSols() with NaN value should return error")
    }

    _, err = factory.FromSols(math.Inf(1))
    if err == nil {
        t.Error("FromSols() with +Inf value should return error")
    }

    _, err = factory.FromSols(math.Inf(-1))
    if err == nil {
        t.Error("FromSols() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSols(0)
    if err != nil {
        t.Errorf("FromSols() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DurationSol)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSols() with zero value = %v, want 0", converted)
    }
}
// Test FromNanoseconds function
func TestDurationFactory_FromNanoseconds(t *testing.T) {
    factory := units.DurationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanoseconds(100)
    if err != nil {
        t.Errorf("FromNanoseconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DurationNanosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanoseconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanoseconds(math.NaN())
    if err == nil {
        t.Error("FromNanoseconds() with NaN value should return error")
    }

    _, err = factory.FromNanoseconds(math.Inf(1))
    if err == nil {
        t.Error("FromNanoseconds() with +Inf value should return error")
    }

    _, err = factory.FromNanoseconds(math.Inf(-1))
    if err == nil {
        t.Error("FromNanoseconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanoseconds(0)
    if err != nil {
        t.Errorf("FromNanoseconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DurationNanosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanoseconds() with zero value = %v, want 0", converted)
    }
}
// Test FromMicroseconds function
func TestDurationFactory_FromMicroseconds(t *testing.T) {
    factory := units.DurationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicroseconds(100)
    if err != nil {
        t.Errorf("FromMicroseconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DurationMicrosecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicroseconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicroseconds(math.NaN())
    if err == nil {
        t.Error("FromMicroseconds() with NaN value should return error")
    }

    _, err = factory.FromMicroseconds(math.Inf(1))
    if err == nil {
        t.Error("FromMicroseconds() with +Inf value should return error")
    }

    _, err = factory.FromMicroseconds(math.Inf(-1))
    if err == nil {
        t.Error("FromMicroseconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicroseconds(0)
    if err != nil {
        t.Errorf("FromMicroseconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DurationMicrosecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicroseconds() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliseconds function
func TestDurationFactory_FromMilliseconds(t *testing.T) {
    factory := units.DurationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliseconds(100)
    if err != nil {
        t.Errorf("FromMilliseconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DurationMillisecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliseconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliseconds(math.NaN())
    if err == nil {
        t.Error("FromMilliseconds() with NaN value should return error")
    }

    _, err = factory.FromMilliseconds(math.Inf(1))
    if err == nil {
        t.Error("FromMilliseconds() with +Inf value should return error")
    }

    _, err = factory.FromMilliseconds(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliseconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliseconds(0)
    if err != nil {
        t.Errorf("FromMilliseconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DurationMillisecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliseconds() with zero value = %v, want 0", converted)
    }
}

func TestDurationToString(t *testing.T) {
	factory := units.DurationFactory{}
	a, err := factory.CreateDuration(45, units.DurationSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.DurationSecond, 2)
	expected := "45.00 " + units.GetDurationAbbreviation(units.DurationSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.DurationSecond, -1)
	expected = "45 " + units.GetDurationAbbreviation(units.DurationSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestDuration_EqualityAndComparison(t *testing.T) {
	factory := units.DurationFactory{}
	a1, _ := factory.CreateDuration(60, units.DurationSecond)
	a2, _ := factory.CreateDuration(60, units.DurationSecond)
	a3, _ := factory.CreateDuration(120, units.DurationSecond)

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

func TestDuration_Arithmetic(t *testing.T) {
	factory := units.DurationFactory{}
	a1, _ := factory.CreateDuration(30, units.DurationSecond)
	a2, _ := factory.CreateDuration(45, units.DurationSecond)

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


func TestGetDurationAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.DurationUnits
        want string
    }{
        {
            name: "Year365 abbreviation",
            unit: units.DurationYear365,
            want: "yr",
        },
        {
            name: "Month30 abbreviation",
            unit: units.DurationMonth30,
            want: "mo",
        },
        {
            name: "Week abbreviation",
            unit: units.DurationWeek,
            want: "wk",
        },
        {
            name: "Day abbreviation",
            unit: units.DurationDay,
            want: "d",
        },
        {
            name: "Hour abbreviation",
            unit: units.DurationHour,
            want: "h",
        },
        {
            name: "Minute abbreviation",
            unit: units.DurationMinute,
            want: "m",
        },
        {
            name: "Second abbreviation",
            unit: units.DurationSecond,
            want: "s",
        },
        {
            name: "JulianYear abbreviation",
            unit: units.DurationJulianYear,
            want: "jyr",
        },
        {
            name: "Sol abbreviation",
            unit: units.DurationSol,
            want: "sol",
        },
        {
            name: "Nanosecond abbreviation",
            unit: units.DurationNanosecond,
            want: "ns",
        },
        {
            name: "Microsecond abbreviation",
            unit: units.DurationMicrosecond,
            want: "μs",
        },
        {
            name: "Millisecond abbreviation",
            unit: units.DurationMillisecond,
            want: "ms",
        },
        {
            name: "invalid unit",
            unit: units.DurationUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetDurationAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetDurationAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestDuration_String(t *testing.T) {
    factory := units.DurationFactory{}
    
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
            unit, err := factory.CreateDuration(tt.value, units.DurationSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Duration.String() = %v, want %v", got, tt.want)
            }
        })
    }
}