// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestScalarDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Amount"}`
	
	factory := units.ScalarDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ScalarAmount {
		t.Errorf("expected unit %v, got %v", units.ScalarAmount, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Amount"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestScalarDto_ToJSON(t *testing.T) {
	dto := units.ScalarDto{
		Value: 45,
		Unit:  units.ScalarAmount,
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
	if result["unit"].(string) != string(units.ScalarAmount) {
		t.Errorf("expected unit %s, got %v", units.ScalarAmount, result["unit"])
	}
}

func TestNewScalar_InvalidValue(t *testing.T) {
	factory := units.ScalarFactory{}
	// NaN value should return an error.
	_, err := factory.CreateScalar(math.NaN(), units.ScalarAmount)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateScalar(math.Inf(1), units.ScalarAmount)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestScalarConversions(t *testing.T) {
	factory := units.ScalarFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateScalar(180, units.ScalarAmount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Amount.
		// No expected conversion value provided for Amount, verifying result is not NaN.
		result := a.Amount()
		cacheResult := a.Amount()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Amount returned NaN")
		}
	}
}

func TestScalar_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ScalarFactory{}
	a, err := factory.CreateScalar(90, units.ScalarAmount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ScalarAmount {
		t.Errorf("expected default unit Amount, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ScalarAmount
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ScalarDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ScalarAmount {
		t.Errorf("expected unit Amount, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestScalarFactory_FromDto(t *testing.T) {
    factory := units.ScalarFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ScalarDto{
        Value: 100,
        Unit:  units.ScalarAmount,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ScalarDto{
        Value: math.NaN(),
        Unit:  units.ScalarAmount,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Amount conversion
    amountDto := units.ScalarDto{
        Value: 100,
        Unit:  units.ScalarAmount,
    }
    
    var amountResult *units.Scalar
    amountResult, err = factory.FromDto(amountDto)
    if err != nil {
        t.Errorf("FromDto() with Amount returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amountResult.Convert(units.ScalarAmount)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Amount = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ScalarDto{
        Value: 0,
        Unit:  units.ScalarAmount,
    }
    
    var zeroResult *units.Scalar
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestScalarFactory_FromDtoJSON(t *testing.T) {
    factory := units.ScalarFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Amount"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Amount"}`)
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
    nanJSON, _ := json.Marshal(units.ScalarDto{
        Value: nanValue,
        Unit:  units.ScalarAmount,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Amount unit
    amountJSON := []byte(`{"value": 100, "unit": "Amount"}`)
    amountResult, err := factory.FromDtoJSON(amountJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Amount unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amountResult.Convert(units.ScalarAmount)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Amount = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Amount"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromAmount function
func TestScalarFactory_FromAmount(t *testing.T) {
    factory := units.ScalarFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAmount(100)
    if err != nil {
        t.Errorf("FromAmount() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ScalarAmount)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAmount() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAmount(math.NaN())
    if err == nil {
        t.Error("FromAmount() with NaN value should return error")
    }

    _, err = factory.FromAmount(math.Inf(1))
    if err == nil {
        t.Error("FromAmount() with +Inf value should return error")
    }

    _, err = factory.FromAmount(math.Inf(-1))
    if err == nil {
        t.Error("FromAmount() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAmount(0)
    if err != nil {
        t.Errorf("FromAmount() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ScalarAmount)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAmount() with zero value = %v, want 0", converted)
    }
}

func TestScalarToString(t *testing.T) {
	factory := units.ScalarFactory{}
	a, err := factory.CreateScalar(45, units.ScalarAmount)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ScalarAmount, 2)
	expected := "45.00 " + units.GetScalarAbbreviation(units.ScalarAmount)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ScalarAmount, -1)
	expected = "45 " + units.GetScalarAbbreviation(units.ScalarAmount)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestScalar_EqualityAndComparison(t *testing.T) {
	factory := units.ScalarFactory{}
	a1, _ := factory.CreateScalar(60, units.ScalarAmount)
	a2, _ := factory.CreateScalar(60, units.ScalarAmount)
	a3, _ := factory.CreateScalar(120, units.ScalarAmount)

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

func TestScalar_Arithmetic(t *testing.T) {
	factory := units.ScalarFactory{}
	a1, _ := factory.CreateScalar(30, units.ScalarAmount)
	a2, _ := factory.CreateScalar(45, units.ScalarAmount)

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


func TestGetScalarAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ScalarUnits
        want string
    }{
        {
            name: "Amount abbreviation",
            unit: units.ScalarAmount,
            want: "",
        },
        {
            name: "invalid unit",
            unit: units.ScalarUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetScalarAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetScalarAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestScalar_String(t *testing.T) {
    factory := units.ScalarFactory{}
    
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
            unit, err := factory.CreateScalar(tt.value, units.ScalarAmount)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Scalar.String() = %v, want %v", got, tt.want)
            }
        })
    }
}