// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestSpecificFuelConsumptionDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "GramPerKilonewtonSecond"}`
	
	factory := units.SpecificFuelConsumptionDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.SpecificFuelConsumptionGramPerKilonewtonSecond {
		t.Errorf("expected unit %v, got %v", units.SpecificFuelConsumptionGramPerKilonewtonSecond, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "GramPerKilonewtonSecond"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestSpecificFuelConsumptionDto_ToJSON(t *testing.T) {
	dto := units.SpecificFuelConsumptionDto{
		Value: 45,
		Unit:  units.SpecificFuelConsumptionGramPerKilonewtonSecond,
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
	if result["unit"].(string) != string(units.SpecificFuelConsumptionGramPerKilonewtonSecond) {
		t.Errorf("expected unit %s, got %v", units.SpecificFuelConsumptionGramPerKilonewtonSecond, result["unit"])
	}
}

func TestNewSpecificFuelConsumption_InvalidValue(t *testing.T) {
	factory := units.SpecificFuelConsumptionFactory{}
	// NaN value should return an error.
	_, err := factory.CreateSpecificFuelConsumption(math.NaN(), units.SpecificFuelConsumptionGramPerKilonewtonSecond)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateSpecificFuelConsumption(math.Inf(1), units.SpecificFuelConsumptionGramPerKilonewtonSecond)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestSpecificFuelConsumptionConversions(t *testing.T) {
	factory := units.SpecificFuelConsumptionFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateSpecificFuelConsumption(180, units.SpecificFuelConsumptionGramPerKilonewtonSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to PoundsMassPerPoundForceHour.
		// No expected conversion value provided for PoundsMassPerPoundForceHour, verifying result is not NaN.
		result := a.PoundsMassPerPoundForceHour()
		cacheResult := a.PoundsMassPerPoundForceHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsMassPerPoundForceHour returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerKilogramForceHour.
		// No expected conversion value provided for KilogramsPerKilogramForceHour, verifying result is not NaN.
		result := a.KilogramsPerKilogramForceHour()
		cacheResult := a.KilogramsPerKilogramForceHour()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsPerKilogramForceHour returned NaN")
		}
	}
	{
		// Test conversion to GramsPerKilonewtonSecond.
		// No expected conversion value provided for GramsPerKilonewtonSecond, verifying result is not NaN.
		result := a.GramsPerKilonewtonSecond()
		cacheResult := a.GramsPerKilonewtonSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GramsPerKilonewtonSecond returned NaN")
		}
	}
	{
		// Test conversion to KilogramsPerKilonewtonSecond.
		// No expected conversion value provided for KilogramsPerKilonewtonSecond, verifying result is not NaN.
		result := a.KilogramsPerKilonewtonSecond()
		cacheResult := a.KilogramsPerKilonewtonSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsPerKilonewtonSecond returned NaN")
		}
	}
}

func TestSpecificFuelConsumption_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.SpecificFuelConsumptionFactory{}
	a, err := factory.CreateSpecificFuelConsumption(90, units.SpecificFuelConsumptionGramPerKilonewtonSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.SpecificFuelConsumptionGramPerKilonewtonSecond {
		t.Errorf("expected default unit GramPerKilonewtonSecond, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.SpecificFuelConsumptionPoundMassPerPoundForceHour
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.SpecificFuelConsumptionDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.SpecificFuelConsumptionGramPerKilonewtonSecond {
		t.Errorf("expected unit GramPerKilonewtonSecond, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestSpecificFuelConsumptionFactory_FromDto(t *testing.T) {
    factory := units.SpecificFuelConsumptionFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.SpecificFuelConsumptionDto{
        Value: 100,
        Unit:  units.SpecificFuelConsumptionGramPerKilonewtonSecond,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.SpecificFuelConsumptionDto{
        Value: math.NaN(),
        Unit:  units.SpecificFuelConsumptionGramPerKilonewtonSecond,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test PoundMassPerPoundForceHour conversion
    pounds_mass_per_pound_force_hourDto := units.SpecificFuelConsumptionDto{
        Value: 100,
        Unit:  units.SpecificFuelConsumptionPoundMassPerPoundForceHour,
    }
    
    var pounds_mass_per_pound_force_hourResult *units.SpecificFuelConsumption
    pounds_mass_per_pound_force_hourResult, err = factory.FromDto(pounds_mass_per_pound_force_hourDto)
    if err != nil {
        t.Errorf("FromDto() with PoundMassPerPoundForceHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_mass_per_pound_force_hourResult.Convert(units.SpecificFuelConsumptionPoundMassPerPoundForceHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundMassPerPoundForceHour = %v, want %v", converted, 100)
    }
    // Test KilogramPerKilogramForceHour conversion
    kilograms_per_kilogram_force_hourDto := units.SpecificFuelConsumptionDto{
        Value: 100,
        Unit:  units.SpecificFuelConsumptionKilogramPerKilogramForceHour,
    }
    
    var kilograms_per_kilogram_force_hourResult *units.SpecificFuelConsumption
    kilograms_per_kilogram_force_hourResult, err = factory.FromDto(kilograms_per_kilogram_force_hourDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerKilogramForceHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_kilogram_force_hourResult.Convert(units.SpecificFuelConsumptionKilogramPerKilogramForceHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerKilogramForceHour = %v, want %v", converted, 100)
    }
    // Test GramPerKilonewtonSecond conversion
    grams_per_kilonewton_secondDto := units.SpecificFuelConsumptionDto{
        Value: 100,
        Unit:  units.SpecificFuelConsumptionGramPerKilonewtonSecond,
    }
    
    var grams_per_kilonewton_secondResult *units.SpecificFuelConsumption
    grams_per_kilonewton_secondResult, err = factory.FromDto(grams_per_kilonewton_secondDto)
    if err != nil {
        t.Errorf("FromDto() with GramPerKilonewtonSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_kilonewton_secondResult.Convert(units.SpecificFuelConsumptionGramPerKilonewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerKilonewtonSecond = %v, want %v", converted, 100)
    }
    // Test KilogramPerKilonewtonSecond conversion
    kilograms_per_kilonewton_secondDto := units.SpecificFuelConsumptionDto{
        Value: 100,
        Unit:  units.SpecificFuelConsumptionKilogramPerKilonewtonSecond,
    }
    
    var kilograms_per_kilonewton_secondResult *units.SpecificFuelConsumption
    kilograms_per_kilonewton_secondResult, err = factory.FromDto(kilograms_per_kilonewton_secondDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramPerKilonewtonSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_kilonewton_secondResult.Convert(units.SpecificFuelConsumptionKilogramPerKilonewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerKilonewtonSecond = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.SpecificFuelConsumptionDto{
        Value: 0,
        Unit:  units.SpecificFuelConsumptionGramPerKilonewtonSecond,
    }
    
    var zeroResult *units.SpecificFuelConsumption
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestSpecificFuelConsumptionFactory_FromDtoJSON(t *testing.T) {
    factory := units.SpecificFuelConsumptionFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "GramPerKilonewtonSecond"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "GramPerKilonewtonSecond"}`)
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
    nanJSON, _ := json.Marshal(units.SpecificFuelConsumptionDto{
        Value: nanValue,
        Unit:  units.SpecificFuelConsumptionGramPerKilonewtonSecond,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with PoundMassPerPoundForceHour unit
    pounds_mass_per_pound_force_hourJSON := []byte(`{"value": 100, "unit": "PoundMassPerPoundForceHour"}`)
    pounds_mass_per_pound_force_hourResult, err := factory.FromDtoJSON(pounds_mass_per_pound_force_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundMassPerPoundForceHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_mass_per_pound_force_hourResult.Convert(units.SpecificFuelConsumptionPoundMassPerPoundForceHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundMassPerPoundForceHour = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerKilogramForceHour unit
    kilograms_per_kilogram_force_hourJSON := []byte(`{"value": 100, "unit": "KilogramPerKilogramForceHour"}`)
    kilograms_per_kilogram_force_hourResult, err := factory.FromDtoJSON(kilograms_per_kilogram_force_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerKilogramForceHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_kilogram_force_hourResult.Convert(units.SpecificFuelConsumptionKilogramPerKilogramForceHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerKilogramForceHour = %v, want %v", converted, 100)
    }
    // Test JSON with GramPerKilonewtonSecond unit
    grams_per_kilonewton_secondJSON := []byte(`{"value": 100, "unit": "GramPerKilonewtonSecond"}`)
    grams_per_kilonewton_secondResult, err := factory.FromDtoJSON(grams_per_kilonewton_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramPerKilonewtonSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = grams_per_kilonewton_secondResult.Convert(units.SpecificFuelConsumptionGramPerKilonewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramPerKilonewtonSecond = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramPerKilonewtonSecond unit
    kilograms_per_kilonewton_secondJSON := []byte(`{"value": 100, "unit": "KilogramPerKilonewtonSecond"}`)
    kilograms_per_kilonewton_secondResult, err := factory.FromDtoJSON(kilograms_per_kilonewton_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramPerKilonewtonSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_per_kilonewton_secondResult.Convert(units.SpecificFuelConsumptionKilogramPerKilonewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramPerKilonewtonSecond = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "GramPerKilonewtonSecond"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromPoundsMassPerPoundForceHour function
func TestSpecificFuelConsumptionFactory_FromPoundsMassPerPoundForceHour(t *testing.T) {
    factory := units.SpecificFuelConsumptionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsMassPerPoundForceHour(100)
    if err != nil {
        t.Errorf("FromPoundsMassPerPoundForceHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificFuelConsumptionPoundMassPerPoundForceHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsMassPerPoundForceHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsMassPerPoundForceHour(math.NaN())
    if err == nil {
        t.Error("FromPoundsMassPerPoundForceHour() with NaN value should return error")
    }

    _, err = factory.FromPoundsMassPerPoundForceHour(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsMassPerPoundForceHour() with +Inf value should return error")
    }

    _, err = factory.FromPoundsMassPerPoundForceHour(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsMassPerPoundForceHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsMassPerPoundForceHour(0)
    if err != nil {
        t.Errorf("FromPoundsMassPerPoundForceHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificFuelConsumptionPoundMassPerPoundForceHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsMassPerPoundForceHour() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerKilogramForceHour function
func TestSpecificFuelConsumptionFactory_FromKilogramsPerKilogramForceHour(t *testing.T) {
    factory := units.SpecificFuelConsumptionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerKilogramForceHour(100)
    if err != nil {
        t.Errorf("FromKilogramsPerKilogramForceHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificFuelConsumptionKilogramPerKilogramForceHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerKilogramForceHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerKilogramForceHour(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerKilogramForceHour() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerKilogramForceHour(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerKilogramForceHour() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerKilogramForceHour(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerKilogramForceHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerKilogramForceHour(0)
    if err != nil {
        t.Errorf("FromKilogramsPerKilogramForceHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificFuelConsumptionKilogramPerKilogramForceHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerKilogramForceHour() with zero value = %v, want 0", converted)
    }
}
// Test FromGramsPerKilonewtonSecond function
func TestSpecificFuelConsumptionFactory_FromGramsPerKilonewtonSecond(t *testing.T) {
    factory := units.SpecificFuelConsumptionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramsPerKilonewtonSecond(100)
    if err != nil {
        t.Errorf("FromGramsPerKilonewtonSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificFuelConsumptionGramPerKilonewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramsPerKilonewtonSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramsPerKilonewtonSecond(math.NaN())
    if err == nil {
        t.Error("FromGramsPerKilonewtonSecond() with NaN value should return error")
    }

    _, err = factory.FromGramsPerKilonewtonSecond(math.Inf(1))
    if err == nil {
        t.Error("FromGramsPerKilonewtonSecond() with +Inf value should return error")
    }

    _, err = factory.FromGramsPerKilonewtonSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromGramsPerKilonewtonSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramsPerKilonewtonSecond(0)
    if err != nil {
        t.Errorf("FromGramsPerKilonewtonSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificFuelConsumptionGramPerKilonewtonSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramsPerKilonewtonSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsPerKilonewtonSecond function
func TestSpecificFuelConsumptionFactory_FromKilogramsPerKilonewtonSecond(t *testing.T) {
    factory := units.SpecificFuelConsumptionFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsPerKilonewtonSecond(100)
    if err != nil {
        t.Errorf("FromKilogramsPerKilonewtonSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificFuelConsumptionKilogramPerKilonewtonSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsPerKilonewtonSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsPerKilonewtonSecond(math.NaN())
    if err == nil {
        t.Error("FromKilogramsPerKilonewtonSecond() with NaN value should return error")
    }

    _, err = factory.FromKilogramsPerKilonewtonSecond(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsPerKilonewtonSecond() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsPerKilonewtonSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsPerKilonewtonSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsPerKilonewtonSecond(0)
    if err != nil {
        t.Errorf("FromKilogramsPerKilonewtonSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificFuelConsumptionKilogramPerKilonewtonSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsPerKilonewtonSecond() with zero value = %v, want 0", converted)
    }
}

func TestSpecificFuelConsumptionToString(t *testing.T) {
	factory := units.SpecificFuelConsumptionFactory{}
	a, err := factory.CreateSpecificFuelConsumption(45, units.SpecificFuelConsumptionGramPerKilonewtonSecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.SpecificFuelConsumptionGramPerKilonewtonSecond, 2)
	expected := "45.00 " + units.GetSpecificFuelConsumptionAbbreviation(units.SpecificFuelConsumptionGramPerKilonewtonSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.SpecificFuelConsumptionGramPerKilonewtonSecond, -1)
	expected = "45 " + units.GetSpecificFuelConsumptionAbbreviation(units.SpecificFuelConsumptionGramPerKilonewtonSecond)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestSpecificFuelConsumption_EqualityAndComparison(t *testing.T) {
	factory := units.SpecificFuelConsumptionFactory{}
	a1, _ := factory.CreateSpecificFuelConsumption(60, units.SpecificFuelConsumptionGramPerKilonewtonSecond)
	a2, _ := factory.CreateSpecificFuelConsumption(60, units.SpecificFuelConsumptionGramPerKilonewtonSecond)
	a3, _ := factory.CreateSpecificFuelConsumption(120, units.SpecificFuelConsumptionGramPerKilonewtonSecond)

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

func TestSpecificFuelConsumption_Arithmetic(t *testing.T) {
	factory := units.SpecificFuelConsumptionFactory{}
	a1, _ := factory.CreateSpecificFuelConsumption(30, units.SpecificFuelConsumptionGramPerKilonewtonSecond)
	a2, _ := factory.CreateSpecificFuelConsumption(45, units.SpecificFuelConsumptionGramPerKilonewtonSecond)

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


func TestGetSpecificFuelConsumptionAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.SpecificFuelConsumptionUnits
        want string
    }{
        {
            name: "PoundMassPerPoundForceHour abbreviation",
            unit: units.SpecificFuelConsumptionPoundMassPerPoundForceHour,
            want: "lb/(lbf·h)",
        },
        {
            name: "KilogramPerKilogramForceHour abbreviation",
            unit: units.SpecificFuelConsumptionKilogramPerKilogramForceHour,
            want: "kg/(kgf·h)",
        },
        {
            name: "GramPerKilonewtonSecond abbreviation",
            unit: units.SpecificFuelConsumptionGramPerKilonewtonSecond,
            want: "g/(kN·s)",
        },
        {
            name: "KilogramPerKilonewtonSecond abbreviation",
            unit: units.SpecificFuelConsumptionKilogramPerKilonewtonSecond,
            want: "kg/(kN·s)",
        },
        {
            name: "invalid unit",
            unit: units.SpecificFuelConsumptionUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetSpecificFuelConsumptionAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetSpecificFuelConsumptionAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestSpecificFuelConsumption_String(t *testing.T) {
    factory := units.SpecificFuelConsumptionFactory{}
    
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
            unit, err := factory.CreateSpecificFuelConsumption(tt.value, units.SpecificFuelConsumptionGramPerKilonewtonSecond)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("SpecificFuelConsumption.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestSpecificFuelConsumption_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.SpecificFuelConsumptionFactory{}

	_, err := uf.CreateSpecificFuelConsumption(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}