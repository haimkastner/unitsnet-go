// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestVolumePerLengthDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "CubicMeterPerMeter"}`
	
	factory := units.VolumePerLengthDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.VolumePerLengthCubicMeterPerMeter {
		t.Errorf("expected unit %v, got %v", units.VolumePerLengthCubicMeterPerMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "CubicMeterPerMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestVolumePerLengthDto_ToJSON(t *testing.T) {
	dto := units.VolumePerLengthDto{
		Value: 45,
		Unit:  units.VolumePerLengthCubicMeterPerMeter,
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
	if result["unit"].(string) != string(units.VolumePerLengthCubicMeterPerMeter) {
		t.Errorf("expected unit %s, got %v", units.VolumePerLengthCubicMeterPerMeter, result["unit"])
	}
}

func TestNewVolumePerLength_InvalidValue(t *testing.T) {
	factory := units.VolumePerLengthFactory{}
	// NaN value should return an error.
	_, err := factory.CreateVolumePerLength(math.NaN(), units.VolumePerLengthCubicMeterPerMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateVolumePerLength(math.Inf(1), units.VolumePerLengthCubicMeterPerMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestVolumePerLengthConversions(t *testing.T) {
	factory := units.VolumePerLengthFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateVolumePerLength(180, units.VolumePerLengthCubicMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to CubicMetersPerMeter.
		// No expected conversion value provided for CubicMetersPerMeter, verifying result is not NaN.
		result := a.CubicMetersPerMeter()
		cacheResult := a.CubicMetersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicMetersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to LitersPerMeter.
		// No expected conversion value provided for LitersPerMeter, verifying result is not NaN.
		result := a.LitersPerMeter()
		cacheResult := a.LitersPerMeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LitersPerMeter returned NaN")
		}
	}
	{
		// Test conversion to LitersPerKilometer.
		// No expected conversion value provided for LitersPerKilometer, verifying result is not NaN.
		result := a.LitersPerKilometer()
		cacheResult := a.LitersPerKilometer()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LitersPerKilometer returned NaN")
		}
	}
	{
		// Test conversion to LitersPerMillimeter.
		// No expected conversion value provided for LitersPerMillimeter, verifying result is not NaN.
		result := a.LitersPerMillimeter()
		cacheResult := a.LitersPerMillimeter()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LitersPerMillimeter returned NaN")
		}
	}
	{
		// Test conversion to OilBarrelsPerFoot.
		// No expected conversion value provided for OilBarrelsPerFoot, verifying result is not NaN.
		result := a.OilBarrelsPerFoot()
		cacheResult := a.OilBarrelsPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to OilBarrelsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to CubicYardsPerFoot.
		// No expected conversion value provided for CubicYardsPerFoot, verifying result is not NaN.
		result := a.CubicYardsPerFoot()
		cacheResult := a.CubicYardsPerFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicYardsPerFoot returned NaN")
		}
	}
	{
		// Test conversion to CubicYardsPerUsSurveyFoot.
		// No expected conversion value provided for CubicYardsPerUsSurveyFoot, verifying result is not NaN.
		result := a.CubicYardsPerUsSurveyFoot()
		cacheResult := a.CubicYardsPerUsSurveyFoot()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CubicYardsPerUsSurveyFoot returned NaN")
		}
	}
	{
		// Test conversion to UsGallonsPerMile.
		// No expected conversion value provided for UsGallonsPerMile, verifying result is not NaN.
		result := a.UsGallonsPerMile()
		cacheResult := a.UsGallonsPerMile()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsGallonsPerMile returned NaN")
		}
	}
	{
		// Test conversion to ImperialGallonsPerMile.
		// No expected conversion value provided for ImperialGallonsPerMile, verifying result is not NaN.
		result := a.ImperialGallonsPerMile()
		cacheResult := a.ImperialGallonsPerMile()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to ImperialGallonsPerMile returned NaN")
		}
	}
}

func TestVolumePerLength_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.VolumePerLengthFactory{}
	a, err := factory.CreateVolumePerLength(90, units.VolumePerLengthCubicMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.VolumePerLengthCubicMeterPerMeter {
		t.Errorf("expected default unit CubicMeterPerMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.VolumePerLengthCubicMeterPerMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.VolumePerLengthDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.VolumePerLengthCubicMeterPerMeter {
		t.Errorf("expected unit CubicMeterPerMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestVolumePerLengthFactory_FromDto(t *testing.T) {
    factory := units.VolumePerLengthFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.VolumePerLengthDto{
        Value: 100,
        Unit:  units.VolumePerLengthCubicMeterPerMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.VolumePerLengthDto{
        Value: math.NaN(),
        Unit:  units.VolumePerLengthCubicMeterPerMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test CubicMeterPerMeter conversion
    cubic_meters_per_meterDto := units.VolumePerLengthDto{
        Value: 100,
        Unit:  units.VolumePerLengthCubicMeterPerMeter,
    }
    
    var cubic_meters_per_meterResult *units.VolumePerLength
    cubic_meters_per_meterResult, err = factory.FromDto(cubic_meters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with CubicMeterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_meterResult.Convert(units.VolumePerLengthCubicMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test LiterPerMeter conversion
    liters_per_meterDto := units.VolumePerLengthDto{
        Value: 100,
        Unit:  units.VolumePerLengthLiterPerMeter,
    }
    
    var liters_per_meterResult *units.VolumePerLength
    liters_per_meterResult, err = factory.FromDto(liters_per_meterDto)
    if err != nil {
        t.Errorf("FromDto() with LiterPerMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_meterResult.Convert(units.VolumePerLengthLiterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerMeter = %v, want %v", converted, 100)
    }
    // Test LiterPerKilometer conversion
    liters_per_kilometerDto := units.VolumePerLengthDto{
        Value: 100,
        Unit:  units.VolumePerLengthLiterPerKilometer,
    }
    
    var liters_per_kilometerResult *units.VolumePerLength
    liters_per_kilometerResult, err = factory.FromDto(liters_per_kilometerDto)
    if err != nil {
        t.Errorf("FromDto() with LiterPerKilometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_kilometerResult.Convert(units.VolumePerLengthLiterPerKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerKilometer = %v, want %v", converted, 100)
    }
    // Test LiterPerMillimeter conversion
    liters_per_millimeterDto := units.VolumePerLengthDto{
        Value: 100,
        Unit:  units.VolumePerLengthLiterPerMillimeter,
    }
    
    var liters_per_millimeterResult *units.VolumePerLength
    liters_per_millimeterResult, err = factory.FromDto(liters_per_millimeterDto)
    if err != nil {
        t.Errorf("FromDto() with LiterPerMillimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_millimeterResult.Convert(units.VolumePerLengthLiterPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerMillimeter = %v, want %v", converted, 100)
    }
    // Test OilBarrelPerFoot conversion
    oil_barrels_per_footDto := units.VolumePerLengthDto{
        Value: 100,
        Unit:  units.VolumePerLengthOilBarrelPerFoot,
    }
    
    var oil_barrels_per_footResult *units.VolumePerLength
    oil_barrels_per_footResult, err = factory.FromDto(oil_barrels_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with OilBarrelPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = oil_barrels_per_footResult.Convert(units.VolumePerLengthOilBarrelPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OilBarrelPerFoot = %v, want %v", converted, 100)
    }
    // Test CubicYardPerFoot conversion
    cubic_yards_per_footDto := units.VolumePerLengthDto{
        Value: 100,
        Unit:  units.VolumePerLengthCubicYardPerFoot,
    }
    
    var cubic_yards_per_footResult *units.VolumePerLength
    cubic_yards_per_footResult, err = factory.FromDto(cubic_yards_per_footDto)
    if err != nil {
        t.Errorf("FromDto() with CubicYardPerFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yards_per_footResult.Convert(units.VolumePerLengthCubicYardPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYardPerFoot = %v, want %v", converted, 100)
    }
    // Test CubicYardPerUsSurveyFoot conversion
    cubic_yards_per_us_survey_footDto := units.VolumePerLengthDto{
        Value: 100,
        Unit:  units.VolumePerLengthCubicYardPerUsSurveyFoot,
    }
    
    var cubic_yards_per_us_survey_footResult *units.VolumePerLength
    cubic_yards_per_us_survey_footResult, err = factory.FromDto(cubic_yards_per_us_survey_footDto)
    if err != nil {
        t.Errorf("FromDto() with CubicYardPerUsSurveyFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yards_per_us_survey_footResult.Convert(units.VolumePerLengthCubicYardPerUsSurveyFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYardPerUsSurveyFoot = %v, want %v", converted, 100)
    }
    // Test UsGallonPerMile conversion
    us_gallons_per_mileDto := units.VolumePerLengthDto{
        Value: 100,
        Unit:  units.VolumePerLengthUsGallonPerMile,
    }
    
    var us_gallons_per_mileResult *units.VolumePerLength
    us_gallons_per_mileResult, err = factory.FromDto(us_gallons_per_mileDto)
    if err != nil {
        t.Errorf("FromDto() with UsGallonPerMile returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_gallons_per_mileResult.Convert(units.VolumePerLengthUsGallonPerMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsGallonPerMile = %v, want %v", converted, 100)
    }
    // Test ImperialGallonPerMile conversion
    imperial_gallons_per_mileDto := units.VolumePerLengthDto{
        Value: 100,
        Unit:  units.VolumePerLengthImperialGallonPerMile,
    }
    
    var imperial_gallons_per_mileResult *units.VolumePerLength
    imperial_gallons_per_mileResult, err = factory.FromDto(imperial_gallons_per_mileDto)
    if err != nil {
        t.Errorf("FromDto() with ImperialGallonPerMile returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = imperial_gallons_per_mileResult.Convert(units.VolumePerLengthImperialGallonPerMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ImperialGallonPerMile = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.VolumePerLengthDto{
        Value: 0,
        Unit:  units.VolumePerLengthCubicMeterPerMeter,
    }
    
    var zeroResult *units.VolumePerLength
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestVolumePerLengthFactory_FromDtoJSON(t *testing.T) {
    factory := units.VolumePerLengthFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "CubicMeterPerMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "CubicMeterPerMeter"}`)
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
    nanJSON, _ := json.Marshal(units.VolumePerLengthDto{
        Value: nanValue,
        Unit:  units.VolumePerLengthCubicMeterPerMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with CubicMeterPerMeter unit
    cubic_meters_per_meterJSON := []byte(`{"value": 100, "unit": "CubicMeterPerMeter"}`)
    cubic_meters_per_meterResult, err := factory.FromDtoJSON(cubic_meters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicMeterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_meters_per_meterResult.Convert(units.VolumePerLengthCubicMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicMeterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with LiterPerMeter unit
    liters_per_meterJSON := []byte(`{"value": 100, "unit": "LiterPerMeter"}`)
    liters_per_meterResult, err := factory.FromDtoJSON(liters_per_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LiterPerMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_meterResult.Convert(units.VolumePerLengthLiterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerMeter = %v, want %v", converted, 100)
    }
    // Test JSON with LiterPerKilometer unit
    liters_per_kilometerJSON := []byte(`{"value": 100, "unit": "LiterPerKilometer"}`)
    liters_per_kilometerResult, err := factory.FromDtoJSON(liters_per_kilometerJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LiterPerKilometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_kilometerResult.Convert(units.VolumePerLengthLiterPerKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerKilometer = %v, want %v", converted, 100)
    }
    // Test JSON with LiterPerMillimeter unit
    liters_per_millimeterJSON := []byte(`{"value": 100, "unit": "LiterPerMillimeter"}`)
    liters_per_millimeterResult, err := factory.FromDtoJSON(liters_per_millimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LiterPerMillimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = liters_per_millimeterResult.Convert(units.VolumePerLengthLiterPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LiterPerMillimeter = %v, want %v", converted, 100)
    }
    // Test JSON with OilBarrelPerFoot unit
    oil_barrels_per_footJSON := []byte(`{"value": 100, "unit": "OilBarrelPerFoot"}`)
    oil_barrels_per_footResult, err := factory.FromDtoJSON(oil_barrels_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with OilBarrelPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = oil_barrels_per_footResult.Convert(units.VolumePerLengthOilBarrelPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for OilBarrelPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with CubicYardPerFoot unit
    cubic_yards_per_footJSON := []byte(`{"value": 100, "unit": "CubicYardPerFoot"}`)
    cubic_yards_per_footResult, err := factory.FromDtoJSON(cubic_yards_per_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicYardPerFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yards_per_footResult.Convert(units.VolumePerLengthCubicYardPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYardPerFoot = %v, want %v", converted, 100)
    }
    // Test JSON with CubicYardPerUsSurveyFoot unit
    cubic_yards_per_us_survey_footJSON := []byte(`{"value": 100, "unit": "CubicYardPerUsSurveyFoot"}`)
    cubic_yards_per_us_survey_footResult, err := factory.FromDtoJSON(cubic_yards_per_us_survey_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CubicYardPerUsSurveyFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = cubic_yards_per_us_survey_footResult.Convert(units.VolumePerLengthCubicYardPerUsSurveyFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CubicYardPerUsSurveyFoot = %v, want %v", converted, 100)
    }
    // Test JSON with UsGallonPerMile unit
    us_gallons_per_mileJSON := []byte(`{"value": 100, "unit": "UsGallonPerMile"}`)
    us_gallons_per_mileResult, err := factory.FromDtoJSON(us_gallons_per_mileJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsGallonPerMile unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_gallons_per_mileResult.Convert(units.VolumePerLengthUsGallonPerMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsGallonPerMile = %v, want %v", converted, 100)
    }
    // Test JSON with ImperialGallonPerMile unit
    imperial_gallons_per_mileJSON := []byte(`{"value": 100, "unit": "ImperialGallonPerMile"}`)
    imperial_gallons_per_mileResult, err := factory.FromDtoJSON(imperial_gallons_per_mileJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with ImperialGallonPerMile unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = imperial_gallons_per_mileResult.Convert(units.VolumePerLengthImperialGallonPerMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for ImperialGallonPerMile = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "CubicMeterPerMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromCubicMetersPerMeter function
func TestVolumePerLengthFactory_FromCubicMetersPerMeter(t *testing.T) {
    factory := units.VolumePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicMetersPerMeter(100)
    if err != nil {
        t.Errorf("FromCubicMetersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumePerLengthCubicMeterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicMetersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicMetersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromCubicMetersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromCubicMetersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromCubicMetersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromCubicMetersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicMetersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicMetersPerMeter(0)
    if err != nil {
        t.Errorf("FromCubicMetersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumePerLengthCubicMeterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicMetersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromLitersPerMeter function
func TestVolumePerLengthFactory_FromLitersPerMeter(t *testing.T) {
    factory := units.VolumePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLitersPerMeter(100)
    if err != nil {
        t.Errorf("FromLitersPerMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumePerLengthLiterPerMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLitersPerMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLitersPerMeter(math.NaN())
    if err == nil {
        t.Error("FromLitersPerMeter() with NaN value should return error")
    }

    _, err = factory.FromLitersPerMeter(math.Inf(1))
    if err == nil {
        t.Error("FromLitersPerMeter() with +Inf value should return error")
    }

    _, err = factory.FromLitersPerMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromLitersPerMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLitersPerMeter(0)
    if err != nil {
        t.Errorf("FromLitersPerMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumePerLengthLiterPerMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLitersPerMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromLitersPerKilometer function
func TestVolumePerLengthFactory_FromLitersPerKilometer(t *testing.T) {
    factory := units.VolumePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLitersPerKilometer(100)
    if err != nil {
        t.Errorf("FromLitersPerKilometer() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumePerLengthLiterPerKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLitersPerKilometer() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLitersPerKilometer(math.NaN())
    if err == nil {
        t.Error("FromLitersPerKilometer() with NaN value should return error")
    }

    _, err = factory.FromLitersPerKilometer(math.Inf(1))
    if err == nil {
        t.Error("FromLitersPerKilometer() with +Inf value should return error")
    }

    _, err = factory.FromLitersPerKilometer(math.Inf(-1))
    if err == nil {
        t.Error("FromLitersPerKilometer() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLitersPerKilometer(0)
    if err != nil {
        t.Errorf("FromLitersPerKilometer() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumePerLengthLiterPerKilometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLitersPerKilometer() with zero value = %v, want 0", converted)
    }
}
// Test FromLitersPerMillimeter function
func TestVolumePerLengthFactory_FromLitersPerMillimeter(t *testing.T) {
    factory := units.VolumePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLitersPerMillimeter(100)
    if err != nil {
        t.Errorf("FromLitersPerMillimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumePerLengthLiterPerMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLitersPerMillimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLitersPerMillimeter(math.NaN())
    if err == nil {
        t.Error("FromLitersPerMillimeter() with NaN value should return error")
    }

    _, err = factory.FromLitersPerMillimeter(math.Inf(1))
    if err == nil {
        t.Error("FromLitersPerMillimeter() with +Inf value should return error")
    }

    _, err = factory.FromLitersPerMillimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromLitersPerMillimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLitersPerMillimeter(0)
    if err != nil {
        t.Errorf("FromLitersPerMillimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumePerLengthLiterPerMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLitersPerMillimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromOilBarrelsPerFoot function
func TestVolumePerLengthFactory_FromOilBarrelsPerFoot(t *testing.T) {
    factory := units.VolumePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromOilBarrelsPerFoot(100)
    if err != nil {
        t.Errorf("FromOilBarrelsPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumePerLengthOilBarrelPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromOilBarrelsPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromOilBarrelsPerFoot(math.NaN())
    if err == nil {
        t.Error("FromOilBarrelsPerFoot() with NaN value should return error")
    }

    _, err = factory.FromOilBarrelsPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromOilBarrelsPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromOilBarrelsPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromOilBarrelsPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromOilBarrelsPerFoot(0)
    if err != nil {
        t.Errorf("FromOilBarrelsPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumePerLengthOilBarrelPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromOilBarrelsPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicYardsPerFoot function
func TestVolumePerLengthFactory_FromCubicYardsPerFoot(t *testing.T) {
    factory := units.VolumePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicYardsPerFoot(100)
    if err != nil {
        t.Errorf("FromCubicYardsPerFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumePerLengthCubicYardPerFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicYardsPerFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicYardsPerFoot(math.NaN())
    if err == nil {
        t.Error("FromCubicYardsPerFoot() with NaN value should return error")
    }

    _, err = factory.FromCubicYardsPerFoot(math.Inf(1))
    if err == nil {
        t.Error("FromCubicYardsPerFoot() with +Inf value should return error")
    }

    _, err = factory.FromCubicYardsPerFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicYardsPerFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicYardsPerFoot(0)
    if err != nil {
        t.Errorf("FromCubicYardsPerFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumePerLengthCubicYardPerFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicYardsPerFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromCubicYardsPerUsSurveyFoot function
func TestVolumePerLengthFactory_FromCubicYardsPerUsSurveyFoot(t *testing.T) {
    factory := units.VolumePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCubicYardsPerUsSurveyFoot(100)
    if err != nil {
        t.Errorf("FromCubicYardsPerUsSurveyFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumePerLengthCubicYardPerUsSurveyFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCubicYardsPerUsSurveyFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCubicYardsPerUsSurveyFoot(math.NaN())
    if err == nil {
        t.Error("FromCubicYardsPerUsSurveyFoot() with NaN value should return error")
    }

    _, err = factory.FromCubicYardsPerUsSurveyFoot(math.Inf(1))
    if err == nil {
        t.Error("FromCubicYardsPerUsSurveyFoot() with +Inf value should return error")
    }

    _, err = factory.FromCubicYardsPerUsSurveyFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromCubicYardsPerUsSurveyFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCubicYardsPerUsSurveyFoot(0)
    if err != nil {
        t.Errorf("FromCubicYardsPerUsSurveyFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumePerLengthCubicYardPerUsSurveyFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCubicYardsPerUsSurveyFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromUsGallonsPerMile function
func TestVolumePerLengthFactory_FromUsGallonsPerMile(t *testing.T) {
    factory := units.VolumePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsGallonsPerMile(100)
    if err != nil {
        t.Errorf("FromUsGallonsPerMile() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumePerLengthUsGallonPerMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsGallonsPerMile() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsGallonsPerMile(math.NaN())
    if err == nil {
        t.Error("FromUsGallonsPerMile() with NaN value should return error")
    }

    _, err = factory.FromUsGallonsPerMile(math.Inf(1))
    if err == nil {
        t.Error("FromUsGallonsPerMile() with +Inf value should return error")
    }

    _, err = factory.FromUsGallonsPerMile(math.Inf(-1))
    if err == nil {
        t.Error("FromUsGallonsPerMile() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsGallonsPerMile(0)
    if err != nil {
        t.Errorf("FromUsGallonsPerMile() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumePerLengthUsGallonPerMile)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsGallonsPerMile() with zero value = %v, want 0", converted)
    }
}
// Test FromImperialGallonsPerMile function
func TestVolumePerLengthFactory_FromImperialGallonsPerMile(t *testing.T) {
    factory := units.VolumePerLengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromImperialGallonsPerMile(100)
    if err != nil {
        t.Errorf("FromImperialGallonsPerMile() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.VolumePerLengthImperialGallonPerMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromImperialGallonsPerMile() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromImperialGallonsPerMile(math.NaN())
    if err == nil {
        t.Error("FromImperialGallonsPerMile() with NaN value should return error")
    }

    _, err = factory.FromImperialGallonsPerMile(math.Inf(1))
    if err == nil {
        t.Error("FromImperialGallonsPerMile() with +Inf value should return error")
    }

    _, err = factory.FromImperialGallonsPerMile(math.Inf(-1))
    if err == nil {
        t.Error("FromImperialGallonsPerMile() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromImperialGallonsPerMile(0)
    if err != nil {
        t.Errorf("FromImperialGallonsPerMile() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.VolumePerLengthImperialGallonPerMile)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromImperialGallonsPerMile() with zero value = %v, want 0", converted)
    }
}

func TestVolumePerLengthToString(t *testing.T) {
	factory := units.VolumePerLengthFactory{}
	a, err := factory.CreateVolumePerLength(45, units.VolumePerLengthCubicMeterPerMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.VolumePerLengthCubicMeterPerMeter, 2)
	expected := "45.00 " + units.GetVolumePerLengthAbbreviation(units.VolumePerLengthCubicMeterPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.VolumePerLengthCubicMeterPerMeter, -1)
	expected = "45 " + units.GetVolumePerLengthAbbreviation(units.VolumePerLengthCubicMeterPerMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestVolumePerLength_EqualityAndComparison(t *testing.T) {
	factory := units.VolumePerLengthFactory{}
	a1, _ := factory.CreateVolumePerLength(60, units.VolumePerLengthCubicMeterPerMeter)
	a2, _ := factory.CreateVolumePerLength(60, units.VolumePerLengthCubicMeterPerMeter)
	a3, _ := factory.CreateVolumePerLength(120, units.VolumePerLengthCubicMeterPerMeter)

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

func TestVolumePerLength_Arithmetic(t *testing.T) {
	factory := units.VolumePerLengthFactory{}
	a1, _ := factory.CreateVolumePerLength(30, units.VolumePerLengthCubicMeterPerMeter)
	a2, _ := factory.CreateVolumePerLength(45, units.VolumePerLengthCubicMeterPerMeter)

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


func TestGetVolumePerLengthAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.VolumePerLengthUnits
        want string
    }{
        {
            name: "CubicMeterPerMeter abbreviation",
            unit: units.VolumePerLengthCubicMeterPerMeter,
            want: "m³/m",
        },
        {
            name: "LiterPerMeter abbreviation",
            unit: units.VolumePerLengthLiterPerMeter,
            want: "l/m",
        },
        {
            name: "LiterPerKilometer abbreviation",
            unit: units.VolumePerLengthLiterPerKilometer,
            want: "l/km",
        },
        {
            name: "LiterPerMillimeter abbreviation",
            unit: units.VolumePerLengthLiterPerMillimeter,
            want: "l/mm",
        },
        {
            name: "OilBarrelPerFoot abbreviation",
            unit: units.VolumePerLengthOilBarrelPerFoot,
            want: "bbl/ft",
        },
        {
            name: "CubicYardPerFoot abbreviation",
            unit: units.VolumePerLengthCubicYardPerFoot,
            want: "yd³/ft",
        },
        {
            name: "CubicYardPerUsSurveyFoot abbreviation",
            unit: units.VolumePerLengthCubicYardPerUsSurveyFoot,
            want: "yd³/ftUS",
        },
        {
            name: "UsGallonPerMile abbreviation",
            unit: units.VolumePerLengthUsGallonPerMile,
            want: "gal (U.S.)/mi",
        },
        {
            name: "ImperialGallonPerMile abbreviation",
            unit: units.VolumePerLengthImperialGallonPerMile,
            want: "gal (imp.)/mi",
        },
        {
            name: "invalid unit",
            unit: units.VolumePerLengthUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetVolumePerLengthAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetVolumePerLengthAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestVolumePerLength_String(t *testing.T) {
    factory := units.VolumePerLengthFactory{}
    
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
            unit, err := factory.CreateVolumePerLength(tt.value, units.VolumePerLengthCubicMeterPerMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("VolumePerLength.String() = %v, want %v", got, tt.want)
            }
        })
    }
}


func TestVolumePerLength_BrokenCreation(t *testing.T) {
	// Create a factory instance
	uf := units.VolumePerLengthFactory{}

	_, err := uf.CreateVolumePerLength(100, "unknown")

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if strings.Contains(err.Error(), "unknown unit") == false {
		t.Errorf("Expected error message to contain 'Unknown unit'")
	}
}