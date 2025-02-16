// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLinearDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerMeter"}`
	
	factory := units.LinearDensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LinearDensityKilogramPerMeter {
		t.Errorf("expected unit %v, got %v", units.LinearDensityKilogramPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLinearDensityDto_ToJSON(t *testing.T) {
	dto := units.LinearDensityDto{
		Value: 45,
		Unit:  units.LinearDensityKilogramPerMeter,
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
	if result["unit"].(string) != string(units.LinearDensityKilogramPerMeter) {
		t.Errorf("expected unit %s, got %v", units.LinearDensityKilogramPerMeter, result["unit"])
	}
}

func TestNewLinearDensity_InvalidValue(t *testing.T) {
	factory := units.LinearDensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLinearDensity(math.NaN(), units.LinearDensityKilogramPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLinearDensity(math.Inf(1), units.LinearDensityKilogramPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLinearDensityConversions(t *testing.T) {
	factory := units.LinearDensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLinearDensity(180, units.LinearDensityKilogramPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramsPerMillimeter.
		// No expected conversion value provided for GramsPerMillimeter, verifying result is not NaN.
		result := a.GramsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerCentimeter.
		// No expected conversion value provided for GramsPerCentimeter, verifying result is not NaN.
		result := a.GramsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerMeter.
		// No expected conversion value provided for GramsPerMeter, verifying result is not NaN.
		result := a.GramsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerInch.
		// No expected conversion value provided for PoundsPerInch, verifying result is not NaN.
		result := a.PoundsPerInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerFoot.
		// No expected conversion value provided for PoundsPerFoot, verifying result is not NaN.
		result := a.PoundsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to GramsPerFoot.
		// No expected conversion value provided for GramsPerFoot, verifying result is not NaN.
		result := a.GramsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerMillimeter.
		// No expected conversion value provided for MicrogramsPerMillimeter, verifying result is not NaN.
		result := a.MicrogramsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerMillimeter.
		// No expected conversion value provided for MilligramsPerMillimeter, verifying result is not NaN.
		result := a.MilligramsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerMillimeter.
		// No expected conversion value provided for KilogramsPerMillimeter, verifying result is not NaN.
		result := a.KilogramsPerMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerCentimeter.
		// No expected conversion value provided for MicrogramsPerCentimeter, verifying result is not NaN.
		result := a.MicrogramsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerCentimeter.
		// No expected conversion value provided for MilligramsPerCentimeter, verifying result is not NaN.
		result := a.MilligramsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerCentimeter.
		// No expected conversion value provided for KilogramsPerCentimeter, verifying result is not NaN.
		result := a.KilogramsPerCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerMeter.
		// No expected conversion value provided for MicrogramsPerMeter, verifying result is not NaN.
		result := a.MicrogramsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerMeter.
		// No expected conversion value provided for MilligramsPerMeter, verifying result is not NaN.
		result := a.MilligramsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerMeter.
		// No expected conversion value provided for KilogramsPerMeter, verifying result is not NaN.
		result := a.KilogramsPerMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerFoot.
		// No expected conversion value provided for MicrogramsPerFoot, verifying result is not NaN.
		result := a.MicrogramsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerFoot.
		// No expected conversion value provided for MilligramsPerFoot, verifying result is not NaN.
		result := a.MilligramsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerFoot.
		// No expected conversion value provided for KilogramsPerFoot, verifying result is not NaN.
		result := a.KilogramsPerFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerFoot returned NaN")
		}
	}
}

func TestLinearDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LinearDensityFactory{}
	a, err := factory.CreateLinearDensity(90, units.LinearDensityKilogramPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LinearDensityKilogramPerMeter {
		t.Errorf("expected default unit KilogramPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LinearDensityGramPerMillimeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LinearDensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LinearDensityKilogramPerMeter {
		t.Errorf("expected unit KilogramPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLinearDensityFactory_FromDto(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityKilogramPerMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.LinearDensityDto{
        Value: math.NaN(),
        Unit:  units.LinearDensityKilogramPerMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test GramPerMillimeter conversion
    grams_per_millimeterDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityGramPerMillimeter,
    }
    
    var grams_per_millimeterResult *units.LinearDensity
    grams_per_millimeterResult, err = factory.FromDto(grams_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_millimeterResult.Convert(units.LinearDensityGramPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerMillimeter = %v, want %v", converted, 100)
    }
    // Test GramPerCentimeter conversion
    grams_per_centimeterDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityGramPerCentimeter,
    }
    
    var grams_per_centimeterResult *units.LinearDensity
    grams_per_centimeterResult, err = factory.FromDto(grams_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_centimeterResult.Convert(units.LinearDensityGramPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCentimeter = %v, want %v", converted, 100)
    }
    // Test GramPerMeter conversion
    grams_per_meterDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityGramPerMeter,
    }
    
    var grams_per_meterResult *units.LinearDensity
    grams_per_meterResult, err = factory.FromDto(grams_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_meterResult.Convert(units.LinearDensityGramPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerMeter = %v, want %v", converted, 100)
    }
    // Test PoundPerInch conversion
    pounds_per_inchDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityPoundPerInch,
    }
    
    var pounds_per_inchResult *units.LinearDensity
    pounds_per_inchResult, err = factory.FromDto(pounds_per_inchDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_inchResult.Convert(units.LinearDensityPoundPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerInch = %v, want %v", converted, 100)
    }
    // Test PoundPerFoot conversion
    pounds_per_footDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityPoundPerFoot,
    }
    
    var pounds_per_footResult *units.LinearDensity
    pounds_per_footResult, err = factory.FromDto(pounds_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_footResult.Convert(units.LinearDensityPoundPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerFoot = %v, want %v", converted, 100)
    }
    // Test GramPerFoot conversion
    grams_per_footDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityGramPerFoot,
    }
    
    var grams_per_footResult *units.LinearDensity
    grams_per_footResult, err = factory.FromDto(grams_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_footResult.Convert(units.LinearDensityGramPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerFoot = %v, want %v", converted, 100)
    }
    // Test MicrogramPerMillimeter conversion
    micrograms_per_millimeterDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityMicrogramPerMillimeter,
    }
    
    var micrograms_per_millimeterResult *units.LinearDensity
    micrograms_per_millimeterResult, err = factory.FromDto(micrograms_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_millimeterResult.Convert(units.LinearDensityMicrogramPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerMillimeter = %v, want %v", converted, 100)
    }
    // Test MilligramPerMillimeter conversion
    milligrams_per_millimeterDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityMilligramPerMillimeter,
    }
    
    var milligrams_per_millimeterResult *units.LinearDensity
    milligrams_per_millimeterResult, err = factory.FromDto(milligrams_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_millimeterResult.Convert(units.LinearDensityMilligramPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerMillimeter = %v, want %v", converted, 100)
    }
    // Test KilogramPerMillimeter conversion
    kilograms_per_millimeterDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityKilogramPerMillimeter,
    }
    
    var kilograms_per_millimeterResult *units.LinearDensity
    kilograms_per_millimeterResult, err = factory.FromDto(kilograms_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_millimeterResult.Convert(units.LinearDensityKilogramPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerMillimeter = %v, want %v", converted, 100)
    }
    // Test MicrogramPerCentimeter conversion
    micrograms_per_centimeterDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityMicrogramPerCentimeter,
    }
    
    var micrograms_per_centimeterResult *units.LinearDensity
    micrograms_per_centimeterResult, err = factory.FromDto(micrograms_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_centimeterResult.Convert(units.LinearDensityMicrogramPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerCentimeter = %v, want %v", converted, 100)
    }
    // Test MilligramPerCentimeter conversion
    milligrams_per_centimeterDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityMilligramPerCentimeter,
    }
    
    var milligrams_per_centimeterResult *units.LinearDensity
    milligrams_per_centimeterResult, err = factory.FromDto(milligrams_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_centimeterResult.Convert(units.LinearDensityMilligramPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerCentimeter = %v, want %v", converted, 100)
    }
    // Test KilogramPerCentimeter conversion
    kilograms_per_centimeterDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityKilogramPerCentimeter,
    }
    
    var kilograms_per_centimeterResult *units.LinearDensity
    kilograms_per_centimeterResult, err = factory.FromDto(kilograms_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_centimeterResult.Convert(units.LinearDensityKilogramPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerCentimeter = %v, want %v", converted, 100)
    }
    // Test MicrogramPerMeter conversion
    micrograms_per_meterDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityMicrogramPerMeter,
    }
    
    var micrograms_per_meterResult *units.LinearDensity
    micrograms_per_meterResult, err = factory.FromDto(micrograms_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_meterResult.Convert(units.LinearDensityMicrogramPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerMeter = %v, want %v", converted, 100)
    }
    // Test MilligramPerMeter conversion
    milligrams_per_meterDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityMilligramPerMeter,
    }
    
    var milligrams_per_meterResult *units.LinearDensity
    milligrams_per_meterResult, err = factory.FromDto(milligrams_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_meterResult.Convert(units.LinearDensityMilligramPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerMeter = %v, want %v", converted, 100)
    }
    // Test KilogramPerMeter conversion
    kilograms_per_meterDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityKilogramPerMeter,
    }
    
    var kilograms_per_meterResult *units.LinearDensity
    kilograms_per_meterResult, err = factory.FromDto(kilograms_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_meterResult.Convert(units.LinearDensityKilogramPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerMeter = %v, want %v", converted, 100)
    }
    // Test MicrogramPerFoot conversion
    micrograms_per_footDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityMicrogramPerFoot,
    }
    
    var micrograms_per_footResult *units.LinearDensity
    micrograms_per_footResult, err = factory.FromDto(micrograms_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_footResult.Convert(units.LinearDensityMicrogramPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerFoot = %v, want %v", converted, 100)
    }
    // Test MilligramPerFoot conversion
    milligrams_per_footDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityMilligramPerFoot,
    }
    
    var milligrams_per_footResult *units.LinearDensity
    milligrams_per_footResult, err = factory.FromDto(milligrams_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_footResult.Convert(units.LinearDensityMilligramPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerFoot = %v, want %v", converted, 100)
    }
    // Test KilogramPerFoot conversion
    kilograms_per_footDto := units.LinearDensityDto{
        Value: 100,
        Unit:  units.LinearDensityKilogramPerFoot,
    }
    
    var kilograms_per_footResult *units.LinearDensity
    kilograms_per_footResult, err = factory.FromDto(kilograms_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_footResult.Convert(units.LinearDensityKilogramPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerFoot = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.LinearDensityDto{
        Value: 0,
        Unit:  units.LinearDensityKilogramPerMeter,
    }
    
    var zeroResult *units.LinearDensity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestLinearDensityFactory_FromDtoJSON(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "KilogramPerMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "KilogramPerMeter"}`)
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
    nanJSON, _ := json.Marshal(units.LinearDensityDto{
        Value: nanValue,
        Unit:  units.LinearDensityKilogramPerMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with GramPerMillimeter unit
    grams_per_millimeterJSON := []byte(`{"value": 100, "unit": "GramPerMillimeter"}`)
    grams_per_millimeterResult, err := factory.FromDtoJSON(grams_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_millimeterResult.Convert(units.LinearDensityGramPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerCentimeter unit
    grams_per_centimeterJSON := []byte(`{"value": 100, "unit": "GramPerCentimeter"}`)
    grams_per_centimeterResult, err := factory.FromDtoJSON(grams_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_centimeterResult.Convert(units.LinearDensityGramPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerMeter unit
    grams_per_meterJSON := []byte(`{"value": 100, "unit": "GramPerMeter"}`)
    grams_per_meterResult, err := factory.FromDtoJSON(grams_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_meterResult.Convert(units.LinearDensityGramPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerInch unit
    pounds_per_inchJSON := []byte(`{"value": 100, "unit": "PoundPerInch"}`)
    pounds_per_inchResult, err := factory.FromDtoJSON(pounds_per_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_inchResult.Convert(units.LinearDensityPoundPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerInch = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerFoot unit
    pounds_per_footJSON := []byte(`{"value": 100, "unit": "PoundPerFoot"}`)
    pounds_per_footResult, err := factory.FromDtoJSON(pounds_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_footResult.Convert(units.LinearDensityPoundPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerFoot unit
    grams_per_footJSON := []byte(`{"value": 100, "unit": "GramPerFoot"}`)
    grams_per_footResult, err := factory.FromDtoJSON(grams_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_footResult.Convert(units.LinearDensityGramPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerMillimeter unit
    micrograms_per_millimeterJSON := []byte(`{"value": 100, "unit": "MicrogramPerMillimeter"}`)
    micrograms_per_millimeterResult, err := factory.FromDtoJSON(micrograms_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_millimeterResult.Convert(units.LinearDensityMicrogramPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerMillimeter unit
    milligrams_per_millimeterJSON := []byte(`{"value": 100, "unit": "MilligramPerMillimeter"}`)
    milligrams_per_millimeterResult, err := factory.FromDtoJSON(milligrams_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_millimeterResult.Convert(units.LinearDensityMilligramPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerMillimeter unit
    kilograms_per_millimeterJSON := []byte(`{"value": 100, "unit": "KilogramPerMillimeter"}`)
    kilograms_per_millimeterResult, err := factory.FromDtoJSON(kilograms_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_millimeterResult.Convert(units.LinearDensityKilogramPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerCentimeter unit
    micrograms_per_centimeterJSON := []byte(`{"value": 100, "unit": "MicrogramPerCentimeter"}`)
    micrograms_per_centimeterResult, err := factory.FromDtoJSON(micrograms_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_centimeterResult.Convert(units.LinearDensityMicrogramPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerCentimeter unit
    milligrams_per_centimeterJSON := []byte(`{"value": 100, "unit": "MilligramPerCentimeter"}`)
    milligrams_per_centimeterResult, err := factory.FromDtoJSON(milligrams_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_centimeterResult.Convert(units.LinearDensityMilligramPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerCentimeter unit
    kilograms_per_centimeterJSON := []byte(`{"value": 100, "unit": "KilogramPerCentimeter"}`)
    kilograms_per_centimeterResult, err := factory.FromDtoJSON(kilograms_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_centimeterResult.Convert(units.LinearDensityKilogramPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerMeter unit
    micrograms_per_meterJSON := []byte(`{"value": 100, "unit": "MicrogramPerMeter"}`)
    micrograms_per_meterResult, err := factory.FromDtoJSON(micrograms_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_meterResult.Convert(units.LinearDensityMicrogramPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerMeter unit
    milligrams_per_meterJSON := []byte(`{"value": 100, "unit": "MilligramPerMeter"}`)
    milligrams_per_meterResult, err := factory.FromDtoJSON(milligrams_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_meterResult.Convert(units.LinearDensityMilligramPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerMeter unit
    kilograms_per_meterJSON := []byte(`{"value": 100, "unit": "KilogramPerMeter"}`)
    kilograms_per_meterResult, err := factory.FromDtoJSON(kilograms_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_meterResult.Convert(units.LinearDensityKilogramPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerFoot unit
    micrograms_per_footJSON := []byte(`{"value": 100, "unit": "MicrogramPerFoot"}`)
    micrograms_per_footResult, err := factory.FromDtoJSON(micrograms_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_footResult.Convert(units.LinearDensityMicrogramPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerFoot unit
    milligrams_per_footJSON := []byte(`{"value": 100, "unit": "MilligramPerFoot"}`)
    milligrams_per_footResult, err := factory.FromDtoJSON(milligrams_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_footResult.Convert(units.LinearDensityMilligramPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerFoot unit
    kilograms_per_footJSON := []byte(`{"value": 100, "unit": "KilogramPerFoot"}`)
    kilograms_per_footResult, err := factory.FromDtoJSON(kilograms_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_footResult.Convert(units.LinearDensityKilogramPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerFoot = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "KilogramPerMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromGramsPerMillimeter function
func TestLinearDensityFactory_FromGramsPerMillimeter(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromGramsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityGramPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromGramsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityGramPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerCentimeter function
func TestLinearDensityFactory_FromGramsPerCentimeter(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromGramsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityGramPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromGramsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityGramPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerMeter function
func TestLinearDensityFactory_FromGramsPerMeter(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerMeter(100)
    if err != nil {
        t.Errorf("FromGramsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityGramPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerMeter(0)
    if err != nil {
        t.Errorf("FromGramsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityGramPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerInch function
func TestLinearDensityFactory_FromPoundsPerInch(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerInch(100)
    if err != nil {
        t.Errorf("FromPoundsPerInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityPoundPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerInch(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerInch() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerInch(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerInch() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerInch(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerInch(0)
    if err != nil {
        t.Errorf("FromPoundsPerInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityPoundPerInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerInch() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerFoot function
func TestLinearDensityFactory_FromPoundsPerFoot(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerFoot(100)
    if err != nil {
        t.Errorf("FromPoundsPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityPoundPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerFoot(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerFoot() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerFoot(0)
    if err != nil {
        t.Errorf("FromPoundsPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityPoundPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerFoot function
func TestLinearDensityFactory_FromGramsPerFoot(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerFoot(100)
    if err != nil {
        t.Errorf("FromGramsPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityGramPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerFoot(math.NaN())
    if err == nil {
        t.Error("FromGramsPerFoot() with NaN value should return error")
    }

    _, err = factory.FromGramsPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerFoot(0)
    if err != nil {
        t.Errorf("FromGramsPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityGramPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerMillimeter function
func TestLinearDensityFactory_FromMicrogramsPerMillimeter(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityMicrogramPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityMicrogramPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerMillimeter function
func TestLinearDensityFactory_FromMilligramsPerMillimeter(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromMilligramsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityMilligramPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromMilligramsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityMilligramPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerMillimeter function
func TestLinearDensityFactory_FromKilogramsPerMillimeter(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityKilogramPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityKilogramPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerCentimeter function
func TestLinearDensityFactory_FromMicrogramsPerCentimeter(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityMicrogramPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityMicrogramPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerCentimeter function
func TestLinearDensityFactory_FromMilligramsPerCentimeter(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromMilligramsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityMilligramPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromMilligramsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityMilligramPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerCentimeter function
func TestLinearDensityFactory_FromKilogramsPerCentimeter(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityKilogramPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityKilogramPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerMeter function
func TestLinearDensityFactory_FromMicrogramsPerMeter(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerMeter(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityMicrogramPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerMeter(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityMicrogramPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerMeter function
func TestLinearDensityFactory_FromMilligramsPerMeter(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerMeter(100)
    if err != nil {
        t.Errorf("FromMilligramsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityMilligramPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerMeter(0)
    if err != nil {
        t.Errorf("FromMilligramsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityMilligramPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerMeter function
func TestLinearDensityFactory_FromKilogramsPerMeter(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerMeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityKilogramPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerMeter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityKilogramPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerFoot function
func TestLinearDensityFactory_FromMicrogramsPerFoot(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerFoot(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityMicrogramPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerFoot(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerFoot() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerFoot(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityMicrogramPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerFoot function
func TestLinearDensityFactory_FromMilligramsPerFoot(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerFoot(100)
    if err != nil {
        t.Errorf("FromMilligramsPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityMilligramPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerFoot(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerFoot() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerFoot(0)
    if err != nil {
        t.Errorf("FromMilligramsPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityMilligramPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerFoot function
func TestLinearDensityFactory_FromKilogramsPerFoot(t *testing.T) {
    factory := units.LinearDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerFoot(100)
    if err != nil {
        t.Errorf("FromKilogramsPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearDensityKilogramPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerFoot(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerFoot() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerFoot(0)
    if err != nil {
        t.Errorf("FromKilogramsPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearDensityKilogramPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerFoot() with zero value = %v, want 0", converted)
    }
}

func TestLinearDensityToString(t *testing.T) {
	factory := units.LinearDensityFactory{}
	a, err := factory.CreateLinearDensity(45, units.LinearDensityKilogramPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LinearDensityKilogramPerMeter, 2)
	expected := "45.00 " + units.GetLinearDensityAbbreviation(units.LinearDensityKilogramPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LinearDensityKilogramPerMeter, -1)
	expected = "45 " + units.GetLinearDensityAbbreviation(units.LinearDensityKilogramPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLinearDensity_EqualityAndComparison(t *testing.T) {
	factory := units.LinearDensityFactory{}
	a1, _ := factory.CreateLinearDensity(60, units.LinearDensityKilogramPerMeter)
	a2, _ := factory.CreateLinearDensity(60, units.LinearDensityKilogramPerMeter)
	a3, _ := factory.CreateLinearDensity(120, units.LinearDensityKilogramPerMeter)

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

func TestLinearDensity_Arithmetic(t *testing.T) {
	factory := units.LinearDensityFactory{}
	a1, _ := factory.CreateLinearDensity(30, units.LinearDensityKilogramPerMeter)
	a2, _ := factory.CreateLinearDensity(45, units.LinearDensityKilogramPerMeter)

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