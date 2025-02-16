// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRotationalSpeedDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "RadianPerSecond"}`
	
	factory := units.RotationalSpeedDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RotationalSpeedRadianPerSecond {
		t.Errorf("expected unit %v, got %v", units.RotationalSpeedRadianPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "RadianPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRotationalSpeedDto_ToJSON(t *testing.T) {
	dto := units.RotationalSpeedDto{
		Value: 45,
		Unit:  units.RotationalSpeedRadianPerSecond,
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
	if result["unit"].(string) != string(units.RotationalSpeedRadianPerSecond) {
		t.Errorf("expected unit %s, got %v", units.RotationalSpeedRadianPerSecond, result["unit"])
	}
}

func TestNewRotationalSpeed_InvalidValue(t *testing.T) {
	factory := units.RotationalSpeedFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRotationalSpeed(math.NaN(), units.RotationalSpeedRadianPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRotationalSpeed(math.Inf(1), units.RotationalSpeedRadianPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRotationalSpeedConversions(t *testing.T) {
	factory := units.RotationalSpeedFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRotationalSpeed(180, units.RotationalSpeedRadianPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to RadiansPerSecond.
		// No expected conversion value provided for RadiansPerSecond, verifying result is not NaN.
		result := a.RadiansPerSecond()
		cacheResult := a.RadiansPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to RadiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DegreesPerSecond.
		// No expected conversion value provided for DegreesPerSecond, verifying result is not NaN.
		result := a.DegreesPerSecond()
		cacheResult := a.DegreesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DegreesPerMinute.
		// No expected conversion value provided for DegreesPerMinute, verifying result is not NaN.
		result := a.DegreesPerMinute()
		cacheResult := a.DegreesPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to RevolutionsPerSecond.
		// No expected conversion value provided for RevolutionsPerSecond, verifying result is not NaN.
		result := a.RevolutionsPerSecond()
		cacheResult := a.RevolutionsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to RevolutionsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to RevolutionsPerMinute.
		// No expected conversion value provided for RevolutionsPerMinute, verifying result is not NaN.
		result := a.RevolutionsPerMinute()
		cacheResult := a.RevolutionsPerMinute()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to RevolutionsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to NanoradiansPerSecond.
		// No expected conversion value provided for NanoradiansPerSecond, verifying result is not NaN.
		result := a.NanoradiansPerSecond()
		cacheResult := a.NanoradiansPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanoradiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicroradiansPerSecond.
		// No expected conversion value provided for MicroradiansPerSecond, verifying result is not NaN.
		result := a.MicroradiansPerSecond()
		cacheResult := a.MicroradiansPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicroradiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MilliradiansPerSecond.
		// No expected conversion value provided for MilliradiansPerSecond, verifying result is not NaN.
		result := a.MilliradiansPerSecond()
		cacheResult := a.MilliradiansPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliradiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CentiradiansPerSecond.
		// No expected conversion value provided for CentiradiansPerSecond, verifying result is not NaN.
		result := a.CentiradiansPerSecond()
		cacheResult := a.CentiradiansPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentiradiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DeciradiansPerSecond.
		// No expected conversion value provided for DeciradiansPerSecond, verifying result is not NaN.
		result := a.DeciradiansPerSecond()
		cacheResult := a.DeciradiansPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DeciradiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to NanodegreesPerSecond.
		// No expected conversion value provided for NanodegreesPerSecond, verifying result is not NaN.
		result := a.NanodegreesPerSecond()
		cacheResult := a.NanodegreesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanodegreesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicrodegreesPerSecond.
		// No expected conversion value provided for MicrodegreesPerSecond, verifying result is not NaN.
		result := a.MicrodegreesPerSecond()
		cacheResult := a.MicrodegreesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrodegreesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillidegreesPerSecond.
		// No expected conversion value provided for MillidegreesPerSecond, verifying result is not NaN.
		result := a.MillidegreesPerSecond()
		cacheResult := a.MillidegreesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillidegreesPerSecond returned NaN")
		}
	}
}

func TestRotationalSpeed_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RotationalSpeedFactory{}
	a, err := factory.CreateRotationalSpeed(90, units.RotationalSpeedRadianPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RotationalSpeedRadianPerSecond {
		t.Errorf("expected default unit RadianPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RotationalSpeedRadianPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RotationalSpeedDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RotationalSpeedRadianPerSecond {
		t.Errorf("expected unit RadianPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRotationalSpeedFactory_FromDto(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedRadianPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.RotationalSpeedDto{
        Value: math.NaN(),
        Unit:  units.RotationalSpeedRadianPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test RadianPerSecond conversion
    radians_per_secondDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedRadianPerSecond,
    }
    
    var radians_per_secondResult *units.RotationalSpeed
    radians_per_secondResult, err = factory.FromDto(radians_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with RadianPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = radians_per_secondResult.Convert(units.RotationalSpeedRadianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RadianPerSecond = %v, want %v", converted, 100)
    }
    // Test DegreePerSecond conversion
    degrees_per_secondDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedDegreePerSecond,
    }
    
    var degrees_per_secondResult *units.RotationalSpeed
    degrees_per_secondResult, err = factory.FromDto(degrees_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DegreePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_per_secondResult.Convert(units.RotationalSpeedDegreePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreePerSecond = %v, want %v", converted, 100)
    }
    // Test DegreePerMinute conversion
    degrees_per_minuteDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedDegreePerMinute,
    }
    
    var degrees_per_minuteResult *units.RotationalSpeed
    degrees_per_minuteResult, err = factory.FromDto(degrees_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with DegreePerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_per_minuteResult.Convert(units.RotationalSpeedDegreePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreePerMinute = %v, want %v", converted, 100)
    }
    // Test RevolutionPerSecond conversion
    revolutions_per_secondDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedRevolutionPerSecond,
    }
    
    var revolutions_per_secondResult *units.RotationalSpeed
    revolutions_per_secondResult, err = factory.FromDto(revolutions_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with RevolutionPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = revolutions_per_secondResult.Convert(units.RotationalSpeedRevolutionPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RevolutionPerSecond = %v, want %v", converted, 100)
    }
    // Test RevolutionPerMinute conversion
    revolutions_per_minuteDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedRevolutionPerMinute,
    }
    
    var revolutions_per_minuteResult *units.RotationalSpeed
    revolutions_per_minuteResult, err = factory.FromDto(revolutions_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with RevolutionPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = revolutions_per_minuteResult.Convert(units.RotationalSpeedRevolutionPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RevolutionPerMinute = %v, want %v", converted, 100)
    }
    // Test NanoradianPerSecond conversion
    nanoradians_per_secondDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedNanoradianPerSecond,
    }
    
    var nanoradians_per_secondResult *units.RotationalSpeed
    nanoradians_per_secondResult, err = factory.FromDto(nanoradians_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with NanoradianPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoradians_per_secondResult.Convert(units.RotationalSpeedNanoradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoradianPerSecond = %v, want %v", converted, 100)
    }
    // Test MicroradianPerSecond conversion
    microradians_per_secondDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedMicroradianPerSecond,
    }
    
    var microradians_per_secondResult *units.RotationalSpeed
    microradians_per_secondResult, err = factory.FromDto(microradians_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MicroradianPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microradians_per_secondResult.Convert(units.RotationalSpeedMicroradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroradianPerSecond = %v, want %v", converted, 100)
    }
    // Test MilliradianPerSecond conversion
    milliradians_per_secondDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedMilliradianPerSecond,
    }
    
    var milliradians_per_secondResult *units.RotationalSpeed
    milliradians_per_secondResult, err = factory.FromDto(milliradians_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MilliradianPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliradians_per_secondResult.Convert(units.RotationalSpeedMilliradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliradianPerSecond = %v, want %v", converted, 100)
    }
    // Test CentiradianPerSecond conversion
    centiradians_per_secondDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedCentiradianPerSecond,
    }
    
    var centiradians_per_secondResult *units.RotationalSpeed
    centiradians_per_secondResult, err = factory.FromDto(centiradians_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with CentiradianPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiradians_per_secondResult.Convert(units.RotationalSpeedCentiradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiradianPerSecond = %v, want %v", converted, 100)
    }
    // Test DeciradianPerSecond conversion
    deciradians_per_secondDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedDeciradianPerSecond,
    }
    
    var deciradians_per_secondResult *units.RotationalSpeed
    deciradians_per_secondResult, err = factory.FromDto(deciradians_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DeciradianPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciradians_per_secondResult.Convert(units.RotationalSpeedDeciradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciradianPerSecond = %v, want %v", converted, 100)
    }
    // Test NanodegreePerSecond conversion
    nanodegrees_per_secondDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedNanodegreePerSecond,
    }
    
    var nanodegrees_per_secondResult *units.RotationalSpeed
    nanodegrees_per_secondResult, err = factory.FromDto(nanodegrees_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with NanodegreePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanodegrees_per_secondResult.Convert(units.RotationalSpeedNanodegreePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanodegreePerSecond = %v, want %v", converted, 100)
    }
    // Test MicrodegreePerSecond conversion
    microdegrees_per_secondDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedMicrodegreePerSecond,
    }
    
    var microdegrees_per_secondResult *units.RotationalSpeed
    microdegrees_per_secondResult, err = factory.FromDto(microdegrees_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MicrodegreePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microdegrees_per_secondResult.Convert(units.RotationalSpeedMicrodegreePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrodegreePerSecond = %v, want %v", converted, 100)
    }
    // Test MillidegreePerSecond conversion
    millidegrees_per_secondDto := units.RotationalSpeedDto{
        Value: 100,
        Unit:  units.RotationalSpeedMillidegreePerSecond,
    }
    
    var millidegrees_per_secondResult *units.RotationalSpeed
    millidegrees_per_secondResult, err = factory.FromDto(millidegrees_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MillidegreePerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millidegrees_per_secondResult.Convert(units.RotationalSpeedMillidegreePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillidegreePerSecond = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.RotationalSpeedDto{
        Value: 0,
        Unit:  units.RotationalSpeedRadianPerSecond,
    }
    
    var zeroResult *units.RotationalSpeed
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestRotationalSpeedFactory_FromDtoJSON(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "RadianPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "RadianPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.RotationalSpeedDto{
        Value: nanValue,
        Unit:  units.RotationalSpeedRadianPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with RadianPerSecond unit
    radians_per_secondJSON := []byte(`{"value": 100, "unit": "RadianPerSecond"}`)
    radians_per_secondResult, err := factory.FromDtoJSON(radians_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with RadianPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = radians_per_secondResult.Convert(units.RotationalSpeedRadianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RadianPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DegreePerSecond unit
    degrees_per_secondJSON := []byte(`{"value": 100, "unit": "DegreePerSecond"}`)
    degrees_per_secondResult, err := factory.FromDtoJSON(degrees_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_per_secondResult.Convert(units.RotationalSpeedDegreePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DegreePerMinute unit
    degrees_per_minuteJSON := []byte(`{"value": 100, "unit": "DegreePerMinute"}`)
    degrees_per_minuteResult, err := factory.FromDtoJSON(degrees_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreePerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_per_minuteResult.Convert(units.RotationalSpeedDegreePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreePerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with RevolutionPerSecond unit
    revolutions_per_secondJSON := []byte(`{"value": 100, "unit": "RevolutionPerSecond"}`)
    revolutions_per_secondResult, err := factory.FromDtoJSON(revolutions_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with RevolutionPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = revolutions_per_secondResult.Convert(units.RotationalSpeedRevolutionPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RevolutionPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with RevolutionPerMinute unit
    revolutions_per_minuteJSON := []byte(`{"value": 100, "unit": "RevolutionPerMinute"}`)
    revolutions_per_minuteResult, err := factory.FromDtoJSON(revolutions_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with RevolutionPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = revolutions_per_minuteResult.Convert(units.RotationalSpeedRevolutionPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RevolutionPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with NanoradianPerSecond unit
    nanoradians_per_secondJSON := []byte(`{"value": 100, "unit": "NanoradianPerSecond"}`)
    nanoradians_per_secondResult, err := factory.FromDtoJSON(nanoradians_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanoradianPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoradians_per_secondResult.Convert(units.RotationalSpeedNanoradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanoradianPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicroradianPerSecond unit
    microradians_per_secondJSON := []byte(`{"value": 100, "unit": "MicroradianPerSecond"}`)
    microradians_per_secondResult, err := factory.FromDtoJSON(microradians_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicroradianPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microradians_per_secondResult.Convert(units.RotationalSpeedMicroradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicroradianPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MilliradianPerSecond unit
    milliradians_per_secondJSON := []byte(`{"value": 100, "unit": "MilliradianPerSecond"}`)
    milliradians_per_secondResult, err := factory.FromDtoJSON(milliradians_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliradianPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliradians_per_secondResult.Convert(units.RotationalSpeedMilliradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliradianPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with CentiradianPerSecond unit
    centiradians_per_secondJSON := []byte(`{"value": 100, "unit": "CentiradianPerSecond"}`)
    centiradians_per_secondResult, err := factory.FromDtoJSON(centiradians_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentiradianPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiradians_per_secondResult.Convert(units.RotationalSpeedCentiradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentiradianPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DeciradianPerSecond unit
    deciradians_per_secondJSON := []byte(`{"value": 100, "unit": "DeciradianPerSecond"}`)
    deciradians_per_secondResult, err := factory.FromDtoJSON(deciradians_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DeciradianPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciradians_per_secondResult.Convert(units.RotationalSpeedDeciradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciradianPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with NanodegreePerSecond unit
    nanodegrees_per_secondJSON := []byte(`{"value": 100, "unit": "NanodegreePerSecond"}`)
    nanodegrees_per_secondResult, err := factory.FromDtoJSON(nanodegrees_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanodegreePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanodegrees_per_secondResult.Convert(units.RotationalSpeedNanodegreePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanodegreePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicrodegreePerSecond unit
    microdegrees_per_secondJSON := []byte(`{"value": 100, "unit": "MicrodegreePerSecond"}`)
    microdegrees_per_secondResult, err := factory.FromDtoJSON(microdegrees_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrodegreePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microdegrees_per_secondResult.Convert(units.RotationalSpeedMicrodegreePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrodegreePerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MillidegreePerSecond unit
    millidegrees_per_secondJSON := []byte(`{"value": 100, "unit": "MillidegreePerSecond"}`)
    millidegrees_per_secondResult, err := factory.FromDtoJSON(millidegrees_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillidegreePerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millidegrees_per_secondResult.Convert(units.RotationalSpeedMillidegreePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillidegreePerSecond = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "RadianPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromRadiansPerSecond function
func TestRotationalSpeedFactory_FromRadiansPerSecond(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRadiansPerSecond(100)
    if err != nil {
        t.Errorf("FromRadiansPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalSpeedRadianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromRadiansPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromRadiansPerSecond(math.NaN())
    if err == nil {
        t.Error("FromRadiansPerSecond() with NaN value should return error")
    }

    _, err = factory.FromRadiansPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromRadiansPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromRadiansPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromRadiansPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromRadiansPerSecond(0)
    if err != nil {
        t.Errorf("FromRadiansPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalSpeedRadianPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRadiansPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesPerSecond function
func TestRotationalSpeedFactory_FromDegreesPerSecond(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesPerSecond(100)
    if err != nil {
        t.Errorf("FromDegreesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalSpeedDegreePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDegreesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDegreesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDegreesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesPerSecond(0)
    if err != nil {
        t.Errorf("FromDegreesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalSpeedDegreePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesPerMinute function
func TestRotationalSpeedFactory_FromDegreesPerMinute(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesPerMinute(100)
    if err != nil {
        t.Errorf("FromDegreesPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalSpeedDegreePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesPerMinute(math.NaN())
    if err == nil {
        t.Error("FromDegreesPerMinute() with NaN value should return error")
    }

    _, err = factory.FromDegreesPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromDegreesPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesPerMinute(0)
    if err != nil {
        t.Errorf("FromDegreesPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalSpeedDegreePerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromRevolutionsPerSecond function
func TestRotationalSpeedFactory_FromRevolutionsPerSecond(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRevolutionsPerSecond(100)
    if err != nil {
        t.Errorf("FromRevolutionsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalSpeedRevolutionPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromRevolutionsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromRevolutionsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromRevolutionsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromRevolutionsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromRevolutionsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromRevolutionsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromRevolutionsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromRevolutionsPerSecond(0)
    if err != nil {
        t.Errorf("FromRevolutionsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalSpeedRevolutionPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRevolutionsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromRevolutionsPerMinute function
func TestRotationalSpeedFactory_FromRevolutionsPerMinute(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRevolutionsPerMinute(100)
    if err != nil {
        t.Errorf("FromRevolutionsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalSpeedRevolutionPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromRevolutionsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromRevolutionsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromRevolutionsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromRevolutionsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromRevolutionsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromRevolutionsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromRevolutionsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromRevolutionsPerMinute(0)
    if err != nil {
        t.Errorf("FromRevolutionsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalSpeedRevolutionPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRevolutionsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromNanoradiansPerSecond function
func TestRotationalSpeedFactory_FromNanoradiansPerSecond(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanoradiansPerSecond(100)
    if err != nil {
        t.Errorf("FromNanoradiansPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalSpeedNanoradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanoradiansPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanoradiansPerSecond(math.NaN())
    if err == nil {
        t.Error("FromNanoradiansPerSecond() with NaN value should return error")
    }

    _, err = factory.FromNanoradiansPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromNanoradiansPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromNanoradiansPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromNanoradiansPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanoradiansPerSecond(0)
    if err != nil {
        t.Errorf("FromNanoradiansPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalSpeedNanoradianPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanoradiansPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMicroradiansPerSecond function
func TestRotationalSpeedFactory_FromMicroradiansPerSecond(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicroradiansPerSecond(100)
    if err != nil {
        t.Errorf("FromMicroradiansPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalSpeedMicroradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicroradiansPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicroradiansPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMicroradiansPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMicroradiansPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMicroradiansPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMicroradiansPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMicroradiansPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicroradiansPerSecond(0)
    if err != nil {
        t.Errorf("FromMicroradiansPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalSpeedMicroradianPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicroradiansPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliradiansPerSecond function
func TestRotationalSpeedFactory_FromMilliradiansPerSecond(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliradiansPerSecond(100)
    if err != nil {
        t.Errorf("FromMilliradiansPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalSpeedMilliradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliradiansPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliradiansPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMilliradiansPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMilliradiansPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMilliradiansPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMilliradiansPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliradiansPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliradiansPerSecond(0)
    if err != nil {
        t.Errorf("FromMilliradiansPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalSpeedMilliradianPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliradiansPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromCentiradiansPerSecond function
func TestRotationalSpeedFactory_FromCentiradiansPerSecond(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentiradiansPerSecond(100)
    if err != nil {
        t.Errorf("FromCentiradiansPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalSpeedCentiradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentiradiansPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentiradiansPerSecond(math.NaN())
    if err == nil {
        t.Error("FromCentiradiansPerSecond() with NaN value should return error")
    }

    _, err = factory.FromCentiradiansPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromCentiradiansPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromCentiradiansPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromCentiradiansPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentiradiansPerSecond(0)
    if err != nil {
        t.Errorf("FromCentiradiansPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalSpeedCentiradianPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentiradiansPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDeciradiansPerSecond function
func TestRotationalSpeedFactory_FromDeciradiansPerSecond(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDeciradiansPerSecond(100)
    if err != nil {
        t.Errorf("FromDeciradiansPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalSpeedDeciradianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDeciradiansPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDeciradiansPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDeciradiansPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDeciradiansPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDeciradiansPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDeciradiansPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDeciradiansPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDeciradiansPerSecond(0)
    if err != nil {
        t.Errorf("FromDeciradiansPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalSpeedDeciradianPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDeciradiansPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromNanodegreesPerSecond function
func TestRotationalSpeedFactory_FromNanodegreesPerSecond(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanodegreesPerSecond(100)
    if err != nil {
        t.Errorf("FromNanodegreesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalSpeedNanodegreePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanodegreesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanodegreesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromNanodegreesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromNanodegreesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromNanodegreesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromNanodegreesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromNanodegreesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanodegreesPerSecond(0)
    if err != nil {
        t.Errorf("FromNanodegreesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalSpeedNanodegreePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanodegreesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrodegreesPerSecond function
func TestRotationalSpeedFactory_FromMicrodegreesPerSecond(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrodegreesPerSecond(100)
    if err != nil {
        t.Errorf("FromMicrodegreesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalSpeedMicrodegreePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrodegreesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrodegreesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMicrodegreesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMicrodegreesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMicrodegreesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMicrodegreesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrodegreesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrodegreesPerSecond(0)
    if err != nil {
        t.Errorf("FromMicrodegreesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalSpeedMicrodegreePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrodegreesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMillidegreesPerSecond function
func TestRotationalSpeedFactory_FromMillidegreesPerSecond(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillidegreesPerSecond(100)
    if err != nil {
        t.Errorf("FromMillidegreesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RotationalSpeedMillidegreePerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillidegreesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillidegreesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMillidegreesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMillidegreesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMillidegreesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMillidegreesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMillidegreesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillidegreesPerSecond(0)
    if err != nil {
        t.Errorf("FromMillidegreesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RotationalSpeedMillidegreePerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillidegreesPerSecond() with zero value = %v, want 0", converted)
    }
}

func TestRotationalSpeedToString(t *testing.T) {
	factory := units.RotationalSpeedFactory{}
	a, err := factory.CreateRotationalSpeed(45, units.RotationalSpeedRadianPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RotationalSpeedRadianPerSecond, 2)
	expected := "45.00 " + units.GetRotationalSpeedAbbreviation(units.RotationalSpeedRadianPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RotationalSpeedRadianPerSecond, -1)
	expected = "45 " + units.GetRotationalSpeedAbbreviation(units.RotationalSpeedRadianPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRotationalSpeed_EqualityAndComparison(t *testing.T) {
	factory := units.RotationalSpeedFactory{}
	a1, _ := factory.CreateRotationalSpeed(60, units.RotationalSpeedRadianPerSecond)
	a2, _ := factory.CreateRotationalSpeed(60, units.RotationalSpeedRadianPerSecond)
	a3, _ := factory.CreateRotationalSpeed(120, units.RotationalSpeedRadianPerSecond)

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

func TestRotationalSpeed_Arithmetic(t *testing.T) {
	factory := units.RotationalSpeedFactory{}
	a1, _ := factory.CreateRotationalSpeed(30, units.RotationalSpeedRadianPerSecond)
	a2, _ := factory.CreateRotationalSpeed(45, units.RotationalSpeedRadianPerSecond)

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


func TestGetRotationalSpeedAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.RotationalSpeedUnits
        want string
    }{
        {
            name: "RadianPerSecond abbreviation",
            unit: units.RotationalSpeedRadianPerSecond,
            want: "rad/s",
        },
        {
            name: "DegreePerSecond abbreviation",
            unit: units.RotationalSpeedDegreePerSecond,
            want: "°/s",
        },
        {
            name: "DegreePerMinute abbreviation",
            unit: units.RotationalSpeedDegreePerMinute,
            want: "°/min",
        },
        {
            name: "RevolutionPerSecond abbreviation",
            unit: units.RotationalSpeedRevolutionPerSecond,
            want: "r/s",
        },
        {
            name: "RevolutionPerMinute abbreviation",
            unit: units.RotationalSpeedRevolutionPerMinute,
            want: "rpm",
        },
        {
            name: "NanoradianPerSecond abbreviation",
            unit: units.RotationalSpeedNanoradianPerSecond,
            want: "nrad/s",
        },
        {
            name: "MicroradianPerSecond abbreviation",
            unit: units.RotationalSpeedMicroradianPerSecond,
            want: "μrad/s",
        },
        {
            name: "MilliradianPerSecond abbreviation",
            unit: units.RotationalSpeedMilliradianPerSecond,
            want: "mrad/s",
        },
        {
            name: "CentiradianPerSecond abbreviation",
            unit: units.RotationalSpeedCentiradianPerSecond,
            want: "crad/s",
        },
        {
            name: "DeciradianPerSecond abbreviation",
            unit: units.RotationalSpeedDeciradianPerSecond,
            want: "drad/s",
        },
        {
            name: "NanodegreePerSecond abbreviation",
            unit: units.RotationalSpeedNanodegreePerSecond,
            want: "n°/s",
        },
        {
            name: "MicrodegreePerSecond abbreviation",
            unit: units.RotationalSpeedMicrodegreePerSecond,
            want: "μ°/s",
        },
        {
            name: "MillidegreePerSecond abbreviation",
            unit: units.RotationalSpeedMillidegreePerSecond,
            want: "m°/s",
        },
        {
            name: "invalid unit",
            unit: units.RotationalSpeedUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetRotationalSpeedAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetRotationalSpeedAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestRotationalSpeed_String(t *testing.T) {
    factory := units.RotationalSpeedFactory{}
    
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
            unit, err := factory.CreateRotationalSpeed(tt.value, units.RotationalSpeedRadianPerSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("RotationalSpeed.String() = %v, want %v", got, tt.want)
            }
        })
    }
}