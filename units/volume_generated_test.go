// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestVolumeDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CubicMeter"}`
	
	factory := units.VolumeDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.VolumeCubicMeter {
		t.Errorf("expected unit %v, got %v", units.VolumeCubicMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CubicMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestVolumeDto_ToJSON(t *testing.T) {
	dto := units.VolumeDto{
		Value: 45,
		Unit:  units.VolumeCubicMeter,
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
	if result["unit"].(string) != string(units.VolumeCubicMeter) {
		t.Errorf("expected unit %s, got %v", units.VolumeCubicMeter, result["unit"])
	}
}

func TestNewVolume_InvalidValue(t *testing.T) {
	factory := units.VolumeFactory{}
	// NaN value should return an error.
	_, err := factory.CreateVolume(math.NaN(), units.VolumeCubicMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateVolume(math.Inf(1), units.VolumeCubicMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestVolumeConversions(t *testing.T) {
	factory := units.VolumeFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateVolume(180, units.VolumeCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Liters.
		// No expected conversion value provided for Liters, verifying result is not NaN.
		result := a.Liters()
		cacheResult := a.Liters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Liters returned NaN")
		}
	}
	{
		// Test conversion to CubicMeters.
		// No expected conversion value provided for CubicMeters, verifying result is not NaN.
		result := a.CubicMeters()
		cacheResult := a.CubicMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicMeters returned NaN")
		}
	}
	{
		// Test conversion to CubicKilometers.
		// No expected conversion value provided for CubicKilometers, verifying result is not NaN.
		result := a.CubicKilometers()
		cacheResult := a.CubicKilometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicKilometers returned NaN")
		}
	}
	{
		// Test conversion to CubicHectometers.
		// No expected conversion value provided for CubicHectometers, verifying result is not NaN.
		result := a.CubicHectometers()
		cacheResult := a.CubicHectometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicHectometers returned NaN")
		}
	}
	{
		// Test conversion to CubicDecimeters.
		// No expected conversion value provided for CubicDecimeters, verifying result is not NaN.
		result := a.CubicDecimeters()
		cacheResult := a.CubicDecimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicDecimeters returned NaN")
		}
	}
	{
		// Test conversion to CubicCentimeters.
		// No expected conversion value provided for CubicCentimeters, verifying result is not NaN.
		result := a.CubicCentimeters()
		cacheResult := a.CubicCentimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicCentimeters returned NaN")
		}
	}
	{
		// Test conversion to CubicMillimeters.
		// No expected conversion value provided for CubicMillimeters, verifying result is not NaN.
		result := a.CubicMillimeters()
		cacheResult := a.CubicMillimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicMillimeters returned NaN")
		}
	}
	{
		// Test conversion to CubicMicrometers.
		// No expected conversion value provided for CubicMicrometers, verifying result is not NaN.
		result := a.CubicMicrometers()
		cacheResult := a.CubicMicrometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicMicrometers returned NaN")
		}
	}
	{
		// Test conversion to CubicMiles.
		// No expected conversion value provided for CubicMiles, verifying result is not NaN.
		result := a.CubicMiles()
		cacheResult := a.CubicMiles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicMiles returned NaN")
		}
	}
	{
		// Test conversion to CubicYards.
		// No expected conversion value provided for CubicYards, verifying result is not NaN.
		result := a.CubicYards()
		cacheResult := a.CubicYards()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicYards returned NaN")
		}
	}
	{
		// Test conversion to CubicFeet.
		// No expected conversion value provided for CubicFeet, verifying result is not NaN.
		result := a.CubicFeet()
		cacheResult := a.CubicFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicFeet returned NaN")
		}
	}
	{
		// Test conversion to CubicInches.
		// No expected conversion value provided for CubicInches, verifying result is not NaN.
		result := a.CubicInches()
		cacheResult := a.CubicInches()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicInches returned NaN")
		}
	}
	{
		// Test conversion to ImperialGallons.
		// No expected conversion value provided for ImperialGallons, verifying result is not NaN.
		result := a.ImperialGallons()
		cacheResult := a.ImperialGallons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ImperialGallons returned NaN")
		}
	}
	{
		// Test conversion to ImperialOunces.
		// No expected conversion value provided for ImperialOunces, verifying result is not NaN.
		result := a.ImperialOunces()
		cacheResult := a.ImperialOunces()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ImperialOunces returned NaN")
		}
	}
	{
		// Test conversion to UsGallons.
		// No expected conversion value provided for UsGallons, verifying result is not NaN.
		result := a.UsGallons()
		cacheResult := a.UsGallons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsGallons returned NaN")
		}
	}
	{
		// Test conversion to UsOunces.
		// No expected conversion value provided for UsOunces, verifying result is not NaN.
		result := a.UsOunces()
		cacheResult := a.UsOunces()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsOunces returned NaN")
		}
	}
	{
		// Test conversion to UsTablespoons.
		// No expected conversion value provided for UsTablespoons, verifying result is not NaN.
		result := a.UsTablespoons()
		cacheResult := a.UsTablespoons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsTablespoons returned NaN")
		}
	}
	{
		// Test conversion to AuTablespoons.
		// No expected conversion value provided for AuTablespoons, verifying result is not NaN.
		result := a.AuTablespoons()
		cacheResult := a.AuTablespoons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AuTablespoons returned NaN")
		}
	}
	{
		// Test conversion to UkTablespoons.
		// No expected conversion value provided for UkTablespoons, verifying result is not NaN.
		result := a.UkTablespoons()
		cacheResult := a.UkTablespoons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UkTablespoons returned NaN")
		}
	}
	{
		// Test conversion to MetricTeaspoons.
		// No expected conversion value provided for MetricTeaspoons, verifying result is not NaN.
		result := a.MetricTeaspoons()
		cacheResult := a.MetricTeaspoons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MetricTeaspoons returned NaN")
		}
	}
	{
		// Test conversion to UsTeaspoons.
		// No expected conversion value provided for UsTeaspoons, verifying result is not NaN.
		result := a.UsTeaspoons()
		cacheResult := a.UsTeaspoons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsTeaspoons returned NaN")
		}
	}
	{
		// Test conversion to MetricCups.
		// No expected conversion value provided for MetricCups, verifying result is not NaN.
		result := a.MetricCups()
		cacheResult := a.MetricCups()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MetricCups returned NaN")
		}
	}
	{
		// Test conversion to UsCustomaryCups.
		// No expected conversion value provided for UsCustomaryCups, verifying result is not NaN.
		result := a.UsCustomaryCups()
		cacheResult := a.UsCustomaryCups()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsCustomaryCups returned NaN")
		}
	}
	{
		// Test conversion to UsLegalCups.
		// No expected conversion value provided for UsLegalCups, verifying result is not NaN.
		result := a.UsLegalCups()
		cacheResult := a.UsLegalCups()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsLegalCups returned NaN")
		}
	}
	{
		// Test conversion to OilBarrels.
		// No expected conversion value provided for OilBarrels, verifying result is not NaN.
		result := a.OilBarrels()
		cacheResult := a.OilBarrels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to OilBarrels returned NaN")
		}
	}
	{
		// Test conversion to UsBeerBarrels.
		// No expected conversion value provided for UsBeerBarrels, verifying result is not NaN.
		result := a.UsBeerBarrels()
		cacheResult := a.UsBeerBarrels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsBeerBarrels returned NaN")
		}
	}
	{
		// Test conversion to ImperialBeerBarrels.
		// No expected conversion value provided for ImperialBeerBarrels, verifying result is not NaN.
		result := a.ImperialBeerBarrels()
		cacheResult := a.ImperialBeerBarrels()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ImperialBeerBarrels returned NaN")
		}
	}
	{
		// Test conversion to UsQuarts.
		// No expected conversion value provided for UsQuarts, verifying result is not NaN.
		result := a.UsQuarts()
		cacheResult := a.UsQuarts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsQuarts returned NaN")
		}
	}
	{
		// Test conversion to ImperialQuarts.
		// No expected conversion value provided for ImperialQuarts, verifying result is not NaN.
		result := a.ImperialQuarts()
		cacheResult := a.ImperialQuarts()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ImperialQuarts returned NaN")
		}
	}
	{
		// Test conversion to UsPints.
		// No expected conversion value provided for UsPints, verifying result is not NaN.
		result := a.UsPints()
		cacheResult := a.UsPints()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsPints returned NaN")
		}
	}
	{
		// Test conversion to AcreFeet.
		// No expected conversion value provided for AcreFeet, verifying result is not NaN.
		result := a.AcreFeet()
		cacheResult := a.AcreFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AcreFeet returned NaN")
		}
	}
	{
		// Test conversion to ImperialPints.
		// No expected conversion value provided for ImperialPints, verifying result is not NaN.
		result := a.ImperialPints()
		cacheResult := a.ImperialPints()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ImperialPints returned NaN")
		}
	}
	{
		// Test conversion to BoardFeet.
		// No expected conversion value provided for BoardFeet, verifying result is not NaN.
		result := a.BoardFeet()
		cacheResult := a.BoardFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to BoardFeet returned NaN")
		}
	}
	{
		// Test conversion to Nanoliters.
		// No expected conversion value provided for Nanoliters, verifying result is not NaN.
		result := a.Nanoliters()
		cacheResult := a.Nanoliters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanoliters returned NaN")
		}
	}
	{
		// Test conversion to Microliters.
		// No expected conversion value provided for Microliters, verifying result is not NaN.
		result := a.Microliters()
		cacheResult := a.Microliters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microliters returned NaN")
		}
	}
	{
		// Test conversion to Milliliters.
		// No expected conversion value provided for Milliliters, verifying result is not NaN.
		result := a.Milliliters()
		cacheResult := a.Milliliters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Milliliters returned NaN")
		}
	}
	{
		// Test conversion to Centiliters.
		// No expected conversion value provided for Centiliters, verifying result is not NaN.
		result := a.Centiliters()
		cacheResult := a.Centiliters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Centiliters returned NaN")
		}
	}
	{
		// Test conversion to Deciliters.
		// No expected conversion value provided for Deciliters, verifying result is not NaN.
		result := a.Deciliters()
		cacheResult := a.Deciliters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Deciliters returned NaN")
		}
	}
	{
		// Test conversion to Decaliters.
		// No expected conversion value provided for Decaliters, verifying result is not NaN.
		result := a.Decaliters()
		cacheResult := a.Decaliters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Decaliters returned NaN")
		}
	}
	{
		// Test conversion to Hectoliters.
		// No expected conversion value provided for Hectoliters, verifying result is not NaN.
		result := a.Hectoliters()
		cacheResult := a.Hectoliters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Hectoliters returned NaN")
		}
	}
	{
		// Test conversion to Kiloliters.
		// No expected conversion value provided for Kiloliters, verifying result is not NaN.
		result := a.Kiloliters()
		cacheResult := a.Kiloliters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kiloliters returned NaN")
		}
	}
	{
		// Test conversion to Megaliters.
		// No expected conversion value provided for Megaliters, verifying result is not NaN.
		result := a.Megaliters()
		cacheResult := a.Megaliters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megaliters returned NaN")
		}
	}
	{
		// Test conversion to HectocubicMeters.
		// No expected conversion value provided for HectocubicMeters, verifying result is not NaN.
		result := a.HectocubicMeters()
		cacheResult := a.HectocubicMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to HectocubicMeters returned NaN")
		}
	}
	{
		// Test conversion to KilocubicMeters.
		// No expected conversion value provided for KilocubicMeters, verifying result is not NaN.
		result := a.KilocubicMeters()
		cacheResult := a.KilocubicMeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilocubicMeters returned NaN")
		}
	}
	{
		// Test conversion to HectocubicFeet.
		// No expected conversion value provided for HectocubicFeet, verifying result is not NaN.
		result := a.HectocubicFeet()
		cacheResult := a.HectocubicFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to HectocubicFeet returned NaN")
		}
	}
	{
		// Test conversion to KilocubicFeet.
		// No expected conversion value provided for KilocubicFeet, verifying result is not NaN.
		result := a.KilocubicFeet()
		cacheResult := a.KilocubicFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilocubicFeet returned NaN")
		}
	}
	{
		// Test conversion to MegacubicFeet.
		// No expected conversion value provided for MegacubicFeet, verifying result is not NaN.
		result := a.MegacubicFeet()
		cacheResult := a.MegacubicFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegacubicFeet returned NaN")
		}
	}
	{
		// Test conversion to KiloimperialGallons.
		// No expected conversion value provided for KiloimperialGallons, verifying result is not NaN.
		result := a.KiloimperialGallons()
		cacheResult := a.KiloimperialGallons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KiloimperialGallons returned NaN")
		}
	}
	{
		// Test conversion to MegaimperialGallons.
		// No expected conversion value provided for MegaimperialGallons, verifying result is not NaN.
		result := a.MegaimperialGallons()
		cacheResult := a.MegaimperialGallons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegaimperialGallons returned NaN")
		}
	}
	{
		// Test conversion to DecausGallons.
		// No expected conversion value provided for DecausGallons, verifying result is not NaN.
		result := a.DecausGallons()
		cacheResult := a.DecausGallons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecausGallons returned NaN")
		}
	}
	{
		// Test conversion to DeciusGallons.
		// No expected conversion value provided for DeciusGallons, verifying result is not NaN.
		result := a.DeciusGallons()
		cacheResult := a.DeciusGallons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DeciusGallons returned NaN")
		}
	}
	{
		// Test conversion to HectousGallons.
		// No expected conversion value provided for HectousGallons, verifying result is not NaN.
		result := a.HectousGallons()
		cacheResult := a.HectousGallons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to HectousGallons returned NaN")
		}
	}
	{
		// Test conversion to KilousGallons.
		// No expected conversion value provided for KilousGallons, verifying result is not NaN.
		result := a.KilousGallons()
		cacheResult := a.KilousGallons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilousGallons returned NaN")
		}
	}
	{
		// Test conversion to MegausGallons.
		// No expected conversion value provided for MegausGallons, verifying result is not NaN.
		result := a.MegausGallons()
		cacheResult := a.MegausGallons()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegausGallons returned NaN")
		}
	}
}

func TestVolume_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.VolumeFactory{}
	a, err := factory.CreateVolume(90, units.VolumeCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.VolumeCubicMeter {
		t.Errorf("expected default unit CubicMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.VolumeLiter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.VolumeDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.VolumeCubicMeter {
		t.Errorf("expected unit CubicMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestVolumeFactory_FromDto(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeCubicMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.VolumeDto{
        Value: math.NaN(),
        Unit:  units.VolumeCubicMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Liter conversion
    litersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeLiter,
    }
    
    var litersResult *units.Volume
    litersResult, err = factory.FromDto(litersDto)
    if err != nil {
        t.Errorf("FromDto() with Liter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = litersResult.Convert(units.VolumeLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Liter = %v, want %v", converted, 100)
    }
    // Test CubicMeter conversion
    cubic_metersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeCubicMeter,
    }
    
    var cubic_metersResult *units.Volume
    cubic_metersResult, err = factory.FromDto(cubic_metersDto)
    if err != nil {
        t.Errorf("FromDto() with CubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_metersResult.Convert(units.VolumeCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeter = %v, want %v", converted, 100)
    }
    // Test CubicKilometer conversion
    cubic_kilometersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeCubicKilometer,
    }
    
    var cubic_kilometersResult *units.Volume
    cubic_kilometersResult, err = factory.FromDto(cubic_kilometersDto)
    if err != nil {
        t.Errorf("FromDto() with CubicKilometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_kilometersResult.Convert(units.VolumeCubicKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicKilometer = %v, want %v", converted, 100)
    }
    // Test CubicHectometer conversion
    cubic_hectometersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeCubicHectometer,
    }
    
    var cubic_hectometersResult *units.Volume
    cubic_hectometersResult, err = factory.FromDto(cubic_hectometersDto)
    if err != nil {
        t.Errorf("FromDto() with CubicHectometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_hectometersResult.Convert(units.VolumeCubicHectometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicHectometer = %v, want %v", converted, 100)
    }
    // Test CubicDecimeter conversion
    cubic_decimetersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeCubicDecimeter,
    }
    
    var cubic_decimetersResult *units.Volume
    cubic_decimetersResult, err = factory.FromDto(cubic_decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with CubicDecimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_decimetersResult.Convert(units.VolumeCubicDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicDecimeter = %v, want %v", converted, 100)
    }
    // Test CubicCentimeter conversion
    cubic_centimetersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeCubicCentimeter,
    }
    
    var cubic_centimetersResult *units.Volume
    cubic_centimetersResult, err = factory.FromDto(cubic_centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with CubicCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_centimetersResult.Convert(units.VolumeCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicCentimeter = %v, want %v", converted, 100)
    }
    // Test CubicMillimeter conversion
    cubic_millimetersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeCubicMillimeter,
    }
    
    var cubic_millimetersResult *units.Volume
    cubic_millimetersResult, err = factory.FromDto(cubic_millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with CubicMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_millimetersResult.Convert(units.VolumeCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMillimeter = %v, want %v", converted, 100)
    }
    // Test CubicMicrometer conversion
    cubic_micrometersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeCubicMicrometer,
    }
    
    var cubic_micrometersResult *units.Volume
    cubic_micrometersResult, err = factory.FromDto(cubic_micrometersDto)
    if err != nil {
        t.Errorf("FromDto() with CubicMicrometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_micrometersResult.Convert(units.VolumeCubicMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMicrometer = %v, want %v", converted, 100)
    }
    // Test CubicMile conversion
    cubic_milesDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeCubicMile,
    }
    
    var cubic_milesResult *units.Volume
    cubic_milesResult, err = factory.FromDto(cubic_milesDto)
    if err != nil {
        t.Errorf("FromDto() with CubicMile returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_milesResult.Convert(units.VolumeCubicMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMile = %v, want %v", converted, 100)
    }
    // Test CubicYard conversion
    cubic_yardsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeCubicYard,
    }
    
    var cubic_yardsResult *units.Volume
    cubic_yardsResult, err = factory.FromDto(cubic_yardsDto)
    if err != nil {
        t.Errorf("FromDto() with CubicYard returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yardsResult.Convert(units.VolumeCubicYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYard = %v, want %v", converted, 100)
    }
    // Test CubicFoot conversion
    cubic_feetDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeCubicFoot,
    }
    
    var cubic_feetResult *units.Volume
    cubic_feetResult, err = factory.FromDto(cubic_feetDto)
    if err != nil {
        t.Errorf("FromDto() with CubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_feetResult.Convert(units.VolumeCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicFoot = %v, want %v", converted, 100)
    }
    // Test CubicInch conversion
    cubic_inchesDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeCubicInch,
    }
    
    var cubic_inchesResult *units.Volume
    cubic_inchesResult, err = factory.FromDto(cubic_inchesDto)
    if err != nil {
        t.Errorf("FromDto() with CubicInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_inchesResult.Convert(units.VolumeCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicInch = %v, want %v", converted, 100)
    }
    // Test ImperialGallon conversion
    imperial_gallonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeImperialGallon,
    }
    
    var imperial_gallonsResult *units.Volume
    imperial_gallonsResult, err = factory.FromDto(imperial_gallonsDto)
    if err != nil {
        t.Errorf("FromDto() with ImperialGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = imperial_gallonsResult.Convert(units.VolumeImperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ImperialGallon = %v, want %v", converted, 100)
    }
    // Test ImperialOunce conversion
    imperial_ouncesDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeImperialOunce,
    }
    
    var imperial_ouncesResult *units.Volume
    imperial_ouncesResult, err = factory.FromDto(imperial_ouncesDto)
    if err != nil {
        t.Errorf("FromDto() with ImperialOunce returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = imperial_ouncesResult.Convert(units.VolumeImperialOunce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ImperialOunce = %v, want %v", converted, 100)
    }
    // Test UsGallon conversion
    us_gallonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeUsGallon,
    }
    
    var us_gallonsResult *units.Volume
    us_gallonsResult, err = factory.FromDto(us_gallonsDto)
    if err != nil {
        t.Errorf("FromDto() with UsGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_gallonsResult.Convert(units.VolumeUsGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsGallon = %v, want %v", converted, 100)
    }
    // Test UsOunce conversion
    us_ouncesDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeUsOunce,
    }
    
    var us_ouncesResult *units.Volume
    us_ouncesResult, err = factory.FromDto(us_ouncesDto)
    if err != nil {
        t.Errorf("FromDto() with UsOunce returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_ouncesResult.Convert(units.VolumeUsOunce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsOunce = %v, want %v", converted, 100)
    }
    // Test UsTablespoon conversion
    us_tablespoonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeUsTablespoon,
    }
    
    var us_tablespoonsResult *units.Volume
    us_tablespoonsResult, err = factory.FromDto(us_tablespoonsDto)
    if err != nil {
        t.Errorf("FromDto() with UsTablespoon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_tablespoonsResult.Convert(units.VolumeUsTablespoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsTablespoon = %v, want %v", converted, 100)
    }
    // Test AuTablespoon conversion
    au_tablespoonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeAuTablespoon,
    }
    
    var au_tablespoonsResult *units.Volume
    au_tablespoonsResult, err = factory.FromDto(au_tablespoonsDto)
    if err != nil {
        t.Errorf("FromDto() with AuTablespoon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = au_tablespoonsResult.Convert(units.VolumeAuTablespoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AuTablespoon = %v, want %v", converted, 100)
    }
    // Test UkTablespoon conversion
    uk_tablespoonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeUkTablespoon,
    }
    
    var uk_tablespoonsResult *units.Volume
    uk_tablespoonsResult, err = factory.FromDto(uk_tablespoonsDto)
    if err != nil {
        t.Errorf("FromDto() with UkTablespoon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = uk_tablespoonsResult.Convert(units.VolumeUkTablespoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UkTablespoon = %v, want %v", converted, 100)
    }
    // Test MetricTeaspoon conversion
    metric_teaspoonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeMetricTeaspoon,
    }
    
    var metric_teaspoonsResult *units.Volume
    metric_teaspoonsResult, err = factory.FromDto(metric_teaspoonsDto)
    if err != nil {
        t.Errorf("FromDto() with MetricTeaspoon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = metric_teaspoonsResult.Convert(units.VolumeMetricTeaspoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MetricTeaspoon = %v, want %v", converted, 100)
    }
    // Test UsTeaspoon conversion
    us_teaspoonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeUsTeaspoon,
    }
    
    var us_teaspoonsResult *units.Volume
    us_teaspoonsResult, err = factory.FromDto(us_teaspoonsDto)
    if err != nil {
        t.Errorf("FromDto() with UsTeaspoon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_teaspoonsResult.Convert(units.VolumeUsTeaspoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsTeaspoon = %v, want %v", converted, 100)
    }
    // Test MetricCup conversion
    metric_cupsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeMetricCup,
    }
    
    var metric_cupsResult *units.Volume
    metric_cupsResult, err = factory.FromDto(metric_cupsDto)
    if err != nil {
        t.Errorf("FromDto() with MetricCup returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = metric_cupsResult.Convert(units.VolumeMetricCup)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MetricCup = %v, want %v", converted, 100)
    }
    // Test UsCustomaryCup conversion
    us_customary_cupsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeUsCustomaryCup,
    }
    
    var us_customary_cupsResult *units.Volume
    us_customary_cupsResult, err = factory.FromDto(us_customary_cupsDto)
    if err != nil {
        t.Errorf("FromDto() with UsCustomaryCup returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_customary_cupsResult.Convert(units.VolumeUsCustomaryCup)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsCustomaryCup = %v, want %v", converted, 100)
    }
    // Test UsLegalCup conversion
    us_legal_cupsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeUsLegalCup,
    }
    
    var us_legal_cupsResult *units.Volume
    us_legal_cupsResult, err = factory.FromDto(us_legal_cupsDto)
    if err != nil {
        t.Errorf("FromDto() with UsLegalCup returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_legal_cupsResult.Convert(units.VolumeUsLegalCup)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsLegalCup = %v, want %v", converted, 100)
    }
    // Test OilBarrel conversion
    oil_barrelsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeOilBarrel,
    }
    
    var oil_barrelsResult *units.Volume
    oil_barrelsResult, err = factory.FromDto(oil_barrelsDto)
    if err != nil {
        t.Errorf("FromDto() with OilBarrel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = oil_barrelsResult.Convert(units.VolumeOilBarrel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OilBarrel = %v, want %v", converted, 100)
    }
    // Test UsBeerBarrel conversion
    us_beer_barrelsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeUsBeerBarrel,
    }
    
    var us_beer_barrelsResult *units.Volume
    us_beer_barrelsResult, err = factory.FromDto(us_beer_barrelsDto)
    if err != nil {
        t.Errorf("FromDto() with UsBeerBarrel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_beer_barrelsResult.Convert(units.VolumeUsBeerBarrel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsBeerBarrel = %v, want %v", converted, 100)
    }
    // Test ImperialBeerBarrel conversion
    imperial_beer_barrelsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeImperialBeerBarrel,
    }
    
    var imperial_beer_barrelsResult *units.Volume
    imperial_beer_barrelsResult, err = factory.FromDto(imperial_beer_barrelsDto)
    if err != nil {
        t.Errorf("FromDto() with ImperialBeerBarrel returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = imperial_beer_barrelsResult.Convert(units.VolumeImperialBeerBarrel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ImperialBeerBarrel = %v, want %v", converted, 100)
    }
    // Test UsQuart conversion
    us_quartsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeUsQuart,
    }
    
    var us_quartsResult *units.Volume
    us_quartsResult, err = factory.FromDto(us_quartsDto)
    if err != nil {
        t.Errorf("FromDto() with UsQuart returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_quartsResult.Convert(units.VolumeUsQuart)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsQuart = %v, want %v", converted, 100)
    }
    // Test ImperialQuart conversion
    imperial_quartsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeImperialQuart,
    }
    
    var imperial_quartsResult *units.Volume
    imperial_quartsResult, err = factory.FromDto(imperial_quartsDto)
    if err != nil {
        t.Errorf("FromDto() with ImperialQuart returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = imperial_quartsResult.Convert(units.VolumeImperialQuart)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ImperialQuart = %v, want %v", converted, 100)
    }
    // Test UsPint conversion
    us_pintsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeUsPint,
    }
    
    var us_pintsResult *units.Volume
    us_pintsResult, err = factory.FromDto(us_pintsDto)
    if err != nil {
        t.Errorf("FromDto() with UsPint returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_pintsResult.Convert(units.VolumeUsPint)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsPint = %v, want %v", converted, 100)
    }
    // Test AcreFoot conversion
    acre_feetDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeAcreFoot,
    }
    
    var acre_feetResult *units.Volume
    acre_feetResult, err = factory.FromDto(acre_feetDto)
    if err != nil {
        t.Errorf("FromDto() with AcreFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = acre_feetResult.Convert(units.VolumeAcreFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AcreFoot = %v, want %v", converted, 100)
    }
    // Test ImperialPint conversion
    imperial_pintsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeImperialPint,
    }
    
    var imperial_pintsResult *units.Volume
    imperial_pintsResult, err = factory.FromDto(imperial_pintsDto)
    if err != nil {
        t.Errorf("FromDto() with ImperialPint returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = imperial_pintsResult.Convert(units.VolumeImperialPint)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ImperialPint = %v, want %v", converted, 100)
    }
    // Test BoardFoot conversion
    board_feetDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeBoardFoot,
    }
    
    var board_feetResult *units.Volume
    board_feetResult, err = factory.FromDto(board_feetDto)
    if err != nil {
        t.Errorf("FromDto() with BoardFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = board_feetResult.Convert(units.VolumeBoardFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BoardFoot = %v, want %v", converted, 100)
    }
    // Test Nanoliter conversion
    nanolitersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeNanoliter,
    }
    
    var nanolitersResult *units.Volume
    nanolitersResult, err = factory.FromDto(nanolitersDto)
    if err != nil {
        t.Errorf("FromDto() with Nanoliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanolitersResult.Convert(units.VolumeNanoliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanoliter = %v, want %v", converted, 100)
    }
    // Test Microliter conversion
    microlitersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeMicroliter,
    }
    
    var microlitersResult *units.Volume
    microlitersResult, err = factory.FromDto(microlitersDto)
    if err != nil {
        t.Errorf("FromDto() with Microliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microlitersResult.Convert(units.VolumeMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microliter = %v, want %v", converted, 100)
    }
    // Test Milliliter conversion
    millilitersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeMilliliter,
    }
    
    var millilitersResult *units.Volume
    millilitersResult, err = factory.FromDto(millilitersDto)
    if err != nil {
        t.Errorf("FromDto() with Milliliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millilitersResult.Convert(units.VolumeMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliliter = %v, want %v", converted, 100)
    }
    // Test Centiliter conversion
    centilitersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeCentiliter,
    }
    
    var centilitersResult *units.Volume
    centilitersResult, err = factory.FromDto(centilitersDto)
    if err != nil {
        t.Errorf("FromDto() with Centiliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centilitersResult.Convert(units.VolumeCentiliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centiliter = %v, want %v", converted, 100)
    }
    // Test Deciliter conversion
    decilitersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeDeciliter,
    }
    
    var decilitersResult *units.Volume
    decilitersResult, err = factory.FromDto(decilitersDto)
    if err != nil {
        t.Errorf("FromDto() with Deciliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decilitersResult.Convert(units.VolumeDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Deciliter = %v, want %v", converted, 100)
    }
    // Test Decaliter conversion
    decalitersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeDecaliter,
    }
    
    var decalitersResult *units.Volume
    decalitersResult, err = factory.FromDto(decalitersDto)
    if err != nil {
        t.Errorf("FromDto() with Decaliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decalitersResult.Convert(units.VolumeDecaliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decaliter = %v, want %v", converted, 100)
    }
    // Test Hectoliter conversion
    hectolitersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeHectoliter,
    }
    
    var hectolitersResult *units.Volume
    hectolitersResult, err = factory.FromDto(hectolitersDto)
    if err != nil {
        t.Errorf("FromDto() with Hectoliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectolitersResult.Convert(units.VolumeHectoliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hectoliter = %v, want %v", converted, 100)
    }
    // Test Kiloliter conversion
    kilolitersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeKiloliter,
    }
    
    var kilolitersResult *units.Volume
    kilolitersResult, err = factory.FromDto(kilolitersDto)
    if err != nil {
        t.Errorf("FromDto() with Kiloliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilolitersResult.Convert(units.VolumeKiloliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kiloliter = %v, want %v", converted, 100)
    }
    // Test Megaliter conversion
    megalitersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeMegaliter,
    }
    
    var megalitersResult *units.Volume
    megalitersResult, err = factory.FromDto(megalitersDto)
    if err != nil {
        t.Errorf("FromDto() with Megaliter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megalitersResult.Convert(units.VolumeMegaliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megaliter = %v, want %v", converted, 100)
    }
    // Test HectocubicMeter conversion
    hectocubic_metersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeHectocubicMeter,
    }
    
    var hectocubic_metersResult *units.Volume
    hectocubic_metersResult, err = factory.FromDto(hectocubic_metersDto)
    if err != nil {
        t.Errorf("FromDto() with HectocubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectocubic_metersResult.Convert(units.VolumeHectocubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectocubicMeter = %v, want %v", converted, 100)
    }
    // Test KilocubicMeter conversion
    kilocubic_metersDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeKilocubicMeter,
    }
    
    var kilocubic_metersResult *units.Volume
    kilocubic_metersResult, err = factory.FromDto(kilocubic_metersDto)
    if err != nil {
        t.Errorf("FromDto() with KilocubicMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocubic_metersResult.Convert(units.VolumeKilocubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocubicMeter = %v, want %v", converted, 100)
    }
    // Test HectocubicFoot conversion
    hectocubic_feetDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeHectocubicFoot,
    }
    
    var hectocubic_feetResult *units.Volume
    hectocubic_feetResult, err = factory.FromDto(hectocubic_feetDto)
    if err != nil {
        t.Errorf("FromDto() with HectocubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectocubic_feetResult.Convert(units.VolumeHectocubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectocubicFoot = %v, want %v", converted, 100)
    }
    // Test KilocubicFoot conversion
    kilocubic_feetDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeKilocubicFoot,
    }
    
    var kilocubic_feetResult *units.Volume
    kilocubic_feetResult, err = factory.FromDto(kilocubic_feetDto)
    if err != nil {
        t.Errorf("FromDto() with KilocubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocubic_feetResult.Convert(units.VolumeKilocubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocubicFoot = %v, want %v", converted, 100)
    }
    // Test MegacubicFoot conversion
    megacubic_feetDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeMegacubicFoot,
    }
    
    var megacubic_feetResult *units.Volume
    megacubic_feetResult, err = factory.FromDto(megacubic_feetDto)
    if err != nil {
        t.Errorf("FromDto() with MegacubicFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megacubic_feetResult.Convert(units.VolumeMegacubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegacubicFoot = %v, want %v", converted, 100)
    }
    // Test KiloimperialGallon conversion
    kiloimperial_gallonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeKiloimperialGallon,
    }
    
    var kiloimperial_gallonsResult *units.Volume
    kiloimperial_gallonsResult, err = factory.FromDto(kiloimperial_gallonsDto)
    if err != nil {
        t.Errorf("FromDto() with KiloimperialGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloimperial_gallonsResult.Convert(units.VolumeKiloimperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloimperialGallon = %v, want %v", converted, 100)
    }
    // Test MegaimperialGallon conversion
    megaimperial_gallonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeMegaimperialGallon,
    }
    
    var megaimperial_gallonsResult *units.Volume
    megaimperial_gallonsResult, err = factory.FromDto(megaimperial_gallonsDto)
    if err != nil {
        t.Errorf("FromDto() with MegaimperialGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaimperial_gallonsResult.Convert(units.VolumeMegaimperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaimperialGallon = %v, want %v", converted, 100)
    }
    // Test DecausGallon conversion
    decaus_gallonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeDecausGallon,
    }
    
    var decaus_gallonsResult *units.Volume
    decaus_gallonsResult, err = factory.FromDto(decaus_gallonsDto)
    if err != nil {
        t.Errorf("FromDto() with DecausGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decaus_gallonsResult.Convert(units.VolumeDecausGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecausGallon = %v, want %v", converted, 100)
    }
    // Test DeciusGallon conversion
    decius_gallonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeDeciusGallon,
    }
    
    var decius_gallonsResult *units.Volume
    decius_gallonsResult, err = factory.FromDto(decius_gallonsDto)
    if err != nil {
        t.Errorf("FromDto() with DeciusGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decius_gallonsResult.Convert(units.VolumeDeciusGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciusGallon = %v, want %v", converted, 100)
    }
    // Test HectousGallon conversion
    hectous_gallonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeHectousGallon,
    }
    
    var hectous_gallonsResult *units.Volume
    hectous_gallonsResult, err = factory.FromDto(hectous_gallonsDto)
    if err != nil {
        t.Errorf("FromDto() with HectousGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectous_gallonsResult.Convert(units.VolumeHectousGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectousGallon = %v, want %v", converted, 100)
    }
    // Test KilousGallon conversion
    kilous_gallonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeKilousGallon,
    }
    
    var kilous_gallonsResult *units.Volume
    kilous_gallonsResult, err = factory.FromDto(kilous_gallonsDto)
    if err != nil {
        t.Errorf("FromDto() with KilousGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilous_gallonsResult.Convert(units.VolumeKilousGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilousGallon = %v, want %v", converted, 100)
    }
    // Test MegausGallon conversion
    megaus_gallonsDto := units.VolumeDto{
        Value: 100,
        Unit:  units.VolumeMegausGallon,
    }
    
    var megaus_gallonsResult *units.Volume
    megaus_gallonsResult, err = factory.FromDto(megaus_gallonsDto)
    if err != nil {
        t.Errorf("FromDto() with MegausGallon returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaus_gallonsResult.Convert(units.VolumeMegausGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegausGallon = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.VolumeDto{
        Value: 0,
        Unit:  units.VolumeCubicMeter,
    }
    
    var zeroResult *units.Volume
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestVolumeFactory_FromDtoJSON(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "CubicMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "CubicMeter"}`)
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
    nanJSON, _ := json.Marshal(units.VolumeDto{
        Value: nanValue,
        Unit:  units.VolumeCubicMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Liter unit
    litersJSON := []byte(`{"value": 100, "unit": "Liter"}`)
    litersResult, err := factory.FromDtoJSON(litersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Liter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = litersResult.Convert(units.VolumeLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Liter = %v, want %v", converted, 100)
    }
    // Test JSON with CubicMeter unit
    cubic_metersJSON := []byte(`{"value": 100, "unit": "CubicMeter"}`)
    cubic_metersResult, err := factory.FromDtoJSON(cubic_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_metersResult.Convert(units.VolumeCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with CubicKilometer unit
    cubic_kilometersJSON := []byte(`{"value": 100, "unit": "CubicKilometer"}`)
    cubic_kilometersResult, err := factory.FromDtoJSON(cubic_kilometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicKilometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_kilometersResult.Convert(units.VolumeCubicKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicKilometer = %v, want %v", converted, 100)
    }
    // Test JSON with CubicHectometer unit
    cubic_hectometersJSON := []byte(`{"value": 100, "unit": "CubicHectometer"}`)
    cubic_hectometersResult, err := factory.FromDtoJSON(cubic_hectometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicHectometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_hectometersResult.Convert(units.VolumeCubicHectometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicHectometer = %v, want %v", converted, 100)
    }
    // Test JSON with CubicDecimeter unit
    cubic_decimetersJSON := []byte(`{"value": 100, "unit": "CubicDecimeter"}`)
    cubic_decimetersResult, err := factory.FromDtoJSON(cubic_decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicDecimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_decimetersResult.Convert(units.VolumeCubicDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicDecimeter = %v, want %v", converted, 100)
    }
    // Test JSON with CubicCentimeter unit
    cubic_centimetersJSON := []byte(`{"value": 100, "unit": "CubicCentimeter"}`)
    cubic_centimetersResult, err := factory.FromDtoJSON(cubic_centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_centimetersResult.Convert(units.VolumeCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with CubicMillimeter unit
    cubic_millimetersJSON := []byte(`{"value": 100, "unit": "CubicMillimeter"}`)
    cubic_millimetersResult, err := factory.FromDtoJSON(cubic_millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_millimetersResult.Convert(units.VolumeCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with CubicMicrometer unit
    cubic_micrometersJSON := []byte(`{"value": 100, "unit": "CubicMicrometer"}`)
    cubic_micrometersResult, err := factory.FromDtoJSON(cubic_micrometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicMicrometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_micrometersResult.Convert(units.VolumeCubicMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMicrometer = %v, want %v", converted, 100)
    }
    // Test JSON with CubicMile unit
    cubic_milesJSON := []byte(`{"value": 100, "unit": "CubicMile"}`)
    cubic_milesResult, err := factory.FromDtoJSON(cubic_milesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicMile unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_milesResult.Convert(units.VolumeCubicMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMile = %v, want %v", converted, 100)
    }
    // Test JSON with CubicYard unit
    cubic_yardsJSON := []byte(`{"value": 100, "unit": "CubicYard"}`)
    cubic_yardsResult, err := factory.FromDtoJSON(cubic_yardsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicYard unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yardsResult.Convert(units.VolumeCubicYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYard = %v, want %v", converted, 100)
    }
    // Test JSON with CubicFoot unit
    cubic_feetJSON := []byte(`{"value": 100, "unit": "CubicFoot"}`)
    cubic_feetResult, err := factory.FromDtoJSON(cubic_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_feetResult.Convert(units.VolumeCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with CubicInch unit
    cubic_inchesJSON := []byte(`{"value": 100, "unit": "CubicInch"}`)
    cubic_inchesResult, err := factory.FromDtoJSON(cubic_inchesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_inchesResult.Convert(units.VolumeCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicInch = %v, want %v", converted, 100)
    }
    // Test JSON with ImperialGallon unit
    imperial_gallonsJSON := []byte(`{"value": 100, "unit": "ImperialGallon"}`)
    imperial_gallonsResult, err := factory.FromDtoJSON(imperial_gallonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ImperialGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = imperial_gallonsResult.Convert(units.VolumeImperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ImperialGallon = %v, want %v", converted, 100)
    }
    // Test JSON with ImperialOunce unit
    imperial_ouncesJSON := []byte(`{"value": 100, "unit": "ImperialOunce"}`)
    imperial_ouncesResult, err := factory.FromDtoJSON(imperial_ouncesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ImperialOunce unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = imperial_ouncesResult.Convert(units.VolumeImperialOunce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ImperialOunce = %v, want %v", converted, 100)
    }
    // Test JSON with UsGallon unit
    us_gallonsJSON := []byte(`{"value": 100, "unit": "UsGallon"}`)
    us_gallonsResult, err := factory.FromDtoJSON(us_gallonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_gallonsResult.Convert(units.VolumeUsGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsGallon = %v, want %v", converted, 100)
    }
    // Test JSON with UsOunce unit
    us_ouncesJSON := []byte(`{"value": 100, "unit": "UsOunce"}`)
    us_ouncesResult, err := factory.FromDtoJSON(us_ouncesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsOunce unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_ouncesResult.Convert(units.VolumeUsOunce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsOunce = %v, want %v", converted, 100)
    }
    // Test JSON with UsTablespoon unit
    us_tablespoonsJSON := []byte(`{"value": 100, "unit": "UsTablespoon"}`)
    us_tablespoonsResult, err := factory.FromDtoJSON(us_tablespoonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsTablespoon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_tablespoonsResult.Convert(units.VolumeUsTablespoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsTablespoon = %v, want %v", converted, 100)
    }
    // Test JSON with AuTablespoon unit
    au_tablespoonsJSON := []byte(`{"value": 100, "unit": "AuTablespoon"}`)
    au_tablespoonsResult, err := factory.FromDtoJSON(au_tablespoonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AuTablespoon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = au_tablespoonsResult.Convert(units.VolumeAuTablespoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AuTablespoon = %v, want %v", converted, 100)
    }
    // Test JSON with UkTablespoon unit
    uk_tablespoonsJSON := []byte(`{"value": 100, "unit": "UkTablespoon"}`)
    uk_tablespoonsResult, err := factory.FromDtoJSON(uk_tablespoonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UkTablespoon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = uk_tablespoonsResult.Convert(units.VolumeUkTablespoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UkTablespoon = %v, want %v", converted, 100)
    }
    // Test JSON with MetricTeaspoon unit
    metric_teaspoonsJSON := []byte(`{"value": 100, "unit": "MetricTeaspoon"}`)
    metric_teaspoonsResult, err := factory.FromDtoJSON(metric_teaspoonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MetricTeaspoon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = metric_teaspoonsResult.Convert(units.VolumeMetricTeaspoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MetricTeaspoon = %v, want %v", converted, 100)
    }
    // Test JSON with UsTeaspoon unit
    us_teaspoonsJSON := []byte(`{"value": 100, "unit": "UsTeaspoon"}`)
    us_teaspoonsResult, err := factory.FromDtoJSON(us_teaspoonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsTeaspoon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_teaspoonsResult.Convert(units.VolumeUsTeaspoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsTeaspoon = %v, want %v", converted, 100)
    }
    // Test JSON with MetricCup unit
    metric_cupsJSON := []byte(`{"value": 100, "unit": "MetricCup"}`)
    metric_cupsResult, err := factory.FromDtoJSON(metric_cupsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MetricCup unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = metric_cupsResult.Convert(units.VolumeMetricCup)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MetricCup = %v, want %v", converted, 100)
    }
    // Test JSON with UsCustomaryCup unit
    us_customary_cupsJSON := []byte(`{"value": 100, "unit": "UsCustomaryCup"}`)
    us_customary_cupsResult, err := factory.FromDtoJSON(us_customary_cupsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsCustomaryCup unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_customary_cupsResult.Convert(units.VolumeUsCustomaryCup)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsCustomaryCup = %v, want %v", converted, 100)
    }
    // Test JSON with UsLegalCup unit
    us_legal_cupsJSON := []byte(`{"value": 100, "unit": "UsLegalCup"}`)
    us_legal_cupsResult, err := factory.FromDtoJSON(us_legal_cupsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsLegalCup unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_legal_cupsResult.Convert(units.VolumeUsLegalCup)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsLegalCup = %v, want %v", converted, 100)
    }
    // Test JSON with OilBarrel unit
    oil_barrelsJSON := []byte(`{"value": 100, "unit": "OilBarrel"}`)
    oil_barrelsResult, err := factory.FromDtoJSON(oil_barrelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with OilBarrel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = oil_barrelsResult.Convert(units.VolumeOilBarrel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OilBarrel = %v, want %v", converted, 100)
    }
    // Test JSON with UsBeerBarrel unit
    us_beer_barrelsJSON := []byte(`{"value": 100, "unit": "UsBeerBarrel"}`)
    us_beer_barrelsResult, err := factory.FromDtoJSON(us_beer_barrelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsBeerBarrel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_beer_barrelsResult.Convert(units.VolumeUsBeerBarrel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsBeerBarrel = %v, want %v", converted, 100)
    }
    // Test JSON with ImperialBeerBarrel unit
    imperial_beer_barrelsJSON := []byte(`{"value": 100, "unit": "ImperialBeerBarrel"}`)
    imperial_beer_barrelsResult, err := factory.FromDtoJSON(imperial_beer_barrelsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ImperialBeerBarrel unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = imperial_beer_barrelsResult.Convert(units.VolumeImperialBeerBarrel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ImperialBeerBarrel = %v, want %v", converted, 100)
    }
    // Test JSON with UsQuart unit
    us_quartsJSON := []byte(`{"value": 100, "unit": "UsQuart"}`)
    us_quartsResult, err := factory.FromDtoJSON(us_quartsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsQuart unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_quartsResult.Convert(units.VolumeUsQuart)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsQuart = %v, want %v", converted, 100)
    }
    // Test JSON with ImperialQuart unit
    imperial_quartsJSON := []byte(`{"value": 100, "unit": "ImperialQuart"}`)
    imperial_quartsResult, err := factory.FromDtoJSON(imperial_quartsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ImperialQuart unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = imperial_quartsResult.Convert(units.VolumeImperialQuart)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ImperialQuart = %v, want %v", converted, 100)
    }
    // Test JSON with UsPint unit
    us_pintsJSON := []byte(`{"value": 100, "unit": "UsPint"}`)
    us_pintsResult, err := factory.FromDtoJSON(us_pintsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsPint unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_pintsResult.Convert(units.VolumeUsPint)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsPint = %v, want %v", converted, 100)
    }
    // Test JSON with AcreFoot unit
    acre_feetJSON := []byte(`{"value": 100, "unit": "AcreFoot"}`)
    acre_feetResult, err := factory.FromDtoJSON(acre_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AcreFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = acre_feetResult.Convert(units.VolumeAcreFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AcreFoot = %v, want %v", converted, 100)
    }
    // Test JSON with ImperialPint unit
    imperial_pintsJSON := []byte(`{"value": 100, "unit": "ImperialPint"}`)
    imperial_pintsResult, err := factory.FromDtoJSON(imperial_pintsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ImperialPint unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = imperial_pintsResult.Convert(units.VolumeImperialPint)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ImperialPint = %v, want %v", converted, 100)
    }
    // Test JSON with BoardFoot unit
    board_feetJSON := []byte(`{"value": 100, "unit": "BoardFoot"}`)
    board_feetResult, err := factory.FromDtoJSON(board_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with BoardFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = board_feetResult.Convert(units.VolumeBoardFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for BoardFoot = %v, want %v", converted, 100)
    }
    // Test JSON with Nanoliter unit
    nanolitersJSON := []byte(`{"value": 100, "unit": "Nanoliter"}`)
    nanolitersResult, err := factory.FromDtoJSON(nanolitersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanoliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanolitersResult.Convert(units.VolumeNanoliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanoliter = %v, want %v", converted, 100)
    }
    // Test JSON with Microliter unit
    microlitersJSON := []byte(`{"value": 100, "unit": "Microliter"}`)
    microlitersResult, err := factory.FromDtoJSON(microlitersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microlitersResult.Convert(units.VolumeMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microliter = %v, want %v", converted, 100)
    }
    // Test JSON with Milliliter unit
    millilitersJSON := []byte(`{"value": 100, "unit": "Milliliter"}`)
    millilitersResult, err := factory.FromDtoJSON(millilitersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Milliliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millilitersResult.Convert(units.VolumeMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliliter = %v, want %v", converted, 100)
    }
    // Test JSON with Centiliter unit
    centilitersJSON := []byte(`{"value": 100, "unit": "Centiliter"}`)
    centilitersResult, err := factory.FromDtoJSON(centilitersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Centiliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centilitersResult.Convert(units.VolumeCentiliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centiliter = %v, want %v", converted, 100)
    }
    // Test JSON with Deciliter unit
    decilitersJSON := []byte(`{"value": 100, "unit": "Deciliter"}`)
    decilitersResult, err := factory.FromDtoJSON(decilitersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Deciliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decilitersResult.Convert(units.VolumeDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Deciliter = %v, want %v", converted, 100)
    }
    // Test JSON with Decaliter unit
    decalitersJSON := []byte(`{"value": 100, "unit": "Decaliter"}`)
    decalitersResult, err := factory.FromDtoJSON(decalitersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Decaliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decalitersResult.Convert(units.VolumeDecaliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decaliter = %v, want %v", converted, 100)
    }
    // Test JSON with Hectoliter unit
    hectolitersJSON := []byte(`{"value": 100, "unit": "Hectoliter"}`)
    hectolitersResult, err := factory.FromDtoJSON(hectolitersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Hectoliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectolitersResult.Convert(units.VolumeHectoliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hectoliter = %v, want %v", converted, 100)
    }
    // Test JSON with Kiloliter unit
    kilolitersJSON := []byte(`{"value": 100, "unit": "Kiloliter"}`)
    kilolitersResult, err := factory.FromDtoJSON(kilolitersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kiloliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilolitersResult.Convert(units.VolumeKiloliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kiloliter = %v, want %v", converted, 100)
    }
    // Test JSON with Megaliter unit
    megalitersJSON := []byte(`{"value": 100, "unit": "Megaliter"}`)
    megalitersResult, err := factory.FromDtoJSON(megalitersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megaliter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megalitersResult.Convert(units.VolumeMegaliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megaliter = %v, want %v", converted, 100)
    }
    // Test JSON with HectocubicMeter unit
    hectocubic_metersJSON := []byte(`{"value": 100, "unit": "HectocubicMeter"}`)
    hectocubic_metersResult, err := factory.FromDtoJSON(hectocubic_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HectocubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectocubic_metersResult.Convert(units.VolumeHectocubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectocubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilocubicMeter unit
    kilocubic_metersJSON := []byte(`{"value": 100, "unit": "KilocubicMeter"}`)
    kilocubic_metersResult, err := factory.FromDtoJSON(kilocubic_metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilocubicMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocubic_metersResult.Convert(units.VolumeKilocubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocubicMeter = %v, want %v", converted, 100)
    }
    // Test JSON with HectocubicFoot unit
    hectocubic_feetJSON := []byte(`{"value": 100, "unit": "HectocubicFoot"}`)
    hectocubic_feetResult, err := factory.FromDtoJSON(hectocubic_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HectocubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectocubic_feetResult.Convert(units.VolumeHectocubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectocubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with KilocubicFoot unit
    kilocubic_feetJSON := []byte(`{"value": 100, "unit": "KilocubicFoot"}`)
    kilocubic_feetResult, err := factory.FromDtoJSON(kilocubic_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilocubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilocubic_feetResult.Convert(units.VolumeKilocubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilocubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with MegacubicFoot unit
    megacubic_feetJSON := []byte(`{"value": 100, "unit": "MegacubicFoot"}`)
    megacubic_feetResult, err := factory.FromDtoJSON(megacubic_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegacubicFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megacubic_feetResult.Convert(units.VolumeMegacubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegacubicFoot = %v, want %v", converted, 100)
    }
    // Test JSON with KiloimperialGallon unit
    kiloimperial_gallonsJSON := []byte(`{"value": 100, "unit": "KiloimperialGallon"}`)
    kiloimperial_gallonsResult, err := factory.FromDtoJSON(kiloimperial_gallonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KiloimperialGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloimperial_gallonsResult.Convert(units.VolumeKiloimperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KiloimperialGallon = %v, want %v", converted, 100)
    }
    // Test JSON with MegaimperialGallon unit
    megaimperial_gallonsJSON := []byte(`{"value": 100, "unit": "MegaimperialGallon"}`)
    megaimperial_gallonsResult, err := factory.FromDtoJSON(megaimperial_gallonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegaimperialGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaimperial_gallonsResult.Convert(units.VolumeMegaimperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegaimperialGallon = %v, want %v", converted, 100)
    }
    // Test JSON with DecausGallon unit
    decaus_gallonsJSON := []byte(`{"value": 100, "unit": "DecausGallon"}`)
    decaus_gallonsResult, err := factory.FromDtoJSON(decaus_gallonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecausGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decaus_gallonsResult.Convert(units.VolumeDecausGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecausGallon = %v, want %v", converted, 100)
    }
    // Test JSON with DeciusGallon unit
    decius_gallonsJSON := []byte(`{"value": 100, "unit": "DeciusGallon"}`)
    decius_gallonsResult, err := factory.FromDtoJSON(decius_gallonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DeciusGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decius_gallonsResult.Convert(units.VolumeDeciusGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DeciusGallon = %v, want %v", converted, 100)
    }
    // Test JSON with HectousGallon unit
    hectous_gallonsJSON := []byte(`{"value": 100, "unit": "HectousGallon"}`)
    hectous_gallonsResult, err := factory.FromDtoJSON(hectous_gallonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with HectousGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectous_gallonsResult.Convert(units.VolumeHectousGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for HectousGallon = %v, want %v", converted, 100)
    }
    // Test JSON with KilousGallon unit
    kilous_gallonsJSON := []byte(`{"value": 100, "unit": "KilousGallon"}`)
    kilous_gallonsResult, err := factory.FromDtoJSON(kilous_gallonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilousGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilous_gallonsResult.Convert(units.VolumeKilousGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilousGallon = %v, want %v", converted, 100)
    }
    // Test JSON with MegausGallon unit
    megaus_gallonsJSON := []byte(`{"value": 100, "unit": "MegausGallon"}`)
    megaus_gallonsResult, err := factory.FromDtoJSON(megaus_gallonsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegausGallon unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaus_gallonsResult.Convert(units.VolumeMegausGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegausGallon = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "CubicMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromLiters function
func TestVolumeFactory_FromLiters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLiters(100)
    if err != nil {
        t.Errorf("FromLiters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeLiter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLiters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLiters(math.NaN())
    if err == nil {
        t.Error("FromLiters() with NaN value should return error")
    }

    _, err = factory.FromLiters(math.Inf(1))
    if err == nil {
        t.Error("FromLiters() with +Inf value should return error")
    }

    _, err = factory.FromLiters(math.Inf(-1))
    if err == nil {
        t.Error("FromLiters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLiters(0)
    if err != nil {
        t.Errorf("FromLiters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeLiter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLiters() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicMeters function
func TestVolumeFactory_FromCubicMeters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicMeters(100)
    if err != nil {
        t.Errorf("FromCubicMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeCubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicMeters(math.NaN())
    if err == nil {
        t.Error("FromCubicMeters() with NaN value should return error")
    }

    _, err = factory.FromCubicMeters(math.Inf(1))
    if err == nil {
        t.Error("FromCubicMeters() with +Inf value should return error")
    }

    _, err = factory.FromCubicMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicMeters(0)
    if err != nil {
        t.Errorf("FromCubicMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeCubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicKilometers function
func TestVolumeFactory_FromCubicKilometers(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicKilometers(100)
    if err != nil {
        t.Errorf("FromCubicKilometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeCubicKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicKilometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicKilometers(math.NaN())
    if err == nil {
        t.Error("FromCubicKilometers() with NaN value should return error")
    }

    _, err = factory.FromCubicKilometers(math.Inf(1))
    if err == nil {
        t.Error("FromCubicKilometers() with +Inf value should return error")
    }

    _, err = factory.FromCubicKilometers(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicKilometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicKilometers(0)
    if err != nil {
        t.Errorf("FromCubicKilometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeCubicKilometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicKilometers() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicHectometers function
func TestVolumeFactory_FromCubicHectometers(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicHectometers(100)
    if err != nil {
        t.Errorf("FromCubicHectometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeCubicHectometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicHectometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicHectometers(math.NaN())
    if err == nil {
        t.Error("FromCubicHectometers() with NaN value should return error")
    }

    _, err = factory.FromCubicHectometers(math.Inf(1))
    if err == nil {
        t.Error("FromCubicHectometers() with +Inf value should return error")
    }

    _, err = factory.FromCubicHectometers(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicHectometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicHectometers(0)
    if err != nil {
        t.Errorf("FromCubicHectometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeCubicHectometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicHectometers() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicDecimeters function
func TestVolumeFactory_FromCubicDecimeters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicDecimeters(100)
    if err != nil {
        t.Errorf("FromCubicDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeCubicDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicDecimeters(math.NaN())
    if err == nil {
        t.Error("FromCubicDecimeters() with NaN value should return error")
    }

    _, err = factory.FromCubicDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromCubicDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromCubicDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicDecimeters(0)
    if err != nil {
        t.Errorf("FromCubicDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeCubicDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicCentimeters function
func TestVolumeFactory_FromCubicCentimeters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicCentimeters(100)
    if err != nil {
        t.Errorf("FromCubicCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeCubicCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicCentimeters(math.NaN())
    if err == nil {
        t.Error("FromCubicCentimeters() with NaN value should return error")
    }

    _, err = factory.FromCubicCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromCubicCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromCubicCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicCentimeters(0)
    if err != nil {
        t.Errorf("FromCubicCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeCubicCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicMillimeters function
func TestVolumeFactory_FromCubicMillimeters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicMillimeters(100)
    if err != nil {
        t.Errorf("FromCubicMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeCubicMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicMillimeters(math.NaN())
    if err == nil {
        t.Error("FromCubicMillimeters() with NaN value should return error")
    }

    _, err = factory.FromCubicMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromCubicMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromCubicMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicMillimeters(0)
    if err != nil {
        t.Errorf("FromCubicMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeCubicMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicMicrometers function
func TestVolumeFactory_FromCubicMicrometers(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicMicrometers(100)
    if err != nil {
        t.Errorf("FromCubicMicrometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeCubicMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicMicrometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicMicrometers(math.NaN())
    if err == nil {
        t.Error("FromCubicMicrometers() with NaN value should return error")
    }

    _, err = factory.FromCubicMicrometers(math.Inf(1))
    if err == nil {
        t.Error("FromCubicMicrometers() with +Inf value should return error")
    }

    _, err = factory.FromCubicMicrometers(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicMicrometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicMicrometers(0)
    if err != nil {
        t.Errorf("FromCubicMicrometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeCubicMicrometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicMicrometers() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicMiles function
func TestVolumeFactory_FromCubicMiles(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicMiles(100)
    if err != nil {
        t.Errorf("FromCubicMiles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeCubicMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicMiles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicMiles(math.NaN())
    if err == nil {
        t.Error("FromCubicMiles() with NaN value should return error")
    }

    _, err = factory.FromCubicMiles(math.Inf(1))
    if err == nil {
        t.Error("FromCubicMiles() with +Inf value should return error")
    }

    _, err = factory.FromCubicMiles(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicMiles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicMiles(0)
    if err != nil {
        t.Errorf("FromCubicMiles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeCubicMile)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicMiles() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicYards function
func TestVolumeFactory_FromCubicYards(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicYards(100)
    if err != nil {
        t.Errorf("FromCubicYards() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeCubicYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicYards() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicYards(math.NaN())
    if err == nil {
        t.Error("FromCubicYards() with NaN value should return error")
    }

    _, err = factory.FromCubicYards(math.Inf(1))
    if err == nil {
        t.Error("FromCubicYards() with +Inf value should return error")
    }

    _, err = factory.FromCubicYards(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicYards() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicYards(0)
    if err != nil {
        t.Errorf("FromCubicYards() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeCubicYard)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicYards() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicFeet function
func TestVolumeFactory_FromCubicFeet(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicFeet(100)
    if err != nil {
        t.Errorf("FromCubicFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeCubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicFeet(math.NaN())
    if err == nil {
        t.Error("FromCubicFeet() with NaN value should return error")
    }

    _, err = factory.FromCubicFeet(math.Inf(1))
    if err == nil {
        t.Error("FromCubicFeet() with +Inf value should return error")
    }

    _, err = factory.FromCubicFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicFeet(0)
    if err != nil {
        t.Errorf("FromCubicFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeCubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicInches function
func TestVolumeFactory_FromCubicInches(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicInches(100)
    if err != nil {
        t.Errorf("FromCubicInches() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeCubicInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicInches() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicInches(math.NaN())
    if err == nil {
        t.Error("FromCubicInches() with NaN value should return error")
    }

    _, err = factory.FromCubicInches(math.Inf(1))
    if err == nil {
        t.Error("FromCubicInches() with +Inf value should return error")
    }

    _, err = factory.FromCubicInches(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicInches() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicInches(0)
    if err != nil {
        t.Errorf("FromCubicInches() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeCubicInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicInches() with zero value = %v, want 0", converted)
    }
}
// Test FromImperialGallons function
func TestVolumeFactory_FromImperialGallons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromImperialGallons(100)
    if err != nil {
        t.Errorf("FromImperialGallons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeImperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromImperialGallons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromImperialGallons(math.NaN())
    if err == nil {
        t.Error("FromImperialGallons() with NaN value should return error")
    }

    _, err = factory.FromImperialGallons(math.Inf(1))
    if err == nil {
        t.Error("FromImperialGallons() with +Inf value should return error")
    }

    _, err = factory.FromImperialGallons(math.Inf(-1))
    if err == nil {
        t.Error("FromImperialGallons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromImperialGallons(0)
    if err != nil {
        t.Errorf("FromImperialGallons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeImperialGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromImperialGallons() with zero value = %v, want 0", converted)
    }
}
// Test FromImperialOunces function
func TestVolumeFactory_FromImperialOunces(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromImperialOunces(100)
    if err != nil {
        t.Errorf("FromImperialOunces() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeImperialOunce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromImperialOunces() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromImperialOunces(math.NaN())
    if err == nil {
        t.Error("FromImperialOunces() with NaN value should return error")
    }

    _, err = factory.FromImperialOunces(math.Inf(1))
    if err == nil {
        t.Error("FromImperialOunces() with +Inf value should return error")
    }

    _, err = factory.FromImperialOunces(math.Inf(-1))
    if err == nil {
        t.Error("FromImperialOunces() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromImperialOunces(0)
    if err != nil {
        t.Errorf("FromImperialOunces() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeImperialOunce)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromImperialOunces() with zero value = %v, want 0", converted)
    }
}
// Test FromUsGallons function
func TestVolumeFactory_FromUsGallons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsGallons(100)
    if err != nil {
        t.Errorf("FromUsGallons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeUsGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsGallons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsGallons(math.NaN())
    if err == nil {
        t.Error("FromUsGallons() with NaN value should return error")
    }

    _, err = factory.FromUsGallons(math.Inf(1))
    if err == nil {
        t.Error("FromUsGallons() with +Inf value should return error")
    }

    _, err = factory.FromUsGallons(math.Inf(-1))
    if err == nil {
        t.Error("FromUsGallons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsGallons(0)
    if err != nil {
        t.Errorf("FromUsGallons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeUsGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsGallons() with zero value = %v, want 0", converted)
    }
}
// Test FromUsOunces function
func TestVolumeFactory_FromUsOunces(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsOunces(100)
    if err != nil {
        t.Errorf("FromUsOunces() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeUsOunce)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsOunces() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsOunces(math.NaN())
    if err == nil {
        t.Error("FromUsOunces() with NaN value should return error")
    }

    _, err = factory.FromUsOunces(math.Inf(1))
    if err == nil {
        t.Error("FromUsOunces() with +Inf value should return error")
    }

    _, err = factory.FromUsOunces(math.Inf(-1))
    if err == nil {
        t.Error("FromUsOunces() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsOunces(0)
    if err != nil {
        t.Errorf("FromUsOunces() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeUsOunce)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsOunces() with zero value = %v, want 0", converted)
    }
}
// Test FromUsTablespoons function
func TestVolumeFactory_FromUsTablespoons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsTablespoons(100)
    if err != nil {
        t.Errorf("FromUsTablespoons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeUsTablespoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsTablespoons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsTablespoons(math.NaN())
    if err == nil {
        t.Error("FromUsTablespoons() with NaN value should return error")
    }

    _, err = factory.FromUsTablespoons(math.Inf(1))
    if err == nil {
        t.Error("FromUsTablespoons() with +Inf value should return error")
    }

    _, err = factory.FromUsTablespoons(math.Inf(-1))
    if err == nil {
        t.Error("FromUsTablespoons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsTablespoons(0)
    if err != nil {
        t.Errorf("FromUsTablespoons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeUsTablespoon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsTablespoons() with zero value = %v, want 0", converted)
    }
}
// Test FromAuTablespoons function
func TestVolumeFactory_FromAuTablespoons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAuTablespoons(100)
    if err != nil {
        t.Errorf("FromAuTablespoons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeAuTablespoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAuTablespoons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAuTablespoons(math.NaN())
    if err == nil {
        t.Error("FromAuTablespoons() with NaN value should return error")
    }

    _, err = factory.FromAuTablespoons(math.Inf(1))
    if err == nil {
        t.Error("FromAuTablespoons() with +Inf value should return error")
    }

    _, err = factory.FromAuTablespoons(math.Inf(-1))
    if err == nil {
        t.Error("FromAuTablespoons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAuTablespoons(0)
    if err != nil {
        t.Errorf("FromAuTablespoons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeAuTablespoon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAuTablespoons() with zero value = %v, want 0", converted)
    }
}
// Test FromUkTablespoons function
func TestVolumeFactory_FromUkTablespoons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUkTablespoons(100)
    if err != nil {
        t.Errorf("FromUkTablespoons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeUkTablespoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUkTablespoons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUkTablespoons(math.NaN())
    if err == nil {
        t.Error("FromUkTablespoons() with NaN value should return error")
    }

    _, err = factory.FromUkTablespoons(math.Inf(1))
    if err == nil {
        t.Error("FromUkTablespoons() with +Inf value should return error")
    }

    _, err = factory.FromUkTablespoons(math.Inf(-1))
    if err == nil {
        t.Error("FromUkTablespoons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUkTablespoons(0)
    if err != nil {
        t.Errorf("FromUkTablespoons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeUkTablespoon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUkTablespoons() with zero value = %v, want 0", converted)
    }
}
// Test FromMetricTeaspoons function
func TestVolumeFactory_FromMetricTeaspoons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMetricTeaspoons(100)
    if err != nil {
        t.Errorf("FromMetricTeaspoons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeMetricTeaspoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMetricTeaspoons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMetricTeaspoons(math.NaN())
    if err == nil {
        t.Error("FromMetricTeaspoons() with NaN value should return error")
    }

    _, err = factory.FromMetricTeaspoons(math.Inf(1))
    if err == nil {
        t.Error("FromMetricTeaspoons() with +Inf value should return error")
    }

    _, err = factory.FromMetricTeaspoons(math.Inf(-1))
    if err == nil {
        t.Error("FromMetricTeaspoons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMetricTeaspoons(0)
    if err != nil {
        t.Errorf("FromMetricTeaspoons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeMetricTeaspoon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMetricTeaspoons() with zero value = %v, want 0", converted)
    }
}
// Test FromUsTeaspoons function
func TestVolumeFactory_FromUsTeaspoons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsTeaspoons(100)
    if err != nil {
        t.Errorf("FromUsTeaspoons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeUsTeaspoon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsTeaspoons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsTeaspoons(math.NaN())
    if err == nil {
        t.Error("FromUsTeaspoons() with NaN value should return error")
    }

    _, err = factory.FromUsTeaspoons(math.Inf(1))
    if err == nil {
        t.Error("FromUsTeaspoons() with +Inf value should return error")
    }

    _, err = factory.FromUsTeaspoons(math.Inf(-1))
    if err == nil {
        t.Error("FromUsTeaspoons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsTeaspoons(0)
    if err != nil {
        t.Errorf("FromUsTeaspoons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeUsTeaspoon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsTeaspoons() with zero value = %v, want 0", converted)
    }
}
// Test FromMetricCups function
func TestVolumeFactory_FromMetricCups(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMetricCups(100)
    if err != nil {
        t.Errorf("FromMetricCups() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeMetricCup)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMetricCups() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMetricCups(math.NaN())
    if err == nil {
        t.Error("FromMetricCups() with NaN value should return error")
    }

    _, err = factory.FromMetricCups(math.Inf(1))
    if err == nil {
        t.Error("FromMetricCups() with +Inf value should return error")
    }

    _, err = factory.FromMetricCups(math.Inf(-1))
    if err == nil {
        t.Error("FromMetricCups() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMetricCups(0)
    if err != nil {
        t.Errorf("FromMetricCups() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeMetricCup)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMetricCups() with zero value = %v, want 0", converted)
    }
}
// Test FromUsCustomaryCups function
func TestVolumeFactory_FromUsCustomaryCups(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsCustomaryCups(100)
    if err != nil {
        t.Errorf("FromUsCustomaryCups() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeUsCustomaryCup)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsCustomaryCups() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsCustomaryCups(math.NaN())
    if err == nil {
        t.Error("FromUsCustomaryCups() with NaN value should return error")
    }

    _, err = factory.FromUsCustomaryCups(math.Inf(1))
    if err == nil {
        t.Error("FromUsCustomaryCups() with +Inf value should return error")
    }

    _, err = factory.FromUsCustomaryCups(math.Inf(-1))
    if err == nil {
        t.Error("FromUsCustomaryCups() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsCustomaryCups(0)
    if err != nil {
        t.Errorf("FromUsCustomaryCups() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeUsCustomaryCup)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsCustomaryCups() with zero value = %v, want 0", converted)
    }
}
// Test FromUsLegalCups function
func TestVolumeFactory_FromUsLegalCups(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsLegalCups(100)
    if err != nil {
        t.Errorf("FromUsLegalCups() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeUsLegalCup)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsLegalCups() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsLegalCups(math.NaN())
    if err == nil {
        t.Error("FromUsLegalCups() with NaN value should return error")
    }

    _, err = factory.FromUsLegalCups(math.Inf(1))
    if err == nil {
        t.Error("FromUsLegalCups() with +Inf value should return error")
    }

    _, err = factory.FromUsLegalCups(math.Inf(-1))
    if err == nil {
        t.Error("FromUsLegalCups() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsLegalCups(0)
    if err != nil {
        t.Errorf("FromUsLegalCups() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeUsLegalCup)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsLegalCups() with zero value = %v, want 0", converted)
    }
}
// Test FromOilBarrels function
func TestVolumeFactory_FromOilBarrels(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOilBarrels(100)
    if err != nil {
        t.Errorf("FromOilBarrels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeOilBarrel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOilBarrels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOilBarrels(math.NaN())
    if err == nil {
        t.Error("FromOilBarrels() with NaN value should return error")
    }

    _, err = factory.FromOilBarrels(math.Inf(1))
    if err == nil {
        t.Error("FromOilBarrels() with +Inf value should return error")
    }

    _, err = factory.FromOilBarrels(math.Inf(-1))
    if err == nil {
        t.Error("FromOilBarrels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOilBarrels(0)
    if err != nil {
        t.Errorf("FromOilBarrels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeOilBarrel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOilBarrels() with zero value = %v, want 0", converted)
    }
}
// Test FromUsBeerBarrels function
func TestVolumeFactory_FromUsBeerBarrels(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsBeerBarrels(100)
    if err != nil {
        t.Errorf("FromUsBeerBarrels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeUsBeerBarrel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsBeerBarrels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsBeerBarrels(math.NaN())
    if err == nil {
        t.Error("FromUsBeerBarrels() with NaN value should return error")
    }

    _, err = factory.FromUsBeerBarrels(math.Inf(1))
    if err == nil {
        t.Error("FromUsBeerBarrels() with +Inf value should return error")
    }

    _, err = factory.FromUsBeerBarrels(math.Inf(-1))
    if err == nil {
        t.Error("FromUsBeerBarrels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsBeerBarrels(0)
    if err != nil {
        t.Errorf("FromUsBeerBarrels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeUsBeerBarrel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsBeerBarrels() with zero value = %v, want 0", converted)
    }
}
// Test FromImperialBeerBarrels function
func TestVolumeFactory_FromImperialBeerBarrels(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromImperialBeerBarrels(100)
    if err != nil {
        t.Errorf("FromImperialBeerBarrels() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeImperialBeerBarrel)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromImperialBeerBarrels() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromImperialBeerBarrels(math.NaN())
    if err == nil {
        t.Error("FromImperialBeerBarrels() with NaN value should return error")
    }

    _, err = factory.FromImperialBeerBarrels(math.Inf(1))
    if err == nil {
        t.Error("FromImperialBeerBarrels() with +Inf value should return error")
    }

    _, err = factory.FromImperialBeerBarrels(math.Inf(-1))
    if err == nil {
        t.Error("FromImperialBeerBarrels() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromImperialBeerBarrels(0)
    if err != nil {
        t.Errorf("FromImperialBeerBarrels() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeImperialBeerBarrel)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromImperialBeerBarrels() with zero value = %v, want 0", converted)
    }
}
// Test FromUsQuarts function
func TestVolumeFactory_FromUsQuarts(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsQuarts(100)
    if err != nil {
        t.Errorf("FromUsQuarts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeUsQuart)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsQuarts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsQuarts(math.NaN())
    if err == nil {
        t.Error("FromUsQuarts() with NaN value should return error")
    }

    _, err = factory.FromUsQuarts(math.Inf(1))
    if err == nil {
        t.Error("FromUsQuarts() with +Inf value should return error")
    }

    _, err = factory.FromUsQuarts(math.Inf(-1))
    if err == nil {
        t.Error("FromUsQuarts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsQuarts(0)
    if err != nil {
        t.Errorf("FromUsQuarts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeUsQuart)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsQuarts() with zero value = %v, want 0", converted)
    }
}
// Test FromImperialQuarts function
func TestVolumeFactory_FromImperialQuarts(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromImperialQuarts(100)
    if err != nil {
        t.Errorf("FromImperialQuarts() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeImperialQuart)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromImperialQuarts() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromImperialQuarts(math.NaN())
    if err == nil {
        t.Error("FromImperialQuarts() with NaN value should return error")
    }

    _, err = factory.FromImperialQuarts(math.Inf(1))
    if err == nil {
        t.Error("FromImperialQuarts() with +Inf value should return error")
    }

    _, err = factory.FromImperialQuarts(math.Inf(-1))
    if err == nil {
        t.Error("FromImperialQuarts() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromImperialQuarts(0)
    if err != nil {
        t.Errorf("FromImperialQuarts() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeImperialQuart)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromImperialQuarts() with zero value = %v, want 0", converted)
    }
}
// Test FromUsPints function
func TestVolumeFactory_FromUsPints(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsPints(100)
    if err != nil {
        t.Errorf("FromUsPints() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeUsPint)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsPints() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsPints(math.NaN())
    if err == nil {
        t.Error("FromUsPints() with NaN value should return error")
    }

    _, err = factory.FromUsPints(math.Inf(1))
    if err == nil {
        t.Error("FromUsPints() with +Inf value should return error")
    }

    _, err = factory.FromUsPints(math.Inf(-1))
    if err == nil {
        t.Error("FromUsPints() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsPints(0)
    if err != nil {
        t.Errorf("FromUsPints() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeUsPint)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsPints() with zero value = %v, want 0", converted)
    }
}
// Test FromAcreFeet function
func TestVolumeFactory_FromAcreFeet(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAcreFeet(100)
    if err != nil {
        t.Errorf("FromAcreFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeAcreFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAcreFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAcreFeet(math.NaN())
    if err == nil {
        t.Error("FromAcreFeet() with NaN value should return error")
    }

    _, err = factory.FromAcreFeet(math.Inf(1))
    if err == nil {
        t.Error("FromAcreFeet() with +Inf value should return error")
    }

    _, err = factory.FromAcreFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromAcreFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAcreFeet(0)
    if err != nil {
        t.Errorf("FromAcreFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeAcreFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAcreFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromImperialPints function
func TestVolumeFactory_FromImperialPints(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromImperialPints(100)
    if err != nil {
        t.Errorf("FromImperialPints() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeImperialPint)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromImperialPints() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromImperialPints(math.NaN())
    if err == nil {
        t.Error("FromImperialPints() with NaN value should return error")
    }

    _, err = factory.FromImperialPints(math.Inf(1))
    if err == nil {
        t.Error("FromImperialPints() with +Inf value should return error")
    }

    _, err = factory.FromImperialPints(math.Inf(-1))
    if err == nil {
        t.Error("FromImperialPints() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromImperialPints(0)
    if err != nil {
        t.Errorf("FromImperialPints() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeImperialPint)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromImperialPints() with zero value = %v, want 0", converted)
    }
}
// Test FromBoardFeet function
func TestVolumeFactory_FromBoardFeet(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromBoardFeet(100)
    if err != nil {
        t.Errorf("FromBoardFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeBoardFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromBoardFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromBoardFeet(math.NaN())
    if err == nil {
        t.Error("FromBoardFeet() with NaN value should return error")
    }

    _, err = factory.FromBoardFeet(math.Inf(1))
    if err == nil {
        t.Error("FromBoardFeet() with +Inf value should return error")
    }

    _, err = factory.FromBoardFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromBoardFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromBoardFeet(0)
    if err != nil {
        t.Errorf("FromBoardFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeBoardFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromBoardFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromNanoliters function
func TestVolumeFactory_FromNanoliters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanoliters(100)
    if err != nil {
        t.Errorf("FromNanoliters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeNanoliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanoliters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanoliters(math.NaN())
    if err == nil {
        t.Error("FromNanoliters() with NaN value should return error")
    }

    _, err = factory.FromNanoliters(math.Inf(1))
    if err == nil {
        t.Error("FromNanoliters() with +Inf value should return error")
    }

    _, err = factory.FromNanoliters(math.Inf(-1))
    if err == nil {
        t.Error("FromNanoliters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanoliters(0)
    if err != nil {
        t.Errorf("FromNanoliters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeNanoliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanoliters() with zero value = %v, want 0", converted)
    }
}
// Test FromMicroliters function
func TestVolumeFactory_FromMicroliters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicroliters(100)
    if err != nil {
        t.Errorf("FromMicroliters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeMicroliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicroliters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicroliters(math.NaN())
    if err == nil {
        t.Error("FromMicroliters() with NaN value should return error")
    }

    _, err = factory.FromMicroliters(math.Inf(1))
    if err == nil {
        t.Error("FromMicroliters() with +Inf value should return error")
    }

    _, err = factory.FromMicroliters(math.Inf(-1))
    if err == nil {
        t.Error("FromMicroliters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicroliters(0)
    if err != nil {
        t.Errorf("FromMicroliters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeMicroliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicroliters() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliliters function
func TestVolumeFactory_FromMilliliters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliliters(100)
    if err != nil {
        t.Errorf("FromMilliliters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeMilliliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliliters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliliters(math.NaN())
    if err == nil {
        t.Error("FromMilliliters() with NaN value should return error")
    }

    _, err = factory.FromMilliliters(math.Inf(1))
    if err == nil {
        t.Error("FromMilliliters() with +Inf value should return error")
    }

    _, err = factory.FromMilliliters(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliliters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliliters(0)
    if err != nil {
        t.Errorf("FromMilliliters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeMilliliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliliters() with zero value = %v, want 0", converted)
    }
}
// Test FromCentiliters function
func TestVolumeFactory_FromCentiliters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentiliters(100)
    if err != nil {
        t.Errorf("FromCentiliters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeCentiliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentiliters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentiliters(math.NaN())
    if err == nil {
        t.Error("FromCentiliters() with NaN value should return error")
    }

    _, err = factory.FromCentiliters(math.Inf(1))
    if err == nil {
        t.Error("FromCentiliters() with +Inf value should return error")
    }

    _, err = factory.FromCentiliters(math.Inf(-1))
    if err == nil {
        t.Error("FromCentiliters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentiliters(0)
    if err != nil {
        t.Errorf("FromCentiliters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeCentiliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentiliters() with zero value = %v, want 0", converted)
    }
}
// Test FromDeciliters function
func TestVolumeFactory_FromDeciliters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDeciliters(100)
    if err != nil {
        t.Errorf("FromDeciliters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeDeciliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDeciliters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDeciliters(math.NaN())
    if err == nil {
        t.Error("FromDeciliters() with NaN value should return error")
    }

    _, err = factory.FromDeciliters(math.Inf(1))
    if err == nil {
        t.Error("FromDeciliters() with +Inf value should return error")
    }

    _, err = factory.FromDeciliters(math.Inf(-1))
    if err == nil {
        t.Error("FromDeciliters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDeciliters(0)
    if err != nil {
        t.Errorf("FromDeciliters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeDeciliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDeciliters() with zero value = %v, want 0", converted)
    }
}
// Test FromDecaliters function
func TestVolumeFactory_FromDecaliters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecaliters(100)
    if err != nil {
        t.Errorf("FromDecaliters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeDecaliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecaliters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecaliters(math.NaN())
    if err == nil {
        t.Error("FromDecaliters() with NaN value should return error")
    }

    _, err = factory.FromDecaliters(math.Inf(1))
    if err == nil {
        t.Error("FromDecaliters() with +Inf value should return error")
    }

    _, err = factory.FromDecaliters(math.Inf(-1))
    if err == nil {
        t.Error("FromDecaliters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecaliters(0)
    if err != nil {
        t.Errorf("FromDecaliters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeDecaliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecaliters() with zero value = %v, want 0", converted)
    }
}
// Test FromHectoliters function
func TestVolumeFactory_FromHectoliters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectoliters(100)
    if err != nil {
        t.Errorf("FromHectoliters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeHectoliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectoliters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectoliters(math.NaN())
    if err == nil {
        t.Error("FromHectoliters() with NaN value should return error")
    }

    _, err = factory.FromHectoliters(math.Inf(1))
    if err == nil {
        t.Error("FromHectoliters() with +Inf value should return error")
    }

    _, err = factory.FromHectoliters(math.Inf(-1))
    if err == nil {
        t.Error("FromHectoliters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectoliters(0)
    if err != nil {
        t.Errorf("FromHectoliters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeHectoliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectoliters() with zero value = %v, want 0", converted)
    }
}
// Test FromKiloliters function
func TestVolumeFactory_FromKiloliters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKiloliters(100)
    if err != nil {
        t.Errorf("FromKiloliters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeKiloliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKiloliters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKiloliters(math.NaN())
    if err == nil {
        t.Error("FromKiloliters() with NaN value should return error")
    }

    _, err = factory.FromKiloliters(math.Inf(1))
    if err == nil {
        t.Error("FromKiloliters() with +Inf value should return error")
    }

    _, err = factory.FromKiloliters(math.Inf(-1))
    if err == nil {
        t.Error("FromKiloliters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKiloliters(0)
    if err != nil {
        t.Errorf("FromKiloliters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeKiloliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKiloliters() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaliters function
func TestVolumeFactory_FromMegaliters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaliters(100)
    if err != nil {
        t.Errorf("FromMegaliters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeMegaliter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegaliters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegaliters(math.NaN())
    if err == nil {
        t.Error("FromMegaliters() with NaN value should return error")
    }

    _, err = factory.FromMegaliters(math.Inf(1))
    if err == nil {
        t.Error("FromMegaliters() with +Inf value should return error")
    }

    _, err = factory.FromMegaliters(math.Inf(-1))
    if err == nil {
        t.Error("FromMegaliters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegaliters(0)
    if err != nil {
        t.Errorf("FromMegaliters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeMegaliter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaliters() with zero value = %v, want 0", converted)
    }
}
// Test FromHectocubicMeters function
func TestVolumeFactory_FromHectocubicMeters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectocubicMeters(100)
    if err != nil {
        t.Errorf("FromHectocubicMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeHectocubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectocubicMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectocubicMeters(math.NaN())
    if err == nil {
        t.Error("FromHectocubicMeters() with NaN value should return error")
    }

    _, err = factory.FromHectocubicMeters(math.Inf(1))
    if err == nil {
        t.Error("FromHectocubicMeters() with +Inf value should return error")
    }

    _, err = factory.FromHectocubicMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromHectocubicMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectocubicMeters(0)
    if err != nil {
        t.Errorf("FromHectocubicMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeHectocubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectocubicMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromKilocubicMeters function
func TestVolumeFactory_FromKilocubicMeters(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilocubicMeters(100)
    if err != nil {
        t.Errorf("FromKilocubicMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeKilocubicMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilocubicMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilocubicMeters(math.NaN())
    if err == nil {
        t.Error("FromKilocubicMeters() with NaN value should return error")
    }

    _, err = factory.FromKilocubicMeters(math.Inf(1))
    if err == nil {
        t.Error("FromKilocubicMeters() with +Inf value should return error")
    }

    _, err = factory.FromKilocubicMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromKilocubicMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilocubicMeters(0)
    if err != nil {
        t.Errorf("FromKilocubicMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeKilocubicMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilocubicMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromHectocubicFeet function
func TestVolumeFactory_FromHectocubicFeet(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectocubicFeet(100)
    if err != nil {
        t.Errorf("FromHectocubicFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeHectocubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectocubicFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectocubicFeet(math.NaN())
    if err == nil {
        t.Error("FromHectocubicFeet() with NaN value should return error")
    }

    _, err = factory.FromHectocubicFeet(math.Inf(1))
    if err == nil {
        t.Error("FromHectocubicFeet() with +Inf value should return error")
    }

    _, err = factory.FromHectocubicFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromHectocubicFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectocubicFeet(0)
    if err != nil {
        t.Errorf("FromHectocubicFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeHectocubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectocubicFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromKilocubicFeet function
func TestVolumeFactory_FromKilocubicFeet(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilocubicFeet(100)
    if err != nil {
        t.Errorf("FromKilocubicFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeKilocubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilocubicFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilocubicFeet(math.NaN())
    if err == nil {
        t.Error("FromKilocubicFeet() with NaN value should return error")
    }

    _, err = factory.FromKilocubicFeet(math.Inf(1))
    if err == nil {
        t.Error("FromKilocubicFeet() with +Inf value should return error")
    }

    _, err = factory.FromKilocubicFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromKilocubicFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilocubicFeet(0)
    if err != nil {
        t.Errorf("FromKilocubicFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeKilocubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilocubicFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromMegacubicFeet function
func TestVolumeFactory_FromMegacubicFeet(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegacubicFeet(100)
    if err != nil {
        t.Errorf("FromMegacubicFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeMegacubicFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegacubicFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegacubicFeet(math.NaN())
    if err == nil {
        t.Error("FromMegacubicFeet() with NaN value should return error")
    }

    _, err = factory.FromMegacubicFeet(math.Inf(1))
    if err == nil {
        t.Error("FromMegacubicFeet() with +Inf value should return error")
    }

    _, err = factory.FromMegacubicFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromMegacubicFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegacubicFeet(0)
    if err != nil {
        t.Errorf("FromMegacubicFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeMegacubicFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegacubicFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromKiloimperialGallons function
func TestVolumeFactory_FromKiloimperialGallons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKiloimperialGallons(100)
    if err != nil {
        t.Errorf("FromKiloimperialGallons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeKiloimperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKiloimperialGallons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKiloimperialGallons(math.NaN())
    if err == nil {
        t.Error("FromKiloimperialGallons() with NaN value should return error")
    }

    _, err = factory.FromKiloimperialGallons(math.Inf(1))
    if err == nil {
        t.Error("FromKiloimperialGallons() with +Inf value should return error")
    }

    _, err = factory.FromKiloimperialGallons(math.Inf(-1))
    if err == nil {
        t.Error("FromKiloimperialGallons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKiloimperialGallons(0)
    if err != nil {
        t.Errorf("FromKiloimperialGallons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeKiloimperialGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKiloimperialGallons() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaimperialGallons function
func TestVolumeFactory_FromMegaimperialGallons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaimperialGallons(100)
    if err != nil {
        t.Errorf("FromMegaimperialGallons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeMegaimperialGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegaimperialGallons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegaimperialGallons(math.NaN())
    if err == nil {
        t.Error("FromMegaimperialGallons() with NaN value should return error")
    }

    _, err = factory.FromMegaimperialGallons(math.Inf(1))
    if err == nil {
        t.Error("FromMegaimperialGallons() with +Inf value should return error")
    }

    _, err = factory.FromMegaimperialGallons(math.Inf(-1))
    if err == nil {
        t.Error("FromMegaimperialGallons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegaimperialGallons(0)
    if err != nil {
        t.Errorf("FromMegaimperialGallons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeMegaimperialGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaimperialGallons() with zero value = %v, want 0", converted)
    }
}
// Test FromDecausGallons function
func TestVolumeFactory_FromDecausGallons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecausGallons(100)
    if err != nil {
        t.Errorf("FromDecausGallons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeDecausGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecausGallons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecausGallons(math.NaN())
    if err == nil {
        t.Error("FromDecausGallons() with NaN value should return error")
    }

    _, err = factory.FromDecausGallons(math.Inf(1))
    if err == nil {
        t.Error("FromDecausGallons() with +Inf value should return error")
    }

    _, err = factory.FromDecausGallons(math.Inf(-1))
    if err == nil {
        t.Error("FromDecausGallons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecausGallons(0)
    if err != nil {
        t.Errorf("FromDecausGallons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeDecausGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecausGallons() with zero value = %v, want 0", converted)
    }
}
// Test FromDeciusGallons function
func TestVolumeFactory_FromDeciusGallons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDeciusGallons(100)
    if err != nil {
        t.Errorf("FromDeciusGallons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeDeciusGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDeciusGallons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDeciusGallons(math.NaN())
    if err == nil {
        t.Error("FromDeciusGallons() with NaN value should return error")
    }

    _, err = factory.FromDeciusGallons(math.Inf(1))
    if err == nil {
        t.Error("FromDeciusGallons() with +Inf value should return error")
    }

    _, err = factory.FromDeciusGallons(math.Inf(-1))
    if err == nil {
        t.Error("FromDeciusGallons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDeciusGallons(0)
    if err != nil {
        t.Errorf("FromDeciusGallons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeDeciusGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDeciusGallons() with zero value = %v, want 0", converted)
    }
}
// Test FromHectousGallons function
func TestVolumeFactory_FromHectousGallons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectousGallons(100)
    if err != nil {
        t.Errorf("FromHectousGallons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeHectousGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectousGallons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectousGallons(math.NaN())
    if err == nil {
        t.Error("FromHectousGallons() with NaN value should return error")
    }

    _, err = factory.FromHectousGallons(math.Inf(1))
    if err == nil {
        t.Error("FromHectousGallons() with +Inf value should return error")
    }

    _, err = factory.FromHectousGallons(math.Inf(-1))
    if err == nil {
        t.Error("FromHectousGallons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectousGallons(0)
    if err != nil {
        t.Errorf("FromHectousGallons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeHectousGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectousGallons() with zero value = %v, want 0", converted)
    }
}
// Test FromKilousGallons function
func TestVolumeFactory_FromKilousGallons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilousGallons(100)
    if err != nil {
        t.Errorf("FromKilousGallons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeKilousGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilousGallons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilousGallons(math.NaN())
    if err == nil {
        t.Error("FromKilousGallons() with NaN value should return error")
    }

    _, err = factory.FromKilousGallons(math.Inf(1))
    if err == nil {
        t.Error("FromKilousGallons() with +Inf value should return error")
    }

    _, err = factory.FromKilousGallons(math.Inf(-1))
    if err == nil {
        t.Error("FromKilousGallons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilousGallons(0)
    if err != nil {
        t.Errorf("FromKilousGallons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeKilousGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilousGallons() with zero value = %v, want 0", converted)
    }
}
// Test FromMegausGallons function
func TestVolumeFactory_FromMegausGallons(t *testing.T) {
    factory := units.VolumeFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegausGallons(100)
    if err != nil {
        t.Errorf("FromMegausGallons() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumeMegausGallon)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegausGallons() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegausGallons(math.NaN())
    if err == nil {
        t.Error("FromMegausGallons() with NaN value should return error")
    }

    _, err = factory.FromMegausGallons(math.Inf(1))
    if err == nil {
        t.Error("FromMegausGallons() with +Inf value should return error")
    }

    _, err = factory.FromMegausGallons(math.Inf(-1))
    if err == nil {
        t.Error("FromMegausGallons() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegausGallons(0)
    if err != nil {
        t.Errorf("FromMegausGallons() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumeMegausGallon)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegausGallons() with zero value = %v, want 0", converted)
    }
}

func TestVolumeToString(t *testing.T) {
	factory := units.VolumeFactory{}
	a, err := factory.CreateVolume(45, units.VolumeCubicMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.VolumeCubicMeter, 2)
	expected := "45.00 " + units.GetVolumeAbbreviation(units.VolumeCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.VolumeCubicMeter, -1)
	expected = "45 " + units.GetVolumeAbbreviation(units.VolumeCubicMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestVolume_EqualityAndComparison(t *testing.T) {
	factory := units.VolumeFactory{}
	a1, _ := factory.CreateVolume(60, units.VolumeCubicMeter)
	a2, _ := factory.CreateVolume(60, units.VolumeCubicMeter)
	a3, _ := factory.CreateVolume(120, units.VolumeCubicMeter)

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

func TestVolume_Arithmetic(t *testing.T) {
	factory := units.VolumeFactory{}
	a1, _ := factory.CreateVolume(30, units.VolumeCubicMeter)
	a2, _ := factory.CreateVolume(45, units.VolumeCubicMeter)

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


func TestGetVolumeAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.VolumeUnits
        want string
    }{
        {
            name: "Liter abbreviation",
            unit: units.VolumeLiter,
            want: "l",
        },
        {
            name: "CubicMeter abbreviation",
            unit: units.VolumeCubicMeter,
            want: "m³",
        },
        {
            name: "CubicKilometer abbreviation",
            unit: units.VolumeCubicKilometer,
            want: "km³",
        },
        {
            name: "CubicHectometer abbreviation",
            unit: units.VolumeCubicHectometer,
            want: "hm³",
        },
        {
            name: "CubicDecimeter abbreviation",
            unit: units.VolumeCubicDecimeter,
            want: "dm³",
        },
        {
            name: "CubicCentimeter abbreviation",
            unit: units.VolumeCubicCentimeter,
            want: "cm³",
        },
        {
            name: "CubicMillimeter abbreviation",
            unit: units.VolumeCubicMillimeter,
            want: "mm³",
        },
        {
            name: "CubicMicrometer abbreviation",
            unit: units.VolumeCubicMicrometer,
            want: "µm³",
        },
        {
            name: "CubicMile abbreviation",
            unit: units.VolumeCubicMile,
            want: "mi³",
        },
        {
            name: "CubicYard abbreviation",
            unit: units.VolumeCubicYard,
            want: "yd³",
        },
        {
            name: "CubicFoot abbreviation",
            unit: units.VolumeCubicFoot,
            want: "ft³",
        },
        {
            name: "CubicInch abbreviation",
            unit: units.VolumeCubicInch,
            want: "in³",
        },
        {
            name: "ImperialGallon abbreviation",
            unit: units.VolumeImperialGallon,
            want: "gal (imp.)",
        },
        {
            name: "ImperialOunce abbreviation",
            unit: units.VolumeImperialOunce,
            want: "oz (imp.)",
        },
        {
            name: "UsGallon abbreviation",
            unit: units.VolumeUsGallon,
            want: "gal (U.S.)",
        },
        {
            name: "UsOunce abbreviation",
            unit: units.VolumeUsOunce,
            want: "oz (U.S.)",
        },
        {
            name: "UsTablespoon abbreviation",
            unit: units.VolumeUsTablespoon,
            want: "",
        },
        {
            name: "AuTablespoon abbreviation",
            unit: units.VolumeAuTablespoon,
            want: "",
        },
        {
            name: "UkTablespoon abbreviation",
            unit: units.VolumeUkTablespoon,
            want: "",
        },
        {
            name: "MetricTeaspoon abbreviation",
            unit: units.VolumeMetricTeaspoon,
            want: "tsp",
        },
        {
            name: "UsTeaspoon abbreviation",
            unit: units.VolumeUsTeaspoon,
            want: "",
        },
        {
            name: "MetricCup abbreviation",
            unit: units.VolumeMetricCup,
            want: "",
        },
        {
            name: "UsCustomaryCup abbreviation",
            unit: units.VolumeUsCustomaryCup,
            want: "",
        },
        {
            name: "UsLegalCup abbreviation",
            unit: units.VolumeUsLegalCup,
            want: "",
        },
        {
            name: "OilBarrel abbreviation",
            unit: units.VolumeOilBarrel,
            want: "bbl",
        },
        {
            name: "UsBeerBarrel abbreviation",
            unit: units.VolumeUsBeerBarrel,
            want: "bl (U.S.)",
        },
        {
            name: "ImperialBeerBarrel abbreviation",
            unit: units.VolumeImperialBeerBarrel,
            want: "bl (imp.)",
        },
        {
            name: "UsQuart abbreviation",
            unit: units.VolumeUsQuart,
            want: "qt (U.S.)",
        },
        {
            name: "ImperialQuart abbreviation",
            unit: units.VolumeImperialQuart,
            want: "qt (imp.)",
        },
        {
            name: "UsPint abbreviation",
            unit: units.VolumeUsPint,
            want: "pt (U.S.)",
        },
        {
            name: "AcreFoot abbreviation",
            unit: units.VolumeAcreFoot,
            want: "ac-ft",
        },
        {
            name: "ImperialPint abbreviation",
            unit: units.VolumeImperialPint,
            want: "pt (imp.)",
        },
        {
            name: "BoardFoot abbreviation",
            unit: units.VolumeBoardFoot,
            want: "bf",
        },
        {
            name: "Nanoliter abbreviation",
            unit: units.VolumeNanoliter,
            want: "nl",
        },
        {
            name: "Microliter abbreviation",
            unit: units.VolumeMicroliter,
            want: "μl",
        },
        {
            name: "Milliliter abbreviation",
            unit: units.VolumeMilliliter,
            want: "ml",
        },
        {
            name: "Centiliter abbreviation",
            unit: units.VolumeCentiliter,
            want: "cl",
        },
        {
            name: "Deciliter abbreviation",
            unit: units.VolumeDeciliter,
            want: "dl",
        },
        {
            name: "Decaliter abbreviation",
            unit: units.VolumeDecaliter,
            want: "dal",
        },
        {
            name: "Hectoliter abbreviation",
            unit: units.VolumeHectoliter,
            want: "hl",
        },
        {
            name: "Kiloliter abbreviation",
            unit: units.VolumeKiloliter,
            want: "kl",
        },
        {
            name: "Megaliter abbreviation",
            unit: units.VolumeMegaliter,
            want: "Ml",
        },
        {
            name: "HectocubicMeter abbreviation",
            unit: units.VolumeHectocubicMeter,
            want: "hm³",
        },
        {
            name: "KilocubicMeter abbreviation",
            unit: units.VolumeKilocubicMeter,
            want: "km³",
        },
        {
            name: "HectocubicFoot abbreviation",
            unit: units.VolumeHectocubicFoot,
            want: "hft³",
        },
        {
            name: "KilocubicFoot abbreviation",
            unit: units.VolumeKilocubicFoot,
            want: "kft³",
        },
        {
            name: "MegacubicFoot abbreviation",
            unit: units.VolumeMegacubicFoot,
            want: "Mft³",
        },
        {
            name: "KiloimperialGallon abbreviation",
            unit: units.VolumeKiloimperialGallon,
            want: "kgal (imp.)",
        },
        {
            name: "MegaimperialGallon abbreviation",
            unit: units.VolumeMegaimperialGallon,
            want: "Mgal (imp.)",
        },
        {
            name: "DecausGallon abbreviation",
            unit: units.VolumeDecausGallon,
            want: "dagal (U.S.)",
        },
        {
            name: "DeciusGallon abbreviation",
            unit: units.VolumeDeciusGallon,
            want: "dgal (U.S.)",
        },
        {
            name: "HectousGallon abbreviation",
            unit: units.VolumeHectousGallon,
            want: "hgal (U.S.)",
        },
        {
            name: "KilousGallon abbreviation",
            unit: units.VolumeKilousGallon,
            want: "kgal (U.S.)",
        },
        {
            name: "MegausGallon abbreviation",
            unit: units.VolumeMegausGallon,
            want: "Mgal (U.S.)",
        },
        {
            name: "invalid unit",
            unit: units.VolumeUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetVolumeAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetVolumeAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestVolume_String(t *testing.T) {
    factory := units.VolumeFactory{}
    
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
            unit, err := factory.CreateVolume(tt.value, units.VolumeCubicMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Volume.String() = %v, want %v", got, tt.want)
            }
        })
    }
}