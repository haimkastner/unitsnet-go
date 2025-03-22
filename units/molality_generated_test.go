// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMolalityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MolePerKilogram"}`
	
	factory := units.MolalityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MolalityMolePerKilogram {
		t.Errorf("expected unit %v, got %v", units.MolalityMolePerKilogram, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MolePerKilogram"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMolalityDto_ToJSON(t *testing.T) {
	dto := units.MolalityDto{
		Value: 45,
		Unit:  units.MolalityMolePerKilogram,
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
	if result["unit"].(string) != string(units.MolalityMolePerKilogram) {
		t.Errorf("expected unit %s, got %v", units.MolalityMolePerKilogram, result["unit"])
	}
}

func TestNewMolality_InvalidValue(t *testing.T) {
	factory := units.MolalityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMolality(math.NaN(), units.MolalityMolePerKilogram)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMolality(math.Inf(1), units.MolalityMolePerKilogram)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMolalityConversions(t *testing.T) {
	factory := units.MolalityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMolality(180, units.MolalityMolePerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MolesPerKilogram.
		// No expected conversion value provided for MolesPerKilogram, verifying result is not NaN.
		result := a.MolesPerKilogram()
		cacheResult := a.MolesPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MolesPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MolesPerGram.
		// No expected conversion value provided for MolesPerGram, verifying result is not NaN.
		result := a.MolesPerGram()
		cacheResult := a.MolesPerGram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MolesPerGram returned NaN")
		}
	}
	{
		// Test conversion to MillimolesPerKilogram.
		// No expected conversion value provided for MillimolesPerKilogram, verifying result is not NaN.
		result := a.MillimolesPerKilogram()
		cacheResult := a.MillimolesPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimolesPerKilogram returned NaN")
		}
	}
}

func TestMolality_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MolalityFactory{}
	a, err := factory.CreateMolality(90, units.MolalityMolePerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MolalityMolePerKilogram {
		t.Errorf("expected default unit MolePerKilogram, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MolalityMolePerKilogram
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MolalityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MolalityMolePerKilogram {
		t.Errorf("expected unit MolePerKilogram, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMolalityFactory_FromDto(t *testing.T) {
    factory := units.MolalityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MolalityDto{
        Value: 100,
        Unit:  units.MolalityMolePerKilogram,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MolalityDto{
        Value: math.NaN(),
        Unit:  units.MolalityMolePerKilogram,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test MolePerKilogram conversion
    moles_per_kilogramDto := units.MolalityDto{
        Value: 100,
        Unit:  units.MolalityMolePerKilogram,
    }
    
    var moles_per_kilogramResult *units.Molality
    moles_per_kilogramResult, err = factory.FromDto(moles_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with MolePerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_kilogramResult.Convert(units.MolalityMolePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerKilogram = %v, want %v", converted, 100)
    }
    // Test MolePerGram conversion
    moles_per_gramDto := units.MolalityDto{
        Value: 100,
        Unit:  units.MolalityMolePerGram,
    }
    
    var moles_per_gramResult *units.Molality
    moles_per_gramResult, err = factory.FromDto(moles_per_gramDto)
    if err != nil {
        t.Errorf("FromDto() with MolePerGram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_gramResult.Convert(units.MolalityMolePerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerGram = %v, want %v", converted, 100)
    }
    // Test MillimolePerKilogram conversion
    millimoles_per_kilogramDto := units.MolalityDto{
        Value: 100,
        Unit:  units.MolalityMillimolePerKilogram,
    }
    
    var millimoles_per_kilogramResult *units.Molality
    millimoles_per_kilogramResult, err = factory.FromDto(millimoles_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with MillimolePerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimoles_per_kilogramResult.Convert(units.MolalityMillimolePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimolePerKilogram = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MolalityDto{
        Value: 0,
        Unit:  units.MolalityMolePerKilogram,
    }
    
    var zeroResult *units.Molality
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMolalityFactory_FromDtoJSON(t *testing.T) {
    factory := units.MolalityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "MolePerKilogram"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "MolePerKilogram"}`)
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
    nanJSON, _ := json.Marshal(units.MolalityDto{
        Value: nanValue,
        Unit:  units.MolalityMolePerKilogram,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with MolePerKilogram unit
    moles_per_kilogramJSON := []byte(`{"value": 100, "unit": "MolePerKilogram"}`)
    moles_per_kilogramResult, err := factory.FromDtoJSON(moles_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MolePerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_kilogramResult.Convert(units.MolalityMolePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with MolePerGram unit
    moles_per_gramJSON := []byte(`{"value": 100, "unit": "MolePerGram"}`)
    moles_per_gramResult, err := factory.FromDtoJSON(moles_per_gramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MolePerGram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = moles_per_gramResult.Convert(units.MolalityMolePerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MolePerGram = %v, want %v", converted, 100)
    }
    // Test JSON with MillimolePerKilogram unit
    millimoles_per_kilogramJSON := []byte(`{"value": 100, "unit": "MillimolePerKilogram"}`)
    millimoles_per_kilogramResult, err := factory.FromDtoJSON(millimoles_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimolePerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimoles_per_kilogramResult.Convert(units.MolalityMillimolePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimolePerKilogram = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "MolePerKilogram"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromMolesPerKilogram function
func TestMolalityFactory_FromMolesPerKilogram(t *testing.T) {
    factory := units.MolalityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMolesPerKilogram(100)
    if err != nil {
        t.Errorf("FromMolesPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolalityMolePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMolesPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMolesPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromMolesPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromMolesPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromMolesPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromMolesPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromMolesPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMolesPerKilogram(0)
    if err != nil {
        t.Errorf("FromMolesPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolalityMolePerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMolesPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromMolesPerGram function
func TestMolalityFactory_FromMolesPerGram(t *testing.T) {
    factory := units.MolalityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMolesPerGram(100)
    if err != nil {
        t.Errorf("FromMolesPerGram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolalityMolePerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMolesPerGram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMolesPerGram(math.NaN())
    if err == nil {
        t.Error("FromMolesPerGram() with NaN value should return error")
    }

    _, err = factory.FromMolesPerGram(math.Inf(1))
    if err == nil {
        t.Error("FromMolesPerGram() with +Inf value should return error")
    }

    _, err = factory.FromMolesPerGram(math.Inf(-1))
    if err == nil {
        t.Error("FromMolesPerGram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMolesPerGram(0)
    if err != nil {
        t.Errorf("FromMolesPerGram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolalityMolePerGram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMolesPerGram() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimolesPerKilogram function
func TestMolalityFactory_FromMillimolesPerKilogram(t *testing.T) {
    factory := units.MolalityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimolesPerKilogram(100)
    if err != nil {
        t.Errorf("FromMillimolesPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolalityMillimolePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimolesPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimolesPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromMillimolesPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromMillimolesPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromMillimolesPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromMillimolesPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimolesPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimolesPerKilogram(0)
    if err != nil {
        t.Errorf("FromMillimolesPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolalityMillimolePerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimolesPerKilogram() with zero value = %v, want 0", converted)
    }
}

func TestMolalityToString(t *testing.T) {
	factory := units.MolalityFactory{}
	a, err := factory.CreateMolality(45, units.MolalityMolePerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MolalityMolePerKilogram, 2)
	expected := "45.00 " + units.GetMolalityAbbreviation(units.MolalityMolePerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MolalityMolePerKilogram, -1)
	expected = "45 " + units.GetMolalityAbbreviation(units.MolalityMolePerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMolality_EqualityAndComparison(t *testing.T) {
	factory := units.MolalityFactory{}
	a1, _ := factory.CreateMolality(60, units.MolalityMolePerKilogram)
	a2, _ := factory.CreateMolality(60, units.MolalityMolePerKilogram)
	a3, _ := factory.CreateMolality(120, units.MolalityMolePerKilogram)

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

func TestMolality_Arithmetic(t *testing.T) {
	factory := units.MolalityFactory{}
	a1, _ := factory.CreateMolality(30, units.MolalityMolePerKilogram)
	a2, _ := factory.CreateMolality(45, units.MolalityMolePerKilogram)

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


func TestGetMolalityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.MolalityUnits
        want string
    }{
        {
            name: "MolePerKilogram abbreviation",
            unit: units.MolalityMolePerKilogram,
            want: "mol/kg",
        },
        {
            name: "MolePerGram abbreviation",
            unit: units.MolalityMolePerGram,
            want: "mol/g",
        },
        {
            name: "MillimolePerKilogram abbreviation",
            unit: units.MolalityMillimolePerKilogram,
            want: "mmol/kg",
        },
        {
            name: "invalid unit",
            unit: units.MolalityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetMolalityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetMolalityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestMolality_String(t *testing.T) {
    factory := units.MolalityFactory{}
    
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
            unit, err := factory.CreateMolality(tt.value, units.MolalityMolePerKilogram)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Molality.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestMolality_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.MolalityFactory{}

	_, err := uf.CreateMolality(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}