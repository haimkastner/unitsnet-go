// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPowerRatioDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "DecibelWatt"}`
	
	factory := units.PowerRatioDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PowerRatioDecibelWatt {
		t.Errorf("expected unit %v, got %v", units.PowerRatioDecibelWatt, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "DecibelWatt"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPowerRatioDto_ToJSON(t *testing.T) {
	dto := units.PowerRatioDto{
		Value: 45,
		Unit:  units.PowerRatioDecibelWatt,
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
	if result["unit"].(string) != string(units.PowerRatioDecibelWatt) {
		t.Errorf("expected unit %s, got %v", units.PowerRatioDecibelWatt, result["unit"])
	}
}

func TestNewPowerRatio_InvalidValue(t *testing.T) {
	factory := units.PowerRatioFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePowerRatio(math.NaN(), units.PowerRatioDecibelWatt)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePowerRatio(math.Inf(1), units.PowerRatioDecibelWatt)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPowerRatioConversions(t *testing.T) {
	factory := units.PowerRatioFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePowerRatio(180, units.PowerRatioDecibelWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to DecibelWatts.
		// No expected conversion value provided for DecibelWatts, verifying result is not NaN.
		result := a.DecibelWatts()
		cacheResult := a.DecibelWatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecibelWatts returned NaN")
		}
	}
	{
		// Test conversion to DecibelMilliwatts.
		// No expected conversion value provided for DecibelMilliwatts, verifying result is not NaN.
		result := a.DecibelMilliwatts()
		cacheResult := a.DecibelMilliwatts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecibelMilliwatts returned NaN")
		}
	}
}

func TestPowerRatio_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PowerRatioFactory{}
	a, err := factory.CreatePowerRatio(90, units.PowerRatioDecibelWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PowerRatioDecibelWatt {
		t.Errorf("expected default unit DecibelWatt, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PowerRatioDecibelWatt
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PowerRatioDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PowerRatioDecibelWatt {
		t.Errorf("expected unit DecibelWatt, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPowerRatioFactory_FromDto(t *testing.T) {
    factory := units.PowerRatioFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.PowerRatioDto{
        Value: 100,
        Unit:  units.PowerRatioDecibelWatt,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.PowerRatioDto{
        Value: math.NaN(),
        Unit:  units.PowerRatioDecibelWatt,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test DecibelWatt conversion
    decibel_wattsDto := units.PowerRatioDto{
        Value: 100,
        Unit:  units.PowerRatioDecibelWatt,
    }
    
    var decibel_wattsResult *units.PowerRatio
    decibel_wattsResult, err = factory.FromDto(decibel_wattsDto)
    if err != nil {
        t.Errorf("FromDto() with DecibelWatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibel_wattsResult.Convert(units.PowerRatioDecibelWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecibelWatt = %v, want %v", converted, 100)
    }
    // Test DecibelMilliwatt conversion
    decibel_milliwattsDto := units.PowerRatioDto{
        Value: 100,
        Unit:  units.PowerRatioDecibelMilliwatt,
    }
    
    var decibel_milliwattsResult *units.PowerRatio
    decibel_milliwattsResult, err = factory.FromDto(decibel_milliwattsDto)
    if err != nil {
        t.Errorf("FromDto() with DecibelMilliwatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibel_milliwattsResult.Convert(units.PowerRatioDecibelMilliwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecibelMilliwatt = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.PowerRatioDto{
        Value: 0,
        Unit:  units.PowerRatioDecibelWatt,
    }
    
    var zeroResult *units.PowerRatio
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestPowerRatioFactory_FromDtoJSON(t *testing.T) {
    factory := units.PowerRatioFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "DecibelWatt"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "DecibelWatt"}`)
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
    nanJSON, _ := json.Marshal(units.PowerRatioDto{
        Value: nanValue,
        Unit:  units.PowerRatioDecibelWatt,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with DecibelWatt unit
    decibel_wattsJSON := []byte(`{"value": 100, "unit": "DecibelWatt"}`)
    decibel_wattsResult, err := factory.FromDtoJSON(decibel_wattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecibelWatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibel_wattsResult.Convert(units.PowerRatioDecibelWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecibelWatt = %v, want %v", converted, 100)
    }
    // Test JSON with DecibelMilliwatt unit
    decibel_milliwattsJSON := []byte(`{"value": 100, "unit": "DecibelMilliwatt"}`)
    decibel_milliwattsResult, err := factory.FromDtoJSON(decibel_milliwattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecibelMilliwatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibel_milliwattsResult.Convert(units.PowerRatioDecibelMilliwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecibelMilliwatt = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "DecibelWatt"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromDecibelWatts function
func TestPowerRatioFactory_FromDecibelWatts(t *testing.T) {
    factory := units.PowerRatioFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecibelWatts(100)
    if err != nil {
        t.Errorf("FromDecibelWatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerRatioDecibelWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecibelWatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecibelWatts(math.NaN())
    if err == nil {
        t.Error("FromDecibelWatts() with NaN value should return error")
    }

    _, err = factory.FromDecibelWatts(math.Inf(1))
    if err == nil {
        t.Error("FromDecibelWatts() with +Inf value should return error")
    }

    _, err = factory.FromDecibelWatts(math.Inf(-1))
    if err == nil {
        t.Error("FromDecibelWatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecibelWatts(0)
    if err != nil {
        t.Errorf("FromDecibelWatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerRatioDecibelWatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecibelWatts() with zero value = %v, want 0", converted)
    }
}
// Test FromDecibelMilliwatts function
func TestPowerRatioFactory_FromDecibelMilliwatts(t *testing.T) {
    factory := units.PowerRatioFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecibelMilliwatts(100)
    if err != nil {
        t.Errorf("FromDecibelMilliwatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerRatioDecibelMilliwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecibelMilliwatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecibelMilliwatts(math.NaN())
    if err == nil {
        t.Error("FromDecibelMilliwatts() with NaN value should return error")
    }

    _, err = factory.FromDecibelMilliwatts(math.Inf(1))
    if err == nil {
        t.Error("FromDecibelMilliwatts() with +Inf value should return error")
    }

    _, err = factory.FromDecibelMilliwatts(math.Inf(-1))
    if err == nil {
        t.Error("FromDecibelMilliwatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecibelMilliwatts(0)
    if err != nil {
        t.Errorf("FromDecibelMilliwatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerRatioDecibelMilliwatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecibelMilliwatts() with zero value = %v, want 0", converted)
    }
}

func TestPowerRatioToString(t *testing.T) {
	factory := units.PowerRatioFactory{}
	a, err := factory.CreatePowerRatio(45, units.PowerRatioDecibelWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PowerRatioDecibelWatt, 2)
	expected := "45.00 " + units.GetPowerRatioAbbreviation(units.PowerRatioDecibelWatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PowerRatioDecibelWatt, -1)
	expected = "45 " + units.GetPowerRatioAbbreviation(units.PowerRatioDecibelWatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPowerRatio_EqualityAndComparison(t *testing.T) {
	factory := units.PowerRatioFactory{}
	a1, _ := factory.CreatePowerRatio(60, units.PowerRatioDecibelWatt)
	a2, _ := factory.CreatePowerRatio(60, units.PowerRatioDecibelWatt)
	a3, _ := factory.CreatePowerRatio(120, units.PowerRatioDecibelWatt)

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

func TestPowerRatio_Arithmetic(t *testing.T) {
	factory := units.PowerRatioFactory{}
	a1, _ := factory.CreatePowerRatio(30, units.PowerRatioDecibelWatt)
	a2, _ := factory.CreatePowerRatio(45, units.PowerRatioDecibelWatt)

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


func TestGetPowerRatioAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.PowerRatioUnits
        want string
    }{
        {
            name: "DecibelWatt abbreviation",
            unit: units.PowerRatioDecibelWatt,
            want: "dBW",
        },
        {
            name: "DecibelMilliwatt abbreviation",
            unit: units.PowerRatioDecibelMilliwatt,
            want: "dBmW",
        },
        {
            name: "invalid unit",
            unit: units.PowerRatioUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetPowerRatioAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetPowerRatioAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestPowerRatio_String(t *testing.T) {
    factory := units.PowerRatioFactory{}
    
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
            unit, err := factory.CreatePowerRatio(tt.value, units.PowerRatioDecibelWatt)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("PowerRatio.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestPowerRatio_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.PowerRatioFactory{}

	_, err := uf.CreatePowerRatio(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}