// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerCubicMeter"}`
	
	factory := units.DensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.DensityKilogramPerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.DensityKilogramPerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestDensityDto_ToJSON(t *testing.T) {
	dto := units.DensityDto{
		Value: 45,
		Unit:  units.DensityKilogramPerCubicMeter,
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
	if result["unit"].(string) != string(units.DensityKilogramPerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.DensityKilogramPerCubicMeter, result["unit"])
	}
}

func TestNewDensity_InvalidValue(t *testing.T) {
	factory := units.DensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateDensity(math.NaN(), units.DensityKilogramPerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateDensity(math.Inf(1), units.DensityKilogramPerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestDensityConversions(t *testing.T) {
	factory := units.DensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateDensity(180, units.DensityKilogramPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to GramsPerCubicMillimeter.
		// No expected conversion value provided for GramsPerCubicMillimeter, verifying result is not NaN.
		result := a.GramsPerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerCubicCentimeter.
		// No expected conversion value provided for GramsPerCubicCentimeter, verifying result is not NaN.
		result := a.GramsPerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerCubicMeter.
		// No expected conversion value provided for GramsPerCubicMeter, verifying result is not NaN.
		result := a.GramsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicInch.
		// No expected conversion value provided for PoundsPerCubicInch, verifying result is not NaN.
		result := a.PoundsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicFoot.
		// No expected conversion value provided for PoundsPerCubicFoot, verifying result is not NaN.
		result := a.PoundsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicYard.
		// No expected conversion value provided for PoundsPerCubicYard, verifying result is not NaN.
		result := a.PoundsPerCubicYard()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerCubicYard returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerCubicMillimeter.
		// No expected conversion value provided for TonnesPerCubicMillimeter, verifying result is not NaN.
		result := a.TonnesPerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerCubicCentimeter.
		// No expected conversion value provided for TonnesPerCubicCentimeter, verifying result is not NaN.
		result := a.TonnesPerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerCubicMeter.
		// No expected conversion value provided for TonnesPerCubicMeter, verifying result is not NaN.
		result := a.TonnesPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to SlugsPerCubicFoot.
		// No expected conversion value provided for SlugsPerCubicFoot, verifying result is not NaN.
		result := a.SlugsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to GramsPerLiter.
		// No expected conversion value provided for GramsPerLiter, verifying result is not NaN.
		result := a.GramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerDeciLiter.
		// No expected conversion value provided for GramsPerDeciLiter, verifying result is not NaN.
		result := a.GramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerMilliliter.
		// No expected conversion value provided for GramsPerMilliliter, verifying result is not NaN.
		result := a.GramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerUSGallon.
		// No expected conversion value provided for PoundsPerUSGallon, verifying result is not NaN.
		result := a.PoundsPerUSGallon()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerUSGallon returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerImperialGallon.
		// No expected conversion value provided for PoundsPerImperialGallon, verifying result is not NaN.
		result := a.PoundsPerImperialGallon()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerImperialGallon returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerLiter.
		// No expected conversion value provided for KilogramsPerLiter, verifying result is not NaN.
		result := a.KilogramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerCubicFoot.
		// No expected conversion value provided for TonnesPerCubicFoot, verifying result is not NaN.
		result := a.TonnesPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to TonnesPerCubicInch.
		// No expected conversion value provided for TonnesPerCubicInch, verifying result is not NaN.
		result := a.TonnesPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to GramsPerCubicFoot.
		// No expected conversion value provided for GramsPerCubicFoot, verifying result is not NaN.
		result := a.GramsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to GramsPerCubicInch.
		// No expected conversion value provided for GramsPerCubicInch, verifying result is not NaN.
		result := a.GramsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicMeter.
		// No expected conversion value provided for PoundsPerCubicMeter, verifying result is not NaN.
		result := a.PoundsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicCentimeter.
		// No expected conversion value provided for PoundsPerCubicCentimeter, verifying result is not NaN.
		result := a.PoundsPerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerCubicMillimeter.
		// No expected conversion value provided for PoundsPerCubicMillimeter, verifying result is not NaN.
		result := a.PoundsPerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to SlugsPerCubicMeter.
		// No expected conversion value provided for SlugsPerCubicMeter, verifying result is not NaN.
		result := a.SlugsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to SlugsPerCubicCentimeter.
		// No expected conversion value provided for SlugsPerCubicCentimeter, verifying result is not NaN.
		result := a.SlugsPerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to SlugsPerCubicMillimeter.
		// No expected conversion value provided for SlugsPerCubicMillimeter, verifying result is not NaN.
		result := a.SlugsPerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to SlugsPerCubicInch.
		// No expected conversion value provided for SlugsPerCubicInch, verifying result is not NaN.
		result := a.SlugsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to SlugsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerCubicMillimeter.
		// No expected conversion value provided for KilogramsPerCubicMillimeter, verifying result is not NaN.
		result := a.KilogramsPerCubicMillimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerCubicCentimeter.
		// No expected conversion value provided for KilogramsPerCubicCentimeter, verifying result is not NaN.
		result := a.KilogramsPerCubicCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerCubicMeter.
		// No expected conversion value provided for KilogramsPerCubicMeter, verifying result is not NaN.
		result := a.KilogramsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerCubicMeter.
		// No expected conversion value provided for MilligramsPerCubicMeter, verifying result is not NaN.
		result := a.MilligramsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerCubicMeter.
		// No expected conversion value provided for MicrogramsPerCubicMeter, verifying result is not NaN.
		result := a.MicrogramsPerCubicMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsPerCubicInch.
		// No expected conversion value provided for KilopoundsPerCubicInch, verifying result is not NaN.
		result := a.KilopoundsPerCubicInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsPerCubicFoot.
		// No expected conversion value provided for KilopoundsPerCubicFoot, verifying result is not NaN.
		result := a.KilopoundsPerCubicFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsPerCubicYard.
		// No expected conversion value provided for KilopoundsPerCubicYard, verifying result is not NaN.
		result := a.KilopoundsPerCubicYard()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsPerCubicYard returned NaN")
		}
	}
	{
		// Test conversion to FemtogramsPerLiter.
		// No expected conversion value provided for FemtogramsPerLiter, verifying result is not NaN.
		result := a.FemtogramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to FemtogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to PicogramsPerLiter.
		// No expected conversion value provided for PicogramsPerLiter, verifying result is not NaN.
		result := a.PicogramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerLiter.
		// No expected conversion value provided for NanogramsPerLiter, verifying result is not NaN.
		result := a.NanogramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerLiter.
		// No expected conversion value provided for MicrogramsPerLiter, verifying result is not NaN.
		result := a.MicrogramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerLiter.
		// No expected conversion value provided for MilligramsPerLiter, verifying result is not NaN.
		result := a.MilligramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerLiter.
		// No expected conversion value provided for CentigramsPerLiter, verifying result is not NaN.
		result := a.CentigramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerLiter.
		// No expected conversion value provided for DecigramsPerLiter, verifying result is not NaN.
		result := a.DecigramsPerLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigramsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to FemtogramsPerDeciLiter.
		// No expected conversion value provided for FemtogramsPerDeciLiter, verifying result is not NaN.
		result := a.FemtogramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to FemtogramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to PicogramsPerDeciLiter.
		// No expected conversion value provided for PicogramsPerDeciLiter, verifying result is not NaN.
		result := a.PicogramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicogramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerDeciLiter.
		// No expected conversion value provided for NanogramsPerDeciLiter, verifying result is not NaN.
		result := a.NanogramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanogramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerDeciLiter.
		// No expected conversion value provided for MicrogramsPerDeciLiter, verifying result is not NaN.
		result := a.MicrogramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerDeciLiter.
		// No expected conversion value provided for MilligramsPerDeciLiter, verifying result is not NaN.
		result := a.MilligramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerDeciLiter.
		// No expected conversion value provided for CentigramsPerDeciLiter, verifying result is not NaN.
		result := a.CentigramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerDeciLiter.
		// No expected conversion value provided for DecigramsPerDeciLiter, verifying result is not NaN.
		result := a.DecigramsPerDeciLiter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigramsPerDeciLiter returned NaN")
		}
	}
	{
		// Test conversion to FemtogramsPerMilliliter.
		// No expected conversion value provided for FemtogramsPerMilliliter, verifying result is not NaN.
		result := a.FemtogramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to FemtogramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to PicogramsPerMilliliter.
		// No expected conversion value provided for PicogramsPerMilliliter, verifying result is not NaN.
		result := a.PicogramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicogramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to NanogramsPerMilliliter.
		// No expected conversion value provided for NanogramsPerMilliliter, verifying result is not NaN.
		result := a.NanogramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanogramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to MicrogramsPerMilliliter.
		// No expected conversion value provided for MicrogramsPerMilliliter, verifying result is not NaN.
		result := a.MicrogramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrogramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerMilliliter.
		// No expected conversion value provided for MilligramsPerMilliliter, verifying result is not NaN.
		result := a.MilligramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilligramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to CentigramsPerMilliliter.
		// No expected conversion value provided for CentigramsPerMilliliter, verifying result is not NaN.
		result := a.CentigramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentigramsPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to DecigramsPerMilliliter.
		// No expected conversion value provided for DecigramsPerMilliliter, verifying result is not NaN.
		result := a.DecigramsPerMilliliter()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecigramsPerMilliliter returned NaN")
		}
	}
}

func TestDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.DensityFactory{}
	a, err := factory.CreateDensity(90, units.DensityKilogramPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.DensityKilogramPerCubicMeter {
		t.Errorf("expected default unit KilogramPerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.DensityGramPerCubicMillimeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.DensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.DensityKilogramPerCubicMeter {
		t.Errorf("expected unit KilogramPerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestDensityFactory_FromDto(t *testing.T) {
    factory := units.DensityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityKilogramPerCubicMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.DensityDto{
        Value: math.NaN(),
        Unit:  units.DensityKilogramPerCubicMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test GramPerCubicMillimeter conversion
    grams_per_cubic_millimeterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityGramPerCubicMillimeter,
    }
    
    var grams_per_cubic_millimeterResult *units.Density
    grams_per_cubic_millimeterResult, err = factory.FromDto(grams_per_cubic_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerCubicMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_cubic_millimeterResult.Convert(units.DensityGramPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test GramPerCubicCentimeter conversion
    grams_per_cubic_centimeterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityGramPerCubicCentimeter,
    }
    
    var grams_per_cubic_centimeterResult *units.Density
    grams_per_cubic_centimeterResult, err = factory.FromDto(grams_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_cubic_centimeterResult.Convert(units.DensityGramPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test GramPerCubicMeter conversion
    grams_per_cubic_meterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityGramPerCubicMeter,
    }
    
    var grams_per_cubic_meterResult *units.Density
    grams_per_cubic_meterResult, err = factory.FromDto(grams_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_cubic_meterResult.Convert(units.DensityGramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test PoundPerCubicInch conversion
    pounds_per_cubic_inchDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityPoundPerCubicInch,
    }
    
    var pounds_per_cubic_inchResult *units.Density
    pounds_per_cubic_inchResult, err = factory.FromDto(pounds_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_inchResult.Convert(units.DensityPoundPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicInch = %v, want %v", converted, 100)
    }
    // Test PoundPerCubicFoot conversion
    pounds_per_cubic_footDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityPoundPerCubicFoot,
    }
    
    var pounds_per_cubic_footResult *units.Density
    pounds_per_cubic_footResult, err = factory.FromDto(pounds_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_footResult.Convert(units.DensityPoundPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test PoundPerCubicYard conversion
    pounds_per_cubic_yardDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityPoundPerCubicYard,
    }
    
    var pounds_per_cubic_yardResult *units.Density
    pounds_per_cubic_yardResult, err = factory.FromDto(pounds_per_cubic_yardDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerCubicYard returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_yardResult.Convert(units.DensityPoundPerCubicYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicYard = %v, want %v", converted, 100)
    }
    // Test TonnePerCubicMillimeter conversion
    tonnes_per_cubic_millimeterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityTonnePerCubicMillimeter,
    }
    
    var tonnes_per_cubic_millimeterResult *units.Density
    tonnes_per_cubic_millimeterResult, err = factory.FromDto(tonnes_per_cubic_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with TonnePerCubicMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_millimeterResult.Convert(units.DensityTonnePerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test TonnePerCubicCentimeter conversion
    tonnes_per_cubic_centimeterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityTonnePerCubicCentimeter,
    }
    
    var tonnes_per_cubic_centimeterResult *units.Density
    tonnes_per_cubic_centimeterResult, err = factory.FromDto(tonnes_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with TonnePerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_centimeterResult.Convert(units.DensityTonnePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test TonnePerCubicMeter conversion
    tonnes_per_cubic_meterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityTonnePerCubicMeter,
    }
    
    var tonnes_per_cubic_meterResult *units.Density
    tonnes_per_cubic_meterResult, err = factory.FromDto(tonnes_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with TonnePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_meterResult.Convert(units.DensityTonnePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test SlugPerCubicFoot conversion
    slugs_per_cubic_footDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensitySlugPerCubicFoot,
    }
    
    var slugs_per_cubic_footResult *units.Density
    slugs_per_cubic_footResult, err = factory.FromDto(slugs_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with SlugPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugs_per_cubic_footResult.Convert(units.DensitySlugPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test GramPerLiter conversion
    grams_per_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityGramPerLiter,
    }
    
    var grams_per_literResult *units.Density
    grams_per_literResult, err = factory.FromDto(grams_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_literResult.Convert(units.DensityGramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerLiter = %v, want %v", converted, 100)
    }
    // Test GramPerDeciliter conversion
    grams_per_deci_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityGramPerDeciliter,
    }
    
    var grams_per_deci_literResult *units.Density
    grams_per_deci_literResult, err = factory.FromDto(grams_per_deci_literDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_deci_literResult.Convert(units.DensityGramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test GramPerMilliliter conversion
    grams_per_milliliterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityGramPerMilliliter,
    }
    
    var grams_per_milliliterResult *units.Density
    grams_per_milliliterResult, err = factory.FromDto(grams_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_milliliterResult.Convert(units.DensityGramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test PoundPerUSGallon conversion
    pounds_per_us_gallonDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityPoundPerUSGallon,
    }
    
    var pounds_per_us_gallonResult *units.Density
    pounds_per_us_gallonResult, err = factory.FromDto(pounds_per_us_gallonDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerUSGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_us_gallonResult.Convert(units.DensityPoundPerUSGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerUSGallon = %v, want %v", converted, 100)
    }
    // Test PoundPerImperialGallon conversion
    pounds_per_imperial_gallonDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityPoundPerImperialGallon,
    }
    
    var pounds_per_imperial_gallonResult *units.Density
    pounds_per_imperial_gallonResult, err = factory.FromDto(pounds_per_imperial_gallonDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerImperialGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_imperial_gallonResult.Convert(units.DensityPoundPerImperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerImperialGallon = %v, want %v", converted, 100)
    }
    // Test KilogramPerLiter conversion
    kilograms_per_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityKilogramPerLiter,
    }
    
    var kilograms_per_literResult *units.Density
    kilograms_per_literResult, err = factory.FromDto(kilograms_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_literResult.Convert(units.DensityKilogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerLiter = %v, want %v", converted, 100)
    }
    // Test TonnePerCubicFoot conversion
    tonnes_per_cubic_footDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityTonnePerCubicFoot,
    }
    
    var tonnes_per_cubic_footResult *units.Density
    tonnes_per_cubic_footResult, err = factory.FromDto(tonnes_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with TonnePerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_footResult.Convert(units.DensityTonnePerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicFoot = %v, want %v", converted, 100)
    }
    // Test TonnePerCubicInch conversion
    tonnes_per_cubic_inchDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityTonnePerCubicInch,
    }
    
    var tonnes_per_cubic_inchResult *units.Density
    tonnes_per_cubic_inchResult, err = factory.FromDto(tonnes_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with TonnePerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_inchResult.Convert(units.DensityTonnePerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicInch = %v, want %v", converted, 100)
    }
    // Test GramPerCubicFoot conversion
    grams_per_cubic_footDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityGramPerCubicFoot,
    }
    
    var grams_per_cubic_footResult *units.Density
    grams_per_cubic_footResult, err = factory.FromDto(grams_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_cubic_footResult.Convert(units.DensityGramPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test GramPerCubicInch conversion
    grams_per_cubic_inchDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityGramPerCubicInch,
    }
    
    var grams_per_cubic_inchResult *units.Density
    grams_per_cubic_inchResult, err = factory.FromDto(grams_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_cubic_inchResult.Convert(units.DensityGramPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicInch = %v, want %v", converted, 100)
    }
    // Test PoundPerCubicMeter conversion
    pounds_per_cubic_meterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityPoundPerCubicMeter,
    }
    
    var pounds_per_cubic_meterResult *units.Density
    pounds_per_cubic_meterResult, err = factory.FromDto(pounds_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_meterResult.Convert(units.DensityPoundPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test PoundPerCubicCentimeter conversion
    pounds_per_cubic_centimeterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityPoundPerCubicCentimeter,
    }
    
    var pounds_per_cubic_centimeterResult *units.Density
    pounds_per_cubic_centimeterResult, err = factory.FromDto(pounds_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_centimeterResult.Convert(units.DensityPoundPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test PoundPerCubicMillimeter conversion
    pounds_per_cubic_millimeterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityPoundPerCubicMillimeter,
    }
    
    var pounds_per_cubic_millimeterResult *units.Density
    pounds_per_cubic_millimeterResult, err = factory.FromDto(pounds_per_cubic_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerCubicMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_millimeterResult.Convert(units.DensityPoundPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test SlugPerCubicMeter conversion
    slugs_per_cubic_meterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensitySlugPerCubicMeter,
    }
    
    var slugs_per_cubic_meterResult *units.Density
    slugs_per_cubic_meterResult, err = factory.FromDto(slugs_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with SlugPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugs_per_cubic_meterResult.Convert(units.DensitySlugPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test SlugPerCubicCentimeter conversion
    slugs_per_cubic_centimeterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensitySlugPerCubicCentimeter,
    }
    
    var slugs_per_cubic_centimeterResult *units.Density
    slugs_per_cubic_centimeterResult, err = factory.FromDto(slugs_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with SlugPerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugs_per_cubic_centimeterResult.Convert(units.DensitySlugPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test SlugPerCubicMillimeter conversion
    slugs_per_cubic_millimeterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensitySlugPerCubicMillimeter,
    }
    
    var slugs_per_cubic_millimeterResult *units.Density
    slugs_per_cubic_millimeterResult, err = factory.FromDto(slugs_per_cubic_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with SlugPerCubicMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugs_per_cubic_millimeterResult.Convert(units.DensitySlugPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test SlugPerCubicInch conversion
    slugs_per_cubic_inchDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensitySlugPerCubicInch,
    }
    
    var slugs_per_cubic_inchResult *units.Density
    slugs_per_cubic_inchResult, err = factory.FromDto(slugs_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with SlugPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugs_per_cubic_inchResult.Convert(units.DensitySlugPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugPerCubicInch = %v, want %v", converted, 100)
    }
    // Test KilogramPerCubicMillimeter conversion
    kilograms_per_cubic_millimeterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityKilogramPerCubicMillimeter,
    }
    
    var kilograms_per_cubic_millimeterResult *units.Density
    kilograms_per_cubic_millimeterResult, err = factory.FromDto(kilograms_per_cubic_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerCubicMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_cubic_millimeterResult.Convert(units.DensityKilogramPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test KilogramPerCubicCentimeter conversion
    kilograms_per_cubic_centimeterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityKilogramPerCubicCentimeter,
    }
    
    var kilograms_per_cubic_centimeterResult *units.Density
    kilograms_per_cubic_centimeterResult, err = factory.FromDto(kilograms_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_cubic_centimeterResult.Convert(units.DensityKilogramPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test KilogramPerCubicMeter conversion
    kilograms_per_cubic_meterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityKilogramPerCubicMeter,
    }
    
    var kilograms_per_cubic_meterResult *units.Density
    kilograms_per_cubic_meterResult, err = factory.FromDto(kilograms_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_cubic_meterResult.Convert(units.DensityKilogramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test MilligramPerCubicMeter conversion
    milligrams_per_cubic_meterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityMilligramPerCubicMeter,
    }
    
    var milligrams_per_cubic_meterResult *units.Density
    milligrams_per_cubic_meterResult, err = factory.FromDto(milligrams_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_cubic_meterResult.Convert(units.DensityMilligramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test MicrogramPerCubicMeter conversion
    micrograms_per_cubic_meterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityMicrogramPerCubicMeter,
    }
    
    var micrograms_per_cubic_meterResult *units.Density
    micrograms_per_cubic_meterResult, err = factory.FromDto(micrograms_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_cubic_meterResult.Convert(units.DensityMicrogramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test KilopoundPerCubicInch conversion
    kilopounds_per_cubic_inchDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityKilopoundPerCubicInch,
    }
    
    var kilopounds_per_cubic_inchResult *units.Density
    kilopounds_per_cubic_inchResult, err = factory.FromDto(kilopounds_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_per_cubic_inchResult.Convert(units.DensityKilopoundPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundPerCubicInch = %v, want %v", converted, 100)
    }
    // Test KilopoundPerCubicFoot conversion
    kilopounds_per_cubic_footDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityKilopoundPerCubicFoot,
    }
    
    var kilopounds_per_cubic_footResult *units.Density
    kilopounds_per_cubic_footResult, err = factory.FromDto(kilopounds_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_per_cubic_footResult.Convert(units.DensityKilopoundPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test KilopoundPerCubicYard conversion
    kilopounds_per_cubic_yardDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityKilopoundPerCubicYard,
    }
    
    var kilopounds_per_cubic_yardResult *units.Density
    kilopounds_per_cubic_yardResult, err = factory.FromDto(kilopounds_per_cubic_yardDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundPerCubicYard returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_per_cubic_yardResult.Convert(units.DensityKilopoundPerCubicYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundPerCubicYard = %v, want %v", converted, 100)
    }
    // Test FemtogramPerLiter conversion
    femtograms_per_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityFemtogramPerLiter,
    }
    
    var femtograms_per_literResult *units.Density
    femtograms_per_literResult, err = factory.FromDto(femtograms_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with FemtogramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtograms_per_literResult.Convert(units.DensityFemtogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FemtogramPerLiter = %v, want %v", converted, 100)
    }
    // Test PicogramPerLiter conversion
    picograms_per_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityPicogramPerLiter,
    }
    
    var picograms_per_literResult *units.Density
    picograms_per_literResult, err = factory.FromDto(picograms_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with PicogramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_literResult.Convert(units.DensityPicogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicogramPerLiter = %v, want %v", converted, 100)
    }
    // Test NanogramPerLiter conversion
    nanograms_per_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityNanogramPerLiter,
    }
    
    var nanograms_per_literResult *units.Density
    nanograms_per_literResult, err = factory.FromDto(nanograms_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with NanogramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_literResult.Convert(units.DensityNanogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerLiter = %v, want %v", converted, 100)
    }
    // Test MicrogramPerLiter conversion
    micrograms_per_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityMicrogramPerLiter,
    }
    
    var micrograms_per_literResult *units.Density
    micrograms_per_literResult, err = factory.FromDto(micrograms_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_literResult.Convert(units.DensityMicrogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerLiter = %v, want %v", converted, 100)
    }
    // Test MilligramPerLiter conversion
    milligrams_per_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityMilligramPerLiter,
    }
    
    var milligrams_per_literResult *units.Density
    milligrams_per_literResult, err = factory.FromDto(milligrams_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_literResult.Convert(units.DensityMilligramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerLiter = %v, want %v", converted, 100)
    }
    // Test CentigramPerLiter conversion
    centigrams_per_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityCentigramPerLiter,
    }
    
    var centigrams_per_literResult *units.Density
    centigrams_per_literResult, err = factory.FromDto(centigrams_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with CentigramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_literResult.Convert(units.DensityCentigramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerLiter = %v, want %v", converted, 100)
    }
    // Test DecigramPerLiter conversion
    decigrams_per_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityDecigramPerLiter,
    }
    
    var decigrams_per_literResult *units.Density
    decigrams_per_literResult, err = factory.FromDto(decigrams_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with DecigramPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_literResult.Convert(units.DensityDecigramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerLiter = %v, want %v", converted, 100)
    }
    // Test FemtogramPerDeciliter conversion
    femtograms_per_deci_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityFemtogramPerDeciliter,
    }
    
    var femtograms_per_deci_literResult *units.Density
    femtograms_per_deci_literResult, err = factory.FromDto(femtograms_per_deci_literDto)
    if err != nil {
        t.Errorf("FromDto() with FemtogramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtograms_per_deci_literResult.Convert(units.DensityFemtogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FemtogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test PicogramPerDeciliter conversion
    picograms_per_deci_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityPicogramPerDeciliter,
    }
    
    var picograms_per_deci_literResult *units.Density
    picograms_per_deci_literResult, err = factory.FromDto(picograms_per_deci_literDto)
    if err != nil {
        t.Errorf("FromDto() with PicogramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_deci_literResult.Convert(units.DensityPicogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test NanogramPerDeciliter conversion
    nanograms_per_deci_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityNanogramPerDeciliter,
    }
    
    var nanograms_per_deci_literResult *units.Density
    nanograms_per_deci_literResult, err = factory.FromDto(nanograms_per_deci_literDto)
    if err != nil {
        t.Errorf("FromDto() with NanogramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_deci_literResult.Convert(units.DensityNanogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test MicrogramPerDeciliter conversion
    micrograms_per_deci_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityMicrogramPerDeciliter,
    }
    
    var micrograms_per_deci_literResult *units.Density
    micrograms_per_deci_literResult, err = factory.FromDto(micrograms_per_deci_literDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_deci_literResult.Convert(units.DensityMicrogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test MilligramPerDeciliter conversion
    milligrams_per_deci_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityMilligramPerDeciliter,
    }
    
    var milligrams_per_deci_literResult *units.Density
    milligrams_per_deci_literResult, err = factory.FromDto(milligrams_per_deci_literDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_deci_literResult.Convert(units.DensityMilligramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test CentigramPerDeciliter conversion
    centigrams_per_deci_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityCentigramPerDeciliter,
    }
    
    var centigrams_per_deci_literResult *units.Density
    centigrams_per_deci_literResult, err = factory.FromDto(centigrams_per_deci_literDto)
    if err != nil {
        t.Errorf("FromDto() with CentigramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_deci_literResult.Convert(units.DensityCentigramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test DecigramPerDeciliter conversion
    decigrams_per_deci_literDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityDecigramPerDeciliter,
    }
    
    var decigrams_per_deci_literResult *units.Density
    decigrams_per_deci_literResult, err = factory.FromDto(decigrams_per_deci_literDto)
    if err != nil {
        t.Errorf("FromDto() with DecigramPerDeciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_deci_literResult.Convert(units.DensityDecigramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test FemtogramPerMilliliter conversion
    femtograms_per_milliliterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityFemtogramPerMilliliter,
    }
    
    var femtograms_per_milliliterResult *units.Density
    femtograms_per_milliliterResult, err = factory.FromDto(femtograms_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with FemtogramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtograms_per_milliliterResult.Convert(units.DensityFemtogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FemtogramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test PicogramPerMilliliter conversion
    picograms_per_milliliterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityPicogramPerMilliliter,
    }
    
    var picograms_per_milliliterResult *units.Density
    picograms_per_milliliterResult, err = factory.FromDto(picograms_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with PicogramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_milliliterResult.Convert(units.DensityPicogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicogramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test NanogramPerMilliliter conversion
    nanograms_per_milliliterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityNanogramPerMilliliter,
    }
    
    var nanograms_per_milliliterResult *units.Density
    nanograms_per_milliliterResult, err = factory.FromDto(nanograms_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with NanogramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_milliliterResult.Convert(units.DensityNanogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test MicrogramPerMilliliter conversion
    micrograms_per_milliliterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityMicrogramPerMilliliter,
    }
    
    var micrograms_per_milliliterResult *units.Density
    micrograms_per_milliliterResult, err = factory.FromDto(micrograms_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrogramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_milliliterResult.Convert(units.DensityMicrogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test MilligramPerMilliliter conversion
    milligrams_per_milliliterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityMilligramPerMilliliter,
    }
    
    var milligrams_per_milliliterResult *units.Density
    milligrams_per_milliliterResult, err = factory.FromDto(milligrams_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_milliliterResult.Convert(units.DensityMilligramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test CentigramPerMilliliter conversion
    centigrams_per_milliliterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityCentigramPerMilliliter,
    }
    
    var centigrams_per_milliliterResult *units.Density
    centigrams_per_milliliterResult, err = factory.FromDto(centigrams_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with CentigramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_milliliterResult.Convert(units.DensityCentigramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test DecigramPerMilliliter conversion
    decigrams_per_milliliterDto := units.DensityDto{
        Value: 100,
        Unit:  units.DensityDecigramPerMilliliter,
    }
    
    var decigrams_per_milliliterResult *units.Density
    decigrams_per_milliliterResult, err = factory.FromDto(decigrams_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with DecigramPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_milliliterResult.Convert(units.DensityDecigramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerMilliliter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.DensityDto{
        Value: 0,
        Unit:  units.DensityKilogramPerCubicMeter,
    }
    
    var zeroResult *units.Density
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestDensityFactory_FromDtoJSON(t *testing.T) {
    factory := units.DensityFactory{}
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
    nanJSON, _ := json.Marshal(units.DensityDto{
        Value: nanValue,
        Unit:  units.DensityKilogramPerCubicMeter,
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
    converted = grams_per_cubic_millimeterResult.Convert(units.DensityGramPerCubicMillimeter)
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
    converted = grams_per_cubic_centimeterResult.Convert(units.DensityGramPerCubicCentimeter)
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
    converted = grams_per_cubic_meterResult.Convert(units.DensityGramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerCubicInch unit
    pounds_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "PoundPerCubicInch"}`)
    pounds_per_cubic_inchResult, err := factory.FromDtoJSON(pounds_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_inchResult.Convert(units.DensityPoundPerCubicInch)
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
    converted = pounds_per_cubic_footResult.Convert(units.DensityPoundPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerCubicYard unit
    pounds_per_cubic_yardJSON := []byte(`{"value": 100, "unit": "PoundPerCubicYard"}`)
    pounds_per_cubic_yardResult, err := factory.FromDtoJSON(pounds_per_cubic_yardJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerCubicYard unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_yardResult.Convert(units.DensityPoundPerCubicYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicYard = %v, want %v", converted, 100)
    }
    // Test JSON with TonnePerCubicMillimeter unit
    tonnes_per_cubic_millimeterJSON := []byte(`{"value": 100, "unit": "TonnePerCubicMillimeter"}`)
    tonnes_per_cubic_millimeterResult, err := factory.FromDtoJSON(tonnes_per_cubic_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonnePerCubicMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_millimeterResult.Convert(units.DensityTonnePerCubicMillimeter)
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
    converted = tonnes_per_cubic_centimeterResult.Convert(units.DensityTonnePerCubicCentimeter)
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
    converted = tonnes_per_cubic_meterResult.Convert(units.DensityTonnePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with SlugPerCubicFoot unit
    slugs_per_cubic_footJSON := []byte(`{"value": 100, "unit": "SlugPerCubicFoot"}`)
    slugs_per_cubic_footResult, err := factory.FromDtoJSON(slugs_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SlugPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugs_per_cubic_footResult.Convert(units.DensitySlugPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerLiter unit
    grams_per_literJSON := []byte(`{"value": 100, "unit": "GramPerLiter"}`)
    grams_per_literResult, err := factory.FromDtoJSON(grams_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_literResult.Convert(units.DensityGramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerDeciliter unit
    grams_per_deci_literJSON := []byte(`{"value": 100, "unit": "GramPerDeciliter"}`)
    grams_per_deci_literResult, err := factory.FromDtoJSON(grams_per_deci_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_deci_literResult.Convert(units.DensityGramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerMilliliter unit
    grams_per_milliliterJSON := []byte(`{"value": 100, "unit": "GramPerMilliliter"}`)
    grams_per_milliliterResult, err := factory.FromDtoJSON(grams_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_milliliterResult.Convert(units.DensityGramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerUSGallon unit
    pounds_per_us_gallonJSON := []byte(`{"value": 100, "unit": "PoundPerUSGallon"}`)
    pounds_per_us_gallonResult, err := factory.FromDtoJSON(pounds_per_us_gallonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerUSGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_us_gallonResult.Convert(units.DensityPoundPerUSGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerUSGallon = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerImperialGallon unit
    pounds_per_imperial_gallonJSON := []byte(`{"value": 100, "unit": "PoundPerImperialGallon"}`)
    pounds_per_imperial_gallonResult, err := factory.FromDtoJSON(pounds_per_imperial_gallonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerImperialGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_imperial_gallonResult.Convert(units.DensityPoundPerImperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerImperialGallon = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerLiter unit
    kilograms_per_literJSON := []byte(`{"value": 100, "unit": "KilogramPerLiter"}`)
    kilograms_per_literResult, err := factory.FromDtoJSON(kilograms_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_literResult.Convert(units.DensityKilogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with TonnePerCubicFoot unit
    tonnes_per_cubic_footJSON := []byte(`{"value": 100, "unit": "TonnePerCubicFoot"}`)
    tonnes_per_cubic_footResult, err := factory.FromDtoJSON(tonnes_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonnePerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_footResult.Convert(units.DensityTonnePerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with TonnePerCubicInch unit
    tonnes_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "TonnePerCubicInch"}`)
    tonnes_per_cubic_inchResult, err := factory.FromDtoJSON(tonnes_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonnePerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_per_cubic_inchResult.Convert(units.DensityTonnePerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonnePerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerCubicFoot unit
    grams_per_cubic_footJSON := []byte(`{"value": 100, "unit": "GramPerCubicFoot"}`)
    grams_per_cubic_footResult, err := factory.FromDtoJSON(grams_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_cubic_footResult.Convert(units.DensityGramPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerCubicInch unit
    grams_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "GramPerCubicInch"}`)
    grams_per_cubic_inchResult, err := factory.FromDtoJSON(grams_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_cubic_inchResult.Convert(units.DensityGramPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerCubicMeter unit
    pounds_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "PoundPerCubicMeter"}`)
    pounds_per_cubic_meterResult, err := factory.FromDtoJSON(pounds_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_meterResult.Convert(units.DensityPoundPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerCubicCentimeter unit
    pounds_per_cubic_centimeterJSON := []byte(`{"value": 100, "unit": "PoundPerCubicCentimeter"}`)
    pounds_per_cubic_centimeterResult, err := factory.FromDtoJSON(pounds_per_cubic_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerCubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_centimeterResult.Convert(units.DensityPoundPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerCubicMillimeter unit
    pounds_per_cubic_millimeterJSON := []byte(`{"value": 100, "unit": "PoundPerCubicMillimeter"}`)
    pounds_per_cubic_millimeterResult, err := factory.FromDtoJSON(pounds_per_cubic_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerCubicMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_cubic_millimeterResult.Convert(units.DensityPoundPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with SlugPerCubicMeter unit
    slugs_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "SlugPerCubicMeter"}`)
    slugs_per_cubic_meterResult, err := factory.FromDtoJSON(slugs_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SlugPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugs_per_cubic_meterResult.Convert(units.DensitySlugPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with SlugPerCubicCentimeter unit
    slugs_per_cubic_centimeterJSON := []byte(`{"value": 100, "unit": "SlugPerCubicCentimeter"}`)
    slugs_per_cubic_centimeterResult, err := factory.FromDtoJSON(slugs_per_cubic_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SlugPerCubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugs_per_cubic_centimeterResult.Convert(units.DensitySlugPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with SlugPerCubicMillimeter unit
    slugs_per_cubic_millimeterJSON := []byte(`{"value": 100, "unit": "SlugPerCubicMillimeter"}`)
    slugs_per_cubic_millimeterResult, err := factory.FromDtoJSON(slugs_per_cubic_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SlugPerCubicMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugs_per_cubic_millimeterResult.Convert(units.DensitySlugPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with SlugPerCubicInch unit
    slugs_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "SlugPerCubicInch"}`)
    slugs_per_cubic_inchResult, err := factory.FromDtoJSON(slugs_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SlugPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = slugs_per_cubic_inchResult.Convert(units.DensitySlugPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SlugPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerCubicMillimeter unit
    kilograms_per_cubic_millimeterJSON := []byte(`{"value": 100, "unit": "KilogramPerCubicMillimeter"}`)
    kilograms_per_cubic_millimeterResult, err := factory.FromDtoJSON(kilograms_per_cubic_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerCubicMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_cubic_millimeterResult.Convert(units.DensityKilogramPerCubicMillimeter)
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
    converted = kilograms_per_cubic_centimeterResult.Convert(units.DensityKilogramPerCubicCentimeter)
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
    converted = kilograms_per_cubic_meterResult.Convert(units.DensityKilogramPerCubicMeter)
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
    converted = milligrams_per_cubic_meterResult.Convert(units.DensityMilligramPerCubicMeter)
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
    converted = micrograms_per_cubic_meterResult.Convert(units.DensityMicrogramPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundPerCubicInch unit
    kilopounds_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "KilopoundPerCubicInch"}`)
    kilopounds_per_cubic_inchResult, err := factory.FromDtoJSON(kilopounds_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_per_cubic_inchResult.Convert(units.DensityKilopoundPerCubicInch)
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
    converted = kilopounds_per_cubic_footResult.Convert(units.DensityKilopoundPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundPerCubicYard unit
    kilopounds_per_cubic_yardJSON := []byte(`{"value": 100, "unit": "KilopoundPerCubicYard"}`)
    kilopounds_per_cubic_yardResult, err := factory.FromDtoJSON(kilopounds_per_cubic_yardJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundPerCubicYard unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_per_cubic_yardResult.Convert(units.DensityKilopoundPerCubicYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundPerCubicYard = %v, want %v", converted, 100)
    }
    // Test JSON with FemtogramPerLiter unit
    femtograms_per_literJSON := []byte(`{"value": 100, "unit": "FemtogramPerLiter"}`)
    femtograms_per_literResult, err := factory.FromDtoJSON(femtograms_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FemtogramPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtograms_per_literResult.Convert(units.DensityFemtogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FemtogramPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with PicogramPerLiter unit
    picograms_per_literJSON := []byte(`{"value": 100, "unit": "PicogramPerLiter"}`)
    picograms_per_literResult, err := factory.FromDtoJSON(picograms_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicogramPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_literResult.Convert(units.DensityPicogramPerLiter)
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
    converted = nanograms_per_literResult.Convert(units.DensityNanogramPerLiter)
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
    converted = micrograms_per_literResult.Convert(units.DensityMicrogramPerLiter)
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
    converted = milligrams_per_literResult.Convert(units.DensityMilligramPerLiter)
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
    converted = centigrams_per_literResult.Convert(units.DensityCentigramPerLiter)
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
    converted = decigrams_per_literResult.Convert(units.DensityDecigramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with FemtogramPerDeciliter unit
    femtograms_per_deci_literJSON := []byte(`{"value": 100, "unit": "FemtogramPerDeciliter"}`)
    femtograms_per_deci_literResult, err := factory.FromDtoJSON(femtograms_per_deci_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FemtogramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtograms_per_deci_literResult.Convert(units.DensityFemtogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FemtogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with PicogramPerDeciliter unit
    picograms_per_deci_literJSON := []byte(`{"value": 100, "unit": "PicogramPerDeciliter"}`)
    picograms_per_deci_literResult, err := factory.FromDtoJSON(picograms_per_deci_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicogramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_deci_literResult.Convert(units.DensityPicogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with NanogramPerDeciliter unit
    nanograms_per_deci_literJSON := []byte(`{"value": 100, "unit": "NanogramPerDeciliter"}`)
    nanograms_per_deci_literResult, err := factory.FromDtoJSON(nanograms_per_deci_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanogramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanograms_per_deci_literResult.Convert(units.DensityNanogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrogramPerDeciliter unit
    micrograms_per_deci_literJSON := []byte(`{"value": 100, "unit": "MicrogramPerDeciliter"}`)
    micrograms_per_deci_literResult, err := factory.FromDtoJSON(micrograms_per_deci_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrogramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrograms_per_deci_literResult.Convert(units.DensityMicrogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrogramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerDeciliter unit
    milligrams_per_deci_literJSON := []byte(`{"value": 100, "unit": "MilligramPerDeciliter"}`)
    milligrams_per_deci_literResult, err := factory.FromDtoJSON(milligrams_per_deci_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_deci_literResult.Convert(units.DensityMilligramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with CentigramPerDeciliter unit
    centigrams_per_deci_literJSON := []byte(`{"value": 100, "unit": "CentigramPerDeciliter"}`)
    centigrams_per_deci_literResult, err := factory.FromDtoJSON(centigrams_per_deci_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentigramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centigrams_per_deci_literResult.Convert(units.DensityCentigramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentigramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with DecigramPerDeciliter unit
    decigrams_per_deci_literJSON := []byte(`{"value": 100, "unit": "DecigramPerDeciliter"}`)
    decigrams_per_deci_literResult, err := factory.FromDtoJSON(decigrams_per_deci_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecigramPerDeciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decigrams_per_deci_literResult.Convert(units.DensityDecigramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerDeciliter = %v, want %v", converted, 100)
    }
    // Test JSON with FemtogramPerMilliliter unit
    femtograms_per_milliliterJSON := []byte(`{"value": 100, "unit": "FemtogramPerMilliliter"}`)
    femtograms_per_milliliterResult, err := factory.FromDtoJSON(femtograms_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FemtogramPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtograms_per_milliliterResult.Convert(units.DensityFemtogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FemtogramPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with PicogramPerMilliliter unit
    picograms_per_milliliterJSON := []byte(`{"value": 100, "unit": "PicogramPerMilliliter"}`)
    picograms_per_milliliterResult, err := factory.FromDtoJSON(picograms_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicogramPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picograms_per_milliliterResult.Convert(units.DensityPicogramPerMilliliter)
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
    converted = nanograms_per_milliliterResult.Convert(units.DensityNanogramPerMilliliter)
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
    converted = micrograms_per_milliliterResult.Convert(units.DensityMicrogramPerMilliliter)
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
    converted = milligrams_per_milliliterResult.Convert(units.DensityMilligramPerMilliliter)
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
    converted = centigrams_per_milliliterResult.Convert(units.DensityCentigramPerMilliliter)
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
    converted = decigrams_per_milliliterResult.Convert(units.DensityDecigramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecigramPerMilliliter = %v, want %v", converted, 100)
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
func TestDensityFactory_FromGramsPerCubicMillimeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerCubicMillimeter(100)
    if err != nil {
        t.Errorf("FromGramsPerCubicMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityGramPerCubicMillimeter)
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
    converted = zeroResult.Convert(units.DensityGramPerCubicMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerCubicMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerCubicCentimeter function
func TestDensityFactory_FromGramsPerCubicCentimeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromGramsPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityGramPerCubicCentimeter)
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
    converted = zeroResult.Convert(units.DensityGramPerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerCubicMeter function
func TestDensityFactory_FromGramsPerCubicMeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromGramsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityGramPerCubicMeter)
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
    converted = zeroResult.Convert(units.DensityGramPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerCubicInch function
func TestDensityFactory_FromPoundsPerCubicInch(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromPoundsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityPoundPerCubicInch)
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
    converted = zeroResult.Convert(units.DensityPoundPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerCubicFoot function
func TestDensityFactory_FromPoundsPerCubicFoot(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromPoundsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityPoundPerCubicFoot)
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
    converted = zeroResult.Convert(units.DensityPoundPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerCubicYard function
func TestDensityFactory_FromPoundsPerCubicYard(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerCubicYard(100)
    if err != nil {
        t.Errorf("FromPoundsPerCubicYard() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityPoundPerCubicYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerCubicYard() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerCubicYard(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerCubicYard() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerCubicYard(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerCubicYard() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerCubicYard(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerCubicYard() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerCubicYard(0)
    if err != nil {
        t.Errorf("FromPoundsPerCubicYard() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityPoundPerCubicYard)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerCubicYard() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesPerCubicMillimeter function
func TestDensityFactory_FromTonnesPerCubicMillimeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesPerCubicMillimeter(100)
    if err != nil {
        t.Errorf("FromTonnesPerCubicMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityTonnePerCubicMillimeter)
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
    converted = zeroResult.Convert(units.DensityTonnePerCubicMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesPerCubicMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesPerCubicCentimeter function
func TestDensityFactory_FromTonnesPerCubicCentimeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromTonnesPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityTonnePerCubicCentimeter)
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
    converted = zeroResult.Convert(units.DensityTonnePerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesPerCubicMeter function
func TestDensityFactory_FromTonnesPerCubicMeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromTonnesPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityTonnePerCubicMeter)
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
    converted = zeroResult.Convert(units.DensityTonnePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromSlugsPerCubicFoot function
func TestDensityFactory_FromSlugsPerCubicFoot(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSlugsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromSlugsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensitySlugPerCubicFoot)
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
    converted = zeroResult.Convert(units.DensitySlugPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSlugsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerLiter function
func TestDensityFactory_FromGramsPerLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerLiter(100)
    if err != nil {
        t.Errorf("FromGramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityGramPerLiter)
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
    converted = zeroResult.Convert(units.DensityGramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerDeciLiter function
func TestDensityFactory_FromGramsPerDeciLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerDeciLiter(100)
    if err != nil {
        t.Errorf("FromGramsPerDeciLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityGramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerDeciLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerDeciLiter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerDeciLiter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerDeciLiter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerDeciLiter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerDeciLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerDeciLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerDeciLiter(0)
    if err != nil {
        t.Errorf("FromGramsPerDeciLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityGramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerDeciLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerMilliliter function
func TestDensityFactory_FromGramsPerMilliliter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromGramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityGramPerMilliliter)
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
    converted = zeroResult.Convert(units.DensityGramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerUSGallon function
func TestDensityFactory_FromPoundsPerUSGallon(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerUSGallon(100)
    if err != nil {
        t.Errorf("FromPoundsPerUSGallon() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityPoundPerUSGallon)
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
    converted = zeroResult.Convert(units.DensityPoundPerUSGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerUSGallon() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerImperialGallon function
func TestDensityFactory_FromPoundsPerImperialGallon(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerImperialGallon(100)
    if err != nil {
        t.Errorf("FromPoundsPerImperialGallon() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityPoundPerImperialGallon)
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
    converted = zeroResult.Convert(units.DensityPoundPerImperialGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerImperialGallon() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerLiter function
func TestDensityFactory_FromKilogramsPerLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerLiter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityKilogramPerLiter)
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
    converted = zeroResult.Convert(units.DensityKilogramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesPerCubicFoot function
func TestDensityFactory_FromTonnesPerCubicFoot(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromTonnesPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityTonnePerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromTonnesPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromTonnesPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromTonnesPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromTonnesPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityTonnePerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesPerCubicInch function
func TestDensityFactory_FromTonnesPerCubicInch(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesPerCubicInch(100)
    if err != nil {
        t.Errorf("FromTonnesPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityTonnePerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromTonnesPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromTonnesPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromTonnesPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesPerCubicInch(0)
    if err != nil {
        t.Errorf("FromTonnesPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityTonnePerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerCubicFoot function
func TestDensityFactory_FromGramsPerCubicFoot(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromGramsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityGramPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromGramsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromGramsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromGramsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityGramPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerCubicInch function
func TestDensityFactory_FromGramsPerCubicInch(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromGramsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityGramPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromGramsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromGramsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromGramsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityGramPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerCubicMeter function
func TestDensityFactory_FromPoundsPerCubicMeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromPoundsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityPoundPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromPoundsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityPoundPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerCubicCentimeter function
func TestDensityFactory_FromPoundsPerCubicCentimeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromPoundsPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityPoundPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerCubicCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerCubicCentimeter(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerCubicCentimeter() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerCubicCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerCubicCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerCubicCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerCubicCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerCubicCentimeter(0)
    if err != nil {
        t.Errorf("FromPoundsPerCubicCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityPoundPerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerCubicMillimeter function
func TestDensityFactory_FromPoundsPerCubicMillimeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerCubicMillimeter(100)
    if err != nil {
        t.Errorf("FromPoundsPerCubicMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityPoundPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerCubicMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerCubicMillimeter(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerCubicMillimeter() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerCubicMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerCubicMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerCubicMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerCubicMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerCubicMillimeter(0)
    if err != nil {
        t.Errorf("FromPoundsPerCubicMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityPoundPerCubicMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerCubicMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromSlugsPerCubicMeter function
func TestDensityFactory_FromSlugsPerCubicMeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSlugsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromSlugsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensitySlugPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSlugsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSlugsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromSlugsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromSlugsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromSlugsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromSlugsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromSlugsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSlugsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromSlugsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensitySlugPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSlugsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromSlugsPerCubicCentimeter function
func TestDensityFactory_FromSlugsPerCubicCentimeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSlugsPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromSlugsPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensitySlugPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSlugsPerCubicCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSlugsPerCubicCentimeter(math.NaN())
    if err == nil {
        t.Error("FromSlugsPerCubicCentimeter() with NaN value should return error")
    }

    _, err = factory.FromSlugsPerCubicCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromSlugsPerCubicCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromSlugsPerCubicCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromSlugsPerCubicCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSlugsPerCubicCentimeter(0)
    if err != nil {
        t.Errorf("FromSlugsPerCubicCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensitySlugPerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSlugsPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromSlugsPerCubicMillimeter function
func TestDensityFactory_FromSlugsPerCubicMillimeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSlugsPerCubicMillimeter(100)
    if err != nil {
        t.Errorf("FromSlugsPerCubicMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensitySlugPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSlugsPerCubicMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSlugsPerCubicMillimeter(math.NaN())
    if err == nil {
        t.Error("FromSlugsPerCubicMillimeter() with NaN value should return error")
    }

    _, err = factory.FromSlugsPerCubicMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromSlugsPerCubicMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromSlugsPerCubicMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromSlugsPerCubicMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSlugsPerCubicMillimeter(0)
    if err != nil {
        t.Errorf("FromSlugsPerCubicMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensitySlugPerCubicMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSlugsPerCubicMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromSlugsPerCubicInch function
func TestDensityFactory_FromSlugsPerCubicInch(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSlugsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromSlugsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensitySlugPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSlugsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSlugsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromSlugsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromSlugsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromSlugsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromSlugsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromSlugsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSlugsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromSlugsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensitySlugPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSlugsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerCubicMillimeter function
func TestDensityFactory_FromKilogramsPerCubicMillimeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerCubicMillimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerCubicMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityKilogramPerCubicMillimeter)
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
    converted = zeroResult.Convert(units.DensityKilogramPerCubicMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerCubicMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerCubicCentimeter function
func TestDensityFactory_FromKilogramsPerCubicCentimeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityKilogramPerCubicCentimeter)
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
    converted = zeroResult.Convert(units.DensityKilogramPerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerCubicMeter function
func TestDensityFactory_FromKilogramsPerCubicMeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityKilogramPerCubicMeter)
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
    converted = zeroResult.Convert(units.DensityKilogramPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerCubicMeter function
func TestDensityFactory_FromMilligramsPerCubicMeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMilligramsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityMilligramPerCubicMeter)
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
    converted = zeroResult.Convert(units.DensityMilligramPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerCubicMeter function
func TestDensityFactory_FromMicrogramsPerCubicMeter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityMicrogramPerCubicMeter)
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
    converted = zeroResult.Convert(units.DensityMicrogramPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsPerCubicInch function
func TestDensityFactory_FromKilopoundsPerCubicInch(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromKilopoundsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityKilopoundPerCubicInch)
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
    converted = zeroResult.Convert(units.DensityKilopoundPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsPerCubicFoot function
func TestDensityFactory_FromKilopoundsPerCubicFoot(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromKilopoundsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityKilopoundPerCubicFoot)
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
    converted = zeroResult.Convert(units.DensityKilopoundPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsPerCubicYard function
func TestDensityFactory_FromKilopoundsPerCubicYard(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsPerCubicYard(100)
    if err != nil {
        t.Errorf("FromKilopoundsPerCubicYard() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityKilopoundPerCubicYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsPerCubicYard() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsPerCubicYard(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsPerCubicYard() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsPerCubicYard(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsPerCubicYard() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsPerCubicYard(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsPerCubicYard() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsPerCubicYard(0)
    if err != nil {
        t.Errorf("FromKilopoundsPerCubicYard() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityKilopoundPerCubicYard)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsPerCubicYard() with zero value = %v, want 0", converted)
    }
}
// Test FromFemtogramsPerLiter function
func TestDensityFactory_FromFemtogramsPerLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFemtogramsPerLiter(100)
    if err != nil {
        t.Errorf("FromFemtogramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityFemtogramPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFemtogramsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFemtogramsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromFemtogramsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromFemtogramsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromFemtogramsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromFemtogramsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromFemtogramsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFemtogramsPerLiter(0)
    if err != nil {
        t.Errorf("FromFemtogramsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityFemtogramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFemtogramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicogramsPerLiter function
func TestDensityFactory_FromPicogramsPerLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicogramsPerLiter(100)
    if err != nil {
        t.Errorf("FromPicogramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityPicogramPerLiter)
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
    converted = zeroResult.Convert(units.DensityPicogramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicogramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanogramsPerLiter function
func TestDensityFactory_FromNanogramsPerLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanogramsPerLiter(100)
    if err != nil {
        t.Errorf("FromNanogramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityNanogramPerLiter)
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
    converted = zeroResult.Convert(units.DensityNanogramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanogramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerLiter function
func TestDensityFactory_FromMicrogramsPerLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerLiter(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityMicrogramPerLiter)
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
    converted = zeroResult.Convert(units.DensityMicrogramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerLiter function
func TestDensityFactory_FromMilligramsPerLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerLiter(100)
    if err != nil {
        t.Errorf("FromMilligramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityMilligramPerLiter)
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
    converted = zeroResult.Convert(units.DensityMilligramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigramsPerLiter function
func TestDensityFactory_FromCentigramsPerLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigramsPerLiter(100)
    if err != nil {
        t.Errorf("FromCentigramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityCentigramPerLiter)
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
    converted = zeroResult.Convert(units.DensityCentigramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigramsPerLiter function
func TestDensityFactory_FromDecigramsPerLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigramsPerLiter(100)
    if err != nil {
        t.Errorf("FromDecigramsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityDecigramPerLiter)
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
    converted = zeroResult.Convert(units.DensityDecigramPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigramsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromFemtogramsPerDeciLiter function
func TestDensityFactory_FromFemtogramsPerDeciLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFemtogramsPerDeciLiter(100)
    if err != nil {
        t.Errorf("FromFemtogramsPerDeciLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityFemtogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFemtogramsPerDeciLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFemtogramsPerDeciLiter(math.NaN())
    if err == nil {
        t.Error("FromFemtogramsPerDeciLiter() with NaN value should return error")
    }

    _, err = factory.FromFemtogramsPerDeciLiter(math.Inf(1))
    if err == nil {
        t.Error("FromFemtogramsPerDeciLiter() with +Inf value should return error")
    }

    _, err = factory.FromFemtogramsPerDeciLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromFemtogramsPerDeciLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFemtogramsPerDeciLiter(0)
    if err != nil {
        t.Errorf("FromFemtogramsPerDeciLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityFemtogramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFemtogramsPerDeciLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicogramsPerDeciLiter function
func TestDensityFactory_FromPicogramsPerDeciLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicogramsPerDeciLiter(100)
    if err != nil {
        t.Errorf("FromPicogramsPerDeciLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityPicogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicogramsPerDeciLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicogramsPerDeciLiter(math.NaN())
    if err == nil {
        t.Error("FromPicogramsPerDeciLiter() with NaN value should return error")
    }

    _, err = factory.FromPicogramsPerDeciLiter(math.Inf(1))
    if err == nil {
        t.Error("FromPicogramsPerDeciLiter() with +Inf value should return error")
    }

    _, err = factory.FromPicogramsPerDeciLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicogramsPerDeciLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicogramsPerDeciLiter(0)
    if err != nil {
        t.Errorf("FromPicogramsPerDeciLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityPicogramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicogramsPerDeciLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanogramsPerDeciLiter function
func TestDensityFactory_FromNanogramsPerDeciLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanogramsPerDeciLiter(100)
    if err != nil {
        t.Errorf("FromNanogramsPerDeciLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityNanogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanogramsPerDeciLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanogramsPerDeciLiter(math.NaN())
    if err == nil {
        t.Error("FromNanogramsPerDeciLiter() with NaN value should return error")
    }

    _, err = factory.FromNanogramsPerDeciLiter(math.Inf(1))
    if err == nil {
        t.Error("FromNanogramsPerDeciLiter() with +Inf value should return error")
    }

    _, err = factory.FromNanogramsPerDeciLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanogramsPerDeciLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanogramsPerDeciLiter(0)
    if err != nil {
        t.Errorf("FromNanogramsPerDeciLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityNanogramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanogramsPerDeciLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerDeciLiter function
func TestDensityFactory_FromMicrogramsPerDeciLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerDeciLiter(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerDeciLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityMicrogramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrogramsPerDeciLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrogramsPerDeciLiter(math.NaN())
    if err == nil {
        t.Error("FromMicrogramsPerDeciLiter() with NaN value should return error")
    }

    _, err = factory.FromMicrogramsPerDeciLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrogramsPerDeciLiter() with +Inf value should return error")
    }

    _, err = factory.FromMicrogramsPerDeciLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrogramsPerDeciLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrogramsPerDeciLiter(0)
    if err != nil {
        t.Errorf("FromMicrogramsPerDeciLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityMicrogramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerDeciLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerDeciLiter function
func TestDensityFactory_FromMilligramsPerDeciLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerDeciLiter(100)
    if err != nil {
        t.Errorf("FromMilligramsPerDeciLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityMilligramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerDeciLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerDeciLiter(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerDeciLiter() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerDeciLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerDeciLiter() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerDeciLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerDeciLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerDeciLiter(0)
    if err != nil {
        t.Errorf("FromMilligramsPerDeciLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityMilligramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerDeciLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigramsPerDeciLiter function
func TestDensityFactory_FromCentigramsPerDeciLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigramsPerDeciLiter(100)
    if err != nil {
        t.Errorf("FromCentigramsPerDeciLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityCentigramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentigramsPerDeciLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentigramsPerDeciLiter(math.NaN())
    if err == nil {
        t.Error("FromCentigramsPerDeciLiter() with NaN value should return error")
    }

    _, err = factory.FromCentigramsPerDeciLiter(math.Inf(1))
    if err == nil {
        t.Error("FromCentigramsPerDeciLiter() with +Inf value should return error")
    }

    _, err = factory.FromCentigramsPerDeciLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromCentigramsPerDeciLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentigramsPerDeciLiter(0)
    if err != nil {
        t.Errorf("FromCentigramsPerDeciLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityCentigramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigramsPerDeciLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigramsPerDeciLiter function
func TestDensityFactory_FromDecigramsPerDeciLiter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigramsPerDeciLiter(100)
    if err != nil {
        t.Errorf("FromDecigramsPerDeciLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityDecigramPerDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecigramsPerDeciLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecigramsPerDeciLiter(math.NaN())
    if err == nil {
        t.Error("FromDecigramsPerDeciLiter() with NaN value should return error")
    }

    _, err = factory.FromDecigramsPerDeciLiter(math.Inf(1))
    if err == nil {
        t.Error("FromDecigramsPerDeciLiter() with +Inf value should return error")
    }

    _, err = factory.FromDecigramsPerDeciLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecigramsPerDeciLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecigramsPerDeciLiter(0)
    if err != nil {
        t.Errorf("FromDecigramsPerDeciLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityDecigramPerDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigramsPerDeciLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromFemtogramsPerMilliliter function
func TestDensityFactory_FromFemtogramsPerMilliliter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFemtogramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromFemtogramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityFemtogramPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFemtogramsPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFemtogramsPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromFemtogramsPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromFemtogramsPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromFemtogramsPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromFemtogramsPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromFemtogramsPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFemtogramsPerMilliliter(0)
    if err != nil {
        t.Errorf("FromFemtogramsPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DensityFemtogramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFemtogramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicogramsPerMilliliter function
func TestDensityFactory_FromPicogramsPerMilliliter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicogramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromPicogramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityPicogramPerMilliliter)
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
    converted = zeroResult.Convert(units.DensityPicogramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicogramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanogramsPerMilliliter function
func TestDensityFactory_FromNanogramsPerMilliliter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanogramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromNanogramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityNanogramPerMilliliter)
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
    converted = zeroResult.Convert(units.DensityNanogramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanogramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrogramsPerMilliliter function
func TestDensityFactory_FromMicrogramsPerMilliliter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrogramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromMicrogramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityMicrogramPerMilliliter)
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
    converted = zeroResult.Convert(units.DensityMicrogramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrogramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerMilliliter function
func TestDensityFactory_FromMilligramsPerMilliliter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromMilligramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityMilligramPerMilliliter)
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
    converted = zeroResult.Convert(units.DensityMilligramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentigramsPerMilliliter function
func TestDensityFactory_FromCentigramsPerMilliliter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentigramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromCentigramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityCentigramPerMilliliter)
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
    converted = zeroResult.Convert(units.DensityCentigramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentigramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecigramsPerMilliliter function
func TestDensityFactory_FromDecigramsPerMilliliter(t *testing.T) {
    factory := units.DensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecigramsPerMilliliter(100)
    if err != nil {
        t.Errorf("FromDecigramsPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DensityDecigramPerMilliliter)
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
    converted = zeroResult.Convert(units.DensityDecigramPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecigramsPerMilliliter() with zero value = %v, want 0", converted)
    }
}

func TestDensityToString(t *testing.T) {
	factory := units.DensityFactory{}
	a, err := factory.CreateDensity(45, units.DensityKilogramPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.DensityKilogramPerCubicMeter, 2)
	expected := "45.00 " + units.GetDensityAbbreviation(units.DensityKilogramPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.DensityKilogramPerCubicMeter, -1)
	expected = "45 " + units.GetDensityAbbreviation(units.DensityKilogramPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestDensity_EqualityAndComparison(t *testing.T) {
	factory := units.DensityFactory{}
	a1, _ := factory.CreateDensity(60, units.DensityKilogramPerCubicMeter)
	a2, _ := factory.CreateDensity(60, units.DensityKilogramPerCubicMeter)
	a3, _ := factory.CreateDensity(120, units.DensityKilogramPerCubicMeter)

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

func TestDensity_Arithmetic(t *testing.T) {
	factory := units.DensityFactory{}
	a1, _ := factory.CreateDensity(30, units.DensityKilogramPerCubicMeter)
	a2, _ := factory.CreateDensity(45, units.DensityKilogramPerCubicMeter)

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