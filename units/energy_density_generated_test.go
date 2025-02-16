// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestEnergyDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerCubicMeter"}`
	
	factory := units.EnergyDensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.EnergyDensityJoulePerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.EnergyDensityJoulePerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestEnergyDensityDto_ToJSON(t *testing.T) {
	dto := units.EnergyDensityDto{
		Value: 45,
		Unit:  units.EnergyDensityJoulePerCubicMeter,
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
	if result["unit"].(string) != string(units.EnergyDensityJoulePerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.EnergyDensityJoulePerCubicMeter, result["unit"])
	}
}

func TestNewEnergyDensity_InvalidValue(t *testing.T) {
	factory := units.EnergyDensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateEnergyDensity(math.NaN(), units.EnergyDensityJoulePerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateEnergyDensity(math.Inf(1), units.EnergyDensityJoulePerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestEnergyDensityConversions(t *testing.T) {
	factory := units.EnergyDensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateEnergyDensity(180, units.EnergyDensityJoulePerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerCubicMeter.
		// No expected conversion value provided for JoulesPerCubicMeter, verifying result is not NaN.
		result := a.JoulesPerCubicMeter()
		cacheResult := a.JoulesPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to JoulesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to WattHoursPerCubicMeter.
		// No expected conversion value provided for WattHoursPerCubicMeter, verifying result is not NaN.
		result := a.WattHoursPerCubicMeter()
		cacheResult := a.WattHoursPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattHoursPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerCubicMeter.
		// No expected conversion value provided for KilojoulesPerCubicMeter, verifying result is not NaN.
		result := a.KilojoulesPerCubicMeter()
		cacheResult := a.KilojoulesPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilojoulesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerCubicMeter.
		// No expected conversion value provided for MegajoulesPerCubicMeter, verifying result is not NaN.
		result := a.MegajoulesPerCubicMeter()
		cacheResult := a.MegajoulesPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegajoulesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to GigajoulesPerCubicMeter.
		// No expected conversion value provided for GigajoulesPerCubicMeter, verifying result is not NaN.
		result := a.GigajoulesPerCubicMeter()
		cacheResult := a.GigajoulesPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigajoulesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to TerajoulesPerCubicMeter.
		// No expected conversion value provided for TerajoulesPerCubicMeter, verifying result is not NaN.
		result := a.TerajoulesPerCubicMeter()
		cacheResult := a.TerajoulesPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TerajoulesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PetajoulesPerCubicMeter.
		// No expected conversion value provided for PetajoulesPerCubicMeter, verifying result is not NaN.
		result := a.PetajoulesPerCubicMeter()
		cacheResult := a.PetajoulesPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PetajoulesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattHoursPerCubicMeter.
		// No expected conversion value provided for KilowattHoursPerCubicMeter, verifying result is not NaN.
		result := a.KilowattHoursPerCubicMeter()
		cacheResult := a.KilowattHoursPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattHoursPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattHoursPerCubicMeter.
		// No expected conversion value provided for MegawattHoursPerCubicMeter, verifying result is not NaN.
		result := a.MegawattHoursPerCubicMeter()
		cacheResult := a.MegawattHoursPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattHoursPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to GigawattHoursPerCubicMeter.
		// No expected conversion value provided for GigawattHoursPerCubicMeter, verifying result is not NaN.
		result := a.GigawattHoursPerCubicMeter()
		cacheResult := a.GigawattHoursPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattHoursPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to TerawattHoursPerCubicMeter.
		// No expected conversion value provided for TerawattHoursPerCubicMeter, verifying result is not NaN.
		result := a.TerawattHoursPerCubicMeter()
		cacheResult := a.TerawattHoursPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TerawattHoursPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PetawattHoursPerCubicMeter.
		// No expected conversion value provided for PetawattHoursPerCubicMeter, verifying result is not NaN.
		result := a.PetawattHoursPerCubicMeter()
		cacheResult := a.PetawattHoursPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PetawattHoursPerCubicMeter returned NaN")
		}
	}
}

func TestEnergyDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.EnergyDensityFactory{}
	a, err := factory.CreateEnergyDensity(90, units.EnergyDensityJoulePerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.EnergyDensityJoulePerCubicMeter {
		t.Errorf("expected default unit JoulePerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.EnergyDensityJoulePerCubicMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.EnergyDensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.EnergyDensityJoulePerCubicMeter {
		t.Errorf("expected unit JoulePerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestEnergyDensityFactory_FromDto(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.EnergyDensityDto{
        Value: 100,
        Unit:  units.EnergyDensityJoulePerCubicMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.EnergyDensityDto{
        Value: math.NaN(),
        Unit:  units.EnergyDensityJoulePerCubicMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test JoulePerCubicMeter conversion
    joules_per_cubic_meterDto := units.EnergyDensityDto{
        Value: 100,
        Unit:  units.EnergyDensityJoulePerCubicMeter,
    }
    
    var joules_per_cubic_meterResult *units.EnergyDensity
    joules_per_cubic_meterResult, err = factory.FromDto(joules_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_cubic_meterResult.Convert(units.EnergyDensityJoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test WattHourPerCubicMeter conversion
    watt_hours_per_cubic_meterDto := units.EnergyDensityDto{
        Value: 100,
        Unit:  units.EnergyDensityWattHourPerCubicMeter,
    }
    
    var watt_hours_per_cubic_meterResult *units.EnergyDensity
    watt_hours_per_cubic_meterResult, err = factory.FromDto(watt_hours_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with WattHourPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_hours_per_cubic_meterResult.Convert(units.EnergyDensityWattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattHourPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test KilojoulePerCubicMeter conversion
    kilojoules_per_cubic_meterDto := units.EnergyDensityDto{
        Value: 100,
        Unit:  units.EnergyDensityKilojoulePerCubicMeter,
    }
    
    var kilojoules_per_cubic_meterResult *units.EnergyDensity
    kilojoules_per_cubic_meterResult, err = factory.FromDto(kilojoules_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilojoulePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_cubic_meterResult.Convert(units.EnergyDensityKilojoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test MegajoulePerCubicMeter conversion
    megajoules_per_cubic_meterDto := units.EnergyDensityDto{
        Value: 100,
        Unit:  units.EnergyDensityMegajoulePerCubicMeter,
    }
    
    var megajoules_per_cubic_meterResult *units.EnergyDensity
    megajoules_per_cubic_meterResult, err = factory.FromDto(megajoules_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MegajoulePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_cubic_meterResult.Convert(units.EnergyDensityMegajoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test GigajoulePerCubicMeter conversion
    gigajoules_per_cubic_meterDto := units.EnergyDensityDto{
        Value: 100,
        Unit:  units.EnergyDensityGigajoulePerCubicMeter,
    }
    
    var gigajoules_per_cubic_meterResult *units.EnergyDensity
    gigajoules_per_cubic_meterResult, err = factory.FromDto(gigajoules_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with GigajoulePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigajoules_per_cubic_meterResult.Convert(units.EnergyDensityGigajoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigajoulePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test TerajoulePerCubicMeter conversion
    terajoules_per_cubic_meterDto := units.EnergyDensityDto{
        Value: 100,
        Unit:  units.EnergyDensityTerajoulePerCubicMeter,
    }
    
    var terajoules_per_cubic_meterResult *units.EnergyDensity
    terajoules_per_cubic_meterResult, err = factory.FromDto(terajoules_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with TerajoulePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terajoules_per_cubic_meterResult.Convert(units.EnergyDensityTerajoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerajoulePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test PetajoulePerCubicMeter conversion
    petajoules_per_cubic_meterDto := units.EnergyDensityDto{
        Value: 100,
        Unit:  units.EnergyDensityPetajoulePerCubicMeter,
    }
    
    var petajoules_per_cubic_meterResult *units.EnergyDensity
    petajoules_per_cubic_meterResult, err = factory.FromDto(petajoules_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with PetajoulePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petajoules_per_cubic_meterResult.Convert(units.EnergyDensityPetajoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PetajoulePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test KilowattHourPerCubicMeter conversion
    kilowatt_hours_per_cubic_meterDto := units.EnergyDensityDto{
        Value: 100,
        Unit:  units.EnergyDensityKilowattHourPerCubicMeter,
    }
    
    var kilowatt_hours_per_cubic_meterResult *units.EnergyDensity
    kilowatt_hours_per_cubic_meterResult, err = factory.FromDto(kilowatt_hours_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattHourPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_hours_per_cubic_meterResult.Convert(units.EnergyDensityKilowattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattHourPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test MegawattHourPerCubicMeter conversion
    megawatt_hours_per_cubic_meterDto := units.EnergyDensityDto{
        Value: 100,
        Unit:  units.EnergyDensityMegawattHourPerCubicMeter,
    }
    
    var megawatt_hours_per_cubic_meterResult *units.EnergyDensity
    megawatt_hours_per_cubic_meterResult, err = factory.FromDto(megawatt_hours_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattHourPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_hours_per_cubic_meterResult.Convert(units.EnergyDensityMegawattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattHourPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test GigawattHourPerCubicMeter conversion
    gigawatt_hours_per_cubic_meterDto := units.EnergyDensityDto{
        Value: 100,
        Unit:  units.EnergyDensityGigawattHourPerCubicMeter,
    }
    
    var gigawatt_hours_per_cubic_meterResult *units.EnergyDensity
    gigawatt_hours_per_cubic_meterResult, err = factory.FromDto(gigawatt_hours_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattHourPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_hours_per_cubic_meterResult.Convert(units.EnergyDensityGigawattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattHourPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test TerawattHourPerCubicMeter conversion
    terawatt_hours_per_cubic_meterDto := units.EnergyDensityDto{
        Value: 100,
        Unit:  units.EnergyDensityTerawattHourPerCubicMeter,
    }
    
    var terawatt_hours_per_cubic_meterResult *units.EnergyDensity
    terawatt_hours_per_cubic_meterResult, err = factory.FromDto(terawatt_hours_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with TerawattHourPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatt_hours_per_cubic_meterResult.Convert(units.EnergyDensityTerawattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattHourPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test PetawattHourPerCubicMeter conversion
    petawatt_hours_per_cubic_meterDto := units.EnergyDensityDto{
        Value: 100,
        Unit:  units.EnergyDensityPetawattHourPerCubicMeter,
    }
    
    var petawatt_hours_per_cubic_meterResult *units.EnergyDensity
    petawatt_hours_per_cubic_meterResult, err = factory.FromDto(petawatt_hours_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with PetawattHourPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petawatt_hours_per_cubic_meterResult.Convert(units.EnergyDensityPetawattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PetawattHourPerCubicMeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.EnergyDensityDto{
        Value: 0,
        Unit:  units.EnergyDensityJoulePerCubicMeter,
    }
    
    var zeroResult *units.EnergyDensity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestEnergyDensityFactory_FromDtoJSON(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "JoulePerCubicMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "JoulePerCubicMeter"}`)
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
    nanJSON, _ := json.Marshal(units.EnergyDensityDto{
        Value: nanValue,
        Unit:  units.EnergyDensityJoulePerCubicMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with JoulePerCubicMeter unit
    joules_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "JoulePerCubicMeter"}`)
    joules_per_cubic_meterResult, err := factory.FromDtoJSON(joules_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_cubic_meterResult.Convert(units.EnergyDensityJoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with WattHourPerCubicMeter unit
    watt_hours_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "WattHourPerCubicMeter"}`)
    watt_hours_per_cubic_meterResult, err := factory.FromDtoJSON(watt_hours_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattHourPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_hours_per_cubic_meterResult.Convert(units.EnergyDensityWattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattHourPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilojoulePerCubicMeter unit
    kilojoules_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "KilojoulePerCubicMeter"}`)
    kilojoules_per_cubic_meterResult, err := factory.FromDtoJSON(kilojoules_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilojoulePerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_cubic_meterResult.Convert(units.EnergyDensityKilojoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegajoulePerCubicMeter unit
    megajoules_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "MegajoulePerCubicMeter"}`)
    megajoules_per_cubic_meterResult, err := factory.FromDtoJSON(megajoules_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegajoulePerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_cubic_meterResult.Convert(units.EnergyDensityMegajoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with GigajoulePerCubicMeter unit
    gigajoules_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "GigajoulePerCubicMeter"}`)
    gigajoules_per_cubic_meterResult, err := factory.FromDtoJSON(gigajoules_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigajoulePerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigajoules_per_cubic_meterResult.Convert(units.EnergyDensityGigajoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigajoulePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with TerajoulePerCubicMeter unit
    terajoules_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "TerajoulePerCubicMeter"}`)
    terajoules_per_cubic_meterResult, err := factory.FromDtoJSON(terajoules_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TerajoulePerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terajoules_per_cubic_meterResult.Convert(units.EnergyDensityTerajoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerajoulePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PetajoulePerCubicMeter unit
    petajoules_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "PetajoulePerCubicMeter"}`)
    petajoules_per_cubic_meterResult, err := factory.FromDtoJSON(petajoules_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PetajoulePerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petajoules_per_cubic_meterResult.Convert(units.EnergyDensityPetajoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PetajoulePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattHourPerCubicMeter unit
    kilowatt_hours_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "KilowattHourPerCubicMeter"}`)
    kilowatt_hours_per_cubic_meterResult, err := factory.FromDtoJSON(kilowatt_hours_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattHourPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_hours_per_cubic_meterResult.Convert(units.EnergyDensityKilowattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattHourPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattHourPerCubicMeter unit
    megawatt_hours_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "MegawattHourPerCubicMeter"}`)
    megawatt_hours_per_cubic_meterResult, err := factory.FromDtoJSON(megawatt_hours_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattHourPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_hours_per_cubic_meterResult.Convert(units.EnergyDensityMegawattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattHourPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattHourPerCubicMeter unit
    gigawatt_hours_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "GigawattHourPerCubicMeter"}`)
    gigawatt_hours_per_cubic_meterResult, err := factory.FromDtoJSON(gigawatt_hours_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattHourPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_hours_per_cubic_meterResult.Convert(units.EnergyDensityGigawattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattHourPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with TerawattHourPerCubicMeter unit
    terawatt_hours_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "TerawattHourPerCubicMeter"}`)
    terawatt_hours_per_cubic_meterResult, err := factory.FromDtoJSON(terawatt_hours_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TerawattHourPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatt_hours_per_cubic_meterResult.Convert(units.EnergyDensityTerawattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattHourPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PetawattHourPerCubicMeter unit
    petawatt_hours_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "PetawattHourPerCubicMeter"}`)
    petawatt_hours_per_cubic_meterResult, err := factory.FromDtoJSON(petawatt_hours_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PetawattHourPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petawatt_hours_per_cubic_meterResult.Convert(units.EnergyDensityPetawattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PetawattHourPerCubicMeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "JoulePerCubicMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromJoulesPerCubicMeter function
func TestEnergyDensityFactory_FromJoulesPerCubicMeter(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromJoulesPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDensityJoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromJoulesPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDensityJoulePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromWattHoursPerCubicMeter function
func TestEnergyDensityFactory_FromWattHoursPerCubicMeter(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattHoursPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromWattHoursPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDensityWattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattHoursPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattHoursPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromWattHoursPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromWattHoursPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromWattHoursPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromWattHoursPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromWattHoursPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattHoursPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromWattHoursPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDensityWattHourPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattHoursPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilojoulesPerCubicMeter function
func TestEnergyDensityFactory_FromKilojoulesPerCubicMeter(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilojoulesPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromKilojoulesPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDensityKilojoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilojoulesPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilojoulesPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromKilojoulesPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromKilojoulesPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilojoulesPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilojoulesPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilojoulesPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilojoulesPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromKilojoulesPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDensityKilojoulePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilojoulesPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMegajoulesPerCubicMeter function
func TestEnergyDensityFactory_FromMegajoulesPerCubicMeter(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegajoulesPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMegajoulesPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDensityMegajoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegajoulesPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegajoulesPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromMegajoulesPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromMegajoulesPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMegajoulesPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromMegajoulesPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMegajoulesPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegajoulesPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromMegajoulesPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDensityMegajoulePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegajoulesPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGigajoulesPerCubicMeter function
func TestEnergyDensityFactory_FromGigajoulesPerCubicMeter(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigajoulesPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromGigajoulesPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDensityGigajoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigajoulesPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigajoulesPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromGigajoulesPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromGigajoulesPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromGigajoulesPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromGigajoulesPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGigajoulesPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigajoulesPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromGigajoulesPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDensityGigajoulePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigajoulesPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTerajoulesPerCubicMeter function
func TestEnergyDensityFactory_FromTerajoulesPerCubicMeter(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerajoulesPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromTerajoulesPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDensityTerajoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerajoulesPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerajoulesPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromTerajoulesPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromTerajoulesPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromTerajoulesPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromTerajoulesPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTerajoulesPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerajoulesPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromTerajoulesPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDensityTerajoulePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerajoulesPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPetajoulesPerCubicMeter function
func TestEnergyDensityFactory_FromPetajoulesPerCubicMeter(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPetajoulesPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromPetajoulesPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDensityPetajoulePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPetajoulesPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPetajoulesPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromPetajoulesPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromPetajoulesPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromPetajoulesPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromPetajoulesPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromPetajoulesPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPetajoulesPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromPetajoulesPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDensityPetajoulePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPetajoulesPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattHoursPerCubicMeter function
func TestEnergyDensityFactory_FromKilowattHoursPerCubicMeter(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattHoursPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromKilowattHoursPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDensityKilowattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattHoursPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattHoursPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromKilowattHoursPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromKilowattHoursPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattHoursPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilowattHoursPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattHoursPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattHoursPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromKilowattHoursPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDensityKilowattHourPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattHoursPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattHoursPerCubicMeter function
func TestEnergyDensityFactory_FromMegawattHoursPerCubicMeter(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattHoursPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMegawattHoursPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDensityMegawattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattHoursPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattHoursPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromMegawattHoursPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromMegawattHoursPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattHoursPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromMegawattHoursPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattHoursPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattHoursPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromMegawattHoursPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDensityMegawattHourPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattHoursPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattHoursPerCubicMeter function
func TestEnergyDensityFactory_FromGigawattHoursPerCubicMeter(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattHoursPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromGigawattHoursPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDensityGigawattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattHoursPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattHoursPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromGigawattHoursPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromGigawattHoursPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattHoursPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromGigawattHoursPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattHoursPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattHoursPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromGigawattHoursPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDensityGigawattHourPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattHoursPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTerawattHoursPerCubicMeter function
func TestEnergyDensityFactory_FromTerawattHoursPerCubicMeter(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerawattHoursPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromTerawattHoursPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDensityTerawattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerawattHoursPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerawattHoursPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromTerawattHoursPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromTerawattHoursPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromTerawattHoursPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromTerawattHoursPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTerawattHoursPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerawattHoursPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromTerawattHoursPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDensityTerawattHourPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerawattHoursPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPetawattHoursPerCubicMeter function
func TestEnergyDensityFactory_FromPetawattHoursPerCubicMeter(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPetawattHoursPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromPetawattHoursPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDensityPetawattHourPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPetawattHoursPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPetawattHoursPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromPetawattHoursPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromPetawattHoursPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromPetawattHoursPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromPetawattHoursPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromPetawattHoursPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPetawattHoursPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromPetawattHoursPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDensityPetawattHourPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPetawattHoursPerCubicMeter() with zero value = %v, want 0", converted)
    }
}

func TestEnergyDensityToString(t *testing.T) {
	factory := units.EnergyDensityFactory{}
	a, err := factory.CreateEnergyDensity(45, units.EnergyDensityJoulePerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.EnergyDensityJoulePerCubicMeter, 2)
	expected := "45.00 " + units.GetEnergyDensityAbbreviation(units.EnergyDensityJoulePerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.EnergyDensityJoulePerCubicMeter, -1)
	expected = "45 " + units.GetEnergyDensityAbbreviation(units.EnergyDensityJoulePerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestEnergyDensity_EqualityAndComparison(t *testing.T) {
	factory := units.EnergyDensityFactory{}
	a1, _ := factory.CreateEnergyDensity(60, units.EnergyDensityJoulePerCubicMeter)
	a2, _ := factory.CreateEnergyDensity(60, units.EnergyDensityJoulePerCubicMeter)
	a3, _ := factory.CreateEnergyDensity(120, units.EnergyDensityJoulePerCubicMeter)

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

func TestEnergyDensity_Arithmetic(t *testing.T) {
	factory := units.EnergyDensityFactory{}
	a1, _ := factory.CreateEnergyDensity(30, units.EnergyDensityJoulePerCubicMeter)
	a2, _ := factory.CreateEnergyDensity(45, units.EnergyDensityJoulePerCubicMeter)

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


func TestGetEnergyDensityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.EnergyDensityUnits
        want string
    }{
        {
            name: "JoulePerCubicMeter abbreviation",
            unit: units.EnergyDensityJoulePerCubicMeter,
            want: "J/m³",
        },
        {
            name: "WattHourPerCubicMeter abbreviation",
            unit: units.EnergyDensityWattHourPerCubicMeter,
            want: "Wh/m³",
        },
        {
            name: "KilojoulePerCubicMeter abbreviation",
            unit: units.EnergyDensityKilojoulePerCubicMeter,
            want: "kJ/m³",
        },
        {
            name: "MegajoulePerCubicMeter abbreviation",
            unit: units.EnergyDensityMegajoulePerCubicMeter,
            want: "MJ/m³",
        },
        {
            name: "GigajoulePerCubicMeter abbreviation",
            unit: units.EnergyDensityGigajoulePerCubicMeter,
            want: "GJ/m³",
        },
        {
            name: "TerajoulePerCubicMeter abbreviation",
            unit: units.EnergyDensityTerajoulePerCubicMeter,
            want: "TJ/m³",
        },
        {
            name: "PetajoulePerCubicMeter abbreviation",
            unit: units.EnergyDensityPetajoulePerCubicMeter,
            want: "PJ/m³",
        },
        {
            name: "KilowattHourPerCubicMeter abbreviation",
            unit: units.EnergyDensityKilowattHourPerCubicMeter,
            want: "kWh/m³",
        },
        {
            name: "MegawattHourPerCubicMeter abbreviation",
            unit: units.EnergyDensityMegawattHourPerCubicMeter,
            want: "MWh/m³",
        },
        {
            name: "GigawattHourPerCubicMeter abbreviation",
            unit: units.EnergyDensityGigawattHourPerCubicMeter,
            want: "GWh/m³",
        },
        {
            name: "TerawattHourPerCubicMeter abbreviation",
            unit: units.EnergyDensityTerawattHourPerCubicMeter,
            want: "TWh/m³",
        },
        {
            name: "PetawattHourPerCubicMeter abbreviation",
            unit: units.EnergyDensityPetawattHourPerCubicMeter,
            want: "PWh/m³",
        },
        {
            name: "invalid unit",
            unit: units.EnergyDensityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetEnergyDensityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetEnergyDensityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestEnergyDensity_String(t *testing.T) {
    factory := units.EnergyDensityFactory{}
    
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
            unit, err := factory.CreateEnergyDensity(tt.value, units.EnergyDensityJoulePerCubicMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("EnergyDensity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}