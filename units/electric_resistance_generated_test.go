// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricResistanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Ohm"}`
	
	factory := units.ElectricResistanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricResistanceOhm {
		t.Errorf("expected unit %v, got %v", units.ElectricResistanceOhm, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Ohm"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricResistanceDto_ToJSON(t *testing.T) {
	dto := units.ElectricResistanceDto{
		Value: 45,
		Unit:  units.ElectricResistanceOhm,
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
	if result["unit"].(string) != string(units.ElectricResistanceOhm) {
		t.Errorf("expected unit %s, got %v", units.ElectricResistanceOhm, result["unit"])
	}
}

func TestNewElectricResistance_InvalidValue(t *testing.T) {
	factory := units.ElectricResistanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricResistance(math.NaN(), units.ElectricResistanceOhm)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricResistance(math.Inf(1), units.ElectricResistanceOhm)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricResistanceConversions(t *testing.T) {
	factory := units.ElectricResistanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricResistance(180, units.ElectricResistanceOhm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Ohms.
		// No expected conversion value provided for Ohms, verifying result is not NaN.
		result := a.Ohms()
		cacheResult := a.Ohms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Ohms returned NaN")
		}
	}
	{
		// Test conversion to Nanoohms.
		// No expected conversion value provided for Nanoohms, verifying result is not NaN.
		result := a.Nanoohms()
		cacheResult := a.Nanoohms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanoohms returned NaN")
		}
	}
	{
		// Test conversion to Microohms.
		// No expected conversion value provided for Microohms, verifying result is not NaN.
		result := a.Microohms()
		cacheResult := a.Microohms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microohms returned NaN")
		}
	}
	{
		// Test conversion to Milliohms.
		// No expected conversion value provided for Milliohms, verifying result is not NaN.
		result := a.Milliohms()
		cacheResult := a.Milliohms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Milliohms returned NaN")
		}
	}
	{
		// Test conversion to Kiloohms.
		// No expected conversion value provided for Kiloohms, verifying result is not NaN.
		result := a.Kiloohms()
		cacheResult := a.Kiloohms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kiloohms returned NaN")
		}
	}
	{
		// Test conversion to Megaohms.
		// No expected conversion value provided for Megaohms, verifying result is not NaN.
		result := a.Megaohms()
		cacheResult := a.Megaohms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megaohms returned NaN")
		}
	}
	{
		// Test conversion to Gigaohms.
		// No expected conversion value provided for Gigaohms, verifying result is not NaN.
		result := a.Gigaohms()
		cacheResult := a.Gigaohms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gigaohms returned NaN")
		}
	}
	{
		// Test conversion to Teraohms.
		// No expected conversion value provided for Teraohms, verifying result is not NaN.
		result := a.Teraohms()
		cacheResult := a.Teraohms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Teraohms returned NaN")
		}
	}
}

func TestElectricResistance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricResistanceFactory{}
	a, err := factory.CreateElectricResistance(90, units.ElectricResistanceOhm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricResistanceOhm {
		t.Errorf("expected default unit Ohm, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricResistanceOhm
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricResistanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricResistanceOhm {
		t.Errorf("expected unit Ohm, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricResistanceFactory_FromDto(t *testing.T) {
    factory := units.ElectricResistanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricResistanceDto{
        Value: 100,
        Unit:  units.ElectricResistanceOhm,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricResistanceDto{
        Value: math.NaN(),
        Unit:  units.ElectricResistanceOhm,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Ohm conversion
    ohmsDto := units.ElectricResistanceDto{
        Value: 100,
        Unit:  units.ElectricResistanceOhm,
    }
    
    var ohmsResult *units.ElectricResistance
    ohmsResult, err = factory.FromDto(ohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Ohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ohmsResult.Convert(units.ElectricResistanceOhm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Ohm = %v, want %v", converted, 100)
    }
    // Test Nanoohm conversion
    nanoohmsDto := units.ElectricResistanceDto{
        Value: 100,
        Unit:  units.ElectricResistanceNanoohm,
    }
    
    var nanoohmsResult *units.ElectricResistance
    nanoohmsResult, err = factory.FromDto(nanoohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Nanoohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoohmsResult.Convert(units.ElectricResistanceNanoohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanoohm = %v, want %v", converted, 100)
    }
    // Test Microohm conversion
    microohmsDto := units.ElectricResistanceDto{
        Value: 100,
        Unit:  units.ElectricResistanceMicroohm,
    }
    
    var microohmsResult *units.ElectricResistance
    microohmsResult, err = factory.FromDto(microohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Microohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microohmsResult.Convert(units.ElectricResistanceMicroohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microohm = %v, want %v", converted, 100)
    }
    // Test Milliohm conversion
    milliohmsDto := units.ElectricResistanceDto{
        Value: 100,
        Unit:  units.ElectricResistanceMilliohm,
    }
    
    var milliohmsResult *units.ElectricResistance
    milliohmsResult, err = factory.FromDto(milliohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Milliohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliohmsResult.Convert(units.ElectricResistanceMilliohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliohm = %v, want %v", converted, 100)
    }
    // Test Kiloohm conversion
    kiloohmsDto := units.ElectricResistanceDto{
        Value: 100,
        Unit:  units.ElectricResistanceKiloohm,
    }
    
    var kiloohmsResult *units.ElectricResistance
    kiloohmsResult, err = factory.FromDto(kiloohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Kiloohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloohmsResult.Convert(units.ElectricResistanceKiloohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kiloohm = %v, want %v", converted, 100)
    }
    // Test Megaohm conversion
    megaohmsDto := units.ElectricResistanceDto{
        Value: 100,
        Unit:  units.ElectricResistanceMegaohm,
    }
    
    var megaohmsResult *units.ElectricResistance
    megaohmsResult, err = factory.FromDto(megaohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Megaohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaohmsResult.Convert(units.ElectricResistanceMegaohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megaohm = %v, want %v", converted, 100)
    }
    // Test Gigaohm conversion
    gigaohmsDto := units.ElectricResistanceDto{
        Value: 100,
        Unit:  units.ElectricResistanceGigaohm,
    }
    
    var gigaohmsResult *units.ElectricResistance
    gigaohmsResult, err = factory.FromDto(gigaohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Gigaohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigaohmsResult.Convert(units.ElectricResistanceGigaohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigaohm = %v, want %v", converted, 100)
    }
    // Test Teraohm conversion
    teraohmsDto := units.ElectricResistanceDto{
        Value: 100,
        Unit:  units.ElectricResistanceTeraohm,
    }
    
    var teraohmsResult *units.ElectricResistance
    teraohmsResult, err = factory.FromDto(teraohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Teraohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teraohmsResult.Convert(units.ElectricResistanceTeraohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Teraohm = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricResistanceDto{
        Value: 0,
        Unit:  units.ElectricResistanceOhm,
    }
    
    var zeroResult *units.ElectricResistance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricResistanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricResistanceFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Ohm"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Ohm"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricResistanceDto{
        Value: nanValue,
        Unit:  units.ElectricResistanceOhm,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Ohm unit
    ohmsJSON := []byte(`{"value": 100, "unit": "Ohm"}`)
    ohmsResult, err := factory.FromDtoJSON(ohmsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Ohm unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ohmsResult.Convert(units.ElectricResistanceOhm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Ohm = %v, want %v", converted, 100)
    }
    // Test JSON with Nanoohm unit
    nanoohmsJSON := []byte(`{"value": 100, "unit": "Nanoohm"}`)
    nanoohmsResult, err := factory.FromDtoJSON(nanoohmsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanoohm unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoohmsResult.Convert(units.ElectricResistanceNanoohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanoohm = %v, want %v", converted, 100)
    }
    // Test JSON with Microohm unit
    microohmsJSON := []byte(`{"value": 100, "unit": "Microohm"}`)
    microohmsResult, err := factory.FromDtoJSON(microohmsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microohm unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microohmsResult.Convert(units.ElectricResistanceMicroohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microohm = %v, want %v", converted, 100)
    }
    // Test JSON with Milliohm unit
    milliohmsJSON := []byte(`{"value": 100, "unit": "Milliohm"}`)
    milliohmsResult, err := factory.FromDtoJSON(milliohmsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Milliohm unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliohmsResult.Convert(units.ElectricResistanceMilliohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliohm = %v, want %v", converted, 100)
    }
    // Test JSON with Kiloohm unit
    kiloohmsJSON := []byte(`{"value": 100, "unit": "Kiloohm"}`)
    kiloohmsResult, err := factory.FromDtoJSON(kiloohmsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kiloohm unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloohmsResult.Convert(units.ElectricResistanceKiloohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kiloohm = %v, want %v", converted, 100)
    }
    // Test JSON with Megaohm unit
    megaohmsJSON := []byte(`{"value": 100, "unit": "Megaohm"}`)
    megaohmsResult, err := factory.FromDtoJSON(megaohmsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megaohm unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaohmsResult.Convert(units.ElectricResistanceMegaohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megaohm = %v, want %v", converted, 100)
    }
    // Test JSON with Gigaohm unit
    gigaohmsJSON := []byte(`{"value": 100, "unit": "Gigaohm"}`)
    gigaohmsResult, err := factory.FromDtoJSON(gigaohmsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigaohm unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigaohmsResult.Convert(units.ElectricResistanceGigaohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigaohm = %v, want %v", converted, 100)
    }
    // Test JSON with Teraohm unit
    teraohmsJSON := []byte(`{"value": 100, "unit": "Teraohm"}`)
    teraohmsResult, err := factory.FromDtoJSON(teraohmsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Teraohm unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teraohmsResult.Convert(units.ElectricResistanceTeraohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Teraohm = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Ohm"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromOhms function
func TestElectricResistanceFactory_FromOhms(t *testing.T) {
    factory := units.ElectricResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOhms(100)
    if err != nil {
        t.Errorf("FromOhms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistanceOhm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOhms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOhms(math.NaN())
    if err == nil {
        t.Error("FromOhms() with NaN value should return error")
    }

    _, err = factory.FromOhms(math.Inf(1))
    if err == nil {
        t.Error("FromOhms() with +Inf value should return error")
    }

    _, err = factory.FromOhms(math.Inf(-1))
    if err == nil {
        t.Error("FromOhms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOhms(0)
    if err != nil {
        t.Errorf("FromOhms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistanceOhm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOhms() with zero value = %v, want 0", converted)
    }
}
// Test FromNanoohms function
func TestElectricResistanceFactory_FromNanoohms(t *testing.T) {
    factory := units.ElectricResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanoohms(100)
    if err != nil {
        t.Errorf("FromNanoohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistanceNanoohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanoohms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanoohms(math.NaN())
    if err == nil {
        t.Error("FromNanoohms() with NaN value should return error")
    }

    _, err = factory.FromNanoohms(math.Inf(1))
    if err == nil {
        t.Error("FromNanoohms() with +Inf value should return error")
    }

    _, err = factory.FromNanoohms(math.Inf(-1))
    if err == nil {
        t.Error("FromNanoohms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanoohms(0)
    if err != nil {
        t.Errorf("FromNanoohms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistanceNanoohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanoohms() with zero value = %v, want 0", converted)
    }
}
// Test FromMicroohms function
func TestElectricResistanceFactory_FromMicroohms(t *testing.T) {
    factory := units.ElectricResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicroohms(100)
    if err != nil {
        t.Errorf("FromMicroohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistanceMicroohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicroohms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicroohms(math.NaN())
    if err == nil {
        t.Error("FromMicroohms() with NaN value should return error")
    }

    _, err = factory.FromMicroohms(math.Inf(1))
    if err == nil {
        t.Error("FromMicroohms() with +Inf value should return error")
    }

    _, err = factory.FromMicroohms(math.Inf(-1))
    if err == nil {
        t.Error("FromMicroohms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicroohms(0)
    if err != nil {
        t.Errorf("FromMicroohms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistanceMicroohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicroohms() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliohms function
func TestElectricResistanceFactory_FromMilliohms(t *testing.T) {
    factory := units.ElectricResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliohms(100)
    if err != nil {
        t.Errorf("FromMilliohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistanceMilliohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliohms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliohms(math.NaN())
    if err == nil {
        t.Error("FromMilliohms() with NaN value should return error")
    }

    _, err = factory.FromMilliohms(math.Inf(1))
    if err == nil {
        t.Error("FromMilliohms() with +Inf value should return error")
    }

    _, err = factory.FromMilliohms(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliohms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliohms(0)
    if err != nil {
        t.Errorf("FromMilliohms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistanceMilliohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliohms() with zero value = %v, want 0", converted)
    }
}
// Test FromKiloohms function
func TestElectricResistanceFactory_FromKiloohms(t *testing.T) {
    factory := units.ElectricResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKiloohms(100)
    if err != nil {
        t.Errorf("FromKiloohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistanceKiloohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKiloohms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKiloohms(math.NaN())
    if err == nil {
        t.Error("FromKiloohms() with NaN value should return error")
    }

    _, err = factory.FromKiloohms(math.Inf(1))
    if err == nil {
        t.Error("FromKiloohms() with +Inf value should return error")
    }

    _, err = factory.FromKiloohms(math.Inf(-1))
    if err == nil {
        t.Error("FromKiloohms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKiloohms(0)
    if err != nil {
        t.Errorf("FromKiloohms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistanceKiloohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKiloohms() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaohms function
func TestElectricResistanceFactory_FromMegaohms(t *testing.T) {
    factory := units.ElectricResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaohms(100)
    if err != nil {
        t.Errorf("FromMegaohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistanceMegaohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegaohms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegaohms(math.NaN())
    if err == nil {
        t.Error("FromMegaohms() with NaN value should return error")
    }

    _, err = factory.FromMegaohms(math.Inf(1))
    if err == nil {
        t.Error("FromMegaohms() with +Inf value should return error")
    }

    _, err = factory.FromMegaohms(math.Inf(-1))
    if err == nil {
        t.Error("FromMegaohms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegaohms(0)
    if err != nil {
        t.Errorf("FromMegaohms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistanceMegaohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaohms() with zero value = %v, want 0", converted)
    }
}
// Test FromGigaohms function
func TestElectricResistanceFactory_FromGigaohms(t *testing.T) {
    factory := units.ElectricResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigaohms(100)
    if err != nil {
        t.Errorf("FromGigaohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistanceGigaohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigaohms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigaohms(math.NaN())
    if err == nil {
        t.Error("FromGigaohms() with NaN value should return error")
    }

    _, err = factory.FromGigaohms(math.Inf(1))
    if err == nil {
        t.Error("FromGigaohms() with +Inf value should return error")
    }

    _, err = factory.FromGigaohms(math.Inf(-1))
    if err == nil {
        t.Error("FromGigaohms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigaohms(0)
    if err != nil {
        t.Errorf("FromGigaohms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistanceGigaohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigaohms() with zero value = %v, want 0", converted)
    }
}
// Test FromTeraohms function
func TestElectricResistanceFactory_FromTeraohms(t *testing.T) {
    factory := units.ElectricResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTeraohms(100)
    if err != nil {
        t.Errorf("FromTeraohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricResistanceTeraohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTeraohms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTeraohms(math.NaN())
    if err == nil {
        t.Error("FromTeraohms() with NaN value should return error")
    }

    _, err = factory.FromTeraohms(math.Inf(1))
    if err == nil {
        t.Error("FromTeraohms() with +Inf value should return error")
    }

    _, err = factory.FromTeraohms(math.Inf(-1))
    if err == nil {
        t.Error("FromTeraohms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTeraohms(0)
    if err != nil {
        t.Errorf("FromTeraohms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricResistanceTeraohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTeraohms() with zero value = %v, want 0", converted)
    }
}

func TestElectricResistanceToString(t *testing.T) {
	factory := units.ElectricResistanceFactory{}
	a, err := factory.CreateElectricResistance(45, units.ElectricResistanceOhm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricResistanceOhm, 2)
	expected := "45.00 " + units.GetElectricResistanceAbbreviation(units.ElectricResistanceOhm)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricResistanceOhm, -1)
	expected = "45 " + units.GetElectricResistanceAbbreviation(units.ElectricResistanceOhm)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricResistance_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricResistanceFactory{}
	a1, _ := factory.CreateElectricResistance(60, units.ElectricResistanceOhm)
	a2, _ := factory.CreateElectricResistance(60, units.ElectricResistanceOhm)
	a3, _ := factory.CreateElectricResistance(120, units.ElectricResistanceOhm)

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

func TestElectricResistance_Arithmetic(t *testing.T) {
	factory := units.ElectricResistanceFactory{}
	a1, _ := factory.CreateElectricResistance(30, units.ElectricResistanceOhm)
	a2, _ := factory.CreateElectricResistance(45, units.ElectricResistanceOhm)

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


func TestGetElectricResistanceAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricResistanceUnits
        want string
    }{
        {
            name: "Ohm abbreviation",
            unit: units.ElectricResistanceOhm,
            want: "Ω",
        },
        {
            name: "Nanoohm abbreviation",
            unit: units.ElectricResistanceNanoohm,
            want: "nΩ",
        },
        {
            name: "Microohm abbreviation",
            unit: units.ElectricResistanceMicroohm,
            want: "μΩ",
        },
        {
            name: "Milliohm abbreviation",
            unit: units.ElectricResistanceMilliohm,
            want: "mΩ",
        },
        {
            name: "Kiloohm abbreviation",
            unit: units.ElectricResistanceKiloohm,
            want: "kΩ",
        },
        {
            name: "Megaohm abbreviation",
            unit: units.ElectricResistanceMegaohm,
            want: "MΩ",
        },
        {
            name: "Gigaohm abbreviation",
            unit: units.ElectricResistanceGigaohm,
            want: "GΩ",
        },
        {
            name: "Teraohm abbreviation",
            unit: units.ElectricResistanceTeraohm,
            want: "TΩ",
        },
        {
            name: "invalid unit",
            unit: units.ElectricResistanceUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricResistanceAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricResistanceAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricResistance_String(t *testing.T) {
    factory := units.ElectricResistanceFactory{}
    
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
            unit, err := factory.CreateElectricResistance(tt.value, units.ElectricResistanceOhm)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricResistance.String() = %v, want %v", got, tt.want)
            }
        })
    }
}