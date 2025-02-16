// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestApparentPowerDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Voltampere"}`
	
	factory := units.ApparentPowerDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ApparentPowerVoltampere {
		t.Errorf("expected unit %v, got %v", units.ApparentPowerVoltampere, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Voltampere"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestApparentPowerDto_ToJSON(t *testing.T) {
	dto := units.ApparentPowerDto{
		Value: 45,
		Unit:  units.ApparentPowerVoltampere,
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
	if result["unit"].(string) != string(units.ApparentPowerVoltampere) {
		t.Errorf("expected unit %s, got %v", units.ApparentPowerVoltampere, result["unit"])
	}
}

func TestNewApparentPower_InvalidValue(t *testing.T) {
	factory := units.ApparentPowerFactory{}
	// NaN value should return an error.
	_, err := factory.CreateApparentPower(math.NaN(), units.ApparentPowerVoltampere)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateApparentPower(math.Inf(1), units.ApparentPowerVoltampere)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestApparentPowerConversions(t *testing.T) {
	factory := units.ApparentPowerFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateApparentPower(180, units.ApparentPowerVoltampere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Voltamperes.
		// No expected conversion value provided for Voltamperes, verifying result is not NaN.
		result := a.Voltamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Voltamperes returned NaN")
		}
	}
	{
		// Test conversion to Microvoltamperes.
		// No expected conversion value provided for Microvoltamperes, verifying result is not NaN.
		result := a.Microvoltamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microvoltamperes returned NaN")
		}
	}
	{
		// Test conversion to Millivoltamperes.
		// No expected conversion value provided for Millivoltamperes, verifying result is not NaN.
		result := a.Millivoltamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millivoltamperes returned NaN")
		}
	}
	{
		// Test conversion to Kilovoltamperes.
		// No expected conversion value provided for Kilovoltamperes, verifying result is not NaN.
		result := a.Kilovoltamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilovoltamperes returned NaN")
		}
	}
	{
		// Test conversion to Megavoltamperes.
		// No expected conversion value provided for Megavoltamperes, verifying result is not NaN.
		result := a.Megavoltamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megavoltamperes returned NaN")
		}
	}
	{
		// Test conversion to Gigavoltamperes.
		// No expected conversion value provided for Gigavoltamperes, verifying result is not NaN.
		result := a.Gigavoltamperes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigavoltamperes returned NaN")
		}
	}
}

func TestApparentPower_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ApparentPowerFactory{}
	a, err := factory.CreateApparentPower(90, units.ApparentPowerVoltampere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ApparentPowerVoltampere {
		t.Errorf("expected default unit Voltampere, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ApparentPowerVoltampere
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ApparentPowerDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ApparentPowerVoltampere {
		t.Errorf("expected unit Voltampere, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestApparentPowerFactory_FromDto(t *testing.T) {
    factory := units.ApparentPowerFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ApparentPowerDto{
        Value: 100,
        Unit:  units.ApparentPowerVoltampere,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ApparentPowerDto{
        Value: math.NaN(),
        Unit:  units.ApparentPowerVoltampere,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Voltampere conversion
    voltamperesDto := units.ApparentPowerDto{
        Value: 100,
        Unit:  units.ApparentPowerVoltampere,
    }
    
    var voltamperesResult *units.ApparentPower
    voltamperesResult, err = factory.FromDto(voltamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Voltampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = voltamperesResult.Convert(units.ApparentPowerVoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Voltampere = %v, want %v", converted, 100)
    }
    // Test Microvoltampere conversion
    microvoltamperesDto := units.ApparentPowerDto{
        Value: 100,
        Unit:  units.ApparentPowerMicrovoltampere,
    }
    
    var microvoltamperesResult *units.ApparentPower
    microvoltamperesResult, err = factory.FromDto(microvoltamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Microvoltampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvoltamperesResult.Convert(units.ApparentPowerMicrovoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microvoltampere = %v, want %v", converted, 100)
    }
    // Test Millivoltampere conversion
    millivoltamperesDto := units.ApparentPowerDto{
        Value: 100,
        Unit:  units.ApparentPowerMillivoltampere,
    }
    
    var millivoltamperesResult *units.ApparentPower
    millivoltamperesResult, err = factory.FromDto(millivoltamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Millivoltampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivoltamperesResult.Convert(units.ApparentPowerMillivoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millivoltampere = %v, want %v", converted, 100)
    }
    // Test Kilovoltampere conversion
    kilovoltamperesDto := units.ApparentPowerDto{
        Value: 100,
        Unit:  units.ApparentPowerKilovoltampere,
    }
    
    var kilovoltamperesResult *units.ApparentPower
    kilovoltamperesResult, err = factory.FromDto(kilovoltamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Kilovoltampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovoltamperesResult.Convert(units.ApparentPowerKilovoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilovoltampere = %v, want %v", converted, 100)
    }
    // Test Megavoltampere conversion
    megavoltamperesDto := units.ApparentPowerDto{
        Value: 100,
        Unit:  units.ApparentPowerMegavoltampere,
    }
    
    var megavoltamperesResult *units.ApparentPower
    megavoltamperesResult, err = factory.FromDto(megavoltamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Megavoltampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavoltamperesResult.Convert(units.ApparentPowerMegavoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megavoltampere = %v, want %v", converted, 100)
    }
    // Test Gigavoltampere conversion
    gigavoltamperesDto := units.ApparentPowerDto{
        Value: 100,
        Unit:  units.ApparentPowerGigavoltampere,
    }
    
    var gigavoltamperesResult *units.ApparentPower
    gigavoltamperesResult, err = factory.FromDto(gigavoltamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Gigavoltampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigavoltamperesResult.Convert(units.ApparentPowerGigavoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigavoltampere = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ApparentPowerDto{
        Value: 0,
        Unit:  units.ApparentPowerVoltampere,
    }
    
    var zeroResult *units.ApparentPower
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestApparentPowerFactory_FromDtoJSON(t *testing.T) {
    factory := units.ApparentPowerFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Voltampere"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Voltampere"}`)
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
    nanJSON, _ := json.Marshal(units.ApparentPowerDto{
        Value: nanValue,
        Unit:  units.ApparentPowerVoltampere,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Voltampere unit
    voltamperesJSON := []byte(`{"value": 100, "unit": "Voltampere"}`)
    voltamperesResult, err := factory.FromDtoJSON(voltamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Voltampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = voltamperesResult.Convert(units.ApparentPowerVoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Voltampere = %v, want %v", converted, 100)
    }
    // Test JSON with Microvoltampere unit
    microvoltamperesJSON := []byte(`{"value": 100, "unit": "Microvoltampere"}`)
    microvoltamperesResult, err := factory.FromDtoJSON(microvoltamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microvoltampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvoltamperesResult.Convert(units.ApparentPowerMicrovoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microvoltampere = %v, want %v", converted, 100)
    }
    // Test JSON with Millivoltampere unit
    millivoltamperesJSON := []byte(`{"value": 100, "unit": "Millivoltampere"}`)
    millivoltamperesResult, err := factory.FromDtoJSON(millivoltamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millivoltampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivoltamperesResult.Convert(units.ApparentPowerMillivoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millivoltampere = %v, want %v", converted, 100)
    }
    // Test JSON with Kilovoltampere unit
    kilovoltamperesJSON := []byte(`{"value": 100, "unit": "Kilovoltampere"}`)
    kilovoltamperesResult, err := factory.FromDtoJSON(kilovoltamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilovoltampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovoltamperesResult.Convert(units.ApparentPowerKilovoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilovoltampere = %v, want %v", converted, 100)
    }
    // Test JSON with Megavoltampere unit
    megavoltamperesJSON := []byte(`{"value": 100, "unit": "Megavoltampere"}`)
    megavoltamperesResult, err := factory.FromDtoJSON(megavoltamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megavoltampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavoltamperesResult.Convert(units.ApparentPowerMegavoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megavoltampere = %v, want %v", converted, 100)
    }
    // Test JSON with Gigavoltampere unit
    gigavoltamperesJSON := []byte(`{"value": 100, "unit": "Gigavoltampere"}`)
    gigavoltamperesResult, err := factory.FromDtoJSON(gigavoltamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigavoltampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigavoltamperesResult.Convert(units.ApparentPowerGigavoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigavoltampere = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Voltampere"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromVoltamperes function
func TestApparentPowerFactory_FromVoltamperes(t *testing.T) {
    factory := units.ApparentPowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltamperes(100)
    if err != nil {
        t.Errorf("FromVoltamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ApparentPowerVoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltamperes(math.NaN())
    if err == nil {
        t.Error("FromVoltamperes() with NaN value should return error")
    }

    _, err = factory.FromVoltamperes(math.Inf(1))
    if err == nil {
        t.Error("FromVoltamperes() with +Inf value should return error")
    }

    _, err = factory.FromVoltamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltamperes(0)
    if err != nil {
        t.Errorf("FromVoltamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ApparentPowerVoltampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltamperes() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrovoltamperes function
func TestApparentPowerFactory_FromMicrovoltamperes(t *testing.T) {
    factory := units.ApparentPowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrovoltamperes(100)
    if err != nil {
        t.Errorf("FromMicrovoltamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ApparentPowerMicrovoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrovoltamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrovoltamperes(math.NaN())
    if err == nil {
        t.Error("FromMicrovoltamperes() with NaN value should return error")
    }

    _, err = factory.FromMicrovoltamperes(math.Inf(1))
    if err == nil {
        t.Error("FromMicrovoltamperes() with +Inf value should return error")
    }

    _, err = factory.FromMicrovoltamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrovoltamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrovoltamperes(0)
    if err != nil {
        t.Errorf("FromMicrovoltamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ApparentPowerMicrovoltampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrovoltamperes() with zero value = %v, want 0", converted)
    }
}
// Test FromMillivoltamperes function
func TestApparentPowerFactory_FromMillivoltamperes(t *testing.T) {
    factory := units.ApparentPowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillivoltamperes(100)
    if err != nil {
        t.Errorf("FromMillivoltamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ApparentPowerMillivoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillivoltamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillivoltamperes(math.NaN())
    if err == nil {
        t.Error("FromMillivoltamperes() with NaN value should return error")
    }

    _, err = factory.FromMillivoltamperes(math.Inf(1))
    if err == nil {
        t.Error("FromMillivoltamperes() with +Inf value should return error")
    }

    _, err = factory.FromMillivoltamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromMillivoltamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillivoltamperes(0)
    if err != nil {
        t.Errorf("FromMillivoltamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ApparentPowerMillivoltampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillivoltamperes() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltamperes function
func TestApparentPowerFactory_FromKilovoltamperes(t *testing.T) {
    factory := units.ApparentPowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltamperes(100)
    if err != nil {
        t.Errorf("FromKilovoltamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ApparentPowerKilovoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltamperes(math.NaN())
    if err == nil {
        t.Error("FromKilovoltamperes() with NaN value should return error")
    }

    _, err = factory.FromKilovoltamperes(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltamperes() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltamperes(0)
    if err != nil {
        t.Errorf("FromKilovoltamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ApparentPowerKilovoltampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltamperes() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltamperes function
func TestApparentPowerFactory_FromMegavoltamperes(t *testing.T) {
    factory := units.ApparentPowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltamperes(100)
    if err != nil {
        t.Errorf("FromMegavoltamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ApparentPowerMegavoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltamperes(math.NaN())
    if err == nil {
        t.Error("FromMegavoltamperes() with NaN value should return error")
    }

    _, err = factory.FromMegavoltamperes(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltamperes() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltamperes(0)
    if err != nil {
        t.Errorf("FromMegavoltamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ApparentPowerMegavoltampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltamperes() with zero value = %v, want 0", converted)
    }
}
// Test FromGigavoltamperes function
func TestApparentPowerFactory_FromGigavoltamperes(t *testing.T) {
    factory := units.ApparentPowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigavoltamperes(100)
    if err != nil {
        t.Errorf("FromGigavoltamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ApparentPowerGigavoltampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigavoltamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigavoltamperes(math.NaN())
    if err == nil {
        t.Error("FromGigavoltamperes() with NaN value should return error")
    }

    _, err = factory.FromGigavoltamperes(math.Inf(1))
    if err == nil {
        t.Error("FromGigavoltamperes() with +Inf value should return error")
    }

    _, err = factory.FromGigavoltamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromGigavoltamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigavoltamperes(0)
    if err != nil {
        t.Errorf("FromGigavoltamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ApparentPowerGigavoltampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigavoltamperes() with zero value = %v, want 0", converted)
    }
}

func TestApparentPowerToString(t *testing.T) {
	factory := units.ApparentPowerFactory{}
	a, err := factory.CreateApparentPower(45, units.ApparentPowerVoltampere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ApparentPowerVoltampere, 2)
	expected := "45.00 " + units.GetApparentPowerAbbreviation(units.ApparentPowerVoltampere)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ApparentPowerVoltampere, -1)
	expected = "45 " + units.GetApparentPowerAbbreviation(units.ApparentPowerVoltampere)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestApparentPower_EqualityAndComparison(t *testing.T) {
	factory := units.ApparentPowerFactory{}
	a1, _ := factory.CreateApparentPower(60, units.ApparentPowerVoltampere)
	a2, _ := factory.CreateApparentPower(60, units.ApparentPowerVoltampere)
	a3, _ := factory.CreateApparentPower(120, units.ApparentPowerVoltampere)

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

func TestApparentPower_Arithmetic(t *testing.T) {
	factory := units.ApparentPowerFactory{}
	a1, _ := factory.CreateApparentPower(30, units.ApparentPowerVoltampere)
	a2, _ := factory.CreateApparentPower(45, units.ApparentPowerVoltampere)

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