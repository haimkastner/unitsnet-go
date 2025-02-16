// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestCompressibilityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "InversePascal"}`
	
	factory := units.CompressibilityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.CompressibilityInversePascal {
		t.Errorf("expected unit %v, got %v", units.CompressibilityInversePascal, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "InversePascal"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestCompressibilityDto_ToJSON(t *testing.T) {
	dto := units.CompressibilityDto{
		Value: 45,
		Unit:  units.CompressibilityInversePascal,
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
	if result["unit"].(string) != string(units.CompressibilityInversePascal) {
		t.Errorf("expected unit %s, got %v", units.CompressibilityInversePascal, result["unit"])
	}
}

func TestNewCompressibility_InvalidValue(t *testing.T) {
	factory := units.CompressibilityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateCompressibility(math.NaN(), units.CompressibilityInversePascal)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateCompressibility(math.Inf(1), units.CompressibilityInversePascal)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestCompressibilityConversions(t *testing.T) {
	factory := units.CompressibilityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateCompressibility(180, units.CompressibilityInversePascal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to InversePascals.
		// No expected conversion value provided for InversePascals, verifying result is not NaN.
		result := a.InversePascals()
		cacheResult := a.InversePascals()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InversePascals returned NaN")
		}
	}
	{
		// Test conversion to InverseKilopascals.
		// No expected conversion value provided for InverseKilopascals, verifying result is not NaN.
		result := a.InverseKilopascals()
		cacheResult := a.InverseKilopascals()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseKilopascals returned NaN")
		}
	}
	{
		// Test conversion to InverseMegapascals.
		// No expected conversion value provided for InverseMegapascals, verifying result is not NaN.
		result := a.InverseMegapascals()
		cacheResult := a.InverseMegapascals()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseMegapascals returned NaN")
		}
	}
	{
		// Test conversion to InverseAtmospheres.
		// No expected conversion value provided for InverseAtmospheres, verifying result is not NaN.
		result := a.InverseAtmospheres()
		cacheResult := a.InverseAtmospheres()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseAtmospheres returned NaN")
		}
	}
	{
		// Test conversion to InverseMillibars.
		// No expected conversion value provided for InverseMillibars, verifying result is not NaN.
		result := a.InverseMillibars()
		cacheResult := a.InverseMillibars()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseMillibars returned NaN")
		}
	}
	{
		// Test conversion to InverseBars.
		// No expected conversion value provided for InverseBars, verifying result is not NaN.
		result := a.InverseBars()
		cacheResult := a.InverseBars()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InverseBars returned NaN")
		}
	}
	{
		// Test conversion to InversePoundsForcePerSquareInch.
		// No expected conversion value provided for InversePoundsForcePerSquareInch, verifying result is not NaN.
		result := a.InversePoundsForcePerSquareInch()
		cacheResult := a.InversePoundsForcePerSquareInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InversePoundsForcePerSquareInch returned NaN")
		}
	}
}

func TestCompressibility_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.CompressibilityFactory{}
	a, err := factory.CreateCompressibility(90, units.CompressibilityInversePascal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.CompressibilityInversePascal {
		t.Errorf("expected default unit InversePascal, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.CompressibilityInversePascal
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.CompressibilityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.CompressibilityInversePascal {
		t.Errorf("expected unit InversePascal, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestCompressibilityFactory_FromDto(t *testing.T) {
    factory := units.CompressibilityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.CompressibilityDto{
        Value: 100,
        Unit:  units.CompressibilityInversePascal,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.CompressibilityDto{
        Value: math.NaN(),
        Unit:  units.CompressibilityInversePascal,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test InversePascal conversion
    inverse_pascalsDto := units.CompressibilityDto{
        Value: 100,
        Unit:  units.CompressibilityInversePascal,
    }
    
    var inverse_pascalsResult *units.Compressibility
    inverse_pascalsResult, err = factory.FromDto(inverse_pascalsDto)
    if err != nil {
        t.Errorf("FromDto() with InversePascal returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_pascalsResult.Convert(units.CompressibilityInversePascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InversePascal = %v, want %v", converted, 100)
    }
    // Test InverseKilopascal conversion
    inverse_kilopascalsDto := units.CompressibilityDto{
        Value: 100,
        Unit:  units.CompressibilityInverseKilopascal,
    }
    
    var inverse_kilopascalsResult *units.Compressibility
    inverse_kilopascalsResult, err = factory.FromDto(inverse_kilopascalsDto)
    if err != nil {
        t.Errorf("FromDto() with InverseKilopascal returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_kilopascalsResult.Convert(units.CompressibilityInverseKilopascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseKilopascal = %v, want %v", converted, 100)
    }
    // Test InverseMegapascal conversion
    inverse_megapascalsDto := units.CompressibilityDto{
        Value: 100,
        Unit:  units.CompressibilityInverseMegapascal,
    }
    
    var inverse_megapascalsResult *units.Compressibility
    inverse_megapascalsResult, err = factory.FromDto(inverse_megapascalsDto)
    if err != nil {
        t.Errorf("FromDto() with InverseMegapascal returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_megapascalsResult.Convert(units.CompressibilityInverseMegapascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMegapascal = %v, want %v", converted, 100)
    }
    // Test InverseAtmosphere conversion
    inverse_atmospheresDto := units.CompressibilityDto{
        Value: 100,
        Unit:  units.CompressibilityInverseAtmosphere,
    }
    
    var inverse_atmospheresResult *units.Compressibility
    inverse_atmospheresResult, err = factory.FromDto(inverse_atmospheresDto)
    if err != nil {
        t.Errorf("FromDto() with InverseAtmosphere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_atmospheresResult.Convert(units.CompressibilityInverseAtmosphere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseAtmosphere = %v, want %v", converted, 100)
    }
    // Test InverseMillibar conversion
    inverse_millibarsDto := units.CompressibilityDto{
        Value: 100,
        Unit:  units.CompressibilityInverseMillibar,
    }
    
    var inverse_millibarsResult *units.Compressibility
    inverse_millibarsResult, err = factory.FromDto(inverse_millibarsDto)
    if err != nil {
        t.Errorf("FromDto() with InverseMillibar returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_millibarsResult.Convert(units.CompressibilityInverseMillibar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMillibar = %v, want %v", converted, 100)
    }
    // Test InverseBar conversion
    inverse_barsDto := units.CompressibilityDto{
        Value: 100,
        Unit:  units.CompressibilityInverseBar,
    }
    
    var inverse_barsResult *units.Compressibility
    inverse_barsResult, err = factory.FromDto(inverse_barsDto)
    if err != nil {
        t.Errorf("FromDto() with InverseBar returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_barsResult.Convert(units.CompressibilityInverseBar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseBar = %v, want %v", converted, 100)
    }
    // Test InversePoundForcePerSquareInch conversion
    inverse_pounds_force_per_square_inchDto := units.CompressibilityDto{
        Value: 100,
        Unit:  units.CompressibilityInversePoundForcePerSquareInch,
    }
    
    var inverse_pounds_force_per_square_inchResult *units.Compressibility
    inverse_pounds_force_per_square_inchResult, err = factory.FromDto(inverse_pounds_force_per_square_inchDto)
    if err != nil {
        t.Errorf("FromDto() with InversePoundForcePerSquareInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_pounds_force_per_square_inchResult.Convert(units.CompressibilityInversePoundForcePerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InversePoundForcePerSquareInch = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.CompressibilityDto{
        Value: 0,
        Unit:  units.CompressibilityInversePascal,
    }
    
    var zeroResult *units.Compressibility
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestCompressibilityFactory_FromDtoJSON(t *testing.T) {
    factory := units.CompressibilityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "InversePascal"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "InversePascal"}`)
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
    nanJSON, _ := json.Marshal(units.CompressibilityDto{
        Value: nanValue,
        Unit:  units.CompressibilityInversePascal,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with InversePascal unit
    inverse_pascalsJSON := []byte(`{"value": 100, "unit": "InversePascal"}`)
    inverse_pascalsResult, err := factory.FromDtoJSON(inverse_pascalsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InversePascal unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_pascalsResult.Convert(units.CompressibilityInversePascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InversePascal = %v, want %v", converted, 100)
    }
    // Test JSON with InverseKilopascal unit
    inverse_kilopascalsJSON := []byte(`{"value": 100, "unit": "InverseKilopascal"}`)
    inverse_kilopascalsResult, err := factory.FromDtoJSON(inverse_kilopascalsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseKilopascal unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_kilopascalsResult.Convert(units.CompressibilityInverseKilopascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseKilopascal = %v, want %v", converted, 100)
    }
    // Test JSON with InverseMegapascal unit
    inverse_megapascalsJSON := []byte(`{"value": 100, "unit": "InverseMegapascal"}`)
    inverse_megapascalsResult, err := factory.FromDtoJSON(inverse_megapascalsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseMegapascal unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_megapascalsResult.Convert(units.CompressibilityInverseMegapascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMegapascal = %v, want %v", converted, 100)
    }
    // Test JSON with InverseAtmosphere unit
    inverse_atmospheresJSON := []byte(`{"value": 100, "unit": "InverseAtmosphere"}`)
    inverse_atmospheresResult, err := factory.FromDtoJSON(inverse_atmospheresJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseAtmosphere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_atmospheresResult.Convert(units.CompressibilityInverseAtmosphere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseAtmosphere = %v, want %v", converted, 100)
    }
    // Test JSON with InverseMillibar unit
    inverse_millibarsJSON := []byte(`{"value": 100, "unit": "InverseMillibar"}`)
    inverse_millibarsResult, err := factory.FromDtoJSON(inverse_millibarsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseMillibar unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_millibarsResult.Convert(units.CompressibilityInverseMillibar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseMillibar = %v, want %v", converted, 100)
    }
    // Test JSON with InverseBar unit
    inverse_barsJSON := []byte(`{"value": 100, "unit": "InverseBar"}`)
    inverse_barsResult, err := factory.FromDtoJSON(inverse_barsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InverseBar unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_barsResult.Convert(units.CompressibilityInverseBar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InverseBar = %v, want %v", converted, 100)
    }
    // Test JSON with InversePoundForcePerSquareInch unit
    inverse_pounds_force_per_square_inchJSON := []byte(`{"value": 100, "unit": "InversePoundForcePerSquareInch"}`)
    inverse_pounds_force_per_square_inchResult, err := factory.FromDtoJSON(inverse_pounds_force_per_square_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InversePoundForcePerSquareInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inverse_pounds_force_per_square_inchResult.Convert(units.CompressibilityInversePoundForcePerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InversePoundForcePerSquareInch = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "InversePascal"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromInversePascals function
func TestCompressibilityFactory_FromInversePascals(t *testing.T) {
    factory := units.CompressibilityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInversePascals(100)
    if err != nil {
        t.Errorf("FromInversePascals() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CompressibilityInversePascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInversePascals() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInversePascals(math.NaN())
    if err == nil {
        t.Error("FromInversePascals() with NaN value should return error")
    }

    _, err = factory.FromInversePascals(math.Inf(1))
    if err == nil {
        t.Error("FromInversePascals() with +Inf value should return error")
    }

    _, err = factory.FromInversePascals(math.Inf(-1))
    if err == nil {
        t.Error("FromInversePascals() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInversePascals(0)
    if err != nil {
        t.Errorf("FromInversePascals() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CompressibilityInversePascal)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInversePascals() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseKilopascals function
func TestCompressibilityFactory_FromInverseKilopascals(t *testing.T) {
    factory := units.CompressibilityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseKilopascals(100)
    if err != nil {
        t.Errorf("FromInverseKilopascals() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CompressibilityInverseKilopascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseKilopascals() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseKilopascals(math.NaN())
    if err == nil {
        t.Error("FromInverseKilopascals() with NaN value should return error")
    }

    _, err = factory.FromInverseKilopascals(math.Inf(1))
    if err == nil {
        t.Error("FromInverseKilopascals() with +Inf value should return error")
    }

    _, err = factory.FromInverseKilopascals(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseKilopascals() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseKilopascals(0)
    if err != nil {
        t.Errorf("FromInverseKilopascals() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CompressibilityInverseKilopascal)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseKilopascals() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseMegapascals function
func TestCompressibilityFactory_FromInverseMegapascals(t *testing.T) {
    factory := units.CompressibilityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseMegapascals(100)
    if err != nil {
        t.Errorf("FromInverseMegapascals() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CompressibilityInverseMegapascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseMegapascals() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseMegapascals(math.NaN())
    if err == nil {
        t.Error("FromInverseMegapascals() with NaN value should return error")
    }

    _, err = factory.FromInverseMegapascals(math.Inf(1))
    if err == nil {
        t.Error("FromInverseMegapascals() with +Inf value should return error")
    }

    _, err = factory.FromInverseMegapascals(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseMegapascals() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseMegapascals(0)
    if err != nil {
        t.Errorf("FromInverseMegapascals() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CompressibilityInverseMegapascal)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseMegapascals() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseAtmospheres function
func TestCompressibilityFactory_FromInverseAtmospheres(t *testing.T) {
    factory := units.CompressibilityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseAtmospheres(100)
    if err != nil {
        t.Errorf("FromInverseAtmospheres() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CompressibilityInverseAtmosphere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseAtmospheres() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseAtmospheres(math.NaN())
    if err == nil {
        t.Error("FromInverseAtmospheres() with NaN value should return error")
    }

    _, err = factory.FromInverseAtmospheres(math.Inf(1))
    if err == nil {
        t.Error("FromInverseAtmospheres() with +Inf value should return error")
    }

    _, err = factory.FromInverseAtmospheres(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseAtmospheres() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseAtmospheres(0)
    if err != nil {
        t.Errorf("FromInverseAtmospheres() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CompressibilityInverseAtmosphere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseAtmospheres() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseMillibars function
func TestCompressibilityFactory_FromInverseMillibars(t *testing.T) {
    factory := units.CompressibilityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseMillibars(100)
    if err != nil {
        t.Errorf("FromInverseMillibars() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CompressibilityInverseMillibar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseMillibars() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseMillibars(math.NaN())
    if err == nil {
        t.Error("FromInverseMillibars() with NaN value should return error")
    }

    _, err = factory.FromInverseMillibars(math.Inf(1))
    if err == nil {
        t.Error("FromInverseMillibars() with +Inf value should return error")
    }

    _, err = factory.FromInverseMillibars(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseMillibars() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseMillibars(0)
    if err != nil {
        t.Errorf("FromInverseMillibars() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CompressibilityInverseMillibar)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseMillibars() with zero value = %v, want 0", converted)
    }
}
// Test FromInverseBars function
func TestCompressibilityFactory_FromInverseBars(t *testing.T) {
    factory := units.CompressibilityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInverseBars(100)
    if err != nil {
        t.Errorf("FromInverseBars() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CompressibilityInverseBar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInverseBars() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInverseBars(math.NaN())
    if err == nil {
        t.Error("FromInverseBars() with NaN value should return error")
    }

    _, err = factory.FromInverseBars(math.Inf(1))
    if err == nil {
        t.Error("FromInverseBars() with +Inf value should return error")
    }

    _, err = factory.FromInverseBars(math.Inf(-1))
    if err == nil {
        t.Error("FromInverseBars() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInverseBars(0)
    if err != nil {
        t.Errorf("FromInverseBars() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CompressibilityInverseBar)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInverseBars() with zero value = %v, want 0", converted)
    }
}
// Test FromInversePoundsForcePerSquareInch function
func TestCompressibilityFactory_FromInversePoundsForcePerSquareInch(t *testing.T) {
    factory := units.CompressibilityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInversePoundsForcePerSquareInch(100)
    if err != nil {
        t.Errorf("FromInversePoundsForcePerSquareInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.CompressibilityInversePoundForcePerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInversePoundsForcePerSquareInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInversePoundsForcePerSquareInch(math.NaN())
    if err == nil {
        t.Error("FromInversePoundsForcePerSquareInch() with NaN value should return error")
    }

    _, err = factory.FromInversePoundsForcePerSquareInch(math.Inf(1))
    if err == nil {
        t.Error("FromInversePoundsForcePerSquareInch() with +Inf value should return error")
    }

    _, err = factory.FromInversePoundsForcePerSquareInch(math.Inf(-1))
    if err == nil {
        t.Error("FromInversePoundsForcePerSquareInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInversePoundsForcePerSquareInch(0)
    if err != nil {
        t.Errorf("FromInversePoundsForcePerSquareInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.CompressibilityInversePoundForcePerSquareInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInversePoundsForcePerSquareInch() with zero value = %v, want 0", converted)
    }
}

func TestCompressibilityToString(t *testing.T) {
	factory := units.CompressibilityFactory{}
	a, err := factory.CreateCompressibility(45, units.CompressibilityInversePascal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.CompressibilityInversePascal, 2)
	expected := "45.00 " + units.GetCompressibilityAbbreviation(units.CompressibilityInversePascal)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.CompressibilityInversePascal, -1)
	expected = "45 " + units.GetCompressibilityAbbreviation(units.CompressibilityInversePascal)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestCompressibility_EqualityAndComparison(t *testing.T) {
	factory := units.CompressibilityFactory{}
	a1, _ := factory.CreateCompressibility(60, units.CompressibilityInversePascal)
	a2, _ := factory.CreateCompressibility(60, units.CompressibilityInversePascal)
	a3, _ := factory.CreateCompressibility(120, units.CompressibilityInversePascal)

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

func TestCompressibility_Arithmetic(t *testing.T) {
	factory := units.CompressibilityFactory{}
	a1, _ := factory.CreateCompressibility(30, units.CompressibilityInversePascal)
	a2, _ := factory.CreateCompressibility(45, units.CompressibilityInversePascal)

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


func TestGetCompressibilityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.CompressibilityUnits
        want string
    }{
        {
            name: "InversePascal abbreviation",
            unit: units.CompressibilityInversePascal,
            want: "Pa⁻¹",
        },
        {
            name: "InverseKilopascal abbreviation",
            unit: units.CompressibilityInverseKilopascal,
            want: "kPa⁻¹",
        },
        {
            name: "InverseMegapascal abbreviation",
            unit: units.CompressibilityInverseMegapascal,
            want: "MPa⁻¹",
        },
        {
            name: "InverseAtmosphere abbreviation",
            unit: units.CompressibilityInverseAtmosphere,
            want: "atm⁻¹",
        },
        {
            name: "InverseMillibar abbreviation",
            unit: units.CompressibilityInverseMillibar,
            want: "mbar⁻¹",
        },
        {
            name: "InverseBar abbreviation",
            unit: units.CompressibilityInverseBar,
            want: "bar⁻¹",
        },
        {
            name: "InversePoundForcePerSquareInch abbreviation",
            unit: units.CompressibilityInversePoundForcePerSquareInch,
            want: "psi⁻¹",
        },
        {
            name: "invalid unit",
            unit: units.CompressibilityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetCompressibilityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetCompressibilityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestCompressibility_String(t *testing.T) {
    factory := units.CompressibilityFactory{}
    
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
            unit, err := factory.CreateCompressibility(tt.value, units.CompressibilityInversePascal)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Compressibility.String() = %v, want %v", got, tt.want)
            }
        })
    }
}