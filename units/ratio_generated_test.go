// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestRatioDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "DecimalFraction"}`
	
	factory := units.RatioDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.RatioDecimalFraction {
		t.Errorf("expected unit %v, got %v", units.RatioDecimalFraction, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "DecimalFraction"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestRatioDto_ToJSON(t *testing.T) {
	dto := units.RatioDto{
		Value: 45,
		Unit:  units.RatioDecimalFraction,
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
	if result["unit"].(string) != string(units.RatioDecimalFraction) {
		t.Errorf("expected unit %s, got %v", units.RatioDecimalFraction, result["unit"])
	}
}

func TestNewRatio_InvalidValue(t *testing.T) {
	factory := units.RatioFactory{}
	// NaN value should return an error.
	_, err := factory.CreateRatio(math.NaN(), units.RatioDecimalFraction)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateRatio(math.Inf(1), units.RatioDecimalFraction)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestRatioConversions(t *testing.T) {
	factory := units.RatioFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateRatio(180, units.RatioDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to DecimalFractions.
		// No expected conversion value provided for DecimalFractions, verifying result is not NaN.
		result := a.DecimalFractions()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecimalFractions returned NaN")
		}
	}
	{
		// Test conversion to Percent.
		// No expected conversion value provided for Percent, verifying result is not NaN.
		result := a.Percent()
		if math.IsNaN(result) {
			t.Errorf("conversion to Percent returned NaN")
		}
	}
	{
		// Test conversion to PartsPerThousand.
		// No expected conversion value provided for PartsPerThousand, verifying result is not NaN.
		result := a.PartsPerThousand()
		if math.IsNaN(result) {
			t.Errorf("conversion to PartsPerThousand returned NaN")
		}
	}
	{
		// Test conversion to PartsPerMillion.
		// No expected conversion value provided for PartsPerMillion, verifying result is not NaN.
		result := a.PartsPerMillion()
		if math.IsNaN(result) {
			t.Errorf("conversion to PartsPerMillion returned NaN")
		}
	}
	{
		// Test conversion to PartsPerBillion.
		// No expected conversion value provided for PartsPerBillion, verifying result is not NaN.
		result := a.PartsPerBillion()
		if math.IsNaN(result) {
			t.Errorf("conversion to PartsPerBillion returned NaN")
		}
	}
	{
		// Test conversion to PartsPerTrillion.
		// No expected conversion value provided for PartsPerTrillion, verifying result is not NaN.
		result := a.PartsPerTrillion()
		if math.IsNaN(result) {
			t.Errorf("conversion to PartsPerTrillion returned NaN")
		}
	}
}

func TestRatio_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.RatioFactory{}
	a, err := factory.CreateRatio(90, units.RatioDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.RatioDecimalFraction {
		t.Errorf("expected default unit DecimalFraction, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.RatioDecimalFraction
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.RatioDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.RatioDecimalFraction {
		t.Errorf("expected unit DecimalFraction, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestRatioFactory_FromDto(t *testing.T) {
    factory := units.RatioFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.RatioDto{
        Value: 100,
        Unit:  units.RatioDecimalFraction,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.RatioDto{
        Value: math.NaN(),
        Unit:  units.RatioDecimalFraction,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test DecimalFraction conversion
    decimal_fractionsDto := units.RatioDto{
        Value: 100,
        Unit:  units.RatioDecimalFraction,
    }
    
    var decimal_fractionsResult *units.Ratio
    decimal_fractionsResult, err = factory.FromDto(decimal_fractionsDto)
    if err != nil {
        t.Errorf("FromDto() with DecimalFraction returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimal_fractionsResult.Convert(units.RatioDecimalFraction)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimalFraction = %v, want %v", converted, 100)
    }
    // Test Percent conversion
    percentDto := units.RatioDto{
        Value: 100,
        Unit:  units.RatioPercent,
    }
    
    var percentResult *units.Ratio
    percentResult, err = factory.FromDto(percentDto)
    if err != nil {
        t.Errorf("FromDto() with Percent returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = percentResult.Convert(units.RatioPercent)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Percent = %v, want %v", converted, 100)
    }
    // Test PartPerThousand conversion
    parts_per_thousandDto := units.RatioDto{
        Value: 100,
        Unit:  units.RatioPartPerThousand,
    }
    
    var parts_per_thousandResult *units.Ratio
    parts_per_thousandResult, err = factory.FromDto(parts_per_thousandDto)
    if err != nil {
        t.Errorf("FromDto() with PartPerThousand returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_thousandResult.Convert(units.RatioPartPerThousand)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerThousand = %v, want %v", converted, 100)
    }
    // Test PartPerMillion conversion
    parts_per_millionDto := units.RatioDto{
        Value: 100,
        Unit:  units.RatioPartPerMillion,
    }
    
    var parts_per_millionResult *units.Ratio
    parts_per_millionResult, err = factory.FromDto(parts_per_millionDto)
    if err != nil {
        t.Errorf("FromDto() with PartPerMillion returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_millionResult.Convert(units.RatioPartPerMillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerMillion = %v, want %v", converted, 100)
    }
    // Test PartPerBillion conversion
    parts_per_billionDto := units.RatioDto{
        Value: 100,
        Unit:  units.RatioPartPerBillion,
    }
    
    var parts_per_billionResult *units.Ratio
    parts_per_billionResult, err = factory.FromDto(parts_per_billionDto)
    if err != nil {
        t.Errorf("FromDto() with PartPerBillion returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_billionResult.Convert(units.RatioPartPerBillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerBillion = %v, want %v", converted, 100)
    }
    // Test PartPerTrillion conversion
    parts_per_trillionDto := units.RatioDto{
        Value: 100,
        Unit:  units.RatioPartPerTrillion,
    }
    
    var parts_per_trillionResult *units.Ratio
    parts_per_trillionResult, err = factory.FromDto(parts_per_trillionDto)
    if err != nil {
        t.Errorf("FromDto() with PartPerTrillion returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_trillionResult.Convert(units.RatioPartPerTrillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerTrillion = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.RatioDto{
        Value: 0,
        Unit:  units.RatioDecimalFraction,
    }
    
    var zeroResult *units.Ratio
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestRatioFactory_FromDtoJSON(t *testing.T) {
    factory := units.RatioFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "DecimalFraction"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "DecimalFraction"}`)
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
    nanJSON, _ := json.Marshal(units.RatioDto{
        Value: nanValue,
        Unit:  units.RatioDecimalFraction,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with DecimalFraction unit
    decimal_fractionsJSON := []byte(`{"value": 100, "unit": "DecimalFraction"}`)
    decimal_fractionsResult, err := factory.FromDtoJSON(decimal_fractionsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecimalFraction unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimal_fractionsResult.Convert(units.RatioDecimalFraction)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimalFraction = %v, want %v", converted, 100)
    }
    // Test JSON with Percent unit
    percentJSON := []byte(`{"value": 100, "unit": "Percent"}`)
    percentResult, err := factory.FromDtoJSON(percentJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Percent unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = percentResult.Convert(units.RatioPercent)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Percent = %v, want %v", converted, 100)
    }
    // Test JSON with PartPerThousand unit
    parts_per_thousandJSON := []byte(`{"value": 100, "unit": "PartPerThousand"}`)
    parts_per_thousandResult, err := factory.FromDtoJSON(parts_per_thousandJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PartPerThousand unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_thousandResult.Convert(units.RatioPartPerThousand)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerThousand = %v, want %v", converted, 100)
    }
    // Test JSON with PartPerMillion unit
    parts_per_millionJSON := []byte(`{"value": 100, "unit": "PartPerMillion"}`)
    parts_per_millionResult, err := factory.FromDtoJSON(parts_per_millionJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PartPerMillion unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_millionResult.Convert(units.RatioPartPerMillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerMillion = %v, want %v", converted, 100)
    }
    // Test JSON with PartPerBillion unit
    parts_per_billionJSON := []byte(`{"value": 100, "unit": "PartPerBillion"}`)
    parts_per_billionResult, err := factory.FromDtoJSON(parts_per_billionJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PartPerBillion unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_billionResult.Convert(units.RatioPartPerBillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerBillion = %v, want %v", converted, 100)
    }
    // Test JSON with PartPerTrillion unit
    parts_per_trillionJSON := []byte(`{"value": 100, "unit": "PartPerTrillion"}`)
    parts_per_trillionResult, err := factory.FromDtoJSON(parts_per_trillionJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PartPerTrillion unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parts_per_trillionResult.Convert(units.RatioPartPerTrillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PartPerTrillion = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "DecimalFraction"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromDecimalFractions function
func TestRatioFactory_FromDecimalFractions(t *testing.T) {
    factory := units.RatioFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecimalFractions(100)
    if err != nil {
        t.Errorf("FromDecimalFractions() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RatioDecimalFraction)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecimalFractions() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecimalFractions(math.NaN())
    if err == nil {
        t.Error("FromDecimalFractions() with NaN value should return error")
    }

    _, err = factory.FromDecimalFractions(math.Inf(1))
    if err == nil {
        t.Error("FromDecimalFractions() with +Inf value should return error")
    }

    _, err = factory.FromDecimalFractions(math.Inf(-1))
    if err == nil {
        t.Error("FromDecimalFractions() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecimalFractions(0)
    if err != nil {
        t.Errorf("FromDecimalFractions() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RatioDecimalFraction)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecimalFractions() with zero value = %v, want 0", converted)
    }
}
// Test FromPercent function
func TestRatioFactory_FromPercent(t *testing.T) {
    factory := units.RatioFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPercent(100)
    if err != nil {
        t.Errorf("FromPercent() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RatioPercent)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPercent() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPercent(math.NaN())
    if err == nil {
        t.Error("FromPercent() with NaN value should return error")
    }

    _, err = factory.FromPercent(math.Inf(1))
    if err == nil {
        t.Error("FromPercent() with +Inf value should return error")
    }

    _, err = factory.FromPercent(math.Inf(-1))
    if err == nil {
        t.Error("FromPercent() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPercent(0)
    if err != nil {
        t.Errorf("FromPercent() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RatioPercent)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPercent() with zero value = %v, want 0", converted)
    }
}
// Test FromPartsPerThousand function
func TestRatioFactory_FromPartsPerThousand(t *testing.T) {
    factory := units.RatioFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPartsPerThousand(100)
    if err != nil {
        t.Errorf("FromPartsPerThousand() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RatioPartPerThousand)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPartsPerThousand() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPartsPerThousand(math.NaN())
    if err == nil {
        t.Error("FromPartsPerThousand() with NaN value should return error")
    }

    _, err = factory.FromPartsPerThousand(math.Inf(1))
    if err == nil {
        t.Error("FromPartsPerThousand() with +Inf value should return error")
    }

    _, err = factory.FromPartsPerThousand(math.Inf(-1))
    if err == nil {
        t.Error("FromPartsPerThousand() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPartsPerThousand(0)
    if err != nil {
        t.Errorf("FromPartsPerThousand() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RatioPartPerThousand)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPartsPerThousand() with zero value = %v, want 0", converted)
    }
}
// Test FromPartsPerMillion function
func TestRatioFactory_FromPartsPerMillion(t *testing.T) {
    factory := units.RatioFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPartsPerMillion(100)
    if err != nil {
        t.Errorf("FromPartsPerMillion() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RatioPartPerMillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPartsPerMillion() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPartsPerMillion(math.NaN())
    if err == nil {
        t.Error("FromPartsPerMillion() with NaN value should return error")
    }

    _, err = factory.FromPartsPerMillion(math.Inf(1))
    if err == nil {
        t.Error("FromPartsPerMillion() with +Inf value should return error")
    }

    _, err = factory.FromPartsPerMillion(math.Inf(-1))
    if err == nil {
        t.Error("FromPartsPerMillion() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPartsPerMillion(0)
    if err != nil {
        t.Errorf("FromPartsPerMillion() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RatioPartPerMillion)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPartsPerMillion() with zero value = %v, want 0", converted)
    }
}
// Test FromPartsPerBillion function
func TestRatioFactory_FromPartsPerBillion(t *testing.T) {
    factory := units.RatioFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPartsPerBillion(100)
    if err != nil {
        t.Errorf("FromPartsPerBillion() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RatioPartPerBillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPartsPerBillion() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPartsPerBillion(math.NaN())
    if err == nil {
        t.Error("FromPartsPerBillion() with NaN value should return error")
    }

    _, err = factory.FromPartsPerBillion(math.Inf(1))
    if err == nil {
        t.Error("FromPartsPerBillion() with +Inf value should return error")
    }

    _, err = factory.FromPartsPerBillion(math.Inf(-1))
    if err == nil {
        t.Error("FromPartsPerBillion() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPartsPerBillion(0)
    if err != nil {
        t.Errorf("FromPartsPerBillion() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RatioPartPerBillion)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPartsPerBillion() with zero value = %v, want 0", converted)
    }
}
// Test FromPartsPerTrillion function
func TestRatioFactory_FromPartsPerTrillion(t *testing.T) {
    factory := units.RatioFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPartsPerTrillion(100)
    if err != nil {
        t.Errorf("FromPartsPerTrillion() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.RatioPartPerTrillion)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPartsPerTrillion() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPartsPerTrillion(math.NaN())
    if err == nil {
        t.Error("FromPartsPerTrillion() with NaN value should return error")
    }

    _, err = factory.FromPartsPerTrillion(math.Inf(1))
    if err == nil {
        t.Error("FromPartsPerTrillion() with +Inf value should return error")
    }

    _, err = factory.FromPartsPerTrillion(math.Inf(-1))
    if err == nil {
        t.Error("FromPartsPerTrillion() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPartsPerTrillion(0)
    if err != nil {
        t.Errorf("FromPartsPerTrillion() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.RatioPartPerTrillion)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPartsPerTrillion() with zero value = %v, want 0", converted)
    }
}

func TestRatioToString(t *testing.T) {
	factory := units.RatioFactory{}
	a, err := factory.CreateRatio(45, units.RatioDecimalFraction)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.RatioDecimalFraction, 2)
	expected := "45.00 " + units.GetRatioAbbreviation(units.RatioDecimalFraction)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.RatioDecimalFraction, -1)
	expected = "45 " + units.GetRatioAbbreviation(units.RatioDecimalFraction)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestRatio_EqualityAndComparison(t *testing.T) {
	factory := units.RatioFactory{}
	a1, _ := factory.CreateRatio(60, units.RatioDecimalFraction)
	a2, _ := factory.CreateRatio(60, units.RatioDecimalFraction)
	a3, _ := factory.CreateRatio(120, units.RatioDecimalFraction)

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

func TestRatio_Arithmetic(t *testing.T) {
	factory := units.RatioFactory{}
	a1, _ := factory.CreateRatio(30, units.RatioDecimalFraction)
	a2, _ := factory.CreateRatio(45, units.RatioDecimalFraction)

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