// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestMolarEnergyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerMole"}`
	
	factory := units.MolarEnergyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.MolarEnergyJoulePerMole {
		t.Errorf("expected unit %v, got %v", units.MolarEnergyJoulePerMole, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerMole"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestMolarEnergyDto_ToJSON(t *testing.T) {
	dto := units.MolarEnergyDto{
		Value: 45,
		Unit:  units.MolarEnergyJoulePerMole,
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
	if result["unit"].(string) != string(units.MolarEnergyJoulePerMole) {
		t.Errorf("expected unit %s, got %v", units.MolarEnergyJoulePerMole, result["unit"])
	}
}

func TestNewMolarEnergy_InvalidValue(t *testing.T) {
	factory := units.MolarEnergyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateMolarEnergy(math.NaN(), units.MolarEnergyJoulePerMole)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateMolarEnergy(math.Inf(1), units.MolarEnergyJoulePerMole)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestMolarEnergyConversions(t *testing.T) {
	factory := units.MolarEnergyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateMolarEnergy(180, units.MolarEnergyJoulePerMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerMole.
		// No expected conversion value provided for JoulesPerMole, verifying result is not NaN.
		result := a.JoulesPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerMole returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerMole.
		// No expected conversion value provided for KilojoulesPerMole, verifying result is not NaN.
		result := a.KilojoulesPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerMole returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerMole.
		// No expected conversion value provided for MegajoulesPerMole, verifying result is not NaN.
		result := a.MegajoulesPerMole()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegajoulesPerMole returned NaN")
		}
	}
}

func TestMolarEnergy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.MolarEnergyFactory{}
	a, err := factory.CreateMolarEnergy(90, units.MolarEnergyJoulePerMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.MolarEnergyJoulePerMole {
		t.Errorf("expected default unit JoulePerMole, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.MolarEnergyJoulePerMole
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.MolarEnergyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.MolarEnergyJoulePerMole {
		t.Errorf("expected unit JoulePerMole, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestMolarEnergyFactory_FromDto(t *testing.T) {
    factory := units.MolarEnergyFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.MolarEnergyDto{
        Value: 100,
        Unit:  units.MolarEnergyJoulePerMole,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.MolarEnergyDto{
        Value: math.NaN(),
        Unit:  units.MolarEnergyJoulePerMole,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test JoulePerMole conversion
    joules_per_moleDto := units.MolarEnergyDto{
        Value: 100,
        Unit:  units.MolarEnergyJoulePerMole,
    }
    
    var joules_per_moleResult *units.MolarEnergy
    joules_per_moleResult, err = factory.FromDto(joules_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_moleResult.Convert(units.MolarEnergyJoulePerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerMole = %v, want %v", converted, 100)
    }
    // Test KilojoulePerMole conversion
    kilojoules_per_moleDto := units.MolarEnergyDto{
        Value: 100,
        Unit:  units.MolarEnergyKilojoulePerMole,
    }
    
    var kilojoules_per_moleResult *units.MolarEnergy
    kilojoules_per_moleResult, err = factory.FromDto(kilojoules_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with KilojoulePerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_moleResult.Convert(units.MolarEnergyKilojoulePerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerMole = %v, want %v", converted, 100)
    }
    // Test MegajoulePerMole conversion
    megajoules_per_moleDto := units.MolarEnergyDto{
        Value: 100,
        Unit:  units.MolarEnergyMegajoulePerMole,
    }
    
    var megajoules_per_moleResult *units.MolarEnergy
    megajoules_per_moleResult, err = factory.FromDto(megajoules_per_moleDto)
    if err != nil {
        t.Errorf("FromDto() with MegajoulePerMole returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_moleResult.Convert(units.MolarEnergyMegajoulePerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerMole = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.MolarEnergyDto{
        Value: 0,
        Unit:  units.MolarEnergyJoulePerMole,
    }
    
    var zeroResult *units.MolarEnergy
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestMolarEnergyFactory_FromDtoJSON(t *testing.T) {
    factory := units.MolarEnergyFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "JoulePerMole"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "JoulePerMole"}`)
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
    nanJSON, _ := json.Marshal(units.MolarEnergyDto{
        Value: nanValue,
        Unit:  units.MolarEnergyJoulePerMole,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with JoulePerMole unit
    joules_per_moleJSON := []byte(`{"value": 100, "unit": "JoulePerMole"}`)
    joules_per_moleResult, err := factory.FromDtoJSON(joules_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_moleResult.Convert(units.MolarEnergyJoulePerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerMole = %v, want %v", converted, 100)
    }
    // Test JSON with KilojoulePerMole unit
    kilojoules_per_moleJSON := []byte(`{"value": 100, "unit": "KilojoulePerMole"}`)
    kilojoules_per_moleResult, err := factory.FromDtoJSON(kilojoules_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilojoulePerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_moleResult.Convert(units.MolarEnergyKilojoulePerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerMole = %v, want %v", converted, 100)
    }
    // Test JSON with MegajoulePerMole unit
    megajoules_per_moleJSON := []byte(`{"value": 100, "unit": "MegajoulePerMole"}`)
    megajoules_per_moleResult, err := factory.FromDtoJSON(megajoules_per_moleJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegajoulePerMole unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_moleResult.Convert(units.MolarEnergyMegajoulePerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerMole = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "JoulePerMole"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromJoulesPerMole function
func TestMolarEnergyFactory_FromJoulesPerMole(t *testing.T) {
    factory := units.MolarEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerMole(100)
    if err != nil {
        t.Errorf("FromJoulesPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarEnergyJoulePerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerMole(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerMole() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerMole() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerMole(0)
    if err != nil {
        t.Errorf("FromJoulesPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarEnergyJoulePerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerMole() with zero value = %v, want 0", converted)
    }
}
// Test FromKilojoulesPerMole function
func TestMolarEnergyFactory_FromKilojoulesPerMole(t *testing.T) {
    factory := units.MolarEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilojoulesPerMole(100)
    if err != nil {
        t.Errorf("FromKilojoulesPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarEnergyKilojoulePerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilojoulesPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilojoulesPerMole(math.NaN())
    if err == nil {
        t.Error("FromKilojoulesPerMole() with NaN value should return error")
    }

    _, err = factory.FromKilojoulesPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromKilojoulesPerMole() with +Inf value should return error")
    }

    _, err = factory.FromKilojoulesPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromKilojoulesPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilojoulesPerMole(0)
    if err != nil {
        t.Errorf("FromKilojoulesPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarEnergyKilojoulePerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilojoulesPerMole() with zero value = %v, want 0", converted)
    }
}
// Test FromMegajoulesPerMole function
func TestMolarEnergyFactory_FromMegajoulesPerMole(t *testing.T) {
    factory := units.MolarEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegajoulesPerMole(100)
    if err != nil {
        t.Errorf("FromMegajoulesPerMole() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.MolarEnergyMegajoulePerMole)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegajoulesPerMole() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegajoulesPerMole(math.NaN())
    if err == nil {
        t.Error("FromMegajoulesPerMole() with NaN value should return error")
    }

    _, err = factory.FromMegajoulesPerMole(math.Inf(1))
    if err == nil {
        t.Error("FromMegajoulesPerMole() with +Inf value should return error")
    }

    _, err = factory.FromMegajoulesPerMole(math.Inf(-1))
    if err == nil {
        t.Error("FromMegajoulesPerMole() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegajoulesPerMole(0)
    if err != nil {
        t.Errorf("FromMegajoulesPerMole() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.MolarEnergyMegajoulePerMole)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegajoulesPerMole() with zero value = %v, want 0", converted)
    }
}

func TestMolarEnergyToString(t *testing.T) {
	factory := units.MolarEnergyFactory{}
	a, err := factory.CreateMolarEnergy(45, units.MolarEnergyJoulePerMole)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.MolarEnergyJoulePerMole, 2)
	expected := "45.00 " + units.GetMolarEnergyAbbreviation(units.MolarEnergyJoulePerMole)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.MolarEnergyJoulePerMole, -1)
	expected = "45 " + units.GetMolarEnergyAbbreviation(units.MolarEnergyJoulePerMole)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestMolarEnergy_EqualityAndComparison(t *testing.T) {
	factory := units.MolarEnergyFactory{}
	a1, _ := factory.CreateMolarEnergy(60, units.MolarEnergyJoulePerMole)
	a2, _ := factory.CreateMolarEnergy(60, units.MolarEnergyJoulePerMole)
	a3, _ := factory.CreateMolarEnergy(120, units.MolarEnergyJoulePerMole)

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

func TestMolarEnergy_Arithmetic(t *testing.T) {
	factory := units.MolarEnergyFactory{}
	a1, _ := factory.CreateMolarEnergy(30, units.MolarEnergyJoulePerMole)
	a2, _ := factory.CreateMolarEnergy(45, units.MolarEnergyJoulePerMole)

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