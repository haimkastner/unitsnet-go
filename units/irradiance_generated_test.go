// Generated Code - DO NOT EDIT.

package units_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/haimkastner/unitsnet-go/units"

)

func TestIrradianceDtoFactory_FromJSON(t *testing.T) {
	validJSON := `{"value": 90, "unit": "WattPerSquareMeter"}`
	
	factory := units.IrradianceDtoFactory{}
	dto, err := factory.FromJSON([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dto.Value != 90 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}
	if dto.Unit != units.IrradianceWattPerSquareMeter {
		t.Errorf("expected unit %v, got %v", units.IrradianceWattPerSquareMeter, dto.Unit)
	}

	invalidJSON := `{"value": "ninety", "unit": "WattPerSquareMeter"}`

	_, err = factory.FromJSON([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error with invalid JSON, got nil")
	}
}

func TestIrradianceDto_ToJSON(t *testing.T) {
	dto := units.IrradianceDto{
		Value: 45,
		Unit:  units.IrradianceWattPerSquareMeter,
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
	if result["unit"].(string) != string(units.IrradianceWattPerSquareMeter) {
		t.Errorf("expected unit %s, got %v", units.IrradianceWattPerSquareMeter, result["unit"])
	}
}

func TestNewIrradiance_InvalidValue(t *testing.T) {
	factory := units.IrradianceFactory{}
	// NaN value should return an error.
	_, err := factory.CreateIrradiance(math.NaN(), units.IrradianceWattPerSquareMeter)
	if err == nil {
		t.Error("expected error for NaN value")
	}
	// Inf value should return an error.
	_, err = factory.CreateIrradiance(math.Inf(1), units.IrradianceWattPerSquareMeter)
	if err == nil {
		t.Error("expected error for Inf value")
	}
}

func TestIrradianceConversions(t *testing.T) {
	factory := units.IrradianceFactory{}
	// Creating a value of 180 in the base unit.
	a, err := factory.CreateIrradiance(180, units.IrradianceWattPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	{
		// Test conversion to WattsPerSquareMeter.
		// No expected conversion value provided for WattsPerSquareMeter, verifying result is not NaN.
		result := a.WattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to WattsPerSquareCentimeter.
		// No expected conversion value provided for WattsPerSquareCentimeter, verifying result is not NaN.
		result := a.WattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to WattsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to PicowattsPerSquareMeter.
		// No expected conversion value provided for PicowattsPerSquareMeter, verifying result is not NaN.
		result := a.PicowattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerSquareMeter.
		// No expected conversion value provided for NanowattsPerSquareMeter, verifying result is not NaN.
		result := a.NanowattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerSquareMeter.
		// No expected conversion value provided for MicrowattsPerSquareMeter, verifying result is not NaN.
		result := a.MicrowattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerSquareMeter.
		// No expected conversion value provided for MilliwattsPerSquareMeter, verifying result is not NaN.
		result := a.MilliwattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerSquareMeter.
		// No expected conversion value provided for KilowattsPerSquareMeter, verifying result is not NaN.
		result := a.KilowattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerSquareMeter.
		// No expected conversion value provided for MegawattsPerSquareMeter, verifying result is not NaN.
		result := a.MegawattsPerSquareMeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattsPerSquareMeter returned NaN")
		}
	}
	{
		// Test conversion to PicowattsPerSquareCentimeter.
		// No expected conversion value provided for PicowattsPerSquareCentimeter, verifying result is not NaN.
		result := a.PicowattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to PicowattsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to NanowattsPerSquareCentimeter.
		// No expected conversion value provided for NanowattsPerSquareCentimeter, verifying result is not NaN.
		result := a.NanowattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to NanowattsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MicrowattsPerSquareCentimeter.
		// No expected conversion value provided for MicrowattsPerSquareCentimeter, verifying result is not NaN.
		result := a.MicrowattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MicrowattsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MilliwattsPerSquareCentimeter.
		// No expected conversion value provided for MilliwattsPerSquareCentimeter, verifying result is not NaN.
		result := a.MilliwattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MilliwattsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to KilowattsPerSquareCentimeter.
		// No expected conversion value provided for KilowattsPerSquareCentimeter, verifying result is not NaN.
		result := a.KilowattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to KilowattsPerSquareCentimeter returned NaN")
		}
	}
	{
		// Test conversion to MegawattsPerSquareCentimeter.
		// No expected conversion value provided for MegawattsPerSquareCentimeter, verifying result is not NaN.
		result := a.MegawattsPerSquareCentimeter()
		if math.IsNaN(result) {
			t.Errorf("conversion to MegawattsPerSquareCentimeter returned NaN")
		}
	}
}

func TestIrradiance_ToDtoAndToDtoJSON(t *testing.T) {
	factory := units.IrradianceFactory{}
	a, err := factory.CreateIrradiance(90, units.IrradianceWattPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test default conversion (nil unit parameter should use base unit).
	dto := a.ToDto(nil)
	if dto.Unit != units.IrradianceWattPerSquareMeter {
		t.Errorf("expected default unit WattPerSquareMeter, got %v", dto.Unit)
	}
	if math.Abs(dto.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", dto.Value)
	}

	// Explicit conversion using first available method.
	holdUnit := units.IrradianceWattPerSquareMeter
	dto = a.ToDto(&holdUnit)

	jsonData, err := a.ToDtoJSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result units.IrradianceDto
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("failed unmarshalling JSON: %v", err)
	}
	if result.Unit != units.IrradianceWattPerSquareMeter {
		t.Errorf("expected unit WattPerSquareMeter, got %v", result.Unit)
	}
	if math.Abs(result.Value-90) > 1e-9 {
		t.Errorf("expected value 90, got %v", result.Value)
	}
}

func TestIrradianceFactory_FromDto(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error
    
    // Test valid base unit conversion
    baseDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradianceWattPerSquareMeter,
    }
    
    baseResult, err := factory.FromDto(baseDto)
    if err != nil {
        t.Errorf("FromDto() with base unit returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDto() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid values
    invalidDto := units.IrradianceDto{
        Value: math.NaN(),
        Unit:  units.IrradianceWattPerSquareMeter,
    }
    
    _, err = factory.FromDto(invalidDto)
    if err == nil {
        t.Error("FromDto() with NaN value should return error")
    }

	var converted float64
    // Test WattPerSquareMeter conversion
    watts_per_square_meterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradianceWattPerSquareMeter,
    }
    
    var watts_per_square_meterResult *units.Irradiance
    watts_per_square_meterResult, err = factory.FromDto(watts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_meterResult.Convert(units.IrradianceWattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test WattPerSquareCentimeter conversion
    watts_per_square_centimeterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradianceWattPerSquareCentimeter,
    }
    
    var watts_per_square_centimeterResult *units.Irradiance
    watts_per_square_centimeterResult, err = factory.FromDto(watts_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with WattPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_centimeterResult.Convert(units.IrradianceWattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test PicowattPerSquareMeter conversion
    picowatts_per_square_meterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradiancePicowattPerSquareMeter,
    }
    
    var picowatts_per_square_meterResult *units.Irradiance
    picowatts_per_square_meterResult, err = factory.FromDto(picowatts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with PicowattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowatts_per_square_meterResult.Convert(units.IrradiancePicowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test NanowattPerSquareMeter conversion
    nanowatts_per_square_meterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradianceNanowattPerSquareMeter,
    }
    
    var nanowatts_per_square_meterResult *units.Irradiance
    nanowatts_per_square_meterResult, err = factory.FromDto(nanowatts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with NanowattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_square_meterResult.Convert(units.IrradianceNanowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test MicrowattPerSquareMeter conversion
    microwatts_per_square_meterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradianceMicrowattPerSquareMeter,
    }
    
    var microwatts_per_square_meterResult *units.Irradiance
    microwatts_per_square_meterResult, err = factory.FromDto(microwatts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrowattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_square_meterResult.Convert(units.IrradianceMicrowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test MilliwattPerSquareMeter conversion
    milliwatts_per_square_meterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradianceMilliwattPerSquareMeter,
    }
    
    var milliwatts_per_square_meterResult *units.Irradiance
    milliwatts_per_square_meterResult, err = factory.FromDto(milliwatts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MilliwattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_square_meterResult.Convert(units.IrradianceMilliwattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test KilowattPerSquareMeter conversion
    kilowatts_per_square_meterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradianceKilowattPerSquareMeter,
    }
    
    var kilowatts_per_square_meterResult *units.Irradiance
    kilowatts_per_square_meterResult, err = factory.FromDto(kilowatts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_square_meterResult.Convert(units.IrradianceKilowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test MegawattPerSquareMeter conversion
    megawatts_per_square_meterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradianceMegawattPerSquareMeter,
    }
    
    var megawatts_per_square_meterResult *units.Irradiance
    megawatts_per_square_meterResult, err = factory.FromDto(megawatts_per_square_meterDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattPerSquareMeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_square_meterResult.Convert(units.IrradianceMegawattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test PicowattPerSquareCentimeter conversion
    picowatts_per_square_centimeterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradiancePicowattPerSquareCentimeter,
    }
    
    var picowatts_per_square_centimeterResult *units.Irradiance
    picowatts_per_square_centimeterResult, err = factory.FromDto(picowatts_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with PicowattPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowatts_per_square_centimeterResult.Convert(units.IrradiancePicowattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicowattPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test NanowattPerSquareCentimeter conversion
    nanowatts_per_square_centimeterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradianceNanowattPerSquareCentimeter,
    }
    
    var nanowatts_per_square_centimeterResult *units.Irradiance
    nanowatts_per_square_centimeterResult, err = factory.FromDto(nanowatts_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with NanowattPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_square_centimeterResult.Convert(units.IrradianceNanowattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test MicrowattPerSquareCentimeter conversion
    microwatts_per_square_centimeterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradianceMicrowattPerSquareCentimeter,
    }
    
    var microwatts_per_square_centimeterResult *units.Irradiance
    microwatts_per_square_centimeterResult, err = factory.FromDto(microwatts_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MicrowattPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_square_centimeterResult.Convert(units.IrradianceMicrowattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test MilliwattPerSquareCentimeter conversion
    milliwatts_per_square_centimeterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradianceMilliwattPerSquareCentimeter,
    }
    
    var milliwatts_per_square_centimeterResult *units.Irradiance
    milliwatts_per_square_centimeterResult, err = factory.FromDto(milliwatts_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MilliwattPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_square_centimeterResult.Convert(units.IrradianceMilliwattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test KilowattPerSquareCentimeter conversion
    kilowatts_per_square_centimeterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradianceKilowattPerSquareCentimeter,
    }
    
    var kilowatts_per_square_centimeterResult *units.Irradiance
    kilowatts_per_square_centimeterResult, err = factory.FromDto(kilowatts_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with KilowattPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_square_centimeterResult.Convert(units.IrradianceKilowattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test MegawattPerSquareCentimeter conversion
    megawatts_per_square_centimeterDto := units.IrradianceDto{
        Value: 100,
        Unit:  units.IrradianceMegawattPerSquareCentimeter,
    }
    
    var megawatts_per_square_centimeterResult *units.Irradiance
    megawatts_per_square_centimeterResult, err = factory.FromDto(megawatts_per_square_centimeterDto)
    if err != nil {
        t.Errorf("FromDto() with MegawattPerSquareCentimeter returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_square_centimeterResult.Convert(units.IrradianceMegawattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerSquareCentimeter = %v, want %v", converted, 100)
    }

    // Test zero value
    zeroDto := units.IrradianceDto{
        Value: 0,
        Unit:  units.IrradianceWattPerSquareMeter,
    }
    
    var zeroResult *units.Irradiance
    zeroResult, err = factory.FromDto(zeroDto)
    if err != nil {
        t.Errorf("FromDto() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDto() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}

func TestIrradianceFactory_FromDtoJSON(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

	var converted float64

    // Test valid JSON with base unit
    validJSON := []byte(`{"value": 100, "unit": "WattPerSquareMeter"}`)
    baseResult, err := factory.FromDtoJSON(validJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with valid JSON returned error: %v", err)
    }
    if baseResult.BaseValue() != 100 {
        t.Errorf("FromDtoJSON() with base unit = %v, want %v", baseResult.BaseValue(), 100)
    }

    // Test invalid JSON format
    invalidJSON := []byte(`{"value": "not a number", "unit": "WattPerSquareMeter"}`)
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
    nanJSON, _ := json.Marshal(units.IrradianceDto{
        Value: nanValue,
        Unit:  units.IrradianceWattPerSquareMeter,
    })
    _, err = factory.FromDtoJSON(nanJSON)
    if err == nil {
        t.Error("FromDtoJSON() with NaN value should return error")
    }
    // Test JSON with WattPerSquareMeter unit
    watts_per_square_meterJSON := []byte(`{"value": 100, "unit": "WattPerSquareMeter"}`)
    watts_per_square_meterResult, err := factory.FromDtoJSON(watts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_meterResult.Convert(units.IrradianceWattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with WattPerSquareCentimeter unit
    watts_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "WattPerSquareCentimeter"}`)
    watts_per_square_centimeterResult, err := factory.FromDtoJSON(watts_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with WattPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = watts_per_square_centimeterResult.Convert(units.IrradianceWattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for WattPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with PicowattPerSquareMeter unit
    picowatts_per_square_meterJSON := []byte(`{"value": 100, "unit": "PicowattPerSquareMeter"}`)
    picowatts_per_square_meterResult, err := factory.FromDtoJSON(picowatts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicowattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowatts_per_square_meterResult.Convert(units.IrradiancePicowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with NanowattPerSquareMeter unit
    nanowatts_per_square_meterJSON := []byte(`{"value": 100, "unit": "NanowattPerSquareMeter"}`)
    nanowatts_per_square_meterResult, err := factory.FromDtoJSON(nanowatts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanowattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_square_meterResult.Convert(units.IrradianceNanowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrowattPerSquareMeter unit
    microwatts_per_square_meterJSON := []byte(`{"value": 100, "unit": "MicrowattPerSquareMeter"}`)
    microwatts_per_square_meterResult, err := factory.FromDtoJSON(microwatts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrowattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_square_meterResult.Convert(units.IrradianceMicrowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilliwattPerSquareMeter unit
    milliwatts_per_square_meterJSON := []byte(`{"value": 100, "unit": "MilliwattPerSquareMeter"}`)
    milliwatts_per_square_meterResult, err := factory.FromDtoJSON(milliwatts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliwattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_square_meterResult.Convert(units.IrradianceMilliwattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattPerSquareMeter unit
    kilowatts_per_square_meterJSON := []byte(`{"value": 100, "unit": "KilowattPerSquareMeter"}`)
    kilowatts_per_square_meterResult, err := factory.FromDtoJSON(kilowatts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_square_meterResult.Convert(units.IrradianceKilowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattPerSquareMeter unit
    megawatts_per_square_meterJSON := []byte(`{"value": 100, "unit": "MegawattPerSquareMeter"}`)
    megawatts_per_square_meterResult, err := factory.FromDtoJSON(megawatts_per_square_meterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattPerSquareMeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_square_meterResult.Convert(units.IrradianceMegawattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerSquareMeter = %v, want %v", converted, 100)
    }
    // Test JSON with PicowattPerSquareCentimeter unit
    picowatts_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "PicowattPerSquareCentimeter"}`)
    picowatts_per_square_centimeterResult, err := factory.FromDtoJSON(picowatts_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with PicowattPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = picowatts_per_square_centimeterResult.Convert(units.IrradiancePicowattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for PicowattPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with NanowattPerSquareCentimeter unit
    nanowatts_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "NanowattPerSquareCentimeter"}`)
    nanowatts_per_square_centimeterResult, err := factory.FromDtoJSON(nanowatts_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with NanowattPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = nanowatts_per_square_centimeterResult.Convert(units.IrradianceNanowattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for NanowattPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MicrowattPerSquareCentimeter unit
    microwatts_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "MicrowattPerSquareCentimeter"}`)
    microwatts_per_square_centimeterResult, err := factory.FromDtoJSON(microwatts_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MicrowattPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = microwatts_per_square_centimeterResult.Convert(units.IrradianceMicrowattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MicrowattPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MilliwattPerSquareCentimeter unit
    milliwatts_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "MilliwattPerSquareCentimeter"}`)
    milliwatts_per_square_centimeterResult, err := factory.FromDtoJSON(milliwatts_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MilliwattPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = milliwatts_per_square_centimeterResult.Convert(units.IrradianceMilliwattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MilliwattPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with KilowattPerSquareCentimeter unit
    kilowatts_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "KilowattPerSquareCentimeter"}`)
    kilowatts_per_square_centimeterResult, err := factory.FromDtoJSON(kilowatts_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with KilowattPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = kilowatts_per_square_centimeterResult.Convert(units.IrradianceKilowattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for KilowattPerSquareCentimeter = %v, want %v", converted, 100)
    }
    // Test JSON with MegawattPerSquareCentimeter unit
    megawatts_per_square_centimeterJSON := []byte(`{"value": 100, "unit": "MegawattPerSquareCentimeter"}`)
    megawatts_per_square_centimeterResult, err := factory.FromDtoJSON(megawatts_per_square_centimeterJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with MegawattPerSquareCentimeter unit returned error: %v", err)
    }
    
    // Convert back to original unit and compare
    converted = megawatts_per_square_centimeterResult.Convert(units.IrradianceMegawattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("Round-trip conversion for MegawattPerSquareCentimeter = %v, want %v", converted, 100)
    }

    // Test zero value JSON
    zeroJSON := []byte(`{"value": 0, "unit": "WattPerSquareMeter"}`)
    zeroResult, err := factory.FromDtoJSON(zeroJSON)
    if err != nil {
        t.Errorf("FromDtoJSON() with zero value returned error: %v", err)
    }
    if zeroResult.BaseValue() != 0 {
        t.Errorf("FromDtoJSON() with zero value = %v, want 0", zeroResult.BaseValue())
    }
}
// Test FromWattsPerSquareMeter function
func TestIrradianceFactory_FromWattsPerSquareMeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromWattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradianceWattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromWattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromWattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromWattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradianceWattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromWattsPerSquareCentimeter function
func TestIrradianceFactory_FromWattsPerSquareCentimeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromWattsPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromWattsPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradianceWattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromWattsPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromWattsPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromWattsPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromWattsPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromWattsPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromWattsPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromWattsPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromWattsPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromWattsPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradianceWattPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromWattsPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicowattsPerSquareMeter function
func TestIrradianceFactory_FromPicowattsPerSquareMeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicowattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromPicowattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradiancePicowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicowattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicowattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromPicowattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromPicowattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromPicowattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromPicowattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicowattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicowattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromPicowattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradiancePicowattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicowattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanowattsPerSquareMeter function
func TestIrradianceFactory_FromNanowattsPerSquareMeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanowattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromNanowattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradianceNanowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanowattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanowattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromNanowattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromNanowattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromNanowattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromNanowattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanowattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanowattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromNanowattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradianceNanowattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanowattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrowattsPerSquareMeter function
func TestIrradianceFactory_FromMicrowattsPerSquareMeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrowattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromMicrowattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradianceMicrowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrowattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrowattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromMicrowattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromMicrowattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrowattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromMicrowattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrowattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrowattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromMicrowattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradianceMicrowattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrowattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwattsPerSquareMeter function
func TestIrradianceFactory_FromMilliwattsPerSquareMeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromMilliwattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradianceMilliwattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromMilliwattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromMilliwattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromMilliwattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromMilliwattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradianceMilliwattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattsPerSquareMeter function
func TestIrradianceFactory_FromKilowattsPerSquareMeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromKilowattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradianceKilowattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromKilowattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromKilowattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromKilowattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromKilowattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradianceKilowattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattsPerSquareMeter function
func TestIrradianceFactory_FromMegawattsPerSquareMeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattsPerSquareMeter(100)
    if err != nil {
        t.Errorf("FromMegawattsPerSquareMeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradianceMegawattPerSquareMeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattsPerSquareMeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattsPerSquareMeter(math.NaN())
    if err == nil {
        t.Error("FromMegawattsPerSquareMeter() with NaN value should return error")
    }

    _, err = factory.FromMegawattsPerSquareMeter(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattsPerSquareMeter() with +Inf value should return error")
    }

    _, err = factory.FromMegawattsPerSquareMeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattsPerSquareMeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattsPerSquareMeter(0)
    if err != nil {
        t.Errorf("FromMegawattsPerSquareMeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradianceMegawattPerSquareMeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattsPerSquareMeter() with zero value = %v, want 0", converted)
    }
}
// Test FromPicowattsPerSquareCentimeter function
func TestIrradianceFactory_FromPicowattsPerSquareCentimeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromPicowattsPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromPicowattsPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradiancePicowattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromPicowattsPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromPicowattsPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromPicowattsPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromPicowattsPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromPicowattsPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromPicowattsPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromPicowattsPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromPicowattsPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromPicowattsPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradiancePicowattPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromPicowattsPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromNanowattsPerSquareCentimeter function
func TestIrradianceFactory_FromNanowattsPerSquareCentimeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromNanowattsPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromNanowattsPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradianceNanowattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromNanowattsPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromNanowattsPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromNanowattsPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromNanowattsPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromNanowattsPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromNanowattsPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromNanowattsPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromNanowattsPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromNanowattsPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradianceNanowattPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromNanowattsPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMicrowattsPerSquareCentimeter function
func TestIrradianceFactory_FromMicrowattsPerSquareCentimeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMicrowattsPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromMicrowattsPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradianceMicrowattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMicrowattsPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMicrowattsPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMicrowattsPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMicrowattsPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMicrowattsPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMicrowattsPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMicrowattsPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMicrowattsPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromMicrowattsPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradianceMicrowattPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMicrowattsPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMilliwattsPerSquareCentimeter function
func TestIrradianceFactory_FromMilliwattsPerSquareCentimeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMilliwattsPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromMilliwattsPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradianceMilliwattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMilliwattsPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMilliwattsPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMilliwattsPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMilliwattsPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMilliwattsPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMilliwattsPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMilliwattsPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMilliwattsPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromMilliwattsPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradianceMilliwattPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMilliwattsPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromKilowattsPerSquareCentimeter function
func TestIrradianceFactory_FromKilowattsPerSquareCentimeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromKilowattsPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromKilowattsPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradianceKilowattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromKilowattsPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromKilowattsPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromKilowattsPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromKilowattsPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromKilowattsPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromKilowattsPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromKilowattsPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromKilowattsPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromKilowattsPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradianceKilowattPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromKilowattsPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}
// Test FromMegawattsPerSquareCentimeter function
func TestIrradianceFactory_FromMegawattsPerSquareCentimeter(t *testing.T) {
    factory := units.IrradianceFactory{}
    var err error

    // Test valid value
    result, err := factory.FromMegawattsPerSquareCentimeter(100)
    if err != nil {
        t.Errorf("FromMegawattsPerSquareCentimeter() returned error: %v", err)
    }
    
    // Convert back and verify
    converted := result.Convert(units.IrradianceMegawattPerSquareCentimeter)
    if math.Abs(converted - 100) > 1e-6 {
        t.Errorf("FromMegawattsPerSquareCentimeter() round-trip = %v, want %v", converted, 100)
    }

    // Test invalid values
    _, err = factory.FromMegawattsPerSquareCentimeter(math.NaN())
    if err == nil {
        t.Error("FromMegawattsPerSquareCentimeter() with NaN value should return error")
    }

    _, err = factory.FromMegawattsPerSquareCentimeter(math.Inf(1))
    if err == nil {
        t.Error("FromMegawattsPerSquareCentimeter() with +Inf value should return error")
    }

    _, err = factory.FromMegawattsPerSquareCentimeter(math.Inf(-1))
    if err == nil {
        t.Error("FromMegawattsPerSquareCentimeter() with -Inf value should return error")
    }

    // Test zero value
    zeroResult, err := factory.FromMegawattsPerSquareCentimeter(0)
    if err != nil {
        t.Errorf("FromMegawattsPerSquareCentimeter() with zero value returned error: %v", err)
    }
    converted = zeroResult.Convert(units.IrradianceMegawattPerSquareCentimeter)
    if math.Abs(converted) > 1e-6 {
        t.Errorf("FromMegawattsPerSquareCentimeter() with zero value = %v, want 0", converted)
    }
}

func TestIrradianceToString(t *testing.T) {
	factory := units.IrradianceFactory{}
	a, err := factory.CreateIrradiance(45, units.IrradianceWattPerSquareMeter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	formatted := a.ToString(units.IrradianceWattPerSquareMeter, 2)
	expected := "45.00 " + units.GetIrradianceAbbreviation(units.IrradianceWattPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
	formatted = a.ToString(units.IrradianceWattPerSquareMeter, -1)
	expected = "45 " + units.GetIrradianceAbbreviation(units.IrradianceWattPerSquareMeter)
	if formatted != expected {
		t.Errorf("expected '%s', got '%s'", expected, formatted)
	}
}

func TestIrradiance_EqualityAndComparison(t *testing.T) {
	factory := units.IrradianceFactory{}
	a1, _ := factory.CreateIrradiance(60, units.IrradianceWattPerSquareMeter)
	a2, _ := factory.CreateIrradiance(60, units.IrradianceWattPerSquareMeter)
	a3, _ := factory.CreateIrradiance(120, units.IrradianceWattPerSquareMeter)

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

func TestIrradiance_Arithmetic(t *testing.T) {
	factory := units.IrradianceFactory{}
	a1, _ := factory.CreateIrradiance(30, units.IrradianceWattPerSquareMeter)
	a2, _ := factory.CreateIrradiance(45, units.IrradianceWattPerSquareMeter)

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