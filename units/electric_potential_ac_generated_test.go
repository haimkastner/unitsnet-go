// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricPotentialAcDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltAc"}`
	
	factory := units.ElectricPotentialAcDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricPotentialAcVoltAc {
		t.Errorf("expected unit %v, got %v", units.ElectricPotentialAcVoltAc, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltAc"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricPotentialAcDto_ToJSON(t *testing.T) {
	dto := units.ElectricPotentialAcDto{
		Value: 45,
		Unit:  units.ElectricPotentialAcVoltAc,
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
	if result["unit"].(string) != string(units.ElectricPotentialAcVoltAc) {
		t.Errorf("expected unit %s, got %v", units.ElectricPotentialAcVoltAc, result["unit"])
	}
}

func TestNewElectricPotentialAc_InvalidValue(t *testing.T) {
	factory := units.ElectricPotentialAcFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricPotentialAc(math.NaN(), units.ElectricPotentialAcVoltAc)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricPotentialAc(math.Inf(1), units.ElectricPotentialAcVoltAc)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricPotentialAcConversions(t *testing.T) {
	factory := units.ElectricPotentialAcFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricPotentialAc(180, units.ElectricPotentialAcVoltAc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltsAc.
		// No expected conversion value provided for VoltsAc, verifying result is not NaN.
		result := a.VoltsAc()
		cacheResult := a.VoltsAc()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltsAc returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsAc.
		// No expected conversion value provided for MicrovoltsAc, verifying result is not NaN.
		result := a.MicrovoltsAc()
		cacheResult := a.MicrovoltsAc()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrovoltsAc returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsAc.
		// No expected conversion value provided for MillivoltsAc, verifying result is not NaN.
		result := a.MillivoltsAc()
		cacheResult := a.MillivoltsAc()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillivoltsAc returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsAc.
		// No expected conversion value provided for KilovoltsAc, verifying result is not NaN.
		result := a.KilovoltsAc()
		cacheResult := a.KilovoltsAc()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilovoltsAc returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsAc.
		// No expected conversion value provided for MegavoltsAc, verifying result is not NaN.
		result := a.MegavoltsAc()
		cacheResult := a.MegavoltsAc()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegavoltsAc returned NaN")
		}
	}
}

func TestElectricPotentialAc_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricPotentialAcFactory{}
	a, err := factory.CreateElectricPotentialAc(90, units.ElectricPotentialAcVoltAc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricPotentialAcVoltAc {
		t.Errorf("expected default unit VoltAc, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricPotentialAcVoltAc
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricPotentialAcDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricPotentialAcVoltAc {
		t.Errorf("expected unit VoltAc, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricPotentialAcFactory_FromDto(t *testing.T) {
    factory := units.ElectricPotentialAcFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricPotentialAcDto{
        Value: 100,
        Unit:  units.ElectricPotentialAcVoltAc,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricPotentialAcDto{
        Value: math.NaN(),
        Unit:  units.ElectricPotentialAcVoltAc,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test VoltAc conversion
    volts_acDto := units.ElectricPotentialAcDto{
        Value: 100,
        Unit:  units.ElectricPotentialAcVoltAc,
    }
    
    var volts_acResult *units.ElectricPotentialAc
    volts_acResult, err = factory.FromDto(volts_acDto)
    if err != nil {
        t.Errorf("FromDto() with VoltAc returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_acResult.Convert(units.ElectricPotentialAcVoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltAc = %v, want %v", converted, 100)
    }
    // Test MicrovoltAc conversion
    microvolts_acDto := units.ElectricPotentialAcDto{
        Value: 100,
        Unit:  units.ElectricPotentialAcMicrovoltAc,
    }
    
    var microvolts_acResult *units.ElectricPotentialAc
    microvolts_acResult, err = factory.FromDto(microvolts_acDto)
    if err != nil {
        t.Errorf("FromDto() with MicrovoltAc returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_acResult.Convert(units.ElectricPotentialAcMicrovoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltAc = %v, want %v", converted, 100)
    }
    // Test MillivoltAc conversion
    millivolts_acDto := units.ElectricPotentialAcDto{
        Value: 100,
        Unit:  units.ElectricPotentialAcMillivoltAc,
    }
    
    var millivolts_acResult *units.ElectricPotentialAc
    millivolts_acResult, err = factory.FromDto(millivolts_acDto)
    if err != nil {
        t.Errorf("FromDto() with MillivoltAc returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_acResult.Convert(units.ElectricPotentialAcMillivoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltAc = %v, want %v", converted, 100)
    }
    // Test KilovoltAc conversion
    kilovolts_acDto := units.ElectricPotentialAcDto{
        Value: 100,
        Unit:  units.ElectricPotentialAcKilovoltAc,
    }
    
    var kilovolts_acResult *units.ElectricPotentialAc
    kilovolts_acResult, err = factory.FromDto(kilovolts_acDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltAc returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_acResult.Convert(units.ElectricPotentialAcKilovoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltAc = %v, want %v", converted, 100)
    }
    // Test MegavoltAc conversion
    megavolts_acDto := units.ElectricPotentialAcDto{
        Value: 100,
        Unit:  units.ElectricPotentialAcMegavoltAc,
    }
    
    var megavolts_acResult *units.ElectricPotentialAc
    megavolts_acResult, err = factory.FromDto(megavolts_acDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltAc returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_acResult.Convert(units.ElectricPotentialAcMegavoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltAc = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricPotentialAcDto{
        Value: 0,
        Unit:  units.ElectricPotentialAcVoltAc,
    }
    
    var zeroResult *units.ElectricPotentialAc
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricPotentialAcFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricPotentialAcFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "VoltAc"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "VoltAc"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricPotentialAcDto{
        Value: nanValue,
        Unit:  units.ElectricPotentialAcVoltAc,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with VoltAc unit
    volts_acJSON := []byte(`{"value": 100, "unit": "VoltAc"}`)
    volts_acResult, err := factory.FromDtoJSON(volts_acJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltAc unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_acResult.Convert(units.ElectricPotentialAcVoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltAc = %v, want %v", converted, 100)
    }
    // Test JSON with MicrovoltAc unit
    microvolts_acJSON := []byte(`{"value": 100, "unit": "MicrovoltAc"}`)
    microvolts_acResult, err := factory.FromDtoJSON(microvolts_acJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrovoltAc unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_acResult.Convert(units.ElectricPotentialAcMicrovoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltAc = %v, want %v", converted, 100)
    }
    // Test JSON with MillivoltAc unit
    millivolts_acJSON := []byte(`{"value": 100, "unit": "MillivoltAc"}`)
    millivolts_acResult, err := factory.FromDtoJSON(millivolts_acJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillivoltAc unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_acResult.Convert(units.ElectricPotentialAcMillivoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltAc = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltAc unit
    kilovolts_acJSON := []byte(`{"value": 100, "unit": "KilovoltAc"}`)
    kilovolts_acResult, err := factory.FromDtoJSON(kilovolts_acJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltAc unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_acResult.Convert(units.ElectricPotentialAcKilovoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltAc = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltAc unit
    megavolts_acJSON := []byte(`{"value": 100, "unit": "MegavoltAc"}`)
    megavolts_acResult, err := factory.FromDtoJSON(megavolts_acJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltAc unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_acResult.Convert(units.ElectricPotentialAcMegavoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltAc = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "VoltAc"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromVoltsAc function
func TestElectricPotentialAcFactory_FromVoltsAc(t *testing.T) {
    factory := units.ElectricPotentialAcFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltsAc(100)
    if err != nil {
        t.Errorf("FromVoltsAc() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialAcVoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltsAc() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltsAc(math.NaN())
    if err == nil {
        t.Error("FromVoltsAc() with NaN value should return error")
    }

    _, err = factory.FromVoltsAc(math.Inf(1))
    if err == nil {
        t.Error("FromVoltsAc() with +Inf value should return error")
    }

    _, err = factory.FromVoltsAc(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltsAc() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltsAc(0)
    if err != nil {
        t.Errorf("FromVoltsAc() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialAcVoltAc)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltsAc() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrovoltsAc function
func TestElectricPotentialAcFactory_FromMicrovoltsAc(t *testing.T) {
    factory := units.ElectricPotentialAcFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrovoltsAc(100)
    if err != nil {
        t.Errorf("FromMicrovoltsAc() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialAcMicrovoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrovoltsAc() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrovoltsAc(math.NaN())
    if err == nil {
        t.Error("FromMicrovoltsAc() with NaN value should return error")
    }

    _, err = factory.FromMicrovoltsAc(math.Inf(1))
    if err == nil {
        t.Error("FromMicrovoltsAc() with +Inf value should return error")
    }

    _, err = factory.FromMicrovoltsAc(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrovoltsAc() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrovoltsAc(0)
    if err != nil {
        t.Errorf("FromMicrovoltsAc() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialAcMicrovoltAc)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrovoltsAc() with zero value = %v, want 0", converted)
    }
}
// Test FromMillivoltsAc function
func TestElectricPotentialAcFactory_FromMillivoltsAc(t *testing.T) {
    factory := units.ElectricPotentialAcFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillivoltsAc(100)
    if err != nil {
        t.Errorf("FromMillivoltsAc() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialAcMillivoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillivoltsAc() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillivoltsAc(math.NaN())
    if err == nil {
        t.Error("FromMillivoltsAc() with NaN value should return error")
    }

    _, err = factory.FromMillivoltsAc(math.Inf(1))
    if err == nil {
        t.Error("FromMillivoltsAc() with +Inf value should return error")
    }

    _, err = factory.FromMillivoltsAc(math.Inf(-1))
    if err == nil {
        t.Error("FromMillivoltsAc() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillivoltsAc(0)
    if err != nil {
        t.Errorf("FromMillivoltsAc() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialAcMillivoltAc)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillivoltsAc() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltsAc function
func TestElectricPotentialAcFactory_FromKilovoltsAc(t *testing.T) {
    factory := units.ElectricPotentialAcFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltsAc(100)
    if err != nil {
        t.Errorf("FromKilovoltsAc() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialAcKilovoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltsAc() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltsAc(math.NaN())
    if err == nil {
        t.Error("FromKilovoltsAc() with NaN value should return error")
    }

    _, err = factory.FromKilovoltsAc(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltsAc() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltsAc(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltsAc() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltsAc(0)
    if err != nil {
        t.Errorf("FromKilovoltsAc() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialAcKilovoltAc)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltsAc() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltsAc function
func TestElectricPotentialAcFactory_FromMegavoltsAc(t *testing.T) {
    factory := units.ElectricPotentialAcFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltsAc(100)
    if err != nil {
        t.Errorf("FromMegavoltsAc() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialAcMegavoltAc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltsAc() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltsAc(math.NaN())
    if err == nil {
        t.Error("FromMegavoltsAc() with NaN value should return error")
    }

    _, err = factory.FromMegavoltsAc(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltsAc() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltsAc(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltsAc() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltsAc(0)
    if err != nil {
        t.Errorf("FromMegavoltsAc() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialAcMegavoltAc)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltsAc() with zero value = %v, want 0", converted)
    }
}

func TestElectricPotentialAcToString(t *testing.T) {
	factory := units.ElectricPotentialAcFactory{}
	a, err := factory.CreateElectricPotentialAc(45, units.ElectricPotentialAcVoltAc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricPotentialAcVoltAc, 2)
	expected := "45.00 " + units.GetElectricPotentialAcAbbreviation(units.ElectricPotentialAcVoltAc)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricPotentialAcVoltAc, -1)
	expected = "45 " + units.GetElectricPotentialAcAbbreviation(units.ElectricPotentialAcVoltAc)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricPotentialAc_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricPotentialAcFactory{}
	a1, _ := factory.CreateElectricPotentialAc(60, units.ElectricPotentialAcVoltAc)
	a2, _ := factory.CreateElectricPotentialAc(60, units.ElectricPotentialAcVoltAc)
	a3, _ := factory.CreateElectricPotentialAc(120, units.ElectricPotentialAcVoltAc)

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

func TestElectricPotentialAc_Arithmetic(t *testing.T) {
	factory := units.ElectricPotentialAcFactory{}
	a1, _ := factory.CreateElectricPotentialAc(30, units.ElectricPotentialAcVoltAc)
	a2, _ := factory.CreateElectricPotentialAc(45, units.ElectricPotentialAcVoltAc)

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


func TestGetElectricPotentialAcAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricPotentialAcUnits
        want string
    }{
        {
            name: "VoltAc abbreviation",
            unit: units.ElectricPotentialAcVoltAc,
            want: "Vac",
        },
        {
            name: "MicrovoltAc abbreviation",
            unit: units.ElectricPotentialAcMicrovoltAc,
            want: "μVac",
        },
        {
            name: "MillivoltAc abbreviation",
            unit: units.ElectricPotentialAcMillivoltAc,
            want: "mVac",
        },
        {
            name: "KilovoltAc abbreviation",
            unit: units.ElectricPotentialAcKilovoltAc,
            want: "kVac",
        },
        {
            name: "MegavoltAc abbreviation",
            unit: units.ElectricPotentialAcMegavoltAc,
            want: "MVac",
        },
        {
            name: "invalid unit",
            unit: units.ElectricPotentialAcUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricPotentialAcAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricPotentialAcAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricPotentialAc_String(t *testing.T) {
    factory := units.ElectricPotentialAcFactory{}
    
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
            unit, err := factory.CreateElectricPotentialAc(tt.value, units.ElectricPotentialAcVoltAc)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricPotentialAc.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestElectricPotentialAc_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.ElectricPotentialAcFactory{}

	_, err := uf.CreateElectricPotentialAc(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}