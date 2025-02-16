// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPermittivityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "FaradPerMeter"}`
	
	factory := units.PermittivityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PermittivityFaradPerMeter {
		t.Errorf("expected unit %v, got %v", units.PermittivityFaradPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "FaradPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPermittivityDto_ToJSON(t *testing.T) {
	dto := units.PermittivityDto{
		Value: 45,
		Unit:  units.PermittivityFaradPerMeter,
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
	if result["unit"].(string) != string(units.PermittivityFaradPerMeter) {
		t.Errorf("expected unit %s, got %v", units.PermittivityFaradPerMeter, result["unit"])
	}
}

func TestNewPermittivity_InvalidValue(t *testing.T) {
	factory := units.PermittivityFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePermittivity(math.NaN(), units.PermittivityFaradPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePermittivity(math.Inf(1), units.PermittivityFaradPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPermittivityConversions(t *testing.T) {
	factory := units.PermittivityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePermittivity(180, units.PermittivityFaradPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to FaradsPerMeter.
		// No expected conversion value provided for FaradsPerMeter, verifying result is not NaN.
		result := a.FaradsPerMeter()
		cacheResult := a.FaradsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to FaradsPerMeter returned NaN")
		}
	}
}

func TestPermittivity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PermittivityFactory{}
	a, err := factory.CreatePermittivity(90, units.PermittivityFaradPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PermittivityFaradPerMeter {
		t.Errorf("expected default unit FaradPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PermittivityFaradPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PermittivityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PermittivityFaradPerMeter {
		t.Errorf("expected unit FaradPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPermittivityFactory_FromDto(t *testing.T) {
    factory := units.PermittivityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.PermittivityDto{
        Value: 100,
        Unit:  units.PermittivityFaradPerMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.PermittivityDto{
        Value: math.NaN(),
        Unit:  units.PermittivityFaradPerMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test FaradPerMeter conversion
    farads_per_meterDto := units.PermittivityDto{
        Value: 100,
        Unit:  units.PermittivityFaradPerMeter,
    }
    
    var farads_per_meterResult *units.Permittivity
    farads_per_meterResult, err = factory.FromDto(farads_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with FaradPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = farads_per_meterResult.Convert(units.PermittivityFaradPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FaradPerMeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.PermittivityDto{
        Value: 0,
        Unit:  units.PermittivityFaradPerMeter,
    }
    
    var zeroResult *units.Permittivity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestPermittivityFactory_FromDtoJSON(t *testing.T) {
    factory := units.PermittivityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "FaradPerMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "FaradPerMeter"}`)
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
    nanJSON, _ := json.Marshal(units.PermittivityDto{
        Value: nanValue,
        Unit:  units.PermittivityFaradPerMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with FaradPerMeter unit
    farads_per_meterJSON := []byte(`{"value": 100, "unit": "FaradPerMeter"}`)
    farads_per_meterResult, err := factory.FromDtoJSON(farads_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FaradPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = farads_per_meterResult.Convert(units.PermittivityFaradPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FaradPerMeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "FaradPerMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromFaradsPerMeter function
func TestPermittivityFactory_FromFaradsPerMeter(t *testing.T) {
    factory := units.PermittivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFaradsPerMeter(100)
    if err != nil {
        t.Errorf("FromFaradsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PermittivityFaradPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFaradsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFaradsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromFaradsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromFaradsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromFaradsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromFaradsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromFaradsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFaradsPerMeter(0)
    if err != nil {
        t.Errorf("FromFaradsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PermittivityFaradPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFaradsPerMeter() with zero value = %v, want 0", converted)
    }
}

func TestPermittivityToString(t *testing.T) {
	factory := units.PermittivityFactory{}
	a, err := factory.CreatePermittivity(45, units.PermittivityFaradPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PermittivityFaradPerMeter, 2)
	expected := "45.00 " + units.GetPermittivityAbbreviation(units.PermittivityFaradPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PermittivityFaradPerMeter, -1)
	expected = "45 " + units.GetPermittivityAbbreviation(units.PermittivityFaradPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPermittivity_EqualityAndComparison(t *testing.T) {
	factory := units.PermittivityFactory{}
	a1, _ := factory.CreatePermittivity(60, units.PermittivityFaradPerMeter)
	a2, _ := factory.CreatePermittivity(60, units.PermittivityFaradPerMeter)
	a3, _ := factory.CreatePermittivity(120, units.PermittivityFaradPerMeter)

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

func TestPermittivity_Arithmetic(t *testing.T) {
	factory := units.PermittivityFactory{}
	a1, _ := factory.CreatePermittivity(30, units.PermittivityFaradPerMeter)
	a2, _ := factory.CreatePermittivity(45, units.PermittivityFaradPerMeter)

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


func TestGetPermittivityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.PermittivityUnits
        want string
    }{
        {
            name: "FaradPerMeter abbreviation",
            unit: units.PermittivityFaradPerMeter,
            want: "F/m",
        },
        {
            name: "invalid unit",
            unit: units.PermittivityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetPermittivityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetPermittivityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestPermittivity_String(t *testing.T) {
    factory := units.PermittivityFactory{}
    
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
            unit, err := factory.CreatePermittivity(tt.value, units.PermittivityFaradPerMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Permittivity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}