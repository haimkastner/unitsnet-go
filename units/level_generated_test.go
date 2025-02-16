// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLevelDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Decibel"}`
	
	factory := units.LevelDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LevelDecibel {
		t.Errorf("expected unit %v, got %v", units.LevelDecibel, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Decibel"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLevelDto_ToJSON(t *testing.T) {
	dto := units.LevelDto{
		Value: 45,
		Unit:  units.LevelDecibel,
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
	if result["unit"].(string) != string(units.LevelDecibel) {
		t.Errorf("expected unit %s, got %v", units.LevelDecibel, result["unit"])
	}
}

func TestNewLevel_InvalidValue(t *testing.T) {
	factory := units.LevelFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLevel(math.NaN(), units.LevelDecibel)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLevel(math.Inf(1), units.LevelDecibel)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLevelConversions(t *testing.T) {
	factory := units.LevelFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLevel(180, units.LevelDecibel)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Decibels.
		// No expected conversion value provided for Decibels, verifying result is not NaN.
		result := a.Decibels()
		cacheResult := a.Decibels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Decibels returned NaN")
		}
	}
	{
		// Test conversion to Nepers.
		// No expected conversion value provided for Nepers, verifying result is not NaN.
		result := a.Nepers()
		cacheResult := a.Nepers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nepers returned NaN")
		}
	}
}

func TestLevel_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LevelFactory{}
	a, err := factory.CreateLevel(90, units.LevelDecibel)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LevelDecibel {
		t.Errorf("expected default unit Decibel, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LevelDecibel
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LevelDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LevelDecibel {
		t.Errorf("expected unit Decibel, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLevelFactory_FromDto(t *testing.T) {
    factory := units.LevelFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.LevelDto{
        Value: 100,
        Unit:  units.LevelDecibel,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.LevelDto{
        Value: math.NaN(),
        Unit:  units.LevelDecibel,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Decibel conversion
    decibelsDto := units.LevelDto{
        Value: 100,
        Unit:  units.LevelDecibel,
    }
    
    var decibelsResult *units.Level
    decibelsResult, err = factory.FromDto(decibelsDto)
    if err != nil {
        t.Errorf("FromDto() with Decibel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibelsResult.Convert(units.LevelDecibel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decibel = %v, want %v", converted, 100)
    }
    // Test Neper conversion
    nepersDto := units.LevelDto{
        Value: 100,
        Unit:  units.LevelNeper,
    }
    
    var nepersResult *units.Level
    nepersResult, err = factory.FromDto(nepersDto)
    if err != nil {
        t.Errorf("FromDto() with Neper returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nepersResult.Convert(units.LevelNeper)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Neper = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.LevelDto{
        Value: 0,
        Unit:  units.LevelDecibel,
    }
    
    var zeroResult *units.Level
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestLevelFactory_FromDtoJSON(t *testing.T) {
    factory := units.LevelFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Decibel"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Decibel"}`)
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
    nanJSON, _ := json.Marshal(units.LevelDto{
        Value: nanValue,
        Unit:  units.LevelDecibel,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Decibel unit
    decibelsJSON := []byte(`{"value": 100, "unit": "Decibel"}`)
    decibelsResult, err := factory.FromDtoJSON(decibelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Decibel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibelsResult.Convert(units.LevelDecibel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decibel = %v, want %v", converted, 100)
    }
    // Test JSON with Neper unit
    nepersJSON := []byte(`{"value": 100, "unit": "Neper"}`)
    nepersResult, err := factory.FromDtoJSON(nepersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Neper unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nepersResult.Convert(units.LevelNeper)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Neper = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Decibel"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromDecibels function
func TestLevelFactory_FromDecibels(t *testing.T) {
    factory := units.LevelFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecibels(100)
    if err != nil {
        t.Errorf("FromDecibels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LevelDecibel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecibels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecibels(math.NaN())
    if err == nil {
        t.Error("FromDecibels() with NaN value should return error")
    }

    _, err = factory.FromDecibels(math.Inf(1))
    if err == nil {
        t.Error("FromDecibels() with +Inf value should return error")
    }

    _, err = factory.FromDecibels(math.Inf(-1))
    if err == nil {
        t.Error("FromDecibels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecibels(0)
    if err != nil {
        t.Errorf("FromDecibels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LevelDecibel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecibels() with zero value = %v, want 0", converted)
    }
}
// Test FromNepers function
func TestLevelFactory_FromNepers(t *testing.T) {
    factory := units.LevelFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNepers(100)
    if err != nil {
        t.Errorf("FromNepers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LevelNeper)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNepers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNepers(math.NaN())
    if err == nil {
        t.Error("FromNepers() with NaN value should return error")
    }

    _, err = factory.FromNepers(math.Inf(1))
    if err == nil {
        t.Error("FromNepers() with +Inf value should return error")
    }

    _, err = factory.FromNepers(math.Inf(-1))
    if err == nil {
        t.Error("FromNepers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNepers(0)
    if err != nil {
        t.Errorf("FromNepers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LevelNeper)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNepers() with zero value = %v, want 0", converted)
    }
}

func TestLevelToString(t *testing.T) {
	factory := units.LevelFactory{}
	a, err := factory.CreateLevel(45, units.LevelDecibel)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LevelDecibel, 2)
	expected := "45.00 " + units.GetLevelAbbreviation(units.LevelDecibel)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LevelDecibel, -1)
	expected = "45 " + units.GetLevelAbbreviation(units.LevelDecibel)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLevel_EqualityAndComparison(t *testing.T) {
	factory := units.LevelFactory{}
	a1, _ := factory.CreateLevel(60, units.LevelDecibel)
	a2, _ := factory.CreateLevel(60, units.LevelDecibel)
	a3, _ := factory.CreateLevel(120, units.LevelDecibel)

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

func TestLevel_Arithmetic(t *testing.T) {
	factory := units.LevelFactory{}
	a1, _ := factory.CreateLevel(30, units.LevelDecibel)
	a2, _ := factory.CreateLevel(45, units.LevelDecibel)

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


func TestGetLevelAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.LevelUnits
        want string
    }{
        {
            name: "Decibel abbreviation",
            unit: units.LevelDecibel,
            want: "dB",
        },
        {
            name: "Neper abbreviation",
            unit: units.LevelNeper,
            want: "Np",
        },
        {
            name: "invalid unit",
            unit: units.LevelUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetLevelAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetLevelAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestLevel_String(t *testing.T) {
    factory := units.LevelFactory{}
    
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
            unit, err := factory.CreateLevel(tt.value, units.LevelDecibel)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Level.String() = %v, want %v", got, tt.want)
            }
        })
    }
}