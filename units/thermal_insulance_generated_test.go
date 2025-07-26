// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestThermalInsulanceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "SquareMeterKelvinPerKilowatt"}`
	
	factory := units.ThermalInsulanceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.ThermalInsulanceSquareMeterKelvinPerKilowatt {
		t.Errorf("expected unit %v, got %v", units.ThermalInsulanceSquareMeterKelvinPerKilowatt, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "SquareMeterKelvinPerKilowatt"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestThermalInsulanceDto_ToJSON(t *testing.T) {
	dto := units.ThermalInsulanceDto{
		Value: 45,
		Unit:  units.ThermalInsulanceSquareMeterKelvinPerKilowatt,
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
	if result["unit"].(string) != string(units.ThermalInsulanceSquareMeterKelvinPerKilowatt) {
		t.Errorf("expected unit %s, got %v", units.ThermalInsulanceSquareMeterKelvinPerKilowatt, result["unit"])
	}
}

func TestNewThermalInsulance_InvalidValue(t *testing.T) {
	factory := units.ThermalInsulanceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateThermalInsulance(math.NaN(), units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateThermalInsulance(math.Inf(1), units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestThermalInsulanceConversions(t *testing.T) {
	factory := units.ThermalInsulanceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateThermalInsulance(180, units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to SquareMeterKelvinsPerKilowatt.
		// No expected conversion value provided for SquareMeterKelvinsPerKilowatt, verifying result is not NaN.
		result := a.SquareMeterKelvinsPerKilowatt()
		cacheResult := a.SquareMeterKelvinsPerKilowatt()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareMeterKelvinsPerKilowatt returned NaN")
		}
	}
	{
		// Test conversion to SquareMeterKelvinsPerWatt.
		// No expected conversion value provided for SquareMeterKelvinsPerWatt, verifying result is not NaN.
		result := a.SquareMeterKelvinsPerWatt()
		cacheResult := a.SquareMeterKelvinsPerWatt()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareMeterKelvinsPerWatt returned NaN")
		}
	}
	{
		// Test conversion to SquareMeterDegreesCelsiusPerWatt.
		// No expected conversion value provided for SquareMeterDegreesCelsiusPerWatt, verifying result is not NaN.
		result := a.SquareMeterDegreesCelsiusPerWatt()
		cacheResult := a.SquareMeterDegreesCelsiusPerWatt()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareMeterDegreesCelsiusPerWatt returned NaN")
		}
	}
	{
		// Test conversion to SquareCentimeterKelvinsPerWatt.
		// No expected conversion value provided for SquareCentimeterKelvinsPerWatt, verifying result is not NaN.
		result := a.SquareCentimeterKelvinsPerWatt()
		cacheResult := a.SquareCentimeterKelvinsPerWatt()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareCentimeterKelvinsPerWatt returned NaN")
		}
	}
	{
		// Test conversion to SquareMillimeterKelvinsPerWatt.
		// No expected conversion value provided for SquareMillimeterKelvinsPerWatt, verifying result is not NaN.
		result := a.SquareMillimeterKelvinsPerWatt()
		cacheResult := a.SquareMillimeterKelvinsPerWatt()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareMillimeterKelvinsPerWatt returned NaN")
		}
	}
	{
		// Test conversion to SquareCentimeterHourDegreesCelsiusPerKilocalorie.
		// No expected conversion value provided for SquareCentimeterHourDegreesCelsiusPerKilocalorie, verifying result is not NaN.
		result := a.SquareCentimeterHourDegreesCelsiusPerKilocalorie()
		cacheResult := a.SquareCentimeterHourDegreesCelsiusPerKilocalorie()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SquareCentimeterHourDegreesCelsiusPerKilocalorie returned NaN")
		}
	}
	{
		// Test conversion to HourSquareFeetDegreesFahrenheitPerBtu.
		// No expected conversion value provided for HourSquareFeetDegreesFahrenheitPerBtu, verifying result is not NaN.
		result := a.HourSquareFeetDegreesFahrenheitPerBtu()
		cacheResult := a.HourSquareFeetDegreesFahrenheitPerBtu()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to HourSquareFeetDegreesFahrenheitPerBtu returned NaN")
		}
	}
}

func TestThermalInsulance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.ThermalInsulanceFactory{}
	a, err := factory.CreateThermalInsulance(90, units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.ThermalInsulanceSquareMeterKelvinPerKilowatt {
		t.Errorf("expected default unit SquareMeterKelvinPerKilowatt, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.ThermalInsulanceSquareMeterKelvinPerKilowatt
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.ThermalInsulanceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.ThermalInsulanceSquareMeterKelvinPerKilowatt {
		t.Errorf("expected unit SquareMeterKelvinPerKilowatt, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestThermalInsulanceFactory_FromDto(t *testing.T) {
    factory := units.ThermalInsulanceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.ThermalInsulanceDto{
        Value: 100,
        Unit:  units.ThermalInsulanceSquareMeterKelvinPerKilowatt,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.ThermalInsulanceDto{
        Value: math.NaN(),
        Unit:  units.ThermalInsulanceSquareMeterKelvinPerKilowatt,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test SquareMeterKelvinPerKilowatt conversion
    square_meter_kelvins_per_kilowattDto := units.ThermalInsulanceDto{
        Value: 100,
        Unit:  units.ThermalInsulanceSquareMeterKelvinPerKilowatt,
    }
    
    var square_meter_kelvins_per_kilowattResult *units.ThermalInsulance
    square_meter_kelvins_per_kilowattResult, err = factory.FromDto(square_meter_kelvins_per_kilowattDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMeterKelvinPerKilowatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_meter_kelvins_per_kilowattResult.Convert(units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeterKelvinPerKilowatt = %v, want %v", converted, 100)
    }
    // Test SquareMeterKelvinPerWatt conversion
    square_meter_kelvins_per_wattDto := units.ThermalInsulanceDto{
        Value: 100,
        Unit:  units.ThermalInsulanceSquareMeterKelvinPerWatt,
    }
    
    var square_meter_kelvins_per_wattResult *units.ThermalInsulance
    square_meter_kelvins_per_wattResult, err = factory.FromDto(square_meter_kelvins_per_wattDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMeterKelvinPerWatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_meter_kelvins_per_wattResult.Convert(units.ThermalInsulanceSquareMeterKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeterKelvinPerWatt = %v, want %v", converted, 100)
    }
    // Test SquareMeterDegreeCelsiusPerWatt conversion
    square_meter_degrees_celsius_per_wattDto := units.ThermalInsulanceDto{
        Value: 100,
        Unit:  units.ThermalInsulanceSquareMeterDegreeCelsiusPerWatt,
    }
    
    var square_meter_degrees_celsius_per_wattResult *units.ThermalInsulance
    square_meter_degrees_celsius_per_wattResult, err = factory.FromDto(square_meter_degrees_celsius_per_wattDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMeterDegreeCelsiusPerWatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_meter_degrees_celsius_per_wattResult.Convert(units.ThermalInsulanceSquareMeterDegreeCelsiusPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMeterDegreeCelsiusPerWatt = %v, want %v", converted, 100)
    }
    // Test SquareCentimeterKelvinPerWatt conversion
    square_centimeter_kelvins_per_wattDto := units.ThermalInsulanceDto{
        Value: 100,
        Unit:  units.ThermalInsulanceSquareCentimeterKelvinPerWatt,
    }
    
    var square_centimeter_kelvins_per_wattResult *units.ThermalInsulance
    square_centimeter_kelvins_per_wattResult, err = factory.FromDto(square_centimeter_kelvins_per_wattDto)
    if err != nil {
        t.Errorf("FromDto() with SquareCentimeterKelvinPerWatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_centimeter_kelvins_per_wattResult.Convert(units.ThermalInsulanceSquareCentimeterKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareCentimeterKelvinPerWatt = %v, want %v", converted, 100)
    }
    // Test SquareMillimeterKelvinPerWatt conversion
    square_millimeter_kelvins_per_wattDto := units.ThermalInsulanceDto{
        Value: 100,
        Unit:  units.ThermalInsulanceSquareMillimeterKelvinPerWatt,
    }
    
    var square_millimeter_kelvins_per_wattResult *units.ThermalInsulance
    square_millimeter_kelvins_per_wattResult, err = factory.FromDto(square_millimeter_kelvins_per_wattDto)
    if err != nil {
        t.Errorf("FromDto() with SquareMillimeterKelvinPerWatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_millimeter_kelvins_per_wattResult.Convert(units.ThermalInsulanceSquareMillimeterKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMillimeterKelvinPerWatt = %v, want %v", converted, 100)
    }
    // Test SquareCentimeterHourDegreeCelsiusPerKilocalorie conversion
    square_centimeter_hour_degrees_celsius_per_kilocalorieDto := units.ThermalInsulanceDto{
        Value: 100,
        Unit:  units.ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie,
    }
    
    var square_centimeter_hour_degrees_celsius_per_kilocalorieResult *units.ThermalInsulance
    square_centimeter_hour_degrees_celsius_per_kilocalorieResult, err = factory.FromDto(square_centimeter_hour_degrees_celsius_per_kilocalorieDto)
    if err != nil {
        t.Errorf("FromDto() with SquareCentimeterHourDegreeCelsiusPerKilocalorie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_centimeter_hour_degrees_celsius_per_kilocalorieResult.Convert(units.ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareCentimeterHourDegreeCelsiusPerKilocalorie = %v, want %v", converted, 100)
    }
    // Test HourSquareFeetDegreeFahrenheitPerBtu conversion
    hour_square_feet_degrees_fahrenheit_per_btuDto := units.ThermalInsulanceDto{
        Value: 100,
        Unit:  units.ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu,
    }
    
    var hour_square_feet_degrees_fahrenheit_per_btuResult *units.ThermalInsulance
    hour_square_feet_degrees_fahrenheit_per_btuResult, err = factory.FromDto(hour_square_feet_degrees_fahrenheit_per_btuDto)
    if err != nil {
        t.Errorf("FromDto() with HourSquareFeetDegreeFahrenheitPerBtu returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hour_square_feet_degrees_fahrenheit_per_btuResult.Convert(units.ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HourSquareFeetDegreeFahrenheitPerBtu = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.ThermalInsulanceDto{
        Value: 0,
        Unit:  units.ThermalInsulanceSquareMeterKelvinPerKilowatt,
    }
    
    var zeroResult *units.ThermalInsulance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestThermalInsulanceFactory_FromDtoJSON(t *testing.T) {
    factory := units.ThermalInsulanceFactory{}
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
    nanJSON, _ := json.Marshal(units.ThermalInsulanceDto{
        Value: nanValue,
        Unit:  units.ThermalInsulanceSquareMeterKelvinPerKilowatt,
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
    converted = square_meter_kelvins_per_kilowattResult.Convert(units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
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
    converted = square_meter_kelvins_per_wattResult.Convert(units.ThermalInsulanceSquareMeterKelvinPerWatt)
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
    converted = square_meter_degrees_celsius_per_wattResult.Convert(units.ThermalInsulanceSquareMeterDegreeCelsiusPerWatt)
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
    converted = square_centimeter_kelvins_per_wattResult.Convert(units.ThermalInsulanceSquareCentimeterKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareCentimeterKelvinPerWatt = %v, want %v", converted, 100)
    }
    // Test JSON with SquareMillimeterKelvinPerWatt unit
    square_millimeter_kelvins_per_wattJSON := []byte(`{"value": 100, "unit": "SquareMillimeterKelvinPerWatt"}`)
    square_millimeter_kelvins_per_wattResult, err := factory.FromDtoJSON(square_millimeter_kelvins_per_wattJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareMillimeterKelvinPerWatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_millimeter_kelvins_per_wattResult.Convert(units.ThermalInsulanceSquareMillimeterKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SquareMillimeterKelvinPerWatt = %v, want %v", converted, 100)
    }
    // Test JSON with SquareCentimeterHourDegreeCelsiusPerKilocalorie unit
    square_centimeter_hour_degrees_celsius_per_kilocalorieJSON := []byte(`{"value": 100, "unit": "SquareCentimeterHourDegreeCelsiusPerKilocalorie"}`)
    square_centimeter_hour_degrees_celsius_per_kilocalorieResult, err := factory.FromDtoJSON(square_centimeter_hour_degrees_celsius_per_kilocalorieJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SquareCentimeterHourDegreeCelsiusPerKilocalorie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = square_centimeter_hour_degrees_celsius_per_kilocalorieResult.Convert(units.ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
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
    converted = hour_square_feet_degrees_fahrenheit_per_btuResult.Convert(units.ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu)
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
func TestThermalInsulanceFactory_FromSquareMeterKelvinsPerKilowatt(t *testing.T) {
    factory := units.ThermalInsulanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMeterKelvinsPerKilowatt(100)
    if err != nil {
        t.Errorf("FromSquareMeterKelvinsPerKilowatt() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
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
    converted = zeroResult.Convert(units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMeterKelvinsPerKilowatt() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareMeterKelvinsPerWatt function
func TestThermalInsulanceFactory_FromSquareMeterKelvinsPerWatt(t *testing.T) {
    factory := units.ThermalInsulanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMeterKelvinsPerWatt(100)
    if err != nil {
        t.Errorf("FromSquareMeterKelvinsPerWatt() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalInsulanceSquareMeterKelvinPerWatt)
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
    converted = zeroResult.Convert(units.ThermalInsulanceSquareMeterKelvinPerWatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMeterKelvinsPerWatt() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareMeterDegreesCelsiusPerWatt function
func TestThermalInsulanceFactory_FromSquareMeterDegreesCelsiusPerWatt(t *testing.T) {
    factory := units.ThermalInsulanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMeterDegreesCelsiusPerWatt(100)
    if err != nil {
        t.Errorf("FromSquareMeterDegreesCelsiusPerWatt() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalInsulanceSquareMeterDegreeCelsiusPerWatt)
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
    converted = zeroResult.Convert(units.ThermalInsulanceSquareMeterDegreeCelsiusPerWatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMeterDegreesCelsiusPerWatt() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareCentimeterKelvinsPerWatt function
func TestThermalInsulanceFactory_FromSquareCentimeterKelvinsPerWatt(t *testing.T) {
    factory := units.ThermalInsulanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareCentimeterKelvinsPerWatt(100)
    if err != nil {
        t.Errorf("FromSquareCentimeterKelvinsPerWatt() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalInsulanceSquareCentimeterKelvinPerWatt)
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
    converted = zeroResult.Convert(units.ThermalInsulanceSquareCentimeterKelvinPerWatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareCentimeterKelvinsPerWatt() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareMillimeterKelvinsPerWatt function
func TestThermalInsulanceFactory_FromSquareMillimeterKelvinsPerWatt(t *testing.T) {
    factory := units.ThermalInsulanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareMillimeterKelvinsPerWatt(100)
    if err != nil {
        t.Errorf("FromSquareMillimeterKelvinsPerWatt() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalInsulanceSquareMillimeterKelvinPerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSquareMillimeterKelvinsPerWatt() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSquareMillimeterKelvinsPerWatt(math.NaN())
    if err == nil {
        t.Error("FromSquareMillimeterKelvinsPerWatt() with NaN value should return error")
    }

    _, err = factory.FromSquareMillimeterKelvinsPerWatt(math.Inf(1))
    if err == nil {
        t.Error("FromSquareMillimeterKelvinsPerWatt() with +Inf value should return error")
    }

    _, err = factory.FromSquareMillimeterKelvinsPerWatt(math.Inf(-1))
    if err == nil {
        t.Error("FromSquareMillimeterKelvinsPerWatt() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSquareMillimeterKelvinsPerWatt(0)
    if err != nil {
        t.Errorf("FromSquareMillimeterKelvinsPerWatt() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.ThermalInsulanceSquareMillimeterKelvinPerWatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareMillimeterKelvinsPerWatt() with zero value = %v, want 0", converted)
    }
}
// Test FromSquareCentimeterHourDegreesCelsiusPerKilocalorie function
func TestThermalInsulanceFactory_FromSquareCentimeterHourDegreesCelsiusPerKilocalorie(t *testing.T) {
    factory := units.ThermalInsulanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSquareCentimeterHourDegreesCelsiusPerKilocalorie(100)
    if err != nil {
        t.Errorf("FromSquareCentimeterHourDegreesCelsiusPerKilocalorie() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
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
    converted = zeroResult.Convert(units.ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSquareCentimeterHourDegreesCelsiusPerKilocalorie() with zero value = %v, want 0", converted)
    }
}
// Test FromHourSquareFeetDegreesFahrenheitPerBtu function
func TestThermalInsulanceFactory_FromHourSquareFeetDegreesFahrenheitPerBtu(t *testing.T) {
    factory := units.ThermalInsulanceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHourSquareFeetDegreesFahrenheitPerBtu(100)
    if err != nil {
        t.Errorf("FromHourSquareFeetDegreesFahrenheitPerBtu() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu)
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
    converted = zeroResult.Convert(units.ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHourSquareFeetDegreesFahrenheitPerBtu() with zero value = %v, want 0", converted)
    }
}

func TestThermalInsulanceToString(t *testing.T) {
	factory := units.ThermalInsulanceFactory{}
	a, err := factory.CreateThermalInsulance(45, units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.ThermalInsulanceSquareMeterKelvinPerKilowatt, 2)
	expected := "45.00 " + units.GetThermalInsulanceAbbreviation(units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.ThermalInsulanceSquareMeterKelvinPerKilowatt, -1)
	expected = "45 " + units.GetThermalInsulanceAbbreviation(units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestThermalInsulance_EqualityAndComparison(t *testing.T) {
	factory := units.ThermalInsulanceFactory{}
	a1, _ := factory.CreateThermalInsulance(60, units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
	a2, _ := factory.CreateThermalInsulance(60, units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
	a3, _ := factory.CreateThermalInsulance(120, units.ThermalInsulanceSquareMeterKelvinPerKilowatt)

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

func TestThermalInsulance_Arithmetic(t *testing.T) {
	factory := units.ThermalInsulanceFactory{}
	a1, _ := factory.CreateThermalInsulance(30, units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
	a2, _ := factory.CreateThermalInsulance(45, units.ThermalInsulanceSquareMeterKelvinPerKilowatt)

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


func TestGetThermalInsulanceAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.ThermalInsulanceUnits
        want string
    }{
        {
            name: "SquareMeterKelvinPerKilowatt abbreviation",
            unit: units.ThermalInsulanceSquareMeterKelvinPerKilowatt,
            want: "m²K/kW",
        },
        {
            name: "SquareMeterKelvinPerWatt abbreviation",
            unit: units.ThermalInsulanceSquareMeterKelvinPerWatt,
            want: "m²K/W",
        },
        {
            name: "SquareMeterDegreeCelsiusPerWatt abbreviation",
            unit: units.ThermalInsulanceSquareMeterDegreeCelsiusPerWatt,
            want: "m²°C/W",
        },
        {
            name: "SquareCentimeterKelvinPerWatt abbreviation",
            unit: units.ThermalInsulanceSquareCentimeterKelvinPerWatt,
            want: "cm²K/W",
        },
        {
            name: "SquareMillimeterKelvinPerWatt abbreviation",
            unit: units.ThermalInsulanceSquareMillimeterKelvinPerWatt,
            want: "mm²K/W",
        },
        {
            name: "SquareCentimeterHourDegreeCelsiusPerKilocalorie abbreviation",
            unit: units.ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie,
            want: "cm²Hr°C/kcal",
        },
        {
            name: "HourSquareFeetDegreeFahrenheitPerBtu abbreviation",
            unit: units.ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu,
            want: "Hrft²°F/Btu",
        },
        {
            name: "invalid unit",
            unit: units.ThermalInsulanceUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetThermalInsulanceAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetThermalInsulanceAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestThermalInsulance_String(t *testing.T) {
    factory := units.ThermalInsulanceFactory{}
    
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
            unit, err := factory.CreateThermalInsulance(tt.value, units.ThermalInsulanceSquareMeterKelvinPerKilowatt)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("ThermalInsulance.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestThermalInsulance_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.ThermalInsulanceFactory{}

	_, err := uf.CreateThermalInsulance(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}