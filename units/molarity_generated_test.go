// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMolarityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MolePerCubicMeter"}`
	
	factory := units.MolarityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MolarityMolePerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.MolarityMolePerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MolePerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMolarityDto_ToJSON(t *testing.T) {
	dto := units.MolarityDto{
		Value: 45,
		Unit:  units.MolarityMolePerCubicMeter,
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
	if result["unit"].(string) != string(units.MolarityMolePerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.MolarityMolePerCubicMeter, result["unit"])
	}
}

func TestNewMolarity_InvalidValue(t *testing.T) {
	factory := units.MolarityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMolarity(math.NaN(), units.MolarityMolePerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMolarity(math.Inf(1), units.MolarityMolePerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMolarityConversions(t *testing.T) {
	factory := units.MolarityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMolarity(180, units.MolarityMolePerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MolesPerCubicMeter.
		// No expected conversion value provided for MolesPerCubicMeter, verifying result is not NaN.
		result := a.MolesPerCubicMeter()
		cacheResult := a.MolesPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MolesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to MolesPerLiter.
		// No expected conversion value provided for MolesPerLiter, verifying result is not NaN.
		result := a.MolesPerLiter()
		cacheResult := a.MolesPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to PoundMolesPerCubicFoot.
		// No expected conversion value provided for PoundMolesPerCubicFoot, verifying result is not NaN.
		result := a.PoundMolesPerCubicFoot()
		cacheResult := a.PoundMolesPerCubicFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundMolesPerCubicFoot returned NaN")
		}
	}
	{
		// Test conversion to KilomolesPerCubicMeter.
		// No expected conversion value provided for KilomolesPerCubicMeter, verifying result is not NaN.
		result := a.KilomolesPerCubicMeter()
		cacheResult := a.KilomolesPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilomolesPerCubicMeter returned NaN")
		}
	}
	{
		// Test conversion to FemtomolesPerLiter.
		// No expected conversion value provided for FemtomolesPerLiter, verifying result is not NaN.
		result := a.FemtomolesPerLiter()
		cacheResult := a.FemtomolesPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to FemtomolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to PicomolesPerLiter.
		// No expected conversion value provided for PicomolesPerLiter, verifying result is not NaN.
		result := a.PicomolesPerLiter()
		cacheResult := a.PicomolesPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicomolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to NanomolesPerLiter.
		// No expected conversion value provided for NanomolesPerLiter, verifying result is not NaN.
		result := a.NanomolesPerLiter()
		cacheResult := a.NanomolesPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanomolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MicromolesPerLiter.
		// No expected conversion value provided for MicromolesPerLiter, verifying result is not NaN.
		result := a.MicromolesPerLiter()
		cacheResult := a.MicromolesPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicromolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to MillimolesPerLiter.
		// No expected conversion value provided for MillimolesPerLiter, verifying result is not NaN.
		result := a.MillimolesPerLiter()
		cacheResult := a.MillimolesPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to CentimolesPerLiter.
		// No expected conversion value provided for CentimolesPerLiter, verifying result is not NaN.
		result := a.CentimolesPerLiter()
		cacheResult := a.CentimolesPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentimolesPerLiter returned NaN")
		}
	}
	{
		// Test conversion to DecimolesPerLiter.
		// No expected conversion value provided for DecimolesPerLiter, verifying result is not NaN.
		result := a.DecimolesPerLiter()
		cacheResult := a.DecimolesPerLiter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecimolesPerLiter returned NaN")
		}
	}
}

func TestMolarity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MolarityFactory{}
	a, err := factory.CreateMolarity(90, units.MolarityMolePerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MolarityMolePerCubicMeter {
		t.Errorf("expected default unit MolePerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MolarityMolePerCubicMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MolarityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MolarityMolePerCubicMeter {
		t.Errorf("expected unit MolePerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMolarityFactory_FromDto(t *testing.T) {
    factory := units.MolarityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MolarityDto{
        Value: 100,
        Unit:  units.MolarityMolePerCubicMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MolarityDto{
        Value: math.NaN(),
        Unit:  units.MolarityMolePerCubicMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test MolePerCubicMeter conversion
    moles_per_cubic_meterDto := units.MolarityDto{
        Value: 100,
        Unit:  units.MolarityMolePerCubicMeter,
    }
    
    var moles_per_cubic_meterResult *units.Molarity
    moles_per_cubic_meterResult, err = factory.FromDto(moles_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MolePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_cubic_meterResult.Convert(units.MolarityMolePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test MolePerLiter conversion
    moles_per_literDto := units.MolarityDto{
        Value: 100,
        Unit:  units.MolarityMolePerLiter,
    }
    
    var moles_per_literResult *units.Molarity
    moles_per_literResult, err = factory.FromDto(moles_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MolePerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_literResult.Convert(units.MolarityMolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerLiter = %v, want %v", converted, 100)
    }
    // Test PoundMolePerCubicFoot conversion
    pound_moles_per_cubic_footDto := units.MolarityDto{
        Value: 100,
        Unit:  units.MolarityPoundMolePerCubicFoot,
    }
    
    var pound_moles_per_cubic_footResult *units.Molarity
    pound_moles_per_cubic_footResult, err = factory.FromDto(pound_moles_per_cubic_footDto)
    if err != nil {
        t.Errorf("FromDto() with PoundMolePerCubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_moles_per_cubic_footResult.Convert(units.MolarityPoundMolePerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundMolePerCubicFoot = %v, want %v", converted, 100)
    }
    // Test KilomolePerCubicMeter conversion
    kilomoles_per_cubic_meterDto := units.MolarityDto{
        Value: 100,
        Unit:  units.MolarityKilomolePerCubicMeter,
    }
    
    var kilomoles_per_cubic_meterResult *units.Molarity
    kilomoles_per_cubic_meterResult, err = factory.FromDto(kilomoles_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilomolePerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilomoles_per_cubic_meterResult.Convert(units.MolarityKilomolePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilomolePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test FemtomolePerLiter conversion
    femtomoles_per_literDto := units.MolarityDto{
        Value: 100,
        Unit:  units.MolarityFemtomolePerLiter,
    }
    
    var femtomoles_per_literResult *units.Molarity
    femtomoles_per_literResult, err = factory.FromDto(femtomoles_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with FemtomolePerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtomoles_per_literResult.Convert(units.MolarityFemtomolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FemtomolePerLiter = %v, want %v", converted, 100)
    }
    // Test PicomolePerLiter conversion
    picomoles_per_literDto := units.MolarityDto{
        Value: 100,
        Unit:  units.MolarityPicomolePerLiter,
    }
    
    var picomoles_per_literResult *units.Molarity
    picomoles_per_literResult, err = factory.FromDto(picomoles_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with PicomolePerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picomoles_per_literResult.Convert(units.MolarityPicomolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicomolePerLiter = %v, want %v", converted, 100)
    }
    // Test NanomolePerLiter conversion
    nanomoles_per_literDto := units.MolarityDto{
        Value: 100,
        Unit:  units.MolarityNanomolePerLiter,
    }
    
    var nanomoles_per_literResult *units.Molarity
    nanomoles_per_literResult, err = factory.FromDto(nanomoles_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with NanomolePerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanomoles_per_literResult.Convert(units.MolarityNanomolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanomolePerLiter = %v, want %v", converted, 100)
    }
    // Test MicromolePerLiter conversion
    micromoles_per_literDto := units.MolarityDto{
        Value: 100,
        Unit:  units.MolarityMicromolePerLiter,
    }
    
    var micromoles_per_literResult *units.Molarity
    micromoles_per_literResult, err = factory.FromDto(micromoles_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MicromolePerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micromoles_per_literResult.Convert(units.MolarityMicromolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicromolePerLiter = %v, want %v", converted, 100)
    }
    // Test MillimolePerLiter conversion
    millimoles_per_literDto := units.MolarityDto{
        Value: 100,
        Unit:  units.MolarityMillimolePerLiter,
    }
    
    var millimoles_per_literResult *units.Molarity
    millimoles_per_literResult, err = factory.FromDto(millimoles_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with MillimolePerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimoles_per_literResult.Convert(units.MolarityMillimolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimolePerLiter = %v, want %v", converted, 100)
    }
    // Test CentimolePerLiter conversion
    centimoles_per_literDto := units.MolarityDto{
        Value: 100,
        Unit:  units.MolarityCentimolePerLiter,
    }
    
    var centimoles_per_literResult *units.Molarity
    centimoles_per_literResult, err = factory.FromDto(centimoles_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with CentimolePerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimoles_per_literResult.Convert(units.MolarityCentimolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimolePerLiter = %v, want %v", converted, 100)
    }
    // Test DecimolePerLiter conversion
    decimoles_per_literDto := units.MolarityDto{
        Value: 100,
        Unit:  units.MolarityDecimolePerLiter,
    }
    
    var decimoles_per_literResult *units.Molarity
    decimoles_per_literResult, err = factory.FromDto(decimoles_per_literDto)
    if err != nil {
        t.Errorf("FromDto() with DecimolePerLiter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimoles_per_literResult.Convert(units.MolarityDecimolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimolePerLiter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MolarityDto{
        Value: 0,
        Unit:  units.MolarityMolePerCubicMeter,
    }
    
    var zeroResult *units.Molarity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMolarityFactory_FromDtoJSON(t *testing.T) {
    factory := units.MolarityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "MolePerCubicMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "MolePerCubicMeter"}`)
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
    nanJSON, _ := json.Marshal(units.MolarityDto{
        Value: nanValue,
        Unit:  units.MolarityMolePerCubicMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with MolePerCubicMeter unit
    moles_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "MolePerCubicMeter"}`)
    moles_per_cubic_meterResult, err := factory.FromDtoJSON(moles_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MolePerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_cubic_meterResult.Convert(units.MolarityMolePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MolePerLiter unit
    moles_per_literJSON := []byte(`{"value": 100, "unit": "MolePerLiter"}`)
    moles_per_literResult, err := factory.FromDtoJSON(moles_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MolePerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_literResult.Convert(units.MolarityMolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundMolePerCubicFoot unit
    pound_moles_per_cubic_footJSON := []byte(`{"value": 100, "unit": "PoundMolePerCubicFoot"}`)
    pound_moles_per_cubic_footResult, err := factory.FromDtoJSON(pound_moles_per_cubic_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundMolePerCubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_moles_per_cubic_footResult.Convert(units.MolarityPoundMolePerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundMolePerCubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with KilomolePerCubicMeter unit
    kilomoles_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "KilomolePerCubicMeter"}`)
    kilomoles_per_cubic_meterResult, err := factory.FromDtoJSON(kilomoles_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilomolePerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilomoles_per_cubic_meterResult.Convert(units.MolarityKilomolePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilomolePerCubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with FemtomolePerLiter unit
    femtomoles_per_literJSON := []byte(`{"value": 100, "unit": "FemtomolePerLiter"}`)
    femtomoles_per_literResult, err := factory.FromDtoJSON(femtomoles_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FemtomolePerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtomoles_per_literResult.Convert(units.MolarityFemtomolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FemtomolePerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with PicomolePerLiter unit
    picomoles_per_literJSON := []byte(`{"value": 100, "unit": "PicomolePerLiter"}`)
    picomoles_per_literResult, err := factory.FromDtoJSON(picomoles_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicomolePerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picomoles_per_literResult.Convert(units.MolarityPicomolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicomolePerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with NanomolePerLiter unit
    nanomoles_per_literJSON := []byte(`{"value": 100, "unit": "NanomolePerLiter"}`)
    nanomoles_per_literResult, err := factory.FromDtoJSON(nanomoles_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanomolePerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanomoles_per_literResult.Convert(units.MolarityNanomolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanomolePerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with MicromolePerLiter unit
    micromoles_per_literJSON := []byte(`{"value": 100, "unit": "MicromolePerLiter"}`)
    micromoles_per_literResult, err := factory.FromDtoJSON(micromoles_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicromolePerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micromoles_per_literResult.Convert(units.MolarityMicromolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicromolePerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with MillimolePerLiter unit
    millimoles_per_literJSON := []byte(`{"value": 100, "unit": "MillimolePerLiter"}`)
    millimoles_per_literResult, err := factory.FromDtoJSON(millimoles_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimolePerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimoles_per_literResult.Convert(units.MolarityMillimolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimolePerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with CentimolePerLiter unit
    centimoles_per_literJSON := []byte(`{"value": 100, "unit": "CentimolePerLiter"}`)
    centimoles_per_literResult, err := factory.FromDtoJSON(centimoles_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentimolePerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimoles_per_literResult.Convert(units.MolarityCentimolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimolePerLiter = %v, want %v", converted, 100)
    }
    // Test JSON with DecimolePerLiter unit
    decimoles_per_literJSON := []byte(`{"value": 100, "unit": "DecimolePerLiter"}`)
    decimoles_per_literResult, err := factory.FromDtoJSON(decimoles_per_literJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecimolePerLiter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimoles_per_literResult.Convert(units.MolarityDecimolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimolePerLiter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "MolePerCubicMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromMolesPerCubicMeter function
func TestMolarityFactory_FromMolesPerCubicMeter(t *testing.T) {
    factory := units.MolarityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMolesPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromMolesPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarityMolePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMolesPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMolesPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromMolesPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromMolesPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMolesPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromMolesPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMolesPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMolesPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromMolesPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarityMolePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMolesPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMolesPerLiter function
func TestMolarityFactory_FromMolesPerLiter(t *testing.T) {
    factory := units.MolarityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMolesPerLiter(100)
    if err != nil {
        t.Errorf("FromMolesPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarityMolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMolesPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMolesPerLiter(math.NaN())
    if err == nil {
        t.Error("FromMolesPerLiter() with NaN value should return error")
    }

    _, err = factory.FromMolesPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMolesPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromMolesPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMolesPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMolesPerLiter(0)
    if err != nil {
        t.Errorf("FromMolesPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarityMolePerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMolesPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundMolesPerCubicFoot function
func TestMolarityFactory_FromPoundMolesPerCubicFoot(t *testing.T) {
    factory := units.MolarityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundMolesPerCubicFoot(100)
    if err != nil {
        t.Errorf("FromPoundMolesPerCubicFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarityPoundMolePerCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundMolesPerCubicFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundMolesPerCubicFoot(math.NaN())
    if err == nil {
        t.Error("FromPoundMolesPerCubicFoot() with NaN value should return error")
    }

    _, err = factory.FromPoundMolesPerCubicFoot(math.Inf(1))
    if err == nil {
        t.Error("FromPoundMolesPerCubicFoot() with +Inf value should return error")
    }

    _, err = factory.FromPoundMolesPerCubicFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundMolesPerCubicFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundMolesPerCubicFoot(0)
    if err != nil {
        t.Errorf("FromPoundMolesPerCubicFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarityPoundMolePerCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundMolesPerCubicFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromKilomolesPerCubicMeter function
func TestMolarityFactory_FromKilomolesPerCubicMeter(t *testing.T) {
    factory := units.MolarityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilomolesPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromKilomolesPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarityKilomolePerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilomolesPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilomolesPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromKilomolesPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromKilomolesPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilomolesPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilomolesPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilomolesPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilomolesPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromKilomolesPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarityKilomolePerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilomolesPerCubicMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromFemtomolesPerLiter function
func TestMolarityFactory_FromFemtomolesPerLiter(t *testing.T) {
    factory := units.MolarityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFemtomolesPerLiter(100)
    if err != nil {
        t.Errorf("FromFemtomolesPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarityFemtomolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFemtomolesPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFemtomolesPerLiter(math.NaN())
    if err == nil {
        t.Error("FromFemtomolesPerLiter() with NaN value should return error")
    }

    _, err = factory.FromFemtomolesPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromFemtomolesPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromFemtomolesPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromFemtomolesPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFemtomolesPerLiter(0)
    if err != nil {
        t.Errorf("FromFemtomolesPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarityFemtomolePerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFemtomolesPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicomolesPerLiter function
func TestMolarityFactory_FromPicomolesPerLiter(t *testing.T) {
    factory := units.MolarityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicomolesPerLiter(100)
    if err != nil {
        t.Errorf("FromPicomolesPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarityPicomolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicomolesPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicomolesPerLiter(math.NaN())
    if err == nil {
        t.Error("FromPicomolesPerLiter() with NaN value should return error")
    }

    _, err = factory.FromPicomolesPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromPicomolesPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromPicomolesPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicomolesPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicomolesPerLiter(0)
    if err != nil {
        t.Errorf("FromPicomolesPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarityPicomolePerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicomolesPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanomolesPerLiter function
func TestMolarityFactory_FromNanomolesPerLiter(t *testing.T) {
    factory := units.MolarityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanomolesPerLiter(100)
    if err != nil {
        t.Errorf("FromNanomolesPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarityNanomolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanomolesPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanomolesPerLiter(math.NaN())
    if err == nil {
        t.Error("FromNanomolesPerLiter() with NaN value should return error")
    }

    _, err = factory.FromNanomolesPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromNanomolesPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromNanomolesPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanomolesPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanomolesPerLiter(0)
    if err != nil {
        t.Errorf("FromNanomolesPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarityNanomolePerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanomolesPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicromolesPerLiter function
func TestMolarityFactory_FromMicromolesPerLiter(t *testing.T) {
    factory := units.MolarityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicromolesPerLiter(100)
    if err != nil {
        t.Errorf("FromMicromolesPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarityMicromolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicromolesPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicromolesPerLiter(math.NaN())
    if err == nil {
        t.Error("FromMicromolesPerLiter() with NaN value should return error")
    }

    _, err = factory.FromMicromolesPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMicromolesPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromMicromolesPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicromolesPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicromolesPerLiter(0)
    if err != nil {
        t.Errorf("FromMicromolesPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarityMicromolePerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicromolesPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimolesPerLiter function
func TestMolarityFactory_FromMillimolesPerLiter(t *testing.T) {
    factory := units.MolarityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimolesPerLiter(100)
    if err != nil {
        t.Errorf("FromMillimolesPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarityMillimolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimolesPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimolesPerLiter(math.NaN())
    if err == nil {
        t.Error("FromMillimolesPerLiter() with NaN value should return error")
    }

    _, err = factory.FromMillimolesPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromMillimolesPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromMillimolesPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimolesPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimolesPerLiter(0)
    if err != nil {
        t.Errorf("FromMillimolesPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarityMillimolePerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimolesPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromCentimolesPerLiter function
func TestMolarityFactory_FromCentimolesPerLiter(t *testing.T) {
    factory := units.MolarityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentimolesPerLiter(100)
    if err != nil {
        t.Errorf("FromCentimolesPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarityCentimolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentimolesPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentimolesPerLiter(math.NaN())
    if err == nil {
        t.Error("FromCentimolesPerLiter() with NaN value should return error")
    }

    _, err = factory.FromCentimolesPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromCentimolesPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromCentimolesPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromCentimolesPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentimolesPerLiter(0)
    if err != nil {
        t.Errorf("FromCentimolesPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarityCentimolePerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentimolesPerLiter() with zero value = %v, want 0", converted)
    }
}
// Test FromDecimolesPerLiter function
func TestMolarityFactory_FromDecimolesPerLiter(t *testing.T) {
    factory := units.MolarityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecimolesPerLiter(100)
    if err != nil {
        t.Errorf("FromDecimolesPerLiter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarityDecimolePerLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecimolesPerLiter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecimolesPerLiter(math.NaN())
    if err == nil {
        t.Error("FromDecimolesPerLiter() with NaN value should return error")
    }

    _, err = factory.FromDecimolesPerLiter(math.Inf(1))
    if err == nil {
        t.Error("FromDecimolesPerLiter() with +Inf value should return error")
    }

    _, err = factory.FromDecimolesPerLiter(math.Inf(-1))
    if err == nil {
        t.Error("FromDecimolesPerLiter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecimolesPerLiter(0)
    if err != nil {
        t.Errorf("FromDecimolesPerLiter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarityDecimolePerLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecimolesPerLiter() with zero value = %v, want 0", converted)
    }
}

func TestMolarityToString(t *testing.T) {
	factory := units.MolarityFactory{}
	a, err := factory.CreateMolarity(45, units.MolarityMolePerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MolarityMolePerCubicMeter, 2)
	expected := "45.00 " + units.GetMolarityAbbreviation(units.MolarityMolePerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MolarityMolePerCubicMeter, -1)
	expected = "45 " + units.GetMolarityAbbreviation(units.MolarityMolePerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMolarity_EqualityAndComparison(t *testing.T) {
	factory := units.MolarityFactory{}
	a1, _ := factory.CreateMolarity(60, units.MolarityMolePerCubicMeter)
	a2, _ := factory.CreateMolarity(60, units.MolarityMolePerCubicMeter)
	a3, _ := factory.CreateMolarity(120, units.MolarityMolePerCubicMeter)

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

func TestMolarity_Arithmetic(t *testing.T) {
	factory := units.MolarityFactory{}
	a1, _ := factory.CreateMolarity(30, units.MolarityMolePerCubicMeter)
	a2, _ := factory.CreateMolarity(45, units.MolarityMolePerCubicMeter)

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


func TestGetMolarityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.MolarityUnits
        want string
    }{
        {
            name: "MolePerCubicMeter abbreviation",
            unit: units.MolarityMolePerCubicMeter,
            want: "mol/m³",
        },
        {
            name: "MolePerLiter abbreviation",
            unit: units.MolarityMolePerLiter,
            want: "mol/L",
        },
        {
            name: "PoundMolePerCubicFoot abbreviation",
            unit: units.MolarityPoundMolePerCubicFoot,
            want: "lbmol/ft³",
        },
        {
            name: "KilomolePerCubicMeter abbreviation",
            unit: units.MolarityKilomolePerCubicMeter,
            want: "kmol/m³",
        },
        {
            name: "FemtomolePerLiter abbreviation",
            unit: units.MolarityFemtomolePerLiter,
            want: "fmol/L",
        },
        {
            name: "PicomolePerLiter abbreviation",
            unit: units.MolarityPicomolePerLiter,
            want: "pmol/L",
        },
        {
            name: "NanomolePerLiter abbreviation",
            unit: units.MolarityNanomolePerLiter,
            want: "nmol/L",
        },
        {
            name: "MicromolePerLiter abbreviation",
            unit: units.MolarityMicromolePerLiter,
            want: "μmol/L",
        },
        {
            name: "MillimolePerLiter abbreviation",
            unit: units.MolarityMillimolePerLiter,
            want: "mmol/L",
        },
        {
            name: "CentimolePerLiter abbreviation",
            unit: units.MolarityCentimolePerLiter,
            want: "cmol/L",
        },
        {
            name: "DecimolePerLiter abbreviation",
            unit: units.MolarityDecimolePerLiter,
            want: "dmol/L",
        },
        {
            name: "invalid unit",
            unit: units.MolarityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetMolarityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetMolarityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestMolarity_String(t *testing.T) {
    factory := units.MolarityFactory{}
    
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
            unit, err := factory.CreateMolarity(tt.value, units.MolarityMolePerCubicMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Molarity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestMolarity_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.MolarityFactory{}

	_, err := uf.CreateMolarity(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}