// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricPotentialDcDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltDc"}`
	
	factory := units.ElectricPotentialDcDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricPotentialDcVoltDc {
		t.Errorf("expected unit %v, got %v", units.ElectricPotentialDcVoltDc, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltDc"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricPotentialDcDto_ToJSON(t *testing.T) {
	dto := units.ElectricPotentialDcDto{
		Value: 45,
		Unit:  units.ElectricPotentialDcVoltDc,
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
	if result["unit"].(string) != string(units.ElectricPotentialDcVoltDc) {
		t.Errorf("expected unit %s, got %v", units.ElectricPotentialDcVoltDc, result["unit"])
	}
}

func TestNewElectricPotentialDc_InvalidValue(t *testing.T) {
	factory := units.ElectricPotentialDcFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricPotentialDc(math.NaN(), units.ElectricPotentialDcVoltDc)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricPotentialDc(math.Inf(1), units.ElectricPotentialDcVoltDc)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricPotentialDcConversions(t *testing.T) {
	factory := units.ElectricPotentialDcFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricPotentialDc(180, units.ElectricPotentialDcVoltDc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltsDc.
		// No expected conversion value provided for VoltsDc, verifying result is not NaN.
		result := a.VoltsDc()
		cacheResult := a.VoltsDc()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltsDc returned NaN")
		}
	}
	{
		// Test conversion to MicrovoltsDc.
		// No expected conversion value provided for MicrovoltsDc, verifying result is not NaN.
		result := a.MicrovoltsDc()
		cacheResult := a.MicrovoltsDc()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrovoltsDc returned NaN")
		}
	}
	{
		// Test conversion to MillivoltsDc.
		// No expected conversion value provided for MillivoltsDc, verifying result is not NaN.
		result := a.MillivoltsDc()
		cacheResult := a.MillivoltsDc()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillivoltsDc returned NaN")
		}
	}
	{
		// Test conversion to KilovoltsDc.
		// No expected conversion value provided for KilovoltsDc, verifying result is not NaN.
		result := a.KilovoltsDc()
		cacheResult := a.KilovoltsDc()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilovoltsDc returned NaN")
		}
	}
	{
		// Test conversion to MegavoltsDc.
		// No expected conversion value provided for MegavoltsDc, verifying result is not NaN.
		result := a.MegavoltsDc()
		cacheResult := a.MegavoltsDc()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegavoltsDc returned NaN")
		}
	}
}

func TestElectricPotentialDc_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricPotentialDcFactory{}
	a, err := factory.CreateElectricPotentialDc(90, units.ElectricPotentialDcVoltDc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricPotentialDcVoltDc {
		t.Errorf("expected default unit VoltDc, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricPotentialDcVoltDc
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricPotentialDcDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricPotentialDcVoltDc {
		t.Errorf("expected unit VoltDc, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricPotentialDcFactory_FromDto(t *testing.T) {
    factory := units.ElectricPotentialDcFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricPotentialDcDto{
        Value: 100,
        Unit:  units.ElectricPotentialDcVoltDc,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricPotentialDcDto{
        Value: math.NaN(),
        Unit:  units.ElectricPotentialDcVoltDc,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test VoltDc conversion
    volts_dcDto := units.ElectricPotentialDcDto{
        Value: 100,
        Unit:  units.ElectricPotentialDcVoltDc,
    }
    
    var volts_dcResult *units.ElectricPotentialDc
    volts_dcResult, err = factory.FromDto(volts_dcDto)
    if err != nil {
        t.Errorf("FromDto() with VoltDc returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_dcResult.Convert(units.ElectricPotentialDcVoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltDc = %v, want %v", converted, 100)
    }
    // Test MicrovoltDc conversion
    microvolts_dcDto := units.ElectricPotentialDcDto{
        Value: 100,
        Unit:  units.ElectricPotentialDcMicrovoltDc,
    }
    
    var microvolts_dcResult *units.ElectricPotentialDc
    microvolts_dcResult, err = factory.FromDto(microvolts_dcDto)
    if err != nil {
        t.Errorf("FromDto() with MicrovoltDc returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_dcResult.Convert(units.ElectricPotentialDcMicrovoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltDc = %v, want %v", converted, 100)
    }
    // Test MillivoltDc conversion
    millivolts_dcDto := units.ElectricPotentialDcDto{
        Value: 100,
        Unit:  units.ElectricPotentialDcMillivoltDc,
    }
    
    var millivolts_dcResult *units.ElectricPotentialDc
    millivolts_dcResult, err = factory.FromDto(millivolts_dcDto)
    if err != nil {
        t.Errorf("FromDto() with MillivoltDc returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_dcResult.Convert(units.ElectricPotentialDcMillivoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltDc = %v, want %v", converted, 100)
    }
    // Test KilovoltDc conversion
    kilovolts_dcDto := units.ElectricPotentialDcDto{
        Value: 100,
        Unit:  units.ElectricPotentialDcKilovoltDc,
    }
    
    var kilovolts_dcResult *units.ElectricPotentialDc
    kilovolts_dcResult, err = factory.FromDto(kilovolts_dcDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltDc returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_dcResult.Convert(units.ElectricPotentialDcKilovoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltDc = %v, want %v", converted, 100)
    }
    // Test MegavoltDc conversion
    megavolts_dcDto := units.ElectricPotentialDcDto{
        Value: 100,
        Unit:  units.ElectricPotentialDcMegavoltDc,
    }
    
    var megavolts_dcResult *units.ElectricPotentialDc
    megavolts_dcResult, err = factory.FromDto(megavolts_dcDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltDc returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_dcResult.Convert(units.ElectricPotentialDcMegavoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltDc = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricPotentialDcDto{
        Value: 0,
        Unit:  units.ElectricPotentialDcVoltDc,
    }
    
    var zeroResult *units.ElectricPotentialDc
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricPotentialDcFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricPotentialDcFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "VoltDc"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "VoltDc"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricPotentialDcDto{
        Value: nanValue,
        Unit:  units.ElectricPotentialDcVoltDc,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with VoltDc unit
    volts_dcJSON := []byte(`{"value": 100, "unit": "VoltDc"}`)
    volts_dcResult, err := factory.FromDtoJSON(volts_dcJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltDc unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_dcResult.Convert(units.ElectricPotentialDcVoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltDc = %v, want %v", converted, 100)
    }
    // Test JSON with MicrovoltDc unit
    microvolts_dcJSON := []byte(`{"value": 100, "unit": "MicrovoltDc"}`)
    microvolts_dcResult, err := factory.FromDtoJSON(microvolts_dcJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrovoltDc unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvolts_dcResult.Convert(units.ElectricPotentialDcMicrovoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrovoltDc = %v, want %v", converted, 100)
    }
    // Test JSON with MillivoltDc unit
    millivolts_dcJSON := []byte(`{"value": 100, "unit": "MillivoltDc"}`)
    millivolts_dcResult, err := factory.FromDtoJSON(millivolts_dcJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillivoltDc unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivolts_dcResult.Convert(units.ElectricPotentialDcMillivoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillivoltDc = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltDc unit
    kilovolts_dcJSON := []byte(`{"value": 100, "unit": "KilovoltDc"}`)
    kilovolts_dcResult, err := factory.FromDtoJSON(kilovolts_dcJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltDc unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovolts_dcResult.Convert(units.ElectricPotentialDcKilovoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltDc = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltDc unit
    megavolts_dcJSON := []byte(`{"value": 100, "unit": "MegavoltDc"}`)
    megavolts_dcResult, err := factory.FromDtoJSON(megavolts_dcJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltDc unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavolts_dcResult.Convert(units.ElectricPotentialDcMegavoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltDc = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "VoltDc"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromVoltsDc function
func TestElectricPotentialDcFactory_FromVoltsDc(t *testing.T) {
    factory := units.ElectricPotentialDcFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltsDc(100)
    if err != nil {
        t.Errorf("FromVoltsDc() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialDcVoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltsDc() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltsDc(math.NaN())
    if err == nil {
        t.Error("FromVoltsDc() with NaN value should return error")
    }

    _, err = factory.FromVoltsDc(math.Inf(1))
    if err == nil {
        t.Error("FromVoltsDc() with +Inf value should return error")
    }

    _, err = factory.FromVoltsDc(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltsDc() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltsDc(0)
    if err != nil {
        t.Errorf("FromVoltsDc() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialDcVoltDc)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltsDc() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrovoltsDc function
func TestElectricPotentialDcFactory_FromMicrovoltsDc(t *testing.T) {
    factory := units.ElectricPotentialDcFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrovoltsDc(100)
    if err != nil {
        t.Errorf("FromMicrovoltsDc() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialDcMicrovoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrovoltsDc() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrovoltsDc(math.NaN())
    if err == nil {
        t.Error("FromMicrovoltsDc() with NaN value should return error")
    }

    _, err = factory.FromMicrovoltsDc(math.Inf(1))
    if err == nil {
        t.Error("FromMicrovoltsDc() with +Inf value should return error")
    }

    _, err = factory.FromMicrovoltsDc(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrovoltsDc() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrovoltsDc(0)
    if err != nil {
        t.Errorf("FromMicrovoltsDc() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialDcMicrovoltDc)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrovoltsDc() with zero value = %v, want 0", converted)
    }
}
// Test FromMillivoltsDc function
func TestElectricPotentialDcFactory_FromMillivoltsDc(t *testing.T) {
    factory := units.ElectricPotentialDcFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillivoltsDc(100)
    if err != nil {
        t.Errorf("FromMillivoltsDc() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialDcMillivoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillivoltsDc() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillivoltsDc(math.NaN())
    if err == nil {
        t.Error("FromMillivoltsDc() with NaN value should return error")
    }

    _, err = factory.FromMillivoltsDc(math.Inf(1))
    if err == nil {
        t.Error("FromMillivoltsDc() with +Inf value should return error")
    }

    _, err = factory.FromMillivoltsDc(math.Inf(-1))
    if err == nil {
        t.Error("FromMillivoltsDc() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillivoltsDc(0)
    if err != nil {
        t.Errorf("FromMillivoltsDc() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialDcMillivoltDc)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillivoltsDc() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltsDc function
func TestElectricPotentialDcFactory_FromKilovoltsDc(t *testing.T) {
    factory := units.ElectricPotentialDcFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltsDc(100)
    if err != nil {
        t.Errorf("FromKilovoltsDc() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialDcKilovoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltsDc() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltsDc(math.NaN())
    if err == nil {
        t.Error("FromKilovoltsDc() with NaN value should return error")
    }

    _, err = factory.FromKilovoltsDc(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltsDc() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltsDc(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltsDc() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltsDc(0)
    if err != nil {
        t.Errorf("FromKilovoltsDc() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialDcKilovoltDc)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltsDc() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltsDc function
func TestElectricPotentialDcFactory_FromMegavoltsDc(t *testing.T) {
    factory := units.ElectricPotentialDcFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltsDc(100)
    if err != nil {
        t.Errorf("FromMegavoltsDc() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialDcMegavoltDc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltsDc() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltsDc(math.NaN())
    if err == nil {
        t.Error("FromMegavoltsDc() with NaN value should return error")
    }

    _, err = factory.FromMegavoltsDc(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltsDc() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltsDc(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltsDc() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltsDc(0)
    if err != nil {
        t.Errorf("FromMegavoltsDc() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialDcMegavoltDc)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltsDc() with zero value = %v, want 0", converted)
    }
}

func TestElectricPotentialDcToString(t *testing.T) {
	factory := units.ElectricPotentialDcFactory{}
	a, err := factory.CreateElectricPotentialDc(45, units.ElectricPotentialDcVoltDc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricPotentialDcVoltDc, 2)
	expected := "45.00 " + units.GetElectricPotentialDcAbbreviation(units.ElectricPotentialDcVoltDc)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricPotentialDcVoltDc, -1)
	expected = "45 " + units.GetElectricPotentialDcAbbreviation(units.ElectricPotentialDcVoltDc)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricPotentialDc_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricPotentialDcFactory{}
	a1, _ := factory.CreateElectricPotentialDc(60, units.ElectricPotentialDcVoltDc)
	a2, _ := factory.CreateElectricPotentialDc(60, units.ElectricPotentialDcVoltDc)
	a3, _ := factory.CreateElectricPotentialDc(120, units.ElectricPotentialDcVoltDc)

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

func TestElectricPotentialDc_Arithmetic(t *testing.T) {
	factory := units.ElectricPotentialDcFactory{}
	a1, _ := factory.CreateElectricPotentialDc(30, units.ElectricPotentialDcVoltDc)
	a2, _ := factory.CreateElectricPotentialDc(45, units.ElectricPotentialDcVoltDc)

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


func TestGetElectricPotentialDcAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricPotentialDcUnits
        want string
    }{
        {
            name: "VoltDc abbreviation",
            unit: units.ElectricPotentialDcVoltDc,
            want: "Vdc",
        },
        {
            name: "MicrovoltDc abbreviation",
            unit: units.ElectricPotentialDcMicrovoltDc,
            want: "μVdc",
        },
        {
            name: "MillivoltDc abbreviation",
            unit: units.ElectricPotentialDcMillivoltDc,
            want: "mVdc",
        },
        {
            name: "KilovoltDc abbreviation",
            unit: units.ElectricPotentialDcKilovoltDc,
            want: "kVdc",
        },
        {
            name: "MegavoltDc abbreviation",
            unit: units.ElectricPotentialDcMegavoltDc,
            want: "MVdc",
        },
        {
            name: "invalid unit",
            unit: units.ElectricPotentialDcUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricPotentialDcAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricPotentialDcAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricPotentialDc_String(t *testing.T) {
    factory := units.ElectricPotentialDcFactory{}
    
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
            unit, err := factory.CreateElectricPotentialDc(tt.value, units.ElectricPotentialDcVoltDc)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricPotentialDc.String() = %v, want %v", got, tt.want)
            }
        })
    }
}