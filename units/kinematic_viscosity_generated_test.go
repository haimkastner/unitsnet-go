// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestKinematicViscosityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SquareMeterPerSecond"}`
	
	factory := units.KinematicViscosityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.KinematicViscositySquareMeterPerSecond {
		t.Errorf("expected unit %v, got %v", units.KinematicViscositySquareMeterPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SquareMeterPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestKinematicViscosityDto_ToJSON(t *testing.T) {
	dto := units.KinematicViscosityDto{
		Value: 45,
		Unit:  units.KinematicViscositySquareMeterPerSecond,
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
	if result["unit"].(string) != string(units.KinematicViscositySquareMeterPerSecond) {
		t.Errorf("expected unit %s, got %v", units.KinematicViscositySquareMeterPerSecond, result["unit"])
	}
}

func TestNewKinematicViscosity_InvalidValue(t *testing.T) {
	factory := units.KinematicViscosityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateKinematicViscosity(math.NaN(), units.KinematicViscositySquareMeterPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateKinematicViscosity(math.Inf(1), units.KinematicViscositySquareMeterPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestKinematicViscosityConversions(t *testing.T) {
	factory := units.KinematicViscosityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateKinematicViscosity(180, units.KinematicViscositySquareMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to SquareMetersPerSecond.
		// No expected conversion value provided for SquareMetersPerSecond, verifying result is not NaN.
		result := a.SquareMetersPerSecond()
		cacheResult := a.SquareMetersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareMetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to Stokes.
		// No expected conversion value provided for Stokes, verifying result is not NaN.
		result := a.Stokes()
		cacheResult := a.Stokes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Stokes returned NaN")
		}
	}
	{
		// Test conversion to SquareFeetPerSecond.
		// No expected conversion value provided for SquareFeetPerSecond, verifying result is not NaN.
		result := a.SquareFeetPerSecond()
		cacheResult := a.SquareFeetPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareFeetPerSecond returned NaN")
		}
	}
	{
		// Test conversion to Nanostokes.
		// No expected conversion value provided for Nanostokes, verifying result is not NaN.
		result := a.Nanostokes()
		cacheResult := a.Nanostokes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanostokes returned NaN")
		}
	}
	{
		// Test conversion to Microstokes.
		// No expected conversion value provided for Microstokes, verifying result is not NaN.
		result := a.Microstokes()
		cacheResult := a.Microstokes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microstokes returned NaN")
		}
	}
	{
		// Test conversion to Millistokes.
		// No expected conversion value provided for Millistokes, verifying result is not NaN.
		result := a.Millistokes()
		cacheResult := a.Millistokes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millistokes returned NaN")
		}
	}
	{
		// Test conversion to Centistokes.
		// No expected conversion value provided for Centistokes, verifying result is not NaN.
		result := a.Centistokes()
		cacheResult := a.Centistokes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Centistokes returned NaN")
		}
	}
	{
		// Test conversion to Decistokes.
		// No expected conversion value provided for Decistokes, verifying result is not NaN.
		result := a.Decistokes()
		cacheResult := a.Decistokes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Decistokes returned NaN")
		}
	}
	{
		// Test conversion to Kilostokes.
		// No expected conversion value provided for Kilostokes, verifying result is not NaN.
		result := a.Kilostokes()
		cacheResult := a.Kilostokes()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilostokes returned NaN")
		}
	}
}

func TestKinematicViscosity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.KinematicViscosityFactory{}
	a, err := factory.CreateKinematicViscosity(90, units.KinematicViscositySquareMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.KinematicViscositySquareMeterPerSecond {
		t.Errorf("expected default unit SquareMeterPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.KinematicViscositySquareMeterPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.KinematicViscosityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.KinematicViscositySquareMeterPerSecond {
		t.Errorf("expected unit SquareMeterPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestKinematicViscosityFactory_FromDto(t *testing.T) {
    factory := units.KinematicViscosityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.KinematicViscosityDto{
        Value: 100,
        Unit:  units.KinematicViscositySquareMeterPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.KinematicViscosityDto{
        Value: math.NaN(),
        Unit:  units.KinematicViscositySquareMeterPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test SquareMeterPerSecond conversion
    square_meters_per_secondDto := units.KinematicViscosityDto{
        Value: 100,
        Unit:  units.KinematicViscositySquareMeterPerSecond,
    }
    
    var square_meters_per_secondResult *units.KinematicViscosity
    square_meters_per_secondResult, err = factory.FromDto(square_meters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMeterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_meters_per_secondResult.Convert(units.KinematicViscositySquareMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeterPerSecond = %v, want %v", converted, 100)
    }
    // Test Stokes conversion
    stokesDto := units.KinematicViscosityDto{
        Value: 100,
        Unit:  units.KinematicViscosityStokes,
    }
    
    var stokesResult *units.KinematicViscosity
    stokesResult, err = factory.FromDto(stokesDto)
    if err != nil {
        t.Errorf("FromDto() with Stokes returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = stokesResult.Convert(units.KinematicViscosityStokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Stokes = %v, want %v", converted, 100)
    }
    // Test SquareFootPerSecond conversion
    square_feet_per_secondDto := units.KinematicViscosityDto{
        Value: 100,
        Unit:  units.KinematicViscositySquareFootPerSecond,
    }
    
    var square_feet_per_secondResult *units.KinematicViscosity
    square_feet_per_secondResult, err = factory.FromDto(square_feet_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with SquareFootPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_feet_per_secondResult.Convert(units.KinematicViscositySquareFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareFootPerSecond = %v, want %v", converted, 100)
    }
    // Test Nanostokes conversion
    nanostokesDto := units.KinematicViscosityDto{
        Value: 100,
        Unit:  units.KinematicViscosityNanostokes,
    }
    
    var nanostokesResult *units.KinematicViscosity
    nanostokesResult, err = factory.FromDto(nanostokesDto)
    if err != nil {
        t.Errorf("FromDto() with Nanostokes returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanostokesResult.Convert(units.KinematicViscosityNanostokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanostokes = %v, want %v", converted, 100)
    }
    // Test Microstokes conversion
    microstokesDto := units.KinematicViscosityDto{
        Value: 100,
        Unit:  units.KinematicViscosityMicrostokes,
    }
    
    var microstokesResult *units.KinematicViscosity
    microstokesResult, err = factory.FromDto(microstokesDto)
    if err != nil {
        t.Errorf("FromDto() with Microstokes returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microstokesResult.Convert(units.KinematicViscosityMicrostokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microstokes = %v, want %v", converted, 100)
    }
    // Test Millistokes conversion
    millistokesDto := units.KinematicViscosityDto{
        Value: 100,
        Unit:  units.KinematicViscosityMillistokes,
    }
    
    var millistokesResult *units.KinematicViscosity
    millistokesResult, err = factory.FromDto(millistokesDto)
    if err != nil {
        t.Errorf("FromDto() with Millistokes returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millistokesResult.Convert(units.KinematicViscosityMillistokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millistokes = %v, want %v", converted, 100)
    }
    // Test Centistokes conversion
    centistokesDto := units.KinematicViscosityDto{
        Value: 100,
        Unit:  units.KinematicViscosityCentistokes,
    }
    
    var centistokesResult *units.KinematicViscosity
    centistokesResult, err = factory.FromDto(centistokesDto)
    if err != nil {
        t.Errorf("FromDto() with Centistokes returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centistokesResult.Convert(units.KinematicViscosityCentistokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centistokes = %v, want %v", converted, 100)
    }
    // Test Decistokes conversion
    decistokesDto := units.KinematicViscosityDto{
        Value: 100,
        Unit:  units.KinematicViscosityDecistokes,
    }
    
    var decistokesResult *units.KinematicViscosity
    decistokesResult, err = factory.FromDto(decistokesDto)
    if err != nil {
        t.Errorf("FromDto() with Decistokes returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decistokesResult.Convert(units.KinematicViscosityDecistokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decistokes = %v, want %v", converted, 100)
    }
    // Test Kilostokes conversion
    kilostokesDto := units.KinematicViscosityDto{
        Value: 100,
        Unit:  units.KinematicViscosityKilostokes,
    }
    
    var kilostokesResult *units.KinematicViscosity
    kilostokesResult, err = factory.FromDto(kilostokesDto)
    if err != nil {
        t.Errorf("FromDto() with Kilostokes returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilostokesResult.Convert(units.KinematicViscosityKilostokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilostokes = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.KinematicViscosityDto{
        Value: 0,
        Unit:  units.KinematicViscositySquareMeterPerSecond,
    }
    
    var zeroResult *units.KinematicViscosity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestKinematicViscosityFactory_FromDtoJSON(t *testing.T) {
    factory := units.KinematicViscosityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "SquareMeterPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "SquareMeterPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.KinematicViscosityDto{
        Value: nanValue,
        Unit:  units.KinematicViscositySquareMeterPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with SquareMeterPerSecond unit
    square_meters_per_secondJSON := []byte(`{"value": 100, "unit": "SquareMeterPerSecond"}`)
    square_meters_per_secondResult, err := factory.FromDtoJSON(square_meters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareMeterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_meters_per_secondResult.Convert(units.KinematicViscositySquareMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with Stokes unit
    stokesJSON := []byte(`{"value": 100, "unit": "Stokes"}`)
    stokesResult, err := factory.FromDtoJSON(stokesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Stokes unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = stokesResult.Convert(units.KinematicViscosityStokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Stokes = %v, want %v", converted, 100)
    }
    // Test JSON with SquareFootPerSecond unit
    square_feet_per_secondJSON := []byte(`{"value": 100, "unit": "SquareFootPerSecond"}`)
    square_feet_per_secondResult, err := factory.FromDtoJSON(square_feet_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareFootPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_feet_per_secondResult.Convert(units.KinematicViscositySquareFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareFootPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with Nanostokes unit
    nanostokesJSON := []byte(`{"value": 100, "unit": "Nanostokes"}`)
    nanostokesResult, err := factory.FromDtoJSON(nanostokesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanostokes unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanostokesResult.Convert(units.KinematicViscosityNanostokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanostokes = %v, want %v", converted, 100)
    }
    // Test JSON with Microstokes unit
    microstokesJSON := []byte(`{"value": 100, "unit": "Microstokes"}`)
    microstokesResult, err := factory.FromDtoJSON(microstokesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microstokes unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microstokesResult.Convert(units.KinematicViscosityMicrostokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microstokes = %v, want %v", converted, 100)
    }
    // Test JSON with Millistokes unit
    millistokesJSON := []byte(`{"value": 100, "unit": "Millistokes"}`)
    millistokesResult, err := factory.FromDtoJSON(millistokesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millistokes unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millistokesResult.Convert(units.KinematicViscosityMillistokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millistokes = %v, want %v", converted, 100)
    }
    // Test JSON with Centistokes unit
    centistokesJSON := []byte(`{"value": 100, "unit": "Centistokes"}`)
    centistokesResult, err := factory.FromDtoJSON(centistokesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Centistokes unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centistokesResult.Convert(units.KinematicViscosityCentistokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centistokes = %v, want %v", converted, 100)
    }
    // Test JSON with Decistokes unit
    decistokesJSON := []byte(`{"value": 100, "unit": "Decistokes"}`)
    decistokesResult, err := factory.FromDtoJSON(decistokesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Decistokes unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decistokesResult.Convert(units.KinematicViscosityDecistokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decistokes = %v, want %v", converted, 100)
    }
    // Test JSON with Kilostokes unit
    kilostokesJSON := []byte(`{"value": 100, "unit": "Kilostokes"}`)
    kilostokesResult, err := factory.FromDtoJSON(kilostokesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilostokes unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilostokesResult.Convert(units.KinematicViscosityKilostokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilostokes = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "SquareMeterPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromSquareMetersPerSecond function
func TestKinematicViscosityFactory_FromSquareMetersPerSecond(t *testing.T) {
    factory := units.KinematicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMetersPerSecond(100)
    if err != nil {
        t.Errorf("FromSquareMetersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.KinematicViscositySquareMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareMetersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareMetersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromSquareMetersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromSquareMetersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromSquareMetersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromSquareMetersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareMetersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareMetersPerSecond(0)
    if err != nil {
        t.Errorf("FromSquareMetersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.KinematicViscositySquareMeterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMetersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromStokes function
func TestKinematicViscosityFactory_FromStokes(t *testing.T) {
    factory := units.KinematicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromStokes(100)
    if err != nil {
        t.Errorf("FromStokes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.KinematicViscosityStokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromStokes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromStokes(math.NaN())
    if err == nil {
        t.Error("FromStokes() with NaN value should return error")
    }

    _, err = factory.FromStokes(math.Inf(1))
    if err == nil {
        t.Error("FromStokes() with +Inf value should return error")
    }

    _, err = factory.FromStokes(math.Inf(-1))
    if err == nil {
        t.Error("FromStokes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromStokes(0)
    if err != nil {
        t.Errorf("FromStokes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.KinematicViscosityStokes)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromStokes() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareFeetPerSecond function
func TestKinematicViscosityFactory_FromSquareFeetPerSecond(t *testing.T) {
    factory := units.KinematicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareFeetPerSecond(100)
    if err != nil {
        t.Errorf("FromSquareFeetPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.KinematicViscositySquareFootPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareFeetPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareFeetPerSecond(math.NaN())
    if err == nil {
        t.Error("FromSquareFeetPerSecond() with NaN value should return error")
    }

    _, err = factory.FromSquareFeetPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromSquareFeetPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromSquareFeetPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareFeetPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareFeetPerSecond(0)
    if err != nil {
        t.Errorf("FromSquareFeetPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.KinematicViscositySquareFootPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareFeetPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromNanostokes function
func TestKinematicViscosityFactory_FromNanostokes(t *testing.T) {
    factory := units.KinematicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanostokes(100)
    if err != nil {
        t.Errorf("FromNanostokes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.KinematicViscosityNanostokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanostokes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanostokes(math.NaN())
    if err == nil {
        t.Error("FromNanostokes() with NaN value should return error")
    }

    _, err = factory.FromNanostokes(math.Inf(1))
    if err == nil {
        t.Error("FromNanostokes() with +Inf value should return error")
    }

    _, err = factory.FromNanostokes(math.Inf(-1))
    if err == nil {
        t.Error("FromNanostokes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanostokes(0)
    if err != nil {
        t.Errorf("FromNanostokes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.KinematicViscosityNanostokes)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanostokes() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrostokes function
func TestKinematicViscosityFactory_FromMicrostokes(t *testing.T) {
    factory := units.KinematicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrostokes(100)
    if err != nil {
        t.Errorf("FromMicrostokes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.KinematicViscosityMicrostokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrostokes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrostokes(math.NaN())
    if err == nil {
        t.Error("FromMicrostokes() with NaN value should return error")
    }

    _, err = factory.FromMicrostokes(math.Inf(1))
    if err == nil {
        t.Error("FromMicrostokes() with +Inf value should return error")
    }

    _, err = factory.FromMicrostokes(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrostokes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrostokes(0)
    if err != nil {
        t.Errorf("FromMicrostokes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.KinematicViscosityMicrostokes)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrostokes() with zero value = %v, want 0", converted)
    }
}
// Test FromMillistokes function
func TestKinematicViscosityFactory_FromMillistokes(t *testing.T) {
    factory := units.KinematicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillistokes(100)
    if err != nil {
        t.Errorf("FromMillistokes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.KinematicViscosityMillistokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillistokes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillistokes(math.NaN())
    if err == nil {
        t.Error("FromMillistokes() with NaN value should return error")
    }

    _, err = factory.FromMillistokes(math.Inf(1))
    if err == nil {
        t.Error("FromMillistokes() with +Inf value should return error")
    }

    _, err = factory.FromMillistokes(math.Inf(-1))
    if err == nil {
        t.Error("FromMillistokes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillistokes(0)
    if err != nil {
        t.Errorf("FromMillistokes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.KinematicViscosityMillistokes)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillistokes() with zero value = %v, want 0", converted)
    }
}
// Test FromCentistokes function
func TestKinematicViscosityFactory_FromCentistokes(t *testing.T) {
    factory := units.KinematicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentistokes(100)
    if err != nil {
        t.Errorf("FromCentistokes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.KinematicViscosityCentistokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentistokes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentistokes(math.NaN())
    if err == nil {
        t.Error("FromCentistokes() with NaN value should return error")
    }

    _, err = factory.FromCentistokes(math.Inf(1))
    if err == nil {
        t.Error("FromCentistokes() with +Inf value should return error")
    }

    _, err = factory.FromCentistokes(math.Inf(-1))
    if err == nil {
        t.Error("FromCentistokes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentistokes(0)
    if err != nil {
        t.Errorf("FromCentistokes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.KinematicViscosityCentistokes)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentistokes() with zero value = %v, want 0", converted)
    }
}
// Test FromDecistokes function
func TestKinematicViscosityFactory_FromDecistokes(t *testing.T) {
    factory := units.KinematicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecistokes(100)
    if err != nil {
        t.Errorf("FromDecistokes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.KinematicViscosityDecistokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecistokes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecistokes(math.NaN())
    if err == nil {
        t.Error("FromDecistokes() with NaN value should return error")
    }

    _, err = factory.FromDecistokes(math.Inf(1))
    if err == nil {
        t.Error("FromDecistokes() with +Inf value should return error")
    }

    _, err = factory.FromDecistokes(math.Inf(-1))
    if err == nil {
        t.Error("FromDecistokes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecistokes(0)
    if err != nil {
        t.Errorf("FromDecistokes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.KinematicViscosityDecistokes)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecistokes() with zero value = %v, want 0", converted)
    }
}
// Test FromKilostokes function
func TestKinematicViscosityFactory_FromKilostokes(t *testing.T) {
    factory := units.KinematicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilostokes(100)
    if err != nil {
        t.Errorf("FromKilostokes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.KinematicViscosityKilostokes)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilostokes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilostokes(math.NaN())
    if err == nil {
        t.Error("FromKilostokes() with NaN value should return error")
    }

    _, err = factory.FromKilostokes(math.Inf(1))
    if err == nil {
        t.Error("FromKilostokes() with +Inf value should return error")
    }

    _, err = factory.FromKilostokes(math.Inf(-1))
    if err == nil {
        t.Error("FromKilostokes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilostokes(0)
    if err != nil {
        t.Errorf("FromKilostokes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.KinematicViscosityKilostokes)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilostokes() with zero value = %v, want 0", converted)
    }
}

func TestKinematicViscosityToString(t *testing.T) {
	factory := units.KinematicViscosityFactory{}
	a, err := factory.CreateKinematicViscosity(45, units.KinematicViscositySquareMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.KinematicViscositySquareMeterPerSecond, 2)
	expected := "45.00 " + units.GetKinematicViscosityAbbreviation(units.KinematicViscositySquareMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.KinematicViscositySquareMeterPerSecond, -1)
	expected = "45 " + units.GetKinematicViscosityAbbreviation(units.KinematicViscositySquareMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestKinematicViscosity_EqualityAndComparison(t *testing.T) {
	factory := units.KinematicViscosityFactory{}
	a1, _ := factory.CreateKinematicViscosity(60, units.KinematicViscositySquareMeterPerSecond)
	a2, _ := factory.CreateKinematicViscosity(60, units.KinematicViscositySquareMeterPerSecond)
	a3, _ := factory.CreateKinematicViscosity(120, units.KinematicViscositySquareMeterPerSecond)

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

func TestKinematicViscosity_Arithmetic(t *testing.T) {
	factory := units.KinematicViscosityFactory{}
	a1, _ := factory.CreateKinematicViscosity(30, units.KinematicViscositySquareMeterPerSecond)
	a2, _ := factory.CreateKinematicViscosity(45, units.KinematicViscositySquareMeterPerSecond)

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


func TestGetKinematicViscosityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.KinematicViscosityUnits
        want string
    }{
        {
            name: "SquareMeterPerSecond abbreviation",
            unit: units.KinematicViscositySquareMeterPerSecond,
            want: "m²/s",
        },
        {
            name: "Stokes abbreviation",
            unit: units.KinematicViscosityStokes,
            want: "St",
        },
        {
            name: "SquareFootPerSecond abbreviation",
            unit: units.KinematicViscositySquareFootPerSecond,
            want: "ft²/s",
        },
        {
            name: "Nanostokes abbreviation",
            unit: units.KinematicViscosityNanostokes,
            want: "nSt",
        },
        {
            name: "Microstokes abbreviation",
            unit: units.KinematicViscosityMicrostokes,
            want: "μSt",
        },
        {
            name: "Millistokes abbreviation",
            unit: units.KinematicViscosityMillistokes,
            want: "mSt",
        },
        {
            name: "Centistokes abbreviation",
            unit: units.KinematicViscosityCentistokes,
            want: "cSt",
        },
        {
            name: "Decistokes abbreviation",
            unit: units.KinematicViscosityDecistokes,
            want: "dSt",
        },
        {
            name: "Kilostokes abbreviation",
            unit: units.KinematicViscosityKilostokes,
            want: "kSt",
        },
        {
            name: "invalid unit",
            unit: units.KinematicViscosityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetKinematicViscosityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetKinematicViscosityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestKinematicViscosity_String(t *testing.T) {
    factory := units.KinematicViscosityFactory{}
    
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
            unit, err := factory.CreateKinematicViscosity(tt.value, units.KinematicViscositySquareMeterPerSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("KinematicViscosity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestKinematicViscosity_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.KinematicViscosityFactory{}

	_, err := uf.CreateKinematicViscosity(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}