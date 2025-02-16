// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestTurbidityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NTU"}`
	
	factory := units.TurbidityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.TurbidityNTU {
		t.Errorf("expected unit %v, got %v", units.TurbidityNTU, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NTU"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestTurbidityDto_ToJSON(t *testing.T) {
	dto := units.TurbidityDto{
		Value: 45,
		Unit:  units.TurbidityNTU,
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
	if result["unit"].(string) != string(units.TurbidityNTU) {
		t.Errorf("expected unit %s, got %v", units.TurbidityNTU, result["unit"])
	}
}

func TestNewTurbidity_InvalidValue(t *testing.T) {
	factory := units.TurbidityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateTurbidity(math.NaN(), units.TurbidityNTU)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateTurbidity(math.Inf(1), units.TurbidityNTU)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestTurbidityConversions(t *testing.T) {
	factory := units.TurbidityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateTurbidity(180, units.TurbidityNTU)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NTU.
		// No expected conversion value provided for NTU, verifying result is not NaN.
		result := a.NTU()
		cacheResult := a.NTU()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NTU returned NaN")
		}
	}
}

func TestTurbidity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.TurbidityFactory{}
	a, err := factory.CreateTurbidity(90, units.TurbidityNTU)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.TurbidityNTU {
		t.Errorf("expected default unit NTU, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.TurbidityNTU
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.TurbidityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.TurbidityNTU {
		t.Errorf("expected unit NTU, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestTurbidityFactory_FromDto(t *testing.T) {
    factory := units.TurbidityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.TurbidityDto{
        Value: 100,
        Unit:  units.TurbidityNTU,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.TurbidityDto{
        Value: math.NaN(),
        Unit:  units.TurbidityNTU,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test NTU conversion
    ntuDto := units.TurbidityDto{
        Value: 100,
        Unit:  units.TurbidityNTU,
    }
    
    var ntuResult *units.Turbidity
    ntuResult, err = factory.FromDto(ntuDto)
    if err != nil {
        t.Errorf("FromDto() with NTU returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ntuResult.Convert(units.TurbidityNTU)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NTU = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.TurbidityDto{
        Value: 0,
        Unit:  units.TurbidityNTU,
    }
    
    var zeroResult *units.Turbidity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestTurbidityFactory_FromDtoJSON(t *testing.T) {
    factory := units.TurbidityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "NTU"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "NTU"}`)
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
    nanJSON, _ := json.Marshal(units.TurbidityDto{
        Value: nanValue,
        Unit:  units.TurbidityNTU,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with NTU unit
    ntuJSON := []byte(`{"value": 100, "unit": "NTU"}`)
    ntuResult, err := factory.FromDtoJSON(ntuJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NTU unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ntuResult.Convert(units.TurbidityNTU)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NTU = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "NTU"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromNTU function
func TestTurbidityFactory_FromNTU(t *testing.T) {
    factory := units.TurbidityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNTU(100)
    if err != nil {
        t.Errorf("FromNTU() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TurbidityNTU)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNTU() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNTU(math.NaN())
    if err == nil {
        t.Error("FromNTU() with NaN value should return error")
    }

    _, err = factory.FromNTU(math.Inf(1))
    if err == nil {
        t.Error("FromNTU() with +Inf value should return error")
    }

    _, err = factory.FromNTU(math.Inf(-1))
    if err == nil {
        t.Error("FromNTU() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNTU(0)
    if err != nil {
        t.Errorf("FromNTU() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TurbidityNTU)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNTU() with zero value = %v, want 0", converted)
    }
}

func TestTurbidityToString(t *testing.T) {
	factory := units.TurbidityFactory{}
	a, err := factory.CreateTurbidity(45, units.TurbidityNTU)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.TurbidityNTU, 2)
	expected := "45.00 " + units.GetTurbidityAbbreviation(units.TurbidityNTU)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.TurbidityNTU, -1)
	expected = "45 " + units.GetTurbidityAbbreviation(units.TurbidityNTU)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestTurbidity_EqualityAndComparison(t *testing.T) {
	factory := units.TurbidityFactory{}
	a1, _ := factory.CreateTurbidity(60, units.TurbidityNTU)
	a2, _ := factory.CreateTurbidity(60, units.TurbidityNTU)
	a3, _ := factory.CreateTurbidity(120, units.TurbidityNTU)

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

func TestTurbidity_Arithmetic(t *testing.T) {
	factory := units.TurbidityFactory{}
	a1, _ := factory.CreateTurbidity(30, units.TurbidityNTU)
	a2, _ := factory.CreateTurbidity(45, units.TurbidityNTU)

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


func TestGetTurbidityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.TurbidityUnits
        want string
    }{
        {
            name: "NTU abbreviation",
            unit: units.TurbidityNTU,
            want: "NTU",
        },
        {
            name: "invalid unit",
            unit: units.TurbidityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetTurbidityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetTurbidityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestTurbidity_String(t *testing.T) {
    factory := units.TurbidityFactory{}
    
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
            unit, err := factory.CreateTurbidity(tt.value, units.TurbidityNTU)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Turbidity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}