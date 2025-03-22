// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricConductivityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SiemensPerMeter"}`
	
	factory := units.ElectricConductivityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricConductivitySiemensPerMeter {
		t.Errorf("expected unit %v, got %v", units.ElectricConductivitySiemensPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SiemensPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricConductivityDto_ToJSON(t *testing.T) {
	dto := units.ElectricConductivityDto{
		Value: 45,
		Unit:  units.ElectricConductivitySiemensPerMeter,
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
	if result["unit"].(string) != string(units.ElectricConductivitySiemensPerMeter) {
		t.Errorf("expected unit %s, got %v", units.ElectricConductivitySiemensPerMeter, result["unit"])
	}
}

func TestNewElectricConductivity_InvalidValue(t *testing.T) {
	factory := units.ElectricConductivityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricConductivity(math.NaN(), units.ElectricConductivitySiemensPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricConductivity(math.Inf(1), units.ElectricConductivitySiemensPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricConductivityConversions(t *testing.T) {
	factory := units.ElectricConductivityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricConductivity(180, units.ElectricConductivitySiemensPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to SiemensPerMeter.
		// No expected conversion value provided for SiemensPerMeter, verifying result is not NaN.
		result := a.SiemensPerMeter()
		cacheResult := a.SiemensPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SiemensPerMeter returned NaN")
		}
	}
	{
		// Test conversion to SiemensPerInch.
		// No expected conversion value provided for SiemensPerInch, verifying result is not NaN.
		result := a.SiemensPerInch()
		cacheResult := a.SiemensPerInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SiemensPerInch returned NaN")
		}
	}
	{
		// Test conversion to SiemensPerFoot.
		// No expected conversion value provided for SiemensPerFoot, verifying result is not NaN.
		result := a.SiemensPerFoot()
		cacheResult := a.SiemensPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SiemensPerFoot returned NaN")
		}
	}
	{
		// Test conversion to SiemensPerCentimeter.
		// No expected conversion value provided for SiemensPerCentimeter, verifying result is not NaN.
		result := a.SiemensPerCentimeter()
		cacheResult := a.SiemensPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SiemensPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MicrosiemensPerCentimeter.
		// No expected conversion value provided for MicrosiemensPerCentimeter, verifying result is not NaN.
		result := a.MicrosiemensPerCentimeter()
		cacheResult := a.MicrosiemensPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrosiemensPerCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MillisiemensPerCentimeter.
		// No expected conversion value provided for MillisiemensPerCentimeter, verifying result is not NaN.
		result := a.MillisiemensPerCentimeter()
		cacheResult := a.MillisiemensPerCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillisiemensPerCentimeter returned NaN")
		}
	}
}

func TestElectricConductivity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricConductivityFactory{}
	a, err := factory.CreateElectricConductivity(90, units.ElectricConductivitySiemensPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricConductivitySiemensPerMeter {
		t.Errorf("expected default unit SiemensPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricConductivitySiemensPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricConductivityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricConductivitySiemensPerMeter {
		t.Errorf("expected unit SiemensPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricConductivityFactory_FromDto(t *testing.T) {
    factory := units.ElectricConductivityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricConductivityDto{
        Value: 100,
        Unit:  units.ElectricConductivitySiemensPerMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricConductivityDto{
        Value: math.NaN(),
        Unit:  units.ElectricConductivitySiemensPerMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test SiemensPerMeter conversion
    siemens_per_meterDto := units.ElectricConductivityDto{
        Value: 100,
        Unit:  units.ElectricConductivitySiemensPerMeter,
    }
    
    var siemens_per_meterResult *units.ElectricConductivity
    siemens_per_meterResult, err = factory.FromDto(siemens_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with SiemensPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = siemens_per_meterResult.Convert(units.ElectricConductivitySiemensPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SiemensPerMeter = %v, want %v", converted, 100)
    }
    // Test SiemensPerInch conversion
    siemens_per_inchDto := units.ElectricConductivityDto{
        Value: 100,
        Unit:  units.ElectricConductivitySiemensPerInch,
    }
    
    var siemens_per_inchResult *units.ElectricConductivity
    siemens_per_inchResult, err = factory.FromDto(siemens_per_inchDto)
    if err != nil {
        t.Errorf("FromDto() with SiemensPerInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = siemens_per_inchResult.Convert(units.ElectricConductivitySiemensPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SiemensPerInch = %v, want %v", converted, 100)
    }
    // Test SiemensPerFoot conversion
    siemens_per_footDto := units.ElectricConductivityDto{
        Value: 100,
        Unit:  units.ElectricConductivitySiemensPerFoot,
    }
    
    var siemens_per_footResult *units.ElectricConductivity
    siemens_per_footResult, err = factory.FromDto(siemens_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with SiemensPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = siemens_per_footResult.Convert(units.ElectricConductivitySiemensPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SiemensPerFoot = %v, want %v", converted, 100)
    }
    // Test SiemensPerCentimeter conversion
    siemens_per_centimeterDto := units.ElectricConductivityDto{
        Value: 100,
        Unit:  units.ElectricConductivitySiemensPerCentimeter,
    }
    
    var siemens_per_centimeterResult *units.ElectricConductivity
    siemens_per_centimeterResult, err = factory.FromDto(siemens_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with SiemensPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = siemens_per_centimeterResult.Convert(units.ElectricConductivitySiemensPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SiemensPerCentimeter = %v, want %v", converted, 100)
    }
    // Test MicrosiemensPerCentimeter conversion
    microsiemens_per_centimeterDto := units.ElectricConductivityDto{
        Value: 100,
        Unit:  units.ElectricConductivityMicrosiemensPerCentimeter,
    }
    
    var microsiemens_per_centimeterResult *units.ElectricConductivity
    microsiemens_per_centimeterResult, err = factory.FromDto(microsiemens_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrosiemensPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microsiemens_per_centimeterResult.Convert(units.ElectricConductivityMicrosiemensPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrosiemensPerCentimeter = %v, want %v", converted, 100)
    }
    // Test MillisiemensPerCentimeter conversion
    millisiemens_per_centimeterDto := units.ElectricConductivityDto{
        Value: 100,
        Unit:  units.ElectricConductivityMillisiemensPerCentimeter,
    }
    
    var millisiemens_per_centimeterResult *units.ElectricConductivity
    millisiemens_per_centimeterResult, err = factory.FromDto(millisiemens_per_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MillisiemensPerCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millisiemens_per_centimeterResult.Convert(units.ElectricConductivityMillisiemensPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillisiemensPerCentimeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricConductivityDto{
        Value: 0,
        Unit:  units.ElectricConductivitySiemensPerMeter,
    }
    
    var zeroResult *units.ElectricConductivity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricConductivityFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricConductivityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "SiemensPerMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "SiemensPerMeter"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricConductivityDto{
        Value: nanValue,
        Unit:  units.ElectricConductivitySiemensPerMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with SiemensPerMeter unit
    siemens_per_meterJSON := []byte(`{"value": 100, "unit": "SiemensPerMeter"}`)
    siemens_per_meterResult, err := factory.FromDtoJSON(siemens_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SiemensPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = siemens_per_meterResult.Convert(units.ElectricConductivitySiemensPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SiemensPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with SiemensPerInch unit
    siemens_per_inchJSON := []byte(`{"value": 100, "unit": "SiemensPerInch"}`)
    siemens_per_inchResult, err := factory.FromDtoJSON(siemens_per_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SiemensPerInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = siemens_per_inchResult.Convert(units.ElectricConductivitySiemensPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SiemensPerInch = %v, want %v", converted, 100)
    }
    // Test JSON with SiemensPerFoot unit
    siemens_per_footJSON := []byte(`{"value": 100, "unit": "SiemensPerFoot"}`)
    siemens_per_footResult, err := factory.FromDtoJSON(siemens_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SiemensPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = siemens_per_footResult.Convert(units.ElectricConductivitySiemensPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SiemensPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with SiemensPerCentimeter unit
    siemens_per_centimeterJSON := []byte(`{"value": 100, "unit": "SiemensPerCentimeter"}`)
    siemens_per_centimeterResult, err := factory.FromDtoJSON(siemens_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SiemensPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = siemens_per_centimeterResult.Convert(units.ElectricConductivitySiemensPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SiemensPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrosiemensPerCentimeter unit
    microsiemens_per_centimeterJSON := []byte(`{"value": 100, "unit": "MicrosiemensPerCentimeter"}`)
    microsiemens_per_centimeterResult, err := factory.FromDtoJSON(microsiemens_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrosiemensPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microsiemens_per_centimeterResult.Convert(units.ElectricConductivityMicrosiemensPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrosiemensPerCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MillisiemensPerCentimeter unit
    millisiemens_per_centimeterJSON := []byte(`{"value": 100, "unit": "MillisiemensPerCentimeter"}`)
    millisiemens_per_centimeterResult, err := factory.FromDtoJSON(millisiemens_per_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillisiemensPerCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millisiemens_per_centimeterResult.Convert(units.ElectricConductivityMillisiemensPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillisiemensPerCentimeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "SiemensPerMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromSiemensPerMeter function
func TestElectricConductivityFactory_FromSiemensPerMeter(t *testing.T) {
    factory := units.ElectricConductivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSiemensPerMeter(100)
    if err != nil {
        t.Errorf("FromSiemensPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductivitySiemensPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSiemensPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSiemensPerMeter(math.NaN())
    if err == nil {
        t.Error("FromSiemensPerMeter() with NaN value should return error")
    }

    _, err = factory.FromSiemensPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromSiemensPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromSiemensPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromSiemensPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSiemensPerMeter(0)
    if err != nil {
        t.Errorf("FromSiemensPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricConductivitySiemensPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSiemensPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromSiemensPerInch function
func TestElectricConductivityFactory_FromSiemensPerInch(t *testing.T) {
    factory := units.ElectricConductivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSiemensPerInch(100)
    if err != nil {
        t.Errorf("FromSiemensPerInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductivitySiemensPerInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSiemensPerInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSiemensPerInch(math.NaN())
    if err == nil {
        t.Error("FromSiemensPerInch() with NaN value should return error")
    }

    _, err = factory.FromSiemensPerInch(math.Inf(1))
    if err == nil {
        t.Error("FromSiemensPerInch() with +Inf value should return error")
    }

    _, err = factory.FromSiemensPerInch(math.Inf(-1))
    if err == nil {
        t.Error("FromSiemensPerInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSiemensPerInch(0)
    if err != nil {
        t.Errorf("FromSiemensPerInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricConductivitySiemensPerInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSiemensPerInch() with zero value = %v, want 0", converted)
    }
}
// Test FromSiemensPerFoot function
func TestElectricConductivityFactory_FromSiemensPerFoot(t *testing.T) {
    factory := units.ElectricConductivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSiemensPerFoot(100)
    if err != nil {
        t.Errorf("FromSiemensPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductivitySiemensPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSiemensPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSiemensPerFoot(math.NaN())
    if err == nil {
        t.Error("FromSiemensPerFoot() with NaN value should return error")
    }

    _, err = factory.FromSiemensPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromSiemensPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromSiemensPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromSiemensPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSiemensPerFoot(0)
    if err != nil {
        t.Errorf("FromSiemensPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricConductivitySiemensPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSiemensPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromSiemensPerCentimeter function
func TestElectricConductivityFactory_FromSiemensPerCentimeter(t *testing.T) {
    factory := units.ElectricConductivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSiemensPerCentimeter(100)
    if err != nil {
        t.Errorf("FromSiemensPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductivitySiemensPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSiemensPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSiemensPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromSiemensPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromSiemensPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromSiemensPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromSiemensPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromSiemensPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSiemensPerCentimeter(0)
    if err != nil {
        t.Errorf("FromSiemensPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricConductivitySiemensPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSiemensPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrosiemensPerCentimeter function
func TestElectricConductivityFactory_FromMicrosiemensPerCentimeter(t *testing.T) {
    factory := units.ElectricConductivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrosiemensPerCentimeter(100)
    if err != nil {
        t.Errorf("FromMicrosiemensPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductivityMicrosiemensPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrosiemensPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrosiemensPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMicrosiemensPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMicrosiemensPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrosiemensPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMicrosiemensPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrosiemensPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrosiemensPerCentimeter(0)
    if err != nil {
        t.Errorf("FromMicrosiemensPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricConductivityMicrosiemensPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrosiemensPerCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMillisiemensPerCentimeter function
func TestElectricConductivityFactory_FromMillisiemensPerCentimeter(t *testing.T) {
    factory := units.ElectricConductivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillisiemensPerCentimeter(100)
    if err != nil {
        t.Errorf("FromMillisiemensPerCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricConductivityMillisiemensPerCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillisiemensPerCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillisiemensPerCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMillisiemensPerCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMillisiemensPerCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMillisiemensPerCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMillisiemensPerCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMillisiemensPerCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillisiemensPerCentimeter(0)
    if err != nil {
        t.Errorf("FromMillisiemensPerCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricConductivityMillisiemensPerCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillisiemensPerCentimeter() with zero value = %v, want 0", converted)
    }
}

func TestElectricConductivityToString(t *testing.T) {
	factory := units.ElectricConductivityFactory{}
	a, err := factory.CreateElectricConductivity(45, units.ElectricConductivitySiemensPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricConductivitySiemensPerMeter, 2)
	expected := "45.00 " + units.GetElectricConductivityAbbreviation(units.ElectricConductivitySiemensPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricConductivitySiemensPerMeter, -1)
	expected = "45 " + units.GetElectricConductivityAbbreviation(units.ElectricConductivitySiemensPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricConductivity_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricConductivityFactory{}
	a1, _ := factory.CreateElectricConductivity(60, units.ElectricConductivitySiemensPerMeter)
	a2, _ := factory.CreateElectricConductivity(60, units.ElectricConductivitySiemensPerMeter)
	a3, _ := factory.CreateElectricConductivity(120, units.ElectricConductivitySiemensPerMeter)

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

func TestElectricConductivity_Arithmetic(t *testing.T) {
	factory := units.ElectricConductivityFactory{}
	a1, _ := factory.CreateElectricConductivity(30, units.ElectricConductivitySiemensPerMeter)
	a2, _ := factory.CreateElectricConductivity(45, units.ElectricConductivitySiemensPerMeter)

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


func TestGetElectricConductivityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricConductivityUnits
        want string
    }{
        {
            name: "SiemensPerMeter abbreviation",
            unit: units.ElectricConductivitySiemensPerMeter,
            want: "S/m",
        },
        {
            name: "SiemensPerInch abbreviation",
            unit: units.ElectricConductivitySiemensPerInch,
            want: "S/in",
        },
        {
            name: "SiemensPerFoot abbreviation",
            unit: units.ElectricConductivitySiemensPerFoot,
            want: "S/ft",
        },
        {
            name: "SiemensPerCentimeter abbreviation",
            unit: units.ElectricConductivitySiemensPerCentimeter,
            want: "S/cm",
        },
        {
            name: "MicrosiemensPerCentimeter abbreviation",
            unit: units.ElectricConductivityMicrosiemensPerCentimeter,
            want: "μS/cm",
        },
        {
            name: "MillisiemensPerCentimeter abbreviation",
            unit: units.ElectricConductivityMillisiemensPerCentimeter,
            want: "mS/cm",
        },
        {
            name: "invalid unit",
            unit: units.ElectricConductivityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricConductivityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricConductivityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricConductivity_String(t *testing.T) {
    factory := units.ElectricConductivityFactory{}
    
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
            unit, err := factory.CreateElectricConductivity(tt.value, units.ElectricConductivitySiemensPerMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricConductivity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestElectricConductivity_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.ElectricConductivityFactory{}

	_, err := uf.CreateElectricConductivity(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}