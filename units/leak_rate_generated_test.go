// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLeakRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "PascalCubicMeterPerSecond"}`
	
	factory := units.LeakRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LeakRatePascalCubicMeterPerSecond {
		t.Errorf("expected unit %v, got %v", units.LeakRatePascalCubicMeterPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "PascalCubicMeterPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLeakRateDto_ToJSON(t *testing.T) {
	dto := units.LeakRateDto{
		Value: 45,
		Unit:  units.LeakRatePascalCubicMeterPerSecond,
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
	if result["unit"].(string) != string(units.LeakRatePascalCubicMeterPerSecond) {
		t.Errorf("expected unit %s, got %v", units.LeakRatePascalCubicMeterPerSecond, result["unit"])
	}
}

func TestNewLeakRate_InvalidValue(t *testing.T) {
	factory := units.LeakRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLeakRate(math.NaN(), units.LeakRatePascalCubicMeterPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLeakRate(math.Inf(1), units.LeakRatePascalCubicMeterPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLeakRateConversions(t *testing.T) {
	factory := units.LeakRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLeakRate(180, units.LeakRatePascalCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to PascalCubicMetersPerSecond.
		// No expected conversion value provided for PascalCubicMetersPerSecond, verifying result is not NaN.
		result := a.PascalCubicMetersPerSecond()
		cacheResult := a.PascalCubicMetersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PascalCubicMetersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to MillibarLitersPerSecond.
		// No expected conversion value provided for MillibarLitersPerSecond, verifying result is not NaN.
		result := a.MillibarLitersPerSecond()
		cacheResult := a.MillibarLitersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillibarLitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to TorrLitersPerSecond.
		// No expected conversion value provided for TorrLitersPerSecond, verifying result is not NaN.
		result := a.TorrLitersPerSecond()
		cacheResult := a.TorrLitersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TorrLitersPerSecond returned NaN")
		}
	}
	{
		// Test conversion to AtmCubicCentimetersPerSecond.
		// No expected conversion value provided for AtmCubicCentimetersPerSecond, verifying result is not NaN.
		result := a.AtmCubicCentimetersPerSecond()
		cacheResult := a.AtmCubicCentimetersPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AtmCubicCentimetersPerSecond returned NaN")
		}
	}
}

func TestLeakRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LeakRateFactory{}
	a, err := factory.CreateLeakRate(90, units.LeakRatePascalCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LeakRatePascalCubicMeterPerSecond {
		t.Errorf("expected default unit PascalCubicMeterPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LeakRatePascalCubicMeterPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LeakRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LeakRatePascalCubicMeterPerSecond {
		t.Errorf("expected unit PascalCubicMeterPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLeakRateFactory_FromDto(t *testing.T) {
    factory := units.LeakRateFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.LeakRateDto{
        Value: 100,
        Unit:  units.LeakRatePascalCubicMeterPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.LeakRateDto{
        Value: math.NaN(),
        Unit:  units.LeakRatePascalCubicMeterPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test PascalCubicMeterPerSecond conversion
    pascal_cubic_meters_per_secondDto := units.LeakRateDto{
        Value: 100,
        Unit:  units.LeakRatePascalCubicMeterPerSecond,
    }
    
    var pascal_cubic_meters_per_secondResult *units.LeakRate
    pascal_cubic_meters_per_secondResult, err = factory.FromDto(pascal_cubic_meters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PascalCubicMeterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_cubic_meters_per_secondResult.Convert(units.LeakRatePascalCubicMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalCubicMeterPerSecond = %v, want %v", converted, 100)
    }
    // Test MillibarLiterPerSecond conversion
    millibar_liters_per_secondDto := units.LeakRateDto{
        Value: 100,
        Unit:  units.LeakRateMillibarLiterPerSecond,
    }
    
    var millibar_liters_per_secondResult *units.LeakRate
    millibar_liters_per_secondResult, err = factory.FromDto(millibar_liters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MillibarLiterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millibar_liters_per_secondResult.Convert(units.LeakRateMillibarLiterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillibarLiterPerSecond = %v, want %v", converted, 100)
    }
    // Test TorrLiterPerSecond conversion
    torr_liters_per_secondDto := units.LeakRateDto{
        Value: 100,
        Unit:  units.LeakRateTorrLiterPerSecond,
    }
    
    var torr_liters_per_secondResult *units.LeakRate
    torr_liters_per_secondResult, err = factory.FromDto(torr_liters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with TorrLiterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = torr_liters_per_secondResult.Convert(units.LeakRateTorrLiterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TorrLiterPerSecond = %v, want %v", converted, 100)
    }
    // Test AtmCubicCentimeterPerSecond conversion
    atm_cubic_centimeters_per_secondDto := units.LeakRateDto{
        Value: 100,
        Unit:  units.LeakRateAtmCubicCentimeterPerSecond,
    }
    
    var atm_cubic_centimeters_per_secondResult *units.LeakRate
    atm_cubic_centimeters_per_secondResult, err = factory.FromDto(atm_cubic_centimeters_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with AtmCubicCentimeterPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = atm_cubic_centimeters_per_secondResult.Convert(units.LeakRateAtmCubicCentimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AtmCubicCentimeterPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.LeakRateDto{
        Value: 0,
        Unit:  units.LeakRatePascalCubicMeterPerSecond,
    }
    
    var zeroResult *units.LeakRate
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestLeakRateFactory_FromDtoJSON(t *testing.T) {
    factory := units.LeakRateFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "PascalCubicMeterPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "PascalCubicMeterPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.LeakRateDto{
        Value: nanValue,
        Unit:  units.LeakRatePascalCubicMeterPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with PascalCubicMeterPerSecond unit
    pascal_cubic_meters_per_secondJSON := []byte(`{"value": 100, "unit": "PascalCubicMeterPerSecond"}`)
    pascal_cubic_meters_per_secondResult, err := factory.FromDtoJSON(pascal_cubic_meters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PascalCubicMeterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_cubic_meters_per_secondResult.Convert(units.LeakRatePascalCubicMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalCubicMeterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MillibarLiterPerSecond unit
    millibar_liters_per_secondJSON := []byte(`{"value": 100, "unit": "MillibarLiterPerSecond"}`)
    millibar_liters_per_secondResult, err := factory.FromDtoJSON(millibar_liters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillibarLiterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millibar_liters_per_secondResult.Convert(units.LeakRateMillibarLiterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillibarLiterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with TorrLiterPerSecond unit
    torr_liters_per_secondJSON := []byte(`{"value": 100, "unit": "TorrLiterPerSecond"}`)
    torr_liters_per_secondResult, err := factory.FromDtoJSON(torr_liters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TorrLiterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = torr_liters_per_secondResult.Convert(units.LeakRateTorrLiterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TorrLiterPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with AtmCubicCentimeterPerSecond unit
    atm_cubic_centimeters_per_secondJSON := []byte(`{"value": 100, "unit": "AtmCubicCentimeterPerSecond"}`)
    atm_cubic_centimeters_per_secondResult, err := factory.FromDtoJSON(atm_cubic_centimeters_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AtmCubicCentimeterPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = atm_cubic_centimeters_per_secondResult.Convert(units.LeakRateAtmCubicCentimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AtmCubicCentimeterPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "PascalCubicMeterPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromPascalCubicMetersPerSecond function
func TestLeakRateFactory_FromPascalCubicMetersPerSecond(t *testing.T) {
    factory := units.LeakRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPascalCubicMetersPerSecond(100)
    if err != nil {
        t.Errorf("FromPascalCubicMetersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LeakRatePascalCubicMeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPascalCubicMetersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPascalCubicMetersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPascalCubicMetersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPascalCubicMetersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPascalCubicMetersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPascalCubicMetersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPascalCubicMetersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPascalCubicMetersPerSecond(0)
    if err != nil {
        t.Errorf("FromPascalCubicMetersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LeakRatePascalCubicMeterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPascalCubicMetersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMillibarLitersPerSecond function
func TestLeakRateFactory_FromMillibarLitersPerSecond(t *testing.T) {
    factory := units.LeakRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillibarLitersPerSecond(100)
    if err != nil {
        t.Errorf("FromMillibarLitersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LeakRateMillibarLiterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillibarLitersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillibarLitersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMillibarLitersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMillibarLitersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMillibarLitersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMillibarLitersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMillibarLitersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillibarLitersPerSecond(0)
    if err != nil {
        t.Errorf("FromMillibarLitersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LeakRateMillibarLiterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillibarLitersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromTorrLitersPerSecond function
func TestLeakRateFactory_FromTorrLitersPerSecond(t *testing.T) {
    factory := units.LeakRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTorrLitersPerSecond(100)
    if err != nil {
        t.Errorf("FromTorrLitersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LeakRateTorrLiterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTorrLitersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTorrLitersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromTorrLitersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromTorrLitersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromTorrLitersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromTorrLitersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromTorrLitersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTorrLitersPerSecond(0)
    if err != nil {
        t.Errorf("FromTorrLitersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LeakRateTorrLiterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTorrLitersPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromAtmCubicCentimetersPerSecond function
func TestLeakRateFactory_FromAtmCubicCentimetersPerSecond(t *testing.T) {
    factory := units.LeakRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAtmCubicCentimetersPerSecond(100)
    if err != nil {
        t.Errorf("FromAtmCubicCentimetersPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LeakRateAtmCubicCentimeterPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAtmCubicCentimetersPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAtmCubicCentimetersPerSecond(math.NaN())
    if err == nil {
        t.Error("FromAtmCubicCentimetersPerSecond() with NaN value should return error")
    }

    _, err = factory.FromAtmCubicCentimetersPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromAtmCubicCentimetersPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromAtmCubicCentimetersPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromAtmCubicCentimetersPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAtmCubicCentimetersPerSecond(0)
    if err != nil {
        t.Errorf("FromAtmCubicCentimetersPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LeakRateAtmCubicCentimeterPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAtmCubicCentimetersPerSecond() with zero value = %v, want 0", converted)
    }
}

func TestLeakRateToString(t *testing.T) {
	factory := units.LeakRateFactory{}
	a, err := factory.CreateLeakRate(45, units.LeakRatePascalCubicMeterPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LeakRatePascalCubicMeterPerSecond, 2)
	expected := "45.00 " + units.GetLeakRateAbbreviation(units.LeakRatePascalCubicMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LeakRatePascalCubicMeterPerSecond, -1)
	expected = "45 " + units.GetLeakRateAbbreviation(units.LeakRatePascalCubicMeterPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLeakRate_EqualityAndComparison(t *testing.T) {
	factory := units.LeakRateFactory{}
	a1, _ := factory.CreateLeakRate(60, units.LeakRatePascalCubicMeterPerSecond)
	a2, _ := factory.CreateLeakRate(60, units.LeakRatePascalCubicMeterPerSecond)
	a3, _ := factory.CreateLeakRate(120, units.LeakRatePascalCubicMeterPerSecond)

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

func TestLeakRate_Arithmetic(t *testing.T) {
	factory := units.LeakRateFactory{}
	a1, _ := factory.CreateLeakRate(30, units.LeakRatePascalCubicMeterPerSecond)
	a2, _ := factory.CreateLeakRate(45, units.LeakRatePascalCubicMeterPerSecond)

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


func TestGetLeakRateAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.LeakRateUnits
        want string
    }{
        {
            name: "PascalCubicMeterPerSecond abbreviation",
            unit: units.LeakRatePascalCubicMeterPerSecond,
            want: "Pa·m³/s",
        },
        {
            name: "MillibarLiterPerSecond abbreviation",
            unit: units.LeakRateMillibarLiterPerSecond,
            want: "mbar·l/s",
        },
        {
            name: "TorrLiterPerSecond abbreviation",
            unit: units.LeakRateTorrLiterPerSecond,
            want: "Torr·l/s",
        },
        {
            name: "AtmCubicCentimeterPerSecond abbreviation",
            unit: units.LeakRateAtmCubicCentimeterPerSecond,
            want: "atm·cm³/s",
        },
        {
            name: "invalid unit",
            unit: units.LeakRateUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetLeakRateAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetLeakRateAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestLeakRate_String(t *testing.T) {
    factory := units.LeakRateFactory{}
    
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
            unit, err := factory.CreateLeakRate(tt.value, units.LeakRatePascalCubicMeterPerSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("LeakRate.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestLeakRate_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.LeakRateFactory{}

	_, err := uf.CreateLeakRate(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}