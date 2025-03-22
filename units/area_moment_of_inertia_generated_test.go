// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAreaMomentOfInertiaDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MeterToTheFourth"}`
	
	factory := units.AreaMomentOfInertiaDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AreaMomentOfInertiaMeterToTheFourth {
		t.Errorf("expected unit %v, got %v", units.AreaMomentOfInertiaMeterToTheFourth, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MeterToTheFourth"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAreaMomentOfInertiaDto_ToJSON(t *testing.T) {
	dto := units.AreaMomentOfInertiaDto{
		Value: 45,
		Unit:  units.AreaMomentOfInertiaMeterToTheFourth,
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
	if result["unit"].(string) != string(units.AreaMomentOfInertiaMeterToTheFourth) {
		t.Errorf("expected unit %s, got %v", units.AreaMomentOfInertiaMeterToTheFourth, result["unit"])
	}
}

func TestNewAreaMomentOfInertia_InvalidValue(t *testing.T) {
	factory := units.AreaMomentOfInertiaFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAreaMomentOfInertia(math.NaN(), units.AreaMomentOfInertiaMeterToTheFourth)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAreaMomentOfInertia(math.Inf(1), units.AreaMomentOfInertiaMeterToTheFourth)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAreaMomentOfInertiaConversions(t *testing.T) {
	factory := units.AreaMomentOfInertiaFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAreaMomentOfInertia(180, units.AreaMomentOfInertiaMeterToTheFourth)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MetersToTheFourth.
		// No expected conversion value provided for MetersToTheFourth, verifying result is not NaN.
		result := a.MetersToTheFourth()
		cacheResult := a.MetersToTheFourth()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MetersToTheFourth returned NaN")
		}
	}
	{
		// Test conversion to DecimetersToTheFourth.
		// No expected conversion value provided for DecimetersToTheFourth, verifying result is not NaN.
		result := a.DecimetersToTheFourth()
		cacheResult := a.DecimetersToTheFourth()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecimetersToTheFourth returned NaN")
		}
	}
	{
		// Test conversion to CentimetersToTheFourth.
		// No expected conversion value provided for CentimetersToTheFourth, verifying result is not NaN.
		result := a.CentimetersToTheFourth()
		cacheResult := a.CentimetersToTheFourth()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentimetersToTheFourth returned NaN")
		}
	}
	{
		// Test conversion to MillimetersToTheFourth.
		// No expected conversion value provided for MillimetersToTheFourth, verifying result is not NaN.
		result := a.MillimetersToTheFourth()
		cacheResult := a.MillimetersToTheFourth()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimetersToTheFourth returned NaN")
		}
	}
	{
		// Test conversion to FeetToTheFourth.
		// No expected conversion value provided for FeetToTheFourth, verifying result is not NaN.
		result := a.FeetToTheFourth()
		cacheResult := a.FeetToTheFourth()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to FeetToTheFourth returned NaN")
		}
	}
	{
		// Test conversion to InchesToTheFourth.
		// No expected conversion value provided for InchesToTheFourth, verifying result is not NaN.
		result := a.InchesToTheFourth()
		cacheResult := a.InchesToTheFourth()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InchesToTheFourth returned NaN")
		}
	}
}

func TestAreaMomentOfInertia_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AreaMomentOfInertiaFactory{}
	a, err := factory.CreateAreaMomentOfInertia(90, units.AreaMomentOfInertiaMeterToTheFourth)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AreaMomentOfInertiaMeterToTheFourth {
		t.Errorf("expected default unit MeterToTheFourth, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AreaMomentOfInertiaMeterToTheFourth
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AreaMomentOfInertiaDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AreaMomentOfInertiaMeterToTheFourth {
		t.Errorf("expected unit MeterToTheFourth, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAreaMomentOfInertiaFactory_FromDto(t *testing.T) {
    factory := units.AreaMomentOfInertiaFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.AreaMomentOfInertiaDto{
        Value: 100,
        Unit:  units.AreaMomentOfInertiaMeterToTheFourth,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.AreaMomentOfInertiaDto{
        Value: math.NaN(),
        Unit:  units.AreaMomentOfInertiaMeterToTheFourth,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test MeterToTheFourth conversion
    meters_to_the_fourthDto := units.AreaMomentOfInertiaDto{
        Value: 100,
        Unit:  units.AreaMomentOfInertiaMeterToTheFourth,
    }
    
    var meters_to_the_fourthResult *units.AreaMomentOfInertia
    meters_to_the_fourthResult, err = factory.FromDto(meters_to_the_fourthDto)
    if err != nil {
        t.Errorf("FromDto() with MeterToTheFourth returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_to_the_fourthResult.Convert(units.AreaMomentOfInertiaMeterToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterToTheFourth = %v, want %v", converted, 100)
    }
    // Test DecimeterToTheFourth conversion
    decimeters_to_the_fourthDto := units.AreaMomentOfInertiaDto{
        Value: 100,
        Unit:  units.AreaMomentOfInertiaDecimeterToTheFourth,
    }
    
    var decimeters_to_the_fourthResult *units.AreaMomentOfInertia
    decimeters_to_the_fourthResult, err = factory.FromDto(decimeters_to_the_fourthDto)
    if err != nil {
        t.Errorf("FromDto() with DecimeterToTheFourth returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimeters_to_the_fourthResult.Convert(units.AreaMomentOfInertiaDecimeterToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimeterToTheFourth = %v, want %v", converted, 100)
    }
    // Test CentimeterToTheFourth conversion
    centimeters_to_the_fourthDto := units.AreaMomentOfInertiaDto{
        Value: 100,
        Unit:  units.AreaMomentOfInertiaCentimeterToTheFourth,
    }
    
    var centimeters_to_the_fourthResult *units.AreaMomentOfInertia
    centimeters_to_the_fourthResult, err = factory.FromDto(centimeters_to_the_fourthDto)
    if err != nil {
        t.Errorf("FromDto() with CentimeterToTheFourth returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_to_the_fourthResult.Convert(units.AreaMomentOfInertiaCentimeterToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterToTheFourth = %v, want %v", converted, 100)
    }
    // Test MillimeterToTheFourth conversion
    millimeters_to_the_fourthDto := units.AreaMomentOfInertiaDto{
        Value: 100,
        Unit:  units.AreaMomentOfInertiaMillimeterToTheFourth,
    }
    
    var millimeters_to_the_fourthResult *units.AreaMomentOfInertia
    millimeters_to_the_fourthResult, err = factory.FromDto(millimeters_to_the_fourthDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterToTheFourth returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_to_the_fourthResult.Convert(units.AreaMomentOfInertiaMillimeterToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterToTheFourth = %v, want %v", converted, 100)
    }
    // Test FootToTheFourth conversion
    feet_to_the_fourthDto := units.AreaMomentOfInertiaDto{
        Value: 100,
        Unit:  units.AreaMomentOfInertiaFootToTheFourth,
    }
    
    var feet_to_the_fourthResult *units.AreaMomentOfInertia
    feet_to_the_fourthResult, err = factory.FromDto(feet_to_the_fourthDto)
    if err != nil {
        t.Errorf("FromDto() with FootToTheFourth returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_to_the_fourthResult.Convert(units.AreaMomentOfInertiaFootToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootToTheFourth = %v, want %v", converted, 100)
    }
    // Test InchToTheFourth conversion
    inches_to_the_fourthDto := units.AreaMomentOfInertiaDto{
        Value: 100,
        Unit:  units.AreaMomentOfInertiaInchToTheFourth,
    }
    
    var inches_to_the_fourthResult *units.AreaMomentOfInertia
    inches_to_the_fourthResult, err = factory.FromDto(inches_to_the_fourthDto)
    if err != nil {
        t.Errorf("FromDto() with InchToTheFourth returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_to_the_fourthResult.Convert(units.AreaMomentOfInertiaInchToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchToTheFourth = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.AreaMomentOfInertiaDto{
        Value: 0,
        Unit:  units.AreaMomentOfInertiaMeterToTheFourth,
    }
    
    var zeroResult *units.AreaMomentOfInertia
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestAreaMomentOfInertiaFactory_FromDtoJSON(t *testing.T) {
    factory := units.AreaMomentOfInertiaFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "MeterToTheFourth"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "MeterToTheFourth"}`)
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
    nanJSON, _ := json.Marshal(units.AreaMomentOfInertiaDto{
        Value: nanValue,
        Unit:  units.AreaMomentOfInertiaMeterToTheFourth,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with MeterToTheFourth unit
    meters_to_the_fourthJSON := []byte(`{"value": 100, "unit": "MeterToTheFourth"}`)
    meters_to_the_fourthResult, err := factory.FromDtoJSON(meters_to_the_fourthJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeterToTheFourth unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_to_the_fourthResult.Convert(units.AreaMomentOfInertiaMeterToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterToTheFourth = %v, want %v", converted, 100)
    }
    // Test JSON with DecimeterToTheFourth unit
    decimeters_to_the_fourthJSON := []byte(`{"value": 100, "unit": "DecimeterToTheFourth"}`)
    decimeters_to_the_fourthResult, err := factory.FromDtoJSON(decimeters_to_the_fourthJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecimeterToTheFourth unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimeters_to_the_fourthResult.Convert(units.AreaMomentOfInertiaDecimeterToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimeterToTheFourth = %v, want %v", converted, 100)
    }
    // Test JSON with CentimeterToTheFourth unit
    centimeters_to_the_fourthJSON := []byte(`{"value": 100, "unit": "CentimeterToTheFourth"}`)
    centimeters_to_the_fourthResult, err := factory.FromDtoJSON(centimeters_to_the_fourthJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentimeterToTheFourth unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_to_the_fourthResult.Convert(units.AreaMomentOfInertiaCentimeterToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterToTheFourth = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterToTheFourth unit
    millimeters_to_the_fourthJSON := []byte(`{"value": 100, "unit": "MillimeterToTheFourth"}`)
    millimeters_to_the_fourthResult, err := factory.FromDtoJSON(millimeters_to_the_fourthJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterToTheFourth unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_to_the_fourthResult.Convert(units.AreaMomentOfInertiaMillimeterToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterToTheFourth = %v, want %v", converted, 100)
    }
    // Test JSON with FootToTheFourth unit
    feet_to_the_fourthJSON := []byte(`{"value": 100, "unit": "FootToTheFourth"}`)
    feet_to_the_fourthResult, err := factory.FromDtoJSON(feet_to_the_fourthJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FootToTheFourth unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_to_the_fourthResult.Convert(units.AreaMomentOfInertiaFootToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootToTheFourth = %v, want %v", converted, 100)
    }
    // Test JSON with InchToTheFourth unit
    inches_to_the_fourthJSON := []byte(`{"value": 100, "unit": "InchToTheFourth"}`)
    inches_to_the_fourthResult, err := factory.FromDtoJSON(inches_to_the_fourthJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InchToTheFourth unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_to_the_fourthResult.Convert(units.AreaMomentOfInertiaInchToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchToTheFourth = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "MeterToTheFourth"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromMetersToTheFourth function
func TestAreaMomentOfInertiaFactory_FromMetersToTheFourth(t *testing.T) {
    factory := units.AreaMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMetersToTheFourth(100)
    if err != nil {
        t.Errorf("FromMetersToTheFourth() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaMomentOfInertiaMeterToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMetersToTheFourth() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMetersToTheFourth(math.NaN())
    if err == nil {
        t.Error("FromMetersToTheFourth() with NaN value should return error")
    }

    _, err = factory.FromMetersToTheFourth(math.Inf(1))
    if err == nil {
        t.Error("FromMetersToTheFourth() with +Inf value should return error")
    }

    _, err = factory.FromMetersToTheFourth(math.Inf(-1))
    if err == nil {
        t.Error("FromMetersToTheFourth() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMetersToTheFourth(0)
    if err != nil {
        t.Errorf("FromMetersToTheFourth() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaMomentOfInertiaMeterToTheFourth)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMetersToTheFourth() with zero value = %v, want 0", converted)
    }
}
// Test FromDecimetersToTheFourth function
func TestAreaMomentOfInertiaFactory_FromDecimetersToTheFourth(t *testing.T) {
    factory := units.AreaMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecimetersToTheFourth(100)
    if err != nil {
        t.Errorf("FromDecimetersToTheFourth() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaMomentOfInertiaDecimeterToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecimetersToTheFourth() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecimetersToTheFourth(math.NaN())
    if err == nil {
        t.Error("FromDecimetersToTheFourth() with NaN value should return error")
    }

    _, err = factory.FromDecimetersToTheFourth(math.Inf(1))
    if err == nil {
        t.Error("FromDecimetersToTheFourth() with +Inf value should return error")
    }

    _, err = factory.FromDecimetersToTheFourth(math.Inf(-1))
    if err == nil {
        t.Error("FromDecimetersToTheFourth() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecimetersToTheFourth(0)
    if err != nil {
        t.Errorf("FromDecimetersToTheFourth() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaMomentOfInertiaDecimeterToTheFourth)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecimetersToTheFourth() with zero value = %v, want 0", converted)
    }
}
// Test FromCentimetersToTheFourth function
func TestAreaMomentOfInertiaFactory_FromCentimetersToTheFourth(t *testing.T) {
    factory := units.AreaMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentimetersToTheFourth(100)
    if err != nil {
        t.Errorf("FromCentimetersToTheFourth() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaMomentOfInertiaCentimeterToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentimetersToTheFourth() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentimetersToTheFourth(math.NaN())
    if err == nil {
        t.Error("FromCentimetersToTheFourth() with NaN value should return error")
    }

    _, err = factory.FromCentimetersToTheFourth(math.Inf(1))
    if err == nil {
        t.Error("FromCentimetersToTheFourth() with +Inf value should return error")
    }

    _, err = factory.FromCentimetersToTheFourth(math.Inf(-1))
    if err == nil {
        t.Error("FromCentimetersToTheFourth() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentimetersToTheFourth(0)
    if err != nil {
        t.Errorf("FromCentimetersToTheFourth() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaMomentOfInertiaCentimeterToTheFourth)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentimetersToTheFourth() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimetersToTheFourth function
func TestAreaMomentOfInertiaFactory_FromMillimetersToTheFourth(t *testing.T) {
    factory := units.AreaMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimetersToTheFourth(100)
    if err != nil {
        t.Errorf("FromMillimetersToTheFourth() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaMomentOfInertiaMillimeterToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimetersToTheFourth() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimetersToTheFourth(math.NaN())
    if err == nil {
        t.Error("FromMillimetersToTheFourth() with NaN value should return error")
    }

    _, err = factory.FromMillimetersToTheFourth(math.Inf(1))
    if err == nil {
        t.Error("FromMillimetersToTheFourth() with +Inf value should return error")
    }

    _, err = factory.FromMillimetersToTheFourth(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimetersToTheFourth() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimetersToTheFourth(0)
    if err != nil {
        t.Errorf("FromMillimetersToTheFourth() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaMomentOfInertiaMillimeterToTheFourth)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimetersToTheFourth() with zero value = %v, want 0", converted)
    }
}
// Test FromFeetToTheFourth function
func TestAreaMomentOfInertiaFactory_FromFeetToTheFourth(t *testing.T) {
    factory := units.AreaMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFeetToTheFourth(100)
    if err != nil {
        t.Errorf("FromFeetToTheFourth() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaMomentOfInertiaFootToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFeetToTheFourth() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFeetToTheFourth(math.NaN())
    if err == nil {
        t.Error("FromFeetToTheFourth() with NaN value should return error")
    }

    _, err = factory.FromFeetToTheFourth(math.Inf(1))
    if err == nil {
        t.Error("FromFeetToTheFourth() with +Inf value should return error")
    }

    _, err = factory.FromFeetToTheFourth(math.Inf(-1))
    if err == nil {
        t.Error("FromFeetToTheFourth() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFeetToTheFourth(0)
    if err != nil {
        t.Errorf("FromFeetToTheFourth() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaMomentOfInertiaFootToTheFourth)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFeetToTheFourth() with zero value = %v, want 0", converted)
    }
}
// Test FromInchesToTheFourth function
func TestAreaMomentOfInertiaFactory_FromInchesToTheFourth(t *testing.T) {
    factory := units.AreaMomentOfInertiaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInchesToTheFourth(100)
    if err != nil {
        t.Errorf("FromInchesToTheFourth() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaMomentOfInertiaInchToTheFourth)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInchesToTheFourth() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInchesToTheFourth(math.NaN())
    if err == nil {
        t.Error("FromInchesToTheFourth() with NaN value should return error")
    }

    _, err = factory.FromInchesToTheFourth(math.Inf(1))
    if err == nil {
        t.Error("FromInchesToTheFourth() with +Inf value should return error")
    }

    _, err = factory.FromInchesToTheFourth(math.Inf(-1))
    if err == nil {
        t.Error("FromInchesToTheFourth() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInchesToTheFourth(0)
    if err != nil {
        t.Errorf("FromInchesToTheFourth() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaMomentOfInertiaInchToTheFourth)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInchesToTheFourth() with zero value = %v, want 0", converted)
    }
}

func TestAreaMomentOfInertiaToString(t *testing.T) {
	factory := units.AreaMomentOfInertiaFactory{}
	a, err := factory.CreateAreaMomentOfInertia(45, units.AreaMomentOfInertiaMeterToTheFourth)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AreaMomentOfInertiaMeterToTheFourth, 2)
	expected := "45.00 " + units.GetAreaMomentOfInertiaAbbreviation(units.AreaMomentOfInertiaMeterToTheFourth)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AreaMomentOfInertiaMeterToTheFourth, -1)
	expected = "45 " + units.GetAreaMomentOfInertiaAbbreviation(units.AreaMomentOfInertiaMeterToTheFourth)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAreaMomentOfInertia_EqualityAndComparison(t *testing.T) {
	factory := units.AreaMomentOfInertiaFactory{}
	a1, _ := factory.CreateAreaMomentOfInertia(60, units.AreaMomentOfInertiaMeterToTheFourth)
	a2, _ := factory.CreateAreaMomentOfInertia(60, units.AreaMomentOfInertiaMeterToTheFourth)
	a3, _ := factory.CreateAreaMomentOfInertia(120, units.AreaMomentOfInertiaMeterToTheFourth)

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

func TestAreaMomentOfInertia_Arithmetic(t *testing.T) {
	factory := units.AreaMomentOfInertiaFactory{}
	a1, _ := factory.CreateAreaMomentOfInertia(30, units.AreaMomentOfInertiaMeterToTheFourth)
	a2, _ := factory.CreateAreaMomentOfInertia(45, units.AreaMomentOfInertiaMeterToTheFourth)

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


func TestGetAreaMomentOfInertiaAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.AreaMomentOfInertiaUnits
        want string
    }{
        {
            name: "MeterToTheFourth abbreviation",
            unit: units.AreaMomentOfInertiaMeterToTheFourth,
            want: "m⁴",
        },
        {
            name: "DecimeterToTheFourth abbreviation",
            unit: units.AreaMomentOfInertiaDecimeterToTheFourth,
            want: "dm⁴",
        },
        {
            name: "CentimeterToTheFourth abbreviation",
            unit: units.AreaMomentOfInertiaCentimeterToTheFourth,
            want: "cm⁴",
        },
        {
            name: "MillimeterToTheFourth abbreviation",
            unit: units.AreaMomentOfInertiaMillimeterToTheFourth,
            want: "mm⁴",
        },
        {
            name: "FootToTheFourth abbreviation",
            unit: units.AreaMomentOfInertiaFootToTheFourth,
            want: "ft⁴",
        },
        {
            name: "InchToTheFourth abbreviation",
            unit: units.AreaMomentOfInertiaInchToTheFourth,
            want: "in⁴",
        },
        {
            name: "invalid unit",
            unit: units.AreaMomentOfInertiaUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetAreaMomentOfInertiaAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetAreaMomentOfInertiaAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestAreaMomentOfInertia_String(t *testing.T) {
    factory := units.AreaMomentOfInertiaFactory{}
    
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
            unit, err := factory.CreateAreaMomentOfInertia(tt.value, units.AreaMomentOfInertiaMeterToTheFourth)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("AreaMomentOfInertia.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestAreaMomentOfInertia_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.AreaMomentOfInertiaFactory{}

	_, err := uf.CreateAreaMomentOfInertia(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}