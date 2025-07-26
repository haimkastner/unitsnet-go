// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestThermalResistanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KelvinPerWatt"}`
	
	factory := units.ThermalResistanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ThermalResistanceKelvinPerWatt {
		t.Errorf("expected unit %v, got %v", units.ThermalResistanceKelvinPerWatt, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KelvinPerWatt"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestThermalResistanceDto_ToJSON(t *testing.T) {
	dto := units.ThermalResistanceDto{
		Value: 45,
		Unit:  units.ThermalResistanceKelvinPerWatt,
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
	if result["unit"].(string) != string(units.ThermalResistanceKelvinPerWatt) {
		t.Errorf("expected unit %s, got %v", units.ThermalResistanceKelvinPerWatt, result["unit"])
	}
}

func TestNewThermalResistance_InvalidValue(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateThermalResistance(math.NaN(), units.ThermalResistanceKelvinPerWatt)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateThermalResistance(math.Inf(1), units.ThermalResistanceKelvinPerWatt)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestThermalResistanceConversions(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateThermalResistance(180, units.ThermalResistanceKelvinPerWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to KelvinsPerWatt.
		// No expected conversion value provided for KelvinsPerWatt, verifying result is not NaN.
		result := a.KelvinsPerWatt()
		cacheResult := a.KelvinsPerWatt()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KelvinsPerWatt returned NaN")
		}
	}
	{
		// Test conversion to DegreesCelsiusPerWatt.
		// No expected conversion value provided for DegreesCelsiusPerWatt, verifying result is not NaN.
		result := a.DegreesCelsiusPerWatt()
		cacheResult := a.DegreesCelsiusPerWatt()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesCelsiusPerWatt returned NaN")
		}
	}
}

func TestThermalResistance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	a, err := factory.CreateThermalResistance(90, units.ThermalResistanceKelvinPerWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ThermalResistanceKelvinPerWatt {
		t.Errorf("expected default unit KelvinPerWatt, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ThermalResistanceKelvinPerWatt
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ThermalResistanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ThermalResistanceKelvinPerWatt {
		t.Errorf("expected unit KelvinPerWatt, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestThermalResistanceFactory_FromDto(t *testing.T) {
    factory := units.ThermalResistanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ThermalResistanceDto{
        Value: 100,
        Unit:  units.ThermalResistanceKelvinPerWatt,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ThermalResistanceDto{
        Value: math.NaN(),
        Unit:  units.ThermalResistanceKelvinPerWatt,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test KelvinPerWatt conversion
    kelvins_per_wattDto := units.ThermalResistanceDto{
        Value: 100,
        Unit:  units.ThermalResistanceKelvinPerWatt,
    }
    
    var kelvins_per_wattResult *units.ThermalResistance
    kelvins_per_wattResult, err = factory.FromDto(kelvins_per_wattDto)
    if err != nil {
        t.Errorf("FromDto() with KelvinPerWatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kelvins_per_wattResult.Convert(units.ThermalResistanceKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KelvinPerWatt = %v, want %v", converted, 100)
    }
    // Test DegreeCelsiusPerWatt conversion
    degrees_celsius_per_wattDto := units.ThermalResistanceDto{
        Value: 100,
        Unit:  units.ThermalResistanceDegreeCelsiusPerWatt,
    }
    
    var degrees_celsius_per_wattResult *units.ThermalResistance
    degrees_celsius_per_wattResult, err = factory.FromDto(degrees_celsius_per_wattDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeCelsiusPerWatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsius_per_wattResult.Convert(units.ThermalResistanceDegreeCelsiusPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsiusPerWatt = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ThermalResistanceDto{
        Value: 0,
        Unit:  units.ThermalResistanceKelvinPerWatt,
    }
    
    var zeroResult *units.ThermalResistance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestThermalResistanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.ThermalResistanceFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "KelvinPerWatt"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "KelvinPerWatt"}`)
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
    nanJSON, _ := json.Marshal(units.ThermalResistanceDto{
        Value: nanValue,
        Unit:  units.ThermalResistanceKelvinPerWatt,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with KelvinPerWatt unit
    kelvins_per_wattJSON := []byte(`{"value": 100, "unit": "KelvinPerWatt"}`)
    kelvins_per_wattResult, err := factory.FromDtoJSON(kelvins_per_wattJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KelvinPerWatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kelvins_per_wattResult.Convert(units.ThermalResistanceKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KelvinPerWatt = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeCelsiusPerWatt unit
    degrees_celsius_per_wattJSON := []byte(`{"value": 100, "unit": "DegreeCelsiusPerWatt"}`)
    degrees_celsius_per_wattResult, err := factory.FromDtoJSON(degrees_celsius_per_wattJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeCelsiusPerWatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsius_per_wattResult.Convert(units.ThermalResistanceDegreeCelsiusPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsiusPerWatt = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "KelvinPerWatt"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromKelvinsPerWatt function
func TestThermalResistanceFactory_FromKelvinsPerWatt(t *testing.T) {
    factory := units.ThermalResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKelvinsPerWatt(100)
    if err != nil {
        t.Errorf("FromKelvinsPerWatt() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalResistanceKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKelvinsPerWatt() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKelvinsPerWatt(math.NaN())
    if err == nil {
        t.Error("FromKelvinsPerWatt() with NaN value should return error")
    }

    _, err = factory.FromKelvinsPerWatt(math.Inf(1))
    if err == nil {
        t.Error("FromKelvinsPerWatt() with +Inf value should return error")
    }

    _, err = factory.FromKelvinsPerWatt(math.Inf(-1))
    if err == nil {
        t.Error("FromKelvinsPerWatt() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKelvinsPerWatt(0)
    if err != nil {
        t.Errorf("FromKelvinsPerWatt() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ThermalResistanceKelvinPerWatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKelvinsPerWatt() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesCelsiusPerWatt function
func TestThermalResistanceFactory_FromDegreesCelsiusPerWatt(t *testing.T) {
    factory := units.ThermalResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesCelsiusPerWatt(100)
    if err != nil {
        t.Errorf("FromDegreesCelsiusPerWatt() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalResistanceDegreeCelsiusPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesCelsiusPerWatt() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesCelsiusPerWatt(math.NaN())
    if err == nil {
        t.Error("FromDegreesCelsiusPerWatt() with NaN value should return error")
    }

    _, err = factory.FromDegreesCelsiusPerWatt(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesCelsiusPerWatt() with +Inf value should return error")
    }

    _, err = factory.FromDegreesCelsiusPerWatt(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesCelsiusPerWatt() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesCelsiusPerWatt(0)
    if err != nil {
        t.Errorf("FromDegreesCelsiusPerWatt() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ThermalResistanceDegreeCelsiusPerWatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesCelsiusPerWatt() with zero value = %v, want 0", converted)
    }
}

func TestThermalResistanceToString(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	a, err := factory.CreateThermalResistance(45, units.ThermalResistanceKelvinPerWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ThermalResistanceKelvinPerWatt, 2)
	expected := "45.00 " + units.GetThermalResistanceAbbreviation(units.ThermalResistanceKelvinPerWatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ThermalResistanceKelvinPerWatt, -1)
	expected = "45 " + units.GetThermalResistanceAbbreviation(units.ThermalResistanceKelvinPerWatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestThermalResistance_EqualityAndComparison(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	a1, _ := factory.CreateThermalResistance(60, units.ThermalResistanceKelvinPerWatt)
	a2, _ := factory.CreateThermalResistance(60, units.ThermalResistanceKelvinPerWatt)
	a3, _ := factory.CreateThermalResistance(120, units.ThermalResistanceKelvinPerWatt)

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

func TestThermalResistance_Arithmetic(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	a1, _ := factory.CreateThermalResistance(30, units.ThermalResistanceKelvinPerWatt)
	a2, _ := factory.CreateThermalResistance(45, units.ThermalResistanceKelvinPerWatt)

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


func TestGetThermalResistanceAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ThermalResistanceUnits
        want string
    }{
        {
            name: "KelvinPerWatt abbreviation",
            unit: units.ThermalResistanceKelvinPerWatt,
            want: "K/W",
        },
        {
            name: "DegreeCelsiusPerWatt abbreviation",
            unit: units.ThermalResistanceDegreeCelsiusPerWatt,
            want: "°C/W",
        },
        {
            name: "invalid unit",
            unit: units.ThermalResistanceUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetThermalResistanceAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetThermalResistanceAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestThermalResistance_String(t *testing.T) {
    factory := units.ThermalResistanceFactory{}
    
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
            unit, err := factory.CreateThermalResistance(tt.value, units.ThermalResistanceKelvinPerWatt)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ThermalResistance.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestThermalResistance_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.ThermalResistanceFactory{}

	_, err := uf.CreateThermalResistance(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}