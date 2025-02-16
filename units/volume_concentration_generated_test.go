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
		// Test conversion to LitersPerMililiter.
		// No expected conversion value provided for LitersPerMililiter, verifying result is not NaN.
		result := a.LitersPerMililiter()
		cacheResult := a.LitersPerMililiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LitersPerMililiter returned NaN")
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
		// Test conversion to PicolitersPerMililiter.
		// No expected conversion value provided for PicolitersPerMililiter, verifying result is not NaN.
		result := a.PicolitersPerMililiter()
		cacheResult := a.PicolitersPerMililiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicolitersPerMililiter returned NaN")
		}
	}
	{
		// Test conversion to NanolitersPerMililiter.
		// No expected conversion value provided for NanolitersPerMililiter, verifying result is not NaN.
		result := a.NanolitersPerMililiter()
		cacheResult := a.NanolitersPerMililiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanolitersPerMililiter returned NaN")
		}
	}
	{
		// Test conversion to MicrolitersPerMililiter.
		// No expected conversion value provided for MicrolitersPerMililiter, verifying result is not NaN.
		result := a.MicrolitersPerMililiter()
		cacheResult := a.MicrolitersPerMililiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrolitersPerMililiter returned NaN")
		}
	}
	{
		// Test conversion to MillilitersPerMililiter.
		// No expected conversion value provided for MillilitersPerMililiter, verifying result is not NaN.
		result := a.MillilitersPerMililiter()
		cacheResult := a.MillilitersPerMililiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillilitersPerMililiter returned NaN")
		}
	}
	{
		// Test conversion to CentilitersPerMililiter.
		// No expected conversion value provided for CentilitersPerMililiter, verifying result is not NaN.
		result := a.CentilitersPerMililiter()
		cacheResult := a.CentilitersPerMililiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentilitersPerMililiter returned NaN")
		}
	}
	{
		// Test conversion to DecilitersPerMililiter.
		// No expected conversion value provided for DecilitersPerMililiter, verifying result is not NaN.
		result := a.DecilitersPerMililiter()
		cacheResult := a.DecilitersPerMililiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecilitersPerMililiter returned NaN")
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
    // Test LitersPerLiter conversion
    liters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationLitersPerLiter,
    }
    
    var liters_per_literResult *units.VolumeConcentration
    liters_per_literResult, err = factory.FromDto(liters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with LitersPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_literResult.Convert(units.VolumeConcentrationLitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LitersPerLiter = %v, want %v", converted, 100)
    }
    // Test LitersPerMililiter conversion
    liters_per_mililiterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationLitersPerMililiter,
    }
    
    var liters_per_mililiterResult *units.VolumeConcentration
    liters_per_mililiterResult, err = factory.FromDto(liters_per_mililiterDto)
    if err != nil {
        t.Errorf("FromDto() with LitersPerMililiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_mililiterResult.Convert(units.VolumeConcentrationLitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LitersPerMililiter = %v, want %v", converted, 100)
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
    // Test PicolitersPerLiter conversion
    picoliters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationPicolitersPerLiter,
    }
    
    var picoliters_per_literResult *units.VolumeConcentration
    picoliters_per_literResult, err = factory.FromDto(picoliters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with PicolitersPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoliters_per_literResult.Convert(units.VolumeConcentrationPicolitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicolitersPerLiter = %v, want %v", converted, 100)
    }
    // Test NanolitersPerLiter conversion
    nanoliters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationNanolitersPerLiter,
    }
    
    var nanoliters_per_literResult *units.VolumeConcentration
    nanoliters_per_literResult, err = factory.FromDto(nanoliters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with NanolitersPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_literResult.Convert(units.VolumeConcentrationNanolitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanolitersPerLiter = %v, want %v", converted, 100)
    }
    // Test MicrolitersPerLiter conversion
    microliters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationMicrolitersPerLiter,
    }
    
    var microliters_per_literResult *units.VolumeConcentration
    microliters_per_literResult, err = factory.FromDto(microliters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MicrolitersPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_literResult.Convert(units.VolumeConcentrationMicrolitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrolitersPerLiter = %v, want %v", converted, 100)
    }
    // Test MillilitersPerLiter conversion
    milliliters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationMillilitersPerLiter,
    }
    
    var milliliters_per_literResult *units.VolumeConcentration
    milliliters_per_literResult, err = factory.FromDto(milliliters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MillilitersPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_literResult.Convert(units.VolumeConcentrationMillilitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillilitersPerLiter = %v, want %v", converted, 100)
    }
    // Test CentilitersPerLiter conversion
    centiliters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationCentilitersPerLiter,
    }
    
    var centiliters_per_literResult *units.VolumeConcentration
    centiliters_per_literResult, err = factory.FromDto(centiliters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with CentilitersPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_literResult.Convert(units.VolumeConcentrationCentilitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentilitersPerLiter = %v, want %v", converted, 100)
    }
    // Test DecilitersPerLiter conversion
    deciliters_per_literDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationDecilitersPerLiter,
    }
    
    var deciliters_per_literResult *units.VolumeConcentration
    deciliters_per_literResult, err = factory.FromDto(deciliters_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with DecilitersPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_literResult.Convert(units.VolumeConcentrationDecilitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecilitersPerLiter = %v, want %v", converted, 100)
    }
    // Test PicolitersPerMililiter conversion
    picoliters_per_mililiterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationPicolitersPerMililiter,
    }
    
    var picoliters_per_mililiterResult *units.VolumeConcentration
    picoliters_per_mililiterResult, err = factory.FromDto(picoliters_per_mililiterDto)
    if err != nil {
        t.Errorf("FromDto() with PicolitersPerMililiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoliters_per_mililiterResult.Convert(units.VolumeConcentrationPicolitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicolitersPerMililiter = %v, want %v", converted, 100)
    }
    // Test NanolitersPerMililiter conversion
    nanoliters_per_mililiterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationNanolitersPerMililiter,
    }
    
    var nanoliters_per_mililiterResult *units.VolumeConcentration
    nanoliters_per_mililiterResult, err = factory.FromDto(nanoliters_per_mililiterDto)
    if err != nil {
        t.Errorf("FromDto() with NanolitersPerMililiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_mililiterResult.Convert(units.VolumeConcentrationNanolitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanolitersPerMililiter = %v, want %v", converted, 100)
    }
    // Test MicrolitersPerMililiter conversion
    microliters_per_mililiterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationMicrolitersPerMililiter,
    }
    
    var microliters_per_mililiterResult *units.VolumeConcentration
    microliters_per_mililiterResult, err = factory.FromDto(microliters_per_mililiterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrolitersPerMililiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_mililiterResult.Convert(units.VolumeConcentrationMicrolitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrolitersPerMililiter = %v, want %v", converted, 100)
    }
    // Test MillilitersPerMililiter conversion
    milliliters_per_mililiterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationMillilitersPerMililiter,
    }
    
    var milliliters_per_mililiterResult *units.VolumeConcentration
    milliliters_per_mililiterResult, err = factory.FromDto(milliliters_per_mililiterDto)
    if err != nil {
        t.Errorf("FromDto() with MillilitersPerMililiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_mililiterResult.Convert(units.VolumeConcentrationMillilitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillilitersPerMililiter = %v, want %v", converted, 100)
    }
    // Test CentilitersPerMililiter conversion
    centiliters_per_mililiterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationCentilitersPerMililiter,
    }
    
    var centiliters_per_mililiterResult *units.VolumeConcentration
    centiliters_per_mililiterResult, err = factory.FromDto(centiliters_per_mililiterDto)
    if err != nil {
        t.Errorf("FromDto() with CentilitersPerMililiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_mililiterResult.Convert(units.VolumeConcentrationCentilitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentilitersPerMililiter = %v, want %v", converted, 100)
    }
    // Test DecilitersPerMililiter conversion
    deciliters_per_mililiterDto := units.VolumeConcentrationDto{
        Value: 100,
        Unit:  units.VolumeConcentrationDecilitersPerMililiter,
    }
    
    var deciliters_per_mililiterResult *units.VolumeConcentration
    deciliters_per_mililiterResult, err = factory.FromDto(deciliters_per_mililiterDto)
    if err != nil {
        t.Errorf("FromDto() with DecilitersPerMililiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_mililiterResult.Convert(units.VolumeConcentrationDecilitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecilitersPerMililiter = %v, want %v", converted, 100)
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
    // Test JSON with LitersPerLiter unit
    liters_per_literJSON := []byte(`{"value": 100, "unit": "LitersPerLiter"}`)
    liters_per_literResult, err := factory.FromDtoJSON(liters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LitersPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_literResult.Convert(units.VolumeConcentrationLitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LitersPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with LitersPerMililiter unit
    liters_per_mililiterJSON := []byte(`{"value": 100, "unit": "LitersPerMililiter"}`)
    liters_per_mililiterResult, err := factory.FromDtoJSON(liters_per_mililiterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LitersPerMililiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_mililiterResult.Convert(units.VolumeConcentrationLitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LitersPerMililiter = %v, want %v", converted, 100)
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
    // Test JSON with PicolitersPerLiter unit
    picoliters_per_literJSON := []byte(`{"value": 100, "unit": "PicolitersPerLiter"}`)
    picoliters_per_literResult, err := factory.FromDtoJSON(picoliters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicolitersPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoliters_per_literResult.Convert(units.VolumeConcentrationPicolitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicolitersPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with NanolitersPerLiter unit
    nanoliters_per_literJSON := []byte(`{"value": 100, "unit": "NanolitersPerLiter"}`)
    nanoliters_per_literResult, err := factory.FromDtoJSON(nanoliters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanolitersPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_literResult.Convert(units.VolumeConcentrationNanolitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanolitersPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrolitersPerLiter unit
    microliters_per_literJSON := []byte(`{"value": 100, "unit": "MicrolitersPerLiter"}`)
    microliters_per_literResult, err := factory.FromDtoJSON(microliters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrolitersPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_literResult.Convert(units.VolumeConcentrationMicrolitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrolitersPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with MillilitersPerLiter unit
    milliliters_per_literJSON := []byte(`{"value": 100, "unit": "MillilitersPerLiter"}`)
    milliliters_per_literResult, err := factory.FromDtoJSON(milliliters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillilitersPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_literResult.Convert(units.VolumeConcentrationMillilitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillilitersPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with CentilitersPerLiter unit
    centiliters_per_literJSON := []byte(`{"value": 100, "unit": "CentilitersPerLiter"}`)
    centiliters_per_literResult, err := factory.FromDtoJSON(centiliters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentilitersPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_literResult.Convert(units.VolumeConcentrationCentilitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentilitersPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with DecilitersPerLiter unit
    deciliters_per_literJSON := []byte(`{"value": 100, "unit": "DecilitersPerLiter"}`)
    deciliters_per_literResult, err := factory.FromDtoJSON(deciliters_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecilitersPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_literResult.Convert(units.VolumeConcentrationDecilitersPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecilitersPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with PicolitersPerMililiter unit
    picoliters_per_mililiterJSON := []byte(`{"value": 100, "unit": "PicolitersPerMililiter"}`)
    picoliters_per_mililiterResult, err := factory.FromDtoJSON(picoliters_per_mililiterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicolitersPerMililiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picoliters_per_mililiterResult.Convert(units.VolumeConcentrationPicolitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicolitersPerMililiter = %v, want %v", converted, 100)
    }
    // Test JSON with NanolitersPerMililiter unit
    nanoliters_per_mililiterJSON := []byte(`{"value": 100, "unit": "NanolitersPerMililiter"}`)
    nanoliters_per_mililiterResult, err := factory.FromDtoJSON(nanoliters_per_mililiterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanolitersPerMililiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoliters_per_mililiterResult.Convert(units.VolumeConcentrationNanolitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanolitersPerMililiter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrolitersPerMililiter unit
    microliters_per_mililiterJSON := []byte(`{"value": 100, "unit": "MicrolitersPerMililiter"}`)
    microliters_per_mililiterResult, err := factory.FromDtoJSON(microliters_per_mililiterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrolitersPerMililiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microliters_per_mililiterResult.Convert(units.VolumeConcentrationMicrolitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrolitersPerMililiter = %v, want %v", converted, 100)
    }
    // Test JSON with MillilitersPerMililiter unit
    milliliters_per_mililiterJSON := []byte(`{"value": 100, "unit": "MillilitersPerMililiter"}`)
    milliliters_per_mililiterResult, err := factory.FromDtoJSON(milliliters_per_mililiterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillilitersPerMililiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliliters_per_mililiterResult.Convert(units.VolumeConcentrationMillilitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillilitersPerMililiter = %v, want %v", converted, 100)
    }
    // Test JSON with CentilitersPerMililiter unit
    centiliters_per_mililiterJSON := []byte(`{"value": 100, "unit": "CentilitersPerMililiter"}`)
    centiliters_per_mililiterResult, err := factory.FromDtoJSON(centiliters_per_mililiterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentilitersPerMililiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiliters_per_mililiterResult.Convert(units.VolumeConcentrationCentilitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentilitersPerMililiter = %v, want %v", converted, 100)
    }
    // Test JSON with DecilitersPerMililiter unit
    deciliters_per_mililiterJSON := []byte(`{"value": 100, "unit": "DecilitersPerMililiter"}`)
    deciliters_per_mililiterResult, err := factory.FromDtoJSON(deciliters_per_mililiterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecilitersPerMililiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciliters_per_mililiterResult.Convert(units.VolumeConcentrationDecilitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecilitersPerMililiter = %v, want %v", converted, 100)
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
    converted := result.Convert(units.VolumeConcentrationLitersPerLiter)
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
    converted = zeroResult.Convert(units.VolumeConcentrationLitersPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLitersPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromLitersPerMililiter function
func TestVolumeConcentrationFactory_FromLitersPerMililiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLitersPerMililiter(100)
    if err != nil {
        t.Errorf("FromLitersPerMililiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationLitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLitersPerMililiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLitersPerMililiter(math.NaN())
    if err == nil {
        t.Error("FromLitersPerMililiter() with NaN value should return error")
    }

    _, err = factory.FromLitersPerMililiter(math.Inf(1))
    if err == nil {
        t.Error("FromLitersPerMililiter() with +Inf value should return error")
    }

    _, err = factory.FromLitersPerMililiter(math.Inf(-1))
    if err == nil {
        t.Error("FromLitersPerMililiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLitersPerMililiter(0)
    if err != nil {
        t.Errorf("FromLitersPerMililiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationLitersPerMililiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLitersPerMililiter() with zero value = %v, want 0", converted)
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
    converted := result.Convert(units.VolumeConcentrationPicolitersPerLiter)
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
    converted = zeroResult.Convert(units.VolumeConcentrationPicolitersPerLiter)
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
    converted := result.Convert(units.VolumeConcentrationNanolitersPerLiter)
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
    converted = zeroResult.Convert(units.VolumeConcentrationNanolitersPerLiter)
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
    converted := result.Convert(units.VolumeConcentrationMicrolitersPerLiter)
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
    converted = zeroResult.Convert(units.VolumeConcentrationMicrolitersPerLiter)
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
    converted := result.Convert(units.VolumeConcentrationMillilitersPerLiter)
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
    converted = zeroResult.Convert(units.VolumeConcentrationMillilitersPerLiter)
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
    converted := result.Convert(units.VolumeConcentrationCentilitersPerLiter)
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
    converted = zeroResult.Convert(units.VolumeConcentrationCentilitersPerLiter)
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
    converted := result.Convert(units.VolumeConcentrationDecilitersPerLiter)
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
    converted = zeroResult.Convert(units.VolumeConcentrationDecilitersPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecilitersPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicolitersPerMililiter function
func TestVolumeConcentrationFactory_FromPicolitersPerMililiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicolitersPerMililiter(100)
    if err != nil {
        t.Errorf("FromPicolitersPerMililiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationPicolitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicolitersPerMililiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicolitersPerMililiter(math.NaN())
    if err == nil {
        t.Error("FromPicolitersPerMililiter() with NaN value should return error")
    }

    _, err = factory.FromPicolitersPerMililiter(math.Inf(1))
    if err == nil {
        t.Error("FromPicolitersPerMililiter() with +Inf value should return error")
    }

    _, err = factory.FromPicolitersPerMililiter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicolitersPerMililiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicolitersPerMililiter(0)
    if err != nil {
        t.Errorf("FromPicolitersPerMililiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationPicolitersPerMililiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicolitersPerMililiter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanolitersPerMililiter function
func TestVolumeConcentrationFactory_FromNanolitersPerMililiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanolitersPerMililiter(100)
    if err != nil {
        t.Errorf("FromNanolitersPerMililiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationNanolitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanolitersPerMililiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanolitersPerMililiter(math.NaN())
    if err == nil {
        t.Error("FromNanolitersPerMililiter() with NaN value should return error")
    }

    _, err = factory.FromNanolitersPerMililiter(math.Inf(1))
    if err == nil {
        t.Error("FromNanolitersPerMililiter() with +Inf value should return error")
    }

    _, err = factory.FromNanolitersPerMililiter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanolitersPerMililiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanolitersPerMililiter(0)
    if err != nil {
        t.Errorf("FromNanolitersPerMililiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationNanolitersPerMililiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanolitersPerMililiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrolitersPerMililiter function
func TestVolumeConcentrationFactory_FromMicrolitersPerMililiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrolitersPerMililiter(100)
    if err != nil {
        t.Errorf("FromMicrolitersPerMililiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationMicrolitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrolitersPerMililiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrolitersPerMililiter(math.NaN())
    if err == nil {
        t.Error("FromMicrolitersPerMililiter() with NaN value should return error")
    }

    _, err = factory.FromMicrolitersPerMililiter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrolitersPerMililiter() with +Inf value should return error")
    }

    _, err = factory.FromMicrolitersPerMililiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrolitersPerMililiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrolitersPerMililiter(0)
    if err != nil {
        t.Errorf("FromMicrolitersPerMililiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationMicrolitersPerMililiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrolitersPerMililiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillilitersPerMililiter function
func TestVolumeConcentrationFactory_FromMillilitersPerMililiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillilitersPerMililiter(100)
    if err != nil {
        t.Errorf("FromMillilitersPerMililiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationMillilitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillilitersPerMililiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillilitersPerMililiter(math.NaN())
    if err == nil {
        t.Error("FromMillilitersPerMililiter() with NaN value should return error")
    }

    _, err = factory.FromMillilitersPerMililiter(math.Inf(1))
    if err == nil {
        t.Error("FromMillilitersPerMililiter() with +Inf value should return error")
    }

    _, err = factory.FromMillilitersPerMililiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillilitersPerMililiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillilitersPerMililiter(0)
    if err != nil {
        t.Errorf("FromMillilitersPerMililiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationMillilitersPerMililiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillilitersPerMililiter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentilitersPerMililiter function
func TestVolumeConcentrationFactory_FromCentilitersPerMililiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentilitersPerMililiter(100)
    if err != nil {
        t.Errorf("FromCentilitersPerMililiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationCentilitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentilitersPerMililiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentilitersPerMililiter(math.NaN())
    if err == nil {
        t.Error("FromCentilitersPerMililiter() with NaN value should return error")
    }

    _, err = factory.FromCentilitersPerMililiter(math.Inf(1))
    if err == nil {
        t.Error("FromCentilitersPerMililiter() with +Inf value should return error")
    }

    _, err = factory.FromCentilitersPerMililiter(math.Inf(-1))
    if err == nil {
        t.Error("FromCentilitersPerMililiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentilitersPerMililiter(0)
    if err != nil {
        t.Errorf("FromCentilitersPerMililiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationCentilitersPerMililiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentilitersPerMililiter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecilitersPerMililiter function
func TestVolumeConcentrationFactory_FromDecilitersPerMililiter(t *testing.T) {
    factory := units.VolumeConcentrationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecilitersPerMililiter(100)
    if err != nil {
        t.Errorf("FromDecilitersPerMililiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeConcentrationDecilitersPerMililiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecilitersPerMililiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecilitersPerMililiter(math.NaN())
    if err == nil {
        t.Error("FromDecilitersPerMililiter() with NaN value should return error")
    }

    _, err = factory.FromDecilitersPerMililiter(math.Inf(1))
    if err == nil {
        t.Error("FromDecilitersPerMililiter() with +Inf value should return error")
    }

    _, err = factory.FromDecilitersPerMililiter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecilitersPerMililiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecilitersPerMililiter(0)
    if err != nil {
        t.Errorf("FromDecilitersPerMililiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeConcentrationDecilitersPerMililiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecilitersPerMililiter() with zero value = %v, want 0", converted)
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
            name: "LitersPerLiter abbreviation",
            unit: units.VolumeConcentrationLitersPerLiter,
            want: "L/L",
        },
        {
            name: "LitersPerMililiter abbreviation",
            unit: units.VolumeConcentrationLitersPerMililiter,
            want: "L/mL",
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
            name: "PicolitersPerLiter abbreviation",
            unit: units.VolumeConcentrationPicolitersPerLiter,
            want: "pL/L",
        },
        {
            name: "NanolitersPerLiter abbreviation",
            unit: units.VolumeConcentrationNanolitersPerLiter,
            want: "nL/L",
        },
        {
            name: "MicrolitersPerLiter abbreviation",
            unit: units.VolumeConcentrationMicrolitersPerLiter,
            want: "μL/L",
        },
        {
            name: "MillilitersPerLiter abbreviation",
            unit: units.VolumeConcentrationMillilitersPerLiter,
            want: "mL/L",
        },
        {
            name: "CentilitersPerLiter abbreviation",
            unit: units.VolumeConcentrationCentilitersPerLiter,
            want: "cL/L",
        },
        {
            name: "DecilitersPerLiter abbreviation",
            unit: units.VolumeConcentrationDecilitersPerLiter,
            want: "dL/L",
        },
        {
            name: "PicolitersPerMililiter abbreviation",
            unit: units.VolumeConcentrationPicolitersPerMililiter,
            want: "pL/mL",
        },
        {
            name: "NanolitersPerMililiter abbreviation",
            unit: units.VolumeConcentrationNanolitersPerMililiter,
            want: "nL/mL",
        },
        {
            name: "MicrolitersPerMililiter abbreviation",
            unit: units.VolumeConcentrationMicrolitersPerMililiter,
            want: "μL/mL",
        },
        {
            name: "MillilitersPerMililiter abbreviation",
            unit: units.VolumeConcentrationMillilitersPerMililiter,
            want: "mL/mL",
        },
        {
            name: "CentilitersPerMililiter abbreviation",
            unit: units.VolumeConcentrationCentilitersPerMililiter,
            want: "cL/mL",
        },
        {
            name: "DecilitersPerMililiter abbreviation",
            unit: units.VolumeConcentrationDecilitersPerMililiter,
            want: "dL/mL",
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