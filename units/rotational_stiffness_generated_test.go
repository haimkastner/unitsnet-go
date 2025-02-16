// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRotationalStiffnessDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonMeterPerRadian"}`
	
	factory := units.RotationalStiffnessDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RotationalStiffnessNewtonMeterPerRadian {
		t.Errorf("expected unit %v, got %v", units.RotationalStiffnessNewtonMeterPerRadian, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonMeterPerRadian"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRotationalStiffnessDto_ToJSON(t *testing.T) {
	dto := units.RotationalStiffnessDto{
		Value: 45,
		Unit:  units.RotationalStiffnessNewtonMeterPerRadian,
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
	if result["unit"].(string) != string(units.RotationalStiffnessNewtonMeterPerRadian) {
		t.Errorf("expected unit %s, got %v", units.RotationalStiffnessNewtonMeterPerRadian, result["unit"])
	}
}

func TestNewRotationalStiffness_InvalidValue(t *testing.T) {
	factory := units.RotationalStiffnessFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRotationalStiffness(math.NaN(), units.RotationalStiffnessNewtonMeterPerRadian)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRotationalStiffness(math.Inf(1), units.RotationalStiffnessNewtonMeterPerRadian)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRotationalStiffnessConversions(t *testing.T) {
	factory := units.RotationalStiffnessFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRotationalStiffness(180, units.RotationalStiffnessNewtonMeterPerRadian)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonMetersPerRadian.
		// No expected conversion value provided for NewtonMetersPerRadian, verifying result is not NaN.
		result := a.NewtonMetersPerRadian()
		cacheResult := a.NewtonMetersPerRadian()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonMetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to PoundForceFeetPerDegrees.
		// No expected conversion value provided for PoundForceFeetPerDegrees, verifying result is not NaN.
		result := a.PoundForceFeetPerDegrees()
		cacheResult := a.PoundForceFeetPerDegrees()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundForceFeetPerDegrees returned NaN")
		}
	}
	{
		// Test conversion to KilopoundForceFeetPerDegrees.
		// No expected conversion value provided for KilopoundForceFeetPerDegrees, verifying result is not NaN.
		result := a.KilopoundForceFeetPerDegrees()
		cacheResult := a.KilopoundForceFeetPerDegrees()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundForceFeetPerDegrees returned NaN")
		}
	}
	{
		// Test conversion to NewtonMillimetersPerDegree.
		// No expected conversion value provided for NewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.NewtonMillimetersPerDegree()
		cacheResult := a.NewtonMillimetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to NewtonMetersPerDegree.
		// No expected conversion value provided for NewtonMetersPerDegree, verifying result is not NaN.
		result := a.NewtonMetersPerDegree()
		cacheResult := a.NewtonMetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to NewtonMillimetersPerRadian.
		// No expected conversion value provided for NewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.NewtonMillimetersPerRadian()
		cacheResult := a.NewtonMillimetersPerRadian()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to PoundForceFeetPerRadian.
		// No expected conversion value provided for PoundForceFeetPerRadian, verifying result is not NaN.
		result := a.PoundForceFeetPerRadian()
		cacheResult := a.PoundForceFeetPerRadian()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundForceFeetPerRadian returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMetersPerRadian.
		// No expected conversion value provided for KilonewtonMetersPerRadian, verifying result is not NaN.
		result := a.KilonewtonMetersPerRadian()
		cacheResult := a.KilonewtonMetersPerRadian()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonMetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMetersPerRadian.
		// No expected conversion value provided for MeganewtonMetersPerRadian, verifying result is not NaN.
		result := a.MeganewtonMetersPerRadian()
		cacheResult := a.MeganewtonMetersPerRadian()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MeganewtonMetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonMillimetersPerDegree.
		// No expected conversion value provided for NanonewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.NanonewtonMillimetersPerDegree()
		cacheResult := a.NanonewtonMillimetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanonewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonMillimetersPerDegree.
		// No expected conversion value provided for MicronewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.MicronewtonMillimetersPerDegree()
		cacheResult := a.MicronewtonMillimetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicronewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonMillimetersPerDegree.
		// No expected conversion value provided for MillinewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.MillinewtonMillimetersPerDegree()
		cacheResult := a.MillinewtonMillimetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillinewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonMillimetersPerDegree.
		// No expected conversion value provided for CentinewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.CentinewtonMillimetersPerDegree()
		cacheResult := a.CentinewtonMillimetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentinewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonMillimetersPerDegree.
		// No expected conversion value provided for DecinewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.DecinewtonMillimetersPerDegree()
		cacheResult := a.DecinewtonMillimetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecinewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonMillimetersPerDegree.
		// No expected conversion value provided for DecanewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.DecanewtonMillimetersPerDegree()
		cacheResult := a.DecanewtonMillimetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecanewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMillimetersPerDegree.
		// No expected conversion value provided for KilonewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.KilonewtonMillimetersPerDegree()
		cacheResult := a.KilonewtonMillimetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMillimetersPerDegree.
		// No expected conversion value provided for MeganewtonMillimetersPerDegree, verifying result is not NaN.
		result := a.MeganewtonMillimetersPerDegree()
		cacheResult := a.MeganewtonMillimetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MeganewtonMillimetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonMetersPerDegree.
		// No expected conversion value provided for NanonewtonMetersPerDegree, verifying result is not NaN.
		result := a.NanonewtonMetersPerDegree()
		cacheResult := a.NanonewtonMetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanonewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonMetersPerDegree.
		// No expected conversion value provided for MicronewtonMetersPerDegree, verifying result is not NaN.
		result := a.MicronewtonMetersPerDegree()
		cacheResult := a.MicronewtonMetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicronewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonMetersPerDegree.
		// No expected conversion value provided for MillinewtonMetersPerDegree, verifying result is not NaN.
		result := a.MillinewtonMetersPerDegree()
		cacheResult := a.MillinewtonMetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillinewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonMetersPerDegree.
		// No expected conversion value provided for CentinewtonMetersPerDegree, verifying result is not NaN.
		result := a.CentinewtonMetersPerDegree()
		cacheResult := a.CentinewtonMetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentinewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonMetersPerDegree.
		// No expected conversion value provided for DecinewtonMetersPerDegree, verifying result is not NaN.
		result := a.DecinewtonMetersPerDegree()
		cacheResult := a.DecinewtonMetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecinewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonMetersPerDegree.
		// No expected conversion value provided for DecanewtonMetersPerDegree, verifying result is not NaN.
		result := a.DecanewtonMetersPerDegree()
		cacheResult := a.DecanewtonMetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecanewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMetersPerDegree.
		// No expected conversion value provided for KilonewtonMetersPerDegree, verifying result is not NaN.
		result := a.KilonewtonMetersPerDegree()
		cacheResult := a.KilonewtonMetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMetersPerDegree.
		// No expected conversion value provided for MeganewtonMetersPerDegree, verifying result is not NaN.
		result := a.MeganewtonMetersPerDegree()
		cacheResult := a.MeganewtonMetersPerDegree()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MeganewtonMetersPerDegree returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonMillimetersPerRadian.
		// No expected conversion value provided for NanonewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.NanonewtonMillimetersPerRadian()
		cacheResult := a.NanonewtonMillimetersPerRadian()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanonewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonMillimetersPerRadian.
		// No expected conversion value provided for MicronewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.MicronewtonMillimetersPerRadian()
		cacheResult := a.MicronewtonMillimetersPerRadian()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicronewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonMillimetersPerRadian.
		// No expected conversion value provided for MillinewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.MillinewtonMillimetersPerRadian()
		cacheResult := a.MillinewtonMillimetersPerRadian()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillinewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonMillimetersPerRadian.
		// No expected conversion value provided for CentinewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.CentinewtonMillimetersPerRadian()
		cacheResult := a.CentinewtonMillimetersPerRadian()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentinewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonMillimetersPerRadian.
		// No expected conversion value provided for DecinewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.DecinewtonMillimetersPerRadian()
		cacheResult := a.DecinewtonMillimetersPerRadian()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecinewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonMillimetersPerRadian.
		// No expected conversion value provided for DecanewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.DecanewtonMillimetersPerRadian()
		cacheResult := a.DecanewtonMillimetersPerRadian()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecanewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMillimetersPerRadian.
		// No expected conversion value provided for KilonewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.KilonewtonMillimetersPerRadian()
		cacheResult := a.KilonewtonMillimetersPerRadian()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonMillimetersPerRadian returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMillimetersPerRadian.
		// No expected conversion value provided for MeganewtonMillimetersPerRadian, verifying result is not NaN.
		result := a.MeganewtonMillimetersPerRadian()
		cacheResult := a.MeganewtonMillimetersPerRadian()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MeganewtonMillimetersPerRadian returned NaN")
		}
	}
}

func TestRotationalStiffness_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RotationalStiffnessFactory{}
	a, err := factory.CreateRotationalStiffness(90, units.RotationalStiffnessNewtonMeterPerRadian)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RotationalStiffnessNewtonMeterPerRadian {
		t.Errorf("expected default unit NewtonMeterPerRadian, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RotationalStiffnessNewtonMeterPerRadian
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RotationalStiffnessDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RotationalStiffnessNewtonMeterPerRadian {
		t.Errorf("expected unit NewtonMeterPerRadian, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRotationalStiffnessFactory_FromDto(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessNewtonMeterPerRadian,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.RotationalStiffnessDto{
        Value: math.NaN(),
        Unit:  units.RotationalStiffnessNewtonMeterPerRadian,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test NewtonMeterPerRadian conversion
    newton_meters_per_radianDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessNewtonMeterPerRadian,
    }
    
    var newton_meters_per_radianResult *units.RotationalStiffness
    newton_meters_per_radianResult, err = factory.FromDto(newton_meters_per_radianDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonMeterPerRadian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_meters_per_radianResult.Convert(units.RotationalStiffnessNewtonMeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMeterPerRadian = %v, want %v", converted, 100)
    }
    // Test PoundForceFootPerDegrees conversion
    pound_force_feet_per_degreesDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessPoundForceFootPerDegrees,
    }
    
    var pound_force_feet_per_degreesResult *units.RotationalStiffness
    pound_force_feet_per_degreesResult, err = factory.FromDto(pound_force_feet_per_degreesDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForceFootPerDegrees returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_feet_per_degreesResult.Convert(units.RotationalStiffnessPoundForceFootPerDegrees)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceFootPerDegrees = %v, want %v", converted, 100)
    }
    // Test KilopoundForceFootPerDegrees conversion
    kilopound_force_feet_per_degreesDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessKilopoundForceFootPerDegrees,
    }
    
    var kilopound_force_feet_per_degreesResult *units.RotationalStiffness
    kilopound_force_feet_per_degreesResult, err = factory.FromDto(kilopound_force_feet_per_degreesDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForceFootPerDegrees returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_force_feet_per_degreesResult.Convert(units.RotationalStiffnessKilopoundForceFootPerDegrees)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForceFootPerDegrees = %v, want %v", converted, 100)
    }
    // Test NewtonMillimeterPerDegree conversion
    newton_millimeters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessNewtonMillimeterPerDegree,
    }
    
    var newton_millimeters_per_degreeResult *units.RotationalStiffness
    newton_millimeters_per_degreeResult, err = factory.FromDto(newton_millimeters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonMillimeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessNewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test NewtonMeterPerDegree conversion
    newton_meters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessNewtonMeterPerDegree,
    }
    
    var newton_meters_per_degreeResult *units.RotationalStiffness
    newton_meters_per_degreeResult, err = factory.FromDto(newton_meters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonMeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_meters_per_degreeResult.Convert(units.RotationalStiffnessNewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test NewtonMillimeterPerRadian conversion
    newton_millimeters_per_radianDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessNewtonMillimeterPerRadian,
    }
    
    var newton_millimeters_per_radianResult *units.RotationalStiffness
    newton_millimeters_per_radianResult, err = factory.FromDto(newton_millimeters_per_radianDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonMillimeterPerRadian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_millimeters_per_radianResult.Convert(units.RotationalStiffnessNewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test PoundForceFeetPerRadian conversion
    pound_force_feet_per_radianDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessPoundForceFeetPerRadian,
    }
    
    var pound_force_feet_per_radianResult *units.RotationalStiffness
    pound_force_feet_per_radianResult, err = factory.FromDto(pound_force_feet_per_radianDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForceFeetPerRadian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_feet_per_radianResult.Convert(units.RotationalStiffnessPoundForceFeetPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceFeetPerRadian = %v, want %v", converted, 100)
    }
    // Test KilonewtonMeterPerRadian conversion
    kilonewton_meters_per_radianDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessKilonewtonMeterPerRadian,
    }
    
    var kilonewton_meters_per_radianResult *units.RotationalStiffness
    kilonewton_meters_per_radianResult, err = factory.FromDto(kilonewton_meters_per_radianDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonMeterPerRadian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_meters_per_radianResult.Convert(units.RotationalStiffnessKilonewtonMeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMeterPerRadian = %v, want %v", converted, 100)
    }
    // Test MeganewtonMeterPerRadian conversion
    meganewton_meters_per_radianDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessMeganewtonMeterPerRadian,
    }
    
    var meganewton_meters_per_radianResult *units.RotationalStiffness
    meganewton_meters_per_radianResult, err = factory.FromDto(meganewton_meters_per_radianDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonMeterPerRadian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_meters_per_radianResult.Convert(units.RotationalStiffnessMeganewtonMeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMeterPerRadian = %v, want %v", converted, 100)
    }
    // Test NanonewtonMillimeterPerDegree conversion
    nanonewton_millimeters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessNanonewtonMillimeterPerDegree,
    }
    
    var nanonewton_millimeters_per_degreeResult *units.RotationalStiffness
    nanonewton_millimeters_per_degreeResult, err = factory.FromDto(nanonewton_millimeters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with NanonewtonMillimeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessNanonewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test MicronewtonMillimeterPerDegree conversion
    micronewton_millimeters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessMicronewtonMillimeterPerDegree,
    }
    
    var micronewton_millimeters_per_degreeResult *units.RotationalStiffness
    micronewton_millimeters_per_degreeResult, err = factory.FromDto(micronewton_millimeters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with MicronewtonMillimeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessMicronewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test MillinewtonMillimeterPerDegree conversion
    millinewton_millimeters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessMillinewtonMillimeterPerDegree,
    }
    
    var millinewton_millimeters_per_degreeResult *units.RotationalStiffness
    millinewton_millimeters_per_degreeResult, err = factory.FromDto(millinewton_millimeters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with MillinewtonMillimeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessMillinewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test CentinewtonMillimeterPerDegree conversion
    centinewton_millimeters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessCentinewtonMillimeterPerDegree,
    }
    
    var centinewton_millimeters_per_degreeResult *units.RotationalStiffness
    centinewton_millimeters_per_degreeResult, err = factory.FromDto(centinewton_millimeters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with CentinewtonMillimeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessCentinewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test DecinewtonMillimeterPerDegree conversion
    decinewton_millimeters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessDecinewtonMillimeterPerDegree,
    }
    
    var decinewton_millimeters_per_degreeResult *units.RotationalStiffness
    decinewton_millimeters_per_degreeResult, err = factory.FromDto(decinewton_millimeters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with DecinewtonMillimeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessDecinewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test DecanewtonMillimeterPerDegree conversion
    decanewton_millimeters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessDecanewtonMillimeterPerDegree,
    }
    
    var decanewton_millimeters_per_degreeResult *units.RotationalStiffness
    decanewton_millimeters_per_degreeResult, err = factory.FromDto(decanewton_millimeters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with DecanewtonMillimeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessDecanewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test KilonewtonMillimeterPerDegree conversion
    kilonewton_millimeters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessKilonewtonMillimeterPerDegree,
    }
    
    var kilonewton_millimeters_per_degreeResult *units.RotationalStiffness
    kilonewton_millimeters_per_degreeResult, err = factory.FromDto(kilonewton_millimeters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonMillimeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessKilonewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test MeganewtonMillimeterPerDegree conversion
    meganewton_millimeters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessMeganewtonMillimeterPerDegree,
    }
    
    var meganewton_millimeters_per_degreeResult *units.RotationalStiffness
    meganewton_millimeters_per_degreeResult, err = factory.FromDto(meganewton_millimeters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonMillimeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessMeganewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test NanonewtonMeterPerDegree conversion
    nanonewton_meters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessNanonewtonMeterPerDegree,
    }
    
    var nanonewton_meters_per_degreeResult *units.RotationalStiffness
    nanonewton_meters_per_degreeResult, err = factory.FromDto(nanonewton_meters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with NanonewtonMeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewton_meters_per_degreeResult.Convert(units.RotationalStiffnessNanonewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test MicronewtonMeterPerDegree conversion
    micronewton_meters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessMicronewtonMeterPerDegree,
    }
    
    var micronewton_meters_per_degreeResult *units.RotationalStiffness
    micronewton_meters_per_degreeResult, err = factory.FromDto(micronewton_meters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with MicronewtonMeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewton_meters_per_degreeResult.Convert(units.RotationalStiffnessMicronewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test MillinewtonMeterPerDegree conversion
    millinewton_meters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessMillinewtonMeterPerDegree,
    }
    
    var millinewton_meters_per_degreeResult *units.RotationalStiffness
    millinewton_meters_per_degreeResult, err = factory.FromDto(millinewton_meters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with MillinewtonMeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewton_meters_per_degreeResult.Convert(units.RotationalStiffnessMillinewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test CentinewtonMeterPerDegree conversion
    centinewton_meters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessCentinewtonMeterPerDegree,
    }
    
    var centinewton_meters_per_degreeResult *units.RotationalStiffness
    centinewton_meters_per_degreeResult, err = factory.FromDto(centinewton_meters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with CentinewtonMeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewton_meters_per_degreeResult.Convert(units.RotationalStiffnessCentinewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test DecinewtonMeterPerDegree conversion
    decinewton_meters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessDecinewtonMeterPerDegree,
    }
    
    var decinewton_meters_per_degreeResult *units.RotationalStiffness
    decinewton_meters_per_degreeResult, err = factory.FromDto(decinewton_meters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with DecinewtonMeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewton_meters_per_degreeResult.Convert(units.RotationalStiffnessDecinewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test DecanewtonMeterPerDegree conversion
    decanewton_meters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessDecanewtonMeterPerDegree,
    }
    
    var decanewton_meters_per_degreeResult *units.RotationalStiffness
    decanewton_meters_per_degreeResult, err = factory.FromDto(decanewton_meters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with DecanewtonMeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewton_meters_per_degreeResult.Convert(units.RotationalStiffnessDecanewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test KilonewtonMeterPerDegree conversion
    kilonewton_meters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessKilonewtonMeterPerDegree,
    }
    
    var kilonewton_meters_per_degreeResult *units.RotationalStiffness
    kilonewton_meters_per_degreeResult, err = factory.FromDto(kilonewton_meters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonMeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_meters_per_degreeResult.Convert(units.RotationalStiffnessKilonewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test MeganewtonMeterPerDegree conversion
    meganewton_meters_per_degreeDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessMeganewtonMeterPerDegree,
    }
    
    var meganewton_meters_per_degreeResult *units.RotationalStiffness
    meganewton_meters_per_degreeResult, err = factory.FromDto(meganewton_meters_per_degreeDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonMeterPerDegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_meters_per_degreeResult.Convert(units.RotationalStiffnessMeganewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test NanonewtonMillimeterPerRadian conversion
    nanonewton_millimeters_per_radianDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessNanonewtonMillimeterPerRadian,
    }
    
    var nanonewton_millimeters_per_radianResult *units.RotationalStiffness
    nanonewton_millimeters_per_radianResult, err = factory.FromDto(nanonewton_millimeters_per_radianDto)
    if err != nil {
        t.Errorf("FromDto() with NanonewtonMillimeterPerRadian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessNanonewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test MicronewtonMillimeterPerRadian conversion
    micronewton_millimeters_per_radianDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessMicronewtonMillimeterPerRadian,
    }
    
    var micronewton_millimeters_per_radianResult *units.RotationalStiffness
    micronewton_millimeters_per_radianResult, err = factory.FromDto(micronewton_millimeters_per_radianDto)
    if err != nil {
        t.Errorf("FromDto() with MicronewtonMillimeterPerRadian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessMicronewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test MillinewtonMillimeterPerRadian conversion
    millinewton_millimeters_per_radianDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessMillinewtonMillimeterPerRadian,
    }
    
    var millinewton_millimeters_per_radianResult *units.RotationalStiffness
    millinewton_millimeters_per_radianResult, err = factory.FromDto(millinewton_millimeters_per_radianDto)
    if err != nil {
        t.Errorf("FromDto() with MillinewtonMillimeterPerRadian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessMillinewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test CentinewtonMillimeterPerRadian conversion
    centinewton_millimeters_per_radianDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessCentinewtonMillimeterPerRadian,
    }
    
    var centinewton_millimeters_per_radianResult *units.RotationalStiffness
    centinewton_millimeters_per_radianResult, err = factory.FromDto(centinewton_millimeters_per_radianDto)
    if err != nil {
        t.Errorf("FromDto() with CentinewtonMillimeterPerRadian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessCentinewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test DecinewtonMillimeterPerRadian conversion
    decinewton_millimeters_per_radianDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessDecinewtonMillimeterPerRadian,
    }
    
    var decinewton_millimeters_per_radianResult *units.RotationalStiffness
    decinewton_millimeters_per_radianResult, err = factory.FromDto(decinewton_millimeters_per_radianDto)
    if err != nil {
        t.Errorf("FromDto() with DecinewtonMillimeterPerRadian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessDecinewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test DecanewtonMillimeterPerRadian conversion
    decanewton_millimeters_per_radianDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessDecanewtonMillimeterPerRadian,
    }
    
    var decanewton_millimeters_per_radianResult *units.RotationalStiffness
    decanewton_millimeters_per_radianResult, err = factory.FromDto(decanewton_millimeters_per_radianDto)
    if err != nil {
        t.Errorf("FromDto() with DecanewtonMillimeterPerRadian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessDecanewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test KilonewtonMillimeterPerRadian conversion
    kilonewton_millimeters_per_radianDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessKilonewtonMillimeterPerRadian,
    }
    
    var kilonewton_millimeters_per_radianResult *units.RotationalStiffness
    kilonewton_millimeters_per_radianResult, err = factory.FromDto(kilonewton_millimeters_per_radianDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonMillimeterPerRadian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessKilonewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test MeganewtonMillimeterPerRadian conversion
    meganewton_millimeters_per_radianDto := units.RotationalStiffnessDto{
        Value: 100,
        Unit:  units.RotationalStiffnessMeganewtonMillimeterPerRadian,
    }
    
    var meganewton_millimeters_per_radianResult *units.RotationalStiffness
    meganewton_millimeters_per_radianResult, err = factory.FromDto(meganewton_millimeters_per_radianDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonMillimeterPerRadian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessMeganewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.RotationalStiffnessDto{
        Value: 0,
        Unit:  units.RotationalStiffnessNewtonMeterPerRadian,
    }
    
    var zeroResult *units.RotationalStiffness
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestRotationalStiffnessFactory_FromDtoJSON(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "NewtonMeterPerRadian"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "NewtonMeterPerRadian"}`)
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
    nanJSON, _ := json.Marshal(units.RotationalStiffnessDto{
        Value: nanValue,
        Unit:  units.RotationalStiffnessNewtonMeterPerRadian,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with NewtonMeterPerRadian unit
    newton_meters_per_radianJSON := []byte(`{"value": 100, "unit": "NewtonMeterPerRadian"}`)
    newton_meters_per_radianResult, err := factory.FromDtoJSON(newton_meters_per_radianJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonMeterPerRadian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_meters_per_radianResult.Convert(units.RotationalStiffnessNewtonMeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMeterPerRadian = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForceFootPerDegrees unit
    pound_force_feet_per_degreesJSON := []byte(`{"value": 100, "unit": "PoundForceFootPerDegrees"}`)
    pound_force_feet_per_degreesResult, err := factory.FromDtoJSON(pound_force_feet_per_degreesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForceFootPerDegrees unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_feet_per_degreesResult.Convert(units.RotationalStiffnessPoundForceFootPerDegrees)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceFootPerDegrees = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForceFootPerDegrees unit
    kilopound_force_feet_per_degreesJSON := []byte(`{"value": 100, "unit": "KilopoundForceFootPerDegrees"}`)
    kilopound_force_feet_per_degreesResult, err := factory.FromDtoJSON(kilopound_force_feet_per_degreesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForceFootPerDegrees unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_force_feet_per_degreesResult.Convert(units.RotationalStiffnessKilopoundForceFootPerDegrees)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForceFootPerDegrees = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonMillimeterPerDegree unit
    newton_millimeters_per_degreeJSON := []byte(`{"value": 100, "unit": "NewtonMillimeterPerDegree"}`)
    newton_millimeters_per_degreeResult, err := factory.FromDtoJSON(newton_millimeters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonMillimeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessNewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonMeterPerDegree unit
    newton_meters_per_degreeJSON := []byte(`{"value": 100, "unit": "NewtonMeterPerDegree"}`)
    newton_meters_per_degreeResult, err := factory.FromDtoJSON(newton_meters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonMeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_meters_per_degreeResult.Convert(units.RotationalStiffnessNewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonMillimeterPerRadian unit
    newton_millimeters_per_radianJSON := []byte(`{"value": 100, "unit": "NewtonMillimeterPerRadian"}`)
    newton_millimeters_per_radianResult, err := factory.FromDtoJSON(newton_millimeters_per_radianJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonMillimeterPerRadian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_millimeters_per_radianResult.Convert(units.RotationalStiffnessNewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForceFeetPerRadian unit
    pound_force_feet_per_radianJSON := []byte(`{"value": 100, "unit": "PoundForceFeetPerRadian"}`)
    pound_force_feet_per_radianResult, err := factory.FromDtoJSON(pound_force_feet_per_radianJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForceFeetPerRadian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_feet_per_radianResult.Convert(units.RotationalStiffnessPoundForceFeetPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceFeetPerRadian = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonMeterPerRadian unit
    kilonewton_meters_per_radianJSON := []byte(`{"value": 100, "unit": "KilonewtonMeterPerRadian"}`)
    kilonewton_meters_per_radianResult, err := factory.FromDtoJSON(kilonewton_meters_per_radianJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonMeterPerRadian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_meters_per_radianResult.Convert(units.RotationalStiffnessKilonewtonMeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMeterPerRadian = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonMeterPerRadian unit
    meganewton_meters_per_radianJSON := []byte(`{"value": 100, "unit": "MeganewtonMeterPerRadian"}`)
    meganewton_meters_per_radianResult, err := factory.FromDtoJSON(meganewton_meters_per_radianJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonMeterPerRadian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_meters_per_radianResult.Convert(units.RotationalStiffnessMeganewtonMeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMeterPerRadian = %v, want %v", converted, 100)
    }
    // Test JSON with NanonewtonMillimeterPerDegree unit
    nanonewton_millimeters_per_degreeJSON := []byte(`{"value": 100, "unit": "NanonewtonMillimeterPerDegree"}`)
    nanonewton_millimeters_per_degreeResult, err := factory.FromDtoJSON(nanonewton_millimeters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanonewtonMillimeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessNanonewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with MicronewtonMillimeterPerDegree unit
    micronewton_millimeters_per_degreeJSON := []byte(`{"value": 100, "unit": "MicronewtonMillimeterPerDegree"}`)
    micronewton_millimeters_per_degreeResult, err := factory.FromDtoJSON(micronewton_millimeters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicronewtonMillimeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessMicronewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with MillinewtonMillimeterPerDegree unit
    millinewton_millimeters_per_degreeJSON := []byte(`{"value": 100, "unit": "MillinewtonMillimeterPerDegree"}`)
    millinewton_millimeters_per_degreeResult, err := factory.FromDtoJSON(millinewton_millimeters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillinewtonMillimeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessMillinewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with CentinewtonMillimeterPerDegree unit
    centinewton_millimeters_per_degreeJSON := []byte(`{"value": 100, "unit": "CentinewtonMillimeterPerDegree"}`)
    centinewton_millimeters_per_degreeResult, err := factory.FromDtoJSON(centinewton_millimeters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentinewtonMillimeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessCentinewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with DecinewtonMillimeterPerDegree unit
    decinewton_millimeters_per_degreeJSON := []byte(`{"value": 100, "unit": "DecinewtonMillimeterPerDegree"}`)
    decinewton_millimeters_per_degreeResult, err := factory.FromDtoJSON(decinewton_millimeters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecinewtonMillimeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessDecinewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with DecanewtonMillimeterPerDegree unit
    decanewton_millimeters_per_degreeJSON := []byte(`{"value": 100, "unit": "DecanewtonMillimeterPerDegree"}`)
    decanewton_millimeters_per_degreeResult, err := factory.FromDtoJSON(decanewton_millimeters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecanewtonMillimeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessDecanewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonMillimeterPerDegree unit
    kilonewton_millimeters_per_degreeJSON := []byte(`{"value": 100, "unit": "KilonewtonMillimeterPerDegree"}`)
    kilonewton_millimeters_per_degreeResult, err := factory.FromDtoJSON(kilonewton_millimeters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonMillimeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessKilonewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonMillimeterPerDegree unit
    meganewton_millimeters_per_degreeJSON := []byte(`{"value": 100, "unit": "MeganewtonMillimeterPerDegree"}`)
    meganewton_millimeters_per_degreeResult, err := factory.FromDtoJSON(meganewton_millimeters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonMillimeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_millimeters_per_degreeResult.Convert(units.RotationalStiffnessMeganewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMillimeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with NanonewtonMeterPerDegree unit
    nanonewton_meters_per_degreeJSON := []byte(`{"value": 100, "unit": "NanonewtonMeterPerDegree"}`)
    nanonewton_meters_per_degreeResult, err := factory.FromDtoJSON(nanonewton_meters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanonewtonMeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewton_meters_per_degreeResult.Convert(units.RotationalStiffnessNanonewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with MicronewtonMeterPerDegree unit
    micronewton_meters_per_degreeJSON := []byte(`{"value": 100, "unit": "MicronewtonMeterPerDegree"}`)
    micronewton_meters_per_degreeResult, err := factory.FromDtoJSON(micronewton_meters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicronewtonMeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewton_meters_per_degreeResult.Convert(units.RotationalStiffnessMicronewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with MillinewtonMeterPerDegree unit
    millinewton_meters_per_degreeJSON := []byte(`{"value": 100, "unit": "MillinewtonMeterPerDegree"}`)
    millinewton_meters_per_degreeResult, err := factory.FromDtoJSON(millinewton_meters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillinewtonMeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewton_meters_per_degreeResult.Convert(units.RotationalStiffnessMillinewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with CentinewtonMeterPerDegree unit
    centinewton_meters_per_degreeJSON := []byte(`{"value": 100, "unit": "CentinewtonMeterPerDegree"}`)
    centinewton_meters_per_degreeResult, err := factory.FromDtoJSON(centinewton_meters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentinewtonMeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewton_meters_per_degreeResult.Convert(units.RotationalStiffnessCentinewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with DecinewtonMeterPerDegree unit
    decinewton_meters_per_degreeJSON := []byte(`{"value": 100, "unit": "DecinewtonMeterPerDegree"}`)
    decinewton_meters_per_degreeResult, err := factory.FromDtoJSON(decinewton_meters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecinewtonMeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewton_meters_per_degreeResult.Convert(units.RotationalStiffnessDecinewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with DecanewtonMeterPerDegree unit
    decanewton_meters_per_degreeJSON := []byte(`{"value": 100, "unit": "DecanewtonMeterPerDegree"}`)
    decanewton_meters_per_degreeResult, err := factory.FromDtoJSON(decanewton_meters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecanewtonMeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewton_meters_per_degreeResult.Convert(units.RotationalStiffnessDecanewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonMeterPerDegree unit
    kilonewton_meters_per_degreeJSON := []byte(`{"value": 100, "unit": "KilonewtonMeterPerDegree"}`)
    kilonewton_meters_per_degreeResult, err := factory.FromDtoJSON(kilonewton_meters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonMeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_meters_per_degreeResult.Convert(units.RotationalStiffnessKilonewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonMeterPerDegree unit
    meganewton_meters_per_degreeJSON := []byte(`{"value": 100, "unit": "MeganewtonMeterPerDegree"}`)
    meganewton_meters_per_degreeResult, err := factory.FromDtoJSON(meganewton_meters_per_degreeJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonMeterPerDegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_meters_per_degreeResult.Convert(units.RotationalStiffnessMeganewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMeterPerDegree = %v, want %v", converted, 100)
    }
    // Test JSON with NanonewtonMillimeterPerRadian unit
    nanonewton_millimeters_per_radianJSON := []byte(`{"value": 100, "unit": "NanonewtonMillimeterPerRadian"}`)
    nanonewton_millimeters_per_radianResult, err := factory.FromDtoJSON(nanonewton_millimeters_per_radianJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanonewtonMillimeterPerRadian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessNanonewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test JSON with MicronewtonMillimeterPerRadian unit
    micronewton_millimeters_per_radianJSON := []byte(`{"value": 100, "unit": "MicronewtonMillimeterPerRadian"}`)
    micronewton_millimeters_per_radianResult, err := factory.FromDtoJSON(micronewton_millimeters_per_radianJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicronewtonMillimeterPerRadian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessMicronewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test JSON with MillinewtonMillimeterPerRadian unit
    millinewton_millimeters_per_radianJSON := []byte(`{"value": 100, "unit": "MillinewtonMillimeterPerRadian"}`)
    millinewton_millimeters_per_radianResult, err := factory.FromDtoJSON(millinewton_millimeters_per_radianJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillinewtonMillimeterPerRadian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessMillinewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test JSON with CentinewtonMillimeterPerRadian unit
    centinewton_millimeters_per_radianJSON := []byte(`{"value": 100, "unit": "CentinewtonMillimeterPerRadian"}`)
    centinewton_millimeters_per_radianResult, err := factory.FromDtoJSON(centinewton_millimeters_per_radianJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentinewtonMillimeterPerRadian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessCentinewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test JSON with DecinewtonMillimeterPerRadian unit
    decinewton_millimeters_per_radianJSON := []byte(`{"value": 100, "unit": "DecinewtonMillimeterPerRadian"}`)
    decinewton_millimeters_per_radianResult, err := factory.FromDtoJSON(decinewton_millimeters_per_radianJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecinewtonMillimeterPerRadian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessDecinewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test JSON with DecanewtonMillimeterPerRadian unit
    decanewton_millimeters_per_radianJSON := []byte(`{"value": 100, "unit": "DecanewtonMillimeterPerRadian"}`)
    decanewton_millimeters_per_radianResult, err := factory.FromDtoJSON(decanewton_millimeters_per_radianJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecanewtonMillimeterPerRadian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessDecanewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonMillimeterPerRadian unit
    kilonewton_millimeters_per_radianJSON := []byte(`{"value": 100, "unit": "KilonewtonMillimeterPerRadian"}`)
    kilonewton_millimeters_per_radianResult, err := factory.FromDtoJSON(kilonewton_millimeters_per_radianJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonMillimeterPerRadian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessKilonewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonMillimeterPerRadian unit
    meganewton_millimeters_per_radianJSON := []byte(`{"value": 100, "unit": "MeganewtonMillimeterPerRadian"}`)
    meganewton_millimeters_per_radianResult, err := factory.FromDtoJSON(meganewton_millimeters_per_radianJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonMillimeterPerRadian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_millimeters_per_radianResult.Convert(units.RotationalStiffnessMeganewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMillimeterPerRadian = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "NewtonMeterPerRadian"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromNewtonMetersPerRadian function
func TestRotationalStiffnessFactory_FromNewtonMetersPerRadian(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonMetersPerRadian(100)
    if err != nil {
        t.Errorf("FromNewtonMetersPerRadian() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessNewtonMeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonMetersPerRadian() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonMetersPerRadian(math.NaN())
    if err == nil {
        t.Error("FromNewtonMetersPerRadian() with NaN value should return error")
    }

    _, err = factory.FromNewtonMetersPerRadian(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonMetersPerRadian() with +Inf value should return error")
    }

    _, err = factory.FromNewtonMetersPerRadian(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonMetersPerRadian() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonMetersPerRadian(0)
    if err != nil {
        t.Errorf("FromNewtonMetersPerRadian() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessNewtonMeterPerRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonMetersPerRadian() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundForceFeetPerDegrees function
func TestRotationalStiffnessFactory_FromPoundForceFeetPerDegrees(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundForceFeetPerDegrees(100)
    if err != nil {
        t.Errorf("FromPoundForceFeetPerDegrees() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessPoundForceFootPerDegrees)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundForceFeetPerDegrees() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundForceFeetPerDegrees(math.NaN())
    if err == nil {
        t.Error("FromPoundForceFeetPerDegrees() with NaN value should return error")
    }

    _, err = factory.FromPoundForceFeetPerDegrees(math.Inf(1))
    if err == nil {
        t.Error("FromPoundForceFeetPerDegrees() with +Inf value should return error")
    }

    _, err = factory.FromPoundForceFeetPerDegrees(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundForceFeetPerDegrees() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundForceFeetPerDegrees(0)
    if err != nil {
        t.Errorf("FromPoundForceFeetPerDegrees() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessPoundForceFootPerDegrees)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundForceFeetPerDegrees() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundForceFeetPerDegrees function
func TestRotationalStiffnessFactory_FromKilopoundForceFeetPerDegrees(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundForceFeetPerDegrees(100)
    if err != nil {
        t.Errorf("FromKilopoundForceFeetPerDegrees() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessKilopoundForceFootPerDegrees)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundForceFeetPerDegrees() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundForceFeetPerDegrees(math.NaN())
    if err == nil {
        t.Error("FromKilopoundForceFeetPerDegrees() with NaN value should return error")
    }

    _, err = factory.FromKilopoundForceFeetPerDegrees(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundForceFeetPerDegrees() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundForceFeetPerDegrees(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundForceFeetPerDegrees() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundForceFeetPerDegrees(0)
    if err != nil {
        t.Errorf("FromKilopoundForceFeetPerDegrees() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessKilopoundForceFootPerDegrees)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundForceFeetPerDegrees() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonMillimetersPerDegree function
func TestRotationalStiffnessFactory_FromNewtonMillimetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonMillimetersPerDegree(100)
    if err != nil {
        t.Errorf("FromNewtonMillimetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessNewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonMillimetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonMillimetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromNewtonMillimetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromNewtonMillimetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonMillimetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromNewtonMillimetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonMillimetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonMillimetersPerDegree(0)
    if err != nil {
        t.Errorf("FromNewtonMillimetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessNewtonMillimeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonMillimetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonMetersPerDegree function
func TestRotationalStiffnessFactory_FromNewtonMetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonMetersPerDegree(100)
    if err != nil {
        t.Errorf("FromNewtonMetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessNewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonMetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonMetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromNewtonMetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromNewtonMetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonMetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromNewtonMetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonMetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonMetersPerDegree(0)
    if err != nil {
        t.Errorf("FromNewtonMetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessNewtonMeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonMetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonMillimetersPerRadian function
func TestRotationalStiffnessFactory_FromNewtonMillimetersPerRadian(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonMillimetersPerRadian(100)
    if err != nil {
        t.Errorf("FromNewtonMillimetersPerRadian() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessNewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonMillimetersPerRadian() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonMillimetersPerRadian(math.NaN())
    if err == nil {
        t.Error("FromNewtonMillimetersPerRadian() with NaN value should return error")
    }

    _, err = factory.FromNewtonMillimetersPerRadian(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonMillimetersPerRadian() with +Inf value should return error")
    }

    _, err = factory.FromNewtonMillimetersPerRadian(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonMillimetersPerRadian() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonMillimetersPerRadian(0)
    if err != nil {
        t.Errorf("FromNewtonMillimetersPerRadian() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessNewtonMillimeterPerRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonMillimetersPerRadian() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundForceFeetPerRadian function
func TestRotationalStiffnessFactory_FromPoundForceFeetPerRadian(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundForceFeetPerRadian(100)
    if err != nil {
        t.Errorf("FromPoundForceFeetPerRadian() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessPoundForceFeetPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundForceFeetPerRadian() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundForceFeetPerRadian(math.NaN())
    if err == nil {
        t.Error("FromPoundForceFeetPerRadian() with NaN value should return error")
    }

    _, err = factory.FromPoundForceFeetPerRadian(math.Inf(1))
    if err == nil {
        t.Error("FromPoundForceFeetPerRadian() with +Inf value should return error")
    }

    _, err = factory.FromPoundForceFeetPerRadian(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundForceFeetPerRadian() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundForceFeetPerRadian(0)
    if err != nil {
        t.Errorf("FromPoundForceFeetPerRadian() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessPoundForceFeetPerRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundForceFeetPerRadian() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonMetersPerRadian function
func TestRotationalStiffnessFactory_FromKilonewtonMetersPerRadian(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonMetersPerRadian(100)
    if err != nil {
        t.Errorf("FromKilonewtonMetersPerRadian() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessKilonewtonMeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonMetersPerRadian() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonMetersPerRadian(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonMetersPerRadian() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonMetersPerRadian(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonMetersPerRadian() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonMetersPerRadian(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonMetersPerRadian() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonMetersPerRadian(0)
    if err != nil {
        t.Errorf("FromKilonewtonMetersPerRadian() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessKilonewtonMeterPerRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonMetersPerRadian() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonMetersPerRadian function
func TestRotationalStiffnessFactory_FromMeganewtonMetersPerRadian(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonMetersPerRadian(100)
    if err != nil {
        t.Errorf("FromMeganewtonMetersPerRadian() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessMeganewtonMeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonMetersPerRadian() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonMetersPerRadian(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonMetersPerRadian() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonMetersPerRadian(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonMetersPerRadian() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonMetersPerRadian(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonMetersPerRadian() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonMetersPerRadian(0)
    if err != nil {
        t.Errorf("FromMeganewtonMetersPerRadian() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessMeganewtonMeterPerRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonMetersPerRadian() with zero value = %v, want 0", converted)
    }
}
// Test FromNanonewtonMillimetersPerDegree function
func TestRotationalStiffnessFactory_FromNanonewtonMillimetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanonewtonMillimetersPerDegree(100)
    if err != nil {
        t.Errorf("FromNanonewtonMillimetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessNanonewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanonewtonMillimetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanonewtonMillimetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromNanonewtonMillimetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromNanonewtonMillimetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromNanonewtonMillimetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromNanonewtonMillimetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromNanonewtonMillimetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanonewtonMillimetersPerDegree(0)
    if err != nil {
        t.Errorf("FromNanonewtonMillimetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessNanonewtonMillimeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanonewtonMillimetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromMicronewtonMillimetersPerDegree function
func TestRotationalStiffnessFactory_FromMicronewtonMillimetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicronewtonMillimetersPerDegree(100)
    if err != nil {
        t.Errorf("FromMicronewtonMillimetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessMicronewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicronewtonMillimetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicronewtonMillimetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromMicronewtonMillimetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromMicronewtonMillimetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromMicronewtonMillimetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromMicronewtonMillimetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromMicronewtonMillimetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicronewtonMillimetersPerDegree(0)
    if err != nil {
        t.Errorf("FromMicronewtonMillimetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessMicronewtonMillimeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicronewtonMillimetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromMillinewtonMillimetersPerDegree function
func TestRotationalStiffnessFactory_FromMillinewtonMillimetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillinewtonMillimetersPerDegree(100)
    if err != nil {
        t.Errorf("FromMillinewtonMillimetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessMillinewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillinewtonMillimetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillinewtonMillimetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromMillinewtonMillimetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromMillinewtonMillimetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromMillinewtonMillimetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromMillinewtonMillimetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromMillinewtonMillimetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillinewtonMillimetersPerDegree(0)
    if err != nil {
        t.Errorf("FromMillinewtonMillimetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessMillinewtonMillimeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillinewtonMillimetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromCentinewtonMillimetersPerDegree function
func TestRotationalStiffnessFactory_FromCentinewtonMillimetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentinewtonMillimetersPerDegree(100)
    if err != nil {
        t.Errorf("FromCentinewtonMillimetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessCentinewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentinewtonMillimetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentinewtonMillimetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromCentinewtonMillimetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromCentinewtonMillimetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromCentinewtonMillimetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromCentinewtonMillimetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromCentinewtonMillimetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentinewtonMillimetersPerDegree(0)
    if err != nil {
        t.Errorf("FromCentinewtonMillimetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessCentinewtonMillimeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentinewtonMillimetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromDecinewtonMillimetersPerDegree function
func TestRotationalStiffnessFactory_FromDecinewtonMillimetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecinewtonMillimetersPerDegree(100)
    if err != nil {
        t.Errorf("FromDecinewtonMillimetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessDecinewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecinewtonMillimetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecinewtonMillimetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromDecinewtonMillimetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromDecinewtonMillimetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromDecinewtonMillimetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromDecinewtonMillimetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromDecinewtonMillimetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecinewtonMillimetersPerDegree(0)
    if err != nil {
        t.Errorf("FromDecinewtonMillimetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessDecinewtonMillimeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecinewtonMillimetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromDecanewtonMillimetersPerDegree function
func TestRotationalStiffnessFactory_FromDecanewtonMillimetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecanewtonMillimetersPerDegree(100)
    if err != nil {
        t.Errorf("FromDecanewtonMillimetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessDecanewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecanewtonMillimetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecanewtonMillimetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromDecanewtonMillimetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromDecanewtonMillimetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromDecanewtonMillimetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromDecanewtonMillimetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromDecanewtonMillimetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecanewtonMillimetersPerDegree(0)
    if err != nil {
        t.Errorf("FromDecanewtonMillimetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessDecanewtonMillimeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecanewtonMillimetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonMillimetersPerDegree function
func TestRotationalStiffnessFactory_FromKilonewtonMillimetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonMillimetersPerDegree(100)
    if err != nil {
        t.Errorf("FromKilonewtonMillimetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessKilonewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonMillimetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonMillimetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonMillimetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonMillimetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonMillimetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonMillimetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonMillimetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonMillimetersPerDegree(0)
    if err != nil {
        t.Errorf("FromKilonewtonMillimetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessKilonewtonMillimeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonMillimetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonMillimetersPerDegree function
func TestRotationalStiffnessFactory_FromMeganewtonMillimetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonMillimetersPerDegree(100)
    if err != nil {
        t.Errorf("FromMeganewtonMillimetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessMeganewtonMillimeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonMillimetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonMillimetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonMillimetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonMillimetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonMillimetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonMillimetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonMillimetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonMillimetersPerDegree(0)
    if err != nil {
        t.Errorf("FromMeganewtonMillimetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessMeganewtonMillimeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonMillimetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromNanonewtonMetersPerDegree function
func TestRotationalStiffnessFactory_FromNanonewtonMetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanonewtonMetersPerDegree(100)
    if err != nil {
        t.Errorf("FromNanonewtonMetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessNanonewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanonewtonMetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanonewtonMetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromNanonewtonMetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromNanonewtonMetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromNanonewtonMetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromNanonewtonMetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromNanonewtonMetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanonewtonMetersPerDegree(0)
    if err != nil {
        t.Errorf("FromNanonewtonMetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessNanonewtonMeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanonewtonMetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromMicronewtonMetersPerDegree function
func TestRotationalStiffnessFactory_FromMicronewtonMetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicronewtonMetersPerDegree(100)
    if err != nil {
        t.Errorf("FromMicronewtonMetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessMicronewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicronewtonMetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicronewtonMetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromMicronewtonMetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromMicronewtonMetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromMicronewtonMetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromMicronewtonMetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromMicronewtonMetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicronewtonMetersPerDegree(0)
    if err != nil {
        t.Errorf("FromMicronewtonMetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessMicronewtonMeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicronewtonMetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromMillinewtonMetersPerDegree function
func TestRotationalStiffnessFactory_FromMillinewtonMetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillinewtonMetersPerDegree(100)
    if err != nil {
        t.Errorf("FromMillinewtonMetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessMillinewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillinewtonMetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillinewtonMetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromMillinewtonMetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromMillinewtonMetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromMillinewtonMetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromMillinewtonMetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromMillinewtonMetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillinewtonMetersPerDegree(0)
    if err != nil {
        t.Errorf("FromMillinewtonMetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessMillinewtonMeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillinewtonMetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromCentinewtonMetersPerDegree function
func TestRotationalStiffnessFactory_FromCentinewtonMetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentinewtonMetersPerDegree(100)
    if err != nil {
        t.Errorf("FromCentinewtonMetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessCentinewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentinewtonMetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentinewtonMetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromCentinewtonMetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromCentinewtonMetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromCentinewtonMetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromCentinewtonMetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromCentinewtonMetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentinewtonMetersPerDegree(0)
    if err != nil {
        t.Errorf("FromCentinewtonMetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessCentinewtonMeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentinewtonMetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromDecinewtonMetersPerDegree function
func TestRotationalStiffnessFactory_FromDecinewtonMetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecinewtonMetersPerDegree(100)
    if err != nil {
        t.Errorf("FromDecinewtonMetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessDecinewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecinewtonMetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecinewtonMetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromDecinewtonMetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromDecinewtonMetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromDecinewtonMetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromDecinewtonMetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromDecinewtonMetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecinewtonMetersPerDegree(0)
    if err != nil {
        t.Errorf("FromDecinewtonMetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessDecinewtonMeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecinewtonMetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromDecanewtonMetersPerDegree function
func TestRotationalStiffnessFactory_FromDecanewtonMetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecanewtonMetersPerDegree(100)
    if err != nil {
        t.Errorf("FromDecanewtonMetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessDecanewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecanewtonMetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecanewtonMetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromDecanewtonMetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromDecanewtonMetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromDecanewtonMetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromDecanewtonMetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromDecanewtonMetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecanewtonMetersPerDegree(0)
    if err != nil {
        t.Errorf("FromDecanewtonMetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessDecanewtonMeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecanewtonMetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonMetersPerDegree function
func TestRotationalStiffnessFactory_FromKilonewtonMetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonMetersPerDegree(100)
    if err != nil {
        t.Errorf("FromKilonewtonMetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessKilonewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonMetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonMetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonMetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonMetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonMetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonMetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonMetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonMetersPerDegree(0)
    if err != nil {
        t.Errorf("FromKilonewtonMetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessKilonewtonMeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonMetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonMetersPerDegree function
func TestRotationalStiffnessFactory_FromMeganewtonMetersPerDegree(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonMetersPerDegree(100)
    if err != nil {
        t.Errorf("FromMeganewtonMetersPerDegree() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessMeganewtonMeterPerDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonMetersPerDegree() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonMetersPerDegree(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonMetersPerDegree() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonMetersPerDegree(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonMetersPerDegree() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonMetersPerDegree(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonMetersPerDegree() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonMetersPerDegree(0)
    if err != nil {
        t.Errorf("FromMeganewtonMetersPerDegree() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessMeganewtonMeterPerDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonMetersPerDegree() with zero value = %v, want 0", converted)
    }
}
// Test FromNanonewtonMillimetersPerRadian function
func TestRotationalStiffnessFactory_FromNanonewtonMillimetersPerRadian(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanonewtonMillimetersPerRadian(100)
    if err != nil {
        t.Errorf("FromNanonewtonMillimetersPerRadian() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessNanonewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanonewtonMillimetersPerRadian() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanonewtonMillimetersPerRadian(math.NaN())
    if err == nil {
        t.Error("FromNanonewtonMillimetersPerRadian() with NaN value should return error")
    }

    _, err = factory.FromNanonewtonMillimetersPerRadian(math.Inf(1))
    if err == nil {
        t.Error("FromNanonewtonMillimetersPerRadian() with +Inf value should return error")
    }

    _, err = factory.FromNanonewtonMillimetersPerRadian(math.Inf(-1))
    if err == nil {
        t.Error("FromNanonewtonMillimetersPerRadian() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanonewtonMillimetersPerRadian(0)
    if err != nil {
        t.Errorf("FromNanonewtonMillimetersPerRadian() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessNanonewtonMillimeterPerRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanonewtonMillimetersPerRadian() with zero value = %v, want 0", converted)
    }
}
// Test FromMicronewtonMillimetersPerRadian function
func TestRotationalStiffnessFactory_FromMicronewtonMillimetersPerRadian(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicronewtonMillimetersPerRadian(100)
    if err != nil {
        t.Errorf("FromMicronewtonMillimetersPerRadian() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessMicronewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicronewtonMillimetersPerRadian() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicronewtonMillimetersPerRadian(math.NaN())
    if err == nil {
        t.Error("FromMicronewtonMillimetersPerRadian() with NaN value should return error")
    }

    _, err = factory.FromMicronewtonMillimetersPerRadian(math.Inf(1))
    if err == nil {
        t.Error("FromMicronewtonMillimetersPerRadian() with +Inf value should return error")
    }

    _, err = factory.FromMicronewtonMillimetersPerRadian(math.Inf(-1))
    if err == nil {
        t.Error("FromMicronewtonMillimetersPerRadian() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicronewtonMillimetersPerRadian(0)
    if err != nil {
        t.Errorf("FromMicronewtonMillimetersPerRadian() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessMicronewtonMillimeterPerRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicronewtonMillimetersPerRadian() with zero value = %v, want 0", converted)
    }
}
// Test FromMillinewtonMillimetersPerRadian function
func TestRotationalStiffnessFactory_FromMillinewtonMillimetersPerRadian(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillinewtonMillimetersPerRadian(100)
    if err != nil {
        t.Errorf("FromMillinewtonMillimetersPerRadian() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessMillinewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillinewtonMillimetersPerRadian() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillinewtonMillimetersPerRadian(math.NaN())
    if err == nil {
        t.Error("FromMillinewtonMillimetersPerRadian() with NaN value should return error")
    }

    _, err = factory.FromMillinewtonMillimetersPerRadian(math.Inf(1))
    if err == nil {
        t.Error("FromMillinewtonMillimetersPerRadian() with +Inf value should return error")
    }

    _, err = factory.FromMillinewtonMillimetersPerRadian(math.Inf(-1))
    if err == nil {
        t.Error("FromMillinewtonMillimetersPerRadian() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillinewtonMillimetersPerRadian(0)
    if err != nil {
        t.Errorf("FromMillinewtonMillimetersPerRadian() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessMillinewtonMillimeterPerRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillinewtonMillimetersPerRadian() with zero value = %v, want 0", converted)
    }
}
// Test FromCentinewtonMillimetersPerRadian function
func TestRotationalStiffnessFactory_FromCentinewtonMillimetersPerRadian(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentinewtonMillimetersPerRadian(100)
    if err != nil {
        t.Errorf("FromCentinewtonMillimetersPerRadian() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessCentinewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentinewtonMillimetersPerRadian() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentinewtonMillimetersPerRadian(math.NaN())
    if err == nil {
        t.Error("FromCentinewtonMillimetersPerRadian() with NaN value should return error")
    }

    _, err = factory.FromCentinewtonMillimetersPerRadian(math.Inf(1))
    if err == nil {
        t.Error("FromCentinewtonMillimetersPerRadian() with +Inf value should return error")
    }

    _, err = factory.FromCentinewtonMillimetersPerRadian(math.Inf(-1))
    if err == nil {
        t.Error("FromCentinewtonMillimetersPerRadian() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentinewtonMillimetersPerRadian(0)
    if err != nil {
        t.Errorf("FromCentinewtonMillimetersPerRadian() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessCentinewtonMillimeterPerRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentinewtonMillimetersPerRadian() with zero value = %v, want 0", converted)
    }
}
// Test FromDecinewtonMillimetersPerRadian function
func TestRotationalStiffnessFactory_FromDecinewtonMillimetersPerRadian(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecinewtonMillimetersPerRadian(100)
    if err != nil {
        t.Errorf("FromDecinewtonMillimetersPerRadian() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessDecinewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecinewtonMillimetersPerRadian() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecinewtonMillimetersPerRadian(math.NaN())
    if err == nil {
        t.Error("FromDecinewtonMillimetersPerRadian() with NaN value should return error")
    }

    _, err = factory.FromDecinewtonMillimetersPerRadian(math.Inf(1))
    if err == nil {
        t.Error("FromDecinewtonMillimetersPerRadian() with +Inf value should return error")
    }

    _, err = factory.FromDecinewtonMillimetersPerRadian(math.Inf(-1))
    if err == nil {
        t.Error("FromDecinewtonMillimetersPerRadian() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecinewtonMillimetersPerRadian(0)
    if err != nil {
        t.Errorf("FromDecinewtonMillimetersPerRadian() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessDecinewtonMillimeterPerRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecinewtonMillimetersPerRadian() with zero value = %v, want 0", converted)
    }
}
// Test FromDecanewtonMillimetersPerRadian function
func TestRotationalStiffnessFactory_FromDecanewtonMillimetersPerRadian(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecanewtonMillimetersPerRadian(100)
    if err != nil {
        t.Errorf("FromDecanewtonMillimetersPerRadian() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessDecanewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecanewtonMillimetersPerRadian() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecanewtonMillimetersPerRadian(math.NaN())
    if err == nil {
        t.Error("FromDecanewtonMillimetersPerRadian() with NaN value should return error")
    }

    _, err = factory.FromDecanewtonMillimetersPerRadian(math.Inf(1))
    if err == nil {
        t.Error("FromDecanewtonMillimetersPerRadian() with +Inf value should return error")
    }

    _, err = factory.FromDecanewtonMillimetersPerRadian(math.Inf(-1))
    if err == nil {
        t.Error("FromDecanewtonMillimetersPerRadian() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecanewtonMillimetersPerRadian(0)
    if err != nil {
        t.Errorf("FromDecanewtonMillimetersPerRadian() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessDecanewtonMillimeterPerRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecanewtonMillimetersPerRadian() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonMillimetersPerRadian function
func TestRotationalStiffnessFactory_FromKilonewtonMillimetersPerRadian(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonMillimetersPerRadian(100)
    if err != nil {
        t.Errorf("FromKilonewtonMillimetersPerRadian() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessKilonewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonMillimetersPerRadian() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonMillimetersPerRadian(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonMillimetersPerRadian() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonMillimetersPerRadian(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonMillimetersPerRadian() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonMillimetersPerRadian(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonMillimetersPerRadian() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonMillimetersPerRadian(0)
    if err != nil {
        t.Errorf("FromKilonewtonMillimetersPerRadian() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessKilonewtonMillimeterPerRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonMillimetersPerRadian() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonMillimetersPerRadian function
func TestRotationalStiffnessFactory_FromMeganewtonMillimetersPerRadian(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonMillimetersPerRadian(100)
    if err != nil {
        t.Errorf("FromMeganewtonMillimetersPerRadian() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalStiffnessMeganewtonMillimeterPerRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonMillimetersPerRadian() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonMillimetersPerRadian(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonMillimetersPerRadian() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonMillimetersPerRadian(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonMillimetersPerRadian() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonMillimetersPerRadian(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonMillimetersPerRadian() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonMillimetersPerRadian(0)
    if err != nil {
        t.Errorf("FromMeganewtonMillimetersPerRadian() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalStiffnessMeganewtonMillimeterPerRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonMillimetersPerRadian() with zero value = %v, want 0", converted)
    }
}

func TestRotationalStiffnessToString(t *testing.T) {
	factory := units.RotationalStiffnessFactory{}
	a, err := factory.CreateRotationalStiffness(45, units.RotationalStiffnessNewtonMeterPerRadian)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RotationalStiffnessNewtonMeterPerRadian, 2)
	expected := "45.00 " + units.GetRotationalStiffnessAbbreviation(units.RotationalStiffnessNewtonMeterPerRadian)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RotationalStiffnessNewtonMeterPerRadian, -1)
	expected = "45 " + units.GetRotationalStiffnessAbbreviation(units.RotationalStiffnessNewtonMeterPerRadian)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRotationalStiffness_EqualityAndComparison(t *testing.T) {
	factory := units.RotationalStiffnessFactory{}
	a1, _ := factory.CreateRotationalStiffness(60, units.RotationalStiffnessNewtonMeterPerRadian)
	a2, _ := factory.CreateRotationalStiffness(60, units.RotationalStiffnessNewtonMeterPerRadian)
	a3, _ := factory.CreateRotationalStiffness(120, units.RotationalStiffnessNewtonMeterPerRadian)

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

func TestRotationalStiffness_Arithmetic(t *testing.T) {
	factory := units.RotationalStiffnessFactory{}
	a1, _ := factory.CreateRotationalStiffness(30, units.RotationalStiffnessNewtonMeterPerRadian)
	a2, _ := factory.CreateRotationalStiffness(45, units.RotationalStiffnessNewtonMeterPerRadian)

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


func TestGetRotationalStiffnessAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.RotationalStiffnessUnits
        want string
    }{
        {
            name: "NewtonMeterPerRadian abbreviation",
            unit: units.RotationalStiffnessNewtonMeterPerRadian,
            want: "N·m/rad",
        },
        {
            name: "PoundForceFootPerDegrees abbreviation",
            unit: units.RotationalStiffnessPoundForceFootPerDegrees,
            want: "lbf·ft/deg",
        },
        {
            name: "KilopoundForceFootPerDegrees abbreviation",
            unit: units.RotationalStiffnessKilopoundForceFootPerDegrees,
            want: "kipf·ft/°",
        },
        {
            name: "NewtonMillimeterPerDegree abbreviation",
            unit: units.RotationalStiffnessNewtonMillimeterPerDegree,
            want: "N·mm/deg",
        },
        {
            name: "NewtonMeterPerDegree abbreviation",
            unit: units.RotationalStiffnessNewtonMeterPerDegree,
            want: "N·m/deg",
        },
        {
            name: "NewtonMillimeterPerRadian abbreviation",
            unit: units.RotationalStiffnessNewtonMillimeterPerRadian,
            want: "N·mm/rad",
        },
        {
            name: "PoundForceFeetPerRadian abbreviation",
            unit: units.RotationalStiffnessPoundForceFeetPerRadian,
            want: "lbf·ft/rad",
        },
        {
            name: "KilonewtonMeterPerRadian abbreviation",
            unit: units.RotationalStiffnessKilonewtonMeterPerRadian,
            want: "kN·m/rad",
        },
        {
            name: "MeganewtonMeterPerRadian abbreviation",
            unit: units.RotationalStiffnessMeganewtonMeterPerRadian,
            want: "MN·m/rad",
        },
        {
            name: "NanonewtonMillimeterPerDegree abbreviation",
            unit: units.RotationalStiffnessNanonewtonMillimeterPerDegree,
            want: "nN·mm/deg",
        },
        {
            name: "MicronewtonMillimeterPerDegree abbreviation",
            unit: units.RotationalStiffnessMicronewtonMillimeterPerDegree,
            want: "μN·mm/deg",
        },
        {
            name: "MillinewtonMillimeterPerDegree abbreviation",
            unit: units.RotationalStiffnessMillinewtonMillimeterPerDegree,
            want: "mN·mm/deg",
        },
        {
            name: "CentinewtonMillimeterPerDegree abbreviation",
            unit: units.RotationalStiffnessCentinewtonMillimeterPerDegree,
            want: "cN·mm/deg",
        },
        {
            name: "DecinewtonMillimeterPerDegree abbreviation",
            unit: units.RotationalStiffnessDecinewtonMillimeterPerDegree,
            want: "dN·mm/deg",
        },
        {
            name: "DecanewtonMillimeterPerDegree abbreviation",
            unit: units.RotationalStiffnessDecanewtonMillimeterPerDegree,
            want: "daN·mm/deg",
        },
        {
            name: "KilonewtonMillimeterPerDegree abbreviation",
            unit: units.RotationalStiffnessKilonewtonMillimeterPerDegree,
            want: "kN·mm/deg",
        },
        {
            name: "MeganewtonMillimeterPerDegree abbreviation",
            unit: units.RotationalStiffnessMeganewtonMillimeterPerDegree,
            want: "MN·mm/deg",
        },
        {
            name: "NanonewtonMeterPerDegree abbreviation",
            unit: units.RotationalStiffnessNanonewtonMeterPerDegree,
            want: "nN·m/deg",
        },
        {
            name: "MicronewtonMeterPerDegree abbreviation",
            unit: units.RotationalStiffnessMicronewtonMeterPerDegree,
            want: "μN·m/deg",
        },
        {
            name: "MillinewtonMeterPerDegree abbreviation",
            unit: units.RotationalStiffnessMillinewtonMeterPerDegree,
            want: "mN·m/deg",
        },
        {
            name: "CentinewtonMeterPerDegree abbreviation",
            unit: units.RotationalStiffnessCentinewtonMeterPerDegree,
            want: "cN·m/deg",
        },
        {
            name: "DecinewtonMeterPerDegree abbreviation",
            unit: units.RotationalStiffnessDecinewtonMeterPerDegree,
            want: "dN·m/deg",
        },
        {
            name: "DecanewtonMeterPerDegree abbreviation",
            unit: units.RotationalStiffnessDecanewtonMeterPerDegree,
            want: "daN·m/deg",
        },
        {
            name: "KilonewtonMeterPerDegree abbreviation",
            unit: units.RotationalStiffnessKilonewtonMeterPerDegree,
            want: "kN·m/deg",
        },
        {
            name: "MeganewtonMeterPerDegree abbreviation",
            unit: units.RotationalStiffnessMeganewtonMeterPerDegree,
            want: "MN·m/deg",
        },
        {
            name: "NanonewtonMillimeterPerRadian abbreviation",
            unit: units.RotationalStiffnessNanonewtonMillimeterPerRadian,
            want: "nN·mm/rad",
        },
        {
            name: "MicronewtonMillimeterPerRadian abbreviation",
            unit: units.RotationalStiffnessMicronewtonMillimeterPerRadian,
            want: "μN·mm/rad",
        },
        {
            name: "MillinewtonMillimeterPerRadian abbreviation",
            unit: units.RotationalStiffnessMillinewtonMillimeterPerRadian,
            want: "mN·mm/rad",
        },
        {
            name: "CentinewtonMillimeterPerRadian abbreviation",
            unit: units.RotationalStiffnessCentinewtonMillimeterPerRadian,
            want: "cN·mm/rad",
        },
        {
            name: "DecinewtonMillimeterPerRadian abbreviation",
            unit: units.RotationalStiffnessDecinewtonMillimeterPerRadian,
            want: "dN·mm/rad",
        },
        {
            name: "DecanewtonMillimeterPerRadian abbreviation",
            unit: units.RotationalStiffnessDecanewtonMillimeterPerRadian,
            want: "daN·mm/rad",
        },
        {
            name: "KilonewtonMillimeterPerRadian abbreviation",
            unit: units.RotationalStiffnessKilonewtonMillimeterPerRadian,
            want: "kN·mm/rad",
        },
        {
            name: "MeganewtonMillimeterPerRadian abbreviation",
            unit: units.RotationalStiffnessMeganewtonMillimeterPerRadian,
            want: "MN·mm/rad",
        },
        {
            name: "invalid unit",
            unit: units.RotationalStiffnessUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetRotationalStiffnessAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetRotationalStiffnessAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestRotationalStiffness_String(t *testing.T) {
    factory := units.RotationalStiffnessFactory{}
    
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
            unit, err := factory.CreateRotationalStiffness(tt.value, units.RotationalStiffnessNewtonMeterPerRadian)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("RotationalStiffness.String() = %v, want %v", got, tt.want)
            }
        })
    }
}