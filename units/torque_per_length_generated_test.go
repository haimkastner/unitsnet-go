// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestTorquePerLengthDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonMeterPerMeter"}`
	
	factory := units.TorquePerLengthDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.TorquePerLengthNewtonMeterPerMeter {
		t.Errorf("expected unit %v, got %v", units.TorquePerLengthNewtonMeterPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonMeterPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestTorquePerLengthDto_ToJSON(t *testing.T) {
	dto := units.TorquePerLengthDto{
		Value: 45,
		Unit:  units.TorquePerLengthNewtonMeterPerMeter,
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
	if result["unit"].(string) != string(units.TorquePerLengthNewtonMeterPerMeter) {
		t.Errorf("expected unit %s, got %v", units.TorquePerLengthNewtonMeterPerMeter, result["unit"])
	}
}

func TestNewTorquePerLength_InvalidValue(t *testing.T) {
	factory := units.TorquePerLengthFactory{}
	// NaN value should return an error.
	_, err := factory.CreateTorquePerLength(math.NaN(), units.TorquePerLengthNewtonMeterPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateTorquePerLength(math.Inf(1), units.TorquePerLengthNewtonMeterPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestTorquePerLengthConversions(t *testing.T) {
	factory := units.TorquePerLengthFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateTorquePerLength(180, units.TorquePerLengthNewtonMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonMillimetersPerMeter.
		// No expected conversion value provided for NewtonMillimetersPerMeter, verifying result is not NaN.
		result := a.NewtonMillimetersPerMeter()
		cacheResult := a.NewtonMillimetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonMillimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonCentimetersPerMeter.
		// No expected conversion value provided for NewtonCentimetersPerMeter, verifying result is not NaN.
		result := a.NewtonCentimetersPerMeter()
		cacheResult := a.NewtonCentimetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonCentimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonMetersPerMeter.
		// No expected conversion value provided for NewtonMetersPerMeter, verifying result is not NaN.
		result := a.NewtonMetersPerMeter()
		cacheResult := a.NewtonMetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonMetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to PoundForceInchesPerFoot.
		// No expected conversion value provided for PoundForceInchesPerFoot, verifying result is not NaN.
		result := a.PoundForceInchesPerFoot()
		cacheResult := a.PoundForceInchesPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundForceInchesPerFoot returned NaN")
		}
	}
	{
		// Test conversion to PoundForceFeetPerFoot.
		// No expected conversion value provided for PoundForceFeetPerFoot, verifying result is not NaN.
		result := a.PoundForceFeetPerFoot()
		cacheResult := a.PoundForceFeetPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundForceFeetPerFoot returned NaN")
		}
	}
	{
		// Test conversion to KilogramForceMillimetersPerMeter.
		// No expected conversion value provided for KilogramForceMillimetersPerMeter, verifying result is not NaN.
		result := a.KilogramForceMillimetersPerMeter()
		cacheResult := a.KilogramForceMillimetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramForceMillimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramForceCentimetersPerMeter.
		// No expected conversion value provided for KilogramForceCentimetersPerMeter, verifying result is not NaN.
		result := a.KilogramForceCentimetersPerMeter()
		cacheResult := a.KilogramForceCentimetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramForceCentimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramForceMetersPerMeter.
		// No expected conversion value provided for KilogramForceMetersPerMeter, verifying result is not NaN.
		result := a.KilogramForceMetersPerMeter()
		cacheResult := a.KilogramForceMetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramForceMetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to TonneForceMillimetersPerMeter.
		// No expected conversion value provided for TonneForceMillimetersPerMeter, verifying result is not NaN.
		result := a.TonneForceMillimetersPerMeter()
		cacheResult := a.TonneForceMillimetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonneForceMillimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to TonneForceCentimetersPerMeter.
		// No expected conversion value provided for TonneForceCentimetersPerMeter, verifying result is not NaN.
		result := a.TonneForceCentimetersPerMeter()
		cacheResult := a.TonneForceCentimetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonneForceCentimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to TonneForceMetersPerMeter.
		// No expected conversion value provided for TonneForceMetersPerMeter, verifying result is not NaN.
		result := a.TonneForceMetersPerMeter()
		cacheResult := a.TonneForceMetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonneForceMetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMillimetersPerMeter.
		// No expected conversion value provided for KilonewtonMillimetersPerMeter, verifying result is not NaN.
		result := a.KilonewtonMillimetersPerMeter()
		cacheResult := a.KilonewtonMillimetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonMillimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMillimetersPerMeter.
		// No expected conversion value provided for MeganewtonMillimetersPerMeter, verifying result is not NaN.
		result := a.MeganewtonMillimetersPerMeter()
		cacheResult := a.MeganewtonMillimetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MeganewtonMillimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonCentimetersPerMeter.
		// No expected conversion value provided for KilonewtonCentimetersPerMeter, verifying result is not NaN.
		result := a.KilonewtonCentimetersPerMeter()
		cacheResult := a.KilonewtonCentimetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonCentimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonCentimetersPerMeter.
		// No expected conversion value provided for MeganewtonCentimetersPerMeter, verifying result is not NaN.
		result := a.MeganewtonCentimetersPerMeter()
		cacheResult := a.MeganewtonCentimetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MeganewtonCentimetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMetersPerMeter.
		// No expected conversion value provided for KilonewtonMetersPerMeter, verifying result is not NaN.
		result := a.KilonewtonMetersPerMeter()
		cacheResult := a.KilonewtonMetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonMetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMetersPerMeter.
		// No expected conversion value provided for MeganewtonMetersPerMeter, verifying result is not NaN.
		result := a.MeganewtonMetersPerMeter()
		cacheResult := a.MeganewtonMetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MeganewtonMetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to KilopoundForceInchesPerFoot.
		// No expected conversion value provided for KilopoundForceInchesPerFoot, verifying result is not NaN.
		result := a.KilopoundForceInchesPerFoot()
		cacheResult := a.KilopoundForceInchesPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundForceInchesPerFoot returned NaN")
		}
	}
	{
		// Test conversion to MegapoundForceInchesPerFoot.
		// No expected conversion value provided for MegapoundForceInchesPerFoot, verifying result is not NaN.
		result := a.MegapoundForceInchesPerFoot()
		cacheResult := a.MegapoundForceInchesPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegapoundForceInchesPerFoot returned NaN")
		}
	}
	{
		// Test conversion to KilopoundForceFeetPerFoot.
		// No expected conversion value provided for KilopoundForceFeetPerFoot, verifying result is not NaN.
		result := a.KilopoundForceFeetPerFoot()
		cacheResult := a.KilopoundForceFeetPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundForceFeetPerFoot returned NaN")
		}
	}
	{
		// Test conversion to MegapoundForceFeetPerFoot.
		// No expected conversion value provided for MegapoundForceFeetPerFoot, verifying result is not NaN.
		result := a.MegapoundForceFeetPerFoot()
		cacheResult := a.MegapoundForceFeetPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegapoundForceFeetPerFoot returned NaN")
		}
	}
}

func TestTorquePerLength_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.TorquePerLengthFactory{}
	a, err := factory.CreateTorquePerLength(90, units.TorquePerLengthNewtonMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.TorquePerLengthNewtonMeterPerMeter {
		t.Errorf("expected default unit NewtonMeterPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.TorquePerLengthNewtonMillimeterPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.TorquePerLengthDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.TorquePerLengthNewtonMeterPerMeter {
		t.Errorf("expected unit NewtonMeterPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestTorquePerLengthFactory_FromDto(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthNewtonMeterPerMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.TorquePerLengthDto{
        Value: math.NaN(),
        Unit:  units.TorquePerLengthNewtonMeterPerMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test NewtonMillimeterPerMeter conversion
    newton_millimeters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthNewtonMillimeterPerMeter,
    }
    
    var newton_millimeters_per_meterResult *units.TorquePerLength
    newton_millimeters_per_meterResult, err = factory.FromDto(newton_millimeters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonMillimeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_millimeters_per_meterResult.Convert(units.TorquePerLengthNewtonMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMillimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test NewtonCentimeterPerMeter conversion
    newton_centimeters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthNewtonCentimeterPerMeter,
    }
    
    var newton_centimeters_per_meterResult *units.TorquePerLength
    newton_centimeters_per_meterResult, err = factory.FromDto(newton_centimeters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonCentimeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_centimeters_per_meterResult.Convert(units.TorquePerLengthNewtonCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonCentimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test NewtonMeterPerMeter conversion
    newton_meters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthNewtonMeterPerMeter,
    }
    
    var newton_meters_per_meterResult *units.TorquePerLength
    newton_meters_per_meterResult, err = factory.FromDto(newton_meters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonMeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_meters_per_meterResult.Convert(units.TorquePerLengthNewtonMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test PoundForceInchPerFoot conversion
    pound_force_inches_per_footDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthPoundForceInchPerFoot,
    }
    
    var pound_force_inches_per_footResult *units.TorquePerLength
    pound_force_inches_per_footResult, err = factory.FromDto(pound_force_inches_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForceInchPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_inches_per_footResult.Convert(units.TorquePerLengthPoundForceInchPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceInchPerFoot = %v, want %v", converted, 100)
    }
    // Test PoundForceFootPerFoot conversion
    pound_force_feet_per_footDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthPoundForceFootPerFoot,
    }
    
    var pound_force_feet_per_footResult *units.TorquePerLength
    pound_force_feet_per_footResult, err = factory.FromDto(pound_force_feet_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForceFootPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_feet_per_footResult.Convert(units.TorquePerLengthPoundForceFootPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceFootPerFoot = %v, want %v", converted, 100)
    }
    // Test KilogramForceMillimeterPerMeter conversion
    kilogram_force_millimeters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthKilogramForceMillimeterPerMeter,
    }
    
    var kilogram_force_millimeters_per_meterResult *units.TorquePerLength
    kilogram_force_millimeters_per_meterResult, err = factory.FromDto(kilogram_force_millimeters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForceMillimeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_force_millimeters_per_meterResult.Convert(units.TorquePerLengthKilogramForceMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForceMillimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test KilogramForceCentimeterPerMeter conversion
    kilogram_force_centimeters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthKilogramForceCentimeterPerMeter,
    }
    
    var kilogram_force_centimeters_per_meterResult *units.TorquePerLength
    kilogram_force_centimeters_per_meterResult, err = factory.FromDto(kilogram_force_centimeters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForceCentimeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_force_centimeters_per_meterResult.Convert(units.TorquePerLengthKilogramForceCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForceCentimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test KilogramForceMeterPerMeter conversion
    kilogram_force_meters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthKilogramForceMeterPerMeter,
    }
    
    var kilogram_force_meters_per_meterResult *units.TorquePerLength
    kilogram_force_meters_per_meterResult, err = factory.FromDto(kilogram_force_meters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForceMeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_force_meters_per_meterResult.Convert(units.TorquePerLengthKilogramForceMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForceMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test TonneForceMillimeterPerMeter conversion
    tonne_force_millimeters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthTonneForceMillimeterPerMeter,
    }
    
    var tonne_force_millimeters_per_meterResult *units.TorquePerLength
    tonne_force_millimeters_per_meterResult, err = factory.FromDto(tonne_force_millimeters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForceMillimeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_force_millimeters_per_meterResult.Convert(units.TorquePerLengthTonneForceMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForceMillimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test TonneForceCentimeterPerMeter conversion
    tonne_force_centimeters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthTonneForceCentimeterPerMeter,
    }
    
    var tonne_force_centimeters_per_meterResult *units.TorquePerLength
    tonne_force_centimeters_per_meterResult, err = factory.FromDto(tonne_force_centimeters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForceCentimeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_force_centimeters_per_meterResult.Convert(units.TorquePerLengthTonneForceCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForceCentimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test TonneForceMeterPerMeter conversion
    tonne_force_meters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthTonneForceMeterPerMeter,
    }
    
    var tonne_force_meters_per_meterResult *units.TorquePerLength
    tonne_force_meters_per_meterResult, err = factory.FromDto(tonne_force_meters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForceMeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_force_meters_per_meterResult.Convert(units.TorquePerLengthTonneForceMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForceMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonMillimeterPerMeter conversion
    kilonewton_millimeters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthKilonewtonMillimeterPerMeter,
    }
    
    var kilonewton_millimeters_per_meterResult *units.TorquePerLength
    kilonewton_millimeters_per_meterResult, err = factory.FromDto(kilonewton_millimeters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonMillimeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_millimeters_per_meterResult.Convert(units.TorquePerLengthKilonewtonMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMillimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test MeganewtonMillimeterPerMeter conversion
    meganewton_millimeters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthMeganewtonMillimeterPerMeter,
    }
    
    var meganewton_millimeters_per_meterResult *units.TorquePerLength
    meganewton_millimeters_per_meterResult, err = factory.FromDto(meganewton_millimeters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonMillimeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_millimeters_per_meterResult.Convert(units.TorquePerLengthMeganewtonMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMillimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonCentimeterPerMeter conversion
    kilonewton_centimeters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthKilonewtonCentimeterPerMeter,
    }
    
    var kilonewton_centimeters_per_meterResult *units.TorquePerLength
    kilonewton_centimeters_per_meterResult, err = factory.FromDto(kilonewton_centimeters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonCentimeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_centimeters_per_meterResult.Convert(units.TorquePerLengthKilonewtonCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonCentimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test MeganewtonCentimeterPerMeter conversion
    meganewton_centimeters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthMeganewtonCentimeterPerMeter,
    }
    
    var meganewton_centimeters_per_meterResult *units.TorquePerLength
    meganewton_centimeters_per_meterResult, err = factory.FromDto(meganewton_centimeters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonCentimeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_centimeters_per_meterResult.Convert(units.TorquePerLengthMeganewtonCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonCentimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonMeterPerMeter conversion
    kilonewton_meters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthKilonewtonMeterPerMeter,
    }
    
    var kilonewton_meters_per_meterResult *units.TorquePerLength
    kilonewton_meters_per_meterResult, err = factory.FromDto(kilonewton_meters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonMeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_meters_per_meterResult.Convert(units.TorquePerLengthKilonewtonMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test MeganewtonMeterPerMeter conversion
    meganewton_meters_per_meterDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthMeganewtonMeterPerMeter,
    }
    
    var meganewton_meters_per_meterResult *units.TorquePerLength
    meganewton_meters_per_meterResult, err = factory.FromDto(meganewton_meters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonMeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_meters_per_meterResult.Convert(units.TorquePerLengthMeganewtonMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test KilopoundForceInchPerFoot conversion
    kilopound_force_inches_per_footDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthKilopoundForceInchPerFoot,
    }
    
    var kilopound_force_inches_per_footResult *units.TorquePerLength
    kilopound_force_inches_per_footResult, err = factory.FromDto(kilopound_force_inches_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForceInchPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_force_inches_per_footResult.Convert(units.TorquePerLengthKilopoundForceInchPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForceInchPerFoot = %v, want %v", converted, 100)
    }
    // Test MegapoundForceInchPerFoot conversion
    megapound_force_inches_per_footDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthMegapoundForceInchPerFoot,
    }
    
    var megapound_force_inches_per_footResult *units.TorquePerLength
    megapound_force_inches_per_footResult, err = factory.FromDto(megapound_force_inches_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with MegapoundForceInchPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapound_force_inches_per_footResult.Convert(units.TorquePerLengthMegapoundForceInchPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundForceInchPerFoot = %v, want %v", converted, 100)
    }
    // Test KilopoundForceFootPerFoot conversion
    kilopound_force_feet_per_footDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthKilopoundForceFootPerFoot,
    }
    
    var kilopound_force_feet_per_footResult *units.TorquePerLength
    kilopound_force_feet_per_footResult, err = factory.FromDto(kilopound_force_feet_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForceFootPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_force_feet_per_footResult.Convert(units.TorquePerLengthKilopoundForceFootPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForceFootPerFoot = %v, want %v", converted, 100)
    }
    // Test MegapoundForceFootPerFoot conversion
    megapound_force_feet_per_footDto := units.TorquePerLengthDto{
        Value: 100,
        Unit:  units.TorquePerLengthMegapoundForceFootPerFoot,
    }
    
    var megapound_force_feet_per_footResult *units.TorquePerLength
    megapound_force_feet_per_footResult, err = factory.FromDto(megapound_force_feet_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with MegapoundForceFootPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapound_force_feet_per_footResult.Convert(units.TorquePerLengthMegapoundForceFootPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundForceFootPerFoot = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.TorquePerLengthDto{
        Value: 0,
        Unit:  units.TorquePerLengthNewtonMeterPerMeter,
    }
    
    var zeroResult *units.TorquePerLength
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestTorquePerLengthFactory_FromDtoJSON(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "NewtonMeterPerMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "NewtonMeterPerMeter"}`)
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
    nanJSON, _ := json.Marshal(units.TorquePerLengthDto{
        Value: nanValue,
        Unit:  units.TorquePerLengthNewtonMeterPerMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with NewtonMillimeterPerMeter unit
    newton_millimeters_per_meterJSON := []byte(`{"value": 100, "unit": "NewtonMillimeterPerMeter"}`)
    newton_millimeters_per_meterResult, err := factory.FromDtoJSON(newton_millimeters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonMillimeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_millimeters_per_meterResult.Convert(units.TorquePerLengthNewtonMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMillimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonCentimeterPerMeter unit
    newton_centimeters_per_meterJSON := []byte(`{"value": 100, "unit": "NewtonCentimeterPerMeter"}`)
    newton_centimeters_per_meterResult, err := factory.FromDtoJSON(newton_centimeters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonCentimeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_centimeters_per_meterResult.Convert(units.TorquePerLengthNewtonCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonCentimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonMeterPerMeter unit
    newton_meters_per_meterJSON := []byte(`{"value": 100, "unit": "NewtonMeterPerMeter"}`)
    newton_meters_per_meterResult, err := factory.FromDtoJSON(newton_meters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonMeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_meters_per_meterResult.Convert(units.TorquePerLengthNewtonMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForceInchPerFoot unit
    pound_force_inches_per_footJSON := []byte(`{"value": 100, "unit": "PoundForceInchPerFoot"}`)
    pound_force_inches_per_footResult, err := factory.FromDtoJSON(pound_force_inches_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForceInchPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_inches_per_footResult.Convert(units.TorquePerLengthPoundForceInchPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceInchPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForceFootPerFoot unit
    pound_force_feet_per_footJSON := []byte(`{"value": 100, "unit": "PoundForceFootPerFoot"}`)
    pound_force_feet_per_footResult, err := factory.FromDtoJSON(pound_force_feet_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForceFootPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_feet_per_footResult.Convert(units.TorquePerLengthPoundForceFootPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceFootPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForceMillimeterPerMeter unit
    kilogram_force_millimeters_per_meterJSON := []byte(`{"value": 100, "unit": "KilogramForceMillimeterPerMeter"}`)
    kilogram_force_millimeters_per_meterResult, err := factory.FromDtoJSON(kilogram_force_millimeters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForceMillimeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_force_millimeters_per_meterResult.Convert(units.TorquePerLengthKilogramForceMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForceMillimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForceCentimeterPerMeter unit
    kilogram_force_centimeters_per_meterJSON := []byte(`{"value": 100, "unit": "KilogramForceCentimeterPerMeter"}`)
    kilogram_force_centimeters_per_meterResult, err := factory.FromDtoJSON(kilogram_force_centimeters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForceCentimeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_force_centimeters_per_meterResult.Convert(units.TorquePerLengthKilogramForceCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForceCentimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForceMeterPerMeter unit
    kilogram_force_meters_per_meterJSON := []byte(`{"value": 100, "unit": "KilogramForceMeterPerMeter"}`)
    kilogram_force_meters_per_meterResult, err := factory.FromDtoJSON(kilogram_force_meters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForceMeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_force_meters_per_meterResult.Convert(units.TorquePerLengthKilogramForceMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForceMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForceMillimeterPerMeter unit
    tonne_force_millimeters_per_meterJSON := []byte(`{"value": 100, "unit": "TonneForceMillimeterPerMeter"}`)
    tonne_force_millimeters_per_meterResult, err := factory.FromDtoJSON(tonne_force_millimeters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForceMillimeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_force_millimeters_per_meterResult.Convert(units.TorquePerLengthTonneForceMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForceMillimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForceCentimeterPerMeter unit
    tonne_force_centimeters_per_meterJSON := []byte(`{"value": 100, "unit": "TonneForceCentimeterPerMeter"}`)
    tonne_force_centimeters_per_meterResult, err := factory.FromDtoJSON(tonne_force_centimeters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForceCentimeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_force_centimeters_per_meterResult.Convert(units.TorquePerLengthTonneForceCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForceCentimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForceMeterPerMeter unit
    tonne_force_meters_per_meterJSON := []byte(`{"value": 100, "unit": "TonneForceMeterPerMeter"}`)
    tonne_force_meters_per_meterResult, err := factory.FromDtoJSON(tonne_force_meters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForceMeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_force_meters_per_meterResult.Convert(units.TorquePerLengthTonneForceMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForceMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonMillimeterPerMeter unit
    kilonewton_millimeters_per_meterJSON := []byte(`{"value": 100, "unit": "KilonewtonMillimeterPerMeter"}`)
    kilonewton_millimeters_per_meterResult, err := factory.FromDtoJSON(kilonewton_millimeters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonMillimeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_millimeters_per_meterResult.Convert(units.TorquePerLengthKilonewtonMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMillimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonMillimeterPerMeter unit
    meganewton_millimeters_per_meterJSON := []byte(`{"value": 100, "unit": "MeganewtonMillimeterPerMeter"}`)
    meganewton_millimeters_per_meterResult, err := factory.FromDtoJSON(meganewton_millimeters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonMillimeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_millimeters_per_meterResult.Convert(units.TorquePerLengthMeganewtonMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMillimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonCentimeterPerMeter unit
    kilonewton_centimeters_per_meterJSON := []byte(`{"value": 100, "unit": "KilonewtonCentimeterPerMeter"}`)
    kilonewton_centimeters_per_meterResult, err := factory.FromDtoJSON(kilonewton_centimeters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonCentimeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_centimeters_per_meterResult.Convert(units.TorquePerLengthKilonewtonCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonCentimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonCentimeterPerMeter unit
    meganewton_centimeters_per_meterJSON := []byte(`{"value": 100, "unit": "MeganewtonCentimeterPerMeter"}`)
    meganewton_centimeters_per_meterResult, err := factory.FromDtoJSON(meganewton_centimeters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonCentimeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_centimeters_per_meterResult.Convert(units.TorquePerLengthMeganewtonCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonCentimeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonMeterPerMeter unit
    kilonewton_meters_per_meterJSON := []byte(`{"value": 100, "unit": "KilonewtonMeterPerMeter"}`)
    kilonewton_meters_per_meterResult, err := factory.FromDtoJSON(kilonewton_meters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonMeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_meters_per_meterResult.Convert(units.TorquePerLengthKilonewtonMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonMeterPerMeter unit
    meganewton_meters_per_meterJSON := []byte(`{"value": 100, "unit": "MeganewtonMeterPerMeter"}`)
    meganewton_meters_per_meterResult, err := factory.FromDtoJSON(meganewton_meters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonMeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_meters_per_meterResult.Convert(units.TorquePerLengthMeganewtonMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForceInchPerFoot unit
    kilopound_force_inches_per_footJSON := []byte(`{"value": 100, "unit": "KilopoundForceInchPerFoot"}`)
    kilopound_force_inches_per_footResult, err := factory.FromDtoJSON(kilopound_force_inches_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForceInchPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_force_inches_per_footResult.Convert(units.TorquePerLengthKilopoundForceInchPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForceInchPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with MegapoundForceInchPerFoot unit
    megapound_force_inches_per_footJSON := []byte(`{"value": 100, "unit": "MegapoundForceInchPerFoot"}`)
    megapound_force_inches_per_footResult, err := factory.FromDtoJSON(megapound_force_inches_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapoundForceInchPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapound_force_inches_per_footResult.Convert(units.TorquePerLengthMegapoundForceInchPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundForceInchPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForceFootPerFoot unit
    kilopound_force_feet_per_footJSON := []byte(`{"value": 100, "unit": "KilopoundForceFootPerFoot"}`)
    kilopound_force_feet_per_footResult, err := factory.FromDtoJSON(kilopound_force_feet_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForceFootPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_force_feet_per_footResult.Convert(units.TorquePerLengthKilopoundForceFootPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForceFootPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with MegapoundForceFootPerFoot unit
    megapound_force_feet_per_footJSON := []byte(`{"value": 100, "unit": "MegapoundForceFootPerFoot"}`)
    megapound_force_feet_per_footResult, err := factory.FromDtoJSON(megapound_force_feet_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapoundForceFootPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapound_force_feet_per_footResult.Convert(units.TorquePerLengthMegapoundForceFootPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundForceFootPerFoot = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "NewtonMeterPerMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromNewtonMillimetersPerMeter function
func TestTorquePerLengthFactory_FromNewtonMillimetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonMillimetersPerMeter(100)
    if err != nil {
        t.Errorf("FromNewtonMillimetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthNewtonMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonMillimetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonMillimetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromNewtonMillimetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromNewtonMillimetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonMillimetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromNewtonMillimetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonMillimetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonMillimetersPerMeter(0)
    if err != nil {
        t.Errorf("FromNewtonMillimetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthNewtonMillimeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonMillimetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonCentimetersPerMeter function
func TestTorquePerLengthFactory_FromNewtonCentimetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonCentimetersPerMeter(100)
    if err != nil {
        t.Errorf("FromNewtonCentimetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthNewtonCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonCentimetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonCentimetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromNewtonCentimetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromNewtonCentimetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonCentimetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromNewtonCentimetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonCentimetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonCentimetersPerMeter(0)
    if err != nil {
        t.Errorf("FromNewtonCentimetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthNewtonCentimeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonCentimetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonMetersPerMeter function
func TestTorquePerLengthFactory_FromNewtonMetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonMetersPerMeter(100)
    if err != nil {
        t.Errorf("FromNewtonMetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthNewtonMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonMetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonMetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromNewtonMetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromNewtonMetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonMetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromNewtonMetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonMetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonMetersPerMeter(0)
    if err != nil {
        t.Errorf("FromNewtonMetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthNewtonMeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonMetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundForceInchesPerFoot function
func TestTorquePerLengthFactory_FromPoundForceInchesPerFoot(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundForceInchesPerFoot(100)
    if err != nil {
        t.Errorf("FromPoundForceInchesPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthPoundForceInchPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundForceInchesPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundForceInchesPerFoot(math.NaN())
    if err == nil {
        t.Error("FromPoundForceInchesPerFoot() with NaN value should return error")
    }

    _, err = factory.FromPoundForceInchesPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromPoundForceInchesPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromPoundForceInchesPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundForceInchesPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundForceInchesPerFoot(0)
    if err != nil {
        t.Errorf("FromPoundForceInchesPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthPoundForceInchPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundForceInchesPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundForceFeetPerFoot function
func TestTorquePerLengthFactory_FromPoundForceFeetPerFoot(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundForceFeetPerFoot(100)
    if err != nil {
        t.Errorf("FromPoundForceFeetPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthPoundForceFootPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundForceFeetPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundForceFeetPerFoot(math.NaN())
    if err == nil {
        t.Error("FromPoundForceFeetPerFoot() with NaN value should return error")
    }

    _, err = factory.FromPoundForceFeetPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromPoundForceFeetPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromPoundForceFeetPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundForceFeetPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundForceFeetPerFoot(0)
    if err != nil {
        t.Errorf("FromPoundForceFeetPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthPoundForceFootPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundForceFeetPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramForceMillimetersPerMeter function
func TestTorquePerLengthFactory_FromKilogramForceMillimetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramForceMillimetersPerMeter(100)
    if err != nil {
        t.Errorf("FromKilogramForceMillimetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthKilogramForceMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramForceMillimetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramForceMillimetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramForceMillimetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramForceMillimetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramForceMillimetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramForceMillimetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramForceMillimetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramForceMillimetersPerMeter(0)
    if err != nil {
        t.Errorf("FromKilogramForceMillimetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthKilogramForceMillimeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramForceMillimetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramForceCentimetersPerMeter function
func TestTorquePerLengthFactory_FromKilogramForceCentimetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramForceCentimetersPerMeter(100)
    if err != nil {
        t.Errorf("FromKilogramForceCentimetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthKilogramForceCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramForceCentimetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramForceCentimetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramForceCentimetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramForceCentimetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramForceCentimetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramForceCentimetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramForceCentimetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramForceCentimetersPerMeter(0)
    if err != nil {
        t.Errorf("FromKilogramForceCentimetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthKilogramForceCentimeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramForceCentimetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramForceMetersPerMeter function
func TestTorquePerLengthFactory_FromKilogramForceMetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramForceMetersPerMeter(100)
    if err != nil {
        t.Errorf("FromKilogramForceMetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthKilogramForceMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramForceMetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramForceMetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramForceMetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramForceMetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramForceMetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramForceMetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramForceMetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramForceMetersPerMeter(0)
    if err != nil {
        t.Errorf("FromKilogramForceMetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthKilogramForceMeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramForceMetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonneForceMillimetersPerMeter function
func TestTorquePerLengthFactory_FromTonneForceMillimetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonneForceMillimetersPerMeter(100)
    if err != nil {
        t.Errorf("FromTonneForceMillimetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthTonneForceMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonneForceMillimetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonneForceMillimetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromTonneForceMillimetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromTonneForceMillimetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonneForceMillimetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromTonneForceMillimetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonneForceMillimetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonneForceMillimetersPerMeter(0)
    if err != nil {
        t.Errorf("FromTonneForceMillimetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthTonneForceMillimeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonneForceMillimetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonneForceCentimetersPerMeter function
func TestTorquePerLengthFactory_FromTonneForceCentimetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonneForceCentimetersPerMeter(100)
    if err != nil {
        t.Errorf("FromTonneForceCentimetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthTonneForceCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonneForceCentimetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonneForceCentimetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromTonneForceCentimetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromTonneForceCentimetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonneForceCentimetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromTonneForceCentimetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonneForceCentimetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonneForceCentimetersPerMeter(0)
    if err != nil {
        t.Errorf("FromTonneForceCentimetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthTonneForceCentimeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonneForceCentimetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonneForceMetersPerMeter function
func TestTorquePerLengthFactory_FromTonneForceMetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonneForceMetersPerMeter(100)
    if err != nil {
        t.Errorf("FromTonneForceMetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthTonneForceMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonneForceMetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonneForceMetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromTonneForceMetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromTonneForceMetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonneForceMetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromTonneForceMetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonneForceMetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonneForceMetersPerMeter(0)
    if err != nil {
        t.Errorf("FromTonneForceMetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthTonneForceMeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonneForceMetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonMillimetersPerMeter function
func TestTorquePerLengthFactory_FromKilonewtonMillimetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonMillimetersPerMeter(100)
    if err != nil {
        t.Errorf("FromKilonewtonMillimetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthKilonewtonMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonMillimetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonMillimetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonMillimetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonMillimetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonMillimetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonMillimetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonMillimetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonMillimetersPerMeter(0)
    if err != nil {
        t.Errorf("FromKilonewtonMillimetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthKilonewtonMillimeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonMillimetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonMillimetersPerMeter function
func TestTorquePerLengthFactory_FromMeganewtonMillimetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonMillimetersPerMeter(100)
    if err != nil {
        t.Errorf("FromMeganewtonMillimetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthMeganewtonMillimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonMillimetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonMillimetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonMillimetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonMillimetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonMillimetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonMillimetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonMillimetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonMillimetersPerMeter(0)
    if err != nil {
        t.Errorf("FromMeganewtonMillimetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthMeganewtonMillimeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonMillimetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonCentimetersPerMeter function
func TestTorquePerLengthFactory_FromKilonewtonCentimetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonCentimetersPerMeter(100)
    if err != nil {
        t.Errorf("FromKilonewtonCentimetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthKilonewtonCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonCentimetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonCentimetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonCentimetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonCentimetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonCentimetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonCentimetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonCentimetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonCentimetersPerMeter(0)
    if err != nil {
        t.Errorf("FromKilonewtonCentimetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthKilonewtonCentimeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonCentimetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonCentimetersPerMeter function
func TestTorquePerLengthFactory_FromMeganewtonCentimetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonCentimetersPerMeter(100)
    if err != nil {
        t.Errorf("FromMeganewtonCentimetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthMeganewtonCentimeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonCentimetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonCentimetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonCentimetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonCentimetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonCentimetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonCentimetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonCentimetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonCentimetersPerMeter(0)
    if err != nil {
        t.Errorf("FromMeganewtonCentimetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthMeganewtonCentimeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonCentimetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonMetersPerMeter function
func TestTorquePerLengthFactory_FromKilonewtonMetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonMetersPerMeter(100)
    if err != nil {
        t.Errorf("FromKilonewtonMetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthKilonewtonMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonMetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonMetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonMetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonMetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonMetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonMetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonMetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonMetersPerMeter(0)
    if err != nil {
        t.Errorf("FromKilonewtonMetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthKilonewtonMeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonMetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonMetersPerMeter function
func TestTorquePerLengthFactory_FromMeganewtonMetersPerMeter(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonMetersPerMeter(100)
    if err != nil {
        t.Errorf("FromMeganewtonMetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthMeganewtonMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonMetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonMetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonMetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonMetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonMetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonMetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonMetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonMetersPerMeter(0)
    if err != nil {
        t.Errorf("FromMeganewtonMetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthMeganewtonMeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonMetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundForceInchesPerFoot function
func TestTorquePerLengthFactory_FromKilopoundForceInchesPerFoot(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundForceInchesPerFoot(100)
    if err != nil {
        t.Errorf("FromKilopoundForceInchesPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthKilopoundForceInchPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundForceInchesPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundForceInchesPerFoot(math.NaN())
    if err == nil {
        t.Error("FromKilopoundForceInchesPerFoot() with NaN value should return error")
    }

    _, err = factory.FromKilopoundForceInchesPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundForceInchesPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundForceInchesPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundForceInchesPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundForceInchesPerFoot(0)
    if err != nil {
        t.Errorf("FromKilopoundForceInchesPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthKilopoundForceInchPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundForceInchesPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapoundForceInchesPerFoot function
func TestTorquePerLengthFactory_FromMegapoundForceInchesPerFoot(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapoundForceInchesPerFoot(100)
    if err != nil {
        t.Errorf("FromMegapoundForceInchesPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthMegapoundForceInchPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapoundForceInchesPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapoundForceInchesPerFoot(math.NaN())
    if err == nil {
        t.Error("FromMegapoundForceInchesPerFoot() with NaN value should return error")
    }

    _, err = factory.FromMegapoundForceInchesPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromMegapoundForceInchesPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromMegapoundForceInchesPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapoundForceInchesPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapoundForceInchesPerFoot(0)
    if err != nil {
        t.Errorf("FromMegapoundForceInchesPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthMegapoundForceInchPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapoundForceInchesPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundForceFeetPerFoot function
func TestTorquePerLengthFactory_FromKilopoundForceFeetPerFoot(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundForceFeetPerFoot(100)
    if err != nil {
        t.Errorf("FromKilopoundForceFeetPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthKilopoundForceFootPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundForceFeetPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundForceFeetPerFoot(math.NaN())
    if err == nil {
        t.Error("FromKilopoundForceFeetPerFoot() with NaN value should return error")
    }

    _, err = factory.FromKilopoundForceFeetPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundForceFeetPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundForceFeetPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundForceFeetPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundForceFeetPerFoot(0)
    if err != nil {
        t.Errorf("FromKilopoundForceFeetPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthKilopoundForceFootPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundForceFeetPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapoundForceFeetPerFoot function
func TestTorquePerLengthFactory_FromMegapoundForceFeetPerFoot(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapoundForceFeetPerFoot(100)
    if err != nil {
        t.Errorf("FromMegapoundForceFeetPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePerLengthMegapoundForceFootPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapoundForceFeetPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapoundForceFeetPerFoot(math.NaN())
    if err == nil {
        t.Error("FromMegapoundForceFeetPerFoot() with NaN value should return error")
    }

    _, err = factory.FromMegapoundForceFeetPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromMegapoundForceFeetPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromMegapoundForceFeetPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapoundForceFeetPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapoundForceFeetPerFoot(0)
    if err != nil {
        t.Errorf("FromMegapoundForceFeetPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePerLengthMegapoundForceFootPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapoundForceFeetPerFoot() with zero value = %v, want 0", converted)
    }
}

func TestTorquePerLengthToString(t *testing.T) {
	factory := units.TorquePerLengthFactory{}
	a, err := factory.CreateTorquePerLength(45, units.TorquePerLengthNewtonMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.TorquePerLengthNewtonMeterPerMeter, 2)
	expected := "45.00 " + units.GetTorquePerLengthAbbreviation(units.TorquePerLengthNewtonMeterPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.TorquePerLengthNewtonMeterPerMeter, -1)
	expected = "45 " + units.GetTorquePerLengthAbbreviation(units.TorquePerLengthNewtonMeterPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestTorquePerLength_EqualityAndComparison(t *testing.T) {
	factory := units.TorquePerLengthFactory{}
	a1, _ := factory.CreateTorquePerLength(60, units.TorquePerLengthNewtonMeterPerMeter)
	a2, _ := factory.CreateTorquePerLength(60, units.TorquePerLengthNewtonMeterPerMeter)
	a3, _ := factory.CreateTorquePerLength(120, units.TorquePerLengthNewtonMeterPerMeter)

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

func TestTorquePerLength_Arithmetic(t *testing.T) {
	factory := units.TorquePerLengthFactory{}
	a1, _ := factory.CreateTorquePerLength(30, units.TorquePerLengthNewtonMeterPerMeter)
	a2, _ := factory.CreateTorquePerLength(45, units.TorquePerLengthNewtonMeterPerMeter)

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


func TestGetTorquePerLengthAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.TorquePerLengthUnits
        want string
    }{
        {
            name: "NewtonMillimeterPerMeter abbreviation",
            unit: units.TorquePerLengthNewtonMillimeterPerMeter,
            want: "N·mm/m",
        },
        {
            name: "NewtonCentimeterPerMeter abbreviation",
            unit: units.TorquePerLengthNewtonCentimeterPerMeter,
            want: "N·cm/m",
        },
        {
            name: "NewtonMeterPerMeter abbreviation",
            unit: units.TorquePerLengthNewtonMeterPerMeter,
            want: "N·m/m",
        },
        {
            name: "PoundForceInchPerFoot abbreviation",
            unit: units.TorquePerLengthPoundForceInchPerFoot,
            want: "lbf·in/ft",
        },
        {
            name: "PoundForceFootPerFoot abbreviation",
            unit: units.TorquePerLengthPoundForceFootPerFoot,
            want: "lbf·ft/ft",
        },
        {
            name: "KilogramForceMillimeterPerMeter abbreviation",
            unit: units.TorquePerLengthKilogramForceMillimeterPerMeter,
            want: "kgf·mm/m",
        },
        {
            name: "KilogramForceCentimeterPerMeter abbreviation",
            unit: units.TorquePerLengthKilogramForceCentimeterPerMeter,
            want: "kgf·cm/m",
        },
        {
            name: "KilogramForceMeterPerMeter abbreviation",
            unit: units.TorquePerLengthKilogramForceMeterPerMeter,
            want: "kgf·m/m",
        },
        {
            name: "TonneForceMillimeterPerMeter abbreviation",
            unit: units.TorquePerLengthTonneForceMillimeterPerMeter,
            want: "tf·mm/m",
        },
        {
            name: "TonneForceCentimeterPerMeter abbreviation",
            unit: units.TorquePerLengthTonneForceCentimeterPerMeter,
            want: "tf·cm/m",
        },
        {
            name: "TonneForceMeterPerMeter abbreviation",
            unit: units.TorquePerLengthTonneForceMeterPerMeter,
            want: "tf·m/m",
        },
        {
            name: "KilonewtonMillimeterPerMeter abbreviation",
            unit: units.TorquePerLengthKilonewtonMillimeterPerMeter,
            want: "kN·mm/m",
        },
        {
            name: "MeganewtonMillimeterPerMeter abbreviation",
            unit: units.TorquePerLengthMeganewtonMillimeterPerMeter,
            want: "MN·mm/m",
        },
        {
            name: "KilonewtonCentimeterPerMeter abbreviation",
            unit: units.TorquePerLengthKilonewtonCentimeterPerMeter,
            want: "kN·cm/m",
        },
        {
            name: "MeganewtonCentimeterPerMeter abbreviation",
            unit: units.TorquePerLengthMeganewtonCentimeterPerMeter,
            want: "MN·cm/m",
        },
        {
            name: "KilonewtonMeterPerMeter abbreviation",
            unit: units.TorquePerLengthKilonewtonMeterPerMeter,
            want: "kN·m/m",
        },
        {
            name: "MeganewtonMeterPerMeter abbreviation",
            unit: units.TorquePerLengthMeganewtonMeterPerMeter,
            want: "MN·m/m",
        },
        {
            name: "KilopoundForceInchPerFoot abbreviation",
            unit: units.TorquePerLengthKilopoundForceInchPerFoot,
            want: "klbf·in/ft",
        },
        {
            name: "MegapoundForceInchPerFoot abbreviation",
            unit: units.TorquePerLengthMegapoundForceInchPerFoot,
            want: "Mlbf·in/ft",
        },
        {
            name: "KilopoundForceFootPerFoot abbreviation",
            unit: units.TorquePerLengthKilopoundForceFootPerFoot,
            want: "klbf·ft/ft",
        },
        {
            name: "MegapoundForceFootPerFoot abbreviation",
            unit: units.TorquePerLengthMegapoundForceFootPerFoot,
            want: "Mlbf·ft/ft",
        },
        {
            name: "invalid unit",
            unit: units.TorquePerLengthUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetTorquePerLengthAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetTorquePerLengthAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestTorquePerLength_String(t *testing.T) {
    factory := units.TorquePerLengthFactory{}
    
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
            unit, err := factory.CreateTorquePerLength(tt.value, units.TorquePerLengthNewtonMeterPerMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("TorquePerLength.String() = %v, want %v", got, tt.want)
            }
        })
    }
}