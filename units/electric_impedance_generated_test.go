// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricImpedanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Ohm"}`
	
	factory := units.ElectricImpedanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricImpedanceOhm {
		t.Errorf("expected unit %v, got %v", units.ElectricImpedanceOhm, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Ohm"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricImpedanceDto_ToJSON(t *testing.T) {
	dto := units.ElectricImpedanceDto{
		Value: 45,
		Unit:  units.ElectricImpedanceOhm,
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
	if result["unit"].(string) != string(units.ElectricImpedanceOhm) {
		t.Errorf("expected unit %s, got %v", units.ElectricImpedanceOhm, result["unit"])
	}
}

func TestNewElectricImpedance_InvalidValue(t *testing.T) {
	factory := units.ElectricImpedanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricImpedance(math.NaN(), units.ElectricImpedanceOhm)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricImpedance(math.Inf(1), units.ElectricImpedanceOhm)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricImpedanceConversions(t *testing.T) {
	factory := units.ElectricImpedanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricImpedance(180, units.ElectricImpedanceOhm)
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

func TestElectricImpedance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricImpedanceFactory{}
	a, err := factory.CreateElectricImpedance(90, units.ElectricImpedanceOhm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricImpedanceOhm {
		t.Errorf("expected default unit Ohm, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricImpedanceOhm
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricImpedanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricImpedanceOhm {
		t.Errorf("expected unit Ohm, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricImpedanceFactory_FromDto(t *testing.T) {
    factory := units.ElectricImpedanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricImpedanceDto{
        Value: 100,
        Unit:  units.ElectricImpedanceOhm,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricImpedanceDto{
        Value: math.NaN(),
        Unit:  units.ElectricImpedanceOhm,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Ohm conversion
    ohmsDto := units.ElectricImpedanceDto{
        Value: 100,
        Unit:  units.ElectricImpedanceOhm,
    }
    
    var ohmsResult *units.ElectricImpedance
    ohmsResult, err = factory.FromDto(ohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Ohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ohmsResult.Convert(units.ElectricImpedanceOhm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Ohm = %v, want %v", converted, 100)
    }
    // Test Nanoohm conversion
    nanoohmsDto := units.ElectricImpedanceDto{
        Value: 100,
        Unit:  units.ElectricImpedanceNanoohm,
    }
    
    var nanoohmsResult *units.ElectricImpedance
    nanoohmsResult, err = factory.FromDto(nanoohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Nanoohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoohmsResult.Convert(units.ElectricImpedanceNanoohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanoohm = %v, want %v", converted, 100)
    }
    // Test Microohm conversion
    microohmsDto := units.ElectricImpedanceDto{
        Value: 100,
        Unit:  units.ElectricImpedanceMicroohm,
    }
    
    var microohmsResult *units.ElectricImpedance
    microohmsResult, err = factory.FromDto(microohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Microohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microohmsResult.Convert(units.ElectricImpedanceMicroohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microohm = %v, want %v", converted, 100)
    }
    // Test Milliohm conversion
    milliohmsDto := units.ElectricImpedanceDto{
        Value: 100,
        Unit:  units.ElectricImpedanceMilliohm,
    }
    
    var milliohmsResult *units.ElectricImpedance
    milliohmsResult, err = factory.FromDto(milliohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Milliohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliohmsResult.Convert(units.ElectricImpedanceMilliohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliohm = %v, want %v", converted, 100)
    }
    // Test Kiloohm conversion
    kiloohmsDto := units.ElectricImpedanceDto{
        Value: 100,
        Unit:  units.ElectricImpedanceKiloohm,
    }
    
    var kiloohmsResult *units.ElectricImpedance
    kiloohmsResult, err = factory.FromDto(kiloohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Kiloohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloohmsResult.Convert(units.ElectricImpedanceKiloohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kiloohm = %v, want %v", converted, 100)
    }
    // Test Megaohm conversion
    megaohmsDto := units.ElectricImpedanceDto{
        Value: 100,
        Unit:  units.ElectricImpedanceMegaohm,
    }
    
    var megaohmsResult *units.ElectricImpedance
    megaohmsResult, err = factory.FromDto(megaohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Megaohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaohmsResult.Convert(units.ElectricImpedanceMegaohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megaohm = %v, want %v", converted, 100)
    }
    // Test Gigaohm conversion
    gigaohmsDto := units.ElectricImpedanceDto{
        Value: 100,
        Unit:  units.ElectricImpedanceGigaohm,
    }
    
    var gigaohmsResult *units.ElectricImpedance
    gigaohmsResult, err = factory.FromDto(gigaohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Gigaohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigaohmsResult.Convert(units.ElectricImpedanceGigaohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigaohm = %v, want %v", converted, 100)
    }
    // Test Teraohm conversion
    teraohmsDto := units.ElectricImpedanceDto{
        Value: 100,
        Unit:  units.ElectricImpedanceTeraohm,
    }
    
    var teraohmsResult *units.ElectricImpedance
    teraohmsResult, err = factory.FromDto(teraohmsDto)
    if err != nil {
        t.Errorf("FromDto() with Teraohm returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teraohmsResult.Convert(units.ElectricImpedanceTeraohm)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Teraohm = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricImpedanceDto{
        Value: 0,
        Unit:  units.ElectricImpedanceOhm,
    }
    
    var zeroResult *units.ElectricImpedance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricImpedanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricImpedanceFactory{}
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
    nanJSON, _ := json.Marshal(units.ElectricImpedanceDto{
        Value: nanValue,
        Unit:  units.ElectricImpedanceOhm,
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
    converted = ohmsResult.Convert(units.ElectricImpedanceOhm)
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
    converted = nanoohmsResult.Convert(units.ElectricImpedanceNanoohm)
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
    converted = microohmsResult.Convert(units.ElectricImpedanceMicroohm)
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
    converted = milliohmsResult.Convert(units.ElectricImpedanceMilliohm)
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
    converted = kiloohmsResult.Convert(units.ElectricImpedanceKiloohm)
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
    converted = megaohmsResult.Convert(units.ElectricImpedanceMegaohm)
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
    converted = gigaohmsResult.Convert(units.ElectricImpedanceGigaohm)
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
    converted = teraohmsResult.Convert(units.ElectricImpedanceTeraohm)
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
func TestElectricImpedanceFactory_FromOhms(t *testing.T) {
    factory := units.ElectricImpedanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOhms(100)
    if err != nil {
        t.Errorf("FromOhms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricImpedanceOhm)
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
    converted = zeroResult.Convert(units.ElectricImpedanceOhm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOhms() with zero value = %v, want 0", converted)
    }
}
// Test FromNanoohms function
func TestElectricImpedanceFactory_FromNanoohms(t *testing.T) {
    factory := units.ElectricImpedanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanoohms(100)
    if err != nil {
        t.Errorf("FromNanoohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricImpedanceNanoohm)
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
    converted = zeroResult.Convert(units.ElectricImpedanceNanoohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanoohms() with zero value = %v, want 0", converted)
    }
}
// Test FromMicroohms function
func TestElectricImpedanceFactory_FromMicroohms(t *testing.T) {
    factory := units.ElectricImpedanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicroohms(100)
    if err != nil {
        t.Errorf("FromMicroohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricImpedanceMicroohm)
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
    converted = zeroResult.Convert(units.ElectricImpedanceMicroohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicroohms() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliohms function
func TestElectricImpedanceFactory_FromMilliohms(t *testing.T) {
    factory := units.ElectricImpedanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliohms(100)
    if err != nil {
        t.Errorf("FromMilliohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricImpedanceMilliohm)
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
    converted = zeroResult.Convert(units.ElectricImpedanceMilliohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliohms() with zero value = %v, want 0", converted)
    }
}
// Test FromKiloohms function
func TestElectricImpedanceFactory_FromKiloohms(t *testing.T) {
    factory := units.ElectricImpedanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKiloohms(100)
    if err != nil {
        t.Errorf("FromKiloohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricImpedanceKiloohm)
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
    converted = zeroResult.Convert(units.ElectricImpedanceKiloohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKiloohms() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaohms function
func TestElectricImpedanceFactory_FromMegaohms(t *testing.T) {
    factory := units.ElectricImpedanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaohms(100)
    if err != nil {
        t.Errorf("FromMegaohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricImpedanceMegaohm)
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
    converted = zeroResult.Convert(units.ElectricImpedanceMegaohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaohms() with zero value = %v, want 0", converted)
    }
}
// Test FromGigaohms function
func TestElectricImpedanceFactory_FromGigaohms(t *testing.T) {
    factory := units.ElectricImpedanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigaohms(100)
    if err != nil {
        t.Errorf("FromGigaohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricImpedanceGigaohm)
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
    converted = zeroResult.Convert(units.ElectricImpedanceGigaohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigaohms() with zero value = %v, want 0", converted)
    }
}
// Test FromTeraohms function
func TestElectricImpedanceFactory_FromTeraohms(t *testing.T) {
    factory := units.ElectricImpedanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTeraohms(100)
    if err != nil {
        t.Errorf("FromTeraohms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricImpedanceTeraohm)
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
    converted = zeroResult.Convert(units.ElectricImpedanceTeraohm)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTeraohms() with zero value = %v, want 0", converted)
    }
}

func TestElectricImpedanceToString(t *testing.T) {
	factory := units.ElectricImpedanceFactory{}
	a, err := factory.CreateElectricImpedance(45, units.ElectricImpedanceOhm)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricImpedanceOhm, 2)
	expected := "45.00 " + units.GetElectricImpedanceAbbreviation(units.ElectricImpedanceOhm)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricImpedanceOhm, -1)
	expected = "45 " + units.GetElectricImpedanceAbbreviation(units.ElectricImpedanceOhm)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricImpedance_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricImpedanceFactory{}
	a1, _ := factory.CreateElectricImpedance(60, units.ElectricImpedanceOhm)
	a2, _ := factory.CreateElectricImpedance(60, units.ElectricImpedanceOhm)
	a3, _ := factory.CreateElectricImpedance(120, units.ElectricImpedanceOhm)

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

func TestElectricImpedance_Arithmetic(t *testing.T) {
	factory := units.ElectricImpedanceFactory{}
	a1, _ := factory.CreateElectricImpedance(30, units.ElectricImpedanceOhm)
	a2, _ := factory.CreateElectricImpedance(45, units.ElectricImpedanceOhm)

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