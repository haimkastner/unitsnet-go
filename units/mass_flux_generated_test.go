// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMassFluxDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerSecondPerSquareMeter"}`
	
	factory := units.MassFluxDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MassFluxKilogramPerSecondPerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.MassFluxKilogramPerSecondPerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerSecondPerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMassFluxDto_ToJSON(t *testing.T) {
	dto := units.MassFluxDto{
		Value: 45,
		Unit:  units.MassFluxKilogramPerSecondPerSquareMeter,
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
	if result["unit"].(string) != string(units.MassFluxKilogramPerSecondPerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.MassFluxKilogramPerSecondPerSquareMeter, result["unit"])
	}
}

func TestNewMassFlux_InvalidValue(t *testing.T) {
	factory := units.MassFluxFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMassFlux(math.NaN(), units.MassFluxKilogramPerSecondPerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMassFlux(math.Inf(1), units.MassFluxKilogramPerSecondPerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMassFluxConversions(t *testing.T) {
	factory := units.MassFluxFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMassFlux(180, units.MassFluxKilogramPerSecondPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramsPerSecondPerSquareMeter.
		// No expected conversion value provided for GramsPerSecondPerSquareMeter, verifying result is not NaN.
		result := a.GramsPerSecondPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerSecondPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerSecondPerSquareCentimeter.
		// No expected conversion value provided for GramsPerSecondPerSquareCentimeter, verifying result is not NaN.
		result := a.GramsPerSecondPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerSecondPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerSecondPerSquareMillimeter.
		// No expected conversion value provided for GramsPerSecondPerSquareMillimeter, verifying result is not NaN.
		result := a.GramsPerSecondPerSquareMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerSecondPerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerHourPerSquareMeter.
		// No expected conversion value provided for GramsPerHourPerSquareMeter, verifying result is not NaN.
		result := a.GramsPerHourPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerHourPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerHourPerSquareCentimeter.
		// No expected conversion value provided for GramsPerHourPerSquareCentimeter, verifying result is not NaN.
		result := a.GramsPerHourPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerHourPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerHourPerSquareMillimeter.
		// No expected conversion value provided for GramsPerHourPerSquareMillimeter, verifying result is not NaN.
		result := a.GramsPerHourPerSquareMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerHourPerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerSecondPerSquareMeter.
		// No expected conversion value provided for KilogramsPerSecondPerSquareMeter, verifying result is not NaN.
		result := a.KilogramsPerSecondPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerSecondPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerSecondPerSquareCentimeter.
		// No expected conversion value provided for KilogramsPerSecondPerSquareCentimeter, verifying result is not NaN.
		result := a.KilogramsPerSecondPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerSecondPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerSecondPerSquareMillimeter.
		// No expected conversion value provided for KilogramsPerSecondPerSquareMillimeter, verifying result is not NaN.
		result := a.KilogramsPerSecondPerSquareMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerSecondPerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerHourPerSquareMeter.
		// No expected conversion value provided for KilogramsPerHourPerSquareMeter, verifying result is not NaN.
		result := a.KilogramsPerHourPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerHourPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerHourPerSquareCentimeter.
		// No expected conversion value provided for KilogramsPerHourPerSquareCentimeter, verifying result is not NaN.
		result := a.KilogramsPerHourPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerHourPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerHourPerSquareMillimeter.
		// No expected conversion value provided for KilogramsPerHourPerSquareMillimeter, verifying result is not NaN.
		result := a.KilogramsPerHourPerSquareMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerHourPerSquareMillimeter returned NaN")
		}
	}
}

func TestMassFlux_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MassFluxFactory{}
	a, err := factory.CreateMassFlux(90, units.MassFluxKilogramPerSecondPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MassFluxKilogramPerSecondPerSquareMeter {
		t.Errorf("expected default unit KilogramPerSecondPerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MassFluxGramPerSecondPerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MassFluxDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MassFluxKilogramPerSecondPerSquareMeter {
		t.Errorf("expected unit KilogramPerSecondPerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMassFluxFactory_FromDto(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MassFluxDto{
        Value: 100,
        Unit:  units.MassFluxKilogramPerSecondPerSquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MassFluxDto{
        Value: math.NaN(),
        Unit:  units.MassFluxKilogramPerSecondPerSquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test GramPerSecondPerSquareMeter conversion
    grams_per_second_per_square_meterDto := units.MassFluxDto{
        Value: 100,
        Unit:  units.MassFluxGramPerSecondPerSquareMeter,
    }
    
    var grams_per_second_per_square_meterResult *units.MassFlux
    grams_per_second_per_square_meterResult, err = factory.FromDto(grams_per_second_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerSecondPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_second_per_square_meterResult.Convert(units.MassFluxGramPerSecondPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerSecondPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test GramPerSecondPerSquareCentimeter conversion
    grams_per_second_per_square_centimeterDto := units.MassFluxDto{
        Value: 100,
        Unit:  units.MassFluxGramPerSecondPerSquareCentimeter,
    }
    
    var grams_per_second_per_square_centimeterResult *units.MassFlux
    grams_per_second_per_square_centimeterResult, err = factory.FromDto(grams_per_second_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerSecondPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_second_per_square_centimeterResult.Convert(units.MassFluxGramPerSecondPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerSecondPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test GramPerSecondPerSquareMillimeter conversion
    grams_per_second_per_square_millimeterDto := units.MassFluxDto{
        Value: 100,
        Unit:  units.MassFluxGramPerSecondPerSquareMillimeter,
    }
    
    var grams_per_second_per_square_millimeterResult *units.MassFlux
    grams_per_second_per_square_millimeterResult, err = factory.FromDto(grams_per_second_per_square_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerSecondPerSquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_second_per_square_millimeterResult.Convert(units.MassFluxGramPerSecondPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerSecondPerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test GramPerHourPerSquareMeter conversion
    grams_per_hour_per_square_meterDto := units.MassFluxDto{
        Value: 100,
        Unit:  units.MassFluxGramPerHourPerSquareMeter,
    }
    
    var grams_per_hour_per_square_meterResult *units.MassFlux
    grams_per_hour_per_square_meterResult, err = factory.FromDto(grams_per_hour_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerHourPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_hour_per_square_meterResult.Convert(units.MassFluxGramPerHourPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerHourPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test GramPerHourPerSquareCentimeter conversion
    grams_per_hour_per_square_centimeterDto := units.MassFluxDto{
        Value: 100,
        Unit:  units.MassFluxGramPerHourPerSquareCentimeter,
    }
    
    var grams_per_hour_per_square_centimeterResult *units.MassFlux
    grams_per_hour_per_square_centimeterResult, err = factory.FromDto(grams_per_hour_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerHourPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_hour_per_square_centimeterResult.Convert(units.MassFluxGramPerHourPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerHourPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test GramPerHourPerSquareMillimeter conversion
    grams_per_hour_per_square_millimeterDto := units.MassFluxDto{
        Value: 100,
        Unit:  units.MassFluxGramPerHourPerSquareMillimeter,
    }
    
    var grams_per_hour_per_square_millimeterResult *units.MassFlux
    grams_per_hour_per_square_millimeterResult, err = factory.FromDto(grams_per_hour_per_square_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerHourPerSquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_hour_per_square_millimeterResult.Convert(units.MassFluxGramPerHourPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerHourPerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test KilogramPerSecondPerSquareMeter conversion
    kilograms_per_second_per_square_meterDto := units.MassFluxDto{
        Value: 100,
        Unit:  units.MassFluxKilogramPerSecondPerSquareMeter,
    }
    
    var kilograms_per_second_per_square_meterResult *units.MassFlux
    kilograms_per_second_per_square_meterResult, err = factory.FromDto(kilograms_per_second_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerSecondPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_second_per_square_meterResult.Convert(units.MassFluxKilogramPerSecondPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerSecondPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test KilogramPerSecondPerSquareCentimeter conversion
    kilograms_per_second_per_square_centimeterDto := units.MassFluxDto{
        Value: 100,
        Unit:  units.MassFluxKilogramPerSecondPerSquareCentimeter,
    }
    
    var kilograms_per_second_per_square_centimeterResult *units.MassFlux
    kilograms_per_second_per_square_centimeterResult, err = factory.FromDto(kilograms_per_second_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerSecondPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_second_per_square_centimeterResult.Convert(units.MassFluxKilogramPerSecondPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerSecondPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test KilogramPerSecondPerSquareMillimeter conversion
    kilograms_per_second_per_square_millimeterDto := units.MassFluxDto{
        Value: 100,
        Unit:  units.MassFluxKilogramPerSecondPerSquareMillimeter,
    }
    
    var kilograms_per_second_per_square_millimeterResult *units.MassFlux
    kilograms_per_second_per_square_millimeterResult, err = factory.FromDto(kilograms_per_second_per_square_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerSecondPerSquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_second_per_square_millimeterResult.Convert(units.MassFluxKilogramPerSecondPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerSecondPerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test KilogramPerHourPerSquareMeter conversion
    kilograms_per_hour_per_square_meterDto := units.MassFluxDto{
        Value: 100,
        Unit:  units.MassFluxKilogramPerHourPerSquareMeter,
    }
    
    var kilograms_per_hour_per_square_meterResult *units.MassFlux
    kilograms_per_hour_per_square_meterResult, err = factory.FromDto(kilograms_per_hour_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerHourPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_hour_per_square_meterResult.Convert(units.MassFluxKilogramPerHourPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerHourPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test KilogramPerHourPerSquareCentimeter conversion
    kilograms_per_hour_per_square_centimeterDto := units.MassFluxDto{
        Value: 100,
        Unit:  units.MassFluxKilogramPerHourPerSquareCentimeter,
    }
    
    var kilograms_per_hour_per_square_centimeterResult *units.MassFlux
    kilograms_per_hour_per_square_centimeterResult, err = factory.FromDto(kilograms_per_hour_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerHourPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_hour_per_square_centimeterResult.Convert(units.MassFluxKilogramPerHourPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerHourPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test KilogramPerHourPerSquareMillimeter conversion
    kilograms_per_hour_per_square_millimeterDto := units.MassFluxDto{
        Value: 100,
        Unit:  units.MassFluxKilogramPerHourPerSquareMillimeter,
    }
    
    var kilograms_per_hour_per_square_millimeterResult *units.MassFlux
    kilograms_per_hour_per_square_millimeterResult, err = factory.FromDto(kilograms_per_hour_per_square_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerHourPerSquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_hour_per_square_millimeterResult.Convert(units.MassFluxKilogramPerHourPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerHourPerSquareMillimeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MassFluxDto{
        Value: 0,
        Unit:  units.MassFluxKilogramPerSecondPerSquareMeter,
    }
    
    var zeroResult *units.MassFlux
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMassFluxFactory_FromDtoJSON(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "KilogramPerSecondPerSquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "KilogramPerSecondPerSquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.MassFluxDto{
        Value: nanValue,
        Unit:  units.MassFluxKilogramPerSecondPerSquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with GramPerSecondPerSquareMeter unit
    grams_per_second_per_square_meterJSON := []byte(`{"value": 100, "unit": "GramPerSecondPerSquareMeter"}`)
    grams_per_second_per_square_meterResult, err := factory.FromDtoJSON(grams_per_second_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerSecondPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_second_per_square_meterResult.Convert(units.MassFluxGramPerSecondPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerSecondPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerSecondPerSquareCentimeter unit
    grams_per_second_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "GramPerSecondPerSquareCentimeter"}`)
    grams_per_second_per_square_centimeterResult, err := factory.FromDtoJSON(grams_per_second_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerSecondPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_second_per_square_centimeterResult.Convert(units.MassFluxGramPerSecondPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerSecondPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerSecondPerSquareMillimeter unit
    grams_per_second_per_square_millimeterJSON := []byte(`{"value": 100, "unit": "GramPerSecondPerSquareMillimeter"}`)
    grams_per_second_per_square_millimeterResult, err := factory.FromDtoJSON(grams_per_second_per_square_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerSecondPerSquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_second_per_square_millimeterResult.Convert(units.MassFluxGramPerSecondPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerSecondPerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerHourPerSquareMeter unit
    grams_per_hour_per_square_meterJSON := []byte(`{"value": 100, "unit": "GramPerHourPerSquareMeter"}`)
    grams_per_hour_per_square_meterResult, err := factory.FromDtoJSON(grams_per_hour_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerHourPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_hour_per_square_meterResult.Convert(units.MassFluxGramPerHourPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerHourPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerHourPerSquareCentimeter unit
    grams_per_hour_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "GramPerHourPerSquareCentimeter"}`)
    grams_per_hour_per_square_centimeterResult, err := factory.FromDtoJSON(grams_per_hour_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerHourPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_hour_per_square_centimeterResult.Convert(units.MassFluxGramPerHourPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerHourPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerHourPerSquareMillimeter unit
    grams_per_hour_per_square_millimeterJSON := []byte(`{"value": 100, "unit": "GramPerHourPerSquareMillimeter"}`)
    grams_per_hour_per_square_millimeterResult, err := factory.FromDtoJSON(grams_per_hour_per_square_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerHourPerSquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_hour_per_square_millimeterResult.Convert(units.MassFluxGramPerHourPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerHourPerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerSecondPerSquareMeter unit
    kilograms_per_second_per_square_meterJSON := []byte(`{"value": 100, "unit": "KilogramPerSecondPerSquareMeter"}`)
    kilograms_per_second_per_square_meterResult, err := factory.FromDtoJSON(kilograms_per_second_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerSecondPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_second_per_square_meterResult.Convert(units.MassFluxKilogramPerSecondPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerSecondPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerSecondPerSquareCentimeter unit
    kilograms_per_second_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "KilogramPerSecondPerSquareCentimeter"}`)
    kilograms_per_second_per_square_centimeterResult, err := factory.FromDtoJSON(kilograms_per_second_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerSecondPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_second_per_square_centimeterResult.Convert(units.MassFluxKilogramPerSecondPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerSecondPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerSecondPerSquareMillimeter unit
    kilograms_per_second_per_square_millimeterJSON := []byte(`{"value": 100, "unit": "KilogramPerSecondPerSquareMillimeter"}`)
    kilograms_per_second_per_square_millimeterResult, err := factory.FromDtoJSON(kilograms_per_second_per_square_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerSecondPerSquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_second_per_square_millimeterResult.Convert(units.MassFluxKilogramPerSecondPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerSecondPerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerHourPerSquareMeter unit
    kilograms_per_hour_per_square_meterJSON := []byte(`{"value": 100, "unit": "KilogramPerHourPerSquareMeter"}`)
    kilograms_per_hour_per_square_meterResult, err := factory.FromDtoJSON(kilograms_per_hour_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerHourPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_hour_per_square_meterResult.Convert(units.MassFluxKilogramPerHourPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerHourPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerHourPerSquareCentimeter unit
    kilograms_per_hour_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "KilogramPerHourPerSquareCentimeter"}`)
    kilograms_per_hour_per_square_centimeterResult, err := factory.FromDtoJSON(kilograms_per_hour_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerHourPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_hour_per_square_centimeterResult.Convert(units.MassFluxKilogramPerHourPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerHourPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerHourPerSquareMillimeter unit
    kilograms_per_hour_per_square_millimeterJSON := []byte(`{"value": 100, "unit": "KilogramPerHourPerSquareMillimeter"}`)
    kilograms_per_hour_per_square_millimeterResult, err := factory.FromDtoJSON(kilograms_per_hour_per_square_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerHourPerSquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_hour_per_square_millimeterResult.Convert(units.MassFluxKilogramPerHourPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerHourPerSquareMillimeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "KilogramPerSecondPerSquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromGramsPerSecondPerSquareMeter function
func TestMassFluxFactory_FromGramsPerSecondPerSquareMeter(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerSecondPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromGramsPerSecondPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFluxGramPerSecondPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerSecondPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerSecondPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerSecondPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerSecondPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerSecondPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerSecondPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerSecondPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerSecondPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromGramsPerSecondPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFluxGramPerSecondPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerSecondPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerSecondPerSquareCentimeter function
func TestMassFluxFactory_FromGramsPerSecondPerSquareCentimeter(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerSecondPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromGramsPerSecondPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFluxGramPerSecondPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerSecondPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerSecondPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerSecondPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerSecondPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerSecondPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerSecondPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerSecondPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerSecondPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromGramsPerSecondPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFluxGramPerSecondPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerSecondPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerSecondPerSquareMillimeter function
func TestMassFluxFactory_FromGramsPerSecondPerSquareMillimeter(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerSecondPerSquareMillimeter(100)
    if err != nil {
        t.Errorf("FromGramsPerSecondPerSquareMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFluxGramPerSecondPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerSecondPerSquareMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerSecondPerSquareMillimeter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerSecondPerSquareMillimeter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerSecondPerSquareMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerSecondPerSquareMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerSecondPerSquareMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerSecondPerSquareMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerSecondPerSquareMillimeter(0)
    if err != nil {
        t.Errorf("FromGramsPerSecondPerSquareMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFluxGramPerSecondPerSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerSecondPerSquareMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerHourPerSquareMeter function
func TestMassFluxFactory_FromGramsPerHourPerSquareMeter(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerHourPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromGramsPerHourPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFluxGramPerHourPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerHourPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerHourPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerHourPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerHourPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerHourPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerHourPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerHourPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerHourPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromGramsPerHourPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFluxGramPerHourPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerHourPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerHourPerSquareCentimeter function
func TestMassFluxFactory_FromGramsPerHourPerSquareCentimeter(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerHourPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromGramsPerHourPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFluxGramPerHourPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerHourPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerHourPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerHourPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerHourPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerHourPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerHourPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerHourPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerHourPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromGramsPerHourPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFluxGramPerHourPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerHourPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerHourPerSquareMillimeter function
func TestMassFluxFactory_FromGramsPerHourPerSquareMillimeter(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerHourPerSquareMillimeter(100)
    if err != nil {
        t.Errorf("FromGramsPerHourPerSquareMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFluxGramPerHourPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerHourPerSquareMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerHourPerSquareMillimeter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerHourPerSquareMillimeter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerHourPerSquareMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerHourPerSquareMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerHourPerSquareMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerHourPerSquareMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerHourPerSquareMillimeter(0)
    if err != nil {
        t.Errorf("FromGramsPerHourPerSquareMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFluxGramPerHourPerSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerHourPerSquareMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerSecondPerSquareMeter function
func TestMassFluxFactory_FromKilogramsPerSecondPerSquareMeter(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerSecondPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerSecondPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFluxKilogramPerSecondPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerSecondPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerSecondPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerSecondPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerSecondPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerSecondPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerSecondPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerSecondPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerSecondPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerSecondPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFluxKilogramPerSecondPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerSecondPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerSecondPerSquareCentimeter function
func TestMassFluxFactory_FromKilogramsPerSecondPerSquareCentimeter(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerSecondPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerSecondPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFluxKilogramPerSecondPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerSecondPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerSecondPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerSecondPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerSecondPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerSecondPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerSecondPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerSecondPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerSecondPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerSecondPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFluxKilogramPerSecondPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerSecondPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerSecondPerSquareMillimeter function
func TestMassFluxFactory_FromKilogramsPerSecondPerSquareMillimeter(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerSecondPerSquareMillimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerSecondPerSquareMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFluxKilogramPerSecondPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerSecondPerSquareMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerSecondPerSquareMillimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerSecondPerSquareMillimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerSecondPerSquareMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerSecondPerSquareMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerSecondPerSquareMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerSecondPerSquareMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerSecondPerSquareMillimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerSecondPerSquareMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFluxKilogramPerSecondPerSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerSecondPerSquareMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerHourPerSquareMeter function
func TestMassFluxFactory_FromKilogramsPerHourPerSquareMeter(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerHourPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerHourPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFluxKilogramPerHourPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerHourPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerHourPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerHourPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerHourPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerHourPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerHourPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerHourPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerHourPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerHourPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFluxKilogramPerHourPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerHourPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerHourPerSquareCentimeter function
func TestMassFluxFactory_FromKilogramsPerHourPerSquareCentimeter(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerHourPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerHourPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFluxKilogramPerHourPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerHourPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerHourPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerHourPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerHourPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerHourPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerHourPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerHourPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerHourPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerHourPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFluxKilogramPerHourPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerHourPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerHourPerSquareMillimeter function
func TestMassFluxFactory_FromKilogramsPerHourPerSquareMillimeter(t *testing.T) {
    factory := units.MassFluxFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerHourPerSquareMillimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerHourPerSquareMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassFluxKilogramPerHourPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerHourPerSquareMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerHourPerSquareMillimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerHourPerSquareMillimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerHourPerSquareMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerHourPerSquareMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerHourPerSquareMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerHourPerSquareMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerHourPerSquareMillimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerHourPerSquareMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassFluxKilogramPerHourPerSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerHourPerSquareMillimeter() with zero value = %v, want 0", converted)
    }
}

func TestMassFluxToString(t *testing.T) {
	factory := units.MassFluxFactory{}
	a, err := factory.CreateMassFlux(45, units.MassFluxKilogramPerSecondPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MassFluxKilogramPerSecondPerSquareMeter, 2)
	expected := "45.00 " + units.GetMassFluxAbbreviation(units.MassFluxKilogramPerSecondPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MassFluxKilogramPerSecondPerSquareMeter, -1)
	expected = "45 " + units.GetMassFluxAbbreviation(units.MassFluxKilogramPerSecondPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMassFlux_EqualityAndComparison(t *testing.T) {
	factory := units.MassFluxFactory{}
	a1, _ := factory.CreateMassFlux(60, units.MassFluxKilogramPerSecondPerSquareMeter)
	a2, _ := factory.CreateMassFlux(60, units.MassFluxKilogramPerSecondPerSquareMeter)
	a3, _ := factory.CreateMassFlux(120, units.MassFluxKilogramPerSecondPerSquareMeter)

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

func TestMassFlux_Arithmetic(t *testing.T) {
	factory := units.MassFluxFactory{}
	a1, _ := factory.CreateMassFlux(30, units.MassFluxKilogramPerSecondPerSquareMeter)
	a2, _ := factory.CreateMassFlux(45, units.MassFluxKilogramPerSecondPerSquareMeter)

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