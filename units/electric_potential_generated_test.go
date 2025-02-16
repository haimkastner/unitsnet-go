// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricPotentialDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Volt"}`
	
	factory := units.ElectricPotentialDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricPotentialVolt {
		t.Errorf("expected unit %v, got %v", units.ElectricPotentialVolt, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Volt"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricPotentialDto_ToJSON(t *testing.T) {
	dto := units.ElectricPotentialDto{
		Value: 45,
		Unit:  units.ElectricPotentialVolt,
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
	if result["unit"].(string) != string(units.ElectricPotentialVolt) {
		t.Errorf("expected unit %s, got %v", units.ElectricPotentialVolt, result["unit"])
	}
}

func TestNewElectricPotential_InvalidValue(t *testing.T) {
	factory := units.ElectricPotentialFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricPotential(math.NaN(), units.ElectricPotentialVolt)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricPotential(math.Inf(1), units.ElectricPotentialVolt)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricPotentialConversions(t *testing.T) {
	factory := units.ElectricPotentialFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricPotential(180, units.ElectricPotentialVolt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Volts.
		// No expected conversion value provided for Volts, verifying result is not NaN.
		result := a.Volts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Volts returned NaN")
		}
	}
	{
		// Test conversion to Nanovolts.
		// No expected conversion value provided for Nanovolts, verifying result is not NaN.
		result := a.Nanovolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanovolts returned NaN")
		}
	}
	{
		// Test conversion to Microvolts.
		// No expected conversion value provided for Microvolts, verifying result is not NaN.
		result := a.Microvolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microvolts returned NaN")
		}
	}
	{
		// Test conversion to Millivolts.
		// No expected conversion value provided for Millivolts, verifying result is not NaN.
		result := a.Millivolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millivolts returned NaN")
		}
	}
	{
		// Test conversion to Kilovolts.
		// No expected conversion value provided for Kilovolts, verifying result is not NaN.
		result := a.Kilovolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilovolts returned NaN")
		}
	}
	{
		// Test conversion to Megavolts.
		// No expected conversion value provided for Megavolts, verifying result is not NaN.
		result := a.Megavolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megavolts returned NaN")
		}
	}
}

func TestElectricPotential_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricPotentialFactory{}
	a, err := factory.CreateElectricPotential(90, units.ElectricPotentialVolt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricPotentialVolt {
		t.Errorf("expected default unit Volt, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricPotentialVolt
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricPotentialDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricPotentialVolt {
		t.Errorf("expected unit Volt, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricPotentialFactory_FromDto(t *testing.T) {
    factory := units.ElectricPotentialFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricPotentialDto{
        Value: 100,
        Unit:  units.ElectricPotentialVolt,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricPotentialDto{
        Value: math.NaN(),
        Unit:  units.ElectricPotentialVolt,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Volt conversion
    voltsDto := units.ElectricPotentialDto{
        Value: 100,
        Unit:  units.ElectricPotentialVolt,
    }
    
    var voltsResult *units.ElectricPotential
    voltsResult, err = factory.FromDto(voltsDto)
    if err != nil {
        t.Errorf("FromDto() with Volt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = voltsResult.Convert(units.ElectricPotentialVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Volt = %v, want %v", converted, 100)
    }
    // Test Nanovolt conversion
    nanovoltsDto := units.ElectricPotentialDto{
        Value: 100,
        Unit:  units.ElectricPotentialNanovolt,
    }
    
    var nanovoltsResult *units.ElectricPotential
    nanovoltsResult, err = factory.FromDto(nanovoltsDto)
    if err != nil {
        t.Errorf("FromDto() with Nanovolt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanovoltsResult.Convert(units.ElectricPotentialNanovolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanovolt = %v, want %v", converted, 100)
    }
    // Test Microvolt conversion
    microvoltsDto := units.ElectricPotentialDto{
        Value: 100,
        Unit:  units.ElectricPotentialMicrovolt,
    }
    
    var microvoltsResult *units.ElectricPotential
    microvoltsResult, err = factory.FromDto(microvoltsDto)
    if err != nil {
        t.Errorf("FromDto() with Microvolt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvoltsResult.Convert(units.ElectricPotentialMicrovolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microvolt = %v, want %v", converted, 100)
    }
    // Test Millivolt conversion
    millivoltsDto := units.ElectricPotentialDto{
        Value: 100,
        Unit:  units.ElectricPotentialMillivolt,
    }
    
    var millivoltsResult *units.ElectricPotential
    millivoltsResult, err = factory.FromDto(millivoltsDto)
    if err != nil {
        t.Errorf("FromDto() with Millivolt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivoltsResult.Convert(units.ElectricPotentialMillivolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millivolt = %v, want %v", converted, 100)
    }
    // Test Kilovolt conversion
    kilovoltsDto := units.ElectricPotentialDto{
        Value: 100,
        Unit:  units.ElectricPotentialKilovolt,
    }
    
    var kilovoltsResult *units.ElectricPotential
    kilovoltsResult, err = factory.FromDto(kilovoltsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilovolt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovoltsResult.Convert(units.ElectricPotentialKilovolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilovolt = %v, want %v", converted, 100)
    }
    // Test Megavolt conversion
    megavoltsDto := units.ElectricPotentialDto{
        Value: 100,
        Unit:  units.ElectricPotentialMegavolt,
    }
    
    var megavoltsResult *units.ElectricPotential
    megavoltsResult, err = factory.FromDto(megavoltsDto)
    if err != nil {
        t.Errorf("FromDto() with Megavolt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavoltsResult.Convert(units.ElectricPotentialMegavolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megavolt = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricPotentialDto{
        Value: 0,
        Unit:  units.ElectricPotentialVolt,
    }
    
    var zeroResult *units.ElectricPotential
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricPotentialFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricPotentialFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Volt"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Volt"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricPotentialDto{
        Value: nanValue,
        Unit:  units.ElectricPotentialVolt,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Volt unit
    voltsJSON := []byte(`{"value": 100, "unit": "Volt"}`)
    voltsResult, err := factory.FromDtoJSON(voltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Volt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = voltsResult.Convert(units.ElectricPotentialVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Volt = %v, want %v", converted, 100)
    }
    // Test JSON with Nanovolt unit
    nanovoltsJSON := []byte(`{"value": 100, "unit": "Nanovolt"}`)
    nanovoltsResult, err := factory.FromDtoJSON(nanovoltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanovolt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanovoltsResult.Convert(units.ElectricPotentialNanovolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanovolt = %v, want %v", converted, 100)
    }
    // Test JSON with Microvolt unit
    microvoltsJSON := []byte(`{"value": 100, "unit": "Microvolt"}`)
    microvoltsResult, err := factory.FromDtoJSON(microvoltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microvolt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microvoltsResult.Convert(units.ElectricPotentialMicrovolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microvolt = %v, want %v", converted, 100)
    }
    // Test JSON with Millivolt unit
    millivoltsJSON := []byte(`{"value": 100, "unit": "Millivolt"}`)
    millivoltsResult, err := factory.FromDtoJSON(millivoltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millivolt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millivoltsResult.Convert(units.ElectricPotentialMillivolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millivolt = %v, want %v", converted, 100)
    }
    // Test JSON with Kilovolt unit
    kilovoltsJSON := []byte(`{"value": 100, "unit": "Kilovolt"}`)
    kilovoltsResult, err := factory.FromDtoJSON(kilovoltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilovolt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovoltsResult.Convert(units.ElectricPotentialKilovolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilovolt = %v, want %v", converted, 100)
    }
    // Test JSON with Megavolt unit
    megavoltsJSON := []byte(`{"value": 100, "unit": "Megavolt"}`)
    megavoltsResult, err := factory.FromDtoJSON(megavoltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megavolt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavoltsResult.Convert(units.ElectricPotentialMegavolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megavolt = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Volt"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromVolts function
func TestElectricPotentialFactory_FromVolts(t *testing.T) {
    factory := units.ElectricPotentialFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVolts(100)
    if err != nil {
        t.Errorf("FromVolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVolts(math.NaN())
    if err == nil {
        t.Error("FromVolts() with NaN value should return error")
    }

    _, err = factory.FromVolts(math.Inf(1))
    if err == nil {
        t.Error("FromVolts() with +Inf value should return error")
    }

    _, err = factory.FromVolts(math.Inf(-1))
    if err == nil {
        t.Error("FromVolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVolts(0)
    if err != nil {
        t.Errorf("FromVolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialVolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVolts() with zero value = %v, want 0", converted)
    }
}
// Test FromNanovolts function
func TestElectricPotentialFactory_FromNanovolts(t *testing.T) {
    factory := units.ElectricPotentialFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanovolts(100)
    if err != nil {
        t.Errorf("FromNanovolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialNanovolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanovolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanovolts(math.NaN())
    if err == nil {
        t.Error("FromNanovolts() with NaN value should return error")
    }

    _, err = factory.FromNanovolts(math.Inf(1))
    if err == nil {
        t.Error("FromNanovolts() with +Inf value should return error")
    }

    _, err = factory.FromNanovolts(math.Inf(-1))
    if err == nil {
        t.Error("FromNanovolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanovolts(0)
    if err != nil {
        t.Errorf("FromNanovolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialNanovolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanovolts() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrovolts function
func TestElectricPotentialFactory_FromMicrovolts(t *testing.T) {
    factory := units.ElectricPotentialFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrovolts(100)
    if err != nil {
        t.Errorf("FromMicrovolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialMicrovolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrovolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrovolts(math.NaN())
    if err == nil {
        t.Error("FromMicrovolts() with NaN value should return error")
    }

    _, err = factory.FromMicrovolts(math.Inf(1))
    if err == nil {
        t.Error("FromMicrovolts() with +Inf value should return error")
    }

    _, err = factory.FromMicrovolts(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrovolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrovolts(0)
    if err != nil {
        t.Errorf("FromMicrovolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialMicrovolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrovolts() with zero value = %v, want 0", converted)
    }
}
// Test FromMillivolts function
func TestElectricPotentialFactory_FromMillivolts(t *testing.T) {
    factory := units.ElectricPotentialFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillivolts(100)
    if err != nil {
        t.Errorf("FromMillivolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialMillivolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillivolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillivolts(math.NaN())
    if err == nil {
        t.Error("FromMillivolts() with NaN value should return error")
    }

    _, err = factory.FromMillivolts(math.Inf(1))
    if err == nil {
        t.Error("FromMillivolts() with +Inf value should return error")
    }

    _, err = factory.FromMillivolts(math.Inf(-1))
    if err == nil {
        t.Error("FromMillivolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillivolts(0)
    if err != nil {
        t.Errorf("FromMillivolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialMillivolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillivolts() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovolts function
func TestElectricPotentialFactory_FromKilovolts(t *testing.T) {
    factory := units.ElectricPotentialFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovolts(100)
    if err != nil {
        t.Errorf("FromKilovolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialKilovolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovolts(math.NaN())
    if err == nil {
        t.Error("FromKilovolts() with NaN value should return error")
    }

    _, err = factory.FromKilovolts(math.Inf(1))
    if err == nil {
        t.Error("FromKilovolts() with +Inf value should return error")
    }

    _, err = factory.FromKilovolts(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovolts(0)
    if err != nil {
        t.Errorf("FromKilovolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialKilovolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovolts() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavolts function
func TestElectricPotentialFactory_FromMegavolts(t *testing.T) {
    factory := units.ElectricPotentialFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavolts(100)
    if err != nil {
        t.Errorf("FromMegavolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricPotentialMegavolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavolts(math.NaN())
    if err == nil {
        t.Error("FromMegavolts() with NaN value should return error")
    }

    _, err = factory.FromMegavolts(math.Inf(1))
    if err == nil {
        t.Error("FromMegavolts() with +Inf value should return error")
    }

    _, err = factory.FromMegavolts(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavolts(0)
    if err != nil {
        t.Errorf("FromMegavolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricPotentialMegavolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavolts() with zero value = %v, want 0", converted)
    }
}

func TestElectricPotentialToString(t *testing.T) {
	factory := units.ElectricPotentialFactory{}
	a, err := factory.CreateElectricPotential(45, units.ElectricPotentialVolt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricPotentialVolt, 2)
	expected := "45.00 " + units.GetElectricPotentialAbbreviation(units.ElectricPotentialVolt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricPotentialVolt, -1)
	expected = "45 " + units.GetElectricPotentialAbbreviation(units.ElectricPotentialVolt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricPotential_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricPotentialFactory{}
	a1, _ := factory.CreateElectricPotential(60, units.ElectricPotentialVolt)
	a2, _ := factory.CreateElectricPotential(60, units.ElectricPotentialVolt)
	a3, _ := factory.CreateElectricPotential(120, units.ElectricPotentialVolt)

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

func TestElectricPotential_Arithmetic(t *testing.T) {
	factory := units.ElectricPotentialFactory{}
	a1, _ := factory.CreateElectricPotential(30, units.ElectricPotentialVolt)
	a2, _ := factory.CreateElectricPotential(45, units.ElectricPotentialVolt)

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