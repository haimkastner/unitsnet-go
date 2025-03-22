// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestForceChangeRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonPerSecond"}`
	
	factory := units.ForceChangeRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ForceChangeRateNewtonPerSecond {
		t.Errorf("expected unit %v, got %v", units.ForceChangeRateNewtonPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestForceChangeRateDto_ToJSON(t *testing.T) {
	dto := units.ForceChangeRateDto{
		Value: 45,
		Unit:  units.ForceChangeRateNewtonPerSecond,
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
	if result["unit"].(string) != string(units.ForceChangeRateNewtonPerSecond) {
		t.Errorf("expected unit %s, got %v", units.ForceChangeRateNewtonPerSecond, result["unit"])
	}
}

func TestNewForceChangeRate_InvalidValue(t *testing.T) {
	factory := units.ForceChangeRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreateForceChangeRate(math.NaN(), units.ForceChangeRateNewtonPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateForceChangeRate(math.Inf(1), units.ForceChangeRateNewtonPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestForceChangeRateConversions(t *testing.T) {
	factory := units.ForceChangeRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateForceChangeRate(180, units.ForceChangeRateNewtonPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonsPerMinute.
		// No expected conversion value provided for NewtonsPerMinute, verifying result is not NaN.
		result := a.NewtonsPerMinute()
		cacheResult := a.NewtonsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerSecond.
		// No expected conversion value provided for NewtonsPerSecond, verifying result is not NaN.
		result := a.NewtonsPerSecond()
		cacheResult := a.NewtonsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerMinute.
		// No expected conversion value provided for PoundsForcePerMinute, verifying result is not NaN.
		result := a.PoundsForcePerMinute()
		cacheResult := a.PoundsForcePerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsForcePerMinute returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerSecond.
		// No expected conversion value provided for PoundsForcePerSecond, verifying result is not NaN.
		result := a.PoundsForcePerSecond()
		cacheResult := a.PoundsForcePerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsForcePerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonsPerMinute.
		// No expected conversion value provided for DecanewtonsPerMinute, verifying result is not NaN.
		result := a.DecanewtonsPerMinute()
		cacheResult := a.DecanewtonsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecanewtonsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerMinute.
		// No expected conversion value provided for KilonewtonsPerMinute, verifying result is not NaN.
		result := a.KilonewtonsPerMinute()
		cacheResult := a.KilonewtonsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonsPerSecond.
		// No expected conversion value provided for NanonewtonsPerSecond, verifying result is not NaN.
		result := a.NanonewtonsPerSecond()
		cacheResult := a.NanonewtonsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanonewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonsPerSecond.
		// No expected conversion value provided for MicronewtonsPerSecond, verifying result is not NaN.
		result := a.MicronewtonsPerSecond()
		cacheResult := a.MicronewtonsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicronewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonsPerSecond.
		// No expected conversion value provided for MillinewtonsPerSecond, verifying result is not NaN.
		result := a.MillinewtonsPerSecond()
		cacheResult := a.MillinewtonsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillinewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonsPerSecond.
		// No expected conversion value provided for CentinewtonsPerSecond, verifying result is not NaN.
		result := a.CentinewtonsPerSecond()
		cacheResult := a.CentinewtonsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentinewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonsPerSecond.
		// No expected conversion value provided for DecinewtonsPerSecond, verifying result is not NaN.
		result := a.DecinewtonsPerSecond()
		cacheResult := a.DecinewtonsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecinewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonsPerSecond.
		// No expected conversion value provided for DecanewtonsPerSecond, verifying result is not NaN.
		result := a.DecanewtonsPerSecond()
		cacheResult := a.DecanewtonsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecanewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerSecond.
		// No expected conversion value provided for KilonewtonsPerSecond, verifying result is not NaN.
		result := a.KilonewtonsPerSecond()
		cacheResult := a.KilonewtonsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerMinute.
		// No expected conversion value provided for KilopoundsForcePerMinute, verifying result is not NaN.
		result := a.KilopoundsForcePerMinute()
		cacheResult := a.KilopoundsForcePerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsForcePerMinute returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerSecond.
		// No expected conversion value provided for KilopoundsForcePerSecond, verifying result is not NaN.
		result := a.KilopoundsForcePerSecond()
		cacheResult := a.KilopoundsForcePerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsForcePerSecond returned NaN")
		}
	}
}

func TestForceChangeRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ForceChangeRateFactory{}
	a, err := factory.CreateForceChangeRate(90, units.ForceChangeRateNewtonPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ForceChangeRateNewtonPerSecond {
		t.Errorf("expected default unit NewtonPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ForceChangeRateNewtonPerMinute
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ForceChangeRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ForceChangeRateNewtonPerSecond {
		t.Errorf("expected unit NewtonPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestForceChangeRateFactory_FromDto(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateNewtonPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ForceChangeRateDto{
        Value: math.NaN(),
        Unit:  units.ForceChangeRateNewtonPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test NewtonPerMinute conversion
    newtons_per_minuteDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateNewtonPerMinute,
    }
    
    var newtons_per_minuteResult *units.ForceChangeRate
    newtons_per_minuteResult, err = factory.FromDto(newtons_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_minuteResult.Convert(units.ForceChangeRateNewtonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerMinute = %v, want %v", converted, 100)
    }
    // Test NewtonPerSecond conversion
    newtons_per_secondDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateNewtonPerSecond,
    }
    
    var newtons_per_secondResult *units.ForceChangeRate
    newtons_per_secondResult, err = factory.FromDto(newtons_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_secondResult.Convert(units.ForceChangeRateNewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test PoundForcePerMinute conversion
    pounds_force_per_minuteDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRatePoundForcePerMinute,
    }
    
    var pounds_force_per_minuteResult *units.ForceChangeRate
    pounds_force_per_minuteResult, err = factory.FromDto(pounds_force_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForcePerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_minuteResult.Convert(units.ForceChangeRatePoundForcePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerMinute = %v, want %v", converted, 100)
    }
    // Test PoundForcePerSecond conversion
    pounds_force_per_secondDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRatePoundForcePerSecond,
    }
    
    var pounds_force_per_secondResult *units.ForceChangeRate
    pounds_force_per_secondResult, err = factory.FromDto(pounds_force_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForcePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_secondResult.Convert(units.ForceChangeRatePoundForcePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerSecond = %v, want %v", converted, 100)
    }
    // Test DecanewtonPerMinute conversion
    decanewtons_per_minuteDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateDecanewtonPerMinute,
    }
    
    var decanewtons_per_minuteResult *units.ForceChangeRate
    decanewtons_per_minuteResult, err = factory.FromDto(decanewtons_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with DecanewtonPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewtons_per_minuteResult.Convert(units.ForceChangeRateDecanewtonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonPerMinute = %v, want %v", converted, 100)
    }
    // Test KilonewtonPerMinute conversion
    kilonewtons_per_minuteDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateKilonewtonPerMinute,
    }
    
    var kilonewtons_per_minuteResult *units.ForceChangeRate
    kilonewtons_per_minuteResult, err = factory.FromDto(kilonewtons_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_minuteResult.Convert(units.ForceChangeRateKilonewtonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerMinute = %v, want %v", converted, 100)
    }
    // Test NanonewtonPerSecond conversion
    nanonewtons_per_secondDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateNanonewtonPerSecond,
    }
    
    var nanonewtons_per_secondResult *units.ForceChangeRate
    nanonewtons_per_secondResult, err = factory.FromDto(nanonewtons_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with NanonewtonPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewtons_per_secondResult.Convert(units.ForceChangeRateNanonewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test MicronewtonPerSecond conversion
    micronewtons_per_secondDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateMicronewtonPerSecond,
    }
    
    var micronewtons_per_secondResult *units.ForceChangeRate
    micronewtons_per_secondResult, err = factory.FromDto(micronewtons_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MicronewtonPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewtons_per_secondResult.Convert(units.ForceChangeRateMicronewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test MillinewtonPerSecond conversion
    millinewtons_per_secondDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateMillinewtonPerSecond,
    }
    
    var millinewtons_per_secondResult *units.ForceChangeRate
    millinewtons_per_secondResult, err = factory.FromDto(millinewtons_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MillinewtonPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewtons_per_secondResult.Convert(units.ForceChangeRateMillinewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test CentinewtonPerSecond conversion
    centinewtons_per_secondDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateCentinewtonPerSecond,
    }
    
    var centinewtons_per_secondResult *units.ForceChangeRate
    centinewtons_per_secondResult, err = factory.FromDto(centinewtons_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with CentinewtonPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewtons_per_secondResult.Convert(units.ForceChangeRateCentinewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test DecinewtonPerSecond conversion
    decinewtons_per_secondDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateDecinewtonPerSecond,
    }
    
    var decinewtons_per_secondResult *units.ForceChangeRate
    decinewtons_per_secondResult, err = factory.FromDto(decinewtons_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DecinewtonPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewtons_per_secondResult.Convert(units.ForceChangeRateDecinewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test DecanewtonPerSecond conversion
    decanewtons_per_secondDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateDecanewtonPerSecond,
    }
    
    var decanewtons_per_secondResult *units.ForceChangeRate
    decanewtons_per_secondResult, err = factory.FromDto(decanewtons_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DecanewtonPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewtons_per_secondResult.Convert(units.ForceChangeRateDecanewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test KilonewtonPerSecond conversion
    kilonewtons_per_secondDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateKilonewtonPerSecond,
    }
    
    var kilonewtons_per_secondResult *units.ForceChangeRate
    kilonewtons_per_secondResult, err = factory.FromDto(kilonewtons_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_secondResult.Convert(units.ForceChangeRateKilonewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test KilopoundForcePerMinute conversion
    kilopounds_force_per_minuteDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateKilopoundForcePerMinute,
    }
    
    var kilopounds_force_per_minuteResult *units.ForceChangeRate
    kilopounds_force_per_minuteResult, err = factory.FromDto(kilopounds_force_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForcePerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_minuteResult.Convert(units.ForceChangeRateKilopoundForcePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerMinute = %v, want %v", converted, 100)
    }
    // Test KilopoundForcePerSecond conversion
    kilopounds_force_per_secondDto := units.ForceChangeRateDto{
        Value: 100,
        Unit:  units.ForceChangeRateKilopoundForcePerSecond,
    }
    
    var kilopounds_force_per_secondResult *units.ForceChangeRate
    kilopounds_force_per_secondResult, err = factory.FromDto(kilopounds_force_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForcePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_secondResult.Convert(units.ForceChangeRateKilopoundForcePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerSecond = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ForceChangeRateDto{
        Value: 0,
        Unit:  units.ForceChangeRateNewtonPerSecond,
    }
    
    var zeroResult *units.ForceChangeRate
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestForceChangeRateFactory_FromDtoJSON(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "NewtonPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "NewtonPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.ForceChangeRateDto{
        Value: nanValue,
        Unit:  units.ForceChangeRateNewtonPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with NewtonPerMinute unit
    newtons_per_minuteJSON := []byte(`{"value": 100, "unit": "NewtonPerMinute"}`)
    newtons_per_minuteResult, err := factory.FromDtoJSON(newtons_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_minuteResult.Convert(units.ForceChangeRateNewtonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonPerSecond unit
    newtons_per_secondJSON := []byte(`{"value": 100, "unit": "NewtonPerSecond"}`)
    newtons_per_secondResult, err := factory.FromDtoJSON(newtons_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_secondResult.Convert(units.ForceChangeRateNewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForcePerMinute unit
    pounds_force_per_minuteJSON := []byte(`{"value": 100, "unit": "PoundForcePerMinute"}`)
    pounds_force_per_minuteResult, err := factory.FromDtoJSON(pounds_force_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForcePerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_minuteResult.Convert(units.ForceChangeRatePoundForcePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForcePerSecond unit
    pounds_force_per_secondJSON := []byte(`{"value": 100, "unit": "PoundForcePerSecond"}`)
    pounds_force_per_secondResult, err := factory.FromDtoJSON(pounds_force_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForcePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_secondResult.Convert(units.ForceChangeRatePoundForcePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DecanewtonPerMinute unit
    decanewtons_per_minuteJSON := []byte(`{"value": 100, "unit": "DecanewtonPerMinute"}`)
    decanewtons_per_minuteResult, err := factory.FromDtoJSON(decanewtons_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecanewtonPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewtons_per_minuteResult.Convert(units.ForceChangeRateDecanewtonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonPerMinute unit
    kilonewtons_per_minuteJSON := []byte(`{"value": 100, "unit": "KilonewtonPerMinute"}`)
    kilonewtons_per_minuteResult, err := factory.FromDtoJSON(kilonewtons_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_minuteResult.Convert(units.ForceChangeRateKilonewtonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with NanonewtonPerSecond unit
    nanonewtons_per_secondJSON := []byte(`{"value": 100, "unit": "NanonewtonPerSecond"}`)
    nanonewtons_per_secondResult, err := factory.FromDtoJSON(nanonewtons_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanonewtonPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewtons_per_secondResult.Convert(units.ForceChangeRateNanonewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicronewtonPerSecond unit
    micronewtons_per_secondJSON := []byte(`{"value": 100, "unit": "MicronewtonPerSecond"}`)
    micronewtons_per_secondResult, err := factory.FromDtoJSON(micronewtons_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicronewtonPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewtons_per_secondResult.Convert(units.ForceChangeRateMicronewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MillinewtonPerSecond unit
    millinewtons_per_secondJSON := []byte(`{"value": 100, "unit": "MillinewtonPerSecond"}`)
    millinewtons_per_secondResult, err := factory.FromDtoJSON(millinewtons_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillinewtonPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewtons_per_secondResult.Convert(units.ForceChangeRateMillinewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with CentinewtonPerSecond unit
    centinewtons_per_secondJSON := []byte(`{"value": 100, "unit": "CentinewtonPerSecond"}`)
    centinewtons_per_secondResult, err := factory.FromDtoJSON(centinewtons_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentinewtonPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewtons_per_secondResult.Convert(units.ForceChangeRateCentinewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DecinewtonPerSecond unit
    decinewtons_per_secondJSON := []byte(`{"value": 100, "unit": "DecinewtonPerSecond"}`)
    decinewtons_per_secondResult, err := factory.FromDtoJSON(decinewtons_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecinewtonPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewtons_per_secondResult.Convert(units.ForceChangeRateDecinewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DecanewtonPerSecond unit
    decanewtons_per_secondJSON := []byte(`{"value": 100, "unit": "DecanewtonPerSecond"}`)
    decanewtons_per_secondResult, err := factory.FromDtoJSON(decanewtons_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecanewtonPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewtons_per_secondResult.Convert(units.ForceChangeRateDecanewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonPerSecond unit
    kilonewtons_per_secondJSON := []byte(`{"value": 100, "unit": "KilonewtonPerSecond"}`)
    kilonewtons_per_secondResult, err := factory.FromDtoJSON(kilonewtons_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_secondResult.Convert(units.ForceChangeRateKilonewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForcePerMinute unit
    kilopounds_force_per_minuteJSON := []byte(`{"value": 100, "unit": "KilopoundForcePerMinute"}`)
    kilopounds_force_per_minuteResult, err := factory.FromDtoJSON(kilopounds_force_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForcePerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_minuteResult.Convert(units.ForceChangeRateKilopoundForcePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForcePerSecond unit
    kilopounds_force_per_secondJSON := []byte(`{"value": 100, "unit": "KilopoundForcePerSecond"}`)
    kilopounds_force_per_secondResult, err := factory.FromDtoJSON(kilopounds_force_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForcePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_secondResult.Convert(units.ForceChangeRateKilopoundForcePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerSecond = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "NewtonPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromNewtonsPerMinute function
func TestForceChangeRateFactory_FromNewtonsPerMinute(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonsPerMinute(100)
    if err != nil {
        t.Errorf("FromNewtonsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRateNewtonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromNewtonsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromNewtonsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromNewtonsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonsPerMinute(0)
    if err != nil {
        t.Errorf("FromNewtonsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRateNewtonPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonsPerSecond function
func TestForceChangeRateFactory_FromNewtonsPerSecond(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonsPerSecond(100)
    if err != nil {
        t.Errorf("FromNewtonsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRateNewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromNewtonsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromNewtonsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromNewtonsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonsPerSecond(0)
    if err != nil {
        t.Errorf("FromNewtonsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRateNewtonPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForcePerMinute function
func TestForceChangeRateFactory_FromPoundsForcePerMinute(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForcePerMinute(100)
    if err != nil {
        t.Errorf("FromPoundsForcePerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRatePoundForcePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForcePerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForcePerMinute(math.NaN())
    if err == nil {
        t.Error("FromPoundsForcePerMinute() with NaN value should return error")
    }

    _, err = factory.FromPoundsForcePerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForcePerMinute() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForcePerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForcePerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForcePerMinute(0)
    if err != nil {
        t.Errorf("FromPoundsForcePerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRatePoundForcePerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForcePerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForcePerSecond function
func TestForceChangeRateFactory_FromPoundsForcePerSecond(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForcePerSecond(100)
    if err != nil {
        t.Errorf("FromPoundsForcePerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRatePoundForcePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForcePerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForcePerSecond(math.NaN())
    if err == nil {
        t.Error("FromPoundsForcePerSecond() with NaN value should return error")
    }

    _, err = factory.FromPoundsForcePerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForcePerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForcePerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForcePerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForcePerSecond(0)
    if err != nil {
        t.Errorf("FromPoundsForcePerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRatePoundForcePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForcePerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDecanewtonsPerMinute function
func TestForceChangeRateFactory_FromDecanewtonsPerMinute(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecanewtonsPerMinute(100)
    if err != nil {
        t.Errorf("FromDecanewtonsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRateDecanewtonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecanewtonsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecanewtonsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromDecanewtonsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromDecanewtonsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromDecanewtonsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromDecanewtonsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromDecanewtonsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecanewtonsPerMinute(0)
    if err != nil {
        t.Errorf("FromDecanewtonsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRateDecanewtonPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecanewtonsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonsPerMinute function
func TestForceChangeRateFactory_FromKilonewtonsPerMinute(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonsPerMinute(100)
    if err != nil {
        t.Errorf("FromKilonewtonsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRateKilonewtonPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonsPerMinute(0)
    if err != nil {
        t.Errorf("FromKilonewtonsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRateKilonewtonPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromNanonewtonsPerSecond function
func TestForceChangeRateFactory_FromNanonewtonsPerSecond(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanonewtonsPerSecond(100)
    if err != nil {
        t.Errorf("FromNanonewtonsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRateNanonewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanonewtonsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanonewtonsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromNanonewtonsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromNanonewtonsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromNanonewtonsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromNanonewtonsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromNanonewtonsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanonewtonsPerSecond(0)
    if err != nil {
        t.Errorf("FromNanonewtonsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRateNanonewtonPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanonewtonsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMicronewtonsPerSecond function
func TestForceChangeRateFactory_FromMicronewtonsPerSecond(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicronewtonsPerSecond(100)
    if err != nil {
        t.Errorf("FromMicronewtonsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRateMicronewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicronewtonsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicronewtonsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMicronewtonsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMicronewtonsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMicronewtonsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMicronewtonsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMicronewtonsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicronewtonsPerSecond(0)
    if err != nil {
        t.Errorf("FromMicronewtonsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRateMicronewtonPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicronewtonsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMillinewtonsPerSecond function
func TestForceChangeRateFactory_FromMillinewtonsPerSecond(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillinewtonsPerSecond(100)
    if err != nil {
        t.Errorf("FromMillinewtonsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRateMillinewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillinewtonsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillinewtonsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMillinewtonsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMillinewtonsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMillinewtonsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMillinewtonsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMillinewtonsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillinewtonsPerSecond(0)
    if err != nil {
        t.Errorf("FromMillinewtonsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRateMillinewtonPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillinewtonsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromCentinewtonsPerSecond function
func TestForceChangeRateFactory_FromCentinewtonsPerSecond(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentinewtonsPerSecond(100)
    if err != nil {
        t.Errorf("FromCentinewtonsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRateCentinewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentinewtonsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentinewtonsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromCentinewtonsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromCentinewtonsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromCentinewtonsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromCentinewtonsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromCentinewtonsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentinewtonsPerSecond(0)
    if err != nil {
        t.Errorf("FromCentinewtonsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRateCentinewtonPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentinewtonsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDecinewtonsPerSecond function
func TestForceChangeRateFactory_FromDecinewtonsPerSecond(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecinewtonsPerSecond(100)
    if err != nil {
        t.Errorf("FromDecinewtonsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRateDecinewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecinewtonsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecinewtonsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDecinewtonsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDecinewtonsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDecinewtonsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDecinewtonsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDecinewtonsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecinewtonsPerSecond(0)
    if err != nil {
        t.Errorf("FromDecinewtonsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRateDecinewtonPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecinewtonsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDecanewtonsPerSecond function
func TestForceChangeRateFactory_FromDecanewtonsPerSecond(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecanewtonsPerSecond(100)
    if err != nil {
        t.Errorf("FromDecanewtonsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRateDecanewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecanewtonsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecanewtonsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDecanewtonsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDecanewtonsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDecanewtonsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDecanewtonsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDecanewtonsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecanewtonsPerSecond(0)
    if err != nil {
        t.Errorf("FromDecanewtonsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRateDecanewtonPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecanewtonsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonsPerSecond function
func TestForceChangeRateFactory_FromKilonewtonsPerSecond(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonsPerSecond(100)
    if err != nil {
        t.Errorf("FromKilonewtonsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRateKilonewtonPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonsPerSecond(0)
    if err != nil {
        t.Errorf("FromKilonewtonsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRateKilonewtonPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsForcePerMinute function
func TestForceChangeRateFactory_FromKilopoundsForcePerMinute(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsForcePerMinute(100)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRateKilopoundForcePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsForcePerMinute(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsForcePerMinute() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsForcePerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsForcePerMinute() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsForcePerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsForcePerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsForcePerMinute(0)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRateKilopoundForcePerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsForcePerSecond function
func TestForceChangeRateFactory_FromKilopoundsForcePerSecond(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsForcePerSecond(100)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceChangeRateKilopoundForcePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsForcePerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsForcePerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsForcePerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsForcePerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsForcePerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsForcePerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsForcePerSecond(0)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceChangeRateKilopoundForcePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerSecond() with zero value = %v, want 0", converted)
    }
}

func TestForceChangeRateToString(t *testing.T) {
	factory := units.ForceChangeRateFactory{}
	a, err := factory.CreateForceChangeRate(45, units.ForceChangeRateNewtonPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ForceChangeRateNewtonPerSecond, 2)
	expected := "45.00 " + units.GetForceChangeRateAbbreviation(units.ForceChangeRateNewtonPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ForceChangeRateNewtonPerSecond, -1)
	expected = "45 " + units.GetForceChangeRateAbbreviation(units.ForceChangeRateNewtonPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestForceChangeRate_EqualityAndComparison(t *testing.T) {
	factory := units.ForceChangeRateFactory{}
	a1, _ := factory.CreateForceChangeRate(60, units.ForceChangeRateNewtonPerSecond)
	a2, _ := factory.CreateForceChangeRate(60, units.ForceChangeRateNewtonPerSecond)
	a3, _ := factory.CreateForceChangeRate(120, units.ForceChangeRateNewtonPerSecond)

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

func TestForceChangeRate_Arithmetic(t *testing.T) {
	factory := units.ForceChangeRateFactory{}
	a1, _ := factory.CreateForceChangeRate(30, units.ForceChangeRateNewtonPerSecond)
	a2, _ := factory.CreateForceChangeRate(45, units.ForceChangeRateNewtonPerSecond)

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


func TestGetForceChangeRateAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ForceChangeRateUnits
        want string
    }{
        {
            name: "NewtonPerMinute abbreviation",
            unit: units.ForceChangeRateNewtonPerMinute,
            want: "N/min",
        },
        {
            name: "NewtonPerSecond abbreviation",
            unit: units.ForceChangeRateNewtonPerSecond,
            want: "N/s",
        },
        {
            name: "PoundForcePerMinute abbreviation",
            unit: units.ForceChangeRatePoundForcePerMinute,
            want: "lbf/min",
        },
        {
            name: "PoundForcePerSecond abbreviation",
            unit: units.ForceChangeRatePoundForcePerSecond,
            want: "lbf/s",
        },
        {
            name: "DecanewtonPerMinute abbreviation",
            unit: units.ForceChangeRateDecanewtonPerMinute,
            want: "daN/min",
        },
        {
            name: "KilonewtonPerMinute abbreviation",
            unit: units.ForceChangeRateKilonewtonPerMinute,
            want: "kN/min",
        },
        {
            name: "NanonewtonPerSecond abbreviation",
            unit: units.ForceChangeRateNanonewtonPerSecond,
            want: "nN/s",
        },
        {
            name: "MicronewtonPerSecond abbreviation",
            unit: units.ForceChangeRateMicronewtonPerSecond,
            want: "μN/s",
        },
        {
            name: "MillinewtonPerSecond abbreviation",
            unit: units.ForceChangeRateMillinewtonPerSecond,
            want: "mN/s",
        },
        {
            name: "CentinewtonPerSecond abbreviation",
            unit: units.ForceChangeRateCentinewtonPerSecond,
            want: "cN/s",
        },
        {
            name: "DecinewtonPerSecond abbreviation",
            unit: units.ForceChangeRateDecinewtonPerSecond,
            want: "dN/s",
        },
        {
            name: "DecanewtonPerSecond abbreviation",
            unit: units.ForceChangeRateDecanewtonPerSecond,
            want: "daN/s",
        },
        {
            name: "KilonewtonPerSecond abbreviation",
            unit: units.ForceChangeRateKilonewtonPerSecond,
            want: "kN/s",
        },
        {
            name: "KilopoundForcePerMinute abbreviation",
            unit: units.ForceChangeRateKilopoundForcePerMinute,
            want: "klbf/min",
        },
        {
            name: "KilopoundForcePerSecond abbreviation",
            unit: units.ForceChangeRateKilopoundForcePerSecond,
            want: "klbf/s",
        },
        {
            name: "invalid unit",
            unit: units.ForceChangeRateUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetForceChangeRateAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetForceChangeRateAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestForceChangeRate_String(t *testing.T) {
    factory := units.ForceChangeRateFactory{}
    
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
            unit, err := factory.CreateForceChangeRate(tt.value, units.ForceChangeRateNewtonPerSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ForceChangeRate.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestForceChangeRate_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.ForceChangeRateFactory{}

	_, err := uf.CreateForceChangeRate(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}