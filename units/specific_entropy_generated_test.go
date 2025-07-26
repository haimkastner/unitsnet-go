// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestSpecificEntropyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerKilogramKelvin"}`
	
	factory := units.SpecificEntropyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.SpecificEntropyJoulePerKilogramKelvin {
		t.Errorf("expected unit %v, got %v", units.SpecificEntropyJoulePerKilogramKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerKilogramKelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestSpecificEntropyDto_ToJSON(t *testing.T) {
	dto := units.SpecificEntropyDto{
		Value: 45,
		Unit:  units.SpecificEntropyJoulePerKilogramKelvin,
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
	if result["unit"].(string) != string(units.SpecificEntropyJoulePerKilogramKelvin) {
		t.Errorf("expected unit %s, got %v", units.SpecificEntropyJoulePerKilogramKelvin, result["unit"])
	}
}

func TestNewSpecificEntropy_InvalidValue(t *testing.T) {
	factory := units.SpecificEntropyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateSpecificEntropy(math.NaN(), units.SpecificEntropyJoulePerKilogramKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateSpecificEntropy(math.Inf(1), units.SpecificEntropyJoulePerKilogramKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestSpecificEntropyConversions(t *testing.T) {
	factory := units.SpecificEntropyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateSpecificEntropy(180, units.SpecificEntropyJoulePerKilogramKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerKilogramKelvin.
		// No expected conversion value provided for JoulesPerKilogramKelvin, verifying result is not NaN.
		result := a.JoulesPerKilogramKelvin()
		cacheResult := a.JoulesPerKilogramKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to JoulesPerKilogramKelvin returned NaN")
		}
	}
	{
		// Test conversion to JoulesPerKilogramDegreeCelsius.
		// No expected conversion value provided for JoulesPerKilogramDegreeCelsius, verifying result is not NaN.
		result := a.JoulesPerKilogramDegreeCelsius()
		cacheResult := a.JoulesPerKilogramDegreeCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to JoulesPerKilogramDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to CaloriesPerGramKelvin.
		// No expected conversion value provided for CaloriesPerGramKelvin, verifying result is not NaN.
		result := a.CaloriesPerGramKelvin()
		cacheResult := a.CaloriesPerGramKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CaloriesPerGramKelvin returned NaN")
		}
	}
	{
		// Test conversion to BtusPerPoundFahrenheit.
		// No expected conversion value provided for BtusPerPoundFahrenheit, verifying result is not NaN.
		result := a.BtusPerPoundFahrenheit()
		cacheResult := a.BtusPerPoundFahrenheit()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to BtusPerPoundFahrenheit returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerKilogramKelvin.
		// No expected conversion value provided for KilojoulesPerKilogramKelvin, verifying result is not NaN.
		result := a.KilojoulesPerKilogramKelvin()
		cacheResult := a.KilojoulesPerKilogramKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilojoulesPerKilogramKelvin returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerKilogramKelvin.
		// No expected conversion value provided for MegajoulesPerKilogramKelvin, verifying result is not NaN.
		result := a.MegajoulesPerKilogramKelvin()
		cacheResult := a.MegajoulesPerKilogramKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegajoulesPerKilogramKelvin returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerKilogramDegreeCelsius.
		// No expected conversion value provided for KilojoulesPerKilogramDegreeCelsius, verifying result is not NaN.
		result := a.KilojoulesPerKilogramDegreeCelsius()
		cacheResult := a.KilojoulesPerKilogramDegreeCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilojoulesPerKilogramDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerKilogramDegreeCelsius.
		// No expected conversion value provided for MegajoulesPerKilogramDegreeCelsius, verifying result is not NaN.
		result := a.MegajoulesPerKilogramDegreeCelsius()
		cacheResult := a.MegajoulesPerKilogramDegreeCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegajoulesPerKilogramDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerGramKelvin.
		// No expected conversion value provided for KilocaloriesPerGramKelvin, verifying result is not NaN.
		result := a.KilocaloriesPerGramKelvin()
		cacheResult := a.KilocaloriesPerGramKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilocaloriesPerGramKelvin returned NaN")
		}
	}
}

func TestSpecificEntropy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.SpecificEntropyFactory{}
	a, err := factory.CreateSpecificEntropy(90, units.SpecificEntropyJoulePerKilogramKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.SpecificEntropyJoulePerKilogramKelvin {
		t.Errorf("expected default unit JoulePerKilogramKelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.SpecificEntropyJoulePerKilogramKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.SpecificEntropyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.SpecificEntropyJoulePerKilogramKelvin {
		t.Errorf("expected unit JoulePerKilogramKelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestSpecificEntropyFactory_FromDto(t *testing.T) {
    factory := units.SpecificEntropyFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.SpecificEntropyDto{
        Value: 100,
        Unit:  units.SpecificEntropyJoulePerKilogramKelvin,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.SpecificEntropyDto{
        Value: math.NaN(),
        Unit:  units.SpecificEntropyJoulePerKilogramKelvin,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test JoulePerKilogramKelvin conversion
    joules_per_kilogram_kelvinDto := units.SpecificEntropyDto{
        Value: 100,
        Unit:  units.SpecificEntropyJoulePerKilogramKelvin,
    }
    
    var joules_per_kilogram_kelvinResult *units.SpecificEntropy
    joules_per_kilogram_kelvinResult, err = factory.FromDto(joules_per_kilogram_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerKilogramKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_kilogram_kelvinResult.Convert(units.SpecificEntropyJoulePerKilogramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerKilogramKelvin = %v, want %v", converted, 100)
    }
    // Test JoulePerKilogramDegreeCelsius conversion
    joules_per_kilogram_degree_celsiusDto := units.SpecificEntropyDto{
        Value: 100,
        Unit:  units.SpecificEntropyJoulePerKilogramDegreeCelsius,
    }
    
    var joules_per_kilogram_degree_celsiusResult *units.SpecificEntropy
    joules_per_kilogram_degree_celsiusResult, err = factory.FromDto(joules_per_kilogram_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerKilogramDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_kilogram_degree_celsiusResult.Convert(units.SpecificEntropyJoulePerKilogramDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerKilogramDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test CaloriePerGramKelvin conversion
    calories_per_gram_kelvinDto := units.SpecificEntropyDto{
        Value: 100,
        Unit:  units.SpecificEntropyCaloriePerGramKelvin,
    }
    
    var calories_per_gram_kelvinResult *units.SpecificEntropy
    calories_per_gram_kelvinResult, err = factory.FromDto(calories_per_gram_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with CaloriePerGramKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = calories_per_gram_kelvinResult.Convert(units.SpecificEntropyCaloriePerGramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CaloriePerGramKelvin = %v, want %v", converted, 100)
    }
    // Test BtuPerPoundFahrenheit conversion
    btus_per_pound_fahrenheitDto := units.SpecificEntropyDto{
        Value: 100,
        Unit:  units.SpecificEntropyBtuPerPoundFahrenheit,
    }
    
    var btus_per_pound_fahrenheitResult *units.SpecificEntropy
    btus_per_pound_fahrenheitResult, err = factory.FromDto(btus_per_pound_fahrenheitDto)
    if err != nil {
        t.Errorf("FromDto() with BtuPerPoundFahrenheit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_pound_fahrenheitResult.Convert(units.SpecificEntropyBtuPerPoundFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerPoundFahrenheit = %v, want %v", converted, 100)
    }
    // Test KilojoulePerKilogramKelvin conversion
    kilojoules_per_kilogram_kelvinDto := units.SpecificEntropyDto{
        Value: 100,
        Unit:  units.SpecificEntropyKilojoulePerKilogramKelvin,
    }
    
    var kilojoules_per_kilogram_kelvinResult *units.SpecificEntropy
    kilojoules_per_kilogram_kelvinResult, err = factory.FromDto(kilojoules_per_kilogram_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with KilojoulePerKilogramKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_kilogram_kelvinResult.Convert(units.SpecificEntropyKilojoulePerKilogramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerKilogramKelvin = %v, want %v", converted, 100)
    }
    // Test MegajoulePerKilogramKelvin conversion
    megajoules_per_kilogram_kelvinDto := units.SpecificEntropyDto{
        Value: 100,
        Unit:  units.SpecificEntropyMegajoulePerKilogramKelvin,
    }
    
    var megajoules_per_kilogram_kelvinResult *units.SpecificEntropy
    megajoules_per_kilogram_kelvinResult, err = factory.FromDto(megajoules_per_kilogram_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with MegajoulePerKilogramKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_kilogram_kelvinResult.Convert(units.SpecificEntropyMegajoulePerKilogramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerKilogramKelvin = %v, want %v", converted, 100)
    }
    // Test KilojoulePerKilogramDegreeCelsius conversion
    kilojoules_per_kilogram_degree_celsiusDto := units.SpecificEntropyDto{
        Value: 100,
        Unit:  units.SpecificEntropyKilojoulePerKilogramDegreeCelsius,
    }
    
    var kilojoules_per_kilogram_degree_celsiusResult *units.SpecificEntropy
    kilojoules_per_kilogram_degree_celsiusResult, err = factory.FromDto(kilojoules_per_kilogram_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with KilojoulePerKilogramDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_kilogram_degree_celsiusResult.Convert(units.SpecificEntropyKilojoulePerKilogramDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerKilogramDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test MegajoulePerKilogramDegreeCelsius conversion
    megajoules_per_kilogram_degree_celsiusDto := units.SpecificEntropyDto{
        Value: 100,
        Unit:  units.SpecificEntropyMegajoulePerKilogramDegreeCelsius,
    }
    
    var megajoules_per_kilogram_degree_celsiusResult *units.SpecificEntropy
    megajoules_per_kilogram_degree_celsiusResult, err = factory.FromDto(megajoules_per_kilogram_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with MegajoulePerKilogramDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_kilogram_degree_celsiusResult.Convert(units.SpecificEntropyMegajoulePerKilogramDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerKilogramDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test KilocaloriePerGramKelvin conversion
    kilocalories_per_gram_kelvinDto := units.SpecificEntropyDto{
        Value: 100,
        Unit:  units.SpecificEntropyKilocaloriePerGramKelvin,
    }
    
    var kilocalories_per_gram_kelvinResult *units.SpecificEntropy
    kilocalories_per_gram_kelvinResult, err = factory.FromDto(kilocalories_per_gram_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with KilocaloriePerGramKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_gram_kelvinResult.Convert(units.SpecificEntropyKilocaloriePerGramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerGramKelvin = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.SpecificEntropyDto{
        Value: 0,
        Unit:  units.SpecificEntropyJoulePerKilogramKelvin,
    }
    
    var zeroResult *units.SpecificEntropy
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestSpecificEntropyFactory_FromDtoJSON(t *testing.T) {
    factory := units.SpecificEntropyFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "JoulePerKilogramKelvin"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "JoulePerKilogramKelvin"}`)
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
    nanJSON, _ := json.Marshal(units.SpecificEntropyDto{
        Value: nanValue,
        Unit:  units.SpecificEntropyJoulePerKilogramKelvin,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with JoulePerKilogramKelvin unit
    joules_per_kilogram_kelvinJSON := []byte(`{"value": 100, "unit": "JoulePerKilogramKelvin"}`)
    joules_per_kilogram_kelvinResult, err := factory.FromDtoJSON(joules_per_kilogram_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerKilogramKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_kilogram_kelvinResult.Convert(units.SpecificEntropyJoulePerKilogramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerKilogramKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with JoulePerKilogramDegreeCelsius unit
    joules_per_kilogram_degree_celsiusJSON := []byte(`{"value": 100, "unit": "JoulePerKilogramDegreeCelsius"}`)
    joules_per_kilogram_degree_celsiusResult, err := factory.FromDtoJSON(joules_per_kilogram_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerKilogramDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_kilogram_degree_celsiusResult.Convert(units.SpecificEntropyJoulePerKilogramDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerKilogramDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with CaloriePerGramKelvin unit
    calories_per_gram_kelvinJSON := []byte(`{"value": 100, "unit": "CaloriePerGramKelvin"}`)
    calories_per_gram_kelvinResult, err := factory.FromDtoJSON(calories_per_gram_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CaloriePerGramKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = calories_per_gram_kelvinResult.Convert(units.SpecificEntropyCaloriePerGramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CaloriePerGramKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with BtuPerPoundFahrenheit unit
    btus_per_pound_fahrenheitJSON := []byte(`{"value": 100, "unit": "BtuPerPoundFahrenheit"}`)
    btus_per_pound_fahrenheitResult, err := factory.FromDtoJSON(btus_per_pound_fahrenheitJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BtuPerPoundFahrenheit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_pound_fahrenheitResult.Convert(units.SpecificEntropyBtuPerPoundFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerPoundFahrenheit = %v, want %v", converted, 100)
    }
    // Test JSON with KilojoulePerKilogramKelvin unit
    kilojoules_per_kilogram_kelvinJSON := []byte(`{"value": 100, "unit": "KilojoulePerKilogramKelvin"}`)
    kilojoules_per_kilogram_kelvinResult, err := factory.FromDtoJSON(kilojoules_per_kilogram_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilojoulePerKilogramKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_kilogram_kelvinResult.Convert(units.SpecificEntropyKilojoulePerKilogramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerKilogramKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with MegajoulePerKilogramKelvin unit
    megajoules_per_kilogram_kelvinJSON := []byte(`{"value": 100, "unit": "MegajoulePerKilogramKelvin"}`)
    megajoules_per_kilogram_kelvinResult, err := factory.FromDtoJSON(megajoules_per_kilogram_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegajoulePerKilogramKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_kilogram_kelvinResult.Convert(units.SpecificEntropyMegajoulePerKilogramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerKilogramKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with KilojoulePerKilogramDegreeCelsius unit
    kilojoules_per_kilogram_degree_celsiusJSON := []byte(`{"value": 100, "unit": "KilojoulePerKilogramDegreeCelsius"}`)
    kilojoules_per_kilogram_degree_celsiusResult, err := factory.FromDtoJSON(kilojoules_per_kilogram_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilojoulePerKilogramDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_kilogram_degree_celsiusResult.Convert(units.SpecificEntropyKilojoulePerKilogramDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerKilogramDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with MegajoulePerKilogramDegreeCelsius unit
    megajoules_per_kilogram_degree_celsiusJSON := []byte(`{"value": 100, "unit": "MegajoulePerKilogramDegreeCelsius"}`)
    megajoules_per_kilogram_degree_celsiusResult, err := factory.FromDtoJSON(megajoules_per_kilogram_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegajoulePerKilogramDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_kilogram_degree_celsiusResult.Convert(units.SpecificEntropyMegajoulePerKilogramDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerKilogramDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with KilocaloriePerGramKelvin unit
    kilocalories_per_gram_kelvinJSON := []byte(`{"value": 100, "unit": "KilocaloriePerGramKelvin"}`)
    kilocalories_per_gram_kelvinResult, err := factory.FromDtoJSON(kilocalories_per_gram_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilocaloriePerGramKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_gram_kelvinResult.Convert(units.SpecificEntropyKilocaloriePerGramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerGramKelvin = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "JoulePerKilogramKelvin"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromJoulesPerKilogramKelvin function
func TestSpecificEntropyFactory_FromJoulesPerKilogramKelvin(t *testing.T) {
    factory := units.SpecificEntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerKilogramKelvin(100)
    if err != nil {
        t.Errorf("FromJoulesPerKilogramKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEntropyJoulePerKilogramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerKilogramKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerKilogramKelvin(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerKilogramKelvin() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerKilogramKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerKilogramKelvin() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerKilogramKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerKilogramKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerKilogramKelvin(0)
    if err != nil {
        t.Errorf("FromJoulesPerKilogramKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEntropyJoulePerKilogramKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerKilogramKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromJoulesPerKilogramDegreeCelsius function
func TestSpecificEntropyFactory_FromJoulesPerKilogramDegreeCelsius(t *testing.T) {
    factory := units.SpecificEntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerKilogramDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromJoulesPerKilogramDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEntropyJoulePerKilogramDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerKilogramDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerKilogramDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerKilogramDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerKilogramDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerKilogramDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerKilogramDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerKilogramDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerKilogramDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromJoulesPerKilogramDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEntropyJoulePerKilogramDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerKilogramDegreeCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromCaloriesPerGramKelvin function
func TestSpecificEntropyFactory_FromCaloriesPerGramKelvin(t *testing.T) {
    factory := units.SpecificEntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCaloriesPerGramKelvin(100)
    if err != nil {
        t.Errorf("FromCaloriesPerGramKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEntropyCaloriePerGramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCaloriesPerGramKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCaloriesPerGramKelvin(math.NaN())
    if err == nil {
        t.Error("FromCaloriesPerGramKelvin() with NaN value should return error")
    }

    _, err = factory.FromCaloriesPerGramKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromCaloriesPerGramKelvin() with +Inf value should return error")
    }

    _, err = factory.FromCaloriesPerGramKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromCaloriesPerGramKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCaloriesPerGramKelvin(0)
    if err != nil {
        t.Errorf("FromCaloriesPerGramKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEntropyCaloriePerGramKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCaloriesPerGramKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromBtusPerPoundFahrenheit function
func TestSpecificEntropyFactory_FromBtusPerPoundFahrenheit(t *testing.T) {
    factory := units.SpecificEntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBtusPerPoundFahrenheit(100)
    if err != nil {
        t.Errorf("FromBtusPerPoundFahrenheit() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEntropyBtuPerPoundFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBtusPerPoundFahrenheit() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBtusPerPoundFahrenheit(math.NaN())
    if err == nil {
        t.Error("FromBtusPerPoundFahrenheit() with NaN value should return error")
    }

    _, err = factory.FromBtusPerPoundFahrenheit(math.Inf(1))
    if err == nil {
        t.Error("FromBtusPerPoundFahrenheit() with +Inf value should return error")
    }

    _, err = factory.FromBtusPerPoundFahrenheit(math.Inf(-1))
    if err == nil {
        t.Error("FromBtusPerPoundFahrenheit() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBtusPerPoundFahrenheit(0)
    if err != nil {
        t.Errorf("FromBtusPerPoundFahrenheit() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEntropyBtuPerPoundFahrenheit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBtusPerPoundFahrenheit() with zero value = %v, want 0", converted)
    }
}
// Test FromKilojoulesPerKilogramKelvin function
func TestSpecificEntropyFactory_FromKilojoulesPerKilogramKelvin(t *testing.T) {
    factory := units.SpecificEntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilojoulesPerKilogramKelvin(100)
    if err != nil {
        t.Errorf("FromKilojoulesPerKilogramKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEntropyKilojoulePerKilogramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilojoulesPerKilogramKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilojoulesPerKilogramKelvin(math.NaN())
    if err == nil {
        t.Error("FromKilojoulesPerKilogramKelvin() with NaN value should return error")
    }

    _, err = factory.FromKilojoulesPerKilogramKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromKilojoulesPerKilogramKelvin() with +Inf value should return error")
    }

    _, err = factory.FromKilojoulesPerKilogramKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromKilojoulesPerKilogramKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilojoulesPerKilogramKelvin(0)
    if err != nil {
        t.Errorf("FromKilojoulesPerKilogramKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEntropyKilojoulePerKilogramKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilojoulesPerKilogramKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromMegajoulesPerKilogramKelvin function
func TestSpecificEntropyFactory_FromMegajoulesPerKilogramKelvin(t *testing.T) {
    factory := units.SpecificEntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegajoulesPerKilogramKelvin(100)
    if err != nil {
        t.Errorf("FromMegajoulesPerKilogramKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEntropyMegajoulePerKilogramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegajoulesPerKilogramKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegajoulesPerKilogramKelvin(math.NaN())
    if err == nil {
        t.Error("FromMegajoulesPerKilogramKelvin() with NaN value should return error")
    }

    _, err = factory.FromMegajoulesPerKilogramKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromMegajoulesPerKilogramKelvin() with +Inf value should return error")
    }

    _, err = factory.FromMegajoulesPerKilogramKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromMegajoulesPerKilogramKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegajoulesPerKilogramKelvin(0)
    if err != nil {
        t.Errorf("FromMegajoulesPerKilogramKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEntropyMegajoulePerKilogramKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegajoulesPerKilogramKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromKilojoulesPerKilogramDegreeCelsius function
func TestSpecificEntropyFactory_FromKilojoulesPerKilogramDegreeCelsius(t *testing.T) {
    factory := units.SpecificEntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilojoulesPerKilogramDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromKilojoulesPerKilogramDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEntropyKilojoulePerKilogramDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilojoulesPerKilogramDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilojoulesPerKilogramDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromKilojoulesPerKilogramDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromKilojoulesPerKilogramDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromKilojoulesPerKilogramDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromKilojoulesPerKilogramDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromKilojoulesPerKilogramDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilojoulesPerKilogramDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromKilojoulesPerKilogramDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEntropyKilojoulePerKilogramDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilojoulesPerKilogramDegreeCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromMegajoulesPerKilogramDegreeCelsius function
func TestSpecificEntropyFactory_FromMegajoulesPerKilogramDegreeCelsius(t *testing.T) {
    factory := units.SpecificEntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegajoulesPerKilogramDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromMegajoulesPerKilogramDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEntropyMegajoulePerKilogramDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegajoulesPerKilogramDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegajoulesPerKilogramDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromMegajoulesPerKilogramDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromMegajoulesPerKilogramDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromMegajoulesPerKilogramDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromMegajoulesPerKilogramDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromMegajoulesPerKilogramDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegajoulesPerKilogramDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromMegajoulesPerKilogramDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEntropyMegajoulePerKilogramDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegajoulesPerKilogramDegreeCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromKilocaloriesPerGramKelvin function
func TestSpecificEntropyFactory_FromKilocaloriesPerGramKelvin(t *testing.T) {
    factory := units.SpecificEntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilocaloriesPerGramKelvin(100)
    if err != nil {
        t.Errorf("FromKilocaloriesPerGramKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEntropyKilocaloriePerGramKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilocaloriesPerGramKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilocaloriesPerGramKelvin(math.NaN())
    if err == nil {
        t.Error("FromKilocaloriesPerGramKelvin() with NaN value should return error")
    }

    _, err = factory.FromKilocaloriesPerGramKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromKilocaloriesPerGramKelvin() with +Inf value should return error")
    }

    _, err = factory.FromKilocaloriesPerGramKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromKilocaloriesPerGramKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilocaloriesPerGramKelvin(0)
    if err != nil {
        t.Errorf("FromKilocaloriesPerGramKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEntropyKilocaloriePerGramKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilocaloriesPerGramKelvin() with zero value = %v, want 0", converted)
    }
}

func TestSpecificEntropyToString(t *testing.T) {
	factory := units.SpecificEntropyFactory{}
	a, err := factory.CreateSpecificEntropy(45, units.SpecificEntropyJoulePerKilogramKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.SpecificEntropyJoulePerKilogramKelvin, 2)
	expected := "45.00 " + units.GetSpecificEntropyAbbreviation(units.SpecificEntropyJoulePerKilogramKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.SpecificEntropyJoulePerKilogramKelvin, -1)
	expected = "45 " + units.GetSpecificEntropyAbbreviation(units.SpecificEntropyJoulePerKilogramKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestSpecificEntropy_EqualityAndComparison(t *testing.T) {
	factory := units.SpecificEntropyFactory{}
	a1, _ := factory.CreateSpecificEntropy(60, units.SpecificEntropyJoulePerKilogramKelvin)
	a2, _ := factory.CreateSpecificEntropy(60, units.SpecificEntropyJoulePerKilogramKelvin)
	a3, _ := factory.CreateSpecificEntropy(120, units.SpecificEntropyJoulePerKilogramKelvin)

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

func TestSpecificEntropy_Arithmetic(t *testing.T) {
	factory := units.SpecificEntropyFactory{}
	a1, _ := factory.CreateSpecificEntropy(30, units.SpecificEntropyJoulePerKilogramKelvin)
	a2, _ := factory.CreateSpecificEntropy(45, units.SpecificEntropyJoulePerKilogramKelvin)

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


func TestGetSpecificEntropyAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.SpecificEntropyUnits
        want string
    }{
        {
            name: "JoulePerKilogramKelvin abbreviation",
            unit: units.SpecificEntropyJoulePerKilogramKelvin,
            want: "J/kg·K",
        },
        {
            name: "JoulePerKilogramDegreeCelsius abbreviation",
            unit: units.SpecificEntropyJoulePerKilogramDegreeCelsius,
            want: "J/kg·°C",
        },
        {
            name: "CaloriePerGramKelvin abbreviation",
            unit: units.SpecificEntropyCaloriePerGramKelvin,
            want: "cal/g·K",
        },
        {
            name: "BtuPerPoundFahrenheit abbreviation",
            unit: units.SpecificEntropyBtuPerPoundFahrenheit,
            want: "BTU/(lb·°F)",
        },
        {
            name: "KilojoulePerKilogramKelvin abbreviation",
            unit: units.SpecificEntropyKilojoulePerKilogramKelvin,
            want: "kJ/kg·K",
        },
        {
            name: "MegajoulePerKilogramKelvin abbreviation",
            unit: units.SpecificEntropyMegajoulePerKilogramKelvin,
            want: "MJ/kg·K",
        },
        {
            name: "KilojoulePerKilogramDegreeCelsius abbreviation",
            unit: units.SpecificEntropyKilojoulePerKilogramDegreeCelsius,
            want: "kJ/kg·°C",
        },
        {
            name: "MegajoulePerKilogramDegreeCelsius abbreviation",
            unit: units.SpecificEntropyMegajoulePerKilogramDegreeCelsius,
            want: "MJ/kg·°C",
        },
        {
            name: "KilocaloriePerGramKelvin abbreviation",
            unit: units.SpecificEntropyKilocaloriePerGramKelvin,
            want: "kcal/g·K",
        },
        {
            name: "invalid unit",
            unit: units.SpecificEntropyUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetSpecificEntropyAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetSpecificEntropyAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestSpecificEntropy_String(t *testing.T) {
    factory := units.SpecificEntropyFactory{}
    
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
            unit, err := factory.CreateSpecificEntropy(tt.value, units.SpecificEntropyJoulePerKilogramKelvin)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("SpecificEntropy.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestSpecificEntropy_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.SpecificEntropyFactory{}

	_, err := uf.CreateSpecificEntropy(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}