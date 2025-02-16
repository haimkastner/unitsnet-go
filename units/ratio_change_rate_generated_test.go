// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRatioChangeRateDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "DecimalFractionPerSecond"}`
	
	factory := units.RatioChangeRateDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RatioChangeRateDecimalFractionPerSecond {
		t.Errorf("expected unit %v, got %v", units.RatioChangeRateDecimalFractionPerSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "DecimalFractionPerSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRatioChangeRateDto_ToJSON(t *testing.T) {
	dto := units.RatioChangeRateDto{
		Value: 45,
		Unit:  units.RatioChangeRateDecimalFractionPerSecond,
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
	if result["unit"].(string) != string(units.RatioChangeRateDecimalFractionPerSecond) {
		t.Errorf("expected unit %s, got %v", units.RatioChangeRateDecimalFractionPerSecond, result["unit"])
	}
}

func TestNewRatioChangeRate_InvalidValue(t *testing.T) {
	factory := units.RatioChangeRateFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRatioChangeRate(math.NaN(), units.RatioChangeRateDecimalFractionPerSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRatioChangeRate(math.Inf(1), units.RatioChangeRateDecimalFractionPerSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRatioChangeRateConversions(t *testing.T) {
	factory := units.RatioChangeRateFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRatioChangeRate(180, units.RatioChangeRateDecimalFractionPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to PercentsPerSecond.
		// No expected conversion value provided for PercentsPerSecond, verifying result is not NaN.
		result := a.PercentsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PercentsPerSecond returned NaN")
		}
	}
	{
		// Test conversion to DecimalFractionsPerSecond.
		// No expected conversion value provided for DecimalFractionsPerSecond, verifying result is not NaN.
		result := a.DecimalFractionsPerSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecimalFractionsPerSecond returned NaN")
		}
	}
}

func TestRatioChangeRate_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RatioChangeRateFactory{}
	a, err := factory.CreateRatioChangeRate(90, units.RatioChangeRateDecimalFractionPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RatioChangeRateDecimalFractionPerSecond {
		t.Errorf("expected default unit DecimalFractionPerSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RatioChangeRatePercentPerSecond
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RatioChangeRateDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RatioChangeRateDecimalFractionPerSecond {
		t.Errorf("expected unit DecimalFractionPerSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRatioChangeRateFactory_FromDto(t *testing.T) {
    factory := units.RatioChangeRateFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.RatioChangeRateDto{
        Value: 100,
        Unit:  units.RatioChangeRateDecimalFractionPerSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.RatioChangeRateDto{
        Value: math.NaN(),
        Unit:  units.RatioChangeRateDecimalFractionPerSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test PercentPerSecond conversion
    percents_per_secondDto := units.RatioChangeRateDto{
        Value: 100,
        Unit:  units.RatioChangeRatePercentPerSecond,
    }
    
    var percents_per_secondResult *units.RatioChangeRate
    percents_per_secondResult, err = factory.FromDto(percents_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PercentPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = percents_per_secondResult.Convert(units.RatioChangeRatePercentPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PercentPerSecond = %v, want %v", converted, 100)
    }
    // Test DecimalFractionPerSecond conversion
    decimal_fractions_per_secondDto := units.RatioChangeRateDto{
        Value: 100,
        Unit:  units.RatioChangeRateDecimalFractionPerSecond,
    }
    
    var decimal_fractions_per_secondResult *units.RatioChangeRate
    decimal_fractions_per_secondResult, err = factory.FromDto(decimal_fractions_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with DecimalFractionPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimal_fractions_per_secondResult.Convert(units.RatioChangeRateDecimalFractionPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimalFractionPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.RatioChangeRateDto{
        Value: 0,
        Unit:  units.RatioChangeRateDecimalFractionPerSecond,
    }
    
    var zeroResult *units.RatioChangeRate
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestRatioChangeRateFactory_FromDtoJSON(t *testing.T) {
    factory := units.RatioChangeRateFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "DecimalFractionPerSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "DecimalFractionPerSecond"}`)
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
    nanJSON, _ := json.Marshal(units.RatioChangeRateDto{
        Value: nanValue,
        Unit:  units.RatioChangeRateDecimalFractionPerSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with PercentPerSecond unit
    percents_per_secondJSON := []byte(`{"value": 100, "unit": "PercentPerSecond"}`)
    percents_per_secondResult, err := factory.FromDtoJSON(percents_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PercentPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = percents_per_secondResult.Convert(units.RatioChangeRatePercentPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PercentPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with DecimalFractionPerSecond unit
    decimal_fractions_per_secondJSON := []byte(`{"value": 100, "unit": "DecimalFractionPerSecond"}`)
    decimal_fractions_per_secondResult, err := factory.FromDtoJSON(decimal_fractions_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecimalFractionPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimal_fractions_per_secondResult.Convert(units.RatioChangeRateDecimalFractionPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimalFractionPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "DecimalFractionPerSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromPercentsPerSecond function
func TestRatioChangeRateFactory_FromPercentsPerSecond(t *testing.T) {
    factory := units.RatioChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPercentsPerSecond(100)
    if err != nil {
        t.Errorf("FromPercentsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RatioChangeRatePercentPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPercentsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPercentsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromPercentsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromPercentsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPercentsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromPercentsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPercentsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPercentsPerSecond(0)
    if err != nil {
        t.Errorf("FromPercentsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RatioChangeRatePercentPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPercentsPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromDecimalFractionsPerSecond function
func TestRatioChangeRateFactory_FromDecimalFractionsPerSecond(t *testing.T) {
    factory := units.RatioChangeRateFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecimalFractionsPerSecond(100)
    if err != nil {
        t.Errorf("FromDecimalFractionsPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RatioChangeRateDecimalFractionPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecimalFractionsPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecimalFractionsPerSecond(math.NaN())
    if err == nil {
        t.Error("FromDecimalFractionsPerSecond() with NaN value should return error")
    }

    _, err = factory.FromDecimalFractionsPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromDecimalFractionsPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromDecimalFractionsPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromDecimalFractionsPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecimalFractionsPerSecond(0)
    if err != nil {
        t.Errorf("FromDecimalFractionsPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RatioChangeRateDecimalFractionPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecimalFractionsPerSecond() with zero value = %v, want 0", converted)
    }
}

func TestRatioChangeRateToString(t *testing.T) {
	factory := units.RatioChangeRateFactory{}
	a, err := factory.CreateRatioChangeRate(45, units.RatioChangeRateDecimalFractionPerSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RatioChangeRateDecimalFractionPerSecond, 2)
	expected := "45.00 " + units.GetRatioChangeRateAbbreviation(units.RatioChangeRateDecimalFractionPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RatioChangeRateDecimalFractionPerSecond, -1)
	expected = "45 " + units.GetRatioChangeRateAbbreviation(units.RatioChangeRateDecimalFractionPerSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRatioChangeRate_EqualityAndComparison(t *testing.T) {
	factory := units.RatioChangeRateFactory{}
	a1, _ := factory.CreateRatioChangeRate(60, units.RatioChangeRateDecimalFractionPerSecond)
	a2, _ := factory.CreateRatioChangeRate(60, units.RatioChangeRateDecimalFractionPerSecond)
	a3, _ := factory.CreateRatioChangeRate(120, units.RatioChangeRateDecimalFractionPerSecond)

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

func TestRatioChangeRate_Arithmetic(t *testing.T) {
	factory := units.RatioChangeRateFactory{}
	a1, _ := factory.CreateRatioChangeRate(30, units.RatioChangeRateDecimalFractionPerSecond)
	a2, _ := factory.CreateRatioChangeRate(45, units.RatioChangeRateDecimalFractionPerSecond)

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