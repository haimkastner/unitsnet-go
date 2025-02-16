// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestSolidAngleDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Steradian"}`
	
	factory := units.SolidAngleDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.SolidAngleSteradian {
		t.Errorf("expected unit %v, got %v", units.SolidAngleSteradian, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Steradian"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestSolidAngleDto_ToJSON(t *testing.T) {
	dto := units.SolidAngleDto{
		Value: 45,
		Unit:  units.SolidAngleSteradian,
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
	if result["unit"].(string) != string(units.SolidAngleSteradian) {
		t.Errorf("expected unit %s, got %v", units.SolidAngleSteradian, result["unit"])
	}
}

func TestNewSolidAngle_InvalidValue(t *testing.T) {
	factory := units.SolidAngleFactory{}
	// NaN value should return an error.
	_, err := factory.CreateSolidAngle(math.NaN(), units.SolidAngleSteradian)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateSolidAngle(math.Inf(1), units.SolidAngleSteradian)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestSolidAngleConversions(t *testing.T) {
	factory := units.SolidAngleFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateSolidAngle(180, units.SolidAngleSteradian)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Steradians.
		// No expected conversion value provided for Steradians, verifying result is not NaN.
		result := a.Steradians()
		cacheResult := a.Steradians()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Steradians returned NaN")
		}
	}
}

func TestSolidAngle_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.SolidAngleFactory{}
	a, err := factory.CreateSolidAngle(90, units.SolidAngleSteradian)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.SolidAngleSteradian {
		t.Errorf("expected default unit Steradian, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.SolidAngleSteradian
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.SolidAngleDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.SolidAngleSteradian {
		t.Errorf("expected unit Steradian, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestSolidAngleFactory_FromDto(t *testing.T) {
    factory := units.SolidAngleFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.SolidAngleDto{
        Value: 100,
        Unit:  units.SolidAngleSteradian,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.SolidAngleDto{
        Value: math.NaN(),
        Unit:  units.SolidAngleSteradian,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Steradian conversion
    steradiansDto := units.SolidAngleDto{
        Value: 100,
        Unit:  units.SolidAngleSteradian,
    }
    
    var steradiansResult *units.SolidAngle
    steradiansResult, err = factory.FromDto(steradiansDto)
    if err != nil {
        t.Errorf("FromDto() with Steradian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = steradiansResult.Convert(units.SolidAngleSteradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Steradian = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.SolidAngleDto{
        Value: 0,
        Unit:  units.SolidAngleSteradian,
    }
    
    var zeroResult *units.SolidAngle
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestSolidAngleFactory_FromDtoJSON(t *testing.T) {
    factory := units.SolidAngleFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Steradian"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Steradian"}`)
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
    nanJSON, _ := json.Marshal(units.SolidAngleDto{
        Value: nanValue,
        Unit:  units.SolidAngleSteradian,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Steradian unit
    steradiansJSON := []byte(`{"value": 100, "unit": "Steradian"}`)
    steradiansResult, err := factory.FromDtoJSON(steradiansJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Steradian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = steradiansResult.Convert(units.SolidAngleSteradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Steradian = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Steradian"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromSteradians function
func TestSolidAngleFactory_FromSteradians(t *testing.T) {
    factory := units.SolidAngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSteradians(100)
    if err != nil {
        t.Errorf("FromSteradians() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SolidAngleSteradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSteradians() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSteradians(math.NaN())
    if err == nil {
        t.Error("FromSteradians() with NaN value should return error")
    }

    _, err = factory.FromSteradians(math.Inf(1))
    if err == nil {
        t.Error("FromSteradians() with +Inf value should return error")
    }

    _, err = factory.FromSteradians(math.Inf(-1))
    if err == nil {
        t.Error("FromSteradians() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSteradians(0)
    if err != nil {
        t.Errorf("FromSteradians() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SolidAngleSteradian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSteradians() with zero value = %v, want 0", converted)
    }
}

func TestSolidAngleToString(t *testing.T) {
	factory := units.SolidAngleFactory{}
	a, err := factory.CreateSolidAngle(45, units.SolidAngleSteradian)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.SolidAngleSteradian, 2)
	expected := "45.00 " + units.GetSolidAngleAbbreviation(units.SolidAngleSteradian)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.SolidAngleSteradian, -1)
	expected = "45 " + units.GetSolidAngleAbbreviation(units.SolidAngleSteradian)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestSolidAngle_EqualityAndComparison(t *testing.T) {
	factory := units.SolidAngleFactory{}
	a1, _ := factory.CreateSolidAngle(60, units.SolidAngleSteradian)
	a2, _ := factory.CreateSolidAngle(60, units.SolidAngleSteradian)
	a3, _ := factory.CreateSolidAngle(120, units.SolidAngleSteradian)

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

func TestSolidAngle_Arithmetic(t *testing.T) {
	factory := units.SolidAngleFactory{}
	a1, _ := factory.CreateSolidAngle(30, units.SolidAngleSteradian)
	a2, _ := factory.CreateSolidAngle(45, units.SolidAngleSteradian)

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


func TestGetSolidAngleAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.SolidAngleUnits
        want string
    }{
        {
            name: "Steradian abbreviation",
            unit: units.SolidAngleSteradian,
            want: "sr",
        },
        {
            name: "invalid unit",
            unit: units.SolidAngleUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetSolidAngleAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetSolidAngleAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestSolidAngle_String(t *testing.T) {
    factory := units.SolidAngleFactory{}
    
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
            unit, err := factory.CreateSolidAngle(tt.value, units.SolidAngleSteradian)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("SolidAngle.String() = %v, want %v", got, tt.want)
            }
        })
    }
}