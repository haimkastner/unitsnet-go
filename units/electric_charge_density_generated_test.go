// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricChargeDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CoulombPerCubicMeter"}`
	
	factory := units.ElectricChargeDensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricChargeDensityCoulombPerCubicMeter {
		t.Errorf("expected unit %v, got %v", units.ElectricChargeDensityCoulombPerCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CoulombPerCubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricChargeDensityDto_ToJSON(t *testing.T) {
	dto := units.ElectricChargeDensityDto{
		Value: 45,
		Unit:  units.ElectricChargeDensityCoulombPerCubicMeter,
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
	if result["unit"].(string) != string(units.ElectricChargeDensityCoulombPerCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.ElectricChargeDensityCoulombPerCubicMeter, result["unit"])
	}
}

func TestNewElectricChargeDensity_InvalidValue(t *testing.T) {
	factory := units.ElectricChargeDensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricChargeDensity(math.NaN(), units.ElectricChargeDensityCoulombPerCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricChargeDensity(math.Inf(1), units.ElectricChargeDensityCoulombPerCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricChargeDensityConversions(t *testing.T) {
	factory := units.ElectricChargeDensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricChargeDensity(180, units.ElectricChargeDensityCoulombPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to CoulombsPerCubicMeter.
		// No expected conversion value provided for CoulombsPerCubicMeter, verifying result is not NaN.
		result := a.CoulombsPerCubicMeter()
		cacheResult := a.CoulombsPerCubicMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CoulombsPerCubicMeter returned NaN")
		}
	}
}

func TestElectricChargeDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricChargeDensityFactory{}
	a, err := factory.CreateElectricChargeDensity(90, units.ElectricChargeDensityCoulombPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricChargeDensityCoulombPerCubicMeter {
		t.Errorf("expected default unit CoulombPerCubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricChargeDensityCoulombPerCubicMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricChargeDensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricChargeDensityCoulombPerCubicMeter {
		t.Errorf("expected unit CoulombPerCubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricChargeDensityFactory_FromDto(t *testing.T) {
    factory := units.ElectricChargeDensityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricChargeDensityDto{
        Value: 100,
        Unit:  units.ElectricChargeDensityCoulombPerCubicMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricChargeDensityDto{
        Value: math.NaN(),
        Unit:  units.ElectricChargeDensityCoulombPerCubicMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test CoulombPerCubicMeter conversion
    coulombs_per_cubic_meterDto := units.ElectricChargeDensityDto{
        Value: 100,
        Unit:  units.ElectricChargeDensityCoulombPerCubicMeter,
    }
    
    var coulombs_per_cubic_meterResult *units.ElectricChargeDensity
    coulombs_per_cubic_meterResult, err = factory.FromDto(coulombs_per_cubic_meterDto)
    if err != nil {
        t.Errorf("FromDto() with CoulombPerCubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = coulombs_per_cubic_meterResult.Convert(units.ElectricChargeDensityCoulombPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CoulombPerCubicMeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricChargeDensityDto{
        Value: 0,
        Unit:  units.ElectricChargeDensityCoulombPerCubicMeter,
    }
    
    var zeroResult *units.ElectricChargeDensity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricChargeDensityFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricChargeDensityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "CoulombPerCubicMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "CoulombPerCubicMeter"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricChargeDensityDto{
        Value: nanValue,
        Unit:  units.ElectricChargeDensityCoulombPerCubicMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with CoulombPerCubicMeter unit
    coulombs_per_cubic_meterJSON := []byte(`{"value": 100, "unit": "CoulombPerCubicMeter"}`)
    coulombs_per_cubic_meterResult, err := factory.FromDtoJSON(coulombs_per_cubic_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CoulombPerCubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = coulombs_per_cubic_meterResult.Convert(units.ElectricChargeDensityCoulombPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CoulombPerCubicMeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "CoulombPerCubicMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromCoulombsPerCubicMeter function
func TestElectricChargeDensityFactory_FromCoulombsPerCubicMeter(t *testing.T) {
    factory := units.ElectricChargeDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCoulombsPerCubicMeter(100)
    if err != nil {
        t.Errorf("FromCoulombsPerCubicMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricChargeDensityCoulombPerCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCoulombsPerCubicMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCoulombsPerCubicMeter(math.NaN())
    if err == nil {
        t.Error("FromCoulombsPerCubicMeter() with NaN value should return error")
    }

    _, err = factory.FromCoulombsPerCubicMeter(math.Inf(1))
    if err == nil {
        t.Error("FromCoulombsPerCubicMeter() with +Inf value should return error")
    }

    _, err = factory.FromCoulombsPerCubicMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromCoulombsPerCubicMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCoulombsPerCubicMeter(0)
    if err != nil {
        t.Errorf("FromCoulombsPerCubicMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricChargeDensityCoulombPerCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCoulombsPerCubicMeter() with zero value = %v, want 0", converted)
    }
}

func TestElectricChargeDensityToString(t *testing.T) {
	factory := units.ElectricChargeDensityFactory{}
	a, err := factory.CreateElectricChargeDensity(45, units.ElectricChargeDensityCoulombPerCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricChargeDensityCoulombPerCubicMeter, 2)
	expected := "45.00 " + units.GetElectricChargeDensityAbbreviation(units.ElectricChargeDensityCoulombPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricChargeDensityCoulombPerCubicMeter, -1)
	expected = "45 " + units.GetElectricChargeDensityAbbreviation(units.ElectricChargeDensityCoulombPerCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricChargeDensity_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricChargeDensityFactory{}
	a1, _ := factory.CreateElectricChargeDensity(60, units.ElectricChargeDensityCoulombPerCubicMeter)
	a2, _ := factory.CreateElectricChargeDensity(60, units.ElectricChargeDensityCoulombPerCubicMeter)
	a3, _ := factory.CreateElectricChargeDensity(120, units.ElectricChargeDensityCoulombPerCubicMeter)

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

func TestElectricChargeDensity_Arithmetic(t *testing.T) {
	factory := units.ElectricChargeDensityFactory{}
	a1, _ := factory.CreateElectricChargeDensity(30, units.ElectricChargeDensityCoulombPerCubicMeter)
	a2, _ := factory.CreateElectricChargeDensity(45, units.ElectricChargeDensityCoulombPerCubicMeter)

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


func TestGetElectricChargeDensityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricChargeDensityUnits
        want string
    }{
        {
            name: "CoulombPerCubicMeter abbreviation",
            unit: units.ElectricChargeDensityCoulombPerCubicMeter,
            want: "C/m³",
        },
        {
            name: "invalid unit",
            unit: units.ElectricChargeDensityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricChargeDensityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricChargeDensityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricChargeDensity_String(t *testing.T) {
    factory := units.ElectricChargeDensityFactory{}
    
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
            unit, err := factory.CreateElectricChargeDensity(tt.value, units.ElectricChargeDensityCoulombPerCubicMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricChargeDensity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}