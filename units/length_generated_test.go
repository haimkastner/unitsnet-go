// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"
	"strings"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestLengthDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Meter"}`
	
	factory := units.LengthDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.LengthMeter {
		t.Errorf("expected unit %v, got %v", units.LengthMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Meter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestLengthDto_ToJSON(t *testing.T) {
	dto := units.LengthDto{
		Value: 45,
		Unit:  units.LengthMeter,
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
	if result["unit"].(string) != string(units.LengthMeter) {
		t.Errorf("expected unit %s, got %v", units.LengthMeter, result["unit"])
	}
}

func TestNewLength_InvalidValue(t *testing.T) {
	factory := units.LengthFactory{}
	// NaN value should return an error.
	_, err := factory.CreateLength(math.NaN(), units.LengthMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateLength(math.Inf(1), units.LengthMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestLengthConversions(t *testing.T) {
	factory := units.LengthFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateLength(180, units.LengthMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Meters.
		// No expected conversion value provided for Meters, verifying result is not NaN.
		result := a.Meters()
		cacheResult := a.Meters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Meters returned NaN")
		}
	}
	{
		// Test conversion to Miles.
		// No expected conversion value provided for Miles, verifying result is not NaN.
		result := a.Miles()
		cacheResult := a.Miles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Miles returned NaN")
		}
	}
	{
		// Test conversion to Yards.
		// No expected conversion value provided for Yards, verifying result is not NaN.
		result := a.Yards()
		cacheResult := a.Yards()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Yards returned NaN")
		}
	}
	{
		// Test conversion to Feet.
		// No expected conversion value provided for Feet, verifying result is not NaN.
		result := a.Feet()
		cacheResult := a.Feet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Feet returned NaN")
		}
	}
	{
		// Test conversion to UsSurveyFeet.
		// No expected conversion value provided for UsSurveyFeet, verifying result is not NaN.
		result := a.UsSurveyFeet()
		cacheResult := a.UsSurveyFeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to UsSurveyFeet returned NaN")
		}
	}
	{
		// Test conversion to Inches.
		// No expected conversion value provided for Inches, verifying result is not NaN.
		result := a.Inches()
		cacheResult := a.Inches()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Inches returned NaN")
		}
	}
	{
		// Test conversion to Mils.
		// No expected conversion value provided for Mils, verifying result is not NaN.
		result := a.Mils()
		cacheResult := a.Mils()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Mils returned NaN")
		}
	}
	{
		// Test conversion to NauticalMiles.
		// No expected conversion value provided for NauticalMiles, verifying result is not NaN.
		result := a.NauticalMiles()
		cacheResult := a.NauticalMiles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to NauticalMiles returned NaN")
		}
	}
	{
		// Test conversion to Fathoms.
		// No expected conversion value provided for Fathoms, verifying result is not NaN.
		result := a.Fathoms()
		cacheResult := a.Fathoms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Fathoms returned NaN")
		}
	}
	{
		// Test conversion to Shackles.
		// No expected conversion value provided for Shackles, verifying result is not NaN.
		result := a.Shackles()
		cacheResult := a.Shackles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Shackles returned NaN")
		}
	}
	{
		// Test conversion to Microinches.
		// No expected conversion value provided for Microinches, verifying result is not NaN.
		result := a.Microinches()
		cacheResult := a.Microinches()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Microinches returned NaN")
		}
	}
	{
		// Test conversion to PrinterPoints.
		// No expected conversion value provided for PrinterPoints, verifying result is not NaN.
		result := a.PrinterPoints()
		cacheResult := a.PrinterPoints()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PrinterPoints returned NaN")
		}
	}
	{
		// Test conversion to DtpPoints.
		// No expected conversion value provided for DtpPoints, verifying result is not NaN.
		result := a.DtpPoints()
		cacheResult := a.DtpPoints()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DtpPoints returned NaN")
		}
	}
	{
		// Test conversion to PrinterPicas.
		// No expected conversion value provided for PrinterPicas, verifying result is not NaN.
		result := a.PrinterPicas()
		cacheResult := a.PrinterPicas()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to PrinterPicas returned NaN")
		}
	}
	{
		// Test conversion to DtpPicas.
		// No expected conversion value provided for DtpPicas, verifying result is not NaN.
		result := a.DtpPicas()
		cacheResult := a.DtpPicas()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DtpPicas returned NaN")
		}
	}
	{
		// Test conversion to Twips.
		// No expected conversion value provided for Twips, verifying result is not NaN.
		result := a.Twips()
		cacheResult := a.Twips()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Twips returned NaN")
		}
	}
	{
		// Test conversion to Hands.
		// No expected conversion value provided for Hands, verifying result is not NaN.
		result := a.Hands()
		cacheResult := a.Hands()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Hands returned NaN")
		}
	}
	{
		// Test conversion to AstronomicalUnits.
		// No expected conversion value provided for AstronomicalUnits, verifying result is not NaN.
		result := a.AstronomicalUnits()
		cacheResult := a.AstronomicalUnits()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to AstronomicalUnits returned NaN")
		}
	}
	{
		// Test conversion to Parsecs.
		// No expected conversion value provided for Parsecs, verifying result is not NaN.
		result := a.Parsecs()
		cacheResult := a.Parsecs()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Parsecs returned NaN")
		}
	}
	{
		// Test conversion to LightYears.
		// No expected conversion value provided for LightYears, verifying result is not NaN.
		result := a.LightYears()
		cacheResult := a.LightYears()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to LightYears returned NaN")
		}
	}
	{
		// Test conversion to SolarRadiuses.
		// No expected conversion value provided for SolarRadiuses, verifying result is not NaN.
		result := a.SolarRadiuses()
		cacheResult := a.SolarRadiuses()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to SolarRadiuses returned NaN")
		}
	}
	{
		// Test conversion to Chains.
		// No expected conversion value provided for Chains, verifying result is not NaN.
		result := a.Chains()
		cacheResult := a.Chains()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Chains returned NaN")
		}
	}
	{
		// Test conversion to Angstroms.
		// No expected conversion value provided for Angstroms, verifying result is not NaN.
		result := a.Angstroms()
		cacheResult := a.Angstroms()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Angstroms returned NaN")
		}
	}
	{
		// Test conversion to DataMiles.
		// No expected conversion value provided for DataMiles, verifying result is not NaN.
		result := a.DataMiles()
		cacheResult := a.DataMiles()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to DataMiles returned NaN")
		}
	}
	{
		// Test conversion to Femtometers.
		// No expected conversion value provided for Femtometers, verifying result is not NaN.
		result := a.Femtometers()
		cacheResult := a.Femtometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Femtometers returned NaN")
		}
	}
	{
		// Test conversion to Picometers.
		// No expected conversion value provided for Picometers, verifying result is not NaN.
		result := a.Picometers()
		cacheResult := a.Picometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Picometers returned NaN")
		}
	}
	{
		// Test conversion to Nanometers.
		// No expected conversion value provided for Nanometers, verifying result is not NaN.
		result := a.Nanometers()
		cacheResult := a.Nanometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Nanometers returned NaN")
		}
	}
	{
		// Test conversion to Micrometers.
		// No expected conversion value provided for Micrometers, verifying result is not NaN.
		result := a.Micrometers()
		cacheResult := a.Micrometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Micrometers returned NaN")
		}
	}
	{
		// Test conversion to Millimeters.
		// No expected conversion value provided for Millimeters, verifying result is not NaN.
		result := a.Millimeters()
		cacheResult := a.Millimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Millimeters returned NaN")
		}
	}
	{
		// Test conversion to Centimeters.
		// No expected conversion value provided for Centimeters, verifying result is not NaN.
		result := a.Centimeters()
		cacheResult := a.Centimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Centimeters returned NaN")
		}
	}
	{
		// Test conversion to Decimeters.
		// No expected conversion value provided for Decimeters, verifying result is not NaN.
		result := a.Decimeters()
		cacheResult := a.Decimeters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Decimeters returned NaN")
		}
	}
	{
		// Test conversion to Decameters.
		// No expected conversion value provided for Decameters, verifying result is not NaN.
		result := a.Decameters()
		cacheResult := a.Decameters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Decameters returned NaN")
		}
	}
	{
		// Test conversion to Hectometers.
		// No expected conversion value provided for Hectometers, verifying result is not NaN.
		result := a.Hectometers()
		cacheResult := a.Hectometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Hectometers returned NaN")
		}
	}
	{
		// Test conversion to Kilometers.
		// No expected conversion value provided for Kilometers, verifying result is not NaN.
		result := a.Kilometers()
		cacheResult := a.Kilometers()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilometers returned NaN")
		}
	}
	{
		// Test conversion to Megameters.
		// No expected conversion value provided for Megameters, verifying result is not NaN.
		result := a.Megameters()
		cacheResult := a.Megameters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megameters returned NaN")
		}
	}
	{
		// Test conversion to Gigameters.
		// No expected conversion value provided for Gigameters, verifying result is not NaN.
		result := a.Gigameters()
		cacheResult := a.Gigameters()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Gigameters returned NaN")
		}
	}
	{
		// Test conversion to Kiloyards.
		// No expected conversion value provided for Kiloyards, verifying result is not NaN.
		result := a.Kiloyards()
		cacheResult := a.Kiloyards()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kiloyards returned NaN")
		}
	}
	{
		// Test conversion to Kilofeet.
		// No expected conversion value provided for Kilofeet, verifying result is not NaN.
		result := a.Kilofeet()
		cacheResult := a.Kilofeet()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kilofeet returned NaN")
		}
	}
	{
		// Test conversion to Kiloparsecs.
		// No expected conversion value provided for Kiloparsecs, verifying result is not NaN.
		result := a.Kiloparsecs()
		cacheResult := a.Kiloparsecs()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Kiloparsecs returned NaN")
		}
	}
	{
		// Test conversion to Megaparsecs.
		// No expected conversion value provided for Megaparsecs, verifying result is not NaN.
		result := a.Megaparsecs()
		cacheResult := a.Megaparsecs()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to Megaparsecs returned NaN")
		}
	}
	{
		// Test conversion to KilolightYears.
		// No expected conversion value provided for KilolightYears, verifying result is not NaN.
		result := a.KilolightYears()
		cacheResult := a.KilolightYears()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to KilolightYears returned NaN")
		}
	}
	{
		// Test conversion to MegalightYears.
		// No expected conversion value provided for MegalightYears, verifying result is not NaN.
		result := a.MegalightYears()
		cacheResult := a.MegalightYears()
		if math.IsNaN(result) || cacheResult != result {
			t.Errorf("conversion to MegalightYears returned NaN")
		}
	}
}

func TestLength_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.LengthFactory{}
	a, err := factory.CreateLength(90, units.LengthMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.LengthMeter {
		t.Errorf("expected default unit Meter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.LengthMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.LengthDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.LengthMeter {
		t.Errorf("expected unit Meter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestLengthFactory_FromDto(t *testing.T) {
    factory := units.LengthFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.LengthDto{
        Value: math.NaN(),
        Unit:  units.LengthMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Meter conversion
    metersDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthMeter,
    }
    
    var metersResult *units.Length
    metersResult, err = factory.FromDto(metersDto)
    if err != nil {
        t.Errorf("FromDto() with Meter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = metersResult.Convert(units.LengthMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Meter = %v, want %v", converted, 100)
    }
    // Test Mile conversion
    milesDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthMile,
    }
    
    var milesResult *units.Length
    milesResult, err = factory.FromDto(milesDto)
    if err != nil {
        t.Errorf("FromDto() with Mile returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milesResult.Convert(units.LengthMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mile = %v, want %v", converted, 100)
    }
    // Test Yard conversion
    yardsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthYard,
    }
    
    var yardsResult *units.Length
    yardsResult, err = factory.FromDto(yardsDto)
    if err != nil {
        t.Errorf("FromDto() with Yard returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = yardsResult.Convert(units.LengthYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Yard = %v, want %v", converted, 100)
    }
    // Test Foot conversion
    feetDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthFoot,
    }
    
    var feetResult *units.Length
    feetResult, err = factory.FromDto(feetDto)
    if err != nil {
        t.Errorf("FromDto() with Foot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feetResult.Convert(units.LengthFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Foot = %v, want %v", converted, 100)
    }
    // Test UsSurveyFoot conversion
    us_survey_feetDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthUsSurveyFoot,
    }
    
    var us_survey_feetResult *units.Length
    us_survey_feetResult, err = factory.FromDto(us_survey_feetDto)
    if err != nil {
        t.Errorf("FromDto() with UsSurveyFoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_survey_feetResult.Convert(units.LengthUsSurveyFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsSurveyFoot = %v, want %v", converted, 100)
    }
    // Test Inch conversion
    inchesDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthInch,
    }
    
    var inchesResult *units.Length
    inchesResult, err = factory.FromDto(inchesDto)
    if err != nil {
        t.Errorf("FromDto() with Inch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inchesResult.Convert(units.LengthInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Inch = %v, want %v", converted, 100)
    }
    // Test Mil conversion
    milsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthMil,
    }
    
    var milsResult *units.Length
    milsResult, err = factory.FromDto(milsDto)
    if err != nil {
        t.Errorf("FromDto() with Mil returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milsResult.Convert(units.LengthMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mil = %v, want %v", converted, 100)
    }
    // Test NauticalMile conversion
    nautical_milesDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthNauticalMile,
    }
    
    var nautical_milesResult *units.Length
    nautical_milesResult, err = factory.FromDto(nautical_milesDto)
    if err != nil {
        t.Errorf("FromDto() with NauticalMile returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nautical_milesResult.Convert(units.LengthNauticalMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NauticalMile = %v, want %v", converted, 100)
    }
    // Test Fathom conversion
    fathomsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthFathom,
    }
    
    var fathomsResult *units.Length
    fathomsResult, err = factory.FromDto(fathomsDto)
    if err != nil {
        t.Errorf("FromDto() with Fathom returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = fathomsResult.Convert(units.LengthFathom)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Fathom = %v, want %v", converted, 100)
    }
    // Test Shackle conversion
    shacklesDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthShackle,
    }
    
    var shacklesResult *units.Length
    shacklesResult, err = factory.FromDto(shacklesDto)
    if err != nil {
        t.Errorf("FromDto() with Shackle returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = shacklesResult.Convert(units.LengthShackle)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Shackle = %v, want %v", converted, 100)
    }
    // Test Microinch conversion
    microinchesDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthMicroinch,
    }
    
    var microinchesResult *units.Length
    microinchesResult, err = factory.FromDto(microinchesDto)
    if err != nil {
        t.Errorf("FromDto() with Microinch returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microinchesResult.Convert(units.LengthMicroinch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microinch = %v, want %v", converted, 100)
    }
    // Test PrinterPoint conversion
    printer_pointsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthPrinterPoint,
    }
    
    var printer_pointsResult *units.Length
    printer_pointsResult, err = factory.FromDto(printer_pointsDto)
    if err != nil {
        t.Errorf("FromDto() with PrinterPoint returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = printer_pointsResult.Convert(units.LengthPrinterPoint)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PrinterPoint = %v, want %v", converted, 100)
    }
    // Test DtpPoint conversion
    dtp_pointsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthDtpPoint,
    }
    
    var dtp_pointsResult *units.Length
    dtp_pointsResult, err = factory.FromDto(dtp_pointsDto)
    if err != nil {
        t.Errorf("FromDto() with DtpPoint returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = dtp_pointsResult.Convert(units.LengthDtpPoint)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DtpPoint = %v, want %v", converted, 100)
    }
    // Test PrinterPica conversion
    printer_picasDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthPrinterPica,
    }
    
    var printer_picasResult *units.Length
    printer_picasResult, err = factory.FromDto(printer_picasDto)
    if err != nil {
        t.Errorf("FromDto() with PrinterPica returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = printer_picasResult.Convert(units.LengthPrinterPica)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PrinterPica = %v, want %v", converted, 100)
    }
    // Test DtpPica conversion
    dtp_picasDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthDtpPica,
    }
    
    var dtp_picasResult *units.Length
    dtp_picasResult, err = factory.FromDto(dtp_picasDto)
    if err != nil {
        t.Errorf("FromDto() with DtpPica returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = dtp_picasResult.Convert(units.LengthDtpPica)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DtpPica = %v, want %v", converted, 100)
    }
    // Test Twip conversion
    twipsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthTwip,
    }
    
    var twipsResult *units.Length
    twipsResult, err = factory.FromDto(twipsDto)
    if err != nil {
        t.Errorf("FromDto() with Twip returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = twipsResult.Convert(units.LengthTwip)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Twip = %v, want %v", converted, 100)
    }
    // Test Hand conversion
    handsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthHand,
    }
    
    var handsResult *units.Length
    handsResult, err = factory.FromDto(handsDto)
    if err != nil {
        t.Errorf("FromDto() with Hand returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = handsResult.Convert(units.LengthHand)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hand = %v, want %v", converted, 100)
    }
    // Test AstronomicalUnit conversion
    astronomical_unitsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthAstronomicalUnit,
    }
    
    var astronomical_unitsResult *units.Length
    astronomical_unitsResult, err = factory.FromDto(astronomical_unitsDto)
    if err != nil {
        t.Errorf("FromDto() with AstronomicalUnit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = astronomical_unitsResult.Convert(units.LengthAstronomicalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AstronomicalUnit = %v, want %v", converted, 100)
    }
    // Test Parsec conversion
    parsecsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthParsec,
    }
    
    var parsecsResult *units.Length
    parsecsResult, err = factory.FromDto(parsecsDto)
    if err != nil {
        t.Errorf("FromDto() with Parsec returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parsecsResult.Convert(units.LengthParsec)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Parsec = %v, want %v", converted, 100)
    }
    // Test LightYear conversion
    light_yearsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthLightYear,
    }
    
    var light_yearsResult *units.Length
    light_yearsResult, err = factory.FromDto(light_yearsDto)
    if err != nil {
        t.Errorf("FromDto() with LightYear returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = light_yearsResult.Convert(units.LengthLightYear)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LightYear = %v, want %v", converted, 100)
    }
    // Test SolarRadius conversion
    solar_radiusesDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthSolarRadius,
    }
    
    var solar_radiusesResult *units.Length
    solar_radiusesResult, err = factory.FromDto(solar_radiusesDto)
    if err != nil {
        t.Errorf("FromDto() with SolarRadius returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = solar_radiusesResult.Convert(units.LengthSolarRadius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SolarRadius = %v, want %v", converted, 100)
    }
    // Test Chain conversion
    chainsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthChain,
    }
    
    var chainsResult *units.Length
    chainsResult, err = factory.FromDto(chainsDto)
    if err != nil {
        t.Errorf("FromDto() with Chain returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = chainsResult.Convert(units.LengthChain)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Chain = %v, want %v", converted, 100)
    }
    // Test Angstrom conversion
    angstromsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthAngstrom,
    }
    
    var angstromsResult *units.Length
    angstromsResult, err = factory.FromDto(angstromsDto)
    if err != nil {
        t.Errorf("FromDto() with Angstrom returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = angstromsResult.Convert(units.LengthAngstrom)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Angstrom = %v, want %v", converted, 100)
    }
    // Test DataMile conversion
    data_milesDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthDataMile,
    }
    
    var data_milesResult *units.Length
    data_milesResult, err = factory.FromDto(data_milesDto)
    if err != nil {
        t.Errorf("FromDto() with DataMile returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = data_milesResult.Convert(units.LengthDataMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DataMile = %v, want %v", converted, 100)
    }
    // Test Femtometer conversion
    femtometersDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthFemtometer,
    }
    
    var femtometersResult *units.Length
    femtometersResult, err = factory.FromDto(femtometersDto)
    if err != nil {
        t.Errorf("FromDto() with Femtometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtometersResult.Convert(units.LengthFemtometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtometer = %v, want %v", converted, 100)
    }
    // Test Picometer conversion
    picometersDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthPicometer,
    }
    
    var picometersResult *units.Length
    picometersResult, err = factory.FromDto(picometersDto)
    if err != nil {
        t.Errorf("FromDto() with Picometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picometersResult.Convert(units.LengthPicometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picometer = %v, want %v", converted, 100)
    }
    // Test Nanometer conversion
    nanometersDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthNanometer,
    }
    
    var nanometersResult *units.Length
    nanometersResult, err = factory.FromDto(nanometersDto)
    if err != nil {
        t.Errorf("FromDto() with Nanometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanometersResult.Convert(units.LengthNanometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanometer = %v, want %v", converted, 100)
    }
    // Test Micrometer conversion
    micrometersDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthMicrometer,
    }
    
    var micrometersResult *units.Length
    micrometersResult, err = factory.FromDto(micrometersDto)
    if err != nil {
        t.Errorf("FromDto() with Micrometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrometersResult.Convert(units.LengthMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Micrometer = %v, want %v", converted, 100)
    }
    // Test Millimeter conversion
    millimetersDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthMillimeter,
    }
    
    var millimetersResult *units.Length
    millimetersResult, err = factory.FromDto(millimetersDto)
    if err != nil {
        t.Errorf("FromDto() with Millimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimetersResult.Convert(units.LengthMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millimeter = %v, want %v", converted, 100)
    }
    // Test Centimeter conversion
    centimetersDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthCentimeter,
    }
    
    var centimetersResult *units.Length
    centimetersResult, err = factory.FromDto(centimetersDto)
    if err != nil {
        t.Errorf("FromDto() with Centimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimetersResult.Convert(units.LengthCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centimeter = %v, want %v", converted, 100)
    }
    // Test Decimeter conversion
    decimetersDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthDecimeter,
    }
    
    var decimetersResult *units.Length
    decimetersResult, err = factory.FromDto(decimetersDto)
    if err != nil {
        t.Errorf("FromDto() with Decimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimetersResult.Convert(units.LengthDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decimeter = %v, want %v", converted, 100)
    }
    // Test Decameter conversion
    decametersDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthDecameter,
    }
    
    var decametersResult *units.Length
    decametersResult, err = factory.FromDto(decametersDto)
    if err != nil {
        t.Errorf("FromDto() with Decameter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decametersResult.Convert(units.LengthDecameter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decameter = %v, want %v", converted, 100)
    }
    // Test Hectometer conversion
    hectometersDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthHectometer,
    }
    
    var hectometersResult *units.Length
    hectometersResult, err = factory.FromDto(hectometersDto)
    if err != nil {
        t.Errorf("FromDto() with Hectometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectometersResult.Convert(units.LengthHectometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hectometer = %v, want %v", converted, 100)
    }
    // Test Kilometer conversion
    kilometersDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthKilometer,
    }
    
    var kilometersResult *units.Length
    kilometersResult, err = factory.FromDto(kilometersDto)
    if err != nil {
        t.Errorf("FromDto() with Kilometer returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometersResult.Convert(units.LengthKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilometer = %v, want %v", converted, 100)
    }
    // Test Megameter conversion
    megametersDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthMegameter,
    }
    
    var megametersResult *units.Length
    megametersResult, err = factory.FromDto(megametersDto)
    if err != nil {
        t.Errorf("FromDto() with Megameter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megametersResult.Convert(units.LengthMegameter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megameter = %v, want %v", converted, 100)
    }
    // Test Gigameter conversion
    gigametersDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthGigameter,
    }
    
    var gigametersResult *units.Length
    gigametersResult, err = factory.FromDto(gigametersDto)
    if err != nil {
        t.Errorf("FromDto() with Gigameter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigametersResult.Convert(units.LengthGigameter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigameter = %v, want %v", converted, 100)
    }
    // Test Kiloyard conversion
    kiloyardsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthKiloyard,
    }
    
    var kiloyardsResult *units.Length
    kiloyardsResult, err = factory.FromDto(kiloyardsDto)
    if err != nil {
        t.Errorf("FromDto() with Kiloyard returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloyardsResult.Convert(units.LengthKiloyard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kiloyard = %v, want %v", converted, 100)
    }
    // Test Kilofoot conversion
    kilofeetDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthKilofoot,
    }
    
    var kilofeetResult *units.Length
    kilofeetResult, err = factory.FromDto(kilofeetDto)
    if err != nil {
        t.Errorf("FromDto() with Kilofoot returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilofeetResult.Convert(units.LengthKilofoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilofoot = %v, want %v", converted, 100)
    }
    // Test Kiloparsec conversion
    kiloparsecsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthKiloparsec,
    }
    
    var kiloparsecsResult *units.Length
    kiloparsecsResult, err = factory.FromDto(kiloparsecsDto)
    if err != nil {
        t.Errorf("FromDto() with Kiloparsec returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloparsecsResult.Convert(units.LengthKiloparsec)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kiloparsec = %v, want %v", converted, 100)
    }
    // Test Megaparsec conversion
    megaparsecsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthMegaparsec,
    }
    
    var megaparsecsResult *units.Length
    megaparsecsResult, err = factory.FromDto(megaparsecsDto)
    if err != nil {
        t.Errorf("FromDto() with Megaparsec returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaparsecsResult.Convert(units.LengthMegaparsec)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megaparsec = %v, want %v", converted, 100)
    }
    // Test KilolightYear conversion
    kilolight_yearsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthKilolightYear,
    }
    
    var kilolight_yearsResult *units.Length
    kilolight_yearsResult, err = factory.FromDto(kilolight_yearsDto)
    if err != nil {
        t.Errorf("FromDto() with KilolightYear returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilolight_yearsResult.Convert(units.LengthKilolightYear)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilolightYear = %v, want %v", converted, 100)
    }
    // Test MegalightYear conversion
    megalight_yearsDto := units.LengthDto{
        Value: 100,
        Unit:  units.LengthMegalightYear,
    }
    
    var megalight_yearsResult *units.Length
    megalight_yearsResult, err = factory.FromDto(megalight_yearsDto)
    if err != nil {
        t.Errorf("FromDto() with MegalightYear returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megalight_yearsResult.Convert(units.LengthMegalightYear)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegalightYear = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.LengthDto{
        Value: 0,
        Unit:  units.LengthMeter,
    }
    
    var zeroResult *units.Length
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestLengthFactory_FromDtoJSON(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Meter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Meter"}`)
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
    nanJSON, _ := json.Marshal(units.LengthDto{
        Value: nanValue,
        Unit:  units.LengthMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Meter unit
    metersJSON := []byte(`{"value": 100, "unit": "Meter"}`)
    metersResult, err := factory.FromDtoJSON(metersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Meter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = metersResult.Convert(units.LengthMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Meter = %v, want %v", converted, 100)
    }
    // Test JSON with Mile unit
    milesJSON := []byte(`{"value": 100, "unit": "Mile"}`)
    milesResult, err := factory.FromDtoJSON(milesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Mile unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milesResult.Convert(units.LengthMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mile = %v, want %v", converted, 100)
    }
    // Test JSON with Yard unit
    yardsJSON := []byte(`{"value": 100, "unit": "Yard"}`)
    yardsResult, err := factory.FromDtoJSON(yardsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Yard unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = yardsResult.Convert(units.LengthYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Yard = %v, want %v", converted, 100)
    }
    // Test JSON with Foot unit
    feetJSON := []byte(`{"value": 100, "unit": "Foot"}`)
    feetResult, err := factory.FromDtoJSON(feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Foot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = feetResult.Convert(units.LengthFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Foot = %v, want %v", converted, 100)
    }
    // Test JSON with UsSurveyFoot unit
    us_survey_feetJSON := []byte(`{"value": 100, "unit": "UsSurveyFoot"}`)
    us_survey_feetResult, err := factory.FromDtoJSON(us_survey_feetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with UsSurveyFoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = us_survey_feetResult.Convert(units.LengthUsSurveyFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for UsSurveyFoot = %v, want %v", converted, 100)
    }
    // Test JSON with Inch unit
    inchesJSON := []byte(`{"value": 100, "unit": "Inch"}`)
    inchesResult, err := factory.FromDtoJSON(inchesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Inch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = inchesResult.Convert(units.LengthInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Inch = %v, want %v", converted, 100)
    }
    // Test JSON with Mil unit
    milsJSON := []byte(`{"value": 100, "unit": "Mil"}`)
    milsResult, err := factory.FromDtoJSON(milsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Mil unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milsResult.Convert(units.LengthMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Mil = %v, want %v", converted, 100)
    }
    // Test JSON with NauticalMile unit
    nautical_milesJSON := []byte(`{"value": 100, "unit": "NauticalMile"}`)
    nautical_milesResult, err := factory.FromDtoJSON(nautical_milesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NauticalMile unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nautical_milesResult.Convert(units.LengthNauticalMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NauticalMile = %v, want %v", converted, 100)
    }
    // Test JSON with Fathom unit
    fathomsJSON := []byte(`{"value": 100, "unit": "Fathom"}`)
    fathomsResult, err := factory.FromDtoJSON(fathomsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Fathom unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = fathomsResult.Convert(units.LengthFathom)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Fathom = %v, want %v", converted, 100)
    }
    // Test JSON with Shackle unit
    shacklesJSON := []byte(`{"value": 100, "unit": "Shackle"}`)
    shacklesResult, err := factory.FromDtoJSON(shacklesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Shackle unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = shacklesResult.Convert(units.LengthShackle)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Shackle = %v, want %v", converted, 100)
    }
    // Test JSON with Microinch unit
    microinchesJSON := []byte(`{"value": 100, "unit": "Microinch"}`)
    microinchesResult, err := factory.FromDtoJSON(microinchesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microinch unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microinchesResult.Convert(units.LengthMicroinch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microinch = %v, want %v", converted, 100)
    }
    // Test JSON with PrinterPoint unit
    printer_pointsJSON := []byte(`{"value": 100, "unit": "PrinterPoint"}`)
    printer_pointsResult, err := factory.FromDtoJSON(printer_pointsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PrinterPoint unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = printer_pointsResult.Convert(units.LengthPrinterPoint)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PrinterPoint = %v, want %v", converted, 100)
    }
    // Test JSON with DtpPoint unit
    dtp_pointsJSON := []byte(`{"value": 100, "unit": "DtpPoint"}`)
    dtp_pointsResult, err := factory.FromDtoJSON(dtp_pointsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DtpPoint unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = dtp_pointsResult.Convert(units.LengthDtpPoint)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DtpPoint = %v, want %v", converted, 100)
    }
    // Test JSON with PrinterPica unit
    printer_picasJSON := []byte(`{"value": 100, "unit": "PrinterPica"}`)
    printer_picasResult, err := factory.FromDtoJSON(printer_picasJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PrinterPica unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = printer_picasResult.Convert(units.LengthPrinterPica)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PrinterPica = %v, want %v", converted, 100)
    }
    // Test JSON with DtpPica unit
    dtp_picasJSON := []byte(`{"value": 100, "unit": "DtpPica"}`)
    dtp_picasResult, err := factory.FromDtoJSON(dtp_picasJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DtpPica unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = dtp_picasResult.Convert(units.LengthDtpPica)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DtpPica = %v, want %v", converted, 100)
    }
    // Test JSON with Twip unit
    twipsJSON := []byte(`{"value": 100, "unit": "Twip"}`)
    twipsResult, err := factory.FromDtoJSON(twipsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Twip unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = twipsResult.Convert(units.LengthTwip)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Twip = %v, want %v", converted, 100)
    }
    // Test JSON with Hand unit
    handsJSON := []byte(`{"value": 100, "unit": "Hand"}`)
    handsResult, err := factory.FromDtoJSON(handsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Hand unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = handsResult.Convert(units.LengthHand)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hand = %v, want %v", converted, 100)
    }
    // Test JSON with AstronomicalUnit unit
    astronomical_unitsJSON := []byte(`{"value": 100, "unit": "AstronomicalUnit"}`)
    astronomical_unitsResult, err := factory.FromDtoJSON(astronomical_unitsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with AstronomicalUnit unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = astronomical_unitsResult.Convert(units.LengthAstronomicalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for AstronomicalUnit = %v, want %v", converted, 100)
    }
    // Test JSON with Parsec unit
    parsecsJSON := []byte(`{"value": 100, "unit": "Parsec"}`)
    parsecsResult, err := factory.FromDtoJSON(parsecsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Parsec unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = parsecsResult.Convert(units.LengthParsec)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Parsec = %v, want %v", converted, 100)
    }
    // Test JSON with LightYear unit
    light_yearsJSON := []byte(`{"value": 100, "unit": "LightYear"}`)
    light_yearsResult, err := factory.FromDtoJSON(light_yearsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with LightYear unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = light_yearsResult.Convert(units.LengthLightYear)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for LightYear = %v, want %v", converted, 100)
    }
    // Test JSON with SolarRadius unit
    solar_radiusesJSON := []byte(`{"value": 100, "unit": "SolarRadius"}`)
    solar_radiusesResult, err := factory.FromDtoJSON(solar_radiusesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with SolarRadius unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = solar_radiusesResult.Convert(units.LengthSolarRadius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for SolarRadius = %v, want %v", converted, 100)
    }
    // Test JSON with Chain unit
    chainsJSON := []byte(`{"value": 100, "unit": "Chain"}`)
    chainsResult, err := factory.FromDtoJSON(chainsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Chain unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = chainsResult.Convert(units.LengthChain)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Chain = %v, want %v", converted, 100)
    }
    // Test JSON with Angstrom unit
    angstromsJSON := []byte(`{"value": 100, "unit": "Angstrom"}`)
    angstromsResult, err := factory.FromDtoJSON(angstromsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Angstrom unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = angstromsResult.Convert(units.LengthAngstrom)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Angstrom = %v, want %v", converted, 100)
    }
    // Test JSON with DataMile unit
    data_milesJSON := []byte(`{"value": 100, "unit": "DataMile"}`)
    data_milesResult, err := factory.FromDtoJSON(data_milesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with DataMile unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = data_milesResult.Convert(units.LengthDataMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for DataMile = %v, want %v", converted, 100)
    }
    // Test JSON with Femtometer unit
    femtometersJSON := []byte(`{"value": 100, "unit": "Femtometer"}`)
    femtometersResult, err := factory.FromDtoJSON(femtometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Femtometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = femtometersResult.Convert(units.LengthFemtometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Femtometer = %v, want %v", converted, 100)
    }
    // Test JSON with Picometer unit
    picometersJSON := []byte(`{"value": 100, "unit": "Picometer"}`)
    picometersResult, err := factory.FromDtoJSON(picometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Picometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picometersResult.Convert(units.LengthPicometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Picometer = %v, want %v", converted, 100)
    }
    // Test JSON with Nanometer unit
    nanometersJSON := []byte(`{"value": 100, "unit": "Nanometer"}`)
    nanometersResult, err := factory.FromDtoJSON(nanometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanometersResult.Convert(units.LengthNanometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanometer = %v, want %v", converted, 100)
    }
    // Test JSON with Micrometer unit
    micrometersJSON := []byte(`{"value": 100, "unit": "Micrometer"}`)
    micrometersResult, err := factory.FromDtoJSON(micrometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Micrometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = micrometersResult.Convert(units.LengthMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Micrometer = %v, want %v", converted, 100)
    }
    // Test JSON with Millimeter unit
    millimetersJSON := []byte(`{"value": 100, "unit": "Millimeter"}`)
    millimetersResult, err := factory.FromDtoJSON(millimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millimetersResult.Convert(units.LengthMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millimeter = %v, want %v", converted, 100)
    }
    // Test JSON with Centimeter unit
    centimetersJSON := []byte(`{"value": 100, "unit": "Centimeter"}`)
    centimetersResult, err := factory.FromDtoJSON(centimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Centimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centimetersResult.Convert(units.LengthCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centimeter = %v, want %v", converted, 100)
    }
    // Test JSON with Decimeter unit
    decimetersJSON := []byte(`{"value": 100, "unit": "Decimeter"}`)
    decimetersResult, err := factory.FromDtoJSON(decimetersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Decimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decimetersResult.Convert(units.LengthDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decimeter = %v, want %v", converted, 100)
    }
    // Test JSON with Decameter unit
    decametersJSON := []byte(`{"value": 100, "unit": "Decameter"}`)
    decametersResult, err := factory.FromDtoJSON(decametersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Decameter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = decametersResult.Convert(units.LengthDecameter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Decameter = %v, want %v", converted, 100)
    }
    // Test JSON with Hectometer unit
    hectometersJSON := []byte(`{"value": 100, "unit": "Hectometer"}`)
    hectometersResult, err := factory.FromDtoJSON(hectometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Hectometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = hectometersResult.Convert(units.LengthHectometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Hectometer = %v, want %v", converted, 100)
    }
    // Test JSON with Kilometer unit
    kilometersJSON := []byte(`{"value": 100, "unit": "Kilometer"}`)
    kilometersResult, err := factory.FromDtoJSON(kilometersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilometer unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilometersResult.Convert(units.LengthKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilometer = %v, want %v", converted, 100)
    }
    // Test JSON with Megameter unit
    megametersJSON := []byte(`{"value": 100, "unit": "Megameter"}`)
    megametersResult, err := factory.FromDtoJSON(megametersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megameter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megametersResult.Convert(units.LengthMegameter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megameter = %v, want %v", converted, 100)
    }
    // Test JSON with Gigameter unit
    gigametersJSON := []byte(`{"value": 100, "unit": "Gigameter"}`)
    gigametersResult, err := factory.FromDtoJSON(gigametersJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gigameter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gigametersResult.Convert(units.LengthGigameter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gigameter = %v, want %v", converted, 100)
    }
    // Test JSON with Kiloyard unit
    kiloyardsJSON := []byte(`{"value": 100, "unit": "Kiloyard"}`)
    kiloyardsResult, err := factory.FromDtoJSON(kiloyardsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kiloyard unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloyardsResult.Convert(units.LengthKiloyard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kiloyard = %v, want %v", converted, 100)
    }
    // Test JSON with Kilofoot unit
    kilofeetJSON := []byte(`{"value": 100, "unit": "Kilofoot"}`)
    kilofeetResult, err := factory.FromDtoJSON(kilofeetJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kilofoot unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilofeetResult.Convert(units.LengthKilofoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kilofoot = %v, want %v", converted, 100)
    }
    // Test JSON with Kiloparsec unit
    kiloparsecsJSON := []byte(`{"value": 100, "unit": "Kiloparsec"}`)
    kiloparsecsResult, err := factory.FromDtoJSON(kiloparsecsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Kiloparsec unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kiloparsecsResult.Convert(units.LengthKiloparsec)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Kiloparsec = %v, want %v", converted, 100)
    }
    // Test JSON with Megaparsec unit
    megaparsecsJSON := []byte(`{"value": 100, "unit": "Megaparsec"}`)
    megaparsecsResult, err := factory.FromDtoJSON(megaparsecsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Megaparsec unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megaparsecsResult.Convert(units.LengthMegaparsec)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Megaparsec = %v, want %v", converted, 100)
    }
    // Test JSON with KilolightYear unit
    kilolight_yearsJSON := []byte(`{"value": 100, "unit": "KilolightYear"}`)
    kilolight_yearsResult, err := factory.FromDtoJSON(kilolight_yearsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilolightYear unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilolight_yearsResult.Convert(units.LengthKilolightYear)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilolightYear = %v, want %v", converted, 100)
    }
    // Test JSON with MegalightYear unit
    megalight_yearsJSON := []byte(`{"value": 100, "unit": "MegalightYear"}`)
    megalight_yearsResult, err := factory.FromDtoJSON(megalight_yearsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegalightYear unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megalight_yearsResult.Convert(units.LengthMegalightYear)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegalightYear = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Meter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromMeters function
func TestLengthFactory_FromMeters(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMeters(100)
    if err != nil {
        t.Errorf("FromMeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMeters(math.NaN())
    if err == nil {
        t.Error("FromMeters() with NaN value should return error")
    }

    _, err = factory.FromMeters(math.Inf(1))
    if err == nil {
        t.Error("FromMeters() with +Inf value should return error")
    }

    _, err = factory.FromMeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMeters(0)
    if err != nil {
        t.Errorf("FromMeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMeters() with zero value = %v, want 0", converted)
    }
}
// Test FromMiles function
func TestLengthFactory_FromMiles(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMiles(100)
    if err != nil {
        t.Errorf("FromMiles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMiles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMiles(math.NaN())
    if err == nil {
        t.Error("FromMiles() with NaN value should return error")
    }

    _, err = factory.FromMiles(math.Inf(1))
    if err == nil {
        t.Error("FromMiles() with +Inf value should return error")
    }

    _, err = factory.FromMiles(math.Inf(-1))
    if err == nil {
        t.Error("FromMiles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMiles(0)
    if err != nil {
        t.Errorf("FromMiles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthMile)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMiles() with zero value = %v, want 0", converted)
    }
}
// Test FromYards function
func TestLengthFactory_FromYards(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromYards(100)
    if err != nil {
        t.Errorf("FromYards() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthYard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromYards() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromYards(math.NaN())
    if err == nil {
        t.Error("FromYards() with NaN value should return error")
    }

    _, err = factory.FromYards(math.Inf(1))
    if err == nil {
        t.Error("FromYards() with +Inf value should return error")
    }

    _, err = factory.FromYards(math.Inf(-1))
    if err == nil {
        t.Error("FromYards() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromYards(0)
    if err != nil {
        t.Errorf("FromYards() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthYard)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromYards() with zero value = %v, want 0", converted)
    }
}
// Test FromFeet function
func TestLengthFactory_FromFeet(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFeet(100)
    if err != nil {
        t.Errorf("FromFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFeet(math.NaN())
    if err == nil {
        t.Error("FromFeet() with NaN value should return error")
    }

    _, err = factory.FromFeet(math.Inf(1))
    if err == nil {
        t.Error("FromFeet() with +Inf value should return error")
    }

    _, err = factory.FromFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFeet(0)
    if err != nil {
        t.Errorf("FromFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromUsSurveyFeet function
func TestLengthFactory_FromUsSurveyFeet(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromUsSurveyFeet(100)
    if err != nil {
        t.Errorf("FromUsSurveyFeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthUsSurveyFoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromUsSurveyFeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromUsSurveyFeet(math.NaN())
    if err == nil {
        t.Error("FromUsSurveyFeet() with NaN value should return error")
    }

    _, err = factory.FromUsSurveyFeet(math.Inf(1))
    if err == nil {
        t.Error("FromUsSurveyFeet() with +Inf value should return error")
    }

    _, err = factory.FromUsSurveyFeet(math.Inf(-1))
    if err == nil {
        t.Error("FromUsSurveyFeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromUsSurveyFeet(0)
    if err != nil {
        t.Errorf("FromUsSurveyFeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthUsSurveyFoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromUsSurveyFeet() with zero value = %v, want 0", converted)
    }
}
// Test FromInches function
func TestLengthFactory_FromInches(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromInches(100)
    if err != nil {
        t.Errorf("FromInches() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthInch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromInches() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromInches(math.NaN())
    if err == nil {
        t.Error("FromInches() with NaN value should return error")
    }

    _, err = factory.FromInches(math.Inf(1))
    if err == nil {
        t.Error("FromInches() with +Inf value should return error")
    }

    _, err = factory.FromInches(math.Inf(-1))
    if err == nil {
        t.Error("FromInches() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromInches(0)
    if err != nil {
        t.Errorf("FromInches() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthInch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromInches() with zero value = %v, want 0", converted)
    }
}
// Test FromMils function
func TestLengthFactory_FromMils(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMils(100)
    if err != nil {
        t.Errorf("FromMils() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMils() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMils(math.NaN())
    if err == nil {
        t.Error("FromMils() with NaN value should return error")
    }

    _, err = factory.FromMils(math.Inf(1))
    if err == nil {
        t.Error("FromMils() with +Inf value should return error")
    }

    _, err = factory.FromMils(math.Inf(-1))
    if err == nil {
        t.Error("FromMils() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMils(0)
    if err != nil {
        t.Errorf("FromMils() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthMil)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMils() with zero value = %v, want 0", converted)
    }
}
// Test FromNauticalMiles function
func TestLengthFactory_FromNauticalMiles(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNauticalMiles(100)
    if err != nil {
        t.Errorf("FromNauticalMiles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthNauticalMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNauticalMiles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNauticalMiles(math.NaN())
    if err == nil {
        t.Error("FromNauticalMiles() with NaN value should return error")
    }

    _, err = factory.FromNauticalMiles(math.Inf(1))
    if err == nil {
        t.Error("FromNauticalMiles() with +Inf value should return error")
    }

    _, err = factory.FromNauticalMiles(math.Inf(-1))
    if err == nil {
        t.Error("FromNauticalMiles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNauticalMiles(0)
    if err != nil {
        t.Errorf("FromNauticalMiles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthNauticalMile)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNauticalMiles() with zero value = %v, want 0", converted)
    }
}
// Test FromFathoms function
func TestLengthFactory_FromFathoms(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFathoms(100)
    if err != nil {
        t.Errorf("FromFathoms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthFathom)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFathoms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFathoms(math.NaN())
    if err == nil {
        t.Error("FromFathoms() with NaN value should return error")
    }

    _, err = factory.FromFathoms(math.Inf(1))
    if err == nil {
        t.Error("FromFathoms() with +Inf value should return error")
    }

    _, err = factory.FromFathoms(math.Inf(-1))
    if err == nil {
        t.Error("FromFathoms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFathoms(0)
    if err != nil {
        t.Errorf("FromFathoms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthFathom)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFathoms() with zero value = %v, want 0", converted)
    }
}
// Test FromShackles function
func TestLengthFactory_FromShackles(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromShackles(100)
    if err != nil {
        t.Errorf("FromShackles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthShackle)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromShackles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromShackles(math.NaN())
    if err == nil {
        t.Error("FromShackles() with NaN value should return error")
    }

    _, err = factory.FromShackles(math.Inf(1))
    if err == nil {
        t.Error("FromShackles() with +Inf value should return error")
    }

    _, err = factory.FromShackles(math.Inf(-1))
    if err == nil {
        t.Error("FromShackles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromShackles(0)
    if err != nil {
        t.Errorf("FromShackles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthShackle)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromShackles() with zero value = %v, want 0", converted)
    }
}
// Test FromMicroinches function
func TestLengthFactory_FromMicroinches(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicroinches(100)
    if err != nil {
        t.Errorf("FromMicroinches() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthMicroinch)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicroinches() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicroinches(math.NaN())
    if err == nil {
        t.Error("FromMicroinches() with NaN value should return error")
    }

    _, err = factory.FromMicroinches(math.Inf(1))
    if err == nil {
        t.Error("FromMicroinches() with +Inf value should return error")
    }

    _, err = factory.FromMicroinches(math.Inf(-1))
    if err == nil {
        t.Error("FromMicroinches() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicroinches(0)
    if err != nil {
        t.Errorf("FromMicroinches() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthMicroinch)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicroinches() with zero value = %v, want 0", converted)
    }
}
// Test FromPrinterPoints function
func TestLengthFactory_FromPrinterPoints(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPrinterPoints(100)
    if err != nil {
        t.Errorf("FromPrinterPoints() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthPrinterPoint)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPrinterPoints() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPrinterPoints(math.NaN())
    if err == nil {
        t.Error("FromPrinterPoints() with NaN value should return error")
    }

    _, err = factory.FromPrinterPoints(math.Inf(1))
    if err == nil {
        t.Error("FromPrinterPoints() with +Inf value should return error")
    }

    _, err = factory.FromPrinterPoints(math.Inf(-1))
    if err == nil {
        t.Error("FromPrinterPoints() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPrinterPoints(0)
    if err != nil {
        t.Errorf("FromPrinterPoints() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthPrinterPoint)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPrinterPoints() with zero value = %v, want 0", converted)
    }
}
// Test FromDtpPoints function
func TestLengthFactory_FromDtpPoints(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDtpPoints(100)
    if err != nil {
        t.Errorf("FromDtpPoints() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthDtpPoint)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDtpPoints() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDtpPoints(math.NaN())
    if err == nil {
        t.Error("FromDtpPoints() with NaN value should return error")
    }

    _, err = factory.FromDtpPoints(math.Inf(1))
    if err == nil {
        t.Error("FromDtpPoints() with +Inf value should return error")
    }

    _, err = factory.FromDtpPoints(math.Inf(-1))
    if err == nil {
        t.Error("FromDtpPoints() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDtpPoints(0)
    if err != nil {
        t.Errorf("FromDtpPoints() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthDtpPoint)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDtpPoints() with zero value = %v, want 0", converted)
    }
}
// Test FromPrinterPicas function
func TestLengthFactory_FromPrinterPicas(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPrinterPicas(100)
    if err != nil {
        t.Errorf("FromPrinterPicas() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthPrinterPica)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPrinterPicas() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPrinterPicas(math.NaN())
    if err == nil {
        t.Error("FromPrinterPicas() with NaN value should return error")
    }

    _, err = factory.FromPrinterPicas(math.Inf(1))
    if err == nil {
        t.Error("FromPrinterPicas() with +Inf value should return error")
    }

    _, err = factory.FromPrinterPicas(math.Inf(-1))
    if err == nil {
        t.Error("FromPrinterPicas() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPrinterPicas(0)
    if err != nil {
        t.Errorf("FromPrinterPicas() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthPrinterPica)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPrinterPicas() with zero value = %v, want 0", converted)
    }
}
// Test FromDtpPicas function
func TestLengthFactory_FromDtpPicas(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDtpPicas(100)
    if err != nil {
        t.Errorf("FromDtpPicas() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthDtpPica)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDtpPicas() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDtpPicas(math.NaN())
    if err == nil {
        t.Error("FromDtpPicas() with NaN value should return error")
    }

    _, err = factory.FromDtpPicas(math.Inf(1))
    if err == nil {
        t.Error("FromDtpPicas() with +Inf value should return error")
    }

    _, err = factory.FromDtpPicas(math.Inf(-1))
    if err == nil {
        t.Error("FromDtpPicas() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDtpPicas(0)
    if err != nil {
        t.Errorf("FromDtpPicas() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthDtpPica)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDtpPicas() with zero value = %v, want 0", converted)
    }
}
// Test FromTwips function
func TestLengthFactory_FromTwips(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTwips(100)
    if err != nil {
        t.Errorf("FromTwips() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthTwip)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTwips() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTwips(math.NaN())
    if err == nil {
        t.Error("FromTwips() with NaN value should return error")
    }

    _, err = factory.FromTwips(math.Inf(1))
    if err == nil {
        t.Error("FromTwips() with +Inf value should return error")
    }

    _, err = factory.FromTwips(math.Inf(-1))
    if err == nil {
        t.Error("FromTwips() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTwips(0)
    if err != nil {
        t.Errorf("FromTwips() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthTwip)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTwips() with zero value = %v, want 0", converted)
    }
}
// Test FromHands function
func TestLengthFactory_FromHands(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHands(100)
    if err != nil {
        t.Errorf("FromHands() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthHand)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHands() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHands(math.NaN())
    if err == nil {
        t.Error("FromHands() with NaN value should return error")
    }

    _, err = factory.FromHands(math.Inf(1))
    if err == nil {
        t.Error("FromHands() with +Inf value should return error")
    }

    _, err = factory.FromHands(math.Inf(-1))
    if err == nil {
        t.Error("FromHands() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHands(0)
    if err != nil {
        t.Errorf("FromHands() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthHand)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHands() with zero value = %v, want 0", converted)
    }
}
// Test FromAstronomicalUnits function
func TestLengthFactory_FromAstronomicalUnits(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAstronomicalUnits(100)
    if err != nil {
        t.Errorf("FromAstronomicalUnits() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthAstronomicalUnit)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAstronomicalUnits() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAstronomicalUnits(math.NaN())
    if err == nil {
        t.Error("FromAstronomicalUnits() with NaN value should return error")
    }

    _, err = factory.FromAstronomicalUnits(math.Inf(1))
    if err == nil {
        t.Error("FromAstronomicalUnits() with +Inf value should return error")
    }

    _, err = factory.FromAstronomicalUnits(math.Inf(-1))
    if err == nil {
        t.Error("FromAstronomicalUnits() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAstronomicalUnits(0)
    if err != nil {
        t.Errorf("FromAstronomicalUnits() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthAstronomicalUnit)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAstronomicalUnits() with zero value = %v, want 0", converted)
    }
}
// Test FromParsecs function
func TestLengthFactory_FromParsecs(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromParsecs(100)
    if err != nil {
        t.Errorf("FromParsecs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthParsec)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromParsecs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromParsecs(math.NaN())
    if err == nil {
        t.Error("FromParsecs() with NaN value should return error")
    }

    _, err = factory.FromParsecs(math.Inf(1))
    if err == nil {
        t.Error("FromParsecs() with +Inf value should return error")
    }

    _, err = factory.FromParsecs(math.Inf(-1))
    if err == nil {
        t.Error("FromParsecs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromParsecs(0)
    if err != nil {
        t.Errorf("FromParsecs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthParsec)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromParsecs() with zero value = %v, want 0", converted)
    }
}
// Test FromLightYears function
func TestLengthFactory_FromLightYears(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromLightYears(100)
    if err != nil {
        t.Errorf("FromLightYears() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthLightYear)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromLightYears() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromLightYears(math.NaN())
    if err == nil {
        t.Error("FromLightYears() with NaN value should return error")
    }

    _, err = factory.FromLightYears(math.Inf(1))
    if err == nil {
        t.Error("FromLightYears() with +Inf value should return error")
    }

    _, err = factory.FromLightYears(math.Inf(-1))
    if err == nil {
        t.Error("FromLightYears() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromLightYears(0)
    if err != nil {
        t.Errorf("FromLightYears() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthLightYear)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromLightYears() with zero value = %v, want 0", converted)
    }
}
// Test FromSolarRadiuses function
func TestLengthFactory_FromSolarRadiuses(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromSolarRadiuses(100)
    if err != nil {
        t.Errorf("FromSolarRadiuses() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthSolarRadius)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromSolarRadiuses() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromSolarRadiuses(math.NaN())
    if err == nil {
        t.Error("FromSolarRadiuses() with NaN value should return error")
    }

    _, err = factory.FromSolarRadiuses(math.Inf(1))
    if err == nil {
        t.Error("FromSolarRadiuses() with +Inf value should return error")
    }

    _, err = factory.FromSolarRadiuses(math.Inf(-1))
    if err == nil {
        t.Error("FromSolarRadiuses() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromSolarRadiuses(0)
    if err != nil {
        t.Errorf("FromSolarRadiuses() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthSolarRadius)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromSolarRadiuses() with zero value = %v, want 0", converted)
    }
}
// Test FromChains function
func TestLengthFactory_FromChains(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromChains(100)
    if err != nil {
        t.Errorf("FromChains() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthChain)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromChains() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromChains(math.NaN())
    if err == nil {
        t.Error("FromChains() with NaN value should return error")
    }

    _, err = factory.FromChains(math.Inf(1))
    if err == nil {
        t.Error("FromChains() with +Inf value should return error")
    }

    _, err = factory.FromChains(math.Inf(-1))
    if err == nil {
        t.Error("FromChains() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromChains(0)
    if err != nil {
        t.Errorf("FromChains() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthChain)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromChains() with zero value = %v, want 0", converted)
    }
}
// Test FromAngstroms function
func TestLengthFactory_FromAngstroms(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromAngstroms(100)
    if err != nil {
        t.Errorf("FromAngstroms() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthAngstrom)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromAngstroms() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromAngstroms(math.NaN())
    if err == nil {
        t.Error("FromAngstroms() with NaN value should return error")
    }

    _, err = factory.FromAngstroms(math.Inf(1))
    if err == nil {
        t.Error("FromAngstroms() with +Inf value should return error")
    }

    _, err = factory.FromAngstroms(math.Inf(-1))
    if err == nil {
        t.Error("FromAngstroms() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromAngstroms(0)
    if err != nil {
        t.Errorf("FromAngstroms() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthAngstrom)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromAngstroms() with zero value = %v, want 0", converted)
    }
}
// Test FromDataMiles function
func TestLengthFactory_FromDataMiles(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDataMiles(100)
    if err != nil {
        t.Errorf("FromDataMiles() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthDataMile)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDataMiles() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDataMiles(math.NaN())
    if err == nil {
        t.Error("FromDataMiles() with NaN value should return error")
    }

    _, err = factory.FromDataMiles(math.Inf(1))
    if err == nil {
        t.Error("FromDataMiles() with +Inf value should return error")
    }

    _, err = factory.FromDataMiles(math.Inf(-1))
    if err == nil {
        t.Error("FromDataMiles() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDataMiles(0)
    if err != nil {
        t.Errorf("FromDataMiles() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthDataMile)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDataMiles() with zero value = %v, want 0", converted)
    }
}
// Test FromFemtometers function
func TestLengthFactory_FromFemtometers(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromFemtometers(100)
    if err != nil {
        t.Errorf("FromFemtometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthFemtometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromFemtometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromFemtometers(math.NaN())
    if err == nil {
        t.Error("FromFemtometers() with NaN value should return error")
    }

    _, err = factory.FromFemtometers(math.Inf(1))
    if err == nil {
        t.Error("FromFemtometers() with +Inf value should return error")
    }

    _, err = factory.FromFemtometers(math.Inf(-1))
    if err == nil {
        t.Error("FromFemtometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromFemtometers(0)
    if err != nil {
        t.Errorf("FromFemtometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthFemtometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromFemtometers() with zero value = %v, want 0", converted)
    }
}
// Test FromPicometers function
func TestLengthFactory_FromPicometers(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicometers(100)
    if err != nil {
        t.Errorf("FromPicometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthPicometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicometers(math.NaN())
    if err == nil {
        t.Error("FromPicometers() with NaN value should return error")
    }

    _, err = factory.FromPicometers(math.Inf(1))
    if err == nil {
        t.Error("FromPicometers() with +Inf value should return error")
    }

    _, err = factory.FromPicometers(math.Inf(-1))
    if err == nil {
        t.Error("FromPicometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicometers(0)
    if err != nil {
        t.Errorf("FromPicometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthPicometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicometers() with zero value = %v, want 0", converted)
    }
}
// Test FromNanometers function
func TestLengthFactory_FromNanometers(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanometers(100)
    if err != nil {
        t.Errorf("FromNanometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthNanometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanometers(math.NaN())
    if err == nil {
        t.Error("FromNanometers() with NaN value should return error")
    }

    _, err = factory.FromNanometers(math.Inf(1))
    if err == nil {
        t.Error("FromNanometers() with +Inf value should return error")
    }

    _, err = factory.FromNanometers(math.Inf(-1))
    if err == nil {
        t.Error("FromNanometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanometers(0)
    if err != nil {
        t.Errorf("FromNanometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthNanometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanometers() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrometers function
func TestLengthFactory_FromMicrometers(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrometers(100)
    if err != nil {
        t.Errorf("FromMicrometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthMicrometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrometers(math.NaN())
    if err == nil {
        t.Error("FromMicrometers() with NaN value should return error")
    }

    _, err = factory.FromMicrometers(math.Inf(1))
    if err == nil {
        t.Error("FromMicrometers() with +Inf value should return error")
    }

    _, err = factory.FromMicrometers(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrometers(0)
    if err != nil {
        t.Errorf("FromMicrometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthMicrometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrometers() with zero value = %v, want 0", converted)
    }
}
// Test FromMillimeters function
func TestLengthFactory_FromMillimeters(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillimeters(100)
    if err != nil {
        t.Errorf("FromMillimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthMillimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillimeters(math.NaN())
    if err == nil {
        t.Error("FromMillimeters() with NaN value should return error")
    }

    _, err = factory.FromMillimeters(math.Inf(1))
    if err == nil {
        t.Error("FromMillimeters() with +Inf value should return error")
    }

    _, err = factory.FromMillimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromMillimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillimeters(0)
    if err != nil {
        t.Errorf("FromMillimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthMillimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromCentimeters function
func TestLengthFactory_FromCentimeters(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentimeters(100)
    if err != nil {
        t.Errorf("FromCentimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentimeters(math.NaN())
    if err == nil {
        t.Error("FromCentimeters() with NaN value should return error")
    }

    _, err = factory.FromCentimeters(math.Inf(1))
    if err == nil {
        t.Error("FromCentimeters() with +Inf value should return error")
    }

    _, err = factory.FromCentimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromCentimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentimeters(0)
    if err != nil {
        t.Errorf("FromCentimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromDecimeters function
func TestLengthFactory_FromDecimeters(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecimeters(100)
    if err != nil {
        t.Errorf("FromDecimeters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthDecimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecimeters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecimeters(math.NaN())
    if err == nil {
        t.Error("FromDecimeters() with NaN value should return error")
    }

    _, err = factory.FromDecimeters(math.Inf(1))
    if err == nil {
        t.Error("FromDecimeters() with +Inf value should return error")
    }

    _, err = factory.FromDecimeters(math.Inf(-1))
    if err == nil {
        t.Error("FromDecimeters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecimeters(0)
    if err != nil {
        t.Errorf("FromDecimeters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthDecimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecimeters() with zero value = %v, want 0", converted)
    }
}
// Test FromDecameters function
func TestLengthFactory_FromDecameters(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDecameters(100)
    if err != nil {
        t.Errorf("FromDecameters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthDecameter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDecameters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDecameters(math.NaN())
    if err == nil {
        t.Error("FromDecameters() with NaN value should return error")
    }

    _, err = factory.FromDecameters(math.Inf(1))
    if err == nil {
        t.Error("FromDecameters() with +Inf value should return error")
    }

    _, err = factory.FromDecameters(math.Inf(-1))
    if err == nil {
        t.Error("FromDecameters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDecameters(0)
    if err != nil {
        t.Errorf("FromDecameters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthDecameter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDecameters() with zero value = %v, want 0", converted)
    }
}
// Test FromHectometers function
func TestLengthFactory_FromHectometers(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromHectometers(100)
    if err != nil {
        t.Errorf("FromHectometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthHectometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromHectometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromHectometers(math.NaN())
    if err == nil {
        t.Error("FromHectometers() with NaN value should return error")
    }

    _, err = factory.FromHectometers(math.Inf(1))
    if err == nil {
        t.Error("FromHectometers() with +Inf value should return error")
    }

    _, err = factory.FromHectometers(math.Inf(-1))
    if err == nil {
        t.Error("FromHectometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromHectometers(0)
    if err != nil {
        t.Errorf("FromHectometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthHectometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromHectometers() with zero value = %v, want 0", converted)
    }
}
// Test FromKilometers function
func TestLengthFactory_FromKilometers(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilometers(100)
    if err != nil {
        t.Errorf("FromKilometers() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthKilometer)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilometers() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilometers(math.NaN())
    if err == nil {
        t.Error("FromKilometers() with NaN value should return error")
    }

    _, err = factory.FromKilometers(math.Inf(1))
    if err == nil {
        t.Error("FromKilometers() with +Inf value should return error")
    }

    _, err = factory.FromKilometers(math.Inf(-1))
    if err == nil {
        t.Error("FromKilometers() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilometers(0)
    if err != nil {
        t.Errorf("FromKilometers() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthKilometer)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilometers() with zero value = %v, want 0", converted)
    }
}
// Test FromMegameters function
func TestLengthFactory_FromMegameters(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegameters(100)
    if err != nil {
        t.Errorf("FromMegameters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthMegameter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegameters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegameters(math.NaN())
    if err == nil {
        t.Error("FromMegameters() with NaN value should return error")
    }

    _, err = factory.FromMegameters(math.Inf(1))
    if err == nil {
        t.Error("FromMegameters() with +Inf value should return error")
    }

    _, err = factory.FromMegameters(math.Inf(-1))
    if err == nil {
        t.Error("FromMegameters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegameters(0)
    if err != nil {
        t.Errorf("FromMegameters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthMegameter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegameters() with zero value = %v, want 0", converted)
    }
}
// Test FromGigameters function
func TestLengthFactory_FromGigameters(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGigameters(100)
    if err != nil {
        t.Errorf("FromGigameters() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthGigameter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGigameters() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGigameters(math.NaN())
    if err == nil {
        t.Error("FromGigameters() with NaN value should return error")
    }

    _, err = factory.FromGigameters(math.Inf(1))
    if err == nil {
        t.Error("FromGigameters() with +Inf value should return error")
    }

    _, err = factory.FromGigameters(math.Inf(-1))
    if err == nil {
        t.Error("FromGigameters() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGigameters(0)
    if err != nil {
        t.Errorf("FromGigameters() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthGigameter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGigameters() with zero value = %v, want 0", converted)
    }
}
// Test FromKiloyards function
func TestLengthFactory_FromKiloyards(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKiloyards(100)
    if err != nil {
        t.Errorf("FromKiloyards() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthKiloyard)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKiloyards() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKiloyards(math.NaN())
    if err == nil {
        t.Error("FromKiloyards() with NaN value should return error")
    }

    _, err = factory.FromKiloyards(math.Inf(1))
    if err == nil {
        t.Error("FromKiloyards() with +Inf value should return error")
    }

    _, err = factory.FromKiloyards(math.Inf(-1))
    if err == nil {
        t.Error("FromKiloyards() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKiloyards(0)
    if err != nil {
        t.Errorf("FromKiloyards() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthKiloyard)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKiloyards() with zero value = %v, want 0", converted)
    }
}
// Test FromKilofeet function
func TestLengthFactory_FromKilofeet(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilofeet(100)
    if err != nil {
        t.Errorf("FromKilofeet() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthKilofoot)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilofeet() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilofeet(math.NaN())
    if err == nil {
        t.Error("FromKilofeet() with NaN value should return error")
    }

    _, err = factory.FromKilofeet(math.Inf(1))
    if err == nil {
        t.Error("FromKilofeet() with +Inf value should return error")
    }

    _, err = factory.FromKilofeet(math.Inf(-1))
    if err == nil {
        t.Error("FromKilofeet() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilofeet(0)
    if err != nil {
        t.Errorf("FromKilofeet() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthKilofoot)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilofeet() with zero value = %v, want 0", converted)
    }
}
// Test FromKiloparsecs function
func TestLengthFactory_FromKiloparsecs(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKiloparsecs(100)
    if err != nil {
        t.Errorf("FromKiloparsecs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthKiloparsec)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKiloparsecs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKiloparsecs(math.NaN())
    if err == nil {
        t.Error("FromKiloparsecs() with NaN value should return error")
    }

    _, err = factory.FromKiloparsecs(math.Inf(1))
    if err == nil {
        t.Error("FromKiloparsecs() with +Inf value should return error")
    }

    _, err = factory.FromKiloparsecs(math.Inf(-1))
    if err == nil {
        t.Error("FromKiloparsecs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKiloparsecs(0)
    if err != nil {
        t.Errorf("FromKiloparsecs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthKiloparsec)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKiloparsecs() with zero value = %v, want 0", converted)
    }
}
// Test FromMegaparsecs function
func TestLengthFactory_FromMegaparsecs(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegaparsecs(100)
    if err != nil {
        t.Errorf("FromMegaparsecs() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthMegaparsec)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegaparsecs() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegaparsecs(math.NaN())
    if err == nil {
        t.Error("FromMegaparsecs() with NaN value should return error")
    }

    _, err = factory.FromMegaparsecs(math.Inf(1))
    if err == nil {
        t.Error("FromMegaparsecs() with +Inf value should return error")
    }

    _, err = factory.FromMegaparsecs(math.Inf(-1))
    if err == nil {
        t.Error("FromMegaparsecs() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegaparsecs(0)
    if err != nil {
        t.Errorf("FromMegaparsecs() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthMegaparsec)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegaparsecs() with zero value = %v, want 0", converted)
    }
}
// Test FromKilolightYears function
func TestLengthFactory_FromKilolightYears(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilolightYears(100)
    if err != nil {
        t.Errorf("FromKilolightYears() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthKilolightYear)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilolightYears() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilolightYears(math.NaN())
    if err == nil {
        t.Error("FromKilolightYears() with NaN value should return error")
    }

    _, err = factory.FromKilolightYears(math.Inf(1))
    if err == nil {
        t.Error("FromKilolightYears() with +Inf value should return error")
    }

    _, err = factory.FromKilolightYears(math.Inf(-1))
    if err == nil {
        t.Error("FromKilolightYears() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilolightYears(0)
    if err != nil {
        t.Errorf("FromKilolightYears() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthKilolightYear)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilolightYears() with zero value = %v, want 0", converted)
    }
}
// Test FromMegalightYears function
func TestLengthFactory_FromMegalightYears(t *testing.T) {
    factory := units.LengthFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegalightYears(100)
    if err != nil {
        t.Errorf("FromMegalightYears() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.LengthMegalightYear)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegalightYears() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegalightYears(math.NaN())
    if err == nil {
        t.Error("FromMegalightYears() with NaN value should return error")
    }

    _, err = factory.FromMegalightYears(math.Inf(1))
    if err == nil {
        t.Error("FromMegalightYears() with +Inf value should return error")
    }

    _, err = factory.FromMegalightYears(math.Inf(-1))
    if err == nil {
        t.Error("FromMegalightYears() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegalightYears(0)
    if err != nil {
        t.Errorf("FromMegalightYears() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.LengthMegalightYear)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegalightYears() with zero value = %v, want 0", converted)
    }
}

func TestLengthToString(t *testing.T) {
	factory := units.LengthFactory{}
	a, err := factory.CreateLength(45, units.LengthMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.LengthMeter, 2)
	expected := "45.00 " + units.GetLengthAbbreviation(units.LengthMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.LengthMeter, -1)
	expected = "45 " + units.GetLengthAbbreviation(units.LengthMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestLength_EqualityAndComparison(t *testing.T) {
	factory := units.LengthFactory{}
	a1, _ := factory.CreateLength(60, units.LengthMeter)
	a2, _ := factory.CreateLength(60, units.LengthMeter)
	a3, _ := factory.CreateLength(120, units.LengthMeter)

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

func TestLength_Arithmetic(t *testing.T) {
	factory := units.LengthFactory{}
	a1, _ := factory.CreateLength(30, units.LengthMeter)
	a2, _ := factory.CreateLength(45, units.LengthMeter)

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


func TestGetLengthAbbreviation(t *testing.T) {
    tests := []struct {
        name string
        unit units.LengthUnits
        want string
    }{
        {
            name: "Meter abbreviation",
            unit: units.LengthMeter,
            want: "m",
        },
        {
            name: "Mile abbreviation",
            unit: units.LengthMile,
            want: "mi",
        },
        {
            name: "Yard abbreviation",
            unit: units.LengthYard,
            want: "yd",
        },
        {
            name: "Foot abbreviation",
            unit: units.LengthFoot,
            want: "ft",
        },
        {
            name: "UsSurveyFoot abbreviation",
            unit: units.LengthUsSurveyFoot,
            want: "ftUS",
        },
        {
            name: "Inch abbreviation",
            unit: units.LengthInch,
            want: "in",
        },
        {
            name: "Mil abbreviation",
            unit: units.LengthMil,
            want: "mil",
        },
        {
            name: "NauticalMile abbreviation",
            unit: units.LengthNauticalMile,
            want: "NM",
        },
        {
            name: "Fathom abbreviation",
            unit: units.LengthFathom,
            want: "fathom",
        },
        {
            name: "Shackle abbreviation",
            unit: units.LengthShackle,
            want: "shackle",
        },
        {
            name: "Microinch abbreviation",
            unit: units.LengthMicroinch,
            want: "µin",
        },
        {
            name: "PrinterPoint abbreviation",
            unit: units.LengthPrinterPoint,
            want: "pt",
        },
        {
            name: "DtpPoint abbreviation",
            unit: units.LengthDtpPoint,
            want: "pt",
        },
        {
            name: "PrinterPica abbreviation",
            unit: units.LengthPrinterPica,
            want: "pica",
        },
        {
            name: "DtpPica abbreviation",
            unit: units.LengthDtpPica,
            want: "pica",
        },
        {
            name: "Twip abbreviation",
            unit: units.LengthTwip,
            want: "twip",
        },
        {
            name: "Hand abbreviation",
            unit: units.LengthHand,
            want: "h",
        },
        {
            name: "AstronomicalUnit abbreviation",
            unit: units.LengthAstronomicalUnit,
            want: "au",
        },
        {
            name: "Parsec abbreviation",
            unit: units.LengthParsec,
            want: "pc",
        },
        {
            name: "LightYear abbreviation",
            unit: units.LengthLightYear,
            want: "ly",
        },
        {
            name: "SolarRadius abbreviation",
            unit: units.LengthSolarRadius,
            want: "R⊙",
        },
        {
            name: "Chain abbreviation",
            unit: units.LengthChain,
            want: "ch",
        },
        {
            name: "Angstrom abbreviation",
            unit: units.LengthAngstrom,
            want: "Å",
        },
        {
            name: "DataMile abbreviation",
            unit: units.LengthDataMile,
            want: "DM",
        },
        {
            name: "Femtometer abbreviation",
            unit: units.LengthFemtometer,
            want: "fm",
        },
        {
            name: "Picometer abbreviation",
            unit: units.LengthPicometer,
            want: "pm",
        },
        {
            name: "Nanometer abbreviation",
            unit: units.LengthNanometer,
            want: "nm",
        },
        {
            name: "Micrometer abbreviation",
            unit: units.LengthMicrometer,
            want: "μm",
        },
        {
            name: "Millimeter abbreviation",
            unit: units.LengthMillimeter,
            want: "mm",
        },
        {
            name: "Centimeter abbreviation",
            unit: units.LengthCentimeter,
            want: "cm",
        },
        {
            name: "Decimeter abbreviation",
            unit: units.LengthDecimeter,
            want: "dm",
        },
        {
            name: "Decameter abbreviation",
            unit: units.LengthDecameter,
            want: "dam",
        },
        {
            name: "Hectometer abbreviation",
            unit: units.LengthHectometer,
            want: "hm",
        },
        {
            name: "Kilometer abbreviation",
            unit: units.LengthKilometer,
            want: "km",
        },
        {
            name: "Megameter abbreviation",
            unit: units.LengthMegameter,
            want: "Mm",
        },
        {
            name: "Gigameter abbreviation",
            unit: units.LengthGigameter,
            want: "Gm",
        },
        {
            name: "Kiloyard abbreviation",
            unit: units.LengthKiloyard,
            want: "kyd",
        },
        {
            name: "Kilofoot abbreviation",
            unit: units.LengthKilofoot,
            want: "kft",
        },
        {
            name: "Kiloparsec abbreviation",
            unit: units.LengthKiloparsec,
            want: "kpc",
        },
        {
            name: "Megaparsec abbreviation",
            unit: units.LengthMegaparsec,
            want: "Mpc",
        },
        {
            name: "KilolightYear abbreviation",
            unit: units.LengthKilolightYear,
            want: "kly",
        },
        {
            name: "MegalightYear abbreviation",
            unit: units.LengthMegalightYear,
            want: "Mly",
        },
        {
            name: "invalid unit",
            unit: units.LengthUnits("invalid"),
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := units.GetLengthAbbreviation(tt.unit)
            if got != tt.want {
                t.Errorf("GetLengthAbbreviation(%v) = %v, want %v", 
                    tt.unit, got, tt.want)
            }
        })
    }
}

func TestLength_String(t *testing.T) {
    factory := units.LengthFactory{}
    
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
            unit, err := factory.CreateLength(tt.value, units.LengthMeter)
            if err != nil {
                t.Errorf("Failed to create test unit: %v", err)
                return
            }

            got := unit.String()
            if !strings.HasPrefix(got, tt.want) {
                t.Errorf("Length.String() = %v, want %v", got, tt.want)
            }
        })
    }
}