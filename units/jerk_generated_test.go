// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestJerkDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "MeterPerSecondCubed"}`
	
	factory := units.JerkDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.JerkMeterPerSecondCubed {
		t.Errorf("expected unit %v, got %v", units.JerkMeterPerSecondCubed, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "MeterPerSecondCubed"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestJerkDto_ToJSON(t *testing.T) {
	dto := units.JerkDto{
		Value: 45,
		Unit:  units.JerkMeterPerSecondCubed,
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
	if result["unit"].(string) != string(units.JerkMeterPerSecondCubed) {
		t.Errorf("expected unit %s, got %v", units.JerkMeterPerSecondCubed, result["unit"])
	}
}

func TestNewJerk_InvalidValue(t *testing.T) {
	factory := units.JerkFactory{}
	// NaN value should return an error.
	_, err := factory.CreateJerk(math.NaN(), units.JerkMeterPerSecondCubed)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateJerk(math.Inf(1), units.JerkMeterPerSecondCubed)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestJerkConversions(t *testing.T) {
	factory := units.JerkFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateJerk(180, units.JerkMeterPerSecondCubed)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to MetersPerSecondCubed.
		// No expected conversion value provided for MetersPerSecondCubed, verifying result is not NaN.
		result := a.MetersPerSecondCubed()
		cacheResult := a.MetersPerSecondCubed()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MetersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to InchesPerSecondCubed.
		// No expected conversion value provided for InchesPerSecondCubed, verifying result is not NaN.
		result := a.InchesPerSecondCubed()
		cacheResult := a.InchesPerSecondCubed()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to InchesPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to FeetPerSecondCubed.
		// No expected conversion value provided for FeetPerSecondCubed, verifying result is not NaN.
		result := a.FeetPerSecondCubed()
		cacheResult := a.FeetPerSecondCubed()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to FeetPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to StandardGravitiesPerSecond.
		// No expected conversion value provided for StandardGravitiesPerSecond, verifying result is not NaN.
		result := a.StandardGravitiesPerSecond()
		cacheResult := a.StandardGravitiesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to StandardGravitiesPerSecond returned NaN")
		}
	}
	{
		// Test conversion to NanometersPerSecondCubed.
		// No expected conversion value provided for NanometersPerSecondCubed, verifying result is not NaN.
		result := a.NanometersPerSecondCubed()
		cacheResult := a.NanometersPerSecondCubed()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NanometersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to MicrometersPerSecondCubed.
		// No expected conversion value provided for MicrometersPerSecondCubed, verifying result is not NaN.
		result := a.MicrometersPerSecondCubed()
		cacheResult := a.MicrometersPerSecondCubed()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MicrometersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to MillimetersPerSecondCubed.
		// No expected conversion value provided for MillimetersPerSecondCubed, verifying result is not NaN.
		result := a.MillimetersPerSecondCubed()
		cacheResult := a.MillimetersPerSecondCubed()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillimetersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to CentimetersPerSecondCubed.
		// No expected conversion value provided for CentimetersPerSecondCubed, verifying result is not NaN.
		result := a.CentimetersPerSecondCubed()
		cacheResult := a.CentimetersPerSecondCubed()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to CentimetersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to DecimetersPerSecondCubed.
		// No expected conversion value provided for DecimetersPerSecondCubed, verifying result is not NaN.
		result := a.DecimetersPerSecondCubed()
		cacheResult := a.DecimetersPerSecondCubed()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DecimetersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to KilometersPerSecondCubed.
		// No expected conversion value provided for KilometersPerSecondCubed, verifying result is not NaN.
		result := a.KilometersPerSecondCubed()
		cacheResult := a.KilometersPerSecondCubed()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilometersPerSecondCubed returned NaN")
		}
	}
	{
		// Test conversion to MillistandardGravitiesPerSecond.
		// No expected conversion value provided for MillistandardGravitiesPerSecond, verifying result is not NaN.
		result := a.MillistandardGravitiesPerSecond()
		cacheResult := a.MillistandardGravitiesPerSecond()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MillistandardGravitiesPerSecond returned NaN")
		}
	}
}

func TestJerk_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.JerkFactory{}
	a, err := factory.CreateJerk(90, units.JerkMeterPerSecondCubed)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.JerkMeterPerSecondCubed {
		t.Errorf("expected default unit MeterPerSecondCubed, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.JerkMeterPerSecondCubed
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.JerkDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.JerkMeterPerSecondCubed {
		t.Errorf("expected unit MeterPerSecondCubed, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestJerkFactory_FromDto(t *testing.T) {
    factory := units.JerkFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.JerkDto{
        Value: 100,
        Unit:  units.JerkMeterPerSecondCubed,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.JerkDto{
        Value: math.NaN(),
        Unit:  units.JerkMeterPerSecondCubed,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test MeterPerSecondCubed conversion
    meters_per_second_cubedDto := units.JerkDto{
        Value: 100,
        Unit:  units.JerkMeterPerSecondCubed,
    }
    
    var meters_per_second_cubedResult *units.Jerk
    meters_per_second_cubedResult, err = factory.FromDto(meters_per_second_cubedDto)
    if err != nil {
        t.Errorf("FromDto() with MeterPerSecondCubed returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_per_second_cubedResult.Convert(units.JerkMeterPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test InchPerSecondCubed conversion
    inches_per_second_cubedDto := units.JerkDto{
        Value: 100,
        Unit:  units.JerkInchPerSecondCubed,
    }
    
    var inches_per_second_cubedResult *units.Jerk
    inches_per_second_cubedResult, err = factory.FromDto(inches_per_second_cubedDto)
    if err != nil {
        t.Errorf("FromDto() with InchPerSecondCubed returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_per_second_cubedResult.Convert(units.JerkInchPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test FootPerSecondCubed conversion
    feet_per_second_cubedDto := units.JerkDto{
        Value: 100,
        Unit:  units.JerkFootPerSecondCubed,
    }
    
    var feet_per_second_cubedResult *units.Jerk
    feet_per_second_cubedResult, err = factory.FromDto(feet_per_second_cubedDto)
    if err != nil {
        t.Errorf("FromDto() with FootPerSecondCubed returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_per_second_cubedResult.Convert(units.JerkFootPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test StandardGravitiesPerSecond conversion
    standard_gravities_per_secondDto := units.JerkDto{
        Value: 100,
        Unit:  units.JerkStandardGravitiesPerSecond,
    }
    
    var standard_gravities_per_secondResult *units.Jerk
    standard_gravities_per_secondResult, err = factory.FromDto(standard_gravities_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with StandardGravitiesPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_gravities_per_secondResult.Convert(units.JerkStandardGravitiesPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardGravitiesPerSecond = %v, want %v", converted, 100)
    }
    // Test NanometerPerSecondCubed conversion
    nanometers_per_second_cubedDto := units.JerkDto{
        Value: 100,
        Unit:  units.JerkNanometerPerSecondCubed,
    }
    
    var nanometers_per_second_cubedResult *units.Jerk
    nanometers_per_second_cubedResult, err = factory.FromDto(nanometers_per_second_cubedDto)
    if err != nil {
        t.Errorf("FromDto() with NanometerPerSecondCubed returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanometers_per_second_cubedResult.Convert(units.JerkNanometerPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanometerPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test MicrometerPerSecondCubed conversion
    micrometers_per_second_cubedDto := units.JerkDto{
        Value: 100,
        Unit:  units.JerkMicrometerPerSecondCubed,
    }
    
    var micrometers_per_second_cubedResult *units.Jerk
    micrometers_per_second_cubedResult, err = factory.FromDto(micrometers_per_second_cubedDto)
    if err != nil {
        t.Errorf("FromDto() with MicrometerPerSecondCubed returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrometers_per_second_cubedResult.Convert(units.JerkMicrometerPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrometerPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test MillimeterPerSecondCubed conversion
    millimeters_per_second_cubedDto := units.JerkDto{
        Value: 100,
        Unit:  units.JerkMillimeterPerSecondCubed,
    }
    
    var millimeters_per_second_cubedResult *units.Jerk
    millimeters_per_second_cubedResult, err = factory.FromDto(millimeters_per_second_cubedDto)
    if err != nil {
        t.Errorf("FromDto() with MillimeterPerSecondCubed returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_per_second_cubedResult.Convert(units.JerkMillimeterPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test CentimeterPerSecondCubed conversion
    centimeters_per_second_cubedDto := units.JerkDto{
        Value: 100,
        Unit:  units.JerkCentimeterPerSecondCubed,
    }
    
    var centimeters_per_second_cubedResult *units.Jerk
    centimeters_per_second_cubedResult, err = factory.FromDto(centimeters_per_second_cubedDto)
    if err != nil {
        t.Errorf("FromDto() with CentimeterPerSecondCubed returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_per_second_cubedResult.Convert(units.JerkCentimeterPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test DecimeterPerSecondCubed conversion
    decimeters_per_second_cubedDto := units.JerkDto{
        Value: 100,
        Unit:  units.JerkDecimeterPerSecondCubed,
    }
    
    var decimeters_per_second_cubedResult *units.Jerk
    decimeters_per_second_cubedResult, err = factory.FromDto(decimeters_per_second_cubedDto)
    if err != nil {
        t.Errorf("FromDto() with DecimeterPerSecondCubed returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimeters_per_second_cubedResult.Convert(units.JerkDecimeterPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimeterPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test KilometerPerSecondCubed conversion
    kilometers_per_second_cubedDto := units.JerkDto{
        Value: 100,
        Unit:  units.JerkKilometerPerSecondCubed,
    }
    
    var kilometers_per_second_cubedResult *units.Jerk
    kilometers_per_second_cubedResult, err = factory.FromDto(kilometers_per_second_cubedDto)
    if err != nil {
        t.Errorf("FromDto() with KilometerPerSecondCubed returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometers_per_second_cubedResult.Convert(units.JerkKilometerPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilometerPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test MillistandardGravitiesPerSecond conversion
    millistandard_gravities_per_secondDto := units.JerkDto{
        Value: 100,
        Unit:  units.JerkMillistandardGravitiesPerSecond,
    }
    
    var millistandard_gravities_per_secondResult *units.Jerk
    millistandard_gravities_per_secondResult, err = factory.FromDto(millistandard_gravities_per_secondDto)
    if err != nil {
        t.Errorf("FromDto() with MillistandardGravitiesPerSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millistandard_gravities_per_secondResult.Convert(units.JerkMillistandardGravitiesPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillistandardGravitiesPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.JerkDto{
        Value: 0,
        Unit:  units.JerkMeterPerSecondCubed,
    }
    
    var zeroResult *units.Jerk
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestJerkFactory_FromDtoJSON(t *testing.T) {
    factory := units.JerkFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "MeterPerSecondCubed"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "MeterPerSecondCubed"}`)
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
    nanJSON, _ := json.Marshal(units.JerkDto{
        Value: nanValue,
        Unit:  units.JerkMeterPerSecondCubed,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with MeterPerSecondCubed unit
    meters_per_second_cubedJSON := []byte(`{"value": 100, "unit": "MeterPerSecondCubed"}`)
    meters_per_second_cubedResult, err := factory.FromDtoJSON(meters_per_second_cubedJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MeterPerSecondCubed unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = meters_per_second_cubedResult.Convert(units.JerkMeterPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MeterPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test JSON with InchPerSecondCubed unit
    inches_per_second_cubedJSON := []byte(`{"value": 100, "unit": "InchPerSecondCubed"}`)
    inches_per_second_cubedResult, err := factory.FromDtoJSON(inches_per_second_cubedJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with InchPerSecondCubed unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inches_per_second_cubedResult.Convert(units.JerkInchPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for InchPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test JSON with FootPerSecondCubed unit
    feet_per_second_cubedJSON := []byte(`{"value": 100, "unit": "FootPerSecondCubed"}`)
    feet_per_second_cubedResult, err := factory.FromDtoJSON(feet_per_second_cubedJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with FootPerSecondCubed unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feet_per_second_cubedResult.Convert(units.JerkFootPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for FootPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test JSON with StandardGravitiesPerSecond unit
    standard_gravities_per_secondJSON := []byte(`{"value": 100, "unit": "StandardGravitiesPerSecond"}`)
    standard_gravities_per_secondResult, err := factory.FromDtoJSON(standard_gravities_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with StandardGravitiesPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = standard_gravities_per_secondResult.Convert(units.JerkStandardGravitiesPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for StandardGravitiesPerSecond = %v, want %v", converted, 100)
    }
    // Test JSON with NanometerPerSecondCubed unit
    nanometers_per_second_cubedJSON := []byte(`{"value": 100, "unit": "NanometerPerSecondCubed"}`)
    nanometers_per_second_cubedResult, err := factory.FromDtoJSON(nanometers_per_second_cubedJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanometerPerSecondCubed unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanometers_per_second_cubedResult.Convert(units.JerkNanometerPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanometerPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test JSON with MicrometerPerSecondCubed unit
    micrometers_per_second_cubedJSON := []byte(`{"value": 100, "unit": "MicrometerPerSecondCubed"}`)
    micrometers_per_second_cubedResult, err := factory.FromDtoJSON(micrometers_per_second_cubedJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrometerPerSecondCubed unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrometers_per_second_cubedResult.Convert(units.JerkMicrometerPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrometerPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test JSON with MillimeterPerSecondCubed unit
    millimeters_per_second_cubedJSON := []byte(`{"value": 100, "unit": "MillimeterPerSecondCubed"}`)
    millimeters_per_second_cubedResult, err := factory.FromDtoJSON(millimeters_per_second_cubedJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillimeterPerSecondCubed unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimeters_per_second_cubedResult.Convert(units.JerkMillimeterPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillimeterPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test JSON with CentimeterPerSecondCubed unit
    centimeters_per_second_cubedJSON := []byte(`{"value": 100, "unit": "CentimeterPerSecondCubed"}`)
    centimeters_per_second_cubedResult, err := factory.FromDtoJSON(centimeters_per_second_cubedJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with CentimeterPerSecondCubed unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimeters_per_second_cubedResult.Convert(units.JerkCentimeterPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for CentimeterPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test JSON with DecimeterPerSecondCubed unit
    decimeters_per_second_cubedJSON := []byte(`{"value": 100, "unit": "DecimeterPerSecondCubed"}`)
    decimeters_per_second_cubedResult, err := factory.FromDtoJSON(decimeters_per_second_cubedJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DecimeterPerSecondCubed unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimeters_per_second_cubedResult.Convert(units.JerkDecimeterPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DecimeterPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test JSON with KilometerPerSecondCubed unit
    kilometers_per_second_cubedJSON := []byte(`{"value": 100, "unit": "KilometerPerSecondCubed"}`)
    kilometers_per_second_cubedResult, err := factory.FromDtoJSON(kilometers_per_second_cubedJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilometerPerSecondCubed unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometers_per_second_cubedResult.Convert(units.JerkKilometerPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilometerPerSecondCubed = %v, want %v", converted, 100)
    }
    // Test JSON with MillistandardGravitiesPerSecond unit
    millistandard_gravities_per_secondJSON := []byte(`{"value": 100, "unit": "MillistandardGravitiesPerSecond"}`)
    millistandard_gravities_per_secondResult, err := factory.FromDtoJSON(millistandard_gravities_per_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillistandardGravitiesPerSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millistandard_gravities_per_secondResult.Convert(units.JerkMillistandardGravitiesPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillistandardGravitiesPerSecond = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "MeterPerSecondCubed"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromMetersPerSecondCubed function
func TestJerkFactory_FromMetersPerSecondCubed(t *testing.T) {
    factory := units.JerkFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMetersPerSecondCubed(100)
    if err != nil {
        t.Errorf("FromMetersPerSecondCubed() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.JerkMeterPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMetersPerSecondCubed() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMetersPerSecondCubed(math.NaN())
    if err == nil {
        t.Error("FromMetersPerSecondCubed() with NaN value should return error")
    }

    _, err = factory.FromMetersPerSecondCubed(math.Inf(1))
    if err == nil {
        t.Error("FromMetersPerSecondCubed() with +Inf value should return error")
    }

    _, err = factory.FromMetersPerSecondCubed(math.Inf(-1))
    if err == nil {
        t.Error("FromMetersPerSecondCubed() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMetersPerSecondCubed(0)
    if err != nil {
        t.Errorf("FromMetersPerSecondCubed() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.JerkMeterPerSecondCubed)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMetersPerSecondCubed() with zero value = %v, want 0", converted)
    }
}
// Test FromInchesPerSecondCubed function
func TestJerkFactory_FromInchesPerSecondCubed(t *testing.T) {
    factory := units.JerkFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInchesPerSecondCubed(100)
    if err != nil {
        t.Errorf("FromInchesPerSecondCubed() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.JerkInchPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInchesPerSecondCubed() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInchesPerSecondCubed(math.NaN())
    if err == nil {
        t.Error("FromInchesPerSecondCubed() with NaN value should return error")
    }

    _, err = factory.FromInchesPerSecondCubed(math.Inf(1))
    if err == nil {
        t.Error("FromInchesPerSecondCubed() with +Inf value should return error")
    }

    _, err = factory.FromInchesPerSecondCubed(math.Inf(-1))
    if err == nil {
        t.Error("FromInchesPerSecondCubed() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInchesPerSecondCubed(0)
    if err != nil {
        t.Errorf("FromInchesPerSecondCubed() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.JerkInchPerSecondCubed)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInchesPerSecondCubed() with zero value = %v, want 0", converted)
    }
}
// Test FromFeetPerSecondCubed function
func TestJerkFactory_FromFeetPerSecondCubed(t *testing.T) {
    factory := units.JerkFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFeetPerSecondCubed(100)
    if err != nil {
        t.Errorf("FromFeetPerSecondCubed() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.JerkFootPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFeetPerSecondCubed() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFeetPerSecondCubed(math.NaN())
    if err == nil {
        t.Error("FromFeetPerSecondCubed() with NaN value should return error")
    }

    _, err = factory.FromFeetPerSecondCubed(math.Inf(1))
    if err == nil {
        t.Error("FromFeetPerSecondCubed() with +Inf value should return error")
    }

    _, err = factory.FromFeetPerSecondCubed(math.Inf(-1))
    if err == nil {
        t.Error("FromFeetPerSecondCubed() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFeetPerSecondCubed(0)
    if err != nil {
        t.Errorf("FromFeetPerSecondCubed() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.JerkFootPerSecondCubed)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFeetPerSecondCubed() with zero value = %v, want 0", converted)
    }
}
// Test FromStandardGravitiesPerSecond function
func TestJerkFactory_FromStandardGravitiesPerSecond(t *testing.T) {
    factory := units.JerkFactory{}
    var err error

    // Test valid value
    result, err := factory.FromStandardGravitiesPerSecond(100)
    if err != nil {
        t.Errorf("FromStandardGravitiesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.JerkStandardGravitiesPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromStandardGravitiesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromStandardGravitiesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromStandardGravitiesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromStandardGravitiesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromStandardGravitiesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromStandardGravitiesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromStandardGravitiesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromStandardGravitiesPerSecond(0)
    if err != nil {
        t.Errorf("FromStandardGravitiesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.JerkStandardGravitiesPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromStandardGravitiesPerSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromNanometersPerSecondCubed function
func TestJerkFactory_FromNanometersPerSecondCubed(t *testing.T) {
    factory := units.JerkFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanometersPerSecondCubed(100)
    if err != nil {
        t.Errorf("FromNanometersPerSecondCubed() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.JerkNanometerPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanometersPerSecondCubed() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanometersPerSecondCubed(math.NaN())
    if err == nil {
        t.Error("FromNanometersPerSecondCubed() with NaN value should return error")
    }

    _, err = factory.FromNanometersPerSecondCubed(math.Inf(1))
    if err == nil {
        t.Error("FromNanometersPerSecondCubed() with +Inf value should return error")
    }

    _, err = factory.FromNanometersPerSecondCubed(math.Inf(-1))
    if err == nil {
        t.Error("FromNanometersPerSecondCubed() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanometersPerSecondCubed(0)
    if err != nil {
        t.Errorf("FromNanometersPerSecondCubed() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.JerkNanometerPerSecondCubed)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanometersPerSecondCubed() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrometersPerSecondCubed function
func TestJerkFactory_FromMicrometersPerSecondCubed(t *testing.T) {
    factory := units.JerkFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrometersPerSecondCubed(100)
    if err != nil {
        t.Errorf("FromMicrometersPerSecondCubed() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.JerkMicrometerPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrometersPerSecondCubed() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrometersPerSecondCubed(math.NaN())
    if err == nil {
        t.Error("FromMicrometersPerSecondCubed() with NaN value should return error")
    }

    _, err = factory.FromMicrometersPerSecondCubed(math.Inf(1))
    if err == nil {
        t.Error("FromMicrometersPerSecondCubed() with +Inf value should return error")
    }

    _, err = factory.FromMicrometersPerSecondCubed(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrometersPerSecondCubed() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrometersPerSecondCubed(0)
    if err != nil {
        t.Errorf("FromMicrometersPerSecondCubed() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.JerkMicrometerPerSecondCubed)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrometersPerSecondCubed() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimetersPerSecondCubed function
func TestJerkFactory_FromMillimetersPerSecondCubed(t *testing.T) {
    factory := units.JerkFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimetersPerSecondCubed(100)
    if err != nil {
        t.Errorf("FromMillimetersPerSecondCubed() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.JerkMillimeterPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimetersPerSecondCubed() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimetersPerSecondCubed(math.NaN())
    if err == nil {
        t.Error("FromMillimetersPerSecondCubed() with NaN value should return error")
    }

    _, err = factory.FromMillimetersPerSecondCubed(math.Inf(1))
    if err == nil {
        t.Error("FromMillimetersPerSecondCubed() with +Inf value should return error")
    }

    _, err = factory.FromMillimetersPerSecondCubed(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimetersPerSecondCubed() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimetersPerSecondCubed(0)
    if err != nil {
        t.Errorf("FromMillimetersPerSecondCubed() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.JerkMillimeterPerSecondCubed)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimetersPerSecondCubed() with zero value = %v, want 0", converted)
    }
}
// Test FromCentimetersPerSecondCubed function
func TestJerkFactory_FromCentimetersPerSecondCubed(t *testing.T) {
    factory := units.JerkFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentimetersPerSecondCubed(100)
    if err != nil {
        t.Errorf("FromCentimetersPerSecondCubed() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.JerkCentimeterPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentimetersPerSecondCubed() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentimetersPerSecondCubed(math.NaN())
    if err == nil {
        t.Error("FromCentimetersPerSecondCubed() with NaN value should return error")
    }

    _, err = factory.FromCentimetersPerSecondCubed(math.Inf(1))
    if err == nil {
        t.Error("FromCentimetersPerSecondCubed() with +Inf value should return error")
    }

    _, err = factory.FromCentimetersPerSecondCubed(math.Inf(-1))
    if err == nil {
        t.Error("FromCentimetersPerSecondCubed() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentimetersPerSecondCubed(0)
    if err != nil {
        t.Errorf("FromCentimetersPerSecondCubed() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.JerkCentimeterPerSecondCubed)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentimetersPerSecondCubed() with zero value = %v, want 0", converted)
    }
}
// Test FromDecimetersPerSecondCubed function
func TestJerkFactory_FromDecimetersPerSecondCubed(t *testing.T) {
    factory := units.JerkFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecimetersPerSecondCubed(100)
    if err != nil {
        t.Errorf("FromDecimetersPerSecondCubed() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.JerkDecimeterPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecimetersPerSecondCubed() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecimetersPerSecondCubed(math.NaN())
    if err == nil {
        t.Error("FromDecimetersPerSecondCubed() with NaN value should return error")
    }

    _, err = factory.FromDecimetersPerSecondCubed(math.Inf(1))
    if err == nil {
        t.Error("FromDecimetersPerSecondCubed() with +Inf value should return error")
    }

    _, err = factory.FromDecimetersPerSecondCubed(math.Inf(-1))
    if err == nil {
        t.Error("FromDecimetersPerSecondCubed() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecimetersPerSecondCubed(0)
    if err != nil {
        t.Errorf("FromDecimetersPerSecondCubed() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.JerkDecimeterPerSecondCubed)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecimetersPerSecondCubed() with zero value = %v, want 0", converted)
    }
}
// Test FromKilometersPerSecondCubed function
func TestJerkFactory_FromKilometersPerSecondCubed(t *testing.T) {
    factory := units.JerkFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilometersPerSecondCubed(100)
    if err != nil {
        t.Errorf("FromKilometersPerSecondCubed() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.JerkKilometerPerSecondCubed)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilometersPerSecondCubed() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilometersPerSecondCubed(math.NaN())
    if err == nil {
        t.Error("FromKilometersPerSecondCubed() with NaN value should return error")
    }

    _, err = factory.FromKilometersPerSecondCubed(math.Inf(1))
    if err == nil {
        t.Error("FromKilometersPerSecondCubed() with +Inf value should return error")
    }

    _, err = factory.FromKilometersPerSecondCubed(math.Inf(-1))
    if err == nil {
        t.Error("FromKilometersPerSecondCubed() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilometersPerSecondCubed(0)
    if err != nil {
        t.Errorf("FromKilometersPerSecondCubed() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.JerkKilometerPerSecondCubed)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilometersPerSecondCubed() with zero value = %v, want 0", converted)
    }
}
// Test FromMillistandardGravitiesPerSecond function
func TestJerkFactory_FromMillistandardGravitiesPerSecond(t *testing.T) {
    factory := units.JerkFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillistandardGravitiesPerSecond(100)
    if err != nil {
        t.Errorf("FromMillistandardGravitiesPerSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.JerkMillistandardGravitiesPerSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillistandardGravitiesPerSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillistandardGravitiesPerSecond(math.NaN())
    if err == nil {
        t.Error("FromMillistandardGravitiesPerSecond() with NaN value should return error")
    }

    _, err = factory.FromMillistandardGravitiesPerSecond(math.Inf(1))
    if err == nil {
        t.Error("FromMillistandardGravitiesPerSecond() with +Inf value should return error")
    }

    _, err = factory.FromMillistandardGravitiesPerSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromMillistandardGravitiesPerSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillistandardGravitiesPerSecond(0)
    if err != nil {
        t.Errorf("FromMillistandardGravitiesPerSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.JerkMillistandardGravitiesPerSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillistandardGravitiesPerSecond() with zero value = %v, want 0", converted)
    }
}

func TestJerkToString(t *testing.T) {
	factory := units.JerkFactory{}
	a, err := factory.CreateJerk(45, units.JerkMeterPerSecondCubed)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.JerkMeterPerSecondCubed, 2)
	expected := "45.00 " + units.GetJerkAbbreviation(units.JerkMeterPerSecondCubed)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.JerkMeterPerSecondCubed, -1)
	expected = "45 " + units.GetJerkAbbreviation(units.JerkMeterPerSecondCubed)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestJerk_EqualityAndComparison(t *testing.T) {
	factory := units.JerkFactory{}
	a1, _ := factory.CreateJerk(60, units.JerkMeterPerSecondCubed)
	a2, _ := factory.CreateJerk(60, units.JerkMeterPerSecondCubed)
	a3, _ := factory.CreateJerk(120, units.JerkMeterPerSecondCubed)

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

func TestJerk_Arithmetic(t *testing.T) {
	factory := units.JerkFactory{}
	a1, _ := factory.CreateJerk(30, units.JerkMeterPerSecondCubed)
	a2, _ := factory.CreateJerk(45, units.JerkMeterPerSecondCubed)

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


func TestGetJerkAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.JerkUnits
        want string
    }{
        {
            name: "MeterPerSecondCubed abbreviation",
            unit: units.JerkMeterPerSecondCubed,
            want: "m/s³",
        },
        {
            name: "InchPerSecondCubed abbreviation",
            unit: units.JerkInchPerSecondCubed,
            want: "in/s³",
        },
        {
            name: "FootPerSecondCubed abbreviation",
            unit: units.JerkFootPerSecondCubed,
            want: "ft/s³",
        },
        {
            name: "StandardGravitiesPerSecond abbreviation",
            unit: units.JerkStandardGravitiesPerSecond,
            want: "g/s",
        },
        {
            name: "NanometerPerSecondCubed abbreviation",
            unit: units.JerkNanometerPerSecondCubed,
            want: "nm/s³",
        },
        {
            name: "MicrometerPerSecondCubed abbreviation",
            unit: units.JerkMicrometerPerSecondCubed,
            want: "μm/s³",
        },
        {
            name: "MillimeterPerSecondCubed abbreviation",
            unit: units.JerkMillimeterPerSecondCubed,
            want: "mm/s³",
        },
        {
            name: "CentimeterPerSecondCubed abbreviation",
            unit: units.JerkCentimeterPerSecondCubed,
            want: "cm/s³",
        },
        {
            name: "DecimeterPerSecondCubed abbreviation",
            unit: units.JerkDecimeterPerSecondCubed,
            want: "dm/s³",
        },
        {
            name: "KilometerPerSecondCubed abbreviation",
            unit: units.JerkKilometerPerSecondCubed,
            want: "km/s³",
        },
        {
            name: "MillistandardGravitiesPerSecond abbreviation",
            unit: units.JerkMillistandardGravitiesPerSecond,
            want: "mg/s",
        },
        {
            name: "invalid unit",
            unit: units.JerkUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetJerkAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetJerkAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestJerk_String(t *testing.T) {
    factory := units.JerkFactory{}
    
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
            unit, err := factory.CreateJerk(tt.value, units.JerkMeterPerSecondCubed)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Jerk.String() = %v, want %v", got, tt.want)
            }
        })
    }
}