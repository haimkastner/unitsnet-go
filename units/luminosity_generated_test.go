// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLuminosityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Watt"}`
	
	factory := units.LuminosityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LuminosityWatt {
		t.Errorf("expected unit %v, got %v", units.LuminosityWatt, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Watt"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLuminosityDto_ToJSON(t *testing.T) {
	dto := units.LuminosityDto{
		Value: 45,
		Unit:  units.LuminosityWatt,
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
	if result["unit"].(string) != string(units.LuminosityWatt) {
		t.Errorf("expected unit %s, got %v", units.LuminosityWatt, result["unit"])
	}
}

func TestNewLuminosity_InvalidValue(t *testing.T) {
	factory := units.LuminosityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLuminosity(math.NaN(), units.LuminosityWatt)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLuminosity(math.Inf(1), units.LuminosityWatt)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLuminosityConversions(t *testing.T) {
	factory := units.LuminosityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLuminosity(180, units.LuminosityWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Watts.
		// No expected conversion value provided for Watts, verifying result is not NaN.
		result := a.Watts()
		cacheResult := a.Watts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Watts returned NaN")
		}
	}
	{
		// Test conversion to SolarLuminosities.
		// No expected conversion value provided for SolarLuminosities, verifying result is not NaN.
		result := a.SolarLuminosities()
		cacheResult := a.SolarLuminosities()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SolarLuminosities returned NaN")
		}
	}
	{
		// Test conversion to Femtowatts.
		// No expected conversion value provided for Femtowatts, verifying result is not NaN.
		result := a.Femtowatts()
		cacheResult := a.Femtowatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Femtowatts returned NaN")
		}
	}
	{
		// Test conversion to Picowatts.
		// No expected conversion value provided for Picowatts, verifying result is not NaN.
		result := a.Picowatts()
		cacheResult := a.Picowatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Picowatts returned NaN")
		}
	}
	{
		// Test conversion to Nanowatts.
		// No expected conversion value provided for Nanowatts, verifying result is not NaN.
		result := a.Nanowatts()
		cacheResult := a.Nanowatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanowatts returned NaN")
		}
	}
	{
		// Test conversion to Microwatts.
		// No expected conversion value provided for Microwatts, verifying result is not NaN.
		result := a.Microwatts()
		cacheResult := a.Microwatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microwatts returned NaN")
		}
	}
	{
		// Test conversion to Milliwatts.
		// No expected conversion value provided for Milliwatts, verifying result is not NaN.
		result := a.Milliwatts()
		cacheResult := a.Milliwatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Milliwatts returned NaN")
		}
	}
	{
		// Test conversion to Deciwatts.
		// No expected conversion value provided for Deciwatts, verifying result is not NaN.
		result := a.Deciwatts()
		cacheResult := a.Deciwatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Deciwatts returned NaN")
		}
	}
	{
		// Test conversion to Decawatts.
		// No expected conversion value provided for Decawatts, verifying result is not NaN.
		result := a.Decawatts()
		cacheResult := a.Decawatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Decawatts returned NaN")
		}
	}
	{
		// Test conversion to Kilowatts.
		// No expected conversion value provided for Kilowatts, verifying result is not NaN.
		result := a.Kilowatts()
		cacheResult := a.Kilowatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilowatts returned NaN")
		}
	}
	{
		// Test conversion to Megawatts.
		// No expected conversion value provided for Megawatts, verifying result is not NaN.
		result := a.Megawatts()
		cacheResult := a.Megawatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megawatts returned NaN")
		}
	}
	{
		// Test conversion to Gigawatts.
		// No expected conversion value provided for Gigawatts, verifying result is not NaN.
		result := a.Gigawatts()
		cacheResult := a.Gigawatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gigawatts returned NaN")
		}
	}
	{
		// Test conversion to Terawatts.
		// No expected conversion value provided for Terawatts, verifying result is not NaN.
		result := a.Terawatts()
		cacheResult := a.Terawatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Terawatts returned NaN")
		}
	}
	{
		// Test conversion to Petawatts.
		// No expected conversion value provided for Petawatts, verifying result is not NaN.
		result := a.Petawatts()
		cacheResult := a.Petawatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Petawatts returned NaN")
		}
	}
}

func TestLuminosity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LuminosityFactory{}
	a, err := factory.CreateLuminosity(90, units.LuminosityWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LuminosityWatt {
		t.Errorf("expected default unit Watt, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LuminosityWatt
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LuminosityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LuminosityWatt {
		t.Errorf("expected unit Watt, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLuminosityFactory_FromDto(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityWatt,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.LuminosityDto{
        Value: math.NaN(),
        Unit:  units.LuminosityWatt,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Watt conversion
    wattsDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityWatt,
    }
    
    var wattsResult *units.Luminosity
    wattsResult, err = factory.FromDto(wattsDto)
    if err != nil {
        t.Errorf("FromDto() with Watt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = wattsResult.Convert(units.LuminosityWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Watt = %v, want %v", converted, 100)
    }
    // Test SolarLuminosity conversion
    solar_luminositiesDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminositySolarLuminosity,
    }
    
    var solar_luminositiesResult *units.Luminosity
    solar_luminositiesResult, err = factory.FromDto(solar_luminositiesDto)
    if err != nil {
        t.Errorf("FromDto() with SolarLuminosity returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = solar_luminositiesResult.Convert(units.LuminositySolarLuminosity)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SolarLuminosity = %v, want %v", converted, 100)
    }
    // Test Femtowatt conversion
    femtowattsDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityFemtowatt,
    }
    
    var femtowattsResult *units.Luminosity
    femtowattsResult, err = factory.FromDto(femtowattsDto)
    if err != nil {
        t.Errorf("FromDto() with Femtowatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtowattsResult.Convert(units.LuminosityFemtowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtowatt = %v, want %v", converted, 100)
    }
    // Test Picowatt conversion
    picowattsDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityPicowatt,
    }
    
    var picowattsResult *units.Luminosity
    picowattsResult, err = factory.FromDto(picowattsDto)
    if err != nil {
        t.Errorf("FromDto() with Picowatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowattsResult.Convert(units.LuminosityPicowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picowatt = %v, want %v", converted, 100)
    }
    // Test Nanowatt conversion
    nanowattsDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityNanowatt,
    }
    
    var nanowattsResult *units.Luminosity
    nanowattsResult, err = factory.FromDto(nanowattsDto)
    if err != nil {
        t.Errorf("FromDto() with Nanowatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowattsResult.Convert(units.LuminosityNanowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanowatt = %v, want %v", converted, 100)
    }
    // Test Microwatt conversion
    microwattsDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityMicrowatt,
    }
    
    var microwattsResult *units.Luminosity
    microwattsResult, err = factory.FromDto(microwattsDto)
    if err != nil {
        t.Errorf("FromDto() with Microwatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwattsResult.Convert(units.LuminosityMicrowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microwatt = %v, want %v", converted, 100)
    }
    // Test Milliwatt conversion
    milliwattsDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityMilliwatt,
    }
    
    var milliwattsResult *units.Luminosity
    milliwattsResult, err = factory.FromDto(milliwattsDto)
    if err != nil {
        t.Errorf("FromDto() with Milliwatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwattsResult.Convert(units.LuminosityMilliwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliwatt = %v, want %v", converted, 100)
    }
    // Test Deciwatt conversion
    deciwattsDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityDeciwatt,
    }
    
    var deciwattsResult *units.Luminosity
    deciwattsResult, err = factory.FromDto(deciwattsDto)
    if err != nil {
        t.Errorf("FromDto() with Deciwatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwattsResult.Convert(units.LuminosityDeciwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Deciwatt = %v, want %v", converted, 100)
    }
    // Test Decawatt conversion
    decawattsDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityDecawatt,
    }
    
    var decawattsResult *units.Luminosity
    decawattsResult, err = factory.FromDto(decawattsDto)
    if err != nil {
        t.Errorf("FromDto() with Decawatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decawattsResult.Convert(units.LuminosityDecawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decawatt = %v, want %v", converted, 100)
    }
    // Test Kilowatt conversion
    kilowattsDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityKilowatt,
    }
    
    var kilowattsResult *units.Luminosity
    kilowattsResult, err = factory.FromDto(kilowattsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilowatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowattsResult.Convert(units.LuminosityKilowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilowatt = %v, want %v", converted, 100)
    }
    // Test Megawatt conversion
    megawattsDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityMegawatt,
    }
    
    var megawattsResult *units.Luminosity
    megawattsResult, err = factory.FromDto(megawattsDto)
    if err != nil {
        t.Errorf("FromDto() with Megawatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawattsResult.Convert(units.LuminosityMegawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megawatt = %v, want %v", converted, 100)
    }
    // Test Gigawatt conversion
    gigawattsDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityGigawatt,
    }
    
    var gigawattsResult *units.Luminosity
    gigawattsResult, err = factory.FromDto(gigawattsDto)
    if err != nil {
        t.Errorf("FromDto() with Gigawatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawattsResult.Convert(units.LuminosityGigawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigawatt = %v, want %v", converted, 100)
    }
    // Test Terawatt conversion
    terawattsDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityTerawatt,
    }
    
    var terawattsResult *units.Luminosity
    terawattsResult, err = factory.FromDto(terawattsDto)
    if err != nil {
        t.Errorf("FromDto() with Terawatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawattsResult.Convert(units.LuminosityTerawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terawatt = %v, want %v", converted, 100)
    }
    // Test Petawatt conversion
    petawattsDto := units.LuminosityDto{
        Value: 100,
        Unit:  units.LuminosityPetawatt,
    }
    
    var petawattsResult *units.Luminosity
    petawattsResult, err = factory.FromDto(petawattsDto)
    if err != nil {
        t.Errorf("FromDto() with Petawatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petawattsResult.Convert(units.LuminosityPetawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petawatt = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.LuminosityDto{
        Value: 0,
        Unit:  units.LuminosityWatt,
    }
    
    var zeroResult *units.Luminosity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestLuminosityFactory_FromDtoJSON(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Watt"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Watt"}`)
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
    nanJSON, _ := json.Marshal(units.LuminosityDto{
        Value: nanValue,
        Unit:  units.LuminosityWatt,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Watt unit
    wattsJSON := []byte(`{"value": 100, "unit": "Watt"}`)
    wattsResult, err := factory.FromDtoJSON(wattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Watt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = wattsResult.Convert(units.LuminosityWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Watt = %v, want %v", converted, 100)
    }
    // Test JSON with SolarLuminosity unit
    solar_luminositiesJSON := []byte(`{"value": 100, "unit": "SolarLuminosity"}`)
    solar_luminositiesResult, err := factory.FromDtoJSON(solar_luminositiesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SolarLuminosity unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = solar_luminositiesResult.Convert(units.LuminositySolarLuminosity)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SolarLuminosity = %v, want %v", converted, 100)
    }
    // Test JSON with Femtowatt unit
    femtowattsJSON := []byte(`{"value": 100, "unit": "Femtowatt"}`)
    femtowattsResult, err := factory.FromDtoJSON(femtowattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Femtowatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtowattsResult.Convert(units.LuminosityFemtowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtowatt = %v, want %v", converted, 100)
    }
    // Test JSON with Picowatt unit
    picowattsJSON := []byte(`{"value": 100, "unit": "Picowatt"}`)
    picowattsResult, err := factory.FromDtoJSON(picowattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Picowatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowattsResult.Convert(units.LuminosityPicowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picowatt = %v, want %v", converted, 100)
    }
    // Test JSON with Nanowatt unit
    nanowattsJSON := []byte(`{"value": 100, "unit": "Nanowatt"}`)
    nanowattsResult, err := factory.FromDtoJSON(nanowattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanowatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowattsResult.Convert(units.LuminosityNanowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanowatt = %v, want %v", converted, 100)
    }
    // Test JSON with Microwatt unit
    microwattsJSON := []byte(`{"value": 100, "unit": "Microwatt"}`)
    microwattsResult, err := factory.FromDtoJSON(microwattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microwatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwattsResult.Convert(units.LuminosityMicrowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microwatt = %v, want %v", converted, 100)
    }
    // Test JSON with Milliwatt unit
    milliwattsJSON := []byte(`{"value": 100, "unit": "Milliwatt"}`)
    milliwattsResult, err := factory.FromDtoJSON(milliwattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Milliwatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwattsResult.Convert(units.LuminosityMilliwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliwatt = %v, want %v", converted, 100)
    }
    // Test JSON with Deciwatt unit
    deciwattsJSON := []byte(`{"value": 100, "unit": "Deciwatt"}`)
    deciwattsResult, err := factory.FromDtoJSON(deciwattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Deciwatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwattsResult.Convert(units.LuminosityDeciwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Deciwatt = %v, want %v", converted, 100)
    }
    // Test JSON with Decawatt unit
    decawattsJSON := []byte(`{"value": 100, "unit": "Decawatt"}`)
    decawattsResult, err := factory.FromDtoJSON(decawattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Decawatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decawattsResult.Convert(units.LuminosityDecawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decawatt = %v, want %v", converted, 100)
    }
    // Test JSON with Kilowatt unit
    kilowattsJSON := []byte(`{"value": 100, "unit": "Kilowatt"}`)
    kilowattsResult, err := factory.FromDtoJSON(kilowattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilowatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowattsResult.Convert(units.LuminosityKilowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilowatt = %v, want %v", converted, 100)
    }
    // Test JSON with Megawatt unit
    megawattsJSON := []byte(`{"value": 100, "unit": "Megawatt"}`)
    megawattsResult, err := factory.FromDtoJSON(megawattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megawatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawattsResult.Convert(units.LuminosityMegawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megawatt = %v, want %v", converted, 100)
    }
    // Test JSON with Gigawatt unit
    gigawattsJSON := []byte(`{"value": 100, "unit": "Gigawatt"}`)
    gigawattsResult, err := factory.FromDtoJSON(gigawattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigawatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawattsResult.Convert(units.LuminosityGigawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigawatt = %v, want %v", converted, 100)
    }
    // Test JSON with Terawatt unit
    terawattsJSON := []byte(`{"value": 100, "unit": "Terawatt"}`)
    terawattsResult, err := factory.FromDtoJSON(terawattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Terawatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawattsResult.Convert(units.LuminosityTerawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terawatt = %v, want %v", converted, 100)
    }
    // Test JSON with Petawatt unit
    petawattsJSON := []byte(`{"value": 100, "unit": "Petawatt"}`)
    petawattsResult, err := factory.FromDtoJSON(petawattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Petawatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petawattsResult.Convert(units.LuminosityPetawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petawatt = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Watt"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromWatts function
func TestLuminosityFactory_FromWatts(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWatts(100)
    if err != nil {
        t.Errorf("FromWatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminosityWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWatts(math.NaN())
    if err == nil {
        t.Error("FromWatts() with NaN value should return error")
    }

    _, err = factory.FromWatts(math.Inf(1))
    if err == nil {
        t.Error("FromWatts() with +Inf value should return error")
    }

    _, err = factory.FromWatts(math.Inf(-1))
    if err == nil {
        t.Error("FromWatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWatts(0)
    if err != nil {
        t.Errorf("FromWatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminosityWatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWatts() with zero value = %v, want 0", converted)
    }
}
// Test FromSolarLuminosities function
func TestLuminosityFactory_FromSolarLuminosities(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSolarLuminosities(100)
    if err != nil {
        t.Errorf("FromSolarLuminosities() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminositySolarLuminosity)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSolarLuminosities() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSolarLuminosities(math.NaN())
    if err == nil {
        t.Error("FromSolarLuminosities() with NaN value should return error")
    }

    _, err = factory.FromSolarLuminosities(math.Inf(1))
    if err == nil {
        t.Error("FromSolarLuminosities() with +Inf value should return error")
    }

    _, err = factory.FromSolarLuminosities(math.Inf(-1))
    if err == nil {
        t.Error("FromSolarLuminosities() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSolarLuminosities(0)
    if err != nil {
        t.Errorf("FromSolarLuminosities() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminositySolarLuminosity)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSolarLuminosities() with zero value = %v, want 0", converted)
    }
}
// Test FromFemtowatts function
func TestLuminosityFactory_FromFemtowatts(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFemtowatts(100)
    if err != nil {
        t.Errorf("FromFemtowatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminosityFemtowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFemtowatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFemtowatts(math.NaN())
    if err == nil {
        t.Error("FromFemtowatts() with NaN value should return error")
    }

    _, err = factory.FromFemtowatts(math.Inf(1))
    if err == nil {
        t.Error("FromFemtowatts() with +Inf value should return error")
    }

    _, err = factory.FromFemtowatts(math.Inf(-1))
    if err == nil {
        t.Error("FromFemtowatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFemtowatts(0)
    if err != nil {
        t.Errorf("FromFemtowatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminosityFemtowatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFemtowatts() with zero value = %v, want 0", converted)
    }
}
// Test FromPicowatts function
func TestLuminosityFactory_FromPicowatts(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicowatts(100)
    if err != nil {
        t.Errorf("FromPicowatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminosityPicowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicowatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicowatts(math.NaN())
    if err == nil {
        t.Error("FromPicowatts() with NaN value should return error")
    }

    _, err = factory.FromPicowatts(math.Inf(1))
    if err == nil {
        t.Error("FromPicowatts() with +Inf value should return error")
    }

    _, err = factory.FromPicowatts(math.Inf(-1))
    if err == nil {
        t.Error("FromPicowatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicowatts(0)
    if err != nil {
        t.Errorf("FromPicowatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminosityPicowatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicowatts() with zero value = %v, want 0", converted)
    }
}
// Test FromNanowatts function
func TestLuminosityFactory_FromNanowatts(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanowatts(100)
    if err != nil {
        t.Errorf("FromNanowatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminosityNanowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanowatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanowatts(math.NaN())
    if err == nil {
        t.Error("FromNanowatts() with NaN value should return error")
    }

    _, err = factory.FromNanowatts(math.Inf(1))
    if err == nil {
        t.Error("FromNanowatts() with +Inf value should return error")
    }

    _, err = factory.FromNanowatts(math.Inf(-1))
    if err == nil {
        t.Error("FromNanowatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanowatts(0)
    if err != nil {
        t.Errorf("FromNanowatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminosityNanowatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanowatts() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrowatts function
func TestLuminosityFactory_FromMicrowatts(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrowatts(100)
    if err != nil {
        t.Errorf("FromMicrowatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminosityMicrowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrowatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrowatts(math.NaN())
    if err == nil {
        t.Error("FromMicrowatts() with NaN value should return error")
    }

    _, err = factory.FromMicrowatts(math.Inf(1))
    if err == nil {
        t.Error("FromMicrowatts() with +Inf value should return error")
    }

    _, err = factory.FromMicrowatts(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrowatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrowatts(0)
    if err != nil {
        t.Errorf("FromMicrowatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminosityMicrowatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrowatts() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwatts function
func TestLuminosityFactory_FromMilliwatts(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwatts(100)
    if err != nil {
        t.Errorf("FromMilliwatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminosityMilliwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwatts(math.NaN())
    if err == nil {
        t.Error("FromMilliwatts() with NaN value should return error")
    }

    _, err = factory.FromMilliwatts(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwatts() with +Inf value should return error")
    }

    _, err = factory.FromMilliwatts(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwatts(0)
    if err != nil {
        t.Errorf("FromMilliwatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminosityMilliwatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwatts() with zero value = %v, want 0", converted)
    }
}
// Test FromDeciwatts function
func TestLuminosityFactory_FromDeciwatts(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDeciwatts(100)
    if err != nil {
        t.Errorf("FromDeciwatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminosityDeciwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDeciwatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDeciwatts(math.NaN())
    if err == nil {
        t.Error("FromDeciwatts() with NaN value should return error")
    }

    _, err = factory.FromDeciwatts(math.Inf(1))
    if err == nil {
        t.Error("FromDeciwatts() with +Inf value should return error")
    }

    _, err = factory.FromDeciwatts(math.Inf(-1))
    if err == nil {
        t.Error("FromDeciwatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDeciwatts(0)
    if err != nil {
        t.Errorf("FromDeciwatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminosityDeciwatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDeciwatts() with zero value = %v, want 0", converted)
    }
}
// Test FromDecawatts function
func TestLuminosityFactory_FromDecawatts(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecawatts(100)
    if err != nil {
        t.Errorf("FromDecawatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminosityDecawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecawatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecawatts(math.NaN())
    if err == nil {
        t.Error("FromDecawatts() with NaN value should return error")
    }

    _, err = factory.FromDecawatts(math.Inf(1))
    if err == nil {
        t.Error("FromDecawatts() with +Inf value should return error")
    }

    _, err = factory.FromDecawatts(math.Inf(-1))
    if err == nil {
        t.Error("FromDecawatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecawatts(0)
    if err != nil {
        t.Errorf("FromDecawatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminosityDecawatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecawatts() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowatts function
func TestLuminosityFactory_FromKilowatts(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowatts(100)
    if err != nil {
        t.Errorf("FromKilowatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminosityKilowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowatts(math.NaN())
    if err == nil {
        t.Error("FromKilowatts() with NaN value should return error")
    }

    _, err = factory.FromKilowatts(math.Inf(1))
    if err == nil {
        t.Error("FromKilowatts() with +Inf value should return error")
    }

    _, err = factory.FromKilowatts(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowatts(0)
    if err != nil {
        t.Errorf("FromKilowatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminosityKilowatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowatts() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawatts function
func TestLuminosityFactory_FromMegawatts(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawatts(100)
    if err != nil {
        t.Errorf("FromMegawatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminosityMegawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawatts(math.NaN())
    if err == nil {
        t.Error("FromMegawatts() with NaN value should return error")
    }

    _, err = factory.FromMegawatts(math.Inf(1))
    if err == nil {
        t.Error("FromMegawatts() with +Inf value should return error")
    }

    _, err = factory.FromMegawatts(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawatts(0)
    if err != nil {
        t.Errorf("FromMegawatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminosityMegawatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawatts() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawatts function
func TestLuminosityFactory_FromGigawatts(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawatts(100)
    if err != nil {
        t.Errorf("FromGigawatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminosityGigawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawatts(math.NaN())
    if err == nil {
        t.Error("FromGigawatts() with NaN value should return error")
    }

    _, err = factory.FromGigawatts(math.Inf(1))
    if err == nil {
        t.Error("FromGigawatts() with +Inf value should return error")
    }

    _, err = factory.FromGigawatts(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawatts(0)
    if err != nil {
        t.Errorf("FromGigawatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminosityGigawatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawatts() with zero value = %v, want 0", converted)
    }
}
// Test FromTerawatts function
func TestLuminosityFactory_FromTerawatts(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerawatts(100)
    if err != nil {
        t.Errorf("FromTerawatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminosityTerawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerawatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerawatts(math.NaN())
    if err == nil {
        t.Error("FromTerawatts() with NaN value should return error")
    }

    _, err = factory.FromTerawatts(math.Inf(1))
    if err == nil {
        t.Error("FromTerawatts() with +Inf value should return error")
    }

    _, err = factory.FromTerawatts(math.Inf(-1))
    if err == nil {
        t.Error("FromTerawatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerawatts(0)
    if err != nil {
        t.Errorf("FromTerawatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminosityTerawatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerawatts() with zero value = %v, want 0", converted)
    }
}
// Test FromPetawatts function
func TestLuminosityFactory_FromPetawatts(t *testing.T) {
    factory := units.LuminosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPetawatts(100)
    if err != nil {
        t.Errorf("FromPetawatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminosityPetawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPetawatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPetawatts(math.NaN())
    if err == nil {
        t.Error("FromPetawatts() with NaN value should return error")
    }

    _, err = factory.FromPetawatts(math.Inf(1))
    if err == nil {
        t.Error("FromPetawatts() with +Inf value should return error")
    }

    _, err = factory.FromPetawatts(math.Inf(-1))
    if err == nil {
        t.Error("FromPetawatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPetawatts(0)
    if err != nil {
        t.Errorf("FromPetawatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminosityPetawatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPetawatts() with zero value = %v, want 0", converted)
    }
}

func TestLuminosityToString(t *testing.T) {
	factory := units.LuminosityFactory{}
	a, err := factory.CreateLuminosity(45, units.LuminosityWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LuminosityWatt, 2)
	expected := "45.00 " + units.GetLuminosityAbbreviation(units.LuminosityWatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LuminosityWatt, -1)
	expected = "45 " + units.GetLuminosityAbbreviation(units.LuminosityWatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLuminosity_EqualityAndComparison(t *testing.T) {
	factory := units.LuminosityFactory{}
	a1, _ := factory.CreateLuminosity(60, units.LuminosityWatt)
	a2, _ := factory.CreateLuminosity(60, units.LuminosityWatt)
	a3, _ := factory.CreateLuminosity(120, units.LuminosityWatt)

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

func TestLuminosity_Arithmetic(t *testing.T) {
	factory := units.LuminosityFactory{}
	a1, _ := factory.CreateLuminosity(30, units.LuminosityWatt)
	a2, _ := factory.CreateLuminosity(45, units.LuminosityWatt)

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


func TestGetLuminosityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.LuminosityUnits
        want string
    }{
        {
            name: "Watt abbreviation",
            unit: units.LuminosityWatt,
            want: "W",
        },
        {
            name: "SolarLuminosity abbreviation",
            unit: units.LuminositySolarLuminosity,
            want: "L⊙",
        },
        {
            name: "Femtowatt abbreviation",
            unit: units.LuminosityFemtowatt,
            want: "fW",
        },
        {
            name: "Picowatt abbreviation",
            unit: units.LuminosityPicowatt,
            want: "pW",
        },
        {
            name: "Nanowatt abbreviation",
            unit: units.LuminosityNanowatt,
            want: "nW",
        },
        {
            name: "Microwatt abbreviation",
            unit: units.LuminosityMicrowatt,
            want: "μW",
        },
        {
            name: "Milliwatt abbreviation",
            unit: units.LuminosityMilliwatt,
            want: "mW",
        },
        {
            name: "Deciwatt abbreviation",
            unit: units.LuminosityDeciwatt,
            want: "dW",
        },
        {
            name: "Decawatt abbreviation",
            unit: units.LuminosityDecawatt,
            want: "daW",
        },
        {
            name: "Kilowatt abbreviation",
            unit: units.LuminosityKilowatt,
            want: "kW",
        },
        {
            name: "Megawatt abbreviation",
            unit: units.LuminosityMegawatt,
            want: "MW",
        },
        {
            name: "Gigawatt abbreviation",
            unit: units.LuminosityGigawatt,
            want: "GW",
        },
        {
            name: "Terawatt abbreviation",
            unit: units.LuminosityTerawatt,
            want: "TW",
        },
        {
            name: "Petawatt abbreviation",
            unit: units.LuminosityPetawatt,
            want: "PW",
        },
        {
            name: "invalid unit",
            unit: units.LuminosityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetLuminosityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetLuminosityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestLuminosity_String(t *testing.T) {
    factory := units.LuminosityFactory{}
    
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
            unit, err := factory.CreateLuminosity(tt.value, units.LuminosityWatt)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Luminosity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestLuminosity_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.LuminosityFactory{}

	_, err := uf.CreateLuminosity(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}