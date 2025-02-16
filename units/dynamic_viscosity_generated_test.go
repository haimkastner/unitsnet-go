// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestDynamicViscosityDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "NewtonSecondPerMeterSquared"}`
	
	factory := units.DynamicViscosityDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.DynamicViscosityNewtonSecondPerMeterSquared {
		t.Errorf("expected unit %v, got %v", units.DynamicViscosityNewtonSecondPerMeterSquared, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "NewtonSecondPerMeterSquared"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestDynamicViscosityDto_ToJSON(t *testing.T) {
	dto := units.DynamicViscosityDto{
		Value: 45,
		Unit:  units.DynamicViscosityNewtonSecondPerMeterSquared,
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
	if result["unit"].(string) != string(units.DynamicViscosityNewtonSecondPerMeterSquared) {
		t.Errorf("expected unit %s, got %v", units.DynamicViscosityNewtonSecondPerMeterSquared, result["unit"])
	}
}

func TestNewDynamicViscosity_InvalidValue(t *testing.T) {
	factory := units.DynamicViscosityFactory{}
	// NaN value should return an error.
	_, err := factory.CreateDynamicViscosity(math.NaN(), units.DynamicViscosityNewtonSecondPerMeterSquared)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateDynamicViscosity(math.Inf(1), units.DynamicViscosityNewtonSecondPerMeterSquared)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestDynamicViscosityConversions(t *testing.T) {
	factory := units.DynamicViscosityFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateDynamicViscosity(180, units.DynamicViscosityNewtonSecondPerMeterSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to NewtonSecondsPerMeterSquared.
		// No expected conversion value provided for NewtonSecondsPerMeterSquared, verifying result is not NaN.
		result := a.NewtonSecondsPerMeterSquared()
		if math.IsNaN(result) {
			t.Errorf("conversion to NewtonSecondsPerMeterSquared returned NaN")
		}
	}
	{
		// Test conversion to PascalSeconds.
		// No expected conversion value provided for PascalSeconds, verifying result is not NaN.
		result := a.PascalSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to PascalSeconds returned NaN")
		}
	}
	{
		// Test conversion to Poise.
		// No expected conversion value provided for Poise, verifying result is not NaN.
		result := a.Poise()
		if math.IsNaN(result) {
			t.Errorf("conversion to Poise returned NaN")
		}
	}
	{
		// Test conversion to Reyns.
		// No expected conversion value provided for Reyns, verifying result is not NaN.
		result := a.Reyns()
		if math.IsNaN(result) {
			t.Errorf("conversion to Reyns returned NaN")
		}
	}
	{
		// Test conversion to PoundsForceSecondPerSquareInch.
		// No expected conversion value provided for PoundsForceSecondPerSquareInch, verifying result is not NaN.
		result := a.PoundsForceSecondPerSquareInch()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForceSecondPerSquareInch returned NaN")
		}
	}
	{
		// Test conversion to PoundsForceSecondPerSquareFoot.
		// No expected conversion value provided for PoundsForceSecondPerSquareFoot, verifying result is not NaN.
		result := a.PoundsForceSecondPerSquareFoot()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsForceSecondPerSquareFoot returned NaN")
		}
	}
	{
		// Test conversion to PoundsPerFootSecond.
		// No expected conversion value provided for PoundsPerFootSecond, verifying result is not NaN.
		result := a.PoundsPerFootSecond()
		if math.IsNaN(result) {
			t.Errorf("conversion to PoundsPerFootSecond returned NaN")
		}
	}
	{
		// Test conversion to MillipascalSeconds.
		// No expected conversion value provided for MillipascalSeconds, verifying result is not NaN.
		result := a.MillipascalSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MillipascalSeconds returned NaN")
		}
	}
	{
		// Test conversion to MicropascalSeconds.
		// No expected conversion value provided for MicropascalSeconds, verifying result is not NaN.
		result := a.MicropascalSeconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicropascalSeconds returned NaN")
		}
	}
	{
		// Test conversion to Centipoise.
		// No expected conversion value provided for Centipoise, verifying result is not NaN.
		result := a.Centipoise()
		if math.IsNaN(result) {
			t.Errorf("conversion to Centipoise returned NaN")
		}
	}
}

func TestDynamicViscosity_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.DynamicViscosityFactory{}
	a, err := factory.CreateDynamicViscosity(90, units.DynamicViscosityNewtonSecondPerMeterSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.DynamicViscosityNewtonSecondPerMeterSquared {
		t.Errorf("expected default unit NewtonSecondPerMeterSquared, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.DynamicViscosityNewtonSecondPerMeterSquared
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.DynamicViscosityDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.DynamicViscosityNewtonSecondPerMeterSquared {
		t.Errorf("expected unit NewtonSecondPerMeterSquared, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestDynamicViscosityFactory_FromDto(t *testing.T) {
    factory := units.DynamicViscosityFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.DynamicViscosityDto{
        Value: 100,
        Unit:  units.DynamicViscosityNewtonSecondPerMeterSquared,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.DynamicViscosityDto{
        Value: math.NaN(),
        Unit:  units.DynamicViscosityNewtonSecondPerMeterSquared,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test NewtonSecondPerMeterSquared conversion
    newton_seconds_per_meter_squaredDto := units.DynamicViscosityDto{
        Value: 100,
        Unit:  units.DynamicViscosityNewtonSecondPerMeterSquared,
    }
    
    var newton_seconds_per_meter_squaredResult *units.DynamicViscosity
    newton_seconds_per_meter_squaredResult, err = factory.FromDto(newton_seconds_per_meter_squaredDto)
    if err != nil {
        t.Errorf("FromDto() with NewtonSecondPerMeterSquared returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_seconds_per_meter_squaredResult.Convert(units.DynamicViscosityNewtonSecondPerMeterSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonSecondPerMeterSquared = %v, want %v", converted, 100)
    }
    // Test PascalSecond conversion
    pascal_secondsDto := units.DynamicViscosityDto{
        Value: 100,
        Unit:  units.DynamicViscosityPascalSecond,
    }
    
    var pascal_secondsResult *units.DynamicViscosity
    pascal_secondsResult, err = factory.FromDto(pascal_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with PascalSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_secondsResult.Convert(units.DynamicViscosityPascalSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalSecond = %v, want %v", converted, 100)
    }
    // Test Poise conversion
    poiseDto := units.DynamicViscosityDto{
        Value: 100,
        Unit:  units.DynamicViscosityPoise,
    }
    
    var poiseResult *units.DynamicViscosity
    poiseResult, err = factory.FromDto(poiseDto)
    if err != nil {
        t.Errorf("FromDto() with Poise returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = poiseResult.Convert(units.DynamicViscosityPoise)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Poise = %v, want %v", converted, 100)
    }
    // Test Reyn conversion
    reynsDto := units.DynamicViscosityDto{
        Value: 100,
        Unit:  units.DynamicViscosityReyn,
    }
    
    var reynsResult *units.DynamicViscosity
    reynsResult, err = factory.FromDto(reynsDto)
    if err != nil {
        t.Errorf("FromDto() with Reyn returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = reynsResult.Convert(units.DynamicViscosityReyn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Reyn = %v, want %v", converted, 100)
    }
    // Test PoundForceSecondPerSquareInch conversion
    pounds_force_second_per_square_inchDto := units.DynamicViscosityDto{
        Value: 100,
        Unit:  units.DynamicViscosityPoundForceSecondPerSquareInch,
    }
    
    var pounds_force_second_per_square_inchResult *units.DynamicViscosity
    pounds_force_second_per_square_inchResult, err = factory.FromDto(pounds_force_second_per_square_inchDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForceSecondPerSquareInch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_second_per_square_inchResult.Convert(units.DynamicViscosityPoundForceSecondPerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceSecondPerSquareInch = %v, want %v", converted, 100)
    }
    // Test PoundForceSecondPerSquareFoot conversion
    pounds_force_second_per_square_footDto := units.DynamicViscosityDto{
        Value: 100,
        Unit:  units.DynamicViscosityPoundForceSecondPerSquareFoot,
    }
    
    var pounds_force_second_per_square_footResult *units.DynamicViscosity
    pounds_force_second_per_square_footResult, err = factory.FromDto(pounds_force_second_per_square_footDto)
    if err != nil {
        t.Errorf("FromDto() with PoundForceSecondPerSquareFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_second_per_square_footResult.Convert(units.DynamicViscosityPoundForceSecondPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceSecondPerSquareFoot = %v, want %v", converted, 100)
    }
    // Test PoundPerFootSecond conversion
    pounds_per_foot_secondDto := units.DynamicViscosityDto{
        Value: 100,
        Unit:  units.DynamicViscosityPoundPerFootSecond,
    }
    
    var pounds_per_foot_secondResult *units.DynamicViscosity
    pounds_per_foot_secondResult, err = factory.FromDto(pounds_per_foot_secondDto)
    if err != nil {
        t.Errorf("FromDto() with PoundPerFootSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_foot_secondResult.Convert(units.DynamicViscosityPoundPerFootSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerFootSecond = %v, want %v", converted, 100)
    }
    // Test MillipascalSecond conversion
    millipascal_secondsDto := units.DynamicViscosityDto{
        Value: 100,
        Unit:  units.DynamicViscosityMillipascalSecond,
    }
    
    var millipascal_secondsResult *units.DynamicViscosity
    millipascal_secondsResult, err = factory.FromDto(millipascal_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with MillipascalSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millipascal_secondsResult.Convert(units.DynamicViscosityMillipascalSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillipascalSecond = %v, want %v", converted, 100)
    }
    // Test MicropascalSecond conversion
    micropascal_secondsDto := units.DynamicViscosityDto{
        Value: 100,
        Unit:  units.DynamicViscosityMicropascalSecond,
    }
    
    var micropascal_secondsResult *units.DynamicViscosity
    micropascal_secondsResult, err = factory.FromDto(micropascal_secondsDto)
    if err != nil {
        t.Errorf("FromDto() with MicropascalSecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micropascal_secondsResult.Convert(units.DynamicViscosityMicropascalSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicropascalSecond = %v, want %v", converted, 100)
    }
    // Test Centipoise conversion
    centipoiseDto := units.DynamicViscosityDto{
        Value: 100,
        Unit:  units.DynamicViscosityCentipoise,
    }
    
    var centipoiseResult *units.DynamicViscosity
    centipoiseResult, err = factory.FromDto(centipoiseDto)
    if err != nil {
        t.Errorf("FromDto() with Centipoise returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centipoiseResult.Convert(units.DynamicViscosityCentipoise)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centipoise = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.DynamicViscosityDto{
        Value: 0,
        Unit:  units.DynamicViscosityNewtonSecondPerMeterSquared,
    }
    
    var zeroResult *units.DynamicViscosity
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestDynamicViscosityFactory_FromDtoJSON(t *testing.T) {
    factory := units.DynamicViscosityFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "NewtonSecondPerMeterSquared"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "NewtonSecondPerMeterSquared"}`)
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
    nanJSON, _ := json.Marshal(units.DynamicViscosityDto{
        Value: nanValue,
        Unit:  units.DynamicViscosityNewtonSecondPerMeterSquared,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with NewtonSecondPerMeterSquared unit
    newton_seconds_per_meter_squaredJSON := []byte(`{"value": 100, "unit": "NewtonSecondPerMeterSquared"}`)
    newton_seconds_per_meter_squaredResult, err := factory.FromDtoJSON(newton_seconds_per_meter_squaredJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NewtonSecondPerMeterSquared unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = newton_seconds_per_meter_squaredResult.Convert(units.DynamicViscosityNewtonSecondPerMeterSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NewtonSecondPerMeterSquared = %v, want %v", converted, 100)
    }
    // Test JSON with PascalSecond unit
    pascal_secondsJSON := []byte(`{"value": 100, "unit": "PascalSecond"}`)
    pascal_secondsResult, err := factory.FromDtoJSON(pascal_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PascalSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pascal_secondsResult.Convert(units.DynamicViscosityPascalSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PascalSecond = %v, want %v", converted, 100)
    }
    // Test JSON with Poise unit
    poiseJSON := []byte(`{"value": 100, "unit": "Poise"}`)
    poiseResult, err := factory.FromDtoJSON(poiseJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Poise unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = poiseResult.Convert(units.DynamicViscosityPoise)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Poise = %v, want %v", converted, 100)
    }
    // Test JSON with Reyn unit
    reynsJSON := []byte(`{"value": 100, "unit": "Reyn"}`)
    reynsResult, err := factory.FromDtoJSON(reynsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Reyn unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = reynsResult.Convert(units.DynamicViscosityReyn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Reyn = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForceSecondPerSquareInch unit
    pounds_force_second_per_square_inchJSON := []byte(`{"value": 100, "unit": "PoundForceSecondPerSquareInch"}`)
    pounds_force_second_per_square_inchResult, err := factory.FromDtoJSON(pounds_force_second_per_square_inchJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForceSecondPerSquareInch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_second_per_square_inchResult.Convert(units.DynamicViscosityPoundForceSecondPerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceSecondPerSquareInch = %v, want %v", converted, 100)
    }
    // Test JSON with PoundForceSecondPerSquareFoot unit
    pounds_force_second_per_square_footJSON := []byte(`{"value": 100, "unit": "PoundForceSecondPerSquareFoot"}`)
    pounds_force_second_per_square_footResult, err := factory.FromDtoJSON(pounds_force_second_per_square_footJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundForceSecondPerSquareFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_force_second_per_square_footResult.Convert(units.DynamicViscosityPoundForceSecondPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundForceSecondPerSquareFoot = %v, want %v", converted, 100)
    }
    // Test JSON with PoundPerFootSecond unit
    pounds_per_foot_secondJSON := []byte(`{"value": 100, "unit": "PoundPerFootSecond"}`)
    pounds_per_foot_secondResult, err := factory.FromDtoJSON(pounds_per_foot_secondJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PoundPerFootSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = pounds_per_foot_secondResult.Convert(units.DynamicViscosityPoundPerFootSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PoundPerFootSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MillipascalSecond unit
    millipascal_secondsJSON := []byte(`{"value": 100, "unit": "MillipascalSecond"}`)
    millipascal_secondsResult, err := factory.FromDtoJSON(millipascal_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MillipascalSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millipascal_secondsResult.Convert(units.DynamicViscosityMillipascalSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MillipascalSecond = %v, want %v", converted, 100)
    }
    // Test JSON with MicropascalSecond unit
    micropascal_secondsJSON := []byte(`{"value": 100, "unit": "MicropascalSecond"}`)
    micropascal_secondsResult, err := factory.FromDtoJSON(micropascal_secondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicropascalSecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micropascal_secondsResult.Convert(units.DynamicViscosityMicropascalSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicropascalSecond = %v, want %v", converted, 100)
    }
    // Test JSON with Centipoise unit
    centipoiseJSON := []byte(`{"value": 100, "unit": "Centipoise"}`)
    centipoiseResult, err := factory.FromDtoJSON(centipoiseJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Centipoise unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centipoiseResult.Convert(units.DynamicViscosityCentipoise)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centipoise = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "NewtonSecondPerMeterSquared"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromNewtonSecondsPerMeterSquared function
func TestDynamicViscosityFactory_FromNewtonSecondsPerMeterSquared(t *testing.T) {
    factory := units.DynamicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNewtonSecondsPerMeterSquared(100)
    if err != nil {
        t.Errorf("FromNewtonSecondsPerMeterSquared() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DynamicViscosityNewtonSecondPerMeterSquared)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNewtonSecondsPerMeterSquared() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNewtonSecondsPerMeterSquared(math.NaN())
    if err == nil {
        t.Error("FromNewtonSecondsPerMeterSquared() with NaN value should return error")
    }

    _, err = factory.FromNewtonSecondsPerMeterSquared(math.Inf(1))
    if err == nil {
        t.Error("FromNewtonSecondsPerMeterSquared() with +Inf value should return error")
    }

    _, err = factory.FromNewtonSecondsPerMeterSquared(math.Inf(-1))
    if err == nil {
        t.Error("FromNewtonSecondsPerMeterSquared() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNewtonSecondsPerMeterSquared(0)
    if err != nil {
        t.Errorf("FromNewtonSecondsPerMeterSquared() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DynamicViscosityNewtonSecondPerMeterSquared)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNewtonSecondsPerMeterSquared() with zero value = %v, want 0", converted)
    }
}
// Test FromPascalSeconds function
func TestDynamicViscosityFactory_FromPascalSeconds(t *testing.T) {
    factory := units.DynamicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPascalSeconds(100)
    if err != nil {
        t.Errorf("FromPascalSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DynamicViscosityPascalSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPascalSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPascalSeconds(math.NaN())
    if err == nil {
        t.Error("FromPascalSeconds() with NaN value should return error")
    }

    _, err = factory.FromPascalSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromPascalSeconds() with +Inf value should return error")
    }

    _, err = factory.FromPascalSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromPascalSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPascalSeconds(0)
    if err != nil {
        t.Errorf("FromPascalSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DynamicViscosityPascalSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPascalSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromPoise function
func TestDynamicViscosityFactory_FromPoise(t *testing.T) {
    factory := units.DynamicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoise(100)
    if err != nil {
        t.Errorf("FromPoise() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DynamicViscosityPoise)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoise() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoise(math.NaN())
    if err == nil {
        t.Error("FromPoise() with NaN value should return error")
    }

    _, err = factory.FromPoise(math.Inf(1))
    if err == nil {
        t.Error("FromPoise() with +Inf value should return error")
    }

    _, err = factory.FromPoise(math.Inf(-1))
    if err == nil {
        t.Error("FromPoise() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoise(0)
    if err != nil {
        t.Errorf("FromPoise() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DynamicViscosityPoise)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoise() with zero value = %v, want 0", converted)
    }
}
// Test FromReyns function
func TestDynamicViscosityFactory_FromReyns(t *testing.T) {
    factory := units.DynamicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromReyns(100)
    if err != nil {
        t.Errorf("FromReyns() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DynamicViscosityReyn)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromReyns() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromReyns(math.NaN())
    if err == nil {
        t.Error("FromReyns() with NaN value should return error")
    }

    _, err = factory.FromReyns(math.Inf(1))
    if err == nil {
        t.Error("FromReyns() with +Inf value should return error")
    }

    _, err = factory.FromReyns(math.Inf(-1))
    if err == nil {
        t.Error("FromReyns() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromReyns(0)
    if err != nil {
        t.Errorf("FromReyns() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DynamicViscosityReyn)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromReyns() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForceSecondPerSquareInch function
func TestDynamicViscosityFactory_FromPoundsForceSecondPerSquareInch(t *testing.T) {
    factory := units.DynamicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForceSecondPerSquareInch(100)
    if err != nil {
        t.Errorf("FromPoundsForceSecondPerSquareInch() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DynamicViscosityPoundForceSecondPerSquareInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForceSecondPerSquareInch() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForceSecondPerSquareInch(math.NaN())
    if err == nil {
        t.Error("FromPoundsForceSecondPerSquareInch() with NaN value should return error")
    }

    _, err = factory.FromPoundsForceSecondPerSquareInch(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForceSecondPerSquareInch() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForceSecondPerSquareInch(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForceSecondPerSquareInch() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForceSecondPerSquareInch(0)
    if err != nil {
        t.Errorf("FromPoundsForceSecondPerSquareInch() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DynamicViscosityPoundForceSecondPerSquareInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForceSecondPerSquareInch() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsForceSecondPerSquareFoot function
func TestDynamicViscosityFactory_FromPoundsForceSecondPerSquareFoot(t *testing.T) {
    factory := units.DynamicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsForceSecondPerSquareFoot(100)
    if err != nil {
        t.Errorf("FromPoundsForceSecondPerSquareFoot() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DynamicViscosityPoundForceSecondPerSquareFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsForceSecondPerSquareFoot() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsForceSecondPerSquareFoot(math.NaN())
    if err == nil {
        t.Error("FromPoundsForceSecondPerSquareFoot() with NaN value should return error")
    }

    _, err = factory.FromPoundsForceSecondPerSquareFoot(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsForceSecondPerSquareFoot() with +Inf value should return error")
    }

    _, err = factory.FromPoundsForceSecondPerSquareFoot(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsForceSecondPerSquareFoot() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsForceSecondPerSquareFoot(0)
    if err != nil {
        t.Errorf("FromPoundsForceSecondPerSquareFoot() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DynamicViscosityPoundForceSecondPerSquareFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsForceSecondPerSquareFoot() with zero value = %v, want 0", converted)
    }
}
// Test FromPoundsPerFootSecond function
func TestDynamicViscosityFactory_FromPoundsPerFootSecond(t *testing.T) {
    factory := units.DynamicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPoundsPerFootSecond(100)
    if err != nil {
        t.Errorf("FromPoundsPerFootSecond() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DynamicViscosityPoundPerFootSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPoundsPerFootSecond() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPoundsPerFootSecond(math.NaN())
    if err == nil {
        t.Error("FromPoundsPerFootSecond() with NaN value should return error")
    }

    _, err = factory.FromPoundsPerFootSecond(math.Inf(1))
    if err == nil {
        t.Error("FromPoundsPerFootSecond() with +Inf value should return error")
    }

    _, err = factory.FromPoundsPerFootSecond(math.Inf(-1))
    if err == nil {
        t.Error("FromPoundsPerFootSecond() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPoundsPerFootSecond(0)
    if err != nil {
        t.Errorf("FromPoundsPerFootSecond() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DynamicViscosityPoundPerFootSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPoundsPerFootSecond() with zero value = %v, want 0", converted)
    }
}
// Test FromMillipascalSeconds function
func TestDynamicViscosityFactory_FromMillipascalSeconds(t *testing.T) {
    factory := units.DynamicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillipascalSeconds(100)
    if err != nil {
        t.Errorf("FromMillipascalSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DynamicViscosityMillipascalSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillipascalSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillipascalSeconds(math.NaN())
    if err == nil {
        t.Error("FromMillipascalSeconds() with NaN value should return error")
    }

    _, err = factory.FromMillipascalSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromMillipascalSeconds() with +Inf value should return error")
    }

    _, err = factory.FromMillipascalSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromMillipascalSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillipascalSeconds(0)
    if err != nil {
        t.Errorf("FromMillipascalSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DynamicViscosityMillipascalSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillipascalSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromMicropascalSeconds function
func TestDynamicViscosityFactory_FromMicropascalSeconds(t *testing.T) {
    factory := units.DynamicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicropascalSeconds(100)
    if err != nil {
        t.Errorf("FromMicropascalSeconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DynamicViscosityMicropascalSecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicropascalSeconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicropascalSeconds(math.NaN())
    if err == nil {
        t.Error("FromMicropascalSeconds() with NaN value should return error")
    }

    _, err = factory.FromMicropascalSeconds(math.Inf(1))
    if err == nil {
        t.Error("FromMicropascalSeconds() with +Inf value should return error")
    }

    _, err = factory.FromMicropascalSeconds(math.Inf(-1))
    if err == nil {
        t.Error("FromMicropascalSeconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicropascalSeconds(0)
    if err != nil {
        t.Errorf("FromMicropascalSeconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DynamicViscosityMicropascalSecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicropascalSeconds() with zero value = %v, want 0", converted)
    }
}
// Test FromCentipoise function
func TestDynamicViscosityFactory_FromCentipoise(t *testing.T) {
    factory := units.DynamicViscosityFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentipoise(100)
    if err != nil {
        t.Errorf("FromCentipoise() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.DynamicViscosityCentipoise)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentipoise() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentipoise(math.NaN())
    if err == nil {
        t.Error("FromCentipoise() with NaN value should return error")
    }

    _, err = factory.FromCentipoise(math.Inf(1))
    if err == nil {
        t.Error("FromCentipoise() with +Inf value should return error")
    }

    _, err = factory.FromCentipoise(math.Inf(-1))
    if err == nil {
        t.Error("FromCentipoise() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentipoise(0)
    if err != nil {
        t.Errorf("FromCentipoise() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.DynamicViscosityCentipoise)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentipoise() with zero value = %v, want 0", converted)
    }
}

func TestDynamicViscosityToString(t *testing.T) {
	factory := units.DynamicViscosityFactory{}
	a, err := factory.CreateDynamicViscosity(45, units.DynamicViscosityNewtonSecondPerMeterSquared)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.DynamicViscosityNewtonSecondPerMeterSquared, 2)
	expected := "45.00 " + units.GetDynamicViscosityAbbreviation(units.DynamicViscosityNewtonSecondPerMeterSquared)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.DynamicViscosityNewtonSecondPerMeterSquared, -1)
	expected = "45 " + units.GetDynamicViscosityAbbreviation(units.DynamicViscosityNewtonSecondPerMeterSquared)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestDynamicViscosity_EqualityAndComparison(t *testing.T) {
	factory := units.DynamicViscosityFactory{}
	a1, _ := factory.CreateDynamicViscosity(60, units.DynamicViscosityNewtonSecondPerMeterSquared)
	a2, _ := factory.CreateDynamicViscosity(60, units.DynamicViscosityNewtonSecondPerMeterSquared)
	a3, _ := factory.CreateDynamicViscosity(120, units.DynamicViscosityNewtonSecondPerMeterSquared)

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

func TestDynamicViscosity_Arithmetic(t *testing.T) {
	factory := units.DynamicViscosityFactory{}
	a1, _ := factory.CreateDynamicViscosity(30, units.DynamicViscosityNewtonSecondPerMeterSquared)
	a2, _ := factory.CreateDynamicViscosity(45, units.DynamicViscosityNewtonSecondPerMeterSquared)

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