// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestForceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Newton"}`
	
	factory := units.ForceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ForceNewton {
		t.Errorf("expected unit %v, got %v", units.ForceNewton, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Newton"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestForceDto_ToJSON(t *testing.T) {
	dto := units.ForceDto{
		Value: 45,
		Unit:  units.ForceNewton,
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
	if result["unit"].(string) != string(units.ForceNewton) {
		t.Errorf("expected unit %s, got %v", units.ForceNewton, result["unit"])
	}
}

func TestNewForce_InvalidValue(t *testing.T) {
	factory := units.ForceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateForce(math.NaN(), units.ForceNewton)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateForce(math.Inf(1), units.ForceNewton)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestForceConversions(t *testing.T) {
	factory := units.ForceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateForce(180, units.ForceNewton)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Dyne.
		// No expected conversion value provided for Dyne, verifying result is not NaN.
		result := a.Dyne()
		if math.IsNaN(result) {
			t.Errorf("conversion to Dyne returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForce.
		// No expected conversion value provided for KilogramsForce, verifying result is not NaN.
		result := a.KilogramsForce()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramsForce returned NaN")
		}
	}
	{
		// Test conversion to TonnesForce.
		// No expected conversion value provided for TonnesForce, verifying result is not NaN.
		result := a.TonnesForce()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonnesForce returned NaN")
		}
	}
	{
		// Test conversion to Newtons.
		// No expected conversion value provided for Newtons, verifying result is not NaN.
		result := a.Newtons()
		if math.IsNaN(result) {
			t.Errorf("conversion to Newtons returned NaN")
		}
	}
	{
		// Test conversion to KiloPonds.
		// No expected conversion value provided for KiloPonds, verifying result is not NaN.
		result := a.KiloPonds()
		if math.IsNaN(result) {
			t.Errorf("conversion to KiloPonds returned NaN")
		}
	}
	{
		// Test conversion to Poundals.
		// No expected conversion value provided for Poundals, verifying result is not NaN.
		result := a.Poundals()
		if math.IsNaN(result) {
			t.Errorf("conversion to Poundals returned NaN")
		}
	}
	{
		// Test conversion to PoundsForce.
		// No expected conversion value provided for PoundsForce, verifying result is not NaN.
		result := a.PoundsForce()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForce returned NaN")
		}
	}
	{
		// Test conversion to OunceForce.
		// No expected conversion value provided for OunceForce, verifying result is not NaN.
		result := a.OunceForce()
		if math.IsNaN(result) {
			t.Errorf("conversion to OunceForce returned NaN")
		}
	}
	{
		// Test conversion to ShortTonsForce.
		// No expected conversion value provided for ShortTonsForce, verifying result is not NaN.
		result := a.ShortTonsForce()
		if math.IsNaN(result) {
			t.Errorf("conversion to ShortTonsForce returned NaN")
		}
	}
	{
		// Test conversion to Micronewtons.
		// No expected conversion value provided for Micronewtons, verifying result is not NaN.
		result := a.Micronewtons()
		if math.IsNaN(result) {
			t.Errorf("conversion to Micronewtons returned NaN")
		}
	}
	{
		// Test conversion to Millinewtons.
		// No expected conversion value provided for Millinewtons, verifying result is not NaN.
		result := a.Millinewtons()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millinewtons returned NaN")
		}
	}
	{
		// Test conversion to Decanewtons.
		// No expected conversion value provided for Decanewtons, verifying result is not NaN.
		result := a.Decanewtons()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decanewtons returned NaN")
		}
	}
	{
		// Test conversion to Kilonewtons.
		// No expected conversion value provided for Kilonewtons, verifying result is not NaN.
		result := a.Kilonewtons()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilonewtons returned NaN")
		}
	}
	{
		// Test conversion to Meganewtons.
		// No expected conversion value provided for Meganewtons, verifying result is not NaN.
		result := a.Meganewtons()
		if math.IsNaN(result) {
			t.Errorf("conversion to Meganewtons returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForce.
		// No expected conversion value provided for KilopoundsForce, verifying result is not NaN.
		result := a.KilopoundsForce()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundsForce returned NaN")
		}
	}
}

func TestForce_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ForceFactory{}
	a, err := factory.CreateForce(90, units.ForceNewton)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ForceNewton {
		t.Errorf("expected default unit Newton, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ForceDyn
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ForceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ForceNewton {
		t.Errorf("expected unit Newton, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestForceFactory_FromDto(t *testing.T) {
    factory := units.ForceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceNewton,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ForceDto{
        Value: math.NaN(),
        Unit:  units.ForceNewton,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Dyn conversion
    dyneDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceDyn,
    }
    
    var dyneResult *units.Force
    dyneResult, err = factory.FromDto(dyneDto)
    if err != nil {
        t.Errorf("FromDto() with Dyn returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = dyneResult.Convert(units.ForceDyn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Dyn = %v, want %v", converted, 100)
    }
    // Test KilogramForce conversion
    kilograms_forceDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceKilogramForce,
    }
    
    var kilograms_forceResult *units.Force
    kilograms_forceResult, err = factory.FromDto(kilograms_forceDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForce returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_forceResult.Convert(units.ForceKilogramForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForce = %v, want %v", converted, 100)
    }
    // Test TonneForce conversion
    tonnes_forceDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceTonneForce,
    }
    
    var tonnes_forceResult *units.Force
    tonnes_forceResult, err = factory.FromDto(tonnes_forceDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForce returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_forceResult.Convert(units.ForceTonneForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForce = %v, want %v", converted, 100)
    }
    // Test Newton conversion
    newtonsDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceNewton,
    }
    
    var newtonsResult *units.Force
    newtonsResult, err = factory.FromDto(newtonsDto)
    if err != nil {
        t.Errorf("FromDto() with Newton returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtonsResult.Convert(units.ForceNewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Newton = %v, want %v", converted, 100)
    }
    // Test KiloPond conversion
    kilo_pondsDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceKiloPond,
    }
    
    var kilo_pondsResult *units.Force
    kilo_pondsResult, err = factory.FromDto(kilo_pondsDto)
    if err != nil {
        t.Errorf("FromDto() with KiloPond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilo_pondsResult.Convert(units.ForceKiloPond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloPond = %v, want %v", converted, 100)
    }
    // Test Poundal conversion
    poundalsDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForcePoundal,
    }
    
    var poundalsResult *units.Force
    poundalsResult, err = factory.FromDto(poundalsDto)
    if err != nil {
        t.Errorf("FromDto() with Poundal returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = poundalsResult.Convert(units.ForcePoundal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Poundal = %v, want %v", converted, 100)
    }
    // Test PoundForce conversion
    pounds_forceDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForcePoundForce,
    }
    
    var pounds_forceResult *units.Force
    pounds_forceResult, err = factory.FromDto(pounds_forceDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForce returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_forceResult.Convert(units.ForcePoundForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForce = %v, want %v", converted, 100)
    }
    // Test OunceForce conversion
    ounce_forceDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceOunceForce,
    }
    
    var ounce_forceResult *units.Force
    ounce_forceResult, err = factory.FromDto(ounce_forceDto)
    if err != nil {
        t.Errorf("FromDto() with OunceForce returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ounce_forceResult.Convert(units.ForceOunceForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OunceForce = %v, want %v", converted, 100)
    }
    // Test ShortTonForce conversion
    short_tons_forceDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceShortTonForce,
    }
    
    var short_tons_forceResult *units.Force
    short_tons_forceResult, err = factory.FromDto(short_tons_forceDto)
    if err != nil {
        t.Errorf("FromDto() with ShortTonForce returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = short_tons_forceResult.Convert(units.ForceShortTonForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ShortTonForce = %v, want %v", converted, 100)
    }
    // Test Micronewton conversion
    micronewtonsDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceMicronewton,
    }
    
    var micronewtonsResult *units.Force
    micronewtonsResult, err = factory.FromDto(micronewtonsDto)
    if err != nil {
        t.Errorf("FromDto() with Micronewton returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewtonsResult.Convert(units.ForceMicronewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Micronewton = %v, want %v", converted, 100)
    }
    // Test Millinewton conversion
    millinewtonsDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceMillinewton,
    }
    
    var millinewtonsResult *units.Force
    millinewtonsResult, err = factory.FromDto(millinewtonsDto)
    if err != nil {
        t.Errorf("FromDto() with Millinewton returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewtonsResult.Convert(units.ForceMillinewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millinewton = %v, want %v", converted, 100)
    }
    // Test Decanewton conversion
    decanewtonsDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceDecanewton,
    }
    
    var decanewtonsResult *units.Force
    decanewtonsResult, err = factory.FromDto(decanewtonsDto)
    if err != nil {
        t.Errorf("FromDto() with Decanewton returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewtonsResult.Convert(units.ForceDecanewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decanewton = %v, want %v", converted, 100)
    }
    // Test Kilonewton conversion
    kilonewtonsDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceKilonewton,
    }
    
    var kilonewtonsResult *units.Force
    kilonewtonsResult, err = factory.FromDto(kilonewtonsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilonewton returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtonsResult.Convert(units.ForceKilonewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilonewton = %v, want %v", converted, 100)
    }
    // Test Meganewton conversion
    meganewtonsDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceMeganewton,
    }
    
    var meganewtonsResult *units.Force
    meganewtonsResult, err = factory.FromDto(meganewtonsDto)
    if err != nil {
        t.Errorf("FromDto() with Meganewton returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewtonsResult.Convert(units.ForceMeganewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Meganewton = %v, want %v", converted, 100)
    }
    // Test KilopoundForce conversion
    kilopounds_forceDto := units.ForceDto{
        Value: 100,
        Unit:  units.ForceKilopoundForce,
    }
    
    var kilopounds_forceResult *units.Force
    kilopounds_forceResult, err = factory.FromDto(kilopounds_forceDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForce returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_forceResult.Convert(units.ForceKilopoundForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForce = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ForceDto{
        Value: 0,
        Unit:  units.ForceNewton,
    }
    
    var zeroResult *units.Force
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestForceFactory_FromDtoJSON(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Newton"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Newton"}`)
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
    nanJSON, _ := json.Marshal(units.ForceDto{
        Value: nanValue,
        Unit:  units.ForceNewton,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Dyn unit
    dyneJSON := []byte(`{"value": 100, "unit": "Dyn"}`)
    dyneResult, err := factory.FromDtoJSON(dyneJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Dyn unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = dyneResult.Convert(units.ForceDyn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Dyn = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForce unit
    kilograms_forceJSON := []byte(`{"value": 100, "unit": "KilogramForce"}`)
    kilograms_forceResult, err := factory.FromDtoJSON(kilograms_forceJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForce unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_forceResult.Convert(units.ForceKilogramForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForce = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForce unit
    tonnes_forceJSON := []byte(`{"value": 100, "unit": "TonneForce"}`)
    tonnes_forceResult, err := factory.FromDtoJSON(tonnes_forceJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForce unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_forceResult.Convert(units.ForceTonneForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForce = %v, want %v", converted, 100)
    }
    // Test JSON with Newton unit
    newtonsJSON := []byte(`{"value": 100, "unit": "Newton"}`)
    newtonsResult, err := factory.FromDtoJSON(newtonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Newton unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtonsResult.Convert(units.ForceNewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Newton = %v, want %v", converted, 100)
    }
    // Test JSON with KiloPond unit
    kilo_pondsJSON := []byte(`{"value": 100, "unit": "KiloPond"}`)
    kilo_pondsResult, err := factory.FromDtoJSON(kilo_pondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KiloPond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilo_pondsResult.Convert(units.ForceKiloPond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloPond = %v, want %v", converted, 100)
    }
    // Test JSON with Poundal unit
    poundalsJSON := []byte(`{"value": 100, "unit": "Poundal"}`)
    poundalsResult, err := factory.FromDtoJSON(poundalsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Poundal unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = poundalsResult.Convert(units.ForcePoundal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Poundal = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForce unit
    pounds_forceJSON := []byte(`{"value": 100, "unit": "PoundForce"}`)
    pounds_forceResult, err := factory.FromDtoJSON(pounds_forceJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForce unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_forceResult.Convert(units.ForcePoundForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForce = %v, want %v", converted, 100)
    }
    // Test JSON with OunceForce unit
    ounce_forceJSON := []byte(`{"value": 100, "unit": "OunceForce"}`)
    ounce_forceResult, err := factory.FromDtoJSON(ounce_forceJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with OunceForce unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ounce_forceResult.Convert(units.ForceOunceForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OunceForce = %v, want %v", converted, 100)
    }
    // Test JSON with ShortTonForce unit
    short_tons_forceJSON := []byte(`{"value": 100, "unit": "ShortTonForce"}`)
    short_tons_forceResult, err := factory.FromDtoJSON(short_tons_forceJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ShortTonForce unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = short_tons_forceResult.Convert(units.ForceShortTonForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ShortTonForce = %v, want %v", converted, 100)
    }
    // Test JSON with Micronewton unit
    micronewtonsJSON := []byte(`{"value": 100, "unit": "Micronewton"}`)
    micronewtonsResult, err := factory.FromDtoJSON(micronewtonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Micronewton unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micronewtonsResult.Convert(units.ForceMicronewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Micronewton = %v, want %v", converted, 100)
    }
    // Test JSON with Millinewton unit
    millinewtonsJSON := []byte(`{"value": 100, "unit": "Millinewton"}`)
    millinewtonsResult, err := factory.FromDtoJSON(millinewtonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millinewton unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millinewtonsResult.Convert(units.ForceMillinewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millinewton = %v, want %v", converted, 100)
    }
    // Test JSON with Decanewton unit
    decanewtonsJSON := []byte(`{"value": 100, "unit": "Decanewton"}`)
    decanewtonsResult, err := factory.FromDtoJSON(decanewtonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Decanewton unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decanewtonsResult.Convert(units.ForceDecanewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decanewton = %v, want %v", converted, 100)
    }
    // Test JSON with Kilonewton unit
    kilonewtonsJSON := []byte(`{"value": 100, "unit": "Kilonewton"}`)
    kilonewtonsResult, err := factory.FromDtoJSON(kilonewtonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilonewton unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtonsResult.Convert(units.ForceKilonewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilonewton = %v, want %v", converted, 100)
    }
    // Test JSON with Meganewton unit
    meganewtonsJSON := []byte(`{"value": 100, "unit": "Meganewton"}`)
    meganewtonsResult, err := factory.FromDtoJSON(meganewtonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Meganewton unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewtonsResult.Convert(units.ForceMeganewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Meganewton = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForce unit
    kilopounds_forceJSON := []byte(`{"value": 100, "unit": "KilopoundForce"}`)
    kilopounds_forceResult, err := factory.FromDtoJSON(kilopounds_forceJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForce unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_forceResult.Convert(units.ForceKilopoundForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForce = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Newton"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromDyne function
func TestForceFactory_FromDyne(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDyne(100)
    if err != nil {
        t.Errorf("FromDyne() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceDyn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDyne() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDyne(math.NaN())
    if err == nil {
        t.Error("FromDyne() with NaN value should return error")
    }

    _, err = factory.FromDyne(math.Inf(1))
    if err == nil {
        t.Error("FromDyne() with +Inf value should return error")
    }

    _, err = factory.FromDyne(math.Inf(-1))
    if err == nil {
        t.Error("FromDyne() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDyne(0)
    if err != nil {
        t.Errorf("FromDyne() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceDyn)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDyne() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsForce function
func TestForceFactory_FromKilogramsForce(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsForce(100)
    if err != nil {
        t.Errorf("FromKilogramsForce() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceKilogramForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsForce() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsForce(math.NaN())
    if err == nil {
        t.Error("FromKilogramsForce() with NaN value should return error")
    }

    _, err = factory.FromKilogramsForce(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsForce() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsForce(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsForce() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsForce(0)
    if err != nil {
        t.Errorf("FromKilogramsForce() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceKilogramForce)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsForce() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesForce function
func TestForceFactory_FromTonnesForce(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesForce(100)
    if err != nil {
        t.Errorf("FromTonnesForce() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceTonneForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesForce() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesForce(math.NaN())
    if err == nil {
        t.Error("FromTonnesForce() with NaN value should return error")
    }

    _, err = factory.FromTonnesForce(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesForce() with +Inf value should return error")
    }

    _, err = factory.FromTonnesForce(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesForce() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesForce(0)
    if err != nil {
        t.Errorf("FromTonnesForce() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceTonneForce)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesForce() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtons function
func TestForceFactory_FromNewtons(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtons(100)
    if err != nil {
        t.Errorf("FromNewtons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceNewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtons(math.NaN())
    if err == nil {
        t.Error("FromNewtons() with NaN value should return error")
    }

    _, err = factory.FromNewtons(math.Inf(1))
    if err == nil {
        t.Error("FromNewtons() with +Inf value should return error")
    }

    _, err = factory.FromNewtons(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtons(0)
    if err != nil {
        t.Errorf("FromNewtons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceNewton)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtons() with zero value = %v, want 0", converted)
    }
}
// Test FromKiloPonds function
func TestForceFactory_FromKiloPonds(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKiloPonds(100)
    if err != nil {
        t.Errorf("FromKiloPonds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceKiloPond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKiloPonds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKiloPonds(math.NaN())
    if err == nil {
        t.Error("FromKiloPonds() with NaN value should return error")
    }

    _, err = factory.FromKiloPonds(math.Inf(1))
    if err == nil {
        t.Error("FromKiloPonds() with +Inf value should return error")
    }

    _, err = factory.FromKiloPonds(math.Inf(-1))
    if err == nil {
        t.Error("FromKiloPonds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKiloPonds(0)
    if err != nil {
        t.Errorf("FromKiloPonds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceKiloPond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKiloPonds() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundals function
func TestForceFactory_FromPoundals(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundals(100)
    if err != nil {
        t.Errorf("FromPoundals() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePoundal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundals() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundals(math.NaN())
    if err == nil {
        t.Error("FromPoundals() with NaN value should return error")
    }

    _, err = factory.FromPoundals(math.Inf(1))
    if err == nil {
        t.Error("FromPoundals() with +Inf value should return error")
    }

    _, err = factory.FromPoundals(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundals() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundals(0)
    if err != nil {
        t.Errorf("FromPoundals() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePoundal)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundals() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForce function
func TestForceFactory_FromPoundsForce(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForce(100)
    if err != nil {
        t.Errorf("FromPoundsForce() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForcePoundForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForce() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForce(math.NaN())
    if err == nil {
        t.Error("FromPoundsForce() with NaN value should return error")
    }

    _, err = factory.FromPoundsForce(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForce() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForce(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForce() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForce(0)
    if err != nil {
        t.Errorf("FromPoundsForce() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForcePoundForce)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForce() with zero value = %v, want 0", converted)
    }
}
// Test FromOunceForce function
func TestForceFactory_FromOunceForce(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOunceForce(100)
    if err != nil {
        t.Errorf("FromOunceForce() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceOunceForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOunceForce() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOunceForce(math.NaN())
    if err == nil {
        t.Error("FromOunceForce() with NaN value should return error")
    }

    _, err = factory.FromOunceForce(math.Inf(1))
    if err == nil {
        t.Error("FromOunceForce() with +Inf value should return error")
    }

    _, err = factory.FromOunceForce(math.Inf(-1))
    if err == nil {
        t.Error("FromOunceForce() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOunceForce(0)
    if err != nil {
        t.Errorf("FromOunceForce() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceOunceForce)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOunceForce() with zero value = %v, want 0", converted)
    }
}
// Test FromShortTonsForce function
func TestForceFactory_FromShortTonsForce(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromShortTonsForce(100)
    if err != nil {
        t.Errorf("FromShortTonsForce() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceShortTonForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromShortTonsForce() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromShortTonsForce(math.NaN())
    if err == nil {
        t.Error("FromShortTonsForce() with NaN value should return error")
    }

    _, err = factory.FromShortTonsForce(math.Inf(1))
    if err == nil {
        t.Error("FromShortTonsForce() with +Inf value should return error")
    }

    _, err = factory.FromShortTonsForce(math.Inf(-1))
    if err == nil {
        t.Error("FromShortTonsForce() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromShortTonsForce(0)
    if err != nil {
        t.Errorf("FromShortTonsForce() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceShortTonForce)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromShortTonsForce() with zero value = %v, want 0", converted)
    }
}
// Test FromMicronewtons function
func TestForceFactory_FromMicronewtons(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicronewtons(100)
    if err != nil {
        t.Errorf("FromMicronewtons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceMicronewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicronewtons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicronewtons(math.NaN())
    if err == nil {
        t.Error("FromMicronewtons() with NaN value should return error")
    }

    _, err = factory.FromMicronewtons(math.Inf(1))
    if err == nil {
        t.Error("FromMicronewtons() with +Inf value should return error")
    }

    _, err = factory.FromMicronewtons(math.Inf(-1))
    if err == nil {
        t.Error("FromMicronewtons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicronewtons(0)
    if err != nil {
        t.Errorf("FromMicronewtons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceMicronewton)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicronewtons() with zero value = %v, want 0", converted)
    }
}
// Test FromMillinewtons function
func TestForceFactory_FromMillinewtons(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillinewtons(100)
    if err != nil {
        t.Errorf("FromMillinewtons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceMillinewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillinewtons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillinewtons(math.NaN())
    if err == nil {
        t.Error("FromMillinewtons() with NaN value should return error")
    }

    _, err = factory.FromMillinewtons(math.Inf(1))
    if err == nil {
        t.Error("FromMillinewtons() with +Inf value should return error")
    }

    _, err = factory.FromMillinewtons(math.Inf(-1))
    if err == nil {
        t.Error("FromMillinewtons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillinewtons(0)
    if err != nil {
        t.Errorf("FromMillinewtons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceMillinewton)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillinewtons() with zero value = %v, want 0", converted)
    }
}
// Test FromDecanewtons function
func TestForceFactory_FromDecanewtons(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecanewtons(100)
    if err != nil {
        t.Errorf("FromDecanewtons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceDecanewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecanewtons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecanewtons(math.NaN())
    if err == nil {
        t.Error("FromDecanewtons() with NaN value should return error")
    }

    _, err = factory.FromDecanewtons(math.Inf(1))
    if err == nil {
        t.Error("FromDecanewtons() with +Inf value should return error")
    }

    _, err = factory.FromDecanewtons(math.Inf(-1))
    if err == nil {
        t.Error("FromDecanewtons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecanewtons(0)
    if err != nil {
        t.Errorf("FromDecanewtons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceDecanewton)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecanewtons() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtons function
func TestForceFactory_FromKilonewtons(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtons(100)
    if err != nil {
        t.Errorf("FromKilonewtons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceKilonewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtons(math.NaN())
    if err == nil {
        t.Error("FromKilonewtons() with NaN value should return error")
    }

    _, err = factory.FromKilonewtons(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtons() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtons(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtons(0)
    if err != nil {
        t.Errorf("FromKilonewtons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceKilonewton)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtons() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtons function
func TestForceFactory_FromMeganewtons(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtons(100)
    if err != nil {
        t.Errorf("FromMeganewtons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceMeganewton)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtons(math.NaN())
    if err == nil {
        t.Error("FromMeganewtons() with NaN value should return error")
    }

    _, err = factory.FromMeganewtons(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtons() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtons(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtons(0)
    if err != nil {
        t.Errorf("FromMeganewtons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceMeganewton)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtons() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsForce function
func TestForceFactory_FromKilopoundsForce(t *testing.T) {
    factory := units.ForceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsForce(100)
    if err != nil {
        t.Errorf("FromKilopoundsForce() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ForceKilopoundForce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsForce() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsForce(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsForce() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsForce(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsForce() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsForce(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsForce() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsForce(0)
    if err != nil {
        t.Errorf("FromKilopoundsForce() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ForceKilopoundForce)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsForce() with zero value = %v, want 0", converted)
    }
}

func TestForceToString(t *testing.T) {
	factory := units.ForceFactory{}
	a, err := factory.CreateForce(45, units.ForceNewton)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ForceNewton, 2)
	expected := "45.00 " + units.GetForceAbbreviation(units.ForceNewton)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ForceNewton, -1)
	expected = "45 " + units.GetForceAbbreviation(units.ForceNewton)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestForce_EqualityAndComparison(t *testing.T) {
	factory := units.ForceFactory{}
	a1, _ := factory.CreateForce(60, units.ForceNewton)
	a2, _ := factory.CreateForce(60, units.ForceNewton)
	a3, _ := factory.CreateForce(120, units.ForceNewton)

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

func TestForce_Arithmetic(t *testing.T) {
	factory := units.ForceFactory{}
	a1, _ := factory.CreateForce(30, units.ForceNewton)
	a2, _ := factory.CreateForce(45, units.ForceNewton)

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