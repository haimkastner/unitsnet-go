// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAreaDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SquareMeter"}`
	
	factory := units.AreaDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AreaSquareMeter {
		t.Errorf("expected unit %v, got %v", units.AreaSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAreaDto_ToJSON(t *testing.T) {
	dto := units.AreaDto{
		Value: 45,
		Unit:  units.AreaSquareMeter,
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
	if result["unit"].(string) != string(units.AreaSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.AreaSquareMeter, result["unit"])
	}
}

func TestNewArea_InvalidValue(t *testing.T) {
	factory := units.AreaFactory{}
	// NaN value should return an error.
	_, err := factory.CreateArea(math.NaN(), units.AreaSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateArea(math.Inf(1), units.AreaSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAreaConversions(t *testing.T) {
	factory := units.AreaFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateArea(180, units.AreaSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to SquareKilometers.
		// No expected conversion value provided for SquareKilometers, verifying result is not NaN.
		result := a.SquareKilometers()
		cacheResult := a.SquareKilometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareKilometers returned NaN")
		}
	}
	{
		// Test conversion to SquareMeters.
		// No expected conversion value provided for SquareMeters, verifying result is not NaN.
		result := a.SquareMeters()
		cacheResult := a.SquareMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareMeters returned NaN")
		}
	}
	{
		// Test conversion to SquareDecimeters.
		// No expected conversion value provided for SquareDecimeters, verifying result is not NaN.
		result := a.SquareDecimeters()
		cacheResult := a.SquareDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareDecimeters returned NaN")
		}
	}
	{
		// Test conversion to SquareCentimeters.
		// No expected conversion value provided for SquareCentimeters, verifying result is not NaN.
		result := a.SquareCentimeters()
		cacheResult := a.SquareCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareCentimeters returned NaN")
		}
	}
	{
		// Test conversion to SquareMillimeters.
		// No expected conversion value provided for SquareMillimeters, verifying result is not NaN.
		result := a.SquareMillimeters()
		cacheResult := a.SquareMillimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareMillimeters returned NaN")
		}
	}
	{
		// Test conversion to SquareMicrometers.
		// No expected conversion value provided for SquareMicrometers, verifying result is not NaN.
		result := a.SquareMicrometers()
		cacheResult := a.SquareMicrometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareMicrometers returned NaN")
		}
	}
	{
		// Test conversion to SquareMiles.
		// No expected conversion value provided for SquareMiles, verifying result is not NaN.
		result := a.SquareMiles()
		cacheResult := a.SquareMiles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareMiles returned NaN")
		}
	}
	{
		// Test conversion to SquareYards.
		// No expected conversion value provided for SquareYards, verifying result is not NaN.
		result := a.SquareYards()
		cacheResult := a.SquareYards()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareYards returned NaN")
		}
	}
	{
		// Test conversion to SquareFeet.
		// No expected conversion value provided for SquareFeet, verifying result is not NaN.
		result := a.SquareFeet()
		cacheResult := a.SquareFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareFeet returned NaN")
		}
	}
	{
		// Test conversion to UsSurveySquareFeet.
		// No expected conversion value provided for UsSurveySquareFeet, verifying result is not NaN.
		result := a.UsSurveySquareFeet()
		cacheResult := a.UsSurveySquareFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsSurveySquareFeet returned NaN")
		}
	}
	{
		// Test conversion to SquareInches.
		// No expected conversion value provided for SquareInches, verifying result is not NaN.
		result := a.SquareInches()
		cacheResult := a.SquareInches()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareInches returned NaN")
		}
	}
	{
		// Test conversion to Acres.
		// No expected conversion value provided for Acres, verifying result is not NaN.
		result := a.Acres()
		cacheResult := a.Acres()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Acres returned NaN")
		}
	}
	{
		// Test conversion to Hectares.
		// No expected conversion value provided for Hectares, verifying result is not NaN.
		result := a.Hectares()
		cacheResult := a.Hectares()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Hectares returned NaN")
		}
	}
	{
		// Test conversion to SquareNauticalMiles.
		// No expected conversion value provided for SquareNauticalMiles, verifying result is not NaN.
		result := a.SquareNauticalMiles()
		cacheResult := a.SquareNauticalMiles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareNauticalMiles returned NaN")
		}
	}
}

func TestArea_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AreaFactory{}
	a, err := factory.CreateArea(90, units.AreaSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AreaSquareMeter {
		t.Errorf("expected default unit SquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AreaSquareKilometer
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AreaDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AreaSquareMeter {
		t.Errorf("expected unit SquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAreaFactory_FromDto(t *testing.T) {
    factory := units.AreaFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaSquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.AreaDto{
        Value: math.NaN(),
        Unit:  units.AreaSquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test SquareKilometer conversion
    square_kilometersDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaSquareKilometer,
    }
    
    var square_kilometersResult *units.Area
    square_kilometersResult, err = factory.FromDto(square_kilometersDto)
    if err != nil {
        t.Errorf("FromDto() with SquareKilometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_kilometersResult.Convert(units.AreaSquareKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareKilometer = %v, want %v", converted, 100)
    }
    // Test SquareMeter conversion
    square_metersDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaSquareMeter,
    }
    
    var square_metersResult *units.Area
    square_metersResult, err = factory.FromDto(square_metersDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_metersResult.Convert(units.AreaSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeter = %v, want %v", converted, 100)
    }
    // Test SquareDecimeter conversion
    square_decimetersDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaSquareDecimeter,
    }
    
    var square_decimetersResult *units.Area
    square_decimetersResult, err = factory.FromDto(square_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with SquareDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_decimetersResult.Convert(units.AreaSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareDecimeter = %v, want %v", converted, 100)
    }
    // Test SquareCentimeter conversion
    square_centimetersDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaSquareCentimeter,
    }
    
    var square_centimetersResult *units.Area
    square_centimetersResult, err = factory.FromDto(square_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with SquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_centimetersResult.Convert(units.AreaSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareCentimeter = %v, want %v", converted, 100)
    }
    // Test SquareMillimeter conversion
    square_millimetersDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaSquareMillimeter,
    }
    
    var square_millimetersResult *units.Area
    square_millimetersResult, err = factory.FromDto(square_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_millimetersResult.Convert(units.AreaSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMillimeter = %v, want %v", converted, 100)
    }
    // Test SquareMicrometer conversion
    square_micrometersDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaSquareMicrometer,
    }
    
    var square_micrometersResult *units.Area
    square_micrometersResult, err = factory.FromDto(square_micrometersDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMicrometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_micrometersResult.Convert(units.AreaSquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMicrometer = %v, want %v", converted, 100)
    }
    // Test SquareMile conversion
    square_milesDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaSquareMile,
    }
    
    var square_milesResult *units.Area
    square_milesResult, err = factory.FromDto(square_milesDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMile returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_milesResult.Convert(units.AreaSquareMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMile = %v, want %v", converted, 100)
    }
    // Test SquareYard conversion
    square_yardsDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaSquareYard,
    }
    
    var square_yardsResult *units.Area
    square_yardsResult, err = factory.FromDto(square_yardsDto)
    if err != nil {
        t.Errorf("FromDto() with SquareYard returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_yardsResult.Convert(units.AreaSquareYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareYard = %v, want %v", converted, 100)
    }
    // Test SquareFoot conversion
    square_feetDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaSquareFoot,
    }
    
    var square_feetResult *units.Area
    square_feetResult, err = factory.FromDto(square_feetDto)
    if err != nil {
        t.Errorf("FromDto() with SquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_feetResult.Convert(units.AreaSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareFoot = %v, want %v", converted, 100)
    }
    // Test UsSurveySquareFoot conversion
    us_survey_square_feetDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaUsSurveySquareFoot,
    }
    
    var us_survey_square_feetResult *units.Area
    us_survey_square_feetResult, err = factory.FromDto(us_survey_square_feetDto)
    if err != nil {
        t.Errorf("FromDto() with UsSurveySquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_survey_square_feetResult.Convert(units.AreaUsSurveySquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsSurveySquareFoot = %v, want %v", converted, 100)
    }
    // Test SquareInch conversion
    square_inchesDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaSquareInch,
    }
    
    var square_inchesResult *units.Area
    square_inchesResult, err = factory.FromDto(square_inchesDto)
    if err != nil {
        t.Errorf("FromDto() with SquareInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_inchesResult.Convert(units.AreaSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareInch = %v, want %v", converted, 100)
    }
    // Test Acre conversion
    acresDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaAcre,
    }
    
    var acresResult *units.Area
    acresResult, err = factory.FromDto(acresDto)
    if err != nil {
        t.Errorf("FromDto() with Acre returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = acresResult.Convert(units.AreaAcre)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Acre = %v, want %v", converted, 100)
    }
    // Test Hectare conversion
    hectaresDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaHectare,
    }
    
    var hectaresResult *units.Area
    hectaresResult, err = factory.FromDto(hectaresDto)
    if err != nil {
        t.Errorf("FromDto() with Hectare returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectaresResult.Convert(units.AreaHectare)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hectare = %v, want %v", converted, 100)
    }
    // Test SquareNauticalMile conversion
    square_nautical_milesDto := units.AreaDto{
        Value: 100,
        Unit:  units.AreaSquareNauticalMile,
    }
    
    var square_nautical_milesResult *units.Area
    square_nautical_milesResult, err = factory.FromDto(square_nautical_milesDto)
    if err != nil {
        t.Errorf("FromDto() with SquareNauticalMile returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_nautical_milesResult.Convert(units.AreaSquareNauticalMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareNauticalMile = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.AreaDto{
        Value: 0,
        Unit:  units.AreaSquareMeter,
    }
    
    var zeroResult *units.Area
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestAreaFactory_FromDtoJSON(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "SquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "SquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.AreaDto{
        Value: nanValue,
        Unit:  units.AreaSquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with SquareKilometer unit
    square_kilometersJSON := []byte(`{"value": 100, "unit": "SquareKilometer"}`)
    square_kilometersResult, err := factory.FromDtoJSON(square_kilometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareKilometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_kilometersResult.Convert(units.AreaSquareKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareKilometer = %v, want %v", converted, 100)
    }
    // Test JSON with SquareMeter unit
    square_metersJSON := []byte(`{"value": 100, "unit": "SquareMeter"}`)
    square_metersResult, err := factory.FromDtoJSON(square_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_metersResult.Convert(units.AreaSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with SquareDecimeter unit
    square_decimetersJSON := []byte(`{"value": 100, "unit": "SquareDecimeter"}`)
    square_decimetersResult, err := factory.FromDtoJSON(square_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_decimetersResult.Convert(units.AreaSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with SquareCentimeter unit
    square_centimetersJSON := []byte(`{"value": 100, "unit": "SquareCentimeter"}`)
    square_centimetersResult, err := factory.FromDtoJSON(square_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_centimetersResult.Convert(units.AreaSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with SquareMillimeter unit
    square_millimetersJSON := []byte(`{"value": 100, "unit": "SquareMillimeter"}`)
    square_millimetersResult, err := factory.FromDtoJSON(square_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_millimetersResult.Convert(units.AreaSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with SquareMicrometer unit
    square_micrometersJSON := []byte(`{"value": 100, "unit": "SquareMicrometer"}`)
    square_micrometersResult, err := factory.FromDtoJSON(square_micrometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareMicrometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_micrometersResult.Convert(units.AreaSquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMicrometer = %v, want %v", converted, 100)
    }
    // Test JSON with SquareMile unit
    square_milesJSON := []byte(`{"value": 100, "unit": "SquareMile"}`)
    square_milesResult, err := factory.FromDtoJSON(square_milesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareMile unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_milesResult.Convert(units.AreaSquareMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMile = %v, want %v", converted, 100)
    }
    // Test JSON with SquareYard unit
    square_yardsJSON := []byte(`{"value": 100, "unit": "SquareYard"}`)
    square_yardsResult, err := factory.FromDtoJSON(square_yardsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareYard unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_yardsResult.Convert(units.AreaSquareYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareYard = %v, want %v", converted, 100)
    }
    // Test JSON with SquareFoot unit
    square_feetJSON := []byte(`{"value": 100, "unit": "SquareFoot"}`)
    square_feetResult, err := factory.FromDtoJSON(square_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_feetResult.Convert(units.AreaSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with UsSurveySquareFoot unit
    us_survey_square_feetJSON := []byte(`{"value": 100, "unit": "UsSurveySquareFoot"}`)
    us_survey_square_feetResult, err := factory.FromDtoJSON(us_survey_square_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsSurveySquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_survey_square_feetResult.Convert(units.AreaUsSurveySquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsSurveySquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with SquareInch unit
    square_inchesJSON := []byte(`{"value": 100, "unit": "SquareInch"}`)
    square_inchesResult, err := factory.FromDtoJSON(square_inchesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_inchesResult.Convert(units.AreaSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareInch = %v, want %v", converted, 100)
    }
    // Test JSON with Acre unit
    acresJSON := []byte(`{"value": 100, "unit": "Acre"}`)
    acresResult, err := factory.FromDtoJSON(acresJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Acre unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = acresResult.Convert(units.AreaAcre)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Acre = %v, want %v", converted, 100)
    }
    // Test JSON with Hectare unit
    hectaresJSON := []byte(`{"value": 100, "unit": "Hectare"}`)
    hectaresResult, err := factory.FromDtoJSON(hectaresJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Hectare unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectaresResult.Convert(units.AreaHectare)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hectare = %v, want %v", converted, 100)
    }
    // Test JSON with SquareNauticalMile unit
    square_nautical_milesJSON := []byte(`{"value": 100, "unit": "SquareNauticalMile"}`)
    square_nautical_milesResult, err := factory.FromDtoJSON(square_nautical_milesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareNauticalMile unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_nautical_milesResult.Convert(units.AreaSquareNauticalMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareNauticalMile = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "SquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromSquareKilometers function
func TestAreaFactory_FromSquareKilometers(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareKilometers(100)
    if err != nil {
        t.Errorf("FromSquareKilometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaSquareKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareKilometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareKilometers(math.NaN())
    if err == nil {
        t.Error("FromSquareKilometers() with NaN value should return error")
    }

    _, err = factory.FromSquareKilometers(math.Inf(1))
    if err == nil {
        t.Error("FromSquareKilometers() with +Inf value should return error")
    }

    _, err = factory.FromSquareKilometers(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareKilometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareKilometers(0)
    if err != nil {
        t.Errorf("FromSquareKilometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaSquareKilometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareKilometers() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareMeters function
func TestAreaFactory_FromSquareMeters(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMeters(100)
    if err != nil {
        t.Errorf("FromSquareMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareMeters(math.NaN())
    if err == nil {
        t.Error("FromSquareMeters() with NaN value should return error")
    }

    _, err = factory.FromSquareMeters(math.Inf(1))
    if err == nil {
        t.Error("FromSquareMeters() with +Inf value should return error")
    }

    _, err = factory.FromSquareMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareMeters(0)
    if err != nil {
        t.Errorf("FromSquareMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareDecimeters function
func TestAreaFactory_FromSquareDecimeters(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareDecimeters(100)
    if err != nil {
        t.Errorf("FromSquareDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaSquareDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareDecimeters(math.NaN())
    if err == nil {
        t.Error("FromSquareDecimeters() with NaN value should return error")
    }

    _, err = factory.FromSquareDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromSquareDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromSquareDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareDecimeters(0)
    if err != nil {
        t.Errorf("FromSquareDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaSquareDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareCentimeters function
func TestAreaFactory_FromSquareCentimeters(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareCentimeters(100)
    if err != nil {
        t.Errorf("FromSquareCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareCentimeters(math.NaN())
    if err == nil {
        t.Error("FromSquareCentimeters() with NaN value should return error")
    }

    _, err = factory.FromSquareCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromSquareCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromSquareCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareCentimeters(0)
    if err != nil {
        t.Errorf("FromSquareCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareMillimeters function
func TestAreaFactory_FromSquareMillimeters(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMillimeters(100)
    if err != nil {
        t.Errorf("FromSquareMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareMillimeters(math.NaN())
    if err == nil {
        t.Error("FromSquareMillimeters() with NaN value should return error")
    }

    _, err = factory.FromSquareMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromSquareMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromSquareMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareMillimeters(0)
    if err != nil {
        t.Errorf("FromSquareMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareMicrometers function
func TestAreaFactory_FromSquareMicrometers(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMicrometers(100)
    if err != nil {
        t.Errorf("FromSquareMicrometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaSquareMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareMicrometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareMicrometers(math.NaN())
    if err == nil {
        t.Error("FromSquareMicrometers() with NaN value should return error")
    }

    _, err = factory.FromSquareMicrometers(math.Inf(1))
    if err == nil {
        t.Error("FromSquareMicrometers() with +Inf value should return error")
    }

    _, err = factory.FromSquareMicrometers(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareMicrometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareMicrometers(0)
    if err != nil {
        t.Errorf("FromSquareMicrometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaSquareMicrometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMicrometers() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareMiles function
func TestAreaFactory_FromSquareMiles(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMiles(100)
    if err != nil {
        t.Errorf("FromSquareMiles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaSquareMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareMiles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareMiles(math.NaN())
    if err == nil {
        t.Error("FromSquareMiles() with NaN value should return error")
    }

    _, err = factory.FromSquareMiles(math.Inf(1))
    if err == nil {
        t.Error("FromSquareMiles() with +Inf value should return error")
    }

    _, err = factory.FromSquareMiles(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareMiles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareMiles(0)
    if err != nil {
        t.Errorf("FromSquareMiles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaSquareMile)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMiles() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareYards function
func TestAreaFactory_FromSquareYards(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareYards(100)
    if err != nil {
        t.Errorf("FromSquareYards() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaSquareYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareYards() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareYards(math.NaN())
    if err == nil {
        t.Error("FromSquareYards() with NaN value should return error")
    }

    _, err = factory.FromSquareYards(math.Inf(1))
    if err == nil {
        t.Error("FromSquareYards() with +Inf value should return error")
    }

    _, err = factory.FromSquareYards(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareYards() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareYards(0)
    if err != nil {
        t.Errorf("FromSquareYards() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaSquareYard)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareYards() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareFeet function
func TestAreaFactory_FromSquareFeet(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareFeet(100)
    if err != nil {
        t.Errorf("FromSquareFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareFeet(math.NaN())
    if err == nil {
        t.Error("FromSquareFeet() with NaN value should return error")
    }

    _, err = factory.FromSquareFeet(math.Inf(1))
    if err == nil {
        t.Error("FromSquareFeet() with +Inf value should return error")
    }

    _, err = factory.FromSquareFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareFeet(0)
    if err != nil {
        t.Errorf("FromSquareFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromUsSurveySquareFeet function
func TestAreaFactory_FromUsSurveySquareFeet(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsSurveySquareFeet(100)
    if err != nil {
        t.Errorf("FromUsSurveySquareFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaUsSurveySquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsSurveySquareFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsSurveySquareFeet(math.NaN())
    if err == nil {
        t.Error("FromUsSurveySquareFeet() with NaN value should return error")
    }

    _, err = factory.FromUsSurveySquareFeet(math.Inf(1))
    if err == nil {
        t.Error("FromUsSurveySquareFeet() with +Inf value should return error")
    }

    _, err = factory.FromUsSurveySquareFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromUsSurveySquareFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsSurveySquareFeet(0)
    if err != nil {
        t.Errorf("FromUsSurveySquareFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaUsSurveySquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsSurveySquareFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareInches function
func TestAreaFactory_FromSquareInches(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareInches(100)
    if err != nil {
        t.Errorf("FromSquareInches() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareInches() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareInches(math.NaN())
    if err == nil {
        t.Error("FromSquareInches() with NaN value should return error")
    }

    _, err = factory.FromSquareInches(math.Inf(1))
    if err == nil {
        t.Error("FromSquareInches() with +Inf value should return error")
    }

    _, err = factory.FromSquareInches(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareInches() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareInches(0)
    if err != nil {
        t.Errorf("FromSquareInches() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaSquareInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareInches() with zero value = %v, want 0", converted)
    }
}
// Test FromAcres function
func TestAreaFactory_FromAcres(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAcres(100)
    if err != nil {
        t.Errorf("FromAcres() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaAcre)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAcres() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAcres(math.NaN())
    if err == nil {
        t.Error("FromAcres() with NaN value should return error")
    }

    _, err = factory.FromAcres(math.Inf(1))
    if err == nil {
        t.Error("FromAcres() with +Inf value should return error")
    }

    _, err = factory.FromAcres(math.Inf(-1))
    if err == nil {
        t.Error("FromAcres() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAcres(0)
    if err != nil {
        t.Errorf("FromAcres() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaAcre)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAcres() with zero value = %v, want 0", converted)
    }
}
// Test FromHectares function
func TestAreaFactory_FromHectares(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectares(100)
    if err != nil {
        t.Errorf("FromHectares() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaHectare)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectares() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectares(math.NaN())
    if err == nil {
        t.Error("FromHectares() with NaN value should return error")
    }

    _, err = factory.FromHectares(math.Inf(1))
    if err == nil {
        t.Error("FromHectares() with +Inf value should return error")
    }

    _, err = factory.FromHectares(math.Inf(-1))
    if err == nil {
        t.Error("FromHectares() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectares(0)
    if err != nil {
        t.Errorf("FromHectares() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaHectare)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectares() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareNauticalMiles function
func TestAreaFactory_FromSquareNauticalMiles(t *testing.T) {
    factory := units.AreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareNauticalMiles(100)
    if err != nil {
        t.Errorf("FromSquareNauticalMiles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaSquareNauticalMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareNauticalMiles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareNauticalMiles(math.NaN())
    if err == nil {
        t.Error("FromSquareNauticalMiles() with NaN value should return error")
    }

    _, err = factory.FromSquareNauticalMiles(math.Inf(1))
    if err == nil {
        t.Error("FromSquareNauticalMiles() with +Inf value should return error")
    }

    _, err = factory.FromSquareNauticalMiles(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareNauticalMiles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareNauticalMiles(0)
    if err != nil {
        t.Errorf("FromSquareNauticalMiles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaSquareNauticalMile)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareNauticalMiles() with zero value = %v, want 0", converted)
    }
}

func TestAreaToString(t *testing.T) {
	factory := units.AreaFactory{}
	a, err := factory.CreateArea(45, units.AreaSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AreaSquareMeter, 2)
	expected := "45.00 " + units.GetAreaAbbreviation(units.AreaSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AreaSquareMeter, -1)
	expected = "45 " + units.GetAreaAbbreviation(units.AreaSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestArea_EqualityAndComparison(t *testing.T) {
	factory := units.AreaFactory{}
	a1, _ := factory.CreateArea(60, units.AreaSquareMeter)
	a2, _ := factory.CreateArea(60, units.AreaSquareMeter)
	a3, _ := factory.CreateArea(120, units.AreaSquareMeter)

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

func TestArea_Arithmetic(t *testing.T) {
	factory := units.AreaFactory{}
	a1, _ := factory.CreateArea(30, units.AreaSquareMeter)
	a2, _ := factory.CreateArea(45, units.AreaSquareMeter)

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


func TestGetAreaAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.AreaUnits
        want string
    }{
        {
            name: "SquareKilometer abbreviation",
            unit: units.AreaSquareKilometer,
            want: "km²",
        },
        {
            name: "SquareMeter abbreviation",
            unit: units.AreaSquareMeter,
            want: "m²",
        },
        {
            name: "SquareDecimeter abbreviation",
            unit: units.AreaSquareDecimeter,
            want: "dm²",
        },
        {
            name: "SquareCentimeter abbreviation",
            unit: units.AreaSquareCentimeter,
            want: "cm²",
        },
        {
            name: "SquareMillimeter abbreviation",
            unit: units.AreaSquareMillimeter,
            want: "mm²",
        },
        {
            name: "SquareMicrometer abbreviation",
            unit: units.AreaSquareMicrometer,
            want: "µm²",
        },
        {
            name: "SquareMile abbreviation",
            unit: units.AreaSquareMile,
            want: "mi²",
        },
        {
            name: "SquareYard abbreviation",
            unit: units.AreaSquareYard,
            want: "yd²",
        },
        {
            name: "SquareFoot abbreviation",
            unit: units.AreaSquareFoot,
            want: "ft²",
        },
        {
            name: "UsSurveySquareFoot abbreviation",
            unit: units.AreaUsSurveySquareFoot,
            want: "ft² (US)",
        },
        {
            name: "SquareInch abbreviation",
            unit: units.AreaSquareInch,
            want: "in²",
        },
        {
            name: "Acre abbreviation",
            unit: units.AreaAcre,
            want: "ac",
        },
        {
            name: "Hectare abbreviation",
            unit: units.AreaHectare,
            want: "ha",
        },
        {
            name: "SquareNauticalMile abbreviation",
            unit: units.AreaSquareNauticalMile,
            want: "nmi²",
        },
        {
            name: "invalid unit",
            unit: units.AreaUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetAreaAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetAreaAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestArea_String(t *testing.T) {
    factory := units.AreaFactory{}
    
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
            unit, err := factory.CreateArea(tt.value, units.AreaSquareMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Area.String() = %v, want %v", got, tt.want)
            }
        })
    }
}