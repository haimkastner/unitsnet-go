// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestAngleDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "Degree"}`
	
	factory := units.AngleDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.AngleDegree {
		t.Errorf("expected unit %v, got %v", units.AngleDegree, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "Degree"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestAngleDto_ToJSON(t *testing.T) {
	dto := units.AngleDto{
		Value: 45,
		Unit:  units.AngleDegree,
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
	if result["unit"].(string) != string(units.AngleDegree) {
		t.Errorf("expected unit %s, got %v", units.AngleDegree, result["unit"])
	}
}

func TestNewAngle_InvalidValue(t *testing.T) {
	factory := units.AngleFactory{}
	// NaN value should return an error.
	_, err := factory.CreateAngle(math.NaN(), units.AngleDegree)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateAngle(math.Inf(1), units.AngleDegree)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestAngleConversions(t *testing.T) {
	factory := units.AngleFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateAngle(180, units.AngleDegree)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to Radians.
		// No expected conversion value provided for Radians, verifying result is not NaN.
		result := a.Radians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Radians returned NaN")
		}
	}
	{
		// Test conversion to Degrees.
		// No expected conversion value provided for Degrees, verifying result is not NaN.
		result := a.Degrees()
		if math.IsNaN(result) {
			t.Errorf("conversion to Degrees returned NaN")
		}
	}
	{
		// Test conversion to Arcminutes.
		// No expected conversion value provided for Arcminutes, verifying result is not NaN.
		result := a.Arcminutes()
		if math.IsNaN(result) {
			t.Errorf("conversion to Arcminutes returned NaN")
		}
	}
	{
		// Test conversion to Arcseconds.
		// No expected conversion value provided for Arcseconds, verifying result is not NaN.
		result := a.Arcseconds()
		if math.IsNaN(result) {
			t.Errorf("conversion to Arcseconds returned NaN")
		}
	}
	{
		// Test conversion to Gradians.
		// No expected conversion value provided for Gradians, verifying result is not NaN.
		result := a.Gradians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Gradians returned NaN")
		}
	}
	{
		// Test conversion to NatoMils.
		// No expected conversion value provided for NatoMils, verifying result is not NaN.
		result := a.NatoMils()
		if math.IsNaN(result) {
			t.Errorf("conversion to NatoMils returned NaN")
		}
	}
	{
		// Test conversion to Revolutions.
		// No expected conversion value provided for Revolutions, verifying result is not NaN.
		result := a.Revolutions()
		if math.IsNaN(result) {
			t.Errorf("conversion to Revolutions returned NaN")
		}
	}
	{
		// Test conversion to Tilt.
		// No expected conversion value provided for Tilt, verifying result is not NaN.
		result := a.Tilt()
		if math.IsNaN(result) {
			t.Errorf("conversion to Tilt returned NaN")
		}
	}
	{
		// Test conversion to Nanoradians.
		// No expected conversion value provided for Nanoradians, verifying result is not NaN.
		result := a.Nanoradians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanoradians returned NaN")
		}
	}
	{
		// Test conversion to Microradians.
		// No expected conversion value provided for Microradians, verifying result is not NaN.
		result := a.Microradians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microradians returned NaN")
		}
	}
	{
		// Test conversion to Milliradians.
		// No expected conversion value provided for Milliradians, verifying result is not NaN.
		result := a.Milliradians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Milliradians returned NaN")
		}
	}
	{
		// Test conversion to Centiradians.
		// No expected conversion value provided for Centiradians, verifying result is not NaN.
		result := a.Centiradians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Centiradians returned NaN")
		}
	}
	{
		// Test conversion to Deciradians.
		// No expected conversion value provided for Deciradians, verifying result is not NaN.
		result := a.Deciradians()
		if math.IsNaN(result) {
			t.Errorf("conversion to Deciradians returned NaN")
		}
	}
	{
		// Test conversion to Nanodegrees.
		// No expected conversion value provided for Nanodegrees, verifying result is not NaN.
		result := a.Nanodegrees()
		if math.IsNaN(result) {
			t.Errorf("conversion to Nanodegrees returned NaN")
		}
	}
	{
		// Test conversion to Microdegrees.
		// No expected conversion value provided for Microdegrees, verifying result is not NaN.
		result := a.Microdegrees()
		if math.IsNaN(result) {
			t.Errorf("conversion to Microdegrees returned NaN")
		}
	}
	{
		// Test conversion to Millidegrees.
		// No expected conversion value provided for Millidegrees, verifying result is not NaN.
		result := a.Millidegrees()
		if math.IsNaN(result) {
			t.Errorf("conversion to Millidegrees returned NaN")
		}
	}
}

func TestAngle_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.AngleFactory{}
	a, err := factory.CreateAngle(90, units.AngleDegree)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.AngleDegree {
		t.Errorf("expected default unit Degree, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.AngleRadian
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.AngleDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.AngleDegree {
		t.Errorf("expected unit Degree, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestAngleFactory_FromDto(t *testing.T) {
    factory := units.AngleFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleDegree,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.AngleDto{
        Value: math.NaN(),
        Unit:  units.AngleDegree,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test Radian conversion
    radiansDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleRadian,
    }
    
    var radiansResult *units.Angle
    radiansResult, err = factory.FromDto(radiansDto)
    if err != nil {
        t.Errorf("FromDto() with Radian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = radiansResult.Convert(units.AngleRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Radian = %v, want %v", converted, 100)
    }
    // Test Degree conversion
    degreesDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleDegree,
    }
    
    var degreesResult *units.Angle
    degreesResult, err = factory.FromDto(degreesDto)
    if err != nil {
        t.Errorf("FromDto() with Degree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degreesResult.Convert(units.AngleDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Degree = %v, want %v", converted, 100)
    }
    // Test Arcminute conversion
    arcminutesDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleArcminute,
    }
    
    var arcminutesResult *units.Angle
    arcminutesResult, err = factory.FromDto(arcminutesDto)
    if err != nil {
        t.Errorf("FromDto() with Arcminute returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = arcminutesResult.Convert(units.AngleArcminute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Arcminute = %v, want %v", converted, 100)
    }
    // Test Arcsecond conversion
    arcsecondsDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleArcsecond,
    }
    
    var arcsecondsResult *units.Angle
    arcsecondsResult, err = factory.FromDto(arcsecondsDto)
    if err != nil {
        t.Errorf("FromDto() with Arcsecond returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = arcsecondsResult.Convert(units.AngleArcsecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Arcsecond = %v, want %v", converted, 100)
    }
    // Test Gradian conversion
    gradiansDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleGradian,
    }
    
    var gradiansResult *units.Angle
    gradiansResult, err = factory.FromDto(gradiansDto)
    if err != nil {
        t.Errorf("FromDto() with Gradian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gradiansResult.Convert(units.AngleGradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gradian = %v, want %v", converted, 100)
    }
    // Test NatoMil conversion
    nato_milsDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleNatoMil,
    }
    
    var nato_milsResult *units.Angle
    nato_milsResult, err = factory.FromDto(nato_milsDto)
    if err != nil {
        t.Errorf("FromDto() with NatoMil returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nato_milsResult.Convert(units.AngleNatoMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NatoMil = %v, want %v", converted, 100)
    }
    // Test Revolution conversion
    revolutionsDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleRevolution,
    }
    
    var revolutionsResult *units.Angle
    revolutionsResult, err = factory.FromDto(revolutionsDto)
    if err != nil {
        t.Errorf("FromDto() with Revolution returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = revolutionsResult.Convert(units.AngleRevolution)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Revolution = %v, want %v", converted, 100)
    }
    // Test Tilt conversion
    tiltDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleTilt,
    }
    
    var tiltResult *units.Angle
    tiltResult, err = factory.FromDto(tiltDto)
    if err != nil {
        t.Errorf("FromDto() with Tilt returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tiltResult.Convert(units.AngleTilt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Tilt = %v, want %v", converted, 100)
    }
    // Test Nanoradian conversion
    nanoradiansDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleNanoradian,
    }
    
    var nanoradiansResult *units.Angle
    nanoradiansResult, err = factory.FromDto(nanoradiansDto)
    if err != nil {
        t.Errorf("FromDto() with Nanoradian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoradiansResult.Convert(units.AngleNanoradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanoradian = %v, want %v", converted, 100)
    }
    // Test Microradian conversion
    microradiansDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleMicroradian,
    }
    
    var microradiansResult *units.Angle
    microradiansResult, err = factory.FromDto(microradiansDto)
    if err != nil {
        t.Errorf("FromDto() with Microradian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microradiansResult.Convert(units.AngleMicroradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microradian = %v, want %v", converted, 100)
    }
    // Test Milliradian conversion
    milliradiansDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleMilliradian,
    }
    
    var milliradiansResult *units.Angle
    milliradiansResult, err = factory.FromDto(milliradiansDto)
    if err != nil {
        t.Errorf("FromDto() with Milliradian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliradiansResult.Convert(units.AngleMilliradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliradian = %v, want %v", converted, 100)
    }
    // Test Centiradian conversion
    centiradiansDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleCentiradian,
    }
    
    var centiradiansResult *units.Angle
    centiradiansResult, err = factory.FromDto(centiradiansDto)
    if err != nil {
        t.Errorf("FromDto() with Centiradian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiradiansResult.Convert(units.AngleCentiradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centiradian = %v, want %v", converted, 100)
    }
    // Test Deciradian conversion
    deciradiansDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleDeciradian,
    }
    
    var deciradiansResult *units.Angle
    deciradiansResult, err = factory.FromDto(deciradiansDto)
    if err != nil {
        t.Errorf("FromDto() with Deciradian returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciradiansResult.Convert(units.AngleDeciradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Deciradian = %v, want %v", converted, 100)
    }
    // Test Nanodegree conversion
    nanodegreesDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleNanodegree,
    }
    
    var nanodegreesResult *units.Angle
    nanodegreesResult, err = factory.FromDto(nanodegreesDto)
    if err != nil {
        t.Errorf("FromDto() with Nanodegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanodegreesResult.Convert(units.AngleNanodegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanodegree = %v, want %v", converted, 100)
    }
    // Test Microdegree conversion
    microdegreesDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleMicrodegree,
    }
    
    var microdegreesResult *units.Angle
    microdegreesResult, err = factory.FromDto(microdegreesDto)
    if err != nil {
        t.Errorf("FromDto() with Microdegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microdegreesResult.Convert(units.AngleMicrodegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microdegree = %v, want %v", converted, 100)
    }
    // Test Millidegree conversion
    millidegreesDto := units.AngleDto{
        Value: 100,
        Unit:  units.AngleMillidegree,
    }
    
    var millidegreesResult *units.Angle
    millidegreesResult, err = factory.FromDto(millidegreesDto)
    if err != nil {
        t.Errorf("FromDto() with Millidegree returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millidegreesResult.Convert(units.AngleMillidegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millidegree = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.AngleDto{
        Value: 0,
        Unit:  units.AngleDegree,
    }
    
    var zeroResult *units.Angle
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestAngleFactory_FromDtoJSON(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "Degree"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "Degree"}`)
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
    nanJSON, _ := json.Marshal(units.AngleDto{
        Value: nanValue,
        Unit:  units.AngleDegree,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with Radian unit
    radiansJSON := []byte(`{"value": 100, "unit": "Radian"}`)
    radiansResult, err := factory.FromDtoJSON(radiansJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Radian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = radiansResult.Convert(units.AngleRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Radian = %v, want %v", converted, 100)
    }
    // Test JSON with Degree unit
    degreesJSON := []byte(`{"value": 100, "unit": "Degree"}`)
    degreesResult, err := factory.FromDtoJSON(degreesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Degree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = degreesResult.Convert(units.AngleDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Degree = %v, want %v", converted, 100)
    }
    // Test JSON with Arcminute unit
    arcminutesJSON := []byte(`{"value": 100, "unit": "Arcminute"}`)
    arcminutesResult, err := factory.FromDtoJSON(arcminutesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Arcminute unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = arcminutesResult.Convert(units.AngleArcminute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Arcminute = %v, want %v", converted, 100)
    }
    // Test JSON with Arcsecond unit
    arcsecondsJSON := []byte(`{"value": 100, "unit": "Arcsecond"}`)
    arcsecondsResult, err := factory.FromDtoJSON(arcsecondsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Arcsecond unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = arcsecondsResult.Convert(units.AngleArcsecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Arcsecond = %v, want %v", converted, 100)
    }
    // Test JSON with Gradian unit
    gradiansJSON := []byte(`{"value": 100, "unit": "Gradian"}`)
    gradiansResult, err := factory.FromDtoJSON(gradiansJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Gradian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = gradiansResult.Convert(units.AngleGradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Gradian = %v, want %v", converted, 100)
    }
    // Test JSON with NatoMil unit
    nato_milsJSON := []byte(`{"value": 100, "unit": "NatoMil"}`)
    nato_milsResult, err := factory.FromDtoJSON(nato_milsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NatoMil unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nato_milsResult.Convert(units.AngleNatoMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NatoMil = %v, want %v", converted, 100)
    }
    // Test JSON with Revolution unit
    revolutionsJSON := []byte(`{"value": 100, "unit": "Revolution"}`)
    revolutionsResult, err := factory.FromDtoJSON(revolutionsJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Revolution unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = revolutionsResult.Convert(units.AngleRevolution)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Revolution = %v, want %v", converted, 100)
    }
    // Test JSON with Tilt unit
    tiltJSON := []byte(`{"value": 100, "unit": "Tilt"}`)
    tiltResult, err := factory.FromDtoJSON(tiltJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Tilt unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = tiltResult.Convert(units.AngleTilt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Tilt = %v, want %v", converted, 100)
    }
    // Test JSON with Nanoradian unit
    nanoradiansJSON := []byte(`{"value": 100, "unit": "Nanoradian"}`)
    nanoradiansResult, err := factory.FromDtoJSON(nanoradiansJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanoradian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanoradiansResult.Convert(units.AngleNanoradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanoradian = %v, want %v", converted, 100)
    }
    // Test JSON with Microradian unit
    microradiansJSON := []byte(`{"value": 100, "unit": "Microradian"}`)
    microradiansResult, err := factory.FromDtoJSON(microradiansJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microradian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microradiansResult.Convert(units.AngleMicroradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microradian = %v, want %v", converted, 100)
    }
    // Test JSON with Milliradian unit
    milliradiansJSON := []byte(`{"value": 100, "unit": "Milliradian"}`)
    milliradiansResult, err := factory.FromDtoJSON(milliradiansJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Milliradian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliradiansResult.Convert(units.AngleMilliradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Milliradian = %v, want %v", converted, 100)
    }
    // Test JSON with Centiradian unit
    centiradiansJSON := []byte(`{"value": 100, "unit": "Centiradian"}`)
    centiradiansResult, err := factory.FromDtoJSON(centiradiansJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Centiradian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = centiradiansResult.Convert(units.AngleCentiradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Centiradian = %v, want %v", converted, 100)
    }
    // Test JSON with Deciradian unit
    deciradiansJSON := []byte(`{"value": 100, "unit": "Deciradian"}`)
    deciradiansResult, err := factory.FromDtoJSON(deciradiansJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Deciradian unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = deciradiansResult.Convert(units.AngleDeciradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Deciradian = %v, want %v", converted, 100)
    }
    // Test JSON with Nanodegree unit
    nanodegreesJSON := []byte(`{"value": 100, "unit": "Nanodegree"}`)
    nanodegreesResult, err := factory.FromDtoJSON(nanodegreesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Nanodegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanodegreesResult.Convert(units.AngleNanodegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Nanodegree = %v, want %v", converted, 100)
    }
    // Test JSON with Microdegree unit
    microdegreesJSON := []byte(`{"value": 100, "unit": "Microdegree"}`)
    microdegreesResult, err := factory.FromDtoJSON(microdegreesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Microdegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microdegreesResult.Convert(units.AngleMicrodegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Microdegree = %v, want %v", converted, 100)
    }
    // Test JSON with Millidegree unit
    millidegreesJSON := []byte(`{"value": 100, "unit": "Millidegree"}`)
    millidegreesResult, err := factory.FromDtoJSON(millidegreesJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with Millidegree unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = millidegreesResult.Convert(units.AngleMillidegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for Millidegree = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "Degree"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromRadians function
func TestAngleFactory_FromRadians(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRadians(100)
    if err != nil {
        t.Errorf("FromRadians() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleRadian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromRadians() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromRadians(math.NaN())
    if err == nil {
        t.Error("FromRadians() with NaN value should return error")
    }

    _, err = factory.FromRadians(math.Inf(1))
    if err == nil {
        t.Error("FromRadians() with +Inf value should return error")
    }

    _, err = factory.FromRadians(math.Inf(-1))
    if err == nil {
        t.Error("FromRadians() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromRadians(0)
    if err != nil {
        t.Errorf("FromRadians() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleRadian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRadians() with zero value = %v, want 0", converted)
    }
}
// Test FromDegrees function
func TestAngleFactory_FromDegrees(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDegrees(100)
    if err != nil {
        t.Errorf("FromDegrees() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleDegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDegrees() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDegrees(math.NaN())
    if err == nil {
        t.Error("FromDegrees() with NaN value should return error")
    }

    _, err = factory.FromDegrees(math.Inf(1))
    if err == nil {
        t.Error("FromDegrees() with +Inf value should return error")
    }

    _, err = factory.FromDegrees(math.Inf(-1))
    if err == nil {
        t.Error("FromDegrees() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDegrees(0)
    if err != nil {
        t.Errorf("FromDegrees() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleDegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDegrees() with zero value = %v, want 0", converted)
    }
}
// Test FromArcminutes function
func TestAngleFactory_FromArcminutes(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromArcminutes(100)
    if err != nil {
        t.Errorf("FromArcminutes() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleArcminute)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromArcminutes() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromArcminutes(math.NaN())
    if err == nil {
        t.Error("FromArcminutes() with NaN value should return error")
    }

    _, err = factory.FromArcminutes(math.Inf(1))
    if err == nil {
        t.Error("FromArcminutes() with +Inf value should return error")
    }

    _, err = factory.FromArcminutes(math.Inf(-1))
    if err == nil {
        t.Error("FromArcminutes() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromArcminutes(0)
    if err != nil {
        t.Errorf("FromArcminutes() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleArcminute)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromArcminutes() with zero value = %v, want 0", converted)
    }
}
// Test FromArcseconds function
func TestAngleFactory_FromArcseconds(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromArcseconds(100)
    if err != nil {
        t.Errorf("FromArcseconds() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleArcsecond)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromArcseconds() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromArcseconds(math.NaN())
    if err == nil {
        t.Error("FromArcseconds() with NaN value should return error")
    }

    _, err = factory.FromArcseconds(math.Inf(1))
    if err == nil {
        t.Error("FromArcseconds() with +Inf value should return error")
    }

    _, err = factory.FromArcseconds(math.Inf(-1))
    if err == nil {
        t.Error("FromArcseconds() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromArcseconds(0)
    if err != nil {
        t.Errorf("FromArcseconds() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleArcsecond)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromArcseconds() with zero value = %v, want 0", converted)
    }
}
// Test FromGradians function
func TestAngleFactory_FromGradians(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromGradians(100)
    if err != nil {
        t.Errorf("FromGradians() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleGradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromGradians() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromGradians(math.NaN())
    if err == nil {
        t.Error("FromGradians() with NaN value should return error")
    }

    _, err = factory.FromGradians(math.Inf(1))
    if err == nil {
        t.Error("FromGradians() with +Inf value should return error")
    }

    _, err = factory.FromGradians(math.Inf(-1))
    if err == nil {
        t.Error("FromGradians() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromGradians(0)
    if err != nil {
        t.Errorf("FromGradians() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleGradian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromGradians() with zero value = %v, want 0", converted)
    }
}
// Test FromNatoMils function
func TestAngleFactory_FromNatoMils(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNatoMils(100)
    if err != nil {
        t.Errorf("FromNatoMils() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleNatoMil)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNatoMils() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNatoMils(math.NaN())
    if err == nil {
        t.Error("FromNatoMils() with NaN value should return error")
    }

    _, err = factory.FromNatoMils(math.Inf(1))
    if err == nil {
        t.Error("FromNatoMils() with +Inf value should return error")
    }

    _, err = factory.FromNatoMils(math.Inf(-1))
    if err == nil {
        t.Error("FromNatoMils() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNatoMils(0)
    if err != nil {
        t.Errorf("FromNatoMils() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleNatoMil)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNatoMils() with zero value = %v, want 0", converted)
    }
}
// Test FromRevolutions function
func TestAngleFactory_FromRevolutions(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromRevolutions(100)
    if err != nil {
        t.Errorf("FromRevolutions() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleRevolution)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromRevolutions() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromRevolutions(math.NaN())
    if err == nil {
        t.Error("FromRevolutions() with NaN value should return error")
    }

    _, err = factory.FromRevolutions(math.Inf(1))
    if err == nil {
        t.Error("FromRevolutions() with +Inf value should return error")
    }

    _, err = factory.FromRevolutions(math.Inf(-1))
    if err == nil {
        t.Error("FromRevolutions() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromRevolutions(0)
    if err != nil {
        t.Errorf("FromRevolutions() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleRevolution)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromRevolutions() with zero value = %v, want 0", converted)
    }
}
// Test FromTilt function
func TestAngleFactory_FromTilt(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromTilt(100)
    if err != nil {
        t.Errorf("FromTilt() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleTilt)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromTilt() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromTilt(math.NaN())
    if err == nil {
        t.Error("FromTilt() with NaN value should return error")
    }

    _, err = factory.FromTilt(math.Inf(1))
    if err == nil {
        t.Error("FromTilt() with +Inf value should return error")
    }

    _, err = factory.FromTilt(math.Inf(-1))
    if err == nil {
        t.Error("FromTilt() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromTilt(0)
    if err != nil {
        t.Errorf("FromTilt() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleTilt)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromTilt() with zero value = %v, want 0", converted)
    }
}
// Test FromNanoradians function
func TestAngleFactory_FromNanoradians(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanoradians(100)
    if err != nil {
        t.Errorf("FromNanoradians() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleNanoradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanoradians() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanoradians(math.NaN())
    if err == nil {
        t.Error("FromNanoradians() with NaN value should return error")
    }

    _, err = factory.FromNanoradians(math.Inf(1))
    if err == nil {
        t.Error("FromNanoradians() with +Inf value should return error")
    }

    _, err = factory.FromNanoradians(math.Inf(-1))
    if err == nil {
        t.Error("FromNanoradians() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanoradians(0)
    if err != nil {
        t.Errorf("FromNanoradians() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleNanoradian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanoradians() with zero value = %v, want 0", converted)
    }
}
// Test FromMicroradians function
func TestAngleFactory_FromMicroradians(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicroradians(100)
    if err != nil {
        t.Errorf("FromMicroradians() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleMicroradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicroradians() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicroradians(math.NaN())
    if err == nil {
        t.Error("FromMicroradians() with NaN value should return error")
    }

    _, err = factory.FromMicroradians(math.Inf(1))
    if err == nil {
        t.Error("FromMicroradians() with +Inf value should return error")
    }

    _, err = factory.FromMicroradians(math.Inf(-1))
    if err == nil {
        t.Error("FromMicroradians() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicroradians(0)
    if err != nil {
        t.Errorf("FromMicroradians() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleMicroradian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicroradians() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliradians function
func TestAngleFactory_FromMilliradians(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliradians(100)
    if err != nil {
        t.Errorf("FromMilliradians() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleMilliradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliradians() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliradians(math.NaN())
    if err == nil {
        t.Error("FromMilliradians() with NaN value should return error")
    }

    _, err = factory.FromMilliradians(math.Inf(1))
    if err == nil {
        t.Error("FromMilliradians() with +Inf value should return error")
    }

    _, err = factory.FromMilliradians(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliradians() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliradians(0)
    if err != nil {
        t.Errorf("FromMilliradians() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleMilliradian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliradians() with zero value = %v, want 0", converted)
    }
}
// Test FromCentiradians function
func TestAngleFactory_FromCentiradians(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromCentiradians(100)
    if err != nil {
        t.Errorf("FromCentiradians() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleCentiradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromCentiradians() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromCentiradians(math.NaN())
    if err == nil {
        t.Error("FromCentiradians() with NaN value should return error")
    }

    _, err = factory.FromCentiradians(math.Inf(1))
    if err == nil {
        t.Error("FromCentiradians() with +Inf value should return error")
    }

    _, err = factory.FromCentiradians(math.Inf(-1))
    if err == nil {
        t.Error("FromCentiradians() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromCentiradians(0)
    if err != nil {
        t.Errorf("FromCentiradians() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleCentiradian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromCentiradians() with zero value = %v, want 0", converted)
    }
}
// Test FromDeciradians function
func TestAngleFactory_FromDeciradians(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromDeciradians(100)
    if err != nil {
        t.Errorf("FromDeciradians() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleDeciradian)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromDeciradians() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromDeciradians(math.NaN())
    if err == nil {
        t.Error("FromDeciradians() with NaN value should return error")
    }

    _, err = factory.FromDeciradians(math.Inf(1))
    if err == nil {
        t.Error("FromDeciradians() with +Inf value should return error")
    }

    _, err = factory.FromDeciradians(math.Inf(-1))
    if err == nil {
        t.Error("FromDeciradians() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromDeciradians(0)
    if err != nil {
        t.Errorf("FromDeciradians() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleDeciradian)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromDeciradians() with zero value = %v, want 0", converted)
    }
}
// Test FromNanodegrees function
func TestAngleFactory_FromNanodegrees(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanodegrees(100)
    if err != nil {
        t.Errorf("FromNanodegrees() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleNanodegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanodegrees() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanodegrees(math.NaN())
    if err == nil {
        t.Error("FromNanodegrees() with NaN value should return error")
    }

    _, err = factory.FromNanodegrees(math.Inf(1))
    if err == nil {
        t.Error("FromNanodegrees() with +Inf value should return error")
    }

    _, err = factory.FromNanodegrees(math.Inf(-1))
    if err == nil {
        t.Error("FromNanodegrees() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanodegrees(0)
    if err != nil {
        t.Errorf("FromNanodegrees() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleNanodegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanodegrees() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrodegrees function
func TestAngleFactory_FromMicrodegrees(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrodegrees(100)
    if err != nil {
        t.Errorf("FromMicrodegrees() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleMicrodegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrodegrees() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrodegrees(math.NaN())
    if err == nil {
        t.Error("FromMicrodegrees() with NaN value should return error")
    }

    _, err = factory.FromMicrodegrees(math.Inf(1))
    if err == nil {
        t.Error("FromMicrodegrees() with +Inf value should return error")
    }

    _, err = factory.FromMicrodegrees(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrodegrees() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrodegrees(0)
    if err != nil {
        t.Errorf("FromMicrodegrees() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleMicrodegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrodegrees() with zero value = %v, want 0", converted)
    }
}
// Test FromMillidegrees function
func TestAngleFactory_FromMillidegrees(t *testing.T) {
    factory := units.AngleFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMillidegrees(100)
    if err != nil {
        t.Errorf("FromMillidegrees() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.AngleMillidegree)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMillidegrees() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMillidegrees(math.NaN())
    if err == nil {
        t.Error("FromMillidegrees() with NaN value should return error")
    }

    _, err = factory.FromMillidegrees(math.Inf(1))
    if err == nil {
        t.Error("FromMillidegrees() with +Inf value should return error")
    }

    _, err = factory.FromMillidegrees(math.Inf(-1))
    if err == nil {
        t.Error("FromMillidegrees() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMillidegrees(0)
    if err != nil {
        t.Errorf("FromMillidegrees() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.AngleMillidegree)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMillidegrees() with zero value = %v, want 0", converted)
    }
}

func TestAngleToString(t *testing.T) {
	factory := units.AngleFactory{}
	a, err := factory.CreateAngle(45, units.AngleDegree)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.AngleDegree, 2)
	expected := "45.00 " + units.GetAngleAbbreviation(units.AngleDegree)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.AngleDegree, -1)
	expected = "45 " + units.GetAngleAbbreviation(units.AngleDegree)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestAngle_EqualityAndComparison(t *testing.T) {
	factory := units.AngleFactory{}
	a1, _ := factory.CreateAngle(60, units.AngleDegree)
	a2, _ := factory.CreateAngle(60, units.AngleDegree)
	a3, _ := factory.CreateAngle(120, units.AngleDegree)

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

func TestAngle_Arithmetic(t *testing.T) {
	factory := units.AngleFactory{}
	a1, _ := factory.CreateAngle(30, units.AngleDegree)
	a2, _ := factory.CreateAngle(45, units.AngleDegree)

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