// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestWarpingMomentOfInertiaDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MeterToTheSixth"}`
	
	factory := units.WarpingMomentOfInertiaDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.WarpingMomentOfInertiaMeterToTheSixth {
		t.Errorf("expected unit %v, got %v", units.WarpingMomentOfInertiaMeterToTheSixth, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MeterToTheSixth"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestWarpingMomentOfInertiaDto_ToJSON(t *testing.T) {
	dto := units.WarpingMomentOfInertiaDto{
		Value: 45,
		Unit:  units.WarpingMomentOfInertiaMeterToTheSixth,
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
	if result["unit"].(string) != string(units.WarpingMomentOfInertiaMeterToTheSixth) {
		t.Errorf("expected unit %s, got %v", units.WarpingMomentOfInertiaMeterToTheSixth, result["unit"])
	}
}

func TestNewWarpingMomentOfInertia_InvalidValue(t *testing.T) {
	factory := units.WarpingMomentOfInertiaFactory{}
	// NaN value should return an error.
	_, err := factory.CreateWarpingMomentOfInertia(math.NaN(), units.WarpingMomentOfInertiaMeterToTheSixth)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateWarpingMomentOfInertia(math.Inf(1), units.WarpingMomentOfInertiaMeterToTheSixth)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestWarpingMomentOfInertiaConversions(t *testing.T) {
	factory := units.WarpingMomentOfInertiaFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateWarpingMomentOfInertia(180, units.WarpingMomentOfInertiaMeterToTheSixth)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MetersToTheSixth.
		// No expected conversion value provided for MetersToTheSixth, verifying result is not NaN.
		result := a.MetersToTheSixth()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetersToTheSixth returned NaN")
		}
	}
	{
		// Test conversion to DecimetersToTheSixth.
		// No expected conversion value provided for DecimetersToTheSixth, verifying result is not NaN.
		result := a.DecimetersToTheSixth()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecimetersToTheSixth returned NaN")
		}
	}
	{
		// Test conversion to CentimetersToTheSixth.
		// No expected conversion value provided for CentimetersToTheSixth, verifying result is not NaN.
		result := a.CentimetersToTheSixth()
		if math.IsNaN(result) {
			t.Errorf("conversion to CentimetersToTheSixth returned NaN")
		}
	}
	{
		// Test conversion to MillimetersToTheSixth.
		// No expected conversion value provided for MillimetersToTheSixth, verifying result is not NaN.
		result := a.MillimetersToTheSixth()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillimetersToTheSixth returned NaN")
		}
	}
	{
		// Test conversion to FeetToTheSixth.
		// No expected conversion value provided for FeetToTheSixth, verifying result is not NaN.
		result := a.FeetToTheSixth()
		if math.IsNaN(result) {
			t.Errorf("conversion to FeetToTheSixth returned NaN")
		}
	}
	{
		// Test conversion to InchesToTheSixth.
		// No expected conversion value provided for InchesToTheSixth, verifying result is not NaN.
		result := a.InchesToTheSixth()
		if math.IsNaN(result) {
			t.Errorf("conversion to InchesToTheSixth returned NaN")
		}
	}
}

func TestWarpingMomentOfInertia_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.WarpingMomentOfInertiaFactory{}
	a, err := factory.CreateWarpingMomentOfInertia(90, units.WarpingMomentOfInertiaMeterToTheSixth)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.WarpingMomentOfInertiaMeterToTheSixth {
		t.Errorf("expected default unit MeterToTheSixth, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.WarpingMomentOfInertiaMeterToTheSixth
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.WarpingMomentOfInertiaDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.WarpingMomentOfInertiaMeterToTheSixth {
		t.Errorf("expected unit MeterToTheSixth, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestWarpingMomentOfInertiaFactory_FromDto(t *testing.T) {
    factory := units.WarpingMomentOfInertiaFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.WarpingMomentOfInertiaDto{
        Value: 100,
        Unit:  units.WarpingMomentOfInertiaMeterToTheSixth,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.WarpingMomentOfInertiaDto{
        Value: math.NaN(),
        Unit:  units.WarpingMomentOfInertiaMeterToTheSixth,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test MeterToTheSixth conversion
    meters_to_the_sixthDto := units.WarpingMomentOfInertiaDto{
        Value: 100,
        Unit:  units.WarpingMomentOfInertiaMeterToTheSixth,
    }
    
    var meters_to_the_sixthResult *units.WarpingMomentOfInertia
    meters_to_the_sixthResult, err = factory.FromDto(meters_to_the_sixthDto)
    if err != nil {
        t.Errorf("FromDto() with MeterToTheSixth returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_to_the_sixthResult.Convert(units.WarpingMomentOfInertiaMeterToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterToTheSixth = %v, want %v", converted, 100)
    }
    // Test DecimeterToTheSixth conversion
    decimeters_to_the_sixthDto := units.WarpingMomentOfInertiaDto{
        Value: 100,
        Unit:  units.WarpingMomentOfInertiaDecimeterToTheSixth,
    }
    
    var decimeters_to_the_sixthResult *units.WarpingMomentOfInertia
    decimeters_to_the_sixthResult, err = factory.FromDto(decimeters_to_the_sixthDto)
    if err != nil {
        t.Errorf("FromDto() with DecimeterToTheSixth returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimeters_to_the_sixthResult.Convert(units.WarpingMomentOfInertiaDecimeterToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimeterToTheSixth = %v, want %v", converted, 100)
    }
    // Test CentimeterToTheSixth conversion
    centimeters_to_the_sixthDto := units.WarpingMomentOfInertiaDto{
        Value: 100,
        Unit:  units.WarpingMomentOfInertiaCentimeterToTheSixth,
    }
    
    var centimeters_to_the_sixthResult *units.WarpingMomentOfInertia
    centimeters_to_the_sixthResult, err = factory.FromDto(centimeters_to_the_sixthDto)
    if err != nil {
        t.Errorf("FromDto() with CentimeterToTheSixth returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_to_the_sixthResult.Convert(units.WarpingMomentOfInertiaCentimeterToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterToTheSixth = %v, want %v", converted, 100)
    }
    // Test MillimeterToTheSixth conversion
    millimeters_to_the_sixthDto := units.WarpingMomentOfInertiaDto{
        Value: 100,
        Unit:  units.WarpingMomentOfInertiaMillimeterToTheSixth,
    }
    
    var millimeters_to_the_sixthResult *units.WarpingMomentOfInertia
    millimeters_to_the_sixthResult, err = factory.FromDto(millimeters_to_the_sixthDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterToTheSixth returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_to_the_sixthResult.Convert(units.WarpingMomentOfInertiaMillimeterToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterToTheSixth = %v, want %v", converted, 100)
    }
    // Test FootToTheSixth conversion
    feet_to_the_sixthDto := units.WarpingMomentOfInertiaDto{
        Value: 100,
        Unit:  units.WarpingMomentOfInertiaFootToTheSixth,
    }
    
    var feet_to_the_sixthResult *units.WarpingMomentOfInertia
    feet_to_the_sixthResult, err = factory.FromDto(feet_to_the_sixthDto)
    if err != nil {
        t.Errorf("FromDto() with FootToTheSixth returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_to_the_sixthResult.Convert(units.WarpingMomentOfInertiaFootToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootToTheSixth = %v, want %v", converted, 100)
    }
    // Test InchToTheSixth conversion
    inches_to_the_sixthDto := units.WarpingMomentOfInertiaDto{
        Value: 100,
        Unit:  units.WarpingMomentOfInertiaInchToTheSixth,
    }
    
    var inches_to_the_sixthResult *units.WarpingMomentOfInertia
    inches_to_the_sixthResult, err = factory.FromDto(inches_to_the_sixthDto)
    if err != nil {
        t.Errorf("FromDto() with InchToTheSixth returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_to_the_sixthResult.Convert(units.WarpingMomentOfInertiaInchToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchToTheSixth = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.WarpingMomentOfInertiaDto{
        Value: 0,
        Unit:  units.WarpingMomentOfInertiaMeterToTheSixth,
    }
    
    var zeroResult *units.WarpingMomentOfInertia
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestWarpingMomentOfInertiaFactory_FromDtoJSON(t *testing.T) {
    factory := units.WarpingMomentOfInertiaFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "MeterToTheSixth"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "MeterToTheSixth"}`)
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
    nanJSON, _ := json.Marshal(units.WarpingMomentOfInertiaDto{
        Value: nanValue,
        Unit:  units.WarpingMomentOfInertiaMeterToTheSixth,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with MeterToTheSixth unit
    meters_to_the_sixthJSON := []byte(`{"value": 100, "unit": "MeterToTheSixth"}`)
    meters_to_the_sixthResult, err := factory.FromDtoJSON(meters_to_the_sixthJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeterToTheSixth unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_to_the_sixthResult.Convert(units.WarpingMomentOfInertiaMeterToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterToTheSixth = %v, want %v", converted, 100)
    }
    // Test JSON with DecimeterToTheSixth unit
    decimeters_to_the_sixthJSON := []byte(`{"value": 100, "unit": "DecimeterToTheSixth"}`)
    decimeters_to_the_sixthResult, err := factory.FromDtoJSON(decimeters_to_the_sixthJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecimeterToTheSixth unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimeters_to_the_sixthResult.Convert(units.WarpingMomentOfInertiaDecimeterToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimeterToTheSixth = %v, want %v", converted, 100)
    }
    // Test JSON with CentimeterToTheSixth unit
    centimeters_to_the_sixthJSON := []byte(`{"value": 100, "unit": "CentimeterToTheSixth"}`)
    centimeters_to_the_sixthResult, err := factory.FromDtoJSON(centimeters_to_the_sixthJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentimeterToTheSixth unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_to_the_sixthResult.Convert(units.WarpingMomentOfInertiaCentimeterToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterToTheSixth = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterToTheSixth unit
    millimeters_to_the_sixthJSON := []byte(`{"value": 100, "unit": "MillimeterToTheSixth"}`)
    millimeters_to_the_sixthResult, err := factory.FromDtoJSON(millimeters_to_the_sixthJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterToTheSixth unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_to_the_sixthResult.Convert(units.WarpingMomentOfInertiaMillimeterToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterToTheSixth = %v, want %v", converted, 100)
    }
    // Test JSON with FootToTheSixth unit
    feet_to_the_sixthJSON := []byte(`{"value": 100, "unit": "FootToTheSixth"}`)
    feet_to_the_sixthResult, err := factory.FromDtoJSON(feet_to_the_sixthJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FootToTheSixth unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_to_the_sixthResult.Convert(units.WarpingMomentOfInertiaFootToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootToTheSixth = %v, want %v", converted, 100)
    }
    // Test JSON with InchToTheSixth unit
    inches_to_the_sixthJSON := []byte(`{"value": 100, "unit": "InchToTheSixth"}`)
    inches_to_the_sixthResult, err := factory.FromDtoJSON(inches_to_the_sixthJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InchToTheSixth unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_to_the_sixthResult.Convert(units.WarpingMomentOfInertiaInchToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchToTheSixth = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "MeterToTheSixth"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromMetersToTheSixth function
func TestWarpingMomentOfInertiaFactory_FromMetersToTheSixth(t *testing.T) {
    factory := units.WarpingMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMetersToTheSixth(100)
    if err != nil {
        t.Errorf("FromMetersToTheSixth() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.WarpingMomentOfInertiaMeterToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMetersToTheSixth() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMetersToTheSixth(math.NaN())
    if err == nil {
        t.Error("FromMetersToTheSixth() with NaN value should return error")
    }

    _, err = factory.FromMetersToTheSixth(math.Inf(1))
    if err == nil {
        t.Error("FromMetersToTheSixth() with +Inf value should return error")
    }

    _, err = factory.FromMetersToTheSixth(math.Inf(-1))
    if err == nil {
        t.Error("FromMetersToTheSixth() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMetersToTheSixth(0)
    if err != nil {
        t.Errorf("FromMetersToTheSixth() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.WarpingMomentOfInertiaMeterToTheSixth)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMetersToTheSixth() with zero value = %v, want 0", converted)
    }
}
// Test FromDecimetersToTheSixth function
func TestWarpingMomentOfInertiaFactory_FromDecimetersToTheSixth(t *testing.T) {
    factory := units.WarpingMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecimetersToTheSixth(100)
    if err != nil {
        t.Errorf("FromDecimetersToTheSixth() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.WarpingMomentOfInertiaDecimeterToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecimetersToTheSixth() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecimetersToTheSixth(math.NaN())
    if err == nil {
        t.Error("FromDecimetersToTheSixth() with NaN value should return error")
    }

    _, err = factory.FromDecimetersToTheSixth(math.Inf(1))
    if err == nil {
        t.Error("FromDecimetersToTheSixth() with +Inf value should return error")
    }

    _, err = factory.FromDecimetersToTheSixth(math.Inf(-1))
    if err == nil {
        t.Error("FromDecimetersToTheSixth() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecimetersToTheSixth(0)
    if err != nil {
        t.Errorf("FromDecimetersToTheSixth() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.WarpingMomentOfInertiaDecimeterToTheSixth)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecimetersToTheSixth() with zero value = %v, want 0", converted)
    }
}
// Test FromCentimetersToTheSixth function
func TestWarpingMomentOfInertiaFactory_FromCentimetersToTheSixth(t *testing.T) {
    factory := units.WarpingMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentimetersToTheSixth(100)
    if err != nil {
        t.Errorf("FromCentimetersToTheSixth() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.WarpingMomentOfInertiaCentimeterToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentimetersToTheSixth() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentimetersToTheSixth(math.NaN())
    if err == nil {
        t.Error("FromCentimetersToTheSixth() with NaN value should return error")
    }

    _, err = factory.FromCentimetersToTheSixth(math.Inf(1))
    if err == nil {
        t.Error("FromCentimetersToTheSixth() with +Inf value should return error")
    }

    _, err = factory.FromCentimetersToTheSixth(math.Inf(-1))
    if err == nil {
        t.Error("FromCentimetersToTheSixth() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentimetersToTheSixth(0)
    if err != nil {
        t.Errorf("FromCentimetersToTheSixth() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.WarpingMomentOfInertiaCentimeterToTheSixth)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentimetersToTheSixth() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimetersToTheSixth function
func TestWarpingMomentOfInertiaFactory_FromMillimetersToTheSixth(t *testing.T) {
    factory := units.WarpingMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimetersToTheSixth(100)
    if err != nil {
        t.Errorf("FromMillimetersToTheSixth() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.WarpingMomentOfInertiaMillimeterToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimetersToTheSixth() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimetersToTheSixth(math.NaN())
    if err == nil {
        t.Error("FromMillimetersToTheSixth() with NaN value should return error")
    }

    _, err = factory.FromMillimetersToTheSixth(math.Inf(1))
    if err == nil {
        t.Error("FromMillimetersToTheSixth() with +Inf value should return error")
    }

    _, err = factory.FromMillimetersToTheSixth(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimetersToTheSixth() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimetersToTheSixth(0)
    if err != nil {
        t.Errorf("FromMillimetersToTheSixth() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.WarpingMomentOfInertiaMillimeterToTheSixth)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimetersToTheSixth() with zero value = %v, want 0", converted)
    }
}
// Test FromFeetToTheSixth function
func TestWarpingMomentOfInertiaFactory_FromFeetToTheSixth(t *testing.T) {
    factory := units.WarpingMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFeetToTheSixth(100)
    if err != nil {
        t.Errorf("FromFeetToTheSixth() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.WarpingMomentOfInertiaFootToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFeetToTheSixth() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFeetToTheSixth(math.NaN())
    if err == nil {
        t.Error("FromFeetToTheSixth() with NaN value should return error")
    }

    _, err = factory.FromFeetToTheSixth(math.Inf(1))
    if err == nil {
        t.Error("FromFeetToTheSixth() with +Inf value should return error")
    }

    _, err = factory.FromFeetToTheSixth(math.Inf(-1))
    if err == nil {
        t.Error("FromFeetToTheSixth() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFeetToTheSixth(0)
    if err != nil {
        t.Errorf("FromFeetToTheSixth() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.WarpingMomentOfInertiaFootToTheSixth)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFeetToTheSixth() with zero value = %v, want 0", converted)
    }
}
// Test FromInchesToTheSixth function
func TestWarpingMomentOfInertiaFactory_FromInchesToTheSixth(t *testing.T) {
    factory := units.WarpingMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInchesToTheSixth(100)
    if err != nil {
        t.Errorf("FromInchesToTheSixth() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.WarpingMomentOfInertiaInchToTheSixth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInchesToTheSixth() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInchesToTheSixth(math.NaN())
    if err == nil {
        t.Error("FromInchesToTheSixth() with NaN value should return error")
    }

    _, err = factory.FromInchesToTheSixth(math.Inf(1))
    if err == nil {
        t.Error("FromInchesToTheSixth() with +Inf value should return error")
    }

    _, err = factory.FromInchesToTheSixth(math.Inf(-1))
    if err == nil {
        t.Error("FromInchesToTheSixth() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInchesToTheSixth(0)
    if err != nil {
        t.Errorf("FromInchesToTheSixth() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.WarpingMomentOfInertiaInchToTheSixth)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInchesToTheSixth() with zero value = %v, want 0", converted)
    }
}

func TestWarpingMomentOfInertiaToString(t *testing.T) {
	factory := units.WarpingMomentOfInertiaFactory{}
	a, err := factory.CreateWarpingMomentOfInertia(45, units.WarpingMomentOfInertiaMeterToTheSixth)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.WarpingMomentOfInertiaMeterToTheSixth, 2)
	expected := "45.00 " + units.GetWarpingMomentOfInertiaAbbreviation(units.WarpingMomentOfInertiaMeterToTheSixth)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.WarpingMomentOfInertiaMeterToTheSixth, -1)
	expected = "45 " + units.GetWarpingMomentOfInertiaAbbreviation(units.WarpingMomentOfInertiaMeterToTheSixth)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestWarpingMomentOfInertia_EqualityAndComparison(t *testing.T) {
	factory := units.WarpingMomentOfInertiaFactory{}
	a1, _ := factory.CreateWarpingMomentOfInertia(60, units.WarpingMomentOfInertiaMeterToTheSixth)
	a2, _ := factory.CreateWarpingMomentOfInertia(60, units.WarpingMomentOfInertiaMeterToTheSixth)
	a3, _ := factory.CreateWarpingMomentOfInertia(120, units.WarpingMomentOfInertiaMeterToTheSixth)

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

func TestWarpingMomentOfInertia_Arithmetic(t *testing.T) {
	factory := units.WarpingMomentOfInertiaFactory{}
	a1, _ := factory.CreateWarpingMomentOfInertia(30, units.WarpingMomentOfInertiaMeterToTheSixth)
	a2, _ := factory.CreateWarpingMomentOfInertia(45, units.WarpingMomentOfInertiaMeterToTheSixth)

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