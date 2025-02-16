// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricConductanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Siemens"}`
	
	factory := units.ElectricConductanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricConductanceSiemens {
		t.Errorf("expected unit %v, got %v", units.ElectricConductanceSiemens, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Siemens"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricConductanceDto_ToJSON(t *testing.T) {
	dto := units.ElectricConductanceDto{
		Value: 45,
		Unit:  units.ElectricConductanceSiemens,
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
	if result["unit"].(string) != string(units.ElectricConductanceSiemens) {
		t.Errorf("expected unit %s, got %v", units.ElectricConductanceSiemens, result["unit"])
	}
}

func TestNewElectricConductance_InvalidValue(t *testing.T) {
	factory := units.ElectricConductanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricConductance(math.NaN(), units.ElectricConductanceSiemens)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricConductance(math.Inf(1), units.ElectricConductanceSiemens)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricConductanceConversions(t *testing.T) {
	factory := units.ElectricConductanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricConductance(180, units.ElectricConductanceSiemens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Siemens.
		// No expected conversion value provided for Siemens, verifying result is not NaN.
		result := a.Siemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Siemens returned NaN")
		}
	}
	{
		// Test conversion to Mhos.
		// No expected conversion value provided for Mhos, verifying result is not NaN.
		result := a.Mhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Mhos returned NaN")
		}
	}
	{
		// Test conversion to Nanosiemens.
		// No expected conversion value provided for Nanosiemens, verifying result is not NaN.
		result := a.Nanosiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanosiemens returned NaN")
		}
	}
	{
		// Test conversion to Microsiemens.
		// No expected conversion value provided for Microsiemens, verifying result is not NaN.
		result := a.Microsiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microsiemens returned NaN")
		}
	}
	{
		// Test conversion to Millisiemens.
		// No expected conversion value provided for Millisiemens, verifying result is not NaN.
		result := a.Millisiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millisiemens returned NaN")
		}
	}
	{
		// Test conversion to Kilosiemens.
		// No expected conversion value provided for Kilosiemens, verifying result is not NaN.
		result := a.Kilosiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilosiemens returned NaN")
		}
	}
	{
		// Test conversion to Megasiemens.
		// No expected conversion value provided for Megasiemens, verifying result is not NaN.
		result := a.Megasiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megasiemens returned NaN")
		}
	}
	{
		// Test conversion to Gigasiemens.
		// No expected conversion value provided for Gigasiemens, verifying result is not NaN.
		result := a.Gigasiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigasiemens returned NaN")
		}
	}
	{
		// Test conversion to Terasiemens.
		// No expected conversion value provided for Terasiemens, verifying result is not NaN.
		result := a.Terasiemens()
		if math.IsNaN(result) {
			t.Errorf("conversion to Terasiemens returned NaN")
		}
	}
	{
		// Test conversion to Nanomhos.
		// No expected conversion value provided for Nanomhos, verifying result is not NaN.
		result := a.Nanomhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanomhos returned NaN")
		}
	}
	{
		// Test conversion to Micromhos.
		// No expected conversion value provided for Micromhos, verifying result is not NaN.
		result := a.Micromhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Micromhos returned NaN")
		}
	}
	{
		// Test conversion to Millimhos.
		// No expected conversion value provided for Millimhos, verifying result is not NaN.
		result := a.Millimhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millimhos returned NaN")
		}
	}
	{
		// Test conversion to Kilomhos.
		// No expected conversion value provided for Kilomhos, verifying result is not NaN.
		result := a.Kilomhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilomhos returned NaN")
		}
	}
	{
		// Test conversion to Megamhos.
		// No expected conversion value provided for Megamhos, verifying result is not NaN.
		result := a.Megamhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megamhos returned NaN")
		}
	}
	{
		// Test conversion to Gigamhos.
		// No expected conversion value provided for Gigamhos, verifying result is not NaN.
		result := a.Gigamhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigamhos returned NaN")
		}
	}
	{
		// Test conversion to Teramhos.
		// No expected conversion value provided for Teramhos, verifying result is not NaN.
		result := a.Teramhos()
		if math.IsNaN(result) {
			t.Errorf("conversion to Teramhos returned NaN")
		}
	}
}

func TestElectricConductance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricConductanceFactory{}
	a, err := factory.CreateElectricConductance(90, units.ElectricConductanceSiemens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricConductanceSiemens {
		t.Errorf("expected default unit Siemens, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricConductanceSiemens
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricConductanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricConductanceSiemens {
		t.Errorf("expected unit Siemens, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricConductanceFactory_FromDto(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceSiemens,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricConductanceDto{
        Value: math.NaN(),
        Unit:  units.ElectricConductanceSiemens,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Siemens conversion
    siemensDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceSiemens,
    }
    
    var siemensResult *units.ElectricConductance
    siemensResult, err = factory.FromDto(siemensDto)
    if err != nil {
        t.Errorf("FromDto() with Siemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = siemensResult.Convert(units.ElectricConductanceSiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Siemens = %v, want %v", converted, 100)
    }
    // Test Mho conversion
    mhosDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceMho,
    }
    
    var mhosResult *units.ElectricConductance
    mhosResult, err = factory.FromDto(mhosDto)
    if err != nil {
        t.Errorf("FromDto() with Mho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mhosResult.Convert(units.ElectricConductanceMho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mho = %v, want %v", converted, 100)
    }
    // Test Nanosiemens conversion
    nanosiemensDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceNanosiemens,
    }
    
    var nanosiemensResult *units.ElectricConductance
    nanosiemensResult, err = factory.FromDto(nanosiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Nanosiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanosiemensResult.Convert(units.ElectricConductanceNanosiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanosiemens = %v, want %v", converted, 100)
    }
    // Test Microsiemens conversion
    microsiemensDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceMicrosiemens,
    }
    
    var microsiemensResult *units.ElectricConductance
    microsiemensResult, err = factory.FromDto(microsiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Microsiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microsiemensResult.Convert(units.ElectricConductanceMicrosiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microsiemens = %v, want %v", converted, 100)
    }
    // Test Millisiemens conversion
    millisiemensDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceMillisiemens,
    }
    
    var millisiemensResult *units.ElectricConductance
    millisiemensResult, err = factory.FromDto(millisiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Millisiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millisiemensResult.Convert(units.ElectricConductanceMillisiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millisiemens = %v, want %v", converted, 100)
    }
    // Test Kilosiemens conversion
    kilosiemensDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceKilosiemens,
    }
    
    var kilosiemensResult *units.ElectricConductance
    kilosiemensResult, err = factory.FromDto(kilosiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Kilosiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilosiemensResult.Convert(units.ElectricConductanceKilosiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilosiemens = %v, want %v", converted, 100)
    }
    // Test Megasiemens conversion
    megasiemensDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceMegasiemens,
    }
    
    var megasiemensResult *units.ElectricConductance
    megasiemensResult, err = factory.FromDto(megasiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Megasiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megasiemensResult.Convert(units.ElectricConductanceMegasiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megasiemens = %v, want %v", converted, 100)
    }
    // Test Gigasiemens conversion
    gigasiemensDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceGigasiemens,
    }
    
    var gigasiemensResult *units.ElectricConductance
    gigasiemensResult, err = factory.FromDto(gigasiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Gigasiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigasiemensResult.Convert(units.ElectricConductanceGigasiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigasiemens = %v, want %v", converted, 100)
    }
    // Test Terasiemens conversion
    terasiemensDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceTerasiemens,
    }
    
    var terasiemensResult *units.ElectricConductance
    terasiemensResult, err = factory.FromDto(terasiemensDto)
    if err != nil {
        t.Errorf("FromDto() with Terasiemens returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terasiemensResult.Convert(units.ElectricConductanceTerasiemens)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terasiemens = %v, want %v", converted, 100)
    }
    // Test Nanomho conversion
    nanomhosDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceNanomho,
    }
    
    var nanomhosResult *units.ElectricConductance
    nanomhosResult, err = factory.FromDto(nanomhosDto)
    if err != nil {
        t.Errorf("FromDto() with Nanomho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanomhosResult.Convert(units.ElectricConductanceNanomho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanomho = %v, want %v", converted, 100)
    }
    // Test Micromho conversion
    micromhosDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceMicromho,
    }
    
    var micromhosResult *units.ElectricConductance
    micromhosResult, err = factory.FromDto(micromhosDto)
    if err != nil {
        t.Errorf("FromDto() with Micromho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micromhosResult.Convert(units.ElectricConductanceMicromho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Micromho = %v, want %v", converted, 100)
    }
    // Test Millimho conversion
    millimhosDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceMillimho,
    }
    
    var millimhosResult *units.ElectricConductance
    millimhosResult, err = factory.FromDto(millimhosDto)
    if err != nil {
        t.Errorf("FromDto() with Millimho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimhosResult.Convert(units.ElectricConductanceMillimho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millimho = %v, want %v", converted, 100)
    }
    // Test Kilomho conversion
    kilomhosDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceKilomho,
    }
    
    var kilomhosResult *units.ElectricConductance
    kilomhosResult, err = factory.FromDto(kilomhosDto)
    if err != nil {
        t.Errorf("FromDto() with Kilomho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilomhosResult.Convert(units.ElectricConductanceKilomho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilomho = %v, want %v", converted, 100)
    }
    // Test Megamho conversion
    megamhosDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceMegamho,
    }
    
    var megamhosResult *units.ElectricConductance
    megamhosResult, err = factory.FromDto(megamhosDto)
    if err != nil {
        t.Errorf("FromDto() with Megamho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megamhosResult.Convert(units.ElectricConductanceMegamho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megamho = %v, want %v", converted, 100)
    }
    // Test Gigamho conversion
    gigamhosDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceGigamho,
    }
    
    var gigamhosResult *units.ElectricConductance
    gigamhosResult, err = factory.FromDto(gigamhosDto)
    if err != nil {
        t.Errorf("FromDto() with Gigamho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigamhosResult.Convert(units.ElectricConductanceGigamho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigamho = %v, want %v", converted, 100)
    }
    // Test Teramho conversion
    teramhosDto := units.ElectricConductanceDto{
        Value: 100,
        Unit:  units.ElectricConductanceTeramho,
    }
    
    var teramhosResult *units.ElectricConductance
    teramhosResult, err = factory.FromDto(teramhosDto)
    if err != nil {
        t.Errorf("FromDto() with Teramho returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teramhosResult.Convert(units.ElectricConductanceTeramho)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Teramho = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricConductanceDto{
        Value: 0,
        Unit:  units.ElectricConductanceSiemens,
    }
    
    var zeroResult *units.ElectricConductance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricConductanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
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
    nanJSON, _ := json.Marshal(units.ElectricConductanceDto{
        Value: nanValue,
        Unit:  units.ElectricConductanceSiemens,
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
    converted = siemensResult.Convert(units.ElectricConductanceSiemens)
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
    converted = mhosResult.Convert(units.ElectricConductanceMho)
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
    converted = nanosiemensResult.Convert(units.ElectricConductanceNanosiemens)
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
    converted = microsiemensResult.Convert(units.ElectricConductanceMicrosiemens)
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
    converted = millisiemensResult.Convert(units.ElectricConductanceMillisiemens)
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
    converted = kilosiemensResult.Convert(units.ElectricConductanceKilosiemens)
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
    converted = megasiemensResult.Convert(units.ElectricConductanceMegasiemens)
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
    converted = gigasiemensResult.Convert(units.ElectricConductanceGigasiemens)
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
    converted = terasiemensResult.Convert(units.ElectricConductanceTerasiemens)
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
    converted = nanomhosResult.Convert(units.ElectricConductanceNanomho)
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
    converted = micromhosResult.Convert(units.ElectricConductanceMicromho)
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
    converted = millimhosResult.Convert(units.ElectricConductanceMillimho)
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
    converted = kilomhosResult.Convert(units.ElectricConductanceKilomho)
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
    converted = megamhosResult.Convert(units.ElectricConductanceMegamho)
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
    converted = gigamhosResult.Convert(units.ElectricConductanceGigamho)
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
    converted = teramhosResult.Convert(units.ElectricConductanceTeramho)
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
func TestElectricConductanceFactory_FromSiemens(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSiemens(100)
    if err != nil {
        t.Errorf("FromSiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceSiemens)
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
    converted = zeroResult.Convert(units.ElectricConductanceSiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromMhos function
func TestElectricConductanceFactory_FromMhos(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMhos(100)
    if err != nil {
        t.Errorf("FromMhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceMho)
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
    converted = zeroResult.Convert(units.ElectricConductanceMho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMhos() with zero value = %v, want 0", converted)
    }
}
// Test FromNanosiemens function
func TestElectricConductanceFactory_FromNanosiemens(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanosiemens(100)
    if err != nil {
        t.Errorf("FromNanosiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceNanosiemens)
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
    converted = zeroResult.Convert(units.ElectricConductanceNanosiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanosiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrosiemens function
func TestElectricConductanceFactory_FromMicrosiemens(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrosiemens(100)
    if err != nil {
        t.Errorf("FromMicrosiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceMicrosiemens)
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
    converted = zeroResult.Convert(units.ElectricConductanceMicrosiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrosiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromMillisiemens function
func TestElectricConductanceFactory_FromMillisiemens(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillisiemens(100)
    if err != nil {
        t.Errorf("FromMillisiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceMillisiemens)
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
    converted = zeroResult.Convert(units.ElectricConductanceMillisiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillisiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromKilosiemens function
func TestElectricConductanceFactory_FromKilosiemens(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilosiemens(100)
    if err != nil {
        t.Errorf("FromKilosiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceKilosiemens)
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
    converted = zeroResult.Convert(units.ElectricConductanceKilosiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilosiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromMegasiemens function
func TestElectricConductanceFactory_FromMegasiemens(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegasiemens(100)
    if err != nil {
        t.Errorf("FromMegasiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceMegasiemens)
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
    converted = zeroResult.Convert(units.ElectricConductanceMegasiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegasiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromGigasiemens function
func TestElectricConductanceFactory_FromGigasiemens(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigasiemens(100)
    if err != nil {
        t.Errorf("FromGigasiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceGigasiemens)
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
    converted = zeroResult.Convert(units.ElectricConductanceGigasiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigasiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromTerasiemens function
func TestElectricConductanceFactory_FromTerasiemens(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerasiemens(100)
    if err != nil {
        t.Errorf("FromTerasiemens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceTerasiemens)
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
    converted = zeroResult.Convert(units.ElectricConductanceTerasiemens)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerasiemens() with zero value = %v, want 0", converted)
    }
}
// Test FromNanomhos function
func TestElectricConductanceFactory_FromNanomhos(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanomhos(100)
    if err != nil {
        t.Errorf("FromNanomhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceNanomho)
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
    converted = zeroResult.Convert(units.ElectricConductanceNanomho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanomhos() with zero value = %v, want 0", converted)
    }
}
// Test FromMicromhos function
func TestElectricConductanceFactory_FromMicromhos(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicromhos(100)
    if err != nil {
        t.Errorf("FromMicromhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceMicromho)
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
    converted = zeroResult.Convert(units.ElectricConductanceMicromho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicromhos() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimhos function
func TestElectricConductanceFactory_FromMillimhos(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimhos(100)
    if err != nil {
        t.Errorf("FromMillimhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceMillimho)
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
    converted = zeroResult.Convert(units.ElectricConductanceMillimho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimhos() with zero value = %v, want 0", converted)
    }
}
// Test FromKilomhos function
func TestElectricConductanceFactory_FromKilomhos(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilomhos(100)
    if err != nil {
        t.Errorf("FromKilomhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceKilomho)
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
    converted = zeroResult.Convert(units.ElectricConductanceKilomho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilomhos() with zero value = %v, want 0", converted)
    }
}
// Test FromMegamhos function
func TestElectricConductanceFactory_FromMegamhos(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegamhos(100)
    if err != nil {
        t.Errorf("FromMegamhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceMegamho)
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
    converted = zeroResult.Convert(units.ElectricConductanceMegamho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegamhos() with zero value = %v, want 0", converted)
    }
}
// Test FromGigamhos function
func TestElectricConductanceFactory_FromGigamhos(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigamhos(100)
    if err != nil {
        t.Errorf("FromGigamhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceGigamho)
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
    converted = zeroResult.Convert(units.ElectricConductanceGigamho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigamhos() with zero value = %v, want 0", converted)
    }
}
// Test FromTeramhos function
func TestElectricConductanceFactory_FromTeramhos(t *testing.T) {
    factory := units.ElectricConductanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTeramhos(100)
    if err != nil {
        t.Errorf("FromTeramhos() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductanceTeramho)
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
    converted = zeroResult.Convert(units.ElectricConductanceTeramho)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTeramhos() with zero value = %v, want 0", converted)
    }
}

func TestElectricConductanceToString(t *testing.T) {
	factory := units.ElectricConductanceFactory{}
	a, err := factory.CreateElectricConductance(45, units.ElectricConductanceSiemens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricConductanceSiemens, 2)
	expected := "45.00 " + units.GetElectricConductanceAbbreviation(units.ElectricConductanceSiemens)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricConductanceSiemens, -1)
	expected = "45 " + units.GetElectricConductanceAbbreviation(units.ElectricConductanceSiemens)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricConductance_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricConductanceFactory{}
	a1, _ := factory.CreateElectricConductance(60, units.ElectricConductanceSiemens)
	a2, _ := factory.CreateElectricConductance(60, units.ElectricConductanceSiemens)
	a3, _ := factory.CreateElectricConductance(120, units.ElectricConductanceSiemens)

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

func TestElectricConductance_Arithmetic(t *testing.T) {
	factory := units.ElectricConductanceFactory{}
	a1, _ := factory.CreateElectricConductance(30, units.ElectricConductanceSiemens)
	a2, _ := factory.CreateElectricConductance(45, units.ElectricConductanceSiemens)

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