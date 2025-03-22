// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestCapacitanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Farad"}`
	
	factory := units.CapacitanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.CapacitanceFarad {
		t.Errorf("expected unit %v, got %v", units.CapacitanceFarad, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Farad"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestCapacitanceDto_ToJSON(t *testing.T) {
	dto := units.CapacitanceDto{
		Value: 45,
		Unit:  units.CapacitanceFarad,
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
	if result["unit"].(string) != string(units.CapacitanceFarad) {
		t.Errorf("expected unit %s, got %v", units.CapacitanceFarad, result["unit"])
	}
}

func TestNewCapacitance_InvalidValue(t *testing.T) {
	factory := units.CapacitanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateCapacitance(math.NaN(), units.CapacitanceFarad)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateCapacitance(math.Inf(1), units.CapacitanceFarad)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestCapacitanceConversions(t *testing.T) {
	factory := units.CapacitanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateCapacitance(180, units.CapacitanceFarad)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Farads.
		// No expected conversion value provided for Farads, verifying result is not NaN.
		result := a.Farads()
		cacheResult := a.Farads()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Farads returned NaN")
		}
	}
	{
		// Test conversion to Picofarads.
		// No expected conversion value provided for Picofarads, verifying result is not NaN.
		result := a.Picofarads()
		cacheResult := a.Picofarads()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Picofarads returned NaN")
		}
	}
	{
		// Test conversion to Nanofarads.
		// No expected conversion value provided for Nanofarads, verifying result is not NaN.
		result := a.Nanofarads()
		cacheResult := a.Nanofarads()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanofarads returned NaN")
		}
	}
	{
		// Test conversion to Microfarads.
		// No expected conversion value provided for Microfarads, verifying result is not NaN.
		result := a.Microfarads()
		cacheResult := a.Microfarads()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microfarads returned NaN")
		}
	}
	{
		// Test conversion to Millifarads.
		// No expected conversion value provided for Millifarads, verifying result is not NaN.
		result := a.Millifarads()
		cacheResult := a.Millifarads()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millifarads returned NaN")
		}
	}
	{
		// Test conversion to Kilofarads.
		// No expected conversion value provided for Kilofarads, verifying result is not NaN.
		result := a.Kilofarads()
		cacheResult := a.Kilofarads()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilofarads returned NaN")
		}
	}
	{
		// Test conversion to Megafarads.
		// No expected conversion value provided for Megafarads, verifying result is not NaN.
		result := a.Megafarads()
		cacheResult := a.Megafarads()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megafarads returned NaN")
		}
	}
}

func TestCapacitance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.CapacitanceFactory{}
	a, err := factory.CreateCapacitance(90, units.CapacitanceFarad)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.CapacitanceFarad {
		t.Errorf("expected default unit Farad, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.CapacitanceFarad
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.CapacitanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.CapacitanceFarad {
		t.Errorf("expected unit Farad, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestCapacitanceFactory_FromDto(t *testing.T) {
    factory := units.CapacitanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.CapacitanceDto{
        Value: 100,
        Unit:  units.CapacitanceFarad,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.CapacitanceDto{
        Value: math.NaN(),
        Unit:  units.CapacitanceFarad,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Farad conversion
    faradsDto := units.CapacitanceDto{
        Value: 100,
        Unit:  units.CapacitanceFarad,
    }
    
    var faradsResult *units.Capacitance
    faradsResult, err = factory.FromDto(faradsDto)
    if err != nil {
        t.Errorf("FromDto() with Farad returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = faradsResult.Convert(units.CapacitanceFarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Farad = %v, want %v", converted, 100)
    }
    // Test Picofarad conversion
    picofaradsDto := units.CapacitanceDto{
        Value: 100,
        Unit:  units.CapacitancePicofarad,
    }
    
    var picofaradsResult *units.Capacitance
    picofaradsResult, err = factory.FromDto(picofaradsDto)
    if err != nil {
        t.Errorf("FromDto() with Picofarad returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picofaradsResult.Convert(units.CapacitancePicofarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picofarad = %v, want %v", converted, 100)
    }
    // Test Nanofarad conversion
    nanofaradsDto := units.CapacitanceDto{
        Value: 100,
        Unit:  units.CapacitanceNanofarad,
    }
    
    var nanofaradsResult *units.Capacitance
    nanofaradsResult, err = factory.FromDto(nanofaradsDto)
    if err != nil {
        t.Errorf("FromDto() with Nanofarad returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanofaradsResult.Convert(units.CapacitanceNanofarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanofarad = %v, want %v", converted, 100)
    }
    // Test Microfarad conversion
    microfaradsDto := units.CapacitanceDto{
        Value: 100,
        Unit:  units.CapacitanceMicrofarad,
    }
    
    var microfaradsResult *units.Capacitance
    microfaradsResult, err = factory.FromDto(microfaradsDto)
    if err != nil {
        t.Errorf("FromDto() with Microfarad returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microfaradsResult.Convert(units.CapacitanceMicrofarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microfarad = %v, want %v", converted, 100)
    }
    // Test Millifarad conversion
    millifaradsDto := units.CapacitanceDto{
        Value: 100,
        Unit:  units.CapacitanceMillifarad,
    }
    
    var millifaradsResult *units.Capacitance
    millifaradsResult, err = factory.FromDto(millifaradsDto)
    if err != nil {
        t.Errorf("FromDto() with Millifarad returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millifaradsResult.Convert(units.CapacitanceMillifarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millifarad = %v, want %v", converted, 100)
    }
    // Test Kilofarad conversion
    kilofaradsDto := units.CapacitanceDto{
        Value: 100,
        Unit:  units.CapacitanceKilofarad,
    }
    
    var kilofaradsResult *units.Capacitance
    kilofaradsResult, err = factory.FromDto(kilofaradsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilofarad returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilofaradsResult.Convert(units.CapacitanceKilofarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilofarad = %v, want %v", converted, 100)
    }
    // Test Megafarad conversion
    megafaradsDto := units.CapacitanceDto{
        Value: 100,
        Unit:  units.CapacitanceMegafarad,
    }
    
    var megafaradsResult *units.Capacitance
    megafaradsResult, err = factory.FromDto(megafaradsDto)
    if err != nil {
        t.Errorf("FromDto() with Megafarad returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megafaradsResult.Convert(units.CapacitanceMegafarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megafarad = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.CapacitanceDto{
        Value: 0,
        Unit:  units.CapacitanceFarad,
    }
    
    var zeroResult *units.Capacitance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestCapacitanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.CapacitanceFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Farad"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Farad"}`)
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
    nanJSON, _ := json.Marshal(units.CapacitanceDto{
        Value: nanValue,
        Unit:  units.CapacitanceFarad,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Farad unit
    faradsJSON := []byte(`{"value": 100, "unit": "Farad"}`)
    faradsResult, err := factory.FromDtoJSON(faradsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Farad unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = faradsResult.Convert(units.CapacitanceFarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Farad = %v, want %v", converted, 100)
    }
    // Test JSON with Picofarad unit
    picofaradsJSON := []byte(`{"value": 100, "unit": "Picofarad"}`)
    picofaradsResult, err := factory.FromDtoJSON(picofaradsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Picofarad unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picofaradsResult.Convert(units.CapacitancePicofarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picofarad = %v, want %v", converted, 100)
    }
    // Test JSON with Nanofarad unit
    nanofaradsJSON := []byte(`{"value": 100, "unit": "Nanofarad"}`)
    nanofaradsResult, err := factory.FromDtoJSON(nanofaradsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanofarad unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanofaradsResult.Convert(units.CapacitanceNanofarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanofarad = %v, want %v", converted, 100)
    }
    // Test JSON with Microfarad unit
    microfaradsJSON := []byte(`{"value": 100, "unit": "Microfarad"}`)
    microfaradsResult, err := factory.FromDtoJSON(microfaradsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microfarad unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microfaradsResult.Convert(units.CapacitanceMicrofarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microfarad = %v, want %v", converted, 100)
    }
    // Test JSON with Millifarad unit
    millifaradsJSON := []byte(`{"value": 100, "unit": "Millifarad"}`)
    millifaradsResult, err := factory.FromDtoJSON(millifaradsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millifarad unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millifaradsResult.Convert(units.CapacitanceMillifarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millifarad = %v, want %v", converted, 100)
    }
    // Test JSON with Kilofarad unit
    kilofaradsJSON := []byte(`{"value": 100, "unit": "Kilofarad"}`)
    kilofaradsResult, err := factory.FromDtoJSON(kilofaradsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilofarad unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilofaradsResult.Convert(units.CapacitanceKilofarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilofarad = %v, want %v", converted, 100)
    }
    // Test JSON with Megafarad unit
    megafaradsJSON := []byte(`{"value": 100, "unit": "Megafarad"}`)
    megafaradsResult, err := factory.FromDtoJSON(megafaradsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megafarad unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megafaradsResult.Convert(units.CapacitanceMegafarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megafarad = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Farad"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromFarads function
func TestCapacitanceFactory_FromFarads(t *testing.T) {
    factory := units.CapacitanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFarads(100)
    if err != nil {
        t.Errorf("FromFarads() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CapacitanceFarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFarads() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFarads(math.NaN())
    if err == nil {
        t.Error("FromFarads() with NaN value should return error")
    }

    _, err = factory.FromFarads(math.Inf(1))
    if err == nil {
        t.Error("FromFarads() with +Inf value should return error")
    }

    _, err = factory.FromFarads(math.Inf(-1))
    if err == nil {
        t.Error("FromFarads() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFarads(0)
    if err != nil {
        t.Errorf("FromFarads() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CapacitanceFarad)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFarads() with zero value = %v, want 0", converted)
    }
}
// Test FromPicofarads function
func TestCapacitanceFactory_FromPicofarads(t *testing.T) {
    factory := units.CapacitanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicofarads(100)
    if err != nil {
        t.Errorf("FromPicofarads() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CapacitancePicofarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicofarads() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicofarads(math.NaN())
    if err == nil {
        t.Error("FromPicofarads() with NaN value should return error")
    }

    _, err = factory.FromPicofarads(math.Inf(1))
    if err == nil {
        t.Error("FromPicofarads() with +Inf value should return error")
    }

    _, err = factory.FromPicofarads(math.Inf(-1))
    if err == nil {
        t.Error("FromPicofarads() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicofarads(0)
    if err != nil {
        t.Errorf("FromPicofarads() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CapacitancePicofarad)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicofarads() with zero value = %v, want 0", converted)
    }
}
// Test FromNanofarads function
func TestCapacitanceFactory_FromNanofarads(t *testing.T) {
    factory := units.CapacitanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanofarads(100)
    if err != nil {
        t.Errorf("FromNanofarads() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CapacitanceNanofarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanofarads() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanofarads(math.NaN())
    if err == nil {
        t.Error("FromNanofarads() with NaN value should return error")
    }

    _, err = factory.FromNanofarads(math.Inf(1))
    if err == nil {
        t.Error("FromNanofarads() with +Inf value should return error")
    }

    _, err = factory.FromNanofarads(math.Inf(-1))
    if err == nil {
        t.Error("FromNanofarads() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanofarads(0)
    if err != nil {
        t.Errorf("FromNanofarads() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CapacitanceNanofarad)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanofarads() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrofarads function
func TestCapacitanceFactory_FromMicrofarads(t *testing.T) {
    factory := units.CapacitanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrofarads(100)
    if err != nil {
        t.Errorf("FromMicrofarads() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CapacitanceMicrofarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrofarads() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrofarads(math.NaN())
    if err == nil {
        t.Error("FromMicrofarads() with NaN value should return error")
    }

    _, err = factory.FromMicrofarads(math.Inf(1))
    if err == nil {
        t.Error("FromMicrofarads() with +Inf value should return error")
    }

    _, err = factory.FromMicrofarads(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrofarads() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrofarads(0)
    if err != nil {
        t.Errorf("FromMicrofarads() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CapacitanceMicrofarad)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrofarads() with zero value = %v, want 0", converted)
    }
}
// Test FromMillifarads function
func TestCapacitanceFactory_FromMillifarads(t *testing.T) {
    factory := units.CapacitanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillifarads(100)
    if err != nil {
        t.Errorf("FromMillifarads() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CapacitanceMillifarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillifarads() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillifarads(math.NaN())
    if err == nil {
        t.Error("FromMillifarads() with NaN value should return error")
    }

    _, err = factory.FromMillifarads(math.Inf(1))
    if err == nil {
        t.Error("FromMillifarads() with +Inf value should return error")
    }

    _, err = factory.FromMillifarads(math.Inf(-1))
    if err == nil {
        t.Error("FromMillifarads() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillifarads(0)
    if err != nil {
        t.Errorf("FromMillifarads() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CapacitanceMillifarad)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillifarads() with zero value = %v, want 0", converted)
    }
}
// Test FromKilofarads function
func TestCapacitanceFactory_FromKilofarads(t *testing.T) {
    factory := units.CapacitanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilofarads(100)
    if err != nil {
        t.Errorf("FromKilofarads() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CapacitanceKilofarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilofarads() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilofarads(math.NaN())
    if err == nil {
        t.Error("FromKilofarads() with NaN value should return error")
    }

    _, err = factory.FromKilofarads(math.Inf(1))
    if err == nil {
        t.Error("FromKilofarads() with +Inf value should return error")
    }

    _, err = factory.FromKilofarads(math.Inf(-1))
    if err == nil {
        t.Error("FromKilofarads() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilofarads(0)
    if err != nil {
        t.Errorf("FromKilofarads() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CapacitanceKilofarad)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilofarads() with zero value = %v, want 0", converted)
    }
}
// Test FromMegafarads function
func TestCapacitanceFactory_FromMegafarads(t *testing.T) {
    factory := units.CapacitanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegafarads(100)
    if err != nil {
        t.Errorf("FromMegafarads() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CapacitanceMegafarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegafarads() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegafarads(math.NaN())
    if err == nil {
        t.Error("FromMegafarads() with NaN value should return error")
    }

    _, err = factory.FromMegafarads(math.Inf(1))
    if err == nil {
        t.Error("FromMegafarads() with +Inf value should return error")
    }

    _, err = factory.FromMegafarads(math.Inf(-1))
    if err == nil {
        t.Error("FromMegafarads() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegafarads(0)
    if err != nil {
        t.Errorf("FromMegafarads() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CapacitanceMegafarad)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegafarads() with zero value = %v, want 0", converted)
    }
}

func TestCapacitanceToString(t *testing.T) {
	factory := units.CapacitanceFactory{}
	a, err := factory.CreateCapacitance(45, units.CapacitanceFarad)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.CapacitanceFarad, 2)
	expected := "45.00 " + units.GetCapacitanceAbbreviation(units.CapacitanceFarad)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.CapacitanceFarad, -1)
	expected = "45 " + units.GetCapacitanceAbbreviation(units.CapacitanceFarad)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestCapacitance_EqualityAndComparison(t *testing.T) {
	factory := units.CapacitanceFactory{}
	a1, _ := factory.CreateCapacitance(60, units.CapacitanceFarad)
	a2, _ := factory.CreateCapacitance(60, units.CapacitanceFarad)
	a3, _ := factory.CreateCapacitance(120, units.CapacitanceFarad)

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

func TestCapacitance_Arithmetic(t *testing.T) {
	factory := units.CapacitanceFactory{}
	a1, _ := factory.CreateCapacitance(30, units.CapacitanceFarad)
	a2, _ := factory.CreateCapacitance(45, units.CapacitanceFarad)

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


func TestGetCapacitanceAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.CapacitanceUnits
        want string
    }{
        {
            name: "Farad abbreviation",
            unit: units.CapacitanceFarad,
            want: "F",
        },
        {
            name: "Picofarad abbreviation",
            unit: units.CapacitancePicofarad,
            want: "pF",
        },
        {
            name: "Nanofarad abbreviation",
            unit: units.CapacitanceNanofarad,
            want: "nF",
        },
        {
            name: "Microfarad abbreviation",
            unit: units.CapacitanceMicrofarad,
            want: "μF",
        },
        {
            name: "Millifarad abbreviation",
            unit: units.CapacitanceMillifarad,
            want: "mF",
        },
        {
            name: "Kilofarad abbreviation",
            unit: units.CapacitanceKilofarad,
            want: "kF",
        },
        {
            name: "Megafarad abbreviation",
            unit: units.CapacitanceMegafarad,
            want: "MF",
        },
        {
            name: "invalid unit",
            unit: units.CapacitanceUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetCapacitanceAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetCapacitanceAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestCapacitance_String(t *testing.T) {
    factory := units.CapacitanceFactory{}
    
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
            unit, err := factory.CreateCapacitance(tt.value, units.CapacitanceFarad)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Capacitance.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestCapacitance_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.CapacitanceFactory{}

	_, err := uf.CreateCapacitance(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}