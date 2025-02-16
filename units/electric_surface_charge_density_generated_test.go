// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricSurfaceChargeDensityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CoulombPerSquareMeter"}`
	
	factory := units.ElectricSurfaceChargeDensityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricSurfaceChargeDensityCoulombPerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.ElectricSurfaceChargeDensityCoulombPerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CoulombPerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricSurfaceChargeDensityDto_ToJSON(t *testing.T) {
	dto := units.ElectricSurfaceChargeDensityDto{
		Value: 45,
		Unit:  units.ElectricSurfaceChargeDensityCoulombPerSquareMeter,
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
	if result["unit"].(string) != string(units.ElectricSurfaceChargeDensityCoulombPerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.ElectricSurfaceChargeDensityCoulombPerSquareMeter, result["unit"])
	}
}

func TestNewElectricSurfaceChargeDensity_InvalidValue(t *testing.T) {
	factory := units.ElectricSurfaceChargeDensityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricSurfaceChargeDensity(math.NaN(), units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricSurfaceChargeDensity(math.Inf(1), units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricSurfaceChargeDensityConversions(t *testing.T) {
	factory := units.ElectricSurfaceChargeDensityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricSurfaceChargeDensity(180, units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to CoulombsPerSquareMeter.
		// No expected conversion value provided for CoulombsPerSquareMeter, verifying result is not NaN.
		result := a.CoulombsPerSquareMeter()
		cacheResult := a.CoulombsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CoulombsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to CoulombsPerSquareCentimeter.
		// No expected conversion value provided for CoulombsPerSquareCentimeter, verifying result is not NaN.
		result := a.CoulombsPerSquareCentimeter()
		cacheResult := a.CoulombsPerSquareCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CoulombsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to CoulombsPerSquareInch.
		// No expected conversion value provided for CoulombsPerSquareInch, verifying result is not NaN.
		result := a.CoulombsPerSquareInch()
		cacheResult := a.CoulombsPerSquareInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CoulombsPerSquareInch returned NaN")
		}
	}
}

func TestElectricSurfaceChargeDensity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricSurfaceChargeDensityFactory{}
	a, err := factory.CreateElectricSurfaceChargeDensity(90, units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricSurfaceChargeDensityCoulombPerSquareMeter {
		t.Errorf("expected default unit CoulombPerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricSurfaceChargeDensityCoulombPerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricSurfaceChargeDensityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricSurfaceChargeDensityCoulombPerSquareMeter {
		t.Errorf("expected unit CoulombPerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricSurfaceChargeDensityFactory_FromDto(t *testing.T) {
    factory := units.ElectricSurfaceChargeDensityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricSurfaceChargeDensityDto{
        Value: 100,
        Unit:  units.ElectricSurfaceChargeDensityCoulombPerSquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricSurfaceChargeDensityDto{
        Value: math.NaN(),
        Unit:  units.ElectricSurfaceChargeDensityCoulombPerSquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test CoulombPerSquareMeter conversion
    coulombs_per_square_meterDto := units.ElectricSurfaceChargeDensityDto{
        Value: 100,
        Unit:  units.ElectricSurfaceChargeDensityCoulombPerSquareMeter,
    }
    
    var coulombs_per_square_meterResult *units.ElectricSurfaceChargeDensity
    coulombs_per_square_meterResult, err = factory.FromDto(coulombs_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with CoulombPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = coulombs_per_square_meterResult.Convert(units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CoulombPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test CoulombPerSquareCentimeter conversion
    coulombs_per_square_centimeterDto := units.ElectricSurfaceChargeDensityDto{
        Value: 100,
        Unit:  units.ElectricSurfaceChargeDensityCoulombPerSquareCentimeter,
    }
    
    var coulombs_per_square_centimeterResult *units.ElectricSurfaceChargeDensity
    coulombs_per_square_centimeterResult, err = factory.FromDto(coulombs_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with CoulombPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = coulombs_per_square_centimeterResult.Convert(units.ElectricSurfaceChargeDensityCoulombPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CoulombPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test CoulombPerSquareInch conversion
    coulombs_per_square_inchDto := units.ElectricSurfaceChargeDensityDto{
        Value: 100,
        Unit:  units.ElectricSurfaceChargeDensityCoulombPerSquareInch,
    }
    
    var coulombs_per_square_inchResult *units.ElectricSurfaceChargeDensity
    coulombs_per_square_inchResult, err = factory.FromDto(coulombs_per_square_inchDto)
    if err != nil {
        t.Errorf("FromDto() with CoulombPerSquareInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = coulombs_per_square_inchResult.Convert(units.ElectricSurfaceChargeDensityCoulombPerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CoulombPerSquareInch = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricSurfaceChargeDensityDto{
        Value: 0,
        Unit:  units.ElectricSurfaceChargeDensityCoulombPerSquareMeter,
    }
    
    var zeroResult *units.ElectricSurfaceChargeDensity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricSurfaceChargeDensityFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricSurfaceChargeDensityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "CoulombPerSquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "CoulombPerSquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricSurfaceChargeDensityDto{
        Value: nanValue,
        Unit:  units.ElectricSurfaceChargeDensityCoulombPerSquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with CoulombPerSquareMeter unit
    coulombs_per_square_meterJSON := []byte(`{"value": 100, "unit": "CoulombPerSquareMeter"}`)
    coulombs_per_square_meterResult, err := factory.FromDtoJSON(coulombs_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CoulombPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = coulombs_per_square_meterResult.Convert(units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CoulombPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with CoulombPerSquareCentimeter unit
    coulombs_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "CoulombPerSquareCentimeter"}`)
    coulombs_per_square_centimeterResult, err := factory.FromDtoJSON(coulombs_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CoulombPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = coulombs_per_square_centimeterResult.Convert(units.ElectricSurfaceChargeDensityCoulombPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CoulombPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with CoulombPerSquareInch unit
    coulombs_per_square_inchJSON := []byte(`{"value": 100, "unit": "CoulombPerSquareInch"}`)
    coulombs_per_square_inchResult, err := factory.FromDtoJSON(coulombs_per_square_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CoulombPerSquareInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = coulombs_per_square_inchResult.Convert(units.ElectricSurfaceChargeDensityCoulombPerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CoulombPerSquareInch = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "CoulombPerSquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromCoulombsPerSquareMeter function
func TestElectricSurfaceChargeDensityFactory_FromCoulombsPerSquareMeter(t *testing.T) {
    factory := units.ElectricSurfaceChargeDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCoulombsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromCoulombsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCoulombsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCoulombsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromCoulombsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromCoulombsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromCoulombsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromCoulombsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromCoulombsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCoulombsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromCoulombsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCoulombsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromCoulombsPerSquareCentimeter function
func TestElectricSurfaceChargeDensityFactory_FromCoulombsPerSquareCentimeter(t *testing.T) {
    factory := units.ElectricSurfaceChargeDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCoulombsPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromCoulombsPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSurfaceChargeDensityCoulombPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCoulombsPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCoulombsPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromCoulombsPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromCoulombsPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromCoulombsPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromCoulombsPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromCoulombsPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCoulombsPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromCoulombsPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSurfaceChargeDensityCoulombPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCoulombsPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromCoulombsPerSquareInch function
func TestElectricSurfaceChargeDensityFactory_FromCoulombsPerSquareInch(t *testing.T) {
    factory := units.ElectricSurfaceChargeDensityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCoulombsPerSquareInch(100)
    if err != nil {
        t.Errorf("FromCoulombsPerSquareInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricSurfaceChargeDensityCoulombPerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCoulombsPerSquareInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCoulombsPerSquareInch(math.NaN())
    if err == nil {
        t.Error("FromCoulombsPerSquareInch() with NaN value should return error")
    }

    _, err = factory.FromCoulombsPerSquareInch(math.Inf(1))
    if err == nil {
        t.Error("FromCoulombsPerSquareInch() with +Inf value should return error")
    }

    _, err = factory.FromCoulombsPerSquareInch(math.Inf(-1))
    if err == nil {
        t.Error("FromCoulombsPerSquareInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCoulombsPerSquareInch(0)
    if err != nil {
        t.Errorf("FromCoulombsPerSquareInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricSurfaceChargeDensityCoulombPerSquareInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCoulombsPerSquareInch() with zero value = %v, want 0", converted)
    }
}

func TestElectricSurfaceChargeDensityToString(t *testing.T) {
	factory := units.ElectricSurfaceChargeDensityFactory{}
	a, err := factory.CreateElectricSurfaceChargeDensity(45, units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricSurfaceChargeDensityCoulombPerSquareMeter, 2)
	expected := "45.00 " + units.GetElectricSurfaceChargeDensityAbbreviation(units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricSurfaceChargeDensityCoulombPerSquareMeter, -1)
	expected = "45 " + units.GetElectricSurfaceChargeDensityAbbreviation(units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricSurfaceChargeDensity_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricSurfaceChargeDensityFactory{}
	a1, _ := factory.CreateElectricSurfaceChargeDensity(60, units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
	a2, _ := factory.CreateElectricSurfaceChargeDensity(60, units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
	a3, _ := factory.CreateElectricSurfaceChargeDensity(120, units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)

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

func TestElectricSurfaceChargeDensity_Arithmetic(t *testing.T) {
	factory := units.ElectricSurfaceChargeDensityFactory{}
	a1, _ := factory.CreateElectricSurfaceChargeDensity(30, units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
	a2, _ := factory.CreateElectricSurfaceChargeDensity(45, units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)

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


func TestGetElectricSurfaceChargeDensityAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ElectricSurfaceChargeDensityUnits
        want string
    }{
        {
            name: "CoulombPerSquareMeter abbreviation",
            unit: units.ElectricSurfaceChargeDensityCoulombPerSquareMeter,
            want: "C/m²",
        },
        {
            name: "CoulombPerSquareCentimeter abbreviation",
            unit: units.ElectricSurfaceChargeDensityCoulombPerSquareCentimeter,
            want: "C/cm²",
        },
        {
            name: "CoulombPerSquareInch abbreviation",
            unit: units.ElectricSurfaceChargeDensityCoulombPerSquareInch,
            want: "C/in²",
        },
        {
            name: "invalid unit",
            unit: units.ElectricSurfaceChargeDensityUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetElectricSurfaceChargeDensityAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetElectricSurfaceChargeDensityAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestElectricSurfaceChargeDensity_String(t *testing.T) {
    factory := units.ElectricSurfaceChargeDensityFactory{}
    
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
            unit, err := factory.CreateElectricSurfaceChargeDensity(tt.value, units.ElectricSurfaceChargeDensityCoulombPerSquareMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ElectricSurfaceChargeDensity.String() = %v, want %v", got, tt.want)
            }
        })
    }
}