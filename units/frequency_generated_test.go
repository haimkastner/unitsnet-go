// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestFrequencyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Hertz"}`
	
	factory := units.FrequencyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.FrequencyHertz {
		t.Errorf("expected unit %v, got %v", units.FrequencyHertz, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Hertz"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestFrequencyDto_ToJSON(t *testing.T) {
	dto := units.FrequencyDto{
		Value: 45,
		Unit:  units.FrequencyHertz,
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
	if result["unit"].(string) != string(units.FrequencyHertz) {
		t.Errorf("expected unit %s, got %v", units.FrequencyHertz, result["unit"])
	}
}

func TestNewFrequency_InvalidValue(t *testing.T) {
	factory := units.FrequencyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateFrequency(math.NaN(), units.FrequencyHertz)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateFrequency(math.Inf(1), units.FrequencyHertz)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestFrequencyConversions(t *testing.T) {
	factory := units.FrequencyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateFrequency(180, units.FrequencyHertz)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Hertz.
		// No expected conversion value provided for Hertz, verifying result is not NaN.
		result := a.Hertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Hertz returned NaN")
		}
	}
	{
		// Test conversion to RadiansPerSecond.
		// No expected conversion value provided for RadiansPerSecond, verifying result is not NaN.
		result := a.RadiansPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to RadiansPerSecond returned NaN")
		}
	}
	{
		// Test conversion to CyclesPerMinute.
		// No expected conversion value provided for CyclesPerMinute, verifying result is not NaN.
		result := a.CyclesPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to CyclesPerMinute returned NaN")
		}
	}
	{
		// Test conversion to CyclesPerHour.
		// No expected conversion value provided for CyclesPerHour, verifying result is not NaN.
		result := a.CyclesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to CyclesPerHour returned NaN")
		}
	}
	{
		// Test conversion to BeatsPerMinute.
		// No expected conversion value provided for BeatsPerMinute, verifying result is not NaN.
		result := a.BeatsPerMinute()
		if math.IsNaN(result) {
			t.Errorf("conversion to BeatsPerMinute returned NaN")
		}
	}
	{
		// Test conversion to PerSecond.
		// No expected conversion value provided for PerSecond, verifying result is not NaN.
		result := a.PerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PerSecond returned NaN")
		}
	}
	{
		// Test conversion to BUnits.
		// No expected conversion value provided for BUnits, verifying result is not NaN.
		result := a.BUnits()
		if math.IsNaN(result) {
			t.Errorf("conversion to BUnits returned NaN")
		}
	}
	{
		// Test conversion to Microhertz.
		// No expected conversion value provided for Microhertz, verifying result is not NaN.
		result := a.Microhertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microhertz returned NaN")
		}
	}
	{
		// Test conversion to Millihertz.
		// No expected conversion value provided for Millihertz, verifying result is not NaN.
		result := a.Millihertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millihertz returned NaN")
		}
	}
	{
		// Test conversion to Kilohertz.
		// No expected conversion value provided for Kilohertz, verifying result is not NaN.
		result := a.Kilohertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilohertz returned NaN")
		}
	}
	{
		// Test conversion to Megahertz.
		// No expected conversion value provided for Megahertz, verifying result is not NaN.
		result := a.Megahertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megahertz returned NaN")
		}
	}
	{
		// Test conversion to Gigahertz.
		// No expected conversion value provided for Gigahertz, verifying result is not NaN.
		result := a.Gigahertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigahertz returned NaN")
		}
	}
	{
		// Test conversion to Terahertz.
		// No expected conversion value provided for Terahertz, verifying result is not NaN.
		result := a.Terahertz()
		if math.IsNaN(result) {
			t.Errorf("conversion to Terahertz returned NaN")
		}
	}
}

func TestFrequency_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.FrequencyFactory{}
	a, err := factory.CreateFrequency(90, units.FrequencyHertz)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.FrequencyHertz {
		t.Errorf("expected default unit Hertz, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.FrequencyHertz
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.FrequencyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.FrequencyHertz {
		t.Errorf("expected unit Hertz, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestFrequencyFactory_FromDto(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyHertz,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.FrequencyDto{
        Value: math.NaN(),
        Unit:  units.FrequencyHertz,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Hertz conversion
    hertzDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyHertz,
    }
    
    var hertzResult *units.Frequency
    hertzResult, err = factory.FromDto(hertzDto)
    if err != nil {
        t.Errorf("FromDto() with Hertz returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hertzResult.Convert(units.FrequencyHertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hertz = %v, want %v", converted, 100)
    }
    // Test RadianPerSecond conversion
    radians_per_secondDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyRadianPerSecond,
    }
    
    var radians_per_secondResult *units.Frequency
    radians_per_secondResult, err = factory.FromDto(radians_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with RadianPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = radians_per_secondResult.Convert(units.FrequencyRadianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RadianPerSecond = %v, want %v", converted, 100)
    }
    // Test CyclePerMinute conversion
    cycles_per_minuteDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyCyclePerMinute,
    }
    
    var cycles_per_minuteResult *units.Frequency
    cycles_per_minuteResult, err = factory.FromDto(cycles_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with CyclePerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cycles_per_minuteResult.Convert(units.FrequencyCyclePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CyclePerMinute = %v, want %v", converted, 100)
    }
    // Test CyclePerHour conversion
    cycles_per_hourDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyCyclePerHour,
    }
    
    var cycles_per_hourResult *units.Frequency
    cycles_per_hourResult, err = factory.FromDto(cycles_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with CyclePerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cycles_per_hourResult.Convert(units.FrequencyCyclePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CyclePerHour = %v, want %v", converted, 100)
    }
    // Test BeatPerMinute conversion
    beats_per_minuteDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyBeatPerMinute,
    }
    
    var beats_per_minuteResult *units.Frequency
    beats_per_minuteResult, err = factory.FromDto(beats_per_minuteDto)
    if err != nil {
        t.Errorf("FromDto() with BeatPerMinute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = beats_per_minuteResult.Convert(units.FrequencyBeatPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BeatPerMinute = %v, want %v", converted, 100)
    }
    // Test PerSecond conversion
    per_secondDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyPerSecond,
    }
    
    var per_secondResult *units.Frequency
    per_secondResult, err = factory.FromDto(per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = per_secondResult.Convert(units.FrequencyPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PerSecond = %v, want %v", converted, 100)
    }
    // Test BUnit conversion
    b_unitsDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyBUnit,
    }
    
    var b_unitsResult *units.Frequency
    b_unitsResult, err = factory.FromDto(b_unitsDto)
    if err != nil {
        t.Errorf("FromDto() with BUnit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = b_unitsResult.Convert(units.FrequencyBUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BUnit = %v, want %v", converted, 100)
    }
    // Test Microhertz conversion
    microhertzDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyMicrohertz,
    }
    
    var microhertzResult *units.Frequency
    microhertzResult, err = factory.FromDto(microhertzDto)
    if err != nil {
        t.Errorf("FromDto() with Microhertz returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microhertzResult.Convert(units.FrequencyMicrohertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microhertz = %v, want %v", converted, 100)
    }
    // Test Millihertz conversion
    millihertzDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyMillihertz,
    }
    
    var millihertzResult *units.Frequency
    millihertzResult, err = factory.FromDto(millihertzDto)
    if err != nil {
        t.Errorf("FromDto() with Millihertz returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millihertzResult.Convert(units.FrequencyMillihertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millihertz = %v, want %v", converted, 100)
    }
    // Test Kilohertz conversion
    kilohertzDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyKilohertz,
    }
    
    var kilohertzResult *units.Frequency
    kilohertzResult, err = factory.FromDto(kilohertzDto)
    if err != nil {
        t.Errorf("FromDto() with Kilohertz returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilohertzResult.Convert(units.FrequencyKilohertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilohertz = %v, want %v", converted, 100)
    }
    // Test Megahertz conversion
    megahertzDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyMegahertz,
    }
    
    var megahertzResult *units.Frequency
    megahertzResult, err = factory.FromDto(megahertzDto)
    if err != nil {
        t.Errorf("FromDto() with Megahertz returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megahertzResult.Convert(units.FrequencyMegahertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megahertz = %v, want %v", converted, 100)
    }
    // Test Gigahertz conversion
    gigahertzDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyGigahertz,
    }
    
    var gigahertzResult *units.Frequency
    gigahertzResult, err = factory.FromDto(gigahertzDto)
    if err != nil {
        t.Errorf("FromDto() with Gigahertz returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigahertzResult.Convert(units.FrequencyGigahertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigahertz = %v, want %v", converted, 100)
    }
    // Test Terahertz conversion
    terahertzDto := units.FrequencyDto{
        Value: 100,
        Unit:  units.FrequencyTerahertz,
    }
    
    var terahertzResult *units.Frequency
    terahertzResult, err = factory.FromDto(terahertzDto)
    if err != nil {
        t.Errorf("FromDto() with Terahertz returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terahertzResult.Convert(units.FrequencyTerahertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terahertz = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.FrequencyDto{
        Value: 0,
        Unit:  units.FrequencyHertz,
    }
    
    var zeroResult *units.Frequency
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestFrequencyFactory_FromDtoJSON(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Hertz"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Hertz"}`)
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
    nanJSON, _ := json.Marshal(units.FrequencyDto{
        Value: nanValue,
        Unit:  units.FrequencyHertz,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Hertz unit
    hertzJSON := []byte(`{"value": 100, "unit": "Hertz"}`)
    hertzResult, err := factory.FromDtoJSON(hertzJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Hertz unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hertzResult.Convert(units.FrequencyHertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hertz = %v, want %v", converted, 100)
    }
    // Test JSON with RadianPerSecond unit
    radians_per_secondJSON := []byte(`{"value": 100, "unit": "RadianPerSecond"}`)
    radians_per_secondResult, err := factory.FromDtoJSON(radians_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with RadianPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = radians_per_secondResult.Convert(units.FrequencyRadianPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for RadianPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with CyclePerMinute unit
    cycles_per_minuteJSON := []byte(`{"value": 100, "unit": "CyclePerMinute"}`)
    cycles_per_minuteResult, err := factory.FromDtoJSON(cycles_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CyclePerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cycles_per_minuteResult.Convert(units.FrequencyCyclePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CyclePerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with CyclePerHour unit
    cycles_per_hourJSON := []byte(`{"value": 100, "unit": "CyclePerHour"}`)
    cycles_per_hourResult, err := factory.FromDtoJSON(cycles_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CyclePerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cycles_per_hourResult.Convert(units.FrequencyCyclePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CyclePerHour = %v, want %v", converted, 100)
    }
    // Test JSON with BeatPerMinute unit
    beats_per_minuteJSON := []byte(`{"value": 100, "unit": "BeatPerMinute"}`)
    beats_per_minuteResult, err := factory.FromDtoJSON(beats_per_minuteJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BeatPerMinute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = beats_per_minuteResult.Convert(units.FrequencyBeatPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BeatPerMinute = %v, want %v", converted, 100)
    }
    // Test JSON with PerSecond unit
    per_secondJSON := []byte(`{"value": 100, "unit": "PerSecond"}`)
    per_secondResult, err := factory.FromDtoJSON(per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = per_secondResult.Convert(units.FrequencyPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with BUnit unit
    b_unitsJSON := []byte(`{"value": 100, "unit": "BUnit"}`)
    b_unitsResult, err := factory.FromDtoJSON(b_unitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BUnit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = b_unitsResult.Convert(units.FrequencyBUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BUnit = %v, want %v", converted, 100)
    }
    // Test JSON with Microhertz unit
    microhertzJSON := []byte(`{"value": 100, "unit": "Microhertz"}`)
    microhertzResult, err := factory.FromDtoJSON(microhertzJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microhertz unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microhertzResult.Convert(units.FrequencyMicrohertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microhertz = %v, want %v", converted, 100)
    }
    // Test JSON with Millihertz unit
    millihertzJSON := []byte(`{"value": 100, "unit": "Millihertz"}`)
    millihertzResult, err := factory.FromDtoJSON(millihertzJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millihertz unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millihertzResult.Convert(units.FrequencyMillihertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millihertz = %v, want %v", converted, 100)
    }
    // Test JSON with Kilohertz unit
    kilohertzJSON := []byte(`{"value": 100, "unit": "Kilohertz"}`)
    kilohertzResult, err := factory.FromDtoJSON(kilohertzJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilohertz unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilohertzResult.Convert(units.FrequencyKilohertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilohertz = %v, want %v", converted, 100)
    }
    // Test JSON with Megahertz unit
    megahertzJSON := []byte(`{"value": 100, "unit": "Megahertz"}`)
    megahertzResult, err := factory.FromDtoJSON(megahertzJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megahertz unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megahertzResult.Convert(units.FrequencyMegahertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megahertz = %v, want %v", converted, 100)
    }
    // Test JSON with Gigahertz unit
    gigahertzJSON := []byte(`{"value": 100, "unit": "Gigahertz"}`)
    gigahertzResult, err := factory.FromDtoJSON(gigahertzJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigahertz unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigahertzResult.Convert(units.FrequencyGigahertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigahertz = %v, want %v", converted, 100)
    }
    // Test JSON with Terahertz unit
    terahertzJSON := []byte(`{"value": 100, "unit": "Terahertz"}`)
    terahertzResult, err := factory.FromDtoJSON(terahertzJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Terahertz unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terahertzResult.Convert(units.FrequencyTerahertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terahertz = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Hertz"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromHertz function
func TestFrequencyFactory_FromHertz(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHertz(100)
    if err != nil {
        t.Errorf("FromHertz() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FrequencyHertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHertz() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHertz(math.NaN())
    if err == nil {
        t.Error("FromHertz() with NaN value should return error")
    }

    _, err = factory.FromHertz(math.Inf(1))
    if err == nil {
        t.Error("FromHertz() with +Inf value should return error")
    }

    _, err = factory.FromHertz(math.Inf(-1))
    if err == nil {
        t.Error("FromHertz() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHertz(0)
    if err != nil {
        t.Errorf("FromHertz() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FrequencyHertz)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHertz() with zero value = %v, want 0", converted)
    }
}
// Test FromRadiansPerSecond function
func TestFrequencyFactory_FromRadiansPerSecond(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRadiansPerSecond(100)
    if err != nil {
        t.Errorf("FromRadiansPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FrequencyRadianPerSecond)
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
    converted = zeroResult.Convert(units.FrequencyRadianPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRadiansPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromCyclesPerMinute function
func TestFrequencyFactory_FromCyclesPerMinute(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCyclesPerMinute(100)
    if err != nil {
        t.Errorf("FromCyclesPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FrequencyCyclePerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCyclesPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCyclesPerMinute(math.NaN())
    if err == nil {
        t.Error("FromCyclesPerMinute() with NaN value should return error")
    }

    _, err = factory.FromCyclesPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromCyclesPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromCyclesPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromCyclesPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCyclesPerMinute(0)
    if err != nil {
        t.Errorf("FromCyclesPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FrequencyCyclePerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCyclesPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromCyclesPerHour function
func TestFrequencyFactory_FromCyclesPerHour(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCyclesPerHour(100)
    if err != nil {
        t.Errorf("FromCyclesPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FrequencyCyclePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCyclesPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCyclesPerHour(math.NaN())
    if err == nil {
        t.Error("FromCyclesPerHour() with NaN value should return error")
    }

    _, err = factory.FromCyclesPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromCyclesPerHour() with +Inf value should return error")
    }

    _, err = factory.FromCyclesPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromCyclesPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCyclesPerHour(0)
    if err != nil {
        t.Errorf("FromCyclesPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FrequencyCyclePerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCyclesPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromBeatsPerMinute function
func TestFrequencyFactory_FromBeatsPerMinute(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBeatsPerMinute(100)
    if err != nil {
        t.Errorf("FromBeatsPerMinute() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FrequencyBeatPerMinute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBeatsPerMinute() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBeatsPerMinute(math.NaN())
    if err == nil {
        t.Error("FromBeatsPerMinute() with NaN value should return error")
    }

    _, err = factory.FromBeatsPerMinute(math.Inf(1))
    if err == nil {
        t.Error("FromBeatsPerMinute() with +Inf value should return error")
    }

    _, err = factory.FromBeatsPerMinute(math.Inf(-1))
    if err == nil {
        t.Error("FromBeatsPerMinute() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBeatsPerMinute(0)
    if err != nil {
        t.Errorf("FromBeatsPerMinute() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FrequencyBeatPerMinute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBeatsPerMinute() with zero value = %v, want 0", converted)
    }
}
// Test FromPerSecond function
func TestFrequencyFactory_FromPerSecond(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPerSecond(100)
    if err != nil {
        t.Errorf("FromPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FrequencyPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPerSecond(0)
    if err != nil {
        t.Errorf("FromPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FrequencyPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromBUnits function
func TestFrequencyFactory_FromBUnits(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBUnits(100)
    if err != nil {
        t.Errorf("FromBUnits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FrequencyBUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBUnits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBUnits(math.NaN())
    if err == nil {
        t.Error("FromBUnits() with NaN value should return error")
    }

    _, err = factory.FromBUnits(math.Inf(1))
    if err == nil {
        t.Error("FromBUnits() with +Inf value should return error")
    }

    _, err = factory.FromBUnits(math.Inf(-1))
    if err == nil {
        t.Error("FromBUnits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBUnits(0)
    if err != nil {
        t.Errorf("FromBUnits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FrequencyBUnit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBUnits() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrohertz function
func TestFrequencyFactory_FromMicrohertz(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrohertz(100)
    if err != nil {
        t.Errorf("FromMicrohertz() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FrequencyMicrohertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrohertz() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrohertz(math.NaN())
    if err == nil {
        t.Error("FromMicrohertz() with NaN value should return error")
    }

    _, err = factory.FromMicrohertz(math.Inf(1))
    if err == nil {
        t.Error("FromMicrohertz() with +Inf value should return error")
    }

    _, err = factory.FromMicrohertz(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrohertz() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrohertz(0)
    if err != nil {
        t.Errorf("FromMicrohertz() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FrequencyMicrohertz)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrohertz() with zero value = %v, want 0", converted)
    }
}
// Test FromMillihertz function
func TestFrequencyFactory_FromMillihertz(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillihertz(100)
    if err != nil {
        t.Errorf("FromMillihertz() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FrequencyMillihertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillihertz() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillihertz(math.NaN())
    if err == nil {
        t.Error("FromMillihertz() with NaN value should return error")
    }

    _, err = factory.FromMillihertz(math.Inf(1))
    if err == nil {
        t.Error("FromMillihertz() with +Inf value should return error")
    }

    _, err = factory.FromMillihertz(math.Inf(-1))
    if err == nil {
        t.Error("FromMillihertz() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillihertz(0)
    if err != nil {
        t.Errorf("FromMillihertz() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FrequencyMillihertz)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillihertz() with zero value = %v, want 0", converted)
    }
}
// Test FromKilohertz function
func TestFrequencyFactory_FromKilohertz(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilohertz(100)
    if err != nil {
        t.Errorf("FromKilohertz() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FrequencyKilohertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilohertz() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilohertz(math.NaN())
    if err == nil {
        t.Error("FromKilohertz() with NaN value should return error")
    }

    _, err = factory.FromKilohertz(math.Inf(1))
    if err == nil {
        t.Error("FromKilohertz() with +Inf value should return error")
    }

    _, err = factory.FromKilohertz(math.Inf(-1))
    if err == nil {
        t.Error("FromKilohertz() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilohertz(0)
    if err != nil {
        t.Errorf("FromKilohertz() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FrequencyKilohertz)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilohertz() with zero value = %v, want 0", converted)
    }
}
// Test FromMegahertz function
func TestFrequencyFactory_FromMegahertz(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegahertz(100)
    if err != nil {
        t.Errorf("FromMegahertz() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FrequencyMegahertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegahertz() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegahertz(math.NaN())
    if err == nil {
        t.Error("FromMegahertz() with NaN value should return error")
    }

    _, err = factory.FromMegahertz(math.Inf(1))
    if err == nil {
        t.Error("FromMegahertz() with +Inf value should return error")
    }

    _, err = factory.FromMegahertz(math.Inf(-1))
    if err == nil {
        t.Error("FromMegahertz() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegahertz(0)
    if err != nil {
        t.Errorf("FromMegahertz() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FrequencyMegahertz)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegahertz() with zero value = %v, want 0", converted)
    }
}
// Test FromGigahertz function
func TestFrequencyFactory_FromGigahertz(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigahertz(100)
    if err != nil {
        t.Errorf("FromGigahertz() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FrequencyGigahertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigahertz() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigahertz(math.NaN())
    if err == nil {
        t.Error("FromGigahertz() with NaN value should return error")
    }

    _, err = factory.FromGigahertz(math.Inf(1))
    if err == nil {
        t.Error("FromGigahertz() with +Inf value should return error")
    }

    _, err = factory.FromGigahertz(math.Inf(-1))
    if err == nil {
        t.Error("FromGigahertz() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigahertz(0)
    if err != nil {
        t.Errorf("FromGigahertz() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FrequencyGigahertz)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigahertz() with zero value = %v, want 0", converted)
    }
}
// Test FromTerahertz function
func TestFrequencyFactory_FromTerahertz(t *testing.T) {
    factory := units.FrequencyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerahertz(100)
    if err != nil {
        t.Errorf("FromTerahertz() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.FrequencyTerahertz)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerahertz() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerahertz(math.NaN())
    if err == nil {
        t.Error("FromTerahertz() with NaN value should return error")
    }

    _, err = factory.FromTerahertz(math.Inf(1))
    if err == nil {
        t.Error("FromTerahertz() with +Inf value should return error")
    }

    _, err = factory.FromTerahertz(math.Inf(-1))
    if err == nil {
        t.Error("FromTerahertz() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerahertz(0)
    if err != nil {
        t.Errorf("FromTerahertz() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.FrequencyTerahertz)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerahertz() with zero value = %v, want 0", converted)
    }
}

func TestFrequencyToString(t *testing.T) {
	factory := units.FrequencyFactory{}
	a, err := factory.CreateFrequency(45, units.FrequencyHertz)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.FrequencyHertz, 2)
	expected := "45.00 " + units.GetFrequencyAbbreviation(units.FrequencyHertz)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.FrequencyHertz, -1)
	expected = "45 " + units.GetFrequencyAbbreviation(units.FrequencyHertz)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestFrequency_EqualityAndComparison(t *testing.T) {
	factory := units.FrequencyFactory{}
	a1, _ := factory.CreateFrequency(60, units.FrequencyHertz)
	a2, _ := factory.CreateFrequency(60, units.FrequencyHertz)
	a3, _ := factory.CreateFrequency(120, units.FrequencyHertz)

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

func TestFrequency_Arithmetic(t *testing.T) {
	factory := units.FrequencyFactory{}
	a1, _ := factory.CreateFrequency(30, units.FrequencyHertz)
	a2, _ := factory.CreateFrequency(45, units.FrequencyHertz)

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