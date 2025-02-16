// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestTorqueDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonMeter"}`
	
	factory := units.TorqueDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.TorqueNewtonMeter {
		t.Errorf("expected unit %v, got %v", units.TorqueNewtonMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestTorqueDto_ToJSON(t *testing.T) {
	dto := units.TorqueDto{
		Value: 45,
		Unit:  units.TorqueNewtonMeter,
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
	if result["unit"].(string) != string(units.TorqueNewtonMeter) {
		t.Errorf("expected unit %s, got %v", units.TorqueNewtonMeter, result["unit"])
	}
}

func TestNewTorque_InvalidValue(t *testing.T) {
	factory := units.TorqueFactory{}
	// NaN value should return an error.
	_, err := factory.CreateTorque(math.NaN(), units.TorqueNewtonMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateTorque(math.Inf(1), units.TorqueNewtonMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestTorqueConversions(t *testing.T) {
	factory := units.TorqueFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateTorque(180, units.TorqueNewtonMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonMillimeters.
		// No expected conversion value provided for NewtonMillimeters, verifying result is not NaN.
		result := a.NewtonMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonMillimeters returned NaN")
		}
	}
	{
		// Test conversion to NewtonCentimeters.
		// No expected conversion value provided for NewtonCentimeters, verifying result is not NaN.
		result := a.NewtonCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonCentimeters returned NaN")
		}
	}
	{
		// Test conversion to NewtonMeters.
		// No expected conversion value provided for NewtonMeters, verifying result is not NaN.
		result := a.NewtonMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonMeters returned NaN")
		}
	}
	{
		// Test conversion to PoundalFeet.
		// No expected conversion value provided for PoundalFeet, verifying result is not NaN.
		result := a.PoundalFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundalFeet returned NaN")
		}
	}
	{
		// Test conversion to PoundForceInches.
		// No expected conversion value provided for PoundForceInches, verifying result is not NaN.
		result := a.PoundForceInches()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundForceInches returned NaN")
		}
	}
	{
		// Test conversion to PoundForceFeet.
		// No expected conversion value provided for PoundForceFeet, verifying result is not NaN.
		result := a.PoundForceFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundForceFeet returned NaN")
		}
	}
	{
		// Test conversion to GramForceMillimeters.
		// No expected conversion value provided for GramForceMillimeters, verifying result is not NaN.
		result := a.GramForceMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramForceMillimeters returned NaN")
		}
	}
	{
		// Test conversion to GramForceCentimeters.
		// No expected conversion value provided for GramForceCentimeters, verifying result is not NaN.
		result := a.GramForceCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramForceCentimeters returned NaN")
		}
	}
	{
		// Test conversion to GramForceMeters.
		// No expected conversion value provided for GramForceMeters, verifying result is not NaN.
		result := a.GramForceMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to GramForceMeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramForceMillimeters.
		// No expected conversion value provided for KilogramForceMillimeters, verifying result is not NaN.
		result := a.KilogramForceMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramForceMillimeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramForceCentimeters.
		// No expected conversion value provided for KilogramForceCentimeters, verifying result is not NaN.
		result := a.KilogramForceCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramForceCentimeters returned NaN")
		}
	}
	{
		// Test conversion to KilogramForceMeters.
		// No expected conversion value provided for KilogramForceMeters, verifying result is not NaN.
		result := a.KilogramForceMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilogramForceMeters returned NaN")
		}
	}
	{
		// Test conversion to TonneForceMillimeters.
		// No expected conversion value provided for TonneForceMillimeters, verifying result is not NaN.
		result := a.TonneForceMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonneForceMillimeters returned NaN")
		}
	}
	{
		// Test conversion to TonneForceCentimeters.
		// No expected conversion value provided for TonneForceCentimeters, verifying result is not NaN.
		result := a.TonneForceCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonneForceCentimeters returned NaN")
		}
	}
	{
		// Test conversion to TonneForceMeters.
		// No expected conversion value provided for TonneForceMeters, verifying result is not NaN.
		result := a.TonneForceMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to TonneForceMeters returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMillimeters.
		// No expected conversion value provided for KilonewtonMillimeters, verifying result is not NaN.
		result := a.KilonewtonMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonMillimeters returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMillimeters.
		// No expected conversion value provided for MeganewtonMillimeters, verifying result is not NaN.
		result := a.MeganewtonMillimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonMillimeters returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonCentimeters.
		// No expected conversion value provided for KilonewtonCentimeters, verifying result is not NaN.
		result := a.KilonewtonCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonCentimeters returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonCentimeters.
		// No expected conversion value provided for MeganewtonCentimeters, verifying result is not NaN.
		result := a.MeganewtonCentimeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonCentimeters returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonMeters.
		// No expected conversion value provided for KilonewtonMeters, verifying result is not NaN.
		result := a.KilonewtonMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilonewtonMeters returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonMeters.
		// No expected conversion value provided for MeganewtonMeters, verifying result is not NaN.
		result := a.MeganewtonMeters()
		if math.IsNaN(result) {
			t.Errorf("conversion to MeganewtonMeters returned NaN")
		}
	}
	{
		// Test conversion to KilopoundForceInches.
		// No expected conversion value provided for KilopoundForceInches, verifying result is not NaN.
		result := a.KilopoundForceInches()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundForceInches returned NaN")
		}
	}
	{
		// Test conversion to MegapoundForceInches.
		// No expected conversion value provided for MegapoundForceInches, verifying result is not NaN.
		result := a.MegapoundForceInches()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapoundForceInches returned NaN")
		}
	}
	{
		// Test conversion to KilopoundForceFeet.
		// No expected conversion value provided for KilopoundForceFeet, verifying result is not NaN.
		result := a.KilopoundForceFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilopoundForceFeet returned NaN")
		}
	}
	{
		// Test conversion to MegapoundForceFeet.
		// No expected conversion value provided for MegapoundForceFeet, verifying result is not NaN.
		result := a.MegapoundForceFeet()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegapoundForceFeet returned NaN")
		}
	}
}

func TestTorque_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.TorqueFactory{}
	a, err := factory.CreateTorque(90, units.TorqueNewtonMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.TorqueNewtonMeter {
		t.Errorf("expected default unit NewtonMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.TorqueNewtonMillimeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.TorqueDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.TorqueNewtonMeter {
		t.Errorf("expected unit NewtonMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestTorqueFactory_FromDto(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueNewtonMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.TorqueDto{
        Value: math.NaN(),
        Unit:  units.TorqueNewtonMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test NewtonMillimeter conversion
    newton_millimetersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueNewtonMillimeter,
    }
    
    var newton_millimetersResult *units.Torque
    newton_millimetersResult, err = factory.FromDto(newton_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_millimetersResult.Convert(units.TorqueNewtonMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMillimeter = %v, want %v", converted, 100)
    }
    // Test NewtonCentimeter conversion
    newton_centimetersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueNewtonCentimeter,
    }
    
    var newton_centimetersResult *units.Torque
    newton_centimetersResult, err = factory.FromDto(newton_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_centimetersResult.Convert(units.TorqueNewtonCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonCentimeter = %v, want %v", converted, 100)
    }
    // Test NewtonMeter conversion
    newton_metersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueNewtonMeter,
    }
    
    var newton_metersResult *units.Torque
    newton_metersResult, err = factory.FromDto(newton_metersDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_metersResult.Convert(units.TorqueNewtonMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMeter = %v, want %v", converted, 100)
    }
    // Test PoundalFoot conversion
    poundal_feetDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorquePoundalFoot,
    }
    
    var poundal_feetResult *units.Torque
    poundal_feetResult, err = factory.FromDto(poundal_feetDto)
    if err != nil {
        t.Errorf("FromDto() with PoundalFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = poundal_feetResult.Convert(units.TorquePoundalFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundalFoot = %v, want %v", converted, 100)
    }
    // Test PoundForceInch conversion
    pound_force_inchesDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorquePoundForceInch,
    }
    
    var pound_force_inchesResult *units.Torque
    pound_force_inchesResult, err = factory.FromDto(pound_force_inchesDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForceInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_inchesResult.Convert(units.TorquePoundForceInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceInch = %v, want %v", converted, 100)
    }
    // Test PoundForceFoot conversion
    pound_force_feetDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorquePoundForceFoot,
    }
    
    var pound_force_feetResult *units.Torque
    pound_force_feetResult, err = factory.FromDto(pound_force_feetDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForceFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_feetResult.Convert(units.TorquePoundForceFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceFoot = %v, want %v", converted, 100)
    }
    // Test GramForceMillimeter conversion
    gram_force_millimetersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueGramForceMillimeter,
    }
    
    var gram_force_millimetersResult *units.Torque
    gram_force_millimetersResult, err = factory.FromDto(gram_force_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with GramForceMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_force_millimetersResult.Convert(units.TorqueGramForceMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramForceMillimeter = %v, want %v", converted, 100)
    }
    // Test GramForceCentimeter conversion
    gram_force_centimetersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueGramForceCentimeter,
    }
    
    var gram_force_centimetersResult *units.Torque
    gram_force_centimetersResult, err = factory.FromDto(gram_force_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with GramForceCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_force_centimetersResult.Convert(units.TorqueGramForceCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramForceCentimeter = %v, want %v", converted, 100)
    }
    // Test GramForceMeter conversion
    gram_force_metersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueGramForceMeter,
    }
    
    var gram_force_metersResult *units.Torque
    gram_force_metersResult, err = factory.FromDto(gram_force_metersDto)
    if err != nil {
        t.Errorf("FromDto() with GramForceMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_force_metersResult.Convert(units.TorqueGramForceMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramForceMeter = %v, want %v", converted, 100)
    }
    // Test KilogramForceMillimeter conversion
    kilogram_force_millimetersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueKilogramForceMillimeter,
    }
    
    var kilogram_force_millimetersResult *units.Torque
    kilogram_force_millimetersResult, err = factory.FromDto(kilogram_force_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForceMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_force_millimetersResult.Convert(units.TorqueKilogramForceMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForceMillimeter = %v, want %v", converted, 100)
    }
    // Test KilogramForceCentimeter conversion
    kilogram_force_centimetersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueKilogramForceCentimeter,
    }
    
    var kilogram_force_centimetersResult *units.Torque
    kilogram_force_centimetersResult, err = factory.FromDto(kilogram_force_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForceCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_force_centimetersResult.Convert(units.TorqueKilogramForceCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForceCentimeter = %v, want %v", converted, 100)
    }
    // Test KilogramForceMeter conversion
    kilogram_force_metersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueKilogramForceMeter,
    }
    
    var kilogram_force_metersResult *units.Torque
    kilogram_force_metersResult, err = factory.FromDto(kilogram_force_metersDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForceMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_force_metersResult.Convert(units.TorqueKilogramForceMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForceMeter = %v, want %v", converted, 100)
    }
    // Test TonneForceMillimeter conversion
    tonne_force_millimetersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueTonneForceMillimeter,
    }
    
    var tonne_force_millimetersResult *units.Torque
    tonne_force_millimetersResult, err = factory.FromDto(tonne_force_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForceMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_force_millimetersResult.Convert(units.TorqueTonneForceMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForceMillimeter = %v, want %v", converted, 100)
    }
    // Test TonneForceCentimeter conversion
    tonne_force_centimetersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueTonneForceCentimeter,
    }
    
    var tonne_force_centimetersResult *units.Torque
    tonne_force_centimetersResult, err = factory.FromDto(tonne_force_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForceCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_force_centimetersResult.Convert(units.TorqueTonneForceCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForceCentimeter = %v, want %v", converted, 100)
    }
    // Test TonneForceMeter conversion
    tonne_force_metersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueTonneForceMeter,
    }
    
    var tonne_force_metersResult *units.Torque
    tonne_force_metersResult, err = factory.FromDto(tonne_force_metersDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForceMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_force_metersResult.Convert(units.TorqueTonneForceMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForceMeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonMillimeter conversion
    kilonewton_millimetersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueKilonewtonMillimeter,
    }
    
    var kilonewton_millimetersResult *units.Torque
    kilonewton_millimetersResult, err = factory.FromDto(kilonewton_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_millimetersResult.Convert(units.TorqueKilonewtonMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMillimeter = %v, want %v", converted, 100)
    }
    // Test MeganewtonMillimeter conversion
    meganewton_millimetersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueMeganewtonMillimeter,
    }
    
    var meganewton_millimetersResult *units.Torque
    meganewton_millimetersResult, err = factory.FromDto(meganewton_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_millimetersResult.Convert(units.TorqueMeganewtonMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMillimeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonCentimeter conversion
    kilonewton_centimetersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueKilonewtonCentimeter,
    }
    
    var kilonewton_centimetersResult *units.Torque
    kilonewton_centimetersResult, err = factory.FromDto(kilonewton_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_centimetersResult.Convert(units.TorqueKilonewtonCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonCentimeter = %v, want %v", converted, 100)
    }
    // Test MeganewtonCentimeter conversion
    meganewton_centimetersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueMeganewtonCentimeter,
    }
    
    var meganewton_centimetersResult *units.Torque
    meganewton_centimetersResult, err = factory.FromDto(meganewton_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_centimetersResult.Convert(units.TorqueMeganewtonCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonCentimeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonMeter conversion
    kilonewton_metersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueKilonewtonMeter,
    }
    
    var kilonewton_metersResult *units.Torque
    kilonewton_metersResult, err = factory.FromDto(kilonewton_metersDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_metersResult.Convert(units.TorqueKilonewtonMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMeter = %v, want %v", converted, 100)
    }
    // Test MeganewtonMeter conversion
    meganewton_metersDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueMeganewtonMeter,
    }
    
    var meganewton_metersResult *units.Torque
    meganewton_metersResult, err = factory.FromDto(meganewton_metersDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_metersResult.Convert(units.TorqueMeganewtonMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMeter = %v, want %v", converted, 100)
    }
    // Test KilopoundForceInch conversion
    kilopound_force_inchesDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueKilopoundForceInch,
    }
    
    var kilopound_force_inchesResult *units.Torque
    kilopound_force_inchesResult, err = factory.FromDto(kilopound_force_inchesDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForceInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_force_inchesResult.Convert(units.TorqueKilopoundForceInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForceInch = %v, want %v", converted, 100)
    }
    // Test MegapoundForceInch conversion
    megapound_force_inchesDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueMegapoundForceInch,
    }
    
    var megapound_force_inchesResult *units.Torque
    megapound_force_inchesResult, err = factory.FromDto(megapound_force_inchesDto)
    if err != nil {
        t.Errorf("FromDto() with MegapoundForceInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapound_force_inchesResult.Convert(units.TorqueMegapoundForceInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundForceInch = %v, want %v", converted, 100)
    }
    // Test KilopoundForceFoot conversion
    kilopound_force_feetDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueKilopoundForceFoot,
    }
    
    var kilopound_force_feetResult *units.Torque
    kilopound_force_feetResult, err = factory.FromDto(kilopound_force_feetDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForceFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_force_feetResult.Convert(units.TorqueKilopoundForceFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForceFoot = %v, want %v", converted, 100)
    }
    // Test MegapoundForceFoot conversion
    megapound_force_feetDto := units.TorqueDto{
        Value: 100,
        Unit:  units.TorqueMegapoundForceFoot,
    }
    
    var megapound_force_feetResult *units.Torque
    megapound_force_feetResult, err = factory.FromDto(megapound_force_feetDto)
    if err != nil {
        t.Errorf("FromDto() with MegapoundForceFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapound_force_feetResult.Convert(units.TorqueMegapoundForceFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundForceFoot = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.TorqueDto{
        Value: 0,
        Unit:  units.TorqueNewtonMeter,
    }
    
    var zeroResult *units.Torque
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestTorqueFactory_FromDtoJSON(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "NewtonMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "NewtonMeter"}`)
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
    nanJSON, _ := json.Marshal(units.TorqueDto{
        Value: nanValue,
        Unit:  units.TorqueNewtonMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with NewtonMillimeter unit
    newton_millimetersJSON := []byte(`{"value": 100, "unit": "NewtonMillimeter"}`)
    newton_millimetersResult, err := factory.FromDtoJSON(newton_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_millimetersResult.Convert(units.TorqueNewtonMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonCentimeter unit
    newton_centimetersJSON := []byte(`{"value": 100, "unit": "NewtonCentimeter"}`)
    newton_centimetersResult, err := factory.FromDtoJSON(newton_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_centimetersResult.Convert(units.TorqueNewtonCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonMeter unit
    newton_metersJSON := []byte(`{"value": 100, "unit": "NewtonMeter"}`)
    newton_metersResult, err := factory.FromDtoJSON(newton_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_metersResult.Convert(units.TorqueNewtonMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundalFoot unit
    poundal_feetJSON := []byte(`{"value": 100, "unit": "PoundalFoot"}`)
    poundal_feetResult, err := factory.FromDtoJSON(poundal_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundalFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = poundal_feetResult.Convert(units.TorquePoundalFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundalFoot = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForceInch unit
    pound_force_inchesJSON := []byte(`{"value": 100, "unit": "PoundForceInch"}`)
    pound_force_inchesResult, err := factory.FromDtoJSON(pound_force_inchesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForceInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_inchesResult.Convert(units.TorquePoundForceInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceInch = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForceFoot unit
    pound_force_feetJSON := []byte(`{"value": 100, "unit": "PoundForceFoot"}`)
    pound_force_feetResult, err := factory.FromDtoJSON(pound_force_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForceFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pound_force_feetResult.Convert(units.TorquePoundForceFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceFoot = %v, want %v", converted, 100)
    }
    // Test JSON with GramForceMillimeter unit
    gram_force_millimetersJSON := []byte(`{"value": 100, "unit": "GramForceMillimeter"}`)
    gram_force_millimetersResult, err := factory.FromDtoJSON(gram_force_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramForceMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_force_millimetersResult.Convert(units.TorqueGramForceMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramForceMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramForceCentimeter unit
    gram_force_centimetersJSON := []byte(`{"value": 100, "unit": "GramForceCentimeter"}`)
    gram_force_centimetersResult, err := factory.FromDtoJSON(gram_force_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramForceCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_force_centimetersResult.Convert(units.TorqueGramForceCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramForceCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with GramForceMeter unit
    gram_force_metersJSON := []byte(`{"value": 100, "unit": "GramForceMeter"}`)
    gram_force_metersResult, err := factory.FromDtoJSON(gram_force_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with GramForceMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gram_force_metersResult.Convert(units.TorqueGramForceMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for GramForceMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForceMillimeter unit
    kilogram_force_millimetersJSON := []byte(`{"value": 100, "unit": "KilogramForceMillimeter"}`)
    kilogram_force_millimetersResult, err := factory.FromDtoJSON(kilogram_force_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForceMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_force_millimetersResult.Convert(units.TorqueKilogramForceMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForceMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForceCentimeter unit
    kilogram_force_centimetersJSON := []byte(`{"value": 100, "unit": "KilogramForceCentimeter"}`)
    kilogram_force_centimetersResult, err := factory.FromDtoJSON(kilogram_force_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForceCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_force_centimetersResult.Convert(units.TorqueKilogramForceCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForceCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForceMeter unit
    kilogram_force_metersJSON := []byte(`{"value": 100, "unit": "KilogramForceMeter"}`)
    kilogram_force_metersResult, err := factory.FromDtoJSON(kilogram_force_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForceMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilogram_force_metersResult.Convert(units.TorqueKilogramForceMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForceMeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForceMillimeter unit
    tonne_force_millimetersJSON := []byte(`{"value": 100, "unit": "TonneForceMillimeter"}`)
    tonne_force_millimetersResult, err := factory.FromDtoJSON(tonne_force_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForceMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_force_millimetersResult.Convert(units.TorqueTonneForceMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForceMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForceCentimeter unit
    tonne_force_centimetersJSON := []byte(`{"value": 100, "unit": "TonneForceCentimeter"}`)
    tonne_force_centimetersResult, err := factory.FromDtoJSON(tonne_force_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForceCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_force_centimetersResult.Convert(units.TorqueTonneForceCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForceCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForceMeter unit
    tonne_force_metersJSON := []byte(`{"value": 100, "unit": "TonneForceMeter"}`)
    tonne_force_metersResult, err := factory.FromDtoJSON(tonne_force_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForceMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonne_force_metersResult.Convert(units.TorqueTonneForceMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForceMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonMillimeter unit
    kilonewton_millimetersJSON := []byte(`{"value": 100, "unit": "KilonewtonMillimeter"}`)
    kilonewton_millimetersResult, err := factory.FromDtoJSON(kilonewton_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_millimetersResult.Convert(units.TorqueKilonewtonMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonMillimeter unit
    meganewton_millimetersJSON := []byte(`{"value": 100, "unit": "MeganewtonMillimeter"}`)
    meganewton_millimetersResult, err := factory.FromDtoJSON(meganewton_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_millimetersResult.Convert(units.TorqueMeganewtonMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonCentimeter unit
    kilonewton_centimetersJSON := []byte(`{"value": 100, "unit": "KilonewtonCentimeter"}`)
    kilonewton_centimetersResult, err := factory.FromDtoJSON(kilonewton_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_centimetersResult.Convert(units.TorqueKilonewtonCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonCentimeter unit
    meganewton_centimetersJSON := []byte(`{"value": 100, "unit": "MeganewtonCentimeter"}`)
    meganewton_centimetersResult, err := factory.FromDtoJSON(meganewton_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_centimetersResult.Convert(units.TorqueMeganewtonCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonMeter unit
    kilonewton_metersJSON := []byte(`{"value": 100, "unit": "KilonewtonMeter"}`)
    kilonewton_metersResult, err := factory.FromDtoJSON(kilonewton_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewton_metersResult.Convert(units.TorqueKilonewtonMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonMeter unit
    meganewton_metersJSON := []byte(`{"value": 100, "unit": "MeganewtonMeter"}`)
    meganewton_metersResult, err := factory.FromDtoJSON(meganewton_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewton_metersResult.Convert(units.TorqueMeganewtonMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForceInch unit
    kilopound_force_inchesJSON := []byte(`{"value": 100, "unit": "KilopoundForceInch"}`)
    kilopound_force_inchesResult, err := factory.FromDtoJSON(kilopound_force_inchesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForceInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_force_inchesResult.Convert(units.TorqueKilopoundForceInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForceInch = %v, want %v", converted, 100)
    }
    // Test JSON with MegapoundForceInch unit
    megapound_force_inchesJSON := []byte(`{"value": 100, "unit": "MegapoundForceInch"}`)
    megapound_force_inchesResult, err := factory.FromDtoJSON(megapound_force_inchesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapoundForceInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapound_force_inchesResult.Convert(units.TorqueMegapoundForceInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundForceInch = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForceFoot unit
    kilopound_force_feetJSON := []byte(`{"value": 100, "unit": "KilopoundForceFoot"}`)
    kilopound_force_feetResult, err := factory.FromDtoJSON(kilopound_force_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForceFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopound_force_feetResult.Convert(units.TorqueKilopoundForceFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForceFoot = %v, want %v", converted, 100)
    }
    // Test JSON with MegapoundForceFoot unit
    megapound_force_feetJSON := []byte(`{"value": 100, "unit": "MegapoundForceFoot"}`)
    megapound_force_feetResult, err := factory.FromDtoJSON(megapound_force_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegapoundForceFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapound_force_feetResult.Convert(units.TorqueMegapoundForceFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegapoundForceFoot = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "NewtonMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromNewtonMillimeters function
func TestTorqueFactory_FromNewtonMillimeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonMillimeters(100)
    if err != nil {
        t.Errorf("FromNewtonMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueNewtonMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonMillimeters(math.NaN())
    if err == nil {
        t.Error("FromNewtonMillimeters() with NaN value should return error")
    }

    _, err = factory.FromNewtonMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromNewtonMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonMillimeters(0)
    if err != nil {
        t.Errorf("FromNewtonMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueNewtonMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonCentimeters function
func TestTorqueFactory_FromNewtonCentimeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonCentimeters(100)
    if err != nil {
        t.Errorf("FromNewtonCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueNewtonCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonCentimeters(math.NaN())
    if err == nil {
        t.Error("FromNewtonCentimeters() with NaN value should return error")
    }

    _, err = factory.FromNewtonCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromNewtonCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonCentimeters(0)
    if err != nil {
        t.Errorf("FromNewtonCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueNewtonCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonMeters function
func TestTorqueFactory_FromNewtonMeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonMeters(100)
    if err != nil {
        t.Errorf("FromNewtonMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueNewtonMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonMeters(math.NaN())
    if err == nil {
        t.Error("FromNewtonMeters() with NaN value should return error")
    }

    _, err = factory.FromNewtonMeters(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonMeters() with +Inf value should return error")
    }

    _, err = factory.FromNewtonMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonMeters(0)
    if err != nil {
        t.Errorf("FromNewtonMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueNewtonMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundalFeet function
func TestTorqueFactory_FromPoundalFeet(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundalFeet(100)
    if err != nil {
        t.Errorf("FromPoundalFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePoundalFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundalFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundalFeet(math.NaN())
    if err == nil {
        t.Error("FromPoundalFeet() with NaN value should return error")
    }

    _, err = factory.FromPoundalFeet(math.Inf(1))
    if err == nil {
        t.Error("FromPoundalFeet() with +Inf value should return error")
    }

    _, err = factory.FromPoundalFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundalFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundalFeet(0)
    if err != nil {
        t.Errorf("FromPoundalFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePoundalFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundalFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundForceInches function
func TestTorqueFactory_FromPoundForceInches(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundForceInches(100)
    if err != nil {
        t.Errorf("FromPoundForceInches() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePoundForceInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundForceInches() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundForceInches(math.NaN())
    if err == nil {
        t.Error("FromPoundForceInches() with NaN value should return error")
    }

    _, err = factory.FromPoundForceInches(math.Inf(1))
    if err == nil {
        t.Error("FromPoundForceInches() with +Inf value should return error")
    }

    _, err = factory.FromPoundForceInches(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundForceInches() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundForceInches(0)
    if err != nil {
        t.Errorf("FromPoundForceInches() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePoundForceInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundForceInches() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundForceFeet function
func TestTorqueFactory_FromPoundForceFeet(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundForceFeet(100)
    if err != nil {
        t.Errorf("FromPoundForceFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorquePoundForceFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundForceFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundForceFeet(math.NaN())
    if err == nil {
        t.Error("FromPoundForceFeet() with NaN value should return error")
    }

    _, err = factory.FromPoundForceFeet(math.Inf(1))
    if err == nil {
        t.Error("FromPoundForceFeet() with +Inf value should return error")
    }

    _, err = factory.FromPoundForceFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundForceFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundForceFeet(0)
    if err != nil {
        t.Errorf("FromPoundForceFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorquePoundForceFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundForceFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromGramForceMillimeters function
func TestTorqueFactory_FromGramForceMillimeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramForceMillimeters(100)
    if err != nil {
        t.Errorf("FromGramForceMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueGramForceMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramForceMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramForceMillimeters(math.NaN())
    if err == nil {
        t.Error("FromGramForceMillimeters() with NaN value should return error")
    }

    _, err = factory.FromGramForceMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromGramForceMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromGramForceMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromGramForceMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramForceMillimeters(0)
    if err != nil {
        t.Errorf("FromGramForceMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueGramForceMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramForceMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromGramForceCentimeters function
func TestTorqueFactory_FromGramForceCentimeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramForceCentimeters(100)
    if err != nil {
        t.Errorf("FromGramForceCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueGramForceCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramForceCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramForceCentimeters(math.NaN())
    if err == nil {
        t.Error("FromGramForceCentimeters() with NaN value should return error")
    }

    _, err = factory.FromGramForceCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromGramForceCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromGramForceCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromGramForceCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramForceCentimeters(0)
    if err != nil {
        t.Errorf("FromGramForceCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueGramForceCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramForceCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromGramForceMeters function
func TestTorqueFactory_FromGramForceMeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGramForceMeters(100)
    if err != nil {
        t.Errorf("FromGramForceMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueGramForceMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGramForceMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGramForceMeters(math.NaN())
    if err == nil {
        t.Error("FromGramForceMeters() with NaN value should return error")
    }

    _, err = factory.FromGramForceMeters(math.Inf(1))
    if err == nil {
        t.Error("FromGramForceMeters() with +Inf value should return error")
    }

    _, err = factory.FromGramForceMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromGramForceMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGramForceMeters(0)
    if err != nil {
        t.Errorf("FromGramForceMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueGramForceMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGramForceMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramForceMillimeters function
func TestTorqueFactory_FromKilogramForceMillimeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramForceMillimeters(100)
    if err != nil {
        t.Errorf("FromKilogramForceMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueKilogramForceMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramForceMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramForceMillimeters(math.NaN())
    if err == nil {
        t.Error("FromKilogramForceMillimeters() with NaN value should return error")
    }

    _, err = factory.FromKilogramForceMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramForceMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromKilogramForceMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramForceMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramForceMillimeters(0)
    if err != nil {
        t.Errorf("FromKilogramForceMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueKilogramForceMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramForceMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramForceCentimeters function
func TestTorqueFactory_FromKilogramForceCentimeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramForceCentimeters(100)
    if err != nil {
        t.Errorf("FromKilogramForceCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueKilogramForceCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramForceCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramForceCentimeters(math.NaN())
    if err == nil {
        t.Error("FromKilogramForceCentimeters() with NaN value should return error")
    }

    _, err = factory.FromKilogramForceCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramForceCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromKilogramForceCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramForceCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramForceCentimeters(0)
    if err != nil {
        t.Errorf("FromKilogramForceCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueKilogramForceCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramForceCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramForceMeters function
func TestTorqueFactory_FromKilogramForceMeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramForceMeters(100)
    if err != nil {
        t.Errorf("FromKilogramForceMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueKilogramForceMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramForceMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramForceMeters(math.NaN())
    if err == nil {
        t.Error("FromKilogramForceMeters() with NaN value should return error")
    }

    _, err = factory.FromKilogramForceMeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramForceMeters() with +Inf value should return error")
    }

    _, err = factory.FromKilogramForceMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramForceMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramForceMeters(0)
    if err != nil {
        t.Errorf("FromKilogramForceMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueKilogramForceMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramForceMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromTonneForceMillimeters function
func TestTorqueFactory_FromTonneForceMillimeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonneForceMillimeters(100)
    if err != nil {
        t.Errorf("FromTonneForceMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueTonneForceMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonneForceMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonneForceMillimeters(math.NaN())
    if err == nil {
        t.Error("FromTonneForceMillimeters() with NaN value should return error")
    }

    _, err = factory.FromTonneForceMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromTonneForceMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromTonneForceMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromTonneForceMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonneForceMillimeters(0)
    if err != nil {
        t.Errorf("FromTonneForceMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueTonneForceMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonneForceMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromTonneForceCentimeters function
func TestTorqueFactory_FromTonneForceCentimeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonneForceCentimeters(100)
    if err != nil {
        t.Errorf("FromTonneForceCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueTonneForceCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonneForceCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonneForceCentimeters(math.NaN())
    if err == nil {
        t.Error("FromTonneForceCentimeters() with NaN value should return error")
    }

    _, err = factory.FromTonneForceCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromTonneForceCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromTonneForceCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromTonneForceCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonneForceCentimeters(0)
    if err != nil {
        t.Errorf("FromTonneForceCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueTonneForceCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonneForceCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromTonneForceMeters function
func TestTorqueFactory_FromTonneForceMeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonneForceMeters(100)
    if err != nil {
        t.Errorf("FromTonneForceMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueTonneForceMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonneForceMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonneForceMeters(math.NaN())
    if err == nil {
        t.Error("FromTonneForceMeters() with NaN value should return error")
    }

    _, err = factory.FromTonneForceMeters(math.Inf(1))
    if err == nil {
        t.Error("FromTonneForceMeters() with +Inf value should return error")
    }

    _, err = factory.FromTonneForceMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromTonneForceMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonneForceMeters(0)
    if err != nil {
        t.Errorf("FromTonneForceMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueTonneForceMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonneForceMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonMillimeters function
func TestTorqueFactory_FromKilonewtonMillimeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonMillimeters(100)
    if err != nil {
        t.Errorf("FromKilonewtonMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueKilonewtonMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonMillimeters(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonMillimeters() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonMillimeters(0)
    if err != nil {
        t.Errorf("FromKilonewtonMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueKilonewtonMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonMillimeters function
func TestTorqueFactory_FromMeganewtonMillimeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonMillimeters(100)
    if err != nil {
        t.Errorf("FromMeganewtonMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueMeganewtonMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonMillimeters(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonMillimeters() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonMillimeters(0)
    if err != nil {
        t.Errorf("FromMeganewtonMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueMeganewtonMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonCentimeters function
func TestTorqueFactory_FromKilonewtonCentimeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonCentimeters(100)
    if err != nil {
        t.Errorf("FromKilonewtonCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueKilonewtonCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonCentimeters(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonCentimeters() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonCentimeters(0)
    if err != nil {
        t.Errorf("FromKilonewtonCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueKilonewtonCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonCentimeters function
func TestTorqueFactory_FromMeganewtonCentimeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonCentimeters(100)
    if err != nil {
        t.Errorf("FromMeganewtonCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueMeganewtonCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonCentimeters(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonCentimeters() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonCentimeters(0)
    if err != nil {
        t.Errorf("FromMeganewtonCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueMeganewtonCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonMeters function
func TestTorqueFactory_FromKilonewtonMeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonMeters(100)
    if err != nil {
        t.Errorf("FromKilonewtonMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueKilonewtonMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonMeters(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonMeters() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonMeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonMeters() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonMeters(0)
    if err != nil {
        t.Errorf("FromKilonewtonMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueKilonewtonMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonMeters function
func TestTorqueFactory_FromMeganewtonMeters(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonMeters(100)
    if err != nil {
        t.Errorf("FromMeganewtonMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueMeganewtonMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonMeters(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonMeters() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonMeters(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonMeters() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonMeters(0)
    if err != nil {
        t.Errorf("FromMeganewtonMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueMeganewtonMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundForceInches function
func TestTorqueFactory_FromKilopoundForceInches(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundForceInches(100)
    if err != nil {
        t.Errorf("FromKilopoundForceInches() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueKilopoundForceInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundForceInches() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundForceInches(math.NaN())
    if err == nil {
        t.Error("FromKilopoundForceInches() with NaN value should return error")
    }

    _, err = factory.FromKilopoundForceInches(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundForceInches() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundForceInches(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundForceInches() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundForceInches(0)
    if err != nil {
        t.Errorf("FromKilopoundForceInches() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueKilopoundForceInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundForceInches() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapoundForceInches function
func TestTorqueFactory_FromMegapoundForceInches(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapoundForceInches(100)
    if err != nil {
        t.Errorf("FromMegapoundForceInches() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueMegapoundForceInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapoundForceInches() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapoundForceInches(math.NaN())
    if err == nil {
        t.Error("FromMegapoundForceInches() with NaN value should return error")
    }

    _, err = factory.FromMegapoundForceInches(math.Inf(1))
    if err == nil {
        t.Error("FromMegapoundForceInches() with +Inf value should return error")
    }

    _, err = factory.FromMegapoundForceInches(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapoundForceInches() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapoundForceInches(0)
    if err != nil {
        t.Errorf("FromMegapoundForceInches() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueMegapoundForceInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapoundForceInches() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundForceFeet function
func TestTorqueFactory_FromKilopoundForceFeet(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundForceFeet(100)
    if err != nil {
        t.Errorf("FromKilopoundForceFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueKilopoundForceFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundForceFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundForceFeet(math.NaN())
    if err == nil {
        t.Error("FromKilopoundForceFeet() with NaN value should return error")
    }

    _, err = factory.FromKilopoundForceFeet(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundForceFeet() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundForceFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundForceFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundForceFeet(0)
    if err != nil {
        t.Errorf("FromKilopoundForceFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueKilopoundForceFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundForceFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapoundForceFeet function
func TestTorqueFactory_FromMegapoundForceFeet(t *testing.T) {
    factory := units.TorqueFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapoundForceFeet(100)
    if err != nil {
        t.Errorf("FromMegapoundForceFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.TorqueMegapoundForceFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapoundForceFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapoundForceFeet(math.NaN())
    if err == nil {
        t.Error("FromMegapoundForceFeet() with NaN value should return error")
    }

    _, err = factory.FromMegapoundForceFeet(math.Inf(1))
    if err == nil {
        t.Error("FromMegapoundForceFeet() with +Inf value should return error")
    }

    _, err = factory.FromMegapoundForceFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapoundForceFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapoundForceFeet(0)
    if err != nil {
        t.Errorf("FromMegapoundForceFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.TorqueMegapoundForceFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapoundForceFeet() with zero value = %v, want 0", converted)
    }
}

func TestTorqueToString(t *testing.T) {
	factory := units.TorqueFactory{}
	a, err := factory.CreateTorque(45, units.TorqueNewtonMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.TorqueNewtonMeter, 2)
	expected := "45.00 " + units.GetTorqueAbbreviation(units.TorqueNewtonMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.TorqueNewtonMeter, -1)
	expected = "45 " + units.GetTorqueAbbreviation(units.TorqueNewtonMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestTorque_EqualityAndComparison(t *testing.T) {
	factory := units.TorqueFactory{}
	a1, _ := factory.CreateTorque(60, units.TorqueNewtonMeter)
	a2, _ := factory.CreateTorque(60, units.TorqueNewtonMeter)
	a3, _ := factory.CreateTorque(120, units.TorqueNewtonMeter)

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

func TestTorque_Arithmetic(t *testing.T) {
	factory := units.TorqueFactory{}
	a1, _ := factory.CreateTorque(30, units.TorqueNewtonMeter)
	a2, _ := factory.CreateTorque(45, units.TorqueNewtonMeter)

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