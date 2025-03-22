// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLuminanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CandelaPerSquareMeter"}`
	
	factory := units.LuminanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LuminanceCandelaPerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.LuminanceCandelaPerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CandelaPerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLuminanceDto_ToJSON(t *testing.T) {
	dto := units.LuminanceDto{
		Value: 45,
		Unit:  units.LuminanceCandelaPerSquareMeter,
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
	if result["unit"].(string) != string(units.LuminanceCandelaPerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.LuminanceCandelaPerSquareMeter, result["unit"])
	}
}

func TestNewLuminance_InvalidValue(t *testing.T) {
	factory := units.LuminanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLuminance(math.NaN(), units.LuminanceCandelaPerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLuminance(math.Inf(1), units.LuminanceCandelaPerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLuminanceConversions(t *testing.T) {
	factory := units.LuminanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLuminance(180, units.LuminanceCandelaPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to CandelasPerSquareMeter.
		// No expected conversion value provided for CandelasPerSquareMeter, verifying result is not NaN.
		result := a.CandelasPerSquareMeter()
		cacheResult := a.CandelasPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CandelasPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to CandelasPerSquareFoot.
		// No expected conversion value provided for CandelasPerSquareFoot, verifying result is not NaN.
		result := a.CandelasPerSquareFoot()
		cacheResult := a.CandelasPerSquareFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CandelasPerSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to CandelasPerSquareInch.
		// No expected conversion value provided for CandelasPerSquareInch, verifying result is not NaN.
		result := a.CandelasPerSquareInch()
		cacheResult := a.CandelasPerSquareInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CandelasPerSquareInch returned NaN")
		}
	}
	{
		// Test conversion to Nits.
		// No expected conversion value provided for Nits, verifying result is not NaN.
		result := a.Nits()
		cacheResult := a.Nits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nits returned NaN")
		}
	}
	{
		// Test conversion to NanocandelasPerSquareMeter.
		// No expected conversion value provided for NanocandelasPerSquareMeter, verifying result is not NaN.
		result := a.NanocandelasPerSquareMeter()
		cacheResult := a.NanocandelasPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanocandelasPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MicrocandelasPerSquareMeter.
		// No expected conversion value provided for MicrocandelasPerSquareMeter, verifying result is not NaN.
		result := a.MicrocandelasPerSquareMeter()
		cacheResult := a.MicrocandelasPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrocandelasPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MillicandelasPerSquareMeter.
		// No expected conversion value provided for MillicandelasPerSquareMeter, verifying result is not NaN.
		result := a.MillicandelasPerSquareMeter()
		cacheResult := a.MillicandelasPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillicandelasPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to CenticandelasPerSquareMeter.
		// No expected conversion value provided for CenticandelasPerSquareMeter, verifying result is not NaN.
		result := a.CenticandelasPerSquareMeter()
		cacheResult := a.CenticandelasPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CenticandelasPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to DecicandelasPerSquareMeter.
		// No expected conversion value provided for DecicandelasPerSquareMeter, verifying result is not NaN.
		result := a.DecicandelasPerSquareMeter()
		cacheResult := a.DecicandelasPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecicandelasPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilocandelasPerSquareMeter.
		// No expected conversion value provided for KilocandelasPerSquareMeter, verifying result is not NaN.
		result := a.KilocandelasPerSquareMeter()
		cacheResult := a.KilocandelasPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilocandelasPerSquareMeter returned NaN")
		}
	}
}

func TestLuminance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LuminanceFactory{}
	a, err := factory.CreateLuminance(90, units.LuminanceCandelaPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LuminanceCandelaPerSquareMeter {
		t.Errorf("expected default unit CandelaPerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LuminanceCandelaPerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LuminanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LuminanceCandelaPerSquareMeter {
		t.Errorf("expected unit CandelaPerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLuminanceFactory_FromDto(t *testing.T) {
    factory := units.LuminanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.LuminanceDto{
        Value: 100,
        Unit:  units.LuminanceCandelaPerSquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.LuminanceDto{
        Value: math.NaN(),
        Unit:  units.LuminanceCandelaPerSquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test CandelaPerSquareMeter conversion
    candelas_per_square_meterDto := units.LuminanceDto{
        Value: 100,
        Unit:  units.LuminanceCandelaPerSquareMeter,
    }
    
    var candelas_per_square_meterResult *units.Luminance
    candelas_per_square_meterResult, err = factory.FromDto(candelas_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with CandelaPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = candelas_per_square_meterResult.Convert(units.LuminanceCandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CandelaPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test CandelaPerSquareFoot conversion
    candelas_per_square_footDto := units.LuminanceDto{
        Value: 100,
        Unit:  units.LuminanceCandelaPerSquareFoot,
    }
    
    var candelas_per_square_footResult *units.Luminance
    candelas_per_square_footResult, err = factory.FromDto(candelas_per_square_footDto)
    if err != nil {
        t.Errorf("FromDto() with CandelaPerSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = candelas_per_square_footResult.Convert(units.LuminanceCandelaPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CandelaPerSquareFoot = %v, want %v", converted, 100)
    }
    // Test CandelaPerSquareInch conversion
    candelas_per_square_inchDto := units.LuminanceDto{
        Value: 100,
        Unit:  units.LuminanceCandelaPerSquareInch,
    }
    
    var candelas_per_square_inchResult *units.Luminance
    candelas_per_square_inchResult, err = factory.FromDto(candelas_per_square_inchDto)
    if err != nil {
        t.Errorf("FromDto() with CandelaPerSquareInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = candelas_per_square_inchResult.Convert(units.LuminanceCandelaPerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CandelaPerSquareInch = %v, want %v", converted, 100)
    }
    // Test Nit conversion
    nitsDto := units.LuminanceDto{
        Value: 100,
        Unit:  units.LuminanceNit,
    }
    
    var nitsResult *units.Luminance
    nitsResult, err = factory.FromDto(nitsDto)
    if err != nil {
        t.Errorf("FromDto() with Nit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nitsResult.Convert(units.LuminanceNit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nit = %v, want %v", converted, 100)
    }
    // Test NanocandelaPerSquareMeter conversion
    nanocandelas_per_square_meterDto := units.LuminanceDto{
        Value: 100,
        Unit:  units.LuminanceNanocandelaPerSquareMeter,
    }
    
    var nanocandelas_per_square_meterResult *units.Luminance
    nanocandelas_per_square_meterResult, err = factory.FromDto(nanocandelas_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with NanocandelaPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanocandelas_per_square_meterResult.Convert(units.LuminanceNanocandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanocandelaPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test MicrocandelaPerSquareMeter conversion
    microcandelas_per_square_meterDto := units.LuminanceDto{
        Value: 100,
        Unit:  units.LuminanceMicrocandelaPerSquareMeter,
    }
    
    var microcandelas_per_square_meterResult *units.Luminance
    microcandelas_per_square_meterResult, err = factory.FromDto(microcandelas_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrocandelaPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microcandelas_per_square_meterResult.Convert(units.LuminanceMicrocandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrocandelaPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test MillicandelaPerSquareMeter conversion
    millicandelas_per_square_meterDto := units.LuminanceDto{
        Value: 100,
        Unit:  units.LuminanceMillicandelaPerSquareMeter,
    }
    
    var millicandelas_per_square_meterResult *units.Luminance
    millicandelas_per_square_meterResult, err = factory.FromDto(millicandelas_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MillicandelaPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millicandelas_per_square_meterResult.Convert(units.LuminanceMillicandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillicandelaPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test CenticandelaPerSquareMeter conversion
    centicandelas_per_square_meterDto := units.LuminanceDto{
        Value: 100,
        Unit:  units.LuminanceCenticandelaPerSquareMeter,
    }
    
    var centicandelas_per_square_meterResult *units.Luminance
    centicandelas_per_square_meterResult, err = factory.FromDto(centicandelas_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with CenticandelaPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centicandelas_per_square_meterResult.Convert(units.LuminanceCenticandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CenticandelaPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test DecicandelaPerSquareMeter conversion
    decicandelas_per_square_meterDto := units.LuminanceDto{
        Value: 100,
        Unit:  units.LuminanceDecicandelaPerSquareMeter,
    }
    
    var decicandelas_per_square_meterResult *units.Luminance
    decicandelas_per_square_meterResult, err = factory.FromDto(decicandelas_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with DecicandelaPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decicandelas_per_square_meterResult.Convert(units.LuminanceDecicandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecicandelaPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test KilocandelaPerSquareMeter conversion
    kilocandelas_per_square_meterDto := units.LuminanceDto{
        Value: 100,
        Unit:  units.LuminanceKilocandelaPerSquareMeter,
    }
    
    var kilocandelas_per_square_meterResult *units.Luminance
    kilocandelas_per_square_meterResult, err = factory.FromDto(kilocandelas_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilocandelaPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocandelas_per_square_meterResult.Convert(units.LuminanceKilocandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocandelaPerSquareMeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.LuminanceDto{
        Value: 0,
        Unit:  units.LuminanceCandelaPerSquareMeter,
    }
    
    var zeroResult *units.Luminance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestLuminanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.LuminanceFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "CandelaPerSquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "CandelaPerSquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.LuminanceDto{
        Value: nanValue,
        Unit:  units.LuminanceCandelaPerSquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with CandelaPerSquareMeter unit
    candelas_per_square_meterJSON := []byte(`{"value": 100, "unit": "CandelaPerSquareMeter"}`)
    candelas_per_square_meterResult, err := factory.FromDtoJSON(candelas_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CandelaPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = candelas_per_square_meterResult.Convert(units.LuminanceCandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CandelaPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with CandelaPerSquareFoot unit
    candelas_per_square_footJSON := []byte(`{"value": 100, "unit": "CandelaPerSquareFoot"}`)
    candelas_per_square_footResult, err := factory.FromDtoJSON(candelas_per_square_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CandelaPerSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = candelas_per_square_footResult.Convert(units.LuminanceCandelaPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CandelaPerSquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with CandelaPerSquareInch unit
    candelas_per_square_inchJSON := []byte(`{"value": 100, "unit": "CandelaPerSquareInch"}`)
    candelas_per_square_inchResult, err := factory.FromDtoJSON(candelas_per_square_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CandelaPerSquareInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = candelas_per_square_inchResult.Convert(units.LuminanceCandelaPerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CandelaPerSquareInch = %v, want %v", converted, 100)
    }
    // Test JSON with Nit unit
    nitsJSON := []byte(`{"value": 100, "unit": "Nit"}`)
    nitsResult, err := factory.FromDtoJSON(nitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nitsResult.Convert(units.LuminanceNit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nit = %v, want %v", converted, 100)
    }
    // Test JSON with NanocandelaPerSquareMeter unit
    nanocandelas_per_square_meterJSON := []byte(`{"value": 100, "unit": "NanocandelaPerSquareMeter"}`)
    nanocandelas_per_square_meterResult, err := factory.FromDtoJSON(nanocandelas_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanocandelaPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanocandelas_per_square_meterResult.Convert(units.LuminanceNanocandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanocandelaPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrocandelaPerSquareMeter unit
    microcandelas_per_square_meterJSON := []byte(`{"value": 100, "unit": "MicrocandelaPerSquareMeter"}`)
    microcandelas_per_square_meterResult, err := factory.FromDtoJSON(microcandelas_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrocandelaPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microcandelas_per_square_meterResult.Convert(units.LuminanceMicrocandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrocandelaPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MillicandelaPerSquareMeter unit
    millicandelas_per_square_meterJSON := []byte(`{"value": 100, "unit": "MillicandelaPerSquareMeter"}`)
    millicandelas_per_square_meterResult, err := factory.FromDtoJSON(millicandelas_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillicandelaPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millicandelas_per_square_meterResult.Convert(units.LuminanceMillicandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillicandelaPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with CenticandelaPerSquareMeter unit
    centicandelas_per_square_meterJSON := []byte(`{"value": 100, "unit": "CenticandelaPerSquareMeter"}`)
    centicandelas_per_square_meterResult, err := factory.FromDtoJSON(centicandelas_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CenticandelaPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centicandelas_per_square_meterResult.Convert(units.LuminanceCenticandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CenticandelaPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with DecicandelaPerSquareMeter unit
    decicandelas_per_square_meterJSON := []byte(`{"value": 100, "unit": "DecicandelaPerSquareMeter"}`)
    decicandelas_per_square_meterResult, err := factory.FromDtoJSON(decicandelas_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecicandelaPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decicandelas_per_square_meterResult.Convert(units.LuminanceDecicandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecicandelaPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilocandelaPerSquareMeter unit
    kilocandelas_per_square_meterJSON := []byte(`{"value": 100, "unit": "KilocandelaPerSquareMeter"}`)
    kilocandelas_per_square_meterResult, err := factory.FromDtoJSON(kilocandelas_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilocandelaPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocandelas_per_square_meterResult.Convert(units.LuminanceKilocandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocandelaPerSquareMeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "CandelaPerSquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromCandelasPerSquareMeter function
func TestLuminanceFactory_FromCandelasPerSquareMeter(t *testing.T) {
    factory := units.LuminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCandelasPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromCandelasPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminanceCandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCandelasPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCandelasPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromCandelasPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromCandelasPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromCandelasPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromCandelasPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromCandelasPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCandelasPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromCandelasPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminanceCandelaPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCandelasPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromCandelasPerSquareFoot function
func TestLuminanceFactory_FromCandelasPerSquareFoot(t *testing.T) {
    factory := units.LuminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCandelasPerSquareFoot(100)
    if err != nil {
        t.Errorf("FromCandelasPerSquareFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminanceCandelaPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCandelasPerSquareFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCandelasPerSquareFoot(math.NaN())
    if err == nil {
        t.Error("FromCandelasPerSquareFoot() with NaN value should return error")
    }

    _, err = factory.FromCandelasPerSquareFoot(math.Inf(1))
    if err == nil {
        t.Error("FromCandelasPerSquareFoot() with +Inf value should return error")
    }

    _, err = factory.FromCandelasPerSquareFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromCandelasPerSquareFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCandelasPerSquareFoot(0)
    if err != nil {
        t.Errorf("FromCandelasPerSquareFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminanceCandelaPerSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCandelasPerSquareFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromCandelasPerSquareInch function
func TestLuminanceFactory_FromCandelasPerSquareInch(t *testing.T) {
    factory := units.LuminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCandelasPerSquareInch(100)
    if err != nil {
        t.Errorf("FromCandelasPerSquareInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminanceCandelaPerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCandelasPerSquareInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCandelasPerSquareInch(math.NaN())
    if err == nil {
        t.Error("FromCandelasPerSquareInch() with NaN value should return error")
    }

    _, err = factory.FromCandelasPerSquareInch(math.Inf(1))
    if err == nil {
        t.Error("FromCandelasPerSquareInch() with +Inf value should return error")
    }

    _, err = factory.FromCandelasPerSquareInch(math.Inf(-1))
    if err == nil {
        t.Error("FromCandelasPerSquareInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCandelasPerSquareInch(0)
    if err != nil {
        t.Errorf("FromCandelasPerSquareInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminanceCandelaPerSquareInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCandelasPerSquareInch() with zero value = %v, want 0", converted)
    }
}
// Test FromNits function
func TestLuminanceFactory_FromNits(t *testing.T) {
    factory := units.LuminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNits(100)
    if err != nil {
        t.Errorf("FromNits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminanceNit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNits(math.NaN())
    if err == nil {
        t.Error("FromNits() with NaN value should return error")
    }

    _, err = factory.FromNits(math.Inf(1))
    if err == nil {
        t.Error("FromNits() with +Inf value should return error")
    }

    _, err = factory.FromNits(math.Inf(-1))
    if err == nil {
        t.Error("FromNits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNits(0)
    if err != nil {
        t.Errorf("FromNits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminanceNit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNits() with zero value = %v, want 0", converted)
    }
}
// Test FromNanocandelasPerSquareMeter function
func TestLuminanceFactory_FromNanocandelasPerSquareMeter(t *testing.T) {
    factory := units.LuminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanocandelasPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromNanocandelasPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminanceNanocandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanocandelasPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanocandelasPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromNanocandelasPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromNanocandelasPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromNanocandelasPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromNanocandelasPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanocandelasPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanocandelasPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromNanocandelasPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminanceNanocandelaPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanocandelasPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrocandelasPerSquareMeter function
func TestLuminanceFactory_FromMicrocandelasPerSquareMeter(t *testing.T) {
    factory := units.LuminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrocandelasPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromMicrocandelasPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminanceMicrocandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrocandelasPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrocandelasPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromMicrocandelasPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromMicrocandelasPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrocandelasPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromMicrocandelasPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrocandelasPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrocandelasPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromMicrocandelasPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminanceMicrocandelaPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrocandelasPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillicandelasPerSquareMeter function
func TestLuminanceFactory_FromMillicandelasPerSquareMeter(t *testing.T) {
    factory := units.LuminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillicandelasPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromMillicandelasPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminanceMillicandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillicandelasPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillicandelasPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromMillicandelasPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromMillicandelasPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMillicandelasPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromMillicandelasPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillicandelasPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillicandelasPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromMillicandelasPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminanceMillicandelaPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillicandelasPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromCenticandelasPerSquareMeter function
func TestLuminanceFactory_FromCenticandelasPerSquareMeter(t *testing.T) {
    factory := units.LuminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCenticandelasPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromCenticandelasPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminanceCenticandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCenticandelasPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCenticandelasPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromCenticandelasPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromCenticandelasPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromCenticandelasPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromCenticandelasPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromCenticandelasPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCenticandelasPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromCenticandelasPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminanceCenticandelaPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCenticandelasPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecicandelasPerSquareMeter function
func TestLuminanceFactory_FromDecicandelasPerSquareMeter(t *testing.T) {
    factory := units.LuminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecicandelasPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromDecicandelasPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminanceDecicandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecicandelasPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecicandelasPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromDecicandelasPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromDecicandelasPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromDecicandelasPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromDecicandelasPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecicandelasPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecicandelasPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromDecicandelasPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminanceDecicandelaPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecicandelasPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilocandelasPerSquareMeter function
func TestLuminanceFactory_FromKilocandelasPerSquareMeter(t *testing.T) {
    factory := units.LuminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilocandelasPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromKilocandelasPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminanceKilocandelaPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilocandelasPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilocandelasPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromKilocandelasPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromKilocandelasPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilocandelasPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilocandelasPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilocandelasPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilocandelasPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromKilocandelasPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminanceKilocandelaPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilocandelasPerSquareMeter() with zero value = %v, want 0", converted)
    }
}

func TestLuminanceToString(t *testing.T) {
	factory := units.LuminanceFactory{}
	a, err := factory.CreateLuminance(45, units.LuminanceCandelaPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LuminanceCandelaPerSquareMeter, 2)
	expected := "45.00 " + units.GetLuminanceAbbreviation(units.LuminanceCandelaPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LuminanceCandelaPerSquareMeter, -1)
	expected = "45 " + units.GetLuminanceAbbreviation(units.LuminanceCandelaPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLuminance_EqualityAndComparison(t *testing.T) {
	factory := units.LuminanceFactory{}
	a1, _ := factory.CreateLuminance(60, units.LuminanceCandelaPerSquareMeter)
	a2, _ := factory.CreateLuminance(60, units.LuminanceCandelaPerSquareMeter)
	a3, _ := factory.CreateLuminance(120, units.LuminanceCandelaPerSquareMeter)

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

func TestLuminance_Arithmetic(t *testing.T) {
	factory := units.LuminanceFactory{}
	a1, _ := factory.CreateLuminance(30, units.LuminanceCandelaPerSquareMeter)
	a2, _ := factory.CreateLuminance(45, units.LuminanceCandelaPerSquareMeter)

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


func TestGetLuminanceAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.LuminanceUnits
        want string
    }{
        {
            name: "CandelaPerSquareMeter abbreviation",
            unit: units.LuminanceCandelaPerSquareMeter,
            want: "Cd/m²",
        },
        {
            name: "CandelaPerSquareFoot abbreviation",
            unit: units.LuminanceCandelaPerSquareFoot,
            want: "Cd/ft²",
        },
        {
            name: "CandelaPerSquareInch abbreviation",
            unit: units.LuminanceCandelaPerSquareInch,
            want: "Cd/in²",
        },
        {
            name: "Nit abbreviation",
            unit: units.LuminanceNit,
            want: "nt",
        },
        {
            name: "NanocandelaPerSquareMeter abbreviation",
            unit: units.LuminanceNanocandelaPerSquareMeter,
            want: "nCd/m²",
        },
        {
            name: "MicrocandelaPerSquareMeter abbreviation",
            unit: units.LuminanceMicrocandelaPerSquareMeter,
            want: "μCd/m²",
        },
        {
            name: "MillicandelaPerSquareMeter abbreviation",
            unit: units.LuminanceMillicandelaPerSquareMeter,
            want: "mCd/m²",
        },
        {
            name: "CenticandelaPerSquareMeter abbreviation",
            unit: units.LuminanceCenticandelaPerSquareMeter,
            want: "cCd/m²",
        },
        {
            name: "DecicandelaPerSquareMeter abbreviation",
            unit: units.LuminanceDecicandelaPerSquareMeter,
            want: "dCd/m²",
        },
        {
            name: "KilocandelaPerSquareMeter abbreviation",
            unit: units.LuminanceKilocandelaPerSquareMeter,
            want: "kCd/m²",
        },
        {
            name: "invalid unit",
            unit: units.LuminanceUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetLuminanceAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetLuminanceAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestLuminance_String(t *testing.T) {
    factory := units.LuminanceFactory{}
    
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
            unit, err := factory.CreateLuminance(tt.value, units.LuminanceCandelaPerSquareMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Luminance.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestLuminance_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.LuminanceFactory{}

	_, err := uf.CreateLuminance(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}