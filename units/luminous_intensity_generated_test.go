// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLuminousIntensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Candela"}`
	
	factory := units.LuminousIntensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LuminousIntensityCandela {
		t.Errorf("expected unit %v, got %v", units.LuminousIntensityCandela, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Candela"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLuminousIntensityDto_ToJSON(t *testing.T) {
	dto := units.LuminousIntensityDto{
		Value: 45,
		Unit:  units.LuminousIntensityCandela,
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
	if result["unit"].(string) != string(units.LuminousIntensityCandela) {
		t.Errorf("expected unit %s, got %v", units.LuminousIntensityCandela, result["unit"])
	}
}

func TestNewLuminousIntensity_InvalidValue(t *testing.T) {
	factory := units.LuminousIntensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLuminousIntensity(math.NaN(), units.LuminousIntensityCandela)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLuminousIntensity(math.Inf(1), units.LuminousIntensityCandela)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLuminousIntensityConversions(t *testing.T) {
	factory := units.LuminousIntensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLuminousIntensity(180, units.LuminousIntensityCandela)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Candela.
		// No expected conversion value provided for Candela, verifying result is not NaN.
		result := a.Candela()
		cacheResult := a.Candela()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Candela returned NaN")
		}
	}
}

func TestLuminousIntensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LuminousIntensityFactory{}
	a, err := factory.CreateLuminousIntensity(90, units.LuminousIntensityCandela)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LuminousIntensityCandela {
		t.Errorf("expected default unit Candela, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LuminousIntensityCandela
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LuminousIntensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LuminousIntensityCandela {
		t.Errorf("expected unit Candela, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLuminousIntensityFactory_FromDto(t *testing.T) {
    factory := units.LuminousIntensityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.LuminousIntensityDto{
        Value: 100,
        Unit:  units.LuminousIntensityCandela,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.LuminousIntensityDto{
        Value: math.NaN(),
        Unit:  units.LuminousIntensityCandela,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Candela conversion
    candelaDto := units.LuminousIntensityDto{
        Value: 100,
        Unit:  units.LuminousIntensityCandela,
    }
    
    var candelaResult *units.LuminousIntensity
    candelaResult, err = factory.FromDto(candelaDto)
    if err != nil {
        t.Errorf("FromDto() with Candela returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = candelaResult.Convert(units.LuminousIntensityCandela)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Candela = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.LuminousIntensityDto{
        Value: 0,
        Unit:  units.LuminousIntensityCandela,
    }
    
    var zeroResult *units.LuminousIntensity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestLuminousIntensityFactory_FromDtoJSON(t *testing.T) {
    factory := units.LuminousIntensityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Candela"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Candela"}`)
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
    nanJSON, _ := json.Marshal(units.LuminousIntensityDto{
        Value: nanValue,
        Unit:  units.LuminousIntensityCandela,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Candela unit
    candelaJSON := []byte(`{"value": 100, "unit": "Candela"}`)
    candelaResult, err := factory.FromDtoJSON(candelaJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Candela unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = candelaResult.Convert(units.LuminousIntensityCandela)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Candela = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Candela"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromCandela function
func TestLuminousIntensityFactory_FromCandela(t *testing.T) {
    factory := units.LuminousIntensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCandela(100)
    if err != nil {
        t.Errorf("FromCandela() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LuminousIntensityCandela)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCandela() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCandela(math.NaN())
    if err == nil {
        t.Error("FromCandela() with NaN value should return error")
    }

    _, err = factory.FromCandela(math.Inf(1))
    if err == nil {
        t.Error("FromCandela() with +Inf value should return error")
    }

    _, err = factory.FromCandela(math.Inf(-1))
    if err == nil {
        t.Error("FromCandela() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCandela(0)
    if err != nil {
        t.Errorf("FromCandela() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LuminousIntensityCandela)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCandela() with zero value = %v, want 0", converted)
    }
}

func TestLuminousIntensityToString(t *testing.T) {
	factory := units.LuminousIntensityFactory{}
	a, err := factory.CreateLuminousIntensity(45, units.LuminousIntensityCandela)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LuminousIntensityCandela, 2)
	expected := "45.00 " + units.GetLuminousIntensityAbbreviation(units.LuminousIntensityCandela)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LuminousIntensityCandela, -1)
	expected = "45 " + units.GetLuminousIntensityAbbreviation(units.LuminousIntensityCandela)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLuminousIntensity_EqualityAndComparison(t *testing.T) {
	factory := units.LuminousIntensityFactory{}
	a1, _ := factory.CreateLuminousIntensity(60, units.LuminousIntensityCandela)
	a2, _ := factory.CreateLuminousIntensity(60, units.LuminousIntensityCandela)
	a3, _ := factory.CreateLuminousIntensity(120, units.LuminousIntensityCandela)

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

func TestLuminousIntensity_Arithmetic(t *testing.T) {
	factory := units.LuminousIntensityFactory{}
	a1, _ := factory.CreateLuminousIntensity(30, units.LuminousIntensityCandela)
	a2, _ := factory.CreateLuminousIntensity(45, units.LuminousIntensityCandela)

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


func TestGetLuminousIntensityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.LuminousIntensityUnits
        want string
    }{
        {
            name: "Candela abbreviation",
            unit: units.LuminousIntensityCandela,
            want: "cd",
        },
        {
            name: "invalid unit",
            unit: units.LuminousIntensityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetLuminousIntensityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetLuminousIntensityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestLuminousIntensity_String(t *testing.T) {
    factory := units.LuminousIntensityFactory{}
    
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
            unit, err := factory.CreateLuminousIntensity(tt.value, units.LuminousIntensityCandela)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("LuminousIntensity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestLuminousIntensity_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.LuminousIntensityFactory{}

	_, err := uf.CreateLuminousIntensity(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}