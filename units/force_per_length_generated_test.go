// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestForcePerLengthDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonPerMeter"}`
	
	factory := units.ForcePerLengthDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ForcePerLengthNewtonPerMeter {
		t.Errorf("expected unit %v, got %v", units.ForcePerLengthNewtonPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestForcePerLengthDto_ToJSON(t *testing.T) {
	dto := units.ForcePerLengthDto{
		Value: 45,
		Unit:  units.ForcePerLengthNewtonPerMeter,
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
	if result["unit"].(string) != string(units.ForcePerLengthNewtonPerMeter) {
		t.Errorf("expected unit %s, got %v", units.ForcePerLengthNewtonPerMeter, result["unit"])
	}
}

func TestNewForcePerLength_InvalidValue(t *testing.T) {
	factory := units.ForcePerLengthFactory{}
	// NaN value should return an error.
	_, err := factory.CreateForcePerLength(math.NaN(), units.ForcePerLengthNewtonPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateForcePerLength(math.Inf(1), units.ForcePerLengthNewtonPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestForcePerLengthConversions(t *testing.T) {
	factory := units.ForcePerLengthFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateForcePerLength(180, units.ForcePerLengthNewtonPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonsPerMeter.
		// No expected conversion value provided for NewtonsPerMeter, verifying result is not NaN.
		result := a.NewtonsPerMeter()
		cacheResult := a.NewtonsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerCentimeter.
		// No expected conversion value provided for NewtonsPerCentimeter, verifying result is not NaN.
		result := a.NewtonsPerCentimeter()
		cacheResult := a.NewtonsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerMillimeter.
		// No expected conversion value provided for NewtonsPerMillimeter, verifying result is not NaN.
		result := a.NewtonsPerMillimeter()
		cacheResult := a.NewtonsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerMeter.
		// No expected conversion value provided for KilogramsForcePerMeter, verifying result is not NaN.
		result := a.KilogramsForcePerMeter()
		cacheResult := a.KilogramsForcePerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsForcePerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerCentimeter.
		// No expected conversion value provided for KilogramsForcePerCentimeter, verifying result is not NaN.
		result := a.KilogramsForcePerCentimeter()
		cacheResult := a.KilogramsForcePerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsForcePerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerMillimeter.
		// No expected conversion value provided for KilogramsForcePerMillimeter, verifying result is not NaN.
		result := a.KilogramsForcePerMillimeter()
		cacheResult := a.KilogramsForcePerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsForcePerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerMeter.
		// No expected conversion value provided for TonnesForcePerMeter, verifying result is not NaN.
		result := a.TonnesForcePerMeter()
		cacheResult := a.TonnesForcePerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonnesForcePerMeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerCentimeter.
		// No expected conversion value provided for TonnesForcePerCentimeter, verifying result is not NaN.
		result := a.TonnesForcePerCentimeter()
		cacheResult := a.TonnesForcePerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonnesForcePerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerMillimeter.
		// No expected conversion value provided for TonnesForcePerMillimeter, verifying result is not NaN.
		result := a.TonnesForcePerMillimeter()
		cacheResult := a.TonnesForcePerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonnesForcePerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerFoot.
		// No expected conversion value provided for PoundsForcePerFoot, verifying result is not NaN.
		result := a.PoundsForcePerFoot()
		cacheResult := a.PoundsForcePerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsForcePerFoot returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerInch.
		// No expected conversion value provided for PoundsForcePerInch, verifying result is not NaN.
		result := a.PoundsForcePerInch()
		cacheResult := a.PoundsForcePerInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsForcePerInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerYard.
		// No expected conversion value provided for PoundsForcePerYard, verifying result is not NaN.
		result := a.PoundsForcePerYard()
		cacheResult := a.PoundsForcePerYard()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsForcePerYard returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerFoot.
		// No expected conversion value provided for KilopoundsForcePerFoot, verifying result is not NaN.
		result := a.KilopoundsForcePerFoot()
		cacheResult := a.KilopoundsForcePerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsForcePerFoot returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerInch.
		// No expected conversion value provided for KilopoundsForcePerInch, verifying result is not NaN.
		result := a.KilopoundsForcePerInch()
		cacheResult := a.KilopoundsForcePerInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsForcePerInch returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonsPerMeter.
		// No expected conversion value provided for NanonewtonsPerMeter, verifying result is not NaN.
		result := a.NanonewtonsPerMeter()
		cacheResult := a.NanonewtonsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanonewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonsPerMeter.
		// No expected conversion value provided for MicronewtonsPerMeter, verifying result is not NaN.
		result := a.MicronewtonsPerMeter()
		cacheResult := a.MicronewtonsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicronewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonsPerMeter.
		// No expected conversion value provided for MillinewtonsPerMeter, verifying result is not NaN.
		result := a.MillinewtonsPerMeter()
		cacheResult := a.MillinewtonsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillinewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonsPerMeter.
		// No expected conversion value provided for CentinewtonsPerMeter, verifying result is not NaN.
		result := a.CentinewtonsPerMeter()
		cacheResult := a.CentinewtonsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentinewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonsPerMeter.
		// No expected conversion value provided for DecinewtonsPerMeter, verifying result is not NaN.
		result := a.DecinewtonsPerMeter()
		cacheResult := a.DecinewtonsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecinewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonsPerMeter.
		// No expected conversion value provided for DecanewtonsPerMeter, verifying result is not NaN.
		result := a.DecanewtonsPerMeter()
		cacheResult := a.DecanewtonsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecanewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerMeter.
		// No expected conversion value provided for KilonewtonsPerMeter, verifying result is not NaN.
		result := a.KilonewtonsPerMeter()
		cacheResult := a.KilonewtonsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonsPerMeter.
		// No expected conversion value provided for MeganewtonsPerMeter, verifying result is not NaN.
		result := a.MeganewtonsPerMeter()
		cacheResult := a.MeganewtonsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MeganewtonsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonsPerCentimeter.
		// No expected conversion value provided for NanonewtonsPerCentimeter, verifying result is not NaN.
		result := a.NanonewtonsPerCentimeter()
		cacheResult := a.NanonewtonsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanonewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonsPerCentimeter.
		// No expected conversion value provided for MicronewtonsPerCentimeter, verifying result is not NaN.
		result := a.MicronewtonsPerCentimeter()
		cacheResult := a.MicronewtonsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicronewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonsPerCentimeter.
		// No expected conversion value provided for MillinewtonsPerCentimeter, verifying result is not NaN.
		result := a.MillinewtonsPerCentimeter()
		cacheResult := a.MillinewtonsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillinewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonsPerCentimeter.
		// No expected conversion value provided for CentinewtonsPerCentimeter, verifying result is not NaN.
		result := a.CentinewtonsPerCentimeter()
		cacheResult := a.CentinewtonsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentinewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonsPerCentimeter.
		// No expected conversion value provided for DecinewtonsPerCentimeter, verifying result is not NaN.
		result := a.DecinewtonsPerCentimeter()
		cacheResult := a.DecinewtonsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecinewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonsPerCentimeter.
		// No expected conversion value provided for DecanewtonsPerCentimeter, verifying result is not NaN.
		result := a.DecanewtonsPerCentimeter()
		cacheResult := a.DecanewtonsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecanewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerCentimeter.
		// No expected conversion value provided for KilonewtonsPerCentimeter, verifying result is not NaN.
		result := a.KilonewtonsPerCentimeter()
		cacheResult := a.KilonewtonsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonsPerCentimeter.
		// No expected conversion value provided for MeganewtonsPerCentimeter, verifying result is not NaN.
		result := a.MeganewtonsPerCentimeter()
		cacheResult := a.MeganewtonsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MeganewtonsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to NanonewtonsPerMillimeter.
		// No expected conversion value provided for NanonewtonsPerMillimeter, verifying result is not NaN.
		result := a.NanonewtonsPerMillimeter()
		cacheResult := a.NanonewtonsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanonewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MicronewtonsPerMillimeter.
		// No expected conversion value provided for MicronewtonsPerMillimeter, verifying result is not NaN.
		result := a.MicronewtonsPerMillimeter()
		cacheResult := a.MicronewtonsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicronewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MillinewtonsPerMillimeter.
		// No expected conversion value provided for MillinewtonsPerMillimeter, verifying result is not NaN.
		result := a.MillinewtonsPerMillimeter()
		cacheResult := a.MillinewtonsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillinewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to CentinewtonsPerMillimeter.
		// No expected conversion value provided for CentinewtonsPerMillimeter, verifying result is not NaN.
		result := a.CentinewtonsPerMillimeter()
		cacheResult := a.CentinewtonsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentinewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to DecinewtonsPerMillimeter.
		// No expected conversion value provided for DecinewtonsPerMillimeter, verifying result is not NaN.
		result := a.DecinewtonsPerMillimeter()
		cacheResult := a.DecinewtonsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecinewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to DecanewtonsPerMillimeter.
		// No expected conversion value provided for DecanewtonsPerMillimeter, verifying result is not NaN.
		result := a.DecanewtonsPerMillimeter()
		cacheResult := a.DecanewtonsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecanewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerMillimeter.
		// No expected conversion value provided for KilonewtonsPerMillimeter, verifying result is not NaN.
		result := a.KilonewtonsPerMillimeter()
		cacheResult := a.KilonewtonsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonsPerMillimeter.
		// No expected conversion value provided for MeganewtonsPerMillimeter, verifying result is not NaN.
		result := a.MeganewtonsPerMillimeter()
		cacheResult := a.MeganewtonsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MeganewtonsPerMillimeter returned NaN")
		}
	}
}

func TestForcePerLength_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ForcePerLengthFactory{}
	a, err := factory.CreateForcePerLength(90, units.ForcePerLengthNewtonPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ForcePerLengthNewtonPerMeter {
		t.Errorf("expected default unit NewtonPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ForcePerLengthNewtonPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ForcePerLengthDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ForcePerLengthNewtonPerMeter {
		t.Errorf("expected unit NewtonPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestForcePerLengthFactory_FromDto(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthNewtonPerMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ForcePerLengthDto{
        Value: math.NaN(),
        Unit:  units.ForcePerLengthNewtonPerMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test NewtonPerMeter conversion
    newtons_per_meterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthNewtonPerMeter,
    }
    
    var newtons_per_meterResult *units.ForcePerLength
    newtons_per_meterResult, err = factory.FromDto(newtons_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_meterResult.Convert(units.ForcePerLengthNewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test NewtonPerCentimeter conversion
    newtons_per_centimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthNewtonPerCentimeter,
    }
    
    var newtons_per_centimeterResult *units.ForcePerLength
    newtons_per_centimeterResult, err = factory.FromDto(newtons_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_centimeterResult.Convert(units.ForcePerLengthNewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test NewtonPerMillimeter conversion
    newtons_per_millimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthNewtonPerMillimeter,
    }
    
    var newtons_per_millimeterResult *units.ForcePerLength
    newtons_per_millimeterResult, err = factory.FromDto(newtons_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_millimeterResult.Convert(units.ForcePerLengthNewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test KilogramForcePerMeter conversion
    kilograms_force_per_meterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthKilogramForcePerMeter,
    }
    
    var kilograms_force_per_meterResult *units.ForcePerLength
    kilograms_force_per_meterResult, err = factory.FromDto(kilograms_force_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForcePerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_meterResult.Convert(units.ForcePerLengthKilogramForcePerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerMeter = %v, want %v", converted, 100)
    }
    // Test KilogramForcePerCentimeter conversion
    kilograms_force_per_centimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthKilogramForcePerCentimeter,
    }
    
    var kilograms_force_per_centimeterResult *units.ForcePerLength
    kilograms_force_per_centimeterResult, err = factory.FromDto(kilograms_force_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForcePerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_centimeterResult.Convert(units.ForcePerLengthKilogramForcePerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerCentimeter = %v, want %v", converted, 100)
    }
    // Test KilogramForcePerMillimeter conversion
    kilograms_force_per_millimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthKilogramForcePerMillimeter,
    }
    
    var kilograms_force_per_millimeterResult *units.ForcePerLength
    kilograms_force_per_millimeterResult, err = factory.FromDto(kilograms_force_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForcePerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_millimeterResult.Convert(units.ForcePerLengthKilogramForcePerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerMillimeter = %v, want %v", converted, 100)
    }
    // Test TonneForcePerMeter conversion
    tonnes_force_per_meterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthTonneForcePerMeter,
    }
    
    var tonnes_force_per_meterResult *units.ForcePerLength
    tonnes_force_per_meterResult, err = factory.FromDto(tonnes_force_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForcePerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_meterResult.Convert(units.ForcePerLengthTonneForcePerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerMeter = %v, want %v", converted, 100)
    }
    // Test TonneForcePerCentimeter conversion
    tonnes_force_per_centimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthTonneForcePerCentimeter,
    }
    
    var tonnes_force_per_centimeterResult *units.ForcePerLength
    tonnes_force_per_centimeterResult, err = factory.FromDto(tonnes_force_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForcePerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_centimeterResult.Convert(units.ForcePerLengthTonneForcePerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerCentimeter = %v, want %v", converted, 100)
    }
    // Test TonneForcePerMillimeter conversion
    tonnes_force_per_millimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthTonneForcePerMillimeter,
    }
    
    var tonnes_force_per_millimeterResult *units.ForcePerLength
    tonnes_force_per_millimeterResult, err = factory.FromDto(tonnes_force_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForcePerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_millimeterResult.Convert(units.ForcePerLengthTonneForcePerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerMillimeter = %v, want %v", converted, 100)
    }
    // Test PoundForcePerFoot conversion
    pounds_force_per_footDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthPoundForcePerFoot,
    }
    
    var pounds_force_per_footResult *units.ForcePerLength
    pounds_force_per_footResult, err = factory.FromDto(pounds_force_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForcePerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_footResult.Convert(units.ForcePerLengthPoundForcePerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerFoot = %v, want %v", converted, 100)
    }
    // Test PoundForcePerInch conversion
    pounds_force_per_inchDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthPoundForcePerInch,
    }
    
    var pounds_force_per_inchResult *units.ForcePerLength
    pounds_force_per_inchResult, err = factory.FromDto(pounds_force_per_inchDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForcePerInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_inchResult.Convert(units.ForcePerLengthPoundForcePerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerInch = %v, want %v", converted, 100)
    }
    // Test PoundForcePerYard conversion
    pounds_force_per_yardDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthPoundForcePerYard,
    }
    
    var pounds_force_per_yardResult *units.ForcePerLength
    pounds_force_per_yardResult, err = factory.FromDto(pounds_force_per_yardDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForcePerYard returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_yardResult.Convert(units.ForcePerLengthPoundForcePerYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerYard = %v, want %v", converted, 100)
    }
    // Test KilopoundForcePerFoot conversion
    kilopounds_force_per_footDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthKilopoundForcePerFoot,
    }
    
    var kilopounds_force_per_footResult *units.ForcePerLength
    kilopounds_force_per_footResult, err = factory.FromDto(kilopounds_force_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForcePerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_footResult.Convert(units.ForcePerLengthKilopoundForcePerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerFoot = %v, want %v", converted, 100)
    }
    // Test KilopoundForcePerInch conversion
    kilopounds_force_per_inchDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthKilopoundForcePerInch,
    }
    
    var kilopounds_force_per_inchResult *units.ForcePerLength
    kilopounds_force_per_inchResult, err = factory.FromDto(kilopounds_force_per_inchDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForcePerInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_inchResult.Convert(units.ForcePerLengthKilopoundForcePerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerInch = %v, want %v", converted, 100)
    }
    // Test NanonewtonPerMeter conversion
    nanonewtons_per_meterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthNanonewtonPerMeter,
    }
    
    var nanonewtons_per_meterResult *units.ForcePerLength
    nanonewtons_per_meterResult, err = factory.FromDto(nanonewtons_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with NanonewtonPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewtons_per_meterResult.Convert(units.ForcePerLengthNanonewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test MicronewtonPerMeter conversion
    micronewtons_per_meterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthMicronewtonPerMeter,
    }
    
    var micronewtons_per_meterResult *units.ForcePerLength
    micronewtons_per_meterResult, err = factory.FromDto(micronewtons_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MicronewtonPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewtons_per_meterResult.Convert(units.ForcePerLengthMicronewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test MillinewtonPerMeter conversion
    millinewtons_per_meterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthMillinewtonPerMeter,
    }
    
    var millinewtons_per_meterResult *units.ForcePerLength
    millinewtons_per_meterResult, err = factory.FromDto(millinewtons_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MillinewtonPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewtons_per_meterResult.Convert(units.ForcePerLengthMillinewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test CentinewtonPerMeter conversion
    centinewtons_per_meterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthCentinewtonPerMeter,
    }
    
    var centinewtons_per_meterResult *units.ForcePerLength
    centinewtons_per_meterResult, err = factory.FromDto(centinewtons_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with CentinewtonPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewtons_per_meterResult.Convert(units.ForcePerLengthCentinewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test DecinewtonPerMeter conversion
    decinewtons_per_meterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthDecinewtonPerMeter,
    }
    
    var decinewtons_per_meterResult *units.ForcePerLength
    decinewtons_per_meterResult, err = factory.FromDto(decinewtons_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with DecinewtonPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewtons_per_meterResult.Convert(units.ForcePerLengthDecinewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test DecanewtonPerMeter conversion
    decanewtons_per_meterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthDecanewtonPerMeter,
    }
    
    var decanewtons_per_meterResult *units.ForcePerLength
    decanewtons_per_meterResult, err = factory.FromDto(decanewtons_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with DecanewtonPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewtons_per_meterResult.Convert(units.ForcePerLengthDecanewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonPerMeter conversion
    kilonewtons_per_meterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthKilonewtonPerMeter,
    }
    
    var kilonewtons_per_meterResult *units.ForcePerLength
    kilonewtons_per_meterResult, err = factory.FromDto(kilonewtons_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_meterResult.Convert(units.ForcePerLengthKilonewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test MeganewtonPerMeter conversion
    meganewtons_per_meterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthMeganewtonPerMeter,
    }
    
    var meganewtons_per_meterResult *units.ForcePerLength
    meganewtons_per_meterResult, err = factory.FromDto(meganewtons_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewtons_per_meterResult.Convert(units.ForcePerLengthMeganewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test NanonewtonPerCentimeter conversion
    nanonewtons_per_centimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthNanonewtonPerCentimeter,
    }
    
    var nanonewtons_per_centimeterResult *units.ForcePerLength
    nanonewtons_per_centimeterResult, err = factory.FromDto(nanonewtons_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with NanonewtonPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewtons_per_centimeterResult.Convert(units.ForcePerLengthNanonewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test MicronewtonPerCentimeter conversion
    micronewtons_per_centimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthMicronewtonPerCentimeter,
    }
    
    var micronewtons_per_centimeterResult *units.ForcePerLength
    micronewtons_per_centimeterResult, err = factory.FromDto(micronewtons_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MicronewtonPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewtons_per_centimeterResult.Convert(units.ForcePerLengthMicronewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test MillinewtonPerCentimeter conversion
    millinewtons_per_centimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthMillinewtonPerCentimeter,
    }
    
    var millinewtons_per_centimeterResult *units.ForcePerLength
    millinewtons_per_centimeterResult, err = factory.FromDto(millinewtons_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MillinewtonPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewtons_per_centimeterResult.Convert(units.ForcePerLengthMillinewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test CentinewtonPerCentimeter conversion
    centinewtons_per_centimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthCentinewtonPerCentimeter,
    }
    
    var centinewtons_per_centimeterResult *units.ForcePerLength
    centinewtons_per_centimeterResult, err = factory.FromDto(centinewtons_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with CentinewtonPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewtons_per_centimeterResult.Convert(units.ForcePerLengthCentinewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test DecinewtonPerCentimeter conversion
    decinewtons_per_centimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthDecinewtonPerCentimeter,
    }
    
    var decinewtons_per_centimeterResult *units.ForcePerLength
    decinewtons_per_centimeterResult, err = factory.FromDto(decinewtons_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with DecinewtonPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewtons_per_centimeterResult.Convert(units.ForcePerLengthDecinewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test DecanewtonPerCentimeter conversion
    decanewtons_per_centimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthDecanewtonPerCentimeter,
    }
    
    var decanewtons_per_centimeterResult *units.ForcePerLength
    decanewtons_per_centimeterResult, err = factory.FromDto(decanewtons_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with DecanewtonPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewtons_per_centimeterResult.Convert(units.ForcePerLengthDecanewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonPerCentimeter conversion
    kilonewtons_per_centimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthKilonewtonPerCentimeter,
    }
    
    var kilonewtons_per_centimeterResult *units.ForcePerLength
    kilonewtons_per_centimeterResult, err = factory.FromDto(kilonewtons_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_centimeterResult.Convert(units.ForcePerLengthKilonewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test MeganewtonPerCentimeter conversion
    meganewtons_per_centimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthMeganewtonPerCentimeter,
    }
    
    var meganewtons_per_centimeterResult *units.ForcePerLength
    meganewtons_per_centimeterResult, err = factory.FromDto(meganewtons_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewtons_per_centimeterResult.Convert(units.ForcePerLengthMeganewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test NanonewtonPerMillimeter conversion
    nanonewtons_per_millimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthNanonewtonPerMillimeter,
    }
    
    var nanonewtons_per_millimeterResult *units.ForcePerLength
    nanonewtons_per_millimeterResult, err = factory.FromDto(nanonewtons_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with NanonewtonPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewtons_per_millimeterResult.Convert(units.ForcePerLengthNanonewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test MicronewtonPerMillimeter conversion
    micronewtons_per_millimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthMicronewtonPerMillimeter,
    }
    
    var micronewtons_per_millimeterResult *units.ForcePerLength
    micronewtons_per_millimeterResult, err = factory.FromDto(micronewtons_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MicronewtonPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewtons_per_millimeterResult.Convert(units.ForcePerLengthMicronewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test MillinewtonPerMillimeter conversion
    millinewtons_per_millimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthMillinewtonPerMillimeter,
    }
    
    var millinewtons_per_millimeterResult *units.ForcePerLength
    millinewtons_per_millimeterResult, err = factory.FromDto(millinewtons_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MillinewtonPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewtons_per_millimeterResult.Convert(units.ForcePerLengthMillinewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test CentinewtonPerMillimeter conversion
    centinewtons_per_millimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthCentinewtonPerMillimeter,
    }
    
    var centinewtons_per_millimeterResult *units.ForcePerLength
    centinewtons_per_millimeterResult, err = factory.FromDto(centinewtons_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with CentinewtonPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewtons_per_millimeterResult.Convert(units.ForcePerLengthCentinewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test DecinewtonPerMillimeter conversion
    decinewtons_per_millimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthDecinewtonPerMillimeter,
    }
    
    var decinewtons_per_millimeterResult *units.ForcePerLength
    decinewtons_per_millimeterResult, err = factory.FromDto(decinewtons_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with DecinewtonPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewtons_per_millimeterResult.Convert(units.ForcePerLengthDecinewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test DecanewtonPerMillimeter conversion
    decanewtons_per_millimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthDecanewtonPerMillimeter,
    }
    
    var decanewtons_per_millimeterResult *units.ForcePerLength
    decanewtons_per_millimeterResult, err = factory.FromDto(decanewtons_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with DecanewtonPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewtons_per_millimeterResult.Convert(units.ForcePerLengthDecanewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonPerMillimeter conversion
    kilonewtons_per_millimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthKilonewtonPerMillimeter,
    }
    
    var kilonewtons_per_millimeterResult *units.ForcePerLength
    kilonewtons_per_millimeterResult, err = factory.FromDto(kilonewtons_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_millimeterResult.Convert(units.ForcePerLengthKilonewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test MeganewtonPerMillimeter conversion
    meganewtons_per_millimeterDto := units.ForcePerLengthDto{
        Value: 100,
        Unit:  units.ForcePerLengthMeganewtonPerMillimeter,
    }
    
    var meganewtons_per_millimeterResult *units.ForcePerLength
    meganewtons_per_millimeterResult, err = factory.FromDto(meganewtons_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewtons_per_millimeterResult.Convert(units.ForcePerLengthMeganewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonPerMillimeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ForcePerLengthDto{
        Value: 0,
        Unit:  units.ForcePerLengthNewtonPerMeter,
    }
    
    var zeroResult *units.ForcePerLength
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestForcePerLengthFactory_FromDtoJSON(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "NewtonPerMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "NewtonPerMeter"}`)
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
    nanJSON, _ := json.Marshal(units.ForcePerLengthDto{
        Value: nanValue,
        Unit:  units.ForcePerLengthNewtonPerMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with NewtonPerMeter unit
    newtons_per_meterJSON := []byte(`{"value": 100, "unit": "NewtonPerMeter"}`)
    newtons_per_meterResult, err := factory.FromDtoJSON(newtons_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_meterResult.Convert(units.ForcePerLengthNewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonPerCentimeter unit
    newtons_per_centimeterJSON := []byte(`{"value": 100, "unit": "NewtonPerCentimeter"}`)
    newtons_per_centimeterResult, err := factory.FromDtoJSON(newtons_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_centimeterResult.Convert(units.ForcePerLengthNewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonPerMillimeter unit
    newtons_per_millimeterJSON := []byte(`{"value": 100, "unit": "NewtonPerMillimeter"}`)
    newtons_per_millimeterResult, err := factory.FromDtoJSON(newtons_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_millimeterResult.Convert(units.ForcePerLengthNewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForcePerMeter unit
    kilograms_force_per_meterJSON := []byte(`{"value": 100, "unit": "KilogramForcePerMeter"}`)
    kilograms_force_per_meterResult, err := factory.FromDtoJSON(kilograms_force_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForcePerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_meterResult.Convert(units.ForcePerLengthKilogramForcePerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForcePerCentimeter unit
    kilograms_force_per_centimeterJSON := []byte(`{"value": 100, "unit": "KilogramForcePerCentimeter"}`)
    kilograms_force_per_centimeterResult, err := factory.FromDtoJSON(kilograms_force_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForcePerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_centimeterResult.Convert(units.ForcePerLengthKilogramForcePerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForcePerMillimeter unit
    kilograms_force_per_millimeterJSON := []byte(`{"value": 100, "unit": "KilogramForcePerMillimeter"}`)
    kilograms_force_per_millimeterResult, err := factory.FromDtoJSON(kilograms_force_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForcePerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_millimeterResult.Convert(units.ForcePerLengthKilogramForcePerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForcePerMeter unit
    tonnes_force_per_meterJSON := []byte(`{"value": 100, "unit": "TonneForcePerMeter"}`)
    tonnes_force_per_meterResult, err := factory.FromDtoJSON(tonnes_force_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForcePerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_meterResult.Convert(units.ForcePerLengthTonneForcePerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForcePerCentimeter unit
    tonnes_force_per_centimeterJSON := []byte(`{"value": 100, "unit": "TonneForcePerCentimeter"}`)
    tonnes_force_per_centimeterResult, err := factory.FromDtoJSON(tonnes_force_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForcePerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_centimeterResult.Convert(units.ForcePerLengthTonneForcePerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForcePerMillimeter unit
    tonnes_force_per_millimeterJSON := []byte(`{"value": 100, "unit": "TonneForcePerMillimeter"}`)
    tonnes_force_per_millimeterResult, err := factory.FromDtoJSON(tonnes_force_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForcePerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_millimeterResult.Convert(units.ForcePerLengthTonneForcePerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForcePerFoot unit
    pounds_force_per_footJSON := []byte(`{"value": 100, "unit": "PoundForcePerFoot"}`)
    pounds_force_per_footResult, err := factory.FromDtoJSON(pounds_force_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForcePerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_footResult.Convert(units.ForcePerLengthPoundForcePerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForcePerInch unit
    pounds_force_per_inchJSON := []byte(`{"value": 100, "unit": "PoundForcePerInch"}`)
    pounds_force_per_inchResult, err := factory.FromDtoJSON(pounds_force_per_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForcePerInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_inchResult.Convert(units.ForcePerLengthPoundForcePerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerInch = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForcePerYard unit
    pounds_force_per_yardJSON := []byte(`{"value": 100, "unit": "PoundForcePerYard"}`)
    pounds_force_per_yardResult, err := factory.FromDtoJSON(pounds_force_per_yardJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForcePerYard unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_yardResult.Convert(units.ForcePerLengthPoundForcePerYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerYard = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForcePerFoot unit
    kilopounds_force_per_footJSON := []byte(`{"value": 100, "unit": "KilopoundForcePerFoot"}`)
    kilopounds_force_per_footResult, err := factory.FromDtoJSON(kilopounds_force_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForcePerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_footResult.Convert(units.ForcePerLengthKilopoundForcePerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForcePerInch unit
    kilopounds_force_per_inchJSON := []byte(`{"value": 100, "unit": "KilopoundForcePerInch"}`)
    kilopounds_force_per_inchResult, err := factory.FromDtoJSON(kilopounds_force_per_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForcePerInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_inchResult.Convert(units.ForcePerLengthKilopoundForcePerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerInch = %v, want %v", converted, 100)
    }
    // Test JSON with NanonewtonPerMeter unit
    nanonewtons_per_meterJSON := []byte(`{"value": 100, "unit": "NanonewtonPerMeter"}`)
    nanonewtons_per_meterResult, err := factory.FromDtoJSON(nanonewtons_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanonewtonPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewtons_per_meterResult.Convert(units.ForcePerLengthNanonewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicronewtonPerMeter unit
    micronewtons_per_meterJSON := []byte(`{"value": 100, "unit": "MicronewtonPerMeter"}`)
    micronewtons_per_meterResult, err := factory.FromDtoJSON(micronewtons_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicronewtonPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewtons_per_meterResult.Convert(units.ForcePerLengthMicronewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MillinewtonPerMeter unit
    millinewtons_per_meterJSON := []byte(`{"value": 100, "unit": "MillinewtonPerMeter"}`)
    millinewtons_per_meterResult, err := factory.FromDtoJSON(millinewtons_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillinewtonPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewtons_per_meterResult.Convert(units.ForcePerLengthMillinewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with CentinewtonPerMeter unit
    centinewtons_per_meterJSON := []byte(`{"value": 100, "unit": "CentinewtonPerMeter"}`)
    centinewtons_per_meterResult, err := factory.FromDtoJSON(centinewtons_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentinewtonPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewtons_per_meterResult.Convert(units.ForcePerLengthCentinewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with DecinewtonPerMeter unit
    decinewtons_per_meterJSON := []byte(`{"value": 100, "unit": "DecinewtonPerMeter"}`)
    decinewtons_per_meterResult, err := factory.FromDtoJSON(decinewtons_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecinewtonPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewtons_per_meterResult.Convert(units.ForcePerLengthDecinewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with DecanewtonPerMeter unit
    decanewtons_per_meterJSON := []byte(`{"value": 100, "unit": "DecanewtonPerMeter"}`)
    decanewtons_per_meterResult, err := factory.FromDtoJSON(decanewtons_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecanewtonPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewtons_per_meterResult.Convert(units.ForcePerLengthDecanewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonPerMeter unit
    kilonewtons_per_meterJSON := []byte(`{"value": 100, "unit": "KilonewtonPerMeter"}`)
    kilonewtons_per_meterResult, err := factory.FromDtoJSON(kilonewtons_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_meterResult.Convert(units.ForcePerLengthKilonewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonPerMeter unit
    meganewtons_per_meterJSON := []byte(`{"value": 100, "unit": "MeganewtonPerMeter"}`)
    meganewtons_per_meterResult, err := factory.FromDtoJSON(meganewtons_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewtons_per_meterResult.Convert(units.ForcePerLengthMeganewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with NanonewtonPerCentimeter unit
    nanonewtons_per_centimeterJSON := []byte(`{"value": 100, "unit": "NanonewtonPerCentimeter"}`)
    nanonewtons_per_centimeterResult, err := factory.FromDtoJSON(nanonewtons_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanonewtonPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewtons_per_centimeterResult.Convert(units.ForcePerLengthNanonewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicronewtonPerCentimeter unit
    micronewtons_per_centimeterJSON := []byte(`{"value": 100, "unit": "MicronewtonPerCentimeter"}`)
    micronewtons_per_centimeterResult, err := factory.FromDtoJSON(micronewtons_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicronewtonPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewtons_per_centimeterResult.Convert(units.ForcePerLengthMicronewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MillinewtonPerCentimeter unit
    millinewtons_per_centimeterJSON := []byte(`{"value": 100, "unit": "MillinewtonPerCentimeter"}`)
    millinewtons_per_centimeterResult, err := factory.FromDtoJSON(millinewtons_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillinewtonPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewtons_per_centimeterResult.Convert(units.ForcePerLengthMillinewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with CentinewtonPerCentimeter unit
    centinewtons_per_centimeterJSON := []byte(`{"value": 100, "unit": "CentinewtonPerCentimeter"}`)
    centinewtons_per_centimeterResult, err := factory.FromDtoJSON(centinewtons_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentinewtonPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewtons_per_centimeterResult.Convert(units.ForcePerLengthCentinewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with DecinewtonPerCentimeter unit
    decinewtons_per_centimeterJSON := []byte(`{"value": 100, "unit": "DecinewtonPerCentimeter"}`)
    decinewtons_per_centimeterResult, err := factory.FromDtoJSON(decinewtons_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecinewtonPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewtons_per_centimeterResult.Convert(units.ForcePerLengthDecinewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with DecanewtonPerCentimeter unit
    decanewtons_per_centimeterJSON := []byte(`{"value": 100, "unit": "DecanewtonPerCentimeter"}`)
    decanewtons_per_centimeterResult, err := factory.FromDtoJSON(decanewtons_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecanewtonPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewtons_per_centimeterResult.Convert(units.ForcePerLengthDecanewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonPerCentimeter unit
    kilonewtons_per_centimeterJSON := []byte(`{"value": 100, "unit": "KilonewtonPerCentimeter"}`)
    kilonewtons_per_centimeterResult, err := factory.FromDtoJSON(kilonewtons_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_centimeterResult.Convert(units.ForcePerLengthKilonewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonPerCentimeter unit
    meganewtons_per_centimeterJSON := []byte(`{"value": 100, "unit": "MeganewtonPerCentimeter"}`)
    meganewtons_per_centimeterResult, err := factory.FromDtoJSON(meganewtons_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewtons_per_centimeterResult.Convert(units.ForcePerLengthMeganewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with NanonewtonPerMillimeter unit
    nanonewtons_per_millimeterJSON := []byte(`{"value": 100, "unit": "NanonewtonPerMillimeter"}`)
    nanonewtons_per_millimeterResult, err := factory.FromDtoJSON(nanonewtons_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanonewtonPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanonewtons_per_millimeterResult.Convert(units.ForcePerLengthNanonewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanonewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicronewtonPerMillimeter unit
    micronewtons_per_millimeterJSON := []byte(`{"value": 100, "unit": "MicronewtonPerMillimeter"}`)
    micronewtons_per_millimeterResult, err := factory.FromDtoJSON(micronewtons_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicronewtonPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewtons_per_millimeterResult.Convert(units.ForcePerLengthMicronewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicronewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MillinewtonPerMillimeter unit
    millinewtons_per_millimeterJSON := []byte(`{"value": 100, "unit": "MillinewtonPerMillimeter"}`)
    millinewtons_per_millimeterResult, err := factory.FromDtoJSON(millinewtons_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillinewtonPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewtons_per_millimeterResult.Convert(units.ForcePerLengthMillinewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillinewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with CentinewtonPerMillimeter unit
    centinewtons_per_millimeterJSON := []byte(`{"value": 100, "unit": "CentinewtonPerMillimeter"}`)
    centinewtons_per_millimeterResult, err := factory.FromDtoJSON(centinewtons_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentinewtonPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centinewtons_per_millimeterResult.Convert(units.ForcePerLengthCentinewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentinewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with DecinewtonPerMillimeter unit
    decinewtons_per_millimeterJSON := []byte(`{"value": 100, "unit": "DecinewtonPerMillimeter"}`)
    decinewtons_per_millimeterResult, err := factory.FromDtoJSON(decinewtons_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecinewtonPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decinewtons_per_millimeterResult.Convert(units.ForcePerLengthDecinewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecinewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with DecanewtonPerMillimeter unit
    decanewtons_per_millimeterJSON := []byte(`{"value": 100, "unit": "DecanewtonPerMillimeter"}`)
    decanewtons_per_millimeterResult, err := factory.FromDtoJSON(decanewtons_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecanewtonPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewtons_per_millimeterResult.Convert(units.ForcePerLengthDecanewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecanewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonPerMillimeter unit
    kilonewtons_per_millimeterJSON := []byte(`{"value": 100, "unit": "KilonewtonPerMillimeter"}`)
    kilonewtons_per_millimeterResult, err := factory.FromDtoJSON(kilonewtons_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_millimeterResult.Convert(units.ForcePerLengthKilonewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonPerMillimeter unit
    meganewtons_per_millimeterJSON := []byte(`{"value": 100, "unit": "MeganewtonPerMillimeter"}`)
    meganewtons_per_millimeterResult, err := factory.FromDtoJSON(meganewtons_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewtons_per_millimeterResult.Convert(units.ForcePerLengthMeganewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonPerMillimeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "NewtonPerMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromNewtonsPerMeter function
func TestForcePerLengthFactory_FromNewtonsPerMeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonsPerMeter(100)
    if err != nil {
        t.Errorf("FromNewtonsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthNewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromNewtonsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromNewtonsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromNewtonsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonsPerMeter(0)
    if err != nil {
        t.Errorf("FromNewtonsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthNewtonPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonsPerCentimeter function
func TestForcePerLengthFactory_FromNewtonsPerCentimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromNewtonsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthNewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromNewtonsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromNewtonsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromNewtonsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromNewtonsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthNewtonPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonsPerMillimeter function
func TestForcePerLengthFactory_FromNewtonsPerMillimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromNewtonsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthNewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromNewtonsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromNewtonsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromNewtonsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromNewtonsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthNewtonPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsForcePerMeter function
func TestForcePerLengthFactory_FromKilogramsForcePerMeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsForcePerMeter(100)
    if err != nil {
        t.Errorf("FromKilogramsForcePerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthKilogramForcePerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsForcePerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsForcePerMeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsForcePerMeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsForcePerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsForcePerMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsForcePerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsForcePerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsForcePerMeter(0)
    if err != nil {
        t.Errorf("FromKilogramsForcePerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthKilogramForcePerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsForcePerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsForcePerCentimeter function
func TestForcePerLengthFactory_FromKilogramsForcePerCentimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsForcePerCentimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsForcePerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthKilogramForcePerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsForcePerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsForcePerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsForcePerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsForcePerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsForcePerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsForcePerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsForcePerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsForcePerCentimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsForcePerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthKilogramForcePerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsForcePerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsForcePerMillimeter function
func TestForcePerLengthFactory_FromKilogramsForcePerMillimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsForcePerMillimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsForcePerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthKilogramForcePerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsForcePerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsForcePerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsForcePerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsForcePerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsForcePerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsForcePerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsForcePerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsForcePerMillimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsForcePerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthKilogramForcePerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsForcePerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesForcePerMeter function
func TestForcePerLengthFactory_FromTonnesForcePerMeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesForcePerMeter(100)
    if err != nil {
        t.Errorf("FromTonnesForcePerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthTonneForcePerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesForcePerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesForcePerMeter(math.NaN())
    if err == nil {
        t.Error("FromTonnesForcePerMeter() with NaN value should return error")
    }

    _, err = factory.FromTonnesForcePerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesForcePerMeter() with +Inf value should return error")
    }

    _, err = factory.FromTonnesForcePerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesForcePerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesForcePerMeter(0)
    if err != nil {
        t.Errorf("FromTonnesForcePerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthTonneForcePerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesForcePerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesForcePerCentimeter function
func TestForcePerLengthFactory_FromTonnesForcePerCentimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesForcePerCentimeter(100)
    if err != nil {
        t.Errorf("FromTonnesForcePerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthTonneForcePerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesForcePerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesForcePerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromTonnesForcePerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromTonnesForcePerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesForcePerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromTonnesForcePerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesForcePerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesForcePerCentimeter(0)
    if err != nil {
        t.Errorf("FromTonnesForcePerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthTonneForcePerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesForcePerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesForcePerMillimeter function
func TestForcePerLengthFactory_FromTonnesForcePerMillimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesForcePerMillimeter(100)
    if err != nil {
        t.Errorf("FromTonnesForcePerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthTonneForcePerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesForcePerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesForcePerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromTonnesForcePerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromTonnesForcePerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesForcePerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromTonnesForcePerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesForcePerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesForcePerMillimeter(0)
    if err != nil {
        t.Errorf("FromTonnesForcePerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthTonneForcePerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesForcePerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForcePerFoot function
func TestForcePerLengthFactory_FromPoundsForcePerFoot(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForcePerFoot(100)
    if err != nil {
        t.Errorf("FromPoundsForcePerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthPoundForcePerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForcePerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForcePerFoot(math.NaN())
    if err == nil {
        t.Error("FromPoundsForcePerFoot() with NaN value should return error")
    }

    _, err = factory.FromPoundsForcePerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForcePerFoot() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForcePerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForcePerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForcePerFoot(0)
    if err != nil {
        t.Errorf("FromPoundsForcePerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthPoundForcePerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForcePerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForcePerInch function
func TestForcePerLengthFactory_FromPoundsForcePerInch(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForcePerInch(100)
    if err != nil {
        t.Errorf("FromPoundsForcePerInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthPoundForcePerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForcePerInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForcePerInch(math.NaN())
    if err == nil {
        t.Error("FromPoundsForcePerInch() with NaN value should return error")
    }

    _, err = factory.FromPoundsForcePerInch(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForcePerInch() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForcePerInch(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForcePerInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForcePerInch(0)
    if err != nil {
        t.Errorf("FromPoundsForcePerInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthPoundForcePerInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForcePerInch() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForcePerYard function
func TestForcePerLengthFactory_FromPoundsForcePerYard(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForcePerYard(100)
    if err != nil {
        t.Errorf("FromPoundsForcePerYard() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthPoundForcePerYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForcePerYard() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForcePerYard(math.NaN())
    if err == nil {
        t.Error("FromPoundsForcePerYard() with NaN value should return error")
    }

    _, err = factory.FromPoundsForcePerYard(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForcePerYard() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForcePerYard(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForcePerYard() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForcePerYard(0)
    if err != nil {
        t.Errorf("FromPoundsForcePerYard() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthPoundForcePerYard)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForcePerYard() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsForcePerFoot function
func TestForcePerLengthFactory_FromKilopoundsForcePerFoot(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsForcePerFoot(100)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthKilopoundForcePerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsForcePerFoot(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsForcePerFoot() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsForcePerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsForcePerFoot() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsForcePerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsForcePerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsForcePerFoot(0)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthKilopoundForcePerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsForcePerInch function
func TestForcePerLengthFactory_FromKilopoundsForcePerInch(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsForcePerInch(100)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthKilopoundForcePerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsForcePerInch(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsForcePerInch() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsForcePerInch(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsForcePerInch() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsForcePerInch(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsForcePerInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsForcePerInch(0)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthKilopoundForcePerInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerInch() with zero value = %v, want 0", converted)
    }
}
// Test FromNanonewtonsPerMeter function
func TestForcePerLengthFactory_FromNanonewtonsPerMeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanonewtonsPerMeter(100)
    if err != nil {
        t.Errorf("FromNanonewtonsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthNanonewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanonewtonsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanonewtonsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromNanonewtonsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromNanonewtonsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromNanonewtonsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromNanonewtonsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanonewtonsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanonewtonsPerMeter(0)
    if err != nil {
        t.Errorf("FromNanonewtonsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthNanonewtonPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanonewtonsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicronewtonsPerMeter function
func TestForcePerLengthFactory_FromMicronewtonsPerMeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicronewtonsPerMeter(100)
    if err != nil {
        t.Errorf("FromMicronewtonsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthMicronewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicronewtonsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicronewtonsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromMicronewtonsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromMicronewtonsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicronewtonsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromMicronewtonsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicronewtonsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicronewtonsPerMeter(0)
    if err != nil {
        t.Errorf("FromMicronewtonsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthMicronewtonPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicronewtonsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillinewtonsPerMeter function
func TestForcePerLengthFactory_FromMillinewtonsPerMeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillinewtonsPerMeter(100)
    if err != nil {
        t.Errorf("FromMillinewtonsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthMillinewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillinewtonsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillinewtonsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromMillinewtonsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromMillinewtonsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMillinewtonsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromMillinewtonsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillinewtonsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillinewtonsPerMeter(0)
    if err != nil {
        t.Errorf("FromMillinewtonsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthMillinewtonPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillinewtonsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentinewtonsPerMeter function
func TestForcePerLengthFactory_FromCentinewtonsPerMeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentinewtonsPerMeter(100)
    if err != nil {
        t.Errorf("FromCentinewtonsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthCentinewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentinewtonsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentinewtonsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromCentinewtonsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromCentinewtonsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromCentinewtonsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromCentinewtonsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromCentinewtonsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentinewtonsPerMeter(0)
    if err != nil {
        t.Errorf("FromCentinewtonsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthCentinewtonPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentinewtonsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecinewtonsPerMeter function
func TestForcePerLengthFactory_FromDecinewtonsPerMeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecinewtonsPerMeter(100)
    if err != nil {
        t.Errorf("FromDecinewtonsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthDecinewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecinewtonsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecinewtonsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromDecinewtonsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromDecinewtonsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromDecinewtonsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromDecinewtonsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecinewtonsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecinewtonsPerMeter(0)
    if err != nil {
        t.Errorf("FromDecinewtonsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthDecinewtonPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecinewtonsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecanewtonsPerMeter function
func TestForcePerLengthFactory_FromDecanewtonsPerMeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecanewtonsPerMeter(100)
    if err != nil {
        t.Errorf("FromDecanewtonsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthDecanewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecanewtonsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecanewtonsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromDecanewtonsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromDecanewtonsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromDecanewtonsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromDecanewtonsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecanewtonsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecanewtonsPerMeter(0)
    if err != nil {
        t.Errorf("FromDecanewtonsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthDecanewtonPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecanewtonsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonsPerMeter function
func TestForcePerLengthFactory_FromKilonewtonsPerMeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonsPerMeter(100)
    if err != nil {
        t.Errorf("FromKilonewtonsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthKilonewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonsPerMeter(0)
    if err != nil {
        t.Errorf("FromKilonewtonsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthKilonewtonPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonsPerMeter function
func TestForcePerLengthFactory_FromMeganewtonsPerMeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonsPerMeter(100)
    if err != nil {
        t.Errorf("FromMeganewtonsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthMeganewtonPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonsPerMeter(0)
    if err != nil {
        t.Errorf("FromMeganewtonsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthMeganewtonPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanonewtonsPerCentimeter function
func TestForcePerLengthFactory_FromNanonewtonsPerCentimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanonewtonsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromNanonewtonsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthNanonewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanonewtonsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanonewtonsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromNanonewtonsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromNanonewtonsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromNanonewtonsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromNanonewtonsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanonewtonsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanonewtonsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromNanonewtonsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthNanonewtonPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanonewtonsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicronewtonsPerCentimeter function
func TestForcePerLengthFactory_FromMicronewtonsPerCentimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicronewtonsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromMicronewtonsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthMicronewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicronewtonsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicronewtonsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMicronewtonsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMicronewtonsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicronewtonsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMicronewtonsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicronewtonsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicronewtonsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromMicronewtonsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthMicronewtonPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicronewtonsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillinewtonsPerCentimeter function
func TestForcePerLengthFactory_FromMillinewtonsPerCentimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillinewtonsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromMillinewtonsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthMillinewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillinewtonsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillinewtonsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMillinewtonsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMillinewtonsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMillinewtonsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMillinewtonsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillinewtonsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillinewtonsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromMillinewtonsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthMillinewtonPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillinewtonsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentinewtonsPerCentimeter function
func TestForcePerLengthFactory_FromCentinewtonsPerCentimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentinewtonsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromCentinewtonsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthCentinewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentinewtonsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentinewtonsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromCentinewtonsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromCentinewtonsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromCentinewtonsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromCentinewtonsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromCentinewtonsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentinewtonsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromCentinewtonsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthCentinewtonPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentinewtonsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecinewtonsPerCentimeter function
func TestForcePerLengthFactory_FromDecinewtonsPerCentimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecinewtonsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromDecinewtonsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthDecinewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecinewtonsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecinewtonsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromDecinewtonsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromDecinewtonsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromDecinewtonsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromDecinewtonsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecinewtonsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecinewtonsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromDecinewtonsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthDecinewtonPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecinewtonsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecanewtonsPerCentimeter function
func TestForcePerLengthFactory_FromDecanewtonsPerCentimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecanewtonsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromDecanewtonsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthDecanewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecanewtonsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecanewtonsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromDecanewtonsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromDecanewtonsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromDecanewtonsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromDecanewtonsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecanewtonsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecanewtonsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromDecanewtonsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthDecanewtonPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecanewtonsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonsPerCentimeter function
func TestForcePerLengthFactory_FromKilonewtonsPerCentimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromKilonewtonsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthKilonewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromKilonewtonsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthKilonewtonPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonsPerCentimeter function
func TestForcePerLengthFactory_FromMeganewtonsPerCentimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromMeganewtonsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthMeganewtonPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromMeganewtonsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthMeganewtonPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanonewtonsPerMillimeter function
func TestForcePerLengthFactory_FromNanonewtonsPerMillimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanonewtonsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromNanonewtonsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthNanonewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanonewtonsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanonewtonsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromNanonewtonsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromNanonewtonsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromNanonewtonsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromNanonewtonsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanonewtonsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanonewtonsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromNanonewtonsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthNanonewtonPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanonewtonsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicronewtonsPerMillimeter function
func TestForcePerLengthFactory_FromMicronewtonsPerMillimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicronewtonsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromMicronewtonsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthMicronewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicronewtonsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicronewtonsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromMicronewtonsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromMicronewtonsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicronewtonsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromMicronewtonsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicronewtonsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicronewtonsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromMicronewtonsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthMicronewtonPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicronewtonsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillinewtonsPerMillimeter function
func TestForcePerLengthFactory_FromMillinewtonsPerMillimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillinewtonsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromMillinewtonsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthMillinewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillinewtonsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillinewtonsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromMillinewtonsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromMillinewtonsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMillinewtonsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromMillinewtonsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillinewtonsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillinewtonsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromMillinewtonsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthMillinewtonPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillinewtonsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentinewtonsPerMillimeter function
func TestForcePerLengthFactory_FromCentinewtonsPerMillimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentinewtonsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromCentinewtonsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthCentinewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentinewtonsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentinewtonsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromCentinewtonsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromCentinewtonsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromCentinewtonsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromCentinewtonsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromCentinewtonsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentinewtonsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromCentinewtonsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthCentinewtonPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentinewtonsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecinewtonsPerMillimeter function
func TestForcePerLengthFactory_FromDecinewtonsPerMillimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecinewtonsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromDecinewtonsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthDecinewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecinewtonsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecinewtonsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromDecinewtonsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromDecinewtonsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromDecinewtonsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromDecinewtonsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecinewtonsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecinewtonsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromDecinewtonsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthDecinewtonPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecinewtonsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecanewtonsPerMillimeter function
func TestForcePerLengthFactory_FromDecanewtonsPerMillimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecanewtonsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromDecanewtonsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthDecanewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecanewtonsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecanewtonsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromDecanewtonsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromDecanewtonsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromDecanewtonsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromDecanewtonsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecanewtonsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecanewtonsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromDecanewtonsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthDecanewtonPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecanewtonsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonsPerMillimeter function
func TestForcePerLengthFactory_FromKilonewtonsPerMillimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromKilonewtonsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthKilonewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromKilonewtonsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthKilonewtonPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonsPerMillimeter function
func TestForcePerLengthFactory_FromMeganewtonsPerMillimeter(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromMeganewtonsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePerLengthMeganewtonPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromMeganewtonsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePerLengthMeganewtonPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonsPerMillimeter() with zero value = %v, want 0", converted)
    }
}

func TestForcePerLengthToString(t *testing.T) {
	factory := units.ForcePerLengthFactory{}
	a, err := factory.CreateForcePerLength(45, units.ForcePerLengthNewtonPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ForcePerLengthNewtonPerMeter, 2)
	expected := "45.00 " + units.GetForcePerLengthAbbreviation(units.ForcePerLengthNewtonPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ForcePerLengthNewtonPerMeter, -1)
	expected = "45 " + units.GetForcePerLengthAbbreviation(units.ForcePerLengthNewtonPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestForcePerLength_EqualityAndComparison(t *testing.T) {
	factory := units.ForcePerLengthFactory{}
	a1, _ := factory.CreateForcePerLength(60, units.ForcePerLengthNewtonPerMeter)
	a2, _ := factory.CreateForcePerLength(60, units.ForcePerLengthNewtonPerMeter)
	a3, _ := factory.CreateForcePerLength(120, units.ForcePerLengthNewtonPerMeter)

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

func TestForcePerLength_Arithmetic(t *testing.T) {
	factory := units.ForcePerLengthFactory{}
	a1, _ := factory.CreateForcePerLength(30, units.ForcePerLengthNewtonPerMeter)
	a2, _ := factory.CreateForcePerLength(45, units.ForcePerLengthNewtonPerMeter)

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


func TestGetForcePerLengthAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ForcePerLengthUnits
        want string
    }{
        {
            name: "NewtonPerMeter abbreviation",
            unit: units.ForcePerLengthNewtonPerMeter,
            want: "N/m",
        },
        {
            name: "NewtonPerCentimeter abbreviation",
            unit: units.ForcePerLengthNewtonPerCentimeter,
            want: "N/cm",
        },
        {
            name: "NewtonPerMillimeter abbreviation",
            unit: units.ForcePerLengthNewtonPerMillimeter,
            want: "N/mm",
        },
        {
            name: "KilogramForcePerMeter abbreviation",
            unit: units.ForcePerLengthKilogramForcePerMeter,
            want: "kgf/m",
        },
        {
            name: "KilogramForcePerCentimeter abbreviation",
            unit: units.ForcePerLengthKilogramForcePerCentimeter,
            want: "kgf/cm",
        },
        {
            name: "KilogramForcePerMillimeter abbreviation",
            unit: units.ForcePerLengthKilogramForcePerMillimeter,
            want: "kgf/mm",
        },
        {
            name: "TonneForcePerMeter abbreviation",
            unit: units.ForcePerLengthTonneForcePerMeter,
            want: "tf/m",
        },
        {
            name: "TonneForcePerCentimeter abbreviation",
            unit: units.ForcePerLengthTonneForcePerCentimeter,
            want: "tf/cm",
        },
        {
            name: "TonneForcePerMillimeter abbreviation",
            unit: units.ForcePerLengthTonneForcePerMillimeter,
            want: "tf/mm",
        },
        {
            name: "PoundForcePerFoot abbreviation",
            unit: units.ForcePerLengthPoundForcePerFoot,
            want: "lbf/ft",
        },
        {
            name: "PoundForcePerInch abbreviation",
            unit: units.ForcePerLengthPoundForcePerInch,
            want: "lbf/in",
        },
        {
            name: "PoundForcePerYard abbreviation",
            unit: units.ForcePerLengthPoundForcePerYard,
            want: "lbf/yd",
        },
        {
            name: "KilopoundForcePerFoot abbreviation",
            unit: units.ForcePerLengthKilopoundForcePerFoot,
            want: "kipf/ft",
        },
        {
            name: "KilopoundForcePerInch abbreviation",
            unit: units.ForcePerLengthKilopoundForcePerInch,
            want: "kipf/in",
        },
        {
            name: "NanonewtonPerMeter abbreviation",
            unit: units.ForcePerLengthNanonewtonPerMeter,
            want: "nN/m",
        },
        {
            name: "MicronewtonPerMeter abbreviation",
            unit: units.ForcePerLengthMicronewtonPerMeter,
            want: "μN/m",
        },
        {
            name: "MillinewtonPerMeter abbreviation",
            unit: units.ForcePerLengthMillinewtonPerMeter,
            want: "mN/m",
        },
        {
            name: "CentinewtonPerMeter abbreviation",
            unit: units.ForcePerLengthCentinewtonPerMeter,
            want: "cN/m",
        },
        {
            name: "DecinewtonPerMeter abbreviation",
            unit: units.ForcePerLengthDecinewtonPerMeter,
            want: "dN/m",
        },
        {
            name: "DecanewtonPerMeter abbreviation",
            unit: units.ForcePerLengthDecanewtonPerMeter,
            want: "daN/m",
        },
        {
            name: "KilonewtonPerMeter abbreviation",
            unit: units.ForcePerLengthKilonewtonPerMeter,
            want: "kN/m",
        },
        {
            name: "MeganewtonPerMeter abbreviation",
            unit: units.ForcePerLengthMeganewtonPerMeter,
            want: "MN/m",
        },
        {
            name: "NanonewtonPerCentimeter abbreviation",
            unit: units.ForcePerLengthNanonewtonPerCentimeter,
            want: "nN/cm",
        },
        {
            name: "MicronewtonPerCentimeter abbreviation",
            unit: units.ForcePerLengthMicronewtonPerCentimeter,
            want: "μN/cm",
        },
        {
            name: "MillinewtonPerCentimeter abbreviation",
            unit: units.ForcePerLengthMillinewtonPerCentimeter,
            want: "mN/cm",
        },
        {
            name: "CentinewtonPerCentimeter abbreviation",
            unit: units.ForcePerLengthCentinewtonPerCentimeter,
            want: "cN/cm",
        },
        {
            name: "DecinewtonPerCentimeter abbreviation",
            unit: units.ForcePerLengthDecinewtonPerCentimeter,
            want: "dN/cm",
        },
        {
            name: "DecanewtonPerCentimeter abbreviation",
            unit: units.ForcePerLengthDecanewtonPerCentimeter,
            want: "daN/cm",
        },
        {
            name: "KilonewtonPerCentimeter abbreviation",
            unit: units.ForcePerLengthKilonewtonPerCentimeter,
            want: "kN/cm",
        },
        {
            name: "MeganewtonPerCentimeter abbreviation",
            unit: units.ForcePerLengthMeganewtonPerCentimeter,
            want: "MN/cm",
        },
        {
            name: "NanonewtonPerMillimeter abbreviation",
            unit: units.ForcePerLengthNanonewtonPerMillimeter,
            want: "nN/mm",
        },
        {
            name: "MicronewtonPerMillimeter abbreviation",
            unit: units.ForcePerLengthMicronewtonPerMillimeter,
            want: "μN/mm",
        },
        {
            name: "MillinewtonPerMillimeter abbreviation",
            unit: units.ForcePerLengthMillinewtonPerMillimeter,
            want: "mN/mm",
        },
        {
            name: "CentinewtonPerMillimeter abbreviation",
            unit: units.ForcePerLengthCentinewtonPerMillimeter,
            want: "cN/mm",
        },
        {
            name: "DecinewtonPerMillimeter abbreviation",
            unit: units.ForcePerLengthDecinewtonPerMillimeter,
            want: "dN/mm",
        },
        {
            name: "DecanewtonPerMillimeter abbreviation",
            unit: units.ForcePerLengthDecanewtonPerMillimeter,
            want: "daN/mm",
        },
        {
            name: "KilonewtonPerMillimeter abbreviation",
            unit: units.ForcePerLengthKilonewtonPerMillimeter,
            want: "kN/mm",
        },
        {
            name: "MeganewtonPerMillimeter abbreviation",
            unit: units.ForcePerLengthMeganewtonPerMillimeter,
            want: "MN/mm",
        },
        {
            name: "invalid unit",
            unit: units.ForcePerLengthUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetForcePerLengthAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetForcePerLengthAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestForcePerLength_String(t *testing.T) {
    factory := units.ForcePerLengthFactory{}
    
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
            unit, err := factory.CreateForcePerLength(tt.value, units.ForcePerLengthNewtonPerMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ForcePerLength.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestForcePerLength_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.ForcePerLengthFactory{}

	_, err := uf.CreateForcePerLength(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}