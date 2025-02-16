// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestSpecificWeightDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonPerCubicMeter"}`
	
	factory := units.SpecificWeightDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.SpecificWeightNewtonPerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.SpecificWeightNewtonPerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonPerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestSpecificWeightDto_ToJSON(t *testing.T) {
	dto := units.SpecificWeightDto{
		Value: 45,
		Unit:  units.SpecificWeightNewtonPerCubicMeter,
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
	if result["unit"].(string) != string(units.SpecificWeightNewtonPerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.SpecificWeightNewtonPerCubicMeter, result["unit"])
	}
}

func TestNewSpecificWeight_InvalidValue(t *testing.T) {
	factory := units.SpecificWeightFactory{}
	// NaN value should return an error.
	_, err := factory.CreateSpecificWeight(math.NaN(), units.SpecificWeightNewtonPerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateSpecificWeight(math.Inf(1), units.SpecificWeightNewtonPerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestSpecificWeightConversions(t *testing.T) {
	factory := units.SpecificWeightFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateSpecificWeight(180, units.SpecificWeightNewtonPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonsPerCubicMillimeter.
		// No expected conversion value provided for NewtonsPerCubicMillimeter, verifying result is not NaN.
		result := a.NewtonsPerCubicMillimeter()
		cacheResult := a.NewtonsPerCubicMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerCubicCentimeter.
		// No expected conversion value provided for NewtonsPerCubicCentimeter, verifying result is not NaN.
		result := a.NewtonsPerCubicCentimeter()
		cacheResult := a.NewtonsPerCubicCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerCubicMeter.
		// No expected conversion value provided for NewtonsPerCubicMeter, verifying result is not NaN.
		result := a.NewtonsPerCubicMeter()
		cacheResult := a.NewtonsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerCubicMillimeter.
		// No expected conversion value provided for KilogramsForcePerCubicMillimeter, verifying result is not NaN.
		result := a.KilogramsForcePerCubicMillimeter()
		cacheResult := a.KilogramsForcePerCubicMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsForcePerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerCubicCentimeter.
		// No expected conversion value provided for KilogramsForcePerCubicCentimeter, verifying result is not NaN.
		result := a.KilogramsForcePerCubicCentimeter()
		cacheResult := a.KilogramsForcePerCubicCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsForcePerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerCubicMeter.
		// No expected conversion value provided for KilogramsForcePerCubicMeter, verifying result is not NaN.
		result := a.KilogramsForcePerCubicMeter()
		cacheResult := a.KilogramsForcePerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsForcePerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerCubicInch.
		// No expected conversion value provided for PoundsForcePerCubicInch, verifying result is not NaN.
		result := a.PoundsForcePerCubicInch()
		cacheResult := a.PoundsForcePerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsForcePerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerCubicFoot.
		// No expected conversion value provided for PoundsForcePerCubicFoot, verifying result is not NaN.
		result := a.PoundsForcePerCubicFoot()
		cacheResult := a.PoundsForcePerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsForcePerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerCubicMillimeter.
		// No expected conversion value provided for TonnesForcePerCubicMillimeter, verifying result is not NaN.
		result := a.TonnesForcePerCubicMillimeter()
		cacheResult := a.TonnesForcePerCubicMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonnesForcePerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerCubicCentimeter.
		// No expected conversion value provided for TonnesForcePerCubicCentimeter, verifying result is not NaN.
		result := a.TonnesForcePerCubicCentimeter()
		cacheResult := a.TonnesForcePerCubicCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonnesForcePerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerCubicMeter.
		// No expected conversion value provided for TonnesForcePerCubicMeter, verifying result is not NaN.
		result := a.TonnesForcePerCubicMeter()
		cacheResult := a.TonnesForcePerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonnesForcePerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerCubicMillimeter.
		// No expected conversion value provided for KilonewtonsPerCubicMillimeter, verifying result is not NaN.
		result := a.KilonewtonsPerCubicMillimeter()
		cacheResult := a.KilonewtonsPerCubicMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonsPerCubicMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerCubicCentimeter.
		// No expected conversion value provided for KilonewtonsPerCubicCentimeter, verifying result is not NaN.
		result := a.KilonewtonsPerCubicCentimeter()
		cacheResult := a.KilonewtonsPerCubicCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonsPerCubicCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerCubicMeter.
		// No expected conversion value provided for KilonewtonsPerCubicMeter, verifying result is not NaN.
		result := a.KilonewtonsPerCubicMeter()
		cacheResult := a.KilonewtonsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonsPerCubicMeter.
		// No expected conversion value provided for MeganewtonsPerCubicMeter, verifying result is not NaN.
		result := a.MeganewtonsPerCubicMeter()
		cacheResult := a.MeganewtonsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MeganewtonsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerCubicInch.
		// No expected conversion value provided for KilopoundsForcePerCubicInch, verifying result is not NaN.
		result := a.KilopoundsForcePerCubicInch()
		cacheResult := a.KilopoundsForcePerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsForcePerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerCubicFoot.
		// No expected conversion value provided for KilopoundsForcePerCubicFoot, verifying result is not NaN.
		result := a.KilopoundsForcePerCubicFoot()
		cacheResult := a.KilopoundsForcePerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsForcePerCubicFoot returned NaN")
		}
	}
}

func TestSpecificWeight_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.SpecificWeightFactory{}
	a, err := factory.CreateSpecificWeight(90, units.SpecificWeightNewtonPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.SpecificWeightNewtonPerCubicMeter {
		t.Errorf("expected default unit NewtonPerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.SpecificWeightNewtonPerCubicMillimeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.SpecificWeightDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.SpecificWeightNewtonPerCubicMeter {
		t.Errorf("expected unit NewtonPerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestSpecificWeightFactory_FromDto(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightNewtonPerCubicMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.SpecificWeightDto{
        Value: math.NaN(),
        Unit:  units.SpecificWeightNewtonPerCubicMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test NewtonPerCubicMillimeter conversion
    newtons_per_cubic_millimeterDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightNewtonPerCubicMillimeter,
    }
    
    var newtons_per_cubic_millimeterResult *units.SpecificWeight
    newtons_per_cubic_millimeterResult, err = factory.FromDto(newtons_per_cubic_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonPerCubicMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_cubic_millimeterResult.Convert(units.SpecificWeightNewtonPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test NewtonPerCubicCentimeter conversion
    newtons_per_cubic_centimeterDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightNewtonPerCubicCentimeter,
    }
    
    var newtons_per_cubic_centimeterResult *units.SpecificWeight
    newtons_per_cubic_centimeterResult, err = factory.FromDto(newtons_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonPerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_cubic_centimeterResult.Convert(units.SpecificWeightNewtonPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test NewtonPerCubicMeter conversion
    newtons_per_cubic_meterDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightNewtonPerCubicMeter,
    }
    
    var newtons_per_cubic_meterResult *units.SpecificWeight
    newtons_per_cubic_meterResult, err = factory.FromDto(newtons_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_cubic_meterResult.Convert(units.SpecificWeightNewtonPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test KilogramForcePerCubicMillimeter conversion
    kilograms_force_per_cubic_millimeterDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightKilogramForcePerCubicMillimeter,
    }
    
    var kilograms_force_per_cubic_millimeterResult *units.SpecificWeight
    kilograms_force_per_cubic_millimeterResult, err = factory.FromDto(kilograms_force_per_cubic_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForcePerCubicMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_cubic_millimeterResult.Convert(units.SpecificWeightKilogramForcePerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test KilogramForcePerCubicCentimeter conversion
    kilograms_force_per_cubic_centimeterDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightKilogramForcePerCubicCentimeter,
    }
    
    var kilograms_force_per_cubic_centimeterResult *units.SpecificWeight
    kilograms_force_per_cubic_centimeterResult, err = factory.FromDto(kilograms_force_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForcePerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_cubic_centimeterResult.Convert(units.SpecificWeightKilogramForcePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test KilogramForcePerCubicMeter conversion
    kilograms_force_per_cubic_meterDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightKilogramForcePerCubicMeter,
    }
    
    var kilograms_force_per_cubic_meterResult *units.SpecificWeight
    kilograms_force_per_cubic_meterResult, err = factory.FromDto(kilograms_force_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForcePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_cubic_meterResult.Convert(units.SpecificWeightKilogramForcePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test PoundForcePerCubicInch conversion
    pounds_force_per_cubic_inchDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightPoundForcePerCubicInch,
    }
    
    var pounds_force_per_cubic_inchResult *units.SpecificWeight
    pounds_force_per_cubic_inchResult, err = factory.FromDto(pounds_force_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForcePerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_cubic_inchResult.Convert(units.SpecificWeightPoundForcePerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerCubicInch = %v, want %v", converted, 100)
    }
    // Test PoundForcePerCubicFoot conversion
    pounds_force_per_cubic_footDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightPoundForcePerCubicFoot,
    }
    
    var pounds_force_per_cubic_footResult *units.SpecificWeight
    pounds_force_per_cubic_footResult, err = factory.FromDto(pounds_force_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForcePerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_cubic_footResult.Convert(units.SpecificWeightPoundForcePerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerCubicFoot = %v, want %v", converted, 100)
    }
    // Test TonneForcePerCubicMillimeter conversion
    tonnes_force_per_cubic_millimeterDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightTonneForcePerCubicMillimeter,
    }
    
    var tonnes_force_per_cubic_millimeterResult *units.SpecificWeight
    tonnes_force_per_cubic_millimeterResult, err = factory.FromDto(tonnes_force_per_cubic_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForcePerCubicMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_cubic_millimeterResult.Convert(units.SpecificWeightTonneForcePerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test TonneForcePerCubicCentimeter conversion
    tonnes_force_per_cubic_centimeterDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightTonneForcePerCubicCentimeter,
    }
    
    var tonnes_force_per_cubic_centimeterResult *units.SpecificWeight
    tonnes_force_per_cubic_centimeterResult, err = factory.FromDto(tonnes_force_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForcePerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_cubic_centimeterResult.Convert(units.SpecificWeightTonneForcePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test TonneForcePerCubicMeter conversion
    tonnes_force_per_cubic_meterDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightTonneForcePerCubicMeter,
    }
    
    var tonnes_force_per_cubic_meterResult *units.SpecificWeight
    tonnes_force_per_cubic_meterResult, err = factory.FromDto(tonnes_force_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForcePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_cubic_meterResult.Convert(units.SpecificWeightTonneForcePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonPerCubicMillimeter conversion
    kilonewtons_per_cubic_millimeterDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightKilonewtonPerCubicMillimeter,
    }
    
    var kilonewtons_per_cubic_millimeterResult *units.SpecificWeight
    kilonewtons_per_cubic_millimeterResult, err = factory.FromDto(kilonewtons_per_cubic_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonPerCubicMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_cubic_millimeterResult.Convert(units.SpecificWeightKilonewtonPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonPerCubicCentimeter conversion
    kilonewtons_per_cubic_centimeterDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightKilonewtonPerCubicCentimeter,
    }
    
    var kilonewtons_per_cubic_centimeterResult *units.SpecificWeight
    kilonewtons_per_cubic_centimeterResult, err = factory.FromDto(kilonewtons_per_cubic_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonPerCubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_cubic_centimeterResult.Convert(units.SpecificWeightKilonewtonPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonPerCubicMeter conversion
    kilonewtons_per_cubic_meterDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightKilonewtonPerCubicMeter,
    }
    
    var kilonewtons_per_cubic_meterResult *units.SpecificWeight
    kilonewtons_per_cubic_meterResult, err = factory.FromDto(kilonewtons_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_cubic_meterResult.Convert(units.SpecificWeightKilonewtonPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test MeganewtonPerCubicMeter conversion
    meganewtons_per_cubic_meterDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightMeganewtonPerCubicMeter,
    }
    
    var meganewtons_per_cubic_meterResult *units.SpecificWeight
    meganewtons_per_cubic_meterResult, err = factory.FromDto(meganewtons_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewtons_per_cubic_meterResult.Convert(units.SpecificWeightMeganewtonPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test KilopoundForcePerCubicInch conversion
    kilopounds_force_per_cubic_inchDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightKilopoundForcePerCubicInch,
    }
    
    var kilopounds_force_per_cubic_inchResult *units.SpecificWeight
    kilopounds_force_per_cubic_inchResult, err = factory.FromDto(kilopounds_force_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForcePerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_cubic_inchResult.Convert(units.SpecificWeightKilopoundForcePerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerCubicInch = %v, want %v", converted, 100)
    }
    // Test KilopoundForcePerCubicFoot conversion
    kilopounds_force_per_cubic_footDto := units.SpecificWeightDto{
        Value: 100,
        Unit:  units.SpecificWeightKilopoundForcePerCubicFoot,
    }
    
    var kilopounds_force_per_cubic_footResult *units.SpecificWeight
    kilopounds_force_per_cubic_footResult, err = factory.FromDto(kilopounds_force_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForcePerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_cubic_footResult.Convert(units.SpecificWeightKilopoundForcePerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerCubicFoot = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.SpecificWeightDto{
        Value: 0,
        Unit:  units.SpecificWeightNewtonPerCubicMeter,
    }
    
    var zeroResult *units.SpecificWeight
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestSpecificWeightFactory_FromDtoJSON(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "NewtonPerCubicMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "NewtonPerCubicMeter"}`)
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
    nanJSON, _ := json.Marshal(units.SpecificWeightDto{
        Value: nanValue,
        Unit:  units.SpecificWeightNewtonPerCubicMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with NewtonPerCubicMillimeter unit
    newtons_per_cubic_millimeterJSON := []byte(`{"value": 100, "unit": "NewtonPerCubicMillimeter"}`)
    newtons_per_cubic_millimeterResult, err := factory.FromDtoJSON(newtons_per_cubic_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonPerCubicMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_cubic_millimeterResult.Convert(units.SpecificWeightNewtonPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonPerCubicCentimeter unit
    newtons_per_cubic_centimeterJSON := []byte(`{"value": 100, "unit": "NewtonPerCubicCentimeter"}`)
    newtons_per_cubic_centimeterResult, err := factory.FromDtoJSON(newtons_per_cubic_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonPerCubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_cubic_centimeterResult.Convert(units.SpecificWeightNewtonPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonPerCubicMeter unit
    newtons_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "NewtonPerCubicMeter"}`)
    newtons_per_cubic_meterResult, err := factory.FromDtoJSON(newtons_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_cubic_meterResult.Convert(units.SpecificWeightNewtonPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForcePerCubicMillimeter unit
    kilograms_force_per_cubic_millimeterJSON := []byte(`{"value": 100, "unit": "KilogramForcePerCubicMillimeter"}`)
    kilograms_force_per_cubic_millimeterResult, err := factory.FromDtoJSON(kilograms_force_per_cubic_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForcePerCubicMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_cubic_millimeterResult.Convert(units.SpecificWeightKilogramForcePerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForcePerCubicCentimeter unit
    kilograms_force_per_cubic_centimeterJSON := []byte(`{"value": 100, "unit": "KilogramForcePerCubicCentimeter"}`)
    kilograms_force_per_cubic_centimeterResult, err := factory.FromDtoJSON(kilograms_force_per_cubic_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForcePerCubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_cubic_centimeterResult.Convert(units.SpecificWeightKilogramForcePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForcePerCubicMeter unit
    kilograms_force_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "KilogramForcePerCubicMeter"}`)
    kilograms_force_per_cubic_meterResult, err := factory.FromDtoJSON(kilograms_force_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForcePerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_cubic_meterResult.Convert(units.SpecificWeightKilogramForcePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForcePerCubicInch unit
    pounds_force_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "PoundForcePerCubicInch"}`)
    pounds_force_per_cubic_inchResult, err := factory.FromDtoJSON(pounds_force_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForcePerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_cubic_inchResult.Convert(units.SpecificWeightPoundForcePerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForcePerCubicFoot unit
    pounds_force_per_cubic_footJSON := []byte(`{"value": 100, "unit": "PoundForcePerCubicFoot"}`)
    pounds_force_per_cubic_footResult, err := factory.FromDtoJSON(pounds_force_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForcePerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_cubic_footResult.Convert(units.SpecificWeightPoundForcePerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForcePerCubicMillimeter unit
    tonnes_force_per_cubic_millimeterJSON := []byte(`{"value": 100, "unit": "TonneForcePerCubicMillimeter"}`)
    tonnes_force_per_cubic_millimeterResult, err := factory.FromDtoJSON(tonnes_force_per_cubic_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForcePerCubicMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_cubic_millimeterResult.Convert(units.SpecificWeightTonneForcePerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForcePerCubicCentimeter unit
    tonnes_force_per_cubic_centimeterJSON := []byte(`{"value": 100, "unit": "TonneForcePerCubicCentimeter"}`)
    tonnes_force_per_cubic_centimeterResult, err := factory.FromDtoJSON(tonnes_force_per_cubic_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForcePerCubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_cubic_centimeterResult.Convert(units.SpecificWeightTonneForcePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForcePerCubicMeter unit
    tonnes_force_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "TonneForcePerCubicMeter"}`)
    tonnes_force_per_cubic_meterResult, err := factory.FromDtoJSON(tonnes_force_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForcePerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_cubic_meterResult.Convert(units.SpecificWeightTonneForcePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonPerCubicMillimeter unit
    kilonewtons_per_cubic_millimeterJSON := []byte(`{"value": 100, "unit": "KilonewtonPerCubicMillimeter"}`)
    kilonewtons_per_cubic_millimeterResult, err := factory.FromDtoJSON(kilonewtons_per_cubic_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonPerCubicMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_cubic_millimeterResult.Convert(units.SpecificWeightKilonewtonPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerCubicMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonPerCubicCentimeter unit
    kilonewtons_per_cubic_centimeterJSON := []byte(`{"value": 100, "unit": "KilonewtonPerCubicCentimeter"}`)
    kilonewtons_per_cubic_centimeterResult, err := factory.FromDtoJSON(kilonewtons_per_cubic_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonPerCubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_cubic_centimeterResult.Convert(units.SpecificWeightKilonewtonPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerCubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonPerCubicMeter unit
    kilonewtons_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "KilonewtonPerCubicMeter"}`)
    kilonewtons_per_cubic_meterResult, err := factory.FromDtoJSON(kilonewtons_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_cubic_meterResult.Convert(units.SpecificWeightKilonewtonPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonPerCubicMeter unit
    meganewtons_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "MeganewtonPerCubicMeter"}`)
    meganewtons_per_cubic_meterResult, err := factory.FromDtoJSON(meganewtons_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewtons_per_cubic_meterResult.Convert(units.SpecificWeightMeganewtonPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForcePerCubicInch unit
    kilopounds_force_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "KilopoundForcePerCubicInch"}`)
    kilopounds_force_per_cubic_inchResult, err := factory.FromDtoJSON(kilopounds_force_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForcePerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_cubic_inchResult.Convert(units.SpecificWeightKilopoundForcePerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForcePerCubicFoot unit
    kilopounds_force_per_cubic_footJSON := []byte(`{"value": 100, "unit": "KilopoundForcePerCubicFoot"}`)
    kilopounds_force_per_cubic_footResult, err := factory.FromDtoJSON(kilopounds_force_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForcePerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_cubic_footResult.Convert(units.SpecificWeightKilopoundForcePerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerCubicFoot = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "NewtonPerCubicMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromNewtonsPerCubicMillimeter function
func TestSpecificWeightFactory_FromNewtonsPerCubicMillimeter(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonsPerCubicMillimeter(100)
    if err != nil {
        t.Errorf("FromNewtonsPerCubicMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightNewtonPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonsPerCubicMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonsPerCubicMillimeter(math.NaN())
    if err == nil {
        t.Error("FromNewtonsPerCubicMillimeter() with NaN value should return error")
    }

    _, err = factory.FromNewtonsPerCubicMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonsPerCubicMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromNewtonsPerCubicMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonsPerCubicMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonsPerCubicMillimeter(0)
    if err != nil {
        t.Errorf("FromNewtonsPerCubicMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightNewtonPerCubicMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonsPerCubicMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonsPerCubicCentimeter function
func TestSpecificWeightFactory_FromNewtonsPerCubicCentimeter(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonsPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromNewtonsPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightNewtonPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonsPerCubicCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonsPerCubicCentimeter(math.NaN())
    if err == nil {
        t.Error("FromNewtonsPerCubicCentimeter() with NaN value should return error")
    }

    _, err = factory.FromNewtonsPerCubicCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonsPerCubicCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromNewtonsPerCubicCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonsPerCubicCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonsPerCubicCentimeter(0)
    if err != nil {
        t.Errorf("FromNewtonsPerCubicCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightNewtonPerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonsPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonsPerCubicMeter function
func TestSpecificWeightFactory_FromNewtonsPerCubicMeter(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromNewtonsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightNewtonPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromNewtonsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromNewtonsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromNewtonsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromNewtonsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightNewtonPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsForcePerCubicMillimeter function
func TestSpecificWeightFactory_FromKilogramsForcePerCubicMillimeter(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsForcePerCubicMillimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsForcePerCubicMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightKilogramForcePerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsForcePerCubicMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsForcePerCubicMillimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsForcePerCubicMillimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsForcePerCubicMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsForcePerCubicMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsForcePerCubicMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsForcePerCubicMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsForcePerCubicMillimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsForcePerCubicMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightKilogramForcePerCubicMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsForcePerCubicMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsForcePerCubicCentimeter function
func TestSpecificWeightFactory_FromKilogramsForcePerCubicCentimeter(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsForcePerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsForcePerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightKilogramForcePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsForcePerCubicCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsForcePerCubicCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsForcePerCubicCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsForcePerCubicCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsForcePerCubicCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsForcePerCubicCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsForcePerCubicCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsForcePerCubicCentimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsForcePerCubicCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightKilogramForcePerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsForcePerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsForcePerCubicMeter function
func TestSpecificWeightFactory_FromKilogramsForcePerCubicMeter(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsForcePerCubicMeter(100)
    if err != nil {
        t.Errorf("FromKilogramsForcePerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightKilogramForcePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsForcePerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsForcePerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsForcePerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsForcePerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsForcePerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsForcePerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsForcePerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsForcePerCubicMeter(0)
    if err != nil {
        t.Errorf("FromKilogramsForcePerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightKilogramForcePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsForcePerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForcePerCubicInch function
func TestSpecificWeightFactory_FromPoundsForcePerCubicInch(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForcePerCubicInch(100)
    if err != nil {
        t.Errorf("FromPoundsForcePerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightPoundForcePerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForcePerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForcePerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromPoundsForcePerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromPoundsForcePerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForcePerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForcePerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForcePerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForcePerCubicInch(0)
    if err != nil {
        t.Errorf("FromPoundsForcePerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightPoundForcePerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForcePerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForcePerCubicFoot function
func TestSpecificWeightFactory_FromPoundsForcePerCubicFoot(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForcePerCubicFoot(100)
    if err != nil {
        t.Errorf("FromPoundsForcePerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightPoundForcePerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForcePerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForcePerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromPoundsForcePerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromPoundsForcePerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForcePerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForcePerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForcePerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForcePerCubicFoot(0)
    if err != nil {
        t.Errorf("FromPoundsForcePerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightPoundForcePerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForcePerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesForcePerCubicMillimeter function
func TestSpecificWeightFactory_FromTonnesForcePerCubicMillimeter(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesForcePerCubicMillimeter(100)
    if err != nil {
        t.Errorf("FromTonnesForcePerCubicMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightTonneForcePerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesForcePerCubicMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesForcePerCubicMillimeter(math.NaN())
    if err == nil {
        t.Error("FromTonnesForcePerCubicMillimeter() with NaN value should return error")
    }

    _, err = factory.FromTonnesForcePerCubicMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesForcePerCubicMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromTonnesForcePerCubicMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesForcePerCubicMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesForcePerCubicMillimeter(0)
    if err != nil {
        t.Errorf("FromTonnesForcePerCubicMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightTonneForcePerCubicMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesForcePerCubicMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesForcePerCubicCentimeter function
func TestSpecificWeightFactory_FromTonnesForcePerCubicCentimeter(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesForcePerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromTonnesForcePerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightTonneForcePerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesForcePerCubicCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesForcePerCubicCentimeter(math.NaN())
    if err == nil {
        t.Error("FromTonnesForcePerCubicCentimeter() with NaN value should return error")
    }

    _, err = factory.FromTonnesForcePerCubicCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesForcePerCubicCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromTonnesForcePerCubicCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesForcePerCubicCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesForcePerCubicCentimeter(0)
    if err != nil {
        t.Errorf("FromTonnesForcePerCubicCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightTonneForcePerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesForcePerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesForcePerCubicMeter function
func TestSpecificWeightFactory_FromTonnesForcePerCubicMeter(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesForcePerCubicMeter(100)
    if err != nil {
        t.Errorf("FromTonnesForcePerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightTonneForcePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesForcePerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesForcePerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromTonnesForcePerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromTonnesForcePerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesForcePerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromTonnesForcePerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesForcePerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesForcePerCubicMeter(0)
    if err != nil {
        t.Errorf("FromTonnesForcePerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightTonneForcePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesForcePerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonsPerCubicMillimeter function
func TestSpecificWeightFactory_FromKilonewtonsPerCubicMillimeter(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonsPerCubicMillimeter(100)
    if err != nil {
        t.Errorf("FromKilonewtonsPerCubicMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightKilonewtonPerCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonsPerCubicMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonsPerCubicMillimeter(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonsPerCubicMillimeter() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonsPerCubicMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonsPerCubicMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonsPerCubicMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonsPerCubicMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonsPerCubicMillimeter(0)
    if err != nil {
        t.Errorf("FromKilonewtonsPerCubicMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightKilonewtonPerCubicMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonsPerCubicMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonsPerCubicCentimeter function
func TestSpecificWeightFactory_FromKilonewtonsPerCubicCentimeter(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonsPerCubicCentimeter(100)
    if err != nil {
        t.Errorf("FromKilonewtonsPerCubicCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightKilonewtonPerCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonsPerCubicCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonsPerCubicCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonsPerCubicCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonsPerCubicCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonsPerCubicCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonsPerCubicCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonsPerCubicCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonsPerCubicCentimeter(0)
    if err != nil {
        t.Errorf("FromKilonewtonsPerCubicCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightKilonewtonPerCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonsPerCubicCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonsPerCubicMeter function
func TestSpecificWeightFactory_FromKilonewtonsPerCubicMeter(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromKilonewtonsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightKilonewtonPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromKilonewtonsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightKilonewtonPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonsPerCubicMeter function
func TestSpecificWeightFactory_FromMeganewtonsPerCubicMeter(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMeganewtonsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightMeganewtonPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromMeganewtonsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightMeganewtonPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsForcePerCubicInch function
func TestSpecificWeightFactory_FromKilopoundsForcePerCubicInch(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsForcePerCubicInch(100)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightKilopoundForcePerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsForcePerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsForcePerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsForcePerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsForcePerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsForcePerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsForcePerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsForcePerCubicInch(0)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightKilopoundForcePerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsForcePerCubicFoot function
func TestSpecificWeightFactory_FromKilopoundsForcePerCubicFoot(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsForcePerCubicFoot(100)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificWeightKilopoundForcePerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsForcePerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsForcePerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsForcePerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsForcePerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsForcePerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsForcePerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsForcePerCubicFoot(0)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificWeightKilopoundForcePerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerCubicFoot() with zero value = %v, want 0", converted)
    }
}

func TestSpecificWeightToString(t *testing.T) {
	factory := units.SpecificWeightFactory{}
	a, err := factory.CreateSpecificWeight(45, units.SpecificWeightNewtonPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.SpecificWeightNewtonPerCubicMeter, 2)
	expected := "45.00 " + units.GetSpecificWeightAbbreviation(units.SpecificWeightNewtonPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.SpecificWeightNewtonPerCubicMeter, -1)
	expected = "45 " + units.GetSpecificWeightAbbreviation(units.SpecificWeightNewtonPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestSpecificWeight_EqualityAndComparison(t *testing.T) {
	factory := units.SpecificWeightFactory{}
	a1, _ := factory.CreateSpecificWeight(60, units.SpecificWeightNewtonPerCubicMeter)
	a2, _ := factory.CreateSpecificWeight(60, units.SpecificWeightNewtonPerCubicMeter)
	a3, _ := factory.CreateSpecificWeight(120, units.SpecificWeightNewtonPerCubicMeter)

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

func TestSpecificWeight_Arithmetic(t *testing.T) {
	factory := units.SpecificWeightFactory{}
	a1, _ := factory.CreateSpecificWeight(30, units.SpecificWeightNewtonPerCubicMeter)
	a2, _ := factory.CreateSpecificWeight(45, units.SpecificWeightNewtonPerCubicMeter)

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


func TestGetSpecificWeightAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.SpecificWeightUnits
        want string
    }{
        {
            name: "NewtonPerCubicMillimeter abbreviation",
            unit: units.SpecificWeightNewtonPerCubicMillimeter,
            want: "N/mm³",
        },
        {
            name: "NewtonPerCubicCentimeter abbreviation",
            unit: units.SpecificWeightNewtonPerCubicCentimeter,
            want: "N/cm³",
        },
        {
            name: "NewtonPerCubicMeter abbreviation",
            unit: units.SpecificWeightNewtonPerCubicMeter,
            want: "N/m³",
        },
        {
            name: "KilogramForcePerCubicMillimeter abbreviation",
            unit: units.SpecificWeightKilogramForcePerCubicMillimeter,
            want: "kgf/mm³",
        },
        {
            name: "KilogramForcePerCubicCentimeter abbreviation",
            unit: units.SpecificWeightKilogramForcePerCubicCentimeter,
            want: "kgf/cm³",
        },
        {
            name: "KilogramForcePerCubicMeter abbreviation",
            unit: units.SpecificWeightKilogramForcePerCubicMeter,
            want: "kgf/m³",
        },
        {
            name: "PoundForcePerCubicInch abbreviation",
            unit: units.SpecificWeightPoundForcePerCubicInch,
            want: "lbf/in³",
        },
        {
            name: "PoundForcePerCubicFoot abbreviation",
            unit: units.SpecificWeightPoundForcePerCubicFoot,
            want: "lbf/ft³",
        },
        {
            name: "TonneForcePerCubicMillimeter abbreviation",
            unit: units.SpecificWeightTonneForcePerCubicMillimeter,
            want: "tf/mm³",
        },
        {
            name: "TonneForcePerCubicCentimeter abbreviation",
            unit: units.SpecificWeightTonneForcePerCubicCentimeter,
            want: "tf/cm³",
        },
        {
            name: "TonneForcePerCubicMeter abbreviation",
            unit: units.SpecificWeightTonneForcePerCubicMeter,
            want: "tf/m³",
        },
        {
            name: "KilonewtonPerCubicMillimeter abbreviation",
            unit: units.SpecificWeightKilonewtonPerCubicMillimeter,
            want: "kN/mm³",
        },
        {
            name: "KilonewtonPerCubicCentimeter abbreviation",
            unit: units.SpecificWeightKilonewtonPerCubicCentimeter,
            want: "kN/cm³",
        },
        {
            name: "KilonewtonPerCubicMeter abbreviation",
            unit: units.SpecificWeightKilonewtonPerCubicMeter,
            want: "kN/m³",
        },
        {
            name: "MeganewtonPerCubicMeter abbreviation",
            unit: units.SpecificWeightMeganewtonPerCubicMeter,
            want: "MN/m³",
        },
        {
            name: "KilopoundForcePerCubicInch abbreviation",
            unit: units.SpecificWeightKilopoundForcePerCubicInch,
            want: "klbf/in³",
        },
        {
            name: "KilopoundForcePerCubicFoot abbreviation",
            unit: units.SpecificWeightKilopoundForcePerCubicFoot,
            want: "klbf/ft³",
        },
        {
            name: "invalid unit",
            unit: units.SpecificWeightUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetSpecificWeightAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetSpecificWeightAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestSpecificWeight_String(t *testing.T) {
    factory := units.SpecificWeightFactory{}
    
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
            unit, err := factory.CreateSpecificWeight(tt.value, units.SpecificWeightNewtonPerCubicMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("SpecificWeight.String() = %v, want %v", got, tt.want)
            }
        })
    }
}