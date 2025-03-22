// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRadiationExposureDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CoulombPerKilogram"}`
	
	factory := units.RadiationExposureDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RadiationExposureCoulombPerKilogram {
		t.Errorf("expected unit %v, got %v", units.RadiationExposureCoulombPerKilogram, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CoulombPerKilogram"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRadiationExposureDto_ToJSON(t *testing.T) {
	dto := units.RadiationExposureDto{
		Value: 45,
		Unit:  units.RadiationExposureCoulombPerKilogram,
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
	if result["unit"].(string) != string(units.RadiationExposureCoulombPerKilogram) {
		t.Errorf("expected unit %s, got %v", units.RadiationExposureCoulombPerKilogram, result["unit"])
	}
}

func TestNewRadiationExposure_InvalidValue(t *testing.T) {
	factory := units.RadiationExposureFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRadiationExposure(math.NaN(), units.RadiationExposureCoulombPerKilogram)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRadiationExposure(math.Inf(1), units.RadiationExposureCoulombPerKilogram)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRadiationExposureConversions(t *testing.T) {
	factory := units.RadiationExposureFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRadiationExposure(180, units.RadiationExposureCoulombPerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to CoulombsPerKilogram.
		// No expected conversion value provided for CoulombsPerKilogram, verifying result is not NaN.
		result := a.CoulombsPerKilogram()
		cacheResult := a.CoulombsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CoulombsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to Roentgens.
		// No expected conversion value provided for Roentgens, verifying result is not NaN.
		result := a.Roentgens()
		cacheResult := a.Roentgens()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Roentgens returned NaN")
		}
	}
	{
		// Test conversion to PicocoulombsPerKilogram.
		// No expected conversion value provided for PicocoulombsPerKilogram, verifying result is not NaN.
		result := a.PicocoulombsPerKilogram()
		cacheResult := a.PicocoulombsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PicocoulombsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to NanocoulombsPerKilogram.
		// No expected conversion value provided for NanocoulombsPerKilogram, verifying result is not NaN.
		result := a.NanocoulombsPerKilogram()
		cacheResult := a.NanocoulombsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanocoulombsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MicrocoulombsPerKilogram.
		// No expected conversion value provided for MicrocoulombsPerKilogram, verifying result is not NaN.
		result := a.MicrocoulombsPerKilogram()
		cacheResult := a.MicrocoulombsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrocoulombsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MillicoulombsPerKilogram.
		// No expected conversion value provided for MillicoulombsPerKilogram, verifying result is not NaN.
		result := a.MillicoulombsPerKilogram()
		cacheResult := a.MillicoulombsPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillicoulombsPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to Microroentgens.
		// No expected conversion value provided for Microroentgens, verifying result is not NaN.
		result := a.Microroentgens()
		cacheResult := a.Microroentgens()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microroentgens returned NaN")
		}
	}
	{
		// Test conversion to Milliroentgens.
		// No expected conversion value provided for Milliroentgens, verifying result is not NaN.
		result := a.Milliroentgens()
		cacheResult := a.Milliroentgens()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Milliroentgens returned NaN")
		}
	}
}

func TestRadiationExposure_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RadiationExposureFactory{}
	a, err := factory.CreateRadiationExposure(90, units.RadiationExposureCoulombPerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RadiationExposureCoulombPerKilogram {
		t.Errorf("expected default unit CoulombPerKilogram, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RadiationExposureCoulombPerKilogram
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RadiationExposureDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RadiationExposureCoulombPerKilogram {
		t.Errorf("expected unit CoulombPerKilogram, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRadiationExposureFactory_FromDto(t *testing.T) {
    factory := units.RadiationExposureFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.RadiationExposureDto{
        Value: 100,
        Unit:  units.RadiationExposureCoulombPerKilogram,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.RadiationExposureDto{
        Value: math.NaN(),
        Unit:  units.RadiationExposureCoulombPerKilogram,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test CoulombPerKilogram conversion
    coulombs_per_kilogramDto := units.RadiationExposureDto{
        Value: 100,
        Unit:  units.RadiationExposureCoulombPerKilogram,
    }
    
    var coulombs_per_kilogramResult *units.RadiationExposure
    coulombs_per_kilogramResult, err = factory.FromDto(coulombs_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with CoulombPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = coulombs_per_kilogramResult.Convert(units.RadiationExposureCoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CoulombPerKilogram = %v, want %v", converted, 100)
    }
    // Test Roentgen conversion
    roentgensDto := units.RadiationExposureDto{
        Value: 100,
        Unit:  units.RadiationExposureRoentgen,
    }
    
    var roentgensResult *units.RadiationExposure
    roentgensResult, err = factory.FromDto(roentgensDto)
    if err != nil {
        t.Errorf("FromDto() with Roentgen returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = roentgensResult.Convert(units.RadiationExposureRoentgen)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Roentgen = %v, want %v", converted, 100)
    }
    // Test PicocoulombPerKilogram conversion
    picocoulombs_per_kilogramDto := units.RadiationExposureDto{
        Value: 100,
        Unit:  units.RadiationExposurePicocoulombPerKilogram,
    }
    
    var picocoulombs_per_kilogramResult *units.RadiationExposure
    picocoulombs_per_kilogramResult, err = factory.FromDto(picocoulombs_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with PicocoulombPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picocoulombs_per_kilogramResult.Convert(units.RadiationExposurePicocoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicocoulombPerKilogram = %v, want %v", converted, 100)
    }
    // Test NanocoulombPerKilogram conversion
    nanocoulombs_per_kilogramDto := units.RadiationExposureDto{
        Value: 100,
        Unit:  units.RadiationExposureNanocoulombPerKilogram,
    }
    
    var nanocoulombs_per_kilogramResult *units.RadiationExposure
    nanocoulombs_per_kilogramResult, err = factory.FromDto(nanocoulombs_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with NanocoulombPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanocoulombs_per_kilogramResult.Convert(units.RadiationExposureNanocoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanocoulombPerKilogram = %v, want %v", converted, 100)
    }
    // Test MicrocoulombPerKilogram conversion
    microcoulombs_per_kilogramDto := units.RadiationExposureDto{
        Value: 100,
        Unit:  units.RadiationExposureMicrocoulombPerKilogram,
    }
    
    var microcoulombs_per_kilogramResult *units.RadiationExposure
    microcoulombs_per_kilogramResult, err = factory.FromDto(microcoulombs_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with MicrocoulombPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microcoulombs_per_kilogramResult.Convert(units.RadiationExposureMicrocoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrocoulombPerKilogram = %v, want %v", converted, 100)
    }
    // Test MillicoulombPerKilogram conversion
    millicoulombs_per_kilogramDto := units.RadiationExposureDto{
        Value: 100,
        Unit:  units.RadiationExposureMillicoulombPerKilogram,
    }
    
    var millicoulombs_per_kilogramResult *units.RadiationExposure
    millicoulombs_per_kilogramResult, err = factory.FromDto(millicoulombs_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with MillicoulombPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millicoulombs_per_kilogramResult.Convert(units.RadiationExposureMillicoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillicoulombPerKilogram = %v, want %v", converted, 100)
    }
    // Test Microroentgen conversion
    microroentgensDto := units.RadiationExposureDto{
        Value: 100,
        Unit:  units.RadiationExposureMicroroentgen,
    }
    
    var microroentgensResult *units.RadiationExposure
    microroentgensResult, err = factory.FromDto(microroentgensDto)
    if err != nil {
        t.Errorf("FromDto() with Microroentgen returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microroentgensResult.Convert(units.RadiationExposureMicroroentgen)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microroentgen = %v, want %v", converted, 100)
    }
    // Test Milliroentgen conversion
    milliroentgensDto := units.RadiationExposureDto{
        Value: 100,
        Unit:  units.RadiationExposureMilliroentgen,
    }
    
    var milliroentgensResult *units.RadiationExposure
    milliroentgensResult, err = factory.FromDto(milliroentgensDto)
    if err != nil {
        t.Errorf("FromDto() with Milliroentgen returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliroentgensResult.Convert(units.RadiationExposureMilliroentgen)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliroentgen = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.RadiationExposureDto{
        Value: 0,
        Unit:  units.RadiationExposureCoulombPerKilogram,
    }
    
    var zeroResult *units.RadiationExposure
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestRadiationExposureFactory_FromDtoJSON(t *testing.T) {
    factory := units.RadiationExposureFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "CoulombPerKilogram"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "CoulombPerKilogram"}`)
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
    nanJSON, _ := json.Marshal(units.RadiationExposureDto{
        Value: nanValue,
        Unit:  units.RadiationExposureCoulombPerKilogram,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with CoulombPerKilogram unit
    coulombs_per_kilogramJSON := []byte(`{"value": 100, "unit": "CoulombPerKilogram"}`)
    coulombs_per_kilogramResult, err := factory.FromDtoJSON(coulombs_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CoulombPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = coulombs_per_kilogramResult.Convert(units.RadiationExposureCoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CoulombPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with Roentgen unit
    roentgensJSON := []byte(`{"value": 100, "unit": "Roentgen"}`)
    roentgensResult, err := factory.FromDtoJSON(roentgensJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Roentgen unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = roentgensResult.Convert(units.RadiationExposureRoentgen)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Roentgen = %v, want %v", converted, 100)
    }
    // Test JSON with PicocoulombPerKilogram unit
    picocoulombs_per_kilogramJSON := []byte(`{"value": 100, "unit": "PicocoulombPerKilogram"}`)
    picocoulombs_per_kilogramResult, err := factory.FromDtoJSON(picocoulombs_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicocoulombPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picocoulombs_per_kilogramResult.Convert(units.RadiationExposurePicocoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicocoulombPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with NanocoulombPerKilogram unit
    nanocoulombs_per_kilogramJSON := []byte(`{"value": 100, "unit": "NanocoulombPerKilogram"}`)
    nanocoulombs_per_kilogramResult, err := factory.FromDtoJSON(nanocoulombs_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanocoulombPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanocoulombs_per_kilogramResult.Convert(units.RadiationExposureNanocoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanocoulombPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with MicrocoulombPerKilogram unit
    microcoulombs_per_kilogramJSON := []byte(`{"value": 100, "unit": "MicrocoulombPerKilogram"}`)
    microcoulombs_per_kilogramResult, err := factory.FromDtoJSON(microcoulombs_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrocoulombPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microcoulombs_per_kilogramResult.Convert(units.RadiationExposureMicrocoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrocoulombPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with MillicoulombPerKilogram unit
    millicoulombs_per_kilogramJSON := []byte(`{"value": 100, "unit": "MillicoulombPerKilogram"}`)
    millicoulombs_per_kilogramResult, err := factory.FromDtoJSON(millicoulombs_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillicoulombPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millicoulombs_per_kilogramResult.Convert(units.RadiationExposureMillicoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillicoulombPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with Microroentgen unit
    microroentgensJSON := []byte(`{"value": 100, "unit": "Microroentgen"}`)
    microroentgensResult, err := factory.FromDtoJSON(microroentgensJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microroentgen unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microroentgensResult.Convert(units.RadiationExposureMicroroentgen)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microroentgen = %v, want %v", converted, 100)
    }
    // Test JSON with Milliroentgen unit
    milliroentgensJSON := []byte(`{"value": 100, "unit": "Milliroentgen"}`)
    milliroentgensResult, err := factory.FromDtoJSON(milliroentgensJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Milliroentgen unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliroentgensResult.Convert(units.RadiationExposureMilliroentgen)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliroentgen = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "CoulombPerKilogram"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromCoulombsPerKilogram function
func TestRadiationExposureFactory_FromCoulombsPerKilogram(t *testing.T) {
    factory := units.RadiationExposureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCoulombsPerKilogram(100)
    if err != nil {
        t.Errorf("FromCoulombsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationExposureCoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCoulombsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCoulombsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromCoulombsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromCoulombsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromCoulombsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromCoulombsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromCoulombsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCoulombsPerKilogram(0)
    if err != nil {
        t.Errorf("FromCoulombsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationExposureCoulombPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCoulombsPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromRoentgens function
func TestRadiationExposureFactory_FromRoentgens(t *testing.T) {
    factory := units.RadiationExposureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRoentgens(100)
    if err != nil {
        t.Errorf("FromRoentgens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationExposureRoentgen)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromRoentgens() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromRoentgens(math.NaN())
    if err == nil {
        t.Error("FromRoentgens() with NaN value should return error")
    }

    _, err = factory.FromRoentgens(math.Inf(1))
    if err == nil {
        t.Error("FromRoentgens() with +Inf value should return error")
    }

    _, err = factory.FromRoentgens(math.Inf(-1))
    if err == nil {
        t.Error("FromRoentgens() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromRoentgens(0)
    if err != nil {
        t.Errorf("FromRoentgens() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationExposureRoentgen)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRoentgens() with zero value = %v, want 0", converted)
    }
}
// Test FromPicocoulombsPerKilogram function
func TestRadiationExposureFactory_FromPicocoulombsPerKilogram(t *testing.T) {
    factory := units.RadiationExposureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicocoulombsPerKilogram(100)
    if err != nil {
        t.Errorf("FromPicocoulombsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationExposurePicocoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicocoulombsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicocoulombsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromPicocoulombsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromPicocoulombsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromPicocoulombsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromPicocoulombsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromPicocoulombsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicocoulombsPerKilogram(0)
    if err != nil {
        t.Errorf("FromPicocoulombsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationExposurePicocoulombPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicocoulombsPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromNanocoulombsPerKilogram function
func TestRadiationExposureFactory_FromNanocoulombsPerKilogram(t *testing.T) {
    factory := units.RadiationExposureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanocoulombsPerKilogram(100)
    if err != nil {
        t.Errorf("FromNanocoulombsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationExposureNanocoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanocoulombsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanocoulombsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromNanocoulombsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromNanocoulombsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromNanocoulombsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromNanocoulombsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromNanocoulombsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanocoulombsPerKilogram(0)
    if err != nil {
        t.Errorf("FromNanocoulombsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationExposureNanocoulombPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanocoulombsPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrocoulombsPerKilogram function
func TestRadiationExposureFactory_FromMicrocoulombsPerKilogram(t *testing.T) {
    factory := units.RadiationExposureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrocoulombsPerKilogram(100)
    if err != nil {
        t.Errorf("FromMicrocoulombsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationExposureMicrocoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrocoulombsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrocoulombsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromMicrocoulombsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromMicrocoulombsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromMicrocoulombsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromMicrocoulombsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrocoulombsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrocoulombsPerKilogram(0)
    if err != nil {
        t.Errorf("FromMicrocoulombsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationExposureMicrocoulombPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrocoulombsPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromMillicoulombsPerKilogram function
func TestRadiationExposureFactory_FromMillicoulombsPerKilogram(t *testing.T) {
    factory := units.RadiationExposureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillicoulombsPerKilogram(100)
    if err != nil {
        t.Errorf("FromMillicoulombsPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationExposureMillicoulombPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillicoulombsPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillicoulombsPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromMillicoulombsPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromMillicoulombsPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromMillicoulombsPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromMillicoulombsPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromMillicoulombsPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillicoulombsPerKilogram(0)
    if err != nil {
        t.Errorf("FromMillicoulombsPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationExposureMillicoulombPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillicoulombsPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromMicroroentgens function
func TestRadiationExposureFactory_FromMicroroentgens(t *testing.T) {
    factory := units.RadiationExposureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicroroentgens(100)
    if err != nil {
        t.Errorf("FromMicroroentgens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationExposureMicroroentgen)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicroroentgens() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicroroentgens(math.NaN())
    if err == nil {
        t.Error("FromMicroroentgens() with NaN value should return error")
    }

    _, err = factory.FromMicroroentgens(math.Inf(1))
    if err == nil {
        t.Error("FromMicroroentgens() with +Inf value should return error")
    }

    _, err = factory.FromMicroroentgens(math.Inf(-1))
    if err == nil {
        t.Error("FromMicroroentgens() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicroroentgens(0)
    if err != nil {
        t.Errorf("FromMicroroentgens() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationExposureMicroroentgen)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicroroentgens() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliroentgens function
func TestRadiationExposureFactory_FromMilliroentgens(t *testing.T) {
    factory := units.RadiationExposureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliroentgens(100)
    if err != nil {
        t.Errorf("FromMilliroentgens() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationExposureMilliroentgen)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliroentgens() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliroentgens(math.NaN())
    if err == nil {
        t.Error("FromMilliroentgens() with NaN value should return error")
    }

    _, err = factory.FromMilliroentgens(math.Inf(1))
    if err == nil {
        t.Error("FromMilliroentgens() with +Inf value should return error")
    }

    _, err = factory.FromMilliroentgens(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliroentgens() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliroentgens(0)
    if err != nil {
        t.Errorf("FromMilliroentgens() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationExposureMilliroentgen)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliroentgens() with zero value = %v, want 0", converted)
    }
}

func TestRadiationExposureToString(t *testing.T) {
	factory := units.RadiationExposureFactory{}
	a, err := factory.CreateRadiationExposure(45, units.RadiationExposureCoulombPerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RadiationExposureCoulombPerKilogram, 2)
	expected := "45.00 " + units.GetRadiationExposureAbbreviation(units.RadiationExposureCoulombPerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RadiationExposureCoulombPerKilogram, -1)
	expected = "45 " + units.GetRadiationExposureAbbreviation(units.RadiationExposureCoulombPerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRadiationExposure_EqualityAndComparison(t *testing.T) {
	factory := units.RadiationExposureFactory{}
	a1, _ := factory.CreateRadiationExposure(60, units.RadiationExposureCoulombPerKilogram)
	a2, _ := factory.CreateRadiationExposure(60, units.RadiationExposureCoulombPerKilogram)
	a3, _ := factory.CreateRadiationExposure(120, units.RadiationExposureCoulombPerKilogram)

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

func TestRadiationExposure_Arithmetic(t *testing.T) {
	factory := units.RadiationExposureFactory{}
	a1, _ := factory.CreateRadiationExposure(30, units.RadiationExposureCoulombPerKilogram)
	a2, _ := factory.CreateRadiationExposure(45, units.RadiationExposureCoulombPerKilogram)

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


func TestGetRadiationExposureAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.RadiationExposureUnits
        want string
    }{
        {
            name: "CoulombPerKilogram abbreviation",
            unit: units.RadiationExposureCoulombPerKilogram,
            want: "C/kg",
        },
        {
            name: "Roentgen abbreviation",
            unit: units.RadiationExposureRoentgen,
            want: "R",
        },
        {
            name: "PicocoulombPerKilogram abbreviation",
            unit: units.RadiationExposurePicocoulombPerKilogram,
            want: "pC/kg",
        },
        {
            name: "NanocoulombPerKilogram abbreviation",
            unit: units.RadiationExposureNanocoulombPerKilogram,
            want: "nC/kg",
        },
        {
            name: "MicrocoulombPerKilogram abbreviation",
            unit: units.RadiationExposureMicrocoulombPerKilogram,
            want: "μC/kg",
        },
        {
            name: "MillicoulombPerKilogram abbreviation",
            unit: units.RadiationExposureMillicoulombPerKilogram,
            want: "mC/kg",
        },
        {
            name: "Microroentgen abbreviation",
            unit: units.RadiationExposureMicroroentgen,
            want: "μR",
        },
        {
            name: "Milliroentgen abbreviation",
            unit: units.RadiationExposureMilliroentgen,
            want: "mR",
        },
        {
            name: "invalid unit",
            unit: units.RadiationExposureUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetRadiationExposureAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetRadiationExposureAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestRadiationExposure_String(t *testing.T) {
    factory := units.RadiationExposureFactory{}
    
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
            unit, err := factory.CreateRadiationExposure(tt.value, units.RadiationExposureCoulombPerKilogram)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("RadiationExposure.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestRadiationExposure_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.RadiationExposureFactory{}

	_, err := uf.CreateRadiationExposure(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}