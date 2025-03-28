// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricSusceptanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Siemens"}`
	
	factory := units.ElectricSusceptanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricSusceptanceSiemens {
		t.Errorf("expected unit %v, got %v", units.ElectricSusceptanceSiemens, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Siemens"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricSusceptanceDto_ToJSON(t *testing.T) {
	dto := units.ElectricSusceptanceDto{
		Value: 45,
		Unit:  units.ElectricSusceptanceSiemens,
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
	if result["unit"].(string) != string(units.ElectricSusceptanceSiemens) {
		t.Errorf("expected unit %s, got %v", units.ElectricSusceptanceSiemens, result["unit"])
	}
}

func TestNewElectricSusceptance_InvalidValue(t *testing.T) {
	factory := units.ElectricSusceptanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricSusceptance(math.NaN(), units.ElectricSusceptanceSiemens)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricSusceptance(math.Inf(1), units.ElectricSusceptanceSiemens)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricSusceptanceConversions(t *testing.T) {
	factory := units.ElectricSusceptanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricSusceptance(180, units.ElectricSusceptanceSiemens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Siemens.
		// No expected conversion value provided for Siemens, verifying result is not NaN.
		result := a.Siemens()
		cacheResult := a.Siemens()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Siemens returned NaN")
		}
	}
	{
		// Test conversion to Mhos.
		// No expected conversion value provided for Mhos, verifying result is not NaN.
		result := a.Mhos()
		cacheResult := a.Mhos()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Mhos returned NaN")
		}
	}
	{
		// Test conversion to Nanosiemens.
		// No expected conversion value provided for Nanosiemens, verifying result is not NaN.
		result := a.Nanosiemens()
		cacheResult := a.Nanosiemens()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanosiemens returned NaN")
		}
	}
	{
		// Test conversion to Microsiemens.
		// No expected conversion value provided for Microsiemens, verifying result is not NaN.
		result := a.Microsiemens()
		cacheResult := a.Microsiemens()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microsiemens returned NaN")
		}
	}
	{
		// Test conversion to Millisiemens.
		// No expected conversion value provided for Millisiemens, verifying result is not NaN.
		result := a.Millisiemens()
		cacheResult := a.Millisiemens()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millisiemens returned NaN")
		}
	}
	{
		// Test conversion to Kilosiemens.
		// No expected conversion value provided for Kilosiemens, verifying result is not NaN.
		result := a.Kilosiemens()
		cacheResult := a.Kilosiemens()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilosiemens returned NaN")
		}
	}
	{
		// Test conversion to Megasiemens.
		// No expected conversion value provided for Megasiemens, verifying result is not NaN.
		result := a.Megasiemens()
		cacheResult := a.Megasiemens()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megasiemens returned NaN")
		}
	}
	{
		// Test conversion to Gigasiemens.
		// No expected conversion value provided for Gigasiemens, verifying result is not NaN.
		result := a.Gigasiemens()
		cacheResult := a.Gigasiemens()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gigasiemens returned NaN")
		}
	}
	{
		// Test conversion to Terasiemens.
		// No expected conversion value provided for Terasiemens, verifying result is not NaN.
		result := a.Terasiemens()
		cacheResult := a.Terasiemens()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Terasiemens returned NaN")
		}
	}
	{
		// Test conversion to Nanomhos.
		// No expected conversion value provided for Nanomhos, verifying result is not NaN.
		result := a.Nanomhos()
		cacheResult := a.Nanomhos()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanomhos returned NaN")
		}
	}
	{
		// Test conversion to Micromhos.
		// No expected conversion value provided for Micromhos, verifying result is not NaN.
		result := a.Micromhos()
		cacheResult := a.Micromhos()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Micromhos returned NaN")
		}
	}
	{
		// Test conversion to Millimhos.
		// No expected conversion value provided for Millimhos, verifying result is not NaN.
		result := a.Millimhos()
		cacheResult := a.Millimhos()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millimhos returned NaN")
		}
	}
	{
		// Test conversion to Kilomhos.
		// No expected conversion value provided for Kilomhos, verifying result is not NaN.
		result := a.Kilomhos()
		cacheResult := a.Kilomhos()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilomhos returned NaN")
		}
	}
	{
		// Test conversion to Megamhos.
		// No expected conversion value provided for Megamhos, verifying result is not NaN.
		result := a.Megamhos()
		cacheResult := a.Megamhos()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megamhos returned NaN")
		}
	}
	{
		// Test conversion to Gigamhos.
		// No expected conversion value provided for Gigamhos, verifying result is not NaN.
		result := a.Gigamhos()
		cacheResult := a.Gigamhos()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gigamhos returned NaN")
		}
	}
	{
		// Test conversion to Teramhos.
		// No expected conversion value provided for Teramhos, verifying result is not NaN.
		result := a.Teramhos()
		cacheResult := a.Teramhos()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Teramhos returned NaN")
		}
	}
}

func TestElectricSusceptance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricSusceptanceFactory{}
	a, err := factory.CreateElectricSusceptance(90, units.ElectricSusceptanceSiemens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricSusceptanceSiemens {
		t.Errorf("expected default unit Siemens, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricSusceptanceSiemens
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricSusceptanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricSusceptanceSiemens {
		t.Errorf("expected unit Siemens, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricSusceptanceFactory_FromDto(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceSiemens,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricSusceptanceDto{
        Value: math.NaN(),
        Unit:  units.ElectricSusceptanceSiemens,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Siemens conversion
    siemensDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceSiemens,
    }
    
    var siemensResult *units.ElectricSusceptance
    siemensResult, err = factory.FromDto(siemensDto)
    if err != nil {
        t.Errorf("FromDto() with Siemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = siemensResult.Convert(units.ElectricSusceptanceSiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Siemens = %v, want %v", converted, 100)
    }
    // Test Mho conversion
    mhosDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceMho,
    }
    
    var mhosResult *units.ElectricSusceptance
    mhosResult, err = factory.FromDto(mhosDto)
    if err != nil {
        t.Errorf("FromDto() with Mho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mhosResult.Convert(units.ElectricSusceptanceMho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mho = %v, want %v", converted, 100)
    }
    // Test Nanosiemens conversion
    nanosiemensDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceNanosiemens,
    }
    
    var nanosiemensResult *units.ElectricSusceptance
    nanosiemensResult, err = factory.FromDto(nanosiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Nanosiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanosiemensResult.Convert(units.ElectricSusceptanceNanosiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanosiemens = %v, want %v", converted, 100)
    }
    // Test Microsiemens conversion
    microsiemensDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceMicrosiemens,
    }
    
    var microsiemensResult *units.ElectricSusceptance
    microsiemensResult, err = factory.FromDto(microsiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Microsiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microsiemensResult.Convert(units.ElectricSusceptanceMicrosiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microsiemens = %v, want %v", converted, 100)
    }
    // Test Millisiemens conversion
    millisiemensDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceMillisiemens,
    }
    
    var millisiemensResult *units.ElectricSusceptance
    millisiemensResult, err = factory.FromDto(millisiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Millisiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millisiemensResult.Convert(units.ElectricSusceptanceMillisiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millisiemens = %v, want %v", converted, 100)
    }
    // Test Kilosiemens conversion
    kilosiemensDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceKilosiemens,
    }
    
    var kilosiemensResult *units.ElectricSusceptance
    kilosiemensResult, err = factory.FromDto(kilosiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Kilosiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilosiemensResult.Convert(units.ElectricSusceptanceKilosiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilosiemens = %v, want %v", converted, 100)
    }
    // Test Megasiemens conversion
    megasiemensDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceMegasiemens,
    }
    
    var megasiemensResult *units.ElectricSusceptance
    megasiemensResult, err = factory.FromDto(megasiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Megasiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megasiemensResult.Convert(units.ElectricSusceptanceMegasiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megasiemens = %v, want %v", converted, 100)
    }
    // Test Gigasiemens conversion
    gigasiemensDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceGigasiemens,
    }
    
    var gigasiemensResult *units.ElectricSusceptance
    gigasiemensResult, err = factory.FromDto(gigasiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Gigasiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigasiemensResult.Convert(units.ElectricSusceptanceGigasiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigasiemens = %v, want %v", converted, 100)
    }
    // Test Terasiemens conversion
    terasiemensDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceTerasiemens,
    }
    
    var terasiemensResult *units.ElectricSusceptance
    terasiemensResult, err = factory.FromDto(terasiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Terasiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terasiemensResult.Convert(units.ElectricSusceptanceTerasiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terasiemens = %v, want %v", converted, 100)
    }
    // Test Nanomho conversion
    nanomhosDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceNanomho,
    }
    
    var nanomhosResult *units.ElectricSusceptance
    nanomhosResult, err = factory.FromDto(nanomhosDto)
    if err != nil {
        t.Errorf("FromDto() with Nanomho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanomhosResult.Convert(units.ElectricSusceptanceNanomho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanomho = %v, want %v", converted, 100)
    }
    // Test Micromho conversion
    micromhosDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceMicromho,
    }
    
    var micromhosResult *units.ElectricSusceptance
    micromhosResult, err = factory.FromDto(micromhosDto)
    if err != nil {
        t.Errorf("FromDto() with Micromho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micromhosResult.Convert(units.ElectricSusceptanceMicromho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Micromho = %v, want %v", converted, 100)
    }
    // Test Millimho conversion
    millimhosDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceMillimho,
    }
    
    var millimhosResult *units.ElectricSusceptance
    millimhosResult, err = factory.FromDto(millimhosDto)
    if err != nil {
        t.Errorf("FromDto() with Millimho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimhosResult.Convert(units.ElectricSusceptanceMillimho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millimho = %v, want %v", converted, 100)
    }
    // Test Kilomho conversion
    kilomhosDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceKilomho,
    }
    
    var kilomhosResult *units.ElectricSusceptance
    kilomhosResult, err = factory.FromDto(kilomhosDto)
    if err != nil {
        t.Errorf("FromDto() with Kilomho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilomhosResult.Convert(units.ElectricSusceptanceKilomho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilomho = %v, want %v", converted, 100)
    }
    // Test Megamho conversion
    megamhosDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceMegamho,
    }
    
    var megamhosResult *units.ElectricSusceptance
    megamhosResult, err = factory.FromDto(megamhosDto)
    if err != nil {
        t.Errorf("FromDto() with Megamho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megamhosResult.Convert(units.ElectricSusceptanceMegamho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megamho = %v, want %v", converted, 100)
    }
    // Test Gigamho conversion
    gigamhosDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceGigamho,
    }
    
    var gigamhosResult *units.ElectricSusceptance
    gigamhosResult, err = factory.FromDto(gigamhosDto)
    if err != nil {
        t.Errorf("FromDto() with Gigamho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigamhosResult.Convert(units.ElectricSusceptanceGigamho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigamho = %v, want %v", converted, 100)
    }
    // Test Teramho conversion
    teramhosDto := units.ElectricSusceptanceDto{
        Value: 100,
        Unit:  units.ElectricSusceptanceTeramho,
    }
    
    var teramhosResult *units.ElectricSusceptance
    teramhosResult, err = factory.FromDto(teramhosDto)
    if err != nil {
        t.Errorf("FromDto() with Teramho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teramhosResult.Convert(units.ElectricSusceptanceTeramho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Teramho = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricSusceptanceDto{
        Value: 0,
        Unit:  units.ElectricSusceptanceSiemens,
    }
    
    var zeroResult *units.ElectricSusceptance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricSusceptanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Siemens"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Siemens"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricSusceptanceDto{
        Value: nanValue,
        Unit:  units.ElectricSusceptanceSiemens,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Siemens unit
    siemensJSON := []byte(`{"value": 100, "unit": "Siemens"}`)
    siemensResult, err := factory.FromDtoJSON(siemensJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Siemens unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = siemensResult.Convert(units.ElectricSusceptanceSiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Siemens = %v, want %v", converted, 100)
    }
    // Test JSON with Mho unit
    mhosJSON := []byte(`{"value": 100, "unit": "Mho"}`)
    mhosResult, err := factory.FromDtoJSON(mhosJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Mho unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mhosResult.Convert(units.ElectricSusceptanceMho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mho = %v, want %v", converted, 100)
    }
    // Test JSON with Nanosiemens unit
    nanosiemensJSON := []byte(`{"value": 100, "unit": "Nanosiemens"}`)
    nanosiemensResult, err := factory.FromDtoJSON(nanosiemensJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanosiemens unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanosiemensResult.Convert(units.ElectricSusceptanceNanosiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanosiemens = %v, want %v", converted, 100)
    }
    // Test JSON with Microsiemens unit
    microsiemensJSON := []byte(`{"value": 100, "unit": "Microsiemens"}`)
    microsiemensResult, err := factory.FromDtoJSON(microsiemensJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microsiemens unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microsiemensResult.Convert(units.ElectricSusceptanceMicrosiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microsiemens = %v, want %v", converted, 100)
    }
    // Test JSON with Millisiemens unit
    millisiemensJSON := []byte(`{"value": 100, "unit": "Millisiemens"}`)
    millisiemensResult, err := factory.FromDtoJSON(millisiemensJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millisiemens unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millisiemensResult.Convert(units.ElectricSusceptanceMillisiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millisiemens = %v, want %v", converted, 100)
    }
    // Test JSON with Kilosiemens unit
    kilosiemensJSON := []byte(`{"value": 100, "unit": "Kilosiemens"}`)
    kilosiemensResult, err := factory.FromDtoJSON(kilosiemensJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilosiemens unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilosiemensResult.Convert(units.ElectricSusceptanceKilosiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilosiemens = %v, want %v", converted, 100)
    }
    // Test JSON with Megasiemens unit
    megasiemensJSON := []byte(`{"value": 100, "unit": "Megasiemens"}`)
    megasiemensResult, err := factory.FromDtoJSON(megasiemensJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megasiemens unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megasiemensResult.Convert(units.ElectricSusceptanceMegasiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megasiemens = %v, want %v", converted, 100)
    }
    // Test JSON with Gigasiemens unit
    gigasiemensJSON := []byte(`{"value": 100, "unit": "Gigasiemens"}`)
    gigasiemensResult, err := factory.FromDtoJSON(gigasiemensJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigasiemens unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigasiemensResult.Convert(units.ElectricSusceptanceGigasiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigasiemens = %v, want %v", converted, 100)
    }
    // Test JSON with Terasiemens unit
    terasiemensJSON := []byte(`{"value": 100, "unit": "Terasiemens"}`)
    terasiemensResult, err := factory.FromDtoJSON(terasiemensJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Terasiemens unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terasiemensResult.Convert(units.ElectricSusceptanceTerasiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terasiemens = %v, want %v", converted, 100)
    }
    // Test JSON with Nanomho unit
    nanomhosJSON := []byte(`{"value": 100, "unit": "Nanomho"}`)
    nanomhosResult, err := factory.FromDtoJSON(nanomhosJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanomho unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanomhosResult.Convert(units.ElectricSusceptanceNanomho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanomho = %v, want %v", converted, 100)
    }
    // Test JSON with Micromho unit
    micromhosJSON := []byte(`{"value": 100, "unit": "Micromho"}`)
    micromhosResult, err := factory.FromDtoJSON(micromhosJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Micromho unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micromhosResult.Convert(units.ElectricSusceptanceMicromho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Micromho = %v, want %v", converted, 100)
    }
    // Test JSON with Millimho unit
    millimhosJSON := []byte(`{"value": 100, "unit": "Millimho"}`)
    millimhosResult, err := factory.FromDtoJSON(millimhosJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millimho unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimhosResult.Convert(units.ElectricSusceptanceMillimho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millimho = %v, want %v", converted, 100)
    }
    // Test JSON with Kilomho unit
    kilomhosJSON := []byte(`{"value": 100, "unit": "Kilomho"}`)
    kilomhosResult, err := factory.FromDtoJSON(kilomhosJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilomho unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilomhosResult.Convert(units.ElectricSusceptanceKilomho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilomho = %v, want %v", converted, 100)
    }
    // Test JSON with Megamho unit
    megamhosJSON := []byte(`{"value": 100, "unit": "Megamho"}`)
    megamhosResult, err := factory.FromDtoJSON(megamhosJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megamho unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megamhosResult.Convert(units.ElectricSusceptanceMegamho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megamho = %v, want %v", converted, 100)
    }
    // Test JSON with Gigamho unit
    gigamhosJSON := []byte(`{"value": 100, "unit": "Gigamho"}`)
    gigamhosResult, err := factory.FromDtoJSON(gigamhosJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigamho unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigamhosResult.Convert(units.ElectricSusceptanceGigamho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigamho = %v, want %v", converted, 100)
    }
    // Test JSON with Teramho unit
    teramhosJSON := []byte(`{"value": 100, "unit": "Teramho"}`)
    teramhosResult, err := factory.FromDtoJSON(teramhosJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Teramho unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teramhosResult.Convert(units.ElectricSusceptanceTeramho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Teramho = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Siemens"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromSiemens function
func TestElectricSusceptanceFactory_FromSiemens(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSiemens(100)
    if err != nil {
        t.Errorf("FromSiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceSiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSiemens() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSiemens(math.NaN())
    if err == nil {
        t.Error("FromSiemens() with NaN value should return error")
    }

    _, err = factory.FromSiemens(math.Inf(1))
    if err == nil {
        t.Error("FromSiemens() with +Inf value should return error")
    }

    _, err = factory.FromSiemens(math.Inf(-1))
    if err == nil {
        t.Error("FromSiemens() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSiemens(0)
    if err != nil {
        t.Errorf("FromSiemens() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceSiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromMhos function
func TestElectricSusceptanceFactory_FromMhos(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMhos(100)
    if err != nil {
        t.Errorf("FromMhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceMho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMhos() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMhos(math.NaN())
    if err == nil {
        t.Error("FromMhos() with NaN value should return error")
    }

    _, err = factory.FromMhos(math.Inf(1))
    if err == nil {
        t.Error("FromMhos() with +Inf value should return error")
    }

    _, err = factory.FromMhos(math.Inf(-1))
    if err == nil {
        t.Error("FromMhos() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMhos(0)
    if err != nil {
        t.Errorf("FromMhos() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceMho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMhos() with zero value = %v, want 0", converted)
    }
}
// Test FromNanosiemens function
func TestElectricSusceptanceFactory_FromNanosiemens(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanosiemens(100)
    if err != nil {
        t.Errorf("FromNanosiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceNanosiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanosiemens() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanosiemens(math.NaN())
    if err == nil {
        t.Error("FromNanosiemens() with NaN value should return error")
    }

    _, err = factory.FromNanosiemens(math.Inf(1))
    if err == nil {
        t.Error("FromNanosiemens() with +Inf value should return error")
    }

    _, err = factory.FromNanosiemens(math.Inf(-1))
    if err == nil {
        t.Error("FromNanosiemens() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanosiemens(0)
    if err != nil {
        t.Errorf("FromNanosiemens() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceNanosiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanosiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrosiemens function
func TestElectricSusceptanceFactory_FromMicrosiemens(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrosiemens(100)
    if err != nil {
        t.Errorf("FromMicrosiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceMicrosiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrosiemens() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrosiemens(math.NaN())
    if err == nil {
        t.Error("FromMicrosiemens() with NaN value should return error")
    }

    _, err = factory.FromMicrosiemens(math.Inf(1))
    if err == nil {
        t.Error("FromMicrosiemens() with +Inf value should return error")
    }

    _, err = factory.FromMicrosiemens(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrosiemens() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrosiemens(0)
    if err != nil {
        t.Errorf("FromMicrosiemens() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceMicrosiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrosiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromMillisiemens function
func TestElectricSusceptanceFactory_FromMillisiemens(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillisiemens(100)
    if err != nil {
        t.Errorf("FromMillisiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceMillisiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillisiemens() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillisiemens(math.NaN())
    if err == nil {
        t.Error("FromMillisiemens() with NaN value should return error")
    }

    _, err = factory.FromMillisiemens(math.Inf(1))
    if err == nil {
        t.Error("FromMillisiemens() with +Inf value should return error")
    }

    _, err = factory.FromMillisiemens(math.Inf(-1))
    if err == nil {
        t.Error("FromMillisiemens() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillisiemens(0)
    if err != nil {
        t.Errorf("FromMillisiemens() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceMillisiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillisiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromKilosiemens function
func TestElectricSusceptanceFactory_FromKilosiemens(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilosiemens(100)
    if err != nil {
        t.Errorf("FromKilosiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceKilosiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilosiemens() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilosiemens(math.NaN())
    if err == nil {
        t.Error("FromKilosiemens() with NaN value should return error")
    }

    _, err = factory.FromKilosiemens(math.Inf(1))
    if err == nil {
        t.Error("FromKilosiemens() with +Inf value should return error")
    }

    _, err = factory.FromKilosiemens(math.Inf(-1))
    if err == nil {
        t.Error("FromKilosiemens() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilosiemens(0)
    if err != nil {
        t.Errorf("FromKilosiemens() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceKilosiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilosiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromMegasiemens function
func TestElectricSusceptanceFactory_FromMegasiemens(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegasiemens(100)
    if err != nil {
        t.Errorf("FromMegasiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceMegasiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegasiemens() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegasiemens(math.NaN())
    if err == nil {
        t.Error("FromMegasiemens() with NaN value should return error")
    }

    _, err = factory.FromMegasiemens(math.Inf(1))
    if err == nil {
        t.Error("FromMegasiemens() with +Inf value should return error")
    }

    _, err = factory.FromMegasiemens(math.Inf(-1))
    if err == nil {
        t.Error("FromMegasiemens() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegasiemens(0)
    if err != nil {
        t.Errorf("FromMegasiemens() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceMegasiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegasiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromGigasiemens function
func TestElectricSusceptanceFactory_FromGigasiemens(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigasiemens(100)
    if err != nil {
        t.Errorf("FromGigasiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceGigasiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigasiemens() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigasiemens(math.NaN())
    if err == nil {
        t.Error("FromGigasiemens() with NaN value should return error")
    }

    _, err = factory.FromGigasiemens(math.Inf(1))
    if err == nil {
        t.Error("FromGigasiemens() with +Inf value should return error")
    }

    _, err = factory.FromGigasiemens(math.Inf(-1))
    if err == nil {
        t.Error("FromGigasiemens() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigasiemens(0)
    if err != nil {
        t.Errorf("FromGigasiemens() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceGigasiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigasiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromTerasiemens function
func TestElectricSusceptanceFactory_FromTerasiemens(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerasiemens(100)
    if err != nil {
        t.Errorf("FromTerasiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceTerasiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerasiemens() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerasiemens(math.NaN())
    if err == nil {
        t.Error("FromTerasiemens() with NaN value should return error")
    }

    _, err = factory.FromTerasiemens(math.Inf(1))
    if err == nil {
        t.Error("FromTerasiemens() with +Inf value should return error")
    }

    _, err = factory.FromTerasiemens(math.Inf(-1))
    if err == nil {
        t.Error("FromTerasiemens() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerasiemens(0)
    if err != nil {
        t.Errorf("FromTerasiemens() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceTerasiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerasiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromNanomhos function
func TestElectricSusceptanceFactory_FromNanomhos(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanomhos(100)
    if err != nil {
        t.Errorf("FromNanomhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceNanomho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanomhos() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanomhos(math.NaN())
    if err == nil {
        t.Error("FromNanomhos() with NaN value should return error")
    }

    _, err = factory.FromNanomhos(math.Inf(1))
    if err == nil {
        t.Error("FromNanomhos() with +Inf value should return error")
    }

    _, err = factory.FromNanomhos(math.Inf(-1))
    if err == nil {
        t.Error("FromNanomhos() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanomhos(0)
    if err != nil {
        t.Errorf("FromNanomhos() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceNanomho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanomhos() with zero value = %v, want 0", converted)
    }
}
// Test FromMicromhos function
func TestElectricSusceptanceFactory_FromMicromhos(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicromhos(100)
    if err != nil {
        t.Errorf("FromMicromhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceMicromho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicromhos() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicromhos(math.NaN())
    if err == nil {
        t.Error("FromMicromhos() with NaN value should return error")
    }

    _, err = factory.FromMicromhos(math.Inf(1))
    if err == nil {
        t.Error("FromMicromhos() with +Inf value should return error")
    }

    _, err = factory.FromMicromhos(math.Inf(-1))
    if err == nil {
        t.Error("FromMicromhos() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicromhos(0)
    if err != nil {
        t.Errorf("FromMicromhos() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceMicromho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicromhos() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimhos function
func TestElectricSusceptanceFactory_FromMillimhos(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimhos(100)
    if err != nil {
        t.Errorf("FromMillimhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceMillimho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimhos() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimhos(math.NaN())
    if err == nil {
        t.Error("FromMillimhos() with NaN value should return error")
    }

    _, err = factory.FromMillimhos(math.Inf(1))
    if err == nil {
        t.Error("FromMillimhos() with +Inf value should return error")
    }

    _, err = factory.FromMillimhos(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimhos() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimhos(0)
    if err != nil {
        t.Errorf("FromMillimhos() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceMillimho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimhos() with zero value = %v, want 0", converted)
    }
}
// Test FromKilomhos function
func TestElectricSusceptanceFactory_FromKilomhos(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilomhos(100)
    if err != nil {
        t.Errorf("FromKilomhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceKilomho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilomhos() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilomhos(math.NaN())
    if err == nil {
        t.Error("FromKilomhos() with NaN value should return error")
    }

    _, err = factory.FromKilomhos(math.Inf(1))
    if err == nil {
        t.Error("FromKilomhos() with +Inf value should return error")
    }

    _, err = factory.FromKilomhos(math.Inf(-1))
    if err == nil {
        t.Error("FromKilomhos() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilomhos(0)
    if err != nil {
        t.Errorf("FromKilomhos() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceKilomho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilomhos() with zero value = %v, want 0", converted)
    }
}
// Test FromMegamhos function
func TestElectricSusceptanceFactory_FromMegamhos(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegamhos(100)
    if err != nil {
        t.Errorf("FromMegamhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceMegamho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegamhos() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegamhos(math.NaN())
    if err == nil {
        t.Error("FromMegamhos() with NaN value should return error")
    }

    _, err = factory.FromMegamhos(math.Inf(1))
    if err == nil {
        t.Error("FromMegamhos() with +Inf value should return error")
    }

    _, err = factory.FromMegamhos(math.Inf(-1))
    if err == nil {
        t.Error("FromMegamhos() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegamhos(0)
    if err != nil {
        t.Errorf("FromMegamhos() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceMegamho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegamhos() with zero value = %v, want 0", converted)
    }
}
// Test FromGigamhos function
func TestElectricSusceptanceFactory_FromGigamhos(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigamhos(100)
    if err != nil {
        t.Errorf("FromGigamhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceGigamho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigamhos() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigamhos(math.NaN())
    if err == nil {
        t.Error("FromGigamhos() with NaN value should return error")
    }

    _, err = factory.FromGigamhos(math.Inf(1))
    if err == nil {
        t.Error("FromGigamhos() with +Inf value should return error")
    }

    _, err = factory.FromGigamhos(math.Inf(-1))
    if err == nil {
        t.Error("FromGigamhos() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigamhos(0)
    if err != nil {
        t.Errorf("FromGigamhos() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceGigamho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigamhos() with zero value = %v, want 0", converted)
    }
}
// Test FromTeramhos function
func TestElectricSusceptanceFactory_FromTeramhos(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTeramhos(100)
    if err != nil {
        t.Errorf("FromTeramhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSusceptanceTeramho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTeramhos() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTeramhos(math.NaN())
    if err == nil {
        t.Error("FromTeramhos() with NaN value should return error")
    }

    _, err = factory.FromTeramhos(math.Inf(1))
    if err == nil {
        t.Error("FromTeramhos() with +Inf value should return error")
    }

    _, err = factory.FromTeramhos(math.Inf(-1))
    if err == nil {
        t.Error("FromTeramhos() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTeramhos(0)
    if err != nil {
        t.Errorf("FromTeramhos() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSusceptanceTeramho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTeramhos() with zero value = %v, want 0", converted)
    }
}

func TestElectricSusceptanceToString(t *testing.T) {
	factory := units.ElectricSusceptanceFactory{}
	a, err := factory.CreateElectricSusceptance(45, units.ElectricSusceptanceSiemens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricSusceptanceSiemens, 2)
	expected := "45.00 " + units.GetElectricSusceptanceAbbreviation(units.ElectricSusceptanceSiemens)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricSusceptanceSiemens, -1)
	expected = "45 " + units.GetElectricSusceptanceAbbreviation(units.ElectricSusceptanceSiemens)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricSusceptance_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricSusceptanceFactory{}
	a1, _ := factory.CreateElectricSusceptance(60, units.ElectricSusceptanceSiemens)
	a2, _ := factory.CreateElectricSusceptance(60, units.ElectricSusceptanceSiemens)
	a3, _ := factory.CreateElectricSusceptance(120, units.ElectricSusceptanceSiemens)

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

func TestElectricSusceptance_Arithmetic(t *testing.T) {
	factory := units.ElectricSusceptanceFactory{}
	a1, _ := factory.CreateElectricSusceptance(30, units.ElectricSusceptanceSiemens)
	a2, _ := factory.CreateElectricSusceptance(45, units.ElectricSusceptanceSiemens)

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


func TestGetElectricSusceptanceAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricSusceptanceUnits
        want string
    }{
        {
            name: "Siemens abbreviation",
            unit: units.ElectricSusceptanceSiemens,
            want: "S",
        },
        {
            name: "Mho abbreviation",
            unit: units.ElectricSusceptanceMho,
            want: "℧",
        },
        {
            name: "Nanosiemens abbreviation",
            unit: units.ElectricSusceptanceNanosiemens,
            want: "nS",
        },
        {
            name: "Microsiemens abbreviation",
            unit: units.ElectricSusceptanceMicrosiemens,
            want: "μS",
        },
        {
            name: "Millisiemens abbreviation",
            unit: units.ElectricSusceptanceMillisiemens,
            want: "mS",
        },
        {
            name: "Kilosiemens abbreviation",
            unit: units.ElectricSusceptanceKilosiemens,
            want: "kS",
        },
        {
            name: "Megasiemens abbreviation",
            unit: units.ElectricSusceptanceMegasiemens,
            want: "MS",
        },
        {
            name: "Gigasiemens abbreviation",
            unit: units.ElectricSusceptanceGigasiemens,
            want: "GS",
        },
        {
            name: "Terasiemens abbreviation",
            unit: units.ElectricSusceptanceTerasiemens,
            want: "TS",
        },
        {
            name: "Nanomho abbreviation",
            unit: units.ElectricSusceptanceNanomho,
            want: "n℧",
        },
        {
            name: "Micromho abbreviation",
            unit: units.ElectricSusceptanceMicromho,
            want: "μ℧",
        },
        {
            name: "Millimho abbreviation",
            unit: units.ElectricSusceptanceMillimho,
            want: "m℧",
        },
        {
            name: "Kilomho abbreviation",
            unit: units.ElectricSusceptanceKilomho,
            want: "k℧",
        },
        {
            name: "Megamho abbreviation",
            unit: units.ElectricSusceptanceMegamho,
            want: "M℧",
        },
        {
            name: "Gigamho abbreviation",
            unit: units.ElectricSusceptanceGigamho,
            want: "G℧",
        },
        {
            name: "Teramho abbreviation",
            unit: units.ElectricSusceptanceTeramho,
            want: "T℧",
        },
        {
            name: "invalid unit",
            unit: units.ElectricSusceptanceUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricSusceptanceAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricSusceptanceAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricSusceptance_String(t *testing.T) {
    factory := units.ElectricSusceptanceFactory{}
    
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
            unit, err := factory.CreateElectricSusceptance(tt.value, units.ElectricSusceptanceSiemens)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricSusceptance.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestElectricSusceptance_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.ElectricSusceptanceFactory{}

	_, err := uf.CreateElectricSusceptance(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}