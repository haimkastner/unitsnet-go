// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricReactanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Ohm"}`
	
	factory := units.ElectricReactanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricReactanceOhm {
		t.Errorf("expected unit %v, got %v", units.ElectricReactanceOhm, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Ohm"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricReactanceDto_ToJSON(t *testing.T) {
	dto := units.ElectricReactanceDto{
		Value: 45,
		Unit:  units.ElectricReactanceOhm,
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
	if result["unit"].(string) != string(units.ElectricReactanceOhm) {
		t.Errorf("expected unit %s, got %v", units.ElectricReactanceOhm, result["unit"])
	}
}

func TestNewElectricReactance_InvalidValue(t *testing.T) {
	factory := units.ElectricReactanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricReactance(math.NaN(), units.ElectricReactanceOhm)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricReactance(math.Inf(1), units.ElectricReactanceOhm)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricReactanceConversions(t *testing.T) {
	factory := units.ElectricReactanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricReactance(180, units.ElectricReactanceOhm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Ohms.
		// No expected conversion value provided for Ohms, verifying result is not NaN.
		result := a.Ohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Ohms returned NaN")
		}
	}
	{
		// Test conversion to Nanoohms.
		// No expected conversion value provided for Nanoohms, verifying result is not NaN.
		result := a.Nanoohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanoohms returned NaN")
		}
	}
	{
		// Test conversion to Microohms.
		// No expected conversion value provided for Microohms, verifying result is not NaN.
		result := a.Microohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microohms returned NaN")
		}
	}
	{
		// Test conversion to Milliohms.
		// No expected conversion value provided for Milliohms, verifying result is not NaN.
		result := a.Milliohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milliohms returned NaN")
		}
	}
	{
		// Test conversion to Kiloohms.
		// No expected conversion value provided for Kiloohms, verifying result is not NaN.
		result := a.Kiloohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kiloohms returned NaN")
		}
	}
	{
		// Test conversion to Megaohms.
		// No expected conversion value provided for Megaohms, verifying result is not NaN.
		result := a.Megaohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megaohms returned NaN")
		}
	}
	{
		// Test conversion to Gigaohms.
		// No expected conversion value provided for Gigaohms, verifying result is not NaN.
		result := a.Gigaohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigaohms returned NaN")
		}
	}
	{
		// Test conversion to Teraohms.
		// No expected conversion value provided for Teraohms, verifying result is not NaN.
		result := a.Teraohms()
		if math.IsNaN(result) {
			t.Errorf("conversion to Teraohms returned NaN")
		}
	}
}

func TestElectricReactance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricReactanceFactory{}
	a, err := factory.CreateElectricReactance(90, units.ElectricReactanceOhm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricReactanceOhm {
		t.Errorf("expected default unit Ohm, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricReactanceOhm
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricReactanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricReactanceOhm {
		t.Errorf("expected unit Ohm, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricReactanceFactory_FromDto(t *testing.T) {
    factory := units.ElectricReactanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricReactanceDto{
        Value: 100,
        Unit:  units.ElectricReactanceOhm,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricReactanceDto{
        Value: math.NaN(),
        Unit:  units.ElectricReactanceOhm,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Ohm conversion
    ohmsDto := units.ElectricReactanceDto{
        Value: 100,
        Unit:  units.ElectricReactanceOhm,
    }
    
    var ohmsResult *units.ElectricReactance
    ohmsResult, err = factory.FromDto(ohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Ohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ohmsResult.Convert(units.ElectricReactanceOhm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Ohm = %v, want %v", converted, 100)
    }
    // Test Nanoohm conversion
    nanoohmsDto := units.ElectricReactanceDto{
        Value: 100,
        Unit:  units.ElectricReactanceNanoohm,
    }
    
    var nanoohmsResult *units.ElectricReactance
    nanoohmsResult, err = factory.FromDto(nanoohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Nanoohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoohmsResult.Convert(units.ElectricReactanceNanoohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanoohm = %v, want %v", converted, 100)
    }
    // Test Microohm conversion
    microohmsDto := units.ElectricReactanceDto{
        Value: 100,
        Unit:  units.ElectricReactanceMicroohm,
    }
    
    var microohmsResult *units.ElectricReactance
    microohmsResult, err = factory.FromDto(microohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Microohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microohmsResult.Convert(units.ElectricReactanceMicroohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microohm = %v, want %v", converted, 100)
    }
    // Test Milliohm conversion
    milliohmsDto := units.ElectricReactanceDto{
        Value: 100,
        Unit:  units.ElectricReactanceMilliohm,
    }
    
    var milliohmsResult *units.ElectricReactance
    milliohmsResult, err = factory.FromDto(milliohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Milliohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliohmsResult.Convert(units.ElectricReactanceMilliohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliohm = %v, want %v", converted, 100)
    }
    // Test Kiloohm conversion
    kiloohmsDto := units.ElectricReactanceDto{
        Value: 100,
        Unit:  units.ElectricReactanceKiloohm,
    }
    
    var kiloohmsResult *units.ElectricReactance
    kiloohmsResult, err = factory.FromDto(kiloohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Kiloohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloohmsResult.Convert(units.ElectricReactanceKiloohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kiloohm = %v, want %v", converted, 100)
    }
    // Test Megaohm conversion
    megaohmsDto := units.ElectricReactanceDto{
        Value: 100,
        Unit:  units.ElectricReactanceMegaohm,
    }
    
    var megaohmsResult *units.ElectricReactance
    megaohmsResult, err = factory.FromDto(megaohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Megaohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaohmsResult.Convert(units.ElectricReactanceMegaohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megaohm = %v, want %v", converted, 100)
    }
    // Test Gigaohm conversion
    gigaohmsDto := units.ElectricReactanceDto{
        Value: 100,
        Unit:  units.ElectricReactanceGigaohm,
    }
    
    var gigaohmsResult *units.ElectricReactance
    gigaohmsResult, err = factory.FromDto(gigaohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Gigaohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigaohmsResult.Convert(units.ElectricReactanceGigaohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigaohm = %v, want %v", converted, 100)
    }
    // Test Teraohm conversion
    teraohmsDto := units.ElectricReactanceDto{
        Value: 100,
        Unit:  units.ElectricReactanceTeraohm,
    }
    
    var teraohmsResult *units.ElectricReactance
    teraohmsResult, err = factory.FromDto(teraohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Teraohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teraohmsResult.Convert(units.ElectricReactanceTeraohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Teraohm = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricReactanceDto{
        Value: 0,
        Unit:  units.ElectricReactanceOhm,
    }
    
    var zeroResult *units.ElectricReactance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricReactanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricReactanceFactory{}
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
    nanJSON, _ := json.Marshal(units.ElectricReactanceDto{
        Value: nanValue,
        Unit:  units.ElectricReactanceOhm,
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
    converted = ohmsResult.Convert(units.ElectricReactanceOhm)
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
    converted = nanoohmsResult.Convert(units.ElectricReactanceNanoohm)
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
    converted = microohmsResult.Convert(units.ElectricReactanceMicroohm)
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
    converted = milliohmsResult.Convert(units.ElectricReactanceMilliohm)
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
    converted = kiloohmsResult.Convert(units.ElectricReactanceKiloohm)
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
    converted = megaohmsResult.Convert(units.ElectricReactanceMegaohm)
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
    converted = gigaohmsResult.Convert(units.ElectricReactanceGigaohm)
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
    converted = teraohmsResult.Convert(units.ElectricReactanceTeraohm)
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
func TestElectricReactanceFactory_FromOhms(t *testing.T) {
    factory := units.ElectricReactanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOhms(100)
    if err != nil {
        t.Errorf("FromOhms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricReactanceOhm)
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
    converted = zeroResult.Convert(units.ElectricReactanceOhm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOhms() with zero value = %v, want 0", converted)
    }
}
// Test FromNanoohms function
func TestElectricReactanceFactory_FromNanoohms(t *testing.T) {
    factory := units.ElectricReactanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanoohms(100)
    if err != nil {
        t.Errorf("FromNanoohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricReactanceNanoohm)
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
    converted = zeroResult.Convert(units.ElectricReactanceNanoohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanoohms() with zero value = %v, want 0", converted)
    }
}
// Test FromMicroohms function
func TestElectricReactanceFactory_FromMicroohms(t *testing.T) {
    factory := units.ElectricReactanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicroohms(100)
    if err != nil {
        t.Errorf("FromMicroohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricReactanceMicroohm)
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
    converted = zeroResult.Convert(units.ElectricReactanceMicroohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicroohms() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliohms function
func TestElectricReactanceFactory_FromMilliohms(t *testing.T) {
    factory := units.ElectricReactanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliohms(100)
    if err != nil {
        t.Errorf("FromMilliohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricReactanceMilliohm)
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
    converted = zeroResult.Convert(units.ElectricReactanceMilliohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliohms() with zero value = %v, want 0", converted)
    }
}
// Test FromKiloohms function
func TestElectricReactanceFactory_FromKiloohms(t *testing.T) {
    factory := units.ElectricReactanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKiloohms(100)
    if err != nil {
        t.Errorf("FromKiloohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricReactanceKiloohm)
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
    converted = zeroResult.Convert(units.ElectricReactanceKiloohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKiloohms() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaohms function
func TestElectricReactanceFactory_FromMegaohms(t *testing.T) {
    factory := units.ElectricReactanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaohms(100)
    if err != nil {
        t.Errorf("FromMegaohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricReactanceMegaohm)
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
    converted = zeroResult.Convert(units.ElectricReactanceMegaohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaohms() with zero value = %v, want 0", converted)
    }
}
// Test FromGigaohms function
func TestElectricReactanceFactory_FromGigaohms(t *testing.T) {
    factory := units.ElectricReactanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigaohms(100)
    if err != nil {
        t.Errorf("FromGigaohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricReactanceGigaohm)
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
    converted = zeroResult.Convert(units.ElectricReactanceGigaohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigaohms() with zero value = %v, want 0", converted)
    }
}
// Test FromTeraohms function
func TestElectricReactanceFactory_FromTeraohms(t *testing.T) {
    factory := units.ElectricReactanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTeraohms(100)
    if err != nil {
        t.Errorf("FromTeraohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricReactanceTeraohm)
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
    converted = zeroResult.Convert(units.ElectricReactanceTeraohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTeraohms() with zero value = %v, want 0", converted)
    }
}

func TestElectricReactanceToString(t *testing.T) {
	factory := units.ElectricReactanceFactory{}
	a, err := factory.CreateElectricReactance(45, units.ElectricReactanceOhm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricReactanceOhm, 2)
	expected := "45.00 " + units.GetElectricReactanceAbbreviation(units.ElectricReactanceOhm)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricReactanceOhm, -1)
	expected = "45 " + units.GetElectricReactanceAbbreviation(units.ElectricReactanceOhm)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricReactance_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricReactanceFactory{}
	a1, _ := factory.CreateElectricReactance(60, units.ElectricReactanceOhm)
	a2, _ := factory.CreateElectricReactance(60, units.ElectricReactanceOhm)
	a3, _ := factory.CreateElectricReactance(120, units.ElectricReactanceOhm)

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

func TestElectricReactance_Arithmetic(t *testing.T) {
	factory := units.ElectricReactanceFactory{}
	a1, _ := factory.CreateElectricReactance(30, units.ElectricReactanceOhm)
	a2, _ := factory.CreateElectricReactance(45, units.ElectricReactanceOhm)

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