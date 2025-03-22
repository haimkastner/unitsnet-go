// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestApparentEnergyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltampereHour"}`
	
	factory := units.ApparentEnergyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ApparentEnergyVoltampereHour {
		t.Errorf("expected unit %v, got %v", units.ApparentEnergyVoltampereHour, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltampereHour"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestApparentEnergyDto_ToJSON(t *testing.T) {
	dto := units.ApparentEnergyDto{
		Value: 45,
		Unit:  units.ApparentEnergyVoltampereHour,
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
	if result["unit"].(string) != string(units.ApparentEnergyVoltampereHour) {
		t.Errorf("expected unit %s, got %v", units.ApparentEnergyVoltampereHour, result["unit"])
	}
}

func TestNewApparentEnergy_InvalidValue(t *testing.T) {
	factory := units.ApparentEnergyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateApparentEnergy(math.NaN(), units.ApparentEnergyVoltampereHour)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateApparentEnergy(math.Inf(1), units.ApparentEnergyVoltampereHour)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestApparentEnergyConversions(t *testing.T) {
	factory := units.ApparentEnergyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateApparentEnergy(180, units.ApparentEnergyVoltampereHour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltampereHours.
		// No expected conversion value provided for VoltampereHours, verifying result is not NaN.
		result := a.VoltampereHours()
		cacheResult := a.VoltampereHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltampereHours returned NaN")
		}
	}
	{
		// Test conversion to KilovoltampereHours.
		// No expected conversion value provided for KilovoltampereHours, verifying result is not NaN.
		result := a.KilovoltampereHours()
		cacheResult := a.KilovoltampereHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilovoltampereHours returned NaN")
		}
	}
	{
		// Test conversion to MegavoltampereHours.
		// No expected conversion value provided for MegavoltampereHours, verifying result is not NaN.
		result := a.MegavoltampereHours()
		cacheResult := a.MegavoltampereHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegavoltampereHours returned NaN")
		}
	}
}

func TestApparentEnergy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ApparentEnergyFactory{}
	a, err := factory.CreateApparentEnergy(90, units.ApparentEnergyVoltampereHour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ApparentEnergyVoltampereHour {
		t.Errorf("expected default unit VoltampereHour, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ApparentEnergyVoltampereHour
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ApparentEnergyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ApparentEnergyVoltampereHour {
		t.Errorf("expected unit VoltampereHour, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestApparentEnergyFactory_FromDto(t *testing.T) {
    factory := units.ApparentEnergyFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ApparentEnergyDto{
        Value: 100,
        Unit:  units.ApparentEnergyVoltampereHour,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ApparentEnergyDto{
        Value: math.NaN(),
        Unit:  units.ApparentEnergyVoltampereHour,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test VoltampereHour conversion
    voltampere_hoursDto := units.ApparentEnergyDto{
        Value: 100,
        Unit:  units.ApparentEnergyVoltampereHour,
    }
    
    var voltampere_hoursResult *units.ApparentEnergy
    voltampere_hoursResult, err = factory.FromDto(voltampere_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with VoltampereHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = voltampere_hoursResult.Convert(units.ApparentEnergyVoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltampereHour = %v, want %v", converted, 100)
    }
    // Test KilovoltampereHour conversion
    kilovoltampere_hoursDto := units.ApparentEnergyDto{
        Value: 100,
        Unit:  units.ApparentEnergyKilovoltampereHour,
    }
    
    var kilovoltampere_hoursResult *units.ApparentEnergy
    kilovoltampere_hoursResult, err = factory.FromDto(kilovoltampere_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltampereHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovoltampere_hoursResult.Convert(units.ApparentEnergyKilovoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltampereHour = %v, want %v", converted, 100)
    }
    // Test MegavoltampereHour conversion
    megavoltampere_hoursDto := units.ApparentEnergyDto{
        Value: 100,
        Unit:  units.ApparentEnergyMegavoltampereHour,
    }
    
    var megavoltampere_hoursResult *units.ApparentEnergy
    megavoltampere_hoursResult, err = factory.FromDto(megavoltampere_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltampereHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavoltampere_hoursResult.Convert(units.ApparentEnergyMegavoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltampereHour = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ApparentEnergyDto{
        Value: 0,
        Unit:  units.ApparentEnergyVoltampereHour,
    }
    
    var zeroResult *units.ApparentEnergy
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestApparentEnergyFactory_FromDtoJSON(t *testing.T) {
    factory := units.ApparentEnergyFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "VoltampereHour"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "VoltampereHour"}`)
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
    nanJSON, _ := json.Marshal(units.ApparentEnergyDto{
        Value: nanValue,
        Unit:  units.ApparentEnergyVoltampereHour,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with VoltampereHour unit
    voltampere_hoursJSON := []byte(`{"value": 100, "unit": "VoltampereHour"}`)
    voltampere_hoursResult, err := factory.FromDtoJSON(voltampere_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltampereHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = voltampere_hoursResult.Convert(units.ApparentEnergyVoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltampereHour = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltampereHour unit
    kilovoltampere_hoursJSON := []byte(`{"value": 100, "unit": "KilovoltampereHour"}`)
    kilovoltampere_hoursResult, err := factory.FromDtoJSON(kilovoltampere_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltampereHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovoltampere_hoursResult.Convert(units.ApparentEnergyKilovoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltampereHour = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltampereHour unit
    megavoltampere_hoursJSON := []byte(`{"value": 100, "unit": "MegavoltampereHour"}`)
    megavoltampere_hoursResult, err := factory.FromDtoJSON(megavoltampere_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltampereHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavoltampere_hoursResult.Convert(units.ApparentEnergyMegavoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltampereHour = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "VoltampereHour"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromVoltampereHours function
func TestApparentEnergyFactory_FromVoltampereHours(t *testing.T) {
    factory := units.ApparentEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltampereHours(100)
    if err != nil {
        t.Errorf("FromVoltampereHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ApparentEnergyVoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltampereHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltampereHours(math.NaN())
    if err == nil {
        t.Error("FromVoltampereHours() with NaN value should return error")
    }

    _, err = factory.FromVoltampereHours(math.Inf(1))
    if err == nil {
        t.Error("FromVoltampereHours() with +Inf value should return error")
    }

    _, err = factory.FromVoltampereHours(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltampereHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltampereHours(0)
    if err != nil {
        t.Errorf("FromVoltampereHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ApparentEnergyVoltampereHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltampereHours() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltampereHours function
func TestApparentEnergyFactory_FromKilovoltampereHours(t *testing.T) {
    factory := units.ApparentEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltampereHours(100)
    if err != nil {
        t.Errorf("FromKilovoltampereHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ApparentEnergyKilovoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltampereHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltampereHours(math.NaN())
    if err == nil {
        t.Error("FromKilovoltampereHours() with NaN value should return error")
    }

    _, err = factory.FromKilovoltampereHours(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltampereHours() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltampereHours(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltampereHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltampereHours(0)
    if err != nil {
        t.Errorf("FromKilovoltampereHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ApparentEnergyKilovoltampereHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltampereHours() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltampereHours function
func TestApparentEnergyFactory_FromMegavoltampereHours(t *testing.T) {
    factory := units.ApparentEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltampereHours(100)
    if err != nil {
        t.Errorf("FromMegavoltampereHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ApparentEnergyMegavoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltampereHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltampereHours(math.NaN())
    if err == nil {
        t.Error("FromMegavoltampereHours() with NaN value should return error")
    }

    _, err = factory.FromMegavoltampereHours(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltampereHours() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltampereHours(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltampereHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltampereHours(0)
    if err != nil {
        t.Errorf("FromMegavoltampereHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ApparentEnergyMegavoltampereHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltampereHours() with zero value = %v, want 0", converted)
    }
}

func TestApparentEnergyToString(t *testing.T) {
	factory := units.ApparentEnergyFactory{}
	a, err := factory.CreateApparentEnergy(45, units.ApparentEnergyVoltampereHour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ApparentEnergyVoltampereHour, 2)
	expected := "45.00 " + units.GetApparentEnergyAbbreviation(units.ApparentEnergyVoltampereHour)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ApparentEnergyVoltampereHour, -1)
	expected = "45 " + units.GetApparentEnergyAbbreviation(units.ApparentEnergyVoltampereHour)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestApparentEnergy_EqualityAndComparison(t *testing.T) {
	factory := units.ApparentEnergyFactory{}
	a1, _ := factory.CreateApparentEnergy(60, units.ApparentEnergyVoltampereHour)
	a2, _ := factory.CreateApparentEnergy(60, units.ApparentEnergyVoltampereHour)
	a3, _ := factory.CreateApparentEnergy(120, units.ApparentEnergyVoltampereHour)

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

func TestApparentEnergy_Arithmetic(t *testing.T) {
	factory := units.ApparentEnergyFactory{}
	a1, _ := factory.CreateApparentEnergy(30, units.ApparentEnergyVoltampereHour)
	a2, _ := factory.CreateApparentEnergy(45, units.ApparentEnergyVoltampereHour)

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


func TestGetApparentEnergyAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ApparentEnergyUnits
        want string
    }{
        {
            name: "VoltampereHour abbreviation",
            unit: units.ApparentEnergyVoltampereHour,
            want: "VAh",
        },
        {
            name: "KilovoltampereHour abbreviation",
            unit: units.ApparentEnergyKilovoltampereHour,
            want: "kVAh",
        },
        {
            name: "MegavoltampereHour abbreviation",
            unit: units.ApparentEnergyMegavoltampereHour,
            want: "MVAh",
        },
        {
            name: "invalid unit",
            unit: units.ApparentEnergyUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetApparentEnergyAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetApparentEnergyAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestApparentEnergy_String(t *testing.T) {
    factory := units.ApparentEnergyFactory{}
    
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
            unit, err := factory.CreateApparentEnergy(tt.value, units.ApparentEnergyVoltampereHour)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ApparentEnergy.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestApparentEnergy_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.ApparentEnergyFactory{}

	_, err := uf.CreateApparentEnergy(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}