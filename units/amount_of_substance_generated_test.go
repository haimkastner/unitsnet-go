// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAmountOfSubstanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Mole"}`
	
	factory := units.AmountOfSubstanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AmountOfSubstanceMole {
		t.Errorf("expected unit %v, got %v", units.AmountOfSubstanceMole, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Mole"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAmountOfSubstanceDto_ToJSON(t *testing.T) {
	dto := units.AmountOfSubstanceDto{
		Value: 45,
		Unit:  units.AmountOfSubstanceMole,
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
	if result["unit"].(string) != string(units.AmountOfSubstanceMole) {
		t.Errorf("expected unit %s, got %v", units.AmountOfSubstanceMole, result["unit"])
	}
}

func TestNewAmountOfSubstance_InvalidValue(t *testing.T) {
	factory := units.AmountOfSubstanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAmountOfSubstance(math.NaN(), units.AmountOfSubstanceMole)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAmountOfSubstance(math.Inf(1), units.AmountOfSubstanceMole)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAmountOfSubstanceConversions(t *testing.T) {
	factory := units.AmountOfSubstanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAmountOfSubstance(180, units.AmountOfSubstanceMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Moles.
		// No expected conversion value provided for Moles, verifying result is not NaN.
		result := a.Moles()
		cacheResult := a.Moles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Moles returned NaN")
		}
	}
	{
		// Test conversion to PoundMoles.
		// No expected conversion value provided for PoundMoles, verifying result is not NaN.
		result := a.PoundMoles()
		cacheResult := a.PoundMoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundMoles returned NaN")
		}
	}
	{
		// Test conversion to Femtomoles.
		// No expected conversion value provided for Femtomoles, verifying result is not NaN.
		result := a.Femtomoles()
		cacheResult := a.Femtomoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Femtomoles returned NaN")
		}
	}
	{
		// Test conversion to Picomoles.
		// No expected conversion value provided for Picomoles, verifying result is not NaN.
		result := a.Picomoles()
		cacheResult := a.Picomoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Picomoles returned NaN")
		}
	}
	{
		// Test conversion to Nanomoles.
		// No expected conversion value provided for Nanomoles, verifying result is not NaN.
		result := a.Nanomoles()
		cacheResult := a.Nanomoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanomoles returned NaN")
		}
	}
	{
		// Test conversion to Micromoles.
		// No expected conversion value provided for Micromoles, verifying result is not NaN.
		result := a.Micromoles()
		cacheResult := a.Micromoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Micromoles returned NaN")
		}
	}
	{
		// Test conversion to Millimoles.
		// No expected conversion value provided for Millimoles, verifying result is not NaN.
		result := a.Millimoles()
		cacheResult := a.Millimoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millimoles returned NaN")
		}
	}
	{
		// Test conversion to Centimoles.
		// No expected conversion value provided for Centimoles, verifying result is not NaN.
		result := a.Centimoles()
		cacheResult := a.Centimoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Centimoles returned NaN")
		}
	}
	{
		// Test conversion to Decimoles.
		// No expected conversion value provided for Decimoles, verifying result is not NaN.
		result := a.Decimoles()
		cacheResult := a.Decimoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Decimoles returned NaN")
		}
	}
	{
		// Test conversion to Kilomoles.
		// No expected conversion value provided for Kilomoles, verifying result is not NaN.
		result := a.Kilomoles()
		cacheResult := a.Kilomoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilomoles returned NaN")
		}
	}
	{
		// Test conversion to Megamoles.
		// No expected conversion value provided for Megamoles, verifying result is not NaN.
		result := a.Megamoles()
		cacheResult := a.Megamoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megamoles returned NaN")
		}
	}
	{
		// Test conversion to NanopoundMoles.
		// No expected conversion value provided for NanopoundMoles, verifying result is not NaN.
		result := a.NanopoundMoles()
		cacheResult := a.NanopoundMoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanopoundMoles returned NaN")
		}
	}
	{
		// Test conversion to MicropoundMoles.
		// No expected conversion value provided for MicropoundMoles, verifying result is not NaN.
		result := a.MicropoundMoles()
		cacheResult := a.MicropoundMoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicropoundMoles returned NaN")
		}
	}
	{
		// Test conversion to MillipoundMoles.
		// No expected conversion value provided for MillipoundMoles, verifying result is not NaN.
		result := a.MillipoundMoles()
		cacheResult := a.MillipoundMoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillipoundMoles returned NaN")
		}
	}
	{
		// Test conversion to CentipoundMoles.
		// No expected conversion value provided for CentipoundMoles, verifying result is not NaN.
		result := a.CentipoundMoles()
		cacheResult := a.CentipoundMoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentipoundMoles returned NaN")
		}
	}
	{
		// Test conversion to DecipoundMoles.
		// No expected conversion value provided for DecipoundMoles, verifying result is not NaN.
		result := a.DecipoundMoles()
		cacheResult := a.DecipoundMoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecipoundMoles returned NaN")
		}
	}
	{
		// Test conversion to KilopoundMoles.
		// No expected conversion value provided for KilopoundMoles, verifying result is not NaN.
		result := a.KilopoundMoles()
		cacheResult := a.KilopoundMoles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundMoles returned NaN")
		}
	}
}

func TestAmountOfSubstance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AmountOfSubstanceFactory{}
	a, err := factory.CreateAmountOfSubstance(90, units.AmountOfSubstanceMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AmountOfSubstanceMole {
		t.Errorf("expected default unit Mole, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AmountOfSubstanceMole
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AmountOfSubstanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AmountOfSubstanceMole {
		t.Errorf("expected unit Mole, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAmountOfSubstanceFactory_FromDto(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceMole,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.AmountOfSubstanceDto{
        Value: math.NaN(),
        Unit:  units.AmountOfSubstanceMole,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Mole conversion
    molesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceMole,
    }
    
    var molesResult *units.AmountOfSubstance
    molesResult, err = factory.FromDto(molesDto)
    if err != nil {
        t.Errorf("FromDto() with Mole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = molesResult.Convert(units.AmountOfSubstanceMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mole = %v, want %v", converted, 100)
    }
    // Test PoundMole conversion
    pound_molesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstancePoundMole,
    }
    
    var pound_molesResult *units.AmountOfSubstance
    pound_molesResult, err = factory.FromDto(pound_molesDto)
    if err != nil {
        t.Errorf("FromDto() with PoundMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_molesResult.Convert(units.AmountOfSubstancePoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundMole = %v, want %v", converted, 100)
    }
    // Test Femtomole conversion
    femtomolesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceFemtomole,
    }
    
    var femtomolesResult *units.AmountOfSubstance
    femtomolesResult, err = factory.FromDto(femtomolesDto)
    if err != nil {
        t.Errorf("FromDto() with Femtomole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtomolesResult.Convert(units.AmountOfSubstanceFemtomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtomole = %v, want %v", converted, 100)
    }
    // Test Picomole conversion
    picomolesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstancePicomole,
    }
    
    var picomolesResult *units.AmountOfSubstance
    picomolesResult, err = factory.FromDto(picomolesDto)
    if err != nil {
        t.Errorf("FromDto() with Picomole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picomolesResult.Convert(units.AmountOfSubstancePicomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picomole = %v, want %v", converted, 100)
    }
    // Test Nanomole conversion
    nanomolesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceNanomole,
    }
    
    var nanomolesResult *units.AmountOfSubstance
    nanomolesResult, err = factory.FromDto(nanomolesDto)
    if err != nil {
        t.Errorf("FromDto() with Nanomole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanomolesResult.Convert(units.AmountOfSubstanceNanomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanomole = %v, want %v", converted, 100)
    }
    // Test Micromole conversion
    micromolesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceMicromole,
    }
    
    var micromolesResult *units.AmountOfSubstance
    micromolesResult, err = factory.FromDto(micromolesDto)
    if err != nil {
        t.Errorf("FromDto() with Micromole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micromolesResult.Convert(units.AmountOfSubstanceMicromole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Micromole = %v, want %v", converted, 100)
    }
    // Test Millimole conversion
    millimolesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceMillimole,
    }
    
    var millimolesResult *units.AmountOfSubstance
    millimolesResult, err = factory.FromDto(millimolesDto)
    if err != nil {
        t.Errorf("FromDto() with Millimole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimolesResult.Convert(units.AmountOfSubstanceMillimole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millimole = %v, want %v", converted, 100)
    }
    // Test Centimole conversion
    centimolesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceCentimole,
    }
    
    var centimolesResult *units.AmountOfSubstance
    centimolesResult, err = factory.FromDto(centimolesDto)
    if err != nil {
        t.Errorf("FromDto() with Centimole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimolesResult.Convert(units.AmountOfSubstanceCentimole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centimole = %v, want %v", converted, 100)
    }
    // Test Decimole conversion
    decimolesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceDecimole,
    }
    
    var decimolesResult *units.AmountOfSubstance
    decimolesResult, err = factory.FromDto(decimolesDto)
    if err != nil {
        t.Errorf("FromDto() with Decimole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimolesResult.Convert(units.AmountOfSubstanceDecimole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decimole = %v, want %v", converted, 100)
    }
    // Test Kilomole conversion
    kilomolesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceKilomole,
    }
    
    var kilomolesResult *units.AmountOfSubstance
    kilomolesResult, err = factory.FromDto(kilomolesDto)
    if err != nil {
        t.Errorf("FromDto() with Kilomole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilomolesResult.Convert(units.AmountOfSubstanceKilomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilomole = %v, want %v", converted, 100)
    }
    // Test Megamole conversion
    megamolesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceMegamole,
    }
    
    var megamolesResult *units.AmountOfSubstance
    megamolesResult, err = factory.FromDto(megamolesDto)
    if err != nil {
        t.Errorf("FromDto() with Megamole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megamolesResult.Convert(units.AmountOfSubstanceMegamole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megamole = %v, want %v", converted, 100)
    }
    // Test NanopoundMole conversion
    nanopound_molesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceNanopoundMole,
    }
    
    var nanopound_molesResult *units.AmountOfSubstance
    nanopound_molesResult, err = factory.FromDto(nanopound_molesDto)
    if err != nil {
        t.Errorf("FromDto() with NanopoundMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanopound_molesResult.Convert(units.AmountOfSubstanceNanopoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanopoundMole = %v, want %v", converted, 100)
    }
    // Test MicropoundMole conversion
    micropound_molesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceMicropoundMole,
    }
    
    var micropound_molesResult *units.AmountOfSubstance
    micropound_molesResult, err = factory.FromDto(micropound_molesDto)
    if err != nil {
        t.Errorf("FromDto() with MicropoundMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micropound_molesResult.Convert(units.AmountOfSubstanceMicropoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicropoundMole = %v, want %v", converted, 100)
    }
    // Test MillipoundMole conversion
    millipound_molesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceMillipoundMole,
    }
    
    var millipound_molesResult *units.AmountOfSubstance
    millipound_molesResult, err = factory.FromDto(millipound_molesDto)
    if err != nil {
        t.Errorf("FromDto() with MillipoundMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millipound_molesResult.Convert(units.AmountOfSubstanceMillipoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillipoundMole = %v, want %v", converted, 100)
    }
    // Test CentipoundMole conversion
    centipound_molesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceCentipoundMole,
    }
    
    var centipound_molesResult *units.AmountOfSubstance
    centipound_molesResult, err = factory.FromDto(centipound_molesDto)
    if err != nil {
        t.Errorf("FromDto() with CentipoundMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centipound_molesResult.Convert(units.AmountOfSubstanceCentipoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentipoundMole = %v, want %v", converted, 100)
    }
    // Test DecipoundMole conversion
    decipound_molesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceDecipoundMole,
    }
    
    var decipound_molesResult *units.AmountOfSubstance
    decipound_molesResult, err = factory.FromDto(decipound_molesDto)
    if err != nil {
        t.Errorf("FromDto() with DecipoundMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decipound_molesResult.Convert(units.AmountOfSubstanceDecipoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecipoundMole = %v, want %v", converted, 100)
    }
    // Test KilopoundMole conversion
    kilopound_molesDto := units.AmountOfSubstanceDto{
        Value: 100,
        Unit:  units.AmountOfSubstanceKilopoundMole,
    }
    
    var kilopound_molesResult *units.AmountOfSubstance
    kilopound_molesResult, err = factory.FromDto(kilopound_molesDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_molesResult.Convert(units.AmountOfSubstanceKilopoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundMole = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.AmountOfSubstanceDto{
        Value: 0,
        Unit:  units.AmountOfSubstanceMole,
    }
    
    var zeroResult *units.AmountOfSubstance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestAmountOfSubstanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Mole"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Mole"}`)
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
    nanJSON, _ := json.Marshal(units.AmountOfSubstanceDto{
        Value: nanValue,
        Unit:  units.AmountOfSubstanceMole,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Mole unit
    molesJSON := []byte(`{"value": 100, "unit": "Mole"}`)
    molesResult, err := factory.FromDtoJSON(molesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Mole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = molesResult.Convert(units.AmountOfSubstanceMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mole = %v, want %v", converted, 100)
    }
    // Test JSON with PoundMole unit
    pound_molesJSON := []byte(`{"value": 100, "unit": "PoundMole"}`)
    pound_molesResult, err := factory.FromDtoJSON(pound_molesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_molesResult.Convert(units.AmountOfSubstancePoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundMole = %v, want %v", converted, 100)
    }
    // Test JSON with Femtomole unit
    femtomolesJSON := []byte(`{"value": 100, "unit": "Femtomole"}`)
    femtomolesResult, err := factory.FromDtoJSON(femtomolesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Femtomole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtomolesResult.Convert(units.AmountOfSubstanceFemtomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtomole = %v, want %v", converted, 100)
    }
    // Test JSON with Picomole unit
    picomolesJSON := []byte(`{"value": 100, "unit": "Picomole"}`)
    picomolesResult, err := factory.FromDtoJSON(picomolesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Picomole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picomolesResult.Convert(units.AmountOfSubstancePicomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picomole = %v, want %v", converted, 100)
    }
    // Test JSON with Nanomole unit
    nanomolesJSON := []byte(`{"value": 100, "unit": "Nanomole"}`)
    nanomolesResult, err := factory.FromDtoJSON(nanomolesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanomole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanomolesResult.Convert(units.AmountOfSubstanceNanomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanomole = %v, want %v", converted, 100)
    }
    // Test JSON with Micromole unit
    micromolesJSON := []byte(`{"value": 100, "unit": "Micromole"}`)
    micromolesResult, err := factory.FromDtoJSON(micromolesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Micromole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micromolesResult.Convert(units.AmountOfSubstanceMicromole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Micromole = %v, want %v", converted, 100)
    }
    // Test JSON with Millimole unit
    millimolesJSON := []byte(`{"value": 100, "unit": "Millimole"}`)
    millimolesResult, err := factory.FromDtoJSON(millimolesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millimole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimolesResult.Convert(units.AmountOfSubstanceMillimole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millimole = %v, want %v", converted, 100)
    }
    // Test JSON with Centimole unit
    centimolesJSON := []byte(`{"value": 100, "unit": "Centimole"}`)
    centimolesResult, err := factory.FromDtoJSON(centimolesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Centimole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimolesResult.Convert(units.AmountOfSubstanceCentimole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centimole = %v, want %v", converted, 100)
    }
    // Test JSON with Decimole unit
    decimolesJSON := []byte(`{"value": 100, "unit": "Decimole"}`)
    decimolesResult, err := factory.FromDtoJSON(decimolesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Decimole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimolesResult.Convert(units.AmountOfSubstanceDecimole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decimole = %v, want %v", converted, 100)
    }
    // Test JSON with Kilomole unit
    kilomolesJSON := []byte(`{"value": 100, "unit": "Kilomole"}`)
    kilomolesResult, err := factory.FromDtoJSON(kilomolesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilomole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilomolesResult.Convert(units.AmountOfSubstanceKilomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilomole = %v, want %v", converted, 100)
    }
    // Test JSON with Megamole unit
    megamolesJSON := []byte(`{"value": 100, "unit": "Megamole"}`)
    megamolesResult, err := factory.FromDtoJSON(megamolesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megamole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megamolesResult.Convert(units.AmountOfSubstanceMegamole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megamole = %v, want %v", converted, 100)
    }
    // Test JSON with NanopoundMole unit
    nanopound_molesJSON := []byte(`{"value": 100, "unit": "NanopoundMole"}`)
    nanopound_molesResult, err := factory.FromDtoJSON(nanopound_molesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanopoundMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanopound_molesResult.Convert(units.AmountOfSubstanceNanopoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanopoundMole = %v, want %v", converted, 100)
    }
    // Test JSON with MicropoundMole unit
    micropound_molesJSON := []byte(`{"value": 100, "unit": "MicropoundMole"}`)
    micropound_molesResult, err := factory.FromDtoJSON(micropound_molesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicropoundMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micropound_molesResult.Convert(units.AmountOfSubstanceMicropoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicropoundMole = %v, want %v", converted, 100)
    }
    // Test JSON with MillipoundMole unit
    millipound_molesJSON := []byte(`{"value": 100, "unit": "MillipoundMole"}`)
    millipound_molesResult, err := factory.FromDtoJSON(millipound_molesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillipoundMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millipound_molesResult.Convert(units.AmountOfSubstanceMillipoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillipoundMole = %v, want %v", converted, 100)
    }
    // Test JSON with CentipoundMole unit
    centipound_molesJSON := []byte(`{"value": 100, "unit": "CentipoundMole"}`)
    centipound_molesResult, err := factory.FromDtoJSON(centipound_molesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentipoundMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centipound_molesResult.Convert(units.AmountOfSubstanceCentipoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentipoundMole = %v, want %v", converted, 100)
    }
    // Test JSON with DecipoundMole unit
    decipound_molesJSON := []byte(`{"value": 100, "unit": "DecipoundMole"}`)
    decipound_molesResult, err := factory.FromDtoJSON(decipound_molesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecipoundMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decipound_molesResult.Convert(units.AmountOfSubstanceDecipoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecipoundMole = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundMole unit
    kilopound_molesJSON := []byte(`{"value": 100, "unit": "KilopoundMole"}`)
    kilopound_molesResult, err := factory.FromDtoJSON(kilopound_molesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_molesResult.Convert(units.AmountOfSubstanceKilopoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundMole = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Mole"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromMoles function
func TestAmountOfSubstanceFactory_FromMoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMoles(100)
    if err != nil {
        t.Errorf("FromMoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMoles(math.NaN())
    if err == nil {
        t.Error("FromMoles() with NaN value should return error")
    }

    _, err = factory.FromMoles(math.Inf(1))
    if err == nil {
        t.Error("FromMoles() with +Inf value should return error")
    }

    _, err = factory.FromMoles(math.Inf(-1))
    if err == nil {
        t.Error("FromMoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMoles(0)
    if err != nil {
        t.Errorf("FromMoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMoles() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundMoles function
func TestAmountOfSubstanceFactory_FromPoundMoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundMoles(100)
    if err != nil {
        t.Errorf("FromPoundMoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstancePoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundMoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundMoles(math.NaN())
    if err == nil {
        t.Error("FromPoundMoles() with NaN value should return error")
    }

    _, err = factory.FromPoundMoles(math.Inf(1))
    if err == nil {
        t.Error("FromPoundMoles() with +Inf value should return error")
    }

    _, err = factory.FromPoundMoles(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundMoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundMoles(0)
    if err != nil {
        t.Errorf("FromPoundMoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstancePoundMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundMoles() with zero value = %v, want 0", converted)
    }
}
// Test FromFemtomoles function
func TestAmountOfSubstanceFactory_FromFemtomoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFemtomoles(100)
    if err != nil {
        t.Errorf("FromFemtomoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceFemtomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFemtomoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFemtomoles(math.NaN())
    if err == nil {
        t.Error("FromFemtomoles() with NaN value should return error")
    }

    _, err = factory.FromFemtomoles(math.Inf(1))
    if err == nil {
        t.Error("FromFemtomoles() with +Inf value should return error")
    }

    _, err = factory.FromFemtomoles(math.Inf(-1))
    if err == nil {
        t.Error("FromFemtomoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFemtomoles(0)
    if err != nil {
        t.Errorf("FromFemtomoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceFemtomole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFemtomoles() with zero value = %v, want 0", converted)
    }
}
// Test FromPicomoles function
func TestAmountOfSubstanceFactory_FromPicomoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicomoles(100)
    if err != nil {
        t.Errorf("FromPicomoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstancePicomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicomoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicomoles(math.NaN())
    if err == nil {
        t.Error("FromPicomoles() with NaN value should return error")
    }

    _, err = factory.FromPicomoles(math.Inf(1))
    if err == nil {
        t.Error("FromPicomoles() with +Inf value should return error")
    }

    _, err = factory.FromPicomoles(math.Inf(-1))
    if err == nil {
        t.Error("FromPicomoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicomoles(0)
    if err != nil {
        t.Errorf("FromPicomoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstancePicomole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicomoles() with zero value = %v, want 0", converted)
    }
}
// Test FromNanomoles function
func TestAmountOfSubstanceFactory_FromNanomoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanomoles(100)
    if err != nil {
        t.Errorf("FromNanomoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceNanomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanomoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanomoles(math.NaN())
    if err == nil {
        t.Error("FromNanomoles() with NaN value should return error")
    }

    _, err = factory.FromNanomoles(math.Inf(1))
    if err == nil {
        t.Error("FromNanomoles() with +Inf value should return error")
    }

    _, err = factory.FromNanomoles(math.Inf(-1))
    if err == nil {
        t.Error("FromNanomoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanomoles(0)
    if err != nil {
        t.Errorf("FromNanomoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceNanomole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanomoles() with zero value = %v, want 0", converted)
    }
}
// Test FromMicromoles function
func TestAmountOfSubstanceFactory_FromMicromoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicromoles(100)
    if err != nil {
        t.Errorf("FromMicromoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceMicromole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicromoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicromoles(math.NaN())
    if err == nil {
        t.Error("FromMicromoles() with NaN value should return error")
    }

    _, err = factory.FromMicromoles(math.Inf(1))
    if err == nil {
        t.Error("FromMicromoles() with +Inf value should return error")
    }

    _, err = factory.FromMicromoles(math.Inf(-1))
    if err == nil {
        t.Error("FromMicromoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicromoles(0)
    if err != nil {
        t.Errorf("FromMicromoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceMicromole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicromoles() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimoles function
func TestAmountOfSubstanceFactory_FromMillimoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimoles(100)
    if err != nil {
        t.Errorf("FromMillimoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceMillimole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimoles(math.NaN())
    if err == nil {
        t.Error("FromMillimoles() with NaN value should return error")
    }

    _, err = factory.FromMillimoles(math.Inf(1))
    if err == nil {
        t.Error("FromMillimoles() with +Inf value should return error")
    }

    _, err = factory.FromMillimoles(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimoles(0)
    if err != nil {
        t.Errorf("FromMillimoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceMillimole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimoles() with zero value = %v, want 0", converted)
    }
}
// Test FromCentimoles function
func TestAmountOfSubstanceFactory_FromCentimoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentimoles(100)
    if err != nil {
        t.Errorf("FromCentimoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceCentimole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentimoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentimoles(math.NaN())
    if err == nil {
        t.Error("FromCentimoles() with NaN value should return error")
    }

    _, err = factory.FromCentimoles(math.Inf(1))
    if err == nil {
        t.Error("FromCentimoles() with +Inf value should return error")
    }

    _, err = factory.FromCentimoles(math.Inf(-1))
    if err == nil {
        t.Error("FromCentimoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentimoles(0)
    if err != nil {
        t.Errorf("FromCentimoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceCentimole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentimoles() with zero value = %v, want 0", converted)
    }
}
// Test FromDecimoles function
func TestAmountOfSubstanceFactory_FromDecimoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecimoles(100)
    if err != nil {
        t.Errorf("FromDecimoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceDecimole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecimoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecimoles(math.NaN())
    if err == nil {
        t.Error("FromDecimoles() with NaN value should return error")
    }

    _, err = factory.FromDecimoles(math.Inf(1))
    if err == nil {
        t.Error("FromDecimoles() with +Inf value should return error")
    }

    _, err = factory.FromDecimoles(math.Inf(-1))
    if err == nil {
        t.Error("FromDecimoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecimoles(0)
    if err != nil {
        t.Errorf("FromDecimoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceDecimole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecimoles() with zero value = %v, want 0", converted)
    }
}
// Test FromKilomoles function
func TestAmountOfSubstanceFactory_FromKilomoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilomoles(100)
    if err != nil {
        t.Errorf("FromKilomoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceKilomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilomoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilomoles(math.NaN())
    if err == nil {
        t.Error("FromKilomoles() with NaN value should return error")
    }

    _, err = factory.FromKilomoles(math.Inf(1))
    if err == nil {
        t.Error("FromKilomoles() with +Inf value should return error")
    }

    _, err = factory.FromKilomoles(math.Inf(-1))
    if err == nil {
        t.Error("FromKilomoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilomoles(0)
    if err != nil {
        t.Errorf("FromKilomoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceKilomole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilomoles() with zero value = %v, want 0", converted)
    }
}
// Test FromMegamoles function
func TestAmountOfSubstanceFactory_FromMegamoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegamoles(100)
    if err != nil {
        t.Errorf("FromMegamoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceMegamole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegamoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegamoles(math.NaN())
    if err == nil {
        t.Error("FromMegamoles() with NaN value should return error")
    }

    _, err = factory.FromMegamoles(math.Inf(1))
    if err == nil {
        t.Error("FromMegamoles() with +Inf value should return error")
    }

    _, err = factory.FromMegamoles(math.Inf(-1))
    if err == nil {
        t.Error("FromMegamoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegamoles(0)
    if err != nil {
        t.Errorf("FromMegamoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceMegamole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegamoles() with zero value = %v, want 0", converted)
    }
}
// Test FromNanopoundMoles function
func TestAmountOfSubstanceFactory_FromNanopoundMoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanopoundMoles(100)
    if err != nil {
        t.Errorf("FromNanopoundMoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceNanopoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanopoundMoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanopoundMoles(math.NaN())
    if err == nil {
        t.Error("FromNanopoundMoles() with NaN value should return error")
    }

    _, err = factory.FromNanopoundMoles(math.Inf(1))
    if err == nil {
        t.Error("FromNanopoundMoles() with +Inf value should return error")
    }

    _, err = factory.FromNanopoundMoles(math.Inf(-1))
    if err == nil {
        t.Error("FromNanopoundMoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanopoundMoles(0)
    if err != nil {
        t.Errorf("FromNanopoundMoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceNanopoundMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanopoundMoles() with zero value = %v, want 0", converted)
    }
}
// Test FromMicropoundMoles function
func TestAmountOfSubstanceFactory_FromMicropoundMoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicropoundMoles(100)
    if err != nil {
        t.Errorf("FromMicropoundMoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceMicropoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicropoundMoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicropoundMoles(math.NaN())
    if err == nil {
        t.Error("FromMicropoundMoles() with NaN value should return error")
    }

    _, err = factory.FromMicropoundMoles(math.Inf(1))
    if err == nil {
        t.Error("FromMicropoundMoles() with +Inf value should return error")
    }

    _, err = factory.FromMicropoundMoles(math.Inf(-1))
    if err == nil {
        t.Error("FromMicropoundMoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicropoundMoles(0)
    if err != nil {
        t.Errorf("FromMicropoundMoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceMicropoundMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicropoundMoles() with zero value = %v, want 0", converted)
    }
}
// Test FromMillipoundMoles function
func TestAmountOfSubstanceFactory_FromMillipoundMoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillipoundMoles(100)
    if err != nil {
        t.Errorf("FromMillipoundMoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceMillipoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillipoundMoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillipoundMoles(math.NaN())
    if err == nil {
        t.Error("FromMillipoundMoles() with NaN value should return error")
    }

    _, err = factory.FromMillipoundMoles(math.Inf(1))
    if err == nil {
        t.Error("FromMillipoundMoles() with +Inf value should return error")
    }

    _, err = factory.FromMillipoundMoles(math.Inf(-1))
    if err == nil {
        t.Error("FromMillipoundMoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillipoundMoles(0)
    if err != nil {
        t.Errorf("FromMillipoundMoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceMillipoundMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillipoundMoles() with zero value = %v, want 0", converted)
    }
}
// Test FromCentipoundMoles function
func TestAmountOfSubstanceFactory_FromCentipoundMoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentipoundMoles(100)
    if err != nil {
        t.Errorf("FromCentipoundMoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceCentipoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentipoundMoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentipoundMoles(math.NaN())
    if err == nil {
        t.Error("FromCentipoundMoles() with NaN value should return error")
    }

    _, err = factory.FromCentipoundMoles(math.Inf(1))
    if err == nil {
        t.Error("FromCentipoundMoles() with +Inf value should return error")
    }

    _, err = factory.FromCentipoundMoles(math.Inf(-1))
    if err == nil {
        t.Error("FromCentipoundMoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentipoundMoles(0)
    if err != nil {
        t.Errorf("FromCentipoundMoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceCentipoundMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentipoundMoles() with zero value = %v, want 0", converted)
    }
}
// Test FromDecipoundMoles function
func TestAmountOfSubstanceFactory_FromDecipoundMoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecipoundMoles(100)
    if err != nil {
        t.Errorf("FromDecipoundMoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceDecipoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecipoundMoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecipoundMoles(math.NaN())
    if err == nil {
        t.Error("FromDecipoundMoles() with NaN value should return error")
    }

    _, err = factory.FromDecipoundMoles(math.Inf(1))
    if err == nil {
        t.Error("FromDecipoundMoles() with +Inf value should return error")
    }

    _, err = factory.FromDecipoundMoles(math.Inf(-1))
    if err == nil {
        t.Error("FromDecipoundMoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecipoundMoles(0)
    if err != nil {
        t.Errorf("FromDecipoundMoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceDecipoundMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecipoundMoles() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundMoles function
func TestAmountOfSubstanceFactory_FromKilopoundMoles(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundMoles(100)
    if err != nil {
        t.Errorf("FromKilopoundMoles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AmountOfSubstanceKilopoundMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundMoles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundMoles(math.NaN())
    if err == nil {
        t.Error("FromKilopoundMoles() with NaN value should return error")
    }

    _, err = factory.FromKilopoundMoles(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundMoles() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundMoles(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundMoles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundMoles(0)
    if err != nil {
        t.Errorf("FromKilopoundMoles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AmountOfSubstanceKilopoundMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundMoles() with zero value = %v, want 0", converted)
    }
}

func TestAmountOfSubstanceToString(t *testing.T) {
	factory := units.AmountOfSubstanceFactory{}
	a, err := factory.CreateAmountOfSubstance(45, units.AmountOfSubstanceMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AmountOfSubstanceMole, 2)
	expected := "45.00 " + units.GetAmountOfSubstanceAbbreviation(units.AmountOfSubstanceMole)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AmountOfSubstanceMole, -1)
	expected = "45 " + units.GetAmountOfSubstanceAbbreviation(units.AmountOfSubstanceMole)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAmountOfSubstance_EqualityAndComparison(t *testing.T) {
	factory := units.AmountOfSubstanceFactory{}
	a1, _ := factory.CreateAmountOfSubstance(60, units.AmountOfSubstanceMole)
	a2, _ := factory.CreateAmountOfSubstance(60, units.AmountOfSubstanceMole)
	a3, _ := factory.CreateAmountOfSubstance(120, units.AmountOfSubstanceMole)

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

func TestAmountOfSubstance_Arithmetic(t *testing.T) {
	factory := units.AmountOfSubstanceFactory{}
	a1, _ := factory.CreateAmountOfSubstance(30, units.AmountOfSubstanceMole)
	a2, _ := factory.CreateAmountOfSubstance(45, units.AmountOfSubstanceMole)

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


func TestGetAmountOfSubstanceAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.AmountOfSubstanceUnits
        want string
    }{
        {
            name: "Mole abbreviation",
            unit: units.AmountOfSubstanceMole,
            want: "mol",
        },
        {
            name: "PoundMole abbreviation",
            unit: units.AmountOfSubstancePoundMole,
            want: "lbmol",
        },
        {
            name: "Femtomole abbreviation",
            unit: units.AmountOfSubstanceFemtomole,
            want: "fmol",
        },
        {
            name: "Picomole abbreviation",
            unit: units.AmountOfSubstancePicomole,
            want: "pmol",
        },
        {
            name: "Nanomole abbreviation",
            unit: units.AmountOfSubstanceNanomole,
            want: "nmol",
        },
        {
            name: "Micromole abbreviation",
            unit: units.AmountOfSubstanceMicromole,
            want: "μmol",
        },
        {
            name: "Millimole abbreviation",
            unit: units.AmountOfSubstanceMillimole,
            want: "mmol",
        },
        {
            name: "Centimole abbreviation",
            unit: units.AmountOfSubstanceCentimole,
            want: "cmol",
        },
        {
            name: "Decimole abbreviation",
            unit: units.AmountOfSubstanceDecimole,
            want: "dmol",
        },
        {
            name: "Kilomole abbreviation",
            unit: units.AmountOfSubstanceKilomole,
            want: "kmol",
        },
        {
            name: "Megamole abbreviation",
            unit: units.AmountOfSubstanceMegamole,
            want: "Mmol",
        },
        {
            name: "NanopoundMole abbreviation",
            unit: units.AmountOfSubstanceNanopoundMole,
            want: "nlbmol",
        },
        {
            name: "MicropoundMole abbreviation",
            unit: units.AmountOfSubstanceMicropoundMole,
            want: "μlbmol",
        },
        {
            name: "MillipoundMole abbreviation",
            unit: units.AmountOfSubstanceMillipoundMole,
            want: "mlbmol",
        },
        {
            name: "CentipoundMole abbreviation",
            unit: units.AmountOfSubstanceCentipoundMole,
            want: "clbmol",
        },
        {
            name: "DecipoundMole abbreviation",
            unit: units.AmountOfSubstanceDecipoundMole,
            want: "dlbmol",
        },
        {
            name: "KilopoundMole abbreviation",
            unit: units.AmountOfSubstanceKilopoundMole,
            want: "klbmol",
        },
        {
            name: "invalid unit",
            unit: units.AmountOfSubstanceUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetAmountOfSubstanceAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetAmountOfSubstanceAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestAmountOfSubstance_String(t *testing.T) {
    factory := units.AmountOfSubstanceFactory{}
    
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
            unit, err := factory.CreateAmountOfSubstance(tt.value, units.AmountOfSubstanceMole)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("AmountOfSubstance.String() = %v, want %v", got, tt.want)
            }
        })
    }
}