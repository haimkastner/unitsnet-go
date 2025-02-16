// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMolarEntropyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerMoleKelvin"}`
	
	factory := units.MolarEntropyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MolarEntropyJoulePerMoleKelvin {
		t.Errorf("expected unit %v, got %v", units.MolarEntropyJoulePerMoleKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerMoleKelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMolarEntropyDto_ToJSON(t *testing.T) {
	dto := units.MolarEntropyDto{
		Value: 45,
		Unit:  units.MolarEntropyJoulePerMoleKelvin,
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
	if result["unit"].(string) != string(units.MolarEntropyJoulePerMoleKelvin) {
		t.Errorf("expected unit %s, got %v", units.MolarEntropyJoulePerMoleKelvin, result["unit"])
	}
}

func TestNewMolarEntropy_InvalidValue(t *testing.T) {
	factory := units.MolarEntropyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMolarEntropy(math.NaN(), units.MolarEntropyJoulePerMoleKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMolarEntropy(math.Inf(1), units.MolarEntropyJoulePerMoleKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMolarEntropyConversions(t *testing.T) {
	factory := units.MolarEntropyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMolarEntropy(180, units.MolarEntropyJoulePerMoleKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerMoleKelvin.
		// No expected conversion value provided for JoulesPerMoleKelvin, verifying result is not NaN.
		result := a.JoulesPerMoleKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerMoleKelvin returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerMoleKelvin.
		// No expected conversion value provided for KilojoulesPerMoleKelvin, verifying result is not NaN.
		result := a.KilojoulesPerMoleKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerMoleKelvin returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerMoleKelvin.
		// No expected conversion value provided for MegajoulesPerMoleKelvin, verifying result is not NaN.
		result := a.MegajoulesPerMoleKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegajoulesPerMoleKelvin returned NaN")
		}
	}
}

func TestMolarEntropy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MolarEntropyFactory{}
	a, err := factory.CreateMolarEntropy(90, units.MolarEntropyJoulePerMoleKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MolarEntropyJoulePerMoleKelvin {
		t.Errorf("expected default unit JoulePerMoleKelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MolarEntropyJoulePerMoleKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MolarEntropyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MolarEntropyJoulePerMoleKelvin {
		t.Errorf("expected unit JoulePerMoleKelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMolarEntropyFactory_FromDto(t *testing.T) {
    factory := units.MolarEntropyFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MolarEntropyDto{
        Value: 100,
        Unit:  units.MolarEntropyJoulePerMoleKelvin,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MolarEntropyDto{
        Value: math.NaN(),
        Unit:  units.MolarEntropyJoulePerMoleKelvin,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test JoulePerMoleKelvin conversion
    joules_per_mole_kelvinDto := units.MolarEntropyDto{
        Value: 100,
        Unit:  units.MolarEntropyJoulePerMoleKelvin,
    }
    
    var joules_per_mole_kelvinResult *units.MolarEntropy
    joules_per_mole_kelvinResult, err = factory.FromDto(joules_per_mole_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerMoleKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_mole_kelvinResult.Convert(units.MolarEntropyJoulePerMoleKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerMoleKelvin = %v, want %v", converted, 100)
    }
    // Test KilojoulePerMoleKelvin conversion
    kilojoules_per_mole_kelvinDto := units.MolarEntropyDto{
        Value: 100,
        Unit:  units.MolarEntropyKilojoulePerMoleKelvin,
    }
    
    var kilojoules_per_mole_kelvinResult *units.MolarEntropy
    kilojoules_per_mole_kelvinResult, err = factory.FromDto(kilojoules_per_mole_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with KilojoulePerMoleKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_mole_kelvinResult.Convert(units.MolarEntropyKilojoulePerMoleKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerMoleKelvin = %v, want %v", converted, 100)
    }
    // Test MegajoulePerMoleKelvin conversion
    megajoules_per_mole_kelvinDto := units.MolarEntropyDto{
        Value: 100,
        Unit:  units.MolarEntropyMegajoulePerMoleKelvin,
    }
    
    var megajoules_per_mole_kelvinResult *units.MolarEntropy
    megajoules_per_mole_kelvinResult, err = factory.FromDto(megajoules_per_mole_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with MegajoulePerMoleKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_mole_kelvinResult.Convert(units.MolarEntropyMegajoulePerMoleKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerMoleKelvin = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MolarEntropyDto{
        Value: 0,
        Unit:  units.MolarEntropyJoulePerMoleKelvin,
    }
    
    var zeroResult *units.MolarEntropy
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMolarEntropyFactory_FromDtoJSON(t *testing.T) {
    factory := units.MolarEntropyFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "JoulePerMoleKelvin"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "JoulePerMoleKelvin"}`)
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
    nanJSON, _ := json.Marshal(units.MolarEntropyDto{
        Value: nanValue,
        Unit:  units.MolarEntropyJoulePerMoleKelvin,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with JoulePerMoleKelvin unit
    joules_per_mole_kelvinJSON := []byte(`{"value": 100, "unit": "JoulePerMoleKelvin"}`)
    joules_per_mole_kelvinResult, err := factory.FromDtoJSON(joules_per_mole_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerMoleKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_mole_kelvinResult.Convert(units.MolarEntropyJoulePerMoleKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerMoleKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with KilojoulePerMoleKelvin unit
    kilojoules_per_mole_kelvinJSON := []byte(`{"value": 100, "unit": "KilojoulePerMoleKelvin"}`)
    kilojoules_per_mole_kelvinResult, err := factory.FromDtoJSON(kilojoules_per_mole_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilojoulePerMoleKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_mole_kelvinResult.Convert(units.MolarEntropyKilojoulePerMoleKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerMoleKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with MegajoulePerMoleKelvin unit
    megajoules_per_mole_kelvinJSON := []byte(`{"value": 100, "unit": "MegajoulePerMoleKelvin"}`)
    megajoules_per_mole_kelvinResult, err := factory.FromDtoJSON(megajoules_per_mole_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegajoulePerMoleKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_mole_kelvinResult.Convert(units.MolarEntropyMegajoulePerMoleKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerMoleKelvin = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "JoulePerMoleKelvin"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromJoulesPerMoleKelvin function
func TestMolarEntropyFactory_FromJoulesPerMoleKelvin(t *testing.T) {
    factory := units.MolarEntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerMoleKelvin(100)
    if err != nil {
        t.Errorf("FromJoulesPerMoleKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarEntropyJoulePerMoleKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerMoleKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerMoleKelvin(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerMoleKelvin() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerMoleKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerMoleKelvin() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerMoleKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerMoleKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerMoleKelvin(0)
    if err != nil {
        t.Errorf("FromJoulesPerMoleKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarEntropyJoulePerMoleKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerMoleKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromKilojoulesPerMoleKelvin function
func TestMolarEntropyFactory_FromKilojoulesPerMoleKelvin(t *testing.T) {
    factory := units.MolarEntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilojoulesPerMoleKelvin(100)
    if err != nil {
        t.Errorf("FromKilojoulesPerMoleKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarEntropyKilojoulePerMoleKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilojoulesPerMoleKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilojoulesPerMoleKelvin(math.NaN())
    if err == nil {
        t.Error("FromKilojoulesPerMoleKelvin() with NaN value should return error")
    }

    _, err = factory.FromKilojoulesPerMoleKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromKilojoulesPerMoleKelvin() with +Inf value should return error")
    }

    _, err = factory.FromKilojoulesPerMoleKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromKilojoulesPerMoleKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilojoulesPerMoleKelvin(0)
    if err != nil {
        t.Errorf("FromKilojoulesPerMoleKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarEntropyKilojoulePerMoleKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilojoulesPerMoleKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromMegajoulesPerMoleKelvin function
func TestMolarEntropyFactory_FromMegajoulesPerMoleKelvin(t *testing.T) {
    factory := units.MolarEntropyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegajoulesPerMoleKelvin(100)
    if err != nil {
        t.Errorf("FromMegajoulesPerMoleKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarEntropyMegajoulePerMoleKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegajoulesPerMoleKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegajoulesPerMoleKelvin(math.NaN())
    if err == nil {
        t.Error("FromMegajoulesPerMoleKelvin() with NaN value should return error")
    }

    _, err = factory.FromMegajoulesPerMoleKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromMegajoulesPerMoleKelvin() with +Inf value should return error")
    }

    _, err = factory.FromMegajoulesPerMoleKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromMegajoulesPerMoleKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegajoulesPerMoleKelvin(0)
    if err != nil {
        t.Errorf("FromMegajoulesPerMoleKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarEntropyMegajoulePerMoleKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegajoulesPerMoleKelvin() with zero value = %v, want 0", converted)
    }
}

func TestMolarEntropyToString(t *testing.T) {
	factory := units.MolarEntropyFactory{}
	a, err := factory.CreateMolarEntropy(45, units.MolarEntropyJoulePerMoleKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MolarEntropyJoulePerMoleKelvin, 2)
	expected := "45.00 " + units.GetMolarEntropyAbbreviation(units.MolarEntropyJoulePerMoleKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MolarEntropyJoulePerMoleKelvin, -1)
	expected = "45 " + units.GetMolarEntropyAbbreviation(units.MolarEntropyJoulePerMoleKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMolarEntropy_EqualityAndComparison(t *testing.T) {
	factory := units.MolarEntropyFactory{}
	a1, _ := factory.CreateMolarEntropy(60, units.MolarEntropyJoulePerMoleKelvin)
	a2, _ := factory.CreateMolarEntropy(60, units.MolarEntropyJoulePerMoleKelvin)
	a3, _ := factory.CreateMolarEntropy(120, units.MolarEntropyJoulePerMoleKelvin)

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

func TestMolarEntropy_Arithmetic(t *testing.T) {
	factory := units.MolarEntropyFactory{}
	a1, _ := factory.CreateMolarEntropy(30, units.MolarEntropyJoulePerMoleKelvin)
	a2, _ := factory.CreateMolarEntropy(45, units.MolarEntropyJoulePerMoleKelvin)

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