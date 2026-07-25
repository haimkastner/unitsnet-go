// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAreaPerLengthDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SquareMeterPerMeter"}`
	
	factory := units.AreaPerLengthDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AreaPerLengthSquareMeterPerMeter {
		t.Errorf("expected unit %v, got %v", units.AreaPerLengthSquareMeterPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SquareMeterPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAreaPerLengthDto_ToJSON(t *testing.T) {
	dto := units.AreaPerLengthDto{
		Value: 45,
		Unit:  units.AreaPerLengthSquareMeterPerMeter,
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
	if result["unit"].(string) != string(units.AreaPerLengthSquareMeterPerMeter) {
		t.Errorf("expected unit %s, got %v", units.AreaPerLengthSquareMeterPerMeter, result["unit"])
	}
}

func TestNewAreaPerLength_InvalidValue(t *testing.T) {
	factory := units.AreaPerLengthFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAreaPerLength(math.NaN(), units.AreaPerLengthSquareMeterPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAreaPerLength(math.Inf(1), units.AreaPerLengthSquareMeterPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAreaPerLengthConversions(t *testing.T) {
	factory := units.AreaPerLengthFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAreaPerLength(180, units.AreaPerLengthSquareMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to SquareMetersPerMeter.
		// No expected conversion value provided for SquareMetersPerMeter, verifying result is not NaN.
		result := a.SquareMetersPerMeter()
		cacheResult := a.SquareMetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareMetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to SquareCentimetersPerMeter.
		// No expected conversion value provided for SquareCentimetersPerMeter, verifying result is not NaN.
		result := a.SquareCentimetersPerMeter()
		cacheResult := a.SquareCentimetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareCentimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to SquareMillimetersPerMeter.
		// No expected conversion value provided for SquareMillimetersPerMeter, verifying result is not NaN.
		result := a.SquareMillimetersPerMeter()
		cacheResult := a.SquareMillimetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareMillimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to SquareInchesPerFoot.
		// No expected conversion value provided for SquareInchesPerFoot, verifying result is not NaN.
		result := a.SquareInchesPerFoot()
		cacheResult := a.SquareInchesPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareInchesPerFoot returned NaN")
		}
	}
	{
		// Test conversion to SquareInchesPerInch.
		// No expected conversion value provided for SquareInchesPerInch, verifying result is not NaN.
		result := a.SquareInchesPerInch()
		cacheResult := a.SquareInchesPerInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareInchesPerInch returned NaN")
		}
	}
	{
		// Test conversion to SquareFeetPerFoot.
		// No expected conversion value provided for SquareFeetPerFoot, verifying result is not NaN.
		result := a.SquareFeetPerFoot()
		cacheResult := a.SquareFeetPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareFeetPerFoot returned NaN")
		}
	}
}

func TestAreaPerLength_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AreaPerLengthFactory{}
	a, err := factory.CreateAreaPerLength(90, units.AreaPerLengthSquareMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AreaPerLengthSquareMeterPerMeter {
		t.Errorf("expected default unit SquareMeterPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AreaPerLengthSquareMeterPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AreaPerLengthDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AreaPerLengthSquareMeterPerMeter {
		t.Errorf("expected unit SquareMeterPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAreaPerLengthFactory_FromDto(t *testing.T) {
    factory := units.AreaPerLengthFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.AreaPerLengthDto{
        Value: 100,
        Unit:  units.AreaPerLengthSquareMeterPerMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.AreaPerLengthDto{
        Value: math.NaN(),
        Unit:  units.AreaPerLengthSquareMeterPerMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test SquareMeterPerMeter conversion
    square_meters_per_meterDto := units.AreaPerLengthDto{
        Value: 100,
        Unit:  units.AreaPerLengthSquareMeterPerMeter,
    }
    
    var square_meters_per_meterResult *units.AreaPerLength
    square_meters_per_meterResult, err = factory.FromDto(square_meters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_meters_per_meterResult.Convert(units.AreaPerLengthSquareMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test SquareCentimeterPerMeter conversion
    square_centimeters_per_meterDto := units.AreaPerLengthDto{
        Value: 100,
        Unit:  units.AreaPerLengthSquareCentimeterPerMeter,
    }
    
    var square_centimeters_per_meterResult *units.AreaPerLength
    square_centimeters_per_meterResult, err = factory.FromDto(square_centimeters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with SquareCentimeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_centimeters_per_meterResult.Convert(units.AreaPerLengthSquareCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareCentimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test SquareMillimeterPerMeter conversion
    square_millimeters_per_meterDto := units.AreaPerLengthDto{
        Value: 100,
        Unit:  units.AreaPerLengthSquareMillimeterPerMeter,
    }
    
    var square_millimeters_per_meterResult *units.AreaPerLength
    square_millimeters_per_meterResult, err = factory.FromDto(square_millimeters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMillimeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_millimeters_per_meterResult.Convert(units.AreaPerLengthSquareMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMillimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test SquareInchPerFoot conversion
    square_inches_per_footDto := units.AreaPerLengthDto{
        Value: 100,
        Unit:  units.AreaPerLengthSquareInchPerFoot,
    }
    
    var square_inches_per_footResult *units.AreaPerLength
    square_inches_per_footResult, err = factory.FromDto(square_inches_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with SquareInchPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_inches_per_footResult.Convert(units.AreaPerLengthSquareInchPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareInchPerFoot = %v, want %v", converted, 100)
    }
    // Test SquareInchPerInch conversion
    square_inches_per_inchDto := units.AreaPerLengthDto{
        Value: 100,
        Unit:  units.AreaPerLengthSquareInchPerInch,
    }
    
    var square_inches_per_inchResult *units.AreaPerLength
    square_inches_per_inchResult, err = factory.FromDto(square_inches_per_inchDto)
    if err != nil {
        t.Errorf("FromDto() with SquareInchPerInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_inches_per_inchResult.Convert(units.AreaPerLengthSquareInchPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareInchPerInch = %v, want %v", converted, 100)
    }
    // Test SquareFootPerFoot conversion
    square_feet_per_footDto := units.AreaPerLengthDto{
        Value: 100,
        Unit:  units.AreaPerLengthSquareFootPerFoot,
    }
    
    var square_feet_per_footResult *units.AreaPerLength
    square_feet_per_footResult, err = factory.FromDto(square_feet_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with SquareFootPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_feet_per_footResult.Convert(units.AreaPerLengthSquareFootPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareFootPerFoot = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.AreaPerLengthDto{
        Value: 0,
        Unit:  units.AreaPerLengthSquareMeterPerMeter,
    }
    
    var zeroResult *units.AreaPerLength
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestAreaPerLengthFactory_FromDtoJSON(t *testing.T) {
    factory := units.AreaPerLengthFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "SquareMeterPerMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "SquareMeterPerMeter"}`)
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
    nanJSON, _ := json.Marshal(units.AreaPerLengthDto{
        Value: nanValue,
        Unit:  units.AreaPerLengthSquareMeterPerMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with SquareMeterPerMeter unit
    square_meters_per_meterJSON := []byte(`{"value": 100, "unit": "SquareMeterPerMeter"}`)
    square_meters_per_meterResult, err := factory.FromDtoJSON(square_meters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareMeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_meters_per_meterResult.Convert(units.AreaPerLengthSquareMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with SquareCentimeterPerMeter unit
    square_centimeters_per_meterJSON := []byte(`{"value": 100, "unit": "SquareCentimeterPerMeter"}`)
    square_centimeters_per_meterResult, err := factory.FromDtoJSON(square_centimeters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareCentimeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_centimeters_per_meterResult.Convert(units.AreaPerLengthSquareCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareCentimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with SquareMillimeterPerMeter unit
    square_millimeters_per_meterJSON := []byte(`{"value": 100, "unit": "SquareMillimeterPerMeter"}`)
    square_millimeters_per_meterResult, err := factory.FromDtoJSON(square_millimeters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareMillimeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_millimeters_per_meterResult.Convert(units.AreaPerLengthSquareMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMillimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with SquareInchPerFoot unit
    square_inches_per_footJSON := []byte(`{"value": 100, "unit": "SquareInchPerFoot"}`)
    square_inches_per_footResult, err := factory.FromDtoJSON(square_inches_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareInchPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_inches_per_footResult.Convert(units.AreaPerLengthSquareInchPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareInchPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with SquareInchPerInch unit
    square_inches_per_inchJSON := []byte(`{"value": 100, "unit": "SquareInchPerInch"}`)
    square_inches_per_inchResult, err := factory.FromDtoJSON(square_inches_per_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareInchPerInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_inches_per_inchResult.Convert(units.AreaPerLengthSquareInchPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareInchPerInch = %v, want %v", converted, 100)
    }
    // Test JSON with SquareFootPerFoot unit
    square_feet_per_footJSON := []byte(`{"value": 100, "unit": "SquareFootPerFoot"}`)
    square_feet_per_footResult, err := factory.FromDtoJSON(square_feet_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareFootPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_feet_per_footResult.Convert(units.AreaPerLengthSquareFootPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareFootPerFoot = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "SquareMeterPerMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromSquareMetersPerMeter function
func TestAreaPerLengthFactory_FromSquareMetersPerMeter(t *testing.T) {
    factory := units.AreaPerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMetersPerMeter(100)
    if err != nil {
        t.Errorf("FromSquareMetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaPerLengthSquareMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareMetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareMetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromSquareMetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromSquareMetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromSquareMetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromSquareMetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareMetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareMetersPerMeter(0)
    if err != nil {
        t.Errorf("FromSquareMetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaPerLengthSquareMeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareCentimetersPerMeter function
func TestAreaPerLengthFactory_FromSquareCentimetersPerMeter(t *testing.T) {
    factory := units.AreaPerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareCentimetersPerMeter(100)
    if err != nil {
        t.Errorf("FromSquareCentimetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaPerLengthSquareCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareCentimetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareCentimetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromSquareCentimetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromSquareCentimetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromSquareCentimetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromSquareCentimetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareCentimetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareCentimetersPerMeter(0)
    if err != nil {
        t.Errorf("FromSquareCentimetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaPerLengthSquareCentimeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareCentimetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareMillimetersPerMeter function
func TestAreaPerLengthFactory_FromSquareMillimetersPerMeter(t *testing.T) {
    factory := units.AreaPerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMillimetersPerMeter(100)
    if err != nil {
        t.Errorf("FromSquareMillimetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaPerLengthSquareMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareMillimetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareMillimetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromSquareMillimetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromSquareMillimetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromSquareMillimetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromSquareMillimetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareMillimetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareMillimetersPerMeter(0)
    if err != nil {
        t.Errorf("FromSquareMillimetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaPerLengthSquareMillimeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMillimetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareInchesPerFoot function
func TestAreaPerLengthFactory_FromSquareInchesPerFoot(t *testing.T) {
    factory := units.AreaPerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareInchesPerFoot(100)
    if err != nil {
        t.Errorf("FromSquareInchesPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaPerLengthSquareInchPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareInchesPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareInchesPerFoot(math.NaN())
    if err == nil {
        t.Error("FromSquareInchesPerFoot() with NaN value should return error")
    }

    _, err = factory.FromSquareInchesPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromSquareInchesPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromSquareInchesPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareInchesPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareInchesPerFoot(0)
    if err != nil {
        t.Errorf("FromSquareInchesPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaPerLengthSquareInchPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareInchesPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareInchesPerInch function
func TestAreaPerLengthFactory_FromSquareInchesPerInch(t *testing.T) {
    factory := units.AreaPerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareInchesPerInch(100)
    if err != nil {
        t.Errorf("FromSquareInchesPerInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaPerLengthSquareInchPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareInchesPerInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareInchesPerInch(math.NaN())
    if err == nil {
        t.Error("FromSquareInchesPerInch() with NaN value should return error")
    }

    _, err = factory.FromSquareInchesPerInch(math.Inf(1))
    if err == nil {
        t.Error("FromSquareInchesPerInch() with +Inf value should return error")
    }

    _, err = factory.FromSquareInchesPerInch(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareInchesPerInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareInchesPerInch(0)
    if err != nil {
        t.Errorf("FromSquareInchesPerInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaPerLengthSquareInchPerInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareInchesPerInch() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareFeetPerFoot function
func TestAreaPerLengthFactory_FromSquareFeetPerFoot(t *testing.T) {
    factory := units.AreaPerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareFeetPerFoot(100)
    if err != nil {
        t.Errorf("FromSquareFeetPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaPerLengthSquareFootPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareFeetPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareFeetPerFoot(math.NaN())
    if err == nil {
        t.Error("FromSquareFeetPerFoot() with NaN value should return error")
    }

    _, err = factory.FromSquareFeetPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromSquareFeetPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromSquareFeetPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareFeetPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareFeetPerFoot(0)
    if err != nil {
        t.Errorf("FromSquareFeetPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaPerLengthSquareFootPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareFeetPerFoot() with zero value = %v, want 0", converted)
    }
}

func TestAreaPerLengthToString(t *testing.T) {
	factory := units.AreaPerLengthFactory{}
	a, err := factory.CreateAreaPerLength(45, units.AreaPerLengthSquareMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AreaPerLengthSquareMeterPerMeter, 2)
	expected := "45.00 " + units.GetAreaPerLengthAbbreviation(units.AreaPerLengthSquareMeterPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AreaPerLengthSquareMeterPerMeter, -1)
	expected = "45 " + units.GetAreaPerLengthAbbreviation(units.AreaPerLengthSquareMeterPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAreaPerLength_EqualityAndComparison(t *testing.T) {
	factory := units.AreaPerLengthFactory{}
	a1, _ := factory.CreateAreaPerLength(60, units.AreaPerLengthSquareMeterPerMeter)
	a2, _ := factory.CreateAreaPerLength(60, units.AreaPerLengthSquareMeterPerMeter)
	a3, _ := factory.CreateAreaPerLength(120, units.AreaPerLengthSquareMeterPerMeter)

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

func TestAreaPerLength_Arithmetic(t *testing.T) {
	factory := units.AreaPerLengthFactory{}
	a1, _ := factory.CreateAreaPerLength(30, units.AreaPerLengthSquareMeterPerMeter)
	a2, _ := factory.CreateAreaPerLength(45, units.AreaPerLengthSquareMeterPerMeter)

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


func TestGetAreaPerLengthAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.AreaPerLengthUnits
        want string
    }{
        {
            name: "SquareMeterPerMeter abbreviation",
            unit: units.AreaPerLengthSquareMeterPerMeter,
            want: "m²/m",
        },
        {
            name: "SquareCentimeterPerMeter abbreviation",
            unit: units.AreaPerLengthSquareCentimeterPerMeter,
            want: "cm²/m",
        },
        {
            name: "SquareMillimeterPerMeter abbreviation",
            unit: units.AreaPerLengthSquareMillimeterPerMeter,
            want: "mm²/m",
        },
        {
            name: "SquareInchPerFoot abbreviation",
            unit: units.AreaPerLengthSquareInchPerFoot,
            want: "in²/ft",
        },
        {
            name: "SquareInchPerInch abbreviation",
            unit: units.AreaPerLengthSquareInchPerInch,
            want: "in²/in",
        },
        {
            name: "SquareFootPerFoot abbreviation",
            unit: units.AreaPerLengthSquareFootPerFoot,
            want: "ft²/ft",
        },
        {
            name: "invalid unit",
            unit: units.AreaPerLengthUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetAreaPerLengthAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetAreaPerLengthAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestAreaPerLength_String(t *testing.T) {
    factory := units.AreaPerLengthFactory{}
    
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
            unit, err := factory.CreateAreaPerLength(tt.value, units.AreaPerLengthSquareMeterPerMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("AreaPerLength.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestAreaPerLength_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.AreaPerLengthFactory{}

	_, err := uf.CreateAreaPerLength(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}