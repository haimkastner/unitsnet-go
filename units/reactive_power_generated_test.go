// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestReactivePowerDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltampereReactive"}`
	
	factory := units.ReactivePowerDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ReactivePowerVoltampereReactive {
		t.Errorf("expected unit %v, got %v", units.ReactivePowerVoltampereReactive, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltampereReactive"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestReactivePowerDto_ToJSON(t *testing.T) {
	dto := units.ReactivePowerDto{
		Value: 45,
		Unit:  units.ReactivePowerVoltampereReactive,
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
	if result["unit"].(string) != string(units.ReactivePowerVoltampereReactive) {
		t.Errorf("expected unit %s, got %v", units.ReactivePowerVoltampereReactive, result["unit"])
	}
}

func TestNewReactivePower_InvalidValue(t *testing.T) {
	factory := units.ReactivePowerFactory{}
	// NaN value should return an error.
	_, err := factory.CreateReactivePower(math.NaN(), units.ReactivePowerVoltampereReactive)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateReactivePower(math.Inf(1), units.ReactivePowerVoltampereReactive)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestReactivePowerConversions(t *testing.T) {
	factory := units.ReactivePowerFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateReactivePower(180, units.ReactivePowerVoltampereReactive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltamperesReactive.
		// No expected conversion value provided for VoltamperesReactive, verifying result is not NaN.
		result := a.VoltamperesReactive()
		cacheResult := a.VoltamperesReactive()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to VoltamperesReactive returned NaN")
		}
	}
	{
		// Test conversion to KilovoltamperesReactive.
		// No expected conversion value provided for KilovoltamperesReactive, verifying result is not NaN.
		result := a.KilovoltamperesReactive()
		cacheResult := a.KilovoltamperesReactive()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilovoltamperesReactive returned NaN")
		}
	}
	{
		// Test conversion to MegavoltamperesReactive.
		// No expected conversion value provided for MegavoltamperesReactive, verifying result is not NaN.
		result := a.MegavoltamperesReactive()
		cacheResult := a.MegavoltamperesReactive()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegavoltamperesReactive returned NaN")
		}
	}
	{
		// Test conversion to GigavoltamperesReactive.
		// No expected conversion value provided for GigavoltamperesReactive, verifying result is not NaN.
		result := a.GigavoltamperesReactive()
		cacheResult := a.GigavoltamperesReactive()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigavoltamperesReactive returned NaN")
		}
	}
}

func TestReactivePower_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ReactivePowerFactory{}
	a, err := factory.CreateReactivePower(90, units.ReactivePowerVoltampereReactive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ReactivePowerVoltampereReactive {
		t.Errorf("expected default unit VoltampereReactive, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ReactivePowerVoltampereReactive
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ReactivePowerDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ReactivePowerVoltampereReactive {
		t.Errorf("expected unit VoltampereReactive, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestReactivePowerFactory_FromDto(t *testing.T) {
    factory := units.ReactivePowerFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ReactivePowerDto{
        Value: 100,
        Unit:  units.ReactivePowerVoltampereReactive,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ReactivePowerDto{
        Value: math.NaN(),
        Unit:  units.ReactivePowerVoltampereReactive,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test VoltampereReactive conversion
    voltamperes_reactiveDto := units.ReactivePowerDto{
        Value: 100,
        Unit:  units.ReactivePowerVoltampereReactive,
    }
    
    var voltamperes_reactiveResult *units.ReactivePower
    voltamperes_reactiveResult, err = factory.FromDto(voltamperes_reactiveDto)
    if err != nil {
        t.Errorf("FromDto() with VoltampereReactive returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = voltamperes_reactiveResult.Convert(units.ReactivePowerVoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltampereReactive = %v, want %v", converted, 100)
    }
    // Test KilovoltampereReactive conversion
    kilovoltamperes_reactiveDto := units.ReactivePowerDto{
        Value: 100,
        Unit:  units.ReactivePowerKilovoltampereReactive,
    }
    
    var kilovoltamperes_reactiveResult *units.ReactivePower
    kilovoltamperes_reactiveResult, err = factory.FromDto(kilovoltamperes_reactiveDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltampereReactive returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovoltamperes_reactiveResult.Convert(units.ReactivePowerKilovoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltampereReactive = %v, want %v", converted, 100)
    }
    // Test MegavoltampereReactive conversion
    megavoltamperes_reactiveDto := units.ReactivePowerDto{
        Value: 100,
        Unit:  units.ReactivePowerMegavoltampereReactive,
    }
    
    var megavoltamperes_reactiveResult *units.ReactivePower
    megavoltamperes_reactiveResult, err = factory.FromDto(megavoltamperes_reactiveDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltampereReactive returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavoltamperes_reactiveResult.Convert(units.ReactivePowerMegavoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltampereReactive = %v, want %v", converted, 100)
    }
    // Test GigavoltampereReactive conversion
    gigavoltamperes_reactiveDto := units.ReactivePowerDto{
        Value: 100,
        Unit:  units.ReactivePowerGigavoltampereReactive,
    }
    
    var gigavoltamperes_reactiveResult *units.ReactivePower
    gigavoltamperes_reactiveResult, err = factory.FromDto(gigavoltamperes_reactiveDto)
    if err != nil {
        t.Errorf("FromDto() with GigavoltampereReactive returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigavoltamperes_reactiveResult.Convert(units.ReactivePowerGigavoltampereReactive)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigavoltampereReactive = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ReactivePowerDto{
        Value: 0,
        Unit:  units.ReactivePowerVoltampereReactive,
    }
    
    var zeroResult *units.ReactivePower
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestReactivePowerFactory_FromDtoJSON(t *testing.T) {
    factory := units.ReactivePowerFactory{}
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
    nanJSON, _ := json.Marshal(units.ReactivePowerDto{
        Value: nanValue,
        Unit:  units.ReactivePowerVoltampereReactive,
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
    converted = voltamperes_reactiveResult.Convert(units.ReactivePowerVoltampereReactive)
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
    converted = kilovoltamperes_reactiveResult.Convert(units.ReactivePowerKilovoltampereReactive)
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
    converted = megavoltamperes_reactiveResult.Convert(units.ReactivePowerMegavoltampereReactive)
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
    converted = gigavoltamperes_reactiveResult.Convert(units.ReactivePowerGigavoltampereReactive)
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
func TestReactivePowerFactory_FromVoltamperesReactive(t *testing.T) {
    factory := units.ReactivePowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltamperesReactive(100)
    if err != nil {
        t.Errorf("FromVoltamperesReactive() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReactivePowerVoltampereReactive)
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
    converted = zeroResult.Convert(units.ReactivePowerVoltampereReactive)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltamperesReactive() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltamperesReactive function
func TestReactivePowerFactory_FromKilovoltamperesReactive(t *testing.T) {
    factory := units.ReactivePowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltamperesReactive(100)
    if err != nil {
        t.Errorf("FromKilovoltamperesReactive() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReactivePowerKilovoltampereReactive)
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
    converted = zeroResult.Convert(units.ReactivePowerKilovoltampereReactive)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltamperesReactive() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltamperesReactive function
func TestReactivePowerFactory_FromMegavoltamperesReactive(t *testing.T) {
    factory := units.ReactivePowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltamperesReactive(100)
    if err != nil {
        t.Errorf("FromMegavoltamperesReactive() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReactivePowerMegavoltampereReactive)
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
    converted = zeroResult.Convert(units.ReactivePowerMegavoltampereReactive)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltamperesReactive() with zero value = %v, want 0", converted)
    }
}
// Test FromGigavoltamperesReactive function
func TestReactivePowerFactory_FromGigavoltamperesReactive(t *testing.T) {
    factory := units.ReactivePowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigavoltamperesReactive(100)
    if err != nil {
        t.Errorf("FromGigavoltamperesReactive() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ReactivePowerGigavoltampereReactive)
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
    converted = zeroResult.Convert(units.ReactivePowerGigavoltampereReactive)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigavoltamperesReactive() with zero value = %v, want 0", converted)
    }
}

func TestReactivePowerToString(t *testing.T) {
	factory := units.ReactivePowerFactory{}
	a, err := factory.CreateReactivePower(45, units.ReactivePowerVoltampereReactive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ReactivePowerVoltampereReactive, 2)
	expected := "45.00 " + units.GetReactivePowerAbbreviation(units.ReactivePowerVoltampereReactive)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ReactivePowerVoltampereReactive, -1)
	expected = "45 " + units.GetReactivePowerAbbreviation(units.ReactivePowerVoltampereReactive)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestReactivePower_EqualityAndComparison(t *testing.T) {
	factory := units.ReactivePowerFactory{}
	a1, _ := factory.CreateReactivePower(60, units.ReactivePowerVoltampereReactive)
	a2, _ := factory.CreateReactivePower(60, units.ReactivePowerVoltampereReactive)
	a3, _ := factory.CreateReactivePower(120, units.ReactivePowerVoltampereReactive)

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

func TestReactivePower_Arithmetic(t *testing.T) {
	factory := units.ReactivePowerFactory{}
	a1, _ := factory.CreateReactivePower(30, units.ReactivePowerVoltampereReactive)
	a2, _ := factory.CreateReactivePower(45, units.ReactivePowerVoltampereReactive)

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


func TestGetReactivePowerAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ReactivePowerUnits
        want string
    }{
        {
            name: "VoltampereReactive abbreviation",
            unit: units.ReactivePowerVoltampereReactive,
            want: "var",
        },
        {
            name: "KilovoltampereReactive abbreviation",
            unit: units.ReactivePowerKilovoltampereReactive,
            want: "kvar",
        },
        {
            name: "MegavoltampereReactive abbreviation",
            unit: units.ReactivePowerMegavoltampereReactive,
            want: "Mvar",
        },
        {
            name: "GigavoltampereReactive abbreviation",
            unit: units.ReactivePowerGigavoltampereReactive,
            want: "Gvar",
        },
        {
            name: "invalid unit",
            unit: units.ReactivePowerUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetReactivePowerAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetReactivePowerAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestReactivePower_String(t *testing.T) {
    factory := units.ReactivePowerFactory{}
    
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
            unit, err := factory.CreateReactivePower(tt.value, units.ReactivePowerVoltampereReactive)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ReactivePower.String() = %v, want %v", got, tt.want)
            }
        })
    }
}