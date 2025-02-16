// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestVitaminADtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "InternationalUnit"}`
	
	factory := units.VitaminADtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.VitaminAInternationalUnit {
		t.Errorf("expected unit %v, got %v", units.VitaminAInternationalUnit, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "InternationalUnit"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestVitaminADto_ToJSON(t *testing.T) {
	dto := units.VitaminADto{
		Value: 45,
		Unit:  units.VitaminAInternationalUnit,
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
	if result["unit"].(string) != string(units.VitaminAInternationalUnit) {
		t.Errorf("expected unit %s, got %v", units.VitaminAInternationalUnit, result["unit"])
	}
}

func TestNewVitaminA_InvalidValue(t *testing.T) {
	factory := units.VitaminAFactory{}
	// NaN value should return an error.
	_, err := factory.CreateVitaminA(math.NaN(), units.VitaminAInternationalUnit)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateVitaminA(math.Inf(1), units.VitaminAInternationalUnit)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestVitaminAConversions(t *testing.T) {
	factory := units.VitaminAFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateVitaminA(180, units.VitaminAInternationalUnit)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to InternationalUnits.
		// No expected conversion value provided for InternationalUnits, verifying result is not NaN.
		result := a.InternationalUnits()
		if math.IsNaN(result) {
			t.Errorf("conversion to InternationalUnits returned NaN")
		}
	}
}

func TestVitaminA_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.VitaminAFactory{}
	a, err := factory.CreateVitaminA(90, units.VitaminAInternationalUnit)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.VitaminAInternationalUnit {
		t.Errorf("expected default unit InternationalUnit, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.VitaminAInternationalUnit
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.VitaminADto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.VitaminAInternationalUnit {
		t.Errorf("expected unit InternationalUnit, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestVitaminAFactory_FromDto(t *testing.T) {
    factory := units.VitaminAFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.VitaminADto{
        Value: 100,
        Unit:  units.VitaminAInternationalUnit,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.VitaminADto{
        Value: math.NaN(),
        Unit:  units.VitaminAInternationalUnit,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test InternationalUnit conversion
    international_unitsDto := units.VitaminADto{
        Value: 100,
        Unit:  units.VitaminAInternationalUnit,
    }
    
    var international_unitsResult *units.VitaminA
    international_unitsResult, err = factory.FromDto(international_unitsDto)
    if err != nil {
        t.Errorf("FromDto() with InternationalUnit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = international_unitsResult.Convert(units.VitaminAInternationalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InternationalUnit = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.VitaminADto{
        Value: 0,
        Unit:  units.VitaminAInternationalUnit,
    }
    
    var zeroResult *units.VitaminA
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestVitaminAFactory_FromDtoJSON(t *testing.T) {
    factory := units.VitaminAFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "InternationalUnit"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "InternationalUnit"}`)
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
    nanJSON, _ := json.Marshal(units.VitaminADto{
        Value: nanValue,
        Unit:  units.VitaminAInternationalUnit,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with InternationalUnit unit
    international_unitsJSON := []byte(`{"value": 100, "unit": "InternationalUnit"}`)
    international_unitsResult, err := factory.FromDtoJSON(international_unitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InternationalUnit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = international_unitsResult.Convert(units.VitaminAInternationalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InternationalUnit = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "InternationalUnit"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromInternationalUnits function
func TestVitaminAFactory_FromInternationalUnits(t *testing.T) {
    factory := units.VitaminAFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInternationalUnits(100)
    if err != nil {
        t.Errorf("FromInternationalUnits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VitaminAInternationalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInternationalUnits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInternationalUnits(math.NaN())
    if err == nil {
        t.Error("FromInternationalUnits() with NaN value should return error")
    }

    _, err = factory.FromInternationalUnits(math.Inf(1))
    if err == nil {
        t.Error("FromInternationalUnits() with +Inf value should return error")
    }

    _, err = factory.FromInternationalUnits(math.Inf(-1))
    if err == nil {
        t.Error("FromInternationalUnits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInternationalUnits(0)
    if err != nil {
        t.Errorf("FromInternationalUnits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VitaminAInternationalUnit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInternationalUnits() with zero value = %v, want 0", converted)
    }
}

func TestVitaminAToString(t *testing.T) {
	factory := units.VitaminAFactory{}
	a, err := factory.CreateVitaminA(45, units.VitaminAInternationalUnit)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.VitaminAInternationalUnit, 2)
	expected := "45.00 " + units.GetVitaminAAbbreviation(units.VitaminAInternationalUnit)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.VitaminAInternationalUnit, -1)
	expected = "45 " + units.GetVitaminAAbbreviation(units.VitaminAInternationalUnit)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestVitaminA_EqualityAndComparison(t *testing.T) {
	factory := units.VitaminAFactory{}
	a1, _ := factory.CreateVitaminA(60, units.VitaminAInternationalUnit)
	a2, _ := factory.CreateVitaminA(60, units.VitaminAInternationalUnit)
	a3, _ := factory.CreateVitaminA(120, units.VitaminAInternationalUnit)

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

func TestVitaminA_Arithmetic(t *testing.T) {
	factory := units.VitaminAFactory{}
	a1, _ := factory.CreateVitaminA(30, units.VitaminAInternationalUnit)
	a2, _ := factory.CreateVitaminA(45, units.VitaminAInternationalUnit)

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