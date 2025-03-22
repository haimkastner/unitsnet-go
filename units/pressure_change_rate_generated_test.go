// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPressureChangeRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "PascalPerSecond"}`
	
	factory := units.PressureChangeRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PressureChangeRatePascalPerSecond {
		t.Errorf("expected unit %v, got %v", units.PressureChangeRatePascalPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "PascalPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPressureChangeRateDto_ToJSON(t *testing.T) {
	dto := units.PressureChangeRateDto{
		Value: 45,
		Unit:  units.PressureChangeRatePascalPerSecond,
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
	if result["unit"].(string) != string(units.PressureChangeRatePascalPerSecond) {
		t.Errorf("expected unit %s, got %v", units.PressureChangeRatePascalPerSecond, result["unit"])
	}
}

func TestNewPressureChangeRate_InvalidValue(t *testing.T) {
	factory := units.PressureChangeRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePressureChangeRate(math.NaN(), units.PressureChangeRatePascalPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePressureChangeRate(math.Inf(1), units.PressureChangeRatePascalPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPressureChangeRateConversions(t *testing.T) {
	factory := units.PressureChangeRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePressureChangeRate(180, units.PressureChangeRatePascalPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to PascalsPerSecond.
		// No expected conversion value provided for PascalsPerSecond, verifying result is not NaN.
		result := a.PascalsPerSecond()
		cacheResult := a.PascalsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PascalsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PascalsPerMinute.
		// No expected conversion value provided for PascalsPerMinute, verifying result is not NaN.
		result := a.PascalsPerMinute()
		cacheResult := a.PascalsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PascalsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MillimetersOfMercuryPerSecond.
		// No expected conversion value provided for MillimetersOfMercuryPerSecond, verifying result is not NaN.
		result := a.MillimetersOfMercuryPerSecond()
		cacheResult := a.MillimetersOfMercuryPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimetersOfMercuryPerSecond returned NaN")
		}
	}
	{
		// Test conversion to AtmospheresPerSecond.
		// No expected conversion value provided for AtmospheresPerSecond, verifying result is not NaN.
		result := a.AtmospheresPerSecond()
		cacheResult := a.AtmospheresPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AtmospheresPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerSquareInchPerSecond.
		// No expected conversion value provided for PoundsForcePerSquareInchPerSecond, verifying result is not NaN.
		result := a.PoundsForcePerSquareInchPerSecond()
		cacheResult := a.PoundsForcePerSquareInchPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsForcePerSquareInchPerSecond returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerSquareInchPerMinute.
		// No expected conversion value provided for PoundsForcePerSquareInchPerMinute, verifying result is not NaN.
		result := a.PoundsForcePerSquareInchPerMinute()
		cacheResult := a.PoundsForcePerSquareInchPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsForcePerSquareInchPerMinute returned NaN")
		}
	}
	{
		// Test conversion to BarsPerSecond.
		// No expected conversion value provided for BarsPerSecond, verifying result is not NaN.
		result := a.BarsPerSecond()
		cacheResult := a.BarsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to BarsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to BarsPerMinute.
		// No expected conversion value provided for BarsPerMinute, verifying result is not NaN.
		result := a.BarsPerMinute()
		cacheResult := a.BarsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to BarsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to KilopascalsPerSecond.
		// No expected conversion value provided for KilopascalsPerSecond, verifying result is not NaN.
		result := a.KilopascalsPerSecond()
		cacheResult := a.KilopascalsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopascalsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MegapascalsPerSecond.
		// No expected conversion value provided for MegapascalsPerSecond, verifying result is not NaN.
		result := a.MegapascalsPerSecond()
		cacheResult := a.MegapascalsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegapascalsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilopascalsPerMinute.
		// No expected conversion value provided for KilopascalsPerMinute, verifying result is not NaN.
		result := a.KilopascalsPerMinute()
		cacheResult := a.KilopascalsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopascalsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MegapascalsPerMinute.
		// No expected conversion value provided for MegapascalsPerMinute, verifying result is not NaN.
		result := a.MegapascalsPerMinute()
		cacheResult := a.MegapascalsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegapascalsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerSquareInchPerSecond.
		// No expected conversion value provided for KilopoundsForcePerSquareInchPerSecond, verifying result is not NaN.
		result := a.KilopoundsForcePerSquareInchPerSecond()
		cacheResult := a.KilopoundsForcePerSquareInchPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsForcePerSquareInchPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MegapoundsForcePerSquareInchPerSecond.
		// No expected conversion value provided for MegapoundsForcePerSquareInchPerSecond, verifying result is not NaN.
		result := a.MegapoundsForcePerSquareInchPerSecond()
		cacheResult := a.MegapoundsForcePerSquareInchPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegapoundsForcePerSquareInchPerSecond returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerSquareInchPerMinute.
		// No expected conversion value provided for KilopoundsForcePerSquareInchPerMinute, verifying result is not NaN.
		result := a.KilopoundsForcePerSquareInchPerMinute()
		cacheResult := a.KilopoundsForcePerSquareInchPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsForcePerSquareInchPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MegapoundsForcePerSquareInchPerMinute.
		// No expected conversion value provided for MegapoundsForcePerSquareInchPerMinute, verifying result is not NaN.
		result := a.MegapoundsForcePerSquareInchPerMinute()
		cacheResult := a.MegapoundsForcePerSquareInchPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegapoundsForcePerSquareInchPerMinute returned NaN")
		}
	}
	{
		// Test conversion to MillibarsPerSecond.
		// No expected conversion value provided for MillibarsPerSecond, verifying result is not NaN.
		result := a.MillibarsPerSecond()
		cacheResult := a.MillibarsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillibarsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillibarsPerMinute.
		// No expected conversion value provided for MillibarsPerMinute, verifying result is not NaN.
		result := a.MillibarsPerMinute()
		cacheResult := a.MillibarsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillibarsPerMinute returned NaN")
		}
	}
}

func TestPressureChangeRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PressureChangeRateFactory{}
	a, err := factory.CreatePressureChangeRate(90, units.PressureChangeRatePascalPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PressureChangeRatePascalPerSecond {
		t.Errorf("expected default unit PascalPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PressureChangeRatePascalPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PressureChangeRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PressureChangeRatePascalPerSecond {
		t.Errorf("expected unit PascalPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPressureChangeRateFactory_FromDto(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRatePascalPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.PressureChangeRateDto{
        Value: math.NaN(),
        Unit:  units.PressureChangeRatePascalPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test PascalPerSecond conversion
    pascals_per_secondDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRatePascalPerSecond,
    }
    
    var pascals_per_secondResult *units.PressureChangeRate
    pascals_per_secondResult, err = factory.FromDto(pascals_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PascalPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascals_per_secondResult.Convert(units.PressureChangeRatePascalPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalPerSecond = %v, want %v", converted, 100)
    }
    // Test PascalPerMinute conversion
    pascals_per_minuteDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRatePascalPerMinute,
    }
    
    var pascals_per_minuteResult *units.PressureChangeRate
    pascals_per_minuteResult, err = factory.FromDto(pascals_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with PascalPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascals_per_minuteResult.Convert(units.PressureChangeRatePascalPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalPerMinute = %v, want %v", converted, 100)
    }
    // Test MillimeterOfMercuryPerSecond conversion
    millimeters_of_mercury_per_secondDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateMillimeterOfMercuryPerSecond,
    }
    
    var millimeters_of_mercury_per_secondResult *units.PressureChangeRate
    millimeters_of_mercury_per_secondResult, err = factory.FromDto(millimeters_of_mercury_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterOfMercuryPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_of_mercury_per_secondResult.Convert(units.PressureChangeRateMillimeterOfMercuryPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterOfMercuryPerSecond = %v, want %v", converted, 100)
    }
    // Test AtmospherePerSecond conversion
    atmospheres_per_secondDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateAtmospherePerSecond,
    }
    
    var atmospheres_per_secondResult *units.PressureChangeRate
    atmospheres_per_secondResult, err = factory.FromDto(atmospheres_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with AtmospherePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = atmospheres_per_secondResult.Convert(units.PressureChangeRateAtmospherePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AtmospherePerSecond = %v, want %v", converted, 100)
    }
    // Test PoundForcePerSquareInchPerSecond conversion
    pounds_force_per_square_inch_per_secondDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRatePoundForcePerSquareInchPerSecond,
    }
    
    var pounds_force_per_square_inch_per_secondResult *units.PressureChangeRate
    pounds_force_per_square_inch_per_secondResult, err = factory.FromDto(pounds_force_per_square_inch_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForcePerSquareInchPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_square_inch_per_secondResult.Convert(units.PressureChangeRatePoundForcePerSquareInchPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerSquareInchPerSecond = %v, want %v", converted, 100)
    }
    // Test PoundForcePerSquareInchPerMinute conversion
    pounds_force_per_square_inch_per_minuteDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRatePoundForcePerSquareInchPerMinute,
    }
    
    var pounds_force_per_square_inch_per_minuteResult *units.PressureChangeRate
    pounds_force_per_square_inch_per_minuteResult, err = factory.FromDto(pounds_force_per_square_inch_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForcePerSquareInchPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_square_inch_per_minuteResult.Convert(units.PressureChangeRatePoundForcePerSquareInchPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerSquareInchPerMinute = %v, want %v", converted, 100)
    }
    // Test BarPerSecond conversion
    bars_per_secondDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateBarPerSecond,
    }
    
    var bars_per_secondResult *units.PressureChangeRate
    bars_per_secondResult, err = factory.FromDto(bars_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with BarPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = bars_per_secondResult.Convert(units.PressureChangeRateBarPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BarPerSecond = %v, want %v", converted, 100)
    }
    // Test BarPerMinute conversion
    bars_per_minuteDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateBarPerMinute,
    }
    
    var bars_per_minuteResult *units.PressureChangeRate
    bars_per_minuteResult, err = factory.FromDto(bars_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with BarPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = bars_per_minuteResult.Convert(units.PressureChangeRateBarPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BarPerMinute = %v, want %v", converted, 100)
    }
    // Test KilopascalPerSecond conversion
    kilopascals_per_secondDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateKilopascalPerSecond,
    }
    
    var kilopascals_per_secondResult *units.PressureChangeRate
    kilopascals_per_secondResult, err = factory.FromDto(kilopascals_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilopascalPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopascals_per_secondResult.Convert(units.PressureChangeRateKilopascalPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopascalPerSecond = %v, want %v", converted, 100)
    }
    // Test MegapascalPerSecond conversion
    megapascals_per_secondDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateMegapascalPerSecond,
    }
    
    var megapascals_per_secondResult *units.PressureChangeRate
    megapascals_per_secondResult, err = factory.FromDto(megapascals_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MegapascalPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapascals_per_secondResult.Convert(units.PressureChangeRateMegapascalPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapascalPerSecond = %v, want %v", converted, 100)
    }
    // Test KilopascalPerMinute conversion
    kilopascals_per_minuteDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateKilopascalPerMinute,
    }
    
    var kilopascals_per_minuteResult *units.PressureChangeRate
    kilopascals_per_minuteResult, err = factory.FromDto(kilopascals_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with KilopascalPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopascals_per_minuteResult.Convert(units.PressureChangeRateKilopascalPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopascalPerMinute = %v, want %v", converted, 100)
    }
    // Test MegapascalPerMinute conversion
    megapascals_per_minuteDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateMegapascalPerMinute,
    }
    
    var megapascals_per_minuteResult *units.PressureChangeRate
    megapascals_per_minuteResult, err = factory.FromDto(megapascals_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with MegapascalPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapascals_per_minuteResult.Convert(units.PressureChangeRateMegapascalPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapascalPerMinute = %v, want %v", converted, 100)
    }
    // Test KilopoundForcePerSquareInchPerSecond conversion
    kilopounds_force_per_square_inch_per_secondDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateKilopoundForcePerSquareInchPerSecond,
    }
    
    var kilopounds_force_per_square_inch_per_secondResult *units.PressureChangeRate
    kilopounds_force_per_square_inch_per_secondResult, err = factory.FromDto(kilopounds_force_per_square_inch_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForcePerSquareInchPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_square_inch_per_secondResult.Convert(units.PressureChangeRateKilopoundForcePerSquareInchPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerSquareInchPerSecond = %v, want %v", converted, 100)
    }
    // Test MegapoundForcePerSquareInchPerSecond conversion
    megapounds_force_per_square_inch_per_secondDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateMegapoundForcePerSquareInchPerSecond,
    }
    
    var megapounds_force_per_square_inch_per_secondResult *units.PressureChangeRate
    megapounds_force_per_square_inch_per_secondResult, err = factory.FromDto(megapounds_force_per_square_inch_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MegapoundForcePerSquareInchPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_force_per_square_inch_per_secondResult.Convert(units.PressureChangeRateMegapoundForcePerSquareInchPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundForcePerSquareInchPerSecond = %v, want %v", converted, 100)
    }
    // Test KilopoundForcePerSquareInchPerMinute conversion
    kilopounds_force_per_square_inch_per_minuteDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateKilopoundForcePerSquareInchPerMinute,
    }
    
    var kilopounds_force_per_square_inch_per_minuteResult *units.PressureChangeRate
    kilopounds_force_per_square_inch_per_minuteResult, err = factory.FromDto(kilopounds_force_per_square_inch_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForcePerSquareInchPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_square_inch_per_minuteResult.Convert(units.PressureChangeRateKilopoundForcePerSquareInchPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerSquareInchPerMinute = %v, want %v", converted, 100)
    }
    // Test MegapoundForcePerSquareInchPerMinute conversion
    megapounds_force_per_square_inch_per_minuteDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateMegapoundForcePerSquareInchPerMinute,
    }
    
    var megapounds_force_per_square_inch_per_minuteResult *units.PressureChangeRate
    megapounds_force_per_square_inch_per_minuteResult, err = factory.FromDto(megapounds_force_per_square_inch_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with MegapoundForcePerSquareInchPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_force_per_square_inch_per_minuteResult.Convert(units.PressureChangeRateMegapoundForcePerSquareInchPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundForcePerSquareInchPerMinute = %v, want %v", converted, 100)
    }
    // Test MillibarPerSecond conversion
    millibars_per_secondDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateMillibarPerSecond,
    }
    
    var millibars_per_secondResult *units.PressureChangeRate
    millibars_per_secondResult, err = factory.FromDto(millibars_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MillibarPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millibars_per_secondResult.Convert(units.PressureChangeRateMillibarPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillibarPerSecond = %v, want %v", converted, 100)
    }
    // Test MillibarPerMinute conversion
    millibars_per_minuteDto := units.PressureChangeRateDto{
        Value: 100,
        Unit:  units.PressureChangeRateMillibarPerMinute,
    }
    
    var millibars_per_minuteResult *units.PressureChangeRate
    millibars_per_minuteResult, err = factory.FromDto(millibars_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with MillibarPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millibars_per_minuteResult.Convert(units.PressureChangeRateMillibarPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillibarPerMinute = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.PressureChangeRateDto{
        Value: 0,
        Unit:  units.PressureChangeRatePascalPerSecond,
    }
    
    var zeroResult *units.PressureChangeRate
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestPressureChangeRateFactory_FromDtoJSON(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "PascalPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "PascalPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.PressureChangeRateDto{
        Value: nanValue,
        Unit:  units.PressureChangeRatePascalPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with PascalPerSecond unit
    pascals_per_secondJSON := []byte(`{"value": 100, "unit": "PascalPerSecond"}`)
    pascals_per_secondResult, err := factory.FromDtoJSON(pascals_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PascalPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascals_per_secondResult.Convert(units.PressureChangeRatePascalPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PascalPerMinute unit
    pascals_per_minuteJSON := []byte(`{"value": 100, "unit": "PascalPerMinute"}`)
    pascals_per_minuteResult, err := factory.FromDtoJSON(pascals_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PascalPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascals_per_minuteResult.Convert(units.PressureChangeRatePascalPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterOfMercuryPerSecond unit
    millimeters_of_mercury_per_secondJSON := []byte(`{"value": 100, "unit": "MillimeterOfMercuryPerSecond"}`)
    millimeters_of_mercury_per_secondResult, err := factory.FromDtoJSON(millimeters_of_mercury_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterOfMercuryPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_of_mercury_per_secondResult.Convert(units.PressureChangeRateMillimeterOfMercuryPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterOfMercuryPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with AtmospherePerSecond unit
    atmospheres_per_secondJSON := []byte(`{"value": 100, "unit": "AtmospherePerSecond"}`)
    atmospheres_per_secondResult, err := factory.FromDtoJSON(atmospheres_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AtmospherePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = atmospheres_per_secondResult.Convert(units.PressureChangeRateAtmospherePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AtmospherePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForcePerSquareInchPerSecond unit
    pounds_force_per_square_inch_per_secondJSON := []byte(`{"value": 100, "unit": "PoundForcePerSquareInchPerSecond"}`)
    pounds_force_per_square_inch_per_secondResult, err := factory.FromDtoJSON(pounds_force_per_square_inch_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForcePerSquareInchPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_square_inch_per_secondResult.Convert(units.PressureChangeRatePoundForcePerSquareInchPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerSquareInchPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForcePerSquareInchPerMinute unit
    pounds_force_per_square_inch_per_minuteJSON := []byte(`{"value": 100, "unit": "PoundForcePerSquareInchPerMinute"}`)
    pounds_force_per_square_inch_per_minuteResult, err := factory.FromDtoJSON(pounds_force_per_square_inch_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForcePerSquareInchPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_square_inch_per_minuteResult.Convert(units.PressureChangeRatePoundForcePerSquareInchPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerSquareInchPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with BarPerSecond unit
    bars_per_secondJSON := []byte(`{"value": 100, "unit": "BarPerSecond"}`)
    bars_per_secondResult, err := factory.FromDtoJSON(bars_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BarPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = bars_per_secondResult.Convert(units.PressureChangeRateBarPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BarPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with BarPerMinute unit
    bars_per_minuteJSON := []byte(`{"value": 100, "unit": "BarPerMinute"}`)
    bars_per_minuteResult, err := factory.FromDtoJSON(bars_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BarPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = bars_per_minuteResult.Convert(units.PressureChangeRateBarPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BarPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with KilopascalPerSecond unit
    kilopascals_per_secondJSON := []byte(`{"value": 100, "unit": "KilopascalPerSecond"}`)
    kilopascals_per_secondResult, err := factory.FromDtoJSON(kilopascals_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopascalPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopascals_per_secondResult.Convert(units.PressureChangeRateKilopascalPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopascalPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MegapascalPerSecond unit
    megapascals_per_secondJSON := []byte(`{"value": 100, "unit": "MegapascalPerSecond"}`)
    megapascals_per_secondResult, err := factory.FromDtoJSON(megapascals_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapascalPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapascals_per_secondResult.Convert(units.PressureChangeRateMegapascalPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapascalPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilopascalPerMinute unit
    kilopascals_per_minuteJSON := []byte(`{"value": 100, "unit": "KilopascalPerMinute"}`)
    kilopascals_per_minuteResult, err := factory.FromDtoJSON(kilopascals_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopascalPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopascals_per_minuteResult.Convert(units.PressureChangeRateKilopascalPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopascalPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MegapascalPerMinute unit
    megapascals_per_minuteJSON := []byte(`{"value": 100, "unit": "MegapascalPerMinute"}`)
    megapascals_per_minuteResult, err := factory.FromDtoJSON(megapascals_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapascalPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapascals_per_minuteResult.Convert(units.PressureChangeRateMegapascalPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapascalPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForcePerSquareInchPerSecond unit
    kilopounds_force_per_square_inch_per_secondJSON := []byte(`{"value": 100, "unit": "KilopoundForcePerSquareInchPerSecond"}`)
    kilopounds_force_per_square_inch_per_secondResult, err := factory.FromDtoJSON(kilopounds_force_per_square_inch_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForcePerSquareInchPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_square_inch_per_secondResult.Convert(units.PressureChangeRateKilopoundForcePerSquareInchPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerSquareInchPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MegapoundForcePerSquareInchPerSecond unit
    megapounds_force_per_square_inch_per_secondJSON := []byte(`{"value": 100, "unit": "MegapoundForcePerSquareInchPerSecond"}`)
    megapounds_force_per_square_inch_per_secondResult, err := factory.FromDtoJSON(megapounds_force_per_square_inch_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapoundForcePerSquareInchPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_force_per_square_inch_per_secondResult.Convert(units.PressureChangeRateMegapoundForcePerSquareInchPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundForcePerSquareInchPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForcePerSquareInchPerMinute unit
    kilopounds_force_per_square_inch_per_minuteJSON := []byte(`{"value": 100, "unit": "KilopoundForcePerSquareInchPerMinute"}`)
    kilopounds_force_per_square_inch_per_minuteResult, err := factory.FromDtoJSON(kilopounds_force_per_square_inch_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForcePerSquareInchPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_square_inch_per_minuteResult.Convert(units.PressureChangeRateKilopoundForcePerSquareInchPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerSquareInchPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MegapoundForcePerSquareInchPerMinute unit
    megapounds_force_per_square_inch_per_minuteJSON := []byte(`{"value": 100, "unit": "MegapoundForcePerSquareInchPerMinute"}`)
    megapounds_force_per_square_inch_per_minuteResult, err := factory.FromDtoJSON(megapounds_force_per_square_inch_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapoundForcePerSquareInchPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapounds_force_per_square_inch_per_minuteResult.Convert(units.PressureChangeRateMegapoundForcePerSquareInchPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundForcePerSquareInchPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with MillibarPerSecond unit
    millibars_per_secondJSON := []byte(`{"value": 100, "unit": "MillibarPerSecond"}`)
    millibars_per_secondResult, err := factory.FromDtoJSON(millibars_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillibarPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millibars_per_secondResult.Convert(units.PressureChangeRateMillibarPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillibarPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MillibarPerMinute unit
    millibars_per_minuteJSON := []byte(`{"value": 100, "unit": "MillibarPerMinute"}`)
    millibars_per_minuteResult, err := factory.FromDtoJSON(millibars_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillibarPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millibars_per_minuteResult.Convert(units.PressureChangeRateMillibarPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillibarPerMinute = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "PascalPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromPascalsPerSecond function
func TestPressureChangeRateFactory_FromPascalsPerSecond(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPascalsPerSecond(100)
    if err != nil {
        t.Errorf("FromPascalsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRatePascalPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPascalsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPascalsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPascalsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPascalsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPascalsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPascalsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPascalsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPascalsPerSecond(0)
    if err != nil {
        t.Errorf("FromPascalsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRatePascalPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPascalsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromPascalsPerMinute function
func TestPressureChangeRateFactory_FromPascalsPerMinute(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPascalsPerMinute(100)
    if err != nil {
        t.Errorf("FromPascalsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRatePascalPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPascalsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPascalsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromPascalsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromPascalsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromPascalsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromPascalsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromPascalsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPascalsPerMinute(0)
    if err != nil {
        t.Errorf("FromPascalsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRatePascalPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPascalsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimetersOfMercuryPerSecond function
func TestPressureChangeRateFactory_FromMillimetersOfMercuryPerSecond(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimetersOfMercuryPerSecond(100)
    if err != nil {
        t.Errorf("FromMillimetersOfMercuryPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateMillimeterOfMercuryPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimetersOfMercuryPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimetersOfMercuryPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMillimetersOfMercuryPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMillimetersOfMercuryPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMillimetersOfMercuryPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMillimetersOfMercuryPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimetersOfMercuryPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimetersOfMercuryPerSecond(0)
    if err != nil {
        t.Errorf("FromMillimetersOfMercuryPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateMillimeterOfMercuryPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimetersOfMercuryPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromAtmospheresPerSecond function
func TestPressureChangeRateFactory_FromAtmospheresPerSecond(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAtmospheresPerSecond(100)
    if err != nil {
        t.Errorf("FromAtmospheresPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateAtmospherePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAtmospheresPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAtmospheresPerSecond(math.NaN())
    if err == nil {
        t.Error("FromAtmospheresPerSecond() with NaN value should return error")
    }

    _, err = factory.FromAtmospheresPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromAtmospheresPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromAtmospheresPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromAtmospheresPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAtmospheresPerSecond(0)
    if err != nil {
        t.Errorf("FromAtmospheresPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateAtmospherePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAtmospheresPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForcePerSquareInchPerSecond function
func TestPressureChangeRateFactory_FromPoundsForcePerSquareInchPerSecond(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForcePerSquareInchPerSecond(100)
    if err != nil {
        t.Errorf("FromPoundsForcePerSquareInchPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRatePoundForcePerSquareInchPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForcePerSquareInchPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForcePerSquareInchPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPoundsForcePerSquareInchPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPoundsForcePerSquareInchPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForcePerSquareInchPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForcePerSquareInchPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForcePerSquareInchPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForcePerSquareInchPerSecond(0)
    if err != nil {
        t.Errorf("FromPoundsForcePerSquareInchPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRatePoundForcePerSquareInchPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForcePerSquareInchPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForcePerSquareInchPerMinute function
func TestPressureChangeRateFactory_FromPoundsForcePerSquareInchPerMinute(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForcePerSquareInchPerMinute(100)
    if err != nil {
        t.Errorf("FromPoundsForcePerSquareInchPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRatePoundForcePerSquareInchPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForcePerSquareInchPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForcePerSquareInchPerMinute(math.NaN())
    if err == nil {
        t.Error("FromPoundsForcePerSquareInchPerMinute() with NaN value should return error")
    }

    _, err = factory.FromPoundsForcePerSquareInchPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForcePerSquareInchPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForcePerSquareInchPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForcePerSquareInchPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForcePerSquareInchPerMinute(0)
    if err != nil {
        t.Errorf("FromPoundsForcePerSquareInchPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRatePoundForcePerSquareInchPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForcePerSquareInchPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromBarsPerSecond function
func TestPressureChangeRateFactory_FromBarsPerSecond(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBarsPerSecond(100)
    if err != nil {
        t.Errorf("FromBarsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateBarPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBarsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBarsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromBarsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromBarsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromBarsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromBarsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromBarsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBarsPerSecond(0)
    if err != nil {
        t.Errorf("FromBarsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateBarPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBarsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromBarsPerMinute function
func TestPressureChangeRateFactory_FromBarsPerMinute(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBarsPerMinute(100)
    if err != nil {
        t.Errorf("FromBarsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateBarPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBarsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBarsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromBarsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromBarsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromBarsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromBarsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromBarsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBarsPerMinute(0)
    if err != nil {
        t.Errorf("FromBarsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateBarPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBarsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopascalsPerSecond function
func TestPressureChangeRateFactory_FromKilopascalsPerSecond(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopascalsPerSecond(100)
    if err != nil {
        t.Errorf("FromKilopascalsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateKilopascalPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopascalsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopascalsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilopascalsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilopascalsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilopascalsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilopascalsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopascalsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopascalsPerSecond(0)
    if err != nil {
        t.Errorf("FromKilopascalsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateKilopascalPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopascalsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapascalsPerSecond function
func TestPressureChangeRateFactory_FromMegapascalsPerSecond(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapascalsPerSecond(100)
    if err != nil {
        t.Errorf("FromMegapascalsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateMegapascalPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapascalsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapascalsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMegapascalsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMegapascalsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMegapascalsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMegapascalsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapascalsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapascalsPerSecond(0)
    if err != nil {
        t.Errorf("FromMegapascalsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateMegapascalPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapascalsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopascalsPerMinute function
func TestPressureChangeRateFactory_FromKilopascalsPerMinute(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopascalsPerMinute(100)
    if err != nil {
        t.Errorf("FromKilopascalsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateKilopascalPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopascalsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopascalsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromKilopascalsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromKilopascalsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromKilopascalsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromKilopascalsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopascalsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopascalsPerMinute(0)
    if err != nil {
        t.Errorf("FromKilopascalsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateKilopascalPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopascalsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapascalsPerMinute function
func TestPressureChangeRateFactory_FromMegapascalsPerMinute(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapascalsPerMinute(100)
    if err != nil {
        t.Errorf("FromMegapascalsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateMegapascalPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapascalsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapascalsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromMegapascalsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromMegapascalsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromMegapascalsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromMegapascalsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapascalsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapascalsPerMinute(0)
    if err != nil {
        t.Errorf("FromMegapascalsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateMegapascalPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapascalsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsForcePerSquareInchPerSecond function
func TestPressureChangeRateFactory_FromKilopoundsForcePerSquareInchPerSecond(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsForcePerSquareInchPerSecond(100)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerSquareInchPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateKilopoundForcePerSquareInchPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerSquareInchPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsForcePerSquareInchPerSecond(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareInchPerSecond() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsForcePerSquareInchPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareInchPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsForcePerSquareInchPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareInchPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsForcePerSquareInchPerSecond(0)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerSquareInchPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateKilopoundForcePerSquareInchPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerSquareInchPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapoundsForcePerSquareInchPerSecond function
func TestPressureChangeRateFactory_FromMegapoundsForcePerSquareInchPerSecond(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapoundsForcePerSquareInchPerSecond(100)
    if err != nil {
        t.Errorf("FromMegapoundsForcePerSquareInchPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateMegapoundForcePerSquareInchPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapoundsForcePerSquareInchPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapoundsForcePerSquareInchPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMegapoundsForcePerSquareInchPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMegapoundsForcePerSquareInchPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMegapoundsForcePerSquareInchPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMegapoundsForcePerSquareInchPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapoundsForcePerSquareInchPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapoundsForcePerSquareInchPerSecond(0)
    if err != nil {
        t.Errorf("FromMegapoundsForcePerSquareInchPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateMegapoundForcePerSquareInchPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapoundsForcePerSquareInchPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsForcePerSquareInchPerMinute function
func TestPressureChangeRateFactory_FromKilopoundsForcePerSquareInchPerMinute(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsForcePerSquareInchPerMinute(100)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerSquareInchPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateKilopoundForcePerSquareInchPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerSquareInchPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsForcePerSquareInchPerMinute(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareInchPerMinute() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsForcePerSquareInchPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareInchPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsForcePerSquareInchPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareInchPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsForcePerSquareInchPerMinute(0)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerSquareInchPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateKilopoundForcePerSquareInchPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerSquareInchPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapoundsForcePerSquareInchPerMinute function
func TestPressureChangeRateFactory_FromMegapoundsForcePerSquareInchPerMinute(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapoundsForcePerSquareInchPerMinute(100)
    if err != nil {
        t.Errorf("FromMegapoundsForcePerSquareInchPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateMegapoundForcePerSquareInchPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapoundsForcePerSquareInchPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapoundsForcePerSquareInchPerMinute(math.NaN())
    if err == nil {
        t.Error("FromMegapoundsForcePerSquareInchPerMinute() with NaN value should return error")
    }

    _, err = factory.FromMegapoundsForcePerSquareInchPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromMegapoundsForcePerSquareInchPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromMegapoundsForcePerSquareInchPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapoundsForcePerSquareInchPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapoundsForcePerSquareInchPerMinute(0)
    if err != nil {
        t.Errorf("FromMegapoundsForcePerSquareInchPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateMegapoundForcePerSquareInchPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapoundsForcePerSquareInchPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromMillibarsPerSecond function
func TestPressureChangeRateFactory_FromMillibarsPerSecond(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillibarsPerSecond(100)
    if err != nil {
        t.Errorf("FromMillibarsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateMillibarPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillibarsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillibarsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMillibarsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMillibarsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMillibarsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMillibarsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMillibarsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillibarsPerSecond(0)
    if err != nil {
        t.Errorf("FromMillibarsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateMillibarPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillibarsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMillibarsPerMinute function
func TestPressureChangeRateFactory_FromMillibarsPerMinute(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillibarsPerMinute(100)
    if err != nil {
        t.Errorf("FromMillibarsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureChangeRateMillibarPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillibarsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillibarsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromMillibarsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromMillibarsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromMillibarsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromMillibarsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromMillibarsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillibarsPerMinute(0)
    if err != nil {
        t.Errorf("FromMillibarsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureChangeRateMillibarPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillibarsPerMinute() with zero value = %v, want 0", converted)
    }
}

func TestPressureChangeRateToString(t *testing.T) {
	factory := units.PressureChangeRateFactory{}
	a, err := factory.CreatePressureChangeRate(45, units.PressureChangeRatePascalPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PressureChangeRatePascalPerSecond, 2)
	expected := "45.00 " + units.GetPressureChangeRateAbbreviation(units.PressureChangeRatePascalPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PressureChangeRatePascalPerSecond, -1)
	expected = "45 " + units.GetPressureChangeRateAbbreviation(units.PressureChangeRatePascalPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPressureChangeRate_EqualityAndComparison(t *testing.T) {
	factory := units.PressureChangeRateFactory{}
	a1, _ := factory.CreatePressureChangeRate(60, units.PressureChangeRatePascalPerSecond)
	a2, _ := factory.CreatePressureChangeRate(60, units.PressureChangeRatePascalPerSecond)
	a3, _ := factory.CreatePressureChangeRate(120, units.PressureChangeRatePascalPerSecond)

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

func TestPressureChangeRate_Arithmetic(t *testing.T) {
	factory := units.PressureChangeRateFactory{}
	a1, _ := factory.CreatePressureChangeRate(30, units.PressureChangeRatePascalPerSecond)
	a2, _ := factory.CreatePressureChangeRate(45, units.PressureChangeRatePascalPerSecond)

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


func TestGetPressureChangeRateAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.PressureChangeRateUnits
        want string
    }{
        {
            name: "PascalPerSecond abbreviation",
            unit: units.PressureChangeRatePascalPerSecond,
            want: "Pa/s",
        },
        {
            name: "PascalPerMinute abbreviation",
            unit: units.PressureChangeRatePascalPerMinute,
            want: "Pa/min",
        },
        {
            name: "MillimeterOfMercuryPerSecond abbreviation",
            unit: units.PressureChangeRateMillimeterOfMercuryPerSecond,
            want: "mmHg/s",
        },
        {
            name: "AtmospherePerSecond abbreviation",
            unit: units.PressureChangeRateAtmospherePerSecond,
            want: "atm/s",
        },
        {
            name: "PoundForcePerSquareInchPerSecond abbreviation",
            unit: units.PressureChangeRatePoundForcePerSquareInchPerSecond,
            want: "psi/s",
        },
        {
            name: "PoundForcePerSquareInchPerMinute abbreviation",
            unit: units.PressureChangeRatePoundForcePerSquareInchPerMinute,
            want: "psi/min",
        },
        {
            name: "BarPerSecond abbreviation",
            unit: units.PressureChangeRateBarPerSecond,
            want: "bar/s",
        },
        {
            name: "BarPerMinute abbreviation",
            unit: units.PressureChangeRateBarPerMinute,
            want: "bar/min",
        },
        {
            name: "KilopascalPerSecond abbreviation",
            unit: units.PressureChangeRateKilopascalPerSecond,
            want: "kPa/s",
        },
        {
            name: "MegapascalPerSecond abbreviation",
            unit: units.PressureChangeRateMegapascalPerSecond,
            want: "MPa/s",
        },
        {
            name: "KilopascalPerMinute abbreviation",
            unit: units.PressureChangeRateKilopascalPerMinute,
            want: "kPa/min",
        },
        {
            name: "MegapascalPerMinute abbreviation",
            unit: units.PressureChangeRateMegapascalPerMinute,
            want: "MPa/min",
        },
        {
            name: "KilopoundForcePerSquareInchPerSecond abbreviation",
            unit: units.PressureChangeRateKilopoundForcePerSquareInchPerSecond,
            want: "kpsi/s",
        },
        {
            name: "MegapoundForcePerSquareInchPerSecond abbreviation",
            unit: units.PressureChangeRateMegapoundForcePerSquareInchPerSecond,
            want: "Mpsi/s",
        },
        {
            name: "KilopoundForcePerSquareInchPerMinute abbreviation",
            unit: units.PressureChangeRateKilopoundForcePerSquareInchPerMinute,
            want: "kpsi/min",
        },
        {
            name: "MegapoundForcePerSquareInchPerMinute abbreviation",
            unit: units.PressureChangeRateMegapoundForcePerSquareInchPerMinute,
            want: "Mpsi/min",
        },
        {
            name: "MillibarPerSecond abbreviation",
            unit: units.PressureChangeRateMillibarPerSecond,
            want: "mbar/s",
        },
        {
            name: "MillibarPerMinute abbreviation",
            unit: units.PressureChangeRateMillibarPerMinute,
            want: "mbar/min",
        },
        {
            name: "invalid unit",
            unit: units.PressureChangeRateUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetPressureChangeRateAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetPressureChangeRateAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestPressureChangeRate_String(t *testing.T) {
    factory := units.PressureChangeRateFactory{}
    
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
            unit, err := factory.CreatePressureChangeRate(tt.value, units.PressureChangeRatePascalPerSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("PressureChangeRate.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestPressureChangeRate_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.PressureChangeRateFactory{}

	_, err := uf.CreatePressureChangeRate(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}