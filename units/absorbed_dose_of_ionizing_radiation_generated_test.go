// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAbsorbedDoseOfIonizingRadiationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Gray"}`
	
	factory := units.AbsorbedDoseOfIonizingRadiationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AbsorbedDoseOfIonizingRadiationGray {
		t.Errorf("expected unit %v, got %v", units.AbsorbedDoseOfIonizingRadiationGray, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Gray"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAbsorbedDoseOfIonizingRadiationDto_ToJSON(t *testing.T) {
	dto := units.AbsorbedDoseOfIonizingRadiationDto{
		Value: 45,
		Unit:  units.AbsorbedDoseOfIonizingRadiationGray,
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
	if result["unit"].(string) != string(units.AbsorbedDoseOfIonizingRadiationGray) {
		t.Errorf("expected unit %s, got %v", units.AbsorbedDoseOfIonizingRadiationGray, result["unit"])
	}
}

func TestNewAbsorbedDoseOfIonizingRadiation_InvalidValue(t *testing.T) {
	factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAbsorbedDoseOfIonizingRadiation(math.NaN(), units.AbsorbedDoseOfIonizingRadiationGray)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAbsorbedDoseOfIonizingRadiation(math.Inf(1), units.AbsorbedDoseOfIonizingRadiationGray)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAbsorbedDoseOfIonizingRadiationConversions(t *testing.T) {
	factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAbsorbedDoseOfIonizingRadiation(180, units.AbsorbedDoseOfIonizingRadiationGray)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Grays.
		// No expected conversion value provided for Grays, verifying result is not NaN.
		result := a.Grays()
		cacheResult := a.Grays()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Grays returned NaN")
		}
	}
	{
		// Test conversion to Rads.
		// No expected conversion value provided for Rads, verifying result is not NaN.
		result := a.Rads()
		cacheResult := a.Rads()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Rads returned NaN")
		}
	}
	{
		// Test conversion to Femtograys.
		// No expected conversion value provided for Femtograys, verifying result is not NaN.
		result := a.Femtograys()
		cacheResult := a.Femtograys()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Femtograys returned NaN")
		}
	}
	{
		// Test conversion to Picograys.
		// No expected conversion value provided for Picograys, verifying result is not NaN.
		result := a.Picograys()
		cacheResult := a.Picograys()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Picograys returned NaN")
		}
	}
	{
		// Test conversion to Nanograys.
		// No expected conversion value provided for Nanograys, verifying result is not NaN.
		result := a.Nanograys()
		cacheResult := a.Nanograys()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanograys returned NaN")
		}
	}
	{
		// Test conversion to Micrograys.
		// No expected conversion value provided for Micrograys, verifying result is not NaN.
		result := a.Micrograys()
		cacheResult := a.Micrograys()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Micrograys returned NaN")
		}
	}
	{
		// Test conversion to Milligrays.
		// No expected conversion value provided for Milligrays, verifying result is not NaN.
		result := a.Milligrays()
		cacheResult := a.Milligrays()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Milligrays returned NaN")
		}
	}
	{
		// Test conversion to Centigrays.
		// No expected conversion value provided for Centigrays, verifying result is not NaN.
		result := a.Centigrays()
		cacheResult := a.Centigrays()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Centigrays returned NaN")
		}
	}
	{
		// Test conversion to Kilograys.
		// No expected conversion value provided for Kilograys, verifying result is not NaN.
		result := a.Kilograys()
		cacheResult := a.Kilograys()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilograys returned NaN")
		}
	}
	{
		// Test conversion to Megagrays.
		// No expected conversion value provided for Megagrays, verifying result is not NaN.
		result := a.Megagrays()
		cacheResult := a.Megagrays()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megagrays returned NaN")
		}
	}
	{
		// Test conversion to Gigagrays.
		// No expected conversion value provided for Gigagrays, verifying result is not NaN.
		result := a.Gigagrays()
		cacheResult := a.Gigagrays()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gigagrays returned NaN")
		}
	}
	{
		// Test conversion to Teragrays.
		// No expected conversion value provided for Teragrays, verifying result is not NaN.
		result := a.Teragrays()
		cacheResult := a.Teragrays()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Teragrays returned NaN")
		}
	}
	{
		// Test conversion to Petagrays.
		// No expected conversion value provided for Petagrays, verifying result is not NaN.
		result := a.Petagrays()
		cacheResult := a.Petagrays()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Petagrays returned NaN")
		}
	}
	{
		// Test conversion to Millirads.
		// No expected conversion value provided for Millirads, verifying result is not NaN.
		result := a.Millirads()
		cacheResult := a.Millirads()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millirads returned NaN")
		}
	}
	{
		// Test conversion to Kilorads.
		// No expected conversion value provided for Kilorads, verifying result is not NaN.
		result := a.Kilorads()
		cacheResult := a.Kilorads()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilorads returned NaN")
		}
	}
	{
		// Test conversion to Megarads.
		// No expected conversion value provided for Megarads, verifying result is not NaN.
		result := a.Megarads()
		cacheResult := a.Megarads()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megarads returned NaN")
		}
	}
}

func TestAbsorbedDoseOfIonizingRadiation_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
	a, err := factory.CreateAbsorbedDoseOfIonizingRadiation(90, units.AbsorbedDoseOfIonizingRadiationGray)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AbsorbedDoseOfIonizingRadiationGray {
		t.Errorf("expected default unit Gray, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AbsorbedDoseOfIonizingRadiationGray
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AbsorbedDoseOfIonizingRadiationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AbsorbedDoseOfIonizingRadiationGray {
		t.Errorf("expected unit Gray, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAbsorbedDoseOfIonizingRadiationFactory_FromDto(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationGray,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: math.NaN(),
        Unit:  units.AbsorbedDoseOfIonizingRadiationGray,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Gray conversion
    graysDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationGray,
    }
    
    var graysResult *units.AbsorbedDoseOfIonizingRadiation
    graysResult, err = factory.FromDto(graysDto)
    if err != nil {
        t.Errorf("FromDto() with Gray returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = graysResult.Convert(units.AbsorbedDoseOfIonizingRadiationGray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gray = %v, want %v", converted, 100)
    }
    // Test Rad conversion
    radsDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationRad,
    }
    
    var radsResult *units.AbsorbedDoseOfIonizingRadiation
    radsResult, err = factory.FromDto(radsDto)
    if err != nil {
        t.Errorf("FromDto() with Rad returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = radsResult.Convert(units.AbsorbedDoseOfIonizingRadiationRad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Rad = %v, want %v", converted, 100)
    }
    // Test Femtogray conversion
    femtograysDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationFemtogray,
    }
    
    var femtograysResult *units.AbsorbedDoseOfIonizingRadiation
    femtograysResult, err = factory.FromDto(femtograysDto)
    if err != nil {
        t.Errorf("FromDto() with Femtogray returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtograysResult.Convert(units.AbsorbedDoseOfIonizingRadiationFemtogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtogray = %v, want %v", converted, 100)
    }
    // Test Picogray conversion
    picograysDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationPicogray,
    }
    
    var picograysResult *units.AbsorbedDoseOfIonizingRadiation
    picograysResult, err = factory.FromDto(picograysDto)
    if err != nil {
        t.Errorf("FromDto() with Picogray returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograysResult.Convert(units.AbsorbedDoseOfIonizingRadiationPicogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picogray = %v, want %v", converted, 100)
    }
    // Test Nanogray conversion
    nanograysDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationNanogray,
    }
    
    var nanograysResult *units.AbsorbedDoseOfIonizingRadiation
    nanograysResult, err = factory.FromDto(nanograysDto)
    if err != nil {
        t.Errorf("FromDto() with Nanogray returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograysResult.Convert(units.AbsorbedDoseOfIonizingRadiationNanogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanogray = %v, want %v", converted, 100)
    }
    // Test Microgray conversion
    micrograysDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationMicrogray,
    }
    
    var micrograysResult *units.AbsorbedDoseOfIonizingRadiation
    micrograysResult, err = factory.FromDto(micrograysDto)
    if err != nil {
        t.Errorf("FromDto() with Microgray returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograysResult.Convert(units.AbsorbedDoseOfIonizingRadiationMicrogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microgray = %v, want %v", converted, 100)
    }
    // Test Milligray conversion
    milligraysDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationMilligray,
    }
    
    var milligraysResult *units.AbsorbedDoseOfIonizingRadiation
    milligraysResult, err = factory.FromDto(milligraysDto)
    if err != nil {
        t.Errorf("FromDto() with Milligray returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligraysResult.Convert(units.AbsorbedDoseOfIonizingRadiationMilligray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milligray = %v, want %v", converted, 100)
    }
    // Test Centigray conversion
    centigraysDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationCentigray,
    }
    
    var centigraysResult *units.AbsorbedDoseOfIonizingRadiation
    centigraysResult, err = factory.FromDto(centigraysDto)
    if err != nil {
        t.Errorf("FromDto() with Centigray returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigraysResult.Convert(units.AbsorbedDoseOfIonizingRadiationCentigray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centigray = %v, want %v", converted, 100)
    }
    // Test Kilogray conversion
    kilograysDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationKilogray,
    }
    
    var kilograysResult *units.AbsorbedDoseOfIonizingRadiation
    kilograysResult, err = factory.FromDto(kilograysDto)
    if err != nil {
        t.Errorf("FromDto() with Kilogray returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograysResult.Convert(units.AbsorbedDoseOfIonizingRadiationKilogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilogray = %v, want %v", converted, 100)
    }
    // Test Megagray conversion
    megagraysDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationMegagray,
    }
    
    var megagraysResult *units.AbsorbedDoseOfIonizingRadiation
    megagraysResult, err = factory.FromDto(megagraysDto)
    if err != nil {
        t.Errorf("FromDto() with Megagray returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megagraysResult.Convert(units.AbsorbedDoseOfIonizingRadiationMegagray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megagray = %v, want %v", converted, 100)
    }
    // Test Gigagray conversion
    gigagraysDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationGigagray,
    }
    
    var gigagraysResult *units.AbsorbedDoseOfIonizingRadiation
    gigagraysResult, err = factory.FromDto(gigagraysDto)
    if err != nil {
        t.Errorf("FromDto() with Gigagray returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigagraysResult.Convert(units.AbsorbedDoseOfIonizingRadiationGigagray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigagray = %v, want %v", converted, 100)
    }
    // Test Teragray conversion
    teragraysDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationTeragray,
    }
    
    var teragraysResult *units.AbsorbedDoseOfIonizingRadiation
    teragraysResult, err = factory.FromDto(teragraysDto)
    if err != nil {
        t.Errorf("FromDto() with Teragray returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teragraysResult.Convert(units.AbsorbedDoseOfIonizingRadiationTeragray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Teragray = %v, want %v", converted, 100)
    }
    // Test Petagray conversion
    petagraysDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationPetagray,
    }
    
    var petagraysResult *units.AbsorbedDoseOfIonizingRadiation
    petagraysResult, err = factory.FromDto(petagraysDto)
    if err != nil {
        t.Errorf("FromDto() with Petagray returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petagraysResult.Convert(units.AbsorbedDoseOfIonizingRadiationPetagray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petagray = %v, want %v", converted, 100)
    }
    // Test Millirad conversion
    milliradsDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationMillirad,
    }
    
    var milliradsResult *units.AbsorbedDoseOfIonizingRadiation
    milliradsResult, err = factory.FromDto(milliradsDto)
    if err != nil {
        t.Errorf("FromDto() with Millirad returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliradsResult.Convert(units.AbsorbedDoseOfIonizingRadiationMillirad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millirad = %v, want %v", converted, 100)
    }
    // Test Kilorad conversion
    kiloradsDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationKilorad,
    }
    
    var kiloradsResult *units.AbsorbedDoseOfIonizingRadiation
    kiloradsResult, err = factory.FromDto(kiloradsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilorad returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloradsResult.Convert(units.AbsorbedDoseOfIonizingRadiationKilorad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilorad = %v, want %v", converted, 100)
    }
    // Test Megarad conversion
    megaradsDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 100,
        Unit:  units.AbsorbedDoseOfIonizingRadiationMegarad,
    }
    
    var megaradsResult *units.AbsorbedDoseOfIonizingRadiation
    megaradsResult, err = factory.FromDto(megaradsDto)
    if err != nil {
        t.Errorf("FromDto() with Megarad returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaradsResult.Convert(units.AbsorbedDoseOfIonizingRadiationMegarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megarad = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.AbsorbedDoseOfIonizingRadiationDto{
        Value: 0,
        Unit:  units.AbsorbedDoseOfIonizingRadiationGray,
    }
    
    var zeroResult *units.AbsorbedDoseOfIonizingRadiation
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestAbsorbedDoseOfIonizingRadiationFactory_FromDtoJSON(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Gray"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Gray"}`)
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
    nanJSON, _ := json.Marshal(units.AbsorbedDoseOfIonizingRadiationDto{
        Value: nanValue,
        Unit:  units.AbsorbedDoseOfIonizingRadiationGray,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Gray unit
    graysJSON := []byte(`{"value": 100, "unit": "Gray"}`)
    graysResult, err := factory.FromDtoJSON(graysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gray unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = graysResult.Convert(units.AbsorbedDoseOfIonizingRadiationGray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gray = %v, want %v", converted, 100)
    }
    // Test JSON with Rad unit
    radsJSON := []byte(`{"value": 100, "unit": "Rad"}`)
    radsResult, err := factory.FromDtoJSON(radsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Rad unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = radsResult.Convert(units.AbsorbedDoseOfIonizingRadiationRad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Rad = %v, want %v", converted, 100)
    }
    // Test JSON with Femtogray unit
    femtograysJSON := []byte(`{"value": 100, "unit": "Femtogray"}`)
    femtograysResult, err := factory.FromDtoJSON(femtograysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Femtogray unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtograysResult.Convert(units.AbsorbedDoseOfIonizingRadiationFemtogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtogray = %v, want %v", converted, 100)
    }
    // Test JSON with Picogray unit
    picograysJSON := []byte(`{"value": 100, "unit": "Picogray"}`)
    picograysResult, err := factory.FromDtoJSON(picograysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Picogray unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograysResult.Convert(units.AbsorbedDoseOfIonizingRadiationPicogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picogray = %v, want %v", converted, 100)
    }
    // Test JSON with Nanogray unit
    nanograysJSON := []byte(`{"value": 100, "unit": "Nanogray"}`)
    nanograysResult, err := factory.FromDtoJSON(nanograysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanogray unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograysResult.Convert(units.AbsorbedDoseOfIonizingRadiationNanogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanogray = %v, want %v", converted, 100)
    }
    // Test JSON with Microgray unit
    micrograysJSON := []byte(`{"value": 100, "unit": "Microgray"}`)
    micrograysResult, err := factory.FromDtoJSON(micrograysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microgray unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograysResult.Convert(units.AbsorbedDoseOfIonizingRadiationMicrogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microgray = %v, want %v", converted, 100)
    }
    // Test JSON with Milligray unit
    milligraysJSON := []byte(`{"value": 100, "unit": "Milligray"}`)
    milligraysResult, err := factory.FromDtoJSON(milligraysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Milligray unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligraysResult.Convert(units.AbsorbedDoseOfIonizingRadiationMilligray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milligray = %v, want %v", converted, 100)
    }
    // Test JSON with Centigray unit
    centigraysJSON := []byte(`{"value": 100, "unit": "Centigray"}`)
    centigraysResult, err := factory.FromDtoJSON(centigraysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Centigray unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigraysResult.Convert(units.AbsorbedDoseOfIonizingRadiationCentigray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centigray = %v, want %v", converted, 100)
    }
    // Test JSON with Kilogray unit
    kilograysJSON := []byte(`{"value": 100, "unit": "Kilogray"}`)
    kilograysResult, err := factory.FromDtoJSON(kilograysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilogray unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograysResult.Convert(units.AbsorbedDoseOfIonizingRadiationKilogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilogray = %v, want %v", converted, 100)
    }
    // Test JSON with Megagray unit
    megagraysJSON := []byte(`{"value": 100, "unit": "Megagray"}`)
    megagraysResult, err := factory.FromDtoJSON(megagraysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megagray unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megagraysResult.Convert(units.AbsorbedDoseOfIonizingRadiationMegagray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megagray = %v, want %v", converted, 100)
    }
    // Test JSON with Gigagray unit
    gigagraysJSON := []byte(`{"value": 100, "unit": "Gigagray"}`)
    gigagraysResult, err := factory.FromDtoJSON(gigagraysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigagray unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigagraysResult.Convert(units.AbsorbedDoseOfIonizingRadiationGigagray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigagray = %v, want %v", converted, 100)
    }
    // Test JSON with Teragray unit
    teragraysJSON := []byte(`{"value": 100, "unit": "Teragray"}`)
    teragraysResult, err := factory.FromDtoJSON(teragraysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Teragray unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teragraysResult.Convert(units.AbsorbedDoseOfIonizingRadiationTeragray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Teragray = %v, want %v", converted, 100)
    }
    // Test JSON with Petagray unit
    petagraysJSON := []byte(`{"value": 100, "unit": "Petagray"}`)
    petagraysResult, err := factory.FromDtoJSON(petagraysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Petagray unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petagraysResult.Convert(units.AbsorbedDoseOfIonizingRadiationPetagray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petagray = %v, want %v", converted, 100)
    }
    // Test JSON with Millirad unit
    milliradsJSON := []byte(`{"value": 100, "unit": "Millirad"}`)
    milliradsResult, err := factory.FromDtoJSON(milliradsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millirad unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliradsResult.Convert(units.AbsorbedDoseOfIonizingRadiationMillirad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millirad = %v, want %v", converted, 100)
    }
    // Test JSON with Kilorad unit
    kiloradsJSON := []byte(`{"value": 100, "unit": "Kilorad"}`)
    kiloradsResult, err := factory.FromDtoJSON(kiloradsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilorad unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloradsResult.Convert(units.AbsorbedDoseOfIonizingRadiationKilorad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilorad = %v, want %v", converted, 100)
    }
    // Test JSON with Megarad unit
    megaradsJSON := []byte(`{"value": 100, "unit": "Megarad"}`)
    megaradsResult, err := factory.FromDtoJSON(megaradsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megarad unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaradsResult.Convert(units.AbsorbedDoseOfIonizingRadiationMegarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megarad = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Gray"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromGrays function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromGrays(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGrays(100)
    if err != nil {
        t.Errorf("FromGrays() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationGray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGrays() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGrays(math.NaN())
    if err == nil {
        t.Error("FromGrays() with NaN value should return error")
    }

    _, err = factory.FromGrays(math.Inf(1))
    if err == nil {
        t.Error("FromGrays() with +Inf value should return error")
    }

    _, err = factory.FromGrays(math.Inf(-1))
    if err == nil {
        t.Error("FromGrays() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGrays(0)
    if err != nil {
        t.Errorf("FromGrays() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationGray)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGrays() with zero value = %v, want 0", converted)
    }
}
// Test FromRads function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromRads(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRads(100)
    if err != nil {
        t.Errorf("FromRads() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationRad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromRads() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromRads(math.NaN())
    if err == nil {
        t.Error("FromRads() with NaN value should return error")
    }

    _, err = factory.FromRads(math.Inf(1))
    if err == nil {
        t.Error("FromRads() with +Inf value should return error")
    }

    _, err = factory.FromRads(math.Inf(-1))
    if err == nil {
        t.Error("FromRads() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromRads(0)
    if err != nil {
        t.Errorf("FromRads() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationRad)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRads() with zero value = %v, want 0", converted)
    }
}
// Test FromFemtograys function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromFemtograys(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFemtograys(100)
    if err != nil {
        t.Errorf("FromFemtograys() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationFemtogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFemtograys() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFemtograys(math.NaN())
    if err == nil {
        t.Error("FromFemtograys() with NaN value should return error")
    }

    _, err = factory.FromFemtograys(math.Inf(1))
    if err == nil {
        t.Error("FromFemtograys() with +Inf value should return error")
    }

    _, err = factory.FromFemtograys(math.Inf(-1))
    if err == nil {
        t.Error("FromFemtograys() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFemtograys(0)
    if err != nil {
        t.Errorf("FromFemtograys() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationFemtogray)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFemtograys() with zero value = %v, want 0", converted)
    }
}
// Test FromPicograys function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromPicograys(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicograys(100)
    if err != nil {
        t.Errorf("FromPicograys() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationPicogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicograys() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicograys(math.NaN())
    if err == nil {
        t.Error("FromPicograys() with NaN value should return error")
    }

    _, err = factory.FromPicograys(math.Inf(1))
    if err == nil {
        t.Error("FromPicograys() with +Inf value should return error")
    }

    _, err = factory.FromPicograys(math.Inf(-1))
    if err == nil {
        t.Error("FromPicograys() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicograys(0)
    if err != nil {
        t.Errorf("FromPicograys() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationPicogray)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicograys() with zero value = %v, want 0", converted)
    }
}
// Test FromNanograys function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromNanograys(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanograys(100)
    if err != nil {
        t.Errorf("FromNanograys() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationNanogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanograys() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanograys(math.NaN())
    if err == nil {
        t.Error("FromNanograys() with NaN value should return error")
    }

    _, err = factory.FromNanograys(math.Inf(1))
    if err == nil {
        t.Error("FromNanograys() with +Inf value should return error")
    }

    _, err = factory.FromNanograys(math.Inf(-1))
    if err == nil {
        t.Error("FromNanograys() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanograys(0)
    if err != nil {
        t.Errorf("FromNanograys() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationNanogray)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanograys() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrograys function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromMicrograys(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrograys(100)
    if err != nil {
        t.Errorf("FromMicrograys() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationMicrogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrograys() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrograys(math.NaN())
    if err == nil {
        t.Error("FromMicrograys() with NaN value should return error")
    }

    _, err = factory.FromMicrograys(math.Inf(1))
    if err == nil {
        t.Error("FromMicrograys() with +Inf value should return error")
    }

    _, err = factory.FromMicrograys(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrograys() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrograys(0)
    if err != nil {
        t.Errorf("FromMicrograys() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationMicrogray)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrograys() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligrays function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromMilligrays(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligrays(100)
    if err != nil {
        t.Errorf("FromMilligrays() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationMilligray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligrays() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligrays(math.NaN())
    if err == nil {
        t.Error("FromMilligrays() with NaN value should return error")
    }

    _, err = factory.FromMilligrays(math.Inf(1))
    if err == nil {
        t.Error("FromMilligrays() with +Inf value should return error")
    }

    _, err = factory.FromMilligrays(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligrays() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligrays(0)
    if err != nil {
        t.Errorf("FromMilligrays() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationMilligray)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligrays() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigrays function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromCentigrays(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigrays(100)
    if err != nil {
        t.Errorf("FromCentigrays() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationCentigray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigrays() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigrays(math.NaN())
    if err == nil {
        t.Error("FromCentigrays() with NaN value should return error")
    }

    _, err = factory.FromCentigrays(math.Inf(1))
    if err == nil {
        t.Error("FromCentigrays() with +Inf value should return error")
    }

    _, err = factory.FromCentigrays(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigrays() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigrays(0)
    if err != nil {
        t.Errorf("FromCentigrays() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationCentigray)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigrays() with zero value = %v, want 0", converted)
    }
}
// Test FromKilograys function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromKilograys(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilograys(100)
    if err != nil {
        t.Errorf("FromKilograys() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationKilogray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilograys() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilograys(math.NaN())
    if err == nil {
        t.Error("FromKilograys() with NaN value should return error")
    }

    _, err = factory.FromKilograys(math.Inf(1))
    if err == nil {
        t.Error("FromKilograys() with +Inf value should return error")
    }

    _, err = factory.FromKilograys(math.Inf(-1))
    if err == nil {
        t.Error("FromKilograys() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilograys(0)
    if err != nil {
        t.Errorf("FromKilograys() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationKilogray)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilograys() with zero value = %v, want 0", converted)
    }
}
// Test FromMegagrays function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromMegagrays(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegagrays(100)
    if err != nil {
        t.Errorf("FromMegagrays() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationMegagray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegagrays() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegagrays(math.NaN())
    if err == nil {
        t.Error("FromMegagrays() with NaN value should return error")
    }

    _, err = factory.FromMegagrays(math.Inf(1))
    if err == nil {
        t.Error("FromMegagrays() with +Inf value should return error")
    }

    _, err = factory.FromMegagrays(math.Inf(-1))
    if err == nil {
        t.Error("FromMegagrays() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegagrays(0)
    if err != nil {
        t.Errorf("FromMegagrays() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationMegagray)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegagrays() with zero value = %v, want 0", converted)
    }
}
// Test FromGigagrays function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromGigagrays(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigagrays(100)
    if err != nil {
        t.Errorf("FromGigagrays() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationGigagray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigagrays() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigagrays(math.NaN())
    if err == nil {
        t.Error("FromGigagrays() with NaN value should return error")
    }

    _, err = factory.FromGigagrays(math.Inf(1))
    if err == nil {
        t.Error("FromGigagrays() with +Inf value should return error")
    }

    _, err = factory.FromGigagrays(math.Inf(-1))
    if err == nil {
        t.Error("FromGigagrays() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigagrays(0)
    if err != nil {
        t.Errorf("FromGigagrays() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationGigagray)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigagrays() with zero value = %v, want 0", converted)
    }
}
// Test FromTeragrays function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromTeragrays(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTeragrays(100)
    if err != nil {
        t.Errorf("FromTeragrays() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationTeragray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTeragrays() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTeragrays(math.NaN())
    if err == nil {
        t.Error("FromTeragrays() with NaN value should return error")
    }

    _, err = factory.FromTeragrays(math.Inf(1))
    if err == nil {
        t.Error("FromTeragrays() with +Inf value should return error")
    }

    _, err = factory.FromTeragrays(math.Inf(-1))
    if err == nil {
        t.Error("FromTeragrays() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTeragrays(0)
    if err != nil {
        t.Errorf("FromTeragrays() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationTeragray)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTeragrays() with zero value = %v, want 0", converted)
    }
}
// Test FromPetagrays function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromPetagrays(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPetagrays(100)
    if err != nil {
        t.Errorf("FromPetagrays() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationPetagray)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPetagrays() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPetagrays(math.NaN())
    if err == nil {
        t.Error("FromPetagrays() with NaN value should return error")
    }

    _, err = factory.FromPetagrays(math.Inf(1))
    if err == nil {
        t.Error("FromPetagrays() with +Inf value should return error")
    }

    _, err = factory.FromPetagrays(math.Inf(-1))
    if err == nil {
        t.Error("FromPetagrays() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPetagrays(0)
    if err != nil {
        t.Errorf("FromPetagrays() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationPetagray)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPetagrays() with zero value = %v, want 0", converted)
    }
}
// Test FromMillirads function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromMillirads(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillirads(100)
    if err != nil {
        t.Errorf("FromMillirads() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationMillirad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillirads() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillirads(math.NaN())
    if err == nil {
        t.Error("FromMillirads() with NaN value should return error")
    }

    _, err = factory.FromMillirads(math.Inf(1))
    if err == nil {
        t.Error("FromMillirads() with +Inf value should return error")
    }

    _, err = factory.FromMillirads(math.Inf(-1))
    if err == nil {
        t.Error("FromMillirads() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillirads(0)
    if err != nil {
        t.Errorf("FromMillirads() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationMillirad)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillirads() with zero value = %v, want 0", converted)
    }
}
// Test FromKilorads function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromKilorads(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilorads(100)
    if err != nil {
        t.Errorf("FromKilorads() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationKilorad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilorads() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilorads(math.NaN())
    if err == nil {
        t.Error("FromKilorads() with NaN value should return error")
    }

    _, err = factory.FromKilorads(math.Inf(1))
    if err == nil {
        t.Error("FromKilorads() with +Inf value should return error")
    }

    _, err = factory.FromKilorads(math.Inf(-1))
    if err == nil {
        t.Error("FromKilorads() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilorads(0)
    if err != nil {
        t.Errorf("FromKilorads() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationKilorad)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilorads() with zero value = %v, want 0", converted)
    }
}
// Test FromMegarads function
func TestAbsorbedDoseOfIonizingRadiationFactory_FromMegarads(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegarads(100)
    if err != nil {
        t.Errorf("FromMegarads() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AbsorbedDoseOfIonizingRadiationMegarad)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegarads() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegarads(math.NaN())
    if err == nil {
        t.Error("FromMegarads() with NaN value should return error")
    }

    _, err = factory.FromMegarads(math.Inf(1))
    if err == nil {
        t.Error("FromMegarads() with +Inf value should return error")
    }

    _, err = factory.FromMegarads(math.Inf(-1))
    if err == nil {
        t.Error("FromMegarads() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegarads(0)
    if err != nil {
        t.Errorf("FromMegarads() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AbsorbedDoseOfIonizingRadiationMegarad)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegarads() with zero value = %v, want 0", converted)
    }
}

func TestAbsorbedDoseOfIonizingRadiationToString(t *testing.T) {
	factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
	a, err := factory.CreateAbsorbedDoseOfIonizingRadiation(45, units.AbsorbedDoseOfIonizingRadiationGray)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AbsorbedDoseOfIonizingRadiationGray, 2)
	expected := "45.00 " + units.GetAbsorbedDoseOfIonizingRadiationAbbreviation(units.AbsorbedDoseOfIonizingRadiationGray)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AbsorbedDoseOfIonizingRadiationGray, -1)
	expected = "45 " + units.GetAbsorbedDoseOfIonizingRadiationAbbreviation(units.AbsorbedDoseOfIonizingRadiationGray)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAbsorbedDoseOfIonizingRadiation_EqualityAndComparison(t *testing.T) {
	factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
	a1, _ := factory.CreateAbsorbedDoseOfIonizingRadiation(60, units.AbsorbedDoseOfIonizingRadiationGray)
	a2, _ := factory.CreateAbsorbedDoseOfIonizingRadiation(60, units.AbsorbedDoseOfIonizingRadiationGray)
	a3, _ := factory.CreateAbsorbedDoseOfIonizingRadiation(120, units.AbsorbedDoseOfIonizingRadiationGray)

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

func TestAbsorbedDoseOfIonizingRadiation_Arithmetic(t *testing.T) {
	factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
	a1, _ := factory.CreateAbsorbedDoseOfIonizingRadiation(30, units.AbsorbedDoseOfIonizingRadiationGray)
	a2, _ := factory.CreateAbsorbedDoseOfIonizingRadiation(45, units.AbsorbedDoseOfIonizingRadiationGray)

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


func TestGetAbsorbedDoseOfIonizingRadiationAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.AbsorbedDoseOfIonizingRadiationUnits
        want string
    }{
        {
            name: "Gray abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationGray,
            want: "Gy",
        },
        {
            name: "Rad abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationRad,
            want: "rad",
        },
        {
            name: "Femtogray abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationFemtogray,
            want: "fGy",
        },
        {
            name: "Picogray abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationPicogray,
            want: "pGy",
        },
        {
            name: "Nanogray abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationNanogray,
            want: "nGy",
        },
        {
            name: "Microgray abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationMicrogray,
            want: "μGy",
        },
        {
            name: "Milligray abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationMilligray,
            want: "mGy",
        },
        {
            name: "Centigray abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationCentigray,
            want: "cGy",
        },
        {
            name: "Kilogray abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationKilogray,
            want: "kGy",
        },
        {
            name: "Megagray abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationMegagray,
            want: "MGy",
        },
        {
            name: "Gigagray abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationGigagray,
            want: "GGy",
        },
        {
            name: "Teragray abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationTeragray,
            want: "TGy",
        },
        {
            name: "Petagray abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationPetagray,
            want: "PGy",
        },
        {
            name: "Millirad abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationMillirad,
            want: "mrad",
        },
        {
            name: "Kilorad abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationKilorad,
            want: "krad",
        },
        {
            name: "Megarad abbreviation",
            unit: units.AbsorbedDoseOfIonizingRadiationMegarad,
            want: "Mrad",
        },
        {
            name: "invalid unit",
            unit: units.AbsorbedDoseOfIonizingRadiationUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetAbsorbedDoseOfIonizingRadiationAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetAbsorbedDoseOfIonizingRadiationAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestAbsorbedDoseOfIonizingRadiation_String(t *testing.T) {
    factory := units.AbsorbedDoseOfIonizingRadiationFactory{}
    
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
            unit, err := factory.CreateAbsorbedDoseOfIonizingRadiation(tt.value, units.AbsorbedDoseOfIonizingRadiationGray)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("AbsorbedDoseOfIonizingRadiation.String() = %v, want %v", got, tt.want)
            }
        })
    }
}