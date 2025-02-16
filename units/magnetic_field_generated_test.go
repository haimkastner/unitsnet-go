// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMagneticFieldDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Tesla"}`
	
	factory := units.MagneticFieldDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MagneticFieldTesla {
		t.Errorf("expected unit %v, got %v", units.MagneticFieldTesla, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Tesla"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMagneticFieldDto_ToJSON(t *testing.T) {
	dto := units.MagneticFieldDto{
		Value: 45,
		Unit:  units.MagneticFieldTesla,
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
	if result["unit"].(string) != string(units.MagneticFieldTesla) {
		t.Errorf("expected unit %s, got %v", units.MagneticFieldTesla, result["unit"])
	}
}

func TestNewMagneticField_InvalidValue(t *testing.T) {
	factory := units.MagneticFieldFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMagneticField(math.NaN(), units.MagneticFieldTesla)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMagneticField(math.Inf(1), units.MagneticFieldTesla)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMagneticFieldConversions(t *testing.T) {
	factory := units.MagneticFieldFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMagneticField(180, units.MagneticFieldTesla)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Teslas.
		// No expected conversion value provided for Teslas, verifying result is not NaN.
		result := a.Teslas()
		cacheResult := a.Teslas()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Teslas returned NaN")
		}
	}
	{
		// Test conversion to Gausses.
		// No expected conversion value provided for Gausses, verifying result is not NaN.
		result := a.Gausses()
		cacheResult := a.Gausses()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gausses returned NaN")
		}
	}
	{
		// Test conversion to Nanoteslas.
		// No expected conversion value provided for Nanoteslas, verifying result is not NaN.
		result := a.Nanoteslas()
		cacheResult := a.Nanoteslas()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanoteslas returned NaN")
		}
	}
	{
		// Test conversion to Microteslas.
		// No expected conversion value provided for Microteslas, verifying result is not NaN.
		result := a.Microteslas()
		cacheResult := a.Microteslas()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microteslas returned NaN")
		}
	}
	{
		// Test conversion to Milliteslas.
		// No expected conversion value provided for Milliteslas, verifying result is not NaN.
		result := a.Milliteslas()
		cacheResult := a.Milliteslas()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Milliteslas returned NaN")
		}
	}
	{
		// Test conversion to Milligausses.
		// No expected conversion value provided for Milligausses, verifying result is not NaN.
		result := a.Milligausses()
		cacheResult := a.Milligausses()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Milligausses returned NaN")
		}
	}
}

func TestMagneticField_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MagneticFieldFactory{}
	a, err := factory.CreateMagneticField(90, units.MagneticFieldTesla)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MagneticFieldTesla {
		t.Errorf("expected default unit Tesla, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MagneticFieldTesla
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MagneticFieldDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MagneticFieldTesla {
		t.Errorf("expected unit Tesla, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMagneticFieldFactory_FromDto(t *testing.T) {
    factory := units.MagneticFieldFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MagneticFieldDto{
        Value: 100,
        Unit:  units.MagneticFieldTesla,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MagneticFieldDto{
        Value: math.NaN(),
        Unit:  units.MagneticFieldTesla,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Tesla conversion
    teslasDto := units.MagneticFieldDto{
        Value: 100,
        Unit:  units.MagneticFieldTesla,
    }
    
    var teslasResult *units.MagneticField
    teslasResult, err = factory.FromDto(teslasDto)
    if err != nil {
        t.Errorf("FromDto() with Tesla returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teslasResult.Convert(units.MagneticFieldTesla)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Tesla = %v, want %v", converted, 100)
    }
    // Test Gauss conversion
    gaussesDto := units.MagneticFieldDto{
        Value: 100,
        Unit:  units.MagneticFieldGauss,
    }
    
    var gaussesResult *units.MagneticField
    gaussesResult, err = factory.FromDto(gaussesDto)
    if err != nil {
        t.Errorf("FromDto() with Gauss returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gaussesResult.Convert(units.MagneticFieldGauss)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gauss = %v, want %v", converted, 100)
    }
    // Test Nanotesla conversion
    nanoteslasDto := units.MagneticFieldDto{
        Value: 100,
        Unit:  units.MagneticFieldNanotesla,
    }
    
    var nanoteslasResult *units.MagneticField
    nanoteslasResult, err = factory.FromDto(nanoteslasDto)
    if err != nil {
        t.Errorf("FromDto() with Nanotesla returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoteslasResult.Convert(units.MagneticFieldNanotesla)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanotesla = %v, want %v", converted, 100)
    }
    // Test Microtesla conversion
    microteslasDto := units.MagneticFieldDto{
        Value: 100,
        Unit:  units.MagneticFieldMicrotesla,
    }
    
    var microteslasResult *units.MagneticField
    microteslasResult, err = factory.FromDto(microteslasDto)
    if err != nil {
        t.Errorf("FromDto() with Microtesla returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microteslasResult.Convert(units.MagneticFieldMicrotesla)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microtesla = %v, want %v", converted, 100)
    }
    // Test Millitesla conversion
    milliteslasDto := units.MagneticFieldDto{
        Value: 100,
        Unit:  units.MagneticFieldMillitesla,
    }
    
    var milliteslasResult *units.MagneticField
    milliteslasResult, err = factory.FromDto(milliteslasDto)
    if err != nil {
        t.Errorf("FromDto() with Millitesla returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliteslasResult.Convert(units.MagneticFieldMillitesla)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millitesla = %v, want %v", converted, 100)
    }
    // Test Milligauss conversion
    milligaussesDto := units.MagneticFieldDto{
        Value: 100,
        Unit:  units.MagneticFieldMilligauss,
    }
    
    var milligaussesResult *units.MagneticField
    milligaussesResult, err = factory.FromDto(milligaussesDto)
    if err != nil {
        t.Errorf("FromDto() with Milligauss returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligaussesResult.Convert(units.MagneticFieldMilligauss)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milligauss = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MagneticFieldDto{
        Value: 0,
        Unit:  units.MagneticFieldTesla,
    }
    
    var zeroResult *units.MagneticField
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMagneticFieldFactory_FromDtoJSON(t *testing.T) {
    factory := units.MagneticFieldFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Tesla"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Tesla"}`)
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
    nanJSON, _ := json.Marshal(units.MagneticFieldDto{
        Value: nanValue,
        Unit:  units.MagneticFieldTesla,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Tesla unit
    teslasJSON := []byte(`{"value": 100, "unit": "Tesla"}`)
    teslasResult, err := factory.FromDtoJSON(teslasJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Tesla unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teslasResult.Convert(units.MagneticFieldTesla)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Tesla = %v, want %v", converted, 100)
    }
    // Test JSON with Gauss unit
    gaussesJSON := []byte(`{"value": 100, "unit": "Gauss"}`)
    gaussesResult, err := factory.FromDtoJSON(gaussesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gauss unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gaussesResult.Convert(units.MagneticFieldGauss)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gauss = %v, want %v", converted, 100)
    }
    // Test JSON with Nanotesla unit
    nanoteslasJSON := []byte(`{"value": 100, "unit": "Nanotesla"}`)
    nanoteslasResult, err := factory.FromDtoJSON(nanoteslasJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanotesla unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoteslasResult.Convert(units.MagneticFieldNanotesla)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanotesla = %v, want %v", converted, 100)
    }
    // Test JSON with Microtesla unit
    microteslasJSON := []byte(`{"value": 100, "unit": "Microtesla"}`)
    microteslasResult, err := factory.FromDtoJSON(microteslasJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microtesla unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microteslasResult.Convert(units.MagneticFieldMicrotesla)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microtesla = %v, want %v", converted, 100)
    }
    // Test JSON with Millitesla unit
    milliteslasJSON := []byte(`{"value": 100, "unit": "Millitesla"}`)
    milliteslasResult, err := factory.FromDtoJSON(milliteslasJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millitesla unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliteslasResult.Convert(units.MagneticFieldMillitesla)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millitesla = %v, want %v", converted, 100)
    }
    // Test JSON with Milligauss unit
    milligaussesJSON := []byte(`{"value": 100, "unit": "Milligauss"}`)
    milligaussesResult, err := factory.FromDtoJSON(milligaussesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Milligauss unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligaussesResult.Convert(units.MagneticFieldMilligauss)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milligauss = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Tesla"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromTeslas function
func TestMagneticFieldFactory_FromTeslas(t *testing.T) {
    factory := units.MagneticFieldFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTeslas(100)
    if err != nil {
        t.Errorf("FromTeslas() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MagneticFieldTesla)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTeslas() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTeslas(math.NaN())
    if err == nil {
        t.Error("FromTeslas() with NaN value should return error")
    }

    _, err = factory.FromTeslas(math.Inf(1))
    if err == nil {
        t.Error("FromTeslas() with +Inf value should return error")
    }

    _, err = factory.FromTeslas(math.Inf(-1))
    if err == nil {
        t.Error("FromTeslas() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTeslas(0)
    if err != nil {
        t.Errorf("FromTeslas() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MagneticFieldTesla)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTeslas() with zero value = %v, want 0", converted)
    }
}
// Test FromGausses function
func TestMagneticFieldFactory_FromGausses(t *testing.T) {
    factory := units.MagneticFieldFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGausses(100)
    if err != nil {
        t.Errorf("FromGausses() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MagneticFieldGauss)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGausses() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGausses(math.NaN())
    if err == nil {
        t.Error("FromGausses() with NaN value should return error")
    }

    _, err = factory.FromGausses(math.Inf(1))
    if err == nil {
        t.Error("FromGausses() with +Inf value should return error")
    }

    _, err = factory.FromGausses(math.Inf(-1))
    if err == nil {
        t.Error("FromGausses() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGausses(0)
    if err != nil {
        t.Errorf("FromGausses() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MagneticFieldGauss)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGausses() with zero value = %v, want 0", converted)
    }
}
// Test FromNanoteslas function
func TestMagneticFieldFactory_FromNanoteslas(t *testing.T) {
    factory := units.MagneticFieldFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanoteslas(100)
    if err != nil {
        t.Errorf("FromNanoteslas() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MagneticFieldNanotesla)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanoteslas() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanoteslas(math.NaN())
    if err == nil {
        t.Error("FromNanoteslas() with NaN value should return error")
    }

    _, err = factory.FromNanoteslas(math.Inf(1))
    if err == nil {
        t.Error("FromNanoteslas() with +Inf value should return error")
    }

    _, err = factory.FromNanoteslas(math.Inf(-1))
    if err == nil {
        t.Error("FromNanoteslas() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanoteslas(0)
    if err != nil {
        t.Errorf("FromNanoteslas() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MagneticFieldNanotesla)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanoteslas() with zero value = %v, want 0", converted)
    }
}
// Test FromMicroteslas function
func TestMagneticFieldFactory_FromMicroteslas(t *testing.T) {
    factory := units.MagneticFieldFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicroteslas(100)
    if err != nil {
        t.Errorf("FromMicroteslas() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MagneticFieldMicrotesla)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicroteslas() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicroteslas(math.NaN())
    if err == nil {
        t.Error("FromMicroteslas() with NaN value should return error")
    }

    _, err = factory.FromMicroteslas(math.Inf(1))
    if err == nil {
        t.Error("FromMicroteslas() with +Inf value should return error")
    }

    _, err = factory.FromMicroteslas(math.Inf(-1))
    if err == nil {
        t.Error("FromMicroteslas() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicroteslas(0)
    if err != nil {
        t.Errorf("FromMicroteslas() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MagneticFieldMicrotesla)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicroteslas() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliteslas function
func TestMagneticFieldFactory_FromMilliteslas(t *testing.T) {
    factory := units.MagneticFieldFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliteslas(100)
    if err != nil {
        t.Errorf("FromMilliteslas() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MagneticFieldMillitesla)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliteslas() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliteslas(math.NaN())
    if err == nil {
        t.Error("FromMilliteslas() with NaN value should return error")
    }

    _, err = factory.FromMilliteslas(math.Inf(1))
    if err == nil {
        t.Error("FromMilliteslas() with +Inf value should return error")
    }

    _, err = factory.FromMilliteslas(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliteslas() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliteslas(0)
    if err != nil {
        t.Errorf("FromMilliteslas() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MagneticFieldMillitesla)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliteslas() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligausses function
func TestMagneticFieldFactory_FromMilligausses(t *testing.T) {
    factory := units.MagneticFieldFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligausses(100)
    if err != nil {
        t.Errorf("FromMilligausses() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MagneticFieldMilligauss)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligausses() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligausses(math.NaN())
    if err == nil {
        t.Error("FromMilligausses() with NaN value should return error")
    }

    _, err = factory.FromMilligausses(math.Inf(1))
    if err == nil {
        t.Error("FromMilligausses() with +Inf value should return error")
    }

    _, err = factory.FromMilligausses(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligausses() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligausses(0)
    if err != nil {
        t.Errorf("FromMilligausses() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MagneticFieldMilligauss)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligausses() with zero value = %v, want 0", converted)
    }
}

func TestMagneticFieldToString(t *testing.T) {
	factory := units.MagneticFieldFactory{}
	a, err := factory.CreateMagneticField(45, units.MagneticFieldTesla)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MagneticFieldTesla, 2)
	expected := "45.00 " + units.GetMagneticFieldAbbreviation(units.MagneticFieldTesla)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MagneticFieldTesla, -1)
	expected = "45 " + units.GetMagneticFieldAbbreviation(units.MagneticFieldTesla)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMagneticField_EqualityAndComparison(t *testing.T) {
	factory := units.MagneticFieldFactory{}
	a1, _ := factory.CreateMagneticField(60, units.MagneticFieldTesla)
	a2, _ := factory.CreateMagneticField(60, units.MagneticFieldTesla)
	a3, _ := factory.CreateMagneticField(120, units.MagneticFieldTesla)

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

func TestMagneticField_Arithmetic(t *testing.T) {
	factory := units.MagneticFieldFactory{}
	a1, _ := factory.CreateMagneticField(30, units.MagneticFieldTesla)
	a2, _ := factory.CreateMagneticField(45, units.MagneticFieldTesla)

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


func TestGetMagneticFieldAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.MagneticFieldUnits
        want string
    }{
        {
            name: "Tesla abbreviation",
            unit: units.MagneticFieldTesla,
            want: "T",
        },
        {
            name: "Gauss abbreviation",
            unit: units.MagneticFieldGauss,
            want: "G",
        },
        {
            name: "Nanotesla abbreviation",
            unit: units.MagneticFieldNanotesla,
            want: "nT",
        },
        {
            name: "Microtesla abbreviation",
            unit: units.MagneticFieldMicrotesla,
            want: "μT",
        },
        {
            name: "Millitesla abbreviation",
            unit: units.MagneticFieldMillitesla,
            want: "mT",
        },
        {
            name: "Milligauss abbreviation",
            unit: units.MagneticFieldMilligauss,
            want: "mG",
        },
        {
            name: "invalid unit",
            unit: units.MagneticFieldUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetMagneticFieldAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetMagneticFieldAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestMagneticField_String(t *testing.T) {
    factory := units.MagneticFieldFactory{}
    
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
            unit, err := factory.CreateMagneticField(tt.value, units.MagneticFieldTesla)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("MagneticField.String() = %v, want %v", got, tt.want)
            }
        })
    }
}