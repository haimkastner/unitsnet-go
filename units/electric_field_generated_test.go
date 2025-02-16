// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricFieldDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltPerMeter"}`
	
	factory := units.ElectricFieldDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricFieldVoltPerMeter {
		t.Errorf("expected unit %v, got %v", units.ElectricFieldVoltPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricFieldDto_ToJSON(t *testing.T) {
	dto := units.ElectricFieldDto{
		Value: 45,
		Unit:  units.ElectricFieldVoltPerMeter,
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
	if result["unit"].(string) != string(units.ElectricFieldVoltPerMeter) {
		t.Errorf("expected unit %s, got %v", units.ElectricFieldVoltPerMeter, result["unit"])
	}
}

func TestNewElectricField_InvalidValue(t *testing.T) {
	factory := units.ElectricFieldFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricField(math.NaN(), units.ElectricFieldVoltPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricField(math.Inf(1), units.ElectricFieldVoltPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricFieldConversions(t *testing.T) {
	factory := units.ElectricFieldFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricField(180, units.ElectricFieldVoltPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltsPerMeter.
		// No expected conversion value provided for VoltsPerMeter, verifying result is not NaN.
		result := a.VoltsPerMeter()
		cacheResult := a.VoltsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltsPerMeter returned NaN")
		}
	}
}

func TestElectricField_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricFieldFactory{}
	a, err := factory.CreateElectricField(90, units.ElectricFieldVoltPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricFieldVoltPerMeter {
		t.Errorf("expected default unit VoltPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricFieldVoltPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricFieldDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricFieldVoltPerMeter {
		t.Errorf("expected unit VoltPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricFieldFactory_FromDto(t *testing.T) {
    factory := units.ElectricFieldFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricFieldDto{
        Value: 100,
        Unit:  units.ElectricFieldVoltPerMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricFieldDto{
        Value: math.NaN(),
        Unit:  units.ElectricFieldVoltPerMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test VoltPerMeter conversion
    volts_per_meterDto := units.ElectricFieldDto{
        Value: 100,
        Unit:  units.ElectricFieldVoltPerMeter,
    }
    
    var volts_per_meterResult *units.ElectricField
    volts_per_meterResult, err = factory.FromDto(volts_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with VoltPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_meterResult.Convert(units.ElectricFieldVoltPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerMeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricFieldDto{
        Value: 0,
        Unit:  units.ElectricFieldVoltPerMeter,
    }
    
    var zeroResult *units.ElectricField
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricFieldFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricFieldFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "VoltPerMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "VoltPerMeter"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricFieldDto{
        Value: nanValue,
        Unit:  units.ElectricFieldVoltPerMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with VoltPerMeter unit
    volts_per_meterJSON := []byte(`{"value": 100, "unit": "VoltPerMeter"}`)
    volts_per_meterResult, err := factory.FromDtoJSON(volts_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = volts_per_meterResult.Convert(units.ElectricFieldVoltPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltPerMeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "VoltPerMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromVoltsPerMeter function
func TestElectricFieldFactory_FromVoltsPerMeter(t *testing.T) {
    factory := units.ElectricFieldFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltsPerMeter(100)
    if err != nil {
        t.Errorf("FromVoltsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricFieldVoltPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromVoltsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromVoltsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromVoltsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromVoltsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltsPerMeter(0)
    if err != nil {
        t.Errorf("FromVoltsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricFieldVoltPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltsPerMeter() with zero value = %v, want 0", converted)
    }
}

func TestElectricFieldToString(t *testing.T) {
	factory := units.ElectricFieldFactory{}
	a, err := factory.CreateElectricField(45, units.ElectricFieldVoltPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricFieldVoltPerMeter, 2)
	expected := "45.00 " + units.GetElectricFieldAbbreviation(units.ElectricFieldVoltPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricFieldVoltPerMeter, -1)
	expected = "45 " + units.GetElectricFieldAbbreviation(units.ElectricFieldVoltPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricField_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricFieldFactory{}
	a1, _ := factory.CreateElectricField(60, units.ElectricFieldVoltPerMeter)
	a2, _ := factory.CreateElectricField(60, units.ElectricFieldVoltPerMeter)
	a3, _ := factory.CreateElectricField(120, units.ElectricFieldVoltPerMeter)

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

func TestElectricField_Arithmetic(t *testing.T) {
	factory := units.ElectricFieldFactory{}
	a1, _ := factory.CreateElectricField(30, units.ElectricFieldVoltPerMeter)
	a2, _ := factory.CreateElectricField(45, units.ElectricFieldVoltPerMeter)

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


func TestGetElectricFieldAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricFieldUnits
        want string
    }{
        {
            name: "VoltPerMeter abbreviation",
            unit: units.ElectricFieldVoltPerMeter,
            want: "V/m",
        },
        {
            name: "invalid unit",
            unit: units.ElectricFieldUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricFieldAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricFieldAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricField_String(t *testing.T) {
    factory := units.ElectricFieldFactory{}
    
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
            unit, err := factory.CreateElectricField(tt.value, units.ElectricFieldVoltPerMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricField.String() = %v, want %v", got, tt.want)
            }
        })
    }
}