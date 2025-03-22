// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRotationalStiffnessPerLengthDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonMeterPerRadianPerMeter"}`
	
	factory := units.RotationalStiffnessPerLengthDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter {
		t.Errorf("expected unit %v, got %v", units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonMeterPerRadianPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRotationalStiffnessPerLengthDto_ToJSON(t *testing.T) {
	dto := units.RotationalStiffnessPerLengthDto{
		Value: 45,
		Unit:  units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter,
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
	if result["unit"].(string) != string(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter) {
		t.Errorf("expected unit %s, got %v", units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter, result["unit"])
	}
}

func TestNewRotationalStiffnessPerLength_InvalidValue(t *testing.T) {
	factory := units.RotationalStiffnessPerLengthFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRotationalStiffnessPerLength(math.NaN(), units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRotationalStiffnessPerLength(math.Inf(1), units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRotationalStiffnessPerLengthConversions(t *testing.T) {
	factory := units.RotationalStiffnessPerLengthFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRotationalStiffnessPerLength(180, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonMetersPerRadianPerMeter.
		// No expected conversion value provided for NewtonMetersPerRadianPerMeter, verifying result is not NaN.
		result := a.NewtonMetersPerRadianPerMeter()
		cacheResult := a.NewtonMetersPerRadianPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonMetersPerRadianPerMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundForceFeetPerDegreesPerFeet.
		// No expected conversion value provided for PoundForceFeetPerDegreesPerFeet, verifying result is not NaN.
		result := a.PoundForceFeetPerDegreesPerFeet()
		cacheResult := a.PoundForceFeetPerDegreesPerFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundForceFeetPerDegreesPerFeet returned NaN")
		}
	}
	{
		// Test conversion to KilopoundForceFeetPerDegreesPerFeet.
		// No expected conversion value provided for KilopoundForceFeetPerDegreesPerFeet, verifying result is not NaN.
		result := a.KilopoundForceFeetPerDegreesPerFeet()
		cacheResult := a.KilopoundForceFeetPerDegreesPerFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundForceFeetPerDegreesPerFeet returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMetersPerRadianPerMeter.
		// No expected conversion value provided for KilonewtonMetersPerRadianPerMeter, verifying result is not NaN.
		result := a.KilonewtonMetersPerRadianPerMeter()
		cacheResult := a.KilonewtonMetersPerRadianPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonMetersPerRadianPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMetersPerRadianPerMeter.
		// No expected conversion value provided for MeganewtonMetersPerRadianPerMeter, verifying result is not NaN.
		result := a.MeganewtonMetersPerRadianPerMeter()
		cacheResult := a.MeganewtonMetersPerRadianPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MeganewtonMetersPerRadianPerMeter returned NaN")
		}
	}
}

func TestRotationalStiffnessPerLength_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RotationalStiffnessPerLengthFactory{}
	a, err := factory.CreateRotationalStiffnessPerLength(90, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter {
		t.Errorf("expected default unit NewtonMeterPerRadianPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RotationalStiffnessPerLengthDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter {
		t.Errorf("expected unit NewtonMeterPerRadianPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRotationalStiffnessPerLengthFactory_FromDto(t *testing.T) {
    factory := units.RotationalStiffnessPerLengthFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.RotationalStiffnessPerLengthDto{
        Value: 100,
        Unit:  units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.RotationalStiffnessPerLengthDto{
        Value: math.NaN(),
        Unit:  units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test NewtonMeterPerRadianPerMeter conversion
    newton_meters_per_radian_per_meterDto := units.RotationalStiffnessPerLengthDto{
        Value: 100,
        Unit:  units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter,
    }
    
    var newton_meters_per_radian_per_meterResult *units.RotationalStiffnessPerLength
    newton_meters_per_radian_per_meterResult, err = factory.FromDto(newton_meters_per_radian_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonMeterPerRadianPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_meters_per_radian_per_meterResult.Convert(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMeterPerRadianPerMeter = %v, want %v", converted, 100)
    }
    // Test PoundForceFootPerDegreesPerFoot conversion
    pound_force_feet_per_degrees_per_feetDto := units.RotationalStiffnessPerLengthDto{
        Value: 100,
        Unit:  units.RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot,
    }
    
    var pound_force_feet_per_degrees_per_feetResult *units.RotationalStiffnessPerLength
    pound_force_feet_per_degrees_per_feetResult, err = factory.FromDto(pound_force_feet_per_degrees_per_feetDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForceFootPerDegreesPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_feet_per_degrees_per_feetResult.Convert(units.RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceFootPerDegreesPerFoot = %v, want %v", converted, 100)
    }
    // Test KilopoundForceFootPerDegreesPerFoot conversion
    kilopound_force_feet_per_degrees_per_feetDto := units.RotationalStiffnessPerLengthDto{
        Value: 100,
        Unit:  units.RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot,
    }
    
    var kilopound_force_feet_per_degrees_per_feetResult *units.RotationalStiffnessPerLength
    kilopound_force_feet_per_degrees_per_feetResult, err = factory.FromDto(kilopound_force_feet_per_degrees_per_feetDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForceFootPerDegreesPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_force_feet_per_degrees_per_feetResult.Convert(units.RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForceFootPerDegreesPerFoot = %v, want %v", converted, 100)
    }
    // Test KilonewtonMeterPerRadianPerMeter conversion
    kilonewton_meters_per_radian_per_meterDto := units.RotationalStiffnessPerLengthDto{
        Value: 100,
        Unit:  units.RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter,
    }
    
    var kilonewton_meters_per_radian_per_meterResult *units.RotationalStiffnessPerLength
    kilonewton_meters_per_radian_per_meterResult, err = factory.FromDto(kilonewton_meters_per_radian_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonMeterPerRadianPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_meters_per_radian_per_meterResult.Convert(units.RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMeterPerRadianPerMeter = %v, want %v", converted, 100)
    }
    // Test MeganewtonMeterPerRadianPerMeter conversion
    meganewton_meters_per_radian_per_meterDto := units.RotationalStiffnessPerLengthDto{
        Value: 100,
        Unit:  units.RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter,
    }
    
    var meganewton_meters_per_radian_per_meterResult *units.RotationalStiffnessPerLength
    meganewton_meters_per_radian_per_meterResult, err = factory.FromDto(meganewton_meters_per_radian_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonMeterPerRadianPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_meters_per_radian_per_meterResult.Convert(units.RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMeterPerRadianPerMeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.RotationalStiffnessPerLengthDto{
        Value: 0,
        Unit:  units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter,
    }
    
    var zeroResult *units.RotationalStiffnessPerLength
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestRotationalStiffnessPerLengthFactory_FromDtoJSON(t *testing.T) {
    factory := units.RotationalStiffnessPerLengthFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "NewtonMeterPerRadianPerMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "NewtonMeterPerRadianPerMeter"}`)
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
    nanJSON, _ := json.Marshal(units.RotationalStiffnessPerLengthDto{
        Value: nanValue,
        Unit:  units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with NewtonMeterPerRadianPerMeter unit
    newton_meters_per_radian_per_meterJSON := []byte(`{"value": 100, "unit": "NewtonMeterPerRadianPerMeter"}`)
    newton_meters_per_radian_per_meterResult, err := factory.FromDtoJSON(newton_meters_per_radian_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonMeterPerRadianPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_meters_per_radian_per_meterResult.Convert(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMeterPerRadianPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForceFootPerDegreesPerFoot unit
    pound_force_feet_per_degrees_per_feetJSON := []byte(`{"value": 100, "unit": "PoundForceFootPerDegreesPerFoot"}`)
    pound_force_feet_per_degrees_per_feetResult, err := factory.FromDtoJSON(pound_force_feet_per_degrees_per_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForceFootPerDegreesPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_feet_per_degrees_per_feetResult.Convert(units.RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceFootPerDegreesPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForceFootPerDegreesPerFoot unit
    kilopound_force_feet_per_degrees_per_feetJSON := []byte(`{"value": 100, "unit": "KilopoundForceFootPerDegreesPerFoot"}`)
    kilopound_force_feet_per_degrees_per_feetResult, err := factory.FromDtoJSON(kilopound_force_feet_per_degrees_per_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForceFootPerDegreesPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_force_feet_per_degrees_per_feetResult.Convert(units.RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForceFootPerDegreesPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonMeterPerRadianPerMeter unit
    kilonewton_meters_per_radian_per_meterJSON := []byte(`{"value": 100, "unit": "KilonewtonMeterPerRadianPerMeter"}`)
    kilonewton_meters_per_radian_per_meterResult, err := factory.FromDtoJSON(kilonewton_meters_per_radian_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonMeterPerRadianPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_meters_per_radian_per_meterResult.Convert(units.RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMeterPerRadianPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonMeterPerRadianPerMeter unit
    meganewton_meters_per_radian_per_meterJSON := []byte(`{"value": 100, "unit": "MeganewtonMeterPerRadianPerMeter"}`)
    meganewton_meters_per_radian_per_meterResult, err := factory.FromDtoJSON(meganewton_meters_per_radian_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonMeterPerRadianPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_meters_per_radian_per_meterResult.Convert(units.RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMeterPerRadianPerMeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "NewtonMeterPerRadianPerMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromNewtonMetersPerRadianPerMeter function
func TestRotationalStiffnessPerLengthFactory_FromNewtonMetersPerRadianPerMeter(t *testing.T) {
    factory := units.RotationalStiffnessPerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonMetersPerRadianPerMeter(100)
    if err != nil {
        t.Errorf("FromNewtonMetersPerRadianPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonMetersPerRadianPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonMetersPerRadianPerMeter(math.NaN())
    if err == nil {
        t.Error("FromNewtonMetersPerRadianPerMeter() with NaN value should return error")
    }

    _, err = factory.FromNewtonMetersPerRadianPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonMetersPerRadianPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromNewtonMetersPerRadianPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonMetersPerRadianPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonMetersPerRadianPerMeter(0)
    if err != nil {
        t.Errorf("FromNewtonMetersPerRadianPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonMetersPerRadianPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundForceFeetPerDegreesPerFeet function
func TestRotationalStiffnessPerLengthFactory_FromPoundForceFeetPerDegreesPerFeet(t *testing.T) {
    factory := units.RotationalStiffnessPerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundForceFeetPerDegreesPerFeet(100)
    if err != nil {
        t.Errorf("FromPoundForceFeetPerDegreesPerFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundForceFeetPerDegreesPerFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundForceFeetPerDegreesPerFeet(math.NaN())
    if err == nil {
        t.Error("FromPoundForceFeetPerDegreesPerFeet() with NaN value should return error")
    }

    _, err = factory.FromPoundForceFeetPerDegreesPerFeet(math.Inf(1))
    if err == nil {
        t.Error("FromPoundForceFeetPerDegreesPerFeet() with +Inf value should return error")
    }

    _, err = factory.FromPoundForceFeetPerDegreesPerFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundForceFeetPerDegreesPerFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundForceFeetPerDegreesPerFeet(0)
    if err != nil {
        t.Errorf("FromPoundForceFeetPerDegreesPerFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundForceFeetPerDegreesPerFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundForceFeetPerDegreesPerFeet function
func TestRotationalStiffnessPerLengthFactory_FromKilopoundForceFeetPerDegreesPerFeet(t *testing.T) {
    factory := units.RotationalStiffnessPerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundForceFeetPerDegreesPerFeet(100)
    if err != nil {
        t.Errorf("FromKilopoundForceFeetPerDegreesPerFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundForceFeetPerDegreesPerFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundForceFeetPerDegreesPerFeet(math.NaN())
    if err == nil {
        t.Error("FromKilopoundForceFeetPerDegreesPerFeet() with NaN value should return error")
    }

    _, err = factory.FromKilopoundForceFeetPerDegreesPerFeet(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundForceFeetPerDegreesPerFeet() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundForceFeetPerDegreesPerFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundForceFeetPerDegreesPerFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundForceFeetPerDegreesPerFeet(0)
    if err != nil {
        t.Errorf("FromKilopoundForceFeetPerDegreesPerFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundForceFeetPerDegreesPerFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonMetersPerRadianPerMeter function
func TestRotationalStiffnessPerLengthFactory_FromKilonewtonMetersPerRadianPerMeter(t *testing.T) {
    factory := units.RotationalStiffnessPerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonMetersPerRadianPerMeter(100)
    if err != nil {
        t.Errorf("FromKilonewtonMetersPerRadianPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonMetersPerRadianPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonMetersPerRadianPerMeter(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonMetersPerRadianPerMeter() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonMetersPerRadianPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonMetersPerRadianPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonMetersPerRadianPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonMetersPerRadianPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonMetersPerRadianPerMeter(0)
    if err != nil {
        t.Errorf("FromKilonewtonMetersPerRadianPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonMetersPerRadianPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonMetersPerRadianPerMeter function
func TestRotationalStiffnessPerLengthFactory_FromMeganewtonMetersPerRadianPerMeter(t *testing.T) {
    factory := units.RotationalStiffnessPerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonMetersPerRadianPerMeter(100)
    if err != nil {
        t.Errorf("FromMeganewtonMetersPerRadianPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonMetersPerRadianPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonMetersPerRadianPerMeter(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonMetersPerRadianPerMeter() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonMetersPerRadianPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonMetersPerRadianPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonMetersPerRadianPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonMetersPerRadianPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonMetersPerRadianPerMeter(0)
    if err != nil {
        t.Errorf("FromMeganewtonMetersPerRadianPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonMetersPerRadianPerMeter() with zero value = %v, want 0", converted)
    }
}

func TestRotationalStiffnessPerLengthToString(t *testing.T) {
	factory := units.RotationalStiffnessPerLengthFactory{}
	a, err := factory.CreateRotationalStiffnessPerLength(45, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter, 2)
	expected := "45.00 " + units.GetRotationalStiffnessPerLengthAbbreviation(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter, -1)
	expected = "45 " + units.GetRotationalStiffnessPerLengthAbbreviation(units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRotationalStiffnessPerLength_EqualityAndComparison(t *testing.T) {
	factory := units.RotationalStiffnessPerLengthFactory{}
	a1, _ := factory.CreateRotationalStiffnessPerLength(60, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	a2, _ := factory.CreateRotationalStiffnessPerLength(60, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	a3, _ := factory.CreateRotationalStiffnessPerLength(120, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)

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

func TestRotationalStiffnessPerLength_Arithmetic(t *testing.T) {
	factory := units.RotationalStiffnessPerLengthFactory{}
	a1, _ := factory.CreateRotationalStiffnessPerLength(30, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	a2, _ := factory.CreateRotationalStiffnessPerLength(45, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)

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


func TestGetRotationalStiffnessPerLengthAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.RotationalStiffnessPerLengthUnits
        want string
    }{
        {
            name: "NewtonMeterPerRadianPerMeter abbreviation",
            unit: units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter,
            want: "N·m/rad/m",
        },
        {
            name: "PoundForceFootPerDegreesPerFoot abbreviation",
            unit: units.RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot,
            want: "lbf·ft/deg/ft",
        },
        {
            name: "KilopoundForceFootPerDegreesPerFoot abbreviation",
            unit: units.RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot,
            want: "kipf·ft/°/ft",
        },
        {
            name: "KilonewtonMeterPerRadianPerMeter abbreviation",
            unit: units.RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter,
            want: "kN·m/rad/m",
        },
        {
            name: "MeganewtonMeterPerRadianPerMeter abbreviation",
            unit: units.RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter,
            want: "MN·m/rad/m",
        },
        {
            name: "invalid unit",
            unit: units.RotationalStiffnessPerLengthUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetRotationalStiffnessPerLengthAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetRotationalStiffnessPerLengthAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestRotationalStiffnessPerLength_String(t *testing.T) {
    factory := units.RotationalStiffnessPerLengthFactory{}
    
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
            unit, err := factory.CreateRotationalStiffnessPerLength(tt.value, units.RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("RotationalStiffnessPerLength.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestRotationalStiffnessPerLength_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.RotationalStiffnessPerLengthFactory{}

	_, err := uf.CreateRotationalStiffnessPerLength(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}