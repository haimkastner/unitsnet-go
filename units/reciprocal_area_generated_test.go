// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestReciprocalAreaDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "InverseSquareMeter"}`
	
	factory := units.ReciprocalAreaDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ReciprocalAreaInverseSquareMeter {
		t.Errorf("expected unit %v, got %v", units.ReciprocalAreaInverseSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "InverseSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestReciprocalAreaDto_ToJSON(t *testing.T) {
	dto := units.ReciprocalAreaDto{
		Value: 45,
		Unit:  units.ReciprocalAreaInverseSquareMeter,
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
	if result["unit"].(string) != string(units.ReciprocalAreaInverseSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.ReciprocalAreaInverseSquareMeter, result["unit"])
	}
}

func TestNewReciprocalArea_InvalidValue(t *testing.T) {
	factory := units.ReciprocalAreaFactory{}
	// NaN value should return an error.
	_, err := factory.CreateReciprocalArea(math.NaN(), units.ReciprocalAreaInverseSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateReciprocalArea(math.Inf(1), units.ReciprocalAreaInverseSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestReciprocalAreaConversions(t *testing.T) {
	factory := units.ReciprocalAreaFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateReciprocalArea(180, units.ReciprocalAreaInverseSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to InverseSquareMeters.
		// No expected conversion value provided for InverseSquareMeters, verifying result is not NaN.
		result := a.InverseSquareMeters()
		cacheResult := a.InverseSquareMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseSquareMeters returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareKilometers.
		// No expected conversion value provided for InverseSquareKilometers, verifying result is not NaN.
		result := a.InverseSquareKilometers()
		cacheResult := a.InverseSquareKilometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseSquareKilometers returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareDecimeters.
		// No expected conversion value provided for InverseSquareDecimeters, verifying result is not NaN.
		result := a.InverseSquareDecimeters()
		cacheResult := a.InverseSquareDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseSquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareCentimeters.
		// No expected conversion value provided for InverseSquareCentimeters, verifying result is not NaN.
		result := a.InverseSquareCentimeters()
		cacheResult := a.InverseSquareCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseSquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareMillimeters.
		// No expected conversion value provided for InverseSquareMillimeters, verifying result is not NaN.
		result := a.InverseSquareMillimeters()
		cacheResult := a.InverseSquareMillimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseSquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareMicrometers.
		// No expected conversion value provided for InverseSquareMicrometers, verifying result is not NaN.
		result := a.InverseSquareMicrometers()
		cacheResult := a.InverseSquareMicrometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseSquareMicrometers returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareMiles.
		// No expected conversion value provided for InverseSquareMiles, verifying result is not NaN.
		result := a.InverseSquareMiles()
		cacheResult := a.InverseSquareMiles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseSquareMiles returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareYards.
		// No expected conversion value provided for InverseSquareYards, verifying result is not NaN.
		result := a.InverseSquareYards()
		cacheResult := a.InverseSquareYards()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseSquareYards returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareFeet.
		// No expected conversion value provided for InverseSquareFeet, verifying result is not NaN.
		result := a.InverseSquareFeet()
		cacheResult := a.InverseSquareFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseSquareFeet returned NaN")
		}
	}
	{
		// Test conversion to InverseUsSurveySquareFeet.
		// No expected conversion value provided for InverseUsSurveySquareFeet, verifying result is not NaN.
		result := a.InverseUsSurveySquareFeet()
		cacheResult := a.InverseUsSurveySquareFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseUsSurveySquareFeet returned NaN")
		}
	}
	{
		// Test conversion to InverseSquareInches.
		// No expected conversion value provided for InverseSquareInches, verifying result is not NaN.
		result := a.InverseSquareInches()
		cacheResult := a.InverseSquareInches()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseSquareInches returned NaN")
		}
	}
}

func TestReciprocalArea_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ReciprocalAreaFactory{}
	a, err := factory.CreateReciprocalArea(90, units.ReciprocalAreaInverseSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ReciprocalAreaInverseSquareMeter {
		t.Errorf("expected default unit InverseSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ReciprocalAreaInverseSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ReciprocalAreaDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ReciprocalAreaInverseSquareMeter {
		t.Errorf("expected unit InverseSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestReciprocalAreaFactory_FromDto(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ReciprocalAreaDto{
        Value: 100,
        Unit:  units.ReciprocalAreaInverseSquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ReciprocalAreaDto{
        Value: math.NaN(),
        Unit:  units.ReciprocalAreaInverseSquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test InverseSquareMeter conversion
    inverse_square_metersDto := units.ReciprocalAreaDto{
        Value: 100,
        Unit:  units.ReciprocalAreaInverseSquareMeter,
    }
    
    var inverse_square_metersResult *units.ReciprocalArea
    inverse_square_metersResult, err = factory.FromDto(inverse_square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with InverseSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_metersResult.Convert(units.ReciprocalAreaInverseSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareMeter = %v, want %v", converted, 100)
    }
    // Test InverseSquareKilometer conversion
    inverse_square_kilometersDto := units.ReciprocalAreaDto{
        Value: 100,
        Unit:  units.ReciprocalAreaInverseSquareKilometer,
    }
    
    var inverse_square_kilometersResult *units.ReciprocalArea
    inverse_square_kilometersResult, err = factory.FromDto(inverse_square_kilometersDto)
    if err != nil {
        t.Errorf("FromDto() with InverseSquareKilometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_kilometersResult.Convert(units.ReciprocalAreaInverseSquareKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareKilometer = %v, want %v", converted, 100)
    }
    // Test InverseSquareDecimeter conversion
    inverse_square_decimetersDto := units.ReciprocalAreaDto{
        Value: 100,
        Unit:  units.ReciprocalAreaInverseSquareDecimeter,
    }
    
    var inverse_square_decimetersResult *units.ReciprocalArea
    inverse_square_decimetersResult, err = factory.FromDto(inverse_square_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with InverseSquareDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_decimetersResult.Convert(units.ReciprocalAreaInverseSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test InverseSquareCentimeter conversion
    inverse_square_centimetersDto := units.ReciprocalAreaDto{
        Value: 100,
        Unit:  units.ReciprocalAreaInverseSquareCentimeter,
    }
    
    var inverse_square_centimetersResult *units.ReciprocalArea
    inverse_square_centimetersResult, err = factory.FromDto(inverse_square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with InverseSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_centimetersResult.Convert(units.ReciprocalAreaInverseSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test InverseSquareMillimeter conversion
    inverse_square_millimetersDto := units.ReciprocalAreaDto{
        Value: 100,
        Unit:  units.ReciprocalAreaInverseSquareMillimeter,
    }
    
    var inverse_square_millimetersResult *units.ReciprocalArea
    inverse_square_millimetersResult, err = factory.FromDto(inverse_square_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with InverseSquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_millimetersResult.Convert(units.ReciprocalAreaInverseSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test InverseSquareMicrometer conversion
    inverse_square_micrometersDto := units.ReciprocalAreaDto{
        Value: 100,
        Unit:  units.ReciprocalAreaInverseSquareMicrometer,
    }
    
    var inverse_square_micrometersResult *units.ReciprocalArea
    inverse_square_micrometersResult, err = factory.FromDto(inverse_square_micrometersDto)
    if err != nil {
        t.Errorf("FromDto() with InverseSquareMicrometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_micrometersResult.Convert(units.ReciprocalAreaInverseSquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareMicrometer = %v, want %v", converted, 100)
    }
    // Test InverseSquareMile conversion
    inverse_square_milesDto := units.ReciprocalAreaDto{
        Value: 100,
        Unit:  units.ReciprocalAreaInverseSquareMile,
    }
    
    var inverse_square_milesResult *units.ReciprocalArea
    inverse_square_milesResult, err = factory.FromDto(inverse_square_milesDto)
    if err != nil {
        t.Errorf("FromDto() with InverseSquareMile returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_milesResult.Convert(units.ReciprocalAreaInverseSquareMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareMile = %v, want %v", converted, 100)
    }
    // Test InverseSquareYard conversion
    inverse_square_yardsDto := units.ReciprocalAreaDto{
        Value: 100,
        Unit:  units.ReciprocalAreaInverseSquareYard,
    }
    
    var inverse_square_yardsResult *units.ReciprocalArea
    inverse_square_yardsResult, err = factory.FromDto(inverse_square_yardsDto)
    if err != nil {
        t.Errorf("FromDto() with InverseSquareYard returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_yardsResult.Convert(units.ReciprocalAreaInverseSquareYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareYard = %v, want %v", converted, 100)
    }
    // Test InverseSquareFoot conversion
    inverse_square_feetDto := units.ReciprocalAreaDto{
        Value: 100,
        Unit:  units.ReciprocalAreaInverseSquareFoot,
    }
    
    var inverse_square_feetResult *units.ReciprocalArea
    inverse_square_feetResult, err = factory.FromDto(inverse_square_feetDto)
    if err != nil {
        t.Errorf("FromDto() with InverseSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_feetResult.Convert(units.ReciprocalAreaInverseSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareFoot = %v, want %v", converted, 100)
    }
    // Test InverseUsSurveySquareFoot conversion
    inverse_us_survey_square_feetDto := units.ReciprocalAreaDto{
        Value: 100,
        Unit:  units.ReciprocalAreaInverseUsSurveySquareFoot,
    }
    
    var inverse_us_survey_square_feetResult *units.ReciprocalArea
    inverse_us_survey_square_feetResult, err = factory.FromDto(inverse_us_survey_square_feetDto)
    if err != nil {
        t.Errorf("FromDto() with InverseUsSurveySquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_us_survey_square_feetResult.Convert(units.ReciprocalAreaInverseUsSurveySquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseUsSurveySquareFoot = %v, want %v", converted, 100)
    }
    // Test InverseSquareInch conversion
    inverse_square_inchesDto := units.ReciprocalAreaDto{
        Value: 100,
        Unit:  units.ReciprocalAreaInverseSquareInch,
    }
    
    var inverse_square_inchesResult *units.ReciprocalArea
    inverse_square_inchesResult, err = factory.FromDto(inverse_square_inchesDto)
    if err != nil {
        t.Errorf("FromDto() with InverseSquareInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_inchesResult.Convert(units.ReciprocalAreaInverseSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareInch = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ReciprocalAreaDto{
        Value: 0,
        Unit:  units.ReciprocalAreaInverseSquareMeter,
    }
    
    var zeroResult *units.ReciprocalArea
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestReciprocalAreaFactory_FromDtoJSON(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "InverseSquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "InverseSquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.ReciprocalAreaDto{
        Value: nanValue,
        Unit:  units.ReciprocalAreaInverseSquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with InverseSquareMeter unit
    inverse_square_metersJSON := []byte(`{"value": 100, "unit": "InverseSquareMeter"}`)
    inverse_square_metersResult, err := factory.FromDtoJSON(inverse_square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_metersResult.Convert(units.ReciprocalAreaInverseSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with InverseSquareKilometer unit
    inverse_square_kilometersJSON := []byte(`{"value": 100, "unit": "InverseSquareKilometer"}`)
    inverse_square_kilometersResult, err := factory.FromDtoJSON(inverse_square_kilometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseSquareKilometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_kilometersResult.Convert(units.ReciprocalAreaInverseSquareKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareKilometer = %v, want %v", converted, 100)
    }
    // Test JSON with InverseSquareDecimeter unit
    inverse_square_decimetersJSON := []byte(`{"value": 100, "unit": "InverseSquareDecimeter"}`)
    inverse_square_decimetersResult, err := factory.FromDtoJSON(inverse_square_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseSquareDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_decimetersResult.Convert(units.ReciprocalAreaInverseSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with InverseSquareCentimeter unit
    inverse_square_centimetersJSON := []byte(`{"value": 100, "unit": "InverseSquareCentimeter"}`)
    inverse_square_centimetersResult, err := factory.FromDtoJSON(inverse_square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_centimetersResult.Convert(units.ReciprocalAreaInverseSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with InverseSquareMillimeter unit
    inverse_square_millimetersJSON := []byte(`{"value": 100, "unit": "InverseSquareMillimeter"}`)
    inverse_square_millimetersResult, err := factory.FromDtoJSON(inverse_square_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseSquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_millimetersResult.Convert(units.ReciprocalAreaInverseSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with InverseSquareMicrometer unit
    inverse_square_micrometersJSON := []byte(`{"value": 100, "unit": "InverseSquareMicrometer"}`)
    inverse_square_micrometersResult, err := factory.FromDtoJSON(inverse_square_micrometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseSquareMicrometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_micrometersResult.Convert(units.ReciprocalAreaInverseSquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareMicrometer = %v, want %v", converted, 100)
    }
    // Test JSON with InverseSquareMile unit
    inverse_square_milesJSON := []byte(`{"value": 100, "unit": "InverseSquareMile"}`)
    inverse_square_milesResult, err := factory.FromDtoJSON(inverse_square_milesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseSquareMile unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_milesResult.Convert(units.ReciprocalAreaInverseSquareMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareMile = %v, want %v", converted, 100)
    }
    // Test JSON with InverseSquareYard unit
    inverse_square_yardsJSON := []byte(`{"value": 100, "unit": "InverseSquareYard"}`)
    inverse_square_yardsResult, err := factory.FromDtoJSON(inverse_square_yardsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseSquareYard unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_yardsResult.Convert(units.ReciprocalAreaInverseSquareYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareYard = %v, want %v", converted, 100)
    }
    // Test JSON with InverseSquareFoot unit
    inverse_square_feetJSON := []byte(`{"value": 100, "unit": "InverseSquareFoot"}`)
    inverse_square_feetResult, err := factory.FromDtoJSON(inverse_square_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_feetResult.Convert(units.ReciprocalAreaInverseSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with InverseUsSurveySquareFoot unit
    inverse_us_survey_square_feetJSON := []byte(`{"value": 100, "unit": "InverseUsSurveySquareFoot"}`)
    inverse_us_survey_square_feetResult, err := factory.FromDtoJSON(inverse_us_survey_square_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseUsSurveySquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_us_survey_square_feetResult.Convert(units.ReciprocalAreaInverseUsSurveySquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseUsSurveySquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with InverseSquareInch unit
    inverse_square_inchesJSON := []byte(`{"value": 100, "unit": "InverseSquareInch"}`)
    inverse_square_inchesResult, err := factory.FromDtoJSON(inverse_square_inchesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseSquareInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_square_inchesResult.Convert(units.ReciprocalAreaInverseSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseSquareInch = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "InverseSquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromInverseSquareMeters function
func TestReciprocalAreaFactory_FromInverseSquareMeters(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseSquareMeters(100)
    if err != nil {
        t.Errorf("FromInverseSquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalAreaInverseSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseSquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseSquareMeters(math.NaN())
    if err == nil {
        t.Error("FromInverseSquareMeters() with NaN value should return error")
    }

    _, err = factory.FromInverseSquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromInverseSquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromInverseSquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseSquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseSquareMeters(0)
    if err != nil {
        t.Errorf("FromInverseSquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalAreaInverseSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseSquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseSquareKilometers function
func TestReciprocalAreaFactory_FromInverseSquareKilometers(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseSquareKilometers(100)
    if err != nil {
        t.Errorf("FromInverseSquareKilometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalAreaInverseSquareKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseSquareKilometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseSquareKilometers(math.NaN())
    if err == nil {
        t.Error("FromInverseSquareKilometers() with NaN value should return error")
    }

    _, err = factory.FromInverseSquareKilometers(math.Inf(1))
    if err == nil {
        t.Error("FromInverseSquareKilometers() with +Inf value should return error")
    }

    _, err = factory.FromInverseSquareKilometers(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseSquareKilometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseSquareKilometers(0)
    if err != nil {
        t.Errorf("FromInverseSquareKilometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalAreaInverseSquareKilometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseSquareKilometers() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseSquareDecimeters function
func TestReciprocalAreaFactory_FromInverseSquareDecimeters(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseSquareDecimeters(100)
    if err != nil {
        t.Errorf("FromInverseSquareDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalAreaInverseSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseSquareDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseSquareDecimeters(math.NaN())
    if err == nil {
        t.Error("FromInverseSquareDecimeters() with NaN value should return error")
    }

    _, err = factory.FromInverseSquareDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromInverseSquareDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromInverseSquareDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseSquareDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseSquareDecimeters(0)
    if err != nil {
        t.Errorf("FromInverseSquareDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalAreaInverseSquareDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseSquareDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseSquareCentimeters function
func TestReciprocalAreaFactory_FromInverseSquareCentimeters(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseSquareCentimeters(100)
    if err != nil {
        t.Errorf("FromInverseSquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalAreaInverseSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseSquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseSquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromInverseSquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromInverseSquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromInverseSquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromInverseSquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseSquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseSquareCentimeters(0)
    if err != nil {
        t.Errorf("FromInverseSquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalAreaInverseSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseSquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseSquareMillimeters function
func TestReciprocalAreaFactory_FromInverseSquareMillimeters(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseSquareMillimeters(100)
    if err != nil {
        t.Errorf("FromInverseSquareMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalAreaInverseSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseSquareMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseSquareMillimeters(math.NaN())
    if err == nil {
        t.Error("FromInverseSquareMillimeters() with NaN value should return error")
    }

    _, err = factory.FromInverseSquareMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromInverseSquareMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromInverseSquareMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseSquareMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseSquareMillimeters(0)
    if err != nil {
        t.Errorf("FromInverseSquareMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalAreaInverseSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseSquareMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseSquareMicrometers function
func TestReciprocalAreaFactory_FromInverseSquareMicrometers(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseSquareMicrometers(100)
    if err != nil {
        t.Errorf("FromInverseSquareMicrometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalAreaInverseSquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseSquareMicrometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseSquareMicrometers(math.NaN())
    if err == nil {
        t.Error("FromInverseSquareMicrometers() with NaN value should return error")
    }

    _, err = factory.FromInverseSquareMicrometers(math.Inf(1))
    if err == nil {
        t.Error("FromInverseSquareMicrometers() with +Inf value should return error")
    }

    _, err = factory.FromInverseSquareMicrometers(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseSquareMicrometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseSquareMicrometers(0)
    if err != nil {
        t.Errorf("FromInverseSquareMicrometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalAreaInverseSquareMicrometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseSquareMicrometers() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseSquareMiles function
func TestReciprocalAreaFactory_FromInverseSquareMiles(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseSquareMiles(100)
    if err != nil {
        t.Errorf("FromInverseSquareMiles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalAreaInverseSquareMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseSquareMiles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseSquareMiles(math.NaN())
    if err == nil {
        t.Error("FromInverseSquareMiles() with NaN value should return error")
    }

    _, err = factory.FromInverseSquareMiles(math.Inf(1))
    if err == nil {
        t.Error("FromInverseSquareMiles() with +Inf value should return error")
    }

    _, err = factory.FromInverseSquareMiles(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseSquareMiles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseSquareMiles(0)
    if err != nil {
        t.Errorf("FromInverseSquareMiles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalAreaInverseSquareMile)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseSquareMiles() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseSquareYards function
func TestReciprocalAreaFactory_FromInverseSquareYards(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseSquareYards(100)
    if err != nil {
        t.Errorf("FromInverseSquareYards() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalAreaInverseSquareYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseSquareYards() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseSquareYards(math.NaN())
    if err == nil {
        t.Error("FromInverseSquareYards() with NaN value should return error")
    }

    _, err = factory.FromInverseSquareYards(math.Inf(1))
    if err == nil {
        t.Error("FromInverseSquareYards() with +Inf value should return error")
    }

    _, err = factory.FromInverseSquareYards(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseSquareYards() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseSquareYards(0)
    if err != nil {
        t.Errorf("FromInverseSquareYards() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalAreaInverseSquareYard)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseSquareYards() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseSquareFeet function
func TestReciprocalAreaFactory_FromInverseSquareFeet(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseSquareFeet(100)
    if err != nil {
        t.Errorf("FromInverseSquareFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalAreaInverseSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseSquareFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseSquareFeet(math.NaN())
    if err == nil {
        t.Error("FromInverseSquareFeet() with NaN value should return error")
    }

    _, err = factory.FromInverseSquareFeet(math.Inf(1))
    if err == nil {
        t.Error("FromInverseSquareFeet() with +Inf value should return error")
    }

    _, err = factory.FromInverseSquareFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseSquareFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseSquareFeet(0)
    if err != nil {
        t.Errorf("FromInverseSquareFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalAreaInverseSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseSquareFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseUsSurveySquareFeet function
func TestReciprocalAreaFactory_FromInverseUsSurveySquareFeet(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseUsSurveySquareFeet(100)
    if err != nil {
        t.Errorf("FromInverseUsSurveySquareFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalAreaInverseUsSurveySquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseUsSurveySquareFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseUsSurveySquareFeet(math.NaN())
    if err == nil {
        t.Error("FromInverseUsSurveySquareFeet() with NaN value should return error")
    }

    _, err = factory.FromInverseUsSurveySquareFeet(math.Inf(1))
    if err == nil {
        t.Error("FromInverseUsSurveySquareFeet() with +Inf value should return error")
    }

    _, err = factory.FromInverseUsSurveySquareFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseUsSurveySquareFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseUsSurveySquareFeet(0)
    if err != nil {
        t.Errorf("FromInverseUsSurveySquareFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalAreaInverseUsSurveySquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseUsSurveySquareFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseSquareInches function
func TestReciprocalAreaFactory_FromInverseSquareInches(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseSquareInches(100)
    if err != nil {
        t.Errorf("FromInverseSquareInches() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalAreaInverseSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseSquareInches() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseSquareInches(math.NaN())
    if err == nil {
        t.Error("FromInverseSquareInches() with NaN value should return error")
    }

    _, err = factory.FromInverseSquareInches(math.Inf(1))
    if err == nil {
        t.Error("FromInverseSquareInches() with +Inf value should return error")
    }

    _, err = factory.FromInverseSquareInches(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseSquareInches() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseSquareInches(0)
    if err != nil {
        t.Errorf("FromInverseSquareInches() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalAreaInverseSquareInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseSquareInches() with zero value = %v, want 0", converted)
    }
}

func TestReciprocalAreaToString(t *testing.T) {
	factory := units.ReciprocalAreaFactory{}
	a, err := factory.CreateReciprocalArea(45, units.ReciprocalAreaInverseSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ReciprocalAreaInverseSquareMeter, 2)
	expected := "45.00 " + units.GetReciprocalAreaAbbreviation(units.ReciprocalAreaInverseSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ReciprocalAreaInverseSquareMeter, -1)
	expected = "45 " + units.GetReciprocalAreaAbbreviation(units.ReciprocalAreaInverseSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestReciprocalArea_EqualityAndComparison(t *testing.T) {
	factory := units.ReciprocalAreaFactory{}
	a1, _ := factory.CreateReciprocalArea(60, units.ReciprocalAreaInverseSquareMeter)
	a2, _ := factory.CreateReciprocalArea(60, units.ReciprocalAreaInverseSquareMeter)
	a3, _ := factory.CreateReciprocalArea(120, units.ReciprocalAreaInverseSquareMeter)

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

func TestReciprocalArea_Arithmetic(t *testing.T) {
	factory := units.ReciprocalAreaFactory{}
	a1, _ := factory.CreateReciprocalArea(30, units.ReciprocalAreaInverseSquareMeter)
	a2, _ := factory.CreateReciprocalArea(45, units.ReciprocalAreaInverseSquareMeter)

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


func TestGetReciprocalAreaAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ReciprocalAreaUnits
        want string
    }{
        {
            name: "InverseSquareMeter abbreviation",
            unit: units.ReciprocalAreaInverseSquareMeter,
            want: "m⁻²",
        },
        {
            name: "InverseSquareKilometer abbreviation",
            unit: units.ReciprocalAreaInverseSquareKilometer,
            want: "km⁻²",
        },
        {
            name: "InverseSquareDecimeter abbreviation",
            unit: units.ReciprocalAreaInverseSquareDecimeter,
            want: "dm⁻²",
        },
        {
            name: "InverseSquareCentimeter abbreviation",
            unit: units.ReciprocalAreaInverseSquareCentimeter,
            want: "cm⁻²",
        },
        {
            name: "InverseSquareMillimeter abbreviation",
            unit: units.ReciprocalAreaInverseSquareMillimeter,
            want: "mm⁻²",
        },
        {
            name: "InverseSquareMicrometer abbreviation",
            unit: units.ReciprocalAreaInverseSquareMicrometer,
            want: "µm⁻²",
        },
        {
            name: "InverseSquareMile abbreviation",
            unit: units.ReciprocalAreaInverseSquareMile,
            want: "mi⁻²",
        },
        {
            name: "InverseSquareYard abbreviation",
            unit: units.ReciprocalAreaInverseSquareYard,
            want: "yd⁻²",
        },
        {
            name: "InverseSquareFoot abbreviation",
            unit: units.ReciprocalAreaInverseSquareFoot,
            want: "ft⁻²",
        },
        {
            name: "InverseUsSurveySquareFoot abbreviation",
            unit: units.ReciprocalAreaInverseUsSurveySquareFoot,
            want: "ft⁻² (US)",
        },
        {
            name: "InverseSquareInch abbreviation",
            unit: units.ReciprocalAreaInverseSquareInch,
            want: "in⁻²",
        },
        {
            name: "invalid unit",
            unit: units.ReciprocalAreaUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetReciprocalAreaAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetReciprocalAreaAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestReciprocalArea_String(t *testing.T) {
    factory := units.ReciprocalAreaFactory{}
    
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
            unit, err := factory.CreateReciprocalArea(tt.value, units.ReciprocalAreaInverseSquareMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ReciprocalArea.String() = %v, want %v", got, tt.want)
            }
        })
    }
}