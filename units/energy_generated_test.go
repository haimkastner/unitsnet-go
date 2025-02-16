// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestEnergyDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Joule"}`
	
	factory := units.EnergyDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.EnergyJoule {
		t.Errorf("expected unit %v, got %v", units.EnergyJoule, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Joule"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestEnergyDto_ToJSON(t *testing.T) {
	dto := units.EnergyDto{
		Value: 45,
		Unit:  units.EnergyJoule,
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
	if result["unit"].(string) != string(units.EnergyJoule) {
		t.Errorf("expected unit %s, got %v", units.EnergyJoule, result["unit"])
	}
}

func TestNewEnergy_InvalidValue(t *testing.T) {
	factory := units.EnergyFactory{}
	// NaN value should return an error.
	_, err := factory.CreateEnergy(math.NaN(), units.EnergyJoule)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateEnergy(math.Inf(1), units.EnergyJoule)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestEnergyConversions(t *testing.T) {
	factory := units.EnergyFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateEnergy(180, units.EnergyJoule)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Joules.
		// No expected conversion value provided for Joules, verifying result is not NaN.
		result := a.Joules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Joules returned NaN")
		}
	}
	{
		// Test conversion to Calories.
		// No expected conversion value provided for Calories, verifying result is not NaN.
		result := a.Calories()
		if math.IsNaN(result) {
			t.Errorf("conversion to Calories returned NaN")
		}
	}
	{
		// Test conversion to BritishThermalUnits.
		// No expected conversion value provided for BritishThermalUnits, verifying result is not NaN.
		result := a.BritishThermalUnits()
		if math.IsNaN(result) {
			t.Errorf("conversion to BritishThermalUnits returned NaN")
		}
	}
	{
		// Test conversion to ElectronVolts.
		// No expected conversion value provided for ElectronVolts, verifying result is not NaN.
		result := a.ElectronVolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to ElectronVolts returned NaN")
		}
	}
	{
		// Test conversion to FootPounds.
		// No expected conversion value provided for FootPounds, verifying result is not NaN.
		result := a.FootPounds()
		if math.IsNaN(result) {
			t.Errorf("conversion to FootPounds returned NaN")
		}
	}
	{
		// Test conversion to Ergs.
		// No expected conversion value provided for Ergs, verifying result is not NaN.
		result := a.Ergs()
		if math.IsNaN(result) {
			t.Errorf("conversion to Ergs returned NaN")
		}
	}
	{
		// Test conversion to WattHours.
		// No expected conversion value provided for WattHours, verifying result is not NaN.
		result := a.WattHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattHours returned NaN")
		}
	}
	{
		// Test conversion to WattDays.
		// No expected conversion value provided for WattDays, verifying result is not NaN.
		result := a.WattDays()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattDays returned NaN")
		}
	}
	{
		// Test conversion to ThermsEc.
		// No expected conversion value provided for ThermsEc, verifying result is not NaN.
		result := a.ThermsEc()
		if math.IsNaN(result) {
			t.Errorf("conversion to ThermsEc returned NaN")
		}
	}
	{
		// Test conversion to ThermsUs.
		// No expected conversion value provided for ThermsUs, verifying result is not NaN.
		result := a.ThermsUs()
		if math.IsNaN(result) {
			t.Errorf("conversion to ThermsUs returned NaN")
		}
	}
	{
		// Test conversion to ThermsImperial.
		// No expected conversion value provided for ThermsImperial, verifying result is not NaN.
		result := a.ThermsImperial()
		if math.IsNaN(result) {
			t.Errorf("conversion to ThermsImperial returned NaN")
		}
	}
	{
		// Test conversion to HorsepowerHours.
		// No expected conversion value provided for HorsepowerHours, verifying result is not NaN.
		result := a.HorsepowerHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to HorsepowerHours returned NaN")
		}
	}
	{
		// Test conversion to Nanojoules.
		// No expected conversion value provided for Nanojoules, verifying result is not NaN.
		result := a.Nanojoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanojoules returned NaN")
		}
	}
	{
		// Test conversion to Microjoules.
		// No expected conversion value provided for Microjoules, verifying result is not NaN.
		result := a.Microjoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microjoules returned NaN")
		}
	}
	{
		// Test conversion to Millijoules.
		// No expected conversion value provided for Millijoules, verifying result is not NaN.
		result := a.Millijoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millijoules returned NaN")
		}
	}
	{
		// Test conversion to Kilojoules.
		// No expected conversion value provided for Kilojoules, verifying result is not NaN.
		result := a.Kilojoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilojoules returned NaN")
		}
	}
	{
		// Test conversion to Megajoules.
		// No expected conversion value provided for Megajoules, verifying result is not NaN.
		result := a.Megajoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megajoules returned NaN")
		}
	}
	{
		// Test conversion to Gigajoules.
		// No expected conversion value provided for Gigajoules, verifying result is not NaN.
		result := a.Gigajoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigajoules returned NaN")
		}
	}
	{
		// Test conversion to Terajoules.
		// No expected conversion value provided for Terajoules, verifying result is not NaN.
		result := a.Terajoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Terajoules returned NaN")
		}
	}
	{
		// Test conversion to Petajoules.
		// No expected conversion value provided for Petajoules, verifying result is not NaN.
		result := a.Petajoules()
		if math.IsNaN(result) {
			t.Errorf("conversion to Petajoules returned NaN")
		}
	}
	{
		// Test conversion to Kilocalories.
		// No expected conversion value provided for Kilocalories, verifying result is not NaN.
		result := a.Kilocalories()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilocalories returned NaN")
		}
	}
	{
		// Test conversion to Megacalories.
		// No expected conversion value provided for Megacalories, verifying result is not NaN.
		result := a.Megacalories()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megacalories returned NaN")
		}
	}
	{
		// Test conversion to KilobritishThermalUnits.
		// No expected conversion value provided for KilobritishThermalUnits, verifying result is not NaN.
		result := a.KilobritishThermalUnits()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilobritishThermalUnits returned NaN")
		}
	}
	{
		// Test conversion to MegabritishThermalUnits.
		// No expected conversion value provided for MegabritishThermalUnits, verifying result is not NaN.
		result := a.MegabritishThermalUnits()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegabritishThermalUnits returned NaN")
		}
	}
	{
		// Test conversion to GigabritishThermalUnits.
		// No expected conversion value provided for GigabritishThermalUnits, verifying result is not NaN.
		result := a.GigabritishThermalUnits()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigabritishThermalUnits returned NaN")
		}
	}
	{
		// Test conversion to KiloelectronVolts.
		// No expected conversion value provided for KiloelectronVolts, verifying result is not NaN.
		result := a.KiloelectronVolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to KiloelectronVolts returned NaN")
		}
	}
	{
		// Test conversion to MegaelectronVolts.
		// No expected conversion value provided for MegaelectronVolts, verifying result is not NaN.
		result := a.MegaelectronVolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegaelectronVolts returned NaN")
		}
	}
	{
		// Test conversion to GigaelectronVolts.
		// No expected conversion value provided for GigaelectronVolts, verifying result is not NaN.
		result := a.GigaelectronVolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigaelectronVolts returned NaN")
		}
	}
	{
		// Test conversion to TeraelectronVolts.
		// No expected conversion value provided for TeraelectronVolts, verifying result is not NaN.
		result := a.TeraelectronVolts()
		if math.IsNaN(result) {
			t.Errorf("conversion to TeraelectronVolts returned NaN")
		}
	}
	{
		// Test conversion to KilowattHours.
		// No expected conversion value provided for KilowattHours, verifying result is not NaN.
		result := a.KilowattHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattHours returned NaN")
		}
	}
	{
		// Test conversion to MegawattHours.
		// No expected conversion value provided for MegawattHours, verifying result is not NaN.
		result := a.MegawattHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattHours returned NaN")
		}
	}
	{
		// Test conversion to GigawattHours.
		// No expected conversion value provided for GigawattHours, verifying result is not NaN.
		result := a.GigawattHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattHours returned NaN")
		}
	}
	{
		// Test conversion to TerawattHours.
		// No expected conversion value provided for TerawattHours, verifying result is not NaN.
		result := a.TerawattHours()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerawattHours returned NaN")
		}
	}
	{
		// Test conversion to KilowattDays.
		// No expected conversion value provided for KilowattDays, verifying result is not NaN.
		result := a.KilowattDays()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattDays returned NaN")
		}
	}
	{
		// Test conversion to MegawattDays.
		// No expected conversion value provided for MegawattDays, verifying result is not NaN.
		result := a.MegawattDays()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattDays returned NaN")
		}
	}
	{
		// Test conversion to GigawattDays.
		// No expected conversion value provided for GigawattDays, verifying result is not NaN.
		result := a.GigawattDays()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigawattDays returned NaN")
		}
	}
	{
		// Test conversion to TerawattDays.
		// No expected conversion value provided for TerawattDays, verifying result is not NaN.
		result := a.TerawattDays()
		if math.IsNaN(result) {
			t.Errorf("conversion to TerawattDays returned NaN")
		}
	}
	{
		// Test conversion to DecathermsEc.
		// No expected conversion value provided for DecathermsEc, verifying result is not NaN.
		result := a.DecathermsEc()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecathermsEc returned NaN")
		}
	}
	{
		// Test conversion to DecathermsUs.
		// No expected conversion value provided for DecathermsUs, verifying result is not NaN.
		result := a.DecathermsUs()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecathermsUs returned NaN")
		}
	}
	{
		// Test conversion to DecathermsImperial.
		// No expected conversion value provided for DecathermsImperial, verifying result is not NaN.
		result := a.DecathermsImperial()
		if math.IsNaN(result) {
			t.Errorf("conversion to DecathermsImperial returned NaN")
		}
	}
}

func TestEnergy_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.EnergyFactory{}
	a, err := factory.CreateEnergy(90, units.EnergyJoule)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.EnergyJoule {
		t.Errorf("expected default unit Joule, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.EnergyJoule
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.EnergyDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.EnergyJoule {
		t.Errorf("expected unit Joule, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestEnergyFactory_FromDto(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyJoule,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.EnergyDto{
        Value: math.NaN(),
        Unit:  units.EnergyJoule,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Joule conversion
    joulesDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyJoule,
    }
    
    var joulesResult *units.Energy
    joulesResult, err = factory.FromDto(joulesDto)
    if err != nil {
        t.Errorf("FromDto() with Joule returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joulesResult.Convert(units.EnergyJoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Joule = %v, want %v", converted, 100)
    }
    // Test Calorie conversion
    caloriesDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyCalorie,
    }
    
    var caloriesResult *units.Energy
    caloriesResult, err = factory.FromDto(caloriesDto)
    if err != nil {
        t.Errorf("FromDto() with Calorie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = caloriesResult.Convert(units.EnergyCalorie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Calorie = %v, want %v", converted, 100)
    }
    // Test BritishThermalUnit conversion
    british_thermal_unitsDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyBritishThermalUnit,
    }
    
    var british_thermal_unitsResult *units.Energy
    british_thermal_unitsResult, err = factory.FromDto(british_thermal_unitsDto)
    if err != nil {
        t.Errorf("FromDto() with BritishThermalUnit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = british_thermal_unitsResult.Convert(units.EnergyBritishThermalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BritishThermalUnit = %v, want %v", converted, 100)
    }
    // Test ElectronVolt conversion
    electron_voltsDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyElectronVolt,
    }
    
    var electron_voltsResult *units.Energy
    electron_voltsResult, err = factory.FromDto(electron_voltsDto)
    if err != nil {
        t.Errorf("FromDto() with ElectronVolt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = electron_voltsResult.Convert(units.EnergyElectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ElectronVolt = %v, want %v", converted, 100)
    }
    // Test FootPound conversion
    foot_poundsDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyFootPound,
    }
    
    var foot_poundsResult *units.Energy
    foot_poundsResult, err = factory.FromDto(foot_poundsDto)
    if err != nil {
        t.Errorf("FromDto() with FootPound returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = foot_poundsResult.Convert(units.EnergyFootPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootPound = %v, want %v", converted, 100)
    }
    // Test Erg conversion
    ergsDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyErg,
    }
    
    var ergsResult *units.Energy
    ergsResult, err = factory.FromDto(ergsDto)
    if err != nil {
        t.Errorf("FromDto() with Erg returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ergsResult.Convert(units.EnergyErg)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Erg = %v, want %v", converted, 100)
    }
    // Test WattHour conversion
    watt_hoursDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyWattHour,
    }
    
    var watt_hoursResult *units.Energy
    watt_hoursResult, err = factory.FromDto(watt_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with WattHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_hoursResult.Convert(units.EnergyWattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattHour = %v, want %v", converted, 100)
    }
    // Test WattDay conversion
    watt_daysDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyWattDay,
    }
    
    var watt_daysResult *units.Energy
    watt_daysResult, err = factory.FromDto(watt_daysDto)
    if err != nil {
        t.Errorf("FromDto() with WattDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_daysResult.Convert(units.EnergyWattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattDay = %v, want %v", converted, 100)
    }
    // Test ThermEc conversion
    therms_ecDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyThermEc,
    }
    
    var therms_ecResult *units.Energy
    therms_ecResult, err = factory.FromDto(therms_ecDto)
    if err != nil {
        t.Errorf("FromDto() with ThermEc returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = therms_ecResult.Convert(units.EnergyThermEc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ThermEc = %v, want %v", converted, 100)
    }
    // Test ThermUs conversion
    therms_usDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyThermUs,
    }
    
    var therms_usResult *units.Energy
    therms_usResult, err = factory.FromDto(therms_usDto)
    if err != nil {
        t.Errorf("FromDto() with ThermUs returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = therms_usResult.Convert(units.EnergyThermUs)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ThermUs = %v, want %v", converted, 100)
    }
    // Test ThermImperial conversion
    therms_imperialDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyThermImperial,
    }
    
    var therms_imperialResult *units.Energy
    therms_imperialResult, err = factory.FromDto(therms_imperialDto)
    if err != nil {
        t.Errorf("FromDto() with ThermImperial returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = therms_imperialResult.Convert(units.EnergyThermImperial)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ThermImperial = %v, want %v", converted, 100)
    }
    // Test HorsepowerHour conversion
    horsepower_hoursDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyHorsepowerHour,
    }
    
    var horsepower_hoursResult *units.Energy
    horsepower_hoursResult, err = factory.FromDto(horsepower_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with HorsepowerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = horsepower_hoursResult.Convert(units.EnergyHorsepowerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HorsepowerHour = %v, want %v", converted, 100)
    }
    // Test Nanojoule conversion
    nanojoulesDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyNanojoule,
    }
    
    var nanojoulesResult *units.Energy
    nanojoulesResult, err = factory.FromDto(nanojoulesDto)
    if err != nil {
        t.Errorf("FromDto() with Nanojoule returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanojoulesResult.Convert(units.EnergyNanojoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanojoule = %v, want %v", converted, 100)
    }
    // Test Microjoule conversion
    microjoulesDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyMicrojoule,
    }
    
    var microjoulesResult *units.Energy
    microjoulesResult, err = factory.FromDto(microjoulesDto)
    if err != nil {
        t.Errorf("FromDto() with Microjoule returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microjoulesResult.Convert(units.EnergyMicrojoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microjoule = %v, want %v", converted, 100)
    }
    // Test Millijoule conversion
    millijoulesDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyMillijoule,
    }
    
    var millijoulesResult *units.Energy
    millijoulesResult, err = factory.FromDto(millijoulesDto)
    if err != nil {
        t.Errorf("FromDto() with Millijoule returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millijoulesResult.Convert(units.EnergyMillijoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millijoule = %v, want %v", converted, 100)
    }
    // Test Kilojoule conversion
    kilojoulesDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyKilojoule,
    }
    
    var kilojoulesResult *units.Energy
    kilojoulesResult, err = factory.FromDto(kilojoulesDto)
    if err != nil {
        t.Errorf("FromDto() with Kilojoule returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoulesResult.Convert(units.EnergyKilojoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilojoule = %v, want %v", converted, 100)
    }
    // Test Megajoule conversion
    megajoulesDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyMegajoule,
    }
    
    var megajoulesResult *units.Energy
    megajoulesResult, err = factory.FromDto(megajoulesDto)
    if err != nil {
        t.Errorf("FromDto() with Megajoule returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoulesResult.Convert(units.EnergyMegajoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megajoule = %v, want %v", converted, 100)
    }
    // Test Gigajoule conversion
    gigajoulesDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyGigajoule,
    }
    
    var gigajoulesResult *units.Energy
    gigajoulesResult, err = factory.FromDto(gigajoulesDto)
    if err != nil {
        t.Errorf("FromDto() with Gigajoule returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigajoulesResult.Convert(units.EnergyGigajoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigajoule = %v, want %v", converted, 100)
    }
    // Test Terajoule conversion
    terajoulesDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyTerajoule,
    }
    
    var terajoulesResult *units.Energy
    terajoulesResult, err = factory.FromDto(terajoulesDto)
    if err != nil {
        t.Errorf("FromDto() with Terajoule returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terajoulesResult.Convert(units.EnergyTerajoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terajoule = %v, want %v", converted, 100)
    }
    // Test Petajoule conversion
    petajoulesDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyPetajoule,
    }
    
    var petajoulesResult *units.Energy
    petajoulesResult, err = factory.FromDto(petajoulesDto)
    if err != nil {
        t.Errorf("FromDto() with Petajoule returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petajoulesResult.Convert(units.EnergyPetajoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petajoule = %v, want %v", converted, 100)
    }
    // Test Kilocalorie conversion
    kilocaloriesDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyKilocalorie,
    }
    
    var kilocaloriesResult *units.Energy
    kilocaloriesResult, err = factory.FromDto(kilocaloriesDto)
    if err != nil {
        t.Errorf("FromDto() with Kilocalorie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocaloriesResult.Convert(units.EnergyKilocalorie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilocalorie = %v, want %v", converted, 100)
    }
    // Test Megacalorie conversion
    megacaloriesDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyMegacalorie,
    }
    
    var megacaloriesResult *units.Energy
    megacaloriesResult, err = factory.FromDto(megacaloriesDto)
    if err != nil {
        t.Errorf("FromDto() with Megacalorie returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megacaloriesResult.Convert(units.EnergyMegacalorie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megacalorie = %v, want %v", converted, 100)
    }
    // Test KilobritishThermalUnit conversion
    kilobritish_thermal_unitsDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyKilobritishThermalUnit,
    }
    
    var kilobritish_thermal_unitsResult *units.Energy
    kilobritish_thermal_unitsResult, err = factory.FromDto(kilobritish_thermal_unitsDto)
    if err != nil {
        t.Errorf("FromDto() with KilobritishThermalUnit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobritish_thermal_unitsResult.Convert(units.EnergyKilobritishThermalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilobritishThermalUnit = %v, want %v", converted, 100)
    }
    // Test MegabritishThermalUnit conversion
    megabritish_thermal_unitsDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyMegabritishThermalUnit,
    }
    
    var megabritish_thermal_unitsResult *units.Energy
    megabritish_thermal_unitsResult, err = factory.FromDto(megabritish_thermal_unitsDto)
    if err != nil {
        t.Errorf("FromDto() with MegabritishThermalUnit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabritish_thermal_unitsResult.Convert(units.EnergyMegabritishThermalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegabritishThermalUnit = %v, want %v", converted, 100)
    }
    // Test GigabritishThermalUnit conversion
    gigabritish_thermal_unitsDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyGigabritishThermalUnit,
    }
    
    var gigabritish_thermal_unitsResult *units.Energy
    gigabritish_thermal_unitsResult, err = factory.FromDto(gigabritish_thermal_unitsDto)
    if err != nil {
        t.Errorf("FromDto() with GigabritishThermalUnit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigabritish_thermal_unitsResult.Convert(units.EnergyGigabritishThermalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigabritishThermalUnit = %v, want %v", converted, 100)
    }
    // Test KiloelectronVolt conversion
    kiloelectron_voltsDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyKiloelectronVolt,
    }
    
    var kiloelectron_voltsResult *units.Energy
    kiloelectron_voltsResult, err = factory.FromDto(kiloelectron_voltsDto)
    if err != nil {
        t.Errorf("FromDto() with KiloelectronVolt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloelectron_voltsResult.Convert(units.EnergyKiloelectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloelectronVolt = %v, want %v", converted, 100)
    }
    // Test MegaelectronVolt conversion
    megaelectron_voltsDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyMegaelectronVolt,
    }
    
    var megaelectron_voltsResult *units.Energy
    megaelectron_voltsResult, err = factory.FromDto(megaelectron_voltsDto)
    if err != nil {
        t.Errorf("FromDto() with MegaelectronVolt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaelectron_voltsResult.Convert(units.EnergyMegaelectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaelectronVolt = %v, want %v", converted, 100)
    }
    // Test GigaelectronVolt conversion
    gigaelectron_voltsDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyGigaelectronVolt,
    }
    
    var gigaelectron_voltsResult *units.Energy
    gigaelectron_voltsResult, err = factory.FromDto(gigaelectron_voltsDto)
    if err != nil {
        t.Errorf("FromDto() with GigaelectronVolt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigaelectron_voltsResult.Convert(units.EnergyGigaelectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigaelectronVolt = %v, want %v", converted, 100)
    }
    // Test TeraelectronVolt conversion
    teraelectron_voltsDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyTeraelectronVolt,
    }
    
    var teraelectron_voltsResult *units.Energy
    teraelectron_voltsResult, err = factory.FromDto(teraelectron_voltsDto)
    if err != nil {
        t.Errorf("FromDto() with TeraelectronVolt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teraelectron_voltsResult.Convert(units.EnergyTeraelectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TeraelectronVolt = %v, want %v", converted, 100)
    }
    // Test KilowattHour conversion
    kilowatt_hoursDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyKilowattHour,
    }
    
    var kilowatt_hoursResult *units.Energy
    kilowatt_hoursResult, err = factory.FromDto(kilowatt_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_hoursResult.Convert(units.EnergyKilowattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattHour = %v, want %v", converted, 100)
    }
    // Test MegawattHour conversion
    megawatt_hoursDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyMegawattHour,
    }
    
    var megawatt_hoursResult *units.Energy
    megawatt_hoursResult, err = factory.FromDto(megawatt_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_hoursResult.Convert(units.EnergyMegawattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattHour = %v, want %v", converted, 100)
    }
    // Test GigawattHour conversion
    gigawatt_hoursDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyGigawattHour,
    }
    
    var gigawatt_hoursResult *units.Energy
    gigawatt_hoursResult, err = factory.FromDto(gigawatt_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_hoursResult.Convert(units.EnergyGigawattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattHour = %v, want %v", converted, 100)
    }
    // Test TerawattHour conversion
    terawatt_hoursDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyTerawattHour,
    }
    
    var terawatt_hoursResult *units.Energy
    terawatt_hoursResult, err = factory.FromDto(terawatt_hoursDto)
    if err != nil {
        t.Errorf("FromDto() with TerawattHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatt_hoursResult.Convert(units.EnergyTerawattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattHour = %v, want %v", converted, 100)
    }
    // Test KilowattDay conversion
    kilowatt_daysDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyKilowattDay,
    }
    
    var kilowatt_daysResult *units.Energy
    kilowatt_daysResult, err = factory.FromDto(kilowatt_daysDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_daysResult.Convert(units.EnergyKilowattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattDay = %v, want %v", converted, 100)
    }
    // Test MegawattDay conversion
    megawatt_daysDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyMegawattDay,
    }
    
    var megawatt_daysResult *units.Energy
    megawatt_daysResult, err = factory.FromDto(megawatt_daysDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_daysResult.Convert(units.EnergyMegawattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattDay = %v, want %v", converted, 100)
    }
    // Test GigawattDay conversion
    gigawatt_daysDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyGigawattDay,
    }
    
    var gigawatt_daysResult *units.Energy
    gigawatt_daysResult, err = factory.FromDto(gigawatt_daysDto)
    if err != nil {
        t.Errorf("FromDto() with GigawattDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_daysResult.Convert(units.EnergyGigawattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattDay = %v, want %v", converted, 100)
    }
    // Test TerawattDay conversion
    terawatt_daysDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyTerawattDay,
    }
    
    var terawatt_daysResult *units.Energy
    terawatt_daysResult, err = factory.FromDto(terawatt_daysDto)
    if err != nil {
        t.Errorf("FromDto() with TerawattDay returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatt_daysResult.Convert(units.EnergyTerawattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattDay = %v, want %v", converted, 100)
    }
    // Test DecathermEc conversion
    decatherms_ecDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyDecathermEc,
    }
    
    var decatherms_ecResult *units.Energy
    decatherms_ecResult, err = factory.FromDto(decatherms_ecDto)
    if err != nil {
        t.Errorf("FromDto() with DecathermEc returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decatherms_ecResult.Convert(units.EnergyDecathermEc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecathermEc = %v, want %v", converted, 100)
    }
    // Test DecathermUs conversion
    decatherms_usDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyDecathermUs,
    }
    
    var decatherms_usResult *units.Energy
    decatherms_usResult, err = factory.FromDto(decatherms_usDto)
    if err != nil {
        t.Errorf("FromDto() with DecathermUs returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decatherms_usResult.Convert(units.EnergyDecathermUs)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecathermUs = %v, want %v", converted, 100)
    }
    // Test DecathermImperial conversion
    decatherms_imperialDto := units.EnergyDto{
        Value: 100,
        Unit:  units.EnergyDecathermImperial,
    }
    
    var decatherms_imperialResult *units.Energy
    decatherms_imperialResult, err = factory.FromDto(decatherms_imperialDto)
    if err != nil {
        t.Errorf("FromDto() with DecathermImperial returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decatherms_imperialResult.Convert(units.EnergyDecathermImperial)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecathermImperial = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.EnergyDto{
        Value: 0,
        Unit:  units.EnergyJoule,
    }
    
    var zeroResult *units.Energy
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestEnergyFactory_FromDtoJSON(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Joule"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Joule"}`)
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
    nanJSON, _ := json.Marshal(units.EnergyDto{
        Value: nanValue,
        Unit:  units.EnergyJoule,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Joule unit
    joulesJSON := []byte(`{"value": 100, "unit": "Joule"}`)
    joulesResult, err := factory.FromDtoJSON(joulesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Joule unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joulesResult.Convert(units.EnergyJoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Joule = %v, want %v", converted, 100)
    }
    // Test JSON with Calorie unit
    caloriesJSON := []byte(`{"value": 100, "unit": "Calorie"}`)
    caloriesResult, err := factory.FromDtoJSON(caloriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Calorie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = caloriesResult.Convert(units.EnergyCalorie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Calorie = %v, want %v", converted, 100)
    }
    // Test JSON with BritishThermalUnit unit
    british_thermal_unitsJSON := []byte(`{"value": 100, "unit": "BritishThermalUnit"}`)
    british_thermal_unitsResult, err := factory.FromDtoJSON(british_thermal_unitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BritishThermalUnit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = british_thermal_unitsResult.Convert(units.EnergyBritishThermalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BritishThermalUnit = %v, want %v", converted, 100)
    }
    // Test JSON with ElectronVolt unit
    electron_voltsJSON := []byte(`{"value": 100, "unit": "ElectronVolt"}`)
    electron_voltsResult, err := factory.FromDtoJSON(electron_voltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ElectronVolt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = electron_voltsResult.Convert(units.EnergyElectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ElectronVolt = %v, want %v", converted, 100)
    }
    // Test JSON with FootPound unit
    foot_poundsJSON := []byte(`{"value": 100, "unit": "FootPound"}`)
    foot_poundsResult, err := factory.FromDtoJSON(foot_poundsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FootPound unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = foot_poundsResult.Convert(units.EnergyFootPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootPound = %v, want %v", converted, 100)
    }
    // Test JSON with Erg unit
    ergsJSON := []byte(`{"value": 100, "unit": "Erg"}`)
    ergsResult, err := factory.FromDtoJSON(ergsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Erg unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = ergsResult.Convert(units.EnergyErg)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Erg = %v, want %v", converted, 100)
    }
    // Test JSON with WattHour unit
    watt_hoursJSON := []byte(`{"value": 100, "unit": "WattHour"}`)
    watt_hoursResult, err := factory.FromDtoJSON(watt_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_hoursResult.Convert(units.EnergyWattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattHour = %v, want %v", converted, 100)
    }
    // Test JSON with WattDay unit
    watt_daysJSON := []byte(`{"value": 100, "unit": "WattDay"}`)
    watt_daysResult, err := factory.FromDtoJSON(watt_daysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watt_daysResult.Convert(units.EnergyWattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattDay = %v, want %v", converted, 100)
    }
    // Test JSON with ThermEc unit
    therms_ecJSON := []byte(`{"value": 100, "unit": "ThermEc"}`)
    therms_ecResult, err := factory.FromDtoJSON(therms_ecJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ThermEc unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = therms_ecResult.Convert(units.EnergyThermEc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ThermEc = %v, want %v", converted, 100)
    }
    // Test JSON with ThermUs unit
    therms_usJSON := []byte(`{"value": 100, "unit": "ThermUs"}`)
    therms_usResult, err := factory.FromDtoJSON(therms_usJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ThermUs unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = therms_usResult.Convert(units.EnergyThermUs)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ThermUs = %v, want %v", converted, 100)
    }
    // Test JSON with ThermImperial unit
    therms_imperialJSON := []byte(`{"value": 100, "unit": "ThermImperial"}`)
    therms_imperialResult, err := factory.FromDtoJSON(therms_imperialJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ThermImperial unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = therms_imperialResult.Convert(units.EnergyThermImperial)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ThermImperial = %v, want %v", converted, 100)
    }
    // Test JSON with HorsepowerHour unit
    horsepower_hoursJSON := []byte(`{"value": 100, "unit": "HorsepowerHour"}`)
    horsepower_hoursResult, err := factory.FromDtoJSON(horsepower_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HorsepowerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = horsepower_hoursResult.Convert(units.EnergyHorsepowerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HorsepowerHour = %v, want %v", converted, 100)
    }
    // Test JSON with Nanojoule unit
    nanojoulesJSON := []byte(`{"value": 100, "unit": "Nanojoule"}`)
    nanojoulesResult, err := factory.FromDtoJSON(nanojoulesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanojoule unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanojoulesResult.Convert(units.EnergyNanojoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanojoule = %v, want %v", converted, 100)
    }
    // Test JSON with Microjoule unit
    microjoulesJSON := []byte(`{"value": 100, "unit": "Microjoule"}`)
    microjoulesResult, err := factory.FromDtoJSON(microjoulesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microjoule unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microjoulesResult.Convert(units.EnergyMicrojoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microjoule = %v, want %v", converted, 100)
    }
    // Test JSON with Millijoule unit
    millijoulesJSON := []byte(`{"value": 100, "unit": "Millijoule"}`)
    millijoulesResult, err := factory.FromDtoJSON(millijoulesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millijoule unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millijoulesResult.Convert(units.EnergyMillijoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millijoule = %v, want %v", converted, 100)
    }
    // Test JSON with Kilojoule unit
    kilojoulesJSON := []byte(`{"value": 100, "unit": "Kilojoule"}`)
    kilojoulesResult, err := factory.FromDtoJSON(kilojoulesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilojoule unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoulesResult.Convert(units.EnergyKilojoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilojoule = %v, want %v", converted, 100)
    }
    // Test JSON with Megajoule unit
    megajoulesJSON := []byte(`{"value": 100, "unit": "Megajoule"}`)
    megajoulesResult, err := factory.FromDtoJSON(megajoulesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megajoule unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoulesResult.Convert(units.EnergyMegajoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megajoule = %v, want %v", converted, 100)
    }
    // Test JSON with Gigajoule unit
    gigajoulesJSON := []byte(`{"value": 100, "unit": "Gigajoule"}`)
    gigajoulesResult, err := factory.FromDtoJSON(gigajoulesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigajoule unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigajoulesResult.Convert(units.EnergyGigajoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigajoule = %v, want %v", converted, 100)
    }
    // Test JSON with Terajoule unit
    terajoulesJSON := []byte(`{"value": 100, "unit": "Terajoule"}`)
    terajoulesResult, err := factory.FromDtoJSON(terajoulesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Terajoule unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terajoulesResult.Convert(units.EnergyTerajoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terajoule = %v, want %v", converted, 100)
    }
    // Test JSON with Petajoule unit
    petajoulesJSON := []byte(`{"value": 100, "unit": "Petajoule"}`)
    petajoulesResult, err := factory.FromDtoJSON(petajoulesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Petajoule unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petajoulesResult.Convert(units.EnergyPetajoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petajoule = %v, want %v", converted, 100)
    }
    // Test JSON with Kilocalorie unit
    kilocaloriesJSON := []byte(`{"value": 100, "unit": "Kilocalorie"}`)
    kilocaloriesResult, err := factory.FromDtoJSON(kilocaloriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilocalorie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocaloriesResult.Convert(units.EnergyKilocalorie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilocalorie = %v, want %v", converted, 100)
    }
    // Test JSON with Megacalorie unit
    megacaloriesJSON := []byte(`{"value": 100, "unit": "Megacalorie"}`)
    megacaloriesResult, err := factory.FromDtoJSON(megacaloriesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megacalorie unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megacaloriesResult.Convert(units.EnergyMegacalorie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megacalorie = %v, want %v", converted, 100)
    }
    // Test JSON with KilobritishThermalUnit unit
    kilobritish_thermal_unitsJSON := []byte(`{"value": 100, "unit": "KilobritishThermalUnit"}`)
    kilobritish_thermal_unitsResult, err := factory.FromDtoJSON(kilobritish_thermal_unitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilobritishThermalUnit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobritish_thermal_unitsResult.Convert(units.EnergyKilobritishThermalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilobritishThermalUnit = %v, want %v", converted, 100)
    }
    // Test JSON with MegabritishThermalUnit unit
    megabritish_thermal_unitsJSON := []byte(`{"value": 100, "unit": "MegabritishThermalUnit"}`)
    megabritish_thermal_unitsResult, err := factory.FromDtoJSON(megabritish_thermal_unitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegabritishThermalUnit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabritish_thermal_unitsResult.Convert(units.EnergyMegabritishThermalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegabritishThermalUnit = %v, want %v", converted, 100)
    }
    // Test JSON with GigabritishThermalUnit unit
    gigabritish_thermal_unitsJSON := []byte(`{"value": 100, "unit": "GigabritishThermalUnit"}`)
    gigabritish_thermal_unitsResult, err := factory.FromDtoJSON(gigabritish_thermal_unitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigabritishThermalUnit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigabritish_thermal_unitsResult.Convert(units.EnergyGigabritishThermalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigabritishThermalUnit = %v, want %v", converted, 100)
    }
    // Test JSON with KiloelectronVolt unit
    kiloelectron_voltsJSON := []byte(`{"value": 100, "unit": "KiloelectronVolt"}`)
    kiloelectron_voltsResult, err := factory.FromDtoJSON(kiloelectron_voltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KiloelectronVolt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloelectron_voltsResult.Convert(units.EnergyKiloelectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloelectronVolt = %v, want %v", converted, 100)
    }
    // Test JSON with MegaelectronVolt unit
    megaelectron_voltsJSON := []byte(`{"value": 100, "unit": "MegaelectronVolt"}`)
    megaelectron_voltsResult, err := factory.FromDtoJSON(megaelectron_voltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegaelectronVolt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaelectron_voltsResult.Convert(units.EnergyMegaelectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaelectronVolt = %v, want %v", converted, 100)
    }
    // Test JSON with GigaelectronVolt unit
    gigaelectron_voltsJSON := []byte(`{"value": 100, "unit": "GigaelectronVolt"}`)
    gigaelectron_voltsResult, err := factory.FromDtoJSON(gigaelectron_voltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigaelectronVolt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigaelectron_voltsResult.Convert(units.EnergyGigaelectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigaelectronVolt = %v, want %v", converted, 100)
    }
    // Test JSON with TeraelectronVolt unit
    teraelectron_voltsJSON := []byte(`{"value": 100, "unit": "TeraelectronVolt"}`)
    teraelectron_voltsResult, err := factory.FromDtoJSON(teraelectron_voltsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TeraelectronVolt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = teraelectron_voltsResult.Convert(units.EnergyTeraelectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TeraelectronVolt = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattHour unit
    kilowatt_hoursJSON := []byte(`{"value": 100, "unit": "KilowattHour"}`)
    kilowatt_hoursResult, err := factory.FromDtoJSON(kilowatt_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_hoursResult.Convert(units.EnergyKilowattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattHour = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattHour unit
    megawatt_hoursJSON := []byte(`{"value": 100, "unit": "MegawattHour"}`)
    megawatt_hoursResult, err := factory.FromDtoJSON(megawatt_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_hoursResult.Convert(units.EnergyMegawattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattHour = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattHour unit
    gigawatt_hoursJSON := []byte(`{"value": 100, "unit": "GigawattHour"}`)
    gigawatt_hoursResult, err := factory.FromDtoJSON(gigawatt_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_hoursResult.Convert(units.EnergyGigawattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattHour = %v, want %v", converted, 100)
    }
    // Test JSON with TerawattHour unit
    terawatt_hoursJSON := []byte(`{"value": 100, "unit": "TerawattHour"}`)
    terawatt_hoursResult, err := factory.FromDtoJSON(terawatt_hoursJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TerawattHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatt_hoursResult.Convert(units.EnergyTerawattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattHour = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattDay unit
    kilowatt_daysJSON := []byte(`{"value": 100, "unit": "KilowattDay"}`)
    kilowatt_daysResult, err := factory.FromDtoJSON(kilowatt_daysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatt_daysResult.Convert(units.EnergyKilowattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattDay = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattDay unit
    megawatt_daysJSON := []byte(`{"value": 100, "unit": "MegawattDay"}`)
    megawatt_daysResult, err := factory.FromDtoJSON(megawatt_daysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatt_daysResult.Convert(units.EnergyMegawattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattDay = %v, want %v", converted, 100)
    }
    // Test JSON with GigawattDay unit
    gigawatt_daysJSON := []byte(`{"value": 100, "unit": "GigawattDay"}`)
    gigawatt_daysResult, err := factory.FromDtoJSON(gigawatt_daysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigawattDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawatt_daysResult.Convert(units.EnergyGigawattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigawattDay = %v, want %v", converted, 100)
    }
    // Test JSON with TerawattDay unit
    terawatt_daysJSON := []byte(`{"value": 100, "unit": "TerawattDay"}`)
    terawatt_daysResult, err := factory.FromDtoJSON(terawatt_daysJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TerawattDay unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawatt_daysResult.Convert(units.EnergyTerawattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TerawattDay = %v, want %v", converted, 100)
    }
    // Test JSON with DecathermEc unit
    decatherms_ecJSON := []byte(`{"value": 100, "unit": "DecathermEc"}`)
    decatherms_ecResult, err := factory.FromDtoJSON(decatherms_ecJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecathermEc unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decatherms_ecResult.Convert(units.EnergyDecathermEc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecathermEc = %v, want %v", converted, 100)
    }
    // Test JSON with DecathermUs unit
    decatherms_usJSON := []byte(`{"value": 100, "unit": "DecathermUs"}`)
    decatherms_usResult, err := factory.FromDtoJSON(decatherms_usJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecathermUs unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decatherms_usResult.Convert(units.EnergyDecathermUs)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecathermUs = %v, want %v", converted, 100)
    }
    // Test JSON with DecathermImperial unit
    decatherms_imperialJSON := []byte(`{"value": 100, "unit": "DecathermImperial"}`)
    decatherms_imperialResult, err := factory.FromDtoJSON(decatherms_imperialJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecathermImperial unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decatherms_imperialResult.Convert(units.EnergyDecathermImperial)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecathermImperial = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Joule"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromJoules function
func TestEnergyFactory_FromJoules(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoules(100)
    if err != nil {
        t.Errorf("FromJoules() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyJoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoules() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoules(math.NaN())
    if err == nil {
        t.Error("FromJoules() with NaN value should return error")
    }

    _, err = factory.FromJoules(math.Inf(1))
    if err == nil {
        t.Error("FromJoules() with +Inf value should return error")
    }

    _, err = factory.FromJoules(math.Inf(-1))
    if err == nil {
        t.Error("FromJoules() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoules(0)
    if err != nil {
        t.Errorf("FromJoules() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyJoule)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoules() with zero value = %v, want 0", converted)
    }
}
// Test FromCalories function
func TestEnergyFactory_FromCalories(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCalories(100)
    if err != nil {
        t.Errorf("FromCalories() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyCalorie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCalories() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCalories(math.NaN())
    if err == nil {
        t.Error("FromCalories() with NaN value should return error")
    }

    _, err = factory.FromCalories(math.Inf(1))
    if err == nil {
        t.Error("FromCalories() with +Inf value should return error")
    }

    _, err = factory.FromCalories(math.Inf(-1))
    if err == nil {
        t.Error("FromCalories() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCalories(0)
    if err != nil {
        t.Errorf("FromCalories() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyCalorie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCalories() with zero value = %v, want 0", converted)
    }
}
// Test FromBritishThermalUnits function
func TestEnergyFactory_FromBritishThermalUnits(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBritishThermalUnits(100)
    if err != nil {
        t.Errorf("FromBritishThermalUnits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyBritishThermalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBritishThermalUnits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBritishThermalUnits(math.NaN())
    if err == nil {
        t.Error("FromBritishThermalUnits() with NaN value should return error")
    }

    _, err = factory.FromBritishThermalUnits(math.Inf(1))
    if err == nil {
        t.Error("FromBritishThermalUnits() with +Inf value should return error")
    }

    _, err = factory.FromBritishThermalUnits(math.Inf(-1))
    if err == nil {
        t.Error("FromBritishThermalUnits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBritishThermalUnits(0)
    if err != nil {
        t.Errorf("FromBritishThermalUnits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyBritishThermalUnit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBritishThermalUnits() with zero value = %v, want 0", converted)
    }
}
// Test FromElectronVolts function
func TestEnergyFactory_FromElectronVolts(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromElectronVolts(100)
    if err != nil {
        t.Errorf("FromElectronVolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyElectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromElectronVolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromElectronVolts(math.NaN())
    if err == nil {
        t.Error("FromElectronVolts() with NaN value should return error")
    }

    _, err = factory.FromElectronVolts(math.Inf(1))
    if err == nil {
        t.Error("FromElectronVolts() with +Inf value should return error")
    }

    _, err = factory.FromElectronVolts(math.Inf(-1))
    if err == nil {
        t.Error("FromElectronVolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromElectronVolts(0)
    if err != nil {
        t.Errorf("FromElectronVolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyElectronVolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromElectronVolts() with zero value = %v, want 0", converted)
    }
}
// Test FromFootPounds function
func TestEnergyFactory_FromFootPounds(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFootPounds(100)
    if err != nil {
        t.Errorf("FromFootPounds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyFootPound)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFootPounds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFootPounds(math.NaN())
    if err == nil {
        t.Error("FromFootPounds() with NaN value should return error")
    }

    _, err = factory.FromFootPounds(math.Inf(1))
    if err == nil {
        t.Error("FromFootPounds() with +Inf value should return error")
    }

    _, err = factory.FromFootPounds(math.Inf(-1))
    if err == nil {
        t.Error("FromFootPounds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFootPounds(0)
    if err != nil {
        t.Errorf("FromFootPounds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyFootPound)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFootPounds() with zero value = %v, want 0", converted)
    }
}
// Test FromErgs function
func TestEnergyFactory_FromErgs(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromErgs(100)
    if err != nil {
        t.Errorf("FromErgs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyErg)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromErgs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromErgs(math.NaN())
    if err == nil {
        t.Error("FromErgs() with NaN value should return error")
    }

    _, err = factory.FromErgs(math.Inf(1))
    if err == nil {
        t.Error("FromErgs() with +Inf value should return error")
    }

    _, err = factory.FromErgs(math.Inf(-1))
    if err == nil {
        t.Error("FromErgs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromErgs(0)
    if err != nil {
        t.Errorf("FromErgs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyErg)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromErgs() with zero value = %v, want 0", converted)
    }
}
// Test FromWattHours function
func TestEnergyFactory_FromWattHours(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattHours(100)
    if err != nil {
        t.Errorf("FromWattHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyWattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattHours(math.NaN())
    if err == nil {
        t.Error("FromWattHours() with NaN value should return error")
    }

    _, err = factory.FromWattHours(math.Inf(1))
    if err == nil {
        t.Error("FromWattHours() with +Inf value should return error")
    }

    _, err = factory.FromWattHours(math.Inf(-1))
    if err == nil {
        t.Error("FromWattHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattHours(0)
    if err != nil {
        t.Errorf("FromWattHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyWattHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattHours() with zero value = %v, want 0", converted)
    }
}
// Test FromWattDays function
func TestEnergyFactory_FromWattDays(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattDays(100)
    if err != nil {
        t.Errorf("FromWattDays() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyWattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattDays() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattDays(math.NaN())
    if err == nil {
        t.Error("FromWattDays() with NaN value should return error")
    }

    _, err = factory.FromWattDays(math.Inf(1))
    if err == nil {
        t.Error("FromWattDays() with +Inf value should return error")
    }

    _, err = factory.FromWattDays(math.Inf(-1))
    if err == nil {
        t.Error("FromWattDays() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattDays(0)
    if err != nil {
        t.Errorf("FromWattDays() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyWattDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattDays() with zero value = %v, want 0", converted)
    }
}
// Test FromThermsEc function
func TestEnergyFactory_FromThermsEc(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromThermsEc(100)
    if err != nil {
        t.Errorf("FromThermsEc() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyThermEc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromThermsEc() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromThermsEc(math.NaN())
    if err == nil {
        t.Error("FromThermsEc() with NaN value should return error")
    }

    _, err = factory.FromThermsEc(math.Inf(1))
    if err == nil {
        t.Error("FromThermsEc() with +Inf value should return error")
    }

    _, err = factory.FromThermsEc(math.Inf(-1))
    if err == nil {
        t.Error("FromThermsEc() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromThermsEc(0)
    if err != nil {
        t.Errorf("FromThermsEc() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyThermEc)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromThermsEc() with zero value = %v, want 0", converted)
    }
}
// Test FromThermsUs function
func TestEnergyFactory_FromThermsUs(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromThermsUs(100)
    if err != nil {
        t.Errorf("FromThermsUs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyThermUs)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromThermsUs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromThermsUs(math.NaN())
    if err == nil {
        t.Error("FromThermsUs() with NaN value should return error")
    }

    _, err = factory.FromThermsUs(math.Inf(1))
    if err == nil {
        t.Error("FromThermsUs() with +Inf value should return error")
    }

    _, err = factory.FromThermsUs(math.Inf(-1))
    if err == nil {
        t.Error("FromThermsUs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromThermsUs(0)
    if err != nil {
        t.Errorf("FromThermsUs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyThermUs)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromThermsUs() with zero value = %v, want 0", converted)
    }
}
// Test FromThermsImperial function
func TestEnergyFactory_FromThermsImperial(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromThermsImperial(100)
    if err != nil {
        t.Errorf("FromThermsImperial() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyThermImperial)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromThermsImperial() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromThermsImperial(math.NaN())
    if err == nil {
        t.Error("FromThermsImperial() with NaN value should return error")
    }

    _, err = factory.FromThermsImperial(math.Inf(1))
    if err == nil {
        t.Error("FromThermsImperial() with +Inf value should return error")
    }

    _, err = factory.FromThermsImperial(math.Inf(-1))
    if err == nil {
        t.Error("FromThermsImperial() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromThermsImperial(0)
    if err != nil {
        t.Errorf("FromThermsImperial() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyThermImperial)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromThermsImperial() with zero value = %v, want 0", converted)
    }
}
// Test FromHorsepowerHours function
func TestEnergyFactory_FromHorsepowerHours(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHorsepowerHours(100)
    if err != nil {
        t.Errorf("FromHorsepowerHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyHorsepowerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHorsepowerHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHorsepowerHours(math.NaN())
    if err == nil {
        t.Error("FromHorsepowerHours() with NaN value should return error")
    }

    _, err = factory.FromHorsepowerHours(math.Inf(1))
    if err == nil {
        t.Error("FromHorsepowerHours() with +Inf value should return error")
    }

    _, err = factory.FromHorsepowerHours(math.Inf(-1))
    if err == nil {
        t.Error("FromHorsepowerHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHorsepowerHours(0)
    if err != nil {
        t.Errorf("FromHorsepowerHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyHorsepowerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHorsepowerHours() with zero value = %v, want 0", converted)
    }
}
// Test FromNanojoules function
func TestEnergyFactory_FromNanojoules(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanojoules(100)
    if err != nil {
        t.Errorf("FromNanojoules() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyNanojoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanojoules() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanojoules(math.NaN())
    if err == nil {
        t.Error("FromNanojoules() with NaN value should return error")
    }

    _, err = factory.FromNanojoules(math.Inf(1))
    if err == nil {
        t.Error("FromNanojoules() with +Inf value should return error")
    }

    _, err = factory.FromNanojoules(math.Inf(-1))
    if err == nil {
        t.Error("FromNanojoules() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanojoules(0)
    if err != nil {
        t.Errorf("FromNanojoules() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyNanojoule)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanojoules() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrojoules function
func TestEnergyFactory_FromMicrojoules(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrojoules(100)
    if err != nil {
        t.Errorf("FromMicrojoules() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyMicrojoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrojoules() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrojoules(math.NaN())
    if err == nil {
        t.Error("FromMicrojoules() with NaN value should return error")
    }

    _, err = factory.FromMicrojoules(math.Inf(1))
    if err == nil {
        t.Error("FromMicrojoules() with +Inf value should return error")
    }

    _, err = factory.FromMicrojoules(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrojoules() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrojoules(0)
    if err != nil {
        t.Errorf("FromMicrojoules() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyMicrojoule)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrojoules() with zero value = %v, want 0", converted)
    }
}
// Test FromMillijoules function
func TestEnergyFactory_FromMillijoules(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillijoules(100)
    if err != nil {
        t.Errorf("FromMillijoules() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyMillijoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillijoules() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillijoules(math.NaN())
    if err == nil {
        t.Error("FromMillijoules() with NaN value should return error")
    }

    _, err = factory.FromMillijoules(math.Inf(1))
    if err == nil {
        t.Error("FromMillijoules() with +Inf value should return error")
    }

    _, err = factory.FromMillijoules(math.Inf(-1))
    if err == nil {
        t.Error("FromMillijoules() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillijoules(0)
    if err != nil {
        t.Errorf("FromMillijoules() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyMillijoule)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillijoules() with zero value = %v, want 0", converted)
    }
}
// Test FromKilojoules function
func TestEnergyFactory_FromKilojoules(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilojoules(100)
    if err != nil {
        t.Errorf("FromKilojoules() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyKilojoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilojoules() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilojoules(math.NaN())
    if err == nil {
        t.Error("FromKilojoules() with NaN value should return error")
    }

    _, err = factory.FromKilojoules(math.Inf(1))
    if err == nil {
        t.Error("FromKilojoules() with +Inf value should return error")
    }

    _, err = factory.FromKilojoules(math.Inf(-1))
    if err == nil {
        t.Error("FromKilojoules() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilojoules(0)
    if err != nil {
        t.Errorf("FromKilojoules() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyKilojoule)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilojoules() with zero value = %v, want 0", converted)
    }
}
// Test FromMegajoules function
func TestEnergyFactory_FromMegajoules(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegajoules(100)
    if err != nil {
        t.Errorf("FromMegajoules() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyMegajoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegajoules() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegajoules(math.NaN())
    if err == nil {
        t.Error("FromMegajoules() with NaN value should return error")
    }

    _, err = factory.FromMegajoules(math.Inf(1))
    if err == nil {
        t.Error("FromMegajoules() with +Inf value should return error")
    }

    _, err = factory.FromMegajoules(math.Inf(-1))
    if err == nil {
        t.Error("FromMegajoules() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegajoules(0)
    if err != nil {
        t.Errorf("FromMegajoules() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyMegajoule)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegajoules() with zero value = %v, want 0", converted)
    }
}
// Test FromGigajoules function
func TestEnergyFactory_FromGigajoules(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigajoules(100)
    if err != nil {
        t.Errorf("FromGigajoules() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyGigajoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigajoules() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigajoules(math.NaN())
    if err == nil {
        t.Error("FromGigajoules() with NaN value should return error")
    }

    _, err = factory.FromGigajoules(math.Inf(1))
    if err == nil {
        t.Error("FromGigajoules() with +Inf value should return error")
    }

    _, err = factory.FromGigajoules(math.Inf(-1))
    if err == nil {
        t.Error("FromGigajoules() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigajoules(0)
    if err != nil {
        t.Errorf("FromGigajoules() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyGigajoule)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigajoules() with zero value = %v, want 0", converted)
    }
}
// Test FromTerajoules function
func TestEnergyFactory_FromTerajoules(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerajoules(100)
    if err != nil {
        t.Errorf("FromTerajoules() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyTerajoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerajoules() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerajoules(math.NaN())
    if err == nil {
        t.Error("FromTerajoules() with NaN value should return error")
    }

    _, err = factory.FromTerajoules(math.Inf(1))
    if err == nil {
        t.Error("FromTerajoules() with +Inf value should return error")
    }

    _, err = factory.FromTerajoules(math.Inf(-1))
    if err == nil {
        t.Error("FromTerajoules() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerajoules(0)
    if err != nil {
        t.Errorf("FromTerajoules() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyTerajoule)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerajoules() with zero value = %v, want 0", converted)
    }
}
// Test FromPetajoules function
func TestEnergyFactory_FromPetajoules(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPetajoules(100)
    if err != nil {
        t.Errorf("FromPetajoules() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyPetajoule)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPetajoules() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPetajoules(math.NaN())
    if err == nil {
        t.Error("FromPetajoules() with NaN value should return error")
    }

    _, err = factory.FromPetajoules(math.Inf(1))
    if err == nil {
        t.Error("FromPetajoules() with +Inf value should return error")
    }

    _, err = factory.FromPetajoules(math.Inf(-1))
    if err == nil {
        t.Error("FromPetajoules() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPetajoules(0)
    if err != nil {
        t.Errorf("FromPetajoules() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyPetajoule)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPetajoules() with zero value = %v, want 0", converted)
    }
}
// Test FromKilocalories function
func TestEnergyFactory_FromKilocalories(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilocalories(100)
    if err != nil {
        t.Errorf("FromKilocalories() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyKilocalorie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilocalories() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilocalories(math.NaN())
    if err == nil {
        t.Error("FromKilocalories() with NaN value should return error")
    }

    _, err = factory.FromKilocalories(math.Inf(1))
    if err == nil {
        t.Error("FromKilocalories() with +Inf value should return error")
    }

    _, err = factory.FromKilocalories(math.Inf(-1))
    if err == nil {
        t.Error("FromKilocalories() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilocalories(0)
    if err != nil {
        t.Errorf("FromKilocalories() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyKilocalorie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilocalories() with zero value = %v, want 0", converted)
    }
}
// Test FromMegacalories function
func TestEnergyFactory_FromMegacalories(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegacalories(100)
    if err != nil {
        t.Errorf("FromMegacalories() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyMegacalorie)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegacalories() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegacalories(math.NaN())
    if err == nil {
        t.Error("FromMegacalories() with NaN value should return error")
    }

    _, err = factory.FromMegacalories(math.Inf(1))
    if err == nil {
        t.Error("FromMegacalories() with +Inf value should return error")
    }

    _, err = factory.FromMegacalories(math.Inf(-1))
    if err == nil {
        t.Error("FromMegacalories() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegacalories(0)
    if err != nil {
        t.Errorf("FromMegacalories() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyMegacalorie)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegacalories() with zero value = %v, want 0", converted)
    }
}
// Test FromKilobritishThermalUnits function
func TestEnergyFactory_FromKilobritishThermalUnits(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilobritishThermalUnits(100)
    if err != nil {
        t.Errorf("FromKilobritishThermalUnits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyKilobritishThermalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilobritishThermalUnits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilobritishThermalUnits(math.NaN())
    if err == nil {
        t.Error("FromKilobritishThermalUnits() with NaN value should return error")
    }

    _, err = factory.FromKilobritishThermalUnits(math.Inf(1))
    if err == nil {
        t.Error("FromKilobritishThermalUnits() with +Inf value should return error")
    }

    _, err = factory.FromKilobritishThermalUnits(math.Inf(-1))
    if err == nil {
        t.Error("FromKilobritishThermalUnits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilobritishThermalUnits(0)
    if err != nil {
        t.Errorf("FromKilobritishThermalUnits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyKilobritishThermalUnit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilobritishThermalUnits() with zero value = %v, want 0", converted)
    }
}
// Test FromMegabritishThermalUnits function
func TestEnergyFactory_FromMegabritishThermalUnits(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegabritishThermalUnits(100)
    if err != nil {
        t.Errorf("FromMegabritishThermalUnits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyMegabritishThermalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegabritishThermalUnits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegabritishThermalUnits(math.NaN())
    if err == nil {
        t.Error("FromMegabritishThermalUnits() with NaN value should return error")
    }

    _, err = factory.FromMegabritishThermalUnits(math.Inf(1))
    if err == nil {
        t.Error("FromMegabritishThermalUnits() with +Inf value should return error")
    }

    _, err = factory.FromMegabritishThermalUnits(math.Inf(-1))
    if err == nil {
        t.Error("FromMegabritishThermalUnits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegabritishThermalUnits(0)
    if err != nil {
        t.Errorf("FromMegabritishThermalUnits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyMegabritishThermalUnit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegabritishThermalUnits() with zero value = %v, want 0", converted)
    }
}
// Test FromGigabritishThermalUnits function
func TestEnergyFactory_FromGigabritishThermalUnits(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigabritishThermalUnits(100)
    if err != nil {
        t.Errorf("FromGigabritishThermalUnits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyGigabritishThermalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigabritishThermalUnits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigabritishThermalUnits(math.NaN())
    if err == nil {
        t.Error("FromGigabritishThermalUnits() with NaN value should return error")
    }

    _, err = factory.FromGigabritishThermalUnits(math.Inf(1))
    if err == nil {
        t.Error("FromGigabritishThermalUnits() with +Inf value should return error")
    }

    _, err = factory.FromGigabritishThermalUnits(math.Inf(-1))
    if err == nil {
        t.Error("FromGigabritishThermalUnits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigabritishThermalUnits(0)
    if err != nil {
        t.Errorf("FromGigabritishThermalUnits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyGigabritishThermalUnit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigabritishThermalUnits() with zero value = %v, want 0", converted)
    }
}
// Test FromKiloelectronVolts function
func TestEnergyFactory_FromKiloelectronVolts(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKiloelectronVolts(100)
    if err != nil {
        t.Errorf("FromKiloelectronVolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyKiloelectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKiloelectronVolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKiloelectronVolts(math.NaN())
    if err == nil {
        t.Error("FromKiloelectronVolts() with NaN value should return error")
    }

    _, err = factory.FromKiloelectronVolts(math.Inf(1))
    if err == nil {
        t.Error("FromKiloelectronVolts() with +Inf value should return error")
    }

    _, err = factory.FromKiloelectronVolts(math.Inf(-1))
    if err == nil {
        t.Error("FromKiloelectronVolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKiloelectronVolts(0)
    if err != nil {
        t.Errorf("FromKiloelectronVolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyKiloelectronVolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKiloelectronVolts() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaelectronVolts function
func TestEnergyFactory_FromMegaelectronVolts(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaelectronVolts(100)
    if err != nil {
        t.Errorf("FromMegaelectronVolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyMegaelectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegaelectronVolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegaelectronVolts(math.NaN())
    if err == nil {
        t.Error("FromMegaelectronVolts() with NaN value should return error")
    }

    _, err = factory.FromMegaelectronVolts(math.Inf(1))
    if err == nil {
        t.Error("FromMegaelectronVolts() with +Inf value should return error")
    }

    _, err = factory.FromMegaelectronVolts(math.Inf(-1))
    if err == nil {
        t.Error("FromMegaelectronVolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegaelectronVolts(0)
    if err != nil {
        t.Errorf("FromMegaelectronVolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyMegaelectronVolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaelectronVolts() with zero value = %v, want 0", converted)
    }
}
// Test FromGigaelectronVolts function
func TestEnergyFactory_FromGigaelectronVolts(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigaelectronVolts(100)
    if err != nil {
        t.Errorf("FromGigaelectronVolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyGigaelectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigaelectronVolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigaelectronVolts(math.NaN())
    if err == nil {
        t.Error("FromGigaelectronVolts() with NaN value should return error")
    }

    _, err = factory.FromGigaelectronVolts(math.Inf(1))
    if err == nil {
        t.Error("FromGigaelectronVolts() with +Inf value should return error")
    }

    _, err = factory.FromGigaelectronVolts(math.Inf(-1))
    if err == nil {
        t.Error("FromGigaelectronVolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigaelectronVolts(0)
    if err != nil {
        t.Errorf("FromGigaelectronVolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyGigaelectronVolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigaelectronVolts() with zero value = %v, want 0", converted)
    }
}
// Test FromTeraelectronVolts function
func TestEnergyFactory_FromTeraelectronVolts(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTeraelectronVolts(100)
    if err != nil {
        t.Errorf("FromTeraelectronVolts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyTeraelectronVolt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTeraelectronVolts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTeraelectronVolts(math.NaN())
    if err == nil {
        t.Error("FromTeraelectronVolts() with NaN value should return error")
    }

    _, err = factory.FromTeraelectronVolts(math.Inf(1))
    if err == nil {
        t.Error("FromTeraelectronVolts() with +Inf value should return error")
    }

    _, err = factory.FromTeraelectronVolts(math.Inf(-1))
    if err == nil {
        t.Error("FromTeraelectronVolts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTeraelectronVolts(0)
    if err != nil {
        t.Errorf("FromTeraelectronVolts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyTeraelectronVolt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTeraelectronVolts() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattHours function
func TestEnergyFactory_FromKilowattHours(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattHours(100)
    if err != nil {
        t.Errorf("FromKilowattHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyKilowattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattHours(math.NaN())
    if err == nil {
        t.Error("FromKilowattHours() with NaN value should return error")
    }

    _, err = factory.FromKilowattHours(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattHours() with +Inf value should return error")
    }

    _, err = factory.FromKilowattHours(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattHours(0)
    if err != nil {
        t.Errorf("FromKilowattHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyKilowattHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattHours() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattHours function
func TestEnergyFactory_FromMegawattHours(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattHours(100)
    if err != nil {
        t.Errorf("FromMegawattHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyMegawattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattHours(math.NaN())
    if err == nil {
        t.Error("FromMegawattHours() with NaN value should return error")
    }

    _, err = factory.FromMegawattHours(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattHours() with +Inf value should return error")
    }

    _, err = factory.FromMegawattHours(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattHours(0)
    if err != nil {
        t.Errorf("FromMegawattHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyMegawattHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattHours() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattHours function
func TestEnergyFactory_FromGigawattHours(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattHours(100)
    if err != nil {
        t.Errorf("FromGigawattHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyGigawattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattHours(math.NaN())
    if err == nil {
        t.Error("FromGigawattHours() with NaN value should return error")
    }

    _, err = factory.FromGigawattHours(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattHours() with +Inf value should return error")
    }

    _, err = factory.FromGigawattHours(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattHours(0)
    if err != nil {
        t.Errorf("FromGigawattHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyGigawattHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattHours() with zero value = %v, want 0", converted)
    }
}
// Test FromTerawattHours function
func TestEnergyFactory_FromTerawattHours(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerawattHours(100)
    if err != nil {
        t.Errorf("FromTerawattHours() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyTerawattHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerawattHours() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerawattHours(math.NaN())
    if err == nil {
        t.Error("FromTerawattHours() with NaN value should return error")
    }

    _, err = factory.FromTerawattHours(math.Inf(1))
    if err == nil {
        t.Error("FromTerawattHours() with +Inf value should return error")
    }

    _, err = factory.FromTerawattHours(math.Inf(-1))
    if err == nil {
        t.Error("FromTerawattHours() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerawattHours(0)
    if err != nil {
        t.Errorf("FromTerawattHours() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyTerawattHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerawattHours() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattDays function
func TestEnergyFactory_FromKilowattDays(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattDays(100)
    if err != nil {
        t.Errorf("FromKilowattDays() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyKilowattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattDays() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattDays(math.NaN())
    if err == nil {
        t.Error("FromKilowattDays() with NaN value should return error")
    }

    _, err = factory.FromKilowattDays(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattDays() with +Inf value should return error")
    }

    _, err = factory.FromKilowattDays(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattDays() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattDays(0)
    if err != nil {
        t.Errorf("FromKilowattDays() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyKilowattDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattDays() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattDays function
func TestEnergyFactory_FromMegawattDays(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattDays(100)
    if err != nil {
        t.Errorf("FromMegawattDays() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyMegawattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattDays() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattDays(math.NaN())
    if err == nil {
        t.Error("FromMegawattDays() with NaN value should return error")
    }

    _, err = factory.FromMegawattDays(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattDays() with +Inf value should return error")
    }

    _, err = factory.FromMegawattDays(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattDays() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattDays(0)
    if err != nil {
        t.Errorf("FromMegawattDays() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyMegawattDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattDays() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawattDays function
func TestEnergyFactory_FromGigawattDays(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawattDays(100)
    if err != nil {
        t.Errorf("FromGigawattDays() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyGigawattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawattDays() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawattDays(math.NaN())
    if err == nil {
        t.Error("FromGigawattDays() with NaN value should return error")
    }

    _, err = factory.FromGigawattDays(math.Inf(1))
    if err == nil {
        t.Error("FromGigawattDays() with +Inf value should return error")
    }

    _, err = factory.FromGigawattDays(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawattDays() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawattDays(0)
    if err != nil {
        t.Errorf("FromGigawattDays() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyGigawattDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawattDays() with zero value = %v, want 0", converted)
    }
}
// Test FromTerawattDays function
func TestEnergyFactory_FromTerawattDays(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerawattDays(100)
    if err != nil {
        t.Errorf("FromTerawattDays() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyTerawattDay)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerawattDays() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerawattDays(math.NaN())
    if err == nil {
        t.Error("FromTerawattDays() with NaN value should return error")
    }

    _, err = factory.FromTerawattDays(math.Inf(1))
    if err == nil {
        t.Error("FromTerawattDays() with +Inf value should return error")
    }

    _, err = factory.FromTerawattDays(math.Inf(-1))
    if err == nil {
        t.Error("FromTerawattDays() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerawattDays(0)
    if err != nil {
        t.Errorf("FromTerawattDays() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyTerawattDay)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerawattDays() with zero value = %v, want 0", converted)
    }
}
// Test FromDecathermsEc function
func TestEnergyFactory_FromDecathermsEc(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecathermsEc(100)
    if err != nil {
        t.Errorf("FromDecathermsEc() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDecathermEc)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecathermsEc() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecathermsEc(math.NaN())
    if err == nil {
        t.Error("FromDecathermsEc() with NaN value should return error")
    }

    _, err = factory.FromDecathermsEc(math.Inf(1))
    if err == nil {
        t.Error("FromDecathermsEc() with +Inf value should return error")
    }

    _, err = factory.FromDecathermsEc(math.Inf(-1))
    if err == nil {
        t.Error("FromDecathermsEc() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecathermsEc(0)
    if err != nil {
        t.Errorf("FromDecathermsEc() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDecathermEc)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecathermsEc() with zero value = %v, want 0", converted)
    }
}
// Test FromDecathermsUs function
func TestEnergyFactory_FromDecathermsUs(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecathermsUs(100)
    if err != nil {
        t.Errorf("FromDecathermsUs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDecathermUs)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecathermsUs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecathermsUs(math.NaN())
    if err == nil {
        t.Error("FromDecathermsUs() with NaN value should return error")
    }

    _, err = factory.FromDecathermsUs(math.Inf(1))
    if err == nil {
        t.Error("FromDecathermsUs() with +Inf value should return error")
    }

    _, err = factory.FromDecathermsUs(math.Inf(-1))
    if err == nil {
        t.Error("FromDecathermsUs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecathermsUs(0)
    if err != nil {
        t.Errorf("FromDecathermsUs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDecathermUs)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecathermsUs() with zero value = %v, want 0", converted)
    }
}
// Test FromDecathermsImperial function
func TestEnergyFactory_FromDecathermsImperial(t *testing.T) {
    factory := units.EnergyFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecathermsImperial(100)
    if err != nil {
        t.Errorf("FromDecathermsImperial() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.EnergyDecathermImperial)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecathermsImperial() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecathermsImperial(math.NaN())
    if err == nil {
        t.Error("FromDecathermsImperial() with NaN value should return error")
    }

    _, err = factory.FromDecathermsImperial(math.Inf(1))
    if err == nil {
        t.Error("FromDecathermsImperial() with +Inf value should return error")
    }

    _, err = factory.FromDecathermsImperial(math.Inf(-1))
    if err == nil {
        t.Error("FromDecathermsImperial() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecathermsImperial(0)
    if err != nil {
        t.Errorf("FromDecathermsImperial() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.EnergyDecathermImperial)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecathermsImperial() with zero value = %v, want 0", converted)
    }
}

func TestEnergyToString(t *testing.T) {
	factory := units.EnergyFactory{}
	a, err := factory.CreateEnergy(45, units.EnergyJoule)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.EnergyJoule, 2)
	expected := "45.00 " + units.GetEnergyAbbreviation(units.EnergyJoule)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.EnergyJoule, -1)
	expected = "45 " + units.GetEnergyAbbreviation(units.EnergyJoule)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestEnergy_EqualityAndComparison(t *testing.T) {
	factory := units.EnergyFactory{}
	a1, _ := factory.CreateEnergy(60, units.EnergyJoule)
	a2, _ := factory.CreateEnergy(60, units.EnergyJoule)
	a3, _ := factory.CreateEnergy(120, units.EnergyJoule)

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

func TestEnergy_Arithmetic(t *testing.T) {
	factory := units.EnergyFactory{}
	a1, _ := factory.CreateEnergy(30, units.EnergyJoule)
	a2, _ := factory.CreateEnergy(45, units.EnergyJoule)

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