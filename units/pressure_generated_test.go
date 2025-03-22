// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestPressureDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Pascal"}`
	
	factory := units.PressureDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.PressurePascal {
		t.Errorf("expected unit %v, got %v", units.PressurePascal, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Pascal"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestPressureDto_ToJSON(t *testing.T) {
	dto := units.PressureDto{
		Value: 45,
		Unit:  units.PressurePascal,
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
	if result["unit"].(string) != string(units.PressurePascal) {
		t.Errorf("expected unit %s, got %v", units.PressurePascal, result["unit"])
	}
}

func TestNewPressure_InvalidValue(t *testing.T) {
	factory := units.PressureFactory{}
	// NaN value should return an error.
	_, err := factory.CreatePressure(math.NaN(), units.PressurePascal)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreatePressure(math.Inf(1), units.PressurePascal)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestPressureConversions(t *testing.T) {
	factory := units.PressureFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreatePressure(180, units.PressurePascal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Pascals.
		// No expected conversion value provided for Pascals, verifying result is not NaN.
		result := a.Pascals()
		cacheResult := a.Pascals()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Pascals returned NaN")
		}
	}
	{
		// Test conversion to Atmospheres.
		// No expected conversion value provided for Atmospheres, verifying result is not NaN.
		result := a.Atmospheres()
		cacheResult := a.Atmospheres()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Atmospheres returned NaN")
		}
	}
	{
		// Test conversion to Bars.
		// No expected conversion value provided for Bars, verifying result is not NaN.
		result := a.Bars()
		cacheResult := a.Bars()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Bars returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerSquareMeter.
		// No expected conversion value provided for KilogramsForcePerSquareMeter, verifying result is not NaN.
		result := a.KilogramsForcePerSquareMeter()
		cacheResult := a.KilogramsForcePerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsForcePerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerSquareCentimeter.
		// No expected conversion value provided for KilogramsForcePerSquareCentimeter, verifying result is not NaN.
		result := a.KilogramsForcePerSquareCentimeter()
		cacheResult := a.KilogramsForcePerSquareCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsForcePerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilogramsForcePerSquareMillimeter.
		// No expected conversion value provided for KilogramsForcePerSquareMillimeter, verifying result is not NaN.
		result := a.KilogramsForcePerSquareMillimeter()
		cacheResult := a.KilogramsForcePerSquareMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilogramsForcePerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerSquareMeter.
		// No expected conversion value provided for NewtonsPerSquareMeter, verifying result is not NaN.
		result := a.NewtonsPerSquareMeter()
		cacheResult := a.NewtonsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerSquareCentimeter.
		// No expected conversion value provided for NewtonsPerSquareCentimeter, verifying result is not NaN.
		result := a.NewtonsPerSquareCentimeter()
		cacheResult := a.NewtonsPerSquareCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to NewtonsPerSquareMillimeter.
		// No expected conversion value provided for NewtonsPerSquareMillimeter, verifying result is not NaN.
		result := a.NewtonsPerSquareMillimeter()
		cacheResult := a.NewtonsPerSquareMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NewtonsPerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to TechnicalAtmospheres.
		// No expected conversion value provided for TechnicalAtmospheres, verifying result is not NaN.
		result := a.TechnicalAtmospheres()
		cacheResult := a.TechnicalAtmospheres()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TechnicalAtmospheres returned NaN")
		}
	}
	{
		// Test conversion to Torrs.
		// No expected conversion value provided for Torrs, verifying result is not NaN.
		result := a.Torrs()
		cacheResult := a.Torrs()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Torrs returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerSquareInch.
		// No expected conversion value provided for PoundsForcePerSquareInch, verifying result is not NaN.
		result := a.PoundsForcePerSquareInch()
		cacheResult := a.PoundsForcePerSquareInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsForcePerSquareInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerSquareMil.
		// No expected conversion value provided for PoundsForcePerSquareMil, verifying result is not NaN.
		result := a.PoundsForcePerSquareMil()
		cacheResult := a.PoundsForcePerSquareMil()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsForcePerSquareMil returned NaN")
		}
	}
	{
		// Test conversion to PoundsForcePerSquareFoot.
		// No expected conversion value provided for PoundsForcePerSquareFoot, verifying result is not NaN.
		result := a.PoundsForcePerSquareFoot()
		cacheResult := a.PoundsForcePerSquareFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsForcePerSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerSquareMillimeter.
		// No expected conversion value provided for TonnesForcePerSquareMillimeter, verifying result is not NaN.
		result := a.TonnesForcePerSquareMillimeter()
		cacheResult := a.TonnesForcePerSquareMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonnesForcePerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerSquareMeter.
		// No expected conversion value provided for TonnesForcePerSquareMeter, verifying result is not NaN.
		result := a.TonnesForcePerSquareMeter()
		cacheResult := a.TonnesForcePerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonnesForcePerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MetersOfHead.
		// No expected conversion value provided for MetersOfHead, verifying result is not NaN.
		result := a.MetersOfHead()
		cacheResult := a.MetersOfHead()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MetersOfHead returned NaN")
		}
	}
	{
		// Test conversion to TonnesForcePerSquareCentimeter.
		// No expected conversion value provided for TonnesForcePerSquareCentimeter, verifying result is not NaN.
		result := a.TonnesForcePerSquareCentimeter()
		cacheResult := a.TonnesForcePerSquareCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to TonnesForcePerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to FeetOfHead.
		// No expected conversion value provided for FeetOfHead, verifying result is not NaN.
		result := a.FeetOfHead()
		cacheResult := a.FeetOfHead()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to FeetOfHead returned NaN")
		}
	}
	{
		// Test conversion to MillimetersOfMercury.
		// No expected conversion value provided for MillimetersOfMercury, verifying result is not NaN.
		result := a.MillimetersOfMercury()
		cacheResult := a.MillimetersOfMercury()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimetersOfMercury returned NaN")
		}
	}
	{
		// Test conversion to InchesOfMercury.
		// No expected conversion value provided for InchesOfMercury, verifying result is not NaN.
		result := a.InchesOfMercury()
		cacheResult := a.InchesOfMercury()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InchesOfMercury returned NaN")
		}
	}
	{
		// Test conversion to DynesPerSquareCentimeter.
		// No expected conversion value provided for DynesPerSquareCentimeter, verifying result is not NaN.
		result := a.DynesPerSquareCentimeter()
		cacheResult := a.DynesPerSquareCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DynesPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerInchSecondSquared.
		// No expected conversion value provided for PoundsPerInchSecondSquared, verifying result is not NaN.
		result := a.PoundsPerInchSecondSquared()
		cacheResult := a.PoundsPerInchSecondSquared()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PoundsPerInchSecondSquared returned NaN")
		}
	}
	{
		// Test conversion to MetersOfWaterColumn.
		// No expected conversion value provided for MetersOfWaterColumn, verifying result is not NaN.
		result := a.MetersOfWaterColumn()
		cacheResult := a.MetersOfWaterColumn()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MetersOfWaterColumn returned NaN")
		}
	}
	{
		// Test conversion to InchesOfWaterColumn.
		// No expected conversion value provided for InchesOfWaterColumn, verifying result is not NaN.
		result := a.InchesOfWaterColumn()
		cacheResult := a.InchesOfWaterColumn()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InchesOfWaterColumn returned NaN")
		}
	}
	{
		// Test conversion to MetersOfElevation.
		// No expected conversion value provided for MetersOfElevation, verifying result is not NaN.
		result := a.MetersOfElevation()
		cacheResult := a.MetersOfElevation()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MetersOfElevation returned NaN")
		}
	}
	{
		// Test conversion to FeetOfElevation.
		// No expected conversion value provided for FeetOfElevation, verifying result is not NaN.
		result := a.FeetOfElevation()
		cacheResult := a.FeetOfElevation()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to FeetOfElevation returned NaN")
		}
	}
	{
		// Test conversion to Micropascals.
		// No expected conversion value provided for Micropascals, verifying result is not NaN.
		result := a.Micropascals()
		cacheResult := a.Micropascals()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Micropascals returned NaN")
		}
	}
	{
		// Test conversion to Millipascals.
		// No expected conversion value provided for Millipascals, verifying result is not NaN.
		result := a.Millipascals()
		cacheResult := a.Millipascals()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millipascals returned NaN")
		}
	}
	{
		// Test conversion to Decapascals.
		// No expected conversion value provided for Decapascals, verifying result is not NaN.
		result := a.Decapascals()
		cacheResult := a.Decapascals()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Decapascals returned NaN")
		}
	}
	{
		// Test conversion to Hectopascals.
		// No expected conversion value provided for Hectopascals, verifying result is not NaN.
		result := a.Hectopascals()
		cacheResult := a.Hectopascals()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Hectopascals returned NaN")
		}
	}
	{
		// Test conversion to Kilopascals.
		// No expected conversion value provided for Kilopascals, verifying result is not NaN.
		result := a.Kilopascals()
		cacheResult := a.Kilopascals()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilopascals returned NaN")
		}
	}
	{
		// Test conversion to Megapascals.
		// No expected conversion value provided for Megapascals, verifying result is not NaN.
		result := a.Megapascals()
		cacheResult := a.Megapascals()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megapascals returned NaN")
		}
	}
	{
		// Test conversion to Gigapascals.
		// No expected conversion value provided for Gigapascals, verifying result is not NaN.
		result := a.Gigapascals()
		cacheResult := a.Gigapascals()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gigapascals returned NaN")
		}
	}
	{
		// Test conversion to Microbars.
		// No expected conversion value provided for Microbars, verifying result is not NaN.
		result := a.Microbars()
		cacheResult := a.Microbars()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microbars returned NaN")
		}
	}
	{
		// Test conversion to Millibars.
		// No expected conversion value provided for Millibars, verifying result is not NaN.
		result := a.Millibars()
		cacheResult := a.Millibars()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millibars returned NaN")
		}
	}
	{
		// Test conversion to Centibars.
		// No expected conversion value provided for Centibars, verifying result is not NaN.
		result := a.Centibars()
		cacheResult := a.Centibars()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Centibars returned NaN")
		}
	}
	{
		// Test conversion to Decibars.
		// No expected conversion value provided for Decibars, verifying result is not NaN.
		result := a.Decibars()
		cacheResult := a.Decibars()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Decibars returned NaN")
		}
	}
	{
		// Test conversion to Kilobars.
		// No expected conversion value provided for Kilobars, verifying result is not NaN.
		result := a.Kilobars()
		cacheResult := a.Kilobars()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilobars returned NaN")
		}
	}
	{
		// Test conversion to Megabars.
		// No expected conversion value provided for Megabars, verifying result is not NaN.
		result := a.Megabars()
		cacheResult := a.Megabars()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megabars returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerSquareMeter.
		// No expected conversion value provided for KilonewtonsPerSquareMeter, verifying result is not NaN.
		result := a.KilonewtonsPerSquareMeter()
		cacheResult := a.KilonewtonsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MeganewtonsPerSquareMeter.
		// No expected conversion value provided for MeganewtonsPerSquareMeter, verifying result is not NaN.
		result := a.MeganewtonsPerSquareMeter()
		cacheResult := a.MeganewtonsPerSquareMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MeganewtonsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerSquareCentimeter.
		// No expected conversion value provided for KilonewtonsPerSquareCentimeter, verifying result is not NaN.
		result := a.KilonewtonsPerSquareCentimeter()
		cacheResult := a.KilonewtonsPerSquareCentimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilonewtonsPerSquareMillimeter.
		// No expected conversion value provided for KilonewtonsPerSquareMillimeter, verifying result is not NaN.
		result := a.KilonewtonsPerSquareMillimeter()
		cacheResult := a.KilonewtonsPerSquareMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilonewtonsPerSquareMillimeter returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerSquareInch.
		// No expected conversion value provided for KilopoundsForcePerSquareInch, verifying result is not NaN.
		result := a.KilopoundsForcePerSquareInch()
		cacheResult := a.KilopoundsForcePerSquareInch()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsForcePerSquareInch returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerSquareMil.
		// No expected conversion value provided for KilopoundsForcePerSquareMil, verifying result is not NaN.
		result := a.KilopoundsForcePerSquareMil()
		cacheResult := a.KilopoundsForcePerSquareMil()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsForcePerSquareMil returned NaN")
		}
	}
	{
		// Test conversion to KilopoundsForcePerSquareFoot.
		// No expected conversion value provided for KilopoundsForcePerSquareFoot, verifying result is not NaN.
		result := a.KilopoundsForcePerSquareFoot()
		cacheResult := a.KilopoundsForcePerSquareFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilopoundsForcePerSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to MillimetersOfWaterColumn.
		// No expected conversion value provided for MillimetersOfWaterColumn, verifying result is not NaN.
		result := a.MillimetersOfWaterColumn()
		cacheResult := a.MillimetersOfWaterColumn()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimetersOfWaterColumn returned NaN")
		}
	}
	{
		// Test conversion to CentimetersOfWaterColumn.
		// No expected conversion value provided for CentimetersOfWaterColumn, verifying result is not NaN.
		result := a.CentimetersOfWaterColumn()
		cacheResult := a.CentimetersOfWaterColumn()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentimetersOfWaterColumn returned NaN")
		}
	}
}

func TestPressure_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.PressureFactory{}
	a, err := factory.CreatePressure(90, units.PressurePascal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.PressurePascal {
		t.Errorf("expected default unit Pascal, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.PressurePascal
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.PressureDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.PressurePascal {
		t.Errorf("expected unit Pascal, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestPressureFactory_FromDto(t *testing.T) {
    factory := units.PressureFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressurePascal,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.PressureDto{
        Value: math.NaN(),
        Unit:  units.PressurePascal,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Pascal conversion
    pascalsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressurePascal,
    }
    
    var pascalsResult *units.Pressure
    pascalsResult, err = factory.FromDto(pascalsDto)
    if err != nil {
        t.Errorf("FromDto() with Pascal returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascalsResult.Convert(units.PressurePascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Pascal = %v, want %v", converted, 100)
    }
    // Test Atmosphere conversion
    atmospheresDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureAtmosphere,
    }
    
    var atmospheresResult *units.Pressure
    atmospheresResult, err = factory.FromDto(atmospheresDto)
    if err != nil {
        t.Errorf("FromDto() with Atmosphere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = atmospheresResult.Convert(units.PressureAtmosphere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Atmosphere = %v, want %v", converted, 100)
    }
    // Test Bar conversion
    barsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureBar,
    }
    
    var barsResult *units.Pressure
    barsResult, err = factory.FromDto(barsDto)
    if err != nil {
        t.Errorf("FromDto() with Bar returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = barsResult.Convert(units.PressureBar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Bar = %v, want %v", converted, 100)
    }
    // Test KilogramForcePerSquareMeter conversion
    kilograms_force_per_square_meterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureKilogramForcePerSquareMeter,
    }
    
    var kilograms_force_per_square_meterResult *units.Pressure
    kilograms_force_per_square_meterResult, err = factory.FromDto(kilograms_force_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForcePerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_square_meterResult.Convert(units.PressureKilogramForcePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerSquareMeter = %v, want %v", converted, 100)
    }
    // Test KilogramForcePerSquareCentimeter conversion
    kilograms_force_per_square_centimeterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureKilogramForcePerSquareCentimeter,
    }
    
    var kilograms_force_per_square_centimeterResult *units.Pressure
    kilograms_force_per_square_centimeterResult, err = factory.FromDto(kilograms_force_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForcePerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_square_centimeterResult.Convert(units.PressureKilogramForcePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test KilogramForcePerSquareMillimeter conversion
    kilograms_force_per_square_millimeterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureKilogramForcePerSquareMillimeter,
    }
    
    var kilograms_force_per_square_millimeterResult *units.Pressure
    kilograms_force_per_square_millimeterResult, err = factory.FromDto(kilograms_force_per_square_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilogramForcePerSquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_square_millimeterResult.Convert(units.PressureKilogramForcePerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test NewtonPerSquareMeter conversion
    newtons_per_square_meterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureNewtonPerSquareMeter,
    }
    
    var newtons_per_square_meterResult *units.Pressure
    newtons_per_square_meterResult, err = factory.FromDto(newtons_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_square_meterResult.Convert(units.PressureNewtonPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test NewtonPerSquareCentimeter conversion
    newtons_per_square_centimeterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureNewtonPerSquareCentimeter,
    }
    
    var newtons_per_square_centimeterResult *units.Pressure
    newtons_per_square_centimeterResult, err = factory.FromDto(newtons_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_square_centimeterResult.Convert(units.PressureNewtonPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test NewtonPerSquareMillimeter conversion
    newtons_per_square_millimeterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureNewtonPerSquareMillimeter,
    }
    
    var newtons_per_square_millimeterResult *units.Pressure
    newtons_per_square_millimeterResult, err = factory.FromDto(newtons_per_square_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonPerSquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_square_millimeterResult.Convert(units.PressureNewtonPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test TechnicalAtmosphere conversion
    technical_atmospheresDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureTechnicalAtmosphere,
    }
    
    var technical_atmospheresResult *units.Pressure
    technical_atmospheresResult, err = factory.FromDto(technical_atmospheresDto)
    if err != nil {
        t.Errorf("FromDto() with TechnicalAtmosphere returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = technical_atmospheresResult.Convert(units.PressureTechnicalAtmosphere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TechnicalAtmosphere = %v, want %v", converted, 100)
    }
    // Test Torr conversion
    torrsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureTorr,
    }
    
    var torrsResult *units.Pressure
    torrsResult, err = factory.FromDto(torrsDto)
    if err != nil {
        t.Errorf("FromDto() with Torr returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = torrsResult.Convert(units.PressureTorr)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Torr = %v, want %v", converted, 100)
    }
    // Test PoundForcePerSquareInch conversion
    pounds_force_per_square_inchDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressurePoundForcePerSquareInch,
    }
    
    var pounds_force_per_square_inchResult *units.Pressure
    pounds_force_per_square_inchResult, err = factory.FromDto(pounds_force_per_square_inchDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForcePerSquareInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_square_inchResult.Convert(units.PressurePoundForcePerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerSquareInch = %v, want %v", converted, 100)
    }
    // Test PoundForcePerSquareMil conversion
    pounds_force_per_square_milDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressurePoundForcePerSquareMil,
    }
    
    var pounds_force_per_square_milResult *units.Pressure
    pounds_force_per_square_milResult, err = factory.FromDto(pounds_force_per_square_milDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForcePerSquareMil returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_square_milResult.Convert(units.PressurePoundForcePerSquareMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerSquareMil = %v, want %v", converted, 100)
    }
    // Test PoundForcePerSquareFoot conversion
    pounds_force_per_square_footDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressurePoundForcePerSquareFoot,
    }
    
    var pounds_force_per_square_footResult *units.Pressure
    pounds_force_per_square_footResult, err = factory.FromDto(pounds_force_per_square_footDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForcePerSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_square_footResult.Convert(units.PressurePoundForcePerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerSquareFoot = %v, want %v", converted, 100)
    }
    // Test TonneForcePerSquareMillimeter conversion
    tonnes_force_per_square_millimeterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureTonneForcePerSquareMillimeter,
    }
    
    var tonnes_force_per_square_millimeterResult *units.Pressure
    tonnes_force_per_square_millimeterResult, err = factory.FromDto(tonnes_force_per_square_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForcePerSquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_square_millimeterResult.Convert(units.PressureTonneForcePerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test TonneForcePerSquareMeter conversion
    tonnes_force_per_square_meterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureTonneForcePerSquareMeter,
    }
    
    var tonnes_force_per_square_meterResult *units.Pressure
    tonnes_force_per_square_meterResult, err = factory.FromDto(tonnes_force_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForcePerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_square_meterResult.Convert(units.PressureTonneForcePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerSquareMeter = %v, want %v", converted, 100)
    }
    // Test MeterOfHead conversion
    meters_of_headDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureMeterOfHead,
    }
    
    var meters_of_headResult *units.Pressure
    meters_of_headResult, err = factory.FromDto(meters_of_headDto)
    if err != nil {
        t.Errorf("FromDto() with MeterOfHead returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_of_headResult.Convert(units.PressureMeterOfHead)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterOfHead = %v, want %v", converted, 100)
    }
    // Test TonneForcePerSquareCentimeter conversion
    tonnes_force_per_square_centimeterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureTonneForcePerSquareCentimeter,
    }
    
    var tonnes_force_per_square_centimeterResult *units.Pressure
    tonnes_force_per_square_centimeterResult, err = factory.FromDto(tonnes_force_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with TonneForcePerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_square_centimeterResult.Convert(units.PressureTonneForcePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test FootOfHead conversion
    feet_of_headDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureFootOfHead,
    }
    
    var feet_of_headResult *units.Pressure
    feet_of_headResult, err = factory.FromDto(feet_of_headDto)
    if err != nil {
        t.Errorf("FromDto() with FootOfHead returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_of_headResult.Convert(units.PressureFootOfHead)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootOfHead = %v, want %v", converted, 100)
    }
    // Test MillimeterOfMercury conversion
    millimeters_of_mercuryDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureMillimeterOfMercury,
    }
    
    var millimeters_of_mercuryResult *units.Pressure
    millimeters_of_mercuryResult, err = factory.FromDto(millimeters_of_mercuryDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterOfMercury returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_of_mercuryResult.Convert(units.PressureMillimeterOfMercury)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterOfMercury = %v, want %v", converted, 100)
    }
    // Test InchOfMercury conversion
    inches_of_mercuryDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureInchOfMercury,
    }
    
    var inches_of_mercuryResult *units.Pressure
    inches_of_mercuryResult, err = factory.FromDto(inches_of_mercuryDto)
    if err != nil {
        t.Errorf("FromDto() with InchOfMercury returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_of_mercuryResult.Convert(units.PressureInchOfMercury)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchOfMercury = %v, want %v", converted, 100)
    }
    // Test DynePerSquareCentimeter conversion
    dynes_per_square_centimeterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureDynePerSquareCentimeter,
    }
    
    var dynes_per_square_centimeterResult *units.Pressure
    dynes_per_square_centimeterResult, err = factory.FromDto(dynes_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with DynePerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = dynes_per_square_centimeterResult.Convert(units.PressureDynePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DynePerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test PoundPerInchSecondSquared conversion
    pounds_per_inch_second_squaredDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressurePoundPerInchSecondSquared,
    }
    
    var pounds_per_inch_second_squaredResult *units.Pressure
    pounds_per_inch_second_squaredResult, err = factory.FromDto(pounds_per_inch_second_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerInchSecondSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_inch_second_squaredResult.Convert(units.PressurePoundPerInchSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerInchSecondSquared = %v, want %v", converted, 100)
    }
    // Test MeterOfWaterColumn conversion
    meters_of_water_columnDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureMeterOfWaterColumn,
    }
    
    var meters_of_water_columnResult *units.Pressure
    meters_of_water_columnResult, err = factory.FromDto(meters_of_water_columnDto)
    if err != nil {
        t.Errorf("FromDto() with MeterOfWaterColumn returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_of_water_columnResult.Convert(units.PressureMeterOfWaterColumn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterOfWaterColumn = %v, want %v", converted, 100)
    }
    // Test InchOfWaterColumn conversion
    inches_of_water_columnDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureInchOfWaterColumn,
    }
    
    var inches_of_water_columnResult *units.Pressure
    inches_of_water_columnResult, err = factory.FromDto(inches_of_water_columnDto)
    if err != nil {
        t.Errorf("FromDto() with InchOfWaterColumn returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_of_water_columnResult.Convert(units.PressureInchOfWaterColumn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchOfWaterColumn = %v, want %v", converted, 100)
    }
    // Test MeterOfElevation conversion
    meters_of_elevationDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureMeterOfElevation,
    }
    
    var meters_of_elevationResult *units.Pressure
    meters_of_elevationResult, err = factory.FromDto(meters_of_elevationDto)
    if err != nil {
        t.Errorf("FromDto() with MeterOfElevation returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_of_elevationResult.Convert(units.PressureMeterOfElevation)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterOfElevation = %v, want %v", converted, 100)
    }
    // Test FootOfElevation conversion
    feet_of_elevationDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureFootOfElevation,
    }
    
    var feet_of_elevationResult *units.Pressure
    feet_of_elevationResult, err = factory.FromDto(feet_of_elevationDto)
    if err != nil {
        t.Errorf("FromDto() with FootOfElevation returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_of_elevationResult.Convert(units.PressureFootOfElevation)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootOfElevation = %v, want %v", converted, 100)
    }
    // Test Micropascal conversion
    micropascalsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureMicropascal,
    }
    
    var micropascalsResult *units.Pressure
    micropascalsResult, err = factory.FromDto(micropascalsDto)
    if err != nil {
        t.Errorf("FromDto() with Micropascal returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micropascalsResult.Convert(units.PressureMicropascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Micropascal = %v, want %v", converted, 100)
    }
    // Test Millipascal conversion
    millipascalsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureMillipascal,
    }
    
    var millipascalsResult *units.Pressure
    millipascalsResult, err = factory.FromDto(millipascalsDto)
    if err != nil {
        t.Errorf("FromDto() with Millipascal returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millipascalsResult.Convert(units.PressureMillipascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millipascal = %v, want %v", converted, 100)
    }
    // Test Decapascal conversion
    decapascalsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureDecapascal,
    }
    
    var decapascalsResult *units.Pressure
    decapascalsResult, err = factory.FromDto(decapascalsDto)
    if err != nil {
        t.Errorf("FromDto() with Decapascal returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decapascalsResult.Convert(units.PressureDecapascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decapascal = %v, want %v", converted, 100)
    }
    // Test Hectopascal conversion
    hectopascalsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureHectopascal,
    }
    
    var hectopascalsResult *units.Pressure
    hectopascalsResult, err = factory.FromDto(hectopascalsDto)
    if err != nil {
        t.Errorf("FromDto() with Hectopascal returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectopascalsResult.Convert(units.PressureHectopascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hectopascal = %v, want %v", converted, 100)
    }
    // Test Kilopascal conversion
    kilopascalsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureKilopascal,
    }
    
    var kilopascalsResult *units.Pressure
    kilopascalsResult, err = factory.FromDto(kilopascalsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilopascal returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopascalsResult.Convert(units.PressureKilopascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilopascal = %v, want %v", converted, 100)
    }
    // Test Megapascal conversion
    megapascalsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureMegapascal,
    }
    
    var megapascalsResult *units.Pressure
    megapascalsResult, err = factory.FromDto(megapascalsDto)
    if err != nil {
        t.Errorf("FromDto() with Megapascal returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapascalsResult.Convert(units.PressureMegapascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megapascal = %v, want %v", converted, 100)
    }
    // Test Gigapascal conversion
    gigapascalsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureGigapascal,
    }
    
    var gigapascalsResult *units.Pressure
    gigapascalsResult, err = factory.FromDto(gigapascalsDto)
    if err != nil {
        t.Errorf("FromDto() with Gigapascal returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigapascalsResult.Convert(units.PressureGigapascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigapascal = %v, want %v", converted, 100)
    }
    // Test Microbar conversion
    microbarsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureMicrobar,
    }
    
    var microbarsResult *units.Pressure
    microbarsResult, err = factory.FromDto(microbarsDto)
    if err != nil {
        t.Errorf("FromDto() with Microbar returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microbarsResult.Convert(units.PressureMicrobar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microbar = %v, want %v", converted, 100)
    }
    // Test Millibar conversion
    millibarsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureMillibar,
    }
    
    var millibarsResult *units.Pressure
    millibarsResult, err = factory.FromDto(millibarsDto)
    if err != nil {
        t.Errorf("FromDto() with Millibar returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millibarsResult.Convert(units.PressureMillibar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millibar = %v, want %v", converted, 100)
    }
    // Test Centibar conversion
    centibarsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureCentibar,
    }
    
    var centibarsResult *units.Pressure
    centibarsResult, err = factory.FromDto(centibarsDto)
    if err != nil {
        t.Errorf("FromDto() with Centibar returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centibarsResult.Convert(units.PressureCentibar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centibar = %v, want %v", converted, 100)
    }
    // Test Decibar conversion
    decibarsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureDecibar,
    }
    
    var decibarsResult *units.Pressure
    decibarsResult, err = factory.FromDto(decibarsDto)
    if err != nil {
        t.Errorf("FromDto() with Decibar returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibarsResult.Convert(units.PressureDecibar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decibar = %v, want %v", converted, 100)
    }
    // Test Kilobar conversion
    kilobarsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureKilobar,
    }
    
    var kilobarsResult *units.Pressure
    kilobarsResult, err = factory.FromDto(kilobarsDto)
    if err != nil {
        t.Errorf("FromDto() with Kilobar returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobarsResult.Convert(units.PressureKilobar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilobar = %v, want %v", converted, 100)
    }
    // Test Megabar conversion
    megabarsDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureMegabar,
    }
    
    var megabarsResult *units.Pressure
    megabarsResult, err = factory.FromDto(megabarsDto)
    if err != nil {
        t.Errorf("FromDto() with Megabar returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabarsResult.Convert(units.PressureMegabar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megabar = %v, want %v", converted, 100)
    }
    // Test KilonewtonPerSquareMeter conversion
    kilonewtons_per_square_meterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureKilonewtonPerSquareMeter,
    }
    
    var kilonewtons_per_square_meterResult *units.Pressure
    kilonewtons_per_square_meterResult, err = factory.FromDto(kilonewtons_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_square_meterResult.Convert(units.PressureKilonewtonPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test MeganewtonPerSquareMeter conversion
    meganewtons_per_square_meterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureMeganewtonPerSquareMeter,
    }
    
    var meganewtons_per_square_meterResult *units.Pressure
    meganewtons_per_square_meterResult, err = factory.FromDto(meganewtons_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MeganewtonPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewtons_per_square_meterResult.Convert(units.PressureMeganewtonPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonPerSquareCentimeter conversion
    kilonewtons_per_square_centimeterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureKilonewtonPerSquareCentimeter,
    }
    
    var kilonewtons_per_square_centimeterResult *units.Pressure
    kilonewtons_per_square_centimeterResult, err = factory.FromDto(kilonewtons_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_square_centimeterResult.Convert(units.PressureKilonewtonPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test KilonewtonPerSquareMillimeter conversion
    kilonewtons_per_square_millimeterDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureKilonewtonPerSquareMillimeter,
    }
    
    var kilonewtons_per_square_millimeterResult *units.Pressure
    kilonewtons_per_square_millimeterResult, err = factory.FromDto(kilonewtons_per_square_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilonewtonPerSquareMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_square_millimeterResult.Convert(units.PressureKilonewtonPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test KilopoundForcePerSquareInch conversion
    kilopounds_force_per_square_inchDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureKilopoundForcePerSquareInch,
    }
    
    var kilopounds_force_per_square_inchResult *units.Pressure
    kilopounds_force_per_square_inchResult, err = factory.FromDto(kilopounds_force_per_square_inchDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForcePerSquareInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_square_inchResult.Convert(units.PressureKilopoundForcePerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerSquareInch = %v, want %v", converted, 100)
    }
    // Test KilopoundForcePerSquareMil conversion
    kilopounds_force_per_square_milDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureKilopoundForcePerSquareMil,
    }
    
    var kilopounds_force_per_square_milResult *units.Pressure
    kilopounds_force_per_square_milResult, err = factory.FromDto(kilopounds_force_per_square_milDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForcePerSquareMil returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_square_milResult.Convert(units.PressureKilopoundForcePerSquareMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerSquareMil = %v, want %v", converted, 100)
    }
    // Test KilopoundForcePerSquareFoot conversion
    kilopounds_force_per_square_footDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureKilopoundForcePerSquareFoot,
    }
    
    var kilopounds_force_per_square_footResult *units.Pressure
    kilopounds_force_per_square_footResult, err = factory.FromDto(kilopounds_force_per_square_footDto)
    if err != nil {
        t.Errorf("FromDto() with KilopoundForcePerSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_square_footResult.Convert(units.PressureKilopoundForcePerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerSquareFoot = %v, want %v", converted, 100)
    }
    // Test MillimeterOfWaterColumn conversion
    millimeters_of_water_columnDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureMillimeterOfWaterColumn,
    }
    
    var millimeters_of_water_columnResult *units.Pressure
    millimeters_of_water_columnResult, err = factory.FromDto(millimeters_of_water_columnDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterOfWaterColumn returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_of_water_columnResult.Convert(units.PressureMillimeterOfWaterColumn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterOfWaterColumn = %v, want %v", converted, 100)
    }
    // Test CentimeterOfWaterColumn conversion
    centimeters_of_water_columnDto := units.PressureDto{
        Value: 100,
        Unit:  units.PressureCentimeterOfWaterColumn,
    }
    
    var centimeters_of_water_columnResult *units.Pressure
    centimeters_of_water_columnResult, err = factory.FromDto(centimeters_of_water_columnDto)
    if err != nil {
        t.Errorf("FromDto() with CentimeterOfWaterColumn returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_of_water_columnResult.Convert(units.PressureCentimeterOfWaterColumn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterOfWaterColumn = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.PressureDto{
        Value: 0,
        Unit:  units.PressurePascal,
    }
    
    var zeroResult *units.Pressure
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestPressureFactory_FromDtoJSON(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Pascal"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Pascal"}`)
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
    nanJSON, _ := json.Marshal(units.PressureDto{
        Value: nanValue,
        Unit:  units.PressurePascal,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Pascal unit
    pascalsJSON := []byte(`{"value": 100, "unit": "Pascal"}`)
    pascalsResult, err := factory.FromDtoJSON(pascalsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Pascal unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascalsResult.Convert(units.PressurePascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Pascal = %v, want %v", converted, 100)
    }
    // Test JSON with Atmosphere unit
    atmospheresJSON := []byte(`{"value": 100, "unit": "Atmosphere"}`)
    atmospheresResult, err := factory.FromDtoJSON(atmospheresJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Atmosphere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = atmospheresResult.Convert(units.PressureAtmosphere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Atmosphere = %v, want %v", converted, 100)
    }
    // Test JSON with Bar unit
    barsJSON := []byte(`{"value": 100, "unit": "Bar"}`)
    barsResult, err := factory.FromDtoJSON(barsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Bar unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = barsResult.Convert(units.PressureBar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Bar = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForcePerSquareMeter unit
    kilograms_force_per_square_meterJSON := []byte(`{"value": 100, "unit": "KilogramForcePerSquareMeter"}`)
    kilograms_force_per_square_meterResult, err := factory.FromDtoJSON(kilograms_force_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForcePerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_square_meterResult.Convert(units.PressureKilogramForcePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForcePerSquareCentimeter unit
    kilograms_force_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "KilogramForcePerSquareCentimeter"}`)
    kilograms_force_per_square_centimeterResult, err := factory.FromDtoJSON(kilograms_force_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForcePerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_square_centimeterResult.Convert(units.PressureKilogramForcePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilogramForcePerSquareMillimeter unit
    kilograms_force_per_square_millimeterJSON := []byte(`{"value": 100, "unit": "KilogramForcePerSquareMillimeter"}`)
    kilograms_force_per_square_millimeterResult, err := factory.FromDtoJSON(kilograms_force_per_square_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilogramForcePerSquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilograms_force_per_square_millimeterResult.Convert(units.PressureKilogramForcePerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilogramForcePerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonPerSquareMeter unit
    newtons_per_square_meterJSON := []byte(`{"value": 100, "unit": "NewtonPerSquareMeter"}`)
    newtons_per_square_meterResult, err := factory.FromDtoJSON(newtons_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_square_meterResult.Convert(units.PressureNewtonPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonPerSquareCentimeter unit
    newtons_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "NewtonPerSquareCentimeter"}`)
    newtons_per_square_centimeterResult, err := factory.FromDtoJSON(newtons_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_square_centimeterResult.Convert(units.PressureNewtonPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with NewtonPerSquareMillimeter unit
    newtons_per_square_millimeterJSON := []byte(`{"value": 100, "unit": "NewtonPerSquareMillimeter"}`)
    newtons_per_square_millimeterResult, err := factory.FromDtoJSON(newtons_per_square_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonPerSquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newtons_per_square_millimeterResult.Convert(units.PressureNewtonPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonPerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with TechnicalAtmosphere unit
    technical_atmospheresJSON := []byte(`{"value": 100, "unit": "TechnicalAtmosphere"}`)
    technical_atmospheresResult, err := factory.FromDtoJSON(technical_atmospheresJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TechnicalAtmosphere unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = technical_atmospheresResult.Convert(units.PressureTechnicalAtmosphere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TechnicalAtmosphere = %v, want %v", converted, 100)
    }
    // Test JSON with Torr unit
    torrsJSON := []byte(`{"value": 100, "unit": "Torr"}`)
    torrsResult, err := factory.FromDtoJSON(torrsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Torr unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = torrsResult.Convert(units.PressureTorr)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Torr = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForcePerSquareInch unit
    pounds_force_per_square_inchJSON := []byte(`{"value": 100, "unit": "PoundForcePerSquareInch"}`)
    pounds_force_per_square_inchResult, err := factory.FromDtoJSON(pounds_force_per_square_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForcePerSquareInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_square_inchResult.Convert(units.PressurePoundForcePerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerSquareInch = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForcePerSquareMil unit
    pounds_force_per_square_milJSON := []byte(`{"value": 100, "unit": "PoundForcePerSquareMil"}`)
    pounds_force_per_square_milResult, err := factory.FromDtoJSON(pounds_force_per_square_milJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForcePerSquareMil unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_square_milResult.Convert(units.PressurePoundForcePerSquareMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerSquareMil = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForcePerSquareFoot unit
    pounds_force_per_square_footJSON := []byte(`{"value": 100, "unit": "PoundForcePerSquareFoot"}`)
    pounds_force_per_square_footResult, err := factory.FromDtoJSON(pounds_force_per_square_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForcePerSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_per_square_footResult.Convert(units.PressurePoundForcePerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForcePerSquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForcePerSquareMillimeter unit
    tonnes_force_per_square_millimeterJSON := []byte(`{"value": 100, "unit": "TonneForcePerSquareMillimeter"}`)
    tonnes_force_per_square_millimeterResult, err := factory.FromDtoJSON(tonnes_force_per_square_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForcePerSquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_square_millimeterResult.Convert(units.PressureTonneForcePerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForcePerSquareMeter unit
    tonnes_force_per_square_meterJSON := []byte(`{"value": 100, "unit": "TonneForcePerSquareMeter"}`)
    tonnes_force_per_square_meterResult, err := factory.FromDtoJSON(tonnes_force_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForcePerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_square_meterResult.Convert(units.PressureTonneForcePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MeterOfHead unit
    meters_of_headJSON := []byte(`{"value": 100, "unit": "MeterOfHead"}`)
    meters_of_headResult, err := factory.FromDtoJSON(meters_of_headJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeterOfHead unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_of_headResult.Convert(units.PressureMeterOfHead)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterOfHead = %v, want %v", converted, 100)
    }
    // Test JSON with TonneForcePerSquareCentimeter unit
    tonnes_force_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "TonneForcePerSquareCentimeter"}`)
    tonnes_force_per_square_centimeterResult, err := factory.FromDtoJSON(tonnes_force_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with TonneForcePerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tonnes_force_per_square_centimeterResult.Convert(units.PressureTonneForcePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for TonneForcePerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with FootOfHead unit
    feet_of_headJSON := []byte(`{"value": 100, "unit": "FootOfHead"}`)
    feet_of_headResult, err := factory.FromDtoJSON(feet_of_headJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FootOfHead unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_of_headResult.Convert(units.PressureFootOfHead)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootOfHead = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterOfMercury unit
    millimeters_of_mercuryJSON := []byte(`{"value": 100, "unit": "MillimeterOfMercury"}`)
    millimeters_of_mercuryResult, err := factory.FromDtoJSON(millimeters_of_mercuryJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterOfMercury unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_of_mercuryResult.Convert(units.PressureMillimeterOfMercury)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterOfMercury = %v, want %v", converted, 100)
    }
    // Test JSON with InchOfMercury unit
    inches_of_mercuryJSON := []byte(`{"value": 100, "unit": "InchOfMercury"}`)
    inches_of_mercuryResult, err := factory.FromDtoJSON(inches_of_mercuryJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InchOfMercury unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_of_mercuryResult.Convert(units.PressureInchOfMercury)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchOfMercury = %v, want %v", converted, 100)
    }
    // Test JSON with DynePerSquareCentimeter unit
    dynes_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "DynePerSquareCentimeter"}`)
    dynes_per_square_centimeterResult, err := factory.FromDtoJSON(dynes_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DynePerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = dynes_per_square_centimeterResult.Convert(units.PressureDynePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DynePerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerInchSecondSquared unit
    pounds_per_inch_second_squaredJSON := []byte(`{"value": 100, "unit": "PoundPerInchSecondSquared"}`)
    pounds_per_inch_second_squaredResult, err := factory.FromDtoJSON(pounds_per_inch_second_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerInchSecondSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_inch_second_squaredResult.Convert(units.PressurePoundPerInchSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerInchSecondSquared = %v, want %v", converted, 100)
    }
    // Test JSON with MeterOfWaterColumn unit
    meters_of_water_columnJSON := []byte(`{"value": 100, "unit": "MeterOfWaterColumn"}`)
    meters_of_water_columnResult, err := factory.FromDtoJSON(meters_of_water_columnJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeterOfWaterColumn unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_of_water_columnResult.Convert(units.PressureMeterOfWaterColumn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterOfWaterColumn = %v, want %v", converted, 100)
    }
    // Test JSON with InchOfWaterColumn unit
    inches_of_water_columnJSON := []byte(`{"value": 100, "unit": "InchOfWaterColumn"}`)
    inches_of_water_columnResult, err := factory.FromDtoJSON(inches_of_water_columnJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InchOfWaterColumn unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_of_water_columnResult.Convert(units.PressureInchOfWaterColumn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchOfWaterColumn = %v, want %v", converted, 100)
    }
    // Test JSON with MeterOfElevation unit
    meters_of_elevationJSON := []byte(`{"value": 100, "unit": "MeterOfElevation"}`)
    meters_of_elevationResult, err := factory.FromDtoJSON(meters_of_elevationJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeterOfElevation unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_of_elevationResult.Convert(units.PressureMeterOfElevation)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterOfElevation = %v, want %v", converted, 100)
    }
    // Test JSON with FootOfElevation unit
    feet_of_elevationJSON := []byte(`{"value": 100, "unit": "FootOfElevation"}`)
    feet_of_elevationResult, err := factory.FromDtoJSON(feet_of_elevationJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FootOfElevation unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_of_elevationResult.Convert(units.PressureFootOfElevation)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootOfElevation = %v, want %v", converted, 100)
    }
    // Test JSON with Micropascal unit
    micropascalsJSON := []byte(`{"value": 100, "unit": "Micropascal"}`)
    micropascalsResult, err := factory.FromDtoJSON(micropascalsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Micropascal unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micropascalsResult.Convert(units.PressureMicropascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Micropascal = %v, want %v", converted, 100)
    }
    // Test JSON with Millipascal unit
    millipascalsJSON := []byte(`{"value": 100, "unit": "Millipascal"}`)
    millipascalsResult, err := factory.FromDtoJSON(millipascalsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millipascal unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millipascalsResult.Convert(units.PressureMillipascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millipascal = %v, want %v", converted, 100)
    }
    // Test JSON with Decapascal unit
    decapascalsJSON := []byte(`{"value": 100, "unit": "Decapascal"}`)
    decapascalsResult, err := factory.FromDtoJSON(decapascalsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Decapascal unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decapascalsResult.Convert(units.PressureDecapascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decapascal = %v, want %v", converted, 100)
    }
    // Test JSON with Hectopascal unit
    hectopascalsJSON := []byte(`{"value": 100, "unit": "Hectopascal"}`)
    hectopascalsResult, err := factory.FromDtoJSON(hectopascalsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Hectopascal unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectopascalsResult.Convert(units.PressureHectopascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hectopascal = %v, want %v", converted, 100)
    }
    // Test JSON with Kilopascal unit
    kilopascalsJSON := []byte(`{"value": 100, "unit": "Kilopascal"}`)
    kilopascalsResult, err := factory.FromDtoJSON(kilopascalsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilopascal unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopascalsResult.Convert(units.PressureKilopascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilopascal = %v, want %v", converted, 100)
    }
    // Test JSON with Megapascal unit
    megapascalsJSON := []byte(`{"value": 100, "unit": "Megapascal"}`)
    megapascalsResult, err := factory.FromDtoJSON(megapascalsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megapascal unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megapascalsResult.Convert(units.PressureMegapascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megapascal = %v, want %v", converted, 100)
    }
    // Test JSON with Gigapascal unit
    gigapascalsJSON := []byte(`{"value": 100, "unit": "Gigapascal"}`)
    gigapascalsResult, err := factory.FromDtoJSON(gigapascalsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigapascal unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigapascalsResult.Convert(units.PressureGigapascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigapascal = %v, want %v", converted, 100)
    }
    // Test JSON with Microbar unit
    microbarsJSON := []byte(`{"value": 100, "unit": "Microbar"}`)
    microbarsResult, err := factory.FromDtoJSON(microbarsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microbar unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microbarsResult.Convert(units.PressureMicrobar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microbar = %v, want %v", converted, 100)
    }
    // Test JSON with Millibar unit
    millibarsJSON := []byte(`{"value": 100, "unit": "Millibar"}`)
    millibarsResult, err := factory.FromDtoJSON(millibarsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millibar unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millibarsResult.Convert(units.PressureMillibar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millibar = %v, want %v", converted, 100)
    }
    // Test JSON with Centibar unit
    centibarsJSON := []byte(`{"value": 100, "unit": "Centibar"}`)
    centibarsResult, err := factory.FromDtoJSON(centibarsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Centibar unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centibarsResult.Convert(units.PressureCentibar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centibar = %v, want %v", converted, 100)
    }
    // Test JSON with Decibar unit
    decibarsJSON := []byte(`{"value": 100, "unit": "Decibar"}`)
    decibarsResult, err := factory.FromDtoJSON(decibarsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Decibar unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decibarsResult.Convert(units.PressureDecibar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decibar = %v, want %v", converted, 100)
    }
    // Test JSON with Kilobar unit
    kilobarsJSON := []byte(`{"value": 100, "unit": "Kilobar"}`)
    kilobarsResult, err := factory.FromDtoJSON(kilobarsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilobar unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilobarsResult.Convert(units.PressureKilobar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilobar = %v, want %v", converted, 100)
    }
    // Test JSON with Megabar unit
    megabarsJSON := []byte(`{"value": 100, "unit": "Megabar"}`)
    megabarsResult, err := factory.FromDtoJSON(megabarsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megabar unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megabarsResult.Convert(units.PressureMegabar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megabar = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonPerSquareMeter unit
    kilonewtons_per_square_meterJSON := []byte(`{"value": 100, "unit": "KilonewtonPerSquareMeter"}`)
    kilonewtons_per_square_meterResult, err := factory.FromDtoJSON(kilonewtons_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_square_meterResult.Convert(units.PressureKilonewtonPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MeganewtonPerSquareMeter unit
    meganewtons_per_square_meterJSON := []byte(`{"value": 100, "unit": "MeganewtonPerSquareMeter"}`)
    meganewtons_per_square_meterResult, err := factory.FromDtoJSON(meganewtons_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeganewtonPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meganewtons_per_square_meterResult.Convert(units.PressureMeganewtonPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeganewtonPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonPerSquareCentimeter unit
    kilonewtons_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "KilonewtonPerSquareCentimeter"}`)
    kilonewtons_per_square_centimeterResult, err := factory.FromDtoJSON(kilonewtons_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_square_centimeterResult.Convert(units.PressureKilonewtonPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilonewtonPerSquareMillimeter unit
    kilonewtons_per_square_millimeterJSON := []byte(`{"value": 100, "unit": "KilonewtonPerSquareMillimeter"}`)
    kilonewtons_per_square_millimeterResult, err := factory.FromDtoJSON(kilonewtons_per_square_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilonewtonPerSquareMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilonewtons_per_square_millimeterResult.Convert(units.PressureKilonewtonPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilonewtonPerSquareMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForcePerSquareInch unit
    kilopounds_force_per_square_inchJSON := []byte(`{"value": 100, "unit": "KilopoundForcePerSquareInch"}`)
    kilopounds_force_per_square_inchResult, err := factory.FromDtoJSON(kilopounds_force_per_square_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForcePerSquareInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_square_inchResult.Convert(units.PressureKilopoundForcePerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerSquareInch = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForcePerSquareMil unit
    kilopounds_force_per_square_milJSON := []byte(`{"value": 100, "unit": "KilopoundForcePerSquareMil"}`)
    kilopounds_force_per_square_milResult, err := factory.FromDtoJSON(kilopounds_force_per_square_milJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForcePerSquareMil unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_square_milResult.Convert(units.PressureKilopoundForcePerSquareMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerSquareMil = %v, want %v", converted, 100)
    }
    // Test JSON with KilopoundForcePerSquareFoot unit
    kilopounds_force_per_square_footJSON := []byte(`{"value": 100, "unit": "KilopoundForcePerSquareFoot"}`)
    kilopounds_force_per_square_footResult, err := factory.FromDtoJSON(kilopounds_force_per_square_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilopoundForcePerSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilopounds_force_per_square_footResult.Convert(units.PressureKilopoundForcePerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilopoundForcePerSquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterOfWaterColumn unit
    millimeters_of_water_columnJSON := []byte(`{"value": 100, "unit": "MillimeterOfWaterColumn"}`)
    millimeters_of_water_columnResult, err := factory.FromDtoJSON(millimeters_of_water_columnJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterOfWaterColumn unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_of_water_columnResult.Convert(units.PressureMillimeterOfWaterColumn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterOfWaterColumn = %v, want %v", converted, 100)
    }
    // Test JSON with CentimeterOfWaterColumn unit
    centimeters_of_water_columnJSON := []byte(`{"value": 100, "unit": "CentimeterOfWaterColumn"}`)
    centimeters_of_water_columnResult, err := factory.FromDtoJSON(centimeters_of_water_columnJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentimeterOfWaterColumn unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_of_water_columnResult.Convert(units.PressureCentimeterOfWaterColumn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterOfWaterColumn = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Pascal"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromPascals function
func TestPressureFactory_FromPascals(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPascals(100)
    if err != nil {
        t.Errorf("FromPascals() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressurePascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPascals() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPascals(math.NaN())
    if err == nil {
        t.Error("FromPascals() with NaN value should return error")
    }

    _, err = factory.FromPascals(math.Inf(1))
    if err == nil {
        t.Error("FromPascals() with +Inf value should return error")
    }

    _, err = factory.FromPascals(math.Inf(-1))
    if err == nil {
        t.Error("FromPascals() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPascals(0)
    if err != nil {
        t.Errorf("FromPascals() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressurePascal)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPascals() with zero value = %v, want 0", converted)
    }
}
// Test FromAtmospheres function
func TestPressureFactory_FromAtmospheres(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAtmospheres(100)
    if err != nil {
        t.Errorf("FromAtmospheres() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureAtmosphere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAtmospheres() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAtmospheres(math.NaN())
    if err == nil {
        t.Error("FromAtmospheres() with NaN value should return error")
    }

    _, err = factory.FromAtmospheres(math.Inf(1))
    if err == nil {
        t.Error("FromAtmospheres() with +Inf value should return error")
    }

    _, err = factory.FromAtmospheres(math.Inf(-1))
    if err == nil {
        t.Error("FromAtmospheres() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAtmospheres(0)
    if err != nil {
        t.Errorf("FromAtmospheres() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureAtmosphere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAtmospheres() with zero value = %v, want 0", converted)
    }
}
// Test FromBars function
func TestPressureFactory_FromBars(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBars(100)
    if err != nil {
        t.Errorf("FromBars() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureBar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBars() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBars(math.NaN())
    if err == nil {
        t.Error("FromBars() with NaN value should return error")
    }

    _, err = factory.FromBars(math.Inf(1))
    if err == nil {
        t.Error("FromBars() with +Inf value should return error")
    }

    _, err = factory.FromBars(math.Inf(-1))
    if err == nil {
        t.Error("FromBars() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBars(0)
    if err != nil {
        t.Errorf("FromBars() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureBar)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBars() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsForcePerSquareMeter function
func TestPressureFactory_FromKilogramsForcePerSquareMeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsForcePerSquareMeter(100)
    if err != nil {
        t.Errorf("FromKilogramsForcePerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureKilogramForcePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsForcePerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsForcePerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsForcePerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsForcePerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsForcePerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsForcePerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsForcePerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsForcePerSquareMeter(0)
    if err != nil {
        t.Errorf("FromKilogramsForcePerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureKilogramForcePerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsForcePerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsForcePerSquareCentimeter function
func TestPressureFactory_FromKilogramsForcePerSquareCentimeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsForcePerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsForcePerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureKilogramForcePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsForcePerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsForcePerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsForcePerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsForcePerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsForcePerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsForcePerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsForcePerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsForcePerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsForcePerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureKilogramForcePerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsForcePerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilogramsForcePerSquareMillimeter function
func TestPressureFactory_FromKilogramsForcePerSquareMillimeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilogramsForcePerSquareMillimeter(100)
    if err != nil {
        t.Errorf("FromKilogramsForcePerSquareMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureKilogramForcePerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilogramsForcePerSquareMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilogramsForcePerSquareMillimeter(math.NaN())
    if err == nil {
        t.Error("FromKilogramsForcePerSquareMillimeter() with NaN value should return error")
    }

    _, err = factory.FromKilogramsForcePerSquareMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilogramsForcePerSquareMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilogramsForcePerSquareMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilogramsForcePerSquareMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilogramsForcePerSquareMillimeter(0)
    if err != nil {
        t.Errorf("FromKilogramsForcePerSquareMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureKilogramForcePerSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilogramsForcePerSquareMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonsPerSquareMeter function
func TestPressureFactory_FromNewtonsPerSquareMeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromNewtonsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureNewtonPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromNewtonsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromNewtonsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromNewtonsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromNewtonsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureNewtonPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonsPerSquareCentimeter function
func TestPressureFactory_FromNewtonsPerSquareCentimeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonsPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromNewtonsPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureNewtonPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonsPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonsPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromNewtonsPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromNewtonsPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonsPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromNewtonsPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonsPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonsPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromNewtonsPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureNewtonPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonsPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNewtonsPerSquareMillimeter function
func TestPressureFactory_FromNewtonsPerSquareMillimeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonsPerSquareMillimeter(100)
    if err != nil {
        t.Errorf("FromNewtonsPerSquareMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureNewtonPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonsPerSquareMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonsPerSquareMillimeter(math.NaN())
    if err == nil {
        t.Error("FromNewtonsPerSquareMillimeter() with NaN value should return error")
    }

    _, err = factory.FromNewtonsPerSquareMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonsPerSquareMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromNewtonsPerSquareMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonsPerSquareMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonsPerSquareMillimeter(0)
    if err != nil {
        t.Errorf("FromNewtonsPerSquareMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureNewtonPerSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonsPerSquareMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTechnicalAtmospheres function
func TestPressureFactory_FromTechnicalAtmospheres(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTechnicalAtmospheres(100)
    if err != nil {
        t.Errorf("FromTechnicalAtmospheres() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureTechnicalAtmosphere)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTechnicalAtmospheres() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTechnicalAtmospheres(math.NaN())
    if err == nil {
        t.Error("FromTechnicalAtmospheres() with NaN value should return error")
    }

    _, err = factory.FromTechnicalAtmospheres(math.Inf(1))
    if err == nil {
        t.Error("FromTechnicalAtmospheres() with +Inf value should return error")
    }

    _, err = factory.FromTechnicalAtmospheres(math.Inf(-1))
    if err == nil {
        t.Error("FromTechnicalAtmospheres() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTechnicalAtmospheres(0)
    if err != nil {
        t.Errorf("FromTechnicalAtmospheres() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureTechnicalAtmosphere)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTechnicalAtmospheres() with zero value = %v, want 0", converted)
    }
}
// Test FromTorrs function
func TestPressureFactory_FromTorrs(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTorrs(100)
    if err != nil {
        t.Errorf("FromTorrs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureTorr)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTorrs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTorrs(math.NaN())
    if err == nil {
        t.Error("FromTorrs() with NaN value should return error")
    }

    _, err = factory.FromTorrs(math.Inf(1))
    if err == nil {
        t.Error("FromTorrs() with +Inf value should return error")
    }

    _, err = factory.FromTorrs(math.Inf(-1))
    if err == nil {
        t.Error("FromTorrs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTorrs(0)
    if err != nil {
        t.Errorf("FromTorrs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureTorr)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTorrs() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForcePerSquareInch function
func TestPressureFactory_FromPoundsForcePerSquareInch(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForcePerSquareInch(100)
    if err != nil {
        t.Errorf("FromPoundsForcePerSquareInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressurePoundForcePerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForcePerSquareInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForcePerSquareInch(math.NaN())
    if err == nil {
        t.Error("FromPoundsForcePerSquareInch() with NaN value should return error")
    }

    _, err = factory.FromPoundsForcePerSquareInch(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForcePerSquareInch() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForcePerSquareInch(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForcePerSquareInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForcePerSquareInch(0)
    if err != nil {
        t.Errorf("FromPoundsForcePerSquareInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressurePoundForcePerSquareInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForcePerSquareInch() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForcePerSquareMil function
func TestPressureFactory_FromPoundsForcePerSquareMil(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForcePerSquareMil(100)
    if err != nil {
        t.Errorf("FromPoundsForcePerSquareMil() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressurePoundForcePerSquareMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForcePerSquareMil() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForcePerSquareMil(math.NaN())
    if err == nil {
        t.Error("FromPoundsForcePerSquareMil() with NaN value should return error")
    }

    _, err = factory.FromPoundsForcePerSquareMil(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForcePerSquareMil() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForcePerSquareMil(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForcePerSquareMil() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForcePerSquareMil(0)
    if err != nil {
        t.Errorf("FromPoundsForcePerSquareMil() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressurePoundForcePerSquareMil)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForcePerSquareMil() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForcePerSquareFoot function
func TestPressureFactory_FromPoundsForcePerSquareFoot(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForcePerSquareFoot(100)
    if err != nil {
        t.Errorf("FromPoundsForcePerSquareFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressurePoundForcePerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForcePerSquareFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForcePerSquareFoot(math.NaN())
    if err == nil {
        t.Error("FromPoundsForcePerSquareFoot() with NaN value should return error")
    }

    _, err = factory.FromPoundsForcePerSquareFoot(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForcePerSquareFoot() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForcePerSquareFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForcePerSquareFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForcePerSquareFoot(0)
    if err != nil {
        t.Errorf("FromPoundsForcePerSquareFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressurePoundForcePerSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForcePerSquareFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesForcePerSquareMillimeter function
func TestPressureFactory_FromTonnesForcePerSquareMillimeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesForcePerSquareMillimeter(100)
    if err != nil {
        t.Errorf("FromTonnesForcePerSquareMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureTonneForcePerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesForcePerSquareMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesForcePerSquareMillimeter(math.NaN())
    if err == nil {
        t.Error("FromTonnesForcePerSquareMillimeter() with NaN value should return error")
    }

    _, err = factory.FromTonnesForcePerSquareMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesForcePerSquareMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromTonnesForcePerSquareMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesForcePerSquareMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesForcePerSquareMillimeter(0)
    if err != nil {
        t.Errorf("FromTonnesForcePerSquareMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureTonneForcePerSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesForcePerSquareMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesForcePerSquareMeter function
func TestPressureFactory_FromTonnesForcePerSquareMeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesForcePerSquareMeter(100)
    if err != nil {
        t.Errorf("FromTonnesForcePerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureTonneForcePerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesForcePerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesForcePerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromTonnesForcePerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromTonnesForcePerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesForcePerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromTonnesForcePerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesForcePerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesForcePerSquareMeter(0)
    if err != nil {
        t.Errorf("FromTonnesForcePerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureTonneForcePerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesForcePerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMetersOfHead function
func TestPressureFactory_FromMetersOfHead(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMetersOfHead(100)
    if err != nil {
        t.Errorf("FromMetersOfHead() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureMeterOfHead)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMetersOfHead() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMetersOfHead(math.NaN())
    if err == nil {
        t.Error("FromMetersOfHead() with NaN value should return error")
    }

    _, err = factory.FromMetersOfHead(math.Inf(1))
    if err == nil {
        t.Error("FromMetersOfHead() with +Inf value should return error")
    }

    _, err = factory.FromMetersOfHead(math.Inf(-1))
    if err == nil {
        t.Error("FromMetersOfHead() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMetersOfHead(0)
    if err != nil {
        t.Errorf("FromMetersOfHead() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureMeterOfHead)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMetersOfHead() with zero value = %v, want 0", converted)
    }
}
// Test FromTonnesForcePerSquareCentimeter function
func TestPressureFactory_FromTonnesForcePerSquareCentimeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTonnesForcePerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromTonnesForcePerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureTonneForcePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTonnesForcePerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTonnesForcePerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromTonnesForcePerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromTonnesForcePerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromTonnesForcePerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromTonnesForcePerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromTonnesForcePerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTonnesForcePerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromTonnesForcePerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureTonneForcePerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTonnesForcePerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromFeetOfHead function
func TestPressureFactory_FromFeetOfHead(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFeetOfHead(100)
    if err != nil {
        t.Errorf("FromFeetOfHead() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureFootOfHead)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFeetOfHead() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFeetOfHead(math.NaN())
    if err == nil {
        t.Error("FromFeetOfHead() with NaN value should return error")
    }

    _, err = factory.FromFeetOfHead(math.Inf(1))
    if err == nil {
        t.Error("FromFeetOfHead() with +Inf value should return error")
    }

    _, err = factory.FromFeetOfHead(math.Inf(-1))
    if err == nil {
        t.Error("FromFeetOfHead() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFeetOfHead(0)
    if err != nil {
        t.Errorf("FromFeetOfHead() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureFootOfHead)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFeetOfHead() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimetersOfMercury function
func TestPressureFactory_FromMillimetersOfMercury(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimetersOfMercury(100)
    if err != nil {
        t.Errorf("FromMillimetersOfMercury() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureMillimeterOfMercury)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimetersOfMercury() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimetersOfMercury(math.NaN())
    if err == nil {
        t.Error("FromMillimetersOfMercury() with NaN value should return error")
    }

    _, err = factory.FromMillimetersOfMercury(math.Inf(1))
    if err == nil {
        t.Error("FromMillimetersOfMercury() with +Inf value should return error")
    }

    _, err = factory.FromMillimetersOfMercury(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimetersOfMercury() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimetersOfMercury(0)
    if err != nil {
        t.Errorf("FromMillimetersOfMercury() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureMillimeterOfMercury)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimetersOfMercury() with zero value = %v, want 0", converted)
    }
}
// Test FromInchesOfMercury function
func TestPressureFactory_FromInchesOfMercury(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInchesOfMercury(100)
    if err != nil {
        t.Errorf("FromInchesOfMercury() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureInchOfMercury)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInchesOfMercury() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInchesOfMercury(math.NaN())
    if err == nil {
        t.Error("FromInchesOfMercury() with NaN value should return error")
    }

    _, err = factory.FromInchesOfMercury(math.Inf(1))
    if err == nil {
        t.Error("FromInchesOfMercury() with +Inf value should return error")
    }

    _, err = factory.FromInchesOfMercury(math.Inf(-1))
    if err == nil {
        t.Error("FromInchesOfMercury() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInchesOfMercury(0)
    if err != nil {
        t.Errorf("FromInchesOfMercury() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureInchOfMercury)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInchesOfMercury() with zero value = %v, want 0", converted)
    }
}
// Test FromDynesPerSquareCentimeter function
func TestPressureFactory_FromDynesPerSquareCentimeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDynesPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromDynesPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureDynePerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDynesPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDynesPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromDynesPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromDynesPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromDynesPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromDynesPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromDynesPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDynesPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromDynesPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureDynePerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDynesPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerInchSecondSquared function
func TestPressureFactory_FromPoundsPerInchSecondSquared(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerInchSecondSquared(100)
    if err != nil {
        t.Errorf("FromPoundsPerInchSecondSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressurePoundPerInchSecondSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerInchSecondSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerInchSecondSquared(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerInchSecondSquared() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerInchSecondSquared(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerInchSecondSquared() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerInchSecondSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerInchSecondSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerInchSecondSquared(0)
    if err != nil {
        t.Errorf("FromPoundsPerInchSecondSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressurePoundPerInchSecondSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerInchSecondSquared() with zero value = %v, want 0", converted)
    }
}
// Test FromMetersOfWaterColumn function
func TestPressureFactory_FromMetersOfWaterColumn(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMetersOfWaterColumn(100)
    if err != nil {
        t.Errorf("FromMetersOfWaterColumn() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureMeterOfWaterColumn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMetersOfWaterColumn() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMetersOfWaterColumn(math.NaN())
    if err == nil {
        t.Error("FromMetersOfWaterColumn() with NaN value should return error")
    }

    _, err = factory.FromMetersOfWaterColumn(math.Inf(1))
    if err == nil {
        t.Error("FromMetersOfWaterColumn() with +Inf value should return error")
    }

    _, err = factory.FromMetersOfWaterColumn(math.Inf(-1))
    if err == nil {
        t.Error("FromMetersOfWaterColumn() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMetersOfWaterColumn(0)
    if err != nil {
        t.Errorf("FromMetersOfWaterColumn() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureMeterOfWaterColumn)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMetersOfWaterColumn() with zero value = %v, want 0", converted)
    }
}
// Test FromInchesOfWaterColumn function
func TestPressureFactory_FromInchesOfWaterColumn(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInchesOfWaterColumn(100)
    if err != nil {
        t.Errorf("FromInchesOfWaterColumn() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureInchOfWaterColumn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInchesOfWaterColumn() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInchesOfWaterColumn(math.NaN())
    if err == nil {
        t.Error("FromInchesOfWaterColumn() with NaN value should return error")
    }

    _, err = factory.FromInchesOfWaterColumn(math.Inf(1))
    if err == nil {
        t.Error("FromInchesOfWaterColumn() with +Inf value should return error")
    }

    _, err = factory.FromInchesOfWaterColumn(math.Inf(-1))
    if err == nil {
        t.Error("FromInchesOfWaterColumn() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInchesOfWaterColumn(0)
    if err != nil {
        t.Errorf("FromInchesOfWaterColumn() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureInchOfWaterColumn)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInchesOfWaterColumn() with zero value = %v, want 0", converted)
    }
}
// Test FromMetersOfElevation function
func TestPressureFactory_FromMetersOfElevation(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMetersOfElevation(100)
    if err != nil {
        t.Errorf("FromMetersOfElevation() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureMeterOfElevation)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMetersOfElevation() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMetersOfElevation(math.NaN())
    if err == nil {
        t.Error("FromMetersOfElevation() with NaN value should return error")
    }

    _, err = factory.FromMetersOfElevation(math.Inf(1))
    if err == nil {
        t.Error("FromMetersOfElevation() with +Inf value should return error")
    }

    _, err = factory.FromMetersOfElevation(math.Inf(-1))
    if err == nil {
        t.Error("FromMetersOfElevation() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMetersOfElevation(0)
    if err != nil {
        t.Errorf("FromMetersOfElevation() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureMeterOfElevation)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMetersOfElevation() with zero value = %v, want 0", converted)
    }
}
// Test FromFeetOfElevation function
func TestPressureFactory_FromFeetOfElevation(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFeetOfElevation(100)
    if err != nil {
        t.Errorf("FromFeetOfElevation() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureFootOfElevation)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFeetOfElevation() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFeetOfElevation(math.NaN())
    if err == nil {
        t.Error("FromFeetOfElevation() with NaN value should return error")
    }

    _, err = factory.FromFeetOfElevation(math.Inf(1))
    if err == nil {
        t.Error("FromFeetOfElevation() with +Inf value should return error")
    }

    _, err = factory.FromFeetOfElevation(math.Inf(-1))
    if err == nil {
        t.Error("FromFeetOfElevation() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFeetOfElevation(0)
    if err != nil {
        t.Errorf("FromFeetOfElevation() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureFootOfElevation)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFeetOfElevation() with zero value = %v, want 0", converted)
    }
}
// Test FromMicropascals function
func TestPressureFactory_FromMicropascals(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicropascals(100)
    if err != nil {
        t.Errorf("FromMicropascals() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureMicropascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicropascals() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicropascals(math.NaN())
    if err == nil {
        t.Error("FromMicropascals() with NaN value should return error")
    }

    _, err = factory.FromMicropascals(math.Inf(1))
    if err == nil {
        t.Error("FromMicropascals() with +Inf value should return error")
    }

    _, err = factory.FromMicropascals(math.Inf(-1))
    if err == nil {
        t.Error("FromMicropascals() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicropascals(0)
    if err != nil {
        t.Errorf("FromMicropascals() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureMicropascal)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicropascals() with zero value = %v, want 0", converted)
    }
}
// Test FromMillipascals function
func TestPressureFactory_FromMillipascals(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillipascals(100)
    if err != nil {
        t.Errorf("FromMillipascals() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureMillipascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillipascals() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillipascals(math.NaN())
    if err == nil {
        t.Error("FromMillipascals() with NaN value should return error")
    }

    _, err = factory.FromMillipascals(math.Inf(1))
    if err == nil {
        t.Error("FromMillipascals() with +Inf value should return error")
    }

    _, err = factory.FromMillipascals(math.Inf(-1))
    if err == nil {
        t.Error("FromMillipascals() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillipascals(0)
    if err != nil {
        t.Errorf("FromMillipascals() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureMillipascal)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillipascals() with zero value = %v, want 0", converted)
    }
}
// Test FromDecapascals function
func TestPressureFactory_FromDecapascals(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecapascals(100)
    if err != nil {
        t.Errorf("FromDecapascals() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureDecapascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecapascals() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecapascals(math.NaN())
    if err == nil {
        t.Error("FromDecapascals() with NaN value should return error")
    }

    _, err = factory.FromDecapascals(math.Inf(1))
    if err == nil {
        t.Error("FromDecapascals() with +Inf value should return error")
    }

    _, err = factory.FromDecapascals(math.Inf(-1))
    if err == nil {
        t.Error("FromDecapascals() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecapascals(0)
    if err != nil {
        t.Errorf("FromDecapascals() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureDecapascal)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecapascals() with zero value = %v, want 0", converted)
    }
}
// Test FromHectopascals function
func TestPressureFactory_FromHectopascals(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectopascals(100)
    if err != nil {
        t.Errorf("FromHectopascals() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureHectopascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectopascals() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectopascals(math.NaN())
    if err == nil {
        t.Error("FromHectopascals() with NaN value should return error")
    }

    _, err = factory.FromHectopascals(math.Inf(1))
    if err == nil {
        t.Error("FromHectopascals() with +Inf value should return error")
    }

    _, err = factory.FromHectopascals(math.Inf(-1))
    if err == nil {
        t.Error("FromHectopascals() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectopascals(0)
    if err != nil {
        t.Errorf("FromHectopascals() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureHectopascal)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectopascals() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopascals function
func TestPressureFactory_FromKilopascals(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopascals(100)
    if err != nil {
        t.Errorf("FromKilopascals() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureKilopascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopascals() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopascals(math.NaN())
    if err == nil {
        t.Error("FromKilopascals() with NaN value should return error")
    }

    _, err = factory.FromKilopascals(math.Inf(1))
    if err == nil {
        t.Error("FromKilopascals() with +Inf value should return error")
    }

    _, err = factory.FromKilopascals(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopascals() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopascals(0)
    if err != nil {
        t.Errorf("FromKilopascals() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureKilopascal)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopascals() with zero value = %v, want 0", converted)
    }
}
// Test FromMegapascals function
func TestPressureFactory_FromMegapascals(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegapascals(100)
    if err != nil {
        t.Errorf("FromMegapascals() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureMegapascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegapascals() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegapascals(math.NaN())
    if err == nil {
        t.Error("FromMegapascals() with NaN value should return error")
    }

    _, err = factory.FromMegapascals(math.Inf(1))
    if err == nil {
        t.Error("FromMegapascals() with +Inf value should return error")
    }

    _, err = factory.FromMegapascals(math.Inf(-1))
    if err == nil {
        t.Error("FromMegapascals() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegapascals(0)
    if err != nil {
        t.Errorf("FromMegapascals() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureMegapascal)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegapascals() with zero value = %v, want 0", converted)
    }
}
// Test FromGigapascals function
func TestPressureFactory_FromGigapascals(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigapascals(100)
    if err != nil {
        t.Errorf("FromGigapascals() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureGigapascal)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigapascals() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigapascals(math.NaN())
    if err == nil {
        t.Error("FromGigapascals() with NaN value should return error")
    }

    _, err = factory.FromGigapascals(math.Inf(1))
    if err == nil {
        t.Error("FromGigapascals() with +Inf value should return error")
    }

    _, err = factory.FromGigapascals(math.Inf(-1))
    if err == nil {
        t.Error("FromGigapascals() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigapascals(0)
    if err != nil {
        t.Errorf("FromGigapascals() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureGigapascal)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigapascals() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrobars function
func TestPressureFactory_FromMicrobars(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrobars(100)
    if err != nil {
        t.Errorf("FromMicrobars() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureMicrobar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrobars() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrobars(math.NaN())
    if err == nil {
        t.Error("FromMicrobars() with NaN value should return error")
    }

    _, err = factory.FromMicrobars(math.Inf(1))
    if err == nil {
        t.Error("FromMicrobars() with +Inf value should return error")
    }

    _, err = factory.FromMicrobars(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrobars() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrobars(0)
    if err != nil {
        t.Errorf("FromMicrobars() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureMicrobar)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrobars() with zero value = %v, want 0", converted)
    }
}
// Test FromMillibars function
func TestPressureFactory_FromMillibars(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillibars(100)
    if err != nil {
        t.Errorf("FromMillibars() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureMillibar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillibars() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillibars(math.NaN())
    if err == nil {
        t.Error("FromMillibars() with NaN value should return error")
    }

    _, err = factory.FromMillibars(math.Inf(1))
    if err == nil {
        t.Error("FromMillibars() with +Inf value should return error")
    }

    _, err = factory.FromMillibars(math.Inf(-1))
    if err == nil {
        t.Error("FromMillibars() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillibars(0)
    if err != nil {
        t.Errorf("FromMillibars() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureMillibar)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillibars() with zero value = %v, want 0", converted)
    }
}
// Test FromCentibars function
func TestPressureFactory_FromCentibars(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentibars(100)
    if err != nil {
        t.Errorf("FromCentibars() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureCentibar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentibars() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentibars(math.NaN())
    if err == nil {
        t.Error("FromCentibars() with NaN value should return error")
    }

    _, err = factory.FromCentibars(math.Inf(1))
    if err == nil {
        t.Error("FromCentibars() with +Inf value should return error")
    }

    _, err = factory.FromCentibars(math.Inf(-1))
    if err == nil {
        t.Error("FromCentibars() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentibars(0)
    if err != nil {
        t.Errorf("FromCentibars() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureCentibar)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentibars() with zero value = %v, want 0", converted)
    }
}
// Test FromDecibars function
func TestPressureFactory_FromDecibars(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecibars(100)
    if err != nil {
        t.Errorf("FromDecibars() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureDecibar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecibars() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecibars(math.NaN())
    if err == nil {
        t.Error("FromDecibars() with NaN value should return error")
    }

    _, err = factory.FromDecibars(math.Inf(1))
    if err == nil {
        t.Error("FromDecibars() with +Inf value should return error")
    }

    _, err = factory.FromDecibars(math.Inf(-1))
    if err == nil {
        t.Error("FromDecibars() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecibars(0)
    if err != nil {
        t.Errorf("FromDecibars() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureDecibar)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecibars() with zero value = %v, want 0", converted)
    }
}
// Test FromKilobars function
func TestPressureFactory_FromKilobars(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilobars(100)
    if err != nil {
        t.Errorf("FromKilobars() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureKilobar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilobars() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilobars(math.NaN())
    if err == nil {
        t.Error("FromKilobars() with NaN value should return error")
    }

    _, err = factory.FromKilobars(math.Inf(1))
    if err == nil {
        t.Error("FromKilobars() with +Inf value should return error")
    }

    _, err = factory.FromKilobars(math.Inf(-1))
    if err == nil {
        t.Error("FromKilobars() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilobars(0)
    if err != nil {
        t.Errorf("FromKilobars() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureKilobar)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilobars() with zero value = %v, want 0", converted)
    }
}
// Test FromMegabars function
func TestPressureFactory_FromMegabars(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegabars(100)
    if err != nil {
        t.Errorf("FromMegabars() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureMegabar)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegabars() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegabars(math.NaN())
    if err == nil {
        t.Error("FromMegabars() with NaN value should return error")
    }

    _, err = factory.FromMegabars(math.Inf(1))
    if err == nil {
        t.Error("FromMegabars() with +Inf value should return error")
    }

    _, err = factory.FromMegabars(math.Inf(-1))
    if err == nil {
        t.Error("FromMegabars() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegabars(0)
    if err != nil {
        t.Errorf("FromMegabars() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureMegabar)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegabars() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonsPerSquareMeter function
func TestPressureFactory_FromKilonewtonsPerSquareMeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromKilonewtonsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureKilonewtonPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromKilonewtonsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureKilonewtonPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMeganewtonsPerSquareMeter function
func TestPressureFactory_FromMeganewtonsPerSquareMeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeganewtonsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromMeganewtonsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureMeganewtonPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeganewtonsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeganewtonsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromMeganewtonsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromMeganewtonsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMeganewtonsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromMeganewtonsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMeganewtonsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeganewtonsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromMeganewtonsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureMeganewtonPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeganewtonsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonsPerSquareCentimeter function
func TestPressureFactory_FromKilonewtonsPerSquareCentimeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonsPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromKilonewtonsPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureKilonewtonPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonsPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonsPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonsPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonsPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonsPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonsPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonsPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonsPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromKilonewtonsPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureKilonewtonPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonsPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilonewtonsPerSquareMillimeter function
func TestPressureFactory_FromKilonewtonsPerSquareMillimeter(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilonewtonsPerSquareMillimeter(100)
    if err != nil {
        t.Errorf("FromKilonewtonsPerSquareMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureKilonewtonPerSquareMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilonewtonsPerSquareMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilonewtonsPerSquareMillimeter(math.NaN())
    if err == nil {
        t.Error("FromKilonewtonsPerSquareMillimeter() with NaN value should return error")
    }

    _, err = factory.FromKilonewtonsPerSquareMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilonewtonsPerSquareMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilonewtonsPerSquareMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilonewtonsPerSquareMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilonewtonsPerSquareMillimeter(0)
    if err != nil {
        t.Errorf("FromKilonewtonsPerSquareMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureKilonewtonPerSquareMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilonewtonsPerSquareMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsForcePerSquareInch function
func TestPressureFactory_FromKilopoundsForcePerSquareInch(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsForcePerSquareInch(100)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerSquareInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureKilopoundForcePerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerSquareInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsForcePerSquareInch(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareInch() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsForcePerSquareInch(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareInch() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsForcePerSquareInch(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsForcePerSquareInch(0)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerSquareInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureKilopoundForcePerSquareInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerSquareInch() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsForcePerSquareMil function
func TestPressureFactory_FromKilopoundsForcePerSquareMil(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsForcePerSquareMil(100)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerSquareMil() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureKilopoundForcePerSquareMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerSquareMil() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsForcePerSquareMil(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareMil() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsForcePerSquareMil(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareMil() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsForcePerSquareMil(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareMil() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsForcePerSquareMil(0)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerSquareMil() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureKilopoundForcePerSquareMil)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerSquareMil() with zero value = %v, want 0", converted)
    }
}
// Test FromKilopoundsForcePerSquareFoot function
func TestPressureFactory_FromKilopoundsForcePerSquareFoot(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilopoundsForcePerSquareFoot(100)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerSquareFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureKilopoundForcePerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerSquareFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilopoundsForcePerSquareFoot(math.NaN())
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareFoot() with NaN value should return error")
    }

    _, err = factory.FromKilopoundsForcePerSquareFoot(math.Inf(1))
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareFoot() with +Inf value should return error")
    }

    _, err = factory.FromKilopoundsForcePerSquareFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromKilopoundsForcePerSquareFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilopoundsForcePerSquareFoot(0)
    if err != nil {
        t.Errorf("FromKilopoundsForcePerSquareFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureKilopoundForcePerSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilopoundsForcePerSquareFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimetersOfWaterColumn function
func TestPressureFactory_FromMillimetersOfWaterColumn(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimetersOfWaterColumn(100)
    if err != nil {
        t.Errorf("FromMillimetersOfWaterColumn() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureMillimeterOfWaterColumn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimetersOfWaterColumn() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimetersOfWaterColumn(math.NaN())
    if err == nil {
        t.Error("FromMillimetersOfWaterColumn() with NaN value should return error")
    }

    _, err = factory.FromMillimetersOfWaterColumn(math.Inf(1))
    if err == nil {
        t.Error("FromMillimetersOfWaterColumn() with +Inf value should return error")
    }

    _, err = factory.FromMillimetersOfWaterColumn(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimetersOfWaterColumn() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimetersOfWaterColumn(0)
    if err != nil {
        t.Errorf("FromMillimetersOfWaterColumn() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureMillimeterOfWaterColumn)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimetersOfWaterColumn() with zero value = %v, want 0", converted)
    }
}
// Test FromCentimetersOfWaterColumn function
func TestPressureFactory_FromCentimetersOfWaterColumn(t *testing.T) {
    factory := units.PressureFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentimetersOfWaterColumn(100)
    if err != nil {
        t.Errorf("FromCentimetersOfWaterColumn() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.PressureCentimeterOfWaterColumn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentimetersOfWaterColumn() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentimetersOfWaterColumn(math.NaN())
    if err == nil {
        t.Error("FromCentimetersOfWaterColumn() with NaN value should return error")
    }

    _, err = factory.FromCentimetersOfWaterColumn(math.Inf(1))
    if err == nil {
        t.Error("FromCentimetersOfWaterColumn() with +Inf value should return error")
    }

    _, err = factory.FromCentimetersOfWaterColumn(math.Inf(-1))
    if err == nil {
        t.Error("FromCentimetersOfWaterColumn() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentimetersOfWaterColumn(0)
    if err != nil {
        t.Errorf("FromCentimetersOfWaterColumn() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.PressureCentimeterOfWaterColumn)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentimetersOfWaterColumn() with zero value = %v, want 0", converted)
    }
}

func TestPressureToString(t *testing.T) {
	factory := units.PressureFactory{}
	a, err := factory.CreatePressure(45, units.PressurePascal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.PressurePascal, 2)
	expected := "45.00 " + units.GetPressureAbbreviation(units.PressurePascal)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.PressurePascal, -1)
	expected = "45 " + units.GetPressureAbbreviation(units.PressurePascal)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestPressure_EqualityAndComparison(t *testing.T) {
	factory := units.PressureFactory{}
	a1, _ := factory.CreatePressure(60, units.PressurePascal)
	a2, _ := factory.CreatePressure(60, units.PressurePascal)
	a3, _ := factory.CreatePressure(120, units.PressurePascal)

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

func TestPressure_Arithmetic(t *testing.T) {
	factory := units.PressureFactory{}
	a1, _ := factory.CreatePressure(30, units.PressurePascal)
	a2, _ := factory.CreatePressure(45, units.PressurePascal)

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


func TestGetPressureAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.PressureUnits
        want string
    }{
        {
            name: "Pascal abbreviation",
            unit: units.PressurePascal,
            want: "Pa",
        },
        {
            name: "Atmosphere abbreviation",
            unit: units.PressureAtmosphere,
            want: "atm",
        },
        {
            name: "Bar abbreviation",
            unit: units.PressureBar,
            want: "bar",
        },
        {
            name: "KilogramForcePerSquareMeter abbreviation",
            unit: units.PressureKilogramForcePerSquareMeter,
            want: "kgf/m²",
        },
        {
            name: "KilogramForcePerSquareCentimeter abbreviation",
            unit: units.PressureKilogramForcePerSquareCentimeter,
            want: "kgf/cm²",
        },
        {
            name: "KilogramForcePerSquareMillimeter abbreviation",
            unit: units.PressureKilogramForcePerSquareMillimeter,
            want: "kgf/mm²",
        },
        {
            name: "NewtonPerSquareMeter abbreviation",
            unit: units.PressureNewtonPerSquareMeter,
            want: "N/m²",
        },
        {
            name: "NewtonPerSquareCentimeter abbreviation",
            unit: units.PressureNewtonPerSquareCentimeter,
            want: "N/cm²",
        },
        {
            name: "NewtonPerSquareMillimeter abbreviation",
            unit: units.PressureNewtonPerSquareMillimeter,
            want: "N/mm²",
        },
        {
            name: "TechnicalAtmosphere abbreviation",
            unit: units.PressureTechnicalAtmosphere,
            want: "at",
        },
        {
            name: "Torr abbreviation",
            unit: units.PressureTorr,
            want: "torr",
        },
        {
            name: "PoundForcePerSquareInch abbreviation",
            unit: units.PressurePoundForcePerSquareInch,
            want: "psi",
        },
        {
            name: "PoundForcePerSquareMil abbreviation",
            unit: units.PressurePoundForcePerSquareMil,
            want: "lb/mil²",
        },
        {
            name: "PoundForcePerSquareFoot abbreviation",
            unit: units.PressurePoundForcePerSquareFoot,
            want: "lb/ft²",
        },
        {
            name: "TonneForcePerSquareMillimeter abbreviation",
            unit: units.PressureTonneForcePerSquareMillimeter,
            want: "tf/mm²",
        },
        {
            name: "TonneForcePerSquareMeter abbreviation",
            unit: units.PressureTonneForcePerSquareMeter,
            want: "tf/m²",
        },
        {
            name: "MeterOfHead abbreviation",
            unit: units.PressureMeterOfHead,
            want: "m of head",
        },
        {
            name: "TonneForcePerSquareCentimeter abbreviation",
            unit: units.PressureTonneForcePerSquareCentimeter,
            want: "tf/cm²",
        },
        {
            name: "FootOfHead abbreviation",
            unit: units.PressureFootOfHead,
            want: "ft of head",
        },
        {
            name: "MillimeterOfMercury abbreviation",
            unit: units.PressureMillimeterOfMercury,
            want: "mmHg",
        },
        {
            name: "InchOfMercury abbreviation",
            unit: units.PressureInchOfMercury,
            want: "inHg",
        },
        {
            name: "DynePerSquareCentimeter abbreviation",
            unit: units.PressureDynePerSquareCentimeter,
            want: "dyn/cm²",
        },
        {
            name: "PoundPerInchSecondSquared abbreviation",
            unit: units.PressurePoundPerInchSecondSquared,
            want: "lbm/(in·s²)",
        },
        {
            name: "MeterOfWaterColumn abbreviation",
            unit: units.PressureMeterOfWaterColumn,
            want: "mH₂O",
        },
        {
            name: "InchOfWaterColumn abbreviation",
            unit: units.PressureInchOfWaterColumn,
            want: "inH2O",
        },
        {
            name: "MeterOfElevation abbreviation",
            unit: units.PressureMeterOfElevation,
            want: "m of elevation",
        },
        {
            name: "FootOfElevation abbreviation",
            unit: units.PressureFootOfElevation,
            want: "ft of elevation",
        },
        {
            name: "Micropascal abbreviation",
            unit: units.PressureMicropascal,
            want: "μPa",
        },
        {
            name: "Millipascal abbreviation",
            unit: units.PressureMillipascal,
            want: "mPa",
        },
        {
            name: "Decapascal abbreviation",
            unit: units.PressureDecapascal,
            want: "daPa",
        },
        {
            name: "Hectopascal abbreviation",
            unit: units.PressureHectopascal,
            want: "hPa",
        },
        {
            name: "Kilopascal abbreviation",
            unit: units.PressureKilopascal,
            want: "kPa",
        },
        {
            name: "Megapascal abbreviation",
            unit: units.PressureMegapascal,
            want: "MPa",
        },
        {
            name: "Gigapascal abbreviation",
            unit: units.PressureGigapascal,
            want: "GPa",
        },
        {
            name: "Microbar abbreviation",
            unit: units.PressureMicrobar,
            want: "μbar",
        },
        {
            name: "Millibar abbreviation",
            unit: units.PressureMillibar,
            want: "mbar",
        },
        {
            name: "Centibar abbreviation",
            unit: units.PressureCentibar,
            want: "cbar",
        },
        {
            name: "Decibar abbreviation",
            unit: units.PressureDecibar,
            want: "dbar",
        },
        {
            name: "Kilobar abbreviation",
            unit: units.PressureKilobar,
            want: "kbar",
        },
        {
            name: "Megabar abbreviation",
            unit: units.PressureMegabar,
            want: "Mbar",
        },
        {
            name: "KilonewtonPerSquareMeter abbreviation",
            unit: units.PressureKilonewtonPerSquareMeter,
            want: "kN/m²",
        },
        {
            name: "MeganewtonPerSquareMeter abbreviation",
            unit: units.PressureMeganewtonPerSquareMeter,
            want: "MN/m²",
        },
        {
            name: "KilonewtonPerSquareCentimeter abbreviation",
            unit: units.PressureKilonewtonPerSquareCentimeter,
            want: "kN/cm²",
        },
        {
            name: "KilonewtonPerSquareMillimeter abbreviation",
            unit: units.PressureKilonewtonPerSquareMillimeter,
            want: "kN/mm²",
        },
        {
            name: "KilopoundForcePerSquareInch abbreviation",
            unit: units.PressureKilopoundForcePerSquareInch,
            want: "kpsi",
        },
        {
            name: "KilopoundForcePerSquareMil abbreviation",
            unit: units.PressureKilopoundForcePerSquareMil,
            want: "klb/mil²",
        },
        {
            name: "KilopoundForcePerSquareFoot abbreviation",
            unit: units.PressureKilopoundForcePerSquareFoot,
            want: "klb/ft²",
        },
        {
            name: "MillimeterOfWaterColumn abbreviation",
            unit: units.PressureMillimeterOfWaterColumn,
            want: "mmH₂O",
        },
        {
            name: "CentimeterOfWaterColumn abbreviation",
            unit: units.PressureCentimeterOfWaterColumn,
            want: "cmH₂O",
        },
        {
            name: "invalid unit",
            unit: units.PressureUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetPressureAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetPressureAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestPressure_String(t *testing.T) {
    factory := units.PressureFactory{}
    
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
            unit, err := factory.CreatePressure(tt.value, units.PressurePascal)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Pressure.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestPressure_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.PressureFactory{}

	_, err := uf.CreatePressure(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}