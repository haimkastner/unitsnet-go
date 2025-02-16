// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestImpulseDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonSecond"}`
	
	factory := units.ImpulseDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ImpulseNewtonSecond {
		t.Errorf("expected unit %v, got %v", units.ImpulseNewtonSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestImpulseDto_ToJSON(t *testing.T) {
	dto := units.ImpulseDto{
		Value: 45,
		Unit:  units.ImpulseNewtonSecond,
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
	if result["unit"].(string) != string(units.ImpulseNewtonSecond) {
		t.Errorf("expected unit %s, got %v", units.ImpulseNewtonSecond, result["unit"])
	}
}

func TestNewImpulse_InvalidValue(t *testing.T) {
	factory := units.ImpulseFactory{}
	// NaN value should return an error.
	_, err := factory.CreateImpulse(math.NaN(), units.ImpulseNewtonSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateImpulse(math.Inf(1), units.ImpulseNewtonSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestImpulseConversions(t *testing.T) {
	factory := units.ImpulseFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateImpulse(180, units.ImpulseNewtonSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to KilogramMetersPerSecond.
		// No expected conversion value provided for KilogramMetersPerSecond, verifying result is not NaN.
		result := a.KilogramMetersPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramMetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to NewtonSeconds.
		// No expected conversion value provided for NewtonSeconds, verifying result is not NaN.
		result := a.NewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to PoundFeetPerSecond.
		// No expected conversion value provided for PoundFeetPerSecond, verifying result is not NaN.
		result := a.PoundFeetPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PoundForceSeconds.
		// No expected conversion value provided for PoundForceSeconds, verifying result is not NaN.
		result := a.PoundForceSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundForceSeconds returned NaN")
		}
	}
	{
		// Test conversion to SlugFeetPerSecond.
		// No expected conversion value provided for SlugFeetPerSecond, verifying result is not NaN.
		result := a.SlugFeetPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonSeconds.
		// No expected conversion value provided for NanonewtonSeconds, verifying result is not NaN.
		result := a.NanonewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanonewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonSeconds.
		// No expected conversion value provided for MicronewtonSeconds, verifying result is not NaN.
		result := a.MicronewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicronewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonSeconds.
		// No expected conversion value provided for MillinewtonSeconds, verifying result is not NaN.
		result := a.MillinewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillinewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonSeconds.
		// No expected conversion value provided for CentinewtonSeconds, verifying result is not NaN.
		result := a.CentinewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentinewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonSeconds.
		// No expected conversion value provided for DecinewtonSeconds, verifying result is not NaN.
		result := a.DecinewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecinewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonSeconds.
		// No expected conversion value provided for DecanewtonSeconds, verifying result is not NaN.
		result := a.DecanewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecanewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonSeconds.
		// No expected conversion value provided for KilonewtonSeconds, verifying result is not NaN.
		result := a.KilonewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonSeconds returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonSeconds.
		// No expected conversion value provided for MeganewtonSeconds, verifying result is not NaN.
		result := a.MeganewtonSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonSeconds returned NaN")
		}
	}
}

func TestImpulse_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ImpulseFactory{}
	a, err := factory.CreateImpulse(90, units.ImpulseNewtonSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ImpulseNewtonSecond {
		t.Errorf("expected default unit NewtonSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ImpulseKilogramMeterPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ImpulseDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ImpulseNewtonSecond {
		t.Errorf("expected unit NewtonSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestImpulseFactory_FromDto(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulseNewtonSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ImpulseDto{
        Value: math.NaN(),
        Unit:  units.ImpulseNewtonSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test KilogramMeterPerSecond conversion
    kilogram_meters_per_secondDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulseKilogramMeterPerSecond,
    }
    
    var kilogram_meters_per_secondResult *units.Impulse
    kilogram_meters_per_secondResult, err = factory.FromDto(kilogram_meters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramMeterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_meters_per_secondResult.Convert(units.ImpulseKilogramMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramMeterPerSecond = %v, want %v", converted, 100)
    }
    // Test NewtonSecond conversion
    newton_secondsDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulseNewtonSecond,
    }
    
    var newton_secondsResult *units.Impulse
    newton_secondsResult, err = factory.FromDto(newton_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_secondsResult.Convert(units.ImpulseNewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonSecond = %v, want %v", converted, 100)
    }
    // Test PoundFootPerSecond conversion
    pound_feet_per_secondDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulsePoundFootPerSecond,
    }
    
    var pound_feet_per_secondResult *units.Impulse
    pound_feet_per_secondResult, err = factory.FromDto(pound_feet_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PoundFootPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_feet_per_secondResult.Convert(units.ImpulsePoundFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundFootPerSecond = %v, want %v", converted, 100)
    }
    // Test PoundForceSecond conversion
    pound_force_secondsDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulsePoundForceSecond,
    }
    
    var pound_force_secondsResult *units.Impulse
    pound_force_secondsResult, err = factory.FromDto(pound_force_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForceSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_secondsResult.Convert(units.ImpulsePoundForceSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceSecond = %v, want %v", converted, 100)
    }
    // Test SlugFootPerSecond conversion
    slug_feet_per_secondDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulseSlugFootPerSecond,
    }
    
    var slug_feet_per_secondResult *units.Impulse
    slug_feet_per_secondResult, err = factory.FromDto(slug_feet_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with SlugFootPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slug_feet_per_secondResult.Convert(units.ImpulseSlugFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugFootPerSecond = %v, want %v", converted, 100)
    }
    // Test NanonewtonSecond conversion
    nanonewton_secondsDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulseNanonewtonSecond,
    }
    
    var nanonewton_secondsResult *units.Impulse
    nanonewton_secondsResult, err = factory.FromDto(nanonewton_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with NanonewtonSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewton_secondsResult.Convert(units.ImpulseNanonewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonSecond = %v, want %v", converted, 100)
    }
    // Test MicronewtonSecond conversion
    micronewton_secondsDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulseMicronewtonSecond,
    }
    
    var micronewton_secondsResult *units.Impulse
    micronewton_secondsResult, err = factory.FromDto(micronewton_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with MicronewtonSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewton_secondsResult.Convert(units.ImpulseMicronewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonSecond = %v, want %v", converted, 100)
    }
    // Test MillinewtonSecond conversion
    millinewton_secondsDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulseMillinewtonSecond,
    }
    
    var millinewton_secondsResult *units.Impulse
    millinewton_secondsResult, err = factory.FromDto(millinewton_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with MillinewtonSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewton_secondsResult.Convert(units.ImpulseMillinewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonSecond = %v, want %v", converted, 100)
    }
    // Test CentinewtonSecond conversion
    centinewton_secondsDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulseCentinewtonSecond,
    }
    
    var centinewton_secondsResult *units.Impulse
    centinewton_secondsResult, err = factory.FromDto(centinewton_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with CentinewtonSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewton_secondsResult.Convert(units.ImpulseCentinewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonSecond = %v, want %v", converted, 100)
    }
    // Test DecinewtonSecond conversion
    decinewton_secondsDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulseDecinewtonSecond,
    }
    
    var decinewton_secondsResult *units.Impulse
    decinewton_secondsResult, err = factory.FromDto(decinewton_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with DecinewtonSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewton_secondsResult.Convert(units.ImpulseDecinewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonSecond = %v, want %v", converted, 100)
    }
    // Test DecanewtonSecond conversion
    decanewton_secondsDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulseDecanewtonSecond,
    }
    
    var decanewton_secondsResult *units.Impulse
    decanewton_secondsResult, err = factory.FromDto(decanewton_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with DecanewtonSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewton_secondsResult.Convert(units.ImpulseDecanewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonSecond = %v, want %v", converted, 100)
    }
    // Test KilonewtonSecond conversion
    kilonewton_secondsDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulseKilonewtonSecond,
    }
    
    var kilonewton_secondsResult *units.Impulse
    kilonewton_secondsResult, err = factory.FromDto(kilonewton_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_secondsResult.Convert(units.ImpulseKilonewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonSecond = %v, want %v", converted, 100)
    }
    // Test MeganewtonSecond conversion
    meganewton_secondsDto := units.ImpulseDto{
        Value: 100,
        Unit:  units.ImpulseMeganewtonSecond,
    }
    
    var meganewton_secondsResult *units.Impulse
    meganewton_secondsResult, err = factory.FromDto(meganewton_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_secondsResult.Convert(units.ImpulseMeganewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonSecond = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ImpulseDto{
        Value: 0,
        Unit:  units.ImpulseNewtonSecond,
    }
    
    var zeroResult *units.Impulse
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestImpulseFactory_FromDtoJSON(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "NewtonSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "NewtonSecond"}`)
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
    nanJSON, _ := json.Marshal(units.ImpulseDto{
        Value: nanValue,
        Unit:  units.ImpulseNewtonSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with KilogramMeterPerSecond unit
    kilogram_meters_per_secondJSON := []byte(`{"value": 100, "unit": "KilogramMeterPerSecond"}`)
    kilogram_meters_per_secondResult, err := factory.FromDtoJSON(kilogram_meters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramMeterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_meters_per_secondResult.Convert(units.ImpulseKilogramMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramMeterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonSecond unit
    newton_secondsJSON := []byte(`{"value": 100, "unit": "NewtonSecond"}`)
    newton_secondsResult, err := factory.FromDtoJSON(newton_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_secondsResult.Convert(units.ImpulseNewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PoundFootPerSecond unit
    pound_feet_per_secondJSON := []byte(`{"value": 100, "unit": "PoundFootPerSecond"}`)
    pound_feet_per_secondResult, err := factory.FromDtoJSON(pound_feet_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundFootPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_feet_per_secondResult.Convert(units.ImpulsePoundFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundFootPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForceSecond unit
    pound_force_secondsJSON := []byte(`{"value": 100, "unit": "PoundForceSecond"}`)
    pound_force_secondsResult, err := factory.FromDtoJSON(pound_force_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForceSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_secondsResult.Convert(units.ImpulsePoundForceSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceSecond = %v, want %v", converted, 100)
    }
    // Test JSON with SlugFootPerSecond unit
    slug_feet_per_secondJSON := []byte(`{"value": 100, "unit": "SlugFootPerSecond"}`)
    slug_feet_per_secondResult, err := factory.FromDtoJSON(slug_feet_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SlugFootPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slug_feet_per_secondResult.Convert(units.ImpulseSlugFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugFootPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with NanonewtonSecond unit
    nanonewton_secondsJSON := []byte(`{"value": 100, "unit": "NanonewtonSecond"}`)
    nanonewton_secondsResult, err := factory.FromDtoJSON(nanonewton_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanonewtonSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewton_secondsResult.Convert(units.ImpulseNanonewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicronewtonSecond unit
    micronewton_secondsJSON := []byte(`{"value": 100, "unit": "MicronewtonSecond"}`)
    micronewton_secondsResult, err := factory.FromDtoJSON(micronewton_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicronewtonSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewton_secondsResult.Convert(units.ImpulseMicronewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MillinewtonSecond unit
    millinewton_secondsJSON := []byte(`{"value": 100, "unit": "MillinewtonSecond"}`)
    millinewton_secondsResult, err := factory.FromDtoJSON(millinewton_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillinewtonSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewton_secondsResult.Convert(units.ImpulseMillinewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonSecond = %v, want %v", converted, 100)
    }
    // Test JSON with CentinewtonSecond unit
    centinewton_secondsJSON := []byte(`{"value": 100, "unit": "CentinewtonSecond"}`)
    centinewton_secondsResult, err := factory.FromDtoJSON(centinewton_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentinewtonSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewton_secondsResult.Convert(units.ImpulseCentinewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DecinewtonSecond unit
    decinewton_secondsJSON := []byte(`{"value": 100, "unit": "DecinewtonSecond"}`)
    decinewton_secondsResult, err := factory.FromDtoJSON(decinewton_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecinewtonSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewton_secondsResult.Convert(units.ImpulseDecinewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DecanewtonSecond unit
    decanewton_secondsJSON := []byte(`{"value": 100, "unit": "DecanewtonSecond"}`)
    decanewton_secondsResult, err := factory.FromDtoJSON(decanewton_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecanewtonSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewton_secondsResult.Convert(units.ImpulseDecanewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonSecond unit
    kilonewton_secondsJSON := []byte(`{"value": 100, "unit": "KilonewtonSecond"}`)
    kilonewton_secondsResult, err := factory.FromDtoJSON(kilonewton_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_secondsResult.Convert(units.ImpulseKilonewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonSecond unit
    meganewton_secondsJSON := []byte(`{"value": 100, "unit": "MeganewtonSecond"}`)
    meganewton_secondsResult, err := factory.FromDtoJSON(meganewton_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_secondsResult.Convert(units.ImpulseMeganewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonSecond = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "NewtonSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromKilogramMetersPerSecond function
func TestImpulseFactory_FromKilogramMetersPerSecond(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramMetersPerSecond(100)
    if err != nil {
        t.Errorf("FromKilogramMetersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ImpulseKilogramMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramMetersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramMetersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilogramMetersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilogramMetersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramMetersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilogramMetersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramMetersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramMetersPerSecond(0)
    if err != nil {
        t.Errorf("FromKilogramMetersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ImpulseKilogramMeterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramMetersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonSeconds function
func TestImpulseFactory_FromNewtonSeconds(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonSeconds(100)
    if err != nil {
        t.Errorf("FromNewtonSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ImpulseNewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonSeconds(math.NaN())
    if err == nil {
        t.Error("FromNewtonSeconds() with NaN value should return error")
    }

    _, err = factory.FromNewtonSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonSeconds() with +Inf value should return error")
    }

    _, err = factory.FromNewtonSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonSeconds(0)
    if err != nil {
        t.Errorf("FromNewtonSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ImpulseNewtonSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundFeetPerSecond function
func TestImpulseFactory_FromPoundFeetPerSecond(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundFeetPerSecond(100)
    if err != nil {
        t.Errorf("FromPoundFeetPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ImpulsePoundFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundFeetPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundFeetPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPoundFeetPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPoundFeetPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPoundFeetPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPoundFeetPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundFeetPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundFeetPerSecond(0)
    if err != nil {
        t.Errorf("FromPoundFeetPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ImpulsePoundFootPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundFeetPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundForceSeconds function
func TestImpulseFactory_FromPoundForceSeconds(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundForceSeconds(100)
    if err != nil {
        t.Errorf("FromPoundForceSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ImpulsePoundForceSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundForceSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundForceSeconds(math.NaN())
    if err == nil {
        t.Error("FromPoundForceSeconds() with NaN value should return error")
    }

    _, err = factory.FromPoundForceSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromPoundForceSeconds() with +Inf value should return error")
    }

    _, err = factory.FromPoundForceSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundForceSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundForceSeconds(0)
    if err != nil {
        t.Errorf("FromPoundForceSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ImpulsePoundForceSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundForceSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromSlugFeetPerSecond function
func TestImpulseFactory_FromSlugFeetPerSecond(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSlugFeetPerSecond(100)
    if err != nil {
        t.Errorf("FromSlugFeetPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ImpulseSlugFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSlugFeetPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSlugFeetPerSecond(math.NaN())
    if err == nil {
        t.Error("FromSlugFeetPerSecond() with NaN value should return error")
    }

    _, err = factory.FromSlugFeetPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromSlugFeetPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromSlugFeetPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromSlugFeetPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSlugFeetPerSecond(0)
    if err != nil {
        t.Errorf("FromSlugFeetPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ImpulseSlugFootPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSlugFeetPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromNanonewtonSeconds function
func TestImpulseFactory_FromNanonewtonSeconds(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanonewtonSeconds(100)
    if err != nil {
        t.Errorf("FromNanonewtonSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ImpulseNanonewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanonewtonSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanonewtonSeconds(math.NaN())
    if err == nil {
        t.Error("FromNanonewtonSeconds() with NaN value should return error")
    }

    _, err = factory.FromNanonewtonSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromNanonewtonSeconds() with +Inf value should return error")
    }

    _, err = factory.FromNanonewtonSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromNanonewtonSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanonewtonSeconds(0)
    if err != nil {
        t.Errorf("FromNanonewtonSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ImpulseNanonewtonSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanonewtonSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromMicronewtonSeconds function
func TestImpulseFactory_FromMicronewtonSeconds(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicronewtonSeconds(100)
    if err != nil {
        t.Errorf("FromMicronewtonSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ImpulseMicronewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicronewtonSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicronewtonSeconds(math.NaN())
    if err == nil {
        t.Error("FromMicronewtonSeconds() with NaN value should return error")
    }

    _, err = factory.FromMicronewtonSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromMicronewtonSeconds() with +Inf value should return error")
    }

    _, err = factory.FromMicronewtonSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromMicronewtonSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicronewtonSeconds(0)
    if err != nil {
        t.Errorf("FromMicronewtonSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ImpulseMicronewtonSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicronewtonSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromMillinewtonSeconds function
func TestImpulseFactory_FromMillinewtonSeconds(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillinewtonSeconds(100)
    if err != nil {
        t.Errorf("FromMillinewtonSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ImpulseMillinewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillinewtonSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillinewtonSeconds(math.NaN())
    if err == nil {
        t.Error("FromMillinewtonSeconds() with NaN value should return error")
    }

    _, err = factory.FromMillinewtonSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromMillinewtonSeconds() with +Inf value should return error")
    }

    _, err = factory.FromMillinewtonSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromMillinewtonSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillinewtonSeconds(0)
    if err != nil {
        t.Errorf("FromMillinewtonSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ImpulseMillinewtonSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillinewtonSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromCentinewtonSeconds function
func TestImpulseFactory_FromCentinewtonSeconds(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentinewtonSeconds(100)
    if err != nil {
        t.Errorf("FromCentinewtonSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ImpulseCentinewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentinewtonSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentinewtonSeconds(math.NaN())
    if err == nil {
        t.Error("FromCentinewtonSeconds() with NaN value should return error")
    }

    _, err = factory.FromCentinewtonSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromCentinewtonSeconds() with +Inf value should return error")
    }

    _, err = factory.FromCentinewtonSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromCentinewtonSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentinewtonSeconds(0)
    if err != nil {
        t.Errorf("FromCentinewtonSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ImpulseCentinewtonSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentinewtonSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromDecinewtonSeconds function
func TestImpulseFactory_FromDecinewtonSeconds(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecinewtonSeconds(100)
    if err != nil {
        t.Errorf("FromDecinewtonSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ImpulseDecinewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecinewtonSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecinewtonSeconds(math.NaN())
    if err == nil {
        t.Error("FromDecinewtonSeconds() with NaN value should return error")
    }

    _, err = factory.FromDecinewtonSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromDecinewtonSeconds() with +Inf value should return error")
    }

    _, err = factory.FromDecinewtonSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromDecinewtonSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecinewtonSeconds(0)
    if err != nil {
        t.Errorf("FromDecinewtonSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ImpulseDecinewtonSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecinewtonSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromDecanewtonSeconds function
func TestImpulseFactory_FromDecanewtonSeconds(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecanewtonSeconds(100)
    if err != nil {
        t.Errorf("FromDecanewtonSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ImpulseDecanewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecanewtonSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecanewtonSeconds(math.NaN())
    if err == nil {
        t.Error("FromDecanewtonSeconds() with NaN value should return error")
    }

    _, err = factory.FromDecanewtonSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromDecanewtonSeconds() with +Inf value should return error")
    }

    _, err = factory.FromDecanewtonSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromDecanewtonSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecanewtonSeconds(0)
    if err != nil {
        t.Errorf("FromDecanewtonSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ImpulseDecanewtonSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecanewtonSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonSeconds function
func TestImpulseFactory_FromKilonewtonSeconds(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonSeconds(100)
    if err != nil {
        t.Errorf("FromKilonewtonSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ImpulseKilonewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonSeconds(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonSeconds() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonSeconds() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonSeconds(0)
    if err != nil {
        t.Errorf("FromKilonewtonSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ImpulseKilonewtonSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonSeconds function
func TestImpulseFactory_FromMeganewtonSeconds(t *testing.T) {
    factory := units.ImpulseFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonSeconds(100)
    if err != nil {
        t.Errorf("FromMeganewtonSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ImpulseMeganewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonSeconds(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonSeconds() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonSeconds() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonSeconds(0)
    if err != nil {
        t.Errorf("FromMeganewtonSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ImpulseMeganewtonSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonSeconds() with zero value = %v, want 0", converted)
    }
}

func TestImpulseToString(t *testing.T) {
	factory := units.ImpulseFactory{}
	a, err := factory.CreateImpulse(45, units.ImpulseNewtonSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ImpulseNewtonSecond, 2)
	expected := "45.00 " + units.GetImpulseAbbreviation(units.ImpulseNewtonSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ImpulseNewtonSecond, -1)
	expected = "45 " + units.GetImpulseAbbreviation(units.ImpulseNewtonSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestImpulse_EqualityAndComparison(t *testing.T) {
	factory := units.ImpulseFactory{}
	a1, _ := factory.CreateImpulse(60, units.ImpulseNewtonSecond)
	a2, _ := factory.CreateImpulse(60, units.ImpulseNewtonSecond)
	a3, _ := factory.CreateImpulse(120, units.ImpulseNewtonSecond)

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

func TestImpulse_Arithmetic(t *testing.T) {
	factory := units.ImpulseFactory{}
	a1, _ := factory.CreateImpulse(30, units.ImpulseNewtonSecond)
	a2, _ := factory.CreateImpulse(45, units.ImpulseNewtonSecond)

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