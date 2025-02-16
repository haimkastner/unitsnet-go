// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMolarFlowDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MolePerSecond"}`
	
	factory := units.MolarFlowDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MolarFlowMolePerSecond {
		t.Errorf("expected unit %v, got %v", units.MolarFlowMolePerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MolePerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMolarFlowDto_ToJSON(t *testing.T) {
	dto := units.MolarFlowDto{
		Value: 45,
		Unit:  units.MolarFlowMolePerSecond,
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
	if result["unit"].(string) != string(units.MolarFlowMolePerSecond) {
		t.Errorf("expected unit %s, got %v", units.MolarFlowMolePerSecond, result["unit"])
	}
}

func TestNewMolarFlow_InvalidValue(t *testing.T) {
	factory := units.MolarFlowFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMolarFlow(math.NaN(), units.MolarFlowMolePerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMolarFlow(math.Inf(1), units.MolarFlowMolePerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMolarFlowConversions(t *testing.T) {
	factory := units.MolarFlowFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMolarFlow(180, units.MolarFlowMolePerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MolesPerSecond.
		// No expected conversion value provided for MolesPerSecond, verifying result is not NaN.
		result := a.MolesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to MolesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MolesPerMinute.
		// No expected conversion value provided for MolesPerMinute, verifying result is not NaN.
		result := a.MolesPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to MolesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MolesPerHour.
		// No expected conversion value provided for MolesPerHour, verifying result is not NaN.
		result := a.MolesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MolesPerHour returned NaN")
		}
	}
	{
		// Test conversion to PoundMolesPerSecond.
		// No expected conversion value provided for PoundMolesPerSecond, verifying result is not NaN.
		result := a.PoundMolesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundMolesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PoundMolesPerMinute.
		// No expected conversion value provided for PoundMolesPerMinute, verifying result is not NaN.
		result := a.PoundMolesPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundMolesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to PoundMolesPerHour.
		// No expected conversion value provided for PoundMolesPerHour, verifying result is not NaN.
		result := a.PoundMolesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundMolesPerHour returned NaN")
		}
	}
	{
		// Test conversion to KilomolesPerSecond.
		// No expected conversion value provided for KilomolesPerSecond, verifying result is not NaN.
		result := a.KilomolesPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilomolesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilomolesPerMinute.
		// No expected conversion value provided for KilomolesPerMinute, verifying result is not NaN.
		result := a.KilomolesPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilomolesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to KilomolesPerHour.
		// No expected conversion value provided for KilomolesPerHour, verifying result is not NaN.
		result := a.KilomolesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilomolesPerHour returned NaN")
		}
	}
}

func TestMolarFlow_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MolarFlowFactory{}
	a, err := factory.CreateMolarFlow(90, units.MolarFlowMolePerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MolarFlowMolePerSecond {
		t.Errorf("expected default unit MolePerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MolarFlowMolePerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MolarFlowDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MolarFlowMolePerSecond {
		t.Errorf("expected unit MolePerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMolarFlowFactory_FromDto(t *testing.T) {
    factory := units.MolarFlowFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MolarFlowDto{
        Value: 100,
        Unit:  units.MolarFlowMolePerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MolarFlowDto{
        Value: math.NaN(),
        Unit:  units.MolarFlowMolePerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test MolePerSecond conversion
    moles_per_secondDto := units.MolarFlowDto{
        Value: 100,
        Unit:  units.MolarFlowMolePerSecond,
    }
    
    var moles_per_secondResult *units.MolarFlow
    moles_per_secondResult, err = factory.FromDto(moles_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MolePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_secondResult.Convert(units.MolarFlowMolePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerSecond = %v, want %v", converted, 100)
    }
    // Test MolePerMinute conversion
    moles_per_minuteDto := units.MolarFlowDto{
        Value: 100,
        Unit:  units.MolarFlowMolePerMinute,
    }
    
    var moles_per_minuteResult *units.MolarFlow
    moles_per_minuteResult, err = factory.FromDto(moles_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with MolePerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_minuteResult.Convert(units.MolarFlowMolePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerMinute = %v, want %v", converted, 100)
    }
    // Test MolePerHour conversion
    moles_per_hourDto := units.MolarFlowDto{
        Value: 100,
        Unit:  units.MolarFlowMolePerHour,
    }
    
    var moles_per_hourResult *units.MolarFlow
    moles_per_hourResult, err = factory.FromDto(moles_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MolePerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_hourResult.Convert(units.MolarFlowMolePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerHour = %v, want %v", converted, 100)
    }
    // Test PoundMolePerSecond conversion
    pound_moles_per_secondDto := units.MolarFlowDto{
        Value: 100,
        Unit:  units.MolarFlowPoundMolePerSecond,
    }
    
    var pound_moles_per_secondResult *units.MolarFlow
    pound_moles_per_secondResult, err = factory.FromDto(pound_moles_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PoundMolePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_moles_per_secondResult.Convert(units.MolarFlowPoundMolePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundMolePerSecond = %v, want %v", converted, 100)
    }
    // Test PoundMolePerMinute conversion
    pound_moles_per_minuteDto := units.MolarFlowDto{
        Value: 100,
        Unit:  units.MolarFlowPoundMolePerMinute,
    }
    
    var pound_moles_per_minuteResult *units.MolarFlow
    pound_moles_per_minuteResult, err = factory.FromDto(pound_moles_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with PoundMolePerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_moles_per_minuteResult.Convert(units.MolarFlowPoundMolePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundMolePerMinute = %v, want %v", converted, 100)
    }
    // Test PoundMolePerHour conversion
    pound_moles_per_hourDto := units.MolarFlowDto{
        Value: 100,
        Unit:  units.MolarFlowPoundMolePerHour,
    }
    
    var pound_moles_per_hourResult *units.MolarFlow
    pound_moles_per_hourResult, err = factory.FromDto(pound_moles_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with PoundMolePerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_moles_per_hourResult.Convert(units.MolarFlowPoundMolePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundMolePerHour = %v, want %v", converted, 100)
    }
    // Test KilomolePerSecond conversion
    kilomoles_per_secondDto := units.MolarFlowDto{
        Value: 100,
        Unit:  units.MolarFlowKilomolePerSecond,
    }
    
    var kilomoles_per_secondResult *units.MolarFlow
    kilomoles_per_secondResult, err = factory.FromDto(kilomoles_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilomolePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilomoles_per_secondResult.Convert(units.MolarFlowKilomolePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilomolePerSecond = %v, want %v", converted, 100)
    }
    // Test KilomolePerMinute conversion
    kilomoles_per_minuteDto := units.MolarFlowDto{
        Value: 100,
        Unit:  units.MolarFlowKilomolePerMinute,
    }
    
    var kilomoles_per_minuteResult *units.MolarFlow
    kilomoles_per_minuteResult, err = factory.FromDto(kilomoles_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with KilomolePerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilomoles_per_minuteResult.Convert(units.MolarFlowKilomolePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilomolePerMinute = %v, want %v", converted, 100)
    }
    // Test KilomolePerHour conversion
    kilomoles_per_hourDto := units.MolarFlowDto{
        Value: 100,
        Unit:  units.MolarFlowKilomolePerHour,
    }
    
    var kilomoles_per_hourResult *units.MolarFlow
    kilomoles_per_hourResult, err = factory.FromDto(kilomoles_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with KilomolePerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilomoles_per_hourResult.Convert(units.MolarFlowKilomolePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilomolePerHour = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MolarFlowDto{
        Value: 0,
        Unit:  units.MolarFlowMolePerSecond,
    }
    
    var zeroResult *units.MolarFlow
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMolarFlowFactory_FromDtoJSON(t *testing.T) {
    factory := units.MolarFlowFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "MolePerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "MolePerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.MolarFlowDto{
        Value: nanValue,
        Unit:  units.MolarFlowMolePerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with MolePerSecond unit
    moles_per_secondJSON := []byte(`{"value": 100, "unit": "MolePerSecond"}`)
    moles_per_secondResult, err := factory.FromDtoJSON(moles_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MolePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_secondResult.Convert(units.MolarFlowMolePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MolePerMinute unit
    moles_per_minuteJSON := []byte(`{"value": 100, "unit": "MolePerMinute"}`)
    moles_per_minuteResult, err := factory.FromDtoJSON(moles_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MolePerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_minuteResult.Convert(units.MolarFlowMolePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MolePerHour unit
    moles_per_hourJSON := []byte(`{"value": 100, "unit": "MolePerHour"}`)
    moles_per_hourResult, err := factory.FromDtoJSON(moles_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MolePerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_hourResult.Convert(units.MolarFlowMolePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerHour = %v, want %v", converted, 100)
    }
    // Test JSON with PoundMolePerSecond unit
    pound_moles_per_secondJSON := []byte(`{"value": 100, "unit": "PoundMolePerSecond"}`)
    pound_moles_per_secondResult, err := factory.FromDtoJSON(pound_moles_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundMolePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_moles_per_secondResult.Convert(units.MolarFlowPoundMolePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundMolePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PoundMolePerMinute unit
    pound_moles_per_minuteJSON := []byte(`{"value": 100, "unit": "PoundMolePerMinute"}`)
    pound_moles_per_minuteResult, err := factory.FromDtoJSON(pound_moles_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundMolePerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_moles_per_minuteResult.Convert(units.MolarFlowPoundMolePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundMolePerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with PoundMolePerHour unit
    pound_moles_per_hourJSON := []byte(`{"value": 100, "unit": "PoundMolePerHour"}`)
    pound_moles_per_hourResult, err := factory.FromDtoJSON(pound_moles_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundMolePerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_moles_per_hourResult.Convert(units.MolarFlowPoundMolePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundMolePerHour = %v, want %v", converted, 100)
    }
    // Test JSON with KilomolePerSecond unit
    kilomoles_per_secondJSON := []byte(`{"value": 100, "unit": "KilomolePerSecond"}`)
    kilomoles_per_secondResult, err := factory.FromDtoJSON(kilomoles_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilomolePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilomoles_per_secondResult.Convert(units.MolarFlowKilomolePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilomolePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilomolePerMinute unit
    kilomoles_per_minuteJSON := []byte(`{"value": 100, "unit": "KilomolePerMinute"}`)
    kilomoles_per_minuteResult, err := factory.FromDtoJSON(kilomoles_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilomolePerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilomoles_per_minuteResult.Convert(units.MolarFlowKilomolePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilomolePerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with KilomolePerHour unit
    kilomoles_per_hourJSON := []byte(`{"value": 100, "unit": "KilomolePerHour"}`)
    kilomoles_per_hourResult, err := factory.FromDtoJSON(kilomoles_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilomolePerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilomoles_per_hourResult.Convert(units.MolarFlowKilomolePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilomolePerHour = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "MolePerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromMolesPerSecond function
func TestMolarFlowFactory_FromMolesPerSecond(t *testing.T) {
    factory := units.MolarFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMolesPerSecond(100)
    if err != nil {
        t.Errorf("FromMolesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarFlowMolePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMolesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMolesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMolesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMolesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMolesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMolesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMolesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMolesPerSecond(0)
    if err != nil {
        t.Errorf("FromMolesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarFlowMolePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMolesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMolesPerMinute function
func TestMolarFlowFactory_FromMolesPerMinute(t *testing.T) {
    factory := units.MolarFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMolesPerMinute(100)
    if err != nil {
        t.Errorf("FromMolesPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarFlowMolePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMolesPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMolesPerMinute(math.NaN())
    if err == nil {
        t.Error("FromMolesPerMinute() with NaN value should return error")
    }

    _, err = factory.FromMolesPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromMolesPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromMolesPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromMolesPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMolesPerMinute(0)
    if err != nil {
        t.Errorf("FromMolesPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarFlowMolePerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMolesPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromMolesPerHour function
func TestMolarFlowFactory_FromMolesPerHour(t *testing.T) {
    factory := units.MolarFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMolesPerHour(100)
    if err != nil {
        t.Errorf("FromMolesPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarFlowMolePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMolesPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMolesPerHour(math.NaN())
    if err == nil {
        t.Error("FromMolesPerHour() with NaN value should return error")
    }

    _, err = factory.FromMolesPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMolesPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMolesPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMolesPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMolesPerHour(0)
    if err != nil {
        t.Errorf("FromMolesPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarFlowMolePerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMolesPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundMolesPerSecond function
func TestMolarFlowFactory_FromPoundMolesPerSecond(t *testing.T) {
    factory := units.MolarFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundMolesPerSecond(100)
    if err != nil {
        t.Errorf("FromPoundMolesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarFlowPoundMolePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundMolesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundMolesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPoundMolesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPoundMolesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPoundMolesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPoundMolesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundMolesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundMolesPerSecond(0)
    if err != nil {
        t.Errorf("FromPoundMolesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarFlowPoundMolePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundMolesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundMolesPerMinute function
func TestMolarFlowFactory_FromPoundMolesPerMinute(t *testing.T) {
    factory := units.MolarFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundMolesPerMinute(100)
    if err != nil {
        t.Errorf("FromPoundMolesPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarFlowPoundMolePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundMolesPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundMolesPerMinute(math.NaN())
    if err == nil {
        t.Error("FromPoundMolesPerMinute() with NaN value should return error")
    }

    _, err = factory.FromPoundMolesPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromPoundMolesPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromPoundMolesPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundMolesPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundMolesPerMinute(0)
    if err != nil {
        t.Errorf("FromPoundMolesPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarFlowPoundMolePerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundMolesPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundMolesPerHour function
func TestMolarFlowFactory_FromPoundMolesPerHour(t *testing.T) {
    factory := units.MolarFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundMolesPerHour(100)
    if err != nil {
        t.Errorf("FromPoundMolesPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarFlowPoundMolePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundMolesPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundMolesPerHour(math.NaN())
    if err == nil {
        t.Error("FromPoundMolesPerHour() with NaN value should return error")
    }

    _, err = factory.FromPoundMolesPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromPoundMolesPerHour() with +Inf value should return error")
    }

    _, err = factory.FromPoundMolesPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundMolesPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundMolesPerHour(0)
    if err != nil {
        t.Errorf("FromPoundMolesPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarFlowPoundMolePerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundMolesPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromKilomolesPerSecond function
func TestMolarFlowFactory_FromKilomolesPerSecond(t *testing.T) {
    factory := units.MolarFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilomolesPerSecond(100)
    if err != nil {
        t.Errorf("FromKilomolesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarFlowKilomolePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilomolesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilomolesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilomolesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilomolesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilomolesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilomolesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilomolesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilomolesPerSecond(0)
    if err != nil {
        t.Errorf("FromKilomolesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarFlowKilomolePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilomolesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilomolesPerMinute function
func TestMolarFlowFactory_FromKilomolesPerMinute(t *testing.T) {
    factory := units.MolarFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilomolesPerMinute(100)
    if err != nil {
        t.Errorf("FromKilomolesPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarFlowKilomolePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilomolesPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilomolesPerMinute(math.NaN())
    if err == nil {
        t.Error("FromKilomolesPerMinute() with NaN value should return error")
    }

    _, err = factory.FromKilomolesPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromKilomolesPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromKilomolesPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromKilomolesPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilomolesPerMinute(0)
    if err != nil {
        t.Errorf("FromKilomolesPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarFlowKilomolePerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilomolesPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromKilomolesPerHour function
func TestMolarFlowFactory_FromKilomolesPerHour(t *testing.T) {
    factory := units.MolarFlowFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilomolesPerHour(100)
    if err != nil {
        t.Errorf("FromKilomolesPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarFlowKilomolePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilomolesPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilomolesPerHour(math.NaN())
    if err == nil {
        t.Error("FromKilomolesPerHour() with NaN value should return error")
    }

    _, err = factory.FromKilomolesPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromKilomolesPerHour() with +Inf value should return error")
    }

    _, err = factory.FromKilomolesPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromKilomolesPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilomolesPerHour(0)
    if err != nil {
        t.Errorf("FromKilomolesPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarFlowKilomolePerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilomolesPerHour() with zero value = %v, want 0", converted)
    }
}

func TestMolarFlowToString(t *testing.T) {
	factory := units.MolarFlowFactory{}
	a, err := factory.CreateMolarFlow(45, units.MolarFlowMolePerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MolarFlowMolePerSecond, 2)
	expected := "45.00 " + units.GetMolarFlowAbbreviation(units.MolarFlowMolePerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MolarFlowMolePerSecond, -1)
	expected = "45 " + units.GetMolarFlowAbbreviation(units.MolarFlowMolePerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMolarFlow_EqualityAndComparison(t *testing.T) {
	factory := units.MolarFlowFactory{}
	a1, _ := factory.CreateMolarFlow(60, units.MolarFlowMolePerSecond)
	a2, _ := factory.CreateMolarFlow(60, units.MolarFlowMolePerSecond)
	a3, _ := factory.CreateMolarFlow(120, units.MolarFlowMolePerSecond)

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

func TestMolarFlow_Arithmetic(t *testing.T) {
	factory := units.MolarFlowFactory{}
	a1, _ := factory.CreateMolarFlow(30, units.MolarFlowMolePerSecond)
	a2, _ := factory.CreateMolarFlow(45, units.MolarFlowMolePerSecond)

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