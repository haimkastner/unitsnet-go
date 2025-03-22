// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestSpecificEnergyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "JoulePerKilogram"}`
	
	factory := units.SpecificEnergyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.SpecificEnergyJoulePerKilogram {
		t.Errorf("expected unit %v, got %v", units.SpecificEnergyJoulePerKilogram, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "JoulePerKilogram"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestSpecificEnergyDto_ToJSON(t *testing.T) {
	dto := units.SpecificEnergyDto{
		Value: 45,
		Unit:  units.SpecificEnergyJoulePerKilogram,
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
	if result["unit"].(string) != string(units.SpecificEnergyJoulePerKilogram) {
		t.Errorf("expected unit %s, got %v", units.SpecificEnergyJoulePerKilogram, result["unit"])
	}
}

func TestNewSpecificEnergy_InvalidValue(t *testing.T) {
	factory := units.SpecificEnergyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateSpecificEnergy(math.NaN(), units.SpecificEnergyJoulePerKilogram)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateSpecificEnergy(math.Inf(1), units.SpecificEnergyJoulePerKilogram)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestSpecificEnergyConversions(t *testing.T) {
	factory := units.SpecificEnergyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateSpecificEnergy(180, units.SpecificEnergyJoulePerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to JoulesPerKilogram.
		// No expected conversion value provided for JoulesPerKilogram, verifying result is not NaN.
		result := a.JoulesPerKilogram()
		cacheResult := a.JoulesPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to JoulesPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MegaJoulesPerTonne.
		// No expected conversion value provided for MegaJoulesPerTonne, verifying result is not NaN.
		result := a.MegaJoulesPerTonne()
		cacheResult := a.MegaJoulesPerTonne()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegaJoulesPerTonne returned NaN")
		}
	}
	{
		// Test conversion to CaloriesPerGram.
		// No expected conversion value provided for CaloriesPerGram, verifying result is not NaN.
		result := a.CaloriesPerGram()
		cacheResult := a.CaloriesPerGram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CaloriesPerGram returned NaN")
		}
	}
	{
		// Test conversion to WattHoursPerKilogram.
		// No expected conversion value provided for WattHoursPerKilogram, verifying result is not NaN.
		result := a.WattHoursPerKilogram()
		cacheResult := a.WattHoursPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattHoursPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to WattDaysPerKilogram.
		// No expected conversion value provided for WattDaysPerKilogram, verifying result is not NaN.
		result := a.WattDaysPerKilogram()
		cacheResult := a.WattDaysPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattDaysPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to WattDaysPerTonne.
		// No expected conversion value provided for WattDaysPerTonne, verifying result is not NaN.
		result := a.WattDaysPerTonne()
		cacheResult := a.WattDaysPerTonne()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattDaysPerTonne returned NaN")
		}
	}
	{
		// Test conversion to WattDaysPerShortTon.
		// No expected conversion value provided for WattDaysPerShortTon, verifying result is not NaN.
		result := a.WattDaysPerShortTon()
		cacheResult := a.WattDaysPerShortTon()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattDaysPerShortTon returned NaN")
		}
	}
	{
		// Test conversion to WattHoursPerPound.
		// No expected conversion value provided for WattHoursPerPound, verifying result is not NaN.
		result := a.WattHoursPerPound()
		cacheResult := a.WattHoursPerPound()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to WattHoursPerPound returned NaN")
		}
	}
	{
		// Test conversion to BtuPerPound.
		// No expected conversion value provided for BtuPerPound, verifying result is not NaN.
		result := a.BtuPerPound()
		cacheResult := a.BtuPerPound()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to BtuPerPound returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerKilogram.
		// No expected conversion value provided for KilojoulesPerKilogram, verifying result is not NaN.
		result := a.KilojoulesPerKilogram()
		cacheResult := a.KilojoulesPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilojoulesPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerKilogram.
		// No expected conversion value provided for MegajoulesPerKilogram, verifying result is not NaN.
		result := a.MegajoulesPerKilogram()
		cacheResult := a.MegajoulesPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegajoulesPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to KilocaloriesPerGram.
		// No expected conversion value provided for KilocaloriesPerGram, verifying result is not NaN.
		result := a.KilocaloriesPerGram()
		cacheResult := a.KilocaloriesPerGram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilocaloriesPerGram returned NaN")
		}
	}
	{
		// Test conversion to KilowattHoursPerKilogram.
		// No expected conversion value provided for KilowattHoursPerKilogram, verifying result is not NaN.
		result := a.KilowattHoursPerKilogram()
		cacheResult := a.KilowattHoursPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattHoursPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MegawattHoursPerKilogram.
		// No expected conversion value provided for MegawattHoursPerKilogram, verifying result is not NaN.
		result := a.MegawattHoursPerKilogram()
		cacheResult := a.MegawattHoursPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattHoursPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to GigawattHoursPerKilogram.
		// No expected conversion value provided for GigawattHoursPerKilogram, verifying result is not NaN.
		result := a.GigawattHoursPerKilogram()
		cacheResult := a.GigawattHoursPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattHoursPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to KilowattDaysPerKilogram.
		// No expected conversion value provided for KilowattDaysPerKilogram, verifying result is not NaN.
		result := a.KilowattDaysPerKilogram()
		cacheResult := a.KilowattDaysPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattDaysPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to MegawattDaysPerKilogram.
		// No expected conversion value provided for MegawattDaysPerKilogram, verifying result is not NaN.
		result := a.MegawattDaysPerKilogram()
		cacheResult := a.MegawattDaysPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattDaysPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to GigawattDaysPerKilogram.
		// No expected conversion value provided for GigawattDaysPerKilogram, verifying result is not NaN.
		result := a.GigawattDaysPerKilogram()
		cacheResult := a.GigawattDaysPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattDaysPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to TerawattDaysPerKilogram.
		// No expected conversion value provided for TerawattDaysPerKilogram, verifying result is not NaN.
		result := a.TerawattDaysPerKilogram()
		cacheResult := a.TerawattDaysPerKilogram()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TerawattDaysPerKilogram returned NaN")
		}
	}
	{
		// Test conversion to KilowattDaysPerTonne.
		// No expected conversion value provided for KilowattDaysPerTonne, verifying result is not NaN.
		result := a.KilowattDaysPerTonne()
		cacheResult := a.KilowattDaysPerTonne()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattDaysPerTonne returned NaN")
		}
	}
	{
		// Test conversion to MegawattDaysPerTonne.
		// No expected conversion value provided for MegawattDaysPerTonne, verifying result is not NaN.
		result := a.MegawattDaysPerTonne()
		cacheResult := a.MegawattDaysPerTonne()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattDaysPerTonne returned NaN")
		}
	}
	{
		// Test conversion to GigawattDaysPerTonne.
		// No expected conversion value provided for GigawattDaysPerTonne, verifying result is not NaN.
		result := a.GigawattDaysPerTonne()
		cacheResult := a.GigawattDaysPerTonne()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattDaysPerTonne returned NaN")
		}
	}
	{
		// Test conversion to TerawattDaysPerTonne.
		// No expected conversion value provided for TerawattDaysPerTonne, verifying result is not NaN.
		result := a.TerawattDaysPerTonne()
		cacheResult := a.TerawattDaysPerTonne()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TerawattDaysPerTonne returned NaN")
		}
	}
	{
		// Test conversion to KilowattDaysPerShortTon.
		// No expected conversion value provided for KilowattDaysPerShortTon, verifying result is not NaN.
		result := a.KilowattDaysPerShortTon()
		cacheResult := a.KilowattDaysPerShortTon()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattDaysPerShortTon returned NaN")
		}
	}
	{
		// Test conversion to MegawattDaysPerShortTon.
		// No expected conversion value provided for MegawattDaysPerShortTon, verifying result is not NaN.
		result := a.MegawattDaysPerShortTon()
		cacheResult := a.MegawattDaysPerShortTon()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattDaysPerShortTon returned NaN")
		}
	}
	{
		// Test conversion to GigawattDaysPerShortTon.
		// No expected conversion value provided for GigawattDaysPerShortTon, verifying result is not NaN.
		result := a.GigawattDaysPerShortTon()
		cacheResult := a.GigawattDaysPerShortTon()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattDaysPerShortTon returned NaN")
		}
	}
	{
		// Test conversion to TerawattDaysPerShortTon.
		// No expected conversion value provided for TerawattDaysPerShortTon, verifying result is not NaN.
		result := a.TerawattDaysPerShortTon()
		cacheResult := a.TerawattDaysPerShortTon()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TerawattDaysPerShortTon returned NaN")
		}
	}
	{
		// Test conversion to KilowattHoursPerPound.
		// No expected conversion value provided for KilowattHoursPerPound, verifying result is not NaN.
		result := a.KilowattHoursPerPound()
		cacheResult := a.KilowattHoursPerPound()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilowattHoursPerPound returned NaN")
		}
	}
	{
		// Test conversion to MegawattHoursPerPound.
		// No expected conversion value provided for MegawattHoursPerPound, verifying result is not NaN.
		result := a.MegawattHoursPerPound()
		cacheResult := a.MegawattHoursPerPound()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegawattHoursPerPound returned NaN")
		}
	}
	{
		// Test conversion to GigawattHoursPerPound.
		// No expected conversion value provided for GigawattHoursPerPound, verifying result is not NaN.
		result := a.GigawattHoursPerPound()
		cacheResult := a.GigawattHoursPerPound()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to GigawattHoursPerPound returned NaN")
		}
	}
}

func TestSpecificEnergy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.SpecificEnergyFactory{}
	a, err := factory.CreateSpecificEnergy(90, units.SpecificEnergyJoulePerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.SpecificEnergyJoulePerKilogram {
		t.Errorf("expected default unit JoulePerKilogram, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.SpecificEnergyJoulePerKilogram
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.SpecificEnergyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.SpecificEnergyJoulePerKilogram {
		t.Errorf("expected unit JoulePerKilogram, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestSpecificEnergyFactory_FromDto(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyJoulePerKilogram,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.SpecificEnergyDto{
        Value: math.NaN(),
        Unit:  units.SpecificEnergyJoulePerKilogram,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test JoulePerKilogram conversion
    joules_per_kilogramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyJoulePerKilogram,
    }
    
    var joules_per_kilogramResult *units.SpecificEnergy
    joules_per_kilogramResult, err = factory.FromDto(joules_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_kilogramResult.Convert(units.SpecificEnergyJoulePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerKilogram = %v, want %v", converted, 100)
    }
    // Test MegaJoulePerTonne conversion
    mega_joules_per_tonneDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyMegaJoulePerTonne,
    }
    
    var mega_joules_per_tonneResult *units.SpecificEnergy
    mega_joules_per_tonneResult, err = factory.FromDto(mega_joules_per_tonneDto)
    if err != nil {
        t.Errorf("FromDto() with MegaJoulePerTonne returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mega_joules_per_tonneResult.Convert(units.SpecificEnergyMegaJoulePerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaJoulePerTonne = %v, want %v", converted, 100)
    }
    // Test CaloriePerGram conversion
    calories_per_gramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyCaloriePerGram,
    }
    
    var calories_per_gramResult *units.SpecificEnergy
    calories_per_gramResult, err = factory.FromDto(calories_per_gramDto)
    if err != nil {
        t.Errorf("FromDto() with CaloriePerGram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = calories_per_gramResult.Convert(units.SpecificEnergyCaloriePerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CaloriePerGram = %v, want %v", converted, 100)
    }
    // Test WattHourPerKilogram conversion
    watt_hours_per_kilogramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyWattHourPerKilogram,
    }
    
    var watt_hours_per_kilogramResult *units.SpecificEnergy
    watt_hours_per_kilogramResult, err = factory.FromDto(watt_hours_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with WattHourPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_hours_per_kilogramResult.Convert(units.SpecificEnergyWattHourPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattHourPerKilogram = %v, want %v", converted, 100)
    }
    // Test WattDayPerKilogram conversion
    watt_days_per_kilogramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyWattDayPerKilogram,
    }
    
    var watt_days_per_kilogramResult *units.SpecificEnergy
    watt_days_per_kilogramResult, err = factory.FromDto(watt_days_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with WattDayPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_days_per_kilogramResult.Convert(units.SpecificEnergyWattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattDayPerKilogram = %v, want %v", converted, 100)
    }
    // Test WattDayPerTonne conversion
    watt_days_per_tonneDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyWattDayPerTonne,
    }
    
    var watt_days_per_tonneResult *units.SpecificEnergy
    watt_days_per_tonneResult, err = factory.FromDto(watt_days_per_tonneDto)
    if err != nil {
        t.Errorf("FromDto() with WattDayPerTonne returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_days_per_tonneResult.Convert(units.SpecificEnergyWattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattDayPerTonne = %v, want %v", converted, 100)
    }
    // Test WattDayPerShortTon conversion
    watt_days_per_short_tonDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyWattDayPerShortTon,
    }
    
    var watt_days_per_short_tonResult *units.SpecificEnergy
    watt_days_per_short_tonResult, err = factory.FromDto(watt_days_per_short_tonDto)
    if err != nil {
        t.Errorf("FromDto() with WattDayPerShortTon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_days_per_short_tonResult.Convert(units.SpecificEnergyWattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattDayPerShortTon = %v, want %v", converted, 100)
    }
    // Test WattHourPerPound conversion
    watt_hours_per_poundDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyWattHourPerPound,
    }
    
    var watt_hours_per_poundResult *units.SpecificEnergy
    watt_hours_per_poundResult, err = factory.FromDto(watt_hours_per_poundDto)
    if err != nil {
        t.Errorf("FromDto() with WattHourPerPound returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_hours_per_poundResult.Convert(units.SpecificEnergyWattHourPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattHourPerPound = %v, want %v", converted, 100)
    }
    // Test BtuPerPound conversion
    btu_per_poundDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyBtuPerPound,
    }
    
    var btu_per_poundResult *units.SpecificEnergy
    btu_per_poundResult, err = factory.FromDto(btu_per_poundDto)
    if err != nil {
        t.Errorf("FromDto() with BtuPerPound returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btu_per_poundResult.Convert(units.SpecificEnergyBtuPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerPound = %v, want %v", converted, 100)
    }
    // Test KilojoulePerKilogram conversion
    kilojoules_per_kilogramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyKilojoulePerKilogram,
    }
    
    var kilojoules_per_kilogramResult *units.SpecificEnergy
    kilojoules_per_kilogramResult, err = factory.FromDto(kilojoules_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with KilojoulePerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_kilogramResult.Convert(units.SpecificEnergyKilojoulePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerKilogram = %v, want %v", converted, 100)
    }
    // Test MegajoulePerKilogram conversion
    megajoules_per_kilogramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyMegajoulePerKilogram,
    }
    
    var megajoules_per_kilogramResult *units.SpecificEnergy
    megajoules_per_kilogramResult, err = factory.FromDto(megajoules_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with MegajoulePerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_kilogramResult.Convert(units.SpecificEnergyMegajoulePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerKilogram = %v, want %v", converted, 100)
    }
    // Test KilocaloriePerGram conversion
    kilocalories_per_gramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyKilocaloriePerGram,
    }
    
    var kilocalories_per_gramResult *units.SpecificEnergy
    kilocalories_per_gramResult, err = factory.FromDto(kilocalories_per_gramDto)
    if err != nil {
        t.Errorf("FromDto() with KilocaloriePerGram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_gramResult.Convert(units.SpecificEnergyKilocaloriePerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerGram = %v, want %v", converted, 100)
    }
    // Test KilowattHourPerKilogram conversion
    kilowatt_hours_per_kilogramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyKilowattHourPerKilogram,
    }
    
    var kilowatt_hours_per_kilogramResult *units.SpecificEnergy
    kilowatt_hours_per_kilogramResult, err = factory.FromDto(kilowatt_hours_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattHourPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_hours_per_kilogramResult.Convert(units.SpecificEnergyKilowattHourPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattHourPerKilogram = %v, want %v", converted, 100)
    }
    // Test MegawattHourPerKilogram conversion
    megawatt_hours_per_kilogramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyMegawattHourPerKilogram,
    }
    
    var megawatt_hours_per_kilogramResult *units.SpecificEnergy
    megawatt_hours_per_kilogramResult, err = factory.FromDto(megawatt_hours_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattHourPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_hours_per_kilogramResult.Convert(units.SpecificEnergyMegawattHourPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattHourPerKilogram = %v, want %v", converted, 100)
    }
    // Test GigawattHourPerKilogram conversion
    gigawatt_hours_per_kilogramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyGigawattHourPerKilogram,
    }
    
    var gigawatt_hours_per_kilogramResult *units.SpecificEnergy
    gigawatt_hours_per_kilogramResult, err = factory.FromDto(gigawatt_hours_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattHourPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_hours_per_kilogramResult.Convert(units.SpecificEnergyGigawattHourPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattHourPerKilogram = %v, want %v", converted, 100)
    }
    // Test KilowattDayPerKilogram conversion
    kilowatt_days_per_kilogramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyKilowattDayPerKilogram,
    }
    
    var kilowatt_days_per_kilogramResult *units.SpecificEnergy
    kilowatt_days_per_kilogramResult, err = factory.FromDto(kilowatt_days_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattDayPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_days_per_kilogramResult.Convert(units.SpecificEnergyKilowattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattDayPerKilogram = %v, want %v", converted, 100)
    }
    // Test MegawattDayPerKilogram conversion
    megawatt_days_per_kilogramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyMegawattDayPerKilogram,
    }
    
    var megawatt_days_per_kilogramResult *units.SpecificEnergy
    megawatt_days_per_kilogramResult, err = factory.FromDto(megawatt_days_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattDayPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_days_per_kilogramResult.Convert(units.SpecificEnergyMegawattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattDayPerKilogram = %v, want %v", converted, 100)
    }
    // Test GigawattDayPerKilogram conversion
    gigawatt_days_per_kilogramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyGigawattDayPerKilogram,
    }
    
    var gigawatt_days_per_kilogramResult *units.SpecificEnergy
    gigawatt_days_per_kilogramResult, err = factory.FromDto(gigawatt_days_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattDayPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_days_per_kilogramResult.Convert(units.SpecificEnergyGigawattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattDayPerKilogram = %v, want %v", converted, 100)
    }
    // Test TerawattDayPerKilogram conversion
    terawatt_days_per_kilogramDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyTerawattDayPerKilogram,
    }
    
    var terawatt_days_per_kilogramResult *units.SpecificEnergy
    terawatt_days_per_kilogramResult, err = factory.FromDto(terawatt_days_per_kilogramDto)
    if err != nil {
        t.Errorf("FromDto() with TerawattDayPerKilogram returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatt_days_per_kilogramResult.Convert(units.SpecificEnergyTerawattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattDayPerKilogram = %v, want %v", converted, 100)
    }
    // Test KilowattDayPerTonne conversion
    kilowatt_days_per_tonneDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyKilowattDayPerTonne,
    }
    
    var kilowatt_days_per_tonneResult *units.SpecificEnergy
    kilowatt_days_per_tonneResult, err = factory.FromDto(kilowatt_days_per_tonneDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattDayPerTonne returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_days_per_tonneResult.Convert(units.SpecificEnergyKilowattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattDayPerTonne = %v, want %v", converted, 100)
    }
    // Test MegawattDayPerTonne conversion
    megawatt_days_per_tonneDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyMegawattDayPerTonne,
    }
    
    var megawatt_days_per_tonneResult *units.SpecificEnergy
    megawatt_days_per_tonneResult, err = factory.FromDto(megawatt_days_per_tonneDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattDayPerTonne returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_days_per_tonneResult.Convert(units.SpecificEnergyMegawattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattDayPerTonne = %v, want %v", converted, 100)
    }
    // Test GigawattDayPerTonne conversion
    gigawatt_days_per_tonneDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyGigawattDayPerTonne,
    }
    
    var gigawatt_days_per_tonneResult *units.SpecificEnergy
    gigawatt_days_per_tonneResult, err = factory.FromDto(gigawatt_days_per_tonneDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattDayPerTonne returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_days_per_tonneResult.Convert(units.SpecificEnergyGigawattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattDayPerTonne = %v, want %v", converted, 100)
    }
    // Test TerawattDayPerTonne conversion
    terawatt_days_per_tonneDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyTerawattDayPerTonne,
    }
    
    var terawatt_days_per_tonneResult *units.SpecificEnergy
    terawatt_days_per_tonneResult, err = factory.FromDto(terawatt_days_per_tonneDto)
    if err != nil {
        t.Errorf("FromDto() with TerawattDayPerTonne returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatt_days_per_tonneResult.Convert(units.SpecificEnergyTerawattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattDayPerTonne = %v, want %v", converted, 100)
    }
    // Test KilowattDayPerShortTon conversion
    kilowatt_days_per_short_tonDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyKilowattDayPerShortTon,
    }
    
    var kilowatt_days_per_short_tonResult *units.SpecificEnergy
    kilowatt_days_per_short_tonResult, err = factory.FromDto(kilowatt_days_per_short_tonDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattDayPerShortTon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_days_per_short_tonResult.Convert(units.SpecificEnergyKilowattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattDayPerShortTon = %v, want %v", converted, 100)
    }
    // Test MegawattDayPerShortTon conversion
    megawatt_days_per_short_tonDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyMegawattDayPerShortTon,
    }
    
    var megawatt_days_per_short_tonResult *units.SpecificEnergy
    megawatt_days_per_short_tonResult, err = factory.FromDto(megawatt_days_per_short_tonDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattDayPerShortTon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_days_per_short_tonResult.Convert(units.SpecificEnergyMegawattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattDayPerShortTon = %v, want %v", converted, 100)
    }
    // Test GigawattDayPerShortTon conversion
    gigawatt_days_per_short_tonDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyGigawattDayPerShortTon,
    }
    
    var gigawatt_days_per_short_tonResult *units.SpecificEnergy
    gigawatt_days_per_short_tonResult, err = factory.FromDto(gigawatt_days_per_short_tonDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattDayPerShortTon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_days_per_short_tonResult.Convert(units.SpecificEnergyGigawattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattDayPerShortTon = %v, want %v", converted, 100)
    }
    // Test TerawattDayPerShortTon conversion
    terawatt_days_per_short_tonDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyTerawattDayPerShortTon,
    }
    
    var terawatt_days_per_short_tonResult *units.SpecificEnergy
    terawatt_days_per_short_tonResult, err = factory.FromDto(terawatt_days_per_short_tonDto)
    if err != nil {
        t.Errorf("FromDto() with TerawattDayPerShortTon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatt_days_per_short_tonResult.Convert(units.SpecificEnergyTerawattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattDayPerShortTon = %v, want %v", converted, 100)
    }
    // Test KilowattHourPerPound conversion
    kilowatt_hours_per_poundDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyKilowattHourPerPound,
    }
    
    var kilowatt_hours_per_poundResult *units.SpecificEnergy
    kilowatt_hours_per_poundResult, err = factory.FromDto(kilowatt_hours_per_poundDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattHourPerPound returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_hours_per_poundResult.Convert(units.SpecificEnergyKilowattHourPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattHourPerPound = %v, want %v", converted, 100)
    }
    // Test MegawattHourPerPound conversion
    megawatt_hours_per_poundDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyMegawattHourPerPound,
    }
    
    var megawatt_hours_per_poundResult *units.SpecificEnergy
    megawatt_hours_per_poundResult, err = factory.FromDto(megawatt_hours_per_poundDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattHourPerPound returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_hours_per_poundResult.Convert(units.SpecificEnergyMegawattHourPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattHourPerPound = %v, want %v", converted, 100)
    }
    // Test GigawattHourPerPound conversion
    gigawatt_hours_per_poundDto := units.SpecificEnergyDto{
        Value: 100,
        Unit:  units.SpecificEnergyGigawattHourPerPound,
    }
    
    var gigawatt_hours_per_poundResult *units.SpecificEnergy
    gigawatt_hours_per_poundResult, err = factory.FromDto(gigawatt_hours_per_poundDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattHourPerPound returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_hours_per_poundResult.Convert(units.SpecificEnergyGigawattHourPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattHourPerPound = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.SpecificEnergyDto{
        Value: 0,
        Unit:  units.SpecificEnergyJoulePerKilogram,
    }
    
    var zeroResult *units.SpecificEnergy
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestSpecificEnergyFactory_FromDtoJSON(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "JoulePerKilogram"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "JoulePerKilogram"}`)
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
    nanJSON, _ := json.Marshal(units.SpecificEnergyDto{
        Value: nanValue,
        Unit:  units.SpecificEnergyJoulePerKilogram,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with JoulePerKilogram unit
    joules_per_kilogramJSON := []byte(`{"value": 100, "unit": "JoulePerKilogram"}`)
    joules_per_kilogramResult, err := factory.FromDtoJSON(joules_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_kilogramResult.Convert(units.SpecificEnergyJoulePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with MegaJoulePerTonne unit
    mega_joules_per_tonneJSON := []byte(`{"value": 100, "unit": "MegaJoulePerTonne"}`)
    mega_joules_per_tonneResult, err := factory.FromDtoJSON(mega_joules_per_tonneJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegaJoulePerTonne unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mega_joules_per_tonneResult.Convert(units.SpecificEnergyMegaJoulePerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaJoulePerTonne = %v, want %v", converted, 100)
    }
    // Test JSON with CaloriePerGram unit
    calories_per_gramJSON := []byte(`{"value": 100, "unit": "CaloriePerGram"}`)
    calories_per_gramResult, err := factory.FromDtoJSON(calories_per_gramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CaloriePerGram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = calories_per_gramResult.Convert(units.SpecificEnergyCaloriePerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CaloriePerGram = %v, want %v", converted, 100)
    }
    // Test JSON with WattHourPerKilogram unit
    watt_hours_per_kilogramJSON := []byte(`{"value": 100, "unit": "WattHourPerKilogram"}`)
    watt_hours_per_kilogramResult, err := factory.FromDtoJSON(watt_hours_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattHourPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_hours_per_kilogramResult.Convert(units.SpecificEnergyWattHourPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattHourPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with WattDayPerKilogram unit
    watt_days_per_kilogramJSON := []byte(`{"value": 100, "unit": "WattDayPerKilogram"}`)
    watt_days_per_kilogramResult, err := factory.FromDtoJSON(watt_days_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattDayPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_days_per_kilogramResult.Convert(units.SpecificEnergyWattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattDayPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with WattDayPerTonne unit
    watt_days_per_tonneJSON := []byte(`{"value": 100, "unit": "WattDayPerTonne"}`)
    watt_days_per_tonneResult, err := factory.FromDtoJSON(watt_days_per_tonneJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattDayPerTonne unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_days_per_tonneResult.Convert(units.SpecificEnergyWattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattDayPerTonne = %v, want %v", converted, 100)
    }
    // Test JSON with WattDayPerShortTon unit
    watt_days_per_short_tonJSON := []byte(`{"value": 100, "unit": "WattDayPerShortTon"}`)
    watt_days_per_short_tonResult, err := factory.FromDtoJSON(watt_days_per_short_tonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattDayPerShortTon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_days_per_short_tonResult.Convert(units.SpecificEnergyWattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattDayPerShortTon = %v, want %v", converted, 100)
    }
    // Test JSON with WattHourPerPound unit
    watt_hours_per_poundJSON := []byte(`{"value": 100, "unit": "WattHourPerPound"}`)
    watt_hours_per_poundResult, err := factory.FromDtoJSON(watt_hours_per_poundJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattHourPerPound unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_hours_per_poundResult.Convert(units.SpecificEnergyWattHourPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattHourPerPound = %v, want %v", converted, 100)
    }
    // Test JSON with BtuPerPound unit
    btu_per_poundJSON := []byte(`{"value": 100, "unit": "BtuPerPound"}`)
    btu_per_poundResult, err := factory.FromDtoJSON(btu_per_poundJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BtuPerPound unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = btu_per_poundResult.Convert(units.SpecificEnergyBtuPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BtuPerPound = %v, want %v", converted, 100)
    }
    // Test JSON with KilojoulePerKilogram unit
    kilojoules_per_kilogramJSON := []byte(`{"value": 100, "unit": "KilojoulePerKilogram"}`)
    kilojoules_per_kilogramResult, err := factory.FromDtoJSON(kilojoules_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilojoulePerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_kilogramResult.Convert(units.SpecificEnergyKilojoulePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with MegajoulePerKilogram unit
    megajoules_per_kilogramJSON := []byte(`{"value": 100, "unit": "MegajoulePerKilogram"}`)
    megajoules_per_kilogramResult, err := factory.FromDtoJSON(megajoules_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegajoulePerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_kilogramResult.Convert(units.SpecificEnergyMegajoulePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with KilocaloriePerGram unit
    kilocalories_per_gramJSON := []byte(`{"value": 100, "unit": "KilocaloriePerGram"}`)
    kilocalories_per_gramResult, err := factory.FromDtoJSON(kilocalories_per_gramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilocaloriePerGram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocalories_per_gramResult.Convert(units.SpecificEnergyKilocaloriePerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocaloriePerGram = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattHourPerKilogram unit
    kilowatt_hours_per_kilogramJSON := []byte(`{"value": 100, "unit": "KilowattHourPerKilogram"}`)
    kilowatt_hours_per_kilogramResult, err := factory.FromDtoJSON(kilowatt_hours_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattHourPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_hours_per_kilogramResult.Convert(units.SpecificEnergyKilowattHourPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattHourPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattHourPerKilogram unit
    megawatt_hours_per_kilogramJSON := []byte(`{"value": 100, "unit": "MegawattHourPerKilogram"}`)
    megawatt_hours_per_kilogramResult, err := factory.FromDtoJSON(megawatt_hours_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattHourPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_hours_per_kilogramResult.Convert(units.SpecificEnergyMegawattHourPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattHourPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattHourPerKilogram unit
    gigawatt_hours_per_kilogramJSON := []byte(`{"value": 100, "unit": "GigawattHourPerKilogram"}`)
    gigawatt_hours_per_kilogramResult, err := factory.FromDtoJSON(gigawatt_hours_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattHourPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_hours_per_kilogramResult.Convert(units.SpecificEnergyGigawattHourPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattHourPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattDayPerKilogram unit
    kilowatt_days_per_kilogramJSON := []byte(`{"value": 100, "unit": "KilowattDayPerKilogram"}`)
    kilowatt_days_per_kilogramResult, err := factory.FromDtoJSON(kilowatt_days_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattDayPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_days_per_kilogramResult.Convert(units.SpecificEnergyKilowattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattDayPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattDayPerKilogram unit
    megawatt_days_per_kilogramJSON := []byte(`{"value": 100, "unit": "MegawattDayPerKilogram"}`)
    megawatt_days_per_kilogramResult, err := factory.FromDtoJSON(megawatt_days_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattDayPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_days_per_kilogramResult.Convert(units.SpecificEnergyMegawattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattDayPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattDayPerKilogram unit
    gigawatt_days_per_kilogramJSON := []byte(`{"value": 100, "unit": "GigawattDayPerKilogram"}`)
    gigawatt_days_per_kilogramResult, err := factory.FromDtoJSON(gigawatt_days_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattDayPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_days_per_kilogramResult.Convert(units.SpecificEnergyGigawattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattDayPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with TerawattDayPerKilogram unit
    terawatt_days_per_kilogramJSON := []byte(`{"value": 100, "unit": "TerawattDayPerKilogram"}`)
    terawatt_days_per_kilogramResult, err := factory.FromDtoJSON(terawatt_days_per_kilogramJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TerawattDayPerKilogram unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatt_days_per_kilogramResult.Convert(units.SpecificEnergyTerawattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattDayPerKilogram = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattDayPerTonne unit
    kilowatt_days_per_tonneJSON := []byte(`{"value": 100, "unit": "KilowattDayPerTonne"}`)
    kilowatt_days_per_tonneResult, err := factory.FromDtoJSON(kilowatt_days_per_tonneJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattDayPerTonne unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_days_per_tonneResult.Convert(units.SpecificEnergyKilowattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattDayPerTonne = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattDayPerTonne unit
    megawatt_days_per_tonneJSON := []byte(`{"value": 100, "unit": "MegawattDayPerTonne"}`)
    megawatt_days_per_tonneResult, err := factory.FromDtoJSON(megawatt_days_per_tonneJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattDayPerTonne unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_days_per_tonneResult.Convert(units.SpecificEnergyMegawattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattDayPerTonne = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattDayPerTonne unit
    gigawatt_days_per_tonneJSON := []byte(`{"value": 100, "unit": "GigawattDayPerTonne"}`)
    gigawatt_days_per_tonneResult, err := factory.FromDtoJSON(gigawatt_days_per_tonneJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattDayPerTonne unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_days_per_tonneResult.Convert(units.SpecificEnergyGigawattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattDayPerTonne = %v, want %v", converted, 100)
    }
    // Test JSON with TerawattDayPerTonne unit
    terawatt_days_per_tonneJSON := []byte(`{"value": 100, "unit": "TerawattDayPerTonne"}`)
    terawatt_days_per_tonneResult, err := factory.FromDtoJSON(terawatt_days_per_tonneJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TerawattDayPerTonne unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatt_days_per_tonneResult.Convert(units.SpecificEnergyTerawattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattDayPerTonne = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattDayPerShortTon unit
    kilowatt_days_per_short_tonJSON := []byte(`{"value": 100, "unit": "KilowattDayPerShortTon"}`)
    kilowatt_days_per_short_tonResult, err := factory.FromDtoJSON(kilowatt_days_per_short_tonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattDayPerShortTon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_days_per_short_tonResult.Convert(units.SpecificEnergyKilowattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattDayPerShortTon = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattDayPerShortTon unit
    megawatt_days_per_short_tonJSON := []byte(`{"value": 100, "unit": "MegawattDayPerShortTon"}`)
    megawatt_days_per_short_tonResult, err := factory.FromDtoJSON(megawatt_days_per_short_tonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattDayPerShortTon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_days_per_short_tonResult.Convert(units.SpecificEnergyMegawattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattDayPerShortTon = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattDayPerShortTon unit
    gigawatt_days_per_short_tonJSON := []byte(`{"value": 100, "unit": "GigawattDayPerShortTon"}`)
    gigawatt_days_per_short_tonResult, err := factory.FromDtoJSON(gigawatt_days_per_short_tonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattDayPerShortTon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_days_per_short_tonResult.Convert(units.SpecificEnergyGigawattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattDayPerShortTon = %v, want %v", converted, 100)
    }
    // Test JSON with TerawattDayPerShortTon unit
    terawatt_days_per_short_tonJSON := []byte(`{"value": 100, "unit": "TerawattDayPerShortTon"}`)
    terawatt_days_per_short_tonResult, err := factory.FromDtoJSON(terawatt_days_per_short_tonJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TerawattDayPerShortTon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatt_days_per_short_tonResult.Convert(units.SpecificEnergyTerawattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattDayPerShortTon = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattHourPerPound unit
    kilowatt_hours_per_poundJSON := []byte(`{"value": 100, "unit": "KilowattHourPerPound"}`)
    kilowatt_hours_per_poundResult, err := factory.FromDtoJSON(kilowatt_hours_per_poundJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattHourPerPound unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_hours_per_poundResult.Convert(units.SpecificEnergyKilowattHourPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattHourPerPound = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattHourPerPound unit
    megawatt_hours_per_poundJSON := []byte(`{"value": 100, "unit": "MegawattHourPerPound"}`)
    megawatt_hours_per_poundResult, err := factory.FromDtoJSON(megawatt_hours_per_poundJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattHourPerPound unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_hours_per_poundResult.Convert(units.SpecificEnergyMegawattHourPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattHourPerPound = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattHourPerPound unit
    gigawatt_hours_per_poundJSON := []byte(`{"value": 100, "unit": "GigawattHourPerPound"}`)
    gigawatt_hours_per_poundResult, err := factory.FromDtoJSON(gigawatt_hours_per_poundJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattHourPerPound unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_hours_per_poundResult.Convert(units.SpecificEnergyGigawattHourPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattHourPerPound = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "JoulePerKilogram"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromJoulesPerKilogram function
func TestSpecificEnergyFactory_FromJoulesPerKilogram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerKilogram(100)
    if err != nil {
        t.Errorf("FromJoulesPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyJoulePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerKilogram(0)
    if err != nil {
        t.Errorf("FromJoulesPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyJoulePerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaJoulesPerTonne function
func TestSpecificEnergyFactory_FromMegaJoulesPerTonne(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaJoulesPerTonne(100)
    if err != nil {
        t.Errorf("FromMegaJoulesPerTonne() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyMegaJoulePerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegaJoulesPerTonne() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegaJoulesPerTonne(math.NaN())
    if err == nil {
        t.Error("FromMegaJoulesPerTonne() with NaN value should return error")
    }

    _, err = factory.FromMegaJoulesPerTonne(math.Inf(1))
    if err == nil {
        t.Error("FromMegaJoulesPerTonne() with +Inf value should return error")
    }

    _, err = factory.FromMegaJoulesPerTonne(math.Inf(-1))
    if err == nil {
        t.Error("FromMegaJoulesPerTonne() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegaJoulesPerTonne(0)
    if err != nil {
        t.Errorf("FromMegaJoulesPerTonne() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyMegaJoulePerTonne)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaJoulesPerTonne() with zero value = %v, want 0", converted)
    }
}
// Test FromCaloriesPerGram function
func TestSpecificEnergyFactory_FromCaloriesPerGram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCaloriesPerGram(100)
    if err != nil {
        t.Errorf("FromCaloriesPerGram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyCaloriePerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCaloriesPerGram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCaloriesPerGram(math.NaN())
    if err == nil {
        t.Error("FromCaloriesPerGram() with NaN value should return error")
    }

    _, err = factory.FromCaloriesPerGram(math.Inf(1))
    if err == nil {
        t.Error("FromCaloriesPerGram() with +Inf value should return error")
    }

    _, err = factory.FromCaloriesPerGram(math.Inf(-1))
    if err == nil {
        t.Error("FromCaloriesPerGram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCaloriesPerGram(0)
    if err != nil {
        t.Errorf("FromCaloriesPerGram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyCaloriePerGram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCaloriesPerGram() with zero value = %v, want 0", converted)
    }
}
// Test FromWattHoursPerKilogram function
func TestSpecificEnergyFactory_FromWattHoursPerKilogram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattHoursPerKilogram(100)
    if err != nil {
        t.Errorf("FromWattHoursPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyWattHourPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattHoursPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattHoursPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromWattHoursPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromWattHoursPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromWattHoursPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromWattHoursPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromWattHoursPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattHoursPerKilogram(0)
    if err != nil {
        t.Errorf("FromWattHoursPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyWattHourPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattHoursPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromWattDaysPerKilogram function
func TestSpecificEnergyFactory_FromWattDaysPerKilogram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattDaysPerKilogram(100)
    if err != nil {
        t.Errorf("FromWattDaysPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyWattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattDaysPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattDaysPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromWattDaysPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromWattDaysPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromWattDaysPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromWattDaysPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromWattDaysPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattDaysPerKilogram(0)
    if err != nil {
        t.Errorf("FromWattDaysPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyWattDayPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattDaysPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromWattDaysPerTonne function
func TestSpecificEnergyFactory_FromWattDaysPerTonne(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattDaysPerTonne(100)
    if err != nil {
        t.Errorf("FromWattDaysPerTonne() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyWattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattDaysPerTonne() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattDaysPerTonne(math.NaN())
    if err == nil {
        t.Error("FromWattDaysPerTonne() with NaN value should return error")
    }

    _, err = factory.FromWattDaysPerTonne(math.Inf(1))
    if err == nil {
        t.Error("FromWattDaysPerTonne() with +Inf value should return error")
    }

    _, err = factory.FromWattDaysPerTonne(math.Inf(-1))
    if err == nil {
        t.Error("FromWattDaysPerTonne() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattDaysPerTonne(0)
    if err != nil {
        t.Errorf("FromWattDaysPerTonne() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyWattDayPerTonne)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattDaysPerTonne() with zero value = %v, want 0", converted)
    }
}
// Test FromWattDaysPerShortTon function
func TestSpecificEnergyFactory_FromWattDaysPerShortTon(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattDaysPerShortTon(100)
    if err != nil {
        t.Errorf("FromWattDaysPerShortTon() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyWattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattDaysPerShortTon() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattDaysPerShortTon(math.NaN())
    if err == nil {
        t.Error("FromWattDaysPerShortTon() with NaN value should return error")
    }

    _, err = factory.FromWattDaysPerShortTon(math.Inf(1))
    if err == nil {
        t.Error("FromWattDaysPerShortTon() with +Inf value should return error")
    }

    _, err = factory.FromWattDaysPerShortTon(math.Inf(-1))
    if err == nil {
        t.Error("FromWattDaysPerShortTon() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattDaysPerShortTon(0)
    if err != nil {
        t.Errorf("FromWattDaysPerShortTon() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyWattDayPerShortTon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattDaysPerShortTon() with zero value = %v, want 0", converted)
    }
}
// Test FromWattHoursPerPound function
func TestSpecificEnergyFactory_FromWattHoursPerPound(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattHoursPerPound(100)
    if err != nil {
        t.Errorf("FromWattHoursPerPound() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyWattHourPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattHoursPerPound() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattHoursPerPound(math.NaN())
    if err == nil {
        t.Error("FromWattHoursPerPound() with NaN value should return error")
    }

    _, err = factory.FromWattHoursPerPound(math.Inf(1))
    if err == nil {
        t.Error("FromWattHoursPerPound() with +Inf value should return error")
    }

    _, err = factory.FromWattHoursPerPound(math.Inf(-1))
    if err == nil {
        t.Error("FromWattHoursPerPound() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattHoursPerPound(0)
    if err != nil {
        t.Errorf("FromWattHoursPerPound() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyWattHourPerPound)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattHoursPerPound() with zero value = %v, want 0", converted)
    }
}
// Test FromBtuPerPound function
func TestSpecificEnergyFactory_FromBtuPerPound(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBtuPerPound(100)
    if err != nil {
        t.Errorf("FromBtuPerPound() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyBtuPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBtuPerPound() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBtuPerPound(math.NaN())
    if err == nil {
        t.Error("FromBtuPerPound() with NaN value should return error")
    }

    _, err = factory.FromBtuPerPound(math.Inf(1))
    if err == nil {
        t.Error("FromBtuPerPound() with +Inf value should return error")
    }

    _, err = factory.FromBtuPerPound(math.Inf(-1))
    if err == nil {
        t.Error("FromBtuPerPound() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBtuPerPound(0)
    if err != nil {
        t.Errorf("FromBtuPerPound() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyBtuPerPound)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBtuPerPound() with zero value = %v, want 0", converted)
    }
}
// Test FromKilojoulesPerKilogram function
func TestSpecificEnergyFactory_FromKilojoulesPerKilogram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilojoulesPerKilogram(100)
    if err != nil {
        t.Errorf("FromKilojoulesPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyKilojoulePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilojoulesPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilojoulesPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromKilojoulesPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromKilojoulesPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromKilojoulesPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromKilojoulesPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromKilojoulesPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilojoulesPerKilogram(0)
    if err != nil {
        t.Errorf("FromKilojoulesPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyKilojoulePerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilojoulesPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromMegajoulesPerKilogram function
func TestSpecificEnergyFactory_FromMegajoulesPerKilogram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegajoulesPerKilogram(100)
    if err != nil {
        t.Errorf("FromMegajoulesPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyMegajoulePerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegajoulesPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegajoulesPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromMegajoulesPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromMegajoulesPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromMegajoulesPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromMegajoulesPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromMegajoulesPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegajoulesPerKilogram(0)
    if err != nil {
        t.Errorf("FromMegajoulesPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyMegajoulePerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegajoulesPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromKilocaloriesPerGram function
func TestSpecificEnergyFactory_FromKilocaloriesPerGram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilocaloriesPerGram(100)
    if err != nil {
        t.Errorf("FromKilocaloriesPerGram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyKilocaloriePerGram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilocaloriesPerGram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilocaloriesPerGram(math.NaN())
    if err == nil {
        t.Error("FromKilocaloriesPerGram() with NaN value should return error")
    }

    _, err = factory.FromKilocaloriesPerGram(math.Inf(1))
    if err == nil {
        t.Error("FromKilocaloriesPerGram() with +Inf value should return error")
    }

    _, err = factory.FromKilocaloriesPerGram(math.Inf(-1))
    if err == nil {
        t.Error("FromKilocaloriesPerGram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilocaloriesPerGram(0)
    if err != nil {
        t.Errorf("FromKilocaloriesPerGram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyKilocaloriePerGram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilocaloriesPerGram() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattHoursPerKilogram function
func TestSpecificEnergyFactory_FromKilowattHoursPerKilogram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattHoursPerKilogram(100)
    if err != nil {
        t.Errorf("FromKilowattHoursPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyKilowattHourPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattHoursPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattHoursPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromKilowattHoursPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromKilowattHoursPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattHoursPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromKilowattHoursPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattHoursPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattHoursPerKilogram(0)
    if err != nil {
        t.Errorf("FromKilowattHoursPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyKilowattHourPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattHoursPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattHoursPerKilogram function
func TestSpecificEnergyFactory_FromMegawattHoursPerKilogram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattHoursPerKilogram(100)
    if err != nil {
        t.Errorf("FromMegawattHoursPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyMegawattHourPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattHoursPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattHoursPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromMegawattHoursPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromMegawattHoursPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattHoursPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromMegawattHoursPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattHoursPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattHoursPerKilogram(0)
    if err != nil {
        t.Errorf("FromMegawattHoursPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyMegawattHourPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattHoursPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattHoursPerKilogram function
func TestSpecificEnergyFactory_FromGigawattHoursPerKilogram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattHoursPerKilogram(100)
    if err != nil {
        t.Errorf("FromGigawattHoursPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyGigawattHourPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattHoursPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattHoursPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromGigawattHoursPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromGigawattHoursPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattHoursPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromGigawattHoursPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattHoursPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattHoursPerKilogram(0)
    if err != nil {
        t.Errorf("FromGigawattHoursPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyGigawattHourPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattHoursPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattDaysPerKilogram function
func TestSpecificEnergyFactory_FromKilowattDaysPerKilogram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattDaysPerKilogram(100)
    if err != nil {
        t.Errorf("FromKilowattDaysPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyKilowattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattDaysPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattDaysPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromKilowattDaysPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromKilowattDaysPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattDaysPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromKilowattDaysPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattDaysPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattDaysPerKilogram(0)
    if err != nil {
        t.Errorf("FromKilowattDaysPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyKilowattDayPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattDaysPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattDaysPerKilogram function
func TestSpecificEnergyFactory_FromMegawattDaysPerKilogram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattDaysPerKilogram(100)
    if err != nil {
        t.Errorf("FromMegawattDaysPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyMegawattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattDaysPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattDaysPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromMegawattDaysPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromMegawattDaysPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattDaysPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromMegawattDaysPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattDaysPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattDaysPerKilogram(0)
    if err != nil {
        t.Errorf("FromMegawattDaysPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyMegawattDayPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattDaysPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattDaysPerKilogram function
func TestSpecificEnergyFactory_FromGigawattDaysPerKilogram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattDaysPerKilogram(100)
    if err != nil {
        t.Errorf("FromGigawattDaysPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyGigawattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattDaysPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattDaysPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromGigawattDaysPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromGigawattDaysPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattDaysPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromGigawattDaysPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattDaysPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattDaysPerKilogram(0)
    if err != nil {
        t.Errorf("FromGigawattDaysPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyGigawattDayPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattDaysPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromTerawattDaysPerKilogram function
func TestSpecificEnergyFactory_FromTerawattDaysPerKilogram(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerawattDaysPerKilogram(100)
    if err != nil {
        t.Errorf("FromTerawattDaysPerKilogram() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyTerawattDayPerKilogram)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerawattDaysPerKilogram() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerawattDaysPerKilogram(math.NaN())
    if err == nil {
        t.Error("FromTerawattDaysPerKilogram() with NaN value should return error")
    }

    _, err = factory.FromTerawattDaysPerKilogram(math.Inf(1))
    if err == nil {
        t.Error("FromTerawattDaysPerKilogram() with +Inf value should return error")
    }

    _, err = factory.FromTerawattDaysPerKilogram(math.Inf(-1))
    if err == nil {
        t.Error("FromTerawattDaysPerKilogram() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerawattDaysPerKilogram(0)
    if err != nil {
        t.Errorf("FromTerawattDaysPerKilogram() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyTerawattDayPerKilogram)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerawattDaysPerKilogram() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattDaysPerTonne function
func TestSpecificEnergyFactory_FromKilowattDaysPerTonne(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattDaysPerTonne(100)
    if err != nil {
        t.Errorf("FromKilowattDaysPerTonne() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyKilowattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattDaysPerTonne() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattDaysPerTonne(math.NaN())
    if err == nil {
        t.Error("FromKilowattDaysPerTonne() with NaN value should return error")
    }

    _, err = factory.FromKilowattDaysPerTonne(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattDaysPerTonne() with +Inf value should return error")
    }

    _, err = factory.FromKilowattDaysPerTonne(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattDaysPerTonne() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattDaysPerTonne(0)
    if err != nil {
        t.Errorf("FromKilowattDaysPerTonne() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyKilowattDayPerTonne)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattDaysPerTonne() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattDaysPerTonne function
func TestSpecificEnergyFactory_FromMegawattDaysPerTonne(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattDaysPerTonne(100)
    if err != nil {
        t.Errorf("FromMegawattDaysPerTonne() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyMegawattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattDaysPerTonne() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattDaysPerTonne(math.NaN())
    if err == nil {
        t.Error("FromMegawattDaysPerTonne() with NaN value should return error")
    }

    _, err = factory.FromMegawattDaysPerTonne(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattDaysPerTonne() with +Inf value should return error")
    }

    _, err = factory.FromMegawattDaysPerTonne(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattDaysPerTonne() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattDaysPerTonne(0)
    if err != nil {
        t.Errorf("FromMegawattDaysPerTonne() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyMegawattDayPerTonne)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattDaysPerTonne() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattDaysPerTonne function
func TestSpecificEnergyFactory_FromGigawattDaysPerTonne(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattDaysPerTonne(100)
    if err != nil {
        t.Errorf("FromGigawattDaysPerTonne() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyGigawattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattDaysPerTonne() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattDaysPerTonne(math.NaN())
    if err == nil {
        t.Error("FromGigawattDaysPerTonne() with NaN value should return error")
    }

    _, err = factory.FromGigawattDaysPerTonne(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattDaysPerTonne() with +Inf value should return error")
    }

    _, err = factory.FromGigawattDaysPerTonne(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattDaysPerTonne() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattDaysPerTonne(0)
    if err != nil {
        t.Errorf("FromGigawattDaysPerTonne() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyGigawattDayPerTonne)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattDaysPerTonne() with zero value = %v, want 0", converted)
    }
}
// Test FromTerawattDaysPerTonne function
func TestSpecificEnergyFactory_FromTerawattDaysPerTonne(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerawattDaysPerTonne(100)
    if err != nil {
        t.Errorf("FromTerawattDaysPerTonne() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyTerawattDayPerTonne)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerawattDaysPerTonne() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerawattDaysPerTonne(math.NaN())
    if err == nil {
        t.Error("FromTerawattDaysPerTonne() with NaN value should return error")
    }

    _, err = factory.FromTerawattDaysPerTonne(math.Inf(1))
    if err == nil {
        t.Error("FromTerawattDaysPerTonne() with +Inf value should return error")
    }

    _, err = factory.FromTerawattDaysPerTonne(math.Inf(-1))
    if err == nil {
        t.Error("FromTerawattDaysPerTonne() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerawattDaysPerTonne(0)
    if err != nil {
        t.Errorf("FromTerawattDaysPerTonne() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyTerawattDayPerTonne)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerawattDaysPerTonne() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattDaysPerShortTon function
func TestSpecificEnergyFactory_FromKilowattDaysPerShortTon(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattDaysPerShortTon(100)
    if err != nil {
        t.Errorf("FromKilowattDaysPerShortTon() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyKilowattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattDaysPerShortTon() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattDaysPerShortTon(math.NaN())
    if err == nil {
        t.Error("FromKilowattDaysPerShortTon() with NaN value should return error")
    }

    _, err = factory.FromKilowattDaysPerShortTon(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattDaysPerShortTon() with +Inf value should return error")
    }

    _, err = factory.FromKilowattDaysPerShortTon(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattDaysPerShortTon() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattDaysPerShortTon(0)
    if err != nil {
        t.Errorf("FromKilowattDaysPerShortTon() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyKilowattDayPerShortTon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattDaysPerShortTon() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattDaysPerShortTon function
func TestSpecificEnergyFactory_FromMegawattDaysPerShortTon(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattDaysPerShortTon(100)
    if err != nil {
        t.Errorf("FromMegawattDaysPerShortTon() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyMegawattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattDaysPerShortTon() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattDaysPerShortTon(math.NaN())
    if err == nil {
        t.Error("FromMegawattDaysPerShortTon() with NaN value should return error")
    }

    _, err = factory.FromMegawattDaysPerShortTon(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattDaysPerShortTon() with +Inf value should return error")
    }

    _, err = factory.FromMegawattDaysPerShortTon(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattDaysPerShortTon() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattDaysPerShortTon(0)
    if err != nil {
        t.Errorf("FromMegawattDaysPerShortTon() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyMegawattDayPerShortTon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattDaysPerShortTon() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattDaysPerShortTon function
func TestSpecificEnergyFactory_FromGigawattDaysPerShortTon(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattDaysPerShortTon(100)
    if err != nil {
        t.Errorf("FromGigawattDaysPerShortTon() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyGigawattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattDaysPerShortTon() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattDaysPerShortTon(math.NaN())
    if err == nil {
        t.Error("FromGigawattDaysPerShortTon() with NaN value should return error")
    }

    _, err = factory.FromGigawattDaysPerShortTon(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattDaysPerShortTon() with +Inf value should return error")
    }

    _, err = factory.FromGigawattDaysPerShortTon(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattDaysPerShortTon() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattDaysPerShortTon(0)
    if err != nil {
        t.Errorf("FromGigawattDaysPerShortTon() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyGigawattDayPerShortTon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattDaysPerShortTon() with zero value = %v, want 0", converted)
    }
}
// Test FromTerawattDaysPerShortTon function
func TestSpecificEnergyFactory_FromTerawattDaysPerShortTon(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerawattDaysPerShortTon(100)
    if err != nil {
        t.Errorf("FromTerawattDaysPerShortTon() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyTerawattDayPerShortTon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerawattDaysPerShortTon() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerawattDaysPerShortTon(math.NaN())
    if err == nil {
        t.Error("FromTerawattDaysPerShortTon() with NaN value should return error")
    }

    _, err = factory.FromTerawattDaysPerShortTon(math.Inf(1))
    if err == nil {
        t.Error("FromTerawattDaysPerShortTon() with +Inf value should return error")
    }

    _, err = factory.FromTerawattDaysPerShortTon(math.Inf(-1))
    if err == nil {
        t.Error("FromTerawattDaysPerShortTon() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerawattDaysPerShortTon(0)
    if err != nil {
        t.Errorf("FromTerawattDaysPerShortTon() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyTerawattDayPerShortTon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerawattDaysPerShortTon() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattHoursPerPound function
func TestSpecificEnergyFactory_FromKilowattHoursPerPound(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattHoursPerPound(100)
    if err != nil {
        t.Errorf("FromKilowattHoursPerPound() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyKilowattHourPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattHoursPerPound() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattHoursPerPound(math.NaN())
    if err == nil {
        t.Error("FromKilowattHoursPerPound() with NaN value should return error")
    }

    _, err = factory.FromKilowattHoursPerPound(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattHoursPerPound() with +Inf value should return error")
    }

    _, err = factory.FromKilowattHoursPerPound(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattHoursPerPound() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattHoursPerPound(0)
    if err != nil {
        t.Errorf("FromKilowattHoursPerPound() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyKilowattHourPerPound)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattHoursPerPound() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattHoursPerPound function
func TestSpecificEnergyFactory_FromMegawattHoursPerPound(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattHoursPerPound(100)
    if err != nil {
        t.Errorf("FromMegawattHoursPerPound() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyMegawattHourPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattHoursPerPound() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattHoursPerPound(math.NaN())
    if err == nil {
        t.Error("FromMegawattHoursPerPound() with NaN value should return error")
    }

    _, err = factory.FromMegawattHoursPerPound(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattHoursPerPound() with +Inf value should return error")
    }

    _, err = factory.FromMegawattHoursPerPound(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattHoursPerPound() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattHoursPerPound(0)
    if err != nil {
        t.Errorf("FromMegawattHoursPerPound() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyMegawattHourPerPound)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattHoursPerPound() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattHoursPerPound function
func TestSpecificEnergyFactory_FromGigawattHoursPerPound(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattHoursPerPound(100)
    if err != nil {
        t.Errorf("FromGigawattHoursPerPound() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.SpecificEnergyGigawattHourPerPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattHoursPerPound() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattHoursPerPound(math.NaN())
    if err == nil {
        t.Error("FromGigawattHoursPerPound() with NaN value should return error")
    }

    _, err = factory.FromGigawattHoursPerPound(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattHoursPerPound() with +Inf value should return error")
    }

    _, err = factory.FromGigawattHoursPerPound(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattHoursPerPound() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattHoursPerPound(0)
    if err != nil {
        t.Errorf("FromGigawattHoursPerPound() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.SpecificEnergyGigawattHourPerPound)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattHoursPerPound() with zero value = %v, want 0", converted)
    }
}

func TestSpecificEnergyToString(t *testing.T) {
	factory := units.SpecificEnergyFactory{}
	a, err := factory.CreateSpecificEnergy(45, units.SpecificEnergyJoulePerKilogram)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.SpecificEnergyJoulePerKilogram, 2)
	expected := "45.00 " + units.GetSpecificEnergyAbbreviation(units.SpecificEnergyJoulePerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.SpecificEnergyJoulePerKilogram, -1)
	expected = "45 " + units.GetSpecificEnergyAbbreviation(units.SpecificEnergyJoulePerKilogram)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestSpecificEnergy_EqualityAndComparison(t *testing.T) {
	factory := units.SpecificEnergyFactory{}
	a1, _ := factory.CreateSpecificEnergy(60, units.SpecificEnergyJoulePerKilogram)
	a2, _ := factory.CreateSpecificEnergy(60, units.SpecificEnergyJoulePerKilogram)
	a3, _ := factory.CreateSpecificEnergy(120, units.SpecificEnergyJoulePerKilogram)

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

func TestSpecificEnergy_Arithmetic(t *testing.T) {
	factory := units.SpecificEnergyFactory{}
	a1, _ := factory.CreateSpecificEnergy(30, units.SpecificEnergyJoulePerKilogram)
	a2, _ := factory.CreateSpecificEnergy(45, units.SpecificEnergyJoulePerKilogram)

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


func TestGetSpecificEnergyAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.SpecificEnergyUnits
        want string
    }{
        {
            name: "JoulePerKilogram abbreviation",
            unit: units.SpecificEnergyJoulePerKilogram,
            want: "J/kg",
        },
        {
            name: "MegaJoulePerTonne abbreviation",
            unit: units.SpecificEnergyMegaJoulePerTonne,
            want: "MJ/t",
        },
        {
            name: "CaloriePerGram abbreviation",
            unit: units.SpecificEnergyCaloriePerGram,
            want: "cal/g",
        },
        {
            name: "WattHourPerKilogram abbreviation",
            unit: units.SpecificEnergyWattHourPerKilogram,
            want: "Wh/kg",
        },
        {
            name: "WattDayPerKilogram abbreviation",
            unit: units.SpecificEnergyWattDayPerKilogram,
            want: "Wd/kg",
        },
        {
            name: "WattDayPerTonne abbreviation",
            unit: units.SpecificEnergyWattDayPerTonne,
            want: "Wd/t",
        },
        {
            name: "WattDayPerShortTon abbreviation",
            unit: units.SpecificEnergyWattDayPerShortTon,
            want: "Wd/ST",
        },
        {
            name: "WattHourPerPound abbreviation",
            unit: units.SpecificEnergyWattHourPerPound,
            want: "Wh/lbs",
        },
        {
            name: "BtuPerPound abbreviation",
            unit: units.SpecificEnergyBtuPerPound,
            want: "btu/lb",
        },
        {
            name: "KilojoulePerKilogram abbreviation",
            unit: units.SpecificEnergyKilojoulePerKilogram,
            want: "kJ/kg",
        },
        {
            name: "MegajoulePerKilogram abbreviation",
            unit: units.SpecificEnergyMegajoulePerKilogram,
            want: "MJ/kg",
        },
        {
            name: "KilocaloriePerGram abbreviation",
            unit: units.SpecificEnergyKilocaloriePerGram,
            want: "kcal/g",
        },
        {
            name: "KilowattHourPerKilogram abbreviation",
            unit: units.SpecificEnergyKilowattHourPerKilogram,
            want: "kWh/kg",
        },
        {
            name: "MegawattHourPerKilogram abbreviation",
            unit: units.SpecificEnergyMegawattHourPerKilogram,
            want: "MWh/kg",
        },
        {
            name: "GigawattHourPerKilogram abbreviation",
            unit: units.SpecificEnergyGigawattHourPerKilogram,
            want: "GWh/kg",
        },
        {
            name: "KilowattDayPerKilogram abbreviation",
            unit: units.SpecificEnergyKilowattDayPerKilogram,
            want: "kWd/kg",
        },
        {
            name: "MegawattDayPerKilogram abbreviation",
            unit: units.SpecificEnergyMegawattDayPerKilogram,
            want: "MWd/kg",
        },
        {
            name: "GigawattDayPerKilogram abbreviation",
            unit: units.SpecificEnergyGigawattDayPerKilogram,
            want: "GWd/kg",
        },
        {
            name: "TerawattDayPerKilogram abbreviation",
            unit: units.SpecificEnergyTerawattDayPerKilogram,
            want: "TWd/kg",
        },
        {
            name: "KilowattDayPerTonne abbreviation",
            unit: units.SpecificEnergyKilowattDayPerTonne,
            want: "kWd/t",
        },
        {
            name: "MegawattDayPerTonne abbreviation",
            unit: units.SpecificEnergyMegawattDayPerTonne,
            want: "MWd/t",
        },
        {
            name: "GigawattDayPerTonne abbreviation",
            unit: units.SpecificEnergyGigawattDayPerTonne,
            want: "GWd/t",
        },
        {
            name: "TerawattDayPerTonne abbreviation",
            unit: units.SpecificEnergyTerawattDayPerTonne,
            want: "TWd/t",
        },
        {
            name: "KilowattDayPerShortTon abbreviation",
            unit: units.SpecificEnergyKilowattDayPerShortTon,
            want: "kWd/ST",
        },
        {
            name: "MegawattDayPerShortTon abbreviation",
            unit: units.SpecificEnergyMegawattDayPerShortTon,
            want: "MWd/ST",
        },
        {
            name: "GigawattDayPerShortTon abbreviation",
            unit: units.SpecificEnergyGigawattDayPerShortTon,
            want: "GWd/ST",
        },
        {
            name: "TerawattDayPerShortTon abbreviation",
            unit: units.SpecificEnergyTerawattDayPerShortTon,
            want: "TWd/ST",
        },
        {
            name: "KilowattHourPerPound abbreviation",
            unit: units.SpecificEnergyKilowattHourPerPound,
            want: "kWh/lbs",
        },
        {
            name: "MegawattHourPerPound abbreviation",
            unit: units.SpecificEnergyMegawattHourPerPound,
            want: "MWh/lbs",
        },
        {
            name: "GigawattHourPerPound abbreviation",
            unit: units.SpecificEnergyGigawattHourPerPound,
            want: "GWh/lbs",
        },
        {
            name: "invalid unit",
            unit: units.SpecificEnergyUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetSpecificEnergyAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetSpecificEnergyAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestSpecificEnergy_String(t *testing.T) {
    factory := units.SpecificEnergyFactory{}
    
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
            unit, err := factory.CreateSpecificEnergy(tt.value, units.SpecificEnergyJoulePerKilogram)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("SpecificEnergy.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestSpecificEnergy_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.SpecificEnergyFactory{}

	_, err := uf.CreateSpecificEnergy(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}