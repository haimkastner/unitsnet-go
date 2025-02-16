// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestThermalConductivityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "WattPerMeterKelvin"}`
	
	factory := units.ThermalConductivityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ThermalConductivityWattPerMeterKelvin {
		t.Errorf("expected unit %v, got %v", units.ThermalConductivityWattPerMeterKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "WattPerMeterKelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestThermalConductivityDto_ToJSON(t *testing.T) {
	dto := units.ThermalConductivityDto{
		Value: 45,
		Unit:  units.ThermalConductivityWattPerMeterKelvin,
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
	if result["unit"].(string) != string(units.ThermalConductivityWattPerMeterKelvin) {
		t.Errorf("expected unit %s, got %v", units.ThermalConductivityWattPerMeterKelvin, result["unit"])
	}
}

func TestNewThermalConductivity_InvalidValue(t *testing.T) {
	factory := units.ThermalConductivityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateThermalConductivity(math.NaN(), units.ThermalConductivityWattPerMeterKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateThermalConductivity(math.Inf(1), units.ThermalConductivityWattPerMeterKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestThermalConductivityConversions(t *testing.T) {
	factory := units.ThermalConductivityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateThermalConductivity(180, units.ThermalConductivityWattPerMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to WattsPerMeterKelvin.
		// No expected conversion value provided for WattsPerMeterKelvin, verifying result is not NaN.
		result := a.WattsPerMeterKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerMeterKelvin returned NaN")
		}
	}
	{
		// Test conversion to BtusPerHourFootFahrenheit.
		// No expected conversion value provided for BtusPerHourFootFahrenheit, verifying result is not NaN.
		result := a.BtusPerHourFootFahrenheit()
		if math.IsNaN(result) {
			t.Errorf("conversion to BtusPerHourFootFahrenheit returned NaN")
		}
	}
}

func TestThermalConductivity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ThermalConductivityFactory{}
	a, err := factory.CreateThermalConductivity(90, units.ThermalConductivityWattPerMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ThermalConductivityWattPerMeterKelvin {
		t.Errorf("expected default unit WattPerMeterKelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ThermalConductivityWattPerMeterKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ThermalConductivityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ThermalConductivityWattPerMeterKelvin {
		t.Errorf("expected unit WattPerMeterKelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestThermalConductivityFactory_FromDto(t *testing.T) {
    factory := units.ThermalConductivityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ThermalConductivityDto{
        Value: 100,
        Unit:  units.ThermalConductivityWattPerMeterKelvin,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ThermalConductivityDto{
        Value: math.NaN(),
        Unit:  units.ThermalConductivityWattPerMeterKelvin,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test WattPerMeterKelvin conversion
    watts_per_meter_kelvinDto := units.ThermalConductivityDto{
        Value: 100,
        Unit:  units.ThermalConductivityWattPerMeterKelvin,
    }
    
    var watts_per_meter_kelvinResult *units.ThermalConductivity
    watts_per_meter_kelvinResult, err = factory.FromDto(watts_per_meter_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerMeterKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_meter_kelvinResult.Convert(units.ThermalConductivityWattPerMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerMeterKelvin = %v, want %v", converted, 100)
    }
    // Test BtuPerHourFootFahrenheit conversion
    btus_per_hour_foot_fahrenheitDto := units.ThermalConductivityDto{
        Value: 100,
        Unit:  units.ThermalConductivityBtuPerHourFootFahrenheit,
    }
    
    var btus_per_hour_foot_fahrenheitResult *units.ThermalConductivity
    btus_per_hour_foot_fahrenheitResult, err = factory.FromDto(btus_per_hour_foot_fahrenheitDto)
    if err != nil {
        t.Errorf("FromDto() with BtuPerHourFootFahrenheit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_hour_foot_fahrenheitResult.Convert(units.ThermalConductivityBtuPerHourFootFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerHourFootFahrenheit = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ThermalConductivityDto{
        Value: 0,
        Unit:  units.ThermalConductivityWattPerMeterKelvin,
    }
    
    var zeroResult *units.ThermalConductivity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestThermalConductivityFactory_FromDtoJSON(t *testing.T) {
    factory := units.ThermalConductivityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "WattPerMeterKelvin"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "WattPerMeterKelvin"}`)
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
    nanJSON, _ := json.Marshal(units.ThermalConductivityDto{
        Value: nanValue,
        Unit:  units.ThermalConductivityWattPerMeterKelvin,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with WattPerMeterKelvin unit
    watts_per_meter_kelvinJSON := []byte(`{"value": 100, "unit": "WattPerMeterKelvin"}`)
    watts_per_meter_kelvinResult, err := factory.FromDtoJSON(watts_per_meter_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerMeterKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_meter_kelvinResult.Convert(units.ThermalConductivityWattPerMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerMeterKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with BtuPerHourFootFahrenheit unit
    btus_per_hour_foot_fahrenheitJSON := []byte(`{"value": 100, "unit": "BtuPerHourFootFahrenheit"}`)
    btus_per_hour_foot_fahrenheitResult, err := factory.FromDtoJSON(btus_per_hour_foot_fahrenheitJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BtuPerHourFootFahrenheit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_hour_foot_fahrenheitResult.Convert(units.ThermalConductivityBtuPerHourFootFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerHourFootFahrenheit = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "WattPerMeterKelvin"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromWattsPerMeterKelvin function
func TestThermalConductivityFactory_FromWattsPerMeterKelvin(t *testing.T) {
    factory := units.ThermalConductivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerMeterKelvin(100)
    if err != nil {
        t.Errorf("FromWattsPerMeterKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalConductivityWattPerMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerMeterKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerMeterKelvin(math.NaN())
    if err == nil {
        t.Error("FromWattsPerMeterKelvin() with NaN value should return error")
    }

    _, err = factory.FromWattsPerMeterKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerMeterKelvin() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerMeterKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerMeterKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerMeterKelvin(0)
    if err != nil {
        t.Errorf("FromWattsPerMeterKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ThermalConductivityWattPerMeterKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerMeterKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromBtusPerHourFootFahrenheit function
func TestThermalConductivityFactory_FromBtusPerHourFootFahrenheit(t *testing.T) {
    factory := units.ThermalConductivityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBtusPerHourFootFahrenheit(100)
    if err != nil {
        t.Errorf("FromBtusPerHourFootFahrenheit() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalConductivityBtuPerHourFootFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBtusPerHourFootFahrenheit() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBtusPerHourFootFahrenheit(math.NaN())
    if err == nil {
        t.Error("FromBtusPerHourFootFahrenheit() with NaN value should return error")
    }

    _, err = factory.FromBtusPerHourFootFahrenheit(math.Inf(1))
    if err == nil {
        t.Error("FromBtusPerHourFootFahrenheit() with +Inf value should return error")
    }

    _, err = factory.FromBtusPerHourFootFahrenheit(math.Inf(-1))
    if err == nil {
        t.Error("FromBtusPerHourFootFahrenheit() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBtusPerHourFootFahrenheit(0)
    if err != nil {
        t.Errorf("FromBtusPerHourFootFahrenheit() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ThermalConductivityBtuPerHourFootFahrenheit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBtusPerHourFootFahrenheit() with zero value = %v, want 0", converted)
    }
}

func TestThermalConductivityToString(t *testing.T) {
	factory := units.ThermalConductivityFactory{}
	a, err := factory.CreateThermalConductivity(45, units.ThermalConductivityWattPerMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ThermalConductivityWattPerMeterKelvin, 2)
	expected := "45.00 " + units.GetThermalConductivityAbbreviation(units.ThermalConductivityWattPerMeterKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ThermalConductivityWattPerMeterKelvin, -1)
	expected = "45 " + units.GetThermalConductivityAbbreviation(units.ThermalConductivityWattPerMeterKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestThermalConductivity_EqualityAndComparison(t *testing.T) {
	factory := units.ThermalConductivityFactory{}
	a1, _ := factory.CreateThermalConductivity(60, units.ThermalConductivityWattPerMeterKelvin)
	a2, _ := factory.CreateThermalConductivity(60, units.ThermalConductivityWattPerMeterKelvin)
	a3, _ := factory.CreateThermalConductivity(120, units.ThermalConductivityWattPerMeterKelvin)

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

func TestThermalConductivity_Arithmetic(t *testing.T) {
	factory := units.ThermalConductivityFactory{}
	a1, _ := factory.CreateThermalConductivity(30, units.ThermalConductivityWattPerMeterKelvin)
	a2, _ := factory.CreateThermalConductivity(45, units.ThermalConductivityWattPerMeterKelvin)

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