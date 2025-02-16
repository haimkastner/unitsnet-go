// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestReactiveEnergyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltampereReactiveHour"}`
	
	factory := units.ReactiveEnergyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ReactiveEnergyVoltampereReactiveHour {
		t.Errorf("expected unit %v, got %v", units.ReactiveEnergyVoltampereReactiveHour, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltampereReactiveHour"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestReactiveEnergyDto_ToJSON(t *testing.T) {
	dto := units.ReactiveEnergyDto{
		Value: 45,
		Unit:  units.ReactiveEnergyVoltampereReactiveHour,
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
	if result["unit"].(string) != string(units.ReactiveEnergyVoltampereReactiveHour) {
		t.Errorf("expected unit %s, got %v", units.ReactiveEnergyVoltampereReactiveHour, result["unit"])
	}
}

func TestNewReactiveEnergy_InvalidValue(t *testing.T) {
	factory := units.ReactiveEnergyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateReactiveEnergy(math.NaN(), units.ReactiveEnergyVoltampereReactiveHour)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateReactiveEnergy(math.Inf(1), units.ReactiveEnergyVoltampereReactiveHour)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestReactiveEnergyConversions(t *testing.T) {
	factory := units.ReactiveEnergyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateReactiveEnergy(180, units.ReactiveEnergyVoltampereReactiveHour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltampereReactiveHours.
		// No expected conversion value provided for VoltampereReactiveHours, verifying result is not NaN.
		result := a.VoltampereReactiveHours()
		cacheResult := a.VoltampereReactiveHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltampereReactiveHours returned NaN")
		}
	}
	{
		// Test conversion to KilovoltampereReactiveHours.
		// No expected conversion value provided for KilovoltampereReactiveHours, verifying result is not NaN.
		result := a.KilovoltampereReactiveHours()
		cacheResult := a.KilovoltampereReactiveHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilovoltampereReactiveHours returned NaN")
		}
	}
	{
		// Test conversion to MegavoltampereReactiveHours.
		// No expected conversion value provided for MegavoltampereReactiveHours, verifying result is not NaN.
		result := a.MegavoltampereReactiveHours()
		cacheResult := a.MegavoltampereReactiveHours()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegavoltampereReactiveHours returned NaN")
		}
	}
}

func TestReactiveEnergy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ReactiveEnergyFactory{}
	a, err := factory.CreateReactiveEnergy(90, units.ReactiveEnergyVoltampereReactiveHour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ReactiveEnergyVoltampereReactiveHour {
		t.Errorf("expected default unit VoltampereReactiveHour, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ReactiveEnergyVoltampereReactiveHour
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ReactiveEnergyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ReactiveEnergyVoltampereReactiveHour {
		t.Errorf("expected unit VoltampereReactiveHour, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestReactiveEnergyFactory_FromDto(t *testing.T) {
    factory := units.ReactiveEnergyFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ReactiveEnergyDto{
        Value: 100,
        Unit:  units.ReactiveEnergyVoltampereReactiveHour,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ReactiveEnergyDto{
        Value: math.NaN(),
        Unit:  units.ReactiveEnergyVoltampereReactiveHour,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test VoltampereReactiveHour conversion
    voltampere_reactive_hoursDto := units.ReactiveEnergyDto{
        Value: 100,
        Unit:  units.ReactiveEnergyVoltampereReactiveHour,
    }
    
    var voltampere_reactive_hoursResult *units.ReactiveEnergy
    voltampere_reactive_hoursResult, err = factory.FromDto(voltampere_reactive_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with VoltampereReactiveHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = voltampere_reactive_hoursResult.Convert(units.ReactiveEnergyVoltampereReactiveHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltampereReactiveHour = %v, want %v", converted, 100)
    }
    // Test KilovoltampereReactiveHour conversion
    kilovoltampere_reactive_hoursDto := units.ReactiveEnergyDto{
        Value: 100,
        Unit:  units.ReactiveEnergyKilovoltampereReactiveHour,
    }
    
    var kilovoltampere_reactive_hoursResult *units.ReactiveEnergy
    kilovoltampere_reactive_hoursResult, err = factory.FromDto(kilovoltampere_reactive_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltampereReactiveHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovoltampere_reactive_hoursResult.Convert(units.ReactiveEnergyKilovoltampereReactiveHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltampereReactiveHour = %v, want %v", converted, 100)
    }
    // Test MegavoltampereReactiveHour conversion
    megavoltampere_reactive_hoursDto := units.ReactiveEnergyDto{
        Value: 100,
        Unit:  units.ReactiveEnergyMegavoltampereReactiveHour,
    }
    
    var megavoltampere_reactive_hoursResult *units.ReactiveEnergy
    megavoltampere_reactive_hoursResult, err = factory.FromDto(megavoltampere_reactive_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltampereReactiveHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavoltampere_reactive_hoursResult.Convert(units.ReactiveEnergyMegavoltampereReactiveHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltampereReactiveHour = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ReactiveEnergyDto{
        Value: 0,
        Unit:  units.ReactiveEnergyVoltampereReactiveHour,
    }
    
    var zeroResult *units.ReactiveEnergy
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestReactiveEnergyFactory_FromDtoJSON(t *testing.T) {
    factory := units.ReactiveEnergyFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "VoltampereReactiveHour"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "VoltampereReactiveHour"}`)
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
    nanJSON, _ := json.Marshal(units.ReactiveEnergyDto{
        Value: nanValue,
        Unit:  units.ReactiveEnergyVoltampereReactiveHour,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with VoltampereReactiveHour unit
    voltampere_reactive_hoursJSON := []byte(`{"value": 100, "unit": "VoltampereReactiveHour"}`)
    voltampere_reactive_hoursResult, err := factory.FromDtoJSON(voltampere_reactive_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltampereReactiveHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = voltampere_reactive_hoursResult.Convert(units.ReactiveEnergyVoltampereReactiveHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltampereReactiveHour = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltampereReactiveHour unit
    kilovoltampere_reactive_hoursJSON := []byte(`{"value": 100, "unit": "KilovoltampereReactiveHour"}`)
    kilovoltampere_reactive_hoursResult, err := factory.FromDtoJSON(kilovoltampere_reactive_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltampereReactiveHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovoltampere_reactive_hoursResult.Convert(units.ReactiveEnergyKilovoltampereReactiveHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltampereReactiveHour = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltampereReactiveHour unit
    megavoltampere_reactive_hoursJSON := []byte(`{"value": 100, "unit": "MegavoltampereReactiveHour"}`)
    megavoltampere_reactive_hoursResult, err := factory.FromDtoJSON(megavoltampere_reactive_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltampereReactiveHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavoltampere_reactive_hoursResult.Convert(units.ReactiveEnergyMegavoltampereReactiveHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltampereReactiveHour = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "VoltampereReactiveHour"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromVoltampereReactiveHours function
func TestReactiveEnergyFactory_FromVoltampereReactiveHours(t *testing.T) {
    factory := units.ReactiveEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltampereReactiveHours(100)
    if err != nil {
        t.Errorf("FromVoltampereReactiveHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReactiveEnergyVoltampereReactiveHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltampereReactiveHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltampereReactiveHours(math.NaN())
    if err == nil {
        t.Error("FromVoltampereReactiveHours() with NaN value should return error")
    }

    _, err = factory.FromVoltampereReactiveHours(math.Inf(1))
    if err == nil {
        t.Error("FromVoltampereReactiveHours() with +Inf value should return error")
    }

    _, err = factory.FromVoltampereReactiveHours(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltampereReactiveHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltampereReactiveHours(0)
    if err != nil {
        t.Errorf("FromVoltampereReactiveHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReactiveEnergyVoltampereReactiveHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltampereReactiveHours() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltampereReactiveHours function
func TestReactiveEnergyFactory_FromKilovoltampereReactiveHours(t *testing.T) {
    factory := units.ReactiveEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltampereReactiveHours(100)
    if err != nil {
        t.Errorf("FromKilovoltampereReactiveHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReactiveEnergyKilovoltampereReactiveHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltampereReactiveHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltampereReactiveHours(math.NaN())
    if err == nil {
        t.Error("FromKilovoltampereReactiveHours() with NaN value should return error")
    }

    _, err = factory.FromKilovoltampereReactiveHours(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltampereReactiveHours() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltampereReactiveHours(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltampereReactiveHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltampereReactiveHours(0)
    if err != nil {
        t.Errorf("FromKilovoltampereReactiveHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReactiveEnergyKilovoltampereReactiveHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltampereReactiveHours() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltampereReactiveHours function
func TestReactiveEnergyFactory_FromMegavoltampereReactiveHours(t *testing.T) {
    factory := units.ReactiveEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltampereReactiveHours(100)
    if err != nil {
        t.Errorf("FromMegavoltampereReactiveHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReactiveEnergyMegavoltampereReactiveHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltampereReactiveHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltampereReactiveHours(math.NaN())
    if err == nil {
        t.Error("FromMegavoltampereReactiveHours() with NaN value should return error")
    }

    _, err = factory.FromMegavoltampereReactiveHours(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltampereReactiveHours() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltampereReactiveHours(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltampereReactiveHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltampereReactiveHours(0)
    if err != nil {
        t.Errorf("FromMegavoltampereReactiveHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ReactiveEnergyMegavoltampereReactiveHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltampereReactiveHours() with zero value = %v, want 0", converted)
    }
}

func TestReactiveEnergyToString(t *testing.T) {
	factory := units.ReactiveEnergyFactory{}
	a, err := factory.CreateReactiveEnergy(45, units.ReactiveEnergyVoltampereReactiveHour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ReactiveEnergyVoltampereReactiveHour, 2)
	expected := "45.00 " + units.GetReactiveEnergyAbbreviation(units.ReactiveEnergyVoltampereReactiveHour)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ReactiveEnergyVoltampereReactiveHour, -1)
	expected = "45 " + units.GetReactiveEnergyAbbreviation(units.ReactiveEnergyVoltampereReactiveHour)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestReactiveEnergy_EqualityAndComparison(t *testing.T) {
	factory := units.ReactiveEnergyFactory{}
	a1, _ := factory.CreateReactiveEnergy(60, units.ReactiveEnergyVoltampereReactiveHour)
	a2, _ := factory.CreateReactiveEnergy(60, units.ReactiveEnergyVoltampereReactiveHour)
	a3, _ := factory.CreateReactiveEnergy(120, units.ReactiveEnergyVoltampereReactiveHour)

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

func TestReactiveEnergy_Arithmetic(t *testing.T) {
	factory := units.ReactiveEnergyFactory{}
	a1, _ := factory.CreateReactiveEnergy(30, units.ReactiveEnergyVoltampereReactiveHour)
	a2, _ := factory.CreateReactiveEnergy(45, units.ReactiveEnergyVoltampereReactiveHour)

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


func TestGetReactiveEnergyAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ReactiveEnergyUnits
        want string
    }{
        {
            name: "VoltampereReactiveHour abbreviation",
            unit: units.ReactiveEnergyVoltampereReactiveHour,
            want: "varh",
        },
        {
            name: "KilovoltampereReactiveHour abbreviation",
            unit: units.ReactiveEnergyKilovoltampereReactiveHour,
            want: "kvarh",
        },
        {
            name: "MegavoltampereReactiveHour abbreviation",
            unit: units.ReactiveEnergyMegavoltampereReactiveHour,
            want: "Mvarh",
        },
        {
            name: "invalid unit",
            unit: units.ReactiveEnergyUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetReactiveEnergyAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetReactiveEnergyAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestReactiveEnergy_String(t *testing.T) {
    factory := units.ReactiveEnergyFactory{}
    
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
            unit, err := factory.CreateReactiveEnergy(tt.value, units.ReactiveEnergyVoltampereReactiveHour)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ReactiveEnergy.String() = %v, want %v", got, tt.want)
            }
        })
    }
}