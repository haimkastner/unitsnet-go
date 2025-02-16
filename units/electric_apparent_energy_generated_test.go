// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestElectricApparentEnergyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "VoltampereHour"}`
	
	factory := units.ElectricApparentEnergyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ElectricApparentEnergyVoltampereHour {
		t.Errorf("expected unit %v, got %v", units.ElectricApparentEnergyVoltampereHour, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "VoltampereHour"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestElectricApparentEnergyDto_ToJSON(t *testing.T) {
	dto := units.ElectricApparentEnergyDto{
		Value: 45,
		Unit:  units.ElectricApparentEnergyVoltampereHour,
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
	if result["unit"].(string) != string(units.ElectricApparentEnergyVoltampereHour) {
		t.Errorf("expected unit %s, got %v", units.ElectricApparentEnergyVoltampereHour, result["unit"])
	}
}

func TestNewElectricApparentEnergy_InvalidValue(t *testing.T) {
	factory := units.ElectricApparentEnergyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateElectricApparentEnergy(math.NaN(), units.ElectricApparentEnergyVoltampereHour)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateElectricApparentEnergy(math.Inf(1), units.ElectricApparentEnergyVoltampereHour)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestElectricApparentEnergyConversions(t *testing.T) {
	factory := units.ElectricApparentEnergyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateElectricApparentEnergy(180, units.ElectricApparentEnergyVoltampereHour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to VoltampereHours.
		// No expected conversion value provided for VoltampereHours, verifying result is not NaN.
		result := a.VoltampereHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to VoltampereHours returned NaN")
		}
	}
	{
		// Test conversion to KilovoltampereHours.
		// No expected conversion value provided for KilovoltampereHours, verifying result is not NaN.
		result := a.KilovoltampereHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilovoltampereHours returned NaN")
		}
	}
	{
		// Test conversion to MegavoltampereHours.
		// No expected conversion value provided for MegavoltampereHours, verifying result is not NaN.
		result := a.MegavoltampereHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegavoltampereHours returned NaN")
		}
	}
}

func TestElectricApparentEnergy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ElectricApparentEnergyFactory{}
	a, err := factory.CreateElectricApparentEnergy(90, units.ElectricApparentEnergyVoltampereHour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ElectricApparentEnergyVoltampereHour {
		t.Errorf("expected default unit VoltampereHour, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ElectricApparentEnergyVoltampereHour
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ElectricApparentEnergyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ElectricApparentEnergyVoltampereHour {
		t.Errorf("expected unit VoltampereHour, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestElectricApparentEnergyFactory_FromDto(t *testing.T) {
    factory := units.ElectricApparentEnergyFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ElectricApparentEnergyDto{
        Value: 100,
        Unit:  units.ElectricApparentEnergyVoltampereHour,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ElectricApparentEnergyDto{
        Value: math.NaN(),
        Unit:  units.ElectricApparentEnergyVoltampereHour,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test VoltampereHour conversion
    voltampere_hoursDto := units.ElectricApparentEnergyDto{
        Value: 100,
        Unit:  units.ElectricApparentEnergyVoltampereHour,
    }
    
    var voltampere_hoursResult *units.ElectricApparentEnergy
    voltampere_hoursResult, err = factory.FromDto(voltampere_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with VoltampereHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = voltampere_hoursResult.Convert(units.ElectricApparentEnergyVoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltampereHour = %v, want %v", converted, 100)
    }
    // Test KilovoltampereHour conversion
    kilovoltampere_hoursDto := units.ElectricApparentEnergyDto{
        Value: 100,
        Unit:  units.ElectricApparentEnergyKilovoltampereHour,
    }
    
    var kilovoltampere_hoursResult *units.ElectricApparentEnergy
    kilovoltampere_hoursResult, err = factory.FromDto(kilovoltampere_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with KilovoltampereHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovoltampere_hoursResult.Convert(units.ElectricApparentEnergyKilovoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltampereHour = %v, want %v", converted, 100)
    }
    // Test MegavoltampereHour conversion
    megavoltampere_hoursDto := units.ElectricApparentEnergyDto{
        Value: 100,
        Unit:  units.ElectricApparentEnergyMegavoltampereHour,
    }
    
    var megavoltampere_hoursResult *units.ElectricApparentEnergy
    megavoltampere_hoursResult, err = factory.FromDto(megavoltampere_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with MegavoltampereHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavoltampere_hoursResult.Convert(units.ElectricApparentEnergyMegavoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltampereHour = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ElectricApparentEnergyDto{
        Value: 0,
        Unit:  units.ElectricApparentEnergyVoltampereHour,
    }
    
    var zeroResult *units.ElectricApparentEnergy
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestElectricApparentEnergyFactory_FromDtoJSON(t *testing.T) {
    factory := units.ElectricApparentEnergyFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "VoltampereHour"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "VoltampereHour"}`)
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
    nanJSON, _ := json.Marshal(units.ElectricApparentEnergyDto{
        Value: nanValue,
        Unit:  units.ElectricApparentEnergyVoltampereHour,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with VoltampereHour unit
    voltampere_hoursJSON := []byte(`{"value": 100, "unit": "VoltampereHour"}`)
    voltampere_hoursResult, err := factory.FromDtoJSON(voltampere_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with VoltampereHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = voltampere_hoursResult.Convert(units.ElectricApparentEnergyVoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for VoltampereHour = %v, want %v", converted, 100)
    }
    // Test JSON with KilovoltampereHour unit
    kilovoltampere_hoursJSON := []byte(`{"value": 100, "unit": "KilovoltampereHour"}`)
    kilovoltampere_hoursResult, err := factory.FromDtoJSON(kilovoltampere_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilovoltampereHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilovoltampere_hoursResult.Convert(units.ElectricApparentEnergyKilovoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilovoltampereHour = %v, want %v", converted, 100)
    }
    // Test JSON with MegavoltampereHour unit
    megavoltampere_hoursJSON := []byte(`{"value": 100, "unit": "MegavoltampereHour"}`)
    megavoltampere_hoursResult, err := factory.FromDtoJSON(megavoltampere_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegavoltampereHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megavoltampere_hoursResult.Convert(units.ElectricApparentEnergyMegavoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegavoltampereHour = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "VoltampereHour"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromVoltampereHours function
func TestElectricApparentEnergyFactory_FromVoltampereHours(t *testing.T) {
    factory := units.ElectricApparentEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromVoltampereHours(100)
    if err != nil {
        t.Errorf("FromVoltampereHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricApparentEnergyVoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromVoltampereHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromVoltampereHours(math.NaN())
    if err == nil {
        t.Error("FromVoltampereHours() with NaN value should return error")
    }

    _, err = factory.FromVoltampereHours(math.Inf(1))
    if err == nil {
        t.Error("FromVoltampereHours() with +Inf value should return error")
    }

    _, err = factory.FromVoltampereHours(math.Inf(-1))
    if err == nil {
        t.Error("FromVoltampereHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromVoltampereHours(0)
    if err != nil {
        t.Errorf("FromVoltampereHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricApparentEnergyVoltampereHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromVoltampereHours() with zero value = %v, want 0", converted)
    }
}
// Test FromKilovoltampereHours function
func TestElectricApparentEnergyFactory_FromKilovoltampereHours(t *testing.T) {
    factory := units.ElectricApparentEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilovoltampereHours(100)
    if err != nil {
        t.Errorf("FromKilovoltampereHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricApparentEnergyKilovoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilovoltampereHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilovoltampereHours(math.NaN())
    if err == nil {
        t.Error("FromKilovoltampereHours() with NaN value should return error")
    }

    _, err = factory.FromKilovoltampereHours(math.Inf(1))
    if err == nil {
        t.Error("FromKilovoltampereHours() with +Inf value should return error")
    }

    _, err = factory.FromKilovoltampereHours(math.Inf(-1))
    if err == nil {
        t.Error("FromKilovoltampereHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilovoltampereHours(0)
    if err != nil {
        t.Errorf("FromKilovoltampereHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricApparentEnergyKilovoltampereHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilovoltampereHours() with zero value = %v, want 0", converted)
    }
}
// Test FromMegavoltampereHours function
func TestElectricApparentEnergyFactory_FromMegavoltampereHours(t *testing.T) {
    factory := units.ElectricApparentEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegavoltampereHours(100)
    if err != nil {
        t.Errorf("FromMegavoltampereHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ElectricApparentEnergyMegavoltampereHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegavoltampereHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegavoltampereHours(math.NaN())
    if err == nil {
        t.Error("FromMegavoltampereHours() with NaN value should return error")
    }

    _, err = factory.FromMegavoltampereHours(math.Inf(1))
    if err == nil {
        t.Error("FromMegavoltampereHours() with +Inf value should return error")
    }

    _, err = factory.FromMegavoltampereHours(math.Inf(-1))
    if err == nil {
        t.Error("FromMegavoltampereHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegavoltampereHours(0)
    if err != nil {
        t.Errorf("FromMegavoltampereHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ElectricApparentEnergyMegavoltampereHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegavoltampereHours() with zero value = %v, want 0", converted)
    }
}

func TestElectricApparentEnergyToString(t *testing.T) {
	factory := units.ElectricApparentEnergyFactory{}
	a, err := factory.CreateElectricApparentEnergy(45, units.ElectricApparentEnergyVoltampereHour)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ElectricApparentEnergyVoltampereHour, 2)
	expected := "45.00 " + units.GetElectricApparentEnergyAbbreviation(units.ElectricApparentEnergyVoltampereHour)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ElectricApparentEnergyVoltampereHour, -1)
	expected = "45 " + units.GetElectricApparentEnergyAbbreviation(units.ElectricApparentEnergyVoltampereHour)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestElectricApparentEnergy_EqualityAndComparison(t *testing.T) {
	factory := units.ElectricApparentEnergyFactory{}
	a1, _ := factory.CreateElectricApparentEnergy(60, units.ElectricApparentEnergyVoltampereHour)
	a2, _ := factory.CreateElectricApparentEnergy(60, units.ElectricApparentEnergyVoltampereHour)
	a3, _ := factory.CreateElectricApparentEnergy(120, units.ElectricApparentEnergyVoltampereHour)

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

func TestElectricApparentEnergy_Arithmetic(t *testing.T) {
	factory := units.ElectricApparentEnergyFactory{}
	a1, _ := factory.CreateElectricApparentEnergy(30, units.ElectricApparentEnergyVoltampereHour)
	a2, _ := factory.CreateElectricApparentEnergy(45, units.ElectricApparentEnergyVoltampereHour)

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