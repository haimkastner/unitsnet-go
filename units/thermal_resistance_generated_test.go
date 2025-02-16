// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestThermalResistanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SquareMeterKelvinPerKilowatt"}`
	
	factory := units.ThermalResistanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ThermalResistanceSquareMeterKelvinPerKilowatt {
		t.Errorf("expected unit %v, got %v", units.ThermalResistanceSquareMeterKelvinPerKilowatt, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SquareMeterKelvinPerKilowatt"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestThermalResistanceDto_ToJSON(t *testing.T) {
	dto := units.ThermalResistanceDto{
		Value: 45,
		Unit:  units.ThermalResistanceSquareMeterKelvinPerKilowatt,
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
	if result["unit"].(string) != string(units.ThermalResistanceSquareMeterKelvinPerKilowatt) {
		t.Errorf("expected unit %s, got %v", units.ThermalResistanceSquareMeterKelvinPerKilowatt, result["unit"])
	}
}

func TestNewThermalResistance_InvalidValue(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateThermalResistance(math.NaN(), units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateThermalResistance(math.Inf(1), units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestThermalResistanceConversions(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateThermalResistance(180, units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to SquareMeterKelvinsPerKilowatt.
		// No expected conversion value provided for SquareMeterKelvinsPerKilowatt, verifying result is not NaN.
		result := a.SquareMeterKelvinsPerKilowatt()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareMeterKelvinsPerKilowatt returned NaN")
		}
	}
	{
		// Test conversion to SquareMeterKelvinsPerWatt.
		// No expected conversion value provided for SquareMeterKelvinsPerWatt, verifying result is not NaN.
		result := a.SquareMeterKelvinsPerWatt()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareMeterKelvinsPerWatt returned NaN")
		}
	}
	{
		// Test conversion to SquareMeterDegreesCelsiusPerWatt.
		// No expected conversion value provided for SquareMeterDegreesCelsiusPerWatt, verifying result is not NaN.
		result := a.SquareMeterDegreesCelsiusPerWatt()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareMeterDegreesCelsiusPerWatt returned NaN")
		}
	}
	{
		// Test conversion to SquareCentimeterKelvinsPerWatt.
		// No expected conversion value provided for SquareCentimeterKelvinsPerWatt, verifying result is not NaN.
		result := a.SquareCentimeterKelvinsPerWatt()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareCentimeterKelvinsPerWatt returned NaN")
		}
	}
	{
		// Test conversion to SquareCentimeterHourDegreesCelsiusPerKilocalorie.
		// No expected conversion value provided for SquareCentimeterHourDegreesCelsiusPerKilocalorie, verifying result is not NaN.
		result := a.SquareCentimeterHourDegreesCelsiusPerKilocalorie()
		if math.IsNaN(result) {
			t.Errorf("conversion to SquareCentimeterHourDegreesCelsiusPerKilocalorie returned NaN")
		}
	}
	{
		// Test conversion to HourSquareFeetDegreesFahrenheitPerBtu.
		// No expected conversion value provided for HourSquareFeetDegreesFahrenheitPerBtu, verifying result is not NaN.
		result := a.HourSquareFeetDegreesFahrenheitPerBtu()
		if math.IsNaN(result) {
			t.Errorf("conversion to HourSquareFeetDegreesFahrenheitPerBtu returned NaN")
		}
	}
}

func TestThermalResistance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	a, err := factory.CreateThermalResistance(90, units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ThermalResistanceSquareMeterKelvinPerKilowatt {
		t.Errorf("expected default unit SquareMeterKelvinPerKilowatt, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ThermalResistanceSquareMeterKelvinPerKilowatt
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ThermalResistanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ThermalResistanceSquareMeterKelvinPerKilowatt {
		t.Errorf("expected unit SquareMeterKelvinPerKilowatt, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestThermalResistanceFactory_FromDto(t *testing.T) {
    factory := units.ThermalResistanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ThermalResistanceDto{
        Value: 100,
        Unit:  units.ThermalResistanceSquareMeterKelvinPerKilowatt,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ThermalResistanceDto{
        Value: math.NaN(),
        Unit:  units.ThermalResistanceSquareMeterKelvinPerKilowatt,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test SquareMeterKelvinPerKilowatt conversion
    square_meter_kelvins_per_kilowattDto := units.ThermalResistanceDto{
        Value: 100,
        Unit:  units.ThermalResistanceSquareMeterKelvinPerKilowatt,
    }
    
    var square_meter_kelvins_per_kilowattResult *units.ThermalResistance
    square_meter_kelvins_per_kilowattResult, err = factory.FromDto(square_meter_kelvins_per_kilowattDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMeterKelvinPerKilowatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_meter_kelvins_per_kilowattResult.Convert(units.ThermalResistanceSquareMeterKelvinPerKilowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeterKelvinPerKilowatt = %v, want %v", converted, 100)
    }
    // Test SquareMeterKelvinPerWatt conversion
    square_meter_kelvins_per_wattDto := units.ThermalResistanceDto{
        Value: 100,
        Unit:  units.ThermalResistanceSquareMeterKelvinPerWatt,
    }
    
    var square_meter_kelvins_per_wattResult *units.ThermalResistance
    square_meter_kelvins_per_wattResult, err = factory.FromDto(square_meter_kelvins_per_wattDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMeterKelvinPerWatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_meter_kelvins_per_wattResult.Convert(units.ThermalResistanceSquareMeterKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeterKelvinPerWatt = %v, want %v", converted, 100)
    }
    // Test SquareMeterDegreeCelsiusPerWatt conversion
    square_meter_degrees_celsius_per_wattDto := units.ThermalResistanceDto{
        Value: 100,
        Unit:  units.ThermalResistanceSquareMeterDegreeCelsiusPerWatt,
    }
    
    var square_meter_degrees_celsius_per_wattResult *units.ThermalResistance
    square_meter_degrees_celsius_per_wattResult, err = factory.FromDto(square_meter_degrees_celsius_per_wattDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMeterDegreeCelsiusPerWatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_meter_degrees_celsius_per_wattResult.Convert(units.ThermalResistanceSquareMeterDegreeCelsiusPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeterDegreeCelsiusPerWatt = %v, want %v", converted, 100)
    }
    // Test SquareCentimeterKelvinPerWatt conversion
    square_centimeter_kelvins_per_wattDto := units.ThermalResistanceDto{
        Value: 100,
        Unit:  units.ThermalResistanceSquareCentimeterKelvinPerWatt,
    }
    
    var square_centimeter_kelvins_per_wattResult *units.ThermalResistance
    square_centimeter_kelvins_per_wattResult, err = factory.FromDto(square_centimeter_kelvins_per_wattDto)
    if err != nil {
        t.Errorf("FromDto() with SquareCentimeterKelvinPerWatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_centimeter_kelvins_per_wattResult.Convert(units.ThermalResistanceSquareCentimeterKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareCentimeterKelvinPerWatt = %v, want %v", converted, 100)
    }
    // Test SquareCentimeterHourDegreeCelsiusPerKilocalorie conversion
    square_centimeter_hour_degrees_celsius_per_kilocalorieDto := units.ThermalResistanceDto{
        Value: 100,
        Unit:  units.ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie,
    }
    
    var square_centimeter_hour_degrees_celsius_per_kilocalorieResult *units.ThermalResistance
    square_centimeter_hour_degrees_celsius_per_kilocalorieResult, err = factory.FromDto(square_centimeter_hour_degrees_celsius_per_kilocalorieDto)
    if err != nil {
        t.Errorf("FromDto() with SquareCentimeterHourDegreeCelsiusPerKilocalorie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_centimeter_hour_degrees_celsius_per_kilocalorieResult.Convert(units.ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareCentimeterHourDegreeCelsiusPerKilocalorie = %v, want %v", converted, 100)
    }
    // Test HourSquareFeetDegreeFahrenheitPerBtu conversion
    hour_square_feet_degrees_fahrenheit_per_btuDto := units.ThermalResistanceDto{
        Value: 100,
        Unit:  units.ThermalResistanceHourSquareFeetDegreeFahrenheitPerBtu,
    }
    
    var hour_square_feet_degrees_fahrenheit_per_btuResult *units.ThermalResistance
    hour_square_feet_degrees_fahrenheit_per_btuResult, err = factory.FromDto(hour_square_feet_degrees_fahrenheit_per_btuDto)
    if err != nil {
        t.Errorf("FromDto() with HourSquareFeetDegreeFahrenheitPerBtu returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hour_square_feet_degrees_fahrenheit_per_btuResult.Convert(units.ThermalResistanceHourSquareFeetDegreeFahrenheitPerBtu)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HourSquareFeetDegreeFahrenheitPerBtu = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ThermalResistanceDto{
        Value: 0,
        Unit:  units.ThermalResistanceSquareMeterKelvinPerKilowatt,
    }
    
    var zeroResult *units.ThermalResistance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestThermalResistanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.ThermalResistanceFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "SquareMeterKelvinPerKilowatt"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "SquareMeterKelvinPerKilowatt"}`)
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
    nanJSON, _ := json.Marshal(units.ThermalResistanceDto{
        Value: nanValue,
        Unit:  units.ThermalResistanceSquareMeterKelvinPerKilowatt,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with SquareMeterKelvinPerKilowatt unit
    square_meter_kelvins_per_kilowattJSON := []byte(`{"value": 100, "unit": "SquareMeterKelvinPerKilowatt"}`)
    square_meter_kelvins_per_kilowattResult, err := factory.FromDtoJSON(square_meter_kelvins_per_kilowattJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareMeterKelvinPerKilowatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_meter_kelvins_per_kilowattResult.Convert(units.ThermalResistanceSquareMeterKelvinPerKilowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeterKelvinPerKilowatt = %v, want %v", converted, 100)
    }
    // Test JSON with SquareMeterKelvinPerWatt unit
    square_meter_kelvins_per_wattJSON := []byte(`{"value": 100, "unit": "SquareMeterKelvinPerWatt"}`)
    square_meter_kelvins_per_wattResult, err := factory.FromDtoJSON(square_meter_kelvins_per_wattJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareMeterKelvinPerWatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_meter_kelvins_per_wattResult.Convert(units.ThermalResistanceSquareMeterKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeterKelvinPerWatt = %v, want %v", converted, 100)
    }
    // Test JSON with SquareMeterDegreeCelsiusPerWatt unit
    square_meter_degrees_celsius_per_wattJSON := []byte(`{"value": 100, "unit": "SquareMeterDegreeCelsiusPerWatt"}`)
    square_meter_degrees_celsius_per_wattResult, err := factory.FromDtoJSON(square_meter_degrees_celsius_per_wattJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareMeterDegreeCelsiusPerWatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_meter_degrees_celsius_per_wattResult.Convert(units.ThermalResistanceSquareMeterDegreeCelsiusPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeterDegreeCelsiusPerWatt = %v, want %v", converted, 100)
    }
    // Test JSON with SquareCentimeterKelvinPerWatt unit
    square_centimeter_kelvins_per_wattJSON := []byte(`{"value": 100, "unit": "SquareCentimeterKelvinPerWatt"}`)
    square_centimeter_kelvins_per_wattResult, err := factory.FromDtoJSON(square_centimeter_kelvins_per_wattJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareCentimeterKelvinPerWatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_centimeter_kelvins_per_wattResult.Convert(units.ThermalResistanceSquareCentimeterKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareCentimeterKelvinPerWatt = %v, want %v", converted, 100)
    }
    // Test JSON with SquareCentimeterHourDegreeCelsiusPerKilocalorie unit
    square_centimeter_hour_degrees_celsius_per_kilocalorieJSON := []byte(`{"value": 100, "unit": "SquareCentimeterHourDegreeCelsiusPerKilocalorie"}`)
    square_centimeter_hour_degrees_celsius_per_kilocalorieResult, err := factory.FromDtoJSON(square_centimeter_hour_degrees_celsius_per_kilocalorieJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareCentimeterHourDegreeCelsiusPerKilocalorie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_centimeter_hour_degrees_celsius_per_kilocalorieResult.Convert(units.ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareCentimeterHourDegreeCelsiusPerKilocalorie = %v, want %v", converted, 100)
    }
    // Test JSON with HourSquareFeetDegreeFahrenheitPerBtu unit
    hour_square_feet_degrees_fahrenheit_per_btuJSON := []byte(`{"value": 100, "unit": "HourSquareFeetDegreeFahrenheitPerBtu"}`)
    hour_square_feet_degrees_fahrenheit_per_btuResult, err := factory.FromDtoJSON(hour_square_feet_degrees_fahrenheit_per_btuJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HourSquareFeetDegreeFahrenheitPerBtu unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hour_square_feet_degrees_fahrenheit_per_btuResult.Convert(units.ThermalResistanceHourSquareFeetDegreeFahrenheitPerBtu)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HourSquareFeetDegreeFahrenheitPerBtu = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "SquareMeterKelvinPerKilowatt"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromSquareMeterKelvinsPerKilowatt function
func TestThermalResistanceFactory_FromSquareMeterKelvinsPerKilowatt(t *testing.T) {
    factory := units.ThermalResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMeterKelvinsPerKilowatt(100)
    if err != nil {
        t.Errorf("FromSquareMeterKelvinsPerKilowatt() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalResistanceSquareMeterKelvinPerKilowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareMeterKelvinsPerKilowatt() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareMeterKelvinsPerKilowatt(math.NaN())
    if err == nil {
        t.Error("FromSquareMeterKelvinsPerKilowatt() with NaN value should return error")
    }

    _, err = factory.FromSquareMeterKelvinsPerKilowatt(math.Inf(1))
    if err == nil {
        t.Error("FromSquareMeterKelvinsPerKilowatt() with +Inf value should return error")
    }

    _, err = factory.FromSquareMeterKelvinsPerKilowatt(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareMeterKelvinsPerKilowatt() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareMeterKelvinsPerKilowatt(0)
    if err != nil {
        t.Errorf("FromSquareMeterKelvinsPerKilowatt() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ThermalResistanceSquareMeterKelvinPerKilowatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMeterKelvinsPerKilowatt() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareMeterKelvinsPerWatt function
func TestThermalResistanceFactory_FromSquareMeterKelvinsPerWatt(t *testing.T) {
    factory := units.ThermalResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMeterKelvinsPerWatt(100)
    if err != nil {
        t.Errorf("FromSquareMeterKelvinsPerWatt() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalResistanceSquareMeterKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareMeterKelvinsPerWatt() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareMeterKelvinsPerWatt(math.NaN())
    if err == nil {
        t.Error("FromSquareMeterKelvinsPerWatt() with NaN value should return error")
    }

    _, err = factory.FromSquareMeterKelvinsPerWatt(math.Inf(1))
    if err == nil {
        t.Error("FromSquareMeterKelvinsPerWatt() with +Inf value should return error")
    }

    _, err = factory.FromSquareMeterKelvinsPerWatt(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareMeterKelvinsPerWatt() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareMeterKelvinsPerWatt(0)
    if err != nil {
        t.Errorf("FromSquareMeterKelvinsPerWatt() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ThermalResistanceSquareMeterKelvinPerWatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMeterKelvinsPerWatt() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareMeterDegreesCelsiusPerWatt function
func TestThermalResistanceFactory_FromSquareMeterDegreesCelsiusPerWatt(t *testing.T) {
    factory := units.ThermalResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMeterDegreesCelsiusPerWatt(100)
    if err != nil {
        t.Errorf("FromSquareMeterDegreesCelsiusPerWatt() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalResistanceSquareMeterDegreeCelsiusPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareMeterDegreesCelsiusPerWatt() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareMeterDegreesCelsiusPerWatt(math.NaN())
    if err == nil {
        t.Error("FromSquareMeterDegreesCelsiusPerWatt() with NaN value should return error")
    }

    _, err = factory.FromSquareMeterDegreesCelsiusPerWatt(math.Inf(1))
    if err == nil {
        t.Error("FromSquareMeterDegreesCelsiusPerWatt() with +Inf value should return error")
    }

    _, err = factory.FromSquareMeterDegreesCelsiusPerWatt(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareMeterDegreesCelsiusPerWatt() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareMeterDegreesCelsiusPerWatt(0)
    if err != nil {
        t.Errorf("FromSquareMeterDegreesCelsiusPerWatt() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ThermalResistanceSquareMeterDegreeCelsiusPerWatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMeterDegreesCelsiusPerWatt() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareCentimeterKelvinsPerWatt function
func TestThermalResistanceFactory_FromSquareCentimeterKelvinsPerWatt(t *testing.T) {
    factory := units.ThermalResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareCentimeterKelvinsPerWatt(100)
    if err != nil {
        t.Errorf("FromSquareCentimeterKelvinsPerWatt() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalResistanceSquareCentimeterKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareCentimeterKelvinsPerWatt() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareCentimeterKelvinsPerWatt(math.NaN())
    if err == nil {
        t.Error("FromSquareCentimeterKelvinsPerWatt() with NaN value should return error")
    }

    _, err = factory.FromSquareCentimeterKelvinsPerWatt(math.Inf(1))
    if err == nil {
        t.Error("FromSquareCentimeterKelvinsPerWatt() with +Inf value should return error")
    }

    _, err = factory.FromSquareCentimeterKelvinsPerWatt(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareCentimeterKelvinsPerWatt() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareCentimeterKelvinsPerWatt(0)
    if err != nil {
        t.Errorf("FromSquareCentimeterKelvinsPerWatt() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ThermalResistanceSquareCentimeterKelvinPerWatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareCentimeterKelvinsPerWatt() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareCentimeterHourDegreesCelsiusPerKilocalorie function
func TestThermalResistanceFactory_FromSquareCentimeterHourDegreesCelsiusPerKilocalorie(t *testing.T) {
    factory := units.ThermalResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareCentimeterHourDegreesCelsiusPerKilocalorie(100)
    if err != nil {
        t.Errorf("FromSquareCentimeterHourDegreesCelsiusPerKilocalorie() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareCentimeterHourDegreesCelsiusPerKilocalorie() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareCentimeterHourDegreesCelsiusPerKilocalorie(math.NaN())
    if err == nil {
        t.Error("FromSquareCentimeterHourDegreesCelsiusPerKilocalorie() with NaN value should return error")
    }

    _, err = factory.FromSquareCentimeterHourDegreesCelsiusPerKilocalorie(math.Inf(1))
    if err == nil {
        t.Error("FromSquareCentimeterHourDegreesCelsiusPerKilocalorie() with +Inf value should return error")
    }

    _, err = factory.FromSquareCentimeterHourDegreesCelsiusPerKilocalorie(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareCentimeterHourDegreesCelsiusPerKilocalorie() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareCentimeterHourDegreesCelsiusPerKilocalorie(0)
    if err != nil {
        t.Errorf("FromSquareCentimeterHourDegreesCelsiusPerKilocalorie() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareCentimeterHourDegreesCelsiusPerKilocalorie() with zero value = %v, want 0", converted)
    }
}
// Test FromHourSquareFeetDegreesFahrenheitPerBtu function
func TestThermalResistanceFactory_FromHourSquareFeetDegreesFahrenheitPerBtu(t *testing.T) {
    factory := units.ThermalResistanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHourSquareFeetDegreesFahrenheitPerBtu(100)
    if err != nil {
        t.Errorf("FromHourSquareFeetDegreesFahrenheitPerBtu() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalResistanceHourSquareFeetDegreeFahrenheitPerBtu)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHourSquareFeetDegreesFahrenheitPerBtu() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHourSquareFeetDegreesFahrenheitPerBtu(math.NaN())
    if err == nil {
        t.Error("FromHourSquareFeetDegreesFahrenheitPerBtu() with NaN value should return error")
    }

    _, err = factory.FromHourSquareFeetDegreesFahrenheitPerBtu(math.Inf(1))
    if err == nil {
        t.Error("FromHourSquareFeetDegreesFahrenheitPerBtu() with +Inf value should return error")
    }

    _, err = factory.FromHourSquareFeetDegreesFahrenheitPerBtu(math.Inf(-1))
    if err == nil {
        t.Error("FromHourSquareFeetDegreesFahrenheitPerBtu() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHourSquareFeetDegreesFahrenheitPerBtu(0)
    if err != nil {
        t.Errorf("FromHourSquareFeetDegreesFahrenheitPerBtu() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ThermalResistanceHourSquareFeetDegreeFahrenheitPerBtu)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHourSquareFeetDegreesFahrenheitPerBtu() with zero value = %v, want 0", converted)
    }
}

func TestThermalResistanceToString(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	a, err := factory.CreateThermalResistance(45, units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ThermalResistanceSquareMeterKelvinPerKilowatt, 2)
	expected := "45.00 " + units.GetThermalResistanceAbbreviation(units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ThermalResistanceSquareMeterKelvinPerKilowatt, -1)
	expected = "45 " + units.GetThermalResistanceAbbreviation(units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestThermalResistance_EqualityAndComparison(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	a1, _ := factory.CreateThermalResistance(60, units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	a2, _ := factory.CreateThermalResistance(60, units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	a3, _ := factory.CreateThermalResistance(120, units.ThermalResistanceSquareMeterKelvinPerKilowatt)

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

func TestThermalResistance_Arithmetic(t *testing.T) {
	factory := units.ThermalResistanceFactory{}
	a1, _ := factory.CreateThermalResistance(30, units.ThermalResistanceSquareMeterKelvinPerKilowatt)
	a2, _ := factory.CreateThermalResistance(45, units.ThermalResistanceSquareMeterKelvinPerKilowatt)

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