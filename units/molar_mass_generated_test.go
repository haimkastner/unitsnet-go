// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMolarMassDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerMole"}`
	
	factory := units.MolarMassDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MolarMassKilogramPerMole {
		t.Errorf("expected unit %v, got %v", units.MolarMassKilogramPerMole, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerMole"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMolarMassDto_ToJSON(t *testing.T) {
	dto := units.MolarMassDto{
		Value: 45,
		Unit:  units.MolarMassKilogramPerMole,
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
	if result["unit"].(string) != string(units.MolarMassKilogramPerMole) {
		t.Errorf("expected unit %s, got %v", units.MolarMassKilogramPerMole, result["unit"])
	}
}

func TestNewMolarMass_InvalidValue(t *testing.T) {
	factory := units.MolarMassFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMolarMass(math.NaN(), units.MolarMassKilogramPerMole)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMolarMass(math.Inf(1), units.MolarMassKilogramPerMole)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMolarMassConversions(t *testing.T) {
	factory := units.MolarMassFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMolarMass(180, units.MolarMassKilogramPerMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramsPerMole.
		// No expected conversion value provided for GramsPerMole, verifying result is not NaN.
		result := a.GramsPerMole()
		cacheResult := a.GramsPerMole()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerKilomole.
		// No expected conversion value provided for KilogramsPerKilomole, verifying result is not NaN.
		result := a.KilogramsPerKilomole()
		cacheResult := a.KilogramsPerKilomole()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsPerKilomole returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerMole.
		// No expected conversion value provided for PoundsPerMole, verifying result is not NaN.
		result := a.PoundsPerMole()
		cacheResult := a.PoundsPerMole()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsPerMole returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerMole.
		// No expected conversion value provided for NanogramsPerMole, verifying result is not NaN.
		result := a.NanogramsPerMole()
		cacheResult := a.NanogramsPerMole()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanogramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerMole.
		// No expected conversion value provided for MicrogramsPerMole, verifying result is not NaN.
		result := a.MicrogramsPerMole()
		cacheResult := a.MicrogramsPerMole()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrogramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerMole.
		// No expected conversion value provided for MilligramsPerMole, verifying result is not NaN.
		result := a.MilligramsPerMole()
		cacheResult := a.MilligramsPerMole()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerMole.
		// No expected conversion value provided for CentigramsPerMole, verifying result is not NaN.
		result := a.CentigramsPerMole()
		cacheResult := a.CentigramsPerMole()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentigramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerMole.
		// No expected conversion value provided for DecigramsPerMole, verifying result is not NaN.
		result := a.DecigramsPerMole()
		cacheResult := a.DecigramsPerMole()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecigramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to DecagramsPerMole.
		// No expected conversion value provided for DecagramsPerMole, verifying result is not NaN.
		result := a.DecagramsPerMole()
		cacheResult := a.DecagramsPerMole()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecagramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to HectogramsPerMole.
		// No expected conversion value provided for HectogramsPerMole, verifying result is not NaN.
		result := a.HectogramsPerMole()
		cacheResult := a.HectogramsPerMole()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to HectogramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerMole.
		// No expected conversion value provided for KilogramsPerMole, verifying result is not NaN.
		result := a.KilogramsPerMole()
		cacheResult := a.KilogramsPerMole()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsPerMole returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsPerMole.
		// No expected conversion value provided for KilopoundsPerMole, verifying result is not NaN.
		result := a.KilopoundsPerMole()
		cacheResult := a.KilopoundsPerMole()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsPerMole returned NaN")
		}
	}
	{
		// Test conversion to MegapoundsPerMole.
		// No expected conversion value provided for MegapoundsPerMole, verifying result is not NaN.
		result := a.MegapoundsPerMole()
		cacheResult := a.MegapoundsPerMole()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegapoundsPerMole returned NaN")
		}
	}
}

func TestMolarMass_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MolarMassFactory{}
	a, err := factory.CreateMolarMass(90, units.MolarMassKilogramPerMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MolarMassKilogramPerMole {
		t.Errorf("expected default unit KilogramPerMole, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MolarMassGramPerMole
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MolarMassDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MolarMassKilogramPerMole {
		t.Errorf("expected unit KilogramPerMole, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMolarMassFactory_FromDto(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassKilogramPerMole,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MolarMassDto{
        Value: math.NaN(),
        Unit:  units.MolarMassKilogramPerMole,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test GramPerMole conversion
    grams_per_moleDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassGramPerMole,
    }
    
    var grams_per_moleResult *units.MolarMass
    grams_per_moleResult, err = factory.FromDto(grams_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_moleResult.Convert(units.MolarMassGramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerMole = %v, want %v", converted, 100)
    }
    // Test KilogramPerKilomole conversion
    kilograms_per_kilomoleDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassKilogramPerKilomole,
    }
    
    var kilograms_per_kilomoleResult *units.MolarMass
    kilograms_per_kilomoleResult, err = factory.FromDto(kilograms_per_kilomoleDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerKilomole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_kilomoleResult.Convert(units.MolarMassKilogramPerKilomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerKilomole = %v, want %v", converted, 100)
    }
    // Test PoundPerMole conversion
    pounds_per_moleDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassPoundPerMole,
    }
    
    var pounds_per_moleResult *units.MolarMass
    pounds_per_moleResult, err = factory.FromDto(pounds_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_moleResult.Convert(units.MolarMassPoundPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerMole = %v, want %v", converted, 100)
    }
    // Test NanogramPerMole conversion
    nanograms_per_moleDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassNanogramPerMole,
    }
    
    var nanograms_per_moleResult *units.MolarMass
    nanograms_per_moleResult, err = factory.FromDto(nanograms_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with NanogramPerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_moleResult.Convert(units.MolarMassNanogramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerMole = %v, want %v", converted, 100)
    }
    // Test MicrogramPerMole conversion
    micrograms_per_moleDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassMicrogramPerMole,
    }
    
    var micrograms_per_moleResult *units.MolarMass
    micrograms_per_moleResult, err = factory.FromDto(micrograms_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_moleResult.Convert(units.MolarMassMicrogramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerMole = %v, want %v", converted, 100)
    }
    // Test MilligramPerMole conversion
    milligrams_per_moleDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassMilligramPerMole,
    }
    
    var milligrams_per_moleResult *units.MolarMass
    milligrams_per_moleResult, err = factory.FromDto(milligrams_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_moleResult.Convert(units.MolarMassMilligramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerMole = %v, want %v", converted, 100)
    }
    // Test CentigramPerMole conversion
    centigrams_per_moleDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassCentigramPerMole,
    }
    
    var centigrams_per_moleResult *units.MolarMass
    centigrams_per_moleResult, err = factory.FromDto(centigrams_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with CentigramPerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_moleResult.Convert(units.MolarMassCentigramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerMole = %v, want %v", converted, 100)
    }
    // Test DecigramPerMole conversion
    decigrams_per_moleDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassDecigramPerMole,
    }
    
    var decigrams_per_moleResult *units.MolarMass
    decigrams_per_moleResult, err = factory.FromDto(decigrams_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with DecigramPerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_moleResult.Convert(units.MolarMassDecigramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerMole = %v, want %v", converted, 100)
    }
    // Test DecagramPerMole conversion
    decagrams_per_moleDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassDecagramPerMole,
    }
    
    var decagrams_per_moleResult *units.MolarMass
    decagrams_per_moleResult, err = factory.FromDto(decagrams_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with DecagramPerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decagrams_per_moleResult.Convert(units.MolarMassDecagramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecagramPerMole = %v, want %v", converted, 100)
    }
    // Test HectogramPerMole conversion
    hectograms_per_moleDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassHectogramPerMole,
    }
    
    var hectograms_per_moleResult *units.MolarMass
    hectograms_per_moleResult, err = factory.FromDto(hectograms_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with HectogramPerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectograms_per_moleResult.Convert(units.MolarMassHectogramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectogramPerMole = %v, want %v", converted, 100)
    }
    // Test KilogramPerMole conversion
    kilograms_per_moleDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassKilogramPerMole,
    }
    
    var kilograms_per_moleResult *units.MolarMass
    kilograms_per_moleResult, err = factory.FromDto(kilograms_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_moleResult.Convert(units.MolarMassKilogramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerMole = %v, want %v", converted, 100)
    }
    // Test KilopoundPerMole conversion
    kilopounds_per_moleDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassKilopoundPerMole,
    }
    
    var kilopounds_per_moleResult *units.MolarMass
    kilopounds_per_moleResult, err = factory.FromDto(kilopounds_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundPerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_per_moleResult.Convert(units.MolarMassKilopoundPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundPerMole = %v, want %v", converted, 100)
    }
    // Test MegapoundPerMole conversion
    megapounds_per_moleDto := units.MolarMassDto{
        Value: 100,
        Unit:  units.MolarMassMegapoundPerMole,
    }
    
    var megapounds_per_moleResult *units.MolarMass
    megapounds_per_moleResult, err = factory.FromDto(megapounds_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with MegapoundPerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_per_moleResult.Convert(units.MolarMassMegapoundPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundPerMole = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MolarMassDto{
        Value: 0,
        Unit:  units.MolarMassKilogramPerMole,
    }
    
    var zeroResult *units.MolarMass
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMolarMassFactory_FromDtoJSON(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "KilogramPerMole"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "KilogramPerMole"}`)
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
    nanJSON, _ := json.Marshal(units.MolarMassDto{
        Value: nanValue,
        Unit:  units.MolarMassKilogramPerMole,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with GramPerMole unit
    grams_per_moleJSON := []byte(`{"value": 100, "unit": "GramPerMole"}`)
    grams_per_moleResult, err := factory.FromDtoJSON(grams_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_moleResult.Convert(units.MolarMassGramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerMole = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerKilomole unit
    kilograms_per_kilomoleJSON := []byte(`{"value": 100, "unit": "KilogramPerKilomole"}`)
    kilograms_per_kilomoleResult, err := factory.FromDtoJSON(kilograms_per_kilomoleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerKilomole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_kilomoleResult.Convert(units.MolarMassKilogramPerKilomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerKilomole = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerMole unit
    pounds_per_moleJSON := []byte(`{"value": 100, "unit": "PoundPerMole"}`)
    pounds_per_moleResult, err := factory.FromDtoJSON(pounds_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_moleResult.Convert(units.MolarMassPoundPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerMole = %v, want %v", converted, 100)
    }
    // Test JSON with NanogramPerMole unit
    nanograms_per_moleJSON := []byte(`{"value": 100, "unit": "NanogramPerMole"}`)
    nanograms_per_moleResult, err := factory.FromDtoJSON(nanograms_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanogramPerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_moleResult.Convert(units.MolarMassNanogramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerMole = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerMole unit
    micrograms_per_moleJSON := []byte(`{"value": 100, "unit": "MicrogramPerMole"}`)
    micrograms_per_moleResult, err := factory.FromDtoJSON(micrograms_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_moleResult.Convert(units.MolarMassMicrogramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerMole = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerMole unit
    milligrams_per_moleJSON := []byte(`{"value": 100, "unit": "MilligramPerMole"}`)
    milligrams_per_moleResult, err := factory.FromDtoJSON(milligrams_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_moleResult.Convert(units.MolarMassMilligramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerMole = %v, want %v", converted, 100)
    }
    // Test JSON with CentigramPerMole unit
    centigrams_per_moleJSON := []byte(`{"value": 100, "unit": "CentigramPerMole"}`)
    centigrams_per_moleResult, err := factory.FromDtoJSON(centigrams_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigramPerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_moleResult.Convert(units.MolarMassCentigramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerMole = %v, want %v", converted, 100)
    }
    // Test JSON with DecigramPerMole unit
    decigrams_per_moleJSON := []byte(`{"value": 100, "unit": "DecigramPerMole"}`)
    decigrams_per_moleResult, err := factory.FromDtoJSON(decigrams_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigramPerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_moleResult.Convert(units.MolarMassDecigramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerMole = %v, want %v", converted, 100)
    }
    // Test JSON with DecagramPerMole unit
    decagrams_per_moleJSON := []byte(`{"value": 100, "unit": "DecagramPerMole"}`)
    decagrams_per_moleResult, err := factory.FromDtoJSON(decagrams_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecagramPerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decagrams_per_moleResult.Convert(units.MolarMassDecagramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecagramPerMole = %v, want %v", converted, 100)
    }
    // Test JSON with HectogramPerMole unit
    hectograms_per_moleJSON := []byte(`{"value": 100, "unit": "HectogramPerMole"}`)
    hectograms_per_moleResult, err := factory.FromDtoJSON(hectograms_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HectogramPerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectograms_per_moleResult.Convert(units.MolarMassHectogramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectogramPerMole = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerMole unit
    kilograms_per_moleJSON := []byte(`{"value": 100, "unit": "KilogramPerMole"}`)
    kilograms_per_moleResult, err := factory.FromDtoJSON(kilograms_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_moleResult.Convert(units.MolarMassKilogramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerMole = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundPerMole unit
    kilopounds_per_moleJSON := []byte(`{"value": 100, "unit": "KilopoundPerMole"}`)
    kilopounds_per_moleResult, err := factory.FromDtoJSON(kilopounds_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundPerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_per_moleResult.Convert(units.MolarMassKilopoundPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundPerMole = %v, want %v", converted, 100)
    }
    // Test JSON with MegapoundPerMole unit
    megapounds_per_moleJSON := []byte(`{"value": 100, "unit": "MegapoundPerMole"}`)
    megapounds_per_moleResult, err := factory.FromDtoJSON(megapounds_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapoundPerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_per_moleResult.Convert(units.MolarMassMegapoundPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundPerMole = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "KilogramPerMole"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromGramsPerMole function
func TestMolarMassFactory_FromGramsPerMole(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerMole(100)
    if err != nil {
        t.Errorf("FromGramsPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarMassGramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerMole(math.NaN())
    if err == nil {
        t.Error("FromGramsPerMole() with NaN value should return error")
    }

    _, err = factory.FromGramsPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerMole() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerMole(0)
    if err != nil {
        t.Errorf("FromGramsPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarMassGramPerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerMole() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerKilomole function
func TestMolarMassFactory_FromKilogramsPerKilomole(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerKilomole(100)
    if err != nil {
        t.Errorf("FromKilogramsPerKilomole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarMassKilogramPerKilomole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerKilomole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerKilomole(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerKilomole() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerKilomole(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerKilomole() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerKilomole(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerKilomole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerKilomole(0)
    if err != nil {
        t.Errorf("FromKilogramsPerKilomole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarMassKilogramPerKilomole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerKilomole() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerMole function
func TestMolarMassFactory_FromPoundsPerMole(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerMole(100)
    if err != nil {
        t.Errorf("FromPoundsPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarMassPoundPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerMole(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerMole() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerMole() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerMole(0)
    if err != nil {
        t.Errorf("FromPoundsPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarMassPoundPerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerMole() with zero value = %v, want 0", converted)
    }
}
// Test FromNanogramsPerMole function
func TestMolarMassFactory_FromNanogramsPerMole(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanogramsPerMole(100)
    if err != nil {
        t.Errorf("FromNanogramsPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarMassNanogramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanogramsPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanogramsPerMole(math.NaN())
    if err == nil {
        t.Error("FromNanogramsPerMole() with NaN value should return error")
    }

    _, err = factory.FromNanogramsPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromNanogramsPerMole() with +Inf value should return error")
    }

    _, err = factory.FromNanogramsPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromNanogramsPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanogramsPerMole(0)
    if err != nil {
        t.Errorf("FromNanogramsPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarMassNanogramPerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanogramsPerMole() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerMole function
func TestMolarMassFactory_FromMicrogramsPerMole(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerMole(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarMassMicrogramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerMole(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerMole() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerMole() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerMole(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarMassMicrogramPerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerMole() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerMole function
func TestMolarMassFactory_FromMilligramsPerMole(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerMole(100)
    if err != nil {
        t.Errorf("FromMilligramsPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarMassMilligramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerMole(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerMole() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerMole() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerMole(0)
    if err != nil {
        t.Errorf("FromMilligramsPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarMassMilligramPerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerMole() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigramsPerMole function
func TestMolarMassFactory_FromCentigramsPerMole(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigramsPerMole(100)
    if err != nil {
        t.Errorf("FromCentigramsPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarMassCentigramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigramsPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigramsPerMole(math.NaN())
    if err == nil {
        t.Error("FromCentigramsPerMole() with NaN value should return error")
    }

    _, err = factory.FromCentigramsPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromCentigramsPerMole() with +Inf value should return error")
    }

    _, err = factory.FromCentigramsPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigramsPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigramsPerMole(0)
    if err != nil {
        t.Errorf("FromCentigramsPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarMassCentigramPerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigramsPerMole() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigramsPerMole function
func TestMolarMassFactory_FromDecigramsPerMole(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigramsPerMole(100)
    if err != nil {
        t.Errorf("FromDecigramsPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarMassDecigramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigramsPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigramsPerMole(math.NaN())
    if err == nil {
        t.Error("FromDecigramsPerMole() with NaN value should return error")
    }

    _, err = factory.FromDecigramsPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromDecigramsPerMole() with +Inf value should return error")
    }

    _, err = factory.FromDecigramsPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigramsPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigramsPerMole(0)
    if err != nil {
        t.Errorf("FromDecigramsPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarMassDecigramPerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigramsPerMole() with zero value = %v, want 0", converted)
    }
}
// Test FromDecagramsPerMole function
func TestMolarMassFactory_FromDecagramsPerMole(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecagramsPerMole(100)
    if err != nil {
        t.Errorf("FromDecagramsPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarMassDecagramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecagramsPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecagramsPerMole(math.NaN())
    if err == nil {
        t.Error("FromDecagramsPerMole() with NaN value should return error")
    }

    _, err = factory.FromDecagramsPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromDecagramsPerMole() with +Inf value should return error")
    }

    _, err = factory.FromDecagramsPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromDecagramsPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecagramsPerMole(0)
    if err != nil {
        t.Errorf("FromDecagramsPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarMassDecagramPerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecagramsPerMole() with zero value = %v, want 0", converted)
    }
}
// Test FromHectogramsPerMole function
func TestMolarMassFactory_FromHectogramsPerMole(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectogramsPerMole(100)
    if err != nil {
        t.Errorf("FromHectogramsPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarMassHectogramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectogramsPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectogramsPerMole(math.NaN())
    if err == nil {
        t.Error("FromHectogramsPerMole() with NaN value should return error")
    }

    _, err = factory.FromHectogramsPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromHectogramsPerMole() with +Inf value should return error")
    }

    _, err = factory.FromHectogramsPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromHectogramsPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectogramsPerMole(0)
    if err != nil {
        t.Errorf("FromHectogramsPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarMassHectogramPerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectogramsPerMole() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerMole function
func TestMolarMassFactory_FromKilogramsPerMole(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerMole(100)
    if err != nil {
        t.Errorf("FromKilogramsPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarMassKilogramPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerMole(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerMole() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerMole() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerMole(0)
    if err != nil {
        t.Errorf("FromKilogramsPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarMassKilogramPerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerMole() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsPerMole function
func TestMolarMassFactory_FromKilopoundsPerMole(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsPerMole(100)
    if err != nil {
        t.Errorf("FromKilopoundsPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarMassKilopoundPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsPerMole(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsPerMole() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsPerMole() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsPerMole(0)
    if err != nil {
        t.Errorf("FromKilopoundsPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarMassKilopoundPerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsPerMole() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapoundsPerMole function
func TestMolarMassFactory_FromMegapoundsPerMole(t *testing.T) {
    factory := units.MolarMassFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapoundsPerMole(100)
    if err != nil {
        t.Errorf("FromMegapoundsPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarMassMegapoundPerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapoundsPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapoundsPerMole(math.NaN())
    if err == nil {
        t.Error("FromMegapoundsPerMole() with NaN value should return error")
    }

    _, err = factory.FromMegapoundsPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromMegapoundsPerMole() with +Inf value should return error")
    }

    _, err = factory.FromMegapoundsPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapoundsPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapoundsPerMole(0)
    if err != nil {
        t.Errorf("FromMegapoundsPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarMassMegapoundPerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapoundsPerMole() with zero value = %v, want 0", converted)
    }
}

func TestMolarMassToString(t *testing.T) {
	factory := units.MolarMassFactory{}
	a, err := factory.CreateMolarMass(45, units.MolarMassKilogramPerMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MolarMassKilogramPerMole, 2)
	expected := "45.00 " + units.GetMolarMassAbbreviation(units.MolarMassKilogramPerMole)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MolarMassKilogramPerMole, -1)
	expected = "45 " + units.GetMolarMassAbbreviation(units.MolarMassKilogramPerMole)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMolarMass_EqualityAndComparison(t *testing.T) {
	factory := units.MolarMassFactory{}
	a1, _ := factory.CreateMolarMass(60, units.MolarMassKilogramPerMole)
	a2, _ := factory.CreateMolarMass(60, units.MolarMassKilogramPerMole)
	a3, _ := factory.CreateMolarMass(120, units.MolarMassKilogramPerMole)

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

func TestMolarMass_Arithmetic(t *testing.T) {
	factory := units.MolarMassFactory{}
	a1, _ := factory.CreateMolarMass(30, units.MolarMassKilogramPerMole)
	a2, _ := factory.CreateMolarMass(45, units.MolarMassKilogramPerMole)

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


func TestGetMolarMassAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.MolarMassUnits
        want string
    }{
        {
            name: "GramPerMole abbreviation",
            unit: units.MolarMassGramPerMole,
            want: "g/mol",
        },
        {
            name: "KilogramPerKilomole abbreviation",
            unit: units.MolarMassKilogramPerKilomole,
            want: "kg/kmol",
        },
        {
            name: "PoundPerMole abbreviation",
            unit: units.MolarMassPoundPerMole,
            want: "lb/mol",
        },
        {
            name: "NanogramPerMole abbreviation",
            unit: units.MolarMassNanogramPerMole,
            want: "ng/mol",
        },
        {
            name: "MicrogramPerMole abbreviation",
            unit: units.MolarMassMicrogramPerMole,
            want: "μg/mol",
        },
        {
            name: "MilligramPerMole abbreviation",
            unit: units.MolarMassMilligramPerMole,
            want: "mg/mol",
        },
        {
            name: "CentigramPerMole abbreviation",
            unit: units.MolarMassCentigramPerMole,
            want: "cg/mol",
        },
        {
            name: "DecigramPerMole abbreviation",
            unit: units.MolarMassDecigramPerMole,
            want: "dg/mol",
        },
        {
            name: "DecagramPerMole abbreviation",
            unit: units.MolarMassDecagramPerMole,
            want: "dag/mol",
        },
        {
            name: "HectogramPerMole abbreviation",
            unit: units.MolarMassHectogramPerMole,
            want: "hg/mol",
        },
        {
            name: "KilogramPerMole abbreviation",
            unit: units.MolarMassKilogramPerMole,
            want: "kg/mol",
        },
        {
            name: "KilopoundPerMole abbreviation",
            unit: units.MolarMassKilopoundPerMole,
            want: "klb/mol",
        },
        {
            name: "MegapoundPerMole abbreviation",
            unit: units.MolarMassMegapoundPerMole,
            want: "Mlb/mol",
        },
        {
            name: "invalid unit",
            unit: units.MolarMassUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetMolarMassAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetMolarMassAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestMolarMass_String(t *testing.T) {
    factory := units.MolarMassFactory{}
    
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
            unit, err := factory.CreateMolarMass(tt.value, units.MolarMassKilogramPerMole)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("MolarMass.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestMolarMass_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.MolarMassFactory{}

	_, err := uf.CreateMolarMass(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}