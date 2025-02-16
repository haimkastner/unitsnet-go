// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricReactivePowerDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltampereReactive"}`
	
	factory := units.ElectricReactivePowerDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricReactivePowerVoltampereReactive {
		t.Errorf("expected unit %v, got %v", units.ElectricReactivePowerVoltampereReactive, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltampereReactive"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricReactivePowerDto_ToJSON(t *testing.T) {
	dto := units.ElectricReactivePowerDto{
		Value: 45,
		Unit:  units.ElectricReactivePowerVoltampereReactive,
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
	if result["unit"].(string) != string(units.ElectricReactivePowerVoltampereReactive) {
		t.Errorf("expected unit %s, got %v", units.ElectricReactivePowerVoltampereReactive, result["unit"])
	}
}

func TestNewElectricReactivePower_InvalidValue(t *testing.T) {
	factory := units.ElectricReactivePowerFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricReactivePower(math.NaN(), units.ElectricReactivePowerVoltampereReactive)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricReactivePower(math.Inf(1), units.ElectricReactivePowerVoltampereReactive)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricReactivePowerConversions(t *testing.T) {
	factory := units.ElectricReactivePowerFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricReactivePower(180, units.ElectricReactivePowerVoltampereReactive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltamperesReactive.
		// No expected conversion value provided for VoltamperesReactive, verifying result is not NaN.
		result := a.VoltamperesReactive()
		if math.IsNaN(result) {
			t.Errorf("conversion to VoltamperesReactive returned NaN")
		}
	}
	{
		// Test conversion to KilovoltamperesReactive.
		// No expected conversion value provided for KilovoltamperesReactive, verifying result is not NaN.
		result := a.KilovoltamperesReactive()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilovoltamperesReactive returned NaN")
		}
	}
	{
		// Test conversion to MegavoltamperesReactive.
		// No expected conversion value provided for MegavoltamperesReactive, verifying result is not NaN.
		result := a.MegavoltamperesReactive()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegavoltamperesReactive returned NaN")
		}
	}
	{
		// Test conversion to GigavoltamperesReactive.
		// No expected conversion value provided for GigavoltamperesReactive, verifying result is not NaN.
		result := a.GigavoltamperesReactive()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigavoltamperesReactive returned NaN")
		}
	}
}

func TestElectricReactivePower_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricReactivePowerFactory{}
	a, err := factory.CreateElectricReactivePower(90, units.ElectricReactivePowerVoltampereReactive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricReactivePowerVoltampereReactive {
		t.Errorf("expected default unit VoltampereReactive, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricReactivePowerVoltampereReactive
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricReactivePowerDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricReactivePowerVoltampereReactive {
		t.Errorf("expected unit VoltampereReactive, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricReactivePowerFactory_FromDto(t *testing.T) {
    factory := units.ElectricReactivePowerFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricReactivePowerDto{
        Value: 100,
        Unit:  units.ElectricReactivePowerVoltampereReactive,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricReactivePowerDto{
        Value: math.NaN(),
        Unit:  units.ElectricReactivePowerVoltampereReactive,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test VoltampereReactive conversion
    voltamperes_reactiveDto := units.ElectricReactivePowerDto{
        Value: 100,
        Unit:  units.ElectricReactivePowerVoltampereReactive,
    }
    
    var voltamperes_reactiveResult *units.ElectricReactivePower
    voltamperes_reactiveResult, err = factory.FromDto(voltamperes_reactiveDto)
    if err != nil {
        t.Errorf("FromDto() with VoltampereReactive returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = voltamperes_reactiveResult.Convert(units.ElectricReactivePowerVoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltampereReactive = %v, want %v", converted, 100)
    }
    // Test KilovoltampereReactive conversion
    kilovoltamperes_reactiveDto := units.ElectricReactivePowerDto{
        Value: 100,
        Unit:  units.ElectricReactivePowerKilovoltampereReactive,
    }
    
    var kilovoltamperes_reactiveResult *units.ElectricReactivePower
    kilovoltamperes_reactiveResult, err = factory.FromDto(kilovoltamperes_reactiveDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltampereReactive returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovoltamperes_reactiveResult.Convert(units.ElectricReactivePowerKilovoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltampereReactive = %v, want %v", converted, 100)
    }
    // Test MegavoltampereReactive conversion
    megavoltamperes_reactiveDto := units.ElectricReactivePowerDto{
        Value: 100,
        Unit:  units.ElectricReactivePowerMegavoltampereReactive,
    }
    
    var megavoltamperes_reactiveResult *units.ElectricReactivePower
    megavoltamperes_reactiveResult, err = factory.FromDto(megavoltamperes_reactiveDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltampereReactive returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavoltamperes_reactiveResult.Convert(units.ElectricReactivePowerMegavoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltampereReactive = %v, want %v", converted, 100)
    }
    // Test GigavoltampereReactive conversion
    gigavoltamperes_reactiveDto := units.ElectricReactivePowerDto{
        Value: 100,
        Unit:  units.ElectricReactivePowerGigavoltampereReactive,
    }
    
    var gigavoltamperes_reactiveResult *units.ElectricReactivePower
    gigavoltamperes_reactiveResult, err = factory.FromDto(gigavoltamperes_reactiveDto)
    if err != nil {
        t.Errorf("FromDto() with GigavoltampereReactive returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigavoltamperes_reactiveResult.Convert(units.ElectricReactivePowerGigavoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigavoltampereReactive = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricReactivePowerDto{
        Value: 0,
        Unit:  units.ElectricReactivePowerVoltampereReactive,
    }
    
    var zeroResult *units.ElectricReactivePower
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricReactivePowerFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricReactivePowerFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "VoltampereReactive"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "VoltampereReactive"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricReactivePowerDto{
        Value: nanValue,
        Unit:  units.ElectricReactivePowerVoltampereReactive,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with VoltampereReactive unit
    voltamperes_reactiveJSON := []byte(`{"value": 100, "unit": "VoltampereReactive"}`)
    voltamperes_reactiveResult, err := factory.FromDtoJSON(voltamperes_reactiveJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltampereReactive unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = voltamperes_reactiveResult.Convert(units.ElectricReactivePowerVoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltampereReactive = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltampereReactive unit
    kilovoltamperes_reactiveJSON := []byte(`{"value": 100, "unit": "KilovoltampereReactive"}`)
    kilovoltamperes_reactiveResult, err := factory.FromDtoJSON(kilovoltamperes_reactiveJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltampereReactive unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovoltamperes_reactiveResult.Convert(units.ElectricReactivePowerKilovoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltampereReactive = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltampereReactive unit
    megavoltamperes_reactiveJSON := []byte(`{"value": 100, "unit": "MegavoltampereReactive"}`)
    megavoltamperes_reactiveResult, err := factory.FromDtoJSON(megavoltamperes_reactiveJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltampereReactive unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavoltamperes_reactiveResult.Convert(units.ElectricReactivePowerMegavoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltampereReactive = %v, want %v", converted, 100)
    }
    // Test JSON with GigavoltampereReactive unit
    gigavoltamperes_reactiveJSON := []byte(`{"value": 100, "unit": "GigavoltampereReactive"}`)
    gigavoltamperes_reactiveResult, err := factory.FromDtoJSON(gigavoltamperes_reactiveJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigavoltampereReactive unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigavoltamperes_reactiveResult.Convert(units.ElectricReactivePowerGigavoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigavoltampereReactive = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "VoltampereReactive"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromVoltamperesReactive function
func TestElectricReactivePowerFactory_FromVoltamperesReactive(t *testing.T) {
    factory := units.ElectricReactivePowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltamperesReactive(100)
    if err != nil {
        t.Errorf("FromVoltamperesReactive() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricReactivePowerVoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltamperesReactive() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltamperesReactive(math.NaN())
    if err == nil {
        t.Error("FromVoltamperesReactive() with NaN value should return error")
    }

    _, err = factory.FromVoltamperesReactive(math.Inf(1))
    if err == nil {
        t.Error("FromVoltamperesReactive() with +Inf value should return error")
    }

    _, err = factory.FromVoltamperesReactive(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltamperesReactive() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltamperesReactive(0)
    if err != nil {
        t.Errorf("FromVoltamperesReactive() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricReactivePowerVoltampereReactive)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltamperesReactive() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltamperesReactive function
func TestElectricReactivePowerFactory_FromKilovoltamperesReactive(t *testing.T) {
    factory := units.ElectricReactivePowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltamperesReactive(100)
    if err != nil {
        t.Errorf("FromKilovoltamperesReactive() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricReactivePowerKilovoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltamperesReactive() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltamperesReactive(math.NaN())
    if err == nil {
        t.Error("FromKilovoltamperesReactive() with NaN value should return error")
    }

    _, err = factory.FromKilovoltamperesReactive(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltamperesReactive() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltamperesReactive(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltamperesReactive() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltamperesReactive(0)
    if err != nil {
        t.Errorf("FromKilovoltamperesReactive() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricReactivePowerKilovoltampereReactive)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltamperesReactive() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltamperesReactive function
func TestElectricReactivePowerFactory_FromMegavoltamperesReactive(t *testing.T) {
    factory := units.ElectricReactivePowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltamperesReactive(100)
    if err != nil {
        t.Errorf("FromMegavoltamperesReactive() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricReactivePowerMegavoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltamperesReactive() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltamperesReactive(math.NaN())
    if err == nil {
        t.Error("FromMegavoltamperesReactive() with NaN value should return error")
    }

    _, err = factory.FromMegavoltamperesReactive(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltamperesReactive() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltamperesReactive(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltamperesReactive() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltamperesReactive(0)
    if err != nil {
        t.Errorf("FromMegavoltamperesReactive() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricReactivePowerMegavoltampereReactive)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltamperesReactive() with zero value = %v, want 0", converted)
    }
}
// Test FromGigavoltamperesReactive function
func TestElectricReactivePowerFactory_FromGigavoltamperesReactive(t *testing.T) {
    factory := units.ElectricReactivePowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigavoltamperesReactive(100)
    if err != nil {
        t.Errorf("FromGigavoltamperesReactive() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricReactivePowerGigavoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigavoltamperesReactive() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigavoltamperesReactive(math.NaN())
    if err == nil {
        t.Error("FromGigavoltamperesReactive() with NaN value should return error")
    }

    _, err = factory.FromGigavoltamperesReactive(math.Inf(1))
    if err == nil {
        t.Error("FromGigavoltamperesReactive() with +Inf value should return error")
    }

    _, err = factory.FromGigavoltamperesReactive(math.Inf(-1))
    if err == nil {
        t.Error("FromGigavoltamperesReactive() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigavoltamperesReactive(0)
    if err != nil {
        t.Errorf("FromGigavoltamperesReactive() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricReactivePowerGigavoltampereReactive)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigavoltamperesReactive() with zero value = %v, want 0", converted)
    }
}

func TestElectricReactivePowerToString(t *testing.T) {
	factory := units.ElectricReactivePowerFactory{}
	a, err := factory.CreateElectricReactivePower(45, units.ElectricReactivePowerVoltampereReactive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricReactivePowerVoltampereReactive, 2)
	expected := "45.00 " + units.GetElectricReactivePowerAbbreviation(units.ElectricReactivePowerVoltampereReactive)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricReactivePowerVoltampereReactive, -1)
	expected = "45 " + units.GetElectricReactivePowerAbbreviation(units.ElectricReactivePowerVoltampereReactive)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricReactivePower_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricReactivePowerFactory{}
	a1, _ := factory.CreateElectricReactivePower(60, units.ElectricReactivePowerVoltampereReactive)
	a2, _ := factory.CreateElectricReactivePower(60, units.ElectricReactivePowerVoltampereReactive)
	a3, _ := factory.CreateElectricReactivePower(120, units.ElectricReactivePowerVoltampereReactive)

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

func TestElectricReactivePower_Arithmetic(t *testing.T) {
	factory := units.ElectricReactivePowerFactory{}
	a1, _ := factory.CreateElectricReactivePower(30, units.ElectricReactivePowerVoltampereReactive)
	a2, _ := factory.CreateElectricReactivePower(45, units.ElectricReactivePowerVoltampereReactive)

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