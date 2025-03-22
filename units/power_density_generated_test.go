// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPowerDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "WattPerCubicMeter"}`
	
	factory := units.PowerDensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PowerDensityWattPerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.PowerDensityWattPerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "WattPerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPowerDensityDto_ToJSON(t *testing.T) {
	dto := units.PowerDensityDto{
		Value: 45,
		Unit:  units.PowerDensityWattPerCubicMeter,
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
	if result["unit"].(string) != string(units.PowerDensityWattPerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.PowerDensityWattPerCubicMeter, result["unit"])
	}
}

func TestNewPowerDensity_InvalidValue(t *testing.T) {
	factory := units.PowerDensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePowerDensity(math.NaN(), units.PowerDensityWattPerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePowerDensity(math.Inf(1), units.PowerDensityWattPerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPowerDensityConversions(t *testing.T) {
	factory := units.PowerDensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePowerDensity(180, units.PowerDensityWattPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to WattsPerCubicMeter.
		// No expected conversion value provided for WattsPerCubicMeter, verifying result is not NaN.
		result := a.WattsPerCubicMeter()
		cacheResult := a.WattsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to WattsPerCubicInch.
		// No expected conversion value provided for WattsPerCubicInch, verifying result is not NaN.
		result := a.WattsPerCubicInch()
		cacheResult := a.WattsPerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to WattsPerCubicFoot.
		// No expected conversion value provided for WattsPerCubicFoot, verifying result is not NaN.
		result := a.WattsPerCubicFoot()
		cacheResult := a.WattsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to WattsPerLiter.
		// No expected conversion value provided for WattsPerLiter, verifying result is not NaN.
		result := a.WattsPerLiter()
		cacheResult := a.WattsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to PicowattsPerCubicMeter.
		// No expected conversion value provided for PicowattsPerCubicMeter, verifying result is not NaN.
		result := a.PicowattsPerCubicMeter()
		cacheResult := a.PicowattsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicowattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerCubicMeter.
		// No expected conversion value provided for NanowattsPerCubicMeter, verifying result is not NaN.
		result := a.NanowattsPerCubicMeter()
		cacheResult := a.NanowattsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanowattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerCubicMeter.
		// No expected conversion value provided for MicrowattsPerCubicMeter, verifying result is not NaN.
		result := a.MicrowattsPerCubicMeter()
		cacheResult := a.MicrowattsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrowattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerCubicMeter.
		// No expected conversion value provided for MilliwattsPerCubicMeter, verifying result is not NaN.
		result := a.MilliwattsPerCubicMeter()
		cacheResult := a.MilliwattsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliwattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to DeciwattsPerCubicMeter.
		// No expected conversion value provided for DeciwattsPerCubicMeter, verifying result is not NaN.
		result := a.DeciwattsPerCubicMeter()
		cacheResult := a.DeciwattsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DeciwattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to DecawattsPerCubicMeter.
		// No expected conversion value provided for DecawattsPerCubicMeter, verifying result is not NaN.
		result := a.DecawattsPerCubicMeter()
		cacheResult := a.DecawattsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecawattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerCubicMeter.
		// No expected conversion value provided for KilowattsPerCubicMeter, verifying result is not NaN.
		result := a.KilowattsPerCubicMeter()
		cacheResult := a.KilowattsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerCubicMeter.
		// No expected conversion value provided for MegawattsPerCubicMeter, verifying result is not NaN.
		result := a.MegawattsPerCubicMeter()
		cacheResult := a.MegawattsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerCubicMeter.
		// No expected conversion value provided for GigawattsPerCubicMeter, verifying result is not NaN.
		result := a.GigawattsPerCubicMeter()
		cacheResult := a.GigawattsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to TerawattsPerCubicMeter.
		// No expected conversion value provided for TerawattsPerCubicMeter, verifying result is not NaN.
		result := a.TerawattsPerCubicMeter()
		cacheResult := a.TerawattsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TerawattsPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to PicowattsPerCubicInch.
		// No expected conversion value provided for PicowattsPerCubicInch, verifying result is not NaN.
		result := a.PicowattsPerCubicInch()
		cacheResult := a.PicowattsPerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicowattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerCubicInch.
		// No expected conversion value provided for NanowattsPerCubicInch, verifying result is not NaN.
		result := a.NanowattsPerCubicInch()
		cacheResult := a.NanowattsPerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanowattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerCubicInch.
		// No expected conversion value provided for MicrowattsPerCubicInch, verifying result is not NaN.
		result := a.MicrowattsPerCubicInch()
		cacheResult := a.MicrowattsPerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrowattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerCubicInch.
		// No expected conversion value provided for MilliwattsPerCubicInch, verifying result is not NaN.
		result := a.MilliwattsPerCubicInch()
		cacheResult := a.MilliwattsPerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliwattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to DeciwattsPerCubicInch.
		// No expected conversion value provided for DeciwattsPerCubicInch, verifying result is not NaN.
		result := a.DeciwattsPerCubicInch()
		cacheResult := a.DeciwattsPerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DeciwattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to DecawattsPerCubicInch.
		// No expected conversion value provided for DecawattsPerCubicInch, verifying result is not NaN.
		result := a.DecawattsPerCubicInch()
		cacheResult := a.DecawattsPerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecawattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerCubicInch.
		// No expected conversion value provided for KilowattsPerCubicInch, verifying result is not NaN.
		result := a.KilowattsPerCubicInch()
		cacheResult := a.KilowattsPerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerCubicInch.
		// No expected conversion value provided for MegawattsPerCubicInch, verifying result is not NaN.
		result := a.MegawattsPerCubicInch()
		cacheResult := a.MegawattsPerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerCubicInch.
		// No expected conversion value provided for GigawattsPerCubicInch, verifying result is not NaN.
		result := a.GigawattsPerCubicInch()
		cacheResult := a.GigawattsPerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to TerawattsPerCubicInch.
		// No expected conversion value provided for TerawattsPerCubicInch, verifying result is not NaN.
		result := a.TerawattsPerCubicInch()
		cacheResult := a.TerawattsPerCubicInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TerawattsPerCubicInch returned NaN")
		}
	}
	{
		// Test conversion to PicowattsPerCubicFoot.
		// No expected conversion value provided for PicowattsPerCubicFoot, verifying result is not NaN.
		result := a.PicowattsPerCubicFoot()
		cacheResult := a.PicowattsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicowattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerCubicFoot.
		// No expected conversion value provided for NanowattsPerCubicFoot, verifying result is not NaN.
		result := a.NanowattsPerCubicFoot()
		cacheResult := a.NanowattsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanowattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerCubicFoot.
		// No expected conversion value provided for MicrowattsPerCubicFoot, verifying result is not NaN.
		result := a.MicrowattsPerCubicFoot()
		cacheResult := a.MicrowattsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrowattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerCubicFoot.
		// No expected conversion value provided for MilliwattsPerCubicFoot, verifying result is not NaN.
		result := a.MilliwattsPerCubicFoot()
		cacheResult := a.MilliwattsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliwattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to DeciwattsPerCubicFoot.
		// No expected conversion value provided for DeciwattsPerCubicFoot, verifying result is not NaN.
		result := a.DeciwattsPerCubicFoot()
		cacheResult := a.DeciwattsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DeciwattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to DecawattsPerCubicFoot.
		// No expected conversion value provided for DecawattsPerCubicFoot, verifying result is not NaN.
		result := a.DecawattsPerCubicFoot()
		cacheResult := a.DecawattsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecawattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerCubicFoot.
		// No expected conversion value provided for KilowattsPerCubicFoot, verifying result is not NaN.
		result := a.KilowattsPerCubicFoot()
		cacheResult := a.KilowattsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerCubicFoot.
		// No expected conversion value provided for MegawattsPerCubicFoot, verifying result is not NaN.
		result := a.MegawattsPerCubicFoot()
		cacheResult := a.MegawattsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerCubicFoot.
		// No expected conversion value provided for GigawattsPerCubicFoot, verifying result is not NaN.
		result := a.GigawattsPerCubicFoot()
		cacheResult := a.GigawattsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to TerawattsPerCubicFoot.
		// No expected conversion value provided for TerawattsPerCubicFoot, verifying result is not NaN.
		result := a.TerawattsPerCubicFoot()
		cacheResult := a.TerawattsPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TerawattsPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to PicowattsPerLiter.
		// No expected conversion value provided for PicowattsPerLiter, verifying result is not NaN.
		result := a.PicowattsPerLiter()
		cacheResult := a.PicowattsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicowattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerLiter.
		// No expected conversion value provided for NanowattsPerLiter, verifying result is not NaN.
		result := a.NanowattsPerLiter()
		cacheResult := a.NanowattsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanowattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerLiter.
		// No expected conversion value provided for MicrowattsPerLiter, verifying result is not NaN.
		result := a.MicrowattsPerLiter()
		cacheResult := a.MicrowattsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrowattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerLiter.
		// No expected conversion value provided for MilliwattsPerLiter, verifying result is not NaN.
		result := a.MilliwattsPerLiter()
		cacheResult := a.MilliwattsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliwattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to DeciwattsPerLiter.
		// No expected conversion value provided for DeciwattsPerLiter, verifying result is not NaN.
		result := a.DeciwattsPerLiter()
		cacheResult := a.DeciwattsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DeciwattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to DecawattsPerLiter.
		// No expected conversion value provided for DecawattsPerLiter, verifying result is not NaN.
		result := a.DecawattsPerLiter()
		cacheResult := a.DecawattsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecawattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerLiter.
		// No expected conversion value provided for KilowattsPerLiter, verifying result is not NaN.
		result := a.KilowattsPerLiter()
		cacheResult := a.KilowattsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerLiter.
		// No expected conversion value provided for MegawattsPerLiter, verifying result is not NaN.
		result := a.MegawattsPerLiter()
		cacheResult := a.MegawattsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to GigawattsPerLiter.
		// No expected conversion value provided for GigawattsPerLiter, verifying result is not NaN.
		result := a.GigawattsPerLiter()
		cacheResult := a.GigawattsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattsPerLiter returned NaN")
		}
	}
	{
		// Test conversion to TerawattsPerLiter.
		// No expected conversion value provided for TerawattsPerLiter, verifying result is not NaN.
		result := a.TerawattsPerLiter()
		cacheResult := a.TerawattsPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TerawattsPerLiter returned NaN")
		}
	}
}

func TestPowerDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PowerDensityFactory{}
	a, err := factory.CreatePowerDensity(90, units.PowerDensityWattPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PowerDensityWattPerCubicMeter {
		t.Errorf("expected default unit WattPerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PowerDensityWattPerCubicMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PowerDensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PowerDensityWattPerCubicMeter {
		t.Errorf("expected unit WattPerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPowerDensityFactory_FromDto(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityWattPerCubicMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.PowerDensityDto{
        Value: math.NaN(),
        Unit:  units.PowerDensityWattPerCubicMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test WattPerCubicMeter conversion
    watts_per_cubic_meterDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityWattPerCubicMeter,
    }
    
    var watts_per_cubic_meterResult *units.PowerDensity
    watts_per_cubic_meterResult, err = factory.FromDto(watts_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_cubic_meterResult.Convert(units.PowerDensityWattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test WattPerCubicInch conversion
    watts_per_cubic_inchDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityWattPerCubicInch,
    }
    
    var watts_per_cubic_inchResult *units.PowerDensity
    watts_per_cubic_inchResult, err = factory.FromDto(watts_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_cubic_inchResult.Convert(units.PowerDensityWattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test WattPerCubicFoot conversion
    watts_per_cubic_footDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityWattPerCubicFoot,
    }
    
    var watts_per_cubic_footResult *units.PowerDensity
    watts_per_cubic_footResult, err = factory.FromDto(watts_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_cubic_footResult.Convert(units.PowerDensityWattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test WattPerLiter conversion
    watts_per_literDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityWattPerLiter,
    }
    
    var watts_per_literResult *units.PowerDensity
    watts_per_literResult, err = factory.FromDto(watts_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_literResult.Convert(units.PowerDensityWattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerLiter = %v, want %v", converted, 100)
    }
    // Test PicowattPerCubicMeter conversion
    picowatts_per_cubic_meterDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityPicowattPerCubicMeter,
    }
    
    var picowatts_per_cubic_meterResult *units.PowerDensity
    picowatts_per_cubic_meterResult, err = factory.FromDto(picowatts_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with PicowattPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowatts_per_cubic_meterResult.Convert(units.PowerDensityPicowattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicowattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test NanowattPerCubicMeter conversion
    nanowatts_per_cubic_meterDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityNanowattPerCubicMeter,
    }
    
    var nanowatts_per_cubic_meterResult *units.PowerDensity
    nanowatts_per_cubic_meterResult, err = factory.FromDto(nanowatts_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with NanowattPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_cubic_meterResult.Convert(units.PowerDensityNanowattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test MicrowattPerCubicMeter conversion
    microwatts_per_cubic_meterDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityMicrowattPerCubicMeter,
    }
    
    var microwatts_per_cubic_meterResult *units.PowerDensity
    microwatts_per_cubic_meterResult, err = factory.FromDto(microwatts_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrowattPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_cubic_meterResult.Convert(units.PowerDensityMicrowattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test MilliwattPerCubicMeter conversion
    milliwatts_per_cubic_meterDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityMilliwattPerCubicMeter,
    }
    
    var milliwatts_per_cubic_meterResult *units.PowerDensity
    milliwatts_per_cubic_meterResult, err = factory.FromDto(milliwatts_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MilliwattPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_cubic_meterResult.Convert(units.PowerDensityMilliwattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test DeciwattPerCubicMeter conversion
    deciwatts_per_cubic_meterDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityDeciwattPerCubicMeter,
    }
    
    var deciwatts_per_cubic_meterResult *units.PowerDensity
    deciwatts_per_cubic_meterResult, err = factory.FromDto(deciwatts_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with DeciwattPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwatts_per_cubic_meterResult.Convert(units.PowerDensityDeciwattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciwattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test DecawattPerCubicMeter conversion
    decawatts_per_cubic_meterDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityDecawattPerCubicMeter,
    }
    
    var decawatts_per_cubic_meterResult *units.PowerDensity
    decawatts_per_cubic_meterResult, err = factory.FromDto(decawatts_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with DecawattPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decawatts_per_cubic_meterResult.Convert(units.PowerDensityDecawattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecawattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test KilowattPerCubicMeter conversion
    kilowatts_per_cubic_meterDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityKilowattPerCubicMeter,
    }
    
    var kilowatts_per_cubic_meterResult *units.PowerDensity
    kilowatts_per_cubic_meterResult, err = factory.FromDto(kilowatts_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_cubic_meterResult.Convert(units.PowerDensityKilowattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test MegawattPerCubicMeter conversion
    megawatts_per_cubic_meterDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityMegawattPerCubicMeter,
    }
    
    var megawatts_per_cubic_meterResult *units.PowerDensity
    megawatts_per_cubic_meterResult, err = factory.FromDto(megawatts_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_cubic_meterResult.Convert(units.PowerDensityMegawattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test GigawattPerCubicMeter conversion
    gigawatts_per_cubic_meterDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityGigawattPerCubicMeter,
    }
    
    var gigawatts_per_cubic_meterResult *units.PowerDensity
    gigawatts_per_cubic_meterResult, err = factory.FromDto(gigawatts_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_cubic_meterResult.Convert(units.PowerDensityGigawattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test TerawattPerCubicMeter conversion
    terawatts_per_cubic_meterDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityTerawattPerCubicMeter,
    }
    
    var terawatts_per_cubic_meterResult *units.PowerDensity
    terawatts_per_cubic_meterResult, err = factory.FromDto(terawatts_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with TerawattPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatts_per_cubic_meterResult.Convert(units.PowerDensityTerawattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test PicowattPerCubicInch conversion
    picowatts_per_cubic_inchDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityPicowattPerCubicInch,
    }
    
    var picowatts_per_cubic_inchResult *units.PowerDensity
    picowatts_per_cubic_inchResult, err = factory.FromDto(picowatts_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with PicowattPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowatts_per_cubic_inchResult.Convert(units.PowerDensityPicowattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicowattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test NanowattPerCubicInch conversion
    nanowatts_per_cubic_inchDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityNanowattPerCubicInch,
    }
    
    var nanowatts_per_cubic_inchResult *units.PowerDensity
    nanowatts_per_cubic_inchResult, err = factory.FromDto(nanowatts_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with NanowattPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_cubic_inchResult.Convert(units.PowerDensityNanowattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test MicrowattPerCubicInch conversion
    microwatts_per_cubic_inchDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityMicrowattPerCubicInch,
    }
    
    var microwatts_per_cubic_inchResult *units.PowerDensity
    microwatts_per_cubic_inchResult, err = factory.FromDto(microwatts_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with MicrowattPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_cubic_inchResult.Convert(units.PowerDensityMicrowattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test MilliwattPerCubicInch conversion
    milliwatts_per_cubic_inchDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityMilliwattPerCubicInch,
    }
    
    var milliwatts_per_cubic_inchResult *units.PowerDensity
    milliwatts_per_cubic_inchResult, err = factory.FromDto(milliwatts_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with MilliwattPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_cubic_inchResult.Convert(units.PowerDensityMilliwattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test DeciwattPerCubicInch conversion
    deciwatts_per_cubic_inchDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityDeciwattPerCubicInch,
    }
    
    var deciwatts_per_cubic_inchResult *units.PowerDensity
    deciwatts_per_cubic_inchResult, err = factory.FromDto(deciwatts_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with DeciwattPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwatts_per_cubic_inchResult.Convert(units.PowerDensityDeciwattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciwattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test DecawattPerCubicInch conversion
    decawatts_per_cubic_inchDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityDecawattPerCubicInch,
    }
    
    var decawatts_per_cubic_inchResult *units.PowerDensity
    decawatts_per_cubic_inchResult, err = factory.FromDto(decawatts_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with DecawattPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decawatts_per_cubic_inchResult.Convert(units.PowerDensityDecawattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecawattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test KilowattPerCubicInch conversion
    kilowatts_per_cubic_inchDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityKilowattPerCubicInch,
    }
    
    var kilowatts_per_cubic_inchResult *units.PowerDensity
    kilowatts_per_cubic_inchResult, err = factory.FromDto(kilowatts_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_cubic_inchResult.Convert(units.PowerDensityKilowattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test MegawattPerCubicInch conversion
    megawatts_per_cubic_inchDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityMegawattPerCubicInch,
    }
    
    var megawatts_per_cubic_inchResult *units.PowerDensity
    megawatts_per_cubic_inchResult, err = factory.FromDto(megawatts_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_cubic_inchResult.Convert(units.PowerDensityMegawattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test GigawattPerCubicInch conversion
    gigawatts_per_cubic_inchDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityGigawattPerCubicInch,
    }
    
    var gigawatts_per_cubic_inchResult *units.PowerDensity
    gigawatts_per_cubic_inchResult, err = factory.FromDto(gigawatts_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_cubic_inchResult.Convert(units.PowerDensityGigawattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test TerawattPerCubicInch conversion
    terawatts_per_cubic_inchDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityTerawattPerCubicInch,
    }
    
    var terawatts_per_cubic_inchResult *units.PowerDensity
    terawatts_per_cubic_inchResult, err = factory.FromDto(terawatts_per_cubic_inchDto)
    if err != nil {
        t.Errorf("FromDto() with TerawattPerCubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatts_per_cubic_inchResult.Convert(units.PowerDensityTerawattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test PicowattPerCubicFoot conversion
    picowatts_per_cubic_footDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityPicowattPerCubicFoot,
    }
    
    var picowatts_per_cubic_footResult *units.PowerDensity
    picowatts_per_cubic_footResult, err = factory.FromDto(picowatts_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with PicowattPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowatts_per_cubic_footResult.Convert(units.PowerDensityPicowattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicowattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test NanowattPerCubicFoot conversion
    nanowatts_per_cubic_footDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityNanowattPerCubicFoot,
    }
    
    var nanowatts_per_cubic_footResult *units.PowerDensity
    nanowatts_per_cubic_footResult, err = factory.FromDto(nanowatts_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with NanowattPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_cubic_footResult.Convert(units.PowerDensityNanowattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test MicrowattPerCubicFoot conversion
    microwatts_per_cubic_footDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityMicrowattPerCubicFoot,
    }
    
    var microwatts_per_cubic_footResult *units.PowerDensity
    microwatts_per_cubic_footResult, err = factory.FromDto(microwatts_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with MicrowattPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_cubic_footResult.Convert(units.PowerDensityMicrowattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test MilliwattPerCubicFoot conversion
    milliwatts_per_cubic_footDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityMilliwattPerCubicFoot,
    }
    
    var milliwatts_per_cubic_footResult *units.PowerDensity
    milliwatts_per_cubic_footResult, err = factory.FromDto(milliwatts_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with MilliwattPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_cubic_footResult.Convert(units.PowerDensityMilliwattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test DeciwattPerCubicFoot conversion
    deciwatts_per_cubic_footDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityDeciwattPerCubicFoot,
    }
    
    var deciwatts_per_cubic_footResult *units.PowerDensity
    deciwatts_per_cubic_footResult, err = factory.FromDto(deciwatts_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with DeciwattPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwatts_per_cubic_footResult.Convert(units.PowerDensityDeciwattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciwattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test DecawattPerCubicFoot conversion
    decawatts_per_cubic_footDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityDecawattPerCubicFoot,
    }
    
    var decawatts_per_cubic_footResult *units.PowerDensity
    decawatts_per_cubic_footResult, err = factory.FromDto(decawatts_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with DecawattPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decawatts_per_cubic_footResult.Convert(units.PowerDensityDecawattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecawattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test KilowattPerCubicFoot conversion
    kilowatts_per_cubic_footDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityKilowattPerCubicFoot,
    }
    
    var kilowatts_per_cubic_footResult *units.PowerDensity
    kilowatts_per_cubic_footResult, err = factory.FromDto(kilowatts_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_cubic_footResult.Convert(units.PowerDensityKilowattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test MegawattPerCubicFoot conversion
    megawatts_per_cubic_footDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityMegawattPerCubicFoot,
    }
    
    var megawatts_per_cubic_footResult *units.PowerDensity
    megawatts_per_cubic_footResult, err = factory.FromDto(megawatts_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_cubic_footResult.Convert(units.PowerDensityMegawattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test GigawattPerCubicFoot conversion
    gigawatts_per_cubic_footDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityGigawattPerCubicFoot,
    }
    
    var gigawatts_per_cubic_footResult *units.PowerDensity
    gigawatts_per_cubic_footResult, err = factory.FromDto(gigawatts_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_cubic_footResult.Convert(units.PowerDensityGigawattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test TerawattPerCubicFoot conversion
    terawatts_per_cubic_footDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityTerawattPerCubicFoot,
    }
    
    var terawatts_per_cubic_footResult *units.PowerDensity
    terawatts_per_cubic_footResult, err = factory.FromDto(terawatts_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with TerawattPerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatts_per_cubic_footResult.Convert(units.PowerDensityTerawattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test PicowattPerLiter conversion
    picowatts_per_literDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityPicowattPerLiter,
    }
    
    var picowatts_per_literResult *units.PowerDensity
    picowatts_per_literResult, err = factory.FromDto(picowatts_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with PicowattPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowatts_per_literResult.Convert(units.PowerDensityPicowattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicowattPerLiter = %v, want %v", converted, 100)
    }
    // Test NanowattPerLiter conversion
    nanowatts_per_literDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityNanowattPerLiter,
    }
    
    var nanowatts_per_literResult *units.PowerDensity
    nanowatts_per_literResult, err = factory.FromDto(nanowatts_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with NanowattPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_literResult.Convert(units.PowerDensityNanowattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerLiter = %v, want %v", converted, 100)
    }
    // Test MicrowattPerLiter conversion
    microwatts_per_literDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityMicrowattPerLiter,
    }
    
    var microwatts_per_literResult *units.PowerDensity
    microwatts_per_literResult, err = factory.FromDto(microwatts_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MicrowattPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_literResult.Convert(units.PowerDensityMicrowattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerLiter = %v, want %v", converted, 100)
    }
    // Test MilliwattPerLiter conversion
    milliwatts_per_literDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityMilliwattPerLiter,
    }
    
    var milliwatts_per_literResult *units.PowerDensity
    milliwatts_per_literResult, err = factory.FromDto(milliwatts_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MilliwattPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_literResult.Convert(units.PowerDensityMilliwattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerLiter = %v, want %v", converted, 100)
    }
    // Test DeciwattPerLiter conversion
    deciwatts_per_literDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityDeciwattPerLiter,
    }
    
    var deciwatts_per_literResult *units.PowerDensity
    deciwatts_per_literResult, err = factory.FromDto(deciwatts_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with DeciwattPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwatts_per_literResult.Convert(units.PowerDensityDeciwattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciwattPerLiter = %v, want %v", converted, 100)
    }
    // Test DecawattPerLiter conversion
    decawatts_per_literDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityDecawattPerLiter,
    }
    
    var decawatts_per_literResult *units.PowerDensity
    decawatts_per_literResult, err = factory.FromDto(decawatts_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with DecawattPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decawatts_per_literResult.Convert(units.PowerDensityDecawattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecawattPerLiter = %v, want %v", converted, 100)
    }
    // Test KilowattPerLiter conversion
    kilowatts_per_literDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityKilowattPerLiter,
    }
    
    var kilowatts_per_literResult *units.PowerDensity
    kilowatts_per_literResult, err = factory.FromDto(kilowatts_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_literResult.Convert(units.PowerDensityKilowattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerLiter = %v, want %v", converted, 100)
    }
    // Test MegawattPerLiter conversion
    megawatts_per_literDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityMegawattPerLiter,
    }
    
    var megawatts_per_literResult *units.PowerDensity
    megawatts_per_literResult, err = factory.FromDto(megawatts_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_literResult.Convert(units.PowerDensityMegawattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerLiter = %v, want %v", converted, 100)
    }
    // Test GigawattPerLiter conversion
    gigawatts_per_literDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityGigawattPerLiter,
    }
    
    var gigawatts_per_literResult *units.PowerDensity
    gigawatts_per_literResult, err = factory.FromDto(gigawatts_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_literResult.Convert(units.PowerDensityGigawattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerLiter = %v, want %v", converted, 100)
    }
    // Test TerawattPerLiter conversion
    terawatts_per_literDto := units.PowerDensityDto{
        Value: 100,
        Unit:  units.PowerDensityTerawattPerLiter,
    }
    
    var terawatts_per_literResult *units.PowerDensity
    terawatts_per_literResult, err = factory.FromDto(terawatts_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with TerawattPerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatts_per_literResult.Convert(units.PowerDensityTerawattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattPerLiter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.PowerDensityDto{
        Value: 0,
        Unit:  units.PowerDensityWattPerCubicMeter,
    }
    
    var zeroResult *units.PowerDensity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestPowerDensityFactory_FromDtoJSON(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "WattPerCubicMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "WattPerCubicMeter"}`)
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
    nanJSON, _ := json.Marshal(units.PowerDensityDto{
        Value: nanValue,
        Unit:  units.PowerDensityWattPerCubicMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with WattPerCubicMeter unit
    watts_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "WattPerCubicMeter"}`)
    watts_per_cubic_meterResult, err := factory.FromDtoJSON(watts_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_cubic_meterResult.Convert(units.PowerDensityWattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with WattPerCubicInch unit
    watts_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "WattPerCubicInch"}`)
    watts_per_cubic_inchResult, err := factory.FromDtoJSON(watts_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_cubic_inchResult.Convert(units.PowerDensityWattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with WattPerCubicFoot unit
    watts_per_cubic_footJSON := []byte(`{"value": 100, "unit": "WattPerCubicFoot"}`)
    watts_per_cubic_footResult, err := factory.FromDtoJSON(watts_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_cubic_footResult.Convert(units.PowerDensityWattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with WattPerLiter unit
    watts_per_literJSON := []byte(`{"value": 100, "unit": "WattPerLiter"}`)
    watts_per_literResult, err := factory.FromDtoJSON(watts_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_literResult.Convert(units.PowerDensityWattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with PicowattPerCubicMeter unit
    picowatts_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "PicowattPerCubicMeter"}`)
    picowatts_per_cubic_meterResult, err := factory.FromDtoJSON(picowatts_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicowattPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowatts_per_cubic_meterResult.Convert(units.PowerDensityPicowattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicowattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with NanowattPerCubicMeter unit
    nanowatts_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "NanowattPerCubicMeter"}`)
    nanowatts_per_cubic_meterResult, err := factory.FromDtoJSON(nanowatts_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanowattPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_cubic_meterResult.Convert(units.PowerDensityNanowattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrowattPerCubicMeter unit
    microwatts_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "MicrowattPerCubicMeter"}`)
    microwatts_per_cubic_meterResult, err := factory.FromDtoJSON(microwatts_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrowattPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_cubic_meterResult.Convert(units.PowerDensityMicrowattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilliwattPerCubicMeter unit
    milliwatts_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "MilliwattPerCubicMeter"}`)
    milliwatts_per_cubic_meterResult, err := factory.FromDtoJSON(milliwatts_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliwattPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_cubic_meterResult.Convert(units.PowerDensityMilliwattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with DeciwattPerCubicMeter unit
    deciwatts_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "DeciwattPerCubicMeter"}`)
    deciwatts_per_cubic_meterResult, err := factory.FromDtoJSON(deciwatts_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DeciwattPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwatts_per_cubic_meterResult.Convert(units.PowerDensityDeciwattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciwattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with DecawattPerCubicMeter unit
    decawatts_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "DecawattPerCubicMeter"}`)
    decawatts_per_cubic_meterResult, err := factory.FromDtoJSON(decawatts_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecawattPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decawatts_per_cubic_meterResult.Convert(units.PowerDensityDecawattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecawattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattPerCubicMeter unit
    kilowatts_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "KilowattPerCubicMeter"}`)
    kilowatts_per_cubic_meterResult, err := factory.FromDtoJSON(kilowatts_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_cubic_meterResult.Convert(units.PowerDensityKilowattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattPerCubicMeter unit
    megawatts_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "MegawattPerCubicMeter"}`)
    megawatts_per_cubic_meterResult, err := factory.FromDtoJSON(megawatts_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_cubic_meterResult.Convert(units.PowerDensityMegawattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattPerCubicMeter unit
    gigawatts_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "GigawattPerCubicMeter"}`)
    gigawatts_per_cubic_meterResult, err := factory.FromDtoJSON(gigawatts_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_cubic_meterResult.Convert(units.PowerDensityGigawattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with TerawattPerCubicMeter unit
    terawatts_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "TerawattPerCubicMeter"}`)
    terawatts_per_cubic_meterResult, err := factory.FromDtoJSON(terawatts_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TerawattPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatts_per_cubic_meterResult.Convert(units.PowerDensityTerawattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattPerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PicowattPerCubicInch unit
    picowatts_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "PicowattPerCubicInch"}`)
    picowatts_per_cubic_inchResult, err := factory.FromDtoJSON(picowatts_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicowattPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowatts_per_cubic_inchResult.Convert(units.PowerDensityPicowattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicowattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with NanowattPerCubicInch unit
    nanowatts_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "NanowattPerCubicInch"}`)
    nanowatts_per_cubic_inchResult, err := factory.FromDtoJSON(nanowatts_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanowattPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_cubic_inchResult.Convert(units.PowerDensityNanowattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with MicrowattPerCubicInch unit
    microwatts_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "MicrowattPerCubicInch"}`)
    microwatts_per_cubic_inchResult, err := factory.FromDtoJSON(microwatts_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrowattPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_cubic_inchResult.Convert(units.PowerDensityMicrowattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with MilliwattPerCubicInch unit
    milliwatts_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "MilliwattPerCubicInch"}`)
    milliwatts_per_cubic_inchResult, err := factory.FromDtoJSON(milliwatts_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliwattPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_cubic_inchResult.Convert(units.PowerDensityMilliwattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with DeciwattPerCubicInch unit
    deciwatts_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "DeciwattPerCubicInch"}`)
    deciwatts_per_cubic_inchResult, err := factory.FromDtoJSON(deciwatts_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DeciwattPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwatts_per_cubic_inchResult.Convert(units.PowerDensityDeciwattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciwattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with DecawattPerCubicInch unit
    decawatts_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "DecawattPerCubicInch"}`)
    decawatts_per_cubic_inchResult, err := factory.FromDtoJSON(decawatts_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecawattPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decawatts_per_cubic_inchResult.Convert(units.PowerDensityDecawattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecawattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattPerCubicInch unit
    kilowatts_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "KilowattPerCubicInch"}`)
    kilowatts_per_cubic_inchResult, err := factory.FromDtoJSON(kilowatts_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_cubic_inchResult.Convert(units.PowerDensityKilowattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattPerCubicInch unit
    megawatts_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "MegawattPerCubicInch"}`)
    megawatts_per_cubic_inchResult, err := factory.FromDtoJSON(megawatts_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_cubic_inchResult.Convert(units.PowerDensityMegawattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattPerCubicInch unit
    gigawatts_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "GigawattPerCubicInch"}`)
    gigawatts_per_cubic_inchResult, err := factory.FromDtoJSON(gigawatts_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_cubic_inchResult.Convert(units.PowerDensityGigawattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with TerawattPerCubicInch unit
    terawatts_per_cubic_inchJSON := []byte(`{"value": 100, "unit": "TerawattPerCubicInch"}`)
    terawatts_per_cubic_inchResult, err := factory.FromDtoJSON(terawatts_per_cubic_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TerawattPerCubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatts_per_cubic_inchResult.Convert(units.PowerDensityTerawattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattPerCubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with PicowattPerCubicFoot unit
    picowatts_per_cubic_footJSON := []byte(`{"value": 100, "unit": "PicowattPerCubicFoot"}`)
    picowatts_per_cubic_footResult, err := factory.FromDtoJSON(picowatts_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicowattPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowatts_per_cubic_footResult.Convert(units.PowerDensityPicowattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicowattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with NanowattPerCubicFoot unit
    nanowatts_per_cubic_footJSON := []byte(`{"value": 100, "unit": "NanowattPerCubicFoot"}`)
    nanowatts_per_cubic_footResult, err := factory.FromDtoJSON(nanowatts_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanowattPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_cubic_footResult.Convert(units.PowerDensityNanowattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with MicrowattPerCubicFoot unit
    microwatts_per_cubic_footJSON := []byte(`{"value": 100, "unit": "MicrowattPerCubicFoot"}`)
    microwatts_per_cubic_footResult, err := factory.FromDtoJSON(microwatts_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrowattPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_cubic_footResult.Convert(units.PowerDensityMicrowattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with MilliwattPerCubicFoot unit
    milliwatts_per_cubic_footJSON := []byte(`{"value": 100, "unit": "MilliwattPerCubicFoot"}`)
    milliwatts_per_cubic_footResult, err := factory.FromDtoJSON(milliwatts_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliwattPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_cubic_footResult.Convert(units.PowerDensityMilliwattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with DeciwattPerCubicFoot unit
    deciwatts_per_cubic_footJSON := []byte(`{"value": 100, "unit": "DeciwattPerCubicFoot"}`)
    deciwatts_per_cubic_footResult, err := factory.FromDtoJSON(deciwatts_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DeciwattPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwatts_per_cubic_footResult.Convert(units.PowerDensityDeciwattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciwattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with DecawattPerCubicFoot unit
    decawatts_per_cubic_footJSON := []byte(`{"value": 100, "unit": "DecawattPerCubicFoot"}`)
    decawatts_per_cubic_footResult, err := factory.FromDtoJSON(decawatts_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecawattPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decawatts_per_cubic_footResult.Convert(units.PowerDensityDecawattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecawattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattPerCubicFoot unit
    kilowatts_per_cubic_footJSON := []byte(`{"value": 100, "unit": "KilowattPerCubicFoot"}`)
    kilowatts_per_cubic_footResult, err := factory.FromDtoJSON(kilowatts_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_cubic_footResult.Convert(units.PowerDensityKilowattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattPerCubicFoot unit
    megawatts_per_cubic_footJSON := []byte(`{"value": 100, "unit": "MegawattPerCubicFoot"}`)
    megawatts_per_cubic_footResult, err := factory.FromDtoJSON(megawatts_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_cubic_footResult.Convert(units.PowerDensityMegawattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattPerCubicFoot unit
    gigawatts_per_cubic_footJSON := []byte(`{"value": 100, "unit": "GigawattPerCubicFoot"}`)
    gigawatts_per_cubic_footResult, err := factory.FromDtoJSON(gigawatts_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_cubic_footResult.Convert(units.PowerDensityGigawattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with TerawattPerCubicFoot unit
    terawatts_per_cubic_footJSON := []byte(`{"value": 100, "unit": "TerawattPerCubicFoot"}`)
    terawatts_per_cubic_footResult, err := factory.FromDtoJSON(terawatts_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TerawattPerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatts_per_cubic_footResult.Convert(units.PowerDensityTerawattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattPerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with PicowattPerLiter unit
    picowatts_per_literJSON := []byte(`{"value": 100, "unit": "PicowattPerLiter"}`)
    picowatts_per_literResult, err := factory.FromDtoJSON(picowatts_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicowattPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowatts_per_literResult.Convert(units.PowerDensityPicowattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicowattPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with NanowattPerLiter unit
    nanowatts_per_literJSON := []byte(`{"value": 100, "unit": "NanowattPerLiter"}`)
    nanowatts_per_literResult, err := factory.FromDtoJSON(nanowatts_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanowattPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_literResult.Convert(units.PowerDensityNanowattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrowattPerLiter unit
    microwatts_per_literJSON := []byte(`{"value": 100, "unit": "MicrowattPerLiter"}`)
    microwatts_per_literResult, err := factory.FromDtoJSON(microwatts_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrowattPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_literResult.Convert(units.PowerDensityMicrowattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with MilliwattPerLiter unit
    milliwatts_per_literJSON := []byte(`{"value": 100, "unit": "MilliwattPerLiter"}`)
    milliwatts_per_literResult, err := factory.FromDtoJSON(milliwatts_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliwattPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_literResult.Convert(units.PowerDensityMilliwattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with DeciwattPerLiter unit
    deciwatts_per_literJSON := []byte(`{"value": 100, "unit": "DeciwattPerLiter"}`)
    deciwatts_per_literResult, err := factory.FromDtoJSON(deciwatts_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DeciwattPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwatts_per_literResult.Convert(units.PowerDensityDeciwattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciwattPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with DecawattPerLiter unit
    decawatts_per_literJSON := []byte(`{"value": 100, "unit": "DecawattPerLiter"}`)
    decawatts_per_literResult, err := factory.FromDtoJSON(decawatts_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecawattPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decawatts_per_literResult.Convert(units.PowerDensityDecawattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecawattPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattPerLiter unit
    kilowatts_per_literJSON := []byte(`{"value": 100, "unit": "KilowattPerLiter"}`)
    kilowatts_per_literResult, err := factory.FromDtoJSON(kilowatts_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_literResult.Convert(units.PowerDensityKilowattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattPerLiter unit
    megawatts_per_literJSON := []byte(`{"value": 100, "unit": "MegawattPerLiter"}`)
    megawatts_per_literResult, err := factory.FromDtoJSON(megawatts_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_literResult.Convert(units.PowerDensityMegawattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattPerLiter unit
    gigawatts_per_literJSON := []byte(`{"value": 100, "unit": "GigawattPerLiter"}`)
    gigawatts_per_literResult, err := factory.FromDtoJSON(gigawatts_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatts_per_literResult.Convert(units.PowerDensityGigawattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattPerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with TerawattPerLiter unit
    terawatts_per_literJSON := []byte(`{"value": 100, "unit": "TerawattPerLiter"}`)
    terawatts_per_literResult, err := factory.FromDtoJSON(terawatts_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TerawattPerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatts_per_literResult.Convert(units.PowerDensityTerawattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattPerLiter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "WattPerCubicMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromWattsPerCubicMeter function
func TestPowerDensityFactory_FromWattsPerCubicMeter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromWattsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityWattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromWattsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromWattsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromWattsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityWattPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromWattsPerCubicInch function
func TestPowerDensityFactory_FromWattsPerCubicInch(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromWattsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityWattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromWattsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromWattsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromWattsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityWattPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromWattsPerCubicFoot function
func TestPowerDensityFactory_FromWattsPerCubicFoot(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromWattsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityWattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromWattsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromWattsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromWattsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityWattPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromWattsPerLiter function
func TestPowerDensityFactory_FromWattsPerLiter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerLiter(100)
    if err != nil {
        t.Errorf("FromWattsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityWattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromWattsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromWattsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerLiter(0)
    if err != nil {
        t.Errorf("FromWattsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityWattPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicowattsPerCubicMeter function
func TestPowerDensityFactory_FromPicowattsPerCubicMeter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicowattsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromPicowattsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityPicowattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicowattsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicowattsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromPicowattsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromPicowattsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromPicowattsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromPicowattsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicowattsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicowattsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromPicowattsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityPicowattPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicowattsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanowattsPerCubicMeter function
func TestPowerDensityFactory_FromNanowattsPerCubicMeter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanowattsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromNanowattsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityNanowattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanowattsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanowattsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromNanowattsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromNanowattsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromNanowattsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromNanowattsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanowattsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanowattsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromNanowattsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityNanowattPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanowattsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrowattsPerCubicMeter function
func TestPowerDensityFactory_FromMicrowattsPerCubicMeter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrowattsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMicrowattsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityMicrowattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrowattsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrowattsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromMicrowattsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromMicrowattsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrowattsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromMicrowattsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrowattsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrowattsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromMicrowattsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityMicrowattPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrowattsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwattsPerCubicMeter function
func TestPowerDensityFactory_FromMilliwattsPerCubicMeter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwattsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMilliwattsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityMilliwattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwattsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwattsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromMilliwattsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromMilliwattsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwattsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromMilliwattsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwattsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwattsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromMilliwattsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityMilliwattPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwattsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromDeciwattsPerCubicMeter function
func TestPowerDensityFactory_FromDeciwattsPerCubicMeter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDeciwattsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromDeciwattsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityDeciwattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDeciwattsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDeciwattsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromDeciwattsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromDeciwattsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromDeciwattsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromDeciwattsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromDeciwattsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDeciwattsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromDeciwattsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityDeciwattPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDeciwattsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecawattsPerCubicMeter function
func TestPowerDensityFactory_FromDecawattsPerCubicMeter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecawattsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromDecawattsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityDecawattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecawattsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecawattsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromDecawattsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromDecawattsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromDecawattsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromDecawattsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecawattsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecawattsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromDecawattsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityDecawattPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecawattsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattsPerCubicMeter function
func TestPowerDensityFactory_FromKilowattsPerCubicMeter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromKilowattsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityKilowattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromKilowattsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromKilowattsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilowattsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromKilowattsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityKilowattPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattsPerCubicMeter function
func TestPowerDensityFactory_FromMegawattsPerCubicMeter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMegawattsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityMegawattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromMegawattsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromMegawattsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromMegawattsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromMegawattsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityMegawattPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattsPerCubicMeter function
func TestPowerDensityFactory_FromGigawattsPerCubicMeter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromGigawattsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityGigawattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromGigawattsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromGigawattsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromGigawattsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromGigawattsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityGigawattPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTerawattsPerCubicMeter function
func TestPowerDensityFactory_FromTerawattsPerCubicMeter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerawattsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromTerawattsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityTerawattPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerawattsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerawattsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromTerawattsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromTerawattsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromTerawattsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromTerawattsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTerawattsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerawattsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromTerawattsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityTerawattPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerawattsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicowattsPerCubicInch function
func TestPowerDensityFactory_FromPicowattsPerCubicInch(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicowattsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromPicowattsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityPicowattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicowattsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicowattsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromPicowattsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromPicowattsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromPicowattsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromPicowattsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromPicowattsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicowattsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromPicowattsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityPicowattPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicowattsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromNanowattsPerCubicInch function
func TestPowerDensityFactory_FromNanowattsPerCubicInch(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanowattsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromNanowattsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityNanowattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanowattsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanowattsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromNanowattsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromNanowattsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromNanowattsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromNanowattsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromNanowattsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanowattsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromNanowattsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityNanowattPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanowattsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrowattsPerCubicInch function
func TestPowerDensityFactory_FromMicrowattsPerCubicInch(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrowattsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromMicrowattsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityMicrowattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrowattsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrowattsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromMicrowattsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromMicrowattsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromMicrowattsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromMicrowattsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrowattsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrowattsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromMicrowattsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityMicrowattPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrowattsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwattsPerCubicInch function
func TestPowerDensityFactory_FromMilliwattsPerCubicInch(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwattsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromMilliwattsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityMilliwattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwattsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwattsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromMilliwattsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromMilliwattsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwattsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromMilliwattsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwattsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwattsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromMilliwattsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityMilliwattPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwattsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromDeciwattsPerCubicInch function
func TestPowerDensityFactory_FromDeciwattsPerCubicInch(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDeciwattsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromDeciwattsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityDeciwattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDeciwattsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDeciwattsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromDeciwattsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromDeciwattsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromDeciwattsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromDeciwattsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromDeciwattsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDeciwattsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromDeciwattsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityDeciwattPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDeciwattsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromDecawattsPerCubicInch function
func TestPowerDensityFactory_FromDecawattsPerCubicInch(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecawattsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromDecawattsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityDecawattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecawattsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecawattsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromDecawattsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromDecawattsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromDecawattsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromDecawattsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromDecawattsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecawattsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromDecawattsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityDecawattPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecawattsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattsPerCubicInch function
func TestPowerDensityFactory_FromKilowattsPerCubicInch(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromKilowattsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityKilowattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromKilowattsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromKilowattsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromKilowattsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromKilowattsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityKilowattPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattsPerCubicInch function
func TestPowerDensityFactory_FromMegawattsPerCubicInch(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromMegawattsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityMegawattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromMegawattsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromMegawattsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromMegawattsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromMegawattsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityMegawattPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattsPerCubicInch function
func TestPowerDensityFactory_FromGigawattsPerCubicInch(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromGigawattsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityGigawattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromGigawattsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromGigawattsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromGigawattsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromGigawattsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityGigawattPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromTerawattsPerCubicInch function
func TestPowerDensityFactory_FromTerawattsPerCubicInch(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerawattsPerCubicInch(100)
    if err != nil {
        t.Errorf("FromTerawattsPerCubicInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityTerawattPerCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerawattsPerCubicInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerawattsPerCubicInch(math.NaN())
    if err == nil {
        t.Error("FromTerawattsPerCubicInch() with NaN value should return error")
    }

    _, err = factory.FromTerawattsPerCubicInch(math.Inf(1))
    if err == nil {
        t.Error("FromTerawattsPerCubicInch() with +Inf value should return error")
    }

    _, err = factory.FromTerawattsPerCubicInch(math.Inf(-1))
    if err == nil {
        t.Error("FromTerawattsPerCubicInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerawattsPerCubicInch(0)
    if err != nil {
        t.Errorf("FromTerawattsPerCubicInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityTerawattPerCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerawattsPerCubicInch() with zero value = %v, want 0", converted)
    }
}
// Test FromPicowattsPerCubicFoot function
func TestPowerDensityFactory_FromPicowattsPerCubicFoot(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicowattsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromPicowattsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityPicowattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicowattsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicowattsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromPicowattsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromPicowattsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromPicowattsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromPicowattsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromPicowattsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicowattsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromPicowattsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityPicowattPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicowattsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromNanowattsPerCubicFoot function
func TestPowerDensityFactory_FromNanowattsPerCubicFoot(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanowattsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromNanowattsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityNanowattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanowattsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanowattsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromNanowattsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromNanowattsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromNanowattsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromNanowattsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromNanowattsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanowattsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromNanowattsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityNanowattPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanowattsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrowattsPerCubicFoot function
func TestPowerDensityFactory_FromMicrowattsPerCubicFoot(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrowattsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromMicrowattsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityMicrowattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrowattsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrowattsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromMicrowattsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromMicrowattsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromMicrowattsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromMicrowattsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrowattsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrowattsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromMicrowattsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityMicrowattPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrowattsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwattsPerCubicFoot function
func TestPowerDensityFactory_FromMilliwattsPerCubicFoot(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwattsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromMilliwattsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityMilliwattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwattsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwattsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromMilliwattsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromMilliwattsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwattsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromMilliwattsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwattsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwattsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromMilliwattsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityMilliwattPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwattsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromDeciwattsPerCubicFoot function
func TestPowerDensityFactory_FromDeciwattsPerCubicFoot(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDeciwattsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromDeciwattsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityDeciwattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDeciwattsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDeciwattsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromDeciwattsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromDeciwattsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromDeciwattsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromDeciwattsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromDeciwattsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDeciwattsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromDeciwattsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityDeciwattPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDeciwattsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromDecawattsPerCubicFoot function
func TestPowerDensityFactory_FromDecawattsPerCubicFoot(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecawattsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromDecawattsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityDecawattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecawattsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecawattsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromDecawattsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromDecawattsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromDecawattsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromDecawattsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromDecawattsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecawattsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromDecawattsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityDecawattPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecawattsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattsPerCubicFoot function
func TestPowerDensityFactory_FromKilowattsPerCubicFoot(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromKilowattsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityKilowattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromKilowattsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromKilowattsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromKilowattsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromKilowattsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityKilowattPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattsPerCubicFoot function
func TestPowerDensityFactory_FromMegawattsPerCubicFoot(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromMegawattsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityMegawattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromMegawattsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromMegawattsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromMegawattsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromMegawattsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityMegawattPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattsPerCubicFoot function
func TestPowerDensityFactory_FromGigawattsPerCubicFoot(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromGigawattsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityGigawattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromGigawattsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromGigawattsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromGigawattsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromGigawattsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityGigawattPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromTerawattsPerCubicFoot function
func TestPowerDensityFactory_FromTerawattsPerCubicFoot(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerawattsPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromTerawattsPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityTerawattPerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerawattsPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerawattsPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromTerawattsPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromTerawattsPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromTerawattsPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromTerawattsPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromTerawattsPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerawattsPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromTerawattsPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityTerawattPerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerawattsPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromPicowattsPerLiter function
func TestPowerDensityFactory_FromPicowattsPerLiter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicowattsPerLiter(100)
    if err != nil {
        t.Errorf("FromPicowattsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityPicowattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicowattsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicowattsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromPicowattsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromPicowattsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromPicowattsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromPicowattsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicowattsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicowattsPerLiter(0)
    if err != nil {
        t.Errorf("FromPicowattsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityPicowattPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicowattsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanowattsPerLiter function
func TestPowerDensityFactory_FromNanowattsPerLiter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanowattsPerLiter(100)
    if err != nil {
        t.Errorf("FromNanowattsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityNanowattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanowattsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanowattsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromNanowattsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromNanowattsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromNanowattsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromNanowattsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanowattsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanowattsPerLiter(0)
    if err != nil {
        t.Errorf("FromNanowattsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityNanowattPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanowattsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrowattsPerLiter function
func TestPowerDensityFactory_FromMicrowattsPerLiter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrowattsPerLiter(100)
    if err != nil {
        t.Errorf("FromMicrowattsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityMicrowattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrowattsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrowattsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromMicrowattsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromMicrowattsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrowattsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromMicrowattsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrowattsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrowattsPerLiter(0)
    if err != nil {
        t.Errorf("FromMicrowattsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityMicrowattPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrowattsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwattsPerLiter function
func TestPowerDensityFactory_FromMilliwattsPerLiter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwattsPerLiter(100)
    if err != nil {
        t.Errorf("FromMilliwattsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityMilliwattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwattsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwattsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromMilliwattsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromMilliwattsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwattsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromMilliwattsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwattsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwattsPerLiter(0)
    if err != nil {
        t.Errorf("FromMilliwattsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityMilliwattPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwattsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromDeciwattsPerLiter function
func TestPowerDensityFactory_FromDeciwattsPerLiter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDeciwattsPerLiter(100)
    if err != nil {
        t.Errorf("FromDeciwattsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityDeciwattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDeciwattsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDeciwattsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromDeciwattsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromDeciwattsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromDeciwattsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromDeciwattsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromDeciwattsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDeciwattsPerLiter(0)
    if err != nil {
        t.Errorf("FromDeciwattsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityDeciwattPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDeciwattsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecawattsPerLiter function
func TestPowerDensityFactory_FromDecawattsPerLiter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecawattsPerLiter(100)
    if err != nil {
        t.Errorf("FromDecawattsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityDecawattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecawattsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecawattsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromDecawattsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromDecawattsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromDecawattsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromDecawattsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecawattsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecawattsPerLiter(0)
    if err != nil {
        t.Errorf("FromDecawattsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityDecawattPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecawattsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattsPerLiter function
func TestPowerDensityFactory_FromKilowattsPerLiter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattsPerLiter(100)
    if err != nil {
        t.Errorf("FromKilowattsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityKilowattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromKilowattsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromKilowattsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromKilowattsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattsPerLiter(0)
    if err != nil {
        t.Errorf("FromKilowattsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityKilowattPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattsPerLiter function
func TestPowerDensityFactory_FromMegawattsPerLiter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattsPerLiter(100)
    if err != nil {
        t.Errorf("FromMegawattsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityMegawattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromMegawattsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromMegawattsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromMegawattsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattsPerLiter(0)
    if err != nil {
        t.Errorf("FromMegawattsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityMegawattPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattsPerLiter function
func TestPowerDensityFactory_FromGigawattsPerLiter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattsPerLiter(100)
    if err != nil {
        t.Errorf("FromGigawattsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityGigawattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromGigawattsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromGigawattsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromGigawattsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattsPerLiter(0)
    if err != nil {
        t.Errorf("FromGigawattsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityGigawattPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattsPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromTerawattsPerLiter function
func TestPowerDensityFactory_FromTerawattsPerLiter(t *testing.T) {
    factory := units.PowerDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerawattsPerLiter(100)
    if err != nil {
        t.Errorf("FromTerawattsPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDensityTerawattPerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerawattsPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerawattsPerLiter(math.NaN())
    if err == nil {
        t.Error("FromTerawattsPerLiter() with NaN value should return error")
    }

    _, err = factory.FromTerawattsPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromTerawattsPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromTerawattsPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromTerawattsPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerawattsPerLiter(0)
    if err != nil {
        t.Errorf("FromTerawattsPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDensityTerawattPerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerawattsPerLiter() with zero value = %v, want 0", converted)
    }
}

func TestPowerDensityToString(t *testing.T) {
	factory := units.PowerDensityFactory{}
	a, err := factory.CreatePowerDensity(45, units.PowerDensityWattPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PowerDensityWattPerCubicMeter, 2)
	expected := "45.00 " + units.GetPowerDensityAbbreviation(units.PowerDensityWattPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PowerDensityWattPerCubicMeter, -1)
	expected = "45 " + units.GetPowerDensityAbbreviation(units.PowerDensityWattPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPowerDensity_EqualityAndComparison(t *testing.T) {
	factory := units.PowerDensityFactory{}
	a1, _ := factory.CreatePowerDensity(60, units.PowerDensityWattPerCubicMeter)
	a2, _ := factory.CreatePowerDensity(60, units.PowerDensityWattPerCubicMeter)
	a3, _ := factory.CreatePowerDensity(120, units.PowerDensityWattPerCubicMeter)

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

func TestPowerDensity_Arithmetic(t *testing.T) {
	factory := units.PowerDensityFactory{}
	a1, _ := factory.CreatePowerDensity(30, units.PowerDensityWattPerCubicMeter)
	a2, _ := factory.CreatePowerDensity(45, units.PowerDensityWattPerCubicMeter)

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


func TestGetPowerDensityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.PowerDensityUnits
        want string
    }{
        {
            name: "WattPerCubicMeter abbreviation",
            unit: units.PowerDensityWattPerCubicMeter,
            want: "W/m³",
        },
        {
            name: "WattPerCubicInch abbreviation",
            unit: units.PowerDensityWattPerCubicInch,
            want: "W/in³",
        },
        {
            name: "WattPerCubicFoot abbreviation",
            unit: units.PowerDensityWattPerCubicFoot,
            want: "W/ft³",
        },
        {
            name: "WattPerLiter abbreviation",
            unit: units.PowerDensityWattPerLiter,
            want: "W/l",
        },
        {
            name: "PicowattPerCubicMeter abbreviation",
            unit: units.PowerDensityPicowattPerCubicMeter,
            want: "pW/m³",
        },
        {
            name: "NanowattPerCubicMeter abbreviation",
            unit: units.PowerDensityNanowattPerCubicMeter,
            want: "nW/m³",
        },
        {
            name: "MicrowattPerCubicMeter abbreviation",
            unit: units.PowerDensityMicrowattPerCubicMeter,
            want: "μW/m³",
        },
        {
            name: "MilliwattPerCubicMeter abbreviation",
            unit: units.PowerDensityMilliwattPerCubicMeter,
            want: "mW/m³",
        },
        {
            name: "DeciwattPerCubicMeter abbreviation",
            unit: units.PowerDensityDeciwattPerCubicMeter,
            want: "dW/m³",
        },
        {
            name: "DecawattPerCubicMeter abbreviation",
            unit: units.PowerDensityDecawattPerCubicMeter,
            want: "daW/m³",
        },
        {
            name: "KilowattPerCubicMeter abbreviation",
            unit: units.PowerDensityKilowattPerCubicMeter,
            want: "kW/m³",
        },
        {
            name: "MegawattPerCubicMeter abbreviation",
            unit: units.PowerDensityMegawattPerCubicMeter,
            want: "MW/m³",
        },
        {
            name: "GigawattPerCubicMeter abbreviation",
            unit: units.PowerDensityGigawattPerCubicMeter,
            want: "GW/m³",
        },
        {
            name: "TerawattPerCubicMeter abbreviation",
            unit: units.PowerDensityTerawattPerCubicMeter,
            want: "TW/m³",
        },
        {
            name: "PicowattPerCubicInch abbreviation",
            unit: units.PowerDensityPicowattPerCubicInch,
            want: "pW/in³",
        },
        {
            name: "NanowattPerCubicInch abbreviation",
            unit: units.PowerDensityNanowattPerCubicInch,
            want: "nW/in³",
        },
        {
            name: "MicrowattPerCubicInch abbreviation",
            unit: units.PowerDensityMicrowattPerCubicInch,
            want: "μW/in³",
        },
        {
            name: "MilliwattPerCubicInch abbreviation",
            unit: units.PowerDensityMilliwattPerCubicInch,
            want: "mW/in³",
        },
        {
            name: "DeciwattPerCubicInch abbreviation",
            unit: units.PowerDensityDeciwattPerCubicInch,
            want: "dW/in³",
        },
        {
            name: "DecawattPerCubicInch abbreviation",
            unit: units.PowerDensityDecawattPerCubicInch,
            want: "daW/in³",
        },
        {
            name: "KilowattPerCubicInch abbreviation",
            unit: units.PowerDensityKilowattPerCubicInch,
            want: "kW/in³",
        },
        {
            name: "MegawattPerCubicInch abbreviation",
            unit: units.PowerDensityMegawattPerCubicInch,
            want: "MW/in³",
        },
        {
            name: "GigawattPerCubicInch abbreviation",
            unit: units.PowerDensityGigawattPerCubicInch,
            want: "GW/in³",
        },
        {
            name: "TerawattPerCubicInch abbreviation",
            unit: units.PowerDensityTerawattPerCubicInch,
            want: "TW/in³",
        },
        {
            name: "PicowattPerCubicFoot abbreviation",
            unit: units.PowerDensityPicowattPerCubicFoot,
            want: "pW/ft³",
        },
        {
            name: "NanowattPerCubicFoot abbreviation",
            unit: units.PowerDensityNanowattPerCubicFoot,
            want: "nW/ft³",
        },
        {
            name: "MicrowattPerCubicFoot abbreviation",
            unit: units.PowerDensityMicrowattPerCubicFoot,
            want: "μW/ft³",
        },
        {
            name: "MilliwattPerCubicFoot abbreviation",
            unit: units.PowerDensityMilliwattPerCubicFoot,
            want: "mW/ft³",
        },
        {
            name: "DeciwattPerCubicFoot abbreviation",
            unit: units.PowerDensityDeciwattPerCubicFoot,
            want: "dW/ft³",
        },
        {
            name: "DecawattPerCubicFoot abbreviation",
            unit: units.PowerDensityDecawattPerCubicFoot,
            want: "daW/ft³",
        },
        {
            name: "KilowattPerCubicFoot abbreviation",
            unit: units.PowerDensityKilowattPerCubicFoot,
            want: "kW/ft³",
        },
        {
            name: "MegawattPerCubicFoot abbreviation",
            unit: units.PowerDensityMegawattPerCubicFoot,
            want: "MW/ft³",
        },
        {
            name: "GigawattPerCubicFoot abbreviation",
            unit: units.PowerDensityGigawattPerCubicFoot,
            want: "GW/ft³",
        },
        {
            name: "TerawattPerCubicFoot abbreviation",
            unit: units.PowerDensityTerawattPerCubicFoot,
            want: "TW/ft³",
        },
        {
            name: "PicowattPerLiter abbreviation",
            unit: units.PowerDensityPicowattPerLiter,
            want: "pW/l",
        },
        {
            name: "NanowattPerLiter abbreviation",
            unit: units.PowerDensityNanowattPerLiter,
            want: "nW/l",
        },
        {
            name: "MicrowattPerLiter abbreviation",
            unit: units.PowerDensityMicrowattPerLiter,
            want: "μW/l",
        },
        {
            name: "MilliwattPerLiter abbreviation",
            unit: units.PowerDensityMilliwattPerLiter,
            want: "mW/l",
        },
        {
            name: "DeciwattPerLiter abbreviation",
            unit: units.PowerDensityDeciwattPerLiter,
            want: "dW/l",
        },
        {
            name: "DecawattPerLiter abbreviation",
            unit: units.PowerDensityDecawattPerLiter,
            want: "daW/l",
        },
        {
            name: "KilowattPerLiter abbreviation",
            unit: units.PowerDensityKilowattPerLiter,
            want: "kW/l",
        },
        {
            name: "MegawattPerLiter abbreviation",
            unit: units.PowerDensityMegawattPerLiter,
            want: "MW/l",
        },
        {
            name: "GigawattPerLiter abbreviation",
            unit: units.PowerDensityGigawattPerLiter,
            want: "GW/l",
        },
        {
            name: "TerawattPerLiter abbreviation",
            unit: units.PowerDensityTerawattPerLiter,
            want: "TW/l",
        },
        {
            name: "invalid unit",
            unit: units.PowerDensityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetPowerDensityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetPowerDensityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestPowerDensity_String(t *testing.T) {
    factory := units.PowerDensityFactory{}
    
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
            unit, err := factory.CreatePowerDensity(tt.value, units.PowerDensityWattPerCubicMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("PowerDensity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestPowerDensity_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.PowerDensityFactory{}

	_, err := uf.CreatePowerDensity(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}