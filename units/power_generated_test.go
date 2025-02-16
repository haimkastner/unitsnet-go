// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPowerDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Watt"}`
	
	factory := units.PowerDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PowerWatt {
		t.Errorf("expected unit %v, got %v", units.PowerWatt, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Watt"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPowerDto_ToJSON(t *testing.T) {
	dto := units.PowerDto{
		Value: 45,
		Unit:  units.PowerWatt,
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
	if result["unit"].(string) != string(units.PowerWatt) {
		t.Errorf("expected unit %s, got %v", units.PowerWatt, result["unit"])
	}
}

func TestNewPower_InvalidValue(t *testing.T) {
	factory := units.PowerFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePower(math.NaN(), units.PowerWatt)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePower(math.Inf(1), units.PowerWatt)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPowerConversions(t *testing.T) {
	factory := units.PowerFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePower(180, units.PowerWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Watts.
		// No expected conversion value provided for Watts, verifying result is not NaN.
		result := a.Watts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Watts returned NaN")
		}
	}
	{
		// Test conversion to MechanicalHorsepower.
		// No expected conversion value provided for MechanicalHorsepower, verifying result is not NaN.
		result := a.MechanicalHorsepower()
		if math.IsNaN(result) {
			t.Errorf("conversion to MechanicalHorsepower returned NaN")
		}
	}
	{
		// Test conversion to MetricHorsepower.
		// No expected conversion value provided for MetricHorsepower, verifying result is not NaN.
		result := a.MetricHorsepower()
		if math.IsNaN(result) {
			t.Errorf("conversion to MetricHorsepower returned NaN")
		}
	}
	{
		// Test conversion to ElectricalHorsepower.
		// No expected conversion value provided for ElectricalHorsepower, verifying result is not NaN.
		result := a.ElectricalHorsepower()
		if math.IsNaN(result) {
			t.Errorf("conversion to ElectricalHorsepower returned NaN")
		}
	}
	{
		// Test conversion to BoilerHorsepower.
		// No expected conversion value provided for BoilerHorsepower, verifying result is not NaN.
		result := a.BoilerHorsepower()
		if math.IsNaN(result) {
			t.Errorf("conversion to BoilerHorsepower returned NaN")
		}
	}
	{
		// Test conversion to HydraulicHorsepower.
		// No expected conversion value provided for HydraulicHorsepower, verifying result is not NaN.
		result := a.HydraulicHorsepower()
		if math.IsNaN(result) {
			t.Errorf("conversion to HydraulicHorsepower returned NaN")
		}
	}
	{
		// Test conversion to BritishThermalUnitsPerHour.
		// No expected conversion value provided for BritishThermalUnitsPerHour, verifying result is not NaN.
		result := a.BritishThermalUnitsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to BritishThermalUnitsPerHour returned NaN")
		}
	}
	{
		// Test conversion to JoulesPerHour.
		// No expected conversion value provided for JoulesPerHour, verifying result is not NaN.
		result := a.JoulesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to JoulesPerHour returned NaN")
		}
	}
	{
		// Test conversion to TonsOfRefrigeration.
		// No expected conversion value provided for TonsOfRefrigeration, verifying result is not NaN.
		result := a.TonsOfRefrigeration()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonsOfRefrigeration returned NaN")
		}
	}
	{
		// Test conversion to Femtowatts.
		// No expected conversion value provided for Femtowatts, verifying result is not NaN.
		result := a.Femtowatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Femtowatts returned NaN")
		}
	}
	{
		// Test conversion to Picowatts.
		// No expected conversion value provided for Picowatts, verifying result is not NaN.
		result := a.Picowatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Picowatts returned NaN")
		}
	}
	{
		// Test conversion to Nanowatts.
		// No expected conversion value provided for Nanowatts, verifying result is not NaN.
		result := a.Nanowatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanowatts returned NaN")
		}
	}
	{
		// Test conversion to Microwatts.
		// No expected conversion value provided for Microwatts, verifying result is not NaN.
		result := a.Microwatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microwatts returned NaN")
		}
	}
	{
		// Test conversion to Milliwatts.
		// No expected conversion value provided for Milliwatts, verifying result is not NaN.
		result := a.Milliwatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milliwatts returned NaN")
		}
	}
	{
		// Test conversion to Deciwatts.
		// No expected conversion value provided for Deciwatts, verifying result is not NaN.
		result := a.Deciwatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Deciwatts returned NaN")
		}
	}
	{
		// Test conversion to Decawatts.
		// No expected conversion value provided for Decawatts, verifying result is not NaN.
		result := a.Decawatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Decawatts returned NaN")
		}
	}
	{
		// Test conversion to Kilowatts.
		// No expected conversion value provided for Kilowatts, verifying result is not NaN.
		result := a.Kilowatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Kilowatts returned NaN")
		}
	}
	{
		// Test conversion to Megawatts.
		// No expected conversion value provided for Megawatts, verifying result is not NaN.
		result := a.Megawatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Megawatts returned NaN")
		}
	}
	{
		// Test conversion to Gigawatts.
		// No expected conversion value provided for Gigawatts, verifying result is not NaN.
		result := a.Gigawatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gigawatts returned NaN")
		}
	}
	{
		// Test conversion to Terawatts.
		// No expected conversion value provided for Terawatts, verifying result is not NaN.
		result := a.Terawatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Terawatts returned NaN")
		}
	}
	{
		// Test conversion to Petawatts.
		// No expected conversion value provided for Petawatts, verifying result is not NaN.
		result := a.Petawatts()
		if math.IsNaN(result) {
			t.Errorf("conversion to Petawatts returned NaN")
		}
	}
	{
		// Test conversion to KilobritishThermalUnitsPerHour.
		// No expected conversion value provided for KilobritishThermalUnitsPerHour, verifying result is not NaN.
		result := a.KilobritishThermalUnitsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilobritishThermalUnitsPerHour returned NaN")
		}
	}
	{
		// Test conversion to MegabritishThermalUnitsPerHour.
		// No expected conversion value provided for MegabritishThermalUnitsPerHour, verifying result is not NaN.
		result := a.MegabritishThermalUnitsPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegabritishThermalUnitsPerHour returned NaN")
		}
	}
	{
		// Test conversion to MillijoulesPerHour.
		// No expected conversion value provided for MillijoulesPerHour, verifying result is not NaN.
		result := a.MillijoulesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillijoulesPerHour returned NaN")
		}
	}
	{
		// Test conversion to KilojoulesPerHour.
		// No expected conversion value provided for KilojoulesPerHour, verifying result is not NaN.
		result := a.KilojoulesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilojoulesPerHour returned NaN")
		}
	}
	{
		// Test conversion to MegajoulesPerHour.
		// No expected conversion value provided for MegajoulesPerHour, verifying result is not NaN.
		result := a.MegajoulesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegajoulesPerHour returned NaN")
		}
	}
	{
		// Test conversion to GigajoulesPerHour.
		// No expected conversion value provided for GigajoulesPerHour, verifying result is not NaN.
		result := a.GigajoulesPerHour()
		if math.IsNaN(result) {
			t.Errorf("conversion to GigajoulesPerHour returned NaN")
		}
	}
}

func TestPower_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PowerFactory{}
	a, err := factory.CreatePower(90, units.PowerWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PowerWatt {
		t.Errorf("expected default unit Watt, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PowerWatt
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PowerDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PowerWatt {
		t.Errorf("expected unit Watt, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPowerFactory_FromDto(t *testing.T) {
    factory := units.PowerFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerWatt,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.PowerDto{
        Value: math.NaN(),
        Unit:  units.PowerWatt,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Watt conversion
    wattsDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerWatt,
    }
    
    var wattsResult *units.Power
    wattsResult, err = factory.FromDto(wattsDto)
    if err != nil {
        t.Errorf("FromDto() with Watt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = wattsResult.Convert(units.PowerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Watt = %v, want %v", converted, 100)
    }
    // Test MechanicalHorsepower conversion
    mechanical_horsepowerDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerMechanicalHorsepower,
    }
    
    var mechanical_horsepowerResult *units.Power
    mechanical_horsepowerResult, err = factory.FromDto(mechanical_horsepowerDto)
    if err != nil {
        t.Errorf("FromDto() with MechanicalHorsepower returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mechanical_horsepowerResult.Convert(units.PowerMechanicalHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MechanicalHorsepower = %v, want %v", converted, 100)
    }
    // Test MetricHorsepower conversion
    metric_horsepowerDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerMetricHorsepower,
    }
    
    var metric_horsepowerResult *units.Power
    metric_horsepowerResult, err = factory.FromDto(metric_horsepowerDto)
    if err != nil {
        t.Errorf("FromDto() with MetricHorsepower returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = metric_horsepowerResult.Convert(units.PowerMetricHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MetricHorsepower = %v, want %v", converted, 100)
    }
    // Test ElectricalHorsepower conversion
    electrical_horsepowerDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerElectricalHorsepower,
    }
    
    var electrical_horsepowerResult *units.Power
    electrical_horsepowerResult, err = factory.FromDto(electrical_horsepowerDto)
    if err != nil {
        t.Errorf("FromDto() with ElectricalHorsepower returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = electrical_horsepowerResult.Convert(units.PowerElectricalHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ElectricalHorsepower = %v, want %v", converted, 100)
    }
    // Test BoilerHorsepower conversion
    boiler_horsepowerDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerBoilerHorsepower,
    }
    
    var boiler_horsepowerResult *units.Power
    boiler_horsepowerResult, err = factory.FromDto(boiler_horsepowerDto)
    if err != nil {
        t.Errorf("FromDto() with BoilerHorsepower returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = boiler_horsepowerResult.Convert(units.PowerBoilerHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BoilerHorsepower = %v, want %v", converted, 100)
    }
    // Test HydraulicHorsepower conversion
    hydraulic_horsepowerDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerHydraulicHorsepower,
    }
    
    var hydraulic_horsepowerResult *units.Power
    hydraulic_horsepowerResult, err = factory.FromDto(hydraulic_horsepowerDto)
    if err != nil {
        t.Errorf("FromDto() with HydraulicHorsepower returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hydraulic_horsepowerResult.Convert(units.PowerHydraulicHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HydraulicHorsepower = %v, want %v", converted, 100)
    }
    // Test BritishThermalUnitPerHour conversion
    british_thermal_units_per_hourDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerBritishThermalUnitPerHour,
    }
    
    var british_thermal_units_per_hourResult *units.Power
    british_thermal_units_per_hourResult, err = factory.FromDto(british_thermal_units_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with BritishThermalUnitPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = british_thermal_units_per_hourResult.Convert(units.PowerBritishThermalUnitPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BritishThermalUnitPerHour = %v, want %v", converted, 100)
    }
    // Test JoulePerHour conversion
    joules_per_hourDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerJoulePerHour,
    }
    
    var joules_per_hourResult *units.Power
    joules_per_hourResult, err = factory.FromDto(joules_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with JoulePerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_hourResult.Convert(units.PowerJoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerHour = %v, want %v", converted, 100)
    }
    // Test TonOfRefrigeration conversion
    tons_of_refrigerationDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerTonOfRefrigeration,
    }
    
    var tons_of_refrigerationResult *units.Power
    tons_of_refrigerationResult, err = factory.FromDto(tons_of_refrigerationDto)
    if err != nil {
        t.Errorf("FromDto() with TonOfRefrigeration returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tons_of_refrigerationResult.Convert(units.PowerTonOfRefrigeration)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonOfRefrigeration = %v, want %v", converted, 100)
    }
    // Test Femtowatt conversion
    femtowattsDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerFemtowatt,
    }
    
    var femtowattsResult *units.Power
    femtowattsResult, err = factory.FromDto(femtowattsDto)
    if err != nil {
        t.Errorf("FromDto() with Femtowatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtowattsResult.Convert(units.PowerFemtowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtowatt = %v, want %v", converted, 100)
    }
    // Test Picowatt conversion
    picowattsDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerPicowatt,
    }
    
    var picowattsResult *units.Power
    picowattsResult, err = factory.FromDto(picowattsDto)
    if err != nil {
        t.Errorf("FromDto() with Picowatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowattsResult.Convert(units.PowerPicowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picowatt = %v, want %v", converted, 100)
    }
    // Test Nanowatt conversion
    nanowattsDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerNanowatt,
    }
    
    var nanowattsResult *units.Power
    nanowattsResult, err = factory.FromDto(nanowattsDto)
    if err != nil {
        t.Errorf("FromDto() with Nanowatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowattsResult.Convert(units.PowerNanowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanowatt = %v, want %v", converted, 100)
    }
    // Test Microwatt conversion
    microwattsDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerMicrowatt,
    }
    
    var microwattsResult *units.Power
    microwattsResult, err = factory.FromDto(microwattsDto)
    if err != nil {
        t.Errorf("FromDto() with Microwatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwattsResult.Convert(units.PowerMicrowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microwatt = %v, want %v", converted, 100)
    }
    // Test Milliwatt conversion
    milliwattsDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerMilliwatt,
    }
    
    var milliwattsResult *units.Power
    milliwattsResult, err = factory.FromDto(milliwattsDto)
    if err != nil {
        t.Errorf("FromDto() with Milliwatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwattsResult.Convert(units.PowerMilliwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliwatt = %v, want %v", converted, 100)
    }
    // Test Deciwatt conversion
    deciwattsDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerDeciwatt,
    }
    
    var deciwattsResult *units.Power
    deciwattsResult, err = factory.FromDto(deciwattsDto)
    if err != nil {
        t.Errorf("FromDto() with Deciwatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwattsResult.Convert(units.PowerDeciwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Deciwatt = %v, want %v", converted, 100)
    }
    // Test Decawatt conversion
    decawattsDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerDecawatt,
    }
    
    var decawattsResult *units.Power
    decawattsResult, err = factory.FromDto(decawattsDto)
    if err != nil {
        t.Errorf("FromDto() with Decawatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decawattsResult.Convert(units.PowerDecawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decawatt = %v, want %v", converted, 100)
    }
    // Test Kilowatt conversion
    kilowattsDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerKilowatt,
    }
    
    var kilowattsResult *units.Power
    kilowattsResult, err = factory.FromDto(kilowattsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilowatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowattsResult.Convert(units.PowerKilowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilowatt = %v, want %v", converted, 100)
    }
    // Test Megawatt conversion
    megawattsDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerMegawatt,
    }
    
    var megawattsResult *units.Power
    megawattsResult, err = factory.FromDto(megawattsDto)
    if err != nil {
        t.Errorf("FromDto() with Megawatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawattsResult.Convert(units.PowerMegawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megawatt = %v, want %v", converted, 100)
    }
    // Test Gigawatt conversion
    gigawattsDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerGigawatt,
    }
    
    var gigawattsResult *units.Power
    gigawattsResult, err = factory.FromDto(gigawattsDto)
    if err != nil {
        t.Errorf("FromDto() with Gigawatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawattsResult.Convert(units.PowerGigawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigawatt = %v, want %v", converted, 100)
    }
    // Test Terawatt conversion
    terawattsDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerTerawatt,
    }
    
    var terawattsResult *units.Power
    terawattsResult, err = factory.FromDto(terawattsDto)
    if err != nil {
        t.Errorf("FromDto() with Terawatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawattsResult.Convert(units.PowerTerawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terawatt = %v, want %v", converted, 100)
    }
    // Test Petawatt conversion
    petawattsDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerPetawatt,
    }
    
    var petawattsResult *units.Power
    petawattsResult, err = factory.FromDto(petawattsDto)
    if err != nil {
        t.Errorf("FromDto() with Petawatt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petawattsResult.Convert(units.PowerPetawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petawatt = %v, want %v", converted, 100)
    }
    // Test KilobritishThermalUnitPerHour conversion
    kilobritish_thermal_units_per_hourDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerKilobritishThermalUnitPerHour,
    }
    
    var kilobritish_thermal_units_per_hourResult *units.Power
    kilobritish_thermal_units_per_hourResult, err = factory.FromDto(kilobritish_thermal_units_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with KilobritishThermalUnitPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobritish_thermal_units_per_hourResult.Convert(units.PowerKilobritishThermalUnitPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilobritishThermalUnitPerHour = %v, want %v", converted, 100)
    }
    // Test MegabritishThermalUnitPerHour conversion
    megabritish_thermal_units_per_hourDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerMegabritishThermalUnitPerHour,
    }
    
    var megabritish_thermal_units_per_hourResult *units.Power
    megabritish_thermal_units_per_hourResult, err = factory.FromDto(megabritish_thermal_units_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MegabritishThermalUnitPerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabritish_thermal_units_per_hourResult.Convert(units.PowerMegabritishThermalUnitPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegabritishThermalUnitPerHour = %v, want %v", converted, 100)
    }
    // Test MillijoulePerHour conversion
    millijoules_per_hourDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerMillijoulePerHour,
    }
    
    var millijoules_per_hourResult *units.Power
    millijoules_per_hourResult, err = factory.FromDto(millijoules_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MillijoulePerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millijoules_per_hourResult.Convert(units.PowerMillijoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillijoulePerHour = %v, want %v", converted, 100)
    }
    // Test KilojoulePerHour conversion
    kilojoules_per_hourDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerKilojoulePerHour,
    }
    
    var kilojoules_per_hourResult *units.Power
    kilojoules_per_hourResult, err = factory.FromDto(kilojoules_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with KilojoulePerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_hourResult.Convert(units.PowerKilojoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerHour = %v, want %v", converted, 100)
    }
    // Test MegajoulePerHour conversion
    megajoules_per_hourDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerMegajoulePerHour,
    }
    
    var megajoules_per_hourResult *units.Power
    megajoules_per_hourResult, err = factory.FromDto(megajoules_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with MegajoulePerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_hourResult.Convert(units.PowerMegajoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerHour = %v, want %v", converted, 100)
    }
    // Test GigajoulePerHour conversion
    gigajoules_per_hourDto := units.PowerDto{
        Value: 100,
        Unit:  units.PowerGigajoulePerHour,
    }
    
    var gigajoules_per_hourResult *units.Power
    gigajoules_per_hourResult, err = factory.FromDto(gigajoules_per_hourDto)
    if err != nil {
        t.Errorf("FromDto() with GigajoulePerHour returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigajoules_per_hourResult.Convert(units.PowerGigajoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigajoulePerHour = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.PowerDto{
        Value: 0,
        Unit:  units.PowerWatt,
    }
    
    var zeroResult *units.Power
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestPowerFactory_FromDtoJSON(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Watt"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Watt"}`)
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
    nanJSON, _ := json.Marshal(units.PowerDto{
        Value: nanValue,
        Unit:  units.PowerWatt,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Watt unit
    wattsJSON := []byte(`{"value": 100, "unit": "Watt"}`)
    wattsResult, err := factory.FromDtoJSON(wattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Watt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = wattsResult.Convert(units.PowerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Watt = %v, want %v", converted, 100)
    }
    // Test JSON with MechanicalHorsepower unit
    mechanical_horsepowerJSON := []byte(`{"value": 100, "unit": "MechanicalHorsepower"}`)
    mechanical_horsepowerResult, err := factory.FromDtoJSON(mechanical_horsepowerJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MechanicalHorsepower unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = mechanical_horsepowerResult.Convert(units.PowerMechanicalHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MechanicalHorsepower = %v, want %v", converted, 100)
    }
    // Test JSON with MetricHorsepower unit
    metric_horsepowerJSON := []byte(`{"value": 100, "unit": "MetricHorsepower"}`)
    metric_horsepowerResult, err := factory.FromDtoJSON(metric_horsepowerJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MetricHorsepower unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = metric_horsepowerResult.Convert(units.PowerMetricHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MetricHorsepower = %v, want %v", converted, 100)
    }
    // Test JSON with ElectricalHorsepower unit
    electrical_horsepowerJSON := []byte(`{"value": 100, "unit": "ElectricalHorsepower"}`)
    electrical_horsepowerResult, err := factory.FromDtoJSON(electrical_horsepowerJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ElectricalHorsepower unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = electrical_horsepowerResult.Convert(units.PowerElectricalHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ElectricalHorsepower = %v, want %v", converted, 100)
    }
    // Test JSON with BoilerHorsepower unit
    boiler_horsepowerJSON := []byte(`{"value": 100, "unit": "BoilerHorsepower"}`)
    boiler_horsepowerResult, err := factory.FromDtoJSON(boiler_horsepowerJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BoilerHorsepower unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = boiler_horsepowerResult.Convert(units.PowerBoilerHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BoilerHorsepower = %v, want %v", converted, 100)
    }
    // Test JSON with HydraulicHorsepower unit
    hydraulic_horsepowerJSON := []byte(`{"value": 100, "unit": "HydraulicHorsepower"}`)
    hydraulic_horsepowerResult, err := factory.FromDtoJSON(hydraulic_horsepowerJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HydraulicHorsepower unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hydraulic_horsepowerResult.Convert(units.PowerHydraulicHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HydraulicHorsepower = %v, want %v", converted, 100)
    }
    // Test JSON with BritishThermalUnitPerHour unit
    british_thermal_units_per_hourJSON := []byte(`{"value": 100, "unit": "BritishThermalUnitPerHour"}`)
    british_thermal_units_per_hourResult, err := factory.FromDtoJSON(british_thermal_units_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BritishThermalUnitPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = british_thermal_units_per_hourResult.Convert(units.PowerBritishThermalUnitPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BritishThermalUnitPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with JoulePerHour unit
    joules_per_hourJSON := []byte(`{"value": 100, "unit": "JoulePerHour"}`)
    joules_per_hourResult, err := factory.FromDtoJSON(joules_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with JoulePerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = joules_per_hourResult.Convert(units.PowerJoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for JoulePerHour = %v, want %v", converted, 100)
    }
    // Test JSON with TonOfRefrigeration unit
    tons_of_refrigerationJSON := []byte(`{"value": 100, "unit": "TonOfRefrigeration"}`)
    tons_of_refrigerationResult, err := factory.FromDtoJSON(tons_of_refrigerationJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonOfRefrigeration unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tons_of_refrigerationResult.Convert(units.PowerTonOfRefrigeration)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonOfRefrigeration = %v, want %v", converted, 100)
    }
    // Test JSON with Femtowatt unit
    femtowattsJSON := []byte(`{"value": 100, "unit": "Femtowatt"}`)
    femtowattsResult, err := factory.FromDtoJSON(femtowattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Femtowatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtowattsResult.Convert(units.PowerFemtowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtowatt = %v, want %v", converted, 100)
    }
    // Test JSON with Picowatt unit
    picowattsJSON := []byte(`{"value": 100, "unit": "Picowatt"}`)
    picowattsResult, err := factory.FromDtoJSON(picowattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Picowatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowattsResult.Convert(units.PowerPicowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picowatt = %v, want %v", converted, 100)
    }
    // Test JSON with Nanowatt unit
    nanowattsJSON := []byte(`{"value": 100, "unit": "Nanowatt"}`)
    nanowattsResult, err := factory.FromDtoJSON(nanowattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanowatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowattsResult.Convert(units.PowerNanowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanowatt = %v, want %v", converted, 100)
    }
    // Test JSON with Microwatt unit
    microwattsJSON := []byte(`{"value": 100, "unit": "Microwatt"}`)
    microwattsResult, err := factory.FromDtoJSON(microwattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microwatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwattsResult.Convert(units.PowerMicrowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microwatt = %v, want %v", converted, 100)
    }
    // Test JSON with Milliwatt unit
    milliwattsJSON := []byte(`{"value": 100, "unit": "Milliwatt"}`)
    milliwattsResult, err := factory.FromDtoJSON(milliwattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Milliwatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwattsResult.Convert(units.PowerMilliwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliwatt = %v, want %v", converted, 100)
    }
    // Test JSON with Deciwatt unit
    deciwattsJSON := []byte(`{"value": 100, "unit": "Deciwatt"}`)
    deciwattsResult, err := factory.FromDtoJSON(deciwattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Deciwatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciwattsResult.Convert(units.PowerDeciwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Deciwatt = %v, want %v", converted, 100)
    }
    // Test JSON with Decawatt unit
    decawattsJSON := []byte(`{"value": 100, "unit": "Decawatt"}`)
    decawattsResult, err := factory.FromDtoJSON(decawattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Decawatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decawattsResult.Convert(units.PowerDecawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decawatt = %v, want %v", converted, 100)
    }
    // Test JSON with Kilowatt unit
    kilowattsJSON := []byte(`{"value": 100, "unit": "Kilowatt"}`)
    kilowattsResult, err := factory.FromDtoJSON(kilowattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilowatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowattsResult.Convert(units.PowerKilowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilowatt = %v, want %v", converted, 100)
    }
    // Test JSON with Megawatt unit
    megawattsJSON := []byte(`{"value": 100, "unit": "Megawatt"}`)
    megawattsResult, err := factory.FromDtoJSON(megawattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megawatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawattsResult.Convert(units.PowerMegawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megawatt = %v, want %v", converted, 100)
    }
    // Test JSON with Gigawatt unit
    gigawattsJSON := []byte(`{"value": 100, "unit": "Gigawatt"}`)
    gigawattsResult, err := factory.FromDtoJSON(gigawattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigawatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigawattsResult.Convert(units.PowerGigawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigawatt = %v, want %v", converted, 100)
    }
    // Test JSON with Terawatt unit
    terawattsJSON := []byte(`{"value": 100, "unit": "Terawatt"}`)
    terawattsResult, err := factory.FromDtoJSON(terawattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Terawatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = terawattsResult.Convert(units.PowerTerawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Terawatt = %v, want %v", converted, 100)
    }
    // Test JSON with Petawatt unit
    petawattsJSON := []byte(`{"value": 100, "unit": "Petawatt"}`)
    petawattsResult, err := factory.FromDtoJSON(petawattsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Petawatt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = petawattsResult.Convert(units.PowerPetawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Petawatt = %v, want %v", converted, 100)
    }
    // Test JSON with KilobritishThermalUnitPerHour unit
    kilobritish_thermal_units_per_hourJSON := []byte(`{"value": 100, "unit": "KilobritishThermalUnitPerHour"}`)
    kilobritish_thermal_units_per_hourResult, err := factory.FromDtoJSON(kilobritish_thermal_units_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilobritishThermalUnitPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobritish_thermal_units_per_hourResult.Convert(units.PowerKilobritishThermalUnitPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilobritishThermalUnitPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MegabritishThermalUnitPerHour unit
    megabritish_thermal_units_per_hourJSON := []byte(`{"value": 100, "unit": "MegabritishThermalUnitPerHour"}`)
    megabritish_thermal_units_per_hourResult, err := factory.FromDtoJSON(megabritish_thermal_units_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegabritishThermalUnitPerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabritish_thermal_units_per_hourResult.Convert(units.PowerMegabritishThermalUnitPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegabritishThermalUnitPerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MillijoulePerHour unit
    millijoules_per_hourJSON := []byte(`{"value": 100, "unit": "MillijoulePerHour"}`)
    millijoules_per_hourResult, err := factory.FromDtoJSON(millijoules_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillijoulePerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millijoules_per_hourResult.Convert(units.PowerMillijoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillijoulePerHour = %v, want %v", converted, 100)
    }
    // Test JSON with KilojoulePerHour unit
    kilojoules_per_hourJSON := []byte(`{"value": 100, "unit": "KilojoulePerHour"}`)
    kilojoules_per_hourResult, err := factory.FromDtoJSON(kilojoules_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilojoulePerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilojoules_per_hourResult.Convert(units.PowerKilojoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilojoulePerHour = %v, want %v", converted, 100)
    }
    // Test JSON with MegajoulePerHour unit
    megajoules_per_hourJSON := []byte(`{"value": 100, "unit": "MegajoulePerHour"}`)
    megajoules_per_hourResult, err := factory.FromDtoJSON(megajoules_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegajoulePerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megajoules_per_hourResult.Convert(units.PowerMegajoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegajoulePerHour = %v, want %v", converted, 100)
    }
    // Test JSON with GigajoulePerHour unit
    gigajoules_per_hourJSON := []byte(`{"value": 100, "unit": "GigajoulePerHour"}`)
    gigajoules_per_hourResult, err := factory.FromDtoJSON(gigajoules_per_hourJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GigajoulePerHour unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigajoules_per_hourResult.Convert(units.PowerGigajoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GigajoulePerHour = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Watt"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromWatts function
func TestPowerFactory_FromWatts(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWatts(100)
    if err != nil {
        t.Errorf("FromWatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerWatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWatts(math.NaN())
    if err == nil {
        t.Error("FromWatts() with NaN value should return error")
    }

    _, err = factory.FromWatts(math.Inf(1))
    if err == nil {
        t.Error("FromWatts() with +Inf value should return error")
    }

    _, err = factory.FromWatts(math.Inf(-1))
    if err == nil {
        t.Error("FromWatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWatts(0)
    if err != nil {
        t.Errorf("FromWatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerWatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWatts() with zero value = %v, want 0", converted)
    }
}
// Test FromMechanicalHorsepower function
func TestPowerFactory_FromMechanicalHorsepower(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMechanicalHorsepower(100)
    if err != nil {
        t.Errorf("FromMechanicalHorsepower() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerMechanicalHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMechanicalHorsepower() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMechanicalHorsepower(math.NaN())
    if err == nil {
        t.Error("FromMechanicalHorsepower() with NaN value should return error")
    }

    _, err = factory.FromMechanicalHorsepower(math.Inf(1))
    if err == nil {
        t.Error("FromMechanicalHorsepower() with +Inf value should return error")
    }

    _, err = factory.FromMechanicalHorsepower(math.Inf(-1))
    if err == nil {
        t.Error("FromMechanicalHorsepower() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMechanicalHorsepower(0)
    if err != nil {
        t.Errorf("FromMechanicalHorsepower() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerMechanicalHorsepower)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMechanicalHorsepower() with zero value = %v, want 0", converted)
    }
}
// Test FromMetricHorsepower function
func TestPowerFactory_FromMetricHorsepower(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMetricHorsepower(100)
    if err != nil {
        t.Errorf("FromMetricHorsepower() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerMetricHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMetricHorsepower() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMetricHorsepower(math.NaN())
    if err == nil {
        t.Error("FromMetricHorsepower() with NaN value should return error")
    }

    _, err = factory.FromMetricHorsepower(math.Inf(1))
    if err == nil {
        t.Error("FromMetricHorsepower() with +Inf value should return error")
    }

    _, err = factory.FromMetricHorsepower(math.Inf(-1))
    if err == nil {
        t.Error("FromMetricHorsepower() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMetricHorsepower(0)
    if err != nil {
        t.Errorf("FromMetricHorsepower() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerMetricHorsepower)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMetricHorsepower() with zero value = %v, want 0", converted)
    }
}
// Test FromElectricalHorsepower function
func TestPowerFactory_FromElectricalHorsepower(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromElectricalHorsepower(100)
    if err != nil {
        t.Errorf("FromElectricalHorsepower() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerElectricalHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromElectricalHorsepower() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromElectricalHorsepower(math.NaN())
    if err == nil {
        t.Error("FromElectricalHorsepower() with NaN value should return error")
    }

    _, err = factory.FromElectricalHorsepower(math.Inf(1))
    if err == nil {
        t.Error("FromElectricalHorsepower() with +Inf value should return error")
    }

    _, err = factory.FromElectricalHorsepower(math.Inf(-1))
    if err == nil {
        t.Error("FromElectricalHorsepower() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromElectricalHorsepower(0)
    if err != nil {
        t.Errorf("FromElectricalHorsepower() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerElectricalHorsepower)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromElectricalHorsepower() with zero value = %v, want 0", converted)
    }
}
// Test FromBoilerHorsepower function
func TestPowerFactory_FromBoilerHorsepower(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBoilerHorsepower(100)
    if err != nil {
        t.Errorf("FromBoilerHorsepower() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerBoilerHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBoilerHorsepower() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBoilerHorsepower(math.NaN())
    if err == nil {
        t.Error("FromBoilerHorsepower() with NaN value should return error")
    }

    _, err = factory.FromBoilerHorsepower(math.Inf(1))
    if err == nil {
        t.Error("FromBoilerHorsepower() with +Inf value should return error")
    }

    _, err = factory.FromBoilerHorsepower(math.Inf(-1))
    if err == nil {
        t.Error("FromBoilerHorsepower() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBoilerHorsepower(0)
    if err != nil {
        t.Errorf("FromBoilerHorsepower() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerBoilerHorsepower)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBoilerHorsepower() with zero value = %v, want 0", converted)
    }
}
// Test FromHydraulicHorsepower function
func TestPowerFactory_FromHydraulicHorsepower(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHydraulicHorsepower(100)
    if err != nil {
        t.Errorf("FromHydraulicHorsepower() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerHydraulicHorsepower)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHydraulicHorsepower() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHydraulicHorsepower(math.NaN())
    if err == nil {
        t.Error("FromHydraulicHorsepower() with NaN value should return error")
    }

    _, err = factory.FromHydraulicHorsepower(math.Inf(1))
    if err == nil {
        t.Error("FromHydraulicHorsepower() with +Inf value should return error")
    }

    _, err = factory.FromHydraulicHorsepower(math.Inf(-1))
    if err == nil {
        t.Error("FromHydraulicHorsepower() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHydraulicHorsepower(0)
    if err != nil {
        t.Errorf("FromHydraulicHorsepower() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerHydraulicHorsepower)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHydraulicHorsepower() with zero value = %v, want 0", converted)
    }
}
// Test FromBritishThermalUnitsPerHour function
func TestPowerFactory_FromBritishThermalUnitsPerHour(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBritishThermalUnitsPerHour(100)
    if err != nil {
        t.Errorf("FromBritishThermalUnitsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerBritishThermalUnitPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBritishThermalUnitsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBritishThermalUnitsPerHour(math.NaN())
    if err == nil {
        t.Error("FromBritishThermalUnitsPerHour() with NaN value should return error")
    }

    _, err = factory.FromBritishThermalUnitsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromBritishThermalUnitsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromBritishThermalUnitsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromBritishThermalUnitsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBritishThermalUnitsPerHour(0)
    if err != nil {
        t.Errorf("FromBritishThermalUnitsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerBritishThermalUnitPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBritishThermalUnitsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromJoulesPerHour function
func TestPowerFactory_FromJoulesPerHour(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromJoulesPerHour(100)
    if err != nil {
        t.Errorf("FromJoulesPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerJoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromJoulesPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromJoulesPerHour(math.NaN())
    if err == nil {
        t.Error("FromJoulesPerHour() with NaN value should return error")
    }

    _, err = factory.FromJoulesPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromJoulesPerHour() with +Inf value should return error")
    }

    _, err = factory.FromJoulesPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromJoulesPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromJoulesPerHour(0)
    if err != nil {
        t.Errorf("FromJoulesPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerJoulePerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromJoulesPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromTonsOfRefrigeration function
func TestPowerFactory_FromTonsOfRefrigeration(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonsOfRefrigeration(100)
    if err != nil {
        t.Errorf("FromTonsOfRefrigeration() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerTonOfRefrigeration)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonsOfRefrigeration() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonsOfRefrigeration(math.NaN())
    if err == nil {
        t.Error("FromTonsOfRefrigeration() with NaN value should return error")
    }

    _, err = factory.FromTonsOfRefrigeration(math.Inf(1))
    if err == nil {
        t.Error("FromTonsOfRefrigeration() with +Inf value should return error")
    }

    _, err = factory.FromTonsOfRefrigeration(math.Inf(-1))
    if err == nil {
        t.Error("FromTonsOfRefrigeration() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonsOfRefrigeration(0)
    if err != nil {
        t.Errorf("FromTonsOfRefrigeration() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerTonOfRefrigeration)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonsOfRefrigeration() with zero value = %v, want 0", converted)
    }
}
// Test FromFemtowatts function
func TestPowerFactory_FromFemtowatts(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFemtowatts(100)
    if err != nil {
        t.Errorf("FromFemtowatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerFemtowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFemtowatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFemtowatts(math.NaN())
    if err == nil {
        t.Error("FromFemtowatts() with NaN value should return error")
    }

    _, err = factory.FromFemtowatts(math.Inf(1))
    if err == nil {
        t.Error("FromFemtowatts() with +Inf value should return error")
    }

    _, err = factory.FromFemtowatts(math.Inf(-1))
    if err == nil {
        t.Error("FromFemtowatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFemtowatts(0)
    if err != nil {
        t.Errorf("FromFemtowatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerFemtowatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFemtowatts() with zero value = %v, want 0", converted)
    }
}
// Test FromPicowatts function
func TestPowerFactory_FromPicowatts(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicowatts(100)
    if err != nil {
        t.Errorf("FromPicowatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerPicowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicowatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicowatts(math.NaN())
    if err == nil {
        t.Error("FromPicowatts() with NaN value should return error")
    }

    _, err = factory.FromPicowatts(math.Inf(1))
    if err == nil {
        t.Error("FromPicowatts() with +Inf value should return error")
    }

    _, err = factory.FromPicowatts(math.Inf(-1))
    if err == nil {
        t.Error("FromPicowatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicowatts(0)
    if err != nil {
        t.Errorf("FromPicowatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerPicowatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicowatts() with zero value = %v, want 0", converted)
    }
}
// Test FromNanowatts function
func TestPowerFactory_FromNanowatts(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanowatts(100)
    if err != nil {
        t.Errorf("FromNanowatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerNanowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanowatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanowatts(math.NaN())
    if err == nil {
        t.Error("FromNanowatts() with NaN value should return error")
    }

    _, err = factory.FromNanowatts(math.Inf(1))
    if err == nil {
        t.Error("FromNanowatts() with +Inf value should return error")
    }

    _, err = factory.FromNanowatts(math.Inf(-1))
    if err == nil {
        t.Error("FromNanowatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanowatts(0)
    if err != nil {
        t.Errorf("FromNanowatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerNanowatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanowatts() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrowatts function
func TestPowerFactory_FromMicrowatts(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrowatts(100)
    if err != nil {
        t.Errorf("FromMicrowatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerMicrowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrowatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrowatts(math.NaN())
    if err == nil {
        t.Error("FromMicrowatts() with NaN value should return error")
    }

    _, err = factory.FromMicrowatts(math.Inf(1))
    if err == nil {
        t.Error("FromMicrowatts() with +Inf value should return error")
    }

    _, err = factory.FromMicrowatts(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrowatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrowatts(0)
    if err != nil {
        t.Errorf("FromMicrowatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerMicrowatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrowatts() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwatts function
func TestPowerFactory_FromMilliwatts(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwatts(100)
    if err != nil {
        t.Errorf("FromMilliwatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerMilliwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwatts(math.NaN())
    if err == nil {
        t.Error("FromMilliwatts() with NaN value should return error")
    }

    _, err = factory.FromMilliwatts(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwatts() with +Inf value should return error")
    }

    _, err = factory.FromMilliwatts(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwatts(0)
    if err != nil {
        t.Errorf("FromMilliwatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerMilliwatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwatts() with zero value = %v, want 0", converted)
    }
}
// Test FromDeciwatts function
func TestPowerFactory_FromDeciwatts(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDeciwatts(100)
    if err != nil {
        t.Errorf("FromDeciwatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDeciwatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDeciwatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDeciwatts(math.NaN())
    if err == nil {
        t.Error("FromDeciwatts() with NaN value should return error")
    }

    _, err = factory.FromDeciwatts(math.Inf(1))
    if err == nil {
        t.Error("FromDeciwatts() with +Inf value should return error")
    }

    _, err = factory.FromDeciwatts(math.Inf(-1))
    if err == nil {
        t.Error("FromDeciwatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDeciwatts(0)
    if err != nil {
        t.Errorf("FromDeciwatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDeciwatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDeciwatts() with zero value = %v, want 0", converted)
    }
}
// Test FromDecawatts function
func TestPowerFactory_FromDecawatts(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecawatts(100)
    if err != nil {
        t.Errorf("FromDecawatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerDecawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecawatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecawatts(math.NaN())
    if err == nil {
        t.Error("FromDecawatts() with NaN value should return error")
    }

    _, err = factory.FromDecawatts(math.Inf(1))
    if err == nil {
        t.Error("FromDecawatts() with +Inf value should return error")
    }

    _, err = factory.FromDecawatts(math.Inf(-1))
    if err == nil {
        t.Error("FromDecawatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecawatts(0)
    if err != nil {
        t.Errorf("FromDecawatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerDecawatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecawatts() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowatts function
func TestPowerFactory_FromKilowatts(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowatts(100)
    if err != nil {
        t.Errorf("FromKilowatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerKilowatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowatts(math.NaN())
    if err == nil {
        t.Error("FromKilowatts() with NaN value should return error")
    }

    _, err = factory.FromKilowatts(math.Inf(1))
    if err == nil {
        t.Error("FromKilowatts() with +Inf value should return error")
    }

    _, err = factory.FromKilowatts(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowatts(0)
    if err != nil {
        t.Errorf("FromKilowatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerKilowatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowatts() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawatts function
func TestPowerFactory_FromMegawatts(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawatts(100)
    if err != nil {
        t.Errorf("FromMegawatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerMegawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawatts(math.NaN())
    if err == nil {
        t.Error("FromMegawatts() with NaN value should return error")
    }

    _, err = factory.FromMegawatts(math.Inf(1))
    if err == nil {
        t.Error("FromMegawatts() with +Inf value should return error")
    }

    _, err = factory.FromMegawatts(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawatts(0)
    if err != nil {
        t.Errorf("FromMegawatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerMegawatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawatts() with zero value = %v, want 0", converted)
    }
}
// Test FromGigawatts function
func TestPowerFactory_FromGigawatts(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigawatts(100)
    if err != nil {
        t.Errorf("FromGigawatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerGigawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigawatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigawatts(math.NaN())
    if err == nil {
        t.Error("FromGigawatts() with NaN value should return error")
    }

    _, err = factory.FromGigawatts(math.Inf(1))
    if err == nil {
        t.Error("FromGigawatts() with +Inf value should return error")
    }

    _, err = factory.FromGigawatts(math.Inf(-1))
    if err == nil {
        t.Error("FromGigawatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigawatts(0)
    if err != nil {
        t.Errorf("FromGigawatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerGigawatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigawatts() with zero value = %v, want 0", converted)
    }
}
// Test FromTerawatts function
func TestPowerFactory_FromTerawatts(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTerawatts(100)
    if err != nil {
        t.Errorf("FromTerawatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerTerawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTerawatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTerawatts(math.NaN())
    if err == nil {
        t.Error("FromTerawatts() with NaN value should return error")
    }

    _, err = factory.FromTerawatts(math.Inf(1))
    if err == nil {
        t.Error("FromTerawatts() with +Inf value should return error")
    }

    _, err = factory.FromTerawatts(math.Inf(-1))
    if err == nil {
        t.Error("FromTerawatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTerawatts(0)
    if err != nil {
        t.Errorf("FromTerawatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerTerawatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTerawatts() with zero value = %v, want 0", converted)
    }
}
// Test FromPetawatts function
func TestPowerFactory_FromPetawatts(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPetawatts(100)
    if err != nil {
        t.Errorf("FromPetawatts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerPetawatt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPetawatts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPetawatts(math.NaN())
    if err == nil {
        t.Error("FromPetawatts() with NaN value should return error")
    }

    _, err = factory.FromPetawatts(math.Inf(1))
    if err == nil {
        t.Error("FromPetawatts() with +Inf value should return error")
    }

    _, err = factory.FromPetawatts(math.Inf(-1))
    if err == nil {
        t.Error("FromPetawatts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPetawatts(0)
    if err != nil {
        t.Errorf("FromPetawatts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerPetawatt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPetawatts() with zero value = %v, want 0", converted)
    }
}
// Test FromKilobritishThermalUnitsPerHour function
func TestPowerFactory_FromKilobritishThermalUnitsPerHour(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilobritishThermalUnitsPerHour(100)
    if err != nil {
        t.Errorf("FromKilobritishThermalUnitsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerKilobritishThermalUnitPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilobritishThermalUnitsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilobritishThermalUnitsPerHour(math.NaN())
    if err == nil {
        t.Error("FromKilobritishThermalUnitsPerHour() with NaN value should return error")
    }

    _, err = factory.FromKilobritishThermalUnitsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromKilobritishThermalUnitsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromKilobritishThermalUnitsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromKilobritishThermalUnitsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilobritishThermalUnitsPerHour(0)
    if err != nil {
        t.Errorf("FromKilobritishThermalUnitsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerKilobritishThermalUnitPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilobritishThermalUnitsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromMegabritishThermalUnitsPerHour function
func TestPowerFactory_FromMegabritishThermalUnitsPerHour(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegabritishThermalUnitsPerHour(100)
    if err != nil {
        t.Errorf("FromMegabritishThermalUnitsPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerMegabritishThermalUnitPerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegabritishThermalUnitsPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegabritishThermalUnitsPerHour(math.NaN())
    if err == nil {
        t.Error("FromMegabritishThermalUnitsPerHour() with NaN value should return error")
    }

    _, err = factory.FromMegabritishThermalUnitsPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMegabritishThermalUnitsPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMegabritishThermalUnitsPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMegabritishThermalUnitsPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegabritishThermalUnitsPerHour(0)
    if err != nil {
        t.Errorf("FromMegabritishThermalUnitsPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerMegabritishThermalUnitPerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegabritishThermalUnitsPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromMillijoulesPerHour function
func TestPowerFactory_FromMillijoulesPerHour(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillijoulesPerHour(100)
    if err != nil {
        t.Errorf("FromMillijoulesPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerMillijoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillijoulesPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillijoulesPerHour(math.NaN())
    if err == nil {
        t.Error("FromMillijoulesPerHour() with NaN value should return error")
    }

    _, err = factory.FromMillijoulesPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMillijoulesPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMillijoulesPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMillijoulesPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillijoulesPerHour(0)
    if err != nil {
        t.Errorf("FromMillijoulesPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerMillijoulePerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillijoulesPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromKilojoulesPerHour function
func TestPowerFactory_FromKilojoulesPerHour(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilojoulesPerHour(100)
    if err != nil {
        t.Errorf("FromKilojoulesPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerKilojoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilojoulesPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilojoulesPerHour(math.NaN())
    if err == nil {
        t.Error("FromKilojoulesPerHour() with NaN value should return error")
    }

    _, err = factory.FromKilojoulesPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromKilojoulesPerHour() with +Inf value should return error")
    }

    _, err = factory.FromKilojoulesPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromKilojoulesPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilojoulesPerHour(0)
    if err != nil {
        t.Errorf("FromKilojoulesPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerKilojoulePerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilojoulesPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromMegajoulesPerHour function
func TestPowerFactory_FromMegajoulesPerHour(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegajoulesPerHour(100)
    if err != nil {
        t.Errorf("FromMegajoulesPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerMegajoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegajoulesPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegajoulesPerHour(math.NaN())
    if err == nil {
        t.Error("FromMegajoulesPerHour() with NaN value should return error")
    }

    _, err = factory.FromMegajoulesPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromMegajoulesPerHour() with +Inf value should return error")
    }

    _, err = factory.FromMegajoulesPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromMegajoulesPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegajoulesPerHour(0)
    if err != nil {
        t.Errorf("FromMegajoulesPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerMegajoulePerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegajoulesPerHour() with zero value = %v, want 0", converted)
    }
}
// Test FromGigajoulesPerHour function
func TestPowerFactory_FromGigajoulesPerHour(t *testing.T) {
    factory := units.PowerFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigajoulesPerHour(100)
    if err != nil {
        t.Errorf("FromGigajoulesPerHour() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PowerGigajoulePerHour)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigajoulesPerHour() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigajoulesPerHour(math.NaN())
    if err == nil {
        t.Error("FromGigajoulesPerHour() with NaN value should return error")
    }

    _, err = factory.FromGigajoulesPerHour(math.Inf(1))
    if err == nil {
        t.Error("FromGigajoulesPerHour() with +Inf value should return error")
    }

    _, err = factory.FromGigajoulesPerHour(math.Inf(-1))
    if err == nil {
        t.Error("FromGigajoulesPerHour() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigajoulesPerHour(0)
    if err != nil {
        t.Errorf("FromGigajoulesPerHour() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PowerGigajoulePerHour)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigajoulesPerHour() with zero value = %v, want 0", converted)
    }
}

func TestPowerToString(t *testing.T) {
	factory := units.PowerFactory{}
	a, err := factory.CreatePower(45, units.PowerWatt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PowerWatt, 2)
	expected := "45.00 " + units.GetPowerAbbreviation(units.PowerWatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PowerWatt, -1)
	expected = "45 " + units.GetPowerAbbreviation(units.PowerWatt)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPower_EqualityAndComparison(t *testing.T) {
	factory := units.PowerFactory{}
	a1, _ := factory.CreatePower(60, units.PowerWatt)
	a2, _ := factory.CreatePower(60, units.PowerWatt)
	a3, _ := factory.CreatePower(120, units.PowerWatt)

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

func TestPower_Arithmetic(t *testing.T) {
	factory := units.PowerFactory{}
	a1, _ := factory.CreatePower(30, units.PowerWatt)
	a2, _ := factory.CreatePower(45, units.PowerWatt)

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