// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRadiationEquivalentDoseRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SievertPerSecond"}`
	
	factory := units.RadiationEquivalentDoseRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RadiationEquivalentDoseRateSievertPerSecond {
		t.Errorf("expected unit %v, got %v", units.RadiationEquivalentDoseRateSievertPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SievertPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRadiationEquivalentDoseRateDto_ToJSON(t *testing.T) {
	dto := units.RadiationEquivalentDoseRateDto{
		Value: 45,
		Unit:  units.RadiationEquivalentDoseRateSievertPerSecond,
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
	if result["unit"].(string) != string(units.RadiationEquivalentDoseRateSievertPerSecond) {
		t.Errorf("expected unit %s, got %v", units.RadiationEquivalentDoseRateSievertPerSecond, result["unit"])
	}
}

func TestNewRadiationEquivalentDoseRate_InvalidValue(t *testing.T) {
	factory := units.RadiationEquivalentDoseRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRadiationEquivalentDoseRate(math.NaN(), units.RadiationEquivalentDoseRateSievertPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRadiationEquivalentDoseRate(math.Inf(1), units.RadiationEquivalentDoseRateSievertPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRadiationEquivalentDoseRateConversions(t *testing.T) {
	factory := units.RadiationEquivalentDoseRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRadiationEquivalentDoseRate(180, units.RadiationEquivalentDoseRateSievertPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to SievertsPerHour.
		// No expected conversion value provided for SievertsPerHour, verifying result is not NaN.
		result := a.SievertsPerHour()
		cacheResult := a.SievertsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SievertsPerHour returned NaN")
		}
	}
	{
		// Test conversion to SievertsPerSecond.
		// No expected conversion value provided for SievertsPerSecond, verifying result is not NaN.
		result := a.SievertsPerSecond()
		cacheResult := a.SievertsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SievertsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to RoentgensEquivalentManPerHour.
		// No expected conversion value provided for RoentgensEquivalentManPerHour, verifying result is not NaN.
		result := a.RoentgensEquivalentManPerHour()
		cacheResult := a.RoentgensEquivalentManPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to RoentgensEquivalentManPerHour returned NaN")
		}
	}
	{
		// Test conversion to NanosievertsPerHour.
		// No expected conversion value provided for NanosievertsPerHour, verifying result is not NaN.
		result := a.NanosievertsPerHour()
		cacheResult := a.NanosievertsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanosievertsPerHour returned NaN")
		}
	}
	{
		// Test conversion to MicrosievertsPerHour.
		// No expected conversion value provided for MicrosievertsPerHour, verifying result is not NaN.
		result := a.MicrosievertsPerHour()
		cacheResult := a.MicrosievertsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrosievertsPerHour returned NaN")
		}
	}
	{
		// Test conversion to MillisievertsPerHour.
		// No expected conversion value provided for MillisievertsPerHour, verifying result is not NaN.
		result := a.MillisievertsPerHour()
		cacheResult := a.MillisievertsPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillisievertsPerHour returned NaN")
		}
	}
	{
		// Test conversion to NanosievertsPerSecond.
		// No expected conversion value provided for NanosievertsPerSecond, verifying result is not NaN.
		result := a.NanosievertsPerSecond()
		cacheResult := a.NanosievertsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanosievertsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MicrosievertsPerSecond.
		// No expected conversion value provided for MicrosievertsPerSecond, verifying result is not NaN.
		result := a.MicrosievertsPerSecond()
		cacheResult := a.MicrosievertsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrosievertsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillisievertsPerSecond.
		// No expected conversion value provided for MillisievertsPerSecond, verifying result is not NaN.
		result := a.MillisievertsPerSecond()
		cacheResult := a.MillisievertsPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillisievertsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MilliroentgensEquivalentManPerHour.
		// No expected conversion value provided for MilliroentgensEquivalentManPerHour, verifying result is not NaN.
		result := a.MilliroentgensEquivalentManPerHour()
		cacheResult := a.MilliroentgensEquivalentManPerHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilliroentgensEquivalentManPerHour returned NaN")
		}
	}
}

func TestRadiationEquivalentDoseRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RadiationEquivalentDoseRateFactory{}
	a, err := factory.CreateRadiationEquivalentDoseRate(90, units.RadiationEquivalentDoseRateSievertPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RadiationEquivalentDoseRateSievertPerSecond {
		t.Errorf("expected default unit SievertPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RadiationEquivalentDoseRateSievertPerHour
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RadiationEquivalentDoseRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RadiationEquivalentDoseRateSievertPerSecond {
		t.Errorf("expected unit SievertPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRadiationEquivalentDoseRateFactory_FromDto(t *testing.T) {
    factory := units.RadiationEquivalentDoseRateFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.RadiationEquivalentDoseRateDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseRateSievertPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.RadiationEquivalentDoseRateDto{
        Value: math.NaN(),
        Unit:  units.RadiationEquivalentDoseRateSievertPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test SievertPerHour conversion
    sieverts_per_hourDto := units.RadiationEquivalentDoseRateDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseRateSievertPerHour,
    }
    
    var sieverts_per_hourResult *units.RadiationEquivalentDoseRate
    sieverts_per_hourResult, err = factory.FromDto(sieverts_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with SievertPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = sieverts_per_hourResult.Convert(units.RadiationEquivalentDoseRateSievertPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SievertPerHour = %v, want %v", converted, 100)
    }
    // Test SievertPerSecond conversion
    sieverts_per_secondDto := units.RadiationEquivalentDoseRateDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseRateSievertPerSecond,
    }
    
    var sieverts_per_secondResult *units.RadiationEquivalentDoseRate
    sieverts_per_secondResult, err = factory.FromDto(sieverts_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with SievertPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = sieverts_per_secondResult.Convert(units.RadiationEquivalentDoseRateSievertPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SievertPerSecond = %v, want %v", converted, 100)
    }
    // Test RoentgenEquivalentManPerHour conversion
    roentgens_equivalent_man_per_hourDto := units.RadiationEquivalentDoseRateDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseRateRoentgenEquivalentManPerHour,
    }
    
    var roentgens_equivalent_man_per_hourResult *units.RadiationEquivalentDoseRate
    roentgens_equivalent_man_per_hourResult, err = factory.FromDto(roentgens_equivalent_man_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with RoentgenEquivalentManPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = roentgens_equivalent_man_per_hourResult.Convert(units.RadiationEquivalentDoseRateRoentgenEquivalentManPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RoentgenEquivalentManPerHour = %v, want %v", converted, 100)
    }
    // Test NanosievertPerHour conversion
    nanosieverts_per_hourDto := units.RadiationEquivalentDoseRateDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseRateNanosievertPerHour,
    }
    
    var nanosieverts_per_hourResult *units.RadiationEquivalentDoseRate
    nanosieverts_per_hourResult, err = factory.FromDto(nanosieverts_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with NanosievertPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanosieverts_per_hourResult.Convert(units.RadiationEquivalentDoseRateNanosievertPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanosievertPerHour = %v, want %v", converted, 100)
    }
    // Test MicrosievertPerHour conversion
    microsieverts_per_hourDto := units.RadiationEquivalentDoseRateDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseRateMicrosievertPerHour,
    }
    
    var microsieverts_per_hourResult *units.RadiationEquivalentDoseRate
    microsieverts_per_hourResult, err = factory.FromDto(microsieverts_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MicrosievertPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microsieverts_per_hourResult.Convert(units.RadiationEquivalentDoseRateMicrosievertPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrosievertPerHour = %v, want %v", converted, 100)
    }
    // Test MillisievertPerHour conversion
    millisieverts_per_hourDto := units.RadiationEquivalentDoseRateDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseRateMillisievertPerHour,
    }
    
    var millisieverts_per_hourResult *units.RadiationEquivalentDoseRate
    millisieverts_per_hourResult, err = factory.FromDto(millisieverts_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MillisievertPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millisieverts_per_hourResult.Convert(units.RadiationEquivalentDoseRateMillisievertPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillisievertPerHour = %v, want %v", converted, 100)
    }
    // Test NanosievertPerSecond conversion
    nanosieverts_per_secondDto := units.RadiationEquivalentDoseRateDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseRateNanosievertPerSecond,
    }
    
    var nanosieverts_per_secondResult *units.RadiationEquivalentDoseRate
    nanosieverts_per_secondResult, err = factory.FromDto(nanosieverts_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with NanosievertPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanosieverts_per_secondResult.Convert(units.RadiationEquivalentDoseRateNanosievertPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanosievertPerSecond = %v, want %v", converted, 100)
    }
    // Test MicrosievertPerSecond conversion
    microsieverts_per_secondDto := units.RadiationEquivalentDoseRateDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseRateMicrosievertPerSecond,
    }
    
    var microsieverts_per_secondResult *units.RadiationEquivalentDoseRate
    microsieverts_per_secondResult, err = factory.FromDto(microsieverts_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MicrosievertPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microsieverts_per_secondResult.Convert(units.RadiationEquivalentDoseRateMicrosievertPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrosievertPerSecond = %v, want %v", converted, 100)
    }
    // Test MillisievertPerSecond conversion
    millisieverts_per_secondDto := units.RadiationEquivalentDoseRateDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseRateMillisievertPerSecond,
    }
    
    var millisieverts_per_secondResult *units.RadiationEquivalentDoseRate
    millisieverts_per_secondResult, err = factory.FromDto(millisieverts_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MillisievertPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millisieverts_per_secondResult.Convert(units.RadiationEquivalentDoseRateMillisievertPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillisievertPerSecond = %v, want %v", converted, 100)
    }
    // Test MilliroentgenEquivalentManPerHour conversion
    milliroentgens_equivalent_man_per_hourDto := units.RadiationEquivalentDoseRateDto{
        Value: 100,
        Unit:  units.RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour,
    }
    
    var milliroentgens_equivalent_man_per_hourResult *units.RadiationEquivalentDoseRate
    milliroentgens_equivalent_man_per_hourResult, err = factory.FromDto(milliroentgens_equivalent_man_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MilliroentgenEquivalentManPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliroentgens_equivalent_man_per_hourResult.Convert(units.RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliroentgenEquivalentManPerHour = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.RadiationEquivalentDoseRateDto{
        Value: 0,
        Unit:  units.RadiationEquivalentDoseRateSievertPerSecond,
    }
    
    var zeroResult *units.RadiationEquivalentDoseRate
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestRadiationEquivalentDoseRateFactory_FromDtoJSON(t *testing.T) {
    factory := units.RadiationEquivalentDoseRateFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "SievertPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "SievertPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.RadiationEquivalentDoseRateDto{
        Value: nanValue,
        Unit:  units.RadiationEquivalentDoseRateSievertPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with SievertPerHour unit
    sieverts_per_hourJSON := []byte(`{"value": 100, "unit": "SievertPerHour"}`)
    sieverts_per_hourResult, err := factory.FromDtoJSON(sieverts_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SievertPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = sieverts_per_hourResult.Convert(units.RadiationEquivalentDoseRateSievertPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SievertPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with SievertPerSecond unit
    sieverts_per_secondJSON := []byte(`{"value": 100, "unit": "SievertPerSecond"}`)
    sieverts_per_secondResult, err := factory.FromDtoJSON(sieverts_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SievertPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = sieverts_per_secondResult.Convert(units.RadiationEquivalentDoseRateSievertPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SievertPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with RoentgenEquivalentManPerHour unit
    roentgens_equivalent_man_per_hourJSON := []byte(`{"value": 100, "unit": "RoentgenEquivalentManPerHour"}`)
    roentgens_equivalent_man_per_hourResult, err := factory.FromDtoJSON(roentgens_equivalent_man_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with RoentgenEquivalentManPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = roentgens_equivalent_man_per_hourResult.Convert(units.RadiationEquivalentDoseRateRoentgenEquivalentManPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RoentgenEquivalentManPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with NanosievertPerHour unit
    nanosieverts_per_hourJSON := []byte(`{"value": 100, "unit": "NanosievertPerHour"}`)
    nanosieverts_per_hourResult, err := factory.FromDtoJSON(nanosieverts_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanosievertPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanosieverts_per_hourResult.Convert(units.RadiationEquivalentDoseRateNanosievertPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanosievertPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MicrosievertPerHour unit
    microsieverts_per_hourJSON := []byte(`{"value": 100, "unit": "MicrosievertPerHour"}`)
    microsieverts_per_hourResult, err := factory.FromDtoJSON(microsieverts_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrosievertPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microsieverts_per_hourResult.Convert(units.RadiationEquivalentDoseRateMicrosievertPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrosievertPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MillisievertPerHour unit
    millisieverts_per_hourJSON := []byte(`{"value": 100, "unit": "MillisievertPerHour"}`)
    millisieverts_per_hourResult, err := factory.FromDtoJSON(millisieverts_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillisievertPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millisieverts_per_hourResult.Convert(units.RadiationEquivalentDoseRateMillisievertPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillisievertPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with NanosievertPerSecond unit
    nanosieverts_per_secondJSON := []byte(`{"value": 100, "unit": "NanosievertPerSecond"}`)
    nanosieverts_per_secondResult, err := factory.FromDtoJSON(nanosieverts_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanosievertPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanosieverts_per_secondResult.Convert(units.RadiationEquivalentDoseRateNanosievertPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanosievertPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicrosievertPerSecond unit
    microsieverts_per_secondJSON := []byte(`{"value": 100, "unit": "MicrosievertPerSecond"}`)
    microsieverts_per_secondResult, err := factory.FromDtoJSON(microsieverts_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrosievertPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microsieverts_per_secondResult.Convert(units.RadiationEquivalentDoseRateMicrosievertPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrosievertPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MillisievertPerSecond unit
    millisieverts_per_secondJSON := []byte(`{"value": 100, "unit": "MillisievertPerSecond"}`)
    millisieverts_per_secondResult, err := factory.FromDtoJSON(millisieverts_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillisievertPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millisieverts_per_secondResult.Convert(units.RadiationEquivalentDoseRateMillisievertPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillisievertPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MilliroentgenEquivalentManPerHour unit
    milliroentgens_equivalent_man_per_hourJSON := []byte(`{"value": 100, "unit": "MilliroentgenEquivalentManPerHour"}`)
    milliroentgens_equivalent_man_per_hourResult, err := factory.FromDtoJSON(milliroentgens_equivalent_man_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliroentgenEquivalentManPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliroentgens_equivalent_man_per_hourResult.Convert(units.RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliroentgenEquivalentManPerHour = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "SievertPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromSievertsPerHour function
func TestRadiationEquivalentDoseRateFactory_FromSievertsPerHour(t *testing.T) {
    factory := units.RadiationEquivalentDoseRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSievertsPerHour(100)
    if err != nil {
        t.Errorf("FromSievertsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseRateSievertPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSievertsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSievertsPerHour(math.NaN())
    if err == nil {
        t.Error("FromSievertsPerHour() with NaN value should return error")
    }

    _, err = factory.FromSievertsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromSievertsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromSievertsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromSievertsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSievertsPerHour(0)
    if err != nil {
        t.Errorf("FromSievertsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseRateSievertPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSievertsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromSievertsPerSecond function
func TestRadiationEquivalentDoseRateFactory_FromSievertsPerSecond(t *testing.T) {
    factory := units.RadiationEquivalentDoseRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSievertsPerSecond(100)
    if err != nil {
        t.Errorf("FromSievertsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseRateSievertPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSievertsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSievertsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromSievertsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromSievertsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromSievertsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromSievertsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromSievertsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSievertsPerSecond(0)
    if err != nil {
        t.Errorf("FromSievertsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseRateSievertPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSievertsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromRoentgensEquivalentManPerHour function
func TestRadiationEquivalentDoseRateFactory_FromRoentgensEquivalentManPerHour(t *testing.T) {
    factory := units.RadiationEquivalentDoseRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRoentgensEquivalentManPerHour(100)
    if err != nil {
        t.Errorf("FromRoentgensEquivalentManPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseRateRoentgenEquivalentManPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromRoentgensEquivalentManPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromRoentgensEquivalentManPerHour(math.NaN())
    if err == nil {
        t.Error("FromRoentgensEquivalentManPerHour() with NaN value should return error")
    }

    _, err = factory.FromRoentgensEquivalentManPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromRoentgensEquivalentManPerHour() with +Inf value should return error")
    }

    _, err = factory.FromRoentgensEquivalentManPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromRoentgensEquivalentManPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromRoentgensEquivalentManPerHour(0)
    if err != nil {
        t.Errorf("FromRoentgensEquivalentManPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseRateRoentgenEquivalentManPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRoentgensEquivalentManPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromNanosievertsPerHour function
func TestRadiationEquivalentDoseRateFactory_FromNanosievertsPerHour(t *testing.T) {
    factory := units.RadiationEquivalentDoseRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanosievertsPerHour(100)
    if err != nil {
        t.Errorf("FromNanosievertsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseRateNanosievertPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanosievertsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanosievertsPerHour(math.NaN())
    if err == nil {
        t.Error("FromNanosievertsPerHour() with NaN value should return error")
    }

    _, err = factory.FromNanosievertsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromNanosievertsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromNanosievertsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromNanosievertsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanosievertsPerHour(0)
    if err != nil {
        t.Errorf("FromNanosievertsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseRateNanosievertPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanosievertsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrosievertsPerHour function
func TestRadiationEquivalentDoseRateFactory_FromMicrosievertsPerHour(t *testing.T) {
    factory := units.RadiationEquivalentDoseRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrosievertsPerHour(100)
    if err != nil {
        t.Errorf("FromMicrosievertsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseRateMicrosievertPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrosievertsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrosievertsPerHour(math.NaN())
    if err == nil {
        t.Error("FromMicrosievertsPerHour() with NaN value should return error")
    }

    _, err = factory.FromMicrosievertsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMicrosievertsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMicrosievertsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrosievertsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrosievertsPerHour(0)
    if err != nil {
        t.Errorf("FromMicrosievertsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseRateMicrosievertPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrosievertsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromMillisievertsPerHour function
func TestRadiationEquivalentDoseRateFactory_FromMillisievertsPerHour(t *testing.T) {
    factory := units.RadiationEquivalentDoseRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillisievertsPerHour(100)
    if err != nil {
        t.Errorf("FromMillisievertsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseRateMillisievertPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillisievertsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillisievertsPerHour(math.NaN())
    if err == nil {
        t.Error("FromMillisievertsPerHour() with NaN value should return error")
    }

    _, err = factory.FromMillisievertsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMillisievertsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMillisievertsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMillisievertsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillisievertsPerHour(0)
    if err != nil {
        t.Errorf("FromMillisievertsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseRateMillisievertPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillisievertsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromNanosievertsPerSecond function
func TestRadiationEquivalentDoseRateFactory_FromNanosievertsPerSecond(t *testing.T) {
    factory := units.RadiationEquivalentDoseRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanosievertsPerSecond(100)
    if err != nil {
        t.Errorf("FromNanosievertsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseRateNanosievertPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanosievertsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanosievertsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromNanosievertsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromNanosievertsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromNanosievertsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromNanosievertsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromNanosievertsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanosievertsPerSecond(0)
    if err != nil {
        t.Errorf("FromNanosievertsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseRateNanosievertPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanosievertsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrosievertsPerSecond function
func TestRadiationEquivalentDoseRateFactory_FromMicrosievertsPerSecond(t *testing.T) {
    factory := units.RadiationEquivalentDoseRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrosievertsPerSecond(100)
    if err != nil {
        t.Errorf("FromMicrosievertsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseRateMicrosievertPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrosievertsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrosievertsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMicrosievertsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMicrosievertsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMicrosievertsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMicrosievertsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrosievertsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrosievertsPerSecond(0)
    if err != nil {
        t.Errorf("FromMicrosievertsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseRateMicrosievertPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrosievertsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMillisievertsPerSecond function
func TestRadiationEquivalentDoseRateFactory_FromMillisievertsPerSecond(t *testing.T) {
    factory := units.RadiationEquivalentDoseRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillisievertsPerSecond(100)
    if err != nil {
        t.Errorf("FromMillisievertsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseRateMillisievertPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillisievertsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillisievertsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMillisievertsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMillisievertsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMillisievertsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMillisievertsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMillisievertsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillisievertsPerSecond(0)
    if err != nil {
        t.Errorf("FromMillisievertsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseRateMillisievertPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillisievertsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliroentgensEquivalentManPerHour function
func TestRadiationEquivalentDoseRateFactory_FromMilliroentgensEquivalentManPerHour(t *testing.T) {
    factory := units.RadiationEquivalentDoseRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliroentgensEquivalentManPerHour(100)
    if err != nil {
        t.Errorf("FromMilliroentgensEquivalentManPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliroentgensEquivalentManPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliroentgensEquivalentManPerHour(math.NaN())
    if err == nil {
        t.Error("FromMilliroentgensEquivalentManPerHour() with NaN value should return error")
    }

    _, err = factory.FromMilliroentgensEquivalentManPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMilliroentgensEquivalentManPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMilliroentgensEquivalentManPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliroentgensEquivalentManPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliroentgensEquivalentManPerHour(0)
    if err != nil {
        t.Errorf("FromMilliroentgensEquivalentManPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliroentgensEquivalentManPerHour() with zero value = %v, want 0", converted)
    }
}

func TestRadiationEquivalentDoseRateToString(t *testing.T) {
	factory := units.RadiationEquivalentDoseRateFactory{}
	a, err := factory.CreateRadiationEquivalentDoseRate(45, units.RadiationEquivalentDoseRateSievertPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RadiationEquivalentDoseRateSievertPerSecond, 2)
	expected := "45.00 " + units.GetRadiationEquivalentDoseRateAbbreviation(units.RadiationEquivalentDoseRateSievertPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RadiationEquivalentDoseRateSievertPerSecond, -1)
	expected = "45 " + units.GetRadiationEquivalentDoseRateAbbreviation(units.RadiationEquivalentDoseRateSievertPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRadiationEquivalentDoseRate_EqualityAndComparison(t *testing.T) {
	factory := units.RadiationEquivalentDoseRateFactory{}
	a1, _ := factory.CreateRadiationEquivalentDoseRate(60, units.RadiationEquivalentDoseRateSievertPerSecond)
	a2, _ := factory.CreateRadiationEquivalentDoseRate(60, units.RadiationEquivalentDoseRateSievertPerSecond)
	a3, _ := factory.CreateRadiationEquivalentDoseRate(120, units.RadiationEquivalentDoseRateSievertPerSecond)

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

func TestRadiationEquivalentDoseRate_Arithmetic(t *testing.T) {
	factory := units.RadiationEquivalentDoseRateFactory{}
	a1, _ := factory.CreateRadiationEquivalentDoseRate(30, units.RadiationEquivalentDoseRateSievertPerSecond)
	a2, _ := factory.CreateRadiationEquivalentDoseRate(45, units.RadiationEquivalentDoseRateSievertPerSecond)

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


func TestGetRadiationEquivalentDoseRateAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.RadiationEquivalentDoseRateUnits
        want string
    }{
        {
            name: "SievertPerHour abbreviation",
            unit: units.RadiationEquivalentDoseRateSievertPerHour,
            want: "Sv/h",
        },
        {
            name: "SievertPerSecond abbreviation",
            unit: units.RadiationEquivalentDoseRateSievertPerSecond,
            want: "Sv/s",
        },
        {
            name: "RoentgenEquivalentManPerHour abbreviation",
            unit: units.RadiationEquivalentDoseRateRoentgenEquivalentManPerHour,
            want: "rem/h",
        },
        {
            name: "NanosievertPerHour abbreviation",
            unit: units.RadiationEquivalentDoseRateNanosievertPerHour,
            want: "nSv/h",
        },
        {
            name: "MicrosievertPerHour abbreviation",
            unit: units.RadiationEquivalentDoseRateMicrosievertPerHour,
            want: "μSv/h",
        },
        {
            name: "MillisievertPerHour abbreviation",
            unit: units.RadiationEquivalentDoseRateMillisievertPerHour,
            want: "mSv/h",
        },
        {
            name: "NanosievertPerSecond abbreviation",
            unit: units.RadiationEquivalentDoseRateNanosievertPerSecond,
            want: "nSv/s",
        },
        {
            name: "MicrosievertPerSecond abbreviation",
            unit: units.RadiationEquivalentDoseRateMicrosievertPerSecond,
            want: "μSv/s",
        },
        {
            name: "MillisievertPerSecond abbreviation",
            unit: units.RadiationEquivalentDoseRateMillisievertPerSecond,
            want: "mSv/s",
        },
        {
            name: "MilliroentgenEquivalentManPerHour abbreviation",
            unit: units.RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour,
            want: "mrem/h",
        },
        {
            name: "invalid unit",
            unit: units.RadiationEquivalentDoseRateUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetRadiationEquivalentDoseRateAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetRadiationEquivalentDoseRateAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestRadiationEquivalentDoseRate_String(t *testing.T) {
    factory := units.RadiationEquivalentDoseRateFactory{}
    
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
            unit, err := factory.CreateRadiationEquivalentDoseRate(tt.value, units.RadiationEquivalentDoseRateSievertPerSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("RadiationEquivalentDoseRate.String() = %v, want %v", got, tt.want)
            }
        })
    }
}