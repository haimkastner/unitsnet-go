// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestVolumeConcentrationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "DecimalFraction"}`
	
	factory := units.VolumeConcentrationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.VolumeConcentrationDecimalFraction {
		t.Errorf("expected unit %v, got %v", units.VolumeConcentrationDecimalFraction, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "DecimalFraction"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestVolumeConcentrationDto_ToJSON(t *testing.T) {
	dto := units.VolumeConcentrationDto{
		Value: 45,
		Unit:  units.VolumeConcentrationDecimalFraction,
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
	if result["unit"].(string) != string(units.VolumeConcentrationDecimalFraction) {
		t.Errorf("expected unit %s, got %v", units.VolumeConcentrationDecimalFraction, result["unit"])
	}
}

func TestNewVolumeConcentration_InvalidValue(t *testing.T) {
	factory := units.VolumeConcentrationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateVolumeConcentration(math.NaN(), units.VolumeConcentrationDecimalFraction)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateVolumeConcentration(math.Inf(1), units.VolumeConcentrationDecimalFraction)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestVolumeConcentrationConversions(t *testing.T) {
	factory := units.VolumeConcentrationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateVolumeConcentration(180, units.VolumeConcentrationDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to DecimalFractions.
		// No expected conversion value provided for DecimalFractions, verifying result is not NaN.
		result := a.DecimalFractions()
		cacheResult := a.DecimalFractions()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecimalFractions returned NaN")
		}
	}
	{
		// Test conversion to LitersPerLiter.
		// No expected conversion value provided for LitersPerLiter, verifying result is not NaN.
		result := a.LitersPerLiter()
		cacheResult := a.LitersPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to LitersPerMilliliter.
		// No expected conversion value provided for LitersPerMilliliter, verifying result is not NaN.
		result := a.LitersPerMilliliter()
		cacheResult := a.LitersPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LitersPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to Percent.
		// No expected conversion value provided for Percent, verifying result is not NaN.
		result := a.Percent()
		cacheResult := a.Percent()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Percent returned NaN")
		}
	}
	{
		// Test conversion to PartsPerThousand.
		// No expected conversion value provided for PartsPerThousand, verifying result is not NaN.
		result := a.PartsPerThousand()
		cacheResult := a.PartsPerThousand()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PartsPerThousand returned NaN")
		}
	}
	{
		// Test conversion to PartsPerMillion.
		// No expected conversion value provided for PartsPerMillion, verifying result is not NaN.
		result := a.PartsPerMillion()
		cacheResult := a.PartsPerMillion()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PartsPerMillion returned NaN")
		}
	}
	{
		// Test conversion to PartsPerBillion.
		// No expected conversion value provided for PartsPerBillion, verifying result is not NaN.
		result := a.PartsPerBillion()
		cacheResult := a.PartsPerBillion()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PartsPerBillion returned NaN")
		}
	}
	{
		// Test conversion to PartsPerTrillion.
		// No expected conversion value provided for PartsPerTrillion, verifying result is not NaN.
		result := a.PartsPerTrillion()
		cacheResult := a.PartsPerTrillion()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PartsPerTrillion returned NaN")
		}
	}
	{
		// Test conversion to PicolitersPerLiter.
		// No expected conversion value provided for PicolitersPerLiter, verifying result is not NaN.
		result := a.PicolitersPerLiter()
		cacheResult := a.PicolitersPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicolitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to NanolitersPerLiter.
		// No expected conversion value provided for NanolitersPerLiter, verifying result is not NaN.
		result := a.NanolitersPerLiter()
		cacheResult := a.NanolitersPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanolitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MicrolitersPerLiter.
		// No expected conversion value provided for MicrolitersPerLiter, verifying result is not NaN.
		result := a.MicrolitersPerLiter()
		cacheResult := a.MicrolitersPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrolitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MillilitersPerLiter.
		// No expected conversion value provided for MillilitersPerLiter, verifying result is not NaN.
		result := a.MillilitersPerLiter()
		cacheResult := a.MillilitersPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillilitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to CentilitersPerLiter.
		// No expected conversion value provided for CentilitersPerLiter, verifying result is not NaN.
		result := a.CentilitersPerLiter()
		cacheResult := a.CentilitersPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentilitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to DecilitersPerLiter.
		// No expected conversion value provided for DecilitersPerLiter, verifying result is not NaN.
		result := a.DecilitersPerLiter()
		cacheResult := a.DecilitersPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecilitersPerLiter returned NaN")
		}
	}
	{
		// Test conversion to PicolitersPerMilliliter.
		// No expected conversion value provided for PicolitersPerMilliliter, verifying result is not NaN.
		result := a.PicolitersPerMilliliter()
		cacheResult := a.PicolitersPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicolitersPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to NanolitersPerMilliliter.
		// No expected conversion value provided for NanolitersPerMilliliter, verifying result is not NaN.
		result := a.NanolitersPerMilliliter()
		cacheResult := a.NanolitersPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanolitersPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to MicrolitersPerMilliliter.
		// No expected conversion value provided for MicrolitersPerMilliliter, verifying result is not NaN.
		result := a.MicrolitersPerMilliliter()
		cacheResult := a.MicrolitersPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrolitersPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to MillilitersPerMilliliter.
		// No expected conversion value provided for MillilitersPerMilliliter, verifying result is not NaN.
		result := a.MillilitersPerMilliliter()
		cacheResult := a.MillilitersPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillilitersPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to CentilitersPerMilliliter.
		// No expected conversion value provided for CentilitersPerMilliliter, verifying result is not NaN.
		result := a.CentilitersPerMilliliter()
		cacheResult := a.CentilitersPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentilitersPerMilliliter returned NaN")
		}
	}
	{
		// Test conversion to DecilitersPerMilliliter.
		// No expected conversion value provided for DecilitersPerMilliliter, verifying result is not NaN.
		result := a.DecilitersPerMilliliter()
		cacheResult := a.DecilitersPerMilliliter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecilitersPerMilliliter returned NaN")
		}
	}
}

func TestVolumeConcentration_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.VolumeConcentrationFactory{}
	a, err := factory.CreateVolumeConcentration(90, units.VolumeConcentrationDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.VolumeConcentrationDecimalFraction {
		t.Errorf("expected default unit DecimalFraction, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.VolumeConcentrationDecimalFraction
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.VolumeConcentrationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.VolumeConcentrationDecimalFraction {
		t.Errorf("expected unit DecimalFraction, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestVolumeConcentrationFactory_FromDto(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationDecimalFraction,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.VolumeConcentrationDto{
        Value: math.NaN(),
        Unit:  units.VolumeConcentrationDecimalFraction,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test DecimalFraction conversion
    decimal_fractionsDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationDecimalFraction,
    }
    
    var decimal_fractionsResult *units.VolumeConcentration
    decimal_fractionsResult, err = factory.FromDto(decimal_fractionsDto)
    if err != nil {
        t.Errorf("FromDto() with DecimalFraction returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimal_fractionsResult.Convert(units.VolumeConcentrationDecimalFraction)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimalFraction = %v, want %v", converted, 100)
    }
    // Test LiterPerLiter conversion
    liters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationLiterPerLiter,
    }
    
    var liters_per_literResult *units.VolumeConcentration
    liters_per_literResult, err = factory.FromDto(liters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with LiterPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_literResult.Convert(units.VolumeConcentrationLiterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerLiter = %v, want %v", converted, 100)
    }
    // Test LiterPerMilliliter conversion
    liters_per_milliliterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationLiterPerMilliliter,
    }
    
    var liters_per_milliliterResult *units.VolumeConcentration
    liters_per_milliliterResult, err = factory.FromDto(liters_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with LiterPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_milliliterResult.Convert(units.VolumeConcentrationLiterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerMilliliter = %v, want %v", converted, 100)
    }
    // Test Percent conversion
    percentDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationPercent,
    }
    
    var percentResult *units.VolumeConcentration
    percentResult, err = factory.FromDto(percentDto)
    if err != nil {
        t.Errorf("FromDto() with Percent returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = percentResult.Convert(units.VolumeConcentrationPercent)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Percent = %v, want %v", converted, 100)
    }
    // Test PartPerThousand conversion
    parts_per_thousandDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationPartPerThousand,
    }
    
    var parts_per_thousandResult *units.VolumeConcentration
    parts_per_thousandResult, err = factory.FromDto(parts_per_thousandDto)
    if err != nil {
        t.Errorf("FromDto() with PartPerThousand returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_thousandResult.Convert(units.VolumeConcentrationPartPerThousand)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerThousand = %v, want %v", converted, 100)
    }
    // Test PartPerMillion conversion
    parts_per_millionDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationPartPerMillion,
    }
    
    var parts_per_millionResult *units.VolumeConcentration
    parts_per_millionResult, err = factory.FromDto(parts_per_millionDto)
    if err != nil {
        t.Errorf("FromDto() with PartPerMillion returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_millionResult.Convert(units.VolumeConcentrationPartPerMillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerMillion = %v, want %v", converted, 100)
    }
    // Test PartPerBillion conversion
    parts_per_billionDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationPartPerBillion,
    }
    
    var parts_per_billionResult *units.VolumeConcentration
    parts_per_billionResult, err = factory.FromDto(parts_per_billionDto)
    if err != nil {
        t.Errorf("FromDto() with PartPerBillion returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_billionResult.Convert(units.VolumeConcentrationPartPerBillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerBillion = %v, want %v", converted, 100)
    }
    // Test PartPerTrillion conversion
    parts_per_trillionDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationPartPerTrillion,
    }
    
    var parts_per_trillionResult *units.VolumeConcentration
    parts_per_trillionResult, err = factory.FromDto(parts_per_trillionDto)
    if err != nil {
        t.Errorf("FromDto() with PartPerTrillion returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_trillionResult.Convert(units.VolumeConcentrationPartPerTrillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerTrillion = %v, want %v", converted, 100)
    }
    // Test PicoliterPerLiter conversion
    picoliters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationPicoliterPerLiter,
    }
    
    var picoliters_per_literResult *units.VolumeConcentration
    picoliters_per_literResult, err = factory.FromDto(picoliters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with PicoliterPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoliters_per_literResult.Convert(units.VolumeConcentrationPicoliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicoliterPerLiter = %v, want %v", converted, 100)
    }
    // Test NanoliterPerLiter conversion
    nanoliters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationNanoliterPerLiter,
    }
    
    var nanoliters_per_literResult *units.VolumeConcentration
    nanoliters_per_literResult, err = factory.FromDto(nanoliters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with NanoliterPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_literResult.Convert(units.VolumeConcentrationNanoliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoliterPerLiter = %v, want %v", converted, 100)
    }
    // Test MicroliterPerLiter conversion
    microliters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationMicroliterPerLiter,
    }
    
    var microliters_per_literResult *units.VolumeConcentration
    microliters_per_literResult, err = factory.FromDto(microliters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MicroliterPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_literResult.Convert(units.VolumeConcentrationMicroliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroliterPerLiter = %v, want %v", converted, 100)
    }
    // Test MilliliterPerLiter conversion
    milliliters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationMilliliterPerLiter,
    }
    
    var milliliters_per_literResult *units.VolumeConcentration
    milliliters_per_literResult, err = factory.FromDto(milliliters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MilliliterPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_literResult.Convert(units.VolumeConcentrationMilliliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliliterPerLiter = %v, want %v", converted, 100)
    }
    // Test CentiliterPerLiter conversion
    centiliters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationCentiliterPerLiter,
    }
    
    var centiliters_per_literResult *units.VolumeConcentration
    centiliters_per_literResult, err = factory.FromDto(centiliters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with CentiliterPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_literResult.Convert(units.VolumeConcentrationCentiliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiliterPerLiter = %v, want %v", converted, 100)
    }
    // Test DeciliterPerLiter conversion
    deciliters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationDeciliterPerLiter,
    }
    
    var deciliters_per_literResult *units.VolumeConcentration
    deciliters_per_literResult, err = factory.FromDto(deciliters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with DeciliterPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_literResult.Convert(units.VolumeConcentrationDeciliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciliterPerLiter = %v, want %v", converted, 100)
    }
    // Test PicoliterPerMilliliter conversion
    picoliters_per_milliliterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationPicoliterPerMilliliter,
    }
    
    var picoliters_per_milliliterResult *units.VolumeConcentration
    picoliters_per_milliliterResult, err = factory.FromDto(picoliters_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with PicoliterPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoliters_per_milliliterResult.Convert(units.VolumeConcentrationPicoliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicoliterPerMilliliter = %v, want %v", converted, 100)
    }
    // Test NanoliterPerMilliliter conversion
    nanoliters_per_milliliterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationNanoliterPerMilliliter,
    }
    
    var nanoliters_per_milliliterResult *units.VolumeConcentration
    nanoliters_per_milliliterResult, err = factory.FromDto(nanoliters_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with NanoliterPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_milliliterResult.Convert(units.VolumeConcentrationNanoliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoliterPerMilliliter = %v, want %v", converted, 100)
    }
    // Test MicroliterPerMilliliter conversion
    microliters_per_milliliterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationMicroliterPerMilliliter,
    }
    
    var microliters_per_milliliterResult *units.VolumeConcentration
    microliters_per_milliliterResult, err = factory.FromDto(microliters_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with MicroliterPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_milliliterResult.Convert(units.VolumeConcentrationMicroliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroliterPerMilliliter = %v, want %v", converted, 100)
    }
    // Test MilliliterPerMilliliter conversion
    milliliters_per_milliliterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationMilliliterPerMilliliter,
    }
    
    var milliliters_per_milliliterResult *units.VolumeConcentration
    milliliters_per_milliliterResult, err = factory.FromDto(milliliters_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with MilliliterPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_milliliterResult.Convert(units.VolumeConcentrationMilliliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliliterPerMilliliter = %v, want %v", converted, 100)
    }
    // Test CentiliterPerMilliliter conversion
    centiliters_per_milliliterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationCentiliterPerMilliliter,
    }
    
    var centiliters_per_milliliterResult *units.VolumeConcentration
    centiliters_per_milliliterResult, err = factory.FromDto(centiliters_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with CentiliterPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_milliliterResult.Convert(units.VolumeConcentrationCentiliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiliterPerMilliliter = %v, want %v", converted, 100)
    }
    // Test DeciliterPerMilliliter conversion
    deciliters_per_milliliterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationDeciliterPerMilliliter,
    }
    
    var deciliters_per_milliliterResult *units.VolumeConcentration
    deciliters_per_milliliterResult, err = factory.FromDto(deciliters_per_milliliterDto)
    if err != nil {
        t.Errorf("FromDto() with DeciliterPerMilliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_milliliterResult.Convert(units.VolumeConcentrationDeciliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciliterPerMilliliter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.VolumeConcentrationDto{
        Value: 0,
        Unit:  units.VolumeConcentrationDecimalFraction,
    }
    
    var zeroResult *units.VolumeConcentration
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestVolumeConcentrationFactory_FromDtoJSON(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "DecimalFraction"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "DecimalFraction"}`)
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
    nanJSON, _ := json.Marshal(units.VolumeConcentrationDto{
        Value: nanValue,
        Unit:  units.VolumeConcentrationDecimalFraction,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with DecimalFraction unit
    decimal_fractionsJSON := []byte(`{"value": 100, "unit": "DecimalFraction"}`)
    decimal_fractionsResult, err := factory.FromDtoJSON(decimal_fractionsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecimalFraction unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimal_fractionsResult.Convert(units.VolumeConcentrationDecimalFraction)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimalFraction = %v, want %v", converted, 100)
    }
    // Test JSON with LiterPerLiter unit
    liters_per_literJSON := []byte(`{"value": 100, "unit": "LiterPerLiter"}`)
    liters_per_literResult, err := factory.FromDtoJSON(liters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LiterPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_literResult.Convert(units.VolumeConcentrationLiterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with LiterPerMilliliter unit
    liters_per_milliliterJSON := []byte(`{"value": 100, "unit": "LiterPerMilliliter"}`)
    liters_per_milliliterResult, err := factory.FromDtoJSON(liters_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LiterPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_milliliterResult.Convert(units.VolumeConcentrationLiterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with Percent unit
    percentJSON := []byte(`{"value": 100, "unit": "Percent"}`)
    percentResult, err := factory.FromDtoJSON(percentJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Percent unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = percentResult.Convert(units.VolumeConcentrationPercent)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Percent = %v, want %v", converted, 100)
    }
    // Test JSON with PartPerThousand unit
    parts_per_thousandJSON := []byte(`{"value": 100, "unit": "PartPerThousand"}`)
    parts_per_thousandResult, err := factory.FromDtoJSON(parts_per_thousandJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PartPerThousand unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_thousandResult.Convert(units.VolumeConcentrationPartPerThousand)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerThousand = %v, want %v", converted, 100)
    }
    // Test JSON with PartPerMillion unit
    parts_per_millionJSON := []byte(`{"value": 100, "unit": "PartPerMillion"}`)
    parts_per_millionResult, err := factory.FromDtoJSON(parts_per_millionJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PartPerMillion unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_millionResult.Convert(units.VolumeConcentrationPartPerMillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerMillion = %v, want %v", converted, 100)
    }
    // Test JSON with PartPerBillion unit
    parts_per_billionJSON := []byte(`{"value": 100, "unit": "PartPerBillion"}`)
    parts_per_billionResult, err := factory.FromDtoJSON(parts_per_billionJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PartPerBillion unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_billionResult.Convert(units.VolumeConcentrationPartPerBillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerBillion = %v, want %v", converted, 100)
    }
    // Test JSON with PartPerTrillion unit
    parts_per_trillionJSON := []byte(`{"value": 100, "unit": "PartPerTrillion"}`)
    parts_per_trillionResult, err := factory.FromDtoJSON(parts_per_trillionJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PartPerTrillion unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_trillionResult.Convert(units.VolumeConcentrationPartPerTrillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerTrillion = %v, want %v", converted, 100)
    }
    // Test JSON with PicoliterPerLiter unit
    picoliters_per_literJSON := []byte(`{"value": 100, "unit": "PicoliterPerLiter"}`)
    picoliters_per_literResult, err := factory.FromDtoJSON(picoliters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicoliterPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoliters_per_literResult.Convert(units.VolumeConcentrationPicoliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicoliterPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with NanoliterPerLiter unit
    nanoliters_per_literJSON := []byte(`{"value": 100, "unit": "NanoliterPerLiter"}`)
    nanoliters_per_literResult, err := factory.FromDtoJSON(nanoliters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanoliterPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_literResult.Convert(units.VolumeConcentrationNanoliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoliterPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with MicroliterPerLiter unit
    microliters_per_literJSON := []byte(`{"value": 100, "unit": "MicroliterPerLiter"}`)
    microliters_per_literResult, err := factory.FromDtoJSON(microliters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicroliterPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_literResult.Convert(units.VolumeConcentrationMicroliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroliterPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with MilliliterPerLiter unit
    milliliters_per_literJSON := []byte(`{"value": 100, "unit": "MilliliterPerLiter"}`)
    milliliters_per_literResult, err := factory.FromDtoJSON(milliliters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliliterPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_literResult.Convert(units.VolumeConcentrationMilliliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliliterPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with CentiliterPerLiter unit
    centiliters_per_literJSON := []byte(`{"value": 100, "unit": "CentiliterPerLiter"}`)
    centiliters_per_literResult, err := factory.FromDtoJSON(centiliters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentiliterPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_literResult.Convert(units.VolumeConcentrationCentiliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiliterPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with DeciliterPerLiter unit
    deciliters_per_literJSON := []byte(`{"value": 100, "unit": "DeciliterPerLiter"}`)
    deciliters_per_literResult, err := factory.FromDtoJSON(deciliters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DeciliterPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_literResult.Convert(units.VolumeConcentrationDeciliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciliterPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with PicoliterPerMilliliter unit
    picoliters_per_milliliterJSON := []byte(`{"value": 100, "unit": "PicoliterPerMilliliter"}`)
    picoliters_per_milliliterResult, err := factory.FromDtoJSON(picoliters_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicoliterPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoliters_per_milliliterResult.Convert(units.VolumeConcentrationPicoliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicoliterPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with NanoliterPerMilliliter unit
    nanoliters_per_milliliterJSON := []byte(`{"value": 100, "unit": "NanoliterPerMilliliter"}`)
    nanoliters_per_milliliterResult, err := factory.FromDtoJSON(nanoliters_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanoliterPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_milliliterResult.Convert(units.VolumeConcentrationNanoliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoliterPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with MicroliterPerMilliliter unit
    microliters_per_milliliterJSON := []byte(`{"value": 100, "unit": "MicroliterPerMilliliter"}`)
    microliters_per_milliliterResult, err := factory.FromDtoJSON(microliters_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicroliterPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_milliliterResult.Convert(units.VolumeConcentrationMicroliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroliterPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with MilliliterPerMilliliter unit
    milliliters_per_milliliterJSON := []byte(`{"value": 100, "unit": "MilliliterPerMilliliter"}`)
    milliliters_per_milliliterResult, err := factory.FromDtoJSON(milliliters_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliliterPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_milliliterResult.Convert(units.VolumeConcentrationMilliliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliliterPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with CentiliterPerMilliliter unit
    centiliters_per_milliliterJSON := []byte(`{"value": 100, "unit": "CentiliterPerMilliliter"}`)
    centiliters_per_milliliterResult, err := factory.FromDtoJSON(centiliters_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentiliterPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_milliliterResult.Convert(units.VolumeConcentrationCentiliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiliterPerMilliliter = %v, want %v", converted, 100)
    }
    // Test JSON with DeciliterPerMilliliter unit
    deciliters_per_milliliterJSON := []byte(`{"value": 100, "unit": "DeciliterPerMilliliter"}`)
    deciliters_per_milliliterResult, err := factory.FromDtoJSON(deciliters_per_milliliterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DeciliterPerMilliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_milliliterResult.Convert(units.VolumeConcentrationDeciliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciliterPerMilliliter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "DecimalFraction"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromDecimalFractions function
func TestVolumeConcentrationFactory_FromDecimalFractions(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecimalFractions(100)
    if err != nil {
        t.Errorf("FromDecimalFractions() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationDecimalFraction)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecimalFractions() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecimalFractions(math.NaN())
    if err == nil {
        t.Error("FromDecimalFractions() with NaN value should return error")
    }

    _, err = factory.FromDecimalFractions(math.Inf(1))
    if err == nil {
        t.Error("FromDecimalFractions() with +Inf value should return error")
    }

    _, err = factory.FromDecimalFractions(math.Inf(-1))
    if err == nil {
        t.Error("FromDecimalFractions() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecimalFractions(0)
    if err != nil {
        t.Errorf("FromDecimalFractions() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationDecimalFraction)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecimalFractions() with zero value = %v, want 0", converted)
    }
}
// Test FromLitersPerLiter function
func TestVolumeConcentrationFactory_FromLitersPerLiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLitersPerLiter(100)
    if err != nil {
        t.Errorf("FromLitersPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationLiterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLitersPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLitersPerLiter(math.NaN())
    if err == nil {
        t.Error("FromLitersPerLiter() with NaN value should return error")
    }

    _, err = factory.FromLitersPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromLitersPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromLitersPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromLitersPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLitersPerLiter(0)
    if err != nil {
        t.Errorf("FromLitersPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationLiterPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLitersPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromLitersPerMilliliter function
func TestVolumeConcentrationFactory_FromLitersPerMilliliter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLitersPerMilliliter(100)
    if err != nil {
        t.Errorf("FromLitersPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationLiterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLitersPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLitersPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromLitersPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromLitersPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromLitersPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromLitersPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromLitersPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLitersPerMilliliter(0)
    if err != nil {
        t.Errorf("FromLitersPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationLiterPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLitersPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromPercent function
func TestVolumeConcentrationFactory_FromPercent(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPercent(100)
    if err != nil {
        t.Errorf("FromPercent() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationPercent)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPercent() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPercent(math.NaN())
    if err == nil {
        t.Error("FromPercent() with NaN value should return error")
    }

    _, err = factory.FromPercent(math.Inf(1))
    if err == nil {
        t.Error("FromPercent() with +Inf value should return error")
    }

    _, err = factory.FromPercent(math.Inf(-1))
    if err == nil {
        t.Error("FromPercent() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPercent(0)
    if err != nil {
        t.Errorf("FromPercent() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationPercent)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPercent() with zero value = %v, want 0", converted)
    }
}
// Test FromPartsPerThousand function
func TestVolumeConcentrationFactory_FromPartsPerThousand(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPartsPerThousand(100)
    if err != nil {
        t.Errorf("FromPartsPerThousand() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationPartPerThousand)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPartsPerThousand() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPartsPerThousand(math.NaN())
    if err == nil {
        t.Error("FromPartsPerThousand() with NaN value should return error")
    }

    _, err = factory.FromPartsPerThousand(math.Inf(1))
    if err == nil {
        t.Error("FromPartsPerThousand() with +Inf value should return error")
    }

    _, err = factory.FromPartsPerThousand(math.Inf(-1))
    if err == nil {
        t.Error("FromPartsPerThousand() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPartsPerThousand(0)
    if err != nil {
        t.Errorf("FromPartsPerThousand() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationPartPerThousand)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPartsPerThousand() with zero value = %v, want 0", converted)
    }
}
// Test FromPartsPerMillion function
func TestVolumeConcentrationFactory_FromPartsPerMillion(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPartsPerMillion(100)
    if err != nil {
        t.Errorf("FromPartsPerMillion() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationPartPerMillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPartsPerMillion() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPartsPerMillion(math.NaN())
    if err == nil {
        t.Error("FromPartsPerMillion() with NaN value should return error")
    }

    _, err = factory.FromPartsPerMillion(math.Inf(1))
    if err == nil {
        t.Error("FromPartsPerMillion() with +Inf value should return error")
    }

    _, err = factory.FromPartsPerMillion(math.Inf(-1))
    if err == nil {
        t.Error("FromPartsPerMillion() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPartsPerMillion(0)
    if err != nil {
        t.Errorf("FromPartsPerMillion() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationPartPerMillion)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPartsPerMillion() with zero value = %v, want 0", converted)
    }
}
// Test FromPartsPerBillion function
func TestVolumeConcentrationFactory_FromPartsPerBillion(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPartsPerBillion(100)
    if err != nil {
        t.Errorf("FromPartsPerBillion() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationPartPerBillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPartsPerBillion() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPartsPerBillion(math.NaN())
    if err == nil {
        t.Error("FromPartsPerBillion() with NaN value should return error")
    }

    _, err = factory.FromPartsPerBillion(math.Inf(1))
    if err == nil {
        t.Error("FromPartsPerBillion() with +Inf value should return error")
    }

    _, err = factory.FromPartsPerBillion(math.Inf(-1))
    if err == nil {
        t.Error("FromPartsPerBillion() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPartsPerBillion(0)
    if err != nil {
        t.Errorf("FromPartsPerBillion() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationPartPerBillion)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPartsPerBillion() with zero value = %v, want 0", converted)
    }
}
// Test FromPartsPerTrillion function
func TestVolumeConcentrationFactory_FromPartsPerTrillion(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPartsPerTrillion(100)
    if err != nil {
        t.Errorf("FromPartsPerTrillion() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationPartPerTrillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPartsPerTrillion() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPartsPerTrillion(math.NaN())
    if err == nil {
        t.Error("FromPartsPerTrillion() with NaN value should return error")
    }

    _, err = factory.FromPartsPerTrillion(math.Inf(1))
    if err == nil {
        t.Error("FromPartsPerTrillion() with +Inf value should return error")
    }

    _, err = factory.FromPartsPerTrillion(math.Inf(-1))
    if err == nil {
        t.Error("FromPartsPerTrillion() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPartsPerTrillion(0)
    if err != nil {
        t.Errorf("FromPartsPerTrillion() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationPartPerTrillion)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPartsPerTrillion() with zero value = %v, want 0", converted)
    }
}
// Test FromPicolitersPerLiter function
func TestVolumeConcentrationFactory_FromPicolitersPerLiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicolitersPerLiter(100)
    if err != nil {
        t.Errorf("FromPicolitersPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationPicoliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicolitersPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicolitersPerLiter(math.NaN())
    if err == nil {
        t.Error("FromPicolitersPerLiter() with NaN value should return error")
    }

    _, err = factory.FromPicolitersPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromPicolitersPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromPicolitersPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicolitersPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicolitersPerLiter(0)
    if err != nil {
        t.Errorf("FromPicolitersPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationPicoliterPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicolitersPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanolitersPerLiter function
func TestVolumeConcentrationFactory_FromNanolitersPerLiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanolitersPerLiter(100)
    if err != nil {
        t.Errorf("FromNanolitersPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationNanoliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanolitersPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanolitersPerLiter(math.NaN())
    if err == nil {
        t.Error("FromNanolitersPerLiter() with NaN value should return error")
    }

    _, err = factory.FromNanolitersPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromNanolitersPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromNanolitersPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanolitersPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanolitersPerLiter(0)
    if err != nil {
        t.Errorf("FromNanolitersPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationNanoliterPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanolitersPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrolitersPerLiter function
func TestVolumeConcentrationFactory_FromMicrolitersPerLiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrolitersPerLiter(100)
    if err != nil {
        t.Errorf("FromMicrolitersPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationMicroliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrolitersPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrolitersPerLiter(math.NaN())
    if err == nil {
        t.Error("FromMicrolitersPerLiter() with NaN value should return error")
    }

    _, err = factory.FromMicrolitersPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrolitersPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromMicrolitersPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrolitersPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrolitersPerLiter(0)
    if err != nil {
        t.Errorf("FromMicrolitersPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationMicroliterPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrolitersPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillilitersPerLiter function
func TestVolumeConcentrationFactory_FromMillilitersPerLiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillilitersPerLiter(100)
    if err != nil {
        t.Errorf("FromMillilitersPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationMilliliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillilitersPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillilitersPerLiter(math.NaN())
    if err == nil {
        t.Error("FromMillilitersPerLiter() with NaN value should return error")
    }

    _, err = factory.FromMillilitersPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMillilitersPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromMillilitersPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillilitersPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillilitersPerLiter(0)
    if err != nil {
        t.Errorf("FromMillilitersPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationMilliliterPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillilitersPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentilitersPerLiter function
func TestVolumeConcentrationFactory_FromCentilitersPerLiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentilitersPerLiter(100)
    if err != nil {
        t.Errorf("FromCentilitersPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationCentiliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentilitersPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentilitersPerLiter(math.NaN())
    if err == nil {
        t.Error("FromCentilitersPerLiter() with NaN value should return error")
    }

    _, err = factory.FromCentilitersPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromCentilitersPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromCentilitersPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromCentilitersPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentilitersPerLiter(0)
    if err != nil {
        t.Errorf("FromCentilitersPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationCentiliterPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentilitersPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecilitersPerLiter function
func TestVolumeConcentrationFactory_FromDecilitersPerLiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecilitersPerLiter(100)
    if err != nil {
        t.Errorf("FromDecilitersPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationDeciliterPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecilitersPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecilitersPerLiter(math.NaN())
    if err == nil {
        t.Error("FromDecilitersPerLiter() with NaN value should return error")
    }

    _, err = factory.FromDecilitersPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromDecilitersPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromDecilitersPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecilitersPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecilitersPerLiter(0)
    if err != nil {
        t.Errorf("FromDecilitersPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationDeciliterPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecilitersPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicolitersPerMilliliter function
func TestVolumeConcentrationFactory_FromPicolitersPerMilliliter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicolitersPerMilliliter(100)
    if err != nil {
        t.Errorf("FromPicolitersPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationPicoliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicolitersPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicolitersPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromPicolitersPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromPicolitersPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromPicolitersPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromPicolitersPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicolitersPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicolitersPerMilliliter(0)
    if err != nil {
        t.Errorf("FromPicolitersPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationPicoliterPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicolitersPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanolitersPerMilliliter function
func TestVolumeConcentrationFactory_FromNanolitersPerMilliliter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanolitersPerMilliliter(100)
    if err != nil {
        t.Errorf("FromNanolitersPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationNanoliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanolitersPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanolitersPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromNanolitersPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromNanolitersPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromNanolitersPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromNanolitersPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanolitersPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanolitersPerMilliliter(0)
    if err != nil {
        t.Errorf("FromNanolitersPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationNanoliterPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanolitersPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrolitersPerMilliliter function
func TestVolumeConcentrationFactory_FromMicrolitersPerMilliliter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrolitersPerMilliliter(100)
    if err != nil {
        t.Errorf("FromMicrolitersPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationMicroliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrolitersPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrolitersPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromMicrolitersPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromMicrolitersPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrolitersPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromMicrolitersPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrolitersPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrolitersPerMilliliter(0)
    if err != nil {
        t.Errorf("FromMicrolitersPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationMicroliterPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrolitersPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillilitersPerMilliliter function
func TestVolumeConcentrationFactory_FromMillilitersPerMilliliter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillilitersPerMilliliter(100)
    if err != nil {
        t.Errorf("FromMillilitersPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationMilliliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillilitersPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillilitersPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromMillilitersPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromMillilitersPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromMillilitersPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromMillilitersPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillilitersPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillilitersPerMilliliter(0)
    if err != nil {
        t.Errorf("FromMillilitersPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationMilliliterPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillilitersPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentilitersPerMilliliter function
func TestVolumeConcentrationFactory_FromCentilitersPerMilliliter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentilitersPerMilliliter(100)
    if err != nil {
        t.Errorf("FromCentilitersPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationCentiliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentilitersPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentilitersPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromCentilitersPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromCentilitersPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromCentilitersPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromCentilitersPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromCentilitersPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentilitersPerMilliliter(0)
    if err != nil {
        t.Errorf("FromCentilitersPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationCentiliterPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentilitersPerMilliliter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecilitersPerMilliliter function
func TestVolumeConcentrationFactory_FromDecilitersPerMilliliter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecilitersPerMilliliter(100)
    if err != nil {
        t.Errorf("FromDecilitersPerMilliliter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationDeciliterPerMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecilitersPerMilliliter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecilitersPerMilliliter(math.NaN())
    if err == nil {
        t.Error("FromDecilitersPerMilliliter() with NaN value should return error")
    }

    _, err = factory.FromDecilitersPerMilliliter(math.Inf(1))
    if err == nil {
        t.Error("FromDecilitersPerMilliliter() with +Inf value should return error")
    }

    _, err = factory.FromDecilitersPerMilliliter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecilitersPerMilliliter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecilitersPerMilliliter(0)
    if err != nil {
        t.Errorf("FromDecilitersPerMilliliter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationDeciliterPerMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecilitersPerMilliliter() with zero value = %v, want 0", converted)
    }
}

func TestVolumeConcentrationToString(t *testing.T) {
	factory := units.VolumeConcentrationFactory{}
	a, err := factory.CreateVolumeConcentration(45, units.VolumeConcentrationDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.VolumeConcentrationDecimalFraction, 2)
	expected := "45.00 " + units.GetVolumeConcentrationAbbreviation(units.VolumeConcentrationDecimalFraction)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.VolumeConcentrationDecimalFraction, -1)
	expected = "45 " + units.GetVolumeConcentrationAbbreviation(units.VolumeConcentrationDecimalFraction)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestVolumeConcentration_EqualityAndComparison(t *testing.T) {
	factory := units.VolumeConcentrationFactory{}
	a1, _ := factory.CreateVolumeConcentration(60, units.VolumeConcentrationDecimalFraction)
	a2, _ := factory.CreateVolumeConcentration(60, units.VolumeConcentrationDecimalFraction)
	a3, _ := factory.CreateVolumeConcentration(120, units.VolumeConcentrationDecimalFraction)

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

func TestVolumeConcentration_Arithmetic(t *testing.T) {
	factory := units.VolumeConcentrationFactory{}
	a1, _ := factory.CreateVolumeConcentration(30, units.VolumeConcentrationDecimalFraction)
	a2, _ := factory.CreateVolumeConcentration(45, units.VolumeConcentrationDecimalFraction)

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


func TestGetVolumeConcentrationAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.VolumeConcentrationUnits
        want string
    }{
        {
            name: "DecimalFraction abbreviation",
            unit: units.VolumeConcentrationDecimalFraction,
            want: "",
        },
        {
            name: "LiterPerLiter abbreviation",
            unit: units.VolumeConcentrationLiterPerLiter,
            want: "l/l",
        },
        {
            name: "LiterPerMilliliter abbreviation",
            unit: units.VolumeConcentrationLiterPerMilliliter,
            want: "l/ml",
        },
        {
            name: "Percent abbreviation",
            unit: units.VolumeConcentrationPercent,
            want: "%",
        },
        {
            name: "PartPerThousand abbreviation",
            unit: units.VolumeConcentrationPartPerThousand,
            want: "‰",
        },
        {
            name: "PartPerMillion abbreviation",
            unit: units.VolumeConcentrationPartPerMillion,
            want: "ppm",
        },
        {
            name: "PartPerBillion abbreviation",
            unit: units.VolumeConcentrationPartPerBillion,
            want: "ppb",
        },
        {
            name: "PartPerTrillion abbreviation",
            unit: units.VolumeConcentrationPartPerTrillion,
            want: "ppt",
        },
        {
            name: "PicoliterPerLiter abbreviation",
            unit: units.VolumeConcentrationPicoliterPerLiter,
            want: "pl/l",
        },
        {
            name: "NanoliterPerLiter abbreviation",
            unit: units.VolumeConcentrationNanoliterPerLiter,
            want: "nl/l",
        },
        {
            name: "MicroliterPerLiter abbreviation",
            unit: units.VolumeConcentrationMicroliterPerLiter,
            want: "μl/l",
        },
        {
            name: "MilliliterPerLiter abbreviation",
            unit: units.VolumeConcentrationMilliliterPerLiter,
            want: "ml/l",
        },
        {
            name: "CentiliterPerLiter abbreviation",
            unit: units.VolumeConcentrationCentiliterPerLiter,
            want: "cl/l",
        },
        {
            name: "DeciliterPerLiter abbreviation",
            unit: units.VolumeConcentrationDeciliterPerLiter,
            want: "dl/l",
        },
        {
            name: "PicoliterPerMilliliter abbreviation",
            unit: units.VolumeConcentrationPicoliterPerMilliliter,
            want: "pl/ml",
        },
        {
            name: "NanoliterPerMilliliter abbreviation",
            unit: units.VolumeConcentrationNanoliterPerMilliliter,
            want: "nl/ml",
        },
        {
            name: "MicroliterPerMilliliter abbreviation",
            unit: units.VolumeConcentrationMicroliterPerMilliliter,
            want: "μl/ml",
        },
        {
            name: "MilliliterPerMilliliter abbreviation",
            unit: units.VolumeConcentrationMilliliterPerMilliliter,
            want: "ml/ml",
        },
        {
            name: "CentiliterPerMilliliter abbreviation",
            unit: units.VolumeConcentrationCentiliterPerMilliliter,
            want: "cl/ml",
        },
        {
            name: "DeciliterPerMilliliter abbreviation",
            unit: units.VolumeConcentrationDeciliterPerMilliliter,
            want: "dl/ml",
        },
        {
            name: "invalid unit",
            unit: units.VolumeConcentrationUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetVolumeConcentrationAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetVolumeConcentrationAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestVolumeConcentration_String(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    
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
            unit, err := factory.CreateVolumeConcentration(tt.value, units.VolumeConcentrationDecimalFraction)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("VolumeConcentration.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestVolumeConcentration_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.VolumeConcentrationFactory{}

	_, err := uf.CreateVolumeConcentration(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}