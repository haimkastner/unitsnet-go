// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricCurrentDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Ampere"}`
	
	factory := units.ElectricCurrentDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricCurrentAmpere {
		t.Errorf("expected unit %v, got %v", units.ElectricCurrentAmpere, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Ampere"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricCurrentDto_ToJSON(t *testing.T) {
	dto := units.ElectricCurrentDto{
		Value: 45,
		Unit:  units.ElectricCurrentAmpere,
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
	if result["unit"].(string) != string(units.ElectricCurrentAmpere) {
		t.Errorf("expected unit %s, got %v", units.ElectricCurrentAmpere, result["unit"])
	}
}

func TestNewElectricCurrent_InvalidValue(t *testing.T) {
	factory := units.ElectricCurrentFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricCurrent(math.NaN(), units.ElectricCurrentAmpere)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricCurrent(math.Inf(1), units.ElectricCurrentAmpere)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricCurrentConversions(t *testing.T) {
	factory := units.ElectricCurrentFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricCurrent(180, units.ElectricCurrentAmpere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Amperes.
		// No expected conversion value provided for Amperes, verifying result is not NaN.
		result := a.Amperes()
		cacheResult := a.Amperes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Amperes returned NaN")
		}
	}
	{
		// Test conversion to Femtoamperes.
		// No expected conversion value provided for Femtoamperes, verifying result is not NaN.
		result := a.Femtoamperes()
		cacheResult := a.Femtoamperes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Femtoamperes returned NaN")
		}
	}
	{
		// Test conversion to Picoamperes.
		// No expected conversion value provided for Picoamperes, verifying result is not NaN.
		result := a.Picoamperes()
		cacheResult := a.Picoamperes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Picoamperes returned NaN")
		}
	}
	{
		// Test conversion to Nanoamperes.
		// No expected conversion value provided for Nanoamperes, verifying result is not NaN.
		result := a.Nanoamperes()
		cacheResult := a.Nanoamperes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanoamperes returned NaN")
		}
	}
	{
		// Test conversion to Microamperes.
		// No expected conversion value provided for Microamperes, verifying result is not NaN.
		result := a.Microamperes()
		cacheResult := a.Microamperes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microamperes returned NaN")
		}
	}
	{
		// Test conversion to Milliamperes.
		// No expected conversion value provided for Milliamperes, verifying result is not NaN.
		result := a.Milliamperes()
		cacheResult := a.Milliamperes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Milliamperes returned NaN")
		}
	}
	{
		// Test conversion to Centiamperes.
		// No expected conversion value provided for Centiamperes, verifying result is not NaN.
		result := a.Centiamperes()
		cacheResult := a.Centiamperes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Centiamperes returned NaN")
		}
	}
	{
		// Test conversion to Kiloamperes.
		// No expected conversion value provided for Kiloamperes, verifying result is not NaN.
		result := a.Kiloamperes()
		cacheResult := a.Kiloamperes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kiloamperes returned NaN")
		}
	}
	{
		// Test conversion to Megaamperes.
		// No expected conversion value provided for Megaamperes, verifying result is not NaN.
		result := a.Megaamperes()
		cacheResult := a.Megaamperes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megaamperes returned NaN")
		}
	}
}

func TestElectricCurrent_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricCurrentFactory{}
	a, err := factory.CreateElectricCurrent(90, units.ElectricCurrentAmpere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricCurrentAmpere {
		t.Errorf("expected default unit Ampere, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricCurrentAmpere
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricCurrentDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricCurrentAmpere {
		t.Errorf("expected unit Ampere, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricCurrentFactory_FromDto(t *testing.T) {
    factory := units.ElectricCurrentFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricCurrentDto{
        Value: 100,
        Unit:  units.ElectricCurrentAmpere,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricCurrentDto{
        Value: math.NaN(),
        Unit:  units.ElectricCurrentAmpere,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Ampere conversion
    amperesDto := units.ElectricCurrentDto{
        Value: 100,
        Unit:  units.ElectricCurrentAmpere,
    }
    
    var amperesResult *units.ElectricCurrent
    amperesResult, err = factory.FromDto(amperesDto)
    if err != nil {
        t.Errorf("FromDto() with Ampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperesResult.Convert(units.ElectricCurrentAmpere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Ampere = %v, want %v", converted, 100)
    }
    // Test Femtoampere conversion
    femtoamperesDto := units.ElectricCurrentDto{
        Value: 100,
        Unit:  units.ElectricCurrentFemtoampere,
    }
    
    var femtoamperesResult *units.ElectricCurrent
    femtoamperesResult, err = factory.FromDto(femtoamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Femtoampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtoamperesResult.Convert(units.ElectricCurrentFemtoampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtoampere = %v, want %v", converted, 100)
    }
    // Test Picoampere conversion
    picoamperesDto := units.ElectricCurrentDto{
        Value: 100,
        Unit:  units.ElectricCurrentPicoampere,
    }
    
    var picoamperesResult *units.ElectricCurrent
    picoamperesResult, err = factory.FromDto(picoamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Picoampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoamperesResult.Convert(units.ElectricCurrentPicoampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picoampere = %v, want %v", converted, 100)
    }
    // Test Nanoampere conversion
    nanoamperesDto := units.ElectricCurrentDto{
        Value: 100,
        Unit:  units.ElectricCurrentNanoampere,
    }
    
    var nanoamperesResult *units.ElectricCurrent
    nanoamperesResult, err = factory.FromDto(nanoamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Nanoampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoamperesResult.Convert(units.ElectricCurrentNanoampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanoampere = %v, want %v", converted, 100)
    }
    // Test Microampere conversion
    microamperesDto := units.ElectricCurrentDto{
        Value: 100,
        Unit:  units.ElectricCurrentMicroampere,
    }
    
    var microamperesResult *units.ElectricCurrent
    microamperesResult, err = factory.FromDto(microamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Microampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microamperesResult.Convert(units.ElectricCurrentMicroampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microampere = %v, want %v", converted, 100)
    }
    // Test Milliampere conversion
    milliamperesDto := units.ElectricCurrentDto{
        Value: 100,
        Unit:  units.ElectricCurrentMilliampere,
    }
    
    var milliamperesResult *units.ElectricCurrent
    milliamperesResult, err = factory.FromDto(milliamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Milliampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliamperesResult.Convert(units.ElectricCurrentMilliampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliampere = %v, want %v", converted, 100)
    }
    // Test Centiampere conversion
    centiamperesDto := units.ElectricCurrentDto{
        Value: 100,
        Unit:  units.ElectricCurrentCentiampere,
    }
    
    var centiamperesResult *units.ElectricCurrent
    centiamperesResult, err = factory.FromDto(centiamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Centiampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiamperesResult.Convert(units.ElectricCurrentCentiampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centiampere = %v, want %v", converted, 100)
    }
    // Test Kiloampere conversion
    kiloamperesDto := units.ElectricCurrentDto{
        Value: 100,
        Unit:  units.ElectricCurrentKiloampere,
    }
    
    var kiloamperesResult *units.ElectricCurrent
    kiloamperesResult, err = factory.FromDto(kiloamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Kiloampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloamperesResult.Convert(units.ElectricCurrentKiloampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kiloampere = %v, want %v", converted, 100)
    }
    // Test Megaampere conversion
    megaamperesDto := units.ElectricCurrentDto{
        Value: 100,
        Unit:  units.ElectricCurrentMegaampere,
    }
    
    var megaamperesResult *units.ElectricCurrent
    megaamperesResult, err = factory.FromDto(megaamperesDto)
    if err != nil {
        t.Errorf("FromDto() with Megaampere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaamperesResult.Convert(units.ElectricCurrentMegaampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megaampere = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricCurrentDto{
        Value: 0,
        Unit:  units.ElectricCurrentAmpere,
    }
    
    var zeroResult *units.ElectricCurrent
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricCurrentFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricCurrentFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Ampere"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Ampere"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricCurrentDto{
        Value: nanValue,
        Unit:  units.ElectricCurrentAmpere,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Ampere unit
    amperesJSON := []byte(`{"value": 100, "unit": "Ampere"}`)
    amperesResult, err := factory.FromDtoJSON(amperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Ampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperesResult.Convert(units.ElectricCurrentAmpere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Ampere = %v, want %v", converted, 100)
    }
    // Test JSON with Femtoampere unit
    femtoamperesJSON := []byte(`{"value": 100, "unit": "Femtoampere"}`)
    femtoamperesResult, err := factory.FromDtoJSON(femtoamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Femtoampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtoamperesResult.Convert(units.ElectricCurrentFemtoampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtoampere = %v, want %v", converted, 100)
    }
    // Test JSON with Picoampere unit
    picoamperesJSON := []byte(`{"value": 100, "unit": "Picoampere"}`)
    picoamperesResult, err := factory.FromDtoJSON(picoamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Picoampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoamperesResult.Convert(units.ElectricCurrentPicoampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picoampere = %v, want %v", converted, 100)
    }
    // Test JSON with Nanoampere unit
    nanoamperesJSON := []byte(`{"value": 100, "unit": "Nanoampere"}`)
    nanoamperesResult, err := factory.FromDtoJSON(nanoamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanoampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoamperesResult.Convert(units.ElectricCurrentNanoampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanoampere = %v, want %v", converted, 100)
    }
    // Test JSON with Microampere unit
    microamperesJSON := []byte(`{"value": 100, "unit": "Microampere"}`)
    microamperesResult, err := factory.FromDtoJSON(microamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microamperesResult.Convert(units.ElectricCurrentMicroampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microampere = %v, want %v", converted, 100)
    }
    // Test JSON with Milliampere unit
    milliamperesJSON := []byte(`{"value": 100, "unit": "Milliampere"}`)
    milliamperesResult, err := factory.FromDtoJSON(milliamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Milliampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliamperesResult.Convert(units.ElectricCurrentMilliampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliampere = %v, want %v", converted, 100)
    }
    // Test JSON with Centiampere unit
    centiamperesJSON := []byte(`{"value": 100, "unit": "Centiampere"}`)
    centiamperesResult, err := factory.FromDtoJSON(centiamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Centiampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiamperesResult.Convert(units.ElectricCurrentCentiampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centiampere = %v, want %v", converted, 100)
    }
    // Test JSON with Kiloampere unit
    kiloamperesJSON := []byte(`{"value": 100, "unit": "Kiloampere"}`)
    kiloamperesResult, err := factory.FromDtoJSON(kiloamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kiloampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloamperesResult.Convert(units.ElectricCurrentKiloampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kiloampere = %v, want %v", converted, 100)
    }
    // Test JSON with Megaampere unit
    megaamperesJSON := []byte(`{"value": 100, "unit": "Megaampere"}`)
    megaamperesResult, err := factory.FromDtoJSON(megaamperesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megaampere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaamperesResult.Convert(units.ElectricCurrentMegaampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megaampere = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Ampere"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromAmperes function
func TestElectricCurrentFactory_FromAmperes(t *testing.T) {
    factory := units.ElectricCurrentFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAmperes(100)
    if err != nil {
        t.Errorf("FromAmperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentAmpere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAmperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAmperes(math.NaN())
    if err == nil {
        t.Error("FromAmperes() with NaN value should return error")
    }

    _, err = factory.FromAmperes(math.Inf(1))
    if err == nil {
        t.Error("FromAmperes() with +Inf value should return error")
    }

    _, err = factory.FromAmperes(math.Inf(-1))
    if err == nil {
        t.Error("FromAmperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAmperes(0)
    if err != nil {
        t.Errorf("FromAmperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentAmpere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAmperes() with zero value = %v, want 0", converted)
    }
}
// Test FromFemtoamperes function
func TestElectricCurrentFactory_FromFemtoamperes(t *testing.T) {
    factory := units.ElectricCurrentFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFemtoamperes(100)
    if err != nil {
        t.Errorf("FromFemtoamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentFemtoampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFemtoamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFemtoamperes(math.NaN())
    if err == nil {
        t.Error("FromFemtoamperes() with NaN value should return error")
    }

    _, err = factory.FromFemtoamperes(math.Inf(1))
    if err == nil {
        t.Error("FromFemtoamperes() with +Inf value should return error")
    }

    _, err = factory.FromFemtoamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromFemtoamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFemtoamperes(0)
    if err != nil {
        t.Errorf("FromFemtoamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentFemtoampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFemtoamperes() with zero value = %v, want 0", converted)
    }
}
// Test FromPicoamperes function
func TestElectricCurrentFactory_FromPicoamperes(t *testing.T) {
    factory := units.ElectricCurrentFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicoamperes(100)
    if err != nil {
        t.Errorf("FromPicoamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentPicoampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicoamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicoamperes(math.NaN())
    if err == nil {
        t.Error("FromPicoamperes() with NaN value should return error")
    }

    _, err = factory.FromPicoamperes(math.Inf(1))
    if err == nil {
        t.Error("FromPicoamperes() with +Inf value should return error")
    }

    _, err = factory.FromPicoamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromPicoamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicoamperes(0)
    if err != nil {
        t.Errorf("FromPicoamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentPicoampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicoamperes() with zero value = %v, want 0", converted)
    }
}
// Test FromNanoamperes function
func TestElectricCurrentFactory_FromNanoamperes(t *testing.T) {
    factory := units.ElectricCurrentFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanoamperes(100)
    if err != nil {
        t.Errorf("FromNanoamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentNanoampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanoamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanoamperes(math.NaN())
    if err == nil {
        t.Error("FromNanoamperes() with NaN value should return error")
    }

    _, err = factory.FromNanoamperes(math.Inf(1))
    if err == nil {
        t.Error("FromNanoamperes() with +Inf value should return error")
    }

    _, err = factory.FromNanoamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromNanoamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanoamperes(0)
    if err != nil {
        t.Errorf("FromNanoamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentNanoampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanoamperes() with zero value = %v, want 0", converted)
    }
}
// Test FromMicroamperes function
func TestElectricCurrentFactory_FromMicroamperes(t *testing.T) {
    factory := units.ElectricCurrentFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicroamperes(100)
    if err != nil {
        t.Errorf("FromMicroamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentMicroampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicroamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicroamperes(math.NaN())
    if err == nil {
        t.Error("FromMicroamperes() with NaN value should return error")
    }

    _, err = factory.FromMicroamperes(math.Inf(1))
    if err == nil {
        t.Error("FromMicroamperes() with +Inf value should return error")
    }

    _, err = factory.FromMicroamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromMicroamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicroamperes(0)
    if err != nil {
        t.Errorf("FromMicroamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentMicroampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicroamperes() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliamperes function
func TestElectricCurrentFactory_FromMilliamperes(t *testing.T) {
    factory := units.ElectricCurrentFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliamperes(100)
    if err != nil {
        t.Errorf("FromMilliamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentMilliampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliamperes(math.NaN())
    if err == nil {
        t.Error("FromMilliamperes() with NaN value should return error")
    }

    _, err = factory.FromMilliamperes(math.Inf(1))
    if err == nil {
        t.Error("FromMilliamperes() with +Inf value should return error")
    }

    _, err = factory.FromMilliamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliamperes(0)
    if err != nil {
        t.Errorf("FromMilliamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentMilliampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliamperes() with zero value = %v, want 0", converted)
    }
}
// Test FromCentiamperes function
func TestElectricCurrentFactory_FromCentiamperes(t *testing.T) {
    factory := units.ElectricCurrentFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentiamperes(100)
    if err != nil {
        t.Errorf("FromCentiamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentCentiampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentiamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentiamperes(math.NaN())
    if err == nil {
        t.Error("FromCentiamperes() with NaN value should return error")
    }

    _, err = factory.FromCentiamperes(math.Inf(1))
    if err == nil {
        t.Error("FromCentiamperes() with +Inf value should return error")
    }

    _, err = factory.FromCentiamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromCentiamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentiamperes(0)
    if err != nil {
        t.Errorf("FromCentiamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentCentiampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentiamperes() with zero value = %v, want 0", converted)
    }
}
// Test FromKiloamperes function
func TestElectricCurrentFactory_FromKiloamperes(t *testing.T) {
    factory := units.ElectricCurrentFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKiloamperes(100)
    if err != nil {
        t.Errorf("FromKiloamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentKiloampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKiloamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKiloamperes(math.NaN())
    if err == nil {
        t.Error("FromKiloamperes() with NaN value should return error")
    }

    _, err = factory.FromKiloamperes(math.Inf(1))
    if err == nil {
        t.Error("FromKiloamperes() with +Inf value should return error")
    }

    _, err = factory.FromKiloamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromKiloamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKiloamperes(0)
    if err != nil {
        t.Errorf("FromKiloamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentKiloampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKiloamperes() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaamperes function
func TestElectricCurrentFactory_FromMegaamperes(t *testing.T) {
    factory := units.ElectricCurrentFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaamperes(100)
    if err != nil {
        t.Errorf("FromMegaamperes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentMegaampere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegaamperes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegaamperes(math.NaN())
    if err == nil {
        t.Error("FromMegaamperes() with NaN value should return error")
    }

    _, err = factory.FromMegaamperes(math.Inf(1))
    if err == nil {
        t.Error("FromMegaamperes() with +Inf value should return error")
    }

    _, err = factory.FromMegaamperes(math.Inf(-1))
    if err == nil {
        t.Error("FromMegaamperes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegaamperes(0)
    if err != nil {
        t.Errorf("FromMegaamperes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentMegaampere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaamperes() with zero value = %v, want 0", converted)
    }
}

func TestElectricCurrentToString(t *testing.T) {
	factory := units.ElectricCurrentFactory{}
	a, err := factory.CreateElectricCurrent(45, units.ElectricCurrentAmpere)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricCurrentAmpere, 2)
	expected := "45.00 " + units.GetElectricCurrentAbbreviation(units.ElectricCurrentAmpere)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricCurrentAmpere, -1)
	expected = "45 " + units.GetElectricCurrentAbbreviation(units.ElectricCurrentAmpere)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricCurrent_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricCurrentFactory{}
	a1, _ := factory.CreateElectricCurrent(60, units.ElectricCurrentAmpere)
	a2, _ := factory.CreateElectricCurrent(60, units.ElectricCurrentAmpere)
	a3, _ := factory.CreateElectricCurrent(120, units.ElectricCurrentAmpere)

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

func TestElectricCurrent_Arithmetic(t *testing.T) {
	factory := units.ElectricCurrentFactory{}
	a1, _ := factory.CreateElectricCurrent(30, units.ElectricCurrentAmpere)
	a2, _ := factory.CreateElectricCurrent(45, units.ElectricCurrentAmpere)

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


func TestGetElectricCurrentAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricCurrentUnits
        want string
    }{
        {
            name: "Ampere abbreviation",
            unit: units.ElectricCurrentAmpere,
            want: "A",
        },
        {
            name: "Femtoampere abbreviation",
            unit: units.ElectricCurrentFemtoampere,
            want: "fA",
        },
        {
            name: "Picoampere abbreviation",
            unit: units.ElectricCurrentPicoampere,
            want: "pA",
        },
        {
            name: "Nanoampere abbreviation",
            unit: units.ElectricCurrentNanoampere,
            want: "nA",
        },
        {
            name: "Microampere abbreviation",
            unit: units.ElectricCurrentMicroampere,
            want: "μA",
        },
        {
            name: "Milliampere abbreviation",
            unit: units.ElectricCurrentMilliampere,
            want: "mA",
        },
        {
            name: "Centiampere abbreviation",
            unit: units.ElectricCurrentCentiampere,
            want: "cA",
        },
        {
            name: "Kiloampere abbreviation",
            unit: units.ElectricCurrentKiloampere,
            want: "kA",
        },
        {
            name: "Megaampere abbreviation",
            unit: units.ElectricCurrentMegaampere,
            want: "MA",
        },
        {
            name: "invalid unit",
            unit: units.ElectricCurrentUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricCurrentAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricCurrentAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricCurrent_String(t *testing.T) {
    factory := units.ElectricCurrentFactory{}
    
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
            unit, err := factory.CreateElectricCurrent(tt.value, units.ElectricCurrentAmpere)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricCurrent.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestElectricCurrent_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.ElectricCurrentFactory{}

	_, err := uf.CreateElectricCurrent(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}