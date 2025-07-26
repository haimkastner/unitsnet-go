// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestVolumetricHeatCapacityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerCubicMeterKelvin"}`
	
	factory := units.VolumetricHeatCapacityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.VolumetricHeatCapacityJoulePerCubicMeterKelvin {
		t.Errorf("expected unit %v, got %v", units.VolumetricHeatCapacityJoulePerCubicMeterKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerCubicMeterKelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestVolumetricHeatCapacityDto_ToJSON(t *testing.T) {
	dto := units.VolumetricHeatCapacityDto{
		Value: 45,
		Unit:  units.VolumetricHeatCapacityJoulePerCubicMeterKelvin,
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
	if result["unit"].(string) != string(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin) {
		t.Errorf("expected unit %s, got %v", units.VolumetricHeatCapacityJoulePerCubicMeterKelvin, result["unit"])
	}
}

func TestNewVolumetricHeatCapacity_InvalidValue(t *testing.T) {
	factory := units.VolumetricHeatCapacityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateVolumetricHeatCapacity(math.NaN(), units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateVolumetricHeatCapacity(math.Inf(1), units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestVolumetricHeatCapacityConversions(t *testing.T) {
	factory := units.VolumetricHeatCapacityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateVolumetricHeatCapacity(180, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerCubicMeterKelvin.
		// No expected conversion value provided for JoulesPerCubicMeterKelvin, verifying result is not NaN.
		result := a.JoulesPerCubicMeterKelvin()
		cacheResult := a.JoulesPerCubicMeterKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to JoulesPerCubicMeterKelvin returned NaN")
		}
	}
	{
		// Test conversion to JoulesPerCubicMeterDegreeCelsius.
		// No expected conversion value provided for JoulesPerCubicMeterDegreeCelsius, verifying result is not NaN.
		result := a.JoulesPerCubicMeterDegreeCelsius()
		cacheResult := a.JoulesPerCubicMeterDegreeCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to JoulesPerCubicMeterDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to CaloriesPerCubicCentimeterDegreeCelsius.
		// No expected conversion value provided for CaloriesPerCubicCentimeterDegreeCelsius, verifying result is not NaN.
		result := a.CaloriesPerCubicCentimeterDegreeCelsius()
		cacheResult := a.CaloriesPerCubicCentimeterDegreeCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CaloriesPerCubicCentimeterDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to BtusPerCubicFootDegreeFahrenheit.
		// No expected conversion value provided for BtusPerCubicFootDegreeFahrenheit, verifying result is not NaN.
		result := a.BtusPerCubicFootDegreeFahrenheit()
		cacheResult := a.BtusPerCubicFootDegreeFahrenheit()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to BtusPerCubicFootDegreeFahrenheit returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerCubicMeterKelvin.
		// No expected conversion value provided for KilojoulesPerCubicMeterKelvin, verifying result is not NaN.
		result := a.KilojoulesPerCubicMeterKelvin()
		cacheResult := a.KilojoulesPerCubicMeterKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilojoulesPerCubicMeterKelvin returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerCubicMeterKelvin.
		// No expected conversion value provided for MegajoulesPerCubicMeterKelvin, verifying result is not NaN.
		result := a.MegajoulesPerCubicMeterKelvin()
		cacheResult := a.MegajoulesPerCubicMeterKelvin()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegajoulesPerCubicMeterKelvin returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerCubicMeterDegreeCelsius.
		// No expected conversion value provided for KilojoulesPerCubicMeterDegreeCelsius, verifying result is not NaN.
		result := a.KilojoulesPerCubicMeterDegreeCelsius()
		cacheResult := a.KilojoulesPerCubicMeterDegreeCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilojoulesPerCubicMeterDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerCubicMeterDegreeCelsius.
		// No expected conversion value provided for MegajoulesPerCubicMeterDegreeCelsius, verifying result is not NaN.
		result := a.MegajoulesPerCubicMeterDegreeCelsius()
		cacheResult := a.MegajoulesPerCubicMeterDegreeCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegajoulesPerCubicMeterDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerCubicCentimeterDegreeCelsius.
		// No expected conversion value provided for KilocaloriesPerCubicCentimeterDegreeCelsius, verifying result is not NaN.
		result := a.KilocaloriesPerCubicCentimeterDegreeCelsius()
		cacheResult := a.KilocaloriesPerCubicCentimeterDegreeCelsius()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilocaloriesPerCubicCentimeterDegreeCelsius returned NaN")
		}
	}
}

func TestVolumetricHeatCapacity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.VolumetricHeatCapacityFactory{}
	a, err := factory.CreateVolumetricHeatCapacity(90, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.VolumetricHeatCapacityJoulePerCubicMeterKelvin {
		t.Errorf("expected default unit JoulePerCubicMeterKelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.VolumetricHeatCapacityJoulePerCubicMeterKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.VolumetricHeatCapacityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.VolumetricHeatCapacityJoulePerCubicMeterKelvin {
		t.Errorf("expected unit JoulePerCubicMeterKelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestVolumetricHeatCapacityFactory_FromDto(t *testing.T) {
    factory := units.VolumetricHeatCapacityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.VolumetricHeatCapacityDto{
        Value: 100,
        Unit:  units.VolumetricHeatCapacityJoulePerCubicMeterKelvin,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.VolumetricHeatCapacityDto{
        Value: math.NaN(),
        Unit:  units.VolumetricHeatCapacityJoulePerCubicMeterKelvin,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test JoulePerCubicMeterKelvin conversion
    joules_per_cubic_meter_kelvinDto := units.VolumetricHeatCapacityDto{
        Value: 100,
        Unit:  units.VolumetricHeatCapacityJoulePerCubicMeterKelvin,
    }
    
    var joules_per_cubic_meter_kelvinResult *units.VolumetricHeatCapacity
    joules_per_cubic_meter_kelvinResult, err = factory.FromDto(joules_per_cubic_meter_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerCubicMeterKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_cubic_meter_kelvinResult.Convert(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerCubicMeterKelvin = %v, want %v", converted, 100)
    }
    // Test JoulePerCubicMeterDegreeCelsius conversion
    joules_per_cubic_meter_degree_celsiusDto := units.VolumetricHeatCapacityDto{
        Value: 100,
        Unit:  units.VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius,
    }
    
    var joules_per_cubic_meter_degree_celsiusResult *units.VolumetricHeatCapacity
    joules_per_cubic_meter_degree_celsiusResult, err = factory.FromDto(joules_per_cubic_meter_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerCubicMeterDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_cubic_meter_degree_celsiusResult.Convert(units.VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerCubicMeterDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test CaloriePerCubicCentimeterDegreeCelsius conversion
    calories_per_cubic_centimeter_degree_celsiusDto := units.VolumetricHeatCapacityDto{
        Value: 100,
        Unit:  units.VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius,
    }
    
    var calories_per_cubic_centimeter_degree_celsiusResult *units.VolumetricHeatCapacity
    calories_per_cubic_centimeter_degree_celsiusResult, err = factory.FromDto(calories_per_cubic_centimeter_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with CaloriePerCubicCentimeterDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = calories_per_cubic_centimeter_degree_celsiusResult.Convert(units.VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CaloriePerCubicCentimeterDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test BtuPerCubicFootDegreeFahrenheit conversion
    btus_per_cubic_foot_degree_fahrenheitDto := units.VolumetricHeatCapacityDto{
        Value: 100,
        Unit:  units.VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit,
    }
    
    var btus_per_cubic_foot_degree_fahrenheitResult *units.VolumetricHeatCapacity
    btus_per_cubic_foot_degree_fahrenheitResult, err = factory.FromDto(btus_per_cubic_foot_degree_fahrenheitDto)
    if err != nil {
        t.Errorf("FromDto() with BtuPerCubicFootDegreeFahrenheit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_cubic_foot_degree_fahrenheitResult.Convert(units.VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerCubicFootDegreeFahrenheit = %v, want %v", converted, 100)
    }
    // Test KilojoulePerCubicMeterKelvin conversion
    kilojoules_per_cubic_meter_kelvinDto := units.VolumetricHeatCapacityDto{
        Value: 100,
        Unit:  units.VolumetricHeatCapacityKilojoulePerCubicMeterKelvin,
    }
    
    var kilojoules_per_cubic_meter_kelvinResult *units.VolumetricHeatCapacity
    kilojoules_per_cubic_meter_kelvinResult, err = factory.FromDto(kilojoules_per_cubic_meter_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with KilojoulePerCubicMeterKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_cubic_meter_kelvinResult.Convert(units.VolumetricHeatCapacityKilojoulePerCubicMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerCubicMeterKelvin = %v, want %v", converted, 100)
    }
    // Test MegajoulePerCubicMeterKelvin conversion
    megajoules_per_cubic_meter_kelvinDto := units.VolumetricHeatCapacityDto{
        Value: 100,
        Unit:  units.VolumetricHeatCapacityMegajoulePerCubicMeterKelvin,
    }
    
    var megajoules_per_cubic_meter_kelvinResult *units.VolumetricHeatCapacity
    megajoules_per_cubic_meter_kelvinResult, err = factory.FromDto(megajoules_per_cubic_meter_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with MegajoulePerCubicMeterKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_cubic_meter_kelvinResult.Convert(units.VolumetricHeatCapacityMegajoulePerCubicMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerCubicMeterKelvin = %v, want %v", converted, 100)
    }
    // Test KilojoulePerCubicMeterDegreeCelsius conversion
    kilojoules_per_cubic_meter_degree_celsiusDto := units.VolumetricHeatCapacityDto{
        Value: 100,
        Unit:  units.VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius,
    }
    
    var kilojoules_per_cubic_meter_degree_celsiusResult *units.VolumetricHeatCapacity
    kilojoules_per_cubic_meter_degree_celsiusResult, err = factory.FromDto(kilojoules_per_cubic_meter_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with KilojoulePerCubicMeterDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_cubic_meter_degree_celsiusResult.Convert(units.VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerCubicMeterDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test MegajoulePerCubicMeterDegreeCelsius conversion
    megajoules_per_cubic_meter_degree_celsiusDto := units.VolumetricHeatCapacityDto{
        Value: 100,
        Unit:  units.VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius,
    }
    
    var megajoules_per_cubic_meter_degree_celsiusResult *units.VolumetricHeatCapacity
    megajoules_per_cubic_meter_degree_celsiusResult, err = factory.FromDto(megajoules_per_cubic_meter_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with MegajoulePerCubicMeterDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_cubic_meter_degree_celsiusResult.Convert(units.VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerCubicMeterDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test KilocaloriePerCubicCentimeterDegreeCelsius conversion
    kilocalories_per_cubic_centimeter_degree_celsiusDto := units.VolumetricHeatCapacityDto{
        Value: 100,
        Unit:  units.VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius,
    }
    
    var kilocalories_per_cubic_centimeter_degree_celsiusResult *units.VolumetricHeatCapacity
    kilocalories_per_cubic_centimeter_degree_celsiusResult, err = factory.FromDto(kilocalories_per_cubic_centimeter_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with KilocaloriePerCubicCentimeterDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_cubic_centimeter_degree_celsiusResult.Convert(units.VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerCubicCentimeterDegreeCelsius = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.VolumetricHeatCapacityDto{
        Value: 0,
        Unit:  units.VolumetricHeatCapacityJoulePerCubicMeterKelvin,
    }
    
    var zeroResult *units.VolumetricHeatCapacity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestVolumetricHeatCapacityFactory_FromDtoJSON(t *testing.T) {
    factory := units.VolumetricHeatCapacityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "JoulePerCubicMeterKelvin"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "JoulePerCubicMeterKelvin"}`)
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
    nanJSON, _ := json.Marshal(units.VolumetricHeatCapacityDto{
        Value: nanValue,
        Unit:  units.VolumetricHeatCapacityJoulePerCubicMeterKelvin,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with JoulePerCubicMeterKelvin unit
    joules_per_cubic_meter_kelvinJSON := []byte(`{"value": 100, "unit": "JoulePerCubicMeterKelvin"}`)
    joules_per_cubic_meter_kelvinResult, err := factory.FromDtoJSON(joules_per_cubic_meter_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerCubicMeterKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_cubic_meter_kelvinResult.Convert(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerCubicMeterKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with JoulePerCubicMeterDegreeCelsius unit
    joules_per_cubic_meter_degree_celsiusJSON := []byte(`{"value": 100, "unit": "JoulePerCubicMeterDegreeCelsius"}`)
    joules_per_cubic_meter_degree_celsiusResult, err := factory.FromDtoJSON(joules_per_cubic_meter_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerCubicMeterDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_cubic_meter_degree_celsiusResult.Convert(units.VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerCubicMeterDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with CaloriePerCubicCentimeterDegreeCelsius unit
    calories_per_cubic_centimeter_degree_celsiusJSON := []byte(`{"value": 100, "unit": "CaloriePerCubicCentimeterDegreeCelsius"}`)
    calories_per_cubic_centimeter_degree_celsiusResult, err := factory.FromDtoJSON(calories_per_cubic_centimeter_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CaloriePerCubicCentimeterDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = calories_per_cubic_centimeter_degree_celsiusResult.Convert(units.VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CaloriePerCubicCentimeterDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with BtuPerCubicFootDegreeFahrenheit unit
    btus_per_cubic_foot_degree_fahrenheitJSON := []byte(`{"value": 100, "unit": "BtuPerCubicFootDegreeFahrenheit"}`)
    btus_per_cubic_foot_degree_fahrenheitResult, err := factory.FromDtoJSON(btus_per_cubic_foot_degree_fahrenheitJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BtuPerCubicFootDegreeFahrenheit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_cubic_foot_degree_fahrenheitResult.Convert(units.VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerCubicFootDegreeFahrenheit = %v, want %v", converted, 100)
    }
    // Test JSON with KilojoulePerCubicMeterKelvin unit
    kilojoules_per_cubic_meter_kelvinJSON := []byte(`{"value": 100, "unit": "KilojoulePerCubicMeterKelvin"}`)
    kilojoules_per_cubic_meter_kelvinResult, err := factory.FromDtoJSON(kilojoules_per_cubic_meter_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilojoulePerCubicMeterKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_cubic_meter_kelvinResult.Convert(units.VolumetricHeatCapacityKilojoulePerCubicMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerCubicMeterKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with MegajoulePerCubicMeterKelvin unit
    megajoules_per_cubic_meter_kelvinJSON := []byte(`{"value": 100, "unit": "MegajoulePerCubicMeterKelvin"}`)
    megajoules_per_cubic_meter_kelvinResult, err := factory.FromDtoJSON(megajoules_per_cubic_meter_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegajoulePerCubicMeterKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_cubic_meter_kelvinResult.Convert(units.VolumetricHeatCapacityMegajoulePerCubicMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerCubicMeterKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with KilojoulePerCubicMeterDegreeCelsius unit
    kilojoules_per_cubic_meter_degree_celsiusJSON := []byte(`{"value": 100, "unit": "KilojoulePerCubicMeterDegreeCelsius"}`)
    kilojoules_per_cubic_meter_degree_celsiusResult, err := factory.FromDtoJSON(kilojoules_per_cubic_meter_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilojoulePerCubicMeterDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_cubic_meter_degree_celsiusResult.Convert(units.VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerCubicMeterDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with MegajoulePerCubicMeterDegreeCelsius unit
    megajoules_per_cubic_meter_degree_celsiusJSON := []byte(`{"value": 100, "unit": "MegajoulePerCubicMeterDegreeCelsius"}`)
    megajoules_per_cubic_meter_degree_celsiusResult, err := factory.FromDtoJSON(megajoules_per_cubic_meter_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegajoulePerCubicMeterDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_cubic_meter_degree_celsiusResult.Convert(units.VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerCubicMeterDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with KilocaloriePerCubicCentimeterDegreeCelsius unit
    kilocalories_per_cubic_centimeter_degree_celsiusJSON := []byte(`{"value": 100, "unit": "KilocaloriePerCubicCentimeterDegreeCelsius"}`)
    kilocalories_per_cubic_centimeter_degree_celsiusResult, err := factory.FromDtoJSON(kilocalories_per_cubic_centimeter_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilocaloriePerCubicCentimeterDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_cubic_centimeter_degree_celsiusResult.Convert(units.VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerCubicCentimeterDegreeCelsius = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "JoulePerCubicMeterKelvin"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromJoulesPerCubicMeterKelvin function
func TestVolumetricHeatCapacityFactory_FromJoulesPerCubicMeterKelvin(t *testing.T) {
    factory := units.VolumetricHeatCapacityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerCubicMeterKelvin(100)
    if err != nil {
        t.Errorf("FromJoulesPerCubicMeterKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerCubicMeterKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerCubicMeterKelvin(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerCubicMeterKelvin() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerCubicMeterKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerCubicMeterKelvin() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerCubicMeterKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerCubicMeterKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerCubicMeterKelvin(0)
    if err != nil {
        t.Errorf("FromJoulesPerCubicMeterKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerCubicMeterKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromJoulesPerCubicMeterDegreeCelsius function
func TestVolumetricHeatCapacityFactory_FromJoulesPerCubicMeterDegreeCelsius(t *testing.T) {
    factory := units.VolumetricHeatCapacityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerCubicMeterDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromJoulesPerCubicMeterDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerCubicMeterDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerCubicMeterDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerCubicMeterDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerCubicMeterDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerCubicMeterDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerCubicMeterDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerCubicMeterDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerCubicMeterDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromJoulesPerCubicMeterDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerCubicMeterDegreeCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromCaloriesPerCubicCentimeterDegreeCelsius function
func TestVolumetricHeatCapacityFactory_FromCaloriesPerCubicCentimeterDegreeCelsius(t *testing.T) {
    factory := units.VolumetricHeatCapacityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCaloriesPerCubicCentimeterDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromCaloriesPerCubicCentimeterDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCaloriesPerCubicCentimeterDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCaloriesPerCubicCentimeterDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromCaloriesPerCubicCentimeterDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromCaloriesPerCubicCentimeterDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromCaloriesPerCubicCentimeterDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromCaloriesPerCubicCentimeterDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromCaloriesPerCubicCentimeterDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCaloriesPerCubicCentimeterDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromCaloriesPerCubicCentimeterDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCaloriesPerCubicCentimeterDegreeCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromBtusPerCubicFootDegreeFahrenheit function
func TestVolumetricHeatCapacityFactory_FromBtusPerCubicFootDegreeFahrenheit(t *testing.T) {
    factory := units.VolumetricHeatCapacityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBtusPerCubicFootDegreeFahrenheit(100)
    if err != nil {
        t.Errorf("FromBtusPerCubicFootDegreeFahrenheit() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBtusPerCubicFootDegreeFahrenheit() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBtusPerCubicFootDegreeFahrenheit(math.NaN())
    if err == nil {
        t.Error("FromBtusPerCubicFootDegreeFahrenheit() with NaN value should return error")
    }

    _, err = factory.FromBtusPerCubicFootDegreeFahrenheit(math.Inf(1))
    if err == nil {
        t.Error("FromBtusPerCubicFootDegreeFahrenheit() with +Inf value should return error")
    }

    _, err = factory.FromBtusPerCubicFootDegreeFahrenheit(math.Inf(-1))
    if err == nil {
        t.Error("FromBtusPerCubicFootDegreeFahrenheit() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBtusPerCubicFootDegreeFahrenheit(0)
    if err != nil {
        t.Errorf("FromBtusPerCubicFootDegreeFahrenheit() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBtusPerCubicFootDegreeFahrenheit() with zero value = %v, want 0", converted)
    }
}
// Test FromKilojoulesPerCubicMeterKelvin function
func TestVolumetricHeatCapacityFactory_FromKilojoulesPerCubicMeterKelvin(t *testing.T) {
    factory := units.VolumetricHeatCapacityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilojoulesPerCubicMeterKelvin(100)
    if err != nil {
        t.Errorf("FromKilojoulesPerCubicMeterKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumetricHeatCapacityKilojoulePerCubicMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilojoulesPerCubicMeterKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilojoulesPerCubicMeterKelvin(math.NaN())
    if err == nil {
        t.Error("FromKilojoulesPerCubicMeterKelvin() with NaN value should return error")
    }

    _, err = factory.FromKilojoulesPerCubicMeterKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromKilojoulesPerCubicMeterKelvin() with +Inf value should return error")
    }

    _, err = factory.FromKilojoulesPerCubicMeterKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromKilojoulesPerCubicMeterKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilojoulesPerCubicMeterKelvin(0)
    if err != nil {
        t.Errorf("FromKilojoulesPerCubicMeterKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumetricHeatCapacityKilojoulePerCubicMeterKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilojoulesPerCubicMeterKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromMegajoulesPerCubicMeterKelvin function
func TestVolumetricHeatCapacityFactory_FromMegajoulesPerCubicMeterKelvin(t *testing.T) {
    factory := units.VolumetricHeatCapacityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegajoulesPerCubicMeterKelvin(100)
    if err != nil {
        t.Errorf("FromMegajoulesPerCubicMeterKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumetricHeatCapacityMegajoulePerCubicMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegajoulesPerCubicMeterKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegajoulesPerCubicMeterKelvin(math.NaN())
    if err == nil {
        t.Error("FromMegajoulesPerCubicMeterKelvin() with NaN value should return error")
    }

    _, err = factory.FromMegajoulesPerCubicMeterKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromMegajoulesPerCubicMeterKelvin() with +Inf value should return error")
    }

    _, err = factory.FromMegajoulesPerCubicMeterKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromMegajoulesPerCubicMeterKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegajoulesPerCubicMeterKelvin(0)
    if err != nil {
        t.Errorf("FromMegajoulesPerCubicMeterKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumetricHeatCapacityMegajoulePerCubicMeterKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegajoulesPerCubicMeterKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromKilojoulesPerCubicMeterDegreeCelsius function
func TestVolumetricHeatCapacityFactory_FromKilojoulesPerCubicMeterDegreeCelsius(t *testing.T) {
    factory := units.VolumetricHeatCapacityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilojoulesPerCubicMeterDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromKilojoulesPerCubicMeterDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilojoulesPerCubicMeterDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilojoulesPerCubicMeterDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromKilojoulesPerCubicMeterDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromKilojoulesPerCubicMeterDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromKilojoulesPerCubicMeterDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromKilojoulesPerCubicMeterDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromKilojoulesPerCubicMeterDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilojoulesPerCubicMeterDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromKilojoulesPerCubicMeterDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilojoulesPerCubicMeterDegreeCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromMegajoulesPerCubicMeterDegreeCelsius function
func TestVolumetricHeatCapacityFactory_FromMegajoulesPerCubicMeterDegreeCelsius(t *testing.T) {
    factory := units.VolumetricHeatCapacityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegajoulesPerCubicMeterDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromMegajoulesPerCubicMeterDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegajoulesPerCubicMeterDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegajoulesPerCubicMeterDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromMegajoulesPerCubicMeterDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromMegajoulesPerCubicMeterDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromMegajoulesPerCubicMeterDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromMegajoulesPerCubicMeterDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromMegajoulesPerCubicMeterDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegajoulesPerCubicMeterDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromMegajoulesPerCubicMeterDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegajoulesPerCubicMeterDegreeCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromKilocaloriesPerCubicCentimeterDegreeCelsius function
func TestVolumetricHeatCapacityFactory_FromKilocaloriesPerCubicCentimeterDegreeCelsius(t *testing.T) {
    factory := units.VolumetricHeatCapacityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilocaloriesPerCubicCentimeterDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromKilocaloriesPerCubicCentimeterDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilocaloriesPerCubicCentimeterDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilocaloriesPerCubicCentimeterDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromKilocaloriesPerCubicCentimeterDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromKilocaloriesPerCubicCentimeterDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromKilocaloriesPerCubicCentimeterDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromKilocaloriesPerCubicCentimeterDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromKilocaloriesPerCubicCentimeterDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilocaloriesPerCubicCentimeterDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromKilocaloriesPerCubicCentimeterDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilocaloriesPerCubicCentimeterDegreeCelsius() with zero value = %v, want 0", converted)
    }
}

func TestVolumetricHeatCapacityToString(t *testing.T) {
	factory := units.VolumetricHeatCapacityFactory{}
	a, err := factory.CreateVolumetricHeatCapacity(45, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin, 2)
	expected := "45.00 " + units.GetVolumetricHeatCapacityAbbreviation(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin, -1)
	expected = "45 " + units.GetVolumetricHeatCapacityAbbreviation(units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestVolumetricHeatCapacity_EqualityAndComparison(t *testing.T) {
	factory := units.VolumetricHeatCapacityFactory{}
	a1, _ := factory.CreateVolumetricHeatCapacity(60, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	a2, _ := factory.CreateVolumetricHeatCapacity(60, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	a3, _ := factory.CreateVolumetricHeatCapacity(120, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)

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

func TestVolumetricHeatCapacity_Arithmetic(t *testing.T) {
	factory := units.VolumetricHeatCapacityFactory{}
	a1, _ := factory.CreateVolumetricHeatCapacity(30, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	a2, _ := factory.CreateVolumetricHeatCapacity(45, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)

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


func TestGetVolumetricHeatCapacityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.VolumetricHeatCapacityUnits
        want string
    }{
        {
            name: "JoulePerCubicMeterKelvin abbreviation",
            unit: units.VolumetricHeatCapacityJoulePerCubicMeterKelvin,
            want: "J/(m³·K)",
        },
        {
            name: "JoulePerCubicMeterDegreeCelsius abbreviation",
            unit: units.VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius,
            want: "J/(m³·°C)",
        },
        {
            name: "CaloriePerCubicCentimeterDegreeCelsius abbreviation",
            unit: units.VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius,
            want: "cal/(cm³·°C)",
        },
        {
            name: "BtuPerCubicFootDegreeFahrenheit abbreviation",
            unit: units.VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit,
            want: "BTU/(ft³·°F)",
        },
        {
            name: "KilojoulePerCubicMeterKelvin abbreviation",
            unit: units.VolumetricHeatCapacityKilojoulePerCubicMeterKelvin,
            want: "kJ/(m³·K)",
        },
        {
            name: "MegajoulePerCubicMeterKelvin abbreviation",
            unit: units.VolumetricHeatCapacityMegajoulePerCubicMeterKelvin,
            want: "MJ/(m³·K)",
        },
        {
            name: "KilojoulePerCubicMeterDegreeCelsius abbreviation",
            unit: units.VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius,
            want: "kJ/(m³·°C)",
        },
        {
            name: "MegajoulePerCubicMeterDegreeCelsius abbreviation",
            unit: units.VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius,
            want: "MJ/(m³·°C)",
        },
        {
            name: "KilocaloriePerCubicCentimeterDegreeCelsius abbreviation",
            unit: units.VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius,
            want: "kcal/(cm³·°C)",
        },
        {
            name: "invalid unit",
            unit: units.VolumetricHeatCapacityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetVolumetricHeatCapacityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetVolumetricHeatCapacityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestVolumetricHeatCapacity_String(t *testing.T) {
    factory := units.VolumetricHeatCapacityFactory{}
    
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
            unit, err := factory.CreateVolumetricHeatCapacity(tt.value, units.VolumetricHeatCapacityJoulePerCubicMeterKelvin)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("VolumetricHeatCapacity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestVolumetricHeatCapacity_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.VolumetricHeatCapacityFactory{}

	_, err := uf.CreateVolumetricHeatCapacity(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}