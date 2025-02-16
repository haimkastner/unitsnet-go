// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestVolumeFlowPerAreaDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CubicMeterPerSecondPerSquareMeter"}`
	
	factory := units.VolumeFlowPerAreaDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CubicMeterPerSecondPerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestVolumeFlowPerAreaDto_ToJSON(t *testing.T) {
	dto := units.VolumeFlowPerAreaDto{
		Value: 45,
		Unit:  units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter,
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
	if result["unit"].(string) != string(units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter, result["unit"])
	}
}

func TestNewVolumeFlowPerArea_InvalidValue(t *testing.T) {
	factory := units.VolumeFlowPerAreaFactory{}
	// NaN value should return an error.
	_, err := factory.CreateVolumeFlowPerArea(math.NaN(), units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateVolumeFlowPerArea(math.Inf(1), units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestVolumeFlowPerAreaConversions(t *testing.T) {
	factory := units.VolumeFlowPerAreaFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateVolumeFlowPerArea(180, units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to CubicMetersPerSecondPerSquareMeter.
		// No expected conversion value provided for CubicMetersPerSecondPerSquareMeter, verifying result is not NaN.
		result := a.CubicMetersPerSecondPerSquareMeter()
		cacheResult := a.CubicMetersPerSecondPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicMetersPerSecondPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to CubicFeetPerMinutePerSquareFoot.
		// No expected conversion value provided for CubicFeetPerMinutePerSquareFoot, verifying result is not NaN.
		result := a.CubicFeetPerMinutePerSquareFoot()
		cacheResult := a.CubicFeetPerMinutePerSquareFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicFeetPerMinutePerSquareFoot returned NaN")
		}
	}
}

func TestVolumeFlowPerArea_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.VolumeFlowPerAreaFactory{}
	a, err := factory.CreateVolumeFlowPerArea(90, units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter {
		t.Errorf("expected default unit CubicMeterPerSecondPerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.VolumeFlowPerAreaDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter {
		t.Errorf("expected unit CubicMeterPerSecondPerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestVolumeFlowPerAreaFactory_FromDto(t *testing.T) {
    factory := units.VolumeFlowPerAreaFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.VolumeFlowPerAreaDto{
        Value: 100,
        Unit:  units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.VolumeFlowPerAreaDto{
        Value: math.NaN(),
        Unit:  units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test CubicMeterPerSecondPerSquareMeter conversion
    cubic_meters_per_second_per_square_meterDto := units.VolumeFlowPerAreaDto{
        Value: 100,
        Unit:  units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter,
    }
    
    var cubic_meters_per_second_per_square_meterResult *units.VolumeFlowPerArea
    cubic_meters_per_second_per_square_meterResult, err = factory.FromDto(cubic_meters_per_second_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with CubicMeterPerSecondPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_second_per_square_meterResult.Convert(units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerSecondPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test CubicFootPerMinutePerSquareFoot conversion
    cubic_feet_per_minute_per_square_footDto := units.VolumeFlowPerAreaDto{
        Value: 100,
        Unit:  units.VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot,
    }
    
    var cubic_feet_per_minute_per_square_footResult *units.VolumeFlowPerArea
    cubic_feet_per_minute_per_square_footResult, err = factory.FromDto(cubic_feet_per_minute_per_square_footDto)
    if err != nil {
        t.Errorf("FromDto() with CubicFootPerMinutePerSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_feet_per_minute_per_square_footResult.Convert(units.VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicFootPerMinutePerSquareFoot = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.VolumeFlowPerAreaDto{
        Value: 0,
        Unit:  units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter,
    }
    
    var zeroResult *units.VolumeFlowPerArea
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestVolumeFlowPerAreaFactory_FromDtoJSON(t *testing.T) {
    factory := units.VolumeFlowPerAreaFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "CubicMeterPerSecondPerSquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "CubicMeterPerSecondPerSquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.VolumeFlowPerAreaDto{
        Value: nanValue,
        Unit:  units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with CubicMeterPerSecondPerSquareMeter unit
    cubic_meters_per_second_per_square_meterJSON := []byte(`{"value": 100, "unit": "CubicMeterPerSecondPerSquareMeter"}`)
    cubic_meters_per_second_per_square_meterResult, err := factory.FromDtoJSON(cubic_meters_per_second_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicMeterPerSecondPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_second_per_square_meterResult.Convert(units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerSecondPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with CubicFootPerMinutePerSquareFoot unit
    cubic_feet_per_minute_per_square_footJSON := []byte(`{"value": 100, "unit": "CubicFootPerMinutePerSquareFoot"}`)
    cubic_feet_per_minute_per_square_footResult, err := factory.FromDtoJSON(cubic_feet_per_minute_per_square_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicFootPerMinutePerSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_feet_per_minute_per_square_footResult.Convert(units.VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicFootPerMinutePerSquareFoot = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "CubicMeterPerSecondPerSquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromCubicMetersPerSecondPerSquareMeter function
func TestVolumeFlowPerAreaFactory_FromCubicMetersPerSecondPerSquareMeter(t *testing.T) {
    factory := units.VolumeFlowPerAreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicMetersPerSecondPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromCubicMetersPerSecondPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicMetersPerSecondPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicMetersPerSecondPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromCubicMetersPerSecondPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromCubicMetersPerSecondPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromCubicMetersPerSecondPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromCubicMetersPerSecondPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicMetersPerSecondPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicMetersPerSecondPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromCubicMetersPerSecondPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicMetersPerSecondPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicFeetPerMinutePerSquareFoot function
func TestVolumeFlowPerAreaFactory_FromCubicFeetPerMinutePerSquareFoot(t *testing.T) {
    factory := units.VolumeFlowPerAreaFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicFeetPerMinutePerSquareFoot(100)
    if err != nil {
        t.Errorf("FromCubicFeetPerMinutePerSquareFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicFeetPerMinutePerSquareFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicFeetPerMinutePerSquareFoot(math.NaN())
    if err == nil {
        t.Error("FromCubicFeetPerMinutePerSquareFoot() with NaN value should return error")
    }

    _, err = factory.FromCubicFeetPerMinutePerSquareFoot(math.Inf(1))
    if err == nil {
        t.Error("FromCubicFeetPerMinutePerSquareFoot() with +Inf value should return error")
    }

    _, err = factory.FromCubicFeetPerMinutePerSquareFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicFeetPerMinutePerSquareFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicFeetPerMinutePerSquareFoot(0)
    if err != nil {
        t.Errorf("FromCubicFeetPerMinutePerSquareFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicFeetPerMinutePerSquareFoot() with zero value = %v, want 0", converted)
    }
}

func TestVolumeFlowPerAreaToString(t *testing.T) {
	factory := units.VolumeFlowPerAreaFactory{}
	a, err := factory.CreateVolumeFlowPerArea(45, units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter, 2)
	expected := "45.00 " + units.GetVolumeFlowPerAreaAbbreviation(units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter, -1)
	expected = "45 " + units.GetVolumeFlowPerAreaAbbreviation(units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestVolumeFlowPerArea_EqualityAndComparison(t *testing.T) {
	factory := units.VolumeFlowPerAreaFactory{}
	a1, _ := factory.CreateVolumeFlowPerArea(60, units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
	a2, _ := factory.CreateVolumeFlowPerArea(60, units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
	a3, _ := factory.CreateVolumeFlowPerArea(120, units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)

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

func TestVolumeFlowPerArea_Arithmetic(t *testing.T) {
	factory := units.VolumeFlowPerAreaFactory{}
	a1, _ := factory.CreateVolumeFlowPerArea(30, units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
	a2, _ := factory.CreateVolumeFlowPerArea(45, units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)

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


func TestGetVolumeFlowPerAreaAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.VolumeFlowPerAreaUnits
        want string
    }{
        {
            name: "CubicMeterPerSecondPerSquareMeter abbreviation",
            unit: units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter,
            want: "m³/(s·m²)",
        },
        {
            name: "CubicFootPerMinutePerSquareFoot abbreviation",
            unit: units.VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot,
            want: "CFM/ft²",
        },
        {
            name: "invalid unit",
            unit: units.VolumeFlowPerAreaUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetVolumeFlowPerAreaAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetVolumeFlowPerAreaAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestVolumeFlowPerArea_String(t *testing.T) {
    factory := units.VolumeFlowPerAreaFactory{}
    
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
            unit, err := factory.CreateVolumeFlowPerArea(tt.value, units.VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("VolumeFlowPerArea.String() = %v, want %v", got, tt.want)
            }
        })
    }
}