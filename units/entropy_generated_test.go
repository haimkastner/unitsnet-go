// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestEntropyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerKelvin"}`
	
	factory := units.EntropyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.EntropyJoulePerKelvin {
		t.Errorf("expected unit %v, got %v", units.EntropyJoulePerKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerKelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestEntropyDto_ToJSON(t *testing.T) {
	dto := units.EntropyDto{
		Value: 45,
		Unit:  units.EntropyJoulePerKelvin,
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
	if result["unit"].(string) != string(units.EntropyJoulePerKelvin) {
		t.Errorf("expected unit %s, got %v", units.EntropyJoulePerKelvin, result["unit"])
	}
}

func TestNewEntropy_InvalidValue(t *testing.T) {
	factory := units.EntropyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateEntropy(math.NaN(), units.EntropyJoulePerKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateEntropy(math.Inf(1), units.EntropyJoulePerKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestEntropyConversions(t *testing.T) {
	factory := units.EntropyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateEntropy(180, units.EntropyJoulePerKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerKelvin.
		// No expected conversion value provided for JoulesPerKelvin, verifying result is not NaN.
		result := a.JoulesPerKelvin()
		cacheResult := a.JoulesPerKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to JoulesPerKelvin returned NaN")
		}
	}
	{
		// Test conversion to CaloriesPerKelvin.
		// No expected conversion value provided for CaloriesPerKelvin, verifying result is not NaN.
		result := a.CaloriesPerKelvin()
		cacheResult := a.CaloriesPerKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CaloriesPerKelvin returned NaN")
		}
	}
	{
		// Test conversion to JoulesPerDegreeCelsius.
		// No expected conversion value provided for JoulesPerDegreeCelsius, verifying result is not NaN.
		result := a.JoulesPerDegreeCelsius()
		cacheResult := a.JoulesPerDegreeCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to JoulesPerDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerKelvin.
		// No expected conversion value provided for KilojoulesPerKelvin, verifying result is not NaN.
		result := a.KilojoulesPerKelvin()
		cacheResult := a.KilojoulesPerKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilojoulesPerKelvin returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerKelvin.
		// No expected conversion value provided for MegajoulesPerKelvin, verifying result is not NaN.
		result := a.MegajoulesPerKelvin()
		cacheResult := a.MegajoulesPerKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegajoulesPerKelvin returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerKelvin.
		// No expected conversion value provided for KilocaloriesPerKelvin, verifying result is not NaN.
		result := a.KilocaloriesPerKelvin()
		cacheResult := a.KilocaloriesPerKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilocaloriesPerKelvin returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerDegreeCelsius.
		// No expected conversion value provided for KilojoulesPerDegreeCelsius, verifying result is not NaN.
		result := a.KilojoulesPerDegreeCelsius()
		cacheResult := a.KilojoulesPerDegreeCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilojoulesPerDegreeCelsius returned NaN")
		}
	}
}

func TestEntropy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.EntropyFactory{}
	a, err := factory.CreateEntropy(90, units.EntropyJoulePerKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.EntropyJoulePerKelvin {
		t.Errorf("expected default unit JoulePerKelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.EntropyJoulePerKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.EntropyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.EntropyJoulePerKelvin {
		t.Errorf("expected unit JoulePerKelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestEntropyFactory_FromDto(t *testing.T) {
    factory := units.EntropyFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.EntropyDto{
        Value: 100,
        Unit:  units.EntropyJoulePerKelvin,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.EntropyDto{
        Value: math.NaN(),
        Unit:  units.EntropyJoulePerKelvin,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test JoulePerKelvin conversion
    joules_per_kelvinDto := units.EntropyDto{
        Value: 100,
        Unit:  units.EntropyJoulePerKelvin,
    }
    
    var joules_per_kelvinResult *units.Entropy
    joules_per_kelvinResult, err = factory.FromDto(joules_per_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_kelvinResult.Convert(units.EntropyJoulePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerKelvin = %v, want %v", converted, 100)
    }
    // Test CaloriePerKelvin conversion
    calories_per_kelvinDto := units.EntropyDto{
        Value: 100,
        Unit:  units.EntropyCaloriePerKelvin,
    }
    
    var calories_per_kelvinResult *units.Entropy
    calories_per_kelvinResult, err = factory.FromDto(calories_per_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with CaloriePerKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = calories_per_kelvinResult.Convert(units.EntropyCaloriePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CaloriePerKelvin = %v, want %v", converted, 100)
    }
    // Test JoulePerDegreeCelsius conversion
    joules_per_degree_celsiusDto := units.EntropyDto{
        Value: 100,
        Unit:  units.EntropyJoulePerDegreeCelsius,
    }
    
    var joules_per_degree_celsiusResult *units.Entropy
    joules_per_degree_celsiusResult, err = factory.FromDto(joules_per_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_degree_celsiusResult.Convert(units.EntropyJoulePerDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test KilojoulePerKelvin conversion
    kilojoules_per_kelvinDto := units.EntropyDto{
        Value: 100,
        Unit:  units.EntropyKilojoulePerKelvin,
    }
    
    var kilojoules_per_kelvinResult *units.Entropy
    kilojoules_per_kelvinResult, err = factory.FromDto(kilojoules_per_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with KilojoulePerKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_kelvinResult.Convert(units.EntropyKilojoulePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerKelvin = %v, want %v", converted, 100)
    }
    // Test MegajoulePerKelvin conversion
    megajoules_per_kelvinDto := units.EntropyDto{
        Value: 100,
        Unit:  units.EntropyMegajoulePerKelvin,
    }
    
    var megajoules_per_kelvinResult *units.Entropy
    megajoules_per_kelvinResult, err = factory.FromDto(megajoules_per_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with MegajoulePerKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_kelvinResult.Convert(units.EntropyMegajoulePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerKelvin = %v, want %v", converted, 100)
    }
    // Test KilocaloriePerKelvin conversion
    kilocalories_per_kelvinDto := units.EntropyDto{
        Value: 100,
        Unit:  units.EntropyKilocaloriePerKelvin,
    }
    
    var kilocalories_per_kelvinResult *units.Entropy
    kilocalories_per_kelvinResult, err = factory.FromDto(kilocalories_per_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with KilocaloriePerKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_kelvinResult.Convert(units.EntropyKilocaloriePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerKelvin = %v, want %v", converted, 100)
    }
    // Test KilojoulePerDegreeCelsius conversion
    kilojoules_per_degree_celsiusDto := units.EntropyDto{
        Value: 100,
        Unit:  units.EntropyKilojoulePerDegreeCelsius,
    }
    
    var kilojoules_per_degree_celsiusResult *units.Entropy
    kilojoules_per_degree_celsiusResult, err = factory.FromDto(kilojoules_per_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with KilojoulePerDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_degree_celsiusResult.Convert(units.EntropyKilojoulePerDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerDegreeCelsius = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.EntropyDto{
        Value: 0,
        Unit:  units.EntropyJoulePerKelvin,
    }
    
    var zeroResult *units.Entropy
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestEntropyFactory_FromDtoJSON(t *testing.T) {
    factory := units.EntropyFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "JoulePerKelvin"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "JoulePerKelvin"}`)
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
    nanJSON, _ := json.Marshal(units.EntropyDto{
        Value: nanValue,
        Unit:  units.EntropyJoulePerKelvin,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with JoulePerKelvin unit
    joules_per_kelvinJSON := []byte(`{"value": 100, "unit": "JoulePerKelvin"}`)
    joules_per_kelvinResult, err := factory.FromDtoJSON(joules_per_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_kelvinResult.Convert(units.EntropyJoulePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with CaloriePerKelvin unit
    calories_per_kelvinJSON := []byte(`{"value": 100, "unit": "CaloriePerKelvin"}`)
    calories_per_kelvinResult, err := factory.FromDtoJSON(calories_per_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CaloriePerKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = calories_per_kelvinResult.Convert(units.EntropyCaloriePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CaloriePerKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with JoulePerDegreeCelsius unit
    joules_per_degree_celsiusJSON := []byte(`{"value": 100, "unit": "JoulePerDegreeCelsius"}`)
    joules_per_degree_celsiusResult, err := factory.FromDtoJSON(joules_per_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_degree_celsiusResult.Convert(units.EntropyJoulePerDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with KilojoulePerKelvin unit
    kilojoules_per_kelvinJSON := []byte(`{"value": 100, "unit": "KilojoulePerKelvin"}`)
    kilojoules_per_kelvinResult, err := factory.FromDtoJSON(kilojoules_per_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilojoulePerKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_kelvinResult.Convert(units.EntropyKilojoulePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with MegajoulePerKelvin unit
    megajoules_per_kelvinJSON := []byte(`{"value": 100, "unit": "MegajoulePerKelvin"}`)
    megajoules_per_kelvinResult, err := factory.FromDtoJSON(megajoules_per_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegajoulePerKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_kelvinResult.Convert(units.EntropyMegajoulePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with KilocaloriePerKelvin unit
    kilocalories_per_kelvinJSON := []byte(`{"value": 100, "unit": "KilocaloriePerKelvin"}`)
    kilocalories_per_kelvinResult, err := factory.FromDtoJSON(kilocalories_per_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilocaloriePerKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_kelvinResult.Convert(units.EntropyKilocaloriePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with KilojoulePerDegreeCelsius unit
    kilojoules_per_degree_celsiusJSON := []byte(`{"value": 100, "unit": "KilojoulePerDegreeCelsius"}`)
    kilojoules_per_degree_celsiusResult, err := factory.FromDtoJSON(kilojoules_per_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilojoulePerDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_degree_celsiusResult.Convert(units.EntropyKilojoulePerDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerDegreeCelsius = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "JoulePerKelvin"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromJoulesPerKelvin function
func TestEntropyFactory_FromJoulesPerKelvin(t *testing.T) {
    factory := units.EntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerKelvin(100)
    if err != nil {
        t.Errorf("FromJoulesPerKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EntropyJoulePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerKelvin(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerKelvin() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerKelvin() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerKelvin(0)
    if err != nil {
        t.Errorf("FromJoulesPerKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EntropyJoulePerKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromCaloriesPerKelvin function
func TestEntropyFactory_FromCaloriesPerKelvin(t *testing.T) {
    factory := units.EntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCaloriesPerKelvin(100)
    if err != nil {
        t.Errorf("FromCaloriesPerKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EntropyCaloriePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCaloriesPerKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCaloriesPerKelvin(math.NaN())
    if err == nil {
        t.Error("FromCaloriesPerKelvin() with NaN value should return error")
    }

    _, err = factory.FromCaloriesPerKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromCaloriesPerKelvin() with +Inf value should return error")
    }

    _, err = factory.FromCaloriesPerKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromCaloriesPerKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCaloriesPerKelvin(0)
    if err != nil {
        t.Errorf("FromCaloriesPerKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EntropyCaloriePerKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCaloriesPerKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromJoulesPerDegreeCelsius function
func TestEntropyFactory_FromJoulesPerDegreeCelsius(t *testing.T) {
    factory := units.EntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromJoulesPerDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EntropyJoulePerDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromJoulesPerDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EntropyJoulePerDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerDegreeCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromKilojoulesPerKelvin function
func TestEntropyFactory_FromKilojoulesPerKelvin(t *testing.T) {
    factory := units.EntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilojoulesPerKelvin(100)
    if err != nil {
        t.Errorf("FromKilojoulesPerKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EntropyKilojoulePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilojoulesPerKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilojoulesPerKelvin(math.NaN())
    if err == nil {
        t.Error("FromKilojoulesPerKelvin() with NaN value should return error")
    }

    _, err = factory.FromKilojoulesPerKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromKilojoulesPerKelvin() with +Inf value should return error")
    }

    _, err = factory.FromKilojoulesPerKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromKilojoulesPerKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilojoulesPerKelvin(0)
    if err != nil {
        t.Errorf("FromKilojoulesPerKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EntropyKilojoulePerKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilojoulesPerKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromMegajoulesPerKelvin function
func TestEntropyFactory_FromMegajoulesPerKelvin(t *testing.T) {
    factory := units.EntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegajoulesPerKelvin(100)
    if err != nil {
        t.Errorf("FromMegajoulesPerKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EntropyMegajoulePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegajoulesPerKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegajoulesPerKelvin(math.NaN())
    if err == nil {
        t.Error("FromMegajoulesPerKelvin() with NaN value should return error")
    }

    _, err = factory.FromMegajoulesPerKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromMegajoulesPerKelvin() with +Inf value should return error")
    }

    _, err = factory.FromMegajoulesPerKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromMegajoulesPerKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegajoulesPerKelvin(0)
    if err != nil {
        t.Errorf("FromMegajoulesPerKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EntropyMegajoulePerKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegajoulesPerKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromKilocaloriesPerKelvin function
func TestEntropyFactory_FromKilocaloriesPerKelvin(t *testing.T) {
    factory := units.EntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilocaloriesPerKelvin(100)
    if err != nil {
        t.Errorf("FromKilocaloriesPerKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EntropyKilocaloriePerKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilocaloriesPerKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilocaloriesPerKelvin(math.NaN())
    if err == nil {
        t.Error("FromKilocaloriesPerKelvin() with NaN value should return error")
    }

    _, err = factory.FromKilocaloriesPerKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromKilocaloriesPerKelvin() with +Inf value should return error")
    }

    _, err = factory.FromKilocaloriesPerKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromKilocaloriesPerKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilocaloriesPerKelvin(0)
    if err != nil {
        t.Errorf("FromKilocaloriesPerKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EntropyKilocaloriePerKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilocaloriesPerKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromKilojoulesPerDegreeCelsius function
func TestEntropyFactory_FromKilojoulesPerDegreeCelsius(t *testing.T) {
    factory := units.EntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilojoulesPerDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromKilojoulesPerDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EntropyKilojoulePerDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilojoulesPerDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilojoulesPerDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromKilojoulesPerDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromKilojoulesPerDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromKilojoulesPerDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromKilojoulesPerDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromKilojoulesPerDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilojoulesPerDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromKilojoulesPerDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EntropyKilojoulePerDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilojoulesPerDegreeCelsius() with zero value = %v, want 0", converted)
    }
}

func TestEntropyToString(t *testing.T) {
	factory := units.EntropyFactory{}
	a, err := factory.CreateEntropy(45, units.EntropyJoulePerKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.EntropyJoulePerKelvin, 2)
	expected := "45.00 " + units.GetEntropyAbbreviation(units.EntropyJoulePerKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.EntropyJoulePerKelvin, -1)
	expected = "45 " + units.GetEntropyAbbreviation(units.EntropyJoulePerKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestEntropy_EqualityAndComparison(t *testing.T) {
	factory := units.EntropyFactory{}
	a1, _ := factory.CreateEntropy(60, units.EntropyJoulePerKelvin)
	a2, _ := factory.CreateEntropy(60, units.EntropyJoulePerKelvin)
	a3, _ := factory.CreateEntropy(120, units.EntropyJoulePerKelvin)

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

func TestEntropy_Arithmetic(t *testing.T) {
	factory := units.EntropyFactory{}
	a1, _ := factory.CreateEntropy(30, units.EntropyJoulePerKelvin)
	a2, _ := factory.CreateEntropy(45, units.EntropyJoulePerKelvin)

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


func TestGetEntropyAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.EntropyUnits
        want string
    }{
        {
            name: "JoulePerKelvin abbreviation",
            unit: units.EntropyJoulePerKelvin,
            want: "J/K",
        },
        {
            name: "CaloriePerKelvin abbreviation",
            unit: units.EntropyCaloriePerKelvin,
            want: "cal/K",
        },
        {
            name: "JoulePerDegreeCelsius abbreviation",
            unit: units.EntropyJoulePerDegreeCelsius,
            want: "J/C",
        },
        {
            name: "KilojoulePerKelvin abbreviation",
            unit: units.EntropyKilojoulePerKelvin,
            want: "kJ/K",
        },
        {
            name: "MegajoulePerKelvin abbreviation",
            unit: units.EntropyMegajoulePerKelvin,
            want: "MJ/K",
        },
        {
            name: "KilocaloriePerKelvin abbreviation",
            unit: units.EntropyKilocaloriePerKelvin,
            want: "kcal/K",
        },
        {
            name: "KilojoulePerDegreeCelsius abbreviation",
            unit: units.EntropyKilojoulePerDegreeCelsius,
            want: "kJ/C",
        },
        {
            name: "invalid unit",
            unit: units.EntropyUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetEntropyAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetEntropyAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestEntropy_String(t *testing.T) {
    factory := units.EntropyFactory{}
    
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
            unit, err := factory.CreateEntropy(tt.value, units.EntropyJoulePerKelvin)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Entropy.String() = %v, want %v", got, tt.want)
            }
        })
    }
}