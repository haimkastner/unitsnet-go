// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestIlluminanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Lux"}`
	
	factory := units.IlluminanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.IlluminanceLux {
		t.Errorf("expected unit %v, got %v", units.IlluminanceLux, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Lux"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestIlluminanceDto_ToJSON(t *testing.T) {
	dto := units.IlluminanceDto{
		Value: 45,
		Unit:  units.IlluminanceLux,
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
	if result["unit"].(string) != string(units.IlluminanceLux) {
		t.Errorf("expected unit %s, got %v", units.IlluminanceLux, result["unit"])
	}
}

func TestNewIlluminance_InvalidValue(t *testing.T) {
	factory := units.IlluminanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateIlluminance(math.NaN(), units.IlluminanceLux)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateIlluminance(math.Inf(1), units.IlluminanceLux)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestIlluminanceConversions(t *testing.T) {
	factory := units.IlluminanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateIlluminance(180, units.IlluminanceLux)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Lux.
		// No expected conversion value provided for Lux, verifying result is not NaN.
		result := a.Lux()
		cacheResult := a.Lux()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Lux returned NaN")
		}
	}
	{
		// Test conversion to Millilux.
		// No expected conversion value provided for Millilux, verifying result is not NaN.
		result := a.Millilux()
		cacheResult := a.Millilux()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millilux returned NaN")
		}
	}
	{
		// Test conversion to Kilolux.
		// No expected conversion value provided for Kilolux, verifying result is not NaN.
		result := a.Kilolux()
		cacheResult := a.Kilolux()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilolux returned NaN")
		}
	}
	{
		// Test conversion to Megalux.
		// No expected conversion value provided for Megalux, verifying result is not NaN.
		result := a.Megalux()
		cacheResult := a.Megalux()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megalux returned NaN")
		}
	}
}

func TestIlluminance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.IlluminanceFactory{}
	a, err := factory.CreateIlluminance(90, units.IlluminanceLux)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.IlluminanceLux {
		t.Errorf("expected default unit Lux, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.IlluminanceLux
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.IlluminanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.IlluminanceLux {
		t.Errorf("expected unit Lux, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestIlluminanceFactory_FromDto(t *testing.T) {
    factory := units.IlluminanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.IlluminanceDto{
        Value: 100,
        Unit:  units.IlluminanceLux,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.IlluminanceDto{
        Value: math.NaN(),
        Unit:  units.IlluminanceLux,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Lux conversion
    luxDto := units.IlluminanceDto{
        Value: 100,
        Unit:  units.IlluminanceLux,
    }
    
    var luxResult *units.Illuminance
    luxResult, err = factory.FromDto(luxDto)
    if err != nil {
        t.Errorf("FromDto() with Lux returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = luxResult.Convert(units.IlluminanceLux)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Lux = %v, want %v", converted, 100)
    }
    // Test Millilux conversion
    milliluxDto := units.IlluminanceDto{
        Value: 100,
        Unit:  units.IlluminanceMillilux,
    }
    
    var milliluxResult *units.Illuminance
    milliluxResult, err = factory.FromDto(milliluxDto)
    if err != nil {
        t.Errorf("FromDto() with Millilux returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliluxResult.Convert(units.IlluminanceMillilux)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millilux = %v, want %v", converted, 100)
    }
    // Test Kilolux conversion
    kiloluxDto := units.IlluminanceDto{
        Value: 100,
        Unit:  units.IlluminanceKilolux,
    }
    
    var kiloluxResult *units.Illuminance
    kiloluxResult, err = factory.FromDto(kiloluxDto)
    if err != nil {
        t.Errorf("FromDto() with Kilolux returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloluxResult.Convert(units.IlluminanceKilolux)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilolux = %v, want %v", converted, 100)
    }
    // Test Megalux conversion
    megaluxDto := units.IlluminanceDto{
        Value: 100,
        Unit:  units.IlluminanceMegalux,
    }
    
    var megaluxResult *units.Illuminance
    megaluxResult, err = factory.FromDto(megaluxDto)
    if err != nil {
        t.Errorf("FromDto() with Megalux returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaluxResult.Convert(units.IlluminanceMegalux)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megalux = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.IlluminanceDto{
        Value: 0,
        Unit:  units.IlluminanceLux,
    }
    
    var zeroResult *units.Illuminance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestIlluminanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.IlluminanceFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Lux"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Lux"}`)
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
    nanJSON, _ := json.Marshal(units.IlluminanceDto{
        Value: nanValue,
        Unit:  units.IlluminanceLux,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Lux unit
    luxJSON := []byte(`{"value": 100, "unit": "Lux"}`)
    luxResult, err := factory.FromDtoJSON(luxJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Lux unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = luxResult.Convert(units.IlluminanceLux)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Lux = %v, want %v", converted, 100)
    }
    // Test JSON with Millilux unit
    milliluxJSON := []byte(`{"value": 100, "unit": "Millilux"}`)
    milliluxResult, err := factory.FromDtoJSON(milliluxJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millilux unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliluxResult.Convert(units.IlluminanceMillilux)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millilux = %v, want %v", converted, 100)
    }
    // Test JSON with Kilolux unit
    kiloluxJSON := []byte(`{"value": 100, "unit": "Kilolux"}`)
    kiloluxResult, err := factory.FromDtoJSON(kiloluxJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilolux unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloluxResult.Convert(units.IlluminanceKilolux)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilolux = %v, want %v", converted, 100)
    }
    // Test JSON with Megalux unit
    megaluxJSON := []byte(`{"value": 100, "unit": "Megalux"}`)
    megaluxResult, err := factory.FromDtoJSON(megaluxJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megalux unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaluxResult.Convert(units.IlluminanceMegalux)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megalux = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Lux"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromLux function
func TestIlluminanceFactory_FromLux(t *testing.T) {
    factory := units.IlluminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLux(100)
    if err != nil {
        t.Errorf("FromLux() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IlluminanceLux)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLux() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLux(math.NaN())
    if err == nil {
        t.Error("FromLux() with NaN value should return error")
    }

    _, err = factory.FromLux(math.Inf(1))
    if err == nil {
        t.Error("FromLux() with +Inf value should return error")
    }

    _, err = factory.FromLux(math.Inf(-1))
    if err == nil {
        t.Error("FromLux() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLux(0)
    if err != nil {
        t.Errorf("FromLux() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IlluminanceLux)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLux() with zero value = %v, want 0", converted)
    }
}
// Test FromMillilux function
func TestIlluminanceFactory_FromMillilux(t *testing.T) {
    factory := units.IlluminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillilux(100)
    if err != nil {
        t.Errorf("FromMillilux() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IlluminanceMillilux)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillilux() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillilux(math.NaN())
    if err == nil {
        t.Error("FromMillilux() with NaN value should return error")
    }

    _, err = factory.FromMillilux(math.Inf(1))
    if err == nil {
        t.Error("FromMillilux() with +Inf value should return error")
    }

    _, err = factory.FromMillilux(math.Inf(-1))
    if err == nil {
        t.Error("FromMillilux() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillilux(0)
    if err != nil {
        t.Errorf("FromMillilux() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IlluminanceMillilux)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillilux() with zero value = %v, want 0", converted)
    }
}
// Test FromKilolux function
func TestIlluminanceFactory_FromKilolux(t *testing.T) {
    factory := units.IlluminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilolux(100)
    if err != nil {
        t.Errorf("FromKilolux() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IlluminanceKilolux)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilolux() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilolux(math.NaN())
    if err == nil {
        t.Error("FromKilolux() with NaN value should return error")
    }

    _, err = factory.FromKilolux(math.Inf(1))
    if err == nil {
        t.Error("FromKilolux() with +Inf value should return error")
    }

    _, err = factory.FromKilolux(math.Inf(-1))
    if err == nil {
        t.Error("FromKilolux() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilolux(0)
    if err != nil {
        t.Errorf("FromKilolux() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IlluminanceKilolux)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilolux() with zero value = %v, want 0", converted)
    }
}
// Test FromMegalux function
func TestIlluminanceFactory_FromMegalux(t *testing.T) {
    factory := units.IlluminanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegalux(100)
    if err != nil {
        t.Errorf("FromMegalux() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IlluminanceMegalux)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegalux() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegalux(math.NaN())
    if err == nil {
        t.Error("FromMegalux() with NaN value should return error")
    }

    _, err = factory.FromMegalux(math.Inf(1))
    if err == nil {
        t.Error("FromMegalux() with +Inf value should return error")
    }

    _, err = factory.FromMegalux(math.Inf(-1))
    if err == nil {
        t.Error("FromMegalux() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegalux(0)
    if err != nil {
        t.Errorf("FromMegalux() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IlluminanceMegalux)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegalux() with zero value = %v, want 0", converted)
    }
}

func TestIlluminanceToString(t *testing.T) {
	factory := units.IlluminanceFactory{}
	a, err := factory.CreateIlluminance(45, units.IlluminanceLux)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.IlluminanceLux, 2)
	expected := "45.00 " + units.GetIlluminanceAbbreviation(units.IlluminanceLux)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.IlluminanceLux, -1)
	expected = "45 " + units.GetIlluminanceAbbreviation(units.IlluminanceLux)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestIlluminance_EqualityAndComparison(t *testing.T) {
	factory := units.IlluminanceFactory{}
	a1, _ := factory.CreateIlluminance(60, units.IlluminanceLux)
	a2, _ := factory.CreateIlluminance(60, units.IlluminanceLux)
	a3, _ := factory.CreateIlluminance(120, units.IlluminanceLux)

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

func TestIlluminance_Arithmetic(t *testing.T) {
	factory := units.IlluminanceFactory{}
	a1, _ := factory.CreateIlluminance(30, units.IlluminanceLux)
	a2, _ := factory.CreateIlluminance(45, units.IlluminanceLux)

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


func TestGetIlluminanceAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.IlluminanceUnits
        want string
    }{
        {
            name: "Lux abbreviation",
            unit: units.IlluminanceLux,
            want: "lx",
        },
        {
            name: "Millilux abbreviation",
            unit: units.IlluminanceMillilux,
            want: "mlx",
        },
        {
            name: "Kilolux abbreviation",
            unit: units.IlluminanceKilolux,
            want: "klx",
        },
        {
            name: "Megalux abbreviation",
            unit: units.IlluminanceMegalux,
            want: "Mlx",
        },
        {
            name: "invalid unit",
            unit: units.IlluminanceUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetIlluminanceAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetIlluminanceAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestIlluminance_String(t *testing.T) {
    factory := units.IlluminanceFactory{}
    
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
            unit, err := factory.CreateIlluminance(tt.value, units.IlluminanceLux)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Illuminance.String() = %v, want %v", got, tt.want)
            }
        })
    }
}