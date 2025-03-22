// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLinearPowerDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "WattPerMeter"}`
	
	factory := units.LinearPowerDensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LinearPowerDensityWattPerMeter {
		t.Errorf("expected unit %v, got %v", units.LinearPowerDensityWattPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "WattPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLinearPowerDensityDto_ToJSON(t *testing.T) {
	dto := units.LinearPowerDensityDto{
		Value: 45,
		Unit:  units.LinearPowerDensityWattPerMeter,
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
	if result["unit"].(string) != string(units.LinearPowerDensityWattPerMeter) {
		t.Errorf("expected unit %s, got %v", units.LinearPowerDensityWattPerMeter, result["unit"])
	}
}

func TestNewLinearPowerDensity_InvalidValue(t *testing.T) {
	factory := units.LinearPowerDensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLinearPowerDensity(math.NaN(), units.LinearPowerDensityWattPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLinearPowerDensity(math.Inf(1), units.LinearPowerDensityWattPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLinearPowerDensityConversions(t *testing.T) {
	factory := units.LinearPowerDensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLinearPowerDensity(180, units.LinearPowerDensityWattPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to WattsPerMeter.
		// No expected conversion value provided for WattsPerMeter, verifying result is not NaN.
		result := a.WattsPerMeter()
		cacheResult := a.WattsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to WattsPerCentimeter.
		// No expected conversion value provided for WattsPerCentimeter, verifying result is not NaN.
		result := a.WattsPerCentimeter()
		cacheResult := a.WattsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to WattsPerMillimeter.
		// No expected conversion value provided for WattsPerMillimeter, verifying result is not NaN.
		result := a.WattsPerMillimeter()
		cacheResult := a.WattsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to WattsPerInch.
		// No expected conversion value provided for WattsPerInch, verifying result is not NaN.
		result := a.WattsPerInch()
		cacheResult := a.WattsPerInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattsPerInch returned NaN")
		}
	}
	{
		// Test conversion to WattsPerFoot.
		// No expected conversion value provided for WattsPerFoot, verifying result is not NaN.
		result := a.WattsPerFoot()
		cacheResult := a.WattsPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerMeter.
		// No expected conversion value provided for MilliwattsPerMeter, verifying result is not NaN.
		result := a.MilliwattsPerMeter()
		cacheResult := a.MilliwattsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliwattsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerMeter.
		// No expected conversion value provided for KilowattsPerMeter, verifying result is not NaN.
		result := a.KilowattsPerMeter()
		cacheResult := a.KilowattsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerMeter.
		// No expected conversion value provided for MegawattsPerMeter, verifying result is not NaN.
		result := a.MegawattsPerMeter()
		cacheResult := a.MegawattsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerMeter.
		// No expected conversion value provided for GigawattsPerMeter, verifying result is not NaN.
		result := a.GigawattsPerMeter()
		cacheResult := a.GigawattsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerCentimeter.
		// No expected conversion value provided for MilliwattsPerCentimeter, verifying result is not NaN.
		result := a.MilliwattsPerCentimeter()
		cacheResult := a.MilliwattsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliwattsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerCentimeter.
		// No expected conversion value provided for KilowattsPerCentimeter, verifying result is not NaN.
		result := a.KilowattsPerCentimeter()
		cacheResult := a.KilowattsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerCentimeter.
		// No expected conversion value provided for MegawattsPerCentimeter, verifying result is not NaN.
		result := a.MegawattsPerCentimeter()
		cacheResult := a.MegawattsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerCentimeter.
		// No expected conversion value provided for GigawattsPerCentimeter, verifying result is not NaN.
		result := a.GigawattsPerCentimeter()
		cacheResult := a.GigawattsPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattsPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerMillimeter.
		// No expected conversion value provided for MilliwattsPerMillimeter, verifying result is not NaN.
		result := a.MilliwattsPerMillimeter()
		cacheResult := a.MilliwattsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliwattsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerMillimeter.
		// No expected conversion value provided for KilowattsPerMillimeter, verifying result is not NaN.
		result := a.KilowattsPerMillimeter()
		cacheResult := a.KilowattsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerMillimeter.
		// No expected conversion value provided for MegawattsPerMillimeter, verifying result is not NaN.
		result := a.MegawattsPerMillimeter()
		cacheResult := a.MegawattsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerMillimeter.
		// No expected conversion value provided for GigawattsPerMillimeter, verifying result is not NaN.
		result := a.GigawattsPerMillimeter()
		cacheResult := a.GigawattsPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattsPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerInch.
		// No expected conversion value provided for MilliwattsPerInch, verifying result is not NaN.
		result := a.MilliwattsPerInch()
		cacheResult := a.MilliwattsPerInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliwattsPerInch returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerInch.
		// No expected conversion value provided for KilowattsPerInch, verifying result is not NaN.
		result := a.KilowattsPerInch()
		cacheResult := a.KilowattsPerInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattsPerInch returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerInch.
		// No expected conversion value provided for MegawattsPerInch, verifying result is not NaN.
		result := a.MegawattsPerInch()
		cacheResult := a.MegawattsPerInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattsPerInch returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerInch.
		// No expected conversion value provided for GigawattsPerInch, verifying result is not NaN.
		result := a.GigawattsPerInch()
		cacheResult := a.GigawattsPerInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattsPerInch returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerFoot.
		// No expected conversion value provided for MilliwattsPerFoot, verifying result is not NaN.
		result := a.MilliwattsPerFoot()
		cacheResult := a.MilliwattsPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliwattsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerFoot.
		// No expected conversion value provided for KilowattsPerFoot, verifying result is not NaN.
		result := a.KilowattsPerFoot()
		cacheResult := a.KilowattsPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerFoot.
		// No expected conversion value provided for MegawattsPerFoot, verifying result is not NaN.
		result := a.MegawattsPerFoot()
		cacheResult := a.MegawattsPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerFoot.
		// No expected conversion value provided for GigawattsPerFoot, verifying result is not NaN.
		result := a.GigawattsPerFoot()
		cacheResult := a.GigawattsPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattsPerFoot returned NaN")
		}
	}
}

func TestLinearPowerDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LinearPowerDensityFactory{}
	a, err := factory.CreateLinearPowerDensity(90, units.LinearPowerDensityWattPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LinearPowerDensityWattPerMeter {
		t.Errorf("expected default unit WattPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LinearPowerDensityWattPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LinearPowerDensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LinearPowerDensityWattPerMeter {
		t.Errorf("expected unit WattPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLinearPowerDensityFactory_FromDto(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityWattPerMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.LinearPowerDensityDto{
        Value: math.NaN(),
        Unit:  units.LinearPowerDensityWattPerMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test WattPerMeter conversion
    watts_per_meterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityWattPerMeter,
    }
    
    var watts_per_meterResult *units.LinearPowerDensity
    watts_per_meterResult, err = factory.FromDto(watts_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_meterResult.Convert(units.LinearPowerDensityWattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerMeter = %v, want %v", converted, 100)
    }
    // Test WattPerCentimeter conversion
    watts_per_centimeterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityWattPerCentimeter,
    }
    
    var watts_per_centimeterResult *units.LinearPowerDensity
    watts_per_centimeterResult, err = factory.FromDto(watts_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_centimeterResult.Convert(units.LinearPowerDensityWattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerCentimeter = %v, want %v", converted, 100)
    }
    // Test WattPerMillimeter conversion
    watts_per_millimeterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityWattPerMillimeter,
    }
    
    var watts_per_millimeterResult *units.LinearPowerDensity
    watts_per_millimeterResult, err = factory.FromDto(watts_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_millimeterResult.Convert(units.LinearPowerDensityWattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerMillimeter = %v, want %v", converted, 100)
    }
    // Test WattPerInch conversion
    watts_per_inchDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityWattPerInch,
    }
    
    var watts_per_inchResult *units.LinearPowerDensity
    watts_per_inchResult, err = factory.FromDto(watts_per_inchDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_inchResult.Convert(units.LinearPowerDensityWattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerInch = %v, want %v", converted, 100)
    }
    // Test WattPerFoot conversion
    watts_per_footDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityWattPerFoot,
    }
    
    var watts_per_footResult *units.LinearPowerDensity
    watts_per_footResult, err = factory.FromDto(watts_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_footResult.Convert(units.LinearPowerDensityWattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerFoot = %v, want %v", converted, 100)
    }
    // Test MilliwattPerMeter conversion
    milliwatts_per_meterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityMilliwattPerMeter,
    }
    
    var milliwatts_per_meterResult *units.LinearPowerDensity
    milliwatts_per_meterResult, err = factory.FromDto(milliwatts_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MilliwattPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_meterResult.Convert(units.LinearPowerDensityMilliwattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerMeter = %v, want %v", converted, 100)
    }
    // Test KilowattPerMeter conversion
    kilowatts_per_meterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityKilowattPerMeter,
    }
    
    var kilowatts_per_meterResult *units.LinearPowerDensity
    kilowatts_per_meterResult, err = factory.FromDto(kilowatts_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_meterResult.Convert(units.LinearPowerDensityKilowattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerMeter = %v, want %v", converted, 100)
    }
    // Test MegawattPerMeter conversion
    megawatts_per_meterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityMegawattPerMeter,
    }
    
    var megawatts_per_meterResult *units.LinearPowerDensity
    megawatts_per_meterResult, err = factory.FromDto(megawatts_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_meterResult.Convert(units.LinearPowerDensityMegawattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerMeter = %v, want %v", converted, 100)
    }
    // Test GigawattPerMeter conversion
    gigawatts_per_meterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityGigawattPerMeter,
    }
    
    var gigawatts_per_meterResult *units.LinearPowerDensity
    gigawatts_per_meterResult, err = factory.FromDto(gigawatts_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_meterResult.Convert(units.LinearPowerDensityGigawattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerMeter = %v, want %v", converted, 100)
    }
    // Test MilliwattPerCentimeter conversion
    milliwatts_per_centimeterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityMilliwattPerCentimeter,
    }
    
    var milliwatts_per_centimeterResult *units.LinearPowerDensity
    milliwatts_per_centimeterResult, err = factory.FromDto(milliwatts_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MilliwattPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_centimeterResult.Convert(units.LinearPowerDensityMilliwattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerCentimeter = %v, want %v", converted, 100)
    }
    // Test KilowattPerCentimeter conversion
    kilowatts_per_centimeterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityKilowattPerCentimeter,
    }
    
    var kilowatts_per_centimeterResult *units.LinearPowerDensity
    kilowatts_per_centimeterResult, err = factory.FromDto(kilowatts_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_centimeterResult.Convert(units.LinearPowerDensityKilowattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerCentimeter = %v, want %v", converted, 100)
    }
    // Test MegawattPerCentimeter conversion
    megawatts_per_centimeterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityMegawattPerCentimeter,
    }
    
    var megawatts_per_centimeterResult *units.LinearPowerDensity
    megawatts_per_centimeterResult, err = factory.FromDto(megawatts_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_centimeterResult.Convert(units.LinearPowerDensityMegawattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerCentimeter = %v, want %v", converted, 100)
    }
    // Test GigawattPerCentimeter conversion
    gigawatts_per_centimeterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityGigawattPerCentimeter,
    }
    
    var gigawatts_per_centimeterResult *units.LinearPowerDensity
    gigawatts_per_centimeterResult, err = factory.FromDto(gigawatts_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_centimeterResult.Convert(units.LinearPowerDensityGigawattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerCentimeter = %v, want %v", converted, 100)
    }
    // Test MilliwattPerMillimeter conversion
    milliwatts_per_millimeterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityMilliwattPerMillimeter,
    }
    
    var milliwatts_per_millimeterResult *units.LinearPowerDensity
    milliwatts_per_millimeterResult, err = factory.FromDto(milliwatts_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MilliwattPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_millimeterResult.Convert(units.LinearPowerDensityMilliwattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerMillimeter = %v, want %v", converted, 100)
    }
    // Test KilowattPerMillimeter conversion
    kilowatts_per_millimeterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityKilowattPerMillimeter,
    }
    
    var kilowatts_per_millimeterResult *units.LinearPowerDensity
    kilowatts_per_millimeterResult, err = factory.FromDto(kilowatts_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_millimeterResult.Convert(units.LinearPowerDensityKilowattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerMillimeter = %v, want %v", converted, 100)
    }
    // Test MegawattPerMillimeter conversion
    megawatts_per_millimeterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityMegawattPerMillimeter,
    }
    
    var megawatts_per_millimeterResult *units.LinearPowerDensity
    megawatts_per_millimeterResult, err = factory.FromDto(megawatts_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_millimeterResult.Convert(units.LinearPowerDensityMegawattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerMillimeter = %v, want %v", converted, 100)
    }
    // Test GigawattPerMillimeter conversion
    gigawatts_per_millimeterDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityGigawattPerMillimeter,
    }
    
    var gigawatts_per_millimeterResult *units.LinearPowerDensity
    gigawatts_per_millimeterResult, err = factory.FromDto(gigawatts_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_millimeterResult.Convert(units.LinearPowerDensityGigawattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerMillimeter = %v, want %v", converted, 100)
    }
    // Test MilliwattPerInch conversion
    milliwatts_per_inchDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityMilliwattPerInch,
    }
    
    var milliwatts_per_inchResult *units.LinearPowerDensity
    milliwatts_per_inchResult, err = factory.FromDto(milliwatts_per_inchDto)
    if err != nil {
        t.Errorf("FromDto() with MilliwattPerInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_inchResult.Convert(units.LinearPowerDensityMilliwattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerInch = %v, want %v", converted, 100)
    }
    // Test KilowattPerInch conversion
    kilowatts_per_inchDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityKilowattPerInch,
    }
    
    var kilowatts_per_inchResult *units.LinearPowerDensity
    kilowatts_per_inchResult, err = factory.FromDto(kilowatts_per_inchDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattPerInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_inchResult.Convert(units.LinearPowerDensityKilowattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerInch = %v, want %v", converted, 100)
    }
    // Test MegawattPerInch conversion
    megawatts_per_inchDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityMegawattPerInch,
    }
    
    var megawatts_per_inchResult *units.LinearPowerDensity
    megawatts_per_inchResult, err = factory.FromDto(megawatts_per_inchDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattPerInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_inchResult.Convert(units.LinearPowerDensityMegawattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerInch = %v, want %v", converted, 100)
    }
    // Test GigawattPerInch conversion
    gigawatts_per_inchDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityGigawattPerInch,
    }
    
    var gigawatts_per_inchResult *units.LinearPowerDensity
    gigawatts_per_inchResult, err = factory.FromDto(gigawatts_per_inchDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattPerInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_inchResult.Convert(units.LinearPowerDensityGigawattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerInch = %v, want %v", converted, 100)
    }
    // Test MilliwattPerFoot conversion
    milliwatts_per_footDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityMilliwattPerFoot,
    }
    
    var milliwatts_per_footResult *units.LinearPowerDensity
    milliwatts_per_footResult, err = factory.FromDto(milliwatts_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with MilliwattPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_footResult.Convert(units.LinearPowerDensityMilliwattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerFoot = %v, want %v", converted, 100)
    }
    // Test KilowattPerFoot conversion
    kilowatts_per_footDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityKilowattPerFoot,
    }
    
    var kilowatts_per_footResult *units.LinearPowerDensity
    kilowatts_per_footResult, err = factory.FromDto(kilowatts_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_footResult.Convert(units.LinearPowerDensityKilowattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerFoot = %v, want %v", converted, 100)
    }
    // Test MegawattPerFoot conversion
    megawatts_per_footDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityMegawattPerFoot,
    }
    
    var megawatts_per_footResult *units.LinearPowerDensity
    megawatts_per_footResult, err = factory.FromDto(megawatts_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_footResult.Convert(units.LinearPowerDensityMegawattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerFoot = %v, want %v", converted, 100)
    }
    // Test GigawattPerFoot conversion
    gigawatts_per_footDto := units.LinearPowerDensityDto{
        Value: 100,
        Unit:  units.LinearPowerDensityGigawattPerFoot,
    }
    
    var gigawatts_per_footResult *units.LinearPowerDensity
    gigawatts_per_footResult, err = factory.FromDto(gigawatts_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_footResult.Convert(units.LinearPowerDensityGigawattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerFoot = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.LinearPowerDensityDto{
        Value: 0,
        Unit:  units.LinearPowerDensityWattPerMeter,
    }
    
    var zeroResult *units.LinearPowerDensity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestLinearPowerDensityFactory_FromDtoJSON(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "WattPerMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "WattPerMeter"}`)
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
    nanJSON, _ := json.Marshal(units.LinearPowerDensityDto{
        Value: nanValue,
        Unit:  units.LinearPowerDensityWattPerMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with WattPerMeter unit
    watts_per_meterJSON := []byte(`{"value": 100, "unit": "WattPerMeter"}`)
    watts_per_meterResult, err := factory.FromDtoJSON(watts_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_meterResult.Convert(units.LinearPowerDensityWattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with WattPerCentimeter unit
    watts_per_centimeterJSON := []byte(`{"value": 100, "unit": "WattPerCentimeter"}`)
    watts_per_centimeterResult, err := factory.FromDtoJSON(watts_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_centimeterResult.Convert(units.LinearPowerDensityWattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with WattPerMillimeter unit
    watts_per_millimeterJSON := []byte(`{"value": 100, "unit": "WattPerMillimeter"}`)
    watts_per_millimeterResult, err := factory.FromDtoJSON(watts_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_millimeterResult.Convert(units.LinearPowerDensityWattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with WattPerInch unit
    watts_per_inchJSON := []byte(`{"value": 100, "unit": "WattPerInch"}`)
    watts_per_inchResult, err := factory.FromDtoJSON(watts_per_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_inchResult.Convert(units.LinearPowerDensityWattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerInch = %v, want %v", converted, 100)
    }
    // Test JSON with WattPerFoot unit
    watts_per_footJSON := []byte(`{"value": 100, "unit": "WattPerFoot"}`)
    watts_per_footResult, err := factory.FromDtoJSON(watts_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_footResult.Convert(units.LinearPowerDensityWattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with MilliwattPerMeter unit
    milliwatts_per_meterJSON := []byte(`{"value": 100, "unit": "MilliwattPerMeter"}`)
    milliwatts_per_meterResult, err := factory.FromDtoJSON(milliwatts_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliwattPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_meterResult.Convert(units.LinearPowerDensityMilliwattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattPerMeter unit
    kilowatts_per_meterJSON := []byte(`{"value": 100, "unit": "KilowattPerMeter"}`)
    kilowatts_per_meterResult, err := factory.FromDtoJSON(kilowatts_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_meterResult.Convert(units.LinearPowerDensityKilowattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattPerMeter unit
    megawatts_per_meterJSON := []byte(`{"value": 100, "unit": "MegawattPerMeter"}`)
    megawatts_per_meterResult, err := factory.FromDtoJSON(megawatts_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_meterResult.Convert(units.LinearPowerDensityMegawattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattPerMeter unit
    gigawatts_per_meterJSON := []byte(`{"value": 100, "unit": "GigawattPerMeter"}`)
    gigawatts_per_meterResult, err := factory.FromDtoJSON(gigawatts_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_meterResult.Convert(units.LinearPowerDensityGigawattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilliwattPerCentimeter unit
    milliwatts_per_centimeterJSON := []byte(`{"value": 100, "unit": "MilliwattPerCentimeter"}`)
    milliwatts_per_centimeterResult, err := factory.FromDtoJSON(milliwatts_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliwattPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_centimeterResult.Convert(units.LinearPowerDensityMilliwattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattPerCentimeter unit
    kilowatts_per_centimeterJSON := []byte(`{"value": 100, "unit": "KilowattPerCentimeter"}`)
    kilowatts_per_centimeterResult, err := factory.FromDtoJSON(kilowatts_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_centimeterResult.Convert(units.LinearPowerDensityKilowattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattPerCentimeter unit
    megawatts_per_centimeterJSON := []byte(`{"value": 100, "unit": "MegawattPerCentimeter"}`)
    megawatts_per_centimeterResult, err := factory.FromDtoJSON(megawatts_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_centimeterResult.Convert(units.LinearPowerDensityMegawattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattPerCentimeter unit
    gigawatts_per_centimeterJSON := []byte(`{"value": 100, "unit": "GigawattPerCentimeter"}`)
    gigawatts_per_centimeterResult, err := factory.FromDtoJSON(gigawatts_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_centimeterResult.Convert(units.LinearPowerDensityGigawattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilliwattPerMillimeter unit
    milliwatts_per_millimeterJSON := []byte(`{"value": 100, "unit": "MilliwattPerMillimeter"}`)
    milliwatts_per_millimeterResult, err := factory.FromDtoJSON(milliwatts_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliwattPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_millimeterResult.Convert(units.LinearPowerDensityMilliwattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattPerMillimeter unit
    kilowatts_per_millimeterJSON := []byte(`{"value": 100, "unit": "KilowattPerMillimeter"}`)
    kilowatts_per_millimeterResult, err := factory.FromDtoJSON(kilowatts_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_millimeterResult.Convert(units.LinearPowerDensityKilowattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattPerMillimeter unit
    megawatts_per_millimeterJSON := []byte(`{"value": 100, "unit": "MegawattPerMillimeter"}`)
    megawatts_per_millimeterResult, err := factory.FromDtoJSON(megawatts_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_millimeterResult.Convert(units.LinearPowerDensityMegawattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattPerMillimeter unit
    gigawatts_per_millimeterJSON := []byte(`{"value": 100, "unit": "GigawattPerMillimeter"}`)
    gigawatts_per_millimeterResult, err := factory.FromDtoJSON(gigawatts_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_millimeterResult.Convert(units.LinearPowerDensityGigawattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilliwattPerInch unit
    milliwatts_per_inchJSON := []byte(`{"value": 100, "unit": "MilliwattPerInch"}`)
    milliwatts_per_inchResult, err := factory.FromDtoJSON(milliwatts_per_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliwattPerInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_inchResult.Convert(units.LinearPowerDensityMilliwattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerInch = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattPerInch unit
    kilowatts_per_inchJSON := []byte(`{"value": 100, "unit": "KilowattPerInch"}`)
    kilowatts_per_inchResult, err := factory.FromDtoJSON(kilowatts_per_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattPerInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_inchResult.Convert(units.LinearPowerDensityKilowattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerInch = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattPerInch unit
    megawatts_per_inchJSON := []byte(`{"value": 100, "unit": "MegawattPerInch"}`)
    megawatts_per_inchResult, err := factory.FromDtoJSON(megawatts_per_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattPerInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_inchResult.Convert(units.LinearPowerDensityMegawattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerInch = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattPerInch unit
    gigawatts_per_inchJSON := []byte(`{"value": 100, "unit": "GigawattPerInch"}`)
    gigawatts_per_inchResult, err := factory.FromDtoJSON(gigawatts_per_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattPerInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_inchResult.Convert(units.LinearPowerDensityGigawattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerInch = %v, want %v", converted, 100)
    }
    // Test JSON with MilliwattPerFoot unit
    milliwatts_per_footJSON := []byte(`{"value": 100, "unit": "MilliwattPerFoot"}`)
    milliwatts_per_footResult, err := factory.FromDtoJSON(milliwatts_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliwattPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_footResult.Convert(units.LinearPowerDensityMilliwattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattPerFoot unit
    kilowatts_per_footJSON := []byte(`{"value": 100, "unit": "KilowattPerFoot"}`)
    kilowatts_per_footResult, err := factory.FromDtoJSON(kilowatts_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_footResult.Convert(units.LinearPowerDensityKilowattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattPerFoot unit
    megawatts_per_footJSON := []byte(`{"value": 100, "unit": "MegawattPerFoot"}`)
    megawatts_per_footResult, err := factory.FromDtoJSON(megawatts_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_footResult.Convert(units.LinearPowerDensityMegawattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattPerFoot unit
    gigawatts_per_footJSON := []byte(`{"value": 100, "unit": "GigawattPerFoot"}`)
    gigawatts_per_footResult, err := factory.FromDtoJSON(gigawatts_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_footResult.Convert(units.LinearPowerDensityGigawattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerFoot = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "WattPerMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromWattsPerMeter function
func TestLinearPowerDensityFactory_FromWattsPerMeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerMeter(100)
    if err != nil {
        t.Errorf("FromWattsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityWattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromWattsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromWattsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerMeter(0)
    if err != nil {
        t.Errorf("FromWattsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityWattPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromWattsPerCentimeter function
func TestLinearPowerDensityFactory_FromWattsPerCentimeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromWattsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityWattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromWattsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromWattsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromWattsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityWattPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromWattsPerMillimeter function
func TestLinearPowerDensityFactory_FromWattsPerMillimeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromWattsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityWattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromWattsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromWattsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromWattsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityWattPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromWattsPerInch function
func TestLinearPowerDensityFactory_FromWattsPerInch(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerInch(100)
    if err != nil {
        t.Errorf("FromWattsPerInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityWattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerInch(math.NaN())
    if err == nil {
        t.Error("FromWattsPerInch() with NaN value should return error")
    }

    _, err = factory.FromWattsPerInch(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerInch() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerInch(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerInch(0)
    if err != nil {
        t.Errorf("FromWattsPerInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityWattPerInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerInch() with zero value = %v, want 0", converted)
    }
}
// Test FromWattsPerFoot function
func TestLinearPowerDensityFactory_FromWattsPerFoot(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerFoot(100)
    if err != nil {
        t.Errorf("FromWattsPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityWattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerFoot(math.NaN())
    if err == nil {
        t.Error("FromWattsPerFoot() with NaN value should return error")
    }

    _, err = factory.FromWattsPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerFoot(0)
    if err != nil {
        t.Errorf("FromWattsPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityWattPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwattsPerMeter function
func TestLinearPowerDensityFactory_FromMilliwattsPerMeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwattsPerMeter(100)
    if err != nil {
        t.Errorf("FromMilliwattsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityMilliwattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwattsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwattsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromMilliwattsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromMilliwattsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwattsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromMilliwattsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwattsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwattsPerMeter(0)
    if err != nil {
        t.Errorf("FromMilliwattsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityMilliwattPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwattsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattsPerMeter function
func TestLinearPowerDensityFactory_FromKilowattsPerMeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattsPerMeter(100)
    if err != nil {
        t.Errorf("FromKilowattsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityKilowattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromKilowattsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromKilowattsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilowattsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattsPerMeter(0)
    if err != nil {
        t.Errorf("FromKilowattsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityKilowattPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattsPerMeter function
func TestLinearPowerDensityFactory_FromMegawattsPerMeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattsPerMeter(100)
    if err != nil {
        t.Errorf("FromMegawattsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityMegawattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromMegawattsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromMegawattsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromMegawattsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattsPerMeter(0)
    if err != nil {
        t.Errorf("FromMegawattsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityMegawattPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattsPerMeter function
func TestLinearPowerDensityFactory_FromGigawattsPerMeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattsPerMeter(100)
    if err != nil {
        t.Errorf("FromGigawattsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityGigawattPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromGigawattsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromGigawattsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromGigawattsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattsPerMeter(0)
    if err != nil {
        t.Errorf("FromGigawattsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityGigawattPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwattsPerCentimeter function
func TestLinearPowerDensityFactory_FromMilliwattsPerCentimeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwattsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromMilliwattsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityMilliwattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwattsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwattsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMilliwattsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMilliwattsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwattsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMilliwattsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwattsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwattsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromMilliwattsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityMilliwattPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwattsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattsPerCentimeter function
func TestLinearPowerDensityFactory_FromKilowattsPerCentimeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromKilowattsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityKilowattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKilowattsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKilowattsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilowattsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromKilowattsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityKilowattPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattsPerCentimeter function
func TestLinearPowerDensityFactory_FromMegawattsPerCentimeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromMegawattsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityMegawattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMegawattsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMegawattsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMegawattsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromMegawattsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityMegawattPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattsPerCentimeter function
func TestLinearPowerDensityFactory_FromGigawattsPerCentimeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattsPerCentimeter(100)
    if err != nil {
        t.Errorf("FromGigawattsPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityGigawattPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattsPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattsPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromGigawattsPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromGigawattsPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattsPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromGigawattsPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattsPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattsPerCentimeter(0)
    if err != nil {
        t.Errorf("FromGigawattsPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityGigawattPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattsPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwattsPerMillimeter function
func TestLinearPowerDensityFactory_FromMilliwattsPerMillimeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwattsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromMilliwattsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityMilliwattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwattsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwattsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromMilliwattsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromMilliwattsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwattsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromMilliwattsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwattsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwattsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromMilliwattsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityMilliwattPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwattsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattsPerMillimeter function
func TestLinearPowerDensityFactory_FromKilowattsPerMillimeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromKilowattsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityKilowattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromKilowattsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromKilowattsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilowattsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromKilowattsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityKilowattPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattsPerMillimeter function
func TestLinearPowerDensityFactory_FromMegawattsPerMillimeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromMegawattsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityMegawattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromMegawattsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromMegawattsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromMegawattsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromMegawattsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityMegawattPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattsPerMillimeter function
func TestLinearPowerDensityFactory_FromGigawattsPerMillimeter(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattsPerMillimeter(100)
    if err != nil {
        t.Errorf("FromGigawattsPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityGigawattPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattsPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattsPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromGigawattsPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromGigawattsPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattsPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromGigawattsPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattsPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattsPerMillimeter(0)
    if err != nil {
        t.Errorf("FromGigawattsPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityGigawattPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattsPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwattsPerInch function
func TestLinearPowerDensityFactory_FromMilliwattsPerInch(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwattsPerInch(100)
    if err != nil {
        t.Errorf("FromMilliwattsPerInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityMilliwattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwattsPerInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwattsPerInch(math.NaN())
    if err == nil {
        t.Error("FromMilliwattsPerInch() with NaN value should return error")
    }

    _, err = factory.FromMilliwattsPerInch(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwattsPerInch() with +Inf value should return error")
    }

    _, err = factory.FromMilliwattsPerInch(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwattsPerInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwattsPerInch(0)
    if err != nil {
        t.Errorf("FromMilliwattsPerInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityMilliwattPerInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwattsPerInch() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattsPerInch function
func TestLinearPowerDensityFactory_FromKilowattsPerInch(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattsPerInch(100)
    if err != nil {
        t.Errorf("FromKilowattsPerInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityKilowattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattsPerInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattsPerInch(math.NaN())
    if err == nil {
        t.Error("FromKilowattsPerInch() with NaN value should return error")
    }

    _, err = factory.FromKilowattsPerInch(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattsPerInch() with +Inf value should return error")
    }

    _, err = factory.FromKilowattsPerInch(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattsPerInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattsPerInch(0)
    if err != nil {
        t.Errorf("FromKilowattsPerInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityKilowattPerInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattsPerInch() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattsPerInch function
func TestLinearPowerDensityFactory_FromMegawattsPerInch(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattsPerInch(100)
    if err != nil {
        t.Errorf("FromMegawattsPerInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityMegawattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattsPerInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattsPerInch(math.NaN())
    if err == nil {
        t.Error("FromMegawattsPerInch() with NaN value should return error")
    }

    _, err = factory.FromMegawattsPerInch(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattsPerInch() with +Inf value should return error")
    }

    _, err = factory.FromMegawattsPerInch(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattsPerInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattsPerInch(0)
    if err != nil {
        t.Errorf("FromMegawattsPerInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityMegawattPerInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattsPerInch() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattsPerInch function
func TestLinearPowerDensityFactory_FromGigawattsPerInch(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattsPerInch(100)
    if err != nil {
        t.Errorf("FromGigawattsPerInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityGigawattPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattsPerInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattsPerInch(math.NaN())
    if err == nil {
        t.Error("FromGigawattsPerInch() with NaN value should return error")
    }

    _, err = factory.FromGigawattsPerInch(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattsPerInch() with +Inf value should return error")
    }

    _, err = factory.FromGigawattsPerInch(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattsPerInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattsPerInch(0)
    if err != nil {
        t.Errorf("FromGigawattsPerInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityGigawattPerInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattsPerInch() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwattsPerFoot function
func TestLinearPowerDensityFactory_FromMilliwattsPerFoot(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwattsPerFoot(100)
    if err != nil {
        t.Errorf("FromMilliwattsPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityMilliwattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwattsPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwattsPerFoot(math.NaN())
    if err == nil {
        t.Error("FromMilliwattsPerFoot() with NaN value should return error")
    }

    _, err = factory.FromMilliwattsPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwattsPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromMilliwattsPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwattsPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwattsPerFoot(0)
    if err != nil {
        t.Errorf("FromMilliwattsPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityMilliwattPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwattsPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattsPerFoot function
func TestLinearPowerDensityFactory_FromKilowattsPerFoot(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattsPerFoot(100)
    if err != nil {
        t.Errorf("FromKilowattsPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityKilowattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattsPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattsPerFoot(math.NaN())
    if err == nil {
        t.Error("FromKilowattsPerFoot() with NaN value should return error")
    }

    _, err = factory.FromKilowattsPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattsPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromKilowattsPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattsPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattsPerFoot(0)
    if err != nil {
        t.Errorf("FromKilowattsPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityKilowattPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattsPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattsPerFoot function
func TestLinearPowerDensityFactory_FromMegawattsPerFoot(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattsPerFoot(100)
    if err != nil {
        t.Errorf("FromMegawattsPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityMegawattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattsPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattsPerFoot(math.NaN())
    if err == nil {
        t.Error("FromMegawattsPerFoot() with NaN value should return error")
    }

    _, err = factory.FromMegawattsPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattsPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromMegawattsPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattsPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattsPerFoot(0)
    if err != nil {
        t.Errorf("FromMegawattsPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityMegawattPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattsPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattsPerFoot function
func TestLinearPowerDensityFactory_FromGigawattsPerFoot(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattsPerFoot(100)
    if err != nil {
        t.Errorf("FromGigawattsPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LinearPowerDensityGigawattPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattsPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattsPerFoot(math.NaN())
    if err == nil {
        t.Error("FromGigawattsPerFoot() with NaN value should return error")
    }

    _, err = factory.FromGigawattsPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattsPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromGigawattsPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattsPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattsPerFoot(0)
    if err != nil {
        t.Errorf("FromGigawattsPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LinearPowerDensityGigawattPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattsPerFoot() with zero value = %v, want 0", converted)
    }
}

func TestLinearPowerDensityToString(t *testing.T) {
	factory := units.LinearPowerDensityFactory{}
	a, err := factory.CreateLinearPowerDensity(45, units.LinearPowerDensityWattPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LinearPowerDensityWattPerMeter, 2)
	expected := "45.00 " + units.GetLinearPowerDensityAbbreviation(units.LinearPowerDensityWattPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LinearPowerDensityWattPerMeter, -1)
	expected = "45 " + units.GetLinearPowerDensityAbbreviation(units.LinearPowerDensityWattPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLinearPowerDensity_EqualityAndComparison(t *testing.T) {
	factory := units.LinearPowerDensityFactory{}
	a1, _ := factory.CreateLinearPowerDensity(60, units.LinearPowerDensityWattPerMeter)
	a2, _ := factory.CreateLinearPowerDensity(60, units.LinearPowerDensityWattPerMeter)
	a3, _ := factory.CreateLinearPowerDensity(120, units.LinearPowerDensityWattPerMeter)

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

func TestLinearPowerDensity_Arithmetic(t *testing.T) {
	factory := units.LinearPowerDensityFactory{}
	a1, _ := factory.CreateLinearPowerDensity(30, units.LinearPowerDensityWattPerMeter)
	a2, _ := factory.CreateLinearPowerDensity(45, units.LinearPowerDensityWattPerMeter)

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


func TestGetLinearPowerDensityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.LinearPowerDensityUnits
        want string
    }{
        {
            name: "WattPerMeter abbreviation",
            unit: units.LinearPowerDensityWattPerMeter,
            want: "W/m",
        },
        {
            name: "WattPerCentimeter abbreviation",
            unit: units.LinearPowerDensityWattPerCentimeter,
            want: "W/cm",
        },
        {
            name: "WattPerMillimeter abbreviation",
            unit: units.LinearPowerDensityWattPerMillimeter,
            want: "W/mm",
        },
        {
            name: "WattPerInch abbreviation",
            unit: units.LinearPowerDensityWattPerInch,
            want: "W/in",
        },
        {
            name: "WattPerFoot abbreviation",
            unit: units.LinearPowerDensityWattPerFoot,
            want: "W/ft",
        },
        {
            name: "MilliwattPerMeter abbreviation",
            unit: units.LinearPowerDensityMilliwattPerMeter,
            want: "mW/m",
        },
        {
            name: "KilowattPerMeter abbreviation",
            unit: units.LinearPowerDensityKilowattPerMeter,
            want: "kW/m",
        },
        {
            name: "MegawattPerMeter abbreviation",
            unit: units.LinearPowerDensityMegawattPerMeter,
            want: "MW/m",
        },
        {
            name: "GigawattPerMeter abbreviation",
            unit: units.LinearPowerDensityGigawattPerMeter,
            want: "GW/m",
        },
        {
            name: "MilliwattPerCentimeter abbreviation",
            unit: units.LinearPowerDensityMilliwattPerCentimeter,
            want: "mW/cm",
        },
        {
            name: "KilowattPerCentimeter abbreviation",
            unit: units.LinearPowerDensityKilowattPerCentimeter,
            want: "kW/cm",
        },
        {
            name: "MegawattPerCentimeter abbreviation",
            unit: units.LinearPowerDensityMegawattPerCentimeter,
            want: "MW/cm",
        },
        {
            name: "GigawattPerCentimeter abbreviation",
            unit: units.LinearPowerDensityGigawattPerCentimeter,
            want: "GW/cm",
        },
        {
            name: "MilliwattPerMillimeter abbreviation",
            unit: units.LinearPowerDensityMilliwattPerMillimeter,
            want: "mW/mm",
        },
        {
            name: "KilowattPerMillimeter abbreviation",
            unit: units.LinearPowerDensityKilowattPerMillimeter,
            want: "kW/mm",
        },
        {
            name: "MegawattPerMillimeter abbreviation",
            unit: units.LinearPowerDensityMegawattPerMillimeter,
            want: "MW/mm",
        },
        {
            name: "GigawattPerMillimeter abbreviation",
            unit: units.LinearPowerDensityGigawattPerMillimeter,
            want: "GW/mm",
        },
        {
            name: "MilliwattPerInch abbreviation",
            unit: units.LinearPowerDensityMilliwattPerInch,
            want: "mW/in",
        },
        {
            name: "KilowattPerInch abbreviation",
            unit: units.LinearPowerDensityKilowattPerInch,
            want: "kW/in",
        },
        {
            name: "MegawattPerInch abbreviation",
            unit: units.LinearPowerDensityMegawattPerInch,
            want: "MW/in",
        },
        {
            name: "GigawattPerInch abbreviation",
            unit: units.LinearPowerDensityGigawattPerInch,
            want: "GW/in",
        },
        {
            name: "MilliwattPerFoot abbreviation",
            unit: units.LinearPowerDensityMilliwattPerFoot,
            want: "mW/ft",
        },
        {
            name: "KilowattPerFoot abbreviation",
            unit: units.LinearPowerDensityKilowattPerFoot,
            want: "kW/ft",
        },
        {
            name: "MegawattPerFoot abbreviation",
            unit: units.LinearPowerDensityMegawattPerFoot,
            want: "MW/ft",
        },
        {
            name: "GigawattPerFoot abbreviation",
            unit: units.LinearPowerDensityGigawattPerFoot,
            want: "GW/ft",
        },
        {
            name: "invalid unit",
            unit: units.LinearPowerDensityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetLinearPowerDensityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetLinearPowerDensityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestLinearPowerDensity_String(t *testing.T) {
    factory := units.LinearPowerDensityFactory{}
    
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
            unit, err := factory.CreateLinearPowerDensity(tt.value, units.LinearPowerDensityWattPerMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("LinearPowerDensity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestLinearPowerDensity_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.LinearPowerDensityFactory{}

	_, err := uf.CreateLinearPowerDensity(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}