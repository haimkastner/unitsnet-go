// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestHeatTransferCoefficientDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "WattPerSquareMeterKelvin"}`
	
	factory := units.HeatTransferCoefficientDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.HeatTransferCoefficientWattPerSquareMeterKelvin {
		t.Errorf("expected unit %v, got %v", units.HeatTransferCoefficientWattPerSquareMeterKelvin, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "WattPerSquareMeterKelvin"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestHeatTransferCoefficientDto_ToJSON(t *testing.T) {
	dto := units.HeatTransferCoefficientDto{
		Value: 45,
		Unit:  units.HeatTransferCoefficientWattPerSquareMeterKelvin,
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
	if result["unit"].(string) != string(units.HeatTransferCoefficientWattPerSquareMeterKelvin) {
		t.Errorf("expected unit %s, got %v", units.HeatTransferCoefficientWattPerSquareMeterKelvin, result["unit"])
	}
}

func TestNewHeatTransferCoefficient_InvalidValue(t *testing.T) {
	factory := units.HeatTransferCoefficientFactory{}
	// NaN value should return an error.
	_, err := factory.CreateHeatTransferCoefficient(math.NaN(), units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateHeatTransferCoefficient(math.Inf(1), units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestHeatTransferCoefficientConversions(t *testing.T) {
	factory := units.HeatTransferCoefficientFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateHeatTransferCoefficient(180, units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to WattsPerSquareMeterKelvin.
		// No expected conversion value provided for WattsPerSquareMeterKelvin, verifying result is not NaN.
		result := a.WattsPerSquareMeterKelvin()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerSquareMeterKelvin returned NaN")
		}
	}
	{
		// Test conversion to WattsPerSquareMeterCelsius.
		// No expected conversion value provided for WattsPerSquareMeterCelsius, verifying result is not NaN.
		result := a.WattsPerSquareMeterCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerSquareMeterCelsius returned NaN")
		}
	}
	{
		// Test conversion to BtusPerHourSquareFootDegreeFahrenheit.
		// No expected conversion value provided for BtusPerHourSquareFootDegreeFahrenheit, verifying result is not NaN.
		result := a.BtusPerHourSquareFootDegreeFahrenheit()
		if math.IsNaN(result) {
			t.Errorf("conversion to BtusPerHourSquareFootDegreeFahrenheit returned NaN")
		}
	}
	{
		// Test conversion to CaloriesPerHourSquareMeterDegreeCelsius.
		// No expected conversion value provided for CaloriesPerHourSquareMeterDegreeCelsius, verifying result is not NaN.
		result := a.CaloriesPerHourSquareMeterDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to CaloriesPerHourSquareMeterDegreeCelsius returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerHourSquareMeterDegreeCelsius.
		// No expected conversion value provided for KilocaloriesPerHourSquareMeterDegreeCelsius, verifying result is not NaN.
		result := a.KilocaloriesPerHourSquareMeterDegreeCelsius()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilocaloriesPerHourSquareMeterDegreeCelsius returned NaN")
		}
	}
}

func TestHeatTransferCoefficient_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.HeatTransferCoefficientFactory{}
	a, err := factory.CreateHeatTransferCoefficient(90, units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.HeatTransferCoefficientWattPerSquareMeterKelvin {
		t.Errorf("expected default unit WattPerSquareMeterKelvin, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.HeatTransferCoefficientWattPerSquareMeterKelvin
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.HeatTransferCoefficientDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.HeatTransferCoefficientWattPerSquareMeterKelvin {
		t.Errorf("expected unit WattPerSquareMeterKelvin, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestHeatTransferCoefficientFactory_FromDto(t *testing.T) {
    factory := units.HeatTransferCoefficientFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.HeatTransferCoefficientDto{
        Value: 100,
        Unit:  units.HeatTransferCoefficientWattPerSquareMeterKelvin,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.HeatTransferCoefficientDto{
        Value: math.NaN(),
        Unit:  units.HeatTransferCoefficientWattPerSquareMeterKelvin,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test WattPerSquareMeterKelvin conversion
    watts_per_square_meter_kelvinDto := units.HeatTransferCoefficientDto{
        Value: 100,
        Unit:  units.HeatTransferCoefficientWattPerSquareMeterKelvin,
    }
    
    var watts_per_square_meter_kelvinResult *units.HeatTransferCoefficient
    watts_per_square_meter_kelvinResult, err = factory.FromDto(watts_per_square_meter_kelvinDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerSquareMeterKelvin returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_meter_kelvinResult.Convert(units.HeatTransferCoefficientWattPerSquareMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareMeterKelvin = %v, want %v", converted, 100)
    }
    // Test WattPerSquareMeterCelsius conversion
    watts_per_square_meter_celsiusDto := units.HeatTransferCoefficientDto{
        Value: 100,
        Unit:  units.HeatTransferCoefficientWattPerSquareMeterCelsius,
    }
    
    var watts_per_square_meter_celsiusResult *units.HeatTransferCoefficient
    watts_per_square_meter_celsiusResult, err = factory.FromDto(watts_per_square_meter_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerSquareMeterCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_meter_celsiusResult.Convert(units.HeatTransferCoefficientWattPerSquareMeterCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareMeterCelsius = %v, want %v", converted, 100)
    }
    // Test BtuPerHourSquareFootDegreeFahrenheit conversion
    btus_per_hour_square_foot_degree_fahrenheitDto := units.HeatTransferCoefficientDto{
        Value: 100,
        Unit:  units.HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit,
    }
    
    var btus_per_hour_square_foot_degree_fahrenheitResult *units.HeatTransferCoefficient
    btus_per_hour_square_foot_degree_fahrenheitResult, err = factory.FromDto(btus_per_hour_square_foot_degree_fahrenheitDto)
    if err != nil {
        t.Errorf("FromDto() with BtuPerHourSquareFootDegreeFahrenheit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_hour_square_foot_degree_fahrenheitResult.Convert(units.HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerHourSquareFootDegreeFahrenheit = %v, want %v", converted, 100)
    }
    // Test CaloriePerHourSquareMeterDegreeCelsius conversion
    calories_per_hour_square_meter_degree_celsiusDto := units.HeatTransferCoefficientDto{
        Value: 100,
        Unit:  units.HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius,
    }
    
    var calories_per_hour_square_meter_degree_celsiusResult *units.HeatTransferCoefficient
    calories_per_hour_square_meter_degree_celsiusResult, err = factory.FromDto(calories_per_hour_square_meter_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with CaloriePerHourSquareMeterDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = calories_per_hour_square_meter_degree_celsiusResult.Convert(units.HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CaloriePerHourSquareMeterDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test KilocaloriePerHourSquareMeterDegreeCelsius conversion
    kilocalories_per_hour_square_meter_degree_celsiusDto := units.HeatTransferCoefficientDto{
        Value: 100,
        Unit:  units.HeatTransferCoefficientKilocaloriePerHourSquareMeterDegreeCelsius,
    }
    
    var kilocalories_per_hour_square_meter_degree_celsiusResult *units.HeatTransferCoefficient
    kilocalories_per_hour_square_meter_degree_celsiusResult, err = factory.FromDto(kilocalories_per_hour_square_meter_degree_celsiusDto)
    if err != nil {
        t.Errorf("FromDto() with KilocaloriePerHourSquareMeterDegreeCelsius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_hour_square_meter_degree_celsiusResult.Convert(units.HeatTransferCoefficientKilocaloriePerHourSquareMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerHourSquareMeterDegreeCelsius = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.HeatTransferCoefficientDto{
        Value: 0,
        Unit:  units.HeatTransferCoefficientWattPerSquareMeterKelvin,
    }
    
    var zeroResult *units.HeatTransferCoefficient
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestHeatTransferCoefficientFactory_FromDtoJSON(t *testing.T) {
    factory := units.HeatTransferCoefficientFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "WattPerSquareMeterKelvin"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "WattPerSquareMeterKelvin"}`)
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
    nanJSON, _ := json.Marshal(units.HeatTransferCoefficientDto{
        Value: nanValue,
        Unit:  units.HeatTransferCoefficientWattPerSquareMeterKelvin,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with WattPerSquareMeterKelvin unit
    watts_per_square_meter_kelvinJSON := []byte(`{"value": 100, "unit": "WattPerSquareMeterKelvin"}`)
    watts_per_square_meter_kelvinResult, err := factory.FromDtoJSON(watts_per_square_meter_kelvinJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerSquareMeterKelvin unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_meter_kelvinResult.Convert(units.HeatTransferCoefficientWattPerSquareMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareMeterKelvin = %v, want %v", converted, 100)
    }
    // Test JSON with WattPerSquareMeterCelsius unit
    watts_per_square_meter_celsiusJSON := []byte(`{"value": 100, "unit": "WattPerSquareMeterCelsius"}`)
    watts_per_square_meter_celsiusResult, err := factory.FromDtoJSON(watts_per_square_meter_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerSquareMeterCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_meter_celsiusResult.Convert(units.HeatTransferCoefficientWattPerSquareMeterCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareMeterCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with BtuPerHourSquareFootDegreeFahrenheit unit
    btus_per_hour_square_foot_degree_fahrenheitJSON := []byte(`{"value": 100, "unit": "BtuPerHourSquareFootDegreeFahrenheit"}`)
    btus_per_hour_square_foot_degree_fahrenheitResult, err := factory.FromDtoJSON(btus_per_hour_square_foot_degree_fahrenheitJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BtuPerHourSquareFootDegreeFahrenheit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btus_per_hour_square_foot_degree_fahrenheitResult.Convert(units.HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerHourSquareFootDegreeFahrenheit = %v, want %v", converted, 100)
    }
    // Test JSON with CaloriePerHourSquareMeterDegreeCelsius unit
    calories_per_hour_square_meter_degree_celsiusJSON := []byte(`{"value": 100, "unit": "CaloriePerHourSquareMeterDegreeCelsius"}`)
    calories_per_hour_square_meter_degree_celsiusResult, err := factory.FromDtoJSON(calories_per_hour_square_meter_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CaloriePerHourSquareMeterDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = calories_per_hour_square_meter_degree_celsiusResult.Convert(units.HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CaloriePerHourSquareMeterDegreeCelsius = %v, want %v", converted, 100)
    }
    // Test JSON with KilocaloriePerHourSquareMeterDegreeCelsius unit
    kilocalories_per_hour_square_meter_degree_celsiusJSON := []byte(`{"value": 100, "unit": "KilocaloriePerHourSquareMeterDegreeCelsius"}`)
    kilocalories_per_hour_square_meter_degree_celsiusResult, err := factory.FromDtoJSON(kilocalories_per_hour_square_meter_degree_celsiusJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilocaloriePerHourSquareMeterDegreeCelsius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_hour_square_meter_degree_celsiusResult.Convert(units.HeatTransferCoefficientKilocaloriePerHourSquareMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerHourSquareMeterDegreeCelsius = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "WattPerSquareMeterKelvin"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromWattsPerSquareMeterKelvin function
func TestHeatTransferCoefficientFactory_FromWattsPerSquareMeterKelvin(t *testing.T) {
    factory := units.HeatTransferCoefficientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerSquareMeterKelvin(100)
    if err != nil {
        t.Errorf("FromWattsPerSquareMeterKelvin() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatTransferCoefficientWattPerSquareMeterKelvin)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerSquareMeterKelvin() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerSquareMeterKelvin(math.NaN())
    if err == nil {
        t.Error("FromWattsPerSquareMeterKelvin() with NaN value should return error")
    }

    _, err = factory.FromWattsPerSquareMeterKelvin(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerSquareMeterKelvin() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerSquareMeterKelvin(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerSquareMeterKelvin() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerSquareMeterKelvin(0)
    if err != nil {
        t.Errorf("FromWattsPerSquareMeterKelvin() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatTransferCoefficientWattPerSquareMeterKelvin)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerSquareMeterKelvin() with zero value = %v, want 0", converted)
    }
}
// Test FromWattsPerSquareMeterCelsius function
func TestHeatTransferCoefficientFactory_FromWattsPerSquareMeterCelsius(t *testing.T) {
    factory := units.HeatTransferCoefficientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerSquareMeterCelsius(100)
    if err != nil {
        t.Errorf("FromWattsPerSquareMeterCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatTransferCoefficientWattPerSquareMeterCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerSquareMeterCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerSquareMeterCelsius(math.NaN())
    if err == nil {
        t.Error("FromWattsPerSquareMeterCelsius() with NaN value should return error")
    }

    _, err = factory.FromWattsPerSquareMeterCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerSquareMeterCelsius() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerSquareMeterCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerSquareMeterCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerSquareMeterCelsius(0)
    if err != nil {
        t.Errorf("FromWattsPerSquareMeterCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatTransferCoefficientWattPerSquareMeterCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerSquareMeterCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromBtusPerHourSquareFootDegreeFahrenheit function
func TestHeatTransferCoefficientFactory_FromBtusPerHourSquareFootDegreeFahrenheit(t *testing.T) {
    factory := units.HeatTransferCoefficientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBtusPerHourSquareFootDegreeFahrenheit(100)
    if err != nil {
        t.Errorf("FromBtusPerHourSquareFootDegreeFahrenheit() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBtusPerHourSquareFootDegreeFahrenheit() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBtusPerHourSquareFootDegreeFahrenheit(math.NaN())
    if err == nil {
        t.Error("FromBtusPerHourSquareFootDegreeFahrenheit() with NaN value should return error")
    }

    _, err = factory.FromBtusPerHourSquareFootDegreeFahrenheit(math.Inf(1))
    if err == nil {
        t.Error("FromBtusPerHourSquareFootDegreeFahrenheit() with +Inf value should return error")
    }

    _, err = factory.FromBtusPerHourSquareFootDegreeFahrenheit(math.Inf(-1))
    if err == nil {
        t.Error("FromBtusPerHourSquareFootDegreeFahrenheit() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBtusPerHourSquareFootDegreeFahrenheit(0)
    if err != nil {
        t.Errorf("FromBtusPerHourSquareFootDegreeFahrenheit() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBtusPerHourSquareFootDegreeFahrenheit() with zero value = %v, want 0", converted)
    }
}
// Test FromCaloriesPerHourSquareMeterDegreeCelsius function
func TestHeatTransferCoefficientFactory_FromCaloriesPerHourSquareMeterDegreeCelsius(t *testing.T) {
    factory := units.HeatTransferCoefficientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCaloriesPerHourSquareMeterDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromCaloriesPerHourSquareMeterDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCaloriesPerHourSquareMeterDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCaloriesPerHourSquareMeterDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromCaloriesPerHourSquareMeterDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromCaloriesPerHourSquareMeterDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromCaloriesPerHourSquareMeterDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromCaloriesPerHourSquareMeterDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromCaloriesPerHourSquareMeterDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCaloriesPerHourSquareMeterDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromCaloriesPerHourSquareMeterDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCaloriesPerHourSquareMeterDegreeCelsius() with zero value = %v, want 0", converted)
    }
}
// Test FromKilocaloriesPerHourSquareMeterDegreeCelsius function
func TestHeatTransferCoefficientFactory_FromKilocaloriesPerHourSquareMeterDegreeCelsius(t *testing.T) {
    factory := units.HeatTransferCoefficientFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilocaloriesPerHourSquareMeterDegreeCelsius(100)
    if err != nil {
        t.Errorf("FromKilocaloriesPerHourSquareMeterDegreeCelsius() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.HeatTransferCoefficientKilocaloriePerHourSquareMeterDegreeCelsius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilocaloriesPerHourSquareMeterDegreeCelsius() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilocaloriesPerHourSquareMeterDegreeCelsius(math.NaN())
    if err == nil {
        t.Error("FromKilocaloriesPerHourSquareMeterDegreeCelsius() with NaN value should return error")
    }

    _, err = factory.FromKilocaloriesPerHourSquareMeterDegreeCelsius(math.Inf(1))
    if err == nil {
        t.Error("FromKilocaloriesPerHourSquareMeterDegreeCelsius() with +Inf value should return error")
    }

    _, err = factory.FromKilocaloriesPerHourSquareMeterDegreeCelsius(math.Inf(-1))
    if err == nil {
        t.Error("FromKilocaloriesPerHourSquareMeterDegreeCelsius() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilocaloriesPerHourSquareMeterDegreeCelsius(0)
    if err != nil {
        t.Errorf("FromKilocaloriesPerHourSquareMeterDegreeCelsius() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.HeatTransferCoefficientKilocaloriePerHourSquareMeterDegreeCelsius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilocaloriesPerHourSquareMeterDegreeCelsius() with zero value = %v, want 0", converted)
    }
}

func TestHeatTransferCoefficientToString(t *testing.T) {
	factory := units.HeatTransferCoefficientFactory{}
	a, err := factory.CreateHeatTransferCoefficient(45, units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.HeatTransferCoefficientWattPerSquareMeterKelvin, 2)
	expected := "45.00 " + units.GetHeatTransferCoefficientAbbreviation(units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.HeatTransferCoefficientWattPerSquareMeterKelvin, -1)
	expected = "45 " + units.GetHeatTransferCoefficientAbbreviation(units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestHeatTransferCoefficient_EqualityAndComparison(t *testing.T) {
	factory := units.HeatTransferCoefficientFactory{}
	a1, _ := factory.CreateHeatTransferCoefficient(60, units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	a2, _ := factory.CreateHeatTransferCoefficient(60, units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	a3, _ := factory.CreateHeatTransferCoefficient(120, units.HeatTransferCoefficientWattPerSquareMeterKelvin)

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

func TestHeatTransferCoefficient_Arithmetic(t *testing.T) {
	factory := units.HeatTransferCoefficientFactory{}
	a1, _ := factory.CreateHeatTransferCoefficient(30, units.HeatTransferCoefficientWattPerSquareMeterKelvin)
	a2, _ := factory.CreateHeatTransferCoefficient(45, units.HeatTransferCoefficientWattPerSquareMeterKelvin)

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