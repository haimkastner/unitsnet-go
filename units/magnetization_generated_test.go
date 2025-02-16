// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMagnetizationDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "AmperePerMeter"}`
	
	factory := units.MagnetizationDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MagnetizationAmperePerMeter {
		t.Errorf("expected unit %v, got %v", units.MagnetizationAmperePerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "AmperePerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMagnetizationDto_ToJSON(t *testing.T) {
	dto := units.MagnetizationDto{
		Value: 45,
		Unit:  units.MagnetizationAmperePerMeter,
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
	if result["unit"].(string) != string(units.MagnetizationAmperePerMeter) {
		t.Errorf("expected unit %s, got %v", units.MagnetizationAmperePerMeter, result["unit"])
	}
}

func TestNewMagnetization_InvalidValue(t *testing.T) {
	factory := units.MagnetizationFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMagnetization(math.NaN(), units.MagnetizationAmperePerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMagnetization(math.Inf(1), units.MagnetizationAmperePerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMagnetizationConversions(t *testing.T) {
	factory := units.MagnetizationFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMagnetization(180, units.MagnetizationAmperePerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to AmperesPerMeter.
		// No expected conversion value provided for AmperesPerMeter, verifying result is not NaN.
		result := a.AmperesPerMeter()
		cacheResult := a.AmperesPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AmperesPerMeter returned NaN")
		}
	}
}

func TestMagnetization_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MagnetizationFactory{}
	a, err := factory.CreateMagnetization(90, units.MagnetizationAmperePerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MagnetizationAmperePerMeter {
		t.Errorf("expected default unit AmperePerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MagnetizationAmperePerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MagnetizationDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MagnetizationAmperePerMeter {
		t.Errorf("expected unit AmperePerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMagnetizationFactory_FromDto(t *testing.T) {
    factory := units.MagnetizationFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MagnetizationDto{
        Value: 100,
        Unit:  units.MagnetizationAmperePerMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MagnetizationDto{
        Value: math.NaN(),
        Unit:  units.MagnetizationAmperePerMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test AmperePerMeter conversion
    amperes_per_meterDto := units.MagnetizationDto{
        Value: 100,
        Unit:  units.MagnetizationAmperePerMeter,
    }
    
    var amperes_per_meterResult *units.Magnetization
    amperes_per_meterResult, err = factory.FromDto(amperes_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with AmperePerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_meterResult.Convert(units.MagnetizationAmperePerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerMeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MagnetizationDto{
        Value: 0,
        Unit:  units.MagnetizationAmperePerMeter,
    }
    
    var zeroResult *units.Magnetization
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMagnetizationFactory_FromDtoJSON(t *testing.T) {
    factory := units.MagnetizationFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "AmperePerMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "AmperePerMeter"}`)
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
    nanJSON, _ := json.Marshal(units.MagnetizationDto{
        Value: nanValue,
        Unit:  units.MagnetizationAmperePerMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with AmperePerMeter unit
    amperes_per_meterJSON := []byte(`{"value": 100, "unit": "AmperePerMeter"}`)
    amperes_per_meterResult, err := factory.FromDtoJSON(amperes_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AmperePerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_meterResult.Convert(units.MagnetizationAmperePerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerMeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "AmperePerMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromAmperesPerMeter function
func TestMagnetizationFactory_FromAmperesPerMeter(t *testing.T) {
    factory := units.MagnetizationFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAmperesPerMeter(100)
    if err != nil {
        t.Errorf("FromAmperesPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MagnetizationAmperePerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAmperesPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAmperesPerMeter(math.NaN())
    if err == nil {
        t.Error("FromAmperesPerMeter() with NaN value should return error")
    }

    _, err = factory.FromAmperesPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromAmperesPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromAmperesPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromAmperesPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAmperesPerMeter(0)
    if err != nil {
        t.Errorf("FromAmperesPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MagnetizationAmperePerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAmperesPerMeter() with zero value = %v, want 0", converted)
    }
}

func TestMagnetizationToString(t *testing.T) {
	factory := units.MagnetizationFactory{}
	a, err := factory.CreateMagnetization(45, units.MagnetizationAmperePerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MagnetizationAmperePerMeter, 2)
	expected := "45.00 " + units.GetMagnetizationAbbreviation(units.MagnetizationAmperePerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MagnetizationAmperePerMeter, -1)
	expected = "45 " + units.GetMagnetizationAbbreviation(units.MagnetizationAmperePerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMagnetization_EqualityAndComparison(t *testing.T) {
	factory := units.MagnetizationFactory{}
	a1, _ := factory.CreateMagnetization(60, units.MagnetizationAmperePerMeter)
	a2, _ := factory.CreateMagnetization(60, units.MagnetizationAmperePerMeter)
	a3, _ := factory.CreateMagnetization(120, units.MagnetizationAmperePerMeter)

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

func TestMagnetization_Arithmetic(t *testing.T) {
	factory := units.MagnetizationFactory{}
	a1, _ := factory.CreateMagnetization(30, units.MagnetizationAmperePerMeter)
	a2, _ := factory.CreateMagnetization(45, units.MagnetizationAmperePerMeter)

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


func TestGetMagnetizationAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.MagnetizationUnits
        want string
    }{
        {
            name: "AmperePerMeter abbreviation",
            unit: units.MagnetizationAmperePerMeter,
            want: "A/m",
        },
        {
            name: "invalid unit",
            unit: units.MagnetizationUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetMagnetizationAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetMagnetizationAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestMagnetization_String(t *testing.T) {
    factory := units.MagnetizationFactory{}
    
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
            unit, err := factory.CreateMagnetization(tt.value, units.MagnetizationAmperePerMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Magnetization.String() = %v, want %v", got, tt.want)
            }
        })
    }
}