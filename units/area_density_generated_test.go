// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAreaDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KilogramPerSquareMeter"}`
	
	factory := units.AreaDensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AreaDensityKilogramPerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.AreaDensityKilogramPerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KilogramPerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAreaDensityDto_ToJSON(t *testing.T) {
	dto := units.AreaDensityDto{
		Value: 45,
		Unit:  units.AreaDensityKilogramPerSquareMeter,
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
	if result["unit"].(string) != string(units.AreaDensityKilogramPerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.AreaDensityKilogramPerSquareMeter, result["unit"])
	}
}

func TestNewAreaDensity_InvalidValue(t *testing.T) {
	factory := units.AreaDensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAreaDensity(math.NaN(), units.AreaDensityKilogramPerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAreaDensity(math.Inf(1), units.AreaDensityKilogramPerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAreaDensityConversions(t *testing.T) {
	factory := units.AreaDensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAreaDensity(180, units.AreaDensityKilogramPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to KilogramsPerSquareMeter.
		// No expected conversion value provided for KilogramsPerSquareMeter, verifying result is not NaN.
		result := a.KilogramsPerSquareMeter()
		cacheResult := a.KilogramsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to GramsPerSquareMeter.
		// No expected conversion value provided for GramsPerSquareMeter, verifying result is not NaN.
		result := a.GramsPerSquareMeter()
		cacheResult := a.GramsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MilligramsPerSquareMeter.
		// No expected conversion value provided for MilligramsPerSquareMeter, verifying result is not NaN.
		result := a.MilligramsPerSquareMeter()
		cacheResult := a.MilligramsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MilligramsPerSquareMeter returned NaN")
		}
	}
}

func TestAreaDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AreaDensityFactory{}
	a, err := factory.CreateAreaDensity(90, units.AreaDensityKilogramPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AreaDensityKilogramPerSquareMeter {
		t.Errorf("expected default unit KilogramPerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AreaDensityKilogramPerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AreaDensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AreaDensityKilogramPerSquareMeter {
		t.Errorf("expected unit KilogramPerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAreaDensityFactory_FromDto(t *testing.T) {
    factory := units.AreaDensityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.AreaDensityDto{
        Value: 100,
        Unit:  units.AreaDensityKilogramPerSquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.AreaDensityDto{
        Value: math.NaN(),
        Unit:  units.AreaDensityKilogramPerSquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test KilogramPerSquareMeter conversion
    kilograms_per_square_meterDto := units.AreaDensityDto{
        Value: 100,
        Unit:  units.AreaDensityKilogramPerSquareMeter,
    }
    
    var kilograms_per_square_meterResult *units.AreaDensity
    kilograms_per_square_meterResult, err = factory.FromDto(kilograms_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_square_meterResult.Convert(units.AreaDensityKilogramPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test GramPerSquareMeter conversion
    grams_per_square_meterDto := units.AreaDensityDto{
        Value: 100,
        Unit:  units.AreaDensityGramPerSquareMeter,
    }
    
    var grams_per_square_meterResult *units.AreaDensity
    grams_per_square_meterResult, err = factory.FromDto(grams_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_square_meterResult.Convert(units.AreaDensityGramPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test MilligramPerSquareMeter conversion
    milligrams_per_square_meterDto := units.AreaDensityDto{
        Value: 100,
        Unit:  units.AreaDensityMilligramPerSquareMeter,
    }
    
    var milligrams_per_square_meterResult *units.AreaDensity
    milligrams_per_square_meterResult, err = factory.FromDto(milligrams_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MilligramPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_square_meterResult.Convert(units.AreaDensityMilligramPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerSquareMeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.AreaDensityDto{
        Value: 0,
        Unit:  units.AreaDensityKilogramPerSquareMeter,
    }
    
    var zeroResult *units.AreaDensity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestAreaDensityFactory_FromDtoJSON(t *testing.T) {
    factory := units.AreaDensityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "KilogramPerSquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "KilogramPerSquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.AreaDensityDto{
        Value: nanValue,
        Unit:  units.AreaDensityKilogramPerSquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with KilogramPerSquareMeter unit
    kilograms_per_square_meterJSON := []byte(`{"value": 100, "unit": "KilogramPerSquareMeter"}`)
    kilograms_per_square_meterResult, err := factory.FromDtoJSON(kilograms_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_square_meterResult.Convert(units.AreaDensityKilogramPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerSquareMeter unit
    grams_per_square_meterJSON := []byte(`{"value": 100, "unit": "GramPerSquareMeter"}`)
    grams_per_square_meterResult, err := factory.FromDtoJSON(grams_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_square_meterResult.Convert(units.AreaDensityGramPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilligramPerSquareMeter unit
    milligrams_per_square_meterJSON := []byte(`{"value": 100, "unit": "MilligramPerSquareMeter"}`)
    milligrams_per_square_meterResult, err := factory.FromDtoJSON(milligrams_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilligramPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milligrams_per_square_meterResult.Convert(units.AreaDensityMilligramPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilligramPerSquareMeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "KilogramPerSquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromKilogramsPerSquareMeter function
func TestAreaDensityFactory_FromKilogramsPerSquareMeter(t *testing.T) {
    factory := units.AreaDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromKilogramsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaDensityKilogramPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromKilogramsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaDensityKilogramPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerSquareMeter function
func TestAreaDensityFactory_FromGramsPerSquareMeter(t *testing.T) {
    factory := units.AreaDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromGramsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaDensityGramPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromGramsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromGramsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromGramsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaDensityGramPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilligramsPerSquareMeter function
func TestAreaDensityFactory_FromMilligramsPerSquareMeter(t *testing.T) {
    factory := units.AreaDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilligramsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromMilligramsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AreaDensityMilligramPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilligramsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilligramsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromMilligramsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromMilligramsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMilligramsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromMilligramsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilligramsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilligramsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromMilligramsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AreaDensityMilligramPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilligramsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}

func TestAreaDensityToString(t *testing.T) {
	factory := units.AreaDensityFactory{}
	a, err := factory.CreateAreaDensity(45, units.AreaDensityKilogramPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AreaDensityKilogramPerSquareMeter, 2)
	expected := "45.00 " + units.GetAreaDensityAbbreviation(units.AreaDensityKilogramPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AreaDensityKilogramPerSquareMeter, -1)
	expected = "45 " + units.GetAreaDensityAbbreviation(units.AreaDensityKilogramPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAreaDensity_EqualityAndComparison(t *testing.T) {
	factory := units.AreaDensityFactory{}
	a1, _ := factory.CreateAreaDensity(60, units.AreaDensityKilogramPerSquareMeter)
	a2, _ := factory.CreateAreaDensity(60, units.AreaDensityKilogramPerSquareMeter)
	a3, _ := factory.CreateAreaDensity(120, units.AreaDensityKilogramPerSquareMeter)

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

func TestAreaDensity_Arithmetic(t *testing.T) {
	factory := units.AreaDensityFactory{}
	a1, _ := factory.CreateAreaDensity(30, units.AreaDensityKilogramPerSquareMeter)
	a2, _ := factory.CreateAreaDensity(45, units.AreaDensityKilogramPerSquareMeter)

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


func TestGetAreaDensityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.AreaDensityUnits
        want string
    }{
        {
            name: "KilogramPerSquareMeter abbreviation",
            unit: units.AreaDensityKilogramPerSquareMeter,
            want: "kg/m²",
        },
        {
            name: "GramPerSquareMeter abbreviation",
            unit: units.AreaDensityGramPerSquareMeter,
            want: "g/m²",
        },
        {
            name: "MilligramPerSquareMeter abbreviation",
            unit: units.AreaDensityMilligramPerSquareMeter,
            want: "mg/m²",
        },
        {
            name: "invalid unit",
            unit: units.AreaDensityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetAreaDensityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetAreaDensityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestAreaDensity_String(t *testing.T) {
    factory := units.AreaDensityFactory{}
    
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
            unit, err := factory.CreateAreaDensity(tt.value, units.AreaDensityKilogramPerSquareMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("AreaDensity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestAreaDensity_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.AreaDensityFactory{}

	_, err := uf.CreateAreaDensity(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}