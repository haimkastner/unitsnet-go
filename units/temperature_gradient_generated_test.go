// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestTemperatureGradientDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "KelvinPerMeter"}`
	
	factory := units.TemperatureGradientDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.TemperatureGradientKelvinPerMeter {
		t.Errorf("expected unit %v, got %v", units.TemperatureGradientKelvinPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "KelvinPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestTemperatureGradientDto_ToJSON(t *testing.T) {
	dto := units.TemperatureGradientDto{
		Value: 45,
		Unit:  units.TemperatureGradientKelvinPerMeter,
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
	if result["unit"].(string) != string(units.TemperatureGradientKelvinPerMeter) {
		t.Errorf("expected unit %s, got %v", units.TemperatureGradientKelvinPerMeter, result["unit"])
	}
}

func TestNewTemperatureGradient_InvalidValue(t *testing.T) {
	factory := units.TemperatureGradientFactory{}
	// NaN value should return an error.
	_, err := factory.CreateTemperatureGradient(math.NaN(), units.TemperatureGradientKelvinPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateTemperatureGradient(math.Inf(1), units.TemperatureGradientKelvinPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestTemperatureGradientConversions(t *testing.T) {
	factory := units.TemperatureGradientFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateTemperatureGradient(180, units.TemperatureGradientKelvinPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to KelvinsPerMeter.
		// No expected conversion value provided for KelvinsPerMeter, verifying result is not NaN.
		result := a.KelvinsPerMeter()
		cacheResult := a.KelvinsPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KelvinsPerMeter returned NaN")
		}
	}
	{
		// Test conversion to DegreesCelsiusPerMeter.
		// No expected conversion value provided for DegreesCelsiusPerMeter, verifying result is not NaN.
		result := a.DegreesCelsiusPerMeter()
		cacheResult := a.DegreesCelsiusPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesCelsiusPerMeter returned NaN")
		}
	}
	{
		// Test conversion to DegreesFahrenheitPerFoot.
		// No expected conversion value provided for DegreesFahrenheitPerFoot, verifying result is not NaN.
		result := a.DegreesFahrenheitPerFoot()
		cacheResult := a.DegreesFahrenheitPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesFahrenheitPerFoot returned NaN")
		}
	}
	{
		// Test conversion to DegreesCelsiusPerKilometer.
		// No expected conversion value provided for DegreesCelsiusPerKilometer, verifying result is not NaN.
		result := a.DegreesCelsiusPerKilometer()
		cacheResult := a.DegreesCelsiusPerKilometer()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DegreesCelsiusPerKilometer returned NaN")
		}
	}
}

func TestTemperatureGradient_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.TemperatureGradientFactory{}
	a, err := factory.CreateTemperatureGradient(90, units.TemperatureGradientKelvinPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.TemperatureGradientKelvinPerMeter {
		t.Errorf("expected default unit KelvinPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.TemperatureGradientKelvinPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.TemperatureGradientDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.TemperatureGradientKelvinPerMeter {
		t.Errorf("expected unit KelvinPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestTemperatureGradientFactory_FromDto(t *testing.T) {
    factory := units.TemperatureGradientFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.TemperatureGradientDto{
        Value: 100,
        Unit:  units.TemperatureGradientKelvinPerMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.TemperatureGradientDto{
        Value: math.NaN(),
        Unit:  units.TemperatureGradientKelvinPerMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test KelvinPerMeter conversion
    kelvins_per_meterDto := units.TemperatureGradientDto{
        Value: 100,
        Unit:  units.TemperatureGradientKelvinPerMeter,
    }
    
    var kelvins_per_meterResult *units.TemperatureGradient
    kelvins_per_meterResult, err = factory.FromDto(kelvins_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KelvinPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kelvins_per_meterResult.Convert(units.TemperatureGradientKelvinPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KelvinPerMeter = %v, want %v", converted, 100)
    }
    // Test DegreeCelsiusPerMeter conversion
    degrees_celsius_per_meterDto := units.TemperatureGradientDto{
        Value: 100,
        Unit:  units.TemperatureGradientDegreeCelsiusPerMeter,
    }
    
    var degrees_celsius_per_meterResult *units.TemperatureGradient
    degrees_celsius_per_meterResult, err = factory.FromDto(degrees_celsius_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeCelsiusPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsius_per_meterResult.Convert(units.TemperatureGradientDegreeCelsiusPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsiusPerMeter = %v, want %v", converted, 100)
    }
    // Test DegreeFahrenheitPerFoot conversion
    degrees_fahrenheit_per_footDto := units.TemperatureGradientDto{
        Value: 100,
        Unit:  units.TemperatureGradientDegreeFahrenheitPerFoot,
    }
    
    var degrees_fahrenheit_per_footResult *units.TemperatureGradient
    degrees_fahrenheit_per_footResult, err = factory.FromDto(degrees_fahrenheit_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeFahrenheitPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_fahrenheit_per_footResult.Convert(units.TemperatureGradientDegreeFahrenheitPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeFahrenheitPerFoot = %v, want %v", converted, 100)
    }
    // Test DegreeCelsiusPerKilometer conversion
    degrees_celsius_per_kilometerDto := units.TemperatureGradientDto{
        Value: 100,
        Unit:  units.TemperatureGradientDegreeCelsiusPerKilometer,
    }
    
    var degrees_celsius_per_kilometerResult *units.TemperatureGradient
    degrees_celsius_per_kilometerResult, err = factory.FromDto(degrees_celsius_per_kilometerDto)
    if err != nil {
        t.Errorf("FromDto() with DegreeCelsiusPerKilometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsius_per_kilometerResult.Convert(units.TemperatureGradientDegreeCelsiusPerKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsiusPerKilometer = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.TemperatureGradientDto{
        Value: 0,
        Unit:  units.TemperatureGradientKelvinPerMeter,
    }
    
    var zeroResult *units.TemperatureGradient
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestTemperatureGradientFactory_FromDtoJSON(t *testing.T) {
    factory := units.TemperatureGradientFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "KelvinPerMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "KelvinPerMeter"}`)
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
    nanJSON, _ := json.Marshal(units.TemperatureGradientDto{
        Value: nanValue,
        Unit:  units.TemperatureGradientKelvinPerMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with KelvinPerMeter unit
    kelvins_per_meterJSON := []byte(`{"value": 100, "unit": "KelvinPerMeter"}`)
    kelvins_per_meterResult, err := factory.FromDtoJSON(kelvins_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KelvinPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kelvins_per_meterResult.Convert(units.TemperatureGradientKelvinPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KelvinPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeCelsiusPerMeter unit
    degrees_celsius_per_meterJSON := []byte(`{"value": 100, "unit": "DegreeCelsiusPerMeter"}`)
    degrees_celsius_per_meterResult, err := factory.FromDtoJSON(degrees_celsius_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeCelsiusPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsius_per_meterResult.Convert(units.TemperatureGradientDegreeCelsiusPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsiusPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeFahrenheitPerFoot unit
    degrees_fahrenheit_per_footJSON := []byte(`{"value": 100, "unit": "DegreeFahrenheitPerFoot"}`)
    degrees_fahrenheit_per_footResult, err := factory.FromDtoJSON(degrees_fahrenheit_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeFahrenheitPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_fahrenheit_per_footResult.Convert(units.TemperatureGradientDegreeFahrenheitPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeFahrenheitPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with DegreeCelsiusPerKilometer unit
    degrees_celsius_per_kilometerJSON := []byte(`{"value": 100, "unit": "DegreeCelsiusPerKilometer"}`)
    degrees_celsius_per_kilometerResult, err := factory.FromDtoJSON(degrees_celsius_per_kilometerJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DegreeCelsiusPerKilometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degrees_celsius_per_kilometerResult.Convert(units.TemperatureGradientDegreeCelsiusPerKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DegreeCelsiusPerKilometer = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "KelvinPerMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromKelvinsPerMeter function
func TestTemperatureGradientFactory_FromKelvinsPerMeter(t *testing.T) {
    factory := units.TemperatureGradientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKelvinsPerMeter(100)
    if err != nil {
        t.Errorf("FromKelvinsPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureGradientKelvinPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKelvinsPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKelvinsPerMeter(math.NaN())
    if err == nil {
        t.Error("FromKelvinsPerMeter() with NaN value should return error")
    }

    _, err = factory.FromKelvinsPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKelvinsPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromKelvinsPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKelvinsPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKelvinsPerMeter(0)
    if err != nil {
        t.Errorf("FromKelvinsPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureGradientKelvinPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKelvinsPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesCelsiusPerMeter function
func TestTemperatureGradientFactory_FromDegreesCelsiusPerMeter(t *testing.T) {
    factory := units.TemperatureGradientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesCelsiusPerMeter(100)
    if err != nil {
        t.Errorf("FromDegreesCelsiusPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureGradientDegreeCelsiusPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesCelsiusPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesCelsiusPerMeter(math.NaN())
    if err == nil {
        t.Error("FromDegreesCelsiusPerMeter() with NaN value should return error")
    }

    _, err = factory.FromDegreesCelsiusPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesCelsiusPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromDegreesCelsiusPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesCelsiusPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesCelsiusPerMeter(0)
    if err != nil {
        t.Errorf("FromDegreesCelsiusPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureGradientDegreeCelsiusPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesCelsiusPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesFahrenheitPerFoot function
func TestTemperatureGradientFactory_FromDegreesFahrenheitPerFoot(t *testing.T) {
    factory := units.TemperatureGradientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesFahrenheitPerFoot(100)
    if err != nil {
        t.Errorf("FromDegreesFahrenheitPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureGradientDegreeFahrenheitPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesFahrenheitPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesFahrenheitPerFoot(math.NaN())
    if err == nil {
        t.Error("FromDegreesFahrenheitPerFoot() with NaN value should return error")
    }

    _, err = factory.FromDegreesFahrenheitPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesFahrenheitPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromDegreesFahrenheitPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesFahrenheitPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesFahrenheitPerFoot(0)
    if err != nil {
        t.Errorf("FromDegreesFahrenheitPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureGradientDegreeFahrenheitPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesFahrenheitPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromDegreesCelsiusPerKilometer function
func TestTemperatureGradientFactory_FromDegreesCelsiusPerKilometer(t *testing.T) {
    factory := units.TemperatureGradientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegreesCelsiusPerKilometer(100)
    if err != nil {
        t.Errorf("FromDegreesCelsiusPerKilometer() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TemperatureGradientDegreeCelsiusPerKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegreesCelsiusPerKilometer() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegreesCelsiusPerKilometer(math.NaN())
    if err == nil {
        t.Error("FromDegreesCelsiusPerKilometer() with NaN value should return error")
    }

    _, err = factory.FromDegreesCelsiusPerKilometer(math.Inf(1))
    if err == nil {
        t.Error("FromDegreesCelsiusPerKilometer() with +Inf value should return error")
    }

    _, err = factory.FromDegreesCelsiusPerKilometer(math.Inf(-1))
    if err == nil {
        t.Error("FromDegreesCelsiusPerKilometer() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegreesCelsiusPerKilometer(0)
    if err != nil {
        t.Errorf("FromDegreesCelsiusPerKilometer() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TemperatureGradientDegreeCelsiusPerKilometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegreesCelsiusPerKilometer() with zero value = %v, want 0", converted)
    }
}

func TestTemperatureGradientToString(t *testing.T) {
	factory := units.TemperatureGradientFactory{}
	a, err := factory.CreateTemperatureGradient(45, units.TemperatureGradientKelvinPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.TemperatureGradientKelvinPerMeter, 2)
	expected := "45.00 " + units.GetTemperatureGradientAbbreviation(units.TemperatureGradientKelvinPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.TemperatureGradientKelvinPerMeter, -1)
	expected = "45 " + units.GetTemperatureGradientAbbreviation(units.TemperatureGradientKelvinPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestTemperatureGradient_EqualityAndComparison(t *testing.T) {
	factory := units.TemperatureGradientFactory{}
	a1, _ := factory.CreateTemperatureGradient(60, units.TemperatureGradientKelvinPerMeter)
	a2, _ := factory.CreateTemperatureGradient(60, units.TemperatureGradientKelvinPerMeter)
	a3, _ := factory.CreateTemperatureGradient(120, units.TemperatureGradientKelvinPerMeter)

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

func TestTemperatureGradient_Arithmetic(t *testing.T) {
	factory := units.TemperatureGradientFactory{}
	a1, _ := factory.CreateTemperatureGradient(30, units.TemperatureGradientKelvinPerMeter)
	a2, _ := factory.CreateTemperatureGradient(45, units.TemperatureGradientKelvinPerMeter)

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


func TestGetTemperatureGradientAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.TemperatureGradientUnits
        want string
    }{
        {
            name: "KelvinPerMeter abbreviation",
            unit: units.TemperatureGradientKelvinPerMeter,
            want: "∆°K/m",
        },
        {
            name: "DegreeCelsiusPerMeter abbreviation",
            unit: units.TemperatureGradientDegreeCelsiusPerMeter,
            want: "∆°C/m",
        },
        {
            name: "DegreeFahrenheitPerFoot abbreviation",
            unit: units.TemperatureGradientDegreeFahrenheitPerFoot,
            want: "∆°F/ft",
        },
        {
            name: "DegreeCelsiusPerKilometer abbreviation",
            unit: units.TemperatureGradientDegreeCelsiusPerKilometer,
            want: "∆°C/km",
        },
        {
            name: "invalid unit",
            unit: units.TemperatureGradientUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetTemperatureGradientAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetTemperatureGradientAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestTemperatureGradient_String(t *testing.T) {
    factory := units.TemperatureGradientFactory{}
    
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
            unit, err := factory.CreateTemperatureGradient(tt.value, units.TemperatureGradientKelvinPerMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("TemperatureGradient.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestTemperatureGradient_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.TemperatureGradientFactory{}

	_, err := uf.CreateTemperatureGradient(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}