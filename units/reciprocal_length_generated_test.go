// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestReciprocalLengthDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "InverseMeter"}`
	
	factory := units.ReciprocalLengthDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ReciprocalLengthInverseMeter {
		t.Errorf("expected unit %v, got %v", units.ReciprocalLengthInverseMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "InverseMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestReciprocalLengthDto_ToJSON(t *testing.T) {
	dto := units.ReciprocalLengthDto{
		Value: 45,
		Unit:  units.ReciprocalLengthInverseMeter,
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
	if result["unit"].(string) != string(units.ReciprocalLengthInverseMeter) {
		t.Errorf("expected unit %s, got %v", units.ReciprocalLengthInverseMeter, result["unit"])
	}
}

func TestNewReciprocalLength_InvalidValue(t *testing.T) {
	factory := units.ReciprocalLengthFactory{}
	// NaN value should return an error.
	_, err := factory.CreateReciprocalLength(math.NaN(), units.ReciprocalLengthInverseMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateReciprocalLength(math.Inf(1), units.ReciprocalLengthInverseMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestReciprocalLengthConversions(t *testing.T) {
	factory := units.ReciprocalLengthFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateReciprocalLength(180, units.ReciprocalLengthInverseMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to InverseMeters.
		// No expected conversion value provided for InverseMeters, verifying result is not NaN.
		result := a.InverseMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseMeters returned NaN")
		}
	}
	{
		// Test conversion to InverseCentimeters.
		// No expected conversion value provided for InverseCentimeters, verifying result is not NaN.
		result := a.InverseCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseCentimeters returned NaN")
		}
	}
	{
		// Test conversion to InverseMillimeters.
		// No expected conversion value provided for InverseMillimeters, verifying result is not NaN.
		result := a.InverseMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseMillimeters returned NaN")
		}
	}
	{
		// Test conversion to InverseMiles.
		// No expected conversion value provided for InverseMiles, verifying result is not NaN.
		result := a.InverseMiles()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseMiles returned NaN")
		}
	}
	{
		// Test conversion to InverseYards.
		// No expected conversion value provided for InverseYards, verifying result is not NaN.
		result := a.InverseYards()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseYards returned NaN")
		}
	}
	{
		// Test conversion to InverseFeet.
		// No expected conversion value provided for InverseFeet, verifying result is not NaN.
		result := a.InverseFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseFeet returned NaN")
		}
	}
	{
		// Test conversion to InverseUsSurveyFeet.
		// No expected conversion value provided for InverseUsSurveyFeet, verifying result is not NaN.
		result := a.InverseUsSurveyFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseUsSurveyFeet returned NaN")
		}
	}
	{
		// Test conversion to InverseInches.
		// No expected conversion value provided for InverseInches, verifying result is not NaN.
		result := a.InverseInches()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseInches returned NaN")
		}
	}
	{
		// Test conversion to InverseMils.
		// No expected conversion value provided for InverseMils, verifying result is not NaN.
		result := a.InverseMils()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseMils returned NaN")
		}
	}
	{
		// Test conversion to InverseMicroinches.
		// No expected conversion value provided for InverseMicroinches, verifying result is not NaN.
		result := a.InverseMicroinches()
		if math.IsNaN(result) {
			t.Errorf("conversion to InverseMicroinches returned NaN")
		}
	}
}

func TestReciprocalLength_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ReciprocalLengthFactory{}
	a, err := factory.CreateReciprocalLength(90, units.ReciprocalLengthInverseMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ReciprocalLengthInverseMeter {
		t.Errorf("expected default unit InverseMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ReciprocalLengthInverseMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ReciprocalLengthDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ReciprocalLengthInverseMeter {
		t.Errorf("expected unit InverseMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestReciprocalLengthFactory_FromDto(t *testing.T) {
    factory := units.ReciprocalLengthFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ReciprocalLengthDto{
        Value: 100,
        Unit:  units.ReciprocalLengthInverseMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ReciprocalLengthDto{
        Value: math.NaN(),
        Unit:  units.ReciprocalLengthInverseMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test InverseMeter conversion
    inverse_metersDto := units.ReciprocalLengthDto{
        Value: 100,
        Unit:  units.ReciprocalLengthInverseMeter,
    }
    
    var inverse_metersResult *units.ReciprocalLength
    inverse_metersResult, err = factory.FromDto(inverse_metersDto)
    if err != nil {
        t.Errorf("FromDto() with InverseMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_metersResult.Convert(units.ReciprocalLengthInverseMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMeter = %v, want %v", converted, 100)
    }
    // Test InverseCentimeter conversion
    inverse_centimetersDto := units.ReciprocalLengthDto{
        Value: 100,
        Unit:  units.ReciprocalLengthInverseCentimeter,
    }
    
    var inverse_centimetersResult *units.ReciprocalLength
    inverse_centimetersResult, err = factory.FromDto(inverse_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with InverseCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_centimetersResult.Convert(units.ReciprocalLengthInverseCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseCentimeter = %v, want %v", converted, 100)
    }
    // Test InverseMillimeter conversion
    inverse_millimetersDto := units.ReciprocalLengthDto{
        Value: 100,
        Unit:  units.ReciprocalLengthInverseMillimeter,
    }
    
    var inverse_millimetersResult *units.ReciprocalLength
    inverse_millimetersResult, err = factory.FromDto(inverse_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with InverseMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_millimetersResult.Convert(units.ReciprocalLengthInverseMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMillimeter = %v, want %v", converted, 100)
    }
    // Test InverseMile conversion
    inverse_milesDto := units.ReciprocalLengthDto{
        Value: 100,
        Unit:  units.ReciprocalLengthInverseMile,
    }
    
    var inverse_milesResult *units.ReciprocalLength
    inverse_milesResult, err = factory.FromDto(inverse_milesDto)
    if err != nil {
        t.Errorf("FromDto() with InverseMile returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_milesResult.Convert(units.ReciprocalLengthInverseMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMile = %v, want %v", converted, 100)
    }
    // Test InverseYard conversion
    inverse_yardsDto := units.ReciprocalLengthDto{
        Value: 100,
        Unit:  units.ReciprocalLengthInverseYard,
    }
    
    var inverse_yardsResult *units.ReciprocalLength
    inverse_yardsResult, err = factory.FromDto(inverse_yardsDto)
    if err != nil {
        t.Errorf("FromDto() with InverseYard returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_yardsResult.Convert(units.ReciprocalLengthInverseYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseYard = %v, want %v", converted, 100)
    }
    // Test InverseFoot conversion
    inverse_feetDto := units.ReciprocalLengthDto{
        Value: 100,
        Unit:  units.ReciprocalLengthInverseFoot,
    }
    
    var inverse_feetResult *units.ReciprocalLength
    inverse_feetResult, err = factory.FromDto(inverse_feetDto)
    if err != nil {
        t.Errorf("FromDto() with InverseFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_feetResult.Convert(units.ReciprocalLengthInverseFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseFoot = %v, want %v", converted, 100)
    }
    // Test InverseUsSurveyFoot conversion
    inverse_us_survey_feetDto := units.ReciprocalLengthDto{
        Value: 100,
        Unit:  units.ReciprocalLengthInverseUsSurveyFoot,
    }
    
    var inverse_us_survey_feetResult *units.ReciprocalLength
    inverse_us_survey_feetResult, err = factory.FromDto(inverse_us_survey_feetDto)
    if err != nil {
        t.Errorf("FromDto() with InverseUsSurveyFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_us_survey_feetResult.Convert(units.ReciprocalLengthInverseUsSurveyFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseUsSurveyFoot = %v, want %v", converted, 100)
    }
    // Test InverseInch conversion
    inverse_inchesDto := units.ReciprocalLengthDto{
        Value: 100,
        Unit:  units.ReciprocalLengthInverseInch,
    }
    
    var inverse_inchesResult *units.ReciprocalLength
    inverse_inchesResult, err = factory.FromDto(inverse_inchesDto)
    if err != nil {
        t.Errorf("FromDto() with InverseInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_inchesResult.Convert(units.ReciprocalLengthInverseInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseInch = %v, want %v", converted, 100)
    }
    // Test InverseMil conversion
    inverse_milsDto := units.ReciprocalLengthDto{
        Value: 100,
        Unit:  units.ReciprocalLengthInverseMil,
    }
    
    var inverse_milsResult *units.ReciprocalLength
    inverse_milsResult, err = factory.FromDto(inverse_milsDto)
    if err != nil {
        t.Errorf("FromDto() with InverseMil returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_milsResult.Convert(units.ReciprocalLengthInverseMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMil = %v, want %v", converted, 100)
    }
    // Test InverseMicroinch conversion
    inverse_microinchesDto := units.ReciprocalLengthDto{
        Value: 100,
        Unit:  units.ReciprocalLengthInverseMicroinch,
    }
    
    var inverse_microinchesResult *units.ReciprocalLength
    inverse_microinchesResult, err = factory.FromDto(inverse_microinchesDto)
    if err != nil {
        t.Errorf("FromDto() with InverseMicroinch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_microinchesResult.Convert(units.ReciprocalLengthInverseMicroinch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMicroinch = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ReciprocalLengthDto{
        Value: 0,
        Unit:  units.ReciprocalLengthInverseMeter,
    }
    
    var zeroResult *units.ReciprocalLength
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestReciprocalLengthFactory_FromDtoJSON(t *testing.T) {
    factory := units.ReciprocalLengthFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "InverseMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "InverseMeter"}`)
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
    nanJSON, _ := json.Marshal(units.ReciprocalLengthDto{
        Value: nanValue,
        Unit:  units.ReciprocalLengthInverseMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with InverseMeter unit
    inverse_metersJSON := []byte(`{"value": 100, "unit": "InverseMeter"}`)
    inverse_metersResult, err := factory.FromDtoJSON(inverse_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_metersResult.Convert(units.ReciprocalLengthInverseMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMeter = %v, want %v", converted, 100)
    }
    // Test JSON with InverseCentimeter unit
    inverse_centimetersJSON := []byte(`{"value": 100, "unit": "InverseCentimeter"}`)
    inverse_centimetersResult, err := factory.FromDtoJSON(inverse_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_centimetersResult.Convert(units.ReciprocalLengthInverseCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with InverseMillimeter unit
    inverse_millimetersJSON := []byte(`{"value": 100, "unit": "InverseMillimeter"}`)
    inverse_millimetersResult, err := factory.FromDtoJSON(inverse_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_millimetersResult.Convert(units.ReciprocalLengthInverseMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with InverseMile unit
    inverse_milesJSON := []byte(`{"value": 100, "unit": "InverseMile"}`)
    inverse_milesResult, err := factory.FromDtoJSON(inverse_milesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseMile unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_milesResult.Convert(units.ReciprocalLengthInverseMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMile = %v, want %v", converted, 100)
    }
    // Test JSON with InverseYard unit
    inverse_yardsJSON := []byte(`{"value": 100, "unit": "InverseYard"}`)
    inverse_yardsResult, err := factory.FromDtoJSON(inverse_yardsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseYard unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_yardsResult.Convert(units.ReciprocalLengthInverseYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseYard = %v, want %v", converted, 100)
    }
    // Test JSON with InverseFoot unit
    inverse_feetJSON := []byte(`{"value": 100, "unit": "InverseFoot"}`)
    inverse_feetResult, err := factory.FromDtoJSON(inverse_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_feetResult.Convert(units.ReciprocalLengthInverseFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseFoot = %v, want %v", converted, 100)
    }
    // Test JSON with InverseUsSurveyFoot unit
    inverse_us_survey_feetJSON := []byte(`{"value": 100, "unit": "InverseUsSurveyFoot"}`)
    inverse_us_survey_feetResult, err := factory.FromDtoJSON(inverse_us_survey_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseUsSurveyFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_us_survey_feetResult.Convert(units.ReciprocalLengthInverseUsSurveyFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseUsSurveyFoot = %v, want %v", converted, 100)
    }
    // Test JSON with InverseInch unit
    inverse_inchesJSON := []byte(`{"value": 100, "unit": "InverseInch"}`)
    inverse_inchesResult, err := factory.FromDtoJSON(inverse_inchesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_inchesResult.Convert(units.ReciprocalLengthInverseInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseInch = %v, want %v", converted, 100)
    }
    // Test JSON with InverseMil unit
    inverse_milsJSON := []byte(`{"value": 100, "unit": "InverseMil"}`)
    inverse_milsResult, err := factory.FromDtoJSON(inverse_milsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseMil unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_milsResult.Convert(units.ReciprocalLengthInverseMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMil = %v, want %v", converted, 100)
    }
    // Test JSON with InverseMicroinch unit
    inverse_microinchesJSON := []byte(`{"value": 100, "unit": "InverseMicroinch"}`)
    inverse_microinchesResult, err := factory.FromDtoJSON(inverse_microinchesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseMicroinch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_microinchesResult.Convert(units.ReciprocalLengthInverseMicroinch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMicroinch = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "InverseMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromInverseMeters function
func TestReciprocalLengthFactory_FromInverseMeters(t *testing.T) {
    factory := units.ReciprocalLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseMeters(100)
    if err != nil {
        t.Errorf("FromInverseMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalLengthInverseMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseMeters(math.NaN())
    if err == nil {
        t.Error("FromInverseMeters() with NaN value should return error")
    }

    _, err = factory.FromInverseMeters(math.Inf(1))
    if err == nil {
        t.Error("FromInverseMeters() with +Inf value should return error")
    }

    _, err = factory.FromInverseMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseMeters(0)
    if err != nil {
        t.Errorf("FromInverseMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalLengthInverseMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseCentimeters function
func TestReciprocalLengthFactory_FromInverseCentimeters(t *testing.T) {
    factory := units.ReciprocalLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseCentimeters(100)
    if err != nil {
        t.Errorf("FromInverseCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalLengthInverseCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseCentimeters(math.NaN())
    if err == nil {
        t.Error("FromInverseCentimeters() with NaN value should return error")
    }

    _, err = factory.FromInverseCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromInverseCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromInverseCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseCentimeters(0)
    if err != nil {
        t.Errorf("FromInverseCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalLengthInverseCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseMillimeters function
func TestReciprocalLengthFactory_FromInverseMillimeters(t *testing.T) {
    factory := units.ReciprocalLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseMillimeters(100)
    if err != nil {
        t.Errorf("FromInverseMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalLengthInverseMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseMillimeters(math.NaN())
    if err == nil {
        t.Error("FromInverseMillimeters() with NaN value should return error")
    }

    _, err = factory.FromInverseMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromInverseMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromInverseMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseMillimeters(0)
    if err != nil {
        t.Errorf("FromInverseMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalLengthInverseMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseMiles function
func TestReciprocalLengthFactory_FromInverseMiles(t *testing.T) {
    factory := units.ReciprocalLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseMiles(100)
    if err != nil {
        t.Errorf("FromInverseMiles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalLengthInverseMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseMiles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseMiles(math.NaN())
    if err == nil {
        t.Error("FromInverseMiles() with NaN value should return error")
    }

    _, err = factory.FromInverseMiles(math.Inf(1))
    if err == nil {
        t.Error("FromInverseMiles() with +Inf value should return error")
    }

    _, err = factory.FromInverseMiles(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseMiles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseMiles(0)
    if err != nil {
        t.Errorf("FromInverseMiles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalLengthInverseMile)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseMiles() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseYards function
func TestReciprocalLengthFactory_FromInverseYards(t *testing.T) {
    factory := units.ReciprocalLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseYards(100)
    if err != nil {
        t.Errorf("FromInverseYards() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalLengthInverseYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseYards() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseYards(math.NaN())
    if err == nil {
        t.Error("FromInverseYards() with NaN value should return error")
    }

    _, err = factory.FromInverseYards(math.Inf(1))
    if err == nil {
        t.Error("FromInverseYards() with +Inf value should return error")
    }

    _, err = factory.FromInverseYards(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseYards() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseYards(0)
    if err != nil {
        t.Errorf("FromInverseYards() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalLengthInverseYard)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseYards() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseFeet function
func TestReciprocalLengthFactory_FromInverseFeet(t *testing.T) {
    factory := units.ReciprocalLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseFeet(100)
    if err != nil {
        t.Errorf("FromInverseFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalLengthInverseFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseFeet(math.NaN())
    if err == nil {
        t.Error("FromInverseFeet() with NaN value should return error")
    }

    _, err = factory.FromInverseFeet(math.Inf(1))
    if err == nil {
        t.Error("FromInverseFeet() with +Inf value should return error")
    }

    _, err = factory.FromInverseFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseFeet(0)
    if err != nil {
        t.Errorf("FromInverseFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalLengthInverseFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseUsSurveyFeet function
func TestReciprocalLengthFactory_FromInverseUsSurveyFeet(t *testing.T) {
    factory := units.ReciprocalLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseUsSurveyFeet(100)
    if err != nil {
        t.Errorf("FromInverseUsSurveyFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalLengthInverseUsSurveyFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseUsSurveyFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseUsSurveyFeet(math.NaN())
    if err == nil {
        t.Error("FromInverseUsSurveyFeet() with NaN value should return error")
    }

    _, err = factory.FromInverseUsSurveyFeet(math.Inf(1))
    if err == nil {
        t.Error("FromInverseUsSurveyFeet() with +Inf value should return error")
    }

    _, err = factory.FromInverseUsSurveyFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseUsSurveyFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseUsSurveyFeet(0)
    if err != nil {
        t.Errorf("FromInverseUsSurveyFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalLengthInverseUsSurveyFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseUsSurveyFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseInches function
func TestReciprocalLengthFactory_FromInverseInches(t *testing.T) {
    factory := units.ReciprocalLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseInches(100)
    if err != nil {
        t.Errorf("FromInverseInches() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalLengthInverseInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseInches() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseInches(math.NaN())
    if err == nil {
        t.Error("FromInverseInches() with NaN value should return error")
    }

    _, err = factory.FromInverseInches(math.Inf(1))
    if err == nil {
        t.Error("FromInverseInches() with +Inf value should return error")
    }

    _, err = factory.FromInverseInches(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseInches() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseInches(0)
    if err != nil {
        t.Errorf("FromInverseInches() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalLengthInverseInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseInches() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseMils function
func TestReciprocalLengthFactory_FromInverseMils(t *testing.T) {
    factory := units.ReciprocalLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseMils(100)
    if err != nil {
        t.Errorf("FromInverseMils() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalLengthInverseMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseMils() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseMils(math.NaN())
    if err == nil {
        t.Error("FromInverseMils() with NaN value should return error")
    }

    _, err = factory.FromInverseMils(math.Inf(1))
    if err == nil {
        t.Error("FromInverseMils() with +Inf value should return error")
    }

    _, err = factory.FromInverseMils(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseMils() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseMils(0)
    if err != nil {
        t.Errorf("FromInverseMils() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalLengthInverseMil)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseMils() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseMicroinches function
func TestReciprocalLengthFactory_FromInverseMicroinches(t *testing.T) {
    factory := units.ReciprocalLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseMicroinches(100)
    if err != nil {
        t.Errorf("FromInverseMicroinches() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReciprocalLengthInverseMicroinch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseMicroinches() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseMicroinches(math.NaN())
    if err == nil {
        t.Error("FromInverseMicroinches() with NaN value should return error")
    }

    _, err = factory.FromInverseMicroinches(math.Inf(1))
    if err == nil {
        t.Error("FromInverseMicroinches() with +Inf value should return error")
    }

    _, err = factory.FromInverseMicroinches(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseMicroinches() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseMicroinches(0)
    if err != nil {
        t.Errorf("FromInverseMicroinches() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReciprocalLengthInverseMicroinch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseMicroinches() with zero value = %v, want 0", converted)
    }
}

func TestReciprocalLengthToString(t *testing.T) {
	factory := units.ReciprocalLengthFactory{}
	a, err := factory.CreateReciprocalLength(45, units.ReciprocalLengthInverseMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ReciprocalLengthInverseMeter, 2)
	expected := "45.00 " + units.GetReciprocalLengthAbbreviation(units.ReciprocalLengthInverseMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ReciprocalLengthInverseMeter, -1)
	expected = "45 " + units.GetReciprocalLengthAbbreviation(units.ReciprocalLengthInverseMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestReciprocalLength_EqualityAndComparison(t *testing.T) {
	factory := units.ReciprocalLengthFactory{}
	a1, _ := factory.CreateReciprocalLength(60, units.ReciprocalLengthInverseMeter)
	a2, _ := factory.CreateReciprocalLength(60, units.ReciprocalLengthInverseMeter)
	a3, _ := factory.CreateReciprocalLength(120, units.ReciprocalLengthInverseMeter)

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

func TestReciprocalLength_Arithmetic(t *testing.T) {
	factory := units.ReciprocalLengthFactory{}
	a1, _ := factory.CreateReciprocalLength(30, units.ReciprocalLengthInverseMeter)
	a2, _ := factory.CreateReciprocalLength(45, units.ReciprocalLengthInverseMeter)

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