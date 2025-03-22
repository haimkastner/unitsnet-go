// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMassConcentrationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerCubicMeter"}`
	
	factory := units.MassConcentrationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MassConcentrationKilogramPerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.MassConcentrationKilogramPerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMassConcentrationDto_ToJSON(t *testing.T) {
	dto := units.MassConcentrationDto{
		Value: 45,
		Unit:  units.MassConcentrationKilogramPerCubicMeter,
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
	if result["unit"].(string) != string(units.MassConcentrationKilogramPerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.MassConcentrationKilogramPerCubicMeter, result["unit"])
	}
}

func TestNewMassConcentration_InvalidValue(t *testing.T) {
	factory := units.MassConcentrationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMassConcentration(math.NaN(), units.MassConcentrationKilogramPerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMassConcentration(math.Inf(1), units.MassConcentrationKilogramPerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMassConcentrationConversions(t *testing.T) {
	factory := units.MassConcentrationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMassConcentration(180, units.MassConcentrationKilogramPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramsPerCubicMillimeter.
		// No expected conversion value provided for GramsPerCubicMillimeter, verifying result is not NaN.
		result := a.GramsPerCubicMillimeter()
		cacheResult := a.GramsPerCubicMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerCubicCentimeter.
		// No expected conversion value provided for GramsPerCubicCentimeter, verifying result is not NaN.
		result := a.GramsPerCubicCentimeter()
		cacheResult := a.GramsPerCubicCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerCubicMeter.
		// No expected conversion value provided for GramsPerCubicMeter, verifying result is not NaN.
		result := a.GramsPerCubicMeter()
		cacheResult := a.GramsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerMicroliter.
		// No expected conversion value provided for GramsPerMicroliter, verifying result is not NaN.
		result := a.GramsPerMicroliter()
		cacheResult := a.GramsPerMicroliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramsPerMicroliter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerMilliliter.
		// No expected conversion value provided for GramsPerMilliliter, verifying result is not NaN.
		result := a.GramsPerMilliliter()
		cacheResult := a.GramsPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerDeciliter.
		// No expected conversion value provided for GramsPerDeciliter, verifying result is not NaN.
		result := a.GramsPerDeciliter()
		cacheResult := a.GramsPerDeciliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramsPerDeciliter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerLiter.
		// No expected conversion value provided for GramsPerLiter, verifying result is not NaN.
		result := a.GramsPerLiter()
		cacheResult := a.GramsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerCubicMillimeter.
		// No expected conversion value provided for TonnesPerCubicMillimeter, verifying result is not NaN.
		result := a.TonnesPerCubicMillimeter()
		cacheResult := a.TonnesPerCubicMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonnesPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerCubicCentimeter.
		// No expected conversion value provided for TonnesPerCubicCentimeter, verifying result is not NaN.
		result := a.TonnesPerCubicCentimeter()
		cacheResult := a.TonnesPerCubicCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonnesPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerCubicMeter.
		// No expected conversion value provided for TonnesPerCubicMeter, verifying result is not NaN.
		result := a.TonnesPerCubicMeter()
		cacheResult := a.TonnesPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonnesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicInch.
		// No expected conversion value provided for PoundsPerCubicInch, verifying result is not NaN.
		result := a.PoundsPerCubicInch()
		cacheResult := a.PoundsPerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicFoot.
		// No expected conversion value provided for PoundsPerCubicFoot, verifying result is not NaN.
		result := a.PoundsPerCubicFoot()
		cacheResult := a.PoundsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to SlugsPerCubicFoot.
		// No expected conversion value provided for SlugsPerCubicFoot, verifying result is not NaN.
		result := a.SlugsPerCubicFoot()
		cacheResult := a.SlugsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SlugsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerUSGallon.
		// No expected conversion value provided for PoundsPerUSGallon, verifying result is not NaN.
		result := a.PoundsPerUSGallon()
		cacheResult := a.PoundsPerUSGallon()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsPerUSGallon returned NaN")
		}
	}
	{
		// Test conversion to OuncesPerUSGallon.
		// No expected conversion value provided for OuncesPerUSGallon, verifying result is not NaN.
		result := a.OuncesPerUSGallon()
		cacheResult := a.OuncesPerUSGallon()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to OuncesPerUSGallon returned NaN")
		}
	}
	{
		// Test conversion to OuncesPerImperialGallon.
		// No expected conversion value provided for OuncesPerImperialGallon, verifying result is not NaN.
		result := a.OuncesPerImperialGallon()
		cacheResult := a.OuncesPerImperialGallon()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to OuncesPerImperialGallon returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerImperialGallon.
		// No expected conversion value provided for PoundsPerImperialGallon, verifying result is not NaN.
		result := a.PoundsPerImperialGallon()
		cacheResult := a.PoundsPerImperialGallon()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsPerImperialGallon returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerCubicMillimeter.
		// No expected conversion value provided for KilogramsPerCubicMillimeter, verifying result is not NaN.
		result := a.KilogramsPerCubicMillimeter()
		cacheResult := a.KilogramsPerCubicMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerCubicCentimeter.
		// No expected conversion value provided for KilogramsPerCubicCentimeter, verifying result is not NaN.
		result := a.KilogramsPerCubicCentimeter()
		cacheResult := a.KilogramsPerCubicCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerCubicMeter.
		// No expected conversion value provided for KilogramsPerCubicMeter, verifying result is not NaN.
		result := a.KilogramsPerCubicMeter()
		cacheResult := a.KilogramsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerCubicMeter.
		// No expected conversion value provided for MilligramsPerCubicMeter, verifying result is not NaN.
		result := a.MilligramsPerCubicMeter()
		cacheResult := a.MilligramsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligramsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerCubicMeter.
		// No expected conversion value provided for MicrogramsPerCubicMeter, verifying result is not NaN.
		result := a.MicrogramsPerCubicMeter()
		cacheResult := a.MicrogramsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrogramsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PicogramsPerMicroliter.
		// No expected conversion value provided for PicogramsPerMicroliter, verifying result is not NaN.
		result := a.PicogramsPerMicroliter()
		cacheResult := a.PicogramsPerMicroliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicogramsPerMicroliter returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerMicroliter.
		// No expected conversion value provided for NanogramsPerMicroliter, verifying result is not NaN.
		result := a.NanogramsPerMicroliter()
		cacheResult := a.NanogramsPerMicroliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanogramsPerMicroliter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerMicroliter.
		// No expected conversion value provided for MicrogramsPerMicroliter, verifying result is not NaN.
		result := a.MicrogramsPerMicroliter()
		cacheResult := a.MicrogramsPerMicroliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrogramsPerMicroliter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerMicroliter.
		// No expected conversion value provided for MilligramsPerMicroliter, verifying result is not NaN.
		result := a.MilligramsPerMicroliter()
		cacheResult := a.MilligramsPerMicroliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligramsPerMicroliter returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerMicroliter.
		// No expected conversion value provided for CentigramsPerMicroliter, verifying result is not NaN.
		result := a.CentigramsPerMicroliter()
		cacheResult := a.CentigramsPerMicroliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentigramsPerMicroliter returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerMicroliter.
		// No expected conversion value provided for DecigramsPerMicroliter, verifying result is not NaN.
		result := a.DecigramsPerMicroliter()
		cacheResult := a.DecigramsPerMicroliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecigramsPerMicroliter returned NaN")
		}
	}
	{
		// Test conversion to PicogramsPerMilliliter.
		// No expected conversion value provided for PicogramsPerMilliliter, verifying result is not NaN.
		result := a.PicogramsPerMilliliter()
		cacheResult := a.PicogramsPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicogramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerMilliliter.
		// No expected conversion value provided for NanogramsPerMilliliter, verifying result is not NaN.
		result := a.NanogramsPerMilliliter()
		cacheResult := a.NanogramsPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanogramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerMilliliter.
		// No expected conversion value provided for MicrogramsPerMilliliter, verifying result is not NaN.
		result := a.MicrogramsPerMilliliter()
		cacheResult := a.MicrogramsPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrogramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerMilliliter.
		// No expected conversion value provided for MilligramsPerMilliliter, verifying result is not NaN.
		result := a.MilligramsPerMilliliter()
		cacheResult := a.MilligramsPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerMilliliter.
		// No expected conversion value provided for CentigramsPerMilliliter, verifying result is not NaN.
		result := a.CentigramsPerMilliliter()
		cacheResult := a.CentigramsPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentigramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerMilliliter.
		// No expected conversion value provided for DecigramsPerMilliliter, verifying result is not NaN.
		result := a.DecigramsPerMilliliter()
		cacheResult := a.DecigramsPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecigramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to PicogramsPerDeciliter.
		// No expected conversion value provided for PicogramsPerDeciliter, verifying result is not NaN.
		result := a.PicogramsPerDeciliter()
		cacheResult := a.PicogramsPerDeciliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicogramsPerDeciliter returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerDeciliter.
		// No expected conversion value provided for NanogramsPerDeciliter, verifying result is not NaN.
		result := a.NanogramsPerDeciliter()
		cacheResult := a.NanogramsPerDeciliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanogramsPerDeciliter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerDeciliter.
		// No expected conversion value provided for MicrogramsPerDeciliter, verifying result is not NaN.
		result := a.MicrogramsPerDeciliter()
		cacheResult := a.MicrogramsPerDeciliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrogramsPerDeciliter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerDeciliter.
		// No expected conversion value provided for MilligramsPerDeciliter, verifying result is not NaN.
		result := a.MilligramsPerDeciliter()
		cacheResult := a.MilligramsPerDeciliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligramsPerDeciliter returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerDeciliter.
		// No expected conversion value provided for CentigramsPerDeciliter, verifying result is not NaN.
		result := a.CentigramsPerDeciliter()
		cacheResult := a.CentigramsPerDeciliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentigramsPerDeciliter returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerDeciliter.
		// No expected conversion value provided for DecigramsPerDeciliter, verifying result is not NaN.
		result := a.DecigramsPerDeciliter()
		cacheResult := a.DecigramsPerDeciliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecigramsPerDeciliter returned NaN")
		}
	}
	{
		// Test conversion to PicogramsPerLiter.
		// No expected conversion value provided for PicogramsPerLiter, verifying result is not NaN.
		result := a.PicogramsPerLiter()
		cacheResult := a.PicogramsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerLiter.
		// No expected conversion value provided for NanogramsPerLiter, verifying result is not NaN.
		result := a.NanogramsPerLiter()
		cacheResult := a.NanogramsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerLiter.
		// No expected conversion value provided for MicrogramsPerLiter, verifying result is not NaN.
		result := a.MicrogramsPerLiter()
		cacheResult := a.MicrogramsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerLiter.
		// No expected conversion value provided for MilligramsPerLiter, verifying result is not NaN.
		result := a.MilligramsPerLiter()
		cacheResult := a.MilligramsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerLiter.
		// No expected conversion value provided for CentigramsPerLiter, verifying result is not NaN.
		result := a.CentigramsPerLiter()
		cacheResult := a.CentigramsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentigramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerLiter.
		// No expected conversion value provided for DecigramsPerLiter, verifying result is not NaN.
		result := a.DecigramsPerLiter()
		cacheResult := a.DecigramsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecigramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerLiter.
		// No expected conversion value provided for KilogramsPerLiter, verifying result is not NaN.
		result := a.KilogramsPerLiter()
		cacheResult := a.KilogramsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsPerCubicInch.
		// No expected conversion value provided for KilopoundsPerCubicInch, verifying result is not NaN.
		result := a.KilopoundsPerCubicInch()
		cacheResult := a.KilopoundsPerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsPerCubicFoot.
		// No expected conversion value provided for KilopoundsPerCubicFoot, verifying result is not NaN.
		result := a.KilopoundsPerCubicFoot()
		cacheResult := a.KilopoundsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsPerCubicFoot returned NaN")
		}
	}
}

func TestMassConcentration_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MassConcentrationFactory{}
	a, err := factory.CreateMassConcentration(90, units.MassConcentrationKilogramPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MassConcentrationKilogramPerCubicMeter {
		t.Errorf("expected default unit KilogramPerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MassConcentrationGramPerCubicMillimeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MassConcentrationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MassConcentrationKilogramPerCubicMeter {
		t.Errorf("expected unit KilogramPerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMassConcentrationFactory_FromDto(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationKilogramPerCubicMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MassConcentrationDto{
        Value: math.NaN(),
        Unit:  units.MassConcentrationKilogramPerCubicMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test GramPerCubicMillimeter conversion
    grams_per_cubic_millimeterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationGramPerCubicMillimeter,
    }
    
    var grams_per_cubic_millimeterResult *units.MassConcentration
    grams_per_cubic_millimeterResult, err = factory.FromDto(grams_per_cubic_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerCubicMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_cubic_millimeterResult.Convert(units.MassConcentrationGramPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test GramPerCubicCentimeter conversion
    grams_per_cubic_centimeterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationGramPerCubicCentimeter,
    }
    
    var grams_per_cubic_centimeterResult *units.MassConcentration
    grams_per_cubic_centimeterResult, err = factory.FromDto(grams_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_cubic_centimeterResult.Convert(units.MassConcentrationGramPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test GramPerCubicMeter conversion
    grams_per_cubic_meterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationGramPerCubicMeter,
    }
    
    var grams_per_cubic_meterResult *units.MassConcentration
    grams_per_cubic_meterResult, err = factory.FromDto(grams_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_cubic_meterResult.Convert(units.MassConcentrationGramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test GramPerMicroliter conversion
    grams_per_microliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationGramPerMicroliter,
    }
    
    var grams_per_microliterResult *units.MassConcentration
    grams_per_microliterResult, err = factory.FromDto(grams_per_microliterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerMicroliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_microliterResult.Convert(units.MassConcentrationGramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test GramPerMilliliter conversion
    grams_per_milliliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationGramPerMilliliter,
    }
    
    var grams_per_milliliterResult *units.MassConcentration
    grams_per_milliliterResult, err = factory.FromDto(grams_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_milliliterResult.Convert(units.MassConcentrationGramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test GramPerDeciliter conversion
    grams_per_deciliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationGramPerDeciliter,
    }
    
    var grams_per_deciliterResult *units.MassConcentration
    grams_per_deciliterResult, err = factory.FromDto(grams_per_deciliterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_deciliterResult.Convert(units.MassConcentrationGramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test GramPerLiter conversion
    grams_per_literDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationGramPerLiter,
    }
    
    var grams_per_literResult *units.MassConcentration
    grams_per_literResult, err = factory.FromDto(grams_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_literResult.Convert(units.MassConcentrationGramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerLiter = %v, want %v", converted, 100)
    }
    // Test TonnePerCubicMillimeter conversion
    tonnes_per_cubic_millimeterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationTonnePerCubicMillimeter,
    }
    
    var tonnes_per_cubic_millimeterResult *units.MassConcentration
    tonnes_per_cubic_millimeterResult, err = factory.FromDto(tonnes_per_cubic_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with TonnePerCubicMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_millimeterResult.Convert(units.MassConcentrationTonnePerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test TonnePerCubicCentimeter conversion
    tonnes_per_cubic_centimeterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationTonnePerCubicCentimeter,
    }
    
    var tonnes_per_cubic_centimeterResult *units.MassConcentration
    tonnes_per_cubic_centimeterResult, err = factory.FromDto(tonnes_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with TonnePerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_centimeterResult.Convert(units.MassConcentrationTonnePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test TonnePerCubicMeter conversion
    tonnes_per_cubic_meterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationTonnePerCubicMeter,
    }
    
    var tonnes_per_cubic_meterResult *units.MassConcentration
    tonnes_per_cubic_meterResult, err = factory.FromDto(tonnes_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with TonnePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_meterResult.Convert(units.MassConcentrationTonnePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test PoundPerCubicInch conversion
    pounds_per_cubic_inchDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationPoundPerCubicInch,
    }
    
    var pounds_per_cubic_inchResult *units.MassConcentration
    pounds_per_cubic_inchResult, err = factory.FromDto(pounds_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_inchResult.Convert(units.MassConcentrationPoundPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicInch = %v, want %v", converted, 100)
    }
    // Test PoundPerCubicFoot conversion
    pounds_per_cubic_footDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationPoundPerCubicFoot,
    }
    
    var pounds_per_cubic_footResult *units.MassConcentration
    pounds_per_cubic_footResult, err = factory.FromDto(pounds_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_footResult.Convert(units.MassConcentrationPoundPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test SlugPerCubicFoot conversion
    slugs_per_cubic_footDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationSlugPerCubicFoot,
    }
    
    var slugs_per_cubic_footResult *units.MassConcentration
    slugs_per_cubic_footResult, err = factory.FromDto(slugs_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with SlugPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugs_per_cubic_footResult.Convert(units.MassConcentrationSlugPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test PoundPerUSGallon conversion
    pounds_per_us_gallonDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationPoundPerUSGallon,
    }
    
    var pounds_per_us_gallonResult *units.MassConcentration
    pounds_per_us_gallonResult, err = factory.FromDto(pounds_per_us_gallonDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerUSGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_us_gallonResult.Convert(units.MassConcentrationPoundPerUSGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerUSGallon = %v, want %v", converted, 100)
    }
    // Test OuncePerUSGallon conversion
    ounces_per_us_gallonDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationOuncePerUSGallon,
    }
    
    var ounces_per_us_gallonResult *units.MassConcentration
    ounces_per_us_gallonResult, err = factory.FromDto(ounces_per_us_gallonDto)
    if err != nil {
        t.Errorf("FromDto() with OuncePerUSGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ounces_per_us_gallonResult.Convert(units.MassConcentrationOuncePerUSGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OuncePerUSGallon = %v, want %v", converted, 100)
    }
    // Test OuncePerImperialGallon conversion
    ounces_per_imperial_gallonDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationOuncePerImperialGallon,
    }
    
    var ounces_per_imperial_gallonResult *units.MassConcentration
    ounces_per_imperial_gallonResult, err = factory.FromDto(ounces_per_imperial_gallonDto)
    if err != nil {
        t.Errorf("FromDto() with OuncePerImperialGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ounces_per_imperial_gallonResult.Convert(units.MassConcentrationOuncePerImperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OuncePerImperialGallon = %v, want %v", converted, 100)
    }
    // Test PoundPerImperialGallon conversion
    pounds_per_imperial_gallonDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationPoundPerImperialGallon,
    }
    
    var pounds_per_imperial_gallonResult *units.MassConcentration
    pounds_per_imperial_gallonResult, err = factory.FromDto(pounds_per_imperial_gallonDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerImperialGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_imperial_gallonResult.Convert(units.MassConcentrationPoundPerImperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerImperialGallon = %v, want %v", converted, 100)
    }
    // Test KilogramPerCubicMillimeter conversion
    kilograms_per_cubic_millimeterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationKilogramPerCubicMillimeter,
    }
    
    var kilograms_per_cubic_millimeterResult *units.MassConcentration
    kilograms_per_cubic_millimeterResult, err = factory.FromDto(kilograms_per_cubic_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerCubicMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_cubic_millimeterResult.Convert(units.MassConcentrationKilogramPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test KilogramPerCubicCentimeter conversion
    kilograms_per_cubic_centimeterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationKilogramPerCubicCentimeter,
    }
    
    var kilograms_per_cubic_centimeterResult *units.MassConcentration
    kilograms_per_cubic_centimeterResult, err = factory.FromDto(kilograms_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_cubic_centimeterResult.Convert(units.MassConcentrationKilogramPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test KilogramPerCubicMeter conversion
    kilograms_per_cubic_meterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationKilogramPerCubicMeter,
    }
    
    var kilograms_per_cubic_meterResult *units.MassConcentration
    kilograms_per_cubic_meterResult, err = factory.FromDto(kilograms_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_cubic_meterResult.Convert(units.MassConcentrationKilogramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test MilligramPerCubicMeter conversion
    milligrams_per_cubic_meterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationMilligramPerCubicMeter,
    }
    
    var milligrams_per_cubic_meterResult *units.MassConcentration
    milligrams_per_cubic_meterResult, err = factory.FromDto(milligrams_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_cubic_meterResult.Convert(units.MassConcentrationMilligramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test MicrogramPerCubicMeter conversion
    micrograms_per_cubic_meterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationMicrogramPerCubicMeter,
    }
    
    var micrograms_per_cubic_meterResult *units.MassConcentration
    micrograms_per_cubic_meterResult, err = factory.FromDto(micrograms_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_cubic_meterResult.Convert(units.MassConcentrationMicrogramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test PicogramPerMicroliter conversion
    picograms_per_microliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationPicogramPerMicroliter,
    }
    
    var picograms_per_microliterResult *units.MassConcentration
    picograms_per_microliterResult, err = factory.FromDto(picograms_per_microliterDto)
    if err != nil {
        t.Errorf("FromDto() with PicogramPerMicroliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_microliterResult.Convert(units.MassConcentrationPicogramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicogramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test NanogramPerMicroliter conversion
    nanograms_per_microliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationNanogramPerMicroliter,
    }
    
    var nanograms_per_microliterResult *units.MassConcentration
    nanograms_per_microliterResult, err = factory.FromDto(nanograms_per_microliterDto)
    if err != nil {
        t.Errorf("FromDto() with NanogramPerMicroliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_microliterResult.Convert(units.MassConcentrationNanogramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test MicrogramPerMicroliter conversion
    micrograms_per_microliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationMicrogramPerMicroliter,
    }
    
    var micrograms_per_microliterResult *units.MassConcentration
    micrograms_per_microliterResult, err = factory.FromDto(micrograms_per_microliterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerMicroliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_microliterResult.Convert(units.MassConcentrationMicrogramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test MilligramPerMicroliter conversion
    milligrams_per_microliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationMilligramPerMicroliter,
    }
    
    var milligrams_per_microliterResult *units.MassConcentration
    milligrams_per_microliterResult, err = factory.FromDto(milligrams_per_microliterDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerMicroliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_microliterResult.Convert(units.MassConcentrationMilligramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test CentigramPerMicroliter conversion
    centigrams_per_microliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationCentigramPerMicroliter,
    }
    
    var centigrams_per_microliterResult *units.MassConcentration
    centigrams_per_microliterResult, err = factory.FromDto(centigrams_per_microliterDto)
    if err != nil {
        t.Errorf("FromDto() with CentigramPerMicroliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_microliterResult.Convert(units.MassConcentrationCentigramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test DecigramPerMicroliter conversion
    decigrams_per_microliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationDecigramPerMicroliter,
    }
    
    var decigrams_per_microliterResult *units.MassConcentration
    decigrams_per_microliterResult, err = factory.FromDto(decigrams_per_microliterDto)
    if err != nil {
        t.Errorf("FromDto() with DecigramPerMicroliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_microliterResult.Convert(units.MassConcentrationDecigramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test PicogramPerMilliliter conversion
    picograms_per_milliliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationPicogramPerMilliliter,
    }
    
    var picograms_per_milliliterResult *units.MassConcentration
    picograms_per_milliliterResult, err = factory.FromDto(picograms_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with PicogramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_milliliterResult.Convert(units.MassConcentrationPicogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicogramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test NanogramPerMilliliter conversion
    nanograms_per_milliliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationNanogramPerMilliliter,
    }
    
    var nanograms_per_milliliterResult *units.MassConcentration
    nanograms_per_milliliterResult, err = factory.FromDto(nanograms_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with NanogramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_milliliterResult.Convert(units.MassConcentrationNanogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test MicrogramPerMilliliter conversion
    micrograms_per_milliliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationMicrogramPerMilliliter,
    }
    
    var micrograms_per_milliliterResult *units.MassConcentration
    micrograms_per_milliliterResult, err = factory.FromDto(micrograms_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_milliliterResult.Convert(units.MassConcentrationMicrogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test MilligramPerMilliliter conversion
    milligrams_per_milliliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationMilligramPerMilliliter,
    }
    
    var milligrams_per_milliliterResult *units.MassConcentration
    milligrams_per_milliliterResult, err = factory.FromDto(milligrams_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_milliliterResult.Convert(units.MassConcentrationMilligramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test CentigramPerMilliliter conversion
    centigrams_per_milliliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationCentigramPerMilliliter,
    }
    
    var centigrams_per_milliliterResult *units.MassConcentration
    centigrams_per_milliliterResult, err = factory.FromDto(centigrams_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with CentigramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_milliliterResult.Convert(units.MassConcentrationCentigramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test DecigramPerMilliliter conversion
    decigrams_per_milliliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationDecigramPerMilliliter,
    }
    
    var decigrams_per_milliliterResult *units.MassConcentration
    decigrams_per_milliliterResult, err = factory.FromDto(decigrams_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with DecigramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_milliliterResult.Convert(units.MassConcentrationDecigramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test PicogramPerDeciliter conversion
    picograms_per_deciliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationPicogramPerDeciliter,
    }
    
    var picograms_per_deciliterResult *units.MassConcentration
    picograms_per_deciliterResult, err = factory.FromDto(picograms_per_deciliterDto)
    if err != nil {
        t.Errorf("FromDto() with PicogramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_deciliterResult.Convert(units.MassConcentrationPicogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test NanogramPerDeciliter conversion
    nanograms_per_deciliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationNanogramPerDeciliter,
    }
    
    var nanograms_per_deciliterResult *units.MassConcentration
    nanograms_per_deciliterResult, err = factory.FromDto(nanograms_per_deciliterDto)
    if err != nil {
        t.Errorf("FromDto() with NanogramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_deciliterResult.Convert(units.MassConcentrationNanogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test MicrogramPerDeciliter conversion
    micrograms_per_deciliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationMicrogramPerDeciliter,
    }
    
    var micrograms_per_deciliterResult *units.MassConcentration
    micrograms_per_deciliterResult, err = factory.FromDto(micrograms_per_deciliterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_deciliterResult.Convert(units.MassConcentrationMicrogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test MilligramPerDeciliter conversion
    milligrams_per_deciliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationMilligramPerDeciliter,
    }
    
    var milligrams_per_deciliterResult *units.MassConcentration
    milligrams_per_deciliterResult, err = factory.FromDto(milligrams_per_deciliterDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_deciliterResult.Convert(units.MassConcentrationMilligramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test CentigramPerDeciliter conversion
    centigrams_per_deciliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationCentigramPerDeciliter,
    }
    
    var centigrams_per_deciliterResult *units.MassConcentration
    centigrams_per_deciliterResult, err = factory.FromDto(centigrams_per_deciliterDto)
    if err != nil {
        t.Errorf("FromDto() with CentigramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_deciliterResult.Convert(units.MassConcentrationCentigramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test DecigramPerDeciliter conversion
    decigrams_per_deciliterDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationDecigramPerDeciliter,
    }
    
    var decigrams_per_deciliterResult *units.MassConcentration
    decigrams_per_deciliterResult, err = factory.FromDto(decigrams_per_deciliterDto)
    if err != nil {
        t.Errorf("FromDto() with DecigramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_deciliterResult.Convert(units.MassConcentrationDecigramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test PicogramPerLiter conversion
    picograms_per_literDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationPicogramPerLiter,
    }
    
    var picograms_per_literResult *units.MassConcentration
    picograms_per_literResult, err = factory.FromDto(picograms_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with PicogramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_literResult.Convert(units.MassConcentrationPicogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicogramPerLiter = %v, want %v", converted, 100)
    }
    // Test NanogramPerLiter conversion
    nanograms_per_literDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationNanogramPerLiter,
    }
    
    var nanograms_per_literResult *units.MassConcentration
    nanograms_per_literResult, err = factory.FromDto(nanograms_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with NanogramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_literResult.Convert(units.MassConcentrationNanogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerLiter = %v, want %v", converted, 100)
    }
    // Test MicrogramPerLiter conversion
    micrograms_per_literDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationMicrogramPerLiter,
    }
    
    var micrograms_per_literResult *units.MassConcentration
    micrograms_per_literResult, err = factory.FromDto(micrograms_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_literResult.Convert(units.MassConcentrationMicrogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerLiter = %v, want %v", converted, 100)
    }
    // Test MilligramPerLiter conversion
    milligrams_per_literDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationMilligramPerLiter,
    }
    
    var milligrams_per_literResult *units.MassConcentration
    milligrams_per_literResult, err = factory.FromDto(milligrams_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_literResult.Convert(units.MassConcentrationMilligramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerLiter = %v, want %v", converted, 100)
    }
    // Test CentigramPerLiter conversion
    centigrams_per_literDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationCentigramPerLiter,
    }
    
    var centigrams_per_literResult *units.MassConcentration
    centigrams_per_literResult, err = factory.FromDto(centigrams_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with CentigramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_literResult.Convert(units.MassConcentrationCentigramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerLiter = %v, want %v", converted, 100)
    }
    // Test DecigramPerLiter conversion
    decigrams_per_literDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationDecigramPerLiter,
    }
    
    var decigrams_per_literResult *units.MassConcentration
    decigrams_per_literResult, err = factory.FromDto(decigrams_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with DecigramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_literResult.Convert(units.MassConcentrationDecigramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerLiter = %v, want %v", converted, 100)
    }
    // Test KilogramPerLiter conversion
    kilograms_per_literDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationKilogramPerLiter,
    }
    
    var kilograms_per_literResult *units.MassConcentration
    kilograms_per_literResult, err = factory.FromDto(kilograms_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_literResult.Convert(units.MassConcentrationKilogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerLiter = %v, want %v", converted, 100)
    }
    // Test KilopoundPerCubicInch conversion
    kilopounds_per_cubic_inchDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationKilopoundPerCubicInch,
    }
    
    var kilopounds_per_cubic_inchResult *units.MassConcentration
    kilopounds_per_cubic_inchResult, err = factory.FromDto(kilopounds_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_per_cubic_inchResult.Convert(units.MassConcentrationKilopoundPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundPerCubicInch = %v, want %v", converted, 100)
    }
    // Test KilopoundPerCubicFoot conversion
    kilopounds_per_cubic_footDto := units.MassConcentrationDto{
        Value: 100,
        Unit:  units.MassConcentrationKilopoundPerCubicFoot,
    }
    
    var kilopounds_per_cubic_footResult *units.MassConcentration
    kilopounds_per_cubic_footResult, err = factory.FromDto(kilopounds_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_per_cubic_footResult.Convert(units.MassConcentrationKilopoundPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundPerCubicFoot = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MassConcentrationDto{
        Value: 0,
        Unit:  units.MassConcentrationKilogramPerCubicMeter,
    }
    
    var zeroResult *units.MassConcentration
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMassConcentrationFactory_FromDtoJSON(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "KilogramPerCubicMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "KilogramPerCubicMeter"}`)
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
    nanJSON, _ := json.Marshal(units.MassConcentrationDto{
        Value: nanValue,
        Unit:  units.MassConcentrationKilogramPerCubicMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with GramPerCubicMillimeter unit
    grams_per_cubic_millimeterJSON := []byte(`{"value": 100, "unit": "GramPerCubicMillimeter"}`)
    grams_per_cubic_millimeterResult, err := factory.FromDtoJSON(grams_per_cubic_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerCubicMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_cubic_millimeterResult.Convert(units.MassConcentrationGramPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerCubicCentimeter unit
    grams_per_cubic_centimeterJSON := []byte(`{"value": 100, "unit": "GramPerCubicCentimeter"}`)
    grams_per_cubic_centimeterResult, err := factory.FromDtoJSON(grams_per_cubic_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerCubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_cubic_centimeterResult.Convert(units.MassConcentrationGramPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerCubicMeter unit
    grams_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "GramPerCubicMeter"}`)
    grams_per_cubic_meterResult, err := factory.FromDtoJSON(grams_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_cubic_meterResult.Convert(units.MassConcentrationGramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerMicroliter unit
    grams_per_microliterJSON := []byte(`{"value": 100, "unit": "GramPerMicroliter"}`)
    grams_per_microliterResult, err := factory.FromDtoJSON(grams_per_microliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerMicroliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_microliterResult.Convert(units.MassConcentrationGramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerMilliliter unit
    grams_per_milliliterJSON := []byte(`{"value": 100, "unit": "GramPerMilliliter"}`)
    grams_per_milliliterResult, err := factory.FromDtoJSON(grams_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_milliliterResult.Convert(units.MassConcentrationGramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerDeciliter unit
    grams_per_deciliterJSON := []byte(`{"value": 100, "unit": "GramPerDeciliter"}`)
    grams_per_deciliterResult, err := factory.FromDtoJSON(grams_per_deciliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_deciliterResult.Convert(units.MassConcentrationGramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerLiter unit
    grams_per_literJSON := []byte(`{"value": 100, "unit": "GramPerLiter"}`)
    grams_per_literResult, err := factory.FromDtoJSON(grams_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_literResult.Convert(units.MassConcentrationGramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with TonnePerCubicMillimeter unit
    tonnes_per_cubic_millimeterJSON := []byte(`{"value": 100, "unit": "TonnePerCubicMillimeter"}`)
    tonnes_per_cubic_millimeterResult, err := factory.FromDtoJSON(tonnes_per_cubic_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonnePerCubicMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_millimeterResult.Convert(units.MassConcentrationTonnePerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonnePerCubicCentimeter unit
    tonnes_per_cubic_centimeterJSON := []byte(`{"value": 100, "unit": "TonnePerCubicCentimeter"}`)
    tonnes_per_cubic_centimeterResult, err := factory.FromDtoJSON(tonnes_per_cubic_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonnePerCubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_centimeterResult.Convert(units.MassConcentrationTonnePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonnePerCubicMeter unit
    tonnes_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "TonnePerCubicMeter"}`)
    tonnes_per_cubic_meterResult, err := factory.FromDtoJSON(tonnes_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonnePerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_meterResult.Convert(units.MassConcentrationTonnePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerCubicInch unit
    pounds_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "PoundPerCubicInch"}`)
    pounds_per_cubic_inchResult, err := factory.FromDtoJSON(pounds_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_inchResult.Convert(units.MassConcentrationPoundPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerCubicFoot unit
    pounds_per_cubic_footJSON := []byte(`{"value": 100, "unit": "PoundPerCubicFoot"}`)
    pounds_per_cubic_footResult, err := factory.FromDtoJSON(pounds_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_footResult.Convert(units.MassConcentrationPoundPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with SlugPerCubicFoot unit
    slugs_per_cubic_footJSON := []byte(`{"value": 100, "unit": "SlugPerCubicFoot"}`)
    slugs_per_cubic_footResult, err := factory.FromDtoJSON(slugs_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SlugPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugs_per_cubic_footResult.Convert(units.MassConcentrationSlugPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerUSGallon unit
    pounds_per_us_gallonJSON := []byte(`{"value": 100, "unit": "PoundPerUSGallon"}`)
    pounds_per_us_gallonResult, err := factory.FromDtoJSON(pounds_per_us_gallonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerUSGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_us_gallonResult.Convert(units.MassConcentrationPoundPerUSGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerUSGallon = %v, want %v", converted, 100)
    }
    // Test JSON with OuncePerUSGallon unit
    ounces_per_us_gallonJSON := []byte(`{"value": 100, "unit": "OuncePerUSGallon"}`)
    ounces_per_us_gallonResult, err := factory.FromDtoJSON(ounces_per_us_gallonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with OuncePerUSGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ounces_per_us_gallonResult.Convert(units.MassConcentrationOuncePerUSGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OuncePerUSGallon = %v, want %v", converted, 100)
    }
    // Test JSON with OuncePerImperialGallon unit
    ounces_per_imperial_gallonJSON := []byte(`{"value": 100, "unit": "OuncePerImperialGallon"}`)
    ounces_per_imperial_gallonResult, err := factory.FromDtoJSON(ounces_per_imperial_gallonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with OuncePerImperialGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ounces_per_imperial_gallonResult.Convert(units.MassConcentrationOuncePerImperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OuncePerImperialGallon = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerImperialGallon unit
    pounds_per_imperial_gallonJSON := []byte(`{"value": 100, "unit": "PoundPerImperialGallon"}`)
    pounds_per_imperial_gallonResult, err := factory.FromDtoJSON(pounds_per_imperial_gallonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerImperialGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_imperial_gallonResult.Convert(units.MassConcentrationPoundPerImperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerImperialGallon = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerCubicMillimeter unit
    kilograms_per_cubic_millimeterJSON := []byte(`{"value": 100, "unit": "KilogramPerCubicMillimeter"}`)
    kilograms_per_cubic_millimeterResult, err := factory.FromDtoJSON(kilograms_per_cubic_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerCubicMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_cubic_millimeterResult.Convert(units.MassConcentrationKilogramPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerCubicCentimeter unit
    kilograms_per_cubic_centimeterJSON := []byte(`{"value": 100, "unit": "KilogramPerCubicCentimeter"}`)
    kilograms_per_cubic_centimeterResult, err := factory.FromDtoJSON(kilograms_per_cubic_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerCubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_cubic_centimeterResult.Convert(units.MassConcentrationKilogramPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerCubicMeter unit
    kilograms_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "KilogramPerCubicMeter"}`)
    kilograms_per_cubic_meterResult, err := factory.FromDtoJSON(kilograms_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_cubic_meterResult.Convert(units.MassConcentrationKilogramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerCubicMeter unit
    milligrams_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "MilligramPerCubicMeter"}`)
    milligrams_per_cubic_meterResult, err := factory.FromDtoJSON(milligrams_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_cubic_meterResult.Convert(units.MassConcentrationMilligramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerCubicMeter unit
    micrograms_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "MicrogramPerCubicMeter"}`)
    micrograms_per_cubic_meterResult, err := factory.FromDtoJSON(micrograms_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_cubic_meterResult.Convert(units.MassConcentrationMicrogramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PicogramPerMicroliter unit
    picograms_per_microliterJSON := []byte(`{"value": 100, "unit": "PicogramPerMicroliter"}`)
    picograms_per_microliterResult, err := factory.FromDtoJSON(picograms_per_microliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicogramPerMicroliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_microliterResult.Convert(units.MassConcentrationPicogramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicogramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test JSON with NanogramPerMicroliter unit
    nanograms_per_microliterJSON := []byte(`{"value": 100, "unit": "NanogramPerMicroliter"}`)
    nanograms_per_microliterResult, err := factory.FromDtoJSON(nanograms_per_microliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanogramPerMicroliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_microliterResult.Convert(units.MassConcentrationNanogramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerMicroliter unit
    micrograms_per_microliterJSON := []byte(`{"value": 100, "unit": "MicrogramPerMicroliter"}`)
    micrograms_per_microliterResult, err := factory.FromDtoJSON(micrograms_per_microliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerMicroliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_microliterResult.Convert(units.MassConcentrationMicrogramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerMicroliter unit
    milligrams_per_microliterJSON := []byte(`{"value": 100, "unit": "MilligramPerMicroliter"}`)
    milligrams_per_microliterResult, err := factory.FromDtoJSON(milligrams_per_microliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerMicroliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_microliterResult.Convert(units.MassConcentrationMilligramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test JSON with CentigramPerMicroliter unit
    centigrams_per_microliterJSON := []byte(`{"value": 100, "unit": "CentigramPerMicroliter"}`)
    centigrams_per_microliterResult, err := factory.FromDtoJSON(centigrams_per_microliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigramPerMicroliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_microliterResult.Convert(units.MassConcentrationCentigramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test JSON with DecigramPerMicroliter unit
    decigrams_per_microliterJSON := []byte(`{"value": 100, "unit": "DecigramPerMicroliter"}`)
    decigrams_per_microliterResult, err := factory.FromDtoJSON(decigrams_per_microliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigramPerMicroliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_microliterResult.Convert(units.MassConcentrationDecigramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerMicroliter = %v, want %v", converted, 100)
    }
    // Test JSON with PicogramPerMilliliter unit
    picograms_per_milliliterJSON := []byte(`{"value": 100, "unit": "PicogramPerMilliliter"}`)
    picograms_per_milliliterResult, err := factory.FromDtoJSON(picograms_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicogramPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_milliliterResult.Convert(units.MassConcentrationPicogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicogramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with NanogramPerMilliliter unit
    nanograms_per_milliliterJSON := []byte(`{"value": 100, "unit": "NanogramPerMilliliter"}`)
    nanograms_per_milliliterResult, err := factory.FromDtoJSON(nanograms_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanogramPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_milliliterResult.Convert(units.MassConcentrationNanogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerMilliliter unit
    micrograms_per_milliliterJSON := []byte(`{"value": 100, "unit": "MicrogramPerMilliliter"}`)
    micrograms_per_milliliterResult, err := factory.FromDtoJSON(micrograms_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_milliliterResult.Convert(units.MassConcentrationMicrogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerMilliliter unit
    milligrams_per_milliliterJSON := []byte(`{"value": 100, "unit": "MilligramPerMilliliter"}`)
    milligrams_per_milliliterResult, err := factory.FromDtoJSON(milligrams_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_milliliterResult.Convert(units.MassConcentrationMilligramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with CentigramPerMilliliter unit
    centigrams_per_milliliterJSON := []byte(`{"value": 100, "unit": "CentigramPerMilliliter"}`)
    centigrams_per_milliliterResult, err := factory.FromDtoJSON(centigrams_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigramPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_milliliterResult.Convert(units.MassConcentrationCentigramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with DecigramPerMilliliter unit
    decigrams_per_milliliterJSON := []byte(`{"value": 100, "unit": "DecigramPerMilliliter"}`)
    decigrams_per_milliliterResult, err := factory.FromDtoJSON(decigrams_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigramPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_milliliterResult.Convert(units.MassConcentrationDecigramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with PicogramPerDeciliter unit
    picograms_per_deciliterJSON := []byte(`{"value": 100, "unit": "PicogramPerDeciliter"}`)
    picograms_per_deciliterResult, err := factory.FromDtoJSON(picograms_per_deciliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicogramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_deciliterResult.Convert(units.MassConcentrationPicogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with NanogramPerDeciliter unit
    nanograms_per_deciliterJSON := []byte(`{"value": 100, "unit": "NanogramPerDeciliter"}`)
    nanograms_per_deciliterResult, err := factory.FromDtoJSON(nanograms_per_deciliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanogramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_deciliterResult.Convert(units.MassConcentrationNanogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerDeciliter unit
    micrograms_per_deciliterJSON := []byte(`{"value": 100, "unit": "MicrogramPerDeciliter"}`)
    micrograms_per_deciliterResult, err := factory.FromDtoJSON(micrograms_per_deciliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_deciliterResult.Convert(units.MassConcentrationMicrogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerDeciliter unit
    milligrams_per_deciliterJSON := []byte(`{"value": 100, "unit": "MilligramPerDeciliter"}`)
    milligrams_per_deciliterResult, err := factory.FromDtoJSON(milligrams_per_deciliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_deciliterResult.Convert(units.MassConcentrationMilligramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with CentigramPerDeciliter unit
    centigrams_per_deciliterJSON := []byte(`{"value": 100, "unit": "CentigramPerDeciliter"}`)
    centigrams_per_deciliterResult, err := factory.FromDtoJSON(centigrams_per_deciliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_deciliterResult.Convert(units.MassConcentrationCentigramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with DecigramPerDeciliter unit
    decigrams_per_deciliterJSON := []byte(`{"value": 100, "unit": "DecigramPerDeciliter"}`)
    decigrams_per_deciliterResult, err := factory.FromDtoJSON(decigrams_per_deciliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_deciliterResult.Convert(units.MassConcentrationDecigramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with PicogramPerLiter unit
    picograms_per_literJSON := []byte(`{"value": 100, "unit": "PicogramPerLiter"}`)
    picograms_per_literResult, err := factory.FromDtoJSON(picograms_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicogramPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_literResult.Convert(units.MassConcentrationPicogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicogramPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with NanogramPerLiter unit
    nanograms_per_literJSON := []byte(`{"value": 100, "unit": "NanogramPerLiter"}`)
    nanograms_per_literResult, err := factory.FromDtoJSON(nanograms_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanogramPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_literResult.Convert(units.MassConcentrationNanogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerLiter unit
    micrograms_per_literJSON := []byte(`{"value": 100, "unit": "MicrogramPerLiter"}`)
    micrograms_per_literResult, err := factory.FromDtoJSON(micrograms_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_literResult.Convert(units.MassConcentrationMicrogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerLiter unit
    milligrams_per_literJSON := []byte(`{"value": 100, "unit": "MilligramPerLiter"}`)
    milligrams_per_literResult, err := factory.FromDtoJSON(milligrams_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_literResult.Convert(units.MassConcentrationMilligramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with CentigramPerLiter unit
    centigrams_per_literJSON := []byte(`{"value": 100, "unit": "CentigramPerLiter"}`)
    centigrams_per_literResult, err := factory.FromDtoJSON(centigrams_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigramPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_literResult.Convert(units.MassConcentrationCentigramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with DecigramPerLiter unit
    decigrams_per_literJSON := []byte(`{"value": 100, "unit": "DecigramPerLiter"}`)
    decigrams_per_literResult, err := factory.FromDtoJSON(decigrams_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigramPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_literResult.Convert(units.MassConcentrationDecigramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerLiter unit
    kilograms_per_literJSON := []byte(`{"value": 100, "unit": "KilogramPerLiter"}`)
    kilograms_per_literResult, err := factory.FromDtoJSON(kilograms_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_literResult.Convert(units.MassConcentrationKilogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundPerCubicInch unit
    kilopounds_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "KilopoundPerCubicInch"}`)
    kilopounds_per_cubic_inchResult, err := factory.FromDtoJSON(kilopounds_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_per_cubic_inchResult.Convert(units.MassConcentrationKilopoundPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundPerCubicFoot unit
    kilopounds_per_cubic_footJSON := []byte(`{"value": 100, "unit": "KilopoundPerCubicFoot"}`)
    kilopounds_per_cubic_footResult, err := factory.FromDtoJSON(kilopounds_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_per_cubic_footResult.Convert(units.MassConcentrationKilopoundPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundPerCubicFoot = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "KilogramPerCubicMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromGramsPerCubicMillimeter function
func TestMassConcentrationFactory_FromGramsPerCubicMillimeter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerCubicMillimeter(100)
    if err != nil {
        t.Errorf("FromGramsPerCubicMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationGramPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerCubicMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerCubicMillimeter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerCubicMillimeter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerCubicMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerCubicMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerCubicMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerCubicMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerCubicMillimeter(0)
    if err != nil {
        t.Errorf("FromGramsPerCubicMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationGramPerCubicMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerCubicMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerCubicCentimeter function
func TestMassConcentrationFactory_FromGramsPerCubicCentimeter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromGramsPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationGramPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerCubicCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerCubicCentimeter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerCubicCentimeter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerCubicCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerCubicCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerCubicCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerCubicCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerCubicCentimeter(0)
    if err != nil {
        t.Errorf("FromGramsPerCubicCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationGramPerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerCubicMeter function
func TestMassConcentrationFactory_FromGramsPerCubicMeter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromGramsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationGramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromGramsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationGramPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerMicroliter function
func TestMassConcentrationFactory_FromGramsPerMicroliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerMicroliter(100)
    if err != nil {
        t.Errorf("FromGramsPerMicroliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationGramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerMicroliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerMicroliter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerMicroliter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerMicroliter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerMicroliter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerMicroliter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerMicroliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerMicroliter(0)
    if err != nil {
        t.Errorf("FromGramsPerMicroliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationGramPerMicroliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerMicroliter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerMilliliter function
func TestMassConcentrationFactory_FromGramsPerMilliliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromGramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationGramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerMilliliter(0)
    if err != nil {
        t.Errorf("FromGramsPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationGramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerDeciliter function
func TestMassConcentrationFactory_FromGramsPerDeciliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerDeciliter(100)
    if err != nil {
        t.Errorf("FromGramsPerDeciliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationGramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerDeciliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerDeciliter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerDeciliter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerDeciliter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerDeciliter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerDeciliter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerDeciliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerDeciliter(0)
    if err != nil {
        t.Errorf("FromGramsPerDeciliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationGramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerDeciliter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerLiter function
func TestMassConcentrationFactory_FromGramsPerLiter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerLiter(100)
    if err != nil {
        t.Errorf("FromGramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationGramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerLiter(0)
    if err != nil {
        t.Errorf("FromGramsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationGramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesPerCubicMillimeter function
func TestMassConcentrationFactory_FromTonnesPerCubicMillimeter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesPerCubicMillimeter(100)
    if err != nil {
        t.Errorf("FromTonnesPerCubicMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationTonnePerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesPerCubicMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesPerCubicMillimeter(math.NaN())
    if err == nil {
        t.Error("FromTonnesPerCubicMillimeter() with NaN value should return error")
    }

    _, err = factory.FromTonnesPerCubicMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesPerCubicMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromTonnesPerCubicMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesPerCubicMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesPerCubicMillimeter(0)
    if err != nil {
        t.Errorf("FromTonnesPerCubicMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationTonnePerCubicMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesPerCubicMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesPerCubicCentimeter function
func TestMassConcentrationFactory_FromTonnesPerCubicCentimeter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromTonnesPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationTonnePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesPerCubicCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesPerCubicCentimeter(math.NaN())
    if err == nil {
        t.Error("FromTonnesPerCubicCentimeter() with NaN value should return error")
    }

    _, err = factory.FromTonnesPerCubicCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesPerCubicCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromTonnesPerCubicCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesPerCubicCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesPerCubicCentimeter(0)
    if err != nil {
        t.Errorf("FromTonnesPerCubicCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationTonnePerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesPerCubicMeter function
func TestMassConcentrationFactory_FromTonnesPerCubicMeter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromTonnesPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationTonnePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromTonnesPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromTonnesPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromTonnesPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromTonnesPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationTonnePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerCubicInch function
func TestMassConcentrationFactory_FromPoundsPerCubicInch(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromPoundsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationPoundPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromPoundsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationPoundPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerCubicFoot function
func TestMassConcentrationFactory_FromPoundsPerCubicFoot(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromPoundsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationPoundPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromPoundsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationPoundPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromSlugsPerCubicFoot function
func TestMassConcentrationFactory_FromSlugsPerCubicFoot(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSlugsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromSlugsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationSlugPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSlugsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSlugsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromSlugsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromSlugsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromSlugsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromSlugsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromSlugsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSlugsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromSlugsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationSlugPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSlugsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerUSGallon function
func TestMassConcentrationFactory_FromPoundsPerUSGallon(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerUSGallon(100)
    if err != nil {
        t.Errorf("FromPoundsPerUSGallon() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationPoundPerUSGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerUSGallon() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerUSGallon(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerUSGallon() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerUSGallon(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerUSGallon() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerUSGallon(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerUSGallon() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerUSGallon(0)
    if err != nil {
        t.Errorf("FromPoundsPerUSGallon() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationPoundPerUSGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerUSGallon() with zero value = %v, want 0", converted)
    }
}
// Test FromOuncesPerUSGallon function
func TestMassConcentrationFactory_FromOuncesPerUSGallon(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOuncesPerUSGallon(100)
    if err != nil {
        t.Errorf("FromOuncesPerUSGallon() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationOuncePerUSGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOuncesPerUSGallon() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOuncesPerUSGallon(math.NaN())
    if err == nil {
        t.Error("FromOuncesPerUSGallon() with NaN value should return error")
    }

    _, err = factory.FromOuncesPerUSGallon(math.Inf(1))
    if err == nil {
        t.Error("FromOuncesPerUSGallon() with +Inf value should return error")
    }

    _, err = factory.FromOuncesPerUSGallon(math.Inf(-1))
    if err == nil {
        t.Error("FromOuncesPerUSGallon() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOuncesPerUSGallon(0)
    if err != nil {
        t.Errorf("FromOuncesPerUSGallon() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationOuncePerUSGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOuncesPerUSGallon() with zero value = %v, want 0", converted)
    }
}
// Test FromOuncesPerImperialGallon function
func TestMassConcentrationFactory_FromOuncesPerImperialGallon(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOuncesPerImperialGallon(100)
    if err != nil {
        t.Errorf("FromOuncesPerImperialGallon() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationOuncePerImperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOuncesPerImperialGallon() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOuncesPerImperialGallon(math.NaN())
    if err == nil {
        t.Error("FromOuncesPerImperialGallon() with NaN value should return error")
    }

    _, err = factory.FromOuncesPerImperialGallon(math.Inf(1))
    if err == nil {
        t.Error("FromOuncesPerImperialGallon() with +Inf value should return error")
    }

    _, err = factory.FromOuncesPerImperialGallon(math.Inf(-1))
    if err == nil {
        t.Error("FromOuncesPerImperialGallon() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOuncesPerImperialGallon(0)
    if err != nil {
        t.Errorf("FromOuncesPerImperialGallon() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationOuncePerImperialGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOuncesPerImperialGallon() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerImperialGallon function
func TestMassConcentrationFactory_FromPoundsPerImperialGallon(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerImperialGallon(100)
    if err != nil {
        t.Errorf("FromPoundsPerImperialGallon() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationPoundPerImperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerImperialGallon() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerImperialGallon(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerImperialGallon() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerImperialGallon(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerImperialGallon() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerImperialGallon(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerImperialGallon() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerImperialGallon(0)
    if err != nil {
        t.Errorf("FromPoundsPerImperialGallon() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationPoundPerImperialGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerImperialGallon() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerCubicMillimeter function
func TestMassConcentrationFactory_FromKilogramsPerCubicMillimeter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerCubicMillimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerCubicMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationKilogramPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerCubicMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerCubicMillimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerCubicMillimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerCubicMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerCubicMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerCubicMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerCubicMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerCubicMillimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerCubicMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationKilogramPerCubicMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerCubicMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerCubicCentimeter function
func TestMassConcentrationFactory_FromKilogramsPerCubicCentimeter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationKilogramPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerCubicCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerCubicCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerCubicCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerCubicCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerCubicCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerCubicCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerCubicCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerCubicCentimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerCubicCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationKilogramPerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerCubicMeter function
func TestMassConcentrationFactory_FromKilogramsPerCubicMeter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationKilogramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationKilogramPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerCubicMeter function
func TestMassConcentrationFactory_FromMilligramsPerCubicMeter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMilligramsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationMilligramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromMilligramsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationMilligramPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerCubicMeter function
func TestMassConcentrationFactory_FromMicrogramsPerCubicMeter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationMicrogramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationMicrogramPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicogramsPerMicroliter function
func TestMassConcentrationFactory_FromPicogramsPerMicroliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicogramsPerMicroliter(100)
    if err != nil {
        t.Errorf("FromPicogramsPerMicroliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationPicogramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicogramsPerMicroliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicogramsPerMicroliter(math.NaN())
    if err == nil {
        t.Error("FromPicogramsPerMicroliter() with NaN value should return error")
    }

    _, err = factory.FromPicogramsPerMicroliter(math.Inf(1))
    if err == nil {
        t.Error("FromPicogramsPerMicroliter() with +Inf value should return error")
    }

    _, err = factory.FromPicogramsPerMicroliter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicogramsPerMicroliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicogramsPerMicroliter(0)
    if err != nil {
        t.Errorf("FromPicogramsPerMicroliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationPicogramPerMicroliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicogramsPerMicroliter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanogramsPerMicroliter function
func TestMassConcentrationFactory_FromNanogramsPerMicroliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanogramsPerMicroliter(100)
    if err != nil {
        t.Errorf("FromNanogramsPerMicroliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationNanogramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanogramsPerMicroliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanogramsPerMicroliter(math.NaN())
    if err == nil {
        t.Error("FromNanogramsPerMicroliter() with NaN value should return error")
    }

    _, err = factory.FromNanogramsPerMicroliter(math.Inf(1))
    if err == nil {
        t.Error("FromNanogramsPerMicroliter() with +Inf value should return error")
    }

    _, err = factory.FromNanogramsPerMicroliter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanogramsPerMicroliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanogramsPerMicroliter(0)
    if err != nil {
        t.Errorf("FromNanogramsPerMicroliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationNanogramPerMicroliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanogramsPerMicroliter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerMicroliter function
func TestMassConcentrationFactory_FromMicrogramsPerMicroliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerMicroliter(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerMicroliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationMicrogramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerMicroliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerMicroliter(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerMicroliter() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerMicroliter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerMicroliter() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerMicroliter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerMicroliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerMicroliter(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerMicroliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationMicrogramPerMicroliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerMicroliter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerMicroliter function
func TestMassConcentrationFactory_FromMilligramsPerMicroliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerMicroliter(100)
    if err != nil {
        t.Errorf("FromMilligramsPerMicroliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationMilligramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerMicroliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerMicroliter(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerMicroliter() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerMicroliter(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerMicroliter() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerMicroliter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerMicroliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerMicroliter(0)
    if err != nil {
        t.Errorf("FromMilligramsPerMicroliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationMilligramPerMicroliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerMicroliter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigramsPerMicroliter function
func TestMassConcentrationFactory_FromCentigramsPerMicroliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigramsPerMicroliter(100)
    if err != nil {
        t.Errorf("FromCentigramsPerMicroliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationCentigramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigramsPerMicroliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigramsPerMicroliter(math.NaN())
    if err == nil {
        t.Error("FromCentigramsPerMicroliter() with NaN value should return error")
    }

    _, err = factory.FromCentigramsPerMicroliter(math.Inf(1))
    if err == nil {
        t.Error("FromCentigramsPerMicroliter() with +Inf value should return error")
    }

    _, err = factory.FromCentigramsPerMicroliter(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigramsPerMicroliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigramsPerMicroliter(0)
    if err != nil {
        t.Errorf("FromCentigramsPerMicroliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationCentigramPerMicroliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigramsPerMicroliter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigramsPerMicroliter function
func TestMassConcentrationFactory_FromDecigramsPerMicroliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigramsPerMicroliter(100)
    if err != nil {
        t.Errorf("FromDecigramsPerMicroliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationDecigramPerMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigramsPerMicroliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigramsPerMicroliter(math.NaN())
    if err == nil {
        t.Error("FromDecigramsPerMicroliter() with NaN value should return error")
    }

    _, err = factory.FromDecigramsPerMicroliter(math.Inf(1))
    if err == nil {
        t.Error("FromDecigramsPerMicroliter() with +Inf value should return error")
    }

    _, err = factory.FromDecigramsPerMicroliter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigramsPerMicroliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigramsPerMicroliter(0)
    if err != nil {
        t.Errorf("FromDecigramsPerMicroliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationDecigramPerMicroliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigramsPerMicroliter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicogramsPerMilliliter function
func TestMassConcentrationFactory_FromPicogramsPerMilliliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicogramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromPicogramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationPicogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicogramsPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicogramsPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromPicogramsPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromPicogramsPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromPicogramsPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromPicogramsPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicogramsPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicogramsPerMilliliter(0)
    if err != nil {
        t.Errorf("FromPicogramsPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationPicogramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicogramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanogramsPerMilliliter function
func TestMassConcentrationFactory_FromNanogramsPerMilliliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanogramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromNanogramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationNanogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanogramsPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanogramsPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromNanogramsPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromNanogramsPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromNanogramsPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromNanogramsPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanogramsPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanogramsPerMilliliter(0)
    if err != nil {
        t.Errorf("FromNanogramsPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationNanogramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanogramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerMilliliter function
func TestMassConcentrationFactory_FromMicrogramsPerMilliliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationMicrogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerMilliliter(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationMicrogramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerMilliliter function
func TestMassConcentrationFactory_FromMilligramsPerMilliliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromMilligramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationMilligramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerMilliliter(0)
    if err != nil {
        t.Errorf("FromMilligramsPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationMilligramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigramsPerMilliliter function
func TestMassConcentrationFactory_FromCentigramsPerMilliliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromCentigramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationCentigramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigramsPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigramsPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromCentigramsPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromCentigramsPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromCentigramsPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromCentigramsPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigramsPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigramsPerMilliliter(0)
    if err != nil {
        t.Errorf("FromCentigramsPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationCentigramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigramsPerMilliliter function
func TestMassConcentrationFactory_FromDecigramsPerMilliliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromDecigramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationDecigramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigramsPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigramsPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromDecigramsPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromDecigramsPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromDecigramsPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromDecigramsPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigramsPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigramsPerMilliliter(0)
    if err != nil {
        t.Errorf("FromDecigramsPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationDecigramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicogramsPerDeciliter function
func TestMassConcentrationFactory_FromPicogramsPerDeciliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicogramsPerDeciliter(100)
    if err != nil {
        t.Errorf("FromPicogramsPerDeciliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationPicogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicogramsPerDeciliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicogramsPerDeciliter(math.NaN())
    if err == nil {
        t.Error("FromPicogramsPerDeciliter() with NaN value should return error")
    }

    _, err = factory.FromPicogramsPerDeciliter(math.Inf(1))
    if err == nil {
        t.Error("FromPicogramsPerDeciliter() with +Inf value should return error")
    }

    _, err = factory.FromPicogramsPerDeciliter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicogramsPerDeciliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicogramsPerDeciliter(0)
    if err != nil {
        t.Errorf("FromPicogramsPerDeciliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationPicogramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicogramsPerDeciliter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanogramsPerDeciliter function
func TestMassConcentrationFactory_FromNanogramsPerDeciliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanogramsPerDeciliter(100)
    if err != nil {
        t.Errorf("FromNanogramsPerDeciliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationNanogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanogramsPerDeciliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanogramsPerDeciliter(math.NaN())
    if err == nil {
        t.Error("FromNanogramsPerDeciliter() with NaN value should return error")
    }

    _, err = factory.FromNanogramsPerDeciliter(math.Inf(1))
    if err == nil {
        t.Error("FromNanogramsPerDeciliter() with +Inf value should return error")
    }

    _, err = factory.FromNanogramsPerDeciliter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanogramsPerDeciliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanogramsPerDeciliter(0)
    if err != nil {
        t.Errorf("FromNanogramsPerDeciliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationNanogramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanogramsPerDeciliter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerDeciliter function
func TestMassConcentrationFactory_FromMicrogramsPerDeciliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerDeciliter(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerDeciliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationMicrogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerDeciliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerDeciliter(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerDeciliter() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerDeciliter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerDeciliter() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerDeciliter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerDeciliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerDeciliter(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerDeciliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationMicrogramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerDeciliter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerDeciliter function
func TestMassConcentrationFactory_FromMilligramsPerDeciliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerDeciliter(100)
    if err != nil {
        t.Errorf("FromMilligramsPerDeciliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationMilligramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerDeciliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerDeciliter(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerDeciliter() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerDeciliter(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerDeciliter() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerDeciliter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerDeciliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerDeciliter(0)
    if err != nil {
        t.Errorf("FromMilligramsPerDeciliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationMilligramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerDeciliter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigramsPerDeciliter function
func TestMassConcentrationFactory_FromCentigramsPerDeciliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigramsPerDeciliter(100)
    if err != nil {
        t.Errorf("FromCentigramsPerDeciliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationCentigramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigramsPerDeciliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigramsPerDeciliter(math.NaN())
    if err == nil {
        t.Error("FromCentigramsPerDeciliter() with NaN value should return error")
    }

    _, err = factory.FromCentigramsPerDeciliter(math.Inf(1))
    if err == nil {
        t.Error("FromCentigramsPerDeciliter() with +Inf value should return error")
    }

    _, err = factory.FromCentigramsPerDeciliter(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigramsPerDeciliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigramsPerDeciliter(0)
    if err != nil {
        t.Errorf("FromCentigramsPerDeciliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationCentigramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigramsPerDeciliter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigramsPerDeciliter function
func TestMassConcentrationFactory_FromDecigramsPerDeciliter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigramsPerDeciliter(100)
    if err != nil {
        t.Errorf("FromDecigramsPerDeciliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationDecigramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigramsPerDeciliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigramsPerDeciliter(math.NaN())
    if err == nil {
        t.Error("FromDecigramsPerDeciliter() with NaN value should return error")
    }

    _, err = factory.FromDecigramsPerDeciliter(math.Inf(1))
    if err == nil {
        t.Error("FromDecigramsPerDeciliter() with +Inf value should return error")
    }

    _, err = factory.FromDecigramsPerDeciliter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigramsPerDeciliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigramsPerDeciliter(0)
    if err != nil {
        t.Errorf("FromDecigramsPerDeciliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationDecigramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigramsPerDeciliter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicogramsPerLiter function
func TestMassConcentrationFactory_FromPicogramsPerLiter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicogramsPerLiter(100)
    if err != nil {
        t.Errorf("FromPicogramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationPicogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicogramsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicogramsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromPicogramsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromPicogramsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromPicogramsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromPicogramsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicogramsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicogramsPerLiter(0)
    if err != nil {
        t.Errorf("FromPicogramsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationPicogramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicogramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanogramsPerLiter function
func TestMassConcentrationFactory_FromNanogramsPerLiter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanogramsPerLiter(100)
    if err != nil {
        t.Errorf("FromNanogramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationNanogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanogramsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanogramsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromNanogramsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromNanogramsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromNanogramsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromNanogramsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanogramsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanogramsPerLiter(0)
    if err != nil {
        t.Errorf("FromNanogramsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationNanogramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanogramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerLiter function
func TestMassConcentrationFactory_FromMicrogramsPerLiter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerLiter(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationMicrogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerLiter(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationMicrogramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerLiter function
func TestMassConcentrationFactory_FromMilligramsPerLiter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerLiter(100)
    if err != nil {
        t.Errorf("FromMilligramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationMilligramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerLiter(0)
    if err != nil {
        t.Errorf("FromMilligramsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationMilligramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigramsPerLiter function
func TestMassConcentrationFactory_FromCentigramsPerLiter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigramsPerLiter(100)
    if err != nil {
        t.Errorf("FromCentigramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationCentigramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigramsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigramsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromCentigramsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromCentigramsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromCentigramsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromCentigramsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigramsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigramsPerLiter(0)
    if err != nil {
        t.Errorf("FromCentigramsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationCentigramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigramsPerLiter function
func TestMassConcentrationFactory_FromDecigramsPerLiter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigramsPerLiter(100)
    if err != nil {
        t.Errorf("FromDecigramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationDecigramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigramsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigramsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromDecigramsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromDecigramsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromDecigramsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromDecigramsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigramsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigramsPerLiter(0)
    if err != nil {
        t.Errorf("FromDecigramsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationDecigramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerLiter function
func TestMassConcentrationFactory_FromKilogramsPerLiter(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerLiter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationKilogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerLiter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationKilogramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsPerCubicInch function
func TestMassConcentrationFactory_FromKilopoundsPerCubicInch(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromKilopoundsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationKilopoundPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromKilopoundsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationKilopoundPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsPerCubicFoot function
func TestMassConcentrationFactory_FromKilopoundsPerCubicFoot(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromKilopoundsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MassConcentrationKilopoundPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromKilopoundsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MassConcentrationKilopoundPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}

func TestMassConcentrationToString(t *testing.T) {
	factory := units.MassConcentrationFactory{}
	a, err := factory.CreateMassConcentration(45, units.MassConcentrationKilogramPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MassConcentrationKilogramPerCubicMeter, 2)
	expected := "45.00 " + units.GetMassConcentrationAbbreviation(units.MassConcentrationKilogramPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MassConcentrationKilogramPerCubicMeter, -1)
	expected = "45 " + units.GetMassConcentrationAbbreviation(units.MassConcentrationKilogramPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMassConcentration_EqualityAndComparison(t *testing.T) {
	factory := units.MassConcentrationFactory{}
	a1, _ := factory.CreateMassConcentration(60, units.MassConcentrationKilogramPerCubicMeter)
	a2, _ := factory.CreateMassConcentration(60, units.MassConcentrationKilogramPerCubicMeter)
	a3, _ := factory.CreateMassConcentration(120, units.MassConcentrationKilogramPerCubicMeter)

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

func TestMassConcentration_Arithmetic(t *testing.T) {
	factory := units.MassConcentrationFactory{}
	a1, _ := factory.CreateMassConcentration(30, units.MassConcentrationKilogramPerCubicMeter)
	a2, _ := factory.CreateMassConcentration(45, units.MassConcentrationKilogramPerCubicMeter)

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


func TestGetMassConcentrationAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.MassConcentrationUnits
        want string
    }{
        {
            name: "GramPerCubicMillimeter abbreviation",
            unit: units.MassConcentrationGramPerCubicMillimeter,
            want: "g/mm³",
        },
        {
            name: "GramPerCubicCentimeter abbreviation",
            unit: units.MassConcentrationGramPerCubicCentimeter,
            want: "g/cm³",
        },
        {
            name: "GramPerCubicMeter abbreviation",
            unit: units.MassConcentrationGramPerCubicMeter,
            want: "g/m³",
        },
        {
            name: "GramPerMicroliter abbreviation",
            unit: units.MassConcentrationGramPerMicroliter,
            want: "g/μL",
        },
        {
            name: "GramPerMilliliter abbreviation",
            unit: units.MassConcentrationGramPerMilliliter,
            want: "g/mL",
        },
        {
            name: "GramPerDeciliter abbreviation",
            unit: units.MassConcentrationGramPerDeciliter,
            want: "g/dL",
        },
        {
            name: "GramPerLiter abbreviation",
            unit: units.MassConcentrationGramPerLiter,
            want: "g/L",
        },
        {
            name: "TonnePerCubicMillimeter abbreviation",
            unit: units.MassConcentrationTonnePerCubicMillimeter,
            want: "t/mm³",
        },
        {
            name: "TonnePerCubicCentimeter abbreviation",
            unit: units.MassConcentrationTonnePerCubicCentimeter,
            want: "t/cm³",
        },
        {
            name: "TonnePerCubicMeter abbreviation",
            unit: units.MassConcentrationTonnePerCubicMeter,
            want: "t/m³",
        },
        {
            name: "PoundPerCubicInch abbreviation",
            unit: units.MassConcentrationPoundPerCubicInch,
            want: "lb/in³",
        },
        {
            name: "PoundPerCubicFoot abbreviation",
            unit: units.MassConcentrationPoundPerCubicFoot,
            want: "lb/ft³",
        },
        {
            name: "SlugPerCubicFoot abbreviation",
            unit: units.MassConcentrationSlugPerCubicFoot,
            want: "slug/ft³",
        },
        {
            name: "PoundPerUSGallon abbreviation",
            unit: units.MassConcentrationPoundPerUSGallon,
            want: "ppg (U.S.)",
        },
        {
            name: "OuncePerUSGallon abbreviation",
            unit: units.MassConcentrationOuncePerUSGallon,
            want: "oz/gal (U.S.)",
        },
        {
            name: "OuncePerImperialGallon abbreviation",
            unit: units.MassConcentrationOuncePerImperialGallon,
            want: "oz/gal (imp.)",
        },
        {
            name: "PoundPerImperialGallon abbreviation",
            unit: units.MassConcentrationPoundPerImperialGallon,
            want: "ppg (imp.)",
        },
        {
            name: "KilogramPerCubicMillimeter abbreviation",
            unit: units.MassConcentrationKilogramPerCubicMillimeter,
            want: "kg/mm³",
        },
        {
            name: "KilogramPerCubicCentimeter abbreviation",
            unit: units.MassConcentrationKilogramPerCubicCentimeter,
            want: "kg/cm³",
        },
        {
            name: "KilogramPerCubicMeter abbreviation",
            unit: units.MassConcentrationKilogramPerCubicMeter,
            want: "kg/m³",
        },
        {
            name: "MilligramPerCubicMeter abbreviation",
            unit: units.MassConcentrationMilligramPerCubicMeter,
            want: "mg/m³",
        },
        {
            name: "MicrogramPerCubicMeter abbreviation",
            unit: units.MassConcentrationMicrogramPerCubicMeter,
            want: "μg/m³",
        },
        {
            name: "PicogramPerMicroliter abbreviation",
            unit: units.MassConcentrationPicogramPerMicroliter,
            want: "pg/μL",
        },
        {
            name: "NanogramPerMicroliter abbreviation",
            unit: units.MassConcentrationNanogramPerMicroliter,
            want: "ng/μL",
        },
        {
            name: "MicrogramPerMicroliter abbreviation",
            unit: units.MassConcentrationMicrogramPerMicroliter,
            want: "μg/μL",
        },
        {
            name: "MilligramPerMicroliter abbreviation",
            unit: units.MassConcentrationMilligramPerMicroliter,
            want: "mg/μL",
        },
        {
            name: "CentigramPerMicroliter abbreviation",
            unit: units.MassConcentrationCentigramPerMicroliter,
            want: "cg/μL",
        },
        {
            name: "DecigramPerMicroliter abbreviation",
            unit: units.MassConcentrationDecigramPerMicroliter,
            want: "dg/μL",
        },
        {
            name: "PicogramPerMilliliter abbreviation",
            unit: units.MassConcentrationPicogramPerMilliliter,
            want: "pg/mL",
        },
        {
            name: "NanogramPerMilliliter abbreviation",
            unit: units.MassConcentrationNanogramPerMilliliter,
            want: "ng/mL",
        },
        {
            name: "MicrogramPerMilliliter abbreviation",
            unit: units.MassConcentrationMicrogramPerMilliliter,
            want: "μg/mL",
        },
        {
            name: "MilligramPerMilliliter abbreviation",
            unit: units.MassConcentrationMilligramPerMilliliter,
            want: "mg/mL",
        },
        {
            name: "CentigramPerMilliliter abbreviation",
            unit: units.MassConcentrationCentigramPerMilliliter,
            want: "cg/mL",
        },
        {
            name: "DecigramPerMilliliter abbreviation",
            unit: units.MassConcentrationDecigramPerMilliliter,
            want: "dg/mL",
        },
        {
            name: "PicogramPerDeciliter abbreviation",
            unit: units.MassConcentrationPicogramPerDeciliter,
            want: "pg/dL",
        },
        {
            name: "NanogramPerDeciliter abbreviation",
            unit: units.MassConcentrationNanogramPerDeciliter,
            want: "ng/dL",
        },
        {
            name: "MicrogramPerDeciliter abbreviation",
            unit: units.MassConcentrationMicrogramPerDeciliter,
            want: "μg/dL",
        },
        {
            name: "MilligramPerDeciliter abbreviation",
            unit: units.MassConcentrationMilligramPerDeciliter,
            want: "mg/dL",
        },
        {
            name: "CentigramPerDeciliter abbreviation",
            unit: units.MassConcentrationCentigramPerDeciliter,
            want: "cg/dL",
        },
        {
            name: "DecigramPerDeciliter abbreviation",
            unit: units.MassConcentrationDecigramPerDeciliter,
            want: "dg/dL",
        },
        {
            name: "PicogramPerLiter abbreviation",
            unit: units.MassConcentrationPicogramPerLiter,
            want: "pg/L",
        },
        {
            name: "NanogramPerLiter abbreviation",
            unit: units.MassConcentrationNanogramPerLiter,
            want: "ng/L",
        },
        {
            name: "MicrogramPerLiter abbreviation",
            unit: units.MassConcentrationMicrogramPerLiter,
            want: "μg/L",
        },
        {
            name: "MilligramPerLiter abbreviation",
            unit: units.MassConcentrationMilligramPerLiter,
            want: "mg/L",
        },
        {
            name: "CentigramPerLiter abbreviation",
            unit: units.MassConcentrationCentigramPerLiter,
            want: "cg/L",
        },
        {
            name: "DecigramPerLiter abbreviation",
            unit: units.MassConcentrationDecigramPerLiter,
            want: "dg/L",
        },
        {
            name: "KilogramPerLiter abbreviation",
            unit: units.MassConcentrationKilogramPerLiter,
            want: "kg/L",
        },
        {
            name: "KilopoundPerCubicInch abbreviation",
            unit: units.MassConcentrationKilopoundPerCubicInch,
            want: "klb/in³",
        },
        {
            name: "KilopoundPerCubicFoot abbreviation",
            unit: units.MassConcentrationKilopoundPerCubicFoot,
            want: "klb/ft³",
        },
        {
            name: "invalid unit",
            unit: units.MassConcentrationUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetMassConcentrationAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetMassConcentrationAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestMassConcentration_String(t *testing.T) {
    factory := units.MassConcentrationFactory{}
    
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
            unit, err := factory.CreateMassConcentration(tt.value, units.MassConcentrationKilogramPerCubicMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("MassConcentration.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestMassConcentration_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.MassConcentrationFactory{}

	_, err := uf.CreateMassConcentration(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}