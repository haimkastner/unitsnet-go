// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricCurrentDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "AmperePerSquareMeter"}`
	
	factory := units.ElectricCurrentDensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricCurrentDensityAmperePerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.ElectricCurrentDensityAmperePerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "AmperePerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricCurrentDensityDto_ToJSON(t *testing.T) {
	dto := units.ElectricCurrentDensityDto{
		Value: 45,
		Unit:  units.ElectricCurrentDensityAmperePerSquareMeter,
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
	if result["unit"].(string) != string(units.ElectricCurrentDensityAmperePerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.ElectricCurrentDensityAmperePerSquareMeter, result["unit"])
	}
}

func TestNewElectricCurrentDensity_InvalidValue(t *testing.T) {
	factory := units.ElectricCurrentDensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricCurrentDensity(math.NaN(), units.ElectricCurrentDensityAmperePerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricCurrentDensity(math.Inf(1), units.ElectricCurrentDensityAmperePerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricCurrentDensityConversions(t *testing.T) {
	factory := units.ElectricCurrentDensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricCurrentDensity(180, units.ElectricCurrentDensityAmperePerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to AmperesPerSquareMeter.
		// No expected conversion value provided for AmperesPerSquareMeter, verifying result is not NaN.
		result := a.AmperesPerSquareMeter()
		cacheResult := a.AmperesPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AmperesPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to AmperesPerSquareInch.
		// No expected conversion value provided for AmperesPerSquareInch, verifying result is not NaN.
		result := a.AmperesPerSquareInch()
		cacheResult := a.AmperesPerSquareInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AmperesPerSquareInch returned NaN")
		}
	}
	{
		// Test conversion to AmperesPerSquareFoot.
		// No expected conversion value provided for AmperesPerSquareFoot, verifying result is not NaN.
		result := a.AmperesPerSquareFoot()
		cacheResult := a.AmperesPerSquareFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AmperesPerSquareFoot returned NaN")
		}
	}
}

func TestElectricCurrentDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricCurrentDensityFactory{}
	a, err := factory.CreateElectricCurrentDensity(90, units.ElectricCurrentDensityAmperePerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricCurrentDensityAmperePerSquareMeter {
		t.Errorf("expected default unit AmperePerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricCurrentDensityAmperePerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricCurrentDensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricCurrentDensityAmperePerSquareMeter {
		t.Errorf("expected unit AmperePerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricCurrentDensityFactory_FromDto(t *testing.T) {
    factory := units.ElectricCurrentDensityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricCurrentDensityDto{
        Value: 100,
        Unit:  units.ElectricCurrentDensityAmperePerSquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricCurrentDensityDto{
        Value: math.NaN(),
        Unit:  units.ElectricCurrentDensityAmperePerSquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test AmperePerSquareMeter conversion
    amperes_per_square_meterDto := units.ElectricCurrentDensityDto{
        Value: 100,
        Unit:  units.ElectricCurrentDensityAmperePerSquareMeter,
    }
    
    var amperes_per_square_meterResult *units.ElectricCurrentDensity
    amperes_per_square_meterResult, err = factory.FromDto(amperes_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with AmperePerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_square_meterResult.Convert(units.ElectricCurrentDensityAmperePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerSquareMeter = %v, want %v", converted, 100)
    }
    // Test AmperePerSquareInch conversion
    amperes_per_square_inchDto := units.ElectricCurrentDensityDto{
        Value: 100,
        Unit:  units.ElectricCurrentDensityAmperePerSquareInch,
    }
    
    var amperes_per_square_inchResult *units.ElectricCurrentDensity
    amperes_per_square_inchResult, err = factory.FromDto(amperes_per_square_inchDto)
    if err != nil {
        t.Errorf("FromDto() with AmperePerSquareInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_square_inchResult.Convert(units.ElectricCurrentDensityAmperePerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerSquareInch = %v, want %v", converted, 100)
    }
    // Test AmperePerSquareFoot conversion
    amperes_per_square_footDto := units.ElectricCurrentDensityDto{
        Value: 100,
        Unit:  units.ElectricCurrentDensityAmperePerSquareFoot,
    }
    
    var amperes_per_square_footResult *units.ElectricCurrentDensity
    amperes_per_square_footResult, err = factory.FromDto(amperes_per_square_footDto)
    if err != nil {
        t.Errorf("FromDto() with AmperePerSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_square_footResult.Convert(units.ElectricCurrentDensityAmperePerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerSquareFoot = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricCurrentDensityDto{
        Value: 0,
        Unit:  units.ElectricCurrentDensityAmperePerSquareMeter,
    }
    
    var zeroResult *units.ElectricCurrentDensity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricCurrentDensityFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricCurrentDensityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "AmperePerSquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "AmperePerSquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricCurrentDensityDto{
        Value: nanValue,
        Unit:  units.ElectricCurrentDensityAmperePerSquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with AmperePerSquareMeter unit
    amperes_per_square_meterJSON := []byte(`{"value": 100, "unit": "AmperePerSquareMeter"}`)
    amperes_per_square_meterResult, err := factory.FromDtoJSON(amperes_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AmperePerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_square_meterResult.Convert(units.ElectricCurrentDensityAmperePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with AmperePerSquareInch unit
    amperes_per_square_inchJSON := []byte(`{"value": 100, "unit": "AmperePerSquareInch"}`)
    amperes_per_square_inchResult, err := factory.FromDtoJSON(amperes_per_square_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AmperePerSquareInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_square_inchResult.Convert(units.ElectricCurrentDensityAmperePerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerSquareInch = %v, want %v", converted, 100)
    }
    // Test JSON with AmperePerSquareFoot unit
    amperes_per_square_footJSON := []byte(`{"value": 100, "unit": "AmperePerSquareFoot"}`)
    amperes_per_square_footResult, err := factory.FromDtoJSON(amperes_per_square_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AmperePerSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = amperes_per_square_footResult.Convert(units.ElectricCurrentDensityAmperePerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AmperePerSquareFoot = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "AmperePerSquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromAmperesPerSquareMeter function
func TestElectricCurrentDensityFactory_FromAmperesPerSquareMeter(t *testing.T) {
    factory := units.ElectricCurrentDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAmperesPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromAmperesPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentDensityAmperePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAmperesPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAmperesPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromAmperesPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromAmperesPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromAmperesPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromAmperesPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromAmperesPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAmperesPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromAmperesPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentDensityAmperePerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAmperesPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromAmperesPerSquareInch function
func TestElectricCurrentDensityFactory_FromAmperesPerSquareInch(t *testing.T) {
    factory := units.ElectricCurrentDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAmperesPerSquareInch(100)
    if err != nil {
        t.Errorf("FromAmperesPerSquareInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentDensityAmperePerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAmperesPerSquareInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAmperesPerSquareInch(math.NaN())
    if err == nil {
        t.Error("FromAmperesPerSquareInch() with NaN value should return error")
    }

    _, err = factory.FromAmperesPerSquareInch(math.Inf(1))
    if err == nil {
        t.Error("FromAmperesPerSquareInch() with +Inf value should return error")
    }

    _, err = factory.FromAmperesPerSquareInch(math.Inf(-1))
    if err == nil {
        t.Error("FromAmperesPerSquareInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAmperesPerSquareInch(0)
    if err != nil {
        t.Errorf("FromAmperesPerSquareInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentDensityAmperePerSquareInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAmperesPerSquareInch() with zero value = %v, want 0", converted)
    }
}
// Test FromAmperesPerSquareFoot function
func TestElectricCurrentDensityFactory_FromAmperesPerSquareFoot(t *testing.T) {
    factory := units.ElectricCurrentDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAmperesPerSquareFoot(100)
    if err != nil {
        t.Errorf("FromAmperesPerSquareFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricCurrentDensityAmperePerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAmperesPerSquareFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAmperesPerSquareFoot(math.NaN())
    if err == nil {
        t.Error("FromAmperesPerSquareFoot() with NaN value should return error")
    }

    _, err = factory.FromAmperesPerSquareFoot(math.Inf(1))
    if err == nil {
        t.Error("FromAmperesPerSquareFoot() with +Inf value should return error")
    }

    _, err = factory.FromAmperesPerSquareFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromAmperesPerSquareFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAmperesPerSquareFoot(0)
    if err != nil {
        t.Errorf("FromAmperesPerSquareFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricCurrentDensityAmperePerSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAmperesPerSquareFoot() with zero value = %v, want 0", converted)
    }
}

func TestElectricCurrentDensityToString(t *testing.T) {
	factory := units.ElectricCurrentDensityFactory{}
	a, err := factory.CreateElectricCurrentDensity(45, units.ElectricCurrentDensityAmperePerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricCurrentDensityAmperePerSquareMeter, 2)
	expected := "45.00 " + units.GetElectricCurrentDensityAbbreviation(units.ElectricCurrentDensityAmperePerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricCurrentDensityAmperePerSquareMeter, -1)
	expected = "45 " + units.GetElectricCurrentDensityAbbreviation(units.ElectricCurrentDensityAmperePerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricCurrentDensity_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricCurrentDensityFactory{}
	a1, _ := factory.CreateElectricCurrentDensity(60, units.ElectricCurrentDensityAmperePerSquareMeter)
	a2, _ := factory.CreateElectricCurrentDensity(60, units.ElectricCurrentDensityAmperePerSquareMeter)
	a3, _ := factory.CreateElectricCurrentDensity(120, units.ElectricCurrentDensityAmperePerSquareMeter)

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

func TestElectricCurrentDensity_Arithmetic(t *testing.T) {
	factory := units.ElectricCurrentDensityFactory{}
	a1, _ := factory.CreateElectricCurrentDensity(30, units.ElectricCurrentDensityAmperePerSquareMeter)
	a2, _ := factory.CreateElectricCurrentDensity(45, units.ElectricCurrentDensityAmperePerSquareMeter)

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


func TestGetElectricCurrentDensityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricCurrentDensityUnits
        want string
    }{
        {
            name: "AmperePerSquareMeter abbreviation",
            unit: units.ElectricCurrentDensityAmperePerSquareMeter,
            want: "A/m²",
        },
        {
            name: "AmperePerSquareInch abbreviation",
            unit: units.ElectricCurrentDensityAmperePerSquareInch,
            want: "A/in²",
        },
        {
            name: "AmperePerSquareFoot abbreviation",
            unit: units.ElectricCurrentDensityAmperePerSquareFoot,
            want: "A/ft²",
        },
        {
            name: "invalid unit",
            unit: units.ElectricCurrentDensityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricCurrentDensityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricCurrentDensityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricCurrentDensity_String(t *testing.T) {
    factory := units.ElectricCurrentDensityFactory{}
    
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
            unit, err := factory.CreateElectricCurrentDensity(tt.value, units.ElectricCurrentDensityAmperePerSquareMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricCurrentDensity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}